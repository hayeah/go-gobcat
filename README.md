# Dumps Golang Gob File

Utility to pretty print serialized gob.

```
$ gobcat -h

gobcat [file]

        Pretty print Golang gob from file or standard input
```

Dump gob from a file:

```
$ gobcat testdata/test.gob
map[string]interface {}{
    "n2":   int(2),
    "key1": "val1",
    "key2": "val2",
    "n1":   int(1),
}
```

Dump gob from a standard input:

```
$ cat testdata/test.gob > gobcat
map[string]interface {}{
    "n2":   int(2),
    "key1": "val1",
    "key2": "val2",
    "n1":   int(1),
}
```