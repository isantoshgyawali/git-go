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

`git cat-file` ➔ [[ read more ]](./PlumbingForBlob.md)<br>

`git hash-object` ➔ [[ read more ]](./PlumbingForBlob.md)<br>

---

`git ls-tree` ➔ [[ read more ]](./PlumbingForTree.md)<br>

`git write-tree` ➔ [[ read more ]](./PlumbingForTree.md)<br>

