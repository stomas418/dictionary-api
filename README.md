# Dictionary Api
Dictionary api is a small project made for practicing writing in Golang and Python. It has two sections: 

 -  In the **json-to-sql** folder, the JSON files that store the dictionary data ordered by starting letter, together with the Python fiile that stores all the data into a MySQL database
 - In the **api** folder, the Golang project to run the http server. 

## Instructions for the project

This project needs Python3, Go and MySQL installed already in your computer for it to work.

### Loading the data into the database

 1. Go into the **json-to-sql** folder.
 2. Create a `.env` file with the variables:
	 - `database` for the name of the database
	 - `host` for the host of the database
	 - `port` for the port to access the database
	 - `user` for your MySQL user
	 - `password` for your MySQL password
3. Open the terminal on the folder and run **json_to_sql.py** with the command `python json_to_sql.py` or `python3 json_to_sql.py`

### Running the Golang server
		

 1. Once the data has been saved into the database, open the **api** folder
 2. Copy the `.env` file created in the **json-to-sql** folder and paste it into the current folder.
 3. Open the terminal on the folder and run the command `go get -d ./...` to install all dependencies necessary for the server.
 4. Once all dependencies have been installed, we run the command `go run main.go` on the terminal and that's it! Now we have a running server, by default it is on `localhost:8080`.
		 

## Api Endpoints

The api accepts only two endpoints: `/:letter` and `/:letter/:word` and only the `GET` method.

### `/:letter`
 - Returns a **page** with 100 words on a JSON array formatted like this:
 
       [
    	       {
    		"word": "A",
    		"meanings": [
                    	"the 1st letter of the Roman alphabet",
                    	"the blood group whose red cells carry the A antigen"
                	]
		},
    		{
    	        "word": "A-HORIZON",
    	        "meanings": [
			"the top layer of a soil profile; usually contains humus"
			]
    	        },
    	    ...
        ]
 - Accepts parameter **letter** which is the starting letter for every word returned
 - Accepts query **page** which sets the page number. Every page has 100 words in alphabetical order and if no page is given, the default is the first hundred words.
 
### `/:letter/:word`
 - Returns a single word formatted the same as the one returned at the `/:letter` endpoint
  - It takes parameters **letter** which is the starting letter for the word returned, and **word** which is the exact word to search for.
  - If the word does not exist, it returns `{"word": "","meanings": ""}`

