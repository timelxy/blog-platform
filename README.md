1. build and run this app
```shell
# under root dir of application, exec this command
# If you don't have Go or docker-compose enviroment, please install it first.
make all
```

then you can try make some requests to it.


2. read data with MongoUI inserted in MongoDB
http://localhost:8081
Username: admin
Passward: pass


3. get Swagger Specification and HTML Docs?

SwaggerUI: http://localhost:8082/swagger/index.html
Swagger API Specification: http://localhost:8082/swagger/doc.json

4. cleanup 
```shell
make clean
```