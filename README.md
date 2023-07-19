# TEST-Project

This test-project is written purely in Go and used MariaDB(more secure than mysql) for its database
needs.Entirely Dockerized this project using Docker and DockerFile

### Frameworks used
Go-Fiber: Whole Project is completed using Go-Fiber,Which is very popular and fast framework which
is used for rapid web development
```
go get -u github.com/gofiber/fiber/v2
```

### Database used:
MariaDB: Project uses MariaDB as main DataBase and use of ORM tool named Gorm provides much better
abstraction over sql queries
```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

### Build & RUN Project
make build: Command creates a Docker image can be used later to run as containerized application which
includes all the dependencies attached to the application

make server: Command Runs the Project using native Go command to run a Go file

### Unit-Testing
SqlMock: SqlMock provides Repo-Level unit testing for projects which simulates a sql driver behavior 
without having any real connection

goMock: goMock provides interface substitution and set behaviors according to our usecases

HttpTest: A native http testing library provided by Golang to test Handlers and Servers using Rest Apis
and Http clients

### Third-Party Libraries
Koanf: Alternative To Viper configuration solution which is very lightweight compared to viper and
produces less memory footprints

Zerolog: Compared to native logger solutions Zerolog is Low to zero allocation based logger-service
which is blazingly fast compared to other solutions in market

### Major Functionalities
In-Memmory Allocation: Provides In-Memmory Allocation which refreshes the Data present in the Ram in 
every 10 seconds.Provides access to the data with less latency

Data-Compression: Data has been Compressed in APIs to reduce memmory allocation using Compresser Middleware

Grace-Full ShutDown: Application listenes for interruption signals(eg:CTRl+C) and safely shut down its processes
and close all connections

Recovery-Middleware: Server can be recovered from any sort of error to avoid unnecessary crashes.This has been
achieved by attaching a recovery middleware to the APIs
