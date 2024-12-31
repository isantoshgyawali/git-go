package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func FileDetails(file string) (*TreeNode, error) {
    info, err := os.Stat(file)
    if err != nil {
        return nil, fmt.Errorf("File not found: %v", err) 
    }

    content, _ := os.ReadFile(file)
    mode := FileModeToGitMode(info.Mode())
    name := fmt.Sprintf("%s", info.Name())
    hash, err := CompressObject("blob", content, nil)
    if err != nil {
        return nil, err 
    }
    
    return &TreeNode{
        Mode: mode, 
        Type: "blob",
        Name: name,
        Hash: hash,
        Path: file,
        IsDir: false,
        Children: nil,
    }, nil 
}

func FindGitRoot() (gitDir string, rootPath string, err error) {
    rootPath, err = os.Getwd() // returns Current Working directory
    if err != nil {
        return "", "", err 
    }

    for {
        gitDir = filepath.Join(rootPath, ".git")
        if _, err := os.Stat(gitDir); err == nil {
            return gitDir, rootPath, nil // found the git root - return the dir
        }
        parentDir := filepath.Dir(rootPath)
        if parentDir == rootPath {
            break
        }
        rootPath = parentDir
    }
    return "", "", fmt.Errorf(".git directory not found.")
}

func ModeToType(mode string) string {
    switch mode {
    case "100644", "100755", "120000":
        return "blob"
    case "040000", "40000":
        return "tree"
    case "160000":
        return "commit"
    default:
        return "unknown"
    }
}

func FileModeToGitMode(mode os.FileMode) string {
    switch {
    case mode.IsDir():
        return "040000"
    case mode.IsRegular():
        return "100644"
    }
    return ""
}
