package args

import (
	"bytes"
	"fmt"

	"github.com/isantoshgyawali/git-go/utils"
)

func LsTree(treeHash string) ([]*utils.TreeEntry, error) {
    // take hash
    // forward hash to ParseTree
    // parse tree returns slices of TreeEntries 
    // return treeentries
    path := utils.GetObjectPath(treeHash)
    content, err := utils.DecompressObject(path)
    if err != nil {
        return nil, err 
    }

    nullByteIndex := bytes.IndexByte(content, 0)
    if nullByteIndex == -1 {
        return nil, fmt.Errorf("Invalid Git Object Format")
    }
    objectContent := content[nullByteIndex+1:]

    treeEntries, err := utils.ParseTree([]byte(objectContent))
    if err != nil {
        return nil, err 
    }

    return treeEntries, nil
}



