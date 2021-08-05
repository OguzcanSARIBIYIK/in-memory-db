 ### In Memory Database
 Basic concept of in memory database.

 ### Installation
 ```git clone https://github.com/OguzcanSARIBIYIK/in-memory-db```
 
 ```go run .```
 ### Listing Data
 Listing all data in database.
 
 `http://localhost:8889/all`
 
 ### Response
 ```
 {
	 "Code": 200,
	 "Data": {}
 }
 ```
 ### Store Data
 
 You can store data by key-value.
 
 `http://localhost:8889/?key=name&value=oguzcan`
 
### Response
 ```
{
	"Code": 200,
	"Data": "Data saved."
}
 ```
 
 Listing data after store.
 
  `http://localhost:8889/all`
 
 ```
 {
	"Code": 200,
	"Data": {
		"name": "oguzcan"
	}
}
 ```
