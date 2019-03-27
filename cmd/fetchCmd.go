// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	aurora "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"os"
	"path"
)

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	panic(err)
}

var extractCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch data files from Common Crawler",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		existed, _ := pathExists(flags.dataFolder)
		if !existed {
			fmt.Println(aurora.Blue("Creating folder: " + flags.dataFolder))
			os.Mkdir(flags.dataFolder, 0740)
		}
		existed, _ = pathExists(flags.matchFolder)
		if !existed {
			fmt.Println(aurora.Blue("Creating folder: " + flags.matchFolder))
			os.Mkdir(flags.matchFolder, 0740)
		}
		scan(flags)
	},
}
var flags Config

func init() {
	rootCmd.AddCommand(extractCmd)

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	extractCmd.PersistentFlags().IntVarP(&flags.start, "start", "s", 0, "First number of file defined in wet.paths to fetch")
	extractCmd.PersistentFlags().IntVarP(&flags.stop, "stop", "e", 1, "Last number minus 1 of file defined in wet.paths to fetch")
	extractCmd.PersistentFlags().StringVarP(&flags.baseURI, "base_uri", "u", "https://commoncrawl.s3.amazonaws.com/", "Base Uri of Common Crawler")
	extractCmd.PersistentFlags().StringVarP(&flags.wetPaths, "wet_paths", "w", path.Join(cwd, "wet.paths"), "wet.paths configuration file path")
	extractCmd.PersistentFlags().StringVarP(&flags.dataFolder, "output_dir", "o", path.Join(cwd, "/output/crawl-data"), "Output directory to store fetched files")
	extractCmd.PersistentFlags().StringVarP(&flags.matchFolder, "match_dir", "m", path.Join(cwd, "/output/match-data"), "Matched directory to store analyzed files information")
}
