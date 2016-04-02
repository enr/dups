# dups

Search for duplicate files in a directory.

It doesn't delete them, it's up to you check that the files are actually to remove.

Files are considered dups if having the same hash (`sha1`).

Default report format is `sha1 | relative path`:

```
da3...80709   1lebsnfoptv8qpa10w6kyy5mp/gradle-2.4-bin.zip.lck
da3...80709   ProjectScript/buildscript/classes/emptyScript.txt
```

You can customize output using:

--names-only:

1lebsnfoptv8qpa10w6kyy5mp/gradle-2.4-bin.zip.lck
ProjectScript/buildscript/classes/emptyScript.txt

--full-path:

/basedir/1lebsnfoptv8qpa10w6kyy5mp/gradle-2.4-bin.zip.lck
/basedir/ProjectScript/buildscript/classes/emptyScript.txt
