package utils

type Object struct {
    Type string
    Content interface{}
    Size int64
    TreeEntries []*TreeEntry
}

type TreeEntry struct {
    Mode string
    Type string
    Name string
    Hash string
}
