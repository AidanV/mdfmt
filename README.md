## Markdown Formatter

- Adds whitespace to the end of files
- Adds empty lines around headers
- Adds spacing for line breaks
- Removes whitespaces from links
- Unindents paragraphs
- Moves first line of text to the beginning of a file

## Usage

mdfmt [filename...]

This can take multiple files as input and format them all.

## Building

If on nixos run:

```Bash
nix develop
```

Then run

```Bash
go run mdfmt
```

To run unit tests run
```Bash
cd mdfmt
go test -v
```
