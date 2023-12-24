## 1. Make sure ports below are available on your computer.
```plain text
8082: API Server

8081: MongoUI

27017: MongoDB
```



## 2. Build and run this app
```shell
# Under root directory of blog-platform, execute this command.
# If you don't have golang or docker-compose enviroment, please install it first.
make all
```

then you can try make some requests to it.

```shell

# create a new post
curl --location 'localhost:8082/posts' \
--header 'Content-Type: application/json' \
--data '{
    "title": "hello world",
    "content": "hello world, this is blog-platform."
}'

# retrieve all posts
curl --location 'localhost:8082/posts'

# retrieve post by id, you need to replace right id in url below.
curl --location 'localhost:8082/posts/${id}'


```


## 3. Read data inside MongoDB with MongoUI

http://localhost:8081

Username: admin

Passward: pass


## 4. Get Swagger Specification and HTML Docs

SwaggerUI: http://localhost:8082/swagger/index.html

Swagger API Specification: http://localhost:8082/swagger/doc.json

## 5. Clean up build output and shut down the container.
```shell
make clean
```