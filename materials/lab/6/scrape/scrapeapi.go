package scrape

// scrapeapi.go HAS TEN TODOS - TODO_5-TODO_14 and an OPTIONAL "ADVANCED" ASK

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"

	"github.com/gorilla/mux"
)

//==========================================================================\\

// Helper function walk function, modfied from Chap 7 BHG to enable passing in of
// additional parameter http responsewriter; also appends items to global Files and
// if responsewriter is passed, outputs to http

// Get root folder of this package and append location into root folder
func GetRootDir(location string) string {
	var (
		_, b, _, _ = runtime.Caller(0)
		rootDir    = filepath.Join(filepath.Dir(b), "../"+location)
	)
	return rootDir
}

// Chec if the file and location is already exists in the files
func isExists(fileName string, location string) (found bool) {
	found = false
	for i := range Files {
		if Files[i].Filename == fileName && Files[i].Location == location {
			// Found!
			found = true
			break
		}
	}
	return found
	// End block
}
func walkFn(w http.ResponseWriter) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		w.Header().Set("Content-Type", "application/json")

		for _, r := range regexes {
			if r.MatchString(path) {
				var tfile FileInfo
				dir, filename := filepath.Split(path)
				tfile.Filename = string(filename)
				tfile.Location = string(dir)

				//TODO_5: As it currently stands the same file can be added to the array more than once
				//TODO_5: Prevent this from happening by checking if the file AND location already exist as a single record
				exists := isExists(tfile.Filename, tfile.Location)
				if !exists {
					Files = append(Files, tfile)

					if w != nil {

						//TODO_6: The current key value is the LEN of Files (this terrible);
						//TODO_6: Create some variable to track how many files have been added
						incTotalFilesAdded()
						w.Write([]byte(`"` + (strconv.FormatInt(int64(len(Files)), 10)) + `":  `))
						json.NewEncoder(w).Encode(tfile)
						w.Write([]byte(`,`))

					}
					if LOG_LEVEL > 1 {
						log.Printf("[+] HIT: %s\n", path)
					}

				}
			}
		}
		return nil
	}
}

//TODO_7: One of the options for the API is a query command
//TODO_7: Create a walkFn2 function based on the walkFn function,
//TODO_7: Instead of using the regexes array, define a single regex
//TODO_7: Hint look at the logic in scrape.go to see how to do that;
//TODO_7: You won't have to itterate through the regexes for loop in this func!

func walkFn2(w http.ResponseWriter, query string) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		w.Header().Set("Content-Type", "application/json")

		if regexp.MustCompile(query).MatchString(path) {
			var tfile FileInfo
			dir, filename := filepath.Split(path)
			tfile.Filename = string(filename)
			tfile.Location = string(dir)

			//TODO_5: As it currently stands the same file can be added to the array more than once
			//TODO_5: Prevent this from happening by checking if the file AND location already exist as a single record
			exists := isExists(tfile.Filename, tfile.Location)
			if !exists {
				Files = append(Files, tfile)

				if w != nil {

					//TODO_6: The current key value is the LEN of Files (this terrible);
					//TODO_6: Create some variable to track how many files have been added
					incTotalFilesAdded()
					w.Write([]byte(`"` + (strconv.FormatInt(int64(len(Files)), 10)) + `":  `))
					json.NewEncoder(w).Encode(tfile)
					w.Write([]byte(`,`))

				}
				if LOG_LEVEL > 1 {
					log.Printf("[+] HIT: %s\n", path)
				}
			}

		}

		return nil

	}
}

//==========================================================================\\

