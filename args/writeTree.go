package args

import (
    "fmt"
    "os"
    "path/filepath"

    "github.com/isantoshgyawali/git-go/utils"
)

func WriteTree(currentPath string) (string, error) {
    fs, err := os.ReadDir(currentPath)
    if err != nil {
        return "", err
    }

    root := &utils.TreeNode{
        Mode:  "040000",
        Type:  "tree",
        Name:  filepath.Base(currentPath),
        Path:  currentPath,
        IsDir: true,
    }

    // FileSystem-Tree traversal
    serializedContent := ""
    for _, entry := range fs {
        // Skipping .git
        if entry.Name() == ".git" {
            continue
        }

        // New current-path for the entry
        newPath := root.Path + "/" + entry.Name()

        if entry.IsDir() {
            childHash, err := WriteTree(newPath)
            if err != nil {
                return "", err
            }

            // Creating a new TreeNode for entry with valid content (non-empty hashes)
            // *** there could be better implementation ***
            if childHash != "" {
                childNode := &utils.TreeNode{
                    Mode:  "040000",
                    Type:  "tree",
                    Name:  entry.Name(),
                    Path:  newPath,
                    IsDir: entry.IsDir(),
                    Hash:  childHash,
                }
                root.Children = append(root.Children, childNode)
                serializedContent += fmt.Sprintf("%s %s\000%s", childNode.Mode, childNode.Name, childHash)
                // fmt.Printf("%d - %s/%s - %v\n", i, newPath, entry.Name(), childHash)
            }
        } else {
            fileHash, err := utils.CompressObject("blob", []byte(newPath), nil)
            if err != nil {
                return "", err
            }

            childNode := &utils.TreeNode{
                Mode:  "100644",
                Type:  "blob",
                Name:  entry.Name(),
                Path:  newPath,
                IsDir: false,
                Hash:  fileHash,
            }
            root.Children = append(root.Children, childNode)
            serializedContent += fmt.Sprintf("%s %s\000%s", childNode.Mode, childNode.Name, fileHash)
            // fmt.Printf("%d - %s/%s - %v\n", i, currentPath, entry.Name(), fileHash)
        }
    }

    // if directory is empty, return an empty hash
     if len(root.Children) == 0 {
        return "", nil
    }

    // Serialize the tree and compute its hash
    treeHash, _ := utils.CompressObject("tree", []byte(serializedContent), nil)
    fmt.Println(treeHash)

    return "", nil
}

