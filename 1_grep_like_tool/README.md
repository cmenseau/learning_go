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
Output control options :
- -o : only matching
- -c : count matching lines in file
- -L : get files without any match
- -l : get files with at least 1 match
Prefix options :
- -H : prefix with filename

Keyword format : Keyword is understood as a regexp (BRE basic format)

Regexo known limitations : 
- \\< (word start)
- \\> (word end)
- | ("|" char)

Run tests with
```
cd src/main
go test ./...
``` 