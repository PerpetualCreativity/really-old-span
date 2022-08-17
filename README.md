# **NB: this version of `span` is now deprecated. Please use this version instead: [PerpetualCreativity/span](https://github.com/PerpetualCreativity/span)**

# `span`

`span` is an easy-to-use static site generator that lets you use the command-line tools you want to use. Each file is processed concurrently.

## Usage

Run `span init` to get a sample configuration file. It will look like the following:

```yaml
# span configuration
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
output: output
```

`span` doesn't care about what commands you want to run, as long as they accept a file path as input. [`pandoc`](https://pandoc.org) and [`sass`](https://sass-lang.com) are just examples.

`span` is currently beta software - configuration or behaviour may change slightly.

Run `span` or `span help` to get help. Run `span init` to generate the sample configuration file above.

Run `span build [folder]` to generate your site according to your configuration. If the path to the folder containing the original files is not found, span will attempt to use the current directory (`.`).

## Feedback / Contributing

If you've found a bug or would like to propose a feature request, first look at [GitHub Issues](https://github.com/PerpetualCreativity/span/issues) to see if someone else has already done so. If so, and if you have additional information, please add that information as a comment. If not, please [open an issue in GitHub](https://github.com/PerpetualCreativity/span/issues/new).

Pull requests are welcome, but please discuss your proposed changes in an issue first.

Thanks!

Ved Thiru (@PerpetualCreativity on GitHub, vulcan#0604 on Discord)
