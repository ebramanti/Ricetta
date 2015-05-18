Ricetta
======

* Discover and share recipes
* Recipe management
* Advanced recipe curation
* A global community of fellow cooks

### Install

1. [Download](http://neo4j.com/download/) and install Neo4j. Select community edition.

2. Start the Neo4j from the bin executable

    ```
    	$ cd neo4j-community-VERSION/
    	$ bin/neo4j start
       
    ```

3. Go to the [Neo4j browser](http://neo4j.com/developer/guide-neo4j-browser/) and change your username and password from the default of `neo4j/neo4j`. *If you don't do this you will wonder why Neo4j/tests are not working.* Just change the password to something more secure.

4. Go into `config/` and change `default.json` and `test.json` to have the updated neo4j credentials. `default` is used when normally running the app. `test` is used when running `npm test`. Should be the same to start off with.

	```
		{
  			"neo4j": {
   				"username": "neo4j",
    			"password": "5pcAIGhnXlyJfZ1L3mrXx5ue"
			}
		}
	```

5. Prevent Git from catching on to you config file changes

    ```
		$ git update-index --assume-unchanged config/*.json
    ```

6. Run the tests. Refrain from obsessing over test coverage.

    ```
        $ npm test
        $ npm run coverage
    ```

