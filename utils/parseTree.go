package utils

import (
	"bytes"
	"fmt"
)

// Function parses the content and retrun entries if content is a tree object
// format of the parsing content: 
// <mode> <name>\0<hash (20 bytes)> <mode> <name>\0<hash (20 bytes)> ... ... .. 
func ParseTree(content []byte) ([]*TreeEntry, error) {
    var entries []*TreeEntry      
    for len(content) > 0 {
        spaceIndex := bytes.IndexByte(content, ' ')
        if spaceIndex == -1 {
            return nil, fmt.Errorf("Invalid tree object format: missing space")
        }
        entryMode := string(content[:spaceIndex])
        entryType := ModeToType(entryMode)
        // updating to escape the mode and space: 
        // new content: <name>\0<hash {20 btes}>
        content = content[spaceIndex+1:] 

        nullIndex := bytes.IndexByte(content, 0)
        if nullIndex == -1 {
            return nil, fmt.Errorf("Invalid tree object format: missing null byte")
        }
        entryName := string(content[:nullIndex])
        // updating to escape the name and \0: 
        // new content: <hash {20 btes}>
        content = content[nullIndex+1:]

        if len(content) < 20 {
            return nil, fmt.Errorf("Invalid tree object format: insufficient data for hash")
        } 
        entryHash := fmt.Sprintf("%x", content[:20])
        if entryMode == "040000" {
            ParseTree([]byte(entryHash))
        }

        entries = append(entries, &TreeEntry{
            Mode: entryMode,
            Type: entryType,
            Name: entryName,
            Hash: entryHash,
        })
        //moving to another entry
        content = content[20:]
    }

    return entries, nil
}

func ModeToType(mode string) string {
    switch mode {
    case "100644", "100755", "120000":
        return "blob"
    case "040000":
        return "tree"
    case "160000":
        return "commit"
    default:
        return "unknown"
    }
}
