package hscan

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

//==========================================================================\\

var shalookup = map[string]string{}
var md5lookup = map[string]string{}

func GuessSingle(sourceHash string, filename string) {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		password := scanner.Text()

		// TODO - From the length of the hash you should know which one of these to check ...
		// add a check and logicial structure

		match, _ := regexp.MatchString("^[a-f0-9]{32}$", sourceHash) // Check for MD5 Hash

		if match {
			hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (MD5): %s\n", password)
			}
		} else {
			hash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (SHA-256): %s\n", password)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

func GenHashMaps(filename string) {

	//TODO
	//itterate through a file (look in the guessSingle function above)
	//rather than check for equality add each hash:passwd entry to a map SHA and MD5 where the key = hash and the value = password

	//TODO at the very least use go subroutines to generate the sha and md5 hashes at the same time
	// use workers to make this even faster

	//TODO create a test in hscan_test.go so that you can time the performance of your implementation
	//Test and record the time it takes to scan to generate these Maps
	// 1. With and without using go subroutines
	// 2. Compute the time per password (hint the number of passwords for each file is listed on the site...)

	startTime := time.Now()
	fmt.Printf("Loading file %q... ", filename)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
		os.Exit(0)
	}
	defer f.Close()

	total := 0
	previousLen := len(md5lookup)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		perPasswordStartTime := time.Now()
		password := strings.TrimSpace(scanner.Text())

		hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		if hash == "" {
			continue
		}
		md5lookup[hash] = password
		hash = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
		if hash == "" {
			continue
		}
		shalookup[hash] = password
		total++
		perPasswordEndTime := time.Now()
		fmt.Printf("%q password pushed into list in %q\n", password, perPasswordEndTime.Sub(perPasswordStartTime))
	}

	endTime := time.Now()
	fmt.Printf("%d lines, %d uniq lines found in %q\n", total, len(md5lookup)-previousLen, endTime.Sub(startTime))

}

func GetSHA(hash string) (string, error) {
	password, ok := shalookup[hash]
	if ok {
		return password, nil

	} else {

		return "", errors.New("password does not exist")

	}
}

//TODO
func GetMD5(hash string) (string, error) {
	password, ok := md5lookup[hash]
	if ok {
		return password, nil

	} else {

		return "", errors.New("password does not exist")

	}
}
