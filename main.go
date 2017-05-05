package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/abhi-bit/gocopy/utils"
)

func main() {
	err := gocopy.CreateArchive(os.Args[1], []string{os.Args[2]})
	if err != nil {
		fmt.Println("CreateArchive failure, err", err)
		return
	}

	err = gocopy.Unarchive(os.Args[1], ".")
	if err != nil {
		fmt.Println("Unarchive failure, err", err)
		return
	}

	o, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("os.Open failure, err", err)
		return
	}

	info, err := o.Stat()
	if err != nil {
		fmt.Println("os.File.Stat failure, err", err)
		return
	}

	md5Hash, err := gocopy.ComputeMD5(info.Name())
	if err != nil {
		fmt.Println("ComputeMD5 failure, err", err)
		return
	}

	fStat := gocopy.FileStat{
		Name:     info.Name(),
		Size:     info.Size(),
		Checksum: md5Hash,
	}

	encodedData, err := json.Marshal(&fStat)
	if err != nil {
		fmt.Println("marshalling failure, err", err)
		return
	}

	fmt.Println(string(encodedData))
}
