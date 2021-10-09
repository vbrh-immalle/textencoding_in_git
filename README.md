# Text-encodings

Demo-project to *research* Unicode-encoding and line endings in text-files.

- Windows Line Endings (`CR LF`)
- Linux Line Endings (`LF`)

- UTF8
- UTF16BE (Big Endian)
- UTF16LE (Little Endian)

# Explanation

Every `.txt`-file has the same content: 
the sentence `This is an example!` followed by a new line.

Git interferes with (i.e. *auto-converts*) line-endings and encodings
To avoid this, we use `.gitattributes`.

The small and primitive tool `showendings.go` shows the last 4 bytes of each text-file. Run it (after installing the go compiler) with

    go run showendings.go

To test the the influence of `.gitattributes`, this repo has 2 tags: `with` git attributes and `without` using git attributes.

Try this:

    git rm *.txt        # delete files so git will re-create them ...
    git checkout with   # ... when we checkout the version with gitattributes
    go run showendings.go
    git rm *.txt
    git checkout without # ... or the version without gitattributes
    go run showendings.go

Or perhaps use a diff-tool (for example on Windows with Powershell):

```powershell
    git rm *.txt
    git checkout with
    go run showendings.go | Out-File with
    git rm *.txt
    go checkout without
    go run showendings.go | Out-File without
    Compare-Object (cat with) (cat without) -IncludeEqual
```

I hope this can help you understand Unicode and line-endings!