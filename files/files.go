package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func directory() {
	_, err := os.Stat("test")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("test", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}

}

func emptyFile() {
	emptyFile, err := os.Create("abc.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(emptyFile)
	emptyFile.Close()
}

func renameFile() {
	oldName := "empty.txt"
	newName := "testing.txt"
	err1 := os.Rename(oldName, newName)
	if err1 != nil {
		log.Fatal(err1)
	}

}

func moveFile() {
	oldLocation := "/Users/gauravsingh/Documents/GitHub/golang/files/testing.txt"
	newLocation := "/Users/gauravsingh/Documents/GitHub/golang/files/test/testing.txt"
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}

func fileStat() {
	fileStat, err := os.Stat("abc.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File Name:", fileStat.Name())        // Base name of the file
	fmt.Println("Size:", fileStat.Size())             // Length in bytes for regular files
	fmt.Println("Permissions:", fileStat.Mode())      // File mode bits
	fmt.Println("Last Modified:", fileStat.ModTime()) // Last modification time
	fmt.Println("Is Directory: ", fileStat.IsDir())   // Abbreviation for Mode().IsDir()

}

func deleteFile() {
	err := os.Remove("/Users/gauravsingh/Documents/GitHub/golang/files/abc.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func readingFileChar() {
	filename := "abc.txt"

	filebuffer, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputdata := string(filebuffer)
	data := bufio.NewScanner(strings.NewReader(inputdata))
	data.Split(bufio.ScanRunes)

	for data.Scan() {
		fmt.Print(data.Text())
	}
}

func truncateFileContent() {
	err := os.Truncate("abc.txt", 100)

	if err != nil {
		log.Fatal(err)
	}
}

func appendContextToFile() {
	message := "Add this content at end"
	filename := "abc.txt"

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer f.Close()

	fmt.Fprintf(f, "%s\n", message)
}

func changeFilePermissionOwnershipTimestamp() {
	_, err := os.Stat("abc.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}
	log.Println("File exist.")

	// Change permissions Linux.
	err = os.Chmod("abc.txt", 0777)
	if err != nil {
		log.Println(err)
	}

	// Change file ownership.
	err = os.Chown("abc.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Println(err)
	}

	// Change file timestamps.
	addOneDayFromNow := time.Now().Add(24 * time.Hour)
	lastAccessTime := addOneDayFromNow
	lastModifyTime := addOneDayFromNow
	err = os.Chtimes("test.txt", lastAccessTime, lastModifyTime)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	//fileStat()
	//deleteFile()
	//emptyFile()
	//readingFileChar()
	//truncateFileContent()
	//appendContextToFile()
	changeFilePermissionOwnershipTimestamp()
}
