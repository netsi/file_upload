# file_upload

File upload is a simple go application that exposes 2 endpoints, 
one for uploading images and the other one for accessing them.

## Prerequisites 

You'll have to install docker to be able to run the application locally see
the [official site](https://docs.docker.com/get-docker/) for the instructions
on how to install it.

## Running Locally

You can simply start the local environment by calling `make start` on the 
project root folder. Then you can use `make stop` to stop the application.

## Endpoints

### Upload Request
`POST /api/v1/file`


`curl -X POST -F "file=@/home/anonymous/Downloads/image.png" http:/localhost:8080/api/v1/file`

### Upload Response

```
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Sat, 29 Jan 2022 19:02:26 GMT
< Content-Length: 105
{
    "file_name":"image.png",
    "asset_url":"http://localhost:8080/api/v1/file/image.png"
}
```

### Fetch Image Request

`GET /api/v1/file/:file_name`

`curl -v -X GET http:/localhost:8080/api/v1/file/image.png --output image_from_server.png`

### Fetch Image Response
#### success
```
< HTTP/1.1 200 OK
< Content-Type: application/octet-stream
< Date: Sat, 29 Jan 2022 19:05:17 GMT
< Transfer-Encoding: chunked
< 
```

#### failure
`curl -v -X GET http:/localhost:8080/api/v1/file/image_other.png --output image_from_server.png`

```
< HTTP/1.1 404 Not Found
< Content-Type: application/json
< Date: Sat, 29 Jan 2022 19:09:03 GMT
< Content-Length: 2
{}
```
