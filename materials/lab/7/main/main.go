package main

import (
	"flag"
	"fmt"
	"hscan/hscan"
)

func main() {

	//To test this with other password files youre going to have to hash
	var md5hash = "77f62e3524cd583d698d51fa24fdff4f"
	var sha256hash = "95a5e1547df73abdd4781b6c9e55f3377c15d08884b11738c2727dbd887d4ced"

	//TODO - Try to find these (you may or may not based on your password lists)
	var drmike1 = "36cb3251dfb3d2b4a559796498a2ac29"
	var drmike2 = "92fd7a8f8a5fd34c953c72ccbc61fe553dae49991f4a0a2579899e2400eb047a"

	// NON CODE - TODO:
	// Download and use bigger password file from: https://weakpass.com/wordlist/tiny  (want to push yourself try /small ; to easy? /big )

	//TODO: Grab the file to use from the command line instead; look at previous lab (e.g., #3 ) for examples of grabbing info from command line

	var DEFAULT_FILE = "main/wordlist.txt"
	var file string
	flag.StringVar(&file, "file", DEFAULT_FILE, "File param") // Pass -file="file name" as paramater e.g go run main/main.go -file="main/wordlist.txt"
	flag.Parse()

	hscan.GuessSingle(md5hash, file)
	hscan.GuessSingle(sha256hash, file)
	hscan.GuessSingle(drmike1, file)
	hscan.GuessSingle(drmike2, file)
	hscan.GenHashMaps(file)
	fmt.Println(hscan.GetSHA(drmike2))
	fmt.Println(hscan.GetMD5(drmike1))
}
