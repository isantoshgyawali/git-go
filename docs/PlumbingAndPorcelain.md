## Plumbing and Porcelain Commands

>Primarily we use Git with 30 or so subcommands such as checkout, branch,
 remote, and so on. But because Git was initially a toolkit for a version control
 system rather than a full user-friendly VCS, it has a number of subcommands that
 do low-level work and were designed to be chained together UNIX-style or called 
 from scripts. These commands are generally referred to as Git’s “plumbing” commands,
 while the more user-friendly commands are called “porcelain” commands.<br>
-- reference [ https://git-scm.com/book/en/v2/Git-Internals-Plumbing-and-Porcelain ]

Basically, Git operations you perform on daily basis are higher-level git operations and are the
parts of the porcelain commands group. eg:  `push` , `checkout` , `branch` , `commit` , `pull` , ....
<br>
Other than those, Commands used underhood certainly on lower-level for the making of git work for us
are the Plumbing commnads, we don't use them on regualar basis but good to know if want to know how 
git internally works

>Context : [ Git uses Zlib as Compression Library  ]

Zlib: It's a compression library that provides in-memory data compression and 
decompression functions, using the [ DEFLATE ](https://en.wikipedia.org/wiki/Deflate) compression algorithm
commonly used for compressing data to save space or to transmit data more
efficiently over a network.

-- Some Plumbing commands that are introduced in this codebase are described here:

`git cat-file`<br>
- allows you to inspect Git objects, view their contents, type, and size

***options:***<br>
-t or --type: Displays the type of the object. <br>
-s or --size: Displays the size of the object in bytes. <br>
-p or --pretty: Pretty prints the contents of the object. <br>

```go
$ git cat-file -p b6fc092b8c43d64c9c2b310d3c2f2f3ff41b9c3a
```
---
`git hash-object`<br>
- reads the inputs from file or standard input (--stdin),
- computes SHA-1 key (hash) of contents provided,
- returns unique key, that represents the content
- optionally, writes the object into your git object database ( -w option ) 

-t \<type> : specifies the type of object (blob, tree, commit, tags)
```go
$ git hash-object -t blob README.md -w
```
