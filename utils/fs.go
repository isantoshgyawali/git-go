package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

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
    return "", "", fmt.Errorf("Error finding git root. Have you initialized the git project?")
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
      if mode&0111 != 0 { return "100755" } // executable file 
      return "100644" // non-executable regular file
    case mode&os.ModeSymlink != 0:
      return "120000" // symbolic links
    // there are other fileTypes too : add more as required : [socket, device file... ]
  }
    return ""
}
