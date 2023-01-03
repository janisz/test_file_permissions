package main

import "os"

func no(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	filePath := "filename.txt"
	f, err := os.Create(filePath)
	no(err)
	no(f.Chmod(0000))
	no(f.Close())
	defer os.Remove(filePath)

	_, err = os.Open(filePath)
	no(err)
}