package args

import (
	"bytes"
	"fmt"

	"github.com/isantoshgyawali/git-go/utils"
)

func CatFile(objectHash string) (*utils.Object, error) {
        path := utils.GetObjectPath(objectHash)
        content, err := utils.DecompressObject(path)
        if err != nil {
            return nil, err 
        }

	// format for object header returned by zlib decompression:
	// <type> <size>\0<content>
	nullByteIndex := bytes.IndexByte(content, 0)
	if nullByteIndex == -1 {
		return nil, fmt.Errorf("Invalid Git Object Format")
	}

        objectType := string(content[:bytes.IndexByte(content, ' ')])
        objectContentStr := string(content[nullByteIndex+1:])
        objectSize := int64(len(objectContentStr))

        obj := &utils.Object{
            Type: objectType,
            Size: objectSize,
            Content: nil,
        }

        switch objectType {
            case "blob": 
                obj.Content = objectContentStr
            case "tree":
                treeEntries, err := utils.ParseTree([]byte(objectContentStr))
                if err != nil {
                    return nil, err 
                }
                obj.TreeEntries = treeEntries
            default:
                // here returning as it is for other type
                // but could be updated for different cases
                obj.Content = string(content[nullByteIndex+1:])
        }

	return obj, nil
}
