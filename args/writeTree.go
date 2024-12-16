package args

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/isantoshgyawali/git-go/utils"
)

func WriteTree() (string, error) {
    gitRoot, err := utils.FindGitRoot()
    if err != nil {
        return "", err 
    }
    rootPath := strings.Trim(gitRoot, ".git")
    var entries []string

    if err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
        if d.IsDir() && d.Name() == ".git"{
            return fs.SkipDir
        }
       fmt.Println(path) 
       return nil
    }); 
    err != nil {

    }

    // <mode> <name>\0<hash (20 bytes)> <mode> <name>\0<hash (20 bytes)> ... ... .. 
    var serializedContent string
    for _, entry := range entries{
        info, err := os.Stat(entry)
        if err != nil {
            return "", err 
        }

        // Recursive handling for the sub-dir
        if info.IsDir() {
            subTreeHash, err := WriteTree()
            if err != nil {
                return "", err 
            }
            serializedContent += fmt.Sprintf("%s %s\000%s", "040000", entry, subTreeHash)
        } else {
            entryDetails, err := utils.FileDetails(entry)
            if err != nil {
                return "", err 
            }
            serializedContent += fmt.Sprintf("%s %s\000%s", entryDetails.Mode, entryDetails.Name, entryDetails.Hash)
        }
    }

    // buf := &bytes.Buffer{}
    // treeHash, _ := utils.CompressObject("tree", []byte(serializedContent), buf)
    // fmt.Println(treeHash)
    
    return "", nil
}
