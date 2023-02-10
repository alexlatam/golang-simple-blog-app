# Golang project for a simple blog 
Test how create a CRUD using Go

### Installation
Create the database using the bd.sql file.

After that run the project usign the next command[on the root folder]:
```
go run ./main.go
```

### Front 
The styles are using the BootstrapV5 framework 

### Folders
/views
    /components
    /home
    /posts

### Packages
For Mysql conection: 
- github.com/go-sql-driver/mysql 
- database/sql

For Logs
- log

For the http server
- net/http

For the Front views
- text/template

For get the format for the dates
- time

