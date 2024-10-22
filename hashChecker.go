package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
	"strings"
)

func printUsage() {
	fmt.Println("Usage: " + os.Args[0] + " <filename> <hash algorithm (md5, sha1, sha256, sha512)> <expected hash>")
	fmt.Println("Example: " + os.Args[0] + " diskimage.iso md5 d41d8cd98f00b204e9800998ecf8427e")
}

func checkArgs() (string, string, string) {
	if len(os.Args) < 4 {
		printUsage()
		os.Exit(1)
	}

	filename := os.Args[1]
	hashAlgo := os.Args[2]
	expectedHash := os.Args[3]
	return filename, hashAlgo, expectedHash
}

func calculateHash(file *os.File, hashFunc func() hash.Hash) ([]byte, error) {
	defer file.Close() // Ensure file closure even on errors

	hasher := hashFunc()

	// Default buffer size for copying is 32*1024 or 32kb per copy
	// Use io.CopyBuffer() if you want to specify the buffer to use
	// It will write 32kb at a time to the digest/hash until EOF
	// The hasher implements a Write() function making it satisfy
	// the writer interface. The Write() function performs the digest
	// at the time the data is copied/written to it. It digests
	// and processes the hash one chunk at a time as it is received.
	_, err := io.Copy(hasher, file)
	if err != nil {
		return nil, err
	}

	// Now get the final sum or checksum.
	// We pass nil to the Sum() function because
	// we already copied the bytes via the Copy to the
	// writer interface and don't need to pass any new bytes
	return hasher.Sum(nil), nil
}

func main() {
	filename, hashAlgo, expectedHash := checkArgs()

	var hashFunc func() hash.Hash
	switch strings.ToLower(hashAlgo) {
	case "md5":
		//Create new hasher, which is a writer interface
		hashFunc = md5.New
	case "sha1":
		hashFunc = sha1.New
	case "sha256":
		hashFunc = sha256.New
	case "sha512":
		hashFunc = sha512.New
	default:
		log.Fatal("Invalid hash algorithm:", hashAlgo)
	}

	// Open file for reading
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	checksum, err := calculateHash(file, hashFunc)
	if err != nil {
		log.Fatal(err)
	}

	calculatedHash := hex.EncodeToString(checksum)

	if calculatedHash == strings.ToLower(expectedHash) {
		fmt.Println("Hashes match!")
	} else {
		fmt.Println("Hashes do not match.")
		fmt.Printf("Calculated hash: %s\n", strings.ToUpper(calculatedHash))
		fmt.Printf("Expected hash:   %s\n", strings.ToUpper(expectedHash))
	}
}
