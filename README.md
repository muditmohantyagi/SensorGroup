# SensorGroup
This is an example  of a sensor group using Golang GIN Gorm(Postgessql)


1)
It will create two containers Postgresql
and PostgreSQL Admin PG admin username: admin@admin.com, password: root
and it will create database: test_db
PG admin URL: http://localhost:5050

docker-compose up -d --build
2) now time to add IP address of the database in .env file
docker ps -a
docker inspect <postgress containere id>
and find the IP address and put it into the .env file(DB_HOST=172.18.0.2)
3)now command inside the project folder to install the dependency
go mod tidy
4) run the migration using Cobra CLI
go run main.go migrat
5) go run main.go
http://localhost:8080
6) run the test
  go test
7) open the swager
http://localhost:8080/docs/index.html

