package utils

type Object struct {
    Type string
    Content interface{}
    Size int64
}

type TreeNode struct {
    Mode string
    Type string
    Name string
    Hash string
    Path string
    IsDir bool
    Children []*TreeNode
}
