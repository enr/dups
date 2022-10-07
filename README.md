# dups

![CI Linux Mac](https://github.com/enr/dups/workflows/CI%20Linux%20Mac/badge.svg)
![CI Windows](https://github.com/enr/dups/workflows/CI%20Windows/badge.svg) 
https://enr.github.io/dups/

Search for duplicate files in a directory.

It doesn't delete them, it's up to you check that the files are actually to remove.

Files are considered dups if having the same hash (`sha1`).

## Usage

Default report format is `sha1 | relative path`:

```console
$ ./bin/dups .git/
Looking for duplicates in /tmp/dups/.git
c2c3cf2f0ce489606d88daa5512693a47dbf1cbf logs/HEAD
c2c3cf2f0ce489606d88daa5512693a47dbf1cbf logs/refs/heads/master
3d9d5a25a252676fe509e29afbad086d6edb3707 refs/heads/master
3d9d5a25a252676fe509e29afbad086d6edb3707 refs/remotes/origin/master
Checked 129 files and found 2 dups in 295ns
```

You can customize output using `--names-only` or `--full-path`:

```console
$ ./bin/dups --names-only .git/
logs/HEAD
logs/refs/heads/master
refs/heads/master
refs/remotes/origin/master

$ ./bin/dups --full-path .git/
Looking for duplicates in /tmp/dups/.git
c2c3cf2f0ce489606d88daa5512693a47dbf1cbf /tmp/dups/logs/HEAD
c2c3cf2f0ce489606d88daa5512693a47dbf1cbf /tmp/dups/logs/refs/heads/master
3d9d5a25a252676fe509e29afbad086d6edb3707 /tmp/dups/refs/heads/master
3d9d5a25a252676fe509e29afbad086d6edb3707 /tmp/dups/refs/remotes/origin/master
Checked 129 files and found 2 dups in 230ns
```

Using `--quiet` option output is suppressed but exit code is the number of duplicates found.

```console
$ ./bin/dups --quiet .git/
$ echo $?
2
```

You can force Dups to set the number of duplicates as exit code using the option `--dups-exit`.

You can exclude certain files or directories using `--exclude` and only include certain **filenames** using `--include`.
Both flags supports patterns, e.g `--include '*.txt'`

```console
$ ./bin/dups . --exclude .git --include '*.txt'
Looking for duplicates in .
f1d2d2f924e986ac86fdf7b36c94bcdf32beec15 testdata/01/01.txt
f1d2d2f924e986ac86fdf7b36c94bcdf32beec15 testdata/01/sub/010.txt
Checked 4 files and found 2 dups in no time
```

To disable output colors use `--no-color` or the `NO_COLOR` environment variable.

To see more: https://enr.github.io/dups/


## Develop

Download or clone repository.

Requires Go 1.19 (or higher).

Build (binaries will be created in `bin/`):

```sh
./.sdlc/build
```

Check (code quality and tests);

```sh
./.sdlc/check
```

## License

Copyright (C) 2020 Dups authors.