func APISTATUS(w http.ResponseWriter, r *http.Request) {

	if LOG_LEVEL == 1 {
		log.Printf("Entering %s end point", r.URL.Path)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if LOG_LEVEL > 1 {
		w.Write([]byte(`{ "status" : "API is up and running ",`))
	}
	var regexstrings []string

	for _, regex := range regexes {
		regexstrings = append(regexstrings, regex.String())
	}
	w.Write([]byte(` "regexs" :`))
	json.NewEncoder(w).Encode(regexstrings)
	w.Write([]byte(`}`))
	log.Println(regexes)

}

func MainPage(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL == 1 {
		log.Printf("Entering %s end point", r.URL.Path)
	}
	w.Header().Set("Content-Type", "text/html")

	w.WriteHeader(http.StatusOK)
	//Task8 - Write out something better than this that describes what this api does
	fmt.Fprintf(w, "<html><body")
	fmt.Fprintf(w, "<H1>Welcome to CEDAR: Security is an investment, not an expense.....</H1>")
	fmt.Fprintf(w, "<p>1: <a href='http://localhost:8080/indexer?location=/' target='_blank' >http://localhost:8080/indexer?location=/ </a> --> [SHAH_ROOTCHECK.HTML]</p>")
	fmt.Fprintf(w, "<p>2. <a href='http://localhost:8080/clear' target='_blank'>http://localhost:8080/clear</a>  ;  <a href='http://localhost:8080/addsearch/go' target='_blank'>http://localhost:8080/addsearch/go</a> ; <a href='http://localhost:8080/api-status' target='_blank'>http://localhost:8080/api-status</a>  --> [SHAH_CLEARSETCHECK.HTML]</p>")
	fmt.Fprintf(w, "<p>3. <a href='http://localhost:8080/search?q=scrapeapi.go' target='_blank'>http://localhost:8080/search?q=scrapeapi.go</a> --> [SHAH_CHECKSEARCH.HTML]</p>")
	fmt.Fprintf(w, "<p>4. <a href='http://localhost:8080/search' target='_blank'>http://localhost:8080/search</a>  --> [SHAH_ALL_PRERESET.HTML]</p>")
	fmt.Fprintf(w, "<p>5. <a href='http://localhost:8080/reset' target='_blank'>http://localhost:8080/reset</a>  ; <a href='http://localhost:8080/indexer?location=/&regex=go' target='_blank'>http://LOCALHOST:8080/indexer?location=/&regex=go</a> --> [SHAH_CUSTOMREGEX.HTML]</p>")
	fmt.Fprintf(w, "<p>6. <a href='http://localhost:8080/search' target='_blank'>http://localhost:8080/search</a> --> [SHAH_FINAL.HTML]</p>")

	fmt.Fprintf(w, "</body>")
}

func FindFile(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL == 1 {
		log.Printf("Entering %s end point", r.URL.Path)
	}
	q, ok := r.URL.Query()["q"]

	w.WriteHeader(http.StatusOK)
	if ok && len(q[0]) > 0 {
		if LOG_LEVEL == 1 {
			log.Printf("Entering search with query=%s", q[0])
		}

		var found = false

		// ADVANCED: Create a function in scrape.go that returns a list of file locations; call and use the result here
		// e.g., func finder(query string) []string { ... }

		for _, File := range Files {
			if File.Filename == q[0] {
				json.NewEncoder(w).Encode(File.Location)
				//consider FOUND = TRUE
				found = true

			}
		}
		//TODO_9: Handle when no matches exist; print a useful json response to the user; hint you might need a "FOUND variable" to check here ...
		if !found {
			w.Write([]byte(`{ "parameters" : {"required": "q"},`))
			w.Write([]byte(`"examples" : { "required": "/search?q=main.md"},`))
			w.Write([]byte(` "status": "File not found!"} `))
			if LOG_LEVEL > 1 {
				log.Printf("File '%s' not found", q[0])
			}
		}

	} else {
		// didn't pass in a search term, show all that you've found
		w.Write([]byte(`"files":`))
		json.NewEncoder(w).Encode(Files)
	}
}

func IndexFiles(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL == 1 {
		log.Printf("Entering %s end point", r.URL.Path)
	}
	w.Header().Set("Content-Type", "application/json")

	location, locOK := r.URL.Query()["location"]

	//TODO_10: Currently there is a huge risk with this code ... namely, we can search from the root /
	//TODO_0: Assume the location passed starts at /home/ (or in Windows pick some "safe?" location)
	//TODO_10: something like ...  rootDir string := "???"
	//TODO_10: create another variable and append location[0] to rootDir (where appropriate) to patch this hole

	// Get root folder of this package and append location[0] into root folder
	rootDir := GetRootDir(location[0])

	if locOK && len(location[0]) > 0 {
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusFailedDependency)
		w.Write([]byte(`{ "parameters" : {"required": "location",`))
		w.Write([]byte(`"optional": "regex"},`))
		w.Write([]byte(`"examples" : { "required": "/indexer?location=/xyz",`))
		w.Write([]byte(`"optional": "/indexer?location=/xyz&regex=(i?).md"}}`))
		return
	}
	location[0] = rootDir

	// reset total files added for new request
	setTotalFilesAdded()

	//wrapper to make "nice json"
	w.Write([]byte(`{ `))

	// TODO_11: Currently the code DOES NOT do anything with an optionally passed regex parameter
	// Define the logic required here to call the new function walkFn2(w,regex[0])
	// Hint, you need to grab the regex parameter (see how it's done for location above...)

	// if regexOK
	//   call filepath.Walk(location[0], walkFn2(w, `(i?)`+regex[0]))
	// else run code to locate files matching stored regular expression

	regex, regexOK := r.URL.Query()["regex"]

	// if regexOK
	if regexOK {
		//   call filepath.Walk(location[0], walkFn2(w, `(i?)`+regex[0]))
		filepath.Walk(location[0], walkFn2(w, `(i?)`+regex[0]))
		if err := filepath.Walk(location[0], walkFn2(w, `(i?)`+regex[0])); err != nil {
			if LOG_LEVEL > 0 {
				log.Panicln(err)
			}
		}
	} else {
		//run code to locate files matching stored regular expression
		if err := filepath.Walk(location[0], walkFn(w)); err != nil {
			if LOG_LEVEL > 1 {
				log.Panicln(err)
			}
		}
	}

	//wrapper to make "nice json"
	w.Write([]byte(` "Total Files Added":  `))
	w.Write([]byte(`"` + (strconv.FormatInt(getTotalFilesAdded(), 10)) + `" `))
	w.Write([]byte(`,`))

	w.Write([]byte(` "status": "completed"} `))

}

