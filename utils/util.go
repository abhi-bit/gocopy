package gocopy

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// ComputeMD5 computes md5 hash of an input file
func ComputeMD5(filePath string) (string, error) {
	var result string
	file, err := os.Open(filePath)
	if err != nil {
		return result, err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return result, err
	}

	result = hex.EncodeToString(hash.Sum(nil)[:16])
	return result, nil
}
