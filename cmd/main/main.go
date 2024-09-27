package main

import (
	"fmt"
	"os"

	"github.com/u/gitclone/args"
)

func main() {
    defer fmt.Print("\n")
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "usage: mygit <command> [<args>...]\n")
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

        fmt.Println("Initialized git directory")

    case "cat-file":
        if len(os.Args) <=3 {
            fmt.Fprintf(os.Stderr, "usage: mygit cat-file -p <object-hash>\n")
            os.Exit(1)
        }

        flag := os.Args[2]
        objectHash := os.Args[3]
        objectType, objectContent, objectSize, err := args.CatFile(objectHash)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error: %s\n", err)
        }

        if (flag == "-t" || flag == "--type") {
            fmt.Print(objectType)
        } else if (flag == "-p" || flag == "--pretty") {
            fmt.Print(objectContent)
        } else if (flag == "-s" || flag == "--size") {
            fmt.Print(objectSize)
        } else {
            fmt.Fprintf(os.Stderr, "Invalid flag for cat-file\nflags:\n\t-p or --pretty: Prints the contents of the object\n\t-s or --size: Displays the size of the object in bytes\n\t-t or --type: Displays the type of the object\n")
        }

    case "hash-object":
        if len(os.Args) <=3 {
            fmt.Fprintf(os.Stderr, "usage: mygit hash-object -w <content>\n")
            os.Exit(1)
        }

        flag := os.Args[2]
        fileName := os.Args[3]
        objectHash, err := args.HashObject(fileName, "blob")
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error Occured - %s\n", err)
        }

        if (flag == "-w") {
            fmt.Println(objectHash)
        }

    default:
        fmt.Fprintf(os.Stderr, "Unknown command %s\n", command)
        os.Exit(1)
    }
}
