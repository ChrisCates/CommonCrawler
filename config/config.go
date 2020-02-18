package config

import (
	"github.com/ChrisCates/CommonCrawler/types"
	"github.com/hashicorp/go-multierror"
	"os"
	"path"
)

func Read() (*types.Config, error) {
	cwd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	return &types.Config{
		Start:       0,
		Stop:        1,
		BaseURI:     "https://commoncrawl.s3.amazonaws.com/",
		WetPaths:    path.Join(cwd, "wet.paths"),
		DataFolder:  path.Join(cwd, "/output/crawl-data"),
		MatchFolder: path.Join(cwd, "/output/match-data"),
	}, nil
}

func DirectorySetup(paths []string) error {
	var err error

	for _, p := range paths {
		if pathErr := os.MkdirAll(p, os.ModePerm); pathErr != nil {
			err = multierror.Append(err, pathErr)
		}
	}

	return err
}