//TODO_12 create endpoint that calls resetRegEx AND *** clears the current Files found; ***
//TODO_12 Make sure to connect the name of your function back to the reset endpoint main.go!
func ResetRegex(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL == 1 {
		log.Printf("Entering %s end point", r.URL.Path)
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	resetRegEx()
	if len(Files) > 0 {
		Files = nil
	}

}

//TODO_13 create endpoint that calls clearRegEx ;
//TODO_12 Make sure to connect the name of your function back to the clear endpoint main.go!
func ClearRegex(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL == 1 {
		log.Printf("Entering %s end point", r.URL.Path)
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	clearRegEx()
	if len(Files) > 0 {
		Files = nil
	}

}

//TODO_14 create endpoint that calls addRegEx ;
//TODO_12 Make sure to connect the name of your function back to the addsearch endpoint in main.go!
// consider using the mux feature
// params := mux.Vars(r)
// params["regex"] should contain your string that you pass to addRegEx
// If you try to pass in (?i) on the command line you'll likely encounter issues
// Suggestion : prepend (?i) to the search query in this endpoint

func AddRegex(w http.ResponseWriter, r *http.Request) {
	if LOG_LEVEL == 1 {
		log.Printf("Entering %s end point", r.URL.Path)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	addRegEx(`(?i)` + params["regex"])
	fmt.Fprintf(w, "Regex added successfully")
	if LOG_LEVEL == 1 {
		log.Printf("Regex %s added successfully", `(?i)`+params["regex"])
	}

}
