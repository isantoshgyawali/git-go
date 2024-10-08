package args

import (
	"bytes"
	"fmt"
)

type LsTreeType struct {
	ObjectId         string
	ObjectName       string
	ObjectType       string
	ObjectPermission int
}

func LsTree(treeHash string) ([]LsTreeType, error) {

    objectType, content, _, err := CatFile(treeHash)
    if err != nil {
        return nil, err
    }

    // checking if the object is tree or not
    if objectType != "tree" {
        return nil, fmt.Errorf("object %s is not a tree", treeHash)
    }


    var lsTreeEntries []LsTreeType
    contentBytes := []byte(content)

    for len(contentBytes) > 0 {
        var entry LsTreeType
        spaceIndex := bytes.IndexByte(contentBytes, ' ')
        if spaceIndex == -1 {
            break
        }

        mode := string(contentBytes[:spaceIndex])
        fmt.Printf(mode)
        contentBytes := contentBytes[spaceIndex+1:]

        nullIndex := bytes.IndexByte(contentBytes, 0)
        if nullIndex == -1 {
            break
        }

        lsTreeEntries = append(lsTreeEntries, entry)
    }


    return lsTreeEntries, nil
}
