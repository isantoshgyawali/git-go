package args

import (
	"compress/zlib"
	"os"
	"path/filepath"
)

type LsTreeType struct {
    ObjectPermission string
    ObjectType       string
    ObjectId         string
    ObjectName       string
}

func readObject(objectId string) ([]byte, error){
    objectpath := filepath.Join(".git","objects", objectId[:2], objectId[2:])
    file, err := os.Open(objectpath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    zlibReader, err := zlib.NewReader(file)
    if err != nil {
        return nil, err
    }

    return 
}

func LsTree(treeHash string) ([]LsTreeType, error) {
}
