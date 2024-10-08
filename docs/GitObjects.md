## Git Objects

```go
Git is a content-addressable filesystem.
i.e. at the core git is "Key-Value" data store.
```
Git stores everything—commits, files, trees, and tags—as objects.and 
each object is represented by a unique SHA-1 hash (now SHA-256 in more recent versions) 
and stored within the .git/objects directory.
 
> *"Unlike regular filesystems, where the name of a file is arbitrary and unrelated to that
file’s contents, the names of files as stored by Git are mathematically derived from 
their contents. This has a very important implication: if a single byte of, say, a text 
file, changes, its internal name will change, too. To put it simply: you don’t modify a
file in git, you create a new file in a different location."<br>
-- refrence : [ https://wyag.thb.lt/#objects ]*

For context:<br>
If the SHA-1 hash of object is abcdef1234567890abcdef1234567890abcdef12, the file path will be:
```bash
.git/objects/ab/cdef1234567890abcdef1234567890abcdef12
```

### Key-Objects
---
#### `Blob Object` : [ Binary Large Object ]
- Blobs contains raw contents of the file without any metadata like the filename or file Permissions
- Blobs are identified by their SHA-1 hash
- If multiple files have the same content, Git will store only one blob for that content

-- ***Each Blob is stored as a seperate file in .git/objects dir***<br>
-- ***File Contains a header and the contents of the blob objects*** <br>
-- ***Compressed using Zlib***

The format of a blob object file looks like this (after Zlib decompression):
```go
blob <size>\0<content>
-- for example --
blob 11\0hello world
```

#### `Tree Objects`
- Tree objects represent directories, and they reference other trees and blobs
- Like other Git objects, a tree object is immutable. Once created, its contents do not change.
- Each entry has:
    - File mode: Indicates the type and permissions (e.g., 100644 for a file, 040000 for a subdirectory).
    - SHA-1 hash (20 bytes): References a blob (file) or another tree (subdirectory).
    - File name or directory name.

>The root tree represents the entire repository at a specific point, and can have nested trees for subdirectories.
A commit object references a root tree object that represents the state of the repository at the time of the commit.

```golang
100644 filename\0<20-byte hash>
```

#### `Commit Objects`
#### `Tag Objects`

