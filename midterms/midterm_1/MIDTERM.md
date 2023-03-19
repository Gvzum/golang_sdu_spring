## Midterm 1

# My team
* Gazimov Dias 200103360

## Project Link
[Dias-Store](https://github.com/Gvzum/dias-store)

## Video Link
[Midterm-1](https://youtu.be/7PE5dUsZ92Q)

## Postman Collection 
[Dias-Store.postman_collection.json](Dias-Store.postman_collection.json)

## How To Run 
I will add docker-compose for more easy run after first midterm. But for now you can run this project by following this commands.

* First You need to clone the actual project. [Dias-Store](https://github.com/Gvzum/dias-store)
* After You need the run postgresql server with credentials what you need. 
* After running you need to specify credentials in .env file on the root directory of project. (I added the example_env file to implement)


Install all the dependencies
```
go mod tidy 
```

Run the migrations 
```
go run cmd/migration/migration.go 
```

Run server 
```
go run cmd/server/server.go       
```