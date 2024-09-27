package args

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func HashObject(fileName string, typeStr string) (string, error) {
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

    // Create header structure: <type> <size>\0<content>
    sizeStr := fmt.Sprintf("%d", len(content))
    // null byte representation type
    // /000: octal representation
    // /x00: hex representation
    // /u0000: hex representation
    // byte(0): byte value representation
    objectHeader := fmt.Sprintf("%s %s\000", typeStr, sizeStr)
    object := append([]byte(objectHeader), content...)

    // -- zlib compression --
    // initializing a bytes.Buffer, to temporarily hold the compressed data
    var buf bytes.Buffer
    zlibWriter := zlib.NewWriter(&buf)

    // Writes the data to the zlib writer. 
    // Write method compresses the data and writes it to buf.
    // It returns the number of bytes written and an error, if any occurred.
    if _, err := zlibWriter.Write(object); err != nil {
        return "", err
    }

    if err := zlibWriter.Close(); err != nil {
        return "", err
    }

    // Hash the compressed data: [ sha1 hashing ]
    hasher := sha1.New()
    hasher.Write(buf.Bytes())
    shaOneHash := hasher.Sum(nil)
    // or - we could just: 
    // shaOneHash := sha1.Sum(object)
    // 
    // which is more like go idiomatic: simple, concise and easy to understand
    // but first implementaion does streaming and maybe "maybe..." more optimized 
    // as entire file isn't being loaded into the memory

    // hexadecimal representation for the HashObject
    // to store in the .git/objects
    objectHash := fmt.Sprintf("%x", shaOneHash)

    // Writing the object to the file system : .git/objects/objectHash[:2]/objectHash[2:]
    objectFilePath := fmt.Sprintf(".git/objects/%s", objectHash[:2])
    if err := os.MkdirAll(objectFilePath, 0755); err != nil {
        return "", err
    }

    objectFile := fmt.Sprintf("%s/%s", objectFilePath, objectHash[2:])
    outFile, err := os.Create(objectFile); 
    if err != nil {
        return "", err
    }
    defer outFile.Close()

    if _, err := outFile.Write(buf.Bytes()); err != nil {
        return "", err
    } 

    return objectHash, nil
}
