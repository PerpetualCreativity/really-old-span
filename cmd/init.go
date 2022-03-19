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
	"io/ioutil"

	"github.com/spf13/cobra"
)

var path *string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates a config file.",
	Long: `'span init' creates a configuration file for your site.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := ioutil.WriteFile(*path, ([]byte(`# span configuration
# programs you want to run on your files
programs:
  - files:
      - .md
    commands:
      - pandoc --template=layout.html
    outputExt: .html
  - files:
      - .scss
      - .sass
    commands:
      - sass
    outputExt: .css
# name of folder to output
output: output`)), 0644)
		fc.ErrCheck(err, "Failed to write to file.")
		fc.Neutral("Successfully created sample configuration file.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	path = initCmd.Flags().String("path", "./.span.yaml", "path to where the config file should be placed")
}

