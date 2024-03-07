Making my own grep-like tool : find words and patterns from a text source

Run with 
```
cd src/main
go run . [options...] keyword file...
```

Supported options :
- -i : search case insensitive
- -w : match word
- -x : match whole line
- -v : reverse (get lines not matching)

Keyword format : Keyword is understood as a regexp (BRE basic format)

Known limitations : 
- \\<
- \\>
- |

Run tests with
```
cd src/main
go test ./...
``` 