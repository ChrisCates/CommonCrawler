package crawler

import (
	"github.com/hashicorp/go-multierror"
	"io"
	"os"
	"path"
)

func NewCrawler(cfg *Config) CommonCrawler {
	return &crawler{cfg}
}

func ReadConfig() (*Config, error) {
	cwd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	return &Config{
		Start:       0,
		Stop:        1,
		BaseURI:     "https://commoncrawl.s3.amazonaws.com/",
		WetPaths:    path.Join(cwd, "wet.paths"),
		DataFolder:  path.Join(cwd, "output/crawl-data"),
		MatchFolder: path.Join(cwd, "output/match-data"),
	}, nil
}

func DirectorySetup(dirPaths ...string) error {
	var err error

	for _, p := range dirPaths {
		if e := os.MkdirAll(p, 0740); e != nil {
			err = multierror.Append(err, e)
		}
	}

	return err
}

type (
	CommonCrawler interface {
		Scan() error
		Extract(path string) error
		Download(uri, path string) error
		Analyze(searchFor string, in io.Reader, matched chan MatchedWarc)
		AnalyzeFile(filepath, path string) error
	}

	MatchedWarc struct {
		Matches  int
		WarcData string
	}

	WarcRecord struct {
		Header string
		Body   string
	}

	Config struct {
		BaseURI     string
		WetPaths    string
		DataFolder  string
		MatchFolder string
		Start       int
		Stop        int
	}
	crawler struct {
		config *Config
	}
)
