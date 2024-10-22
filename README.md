# Hash Checker

This program verifies the integrity of a file by comparing its calculated hash with an expected hash value.

## Features:

- Supports MD5, SHA1, SHA256, and SHA512 hashing algorithms.
- Reads the file's contents for accurate hashing.
- Provides clear output indicating whether the hashes match or not.

## Usage:

```
./hash-checker <filename> <hash algorithm> <expected hash>
```

## Example:

```
./hash-checker diskimage.iso sha256 d41d8cd98f00b204e9800998ecf8427e
```

This command checks the hash of diskimage.iso using the SHA256 algorithm and compares it to the expected hash d41d8cd98f00b204e9800998ecf8427e.

## Dependencies:

Go programming language (tested with version 1.23)

## Instructions:

- Clone the repository
- Run `go build -o hash-checker hashChecker.go` to compile the program.
- Execute the program using the provided usage instructions.

## Output:

The program will print one of the following messages:

- "Hashes match!" if the calculated hash matches the expected hash.
- "Hashes do not match." followed by the calculated hash and expected hash for detailed comparison.

## Error Handling:

The program logs fatal errors for invalid arguments or file access issues for better troubleshooting.

## License:

This project is licensed under the MIT License - see the LICENSE file for details.