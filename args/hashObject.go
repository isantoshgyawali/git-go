package args

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/isantoshgyawali/git-go/utils"
)

// CREATE_HEADER -> HASH -> COMPRESS -> STORE
func HashObject(fileName string, writeToDisk bool) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the Content from file
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
        
        // Initialize it to avoid the nil pointer dereferencing
        buf := &bytes.Buffer{}
        compressedObjectHash, err := utils.CompressObject("blob", content, buf)
        if err != nil {
            return "", err  
        }

        // if "-w" is available: 
	// Writing the object to the file system : .git/objects/objectHash[:2]/objectHash[2:]
        if writeToDisk {
            objectFilePath := fmt.Sprintf(".git/objects/%s", compressedObjectHash[:2])
            if err := os.MkdirAll(objectFilePath, 0755); err != nil {
                    return "", err
            }

            objectFile := fmt.Sprintf("%s/%s", objectFilePath, compressedObjectHash[2:])
            outFile, err := os.Create(objectFile)
            if err != nil {
                    return "", err
            }
            defer outFile.Close()

            // If a file with the same name already exists in the .git/objects directory, 
            // this will overwrite it silently as we want Git ensure "object immutability" by using unique hashes
            if _, err := outFile.Write(buf.Bytes()); err != nil {
                    return "", err
            }
        }

	return compressedObjectHash, nil
}
