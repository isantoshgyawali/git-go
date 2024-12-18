`git ls-tree`<br>
---

Lists the content of a given tree object, like what "/bin/ls -a" does in the CWD.<br>
ie. It contains the reference for other blob and tree objects

>***Note that:*** <br>
The behaviour is slightly different from that of "/bin/ls" in that the <path> denotes just a list of patterns to match, e.g. to specifying directory name (without -r) will behave differently, and order of the arguments does not matter.
<br>- reference [ https://git-scm.com/docs/git-ls-tree ]

- In a tree object file, the SHA hashes are not in hex form. They're 20 bytes long raw bytes.
- In a tree object file, entries are sorted by their name.


>Format: <br>
 \<permission> \<type> \<object-hash>\<tab>\<filename>\0

```golang
$ git ls-treee <tree-sha>

-- Example -- 

~/projects/git-go (master)
$ git ls-tree HEAD
100644 blob e69de29bb2d1d6434b8b29ae775ad8c2e48c5391    .gitignore
100644 blob 5c4221225b3a73b3cd1a5d67978a6fc75ecf5aee    README.md
040000 tree 5dbaefb47d41ca2f884075a8e1446c21f8b58381    args
040000 tree bd271c453fc64f880775f33d19edb47213b8155a    cmd
040000 tree 81fd4f43380c7f3b357cbade1366fc91ef895c74    docs
100644 blob ea7212c1feb7dd6ab7584611b08c4adc9eeb4714    go.mod
100644 blob f9d6d443ec1e90e0b2b6d9503d1ddff1b478cb10    go.sum
100755 blob ed02809a7da815888decce805e095224931c24ee    main
```

>flags: <br>
-d : Show only the named tree entry itself, not its children<br>
-r : list all files and subdirectories recursively<br>
-t : list tree-objects too. (Need to use with -r) <br>
--name-only: List only filenames (instead of the "long" output), one per line<br>
--object-only: List only names of the objects

```golang
~/projects/git-go (master)
$ git ls-tree HEAD -d
040000 tree 5dbaefb47d41ca2f884075a8e1446c21f8b58381    args
040000 tree bd271c453fc64f880775f33d19edb47213b8155a    cmd
040000 tree 81fd4f43380c7f3b357cbade1366fc91ef895c74    docs

~/projects/git-go (master)
$ git ls-tree HEAD -r
100644 blob e69de29bb2d1d6434b8b29ae775ad8c2e48c5391    .gitignore
100644 blob 5c4221225b3a73b3cd1a5d67978a6fc75ecf5aee    README.md
100644 blob c771e6a6ac87d48ceb926a1610dfc587e5320514    args/catFile.go
100644 blob 9306498a50a305f860e0d11e99c0ed0ff005b418    args/hashObject.go
100644 blob 2606d10b4cf54cf9bb9223a47b864b5364916318    cmd/main/main.go
100644 blob 39b71263eb06db0b7089bd5a43e07974cf06ca4b    docs/GitObjects.md
100644 blob a581ad56d4d039b594affd33a8b27f3e8ec93928    docs/PlumbingAndPorcelain.md
100644 blob ea7212c1feb7dd6ab7584611b08c4adc9eeb4714    go.mod
100644 blob f9d6d443ec1e90e0b2b6d9503d1ddff1b478cb10    go.sum
100755 blob ed02809a7da815888decce805e095224931c24ee    main

~/projects/git-go (master)
$ git ls-tree HEAD -r -t
100644 blob e69de29bb2d1d6434b8b29ae775ad8c2e48c5391    .gitignore
100644 blob 5c4221225b3a73b3cd1a5d67978a6fc75ecf5aee    README.md
040000 tree 5dbaefb47d41ca2f884075a8e1446c21f8b58381    args
100644 blob c771e6a6ac87d48ceb926a1610dfc587e5320514    args/catFile.go
100644 blob 9306498a50a305f860e0d11e99c0ed0ff005b418    args/hashObject.go
040000 tree bd271c453fc64f880775f33d19edb47213b8155a    cmd
040000 tree 77fd5229988f2b49699c78ca282115f6eb0a835a    cmd/main
100644 blob 2606d10b4cf54cf9bb9223a47b864b5364916318    cmd/main/main.go
040000 tree 81fd4f43380c7f3b357cbade1366fc91ef895c74    docs
100644 blob 39b71263eb06db0b7089bd5a43e07974cf06ca4b    docs/GitObjects.md
100644 blob a581ad56d4d039b594affd33a8b27f3e8ec93928    docs/PlumbingAndPorcelain.md
100644 blob ea7212c1feb7dd6ab7584611b08c4adc9eeb4714    go.mod
100644 blob f9d6d443ec1e90e0b2b6d9503d1ddff1b478cb10    go.sum
100755 blob ed02809a7da815888decce805e095224931c24ee    main

~/projects/git-go (master)
$ git ls-tree HEAD --name-only
.gitignore
README.md
args
cmd
docs
go.mod
go.sum
main

~/projects/git-go (master)
$ git ls-tree HEAD --object-only
e69de29bb2d1d6434b8b29ae775ad8c2e48c5391
5c4221225b3a73b3cd1a5d67978a6fc75ecf5aee
5dbaefb47d41ca2f884075a8e1446c21f8b58381
bd271c453fc64f880775f33d19edb47213b8155a
81fd4f43380c7f3b357cbade1366fc91ef895c74
ea7212c1feb7dd6ab7584611b08c4adc9eeb4714
f9d6d443ec1e90e0b2b6d9503d1ddff1b478cb10
ed02809a7da815888decce805e095224931c24ee
```

`git write-tree`
---

The "git write-tree" command creates a tree object from the current index of "staging area". 
The staging area is a place where changes go when you run git add.
And, the name of the new tree object is printed to the standard output

```bash 
$ git write-tree

-- Example -- 
# create some files and subdir
~/projects/git-go (master)
$ echo "Content of file1" > file1.txt && echo "Content of file2" > file2.txt

~/projects/git-go (master)
$ mkdir subdir && echo "Content of file3 in subdir" > subdir/file3.txt

# Add the files to staging area
~/projects/git-go (master)
$ git add file1.txt file2.txt subdir/file3.txt

# write-tree
~/projects/git-go (master)
$ git write-tree
3a0c8a1cbbd0b4f4344e947db9e6b11c9075f445

```

>Note: <br> 
This git-go implementation doesn't include the concept of the staging 
area at least for now - here, we assume every file is in the staging area  

```bash
# To inspect and test the subtree we generated we could try `git cat-file` or `git ls-tree`
$ git cat-file -p 3a0c8a1cbbd0b4f4344e947db9e6b11c9075f445
100644 blob a7815fd88b2e91bfa8a9b6619e490cd0cf147cc0    file1.txt
100644 blob b47b13c86524a5501118256fa4387c9e52d7bdc9    file2.txt
040000 tree e1e43b13c1d0eb8e4f69b302be6e1dc62492f97d    subdir
-- or -- 
$ git ls-tree 3a0c8a1cbbd0b4f4344e947db9e6b11c9075f445
100644 blob a7815fd88b2e91bfa8a9b6619e490cd0cf147cc0    file1.txt
100644 blob b47b13c86524a5501118256fa4387c9e52d7bdc9    file2.txt
040000 tree e1e43b13c1d0eb8e4f69b302be6e1dc62492f97d    subdir

```
