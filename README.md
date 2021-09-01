# span

Span is an easy-to-use static site generator that lets you use the command-line tools you want to use. Each file is processed concurrently.

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

Span doesn't care about what commands you want to run, as long as they accept a file path as input. Pandoc and Sass are just examples.

