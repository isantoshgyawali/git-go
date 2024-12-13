package utils

import (
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

func GetObjectPath(objectHash string) string {
	// For Context:
	//  If the SHA-1 hash is abcdef1234567890abcdef1234567890abcdef12, the file path will be:
	//  .git/objects/ab/cdef1234567890abcdef1234567890abcdef12
	return fmt.Sprintf(".git/objects/%v/%v", objectHash[:2], objectHash[2:])
}

func DecompressObject(objectPath string) ([]byte, error){
	// Opening the git object file
	file, err := os.Open(objectPath)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	defer reader.Close()

	// Reading from io.ReadCloser
	content, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

    return content, nil
}
