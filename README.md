# GoBootcamp

There is a REST API services for book store. You can manage different operations with books.
The application works with a database, and both can be run in a docker container.

The kit provides the following features right out of the box:

RESTful endpoints in the widely accepted format
Standard CRUD operations of a database table
Error handling with proper error response generation
Database migration
Data validation
Test coverage

At this time, you have a RESTFUL API server running at http://localhost:8080. It provides the following endpoints:

GET /: just return a message \
GET /books/: return all books from database which amount more than 0 \
GET /books/:id: return book by id \
POST /books: create a new book in the database \
PUT /books/:id: updates an existing book \
DELETE /books/:id: deletes a book by id \

# Run app using makefile

If you want to run application and database in a docker container: \
 ```make docker-compose ```

If you want to run database in a docker container and run app locally on the windows platform: \
 ```make windows-run ```

If you want to run linter: \
 ```make lint ```

If you want to run unit tests: \
 ```make test ```

If you want to run database in a docker container and run app locally on the windows platform: \
```make windows-run```

# Run app manual
1. Build database docker image: \

```docker build . -f db.Dockerfile -t gotraining_db```

2. Run database docker container:
```
   docker run -d --name=book_db \
	       -p 3306:3306 -v \
		/mysql_data:/var/lib/mysql \
		-e MYSQL_RANDOM_ROOT_PASSWORD=secret \
		-e MYSQL_DATABASE=books_database \
		-e MYSQL_USER=test_user \
		-e MYSQL_PASSWORD=secret \
		-it gotraining_db
```

3. Run book app:

   ```go run .```

# Examples of use in postman
Create a book: Method POST URL http://localhost:8080/books Body json {"name":"NewBook", "price":125, "genre": 1, "amount": 1}'\
Get a book by ID: Method GET URL http://localhost:8080/books/1 \
Get all books: Method GET URL http://localhost:8080/books \
Update book by ID: Method PUT URL http://localhost:8080/books/2 Body json {"id": 2, "name":"UpdatedBook", "price": 10.44, "genre": 1, "amount": 5} \
Delete book by ID: Method DELETE URL http://localhost:8080/books/1