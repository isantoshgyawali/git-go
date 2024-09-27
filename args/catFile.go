package args

import (
    "bytes"
    "compress/zlib"
    "fmt"
    "io"
    "os"
)

func CatFile(objectHash string) (objectType string, objectContent string, objectSize int64, err error) {
    // For Context:
    //  If the SHA-1 hash is abcdef1234567890abcdef1234567890abcdef12, the file path will be:
    //  .git/objects/ab/cdef1234567890abcdef1234567890abcdef12
    path := fmt.Sprintf(".git/objects/%v/%v", objectHash[:2], objectHash[2:])

    // Opening the git object file
    file, err := os.Open(path)
    if err != nil {
        return "", "", 0, err
    }
    defer file.Close()

    // [ Git Uses zlib for compression ]
    //
    // zlib: compression library that provides in-memory data compression and 
    // decompression functions, using the DEFLATE compression algorithm
    // commonly used for compressing data to save space or to transmit data more
    // efficiently over a network.

    // Initializes a reader to decompress zlib-compressed data from the provided source.
    // Returns an io.ReadCloser that allows reading decompressed data.
    reader, err := zlib.NewReader(file)
    if err != nil {
        return "", "", 0, err
    }
    defer reader.Close()

    // Reading from io.ReadCloser
    content, _ := io.ReadAll(reader)

    // format for object header returned by zlib decompression:
    // <type> <size>\0<content>
    nullByteIndex := bytes.IndexByte(content, 0)
    if nullByteIndex == -1 {
        return "", "", 0, fmt.Errorf("Invalid Git Object Format")
    }

    objectType = string(content[:bytes.IndexByte(content, ' ')])
    objectContent = string(content[nullByteIndex+1:])
    objectSize = int64(len(objectContent))

    return objectType, objectContent, objectSize, nil
}
