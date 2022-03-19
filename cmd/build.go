/*
Copyright © 2021 Ved Thiru (PerpetualCreativity) <vedthiru@hotmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


var outputFolder, inputFolder string
var programs []map[string][]string

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build [folder]",
	Short: "Generates site. Example: span build my_site",
	Long: `span build generates the site from the source, and outputs.
Example: span build my_site runs the specified programs over the folder my_site.
If folder is not specified, the default is the current folder.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// use current directory as default input.
		inputFolder = "."
		if len(args) == 1 {
			inputFolder = args[0]
		}
		// check if input folder exists
		_, err := os.Stat(inputFolder)
		fc.ErrCheck(err, "folder doesn't exist or is not accessible")

		configFileUsed, err := filepath.Rel(".", viper.ConfigFileUsed())
		fc.ErrCheck(err, "config file not found")

		outputFolder = viper.GetString("output")
		inputFolder = filepath.Clean(inputFolder)

		rawPrograms := viper.Get("programs")

		rawP, ok := rawPrograms.([]interface{})
		fc.ErrNComp(ok, false, "programs must be a YAML list.")

		for _, smss := range rawP {
			programs = append(programs, cast.ToStringMapStringSlice(smss))
		}

		os.RemoveAll(outputFolder)
		os.Mkdir(outputFolder, 0755)

		// array of `done` channels
		var finish []chan bool

		err = filepath.WalkDir(inputFolder, func(path string, info fs.DirEntry, err error) error {
			if path == inputFolder || strings.Contains(path, outputFolder) {
				return nil
			}
			// ignore config file
			if path == configFileUsed {
				return nil
			}

			done := make(chan bool)
			finish = append(finish, done)

			fileRun(done, path, info)
			return nil
		})
		fc.ErrCheck(err, "Error while walking through folder.")

		// only return after all goroutines we started (in fileRun) finish
		for _, done := range finish {
			<- done
			close(done)
		}
		fc.Success("All programs completed.")
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

func fileRun(done chan<- bool, path string, fileInfo fs.DirEntry) {
	go func() {
		dir, name := filepath.Split(path)
		ext := filepath.Ext(path)
		index := -1
		if !fileInfo.IsDir() {
			out:
			for i, program := range programs {
				for _, eon := range program["files"] {
					if name == eon || ext == eon {
						index = i
						break out
					}
				}
			}
			if index == -1 {
				done <- true
				return
			}
		}

		dir = strings.Replace(dir, inputFolder, "", 1)
		dir = filepath.Join(outputFolder, dir)

		if ext != "" {
			name = strings.Replace(name, ext, programs[index]["outputExt"][0], 1)
		}
		outputPath := filepath.Join(dir, name)

		if fileInfo.IsDir() {
			// 0755 for directories
			os.Mkdir(outputPath, 0755)
		} else {
			contents, err := os.ReadFile(path)
			fc.ErrCheck(err, "failed to read folder")
			for _, program := range programs[index]["commands"] {
				programArgs := strings.Split(program, " ")
				p := exec.Command(programArgs[0], append(programArgs[1:], path)...)
				pIn, _ := p.StdinPipe()
				pOut, _ := p.StdoutPipe()
				p.Start()
				pIn.Close()
				pByteOut, err := ioutil.ReadAll(pOut)
				fc.ErrCheck(err, "failed to run program"+program)
				p.Wait()
				contents = pByteOut
			}
			// 0644 for files
			err = os.WriteFile(outputPath, contents, 0644)
			fc.ErrCheck(err, "failed to write to file")
		}
		done <- true
	}()
}

