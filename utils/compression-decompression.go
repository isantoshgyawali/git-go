package utils

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func GetObjectPath(objectHash string) (string, error) {
	// For Context:
	//  If the SHA-1 hash is abcdef1234567890abcdef1234567890abcdef12, the file path will be:
	//  .git/objects/ab/cdef1234567890abcdef1234567890abcdef12
        gitRoot, err := FindGitRoot()
        if err != nil {
            return "", err  
        }
	return fmt.Sprintf("%s/objects/%v/%v", gitRoot, objectHash[:2], objectHash[2:]), nil
}

func CompressObject(objectType string, content []byte, buf *bytes.Buffer) (string, error) {
     if buf == nil {
        buf = &bytes.Buffer{}
    }

    // header format: <type> <size>\0<content>
    header := fmt.Sprintf("%s %d\000", objectType, len(content))
    // null byte representation type
    // \000: octal representation
    // \x00: hex representation
    // \u0000: hex representation
    // byte(0): byte value representation

    var object []byte
    object = append([]byte(header), content...)

    // Hash the uncompressed content with header: [ sha1 hashing ]
    hasher := sha1.New()
    hasher.Write(object)
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

    // -- zlib compression --
    zlibWriter := zlib.NewWriter(buf)

    // Writes the data to the zlib writer.
    // Write method compresses the data and writes it to buf.
    // initializing a bytes.Buffer, to temporarily hold the compressed data
    // It returns the number of bytes written and an error, if any occurred.
    if _, err := zlibWriter.Write(object); err != nil {
        return "", err
    }

    if err := zlibWriter.Close(); err != nil {
        return "", err
    }

    return objectHash, nil
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
