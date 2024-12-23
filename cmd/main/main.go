package main

import (
	"fmt"
	"os"

	"github.com/isantoshgyawali/git-go/args"
	"github.com/isantoshgyawali/git-go/utils"
)

func main() {
    defer fmt.Print("\n")
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "usage: gitgo <command> [<args>...]\n")
        os.Exit(1)
    }

    switch command := os.Args[1]; command {
    // Initializing the git repository
    // 0755 represents file permission where "0" referring to octal
    case "init":
        for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
            if err := os.MkdirAll(dir, 0755); err != nil {
                fmt.Fprintf(os.Stderr, "Error creating directory: %s\n", err)
            }
        }

        headFileContents := []byte("ref: refs/heads/master\n")
        if err := os.WriteFile(".git/HEAD", headFileContents, 0644); err != nil {
            fmt.Fprintf(os.Stderr, "Error writing file: %s\n", err)
        }

        fmt.Printf("Initialized git directory")

    case "cat-file":
        if len(os.Args) <=3 {
            fmt.Fprintf(os.Stderr, "usage: mygit cat-file -p <object-hash>")
            os.Exit(1)
        }

        flag := os.Args[2]
        objectHash := os.Args[3]
        object, err := args.CatFile(objectHash)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error: %s\n", err)
        }

        switch flag {
        case "-t", "--type":
            fmt.Print(object.Type)
        case "-s", "--size":
            fmt.Print(object.Size)
        case "-p", "--pretty":
            if(object.Type == "tree"){
                //checking if the Content is treeNode or blob object
                if treeNodes, ok := object.Content.([]*utils.TreeNode); ok {
                    for _, entry := range treeNodes {
                        fmt.Printf("%s %s %s %s\n", entry.Mode, entry.Type, entry.Hash, entry.Name)
                    }
                } else {
                    fmt.Fprintf(os.Stderr, "Invalid content type for for tree object")
                    }
                } else {
                    fmt.Print(object.Content)
                }
        default: 
            fmt.Fprintf(os.Stderr, "Invalid flag for cat-file\nflags:\n\t-p or --pretty: Prints the contents of the object\n\t-s or --size: Displays the size of the object in bytes\n\t-t or --type: Displays the type of the object\n")
        }

    case "hash-object":
        if len(os.Args) < 3 {
            fmt.Fprintf(os.Stderr, "usage: mygit hash-object <content>\n")
            os.Exit(1)
        }

        var fileName, flag string
        var writeToDisk bool

        if len(os.Args) == 3 {
            fileName = os.Args[2]
            flag = ""
        } else {
            flag = os.Args[2]
            fileName = os.Args[3]
        }

        switch flag {
        case "-w":
            writeToDisk = true
        case "":
            writeToDisk = false
        default:
            fmt.Fprintf(os.Stderr, "Invalid flag for hash-object\n"+
                "Supported flags:\n"+
                "\t-w:\twrite the object into the object database\n")
            os.Exit(1)

        }

        objectHash, err := args.HashObject(fileName, writeToDisk)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error Occured - %s\n", err)
        }
        fmt.Printf(objectHash)

    case "ls-tree":
        if len(os.Args) < 3 {
            fmt.Fprintf(os.Stderr, "usage: mygit ls-tree <tree-sha>\n")
            os.Exit(1)
        }

        var treeHash, flag string
        if len(os.Args) == 3 {
            treeHash = os.Args[2]
        } else {
            flag = os.Args[2]
            treeHash = os.Args[3]
        }

        treeEntries, err := args.LsTree(treeHash)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error Occured - %s\n", err)           
            os.Exit(1)
        }

        switch flag {
        case "--name-only":
            for _, fileName := range treeEntries {
                fmt.Printf("%s\n", fileName.Name)
            }

        case "--object-only":
            for _, object := range treeEntries {
                fmt.Printf("%s\n", object.Hash)
            }

        case "":
            for _, entry := range treeEntries {
                fmt.Printf("%v %v %v %v\n", 
                    entry.Mode,
                    entry.Type,
                    entry.Hash,
                    entry.Name,
                    )
            }

        default:
            fmt.Fprintf(os.Stderr, "Invalid flag for ls-tree\n"+
                "Supported flags:\n"+
                "\t--name-only:\tReturns only filenames\n"+
                "\t--object-only:\tReturns only objects")
            os.Exit(1)
        }

    case "write-tree":
        currentDir, _ := os.Getwd()
        treeHash, _ := args.WriteTree(currentDir)
        fmt.Println(treeHash)
        
    default:
        fmt.Fprintf(os.Stderr, "Unknown command %s\n", command)
        os.Exit(1)
    }
}
