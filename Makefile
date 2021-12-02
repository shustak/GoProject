hello:
	echo "Globant Go training course"

## windows-run: Start the services on windows
windows-run: db-docker-build db-docker-run app-windows-run

## build-linux: Build application on linux
build-linux:
	GOOS=linux GOARCH=386 go build -o bin/GoTraining-linux main.go

## windows-build: Build app on windows
app-windows-build:
	go build -o bin/GoTraining main.go

## windows-run: Run app on windows
app-windows-run:
	go run .

## test: Run unit tests
test:
	go test -v

## db-docker-build: Build docker database image
db-docker-build:
	docker build . -f db.Dockerfile -t gotraining_db

## db-docker-run: Run docker database
db-docker-run:
	docker run -d --name=book_db \
			   -p 3306:3306 -v \
			   /mysql_data:/var/lib/mysql \
			   -e MYSQL_RANDOM_ROOT_PASSWORD=secret \
			   -e MYSQL_DATABASE=books_database \
			   -e MYSQL_USER=test_user \
			   -e MYSQL_PASSWORD=secret \
			   -it gotraining_db

## docker-compose: Run database and app in a docker
docker-compose:
	docker-compose build --no-cache
	docker-compose up

## lint: Start linter
lint:
	golangci-lint run

## clean: Clean build files.
clean:
	rm -f GoTraining
	rm -f GoTraining-linux