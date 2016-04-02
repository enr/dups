# dups

Search for duplicate files in a directory.

It doesn't delete them, it's up to you check that the files are actually to remove.

Files are considered dups if having the same hash (`sha1`).

**Usage**

Default report format is `sha1 | relative path`:

```
$ ./bin/dups .git/
Looking for duplicates in /tmp/dups/.git
c2c3cf2f0ce489606d88daa5512693a47dbf1cbf logs/HEAD
c2c3cf2f0ce489606d88daa5512693a47dbf1cbf logs/refs/heads/master
3d9d5a25a252676fe509e29afbad086d6edb3707 refs/heads/master
3d9d5a25a252676fe509e29afbad086d6edb3707 refs/remotes/origin/master
Checked 129 files and found 2 dups in 295ns
```

You can customize output using `--names-only` or `--full-path`:

```
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

Using `--quiet` option output is suppressed but exit code is 2 if duplicates are
found or 0 otherwise.

```
$ ./bin/dups --quiet .git/
$ echo $?
2
```
