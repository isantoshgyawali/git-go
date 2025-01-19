package args

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"

	"github.com/isantoshgyawali/git-go/utils"
)

// fs - traversal
//  - ignore .git
//  - entries - mode, path, sha1
//  - create tree objects for any subdirectories in the path
//  - add the entry to its parent directory's tree object

// sort entries within the tree by name
// serialize the content
// calculate sha1 hash of the serialized content for each tree
// calculate the final hash
// write the tree object to the object database
func sortGitTree(fs []os.DirEntry) []os.DirEntry{
    // This function in go: takes slice as an input and let's you
    // define your custom sorting method: where i and j could be consecutive slice entries
    // but not compulsorily - sorting is handled by internal algorithm (in this case, I assume 
    // pdqsort_func() as per func def'n)
    sort.Slice(fs, func(i, j int) bool {
        fi, fj := fs[i], fs[j]
        //Directories come first
        if fi.IsDir() && !fj.IsDir() {
            return true
        }
        if !fi.IsDir() && fj.IsDir() {
            return false
        }
        // Lexigraphical order for names (case-sensitive)
        return fi.Name() < fj.Name()
    })

    return fs
}

func WriteTree(filePath string) (string, error) {
    var treeEntries []byte
    fs, err := os.ReadDir(filePath)
    if err != nil {
        return "", err 
    }
    fs = sortGitTree(fs)

    for _, entry := range fs{
        if entry.Name() == ".git" { continue }

        if entry.IsDir() {
            treeHash, err := WriteTree(filepath.Join(filePath, entry.Name())) 
            if err != nil {
                return "", err  
            }
            fmt.Println(entry.Name(), ":", treeHash)
            // Format: "040000 <name>\0<tree-hash>"
            treeEntries = append(treeEntries, fmt.Sprintf("040000 %s\000", entry.Name())...)
            treeEntries = append(treeEntries, []byte(treeHash)...)
            fmt.Println(string(treeEntries))
        } else {
            f, _ := os.Open(filepath.Join(filePath, entry.Name()))
            b, _ := io.ReadAll(f)
            defer f.Close()

            fileHash, _ := utils.CompressObject("blob", b)
            // Format: "fileMode <name>\0<tree-hash>"
            fileInfo, err := f.Stat()
            if err != nil {
              return "", err  
            }
            mode := utils.FileModeToGitMode(fileInfo.Mode())
            treeEntries = append(treeEntries, fmt.Sprintf("%s %s\000", mode, entry.Name())...)
            treeEntries = append(treeEntries, []byte(fileHash)...)
        }
    }

    treeHash, err := utils.CompressObject("tree", treeEntries)
    if err != nil {
        return "", err  
    }

    return treeHash, nil
}
