# last-resort
Are you sick of dubious people who beg for __your__ code? Give it to them... unreadable.

This small tool just takes your files, replaces all the words in them with a given lexeme and saves under a new name.

# Installation
```
go get github.com/andiogenes/last-resort
```

# Usage
```shell script
$ last-resort -l="cut my life" into.pieces
$ last-resort -l="this is my" -o=last.txt resort.txt
$ last-resort -l="suffocation" no-breathing # folder implied
$ last-resort -l="parental advisory" -o=explicit content
```
