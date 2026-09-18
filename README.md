NAME
====

**xs** - experimental strings carver

SYNOPSIS
========

```console
$ cat FILE | xs | uniq | sort > out.txt
```

DESCRIPTION
===========

xs is an experimental fast strings carver. A string is defined as a sequence of at least 3 ASCII characters. Strings consisting of only whitespaces will be omitted. By reading from any input stream, xs is capable of carving raw forensic disk images and memory dumps.

INSTALLATION
============

```console
$ go install go.foxforensics.eu/xs@latest
```

SEE ALSO
========

[**cat(1)**](https://man7.org/linux/man-pages/man1/cat.1.html),
[**uniq(1)**](https://man7.org/linux/man-pages/man1/uniq.1.html),
[**sort(1)**](https://man7.org/linux/man-pages/man1/sort.1.html)
