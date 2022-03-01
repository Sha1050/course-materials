// Build and Use this File to interact with the shodan package
// In this directory lab/3/shodan/main:
// go build main.go
// SHODAN_API_KEY=YOURAPIKEYHERE ./main <search term>

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"shodan/shodan"
)

const DEFAULT_SIZE = "10"
const DEFAULT_SORT = "timestamp"
const DEFAULT_ORDER = "desc"

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: main <searchterm>")
	}
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf(
		"Query Credits: %d\nScan Credits:  %d\n\n",
		info.QueryCredits,
		info.ScanCredits)

	hostSearch, err := s.HostSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("Host Data Dump\n")
	for _, host := range hostSearch.Matches {
		fmt.Println("==== start ", host.IPString, "====")
		h, _ := json.Marshal(host)
		fmt.Println(string(h))
		fmt.Println("==== end ", host.IPString, "====")
		//fmt.Println("Press the Enter Key to continue.")
		//fmt.Scanln()
	}

	fmt.Printf("IP, Port\n")

	for _, host := range hostSearch.Matches {
		fmt.Printf("%s, %d\n", host.IPString, host.Port)
	}

	var query string
	flag.StringVar(&query, "query", "", "Quary param")

	var page string
	flag.StringVar(&page, "page", DEFAULT_SIZE, "Page param")

	var size string
	flag.StringVar(&size, "size", DEFAULT_SIZE, "Size param")

	var sort string
	flag.StringVar(&sort, "sort", DEFAULT_SORT, "Sort param")

	var orderby string
	flag.StringVar(&orderby, "orderby", DEFAULT_ORDER, "Order by param")

	flag.Parse()

	// Call to getQueries to get list of search queries that users have saved in Shodan
	getQueries(page, sort, orderby)

	// Call to searchQueries to get list of popular tags for the saved search queries in Shodan
	searchQueries(query, page, sort, orderby)

	// Call to getQueryTags to get list of popular tags for the saved search queries in Shodan
	getQueryTags(size)

}

// getQueryTags to get list of popular tags for the saved search queries in Shodan
func getQueryTags(pageSize string) {
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	hostSearchQueryTags, err := s.GetQueryTags(pageSize)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("====  Query Tags Start  ====")

	fmt.Printf("Host Data Query tags\n")
	for _, host := range hostSearchQueryTags.Matches {
		fmt.Printf("%d, %s\n", host.Count, host.Value)
	}

	fmt.Printf("Total %d Query tags found\n", hostSearchQueryTags.Total)

	fmt.Println("====  Query Tags End  ====")

}

// getQueries to get list of search queries that users have saved in Shodan

func getQueries(pageSize string, pageSort string, pageOrder string) {

	fmt.Println("====  Get Query Method Start  ====")
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)

	hostSearchQuery, err := s.GetQueries(pageSize, pageSort, pageOrder)
	if err != nil {
		log.Panicln(err)
	}

	for _, host := range hostSearchQuery.Matches {
		fmt.Printf("%d, %s, %s, %s, %s\n", host.Votes, host.Title, host.Description, host.Query, host.Timestamp)
	}

	fmt.Printf("Total %d records found\n", hostSearchQuery.Total)

	fmt.Println("====  Get Query Method End  ====")

}

// searchQueries searches the directory of search queries that users have saved in Shodan.
func searchQueries(query string, pageSize string, pageSort string, pageOrder string) {

	fmt.Println("====  Get Search Query Method Start  ====")
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)

	hostSearchQuery, err := s.SearchQueries(query, pageSize, pageSort, pageOrder)
	if err != nil {
		log.Panicln(err)
	}

	for _, host := range hostSearchQuery.Matches {
		fmt.Printf("%d, %s, %s, %s, %s\n", host.Votes, host.Title, host.Description, host.Query, host.Timestamp)
	}

	fmt.Printf("Total %d records found\n", hostSearchQuery.Total)

	fmt.Println("====  Get Search Query Method End  ====")

}
