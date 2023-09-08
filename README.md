# SecLang ANTLR based parser

This repo contains the ANTLR files for a SecLang parser.

It is still a Proof of Concept, not usable, and incomplete version.

## Why a new parser?

There has been efforts towards having parsers in different languages. Using ANTLR would allow us to have a common parser and generate parsing engines for different languages easily. This way we would consolidate efforts and we can have a more robust parser.

## Usage

Right now this repo contains the ANTLR files and golang basic tests.

To run the tests:

```bash
go generate ./...
go test ./...
```

## Author

Felipe Zipitria <felipe.zipitria@owasp.org>
