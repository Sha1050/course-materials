# Lab 4

## Required [10 points]
- Update wyoassign.go by specifically updating: `func UpdateAssignment`
- Update main.go to enable this new route
- Use postman to exercise/run through ALL of the available endpoints
- Save the responses from each run and place them in one text file


## Option 1
- Create a new set of endpoints for "classes" 
  - create new data structure(s) and 
  - create new endpoints (minimum create, get, delete)
  - update main.go 

## Option 2
- Modify/Harden the existing wyoassign.go endpoints
- Current endpoints lack real testing
- The POST / Create Assignment endpoint is terrible

## Option 3
- Create a wyoassign_test.go file which tests the functionality of the code
----------------------------------------------------------------------------------------------------------------
Check api status:
curl --location --request GET 'http://localhost:8080/api-status'
Response:
API is up and running

Get Job:
curl --location --request GET 'http://localhost:8080/assignment/Mike1A'
Response:
{
    "id": "Mike1A",
    "Title": "assignment",
    "desc": "some information....",
    "points": 20
}

Create New Job:
curl --location --request POST 'http://localhost:8080/assignments?id=001&title=New assignment Title&desc=Here is some description text here&points=40' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id":"0001",
    "title": "Creating title",
    "desc": "Description for the assignment",
    "points":"30"
}'


Update Method:
curl --location --request PUT 'http://localhost:8080/Assignments/001?title=New Job Title Update&desc=Here is some description text here with update&points=50' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id":"0001",
    "title": "Creating title",
    "desc": "Description for the job",
    "points":"30"
}'
Response:
{"id":"001","Title":"New assignment Title Update","desc":"Here is some description text here with update","points":50}

/*******************************************
                STUDENT CLASS
********************************************/

// Student Information 
Get all student api endpoint and its responses
endpoint: http://localhost:8080/students
Method: GET
Response:
[
    {
        "id": "1",
        "Name": "First Student",
        "father_name": "First student father name",
        "class": "O Level",
        "section": "C"
    },
    {
        "id": "2",
        "Name": "Kami",
        "father_name": "Rola EM",
        "class": "A Level",
        "section": "BC"
    },
    {
        "id": "3",
        "Name": "Chirs",
        "father_name": "Gull",
        "class": "5th grade",
        "section": "A"
    },
    {
        "id": "4",
        "Name": "Khan",
        "father_name": "Khan-2",
        "class": "6th grade",
        "section": "D"
    }
]
// Get Single student information api endpoint and its Response
endpoint http://localhost:8080/student/{id} // Replace {id} with student id which is "1" in my case
Method: GET
Response:
{
    "id": "1",
    "Name": "First Student",
    "father_name": "First student father name",
    "class": "O Level",
    "section": "C"
}

// Create New Student api endpoint and its Response
endpoint http://localhost:8080/student
Method: POST
Payload:
{
    "Name": "Kami",
    "father_name": "Rola EM",
    "class": "A Level",
    "section": "BC"
}
Response:
{
    "id": "2",
    "Name": "Kami",
    "father_name": "Rola EM",
    "class": "A Level",
    "section": "BC"
}

// Delete Student Information api endpoint and its Response
endpoint http://localhost:8080/student/{id} // Replace {id} with student id which is "4" in my case
Method: DELETE
Response:
{"status":"Success"}

// Update student information api endpoint and its Response
endpoint http://localhost:8080/student/{id} // Replace {id} with student id which is "2" in my case
Method: PUT
Payload:
{
        "Name": "Kami",
        "father_name": "Rola EM",
        "class": "O Level",
        "section": "F"
}
Response:
{
    "id": "2",
    "Name": "Kami",
    "father_name": "Rola EM",
    "class": "O Level",
    "section": "F"
}
