package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func Test_readWarcRecord(t *testing.T) {

	type testWarcRecord = warcRecord

	var (
		emptyWarcHeader = ""
		emptyWarcBody   = ""
		emptyWarcRecord = emptyWarcHeader + "\r\n" + emptyWarcBody
		emptyWarcReader = bufio.NewReader(strings.NewReader(emptyWarcRecord))
	)

	var (
		simpleWarcHeader = "WARC/1.0\r\nWARC-Type: warcinfo\r\nWARC-Date: 2018-09-26T17:58:38Z\r\nWARC-Filename: CC-MAIN-20180918130631-20180918150631-00000.warc.wet.gz\r\nWARC-Record-ID: <urn:uuid:a327852d-0a00-4a4d-9a92-61d9f9703f5a>\r\nContent-Type: application/warc-fields\r\nContent-Length: 374\r\n"
		simpleWarcBody   = "Software-Info: ia-web-commons.1.1.9-SNAPSHOT-20180911015519\r\nExtracted-Date: Wed, 26 Sep 2018 17:58:38 GMT\r\nrobots: checked via crawler-commons 0.11-SNAPSHOT (https://github.com/crawler-commons/crawler-commons)\r\nisPartOf: CC-MAIN-2018-39\r\noperator: Common Crawl Admin (info@commoncrawl.org)\r\ndescription: Wide crawl of the web for September 2018\r\npublisher: Common Crawl\r\n\r\n"
		simpleWarcRecord = simpleWarcHeader + "\r\n" + simpleWarcBody
		simpleWarcReader = bufio.NewReader(strings.NewReader(simpleWarcRecord))
	)

	//CL for Content-Length
	var (
		wrongCLHeader = "WARC/1.0\r\nContent-Length: 10\r\n"
		wrongCLBody   = "123456789"
		wrongCLRecord = wrongCLHeader + "\r\n" + wrongCLBody
		wrongCLReader = bufio.NewReader(strings.NewReader(wrongCLRecord))
	)

	var (
		wrongVersionHeader  = "WARC/3.0\r\nContent-Length: 10\r\n"
		wrongVersionBody    = "0123456789\r\n\r\n"
		iwrongVersionRecord = wrongVersionHeader + "\r\n" + wrongVersionBody
		wrongVersionReader  = bufio.NewReader(strings.NewReader(iwrongVersionRecord))
	)

	//CL for Content-length
	var (
		noCLHeader = "WARC/1.0\r\n"
		noCLBody   = "whatever\r\n\r\n"
		noCLRecord = noCLHeader + "\r\n" + noCLBody
		noCLReader = bufio.NewReader(strings.NewReader(noCLRecord))
	)

	var (
		emptyContentLengthHeader     = "WARC/1.0\r\nContent-Length:\r\n"
		emptyContentLengthBody       = "body\r\n\r\n"
		emptyContentLengthWarcRecord = emptyContentLengthHeader + "\r\n" + emptyContentLengthBody
		emptyContentLengthReader     = bufio.NewReader(strings.NewReader(emptyContentLengthWarcRecord))
	)

	var (
		nanContentLengthHeader     = "WARC/1.0\r\nContent-Length:azaz\r\n"
		nanContentLengthBody       = "body\r\n\r\n"
		nanContentLengthWarcRecord = nanContentLengthHeader + "\r\n" + nanContentLengthBody
		nanContentLengthReader     = bufio.NewReader(strings.NewReader(nanContentLengthWarcRecord))
	)

	var (
		doubleContentLengthHeader     = "WARC/1.0\r\nContent-Length:4\r\nContent-Length:4\r\n"
		doubleContentLengthBody       = "body\r\n\r\n"
		doubleContentLengthWarcRecord = doubleContentLengthHeader + "\r\n" + doubleContentLengthBody
		doubleContentLengthReader     = bufio.NewReader(strings.NewReader(doubleContentLengthWarcRecord))
	)

	var (
		emptyBodyhHeader = "WARC/1.0\r\nContent-Length:0\r\n"
		emptyBodyRecord  = emptyBodyhHeader + "\r\n"
		emptyBodyReader  = bufio.NewReader(strings.NewReader(emptyBodyRecord))
	)

	var (
		notAWarc       = "WARC/1.0\r\nActually this is not a warc file\r\nbut it starts with warc version\r\n"
		notAWarcReader = bufio.NewReader(strings.NewReader(notAWarc))
	)

	tests := []struct {
		name    string
		args    *bufio.Reader
		want    warcRecord
		wantErr bool
	}{
		{
			name:    "simple correct warc record",
			args:    emptyWarcReader,
			want:    warcRecord{"", ""},
			wantErr: true,
		},
		{
			name:    "simple correnct warc record",
			args:    simpleWarcReader,
			want:    warcRecord{simpleWarcHeader, simpleWarcBody},
			wantErr: false,
		},
		{
			name:    "empty body",
			args:    emptyBodyReader,
			want:    warcRecord{emptyBodyhHeader, ""},
			wantErr: false,
		},
		{
			name:    "content length is more than content",
			args:    wrongCLReader,
			want:    warcRecord{"", ""},
			wantErr: true,
		},
		{
			name:    "invalid WARC version",
			args:    wrongVersionReader,
			want:    warcRecord{"", ""},
			wantErr: true,
		},
		{
			name:    "no Content Length",
			args:    noCLReader,
			want:    warcRecord{"", ""},
			wantErr: true,
		},
		{
			name:    "not a warc revord",
			args:    notAWarcReader,
			want:    warcRecord{"", ""},
			wantErr: true,
		},
		{
			name:    "empty content-length",
			args:    emptyContentLengthReader,
			want:    warcRecord{"", ""},
			wantErr: true,
		},
		{
			name:    "content-lenth is not a number",
			args:    nanContentLengthReader,
			want:    warcRecord{"", ""},
			wantErr: true,
		},
		{
			name:    "double content-length",
			args:    doubleContentLengthReader,
			want:    warcRecord{"", ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readWarcRecord(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("readWarcRecord() error = `%v, wantErr `%v`", err, tt.wantErr)
				return
			}
			if (err != nil) && tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readWarcRecord() = `%v`, want `%v`", got, tt.want)
			}
		})
	}
}
