package entity

type Category struct {
	Id   string
	Name string
}

/*step for make a service get db use mock
step#1
- Make struct name of column from table
- Make interface on repository folder for constant function, variable input, and response
- Make function on xxx_repository file for call db (direct to db, queriying SELECT INSERT UPDATE DELETE etc...)
- Make struct on service folder through interface on repository folder
- Make function on xxx_service file for call function direct to db from repository folder
- Make function on yyy_service file for running login script what you want
*/
