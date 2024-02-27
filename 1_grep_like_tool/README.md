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

Run tests with
```
cd src/main
go test .
``` 