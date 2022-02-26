# Lab 3 - PART 2 


Lab3Dir : ~/course_materials/materials/lab/3

Lab3Dir has two main folders main and shodan
main folder, where you'll run the program from (read top of main.go for details); you'll need to modify this based on which options you choose below
shodan folder, where you'll be adding additional code, do not try to run this code independently


<<<<<<< HEAD

REST API Implementation

Directory Methods
 /shodan/query (function getQueries)
 /shodan/query/search (function searchQueries)
 /shodan/query/tags (function getQueryTags)

implemented in main/main.go

Usage
List the saved search queries
Use this method to obtain a list of search queries 

Command Run: 
SHODAN_API_KEY="Replace your Key" go run main/main.go -query="webcam" -page="20" -size="10" -sort="DESC" -orderby="Field name"

Arguments:
-query="string"
-page="string"
-size="string"
-sort="string"
-orderby="string"

Parameters
query: [String] What to search for in the directory of saved search queries.
page (optional): [Integer] Page number to iterate over results; each page contains 10 items
size (optional): [Integer] The number of tags to return (default: 10).
sort (optional): [String] Sort the list based on a property. Possible values are: votes, timestamp
orderby (optional): [String] Whether to sort the list in ascending or descending order. Possible values are: asc, desc
=======
>>>>>>> 9c8d064e20c879cead8993bc286c8210fe1357de
