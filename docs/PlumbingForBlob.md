*** *Though, mentioned here : some commands are not just limited to blob objects, like "cat-file"* ***

## git cat-file
>allows you to inspect Git objects, view their contents, type, and size

***options:***<br>
-t or --type: Displays the type of the object. <br>
-s or --size: Displays the size of the object in bytes. <br>
-p or --pretty: Pretty prints the contents of the object. <br>

```golang
$ git cat-file -p b6fc092b8c43d64c9c2b310d3c2f2f3ff41b9c3a
```
---

## git hash-object
>➔ reads the inputs from file or standard input (--stdin),<br>
>➔ computes SHA-1 key (hash) of contents provided,<br>
>➔ returns unique key, that represents the content<br>
>➔ optionally, writes the object into your git object database ( -w option ) 

-t \<type> : specifies the type of object (blob, tree, commit, tags)<br>
-w : writes the object into your git object database

```
$ git hash-object -t blob file.txt -w
3b18e512dba79e4c8300dd08aeb37f8e728b8dad

-- also, writes to .git/objects [ as -w ] --

$ tree .git/objects
.git/objects
└── 3b
    └── 18e512dba79e4c8300dd08aeb37f8e728b8dad

```
