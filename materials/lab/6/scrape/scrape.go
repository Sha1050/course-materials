package scrape

import (
	"regexp"
	"sync/atomic"
)

// || GLOBAL DATA STRUCTURES  ||

const LOG_LEVEL int = 2

//CHALLENGE: Replace this Local Structure with a Key-Value DB like REDIS
type FileInfo struct {
	Filename string `json:"filename"`
	Location string `json:"location"`
}

var Files []FileInfo

var regexes = []*regexp.Regexp{
	regexp.MustCompile(`(?i)password`),
	regexp.MustCompile(`(?i).txt`),
}

var totalFilesAdded int64 = 0

// END GLOBAL VARIABLES
//==========================================================================//

//==========================================================================\\
// || HELPER FUNCTIONS TO MANIPULATE THE REGULAR EXPRESSIONS ||

func resetRegEx() {
	regexes = []*regexp.Regexp{
		regexp.MustCompile(`(?i)password`),
		regexp.MustCompile(`(?i)kdb`),
		regexp.MustCompile(`(?i)login`),
	}
}

func clearRegEx() {
	//Task15 - Validate that this works as expected and doesn't cause issues
	if len(regexes) > 0 {
		regexes = nil
	}
}

func addRegEx(regex string) {
	regexes = append(regexes, regexp.MustCompile(regex))
}

// Check if the regex exists in the regexs array else add
func isRegexExists(regex string) (result bool) {
	result = false
	for i := range regexes {
		if regexes[i] == regexp.MustCompile(regex) {
			// Found!
			result = true
			break
		}
	}

	return result
}

//==========================================================================//

// increments the number of files added and returns the new value
func incTotalFilesAdded() int64 {
	return atomic.AddInt64(&totalFilesAdded, 1)
}

// reset totle files added into files array
func setTotalFilesAdded() {
	totalFilesAdded = 0
}

// returns the current value
func getTotalFilesAdded() int64 {
	return atomic.LoadInt64(&totalFilesAdded)
}
