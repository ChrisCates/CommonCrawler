package crawler

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/logrusorgru/aurora"
)

func analyzeFile(filePath string, path string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := os.Create(path)
	if err != nil {
		return err
	}
	defer data.Close()

	matched := make(chan matchedWarc)

	go analyze("ninja", file, matched)

	for m := range matched {
		fmt.Println(aurora.Blue("\t  Found " + strconv.Itoa(m.matches) + " matches in this wet file..."))
		data.WriteString(fmt.Sprintf("%sMATCHES: %d\n\n", m.warcData, m.matches))
	}

	return nil
}

type warcRecord struct {
	header string
	body   string
}

//readWarcRecord reads one warc record from Reader
//  warc-record  = header CRLF
//  block CRLF CRLF
func readWarcRecord(in *bufio.Reader) (warcRecord, error) {

	var ret warcRecord

	line, err := in.ReadBytes('\n')
	if err != nil {
		return ret, err
	}

	firstLine := string(line)

	//Warc record starts with version e.g. "WARC/1.0"
	if firstLine != "WARC/1.0\r\n" {
		return ret, fmt.Errorf("warc version expected '%s' found", firstLine)
	}
	var warcHeaderBuilder strings.Builder

	var contentLength = -1

	//read header till end (\n)
	for ; string(line) != "\r\n"; line, err = in.ReadBytes('\n') {

		if err != nil {
			return ret, err
		}

		//each header must contains Content-Length
		//alse named headers are case insensitive
		if strings.HasPrefix(strings.ToLower(string(line)), "content-length:") {

			if contentLength > 0 {
				return ret, fmt.Errorf("exactly one content-length should be present in a WARC header")
			}

			keyAndValue := strings.SplitN(string(line), ":", 2)
			if len(keyAndValue) != 2 {
				return ret, fmt.Errorf("Content-Length field must contains a value. '%s' found)", line)
			}
			//field value may be preceded by any  amount  of  linear  whitespace
			strValue := strings.TrimSpace(keyAndValue[1])
			contentLength, err = strconv.Atoi(strValue)
			if err != nil {
				return ret, err
			}
		}

		warcHeaderBuilder.Write(line)
	}

	//content length sould be non-negative
	if contentLength < 0 {
		return ret, fmt.Errorf("exactly one content-length should be present in a WARC header. WARC header: %s", warcHeaderBuilder.String())
	}

	//early return if body is empty
	if contentLength == 0 {
		return warcRecord{warcHeaderBuilder.String(), ""}, nil
	}

	//body buffer
	body := make([]byte, contentLength)

	n := 0
	//put reader date to body buffer
	for k, err := in.Read(body); n < contentLength; k, err = in.Read(body[n:]) {
		if err != nil && err != io.EOF {
			return ret, err
		}
		if err == io.EOF && (n+k) < contentLength {
			return ret, fmt.Errorf("WARC record finished unexpectedly. Content-Length : %d, got %d", contentLength, n)
		}
		n += k
	}

	return warcRecord{warcHeaderBuilder.String(), string(body)}, err
}

type matchedWarc struct {
	matches  int
	warcData string
}

func analyze(searchFor string, in io.Reader, matched chan matchedWarc) {
	defer close(matched)
	bufin := bufio.NewReader(in)
	var wg sync.WaitGroup

	for record, err := readWarcRecord(bufin); err == nil; record, err = readWarcRecord(bufin) {
		wg.Add(1)
		go func(r warcRecord) {
			found := strings.Count(r.body, searchFor)

			if found > 0 {
				matched <- matchedWarc{found, r.header}
			}
			wg.Done()
		}(record)
		bufin.ReadBytes('\n')
		bufin.ReadBytes('\n')
	}

	wg.Wait()
}
