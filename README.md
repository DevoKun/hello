hello
=====


* A simple http server, written in go.
* Requests can be sent using GET, POST, PUT, et. al.
* Requests can be sent to **ANY** URL.
* This service will **always** return the json string: ```{ "hello": "world" }```


## HTTP Service Port

* The default port is **8765**
* The http service port can be overwritten by setting the **HTTP_PORT** environment variable
```bash
HTTP_PORT=8765
```

## Build

### Application

* _Ensure golang is installed_
* Run **build.sh**
* build.sh will produce binaries in the **target** directory which can be run on multiple platforms.

```bash
./build.sh
Building darwin-amd64
Building linux-amd64
Building windows-amd64

tree target
target
├── hello-darwin-amd64
├── hello-linux-amd64
└── hello-windows-amd64.exe
```

### Docker Image

* _Ensure docker is installed_
* The **build_docker.sh** script will build the docker image.
* Build the docker image **after** building the application.
* To push to the Docker Registry, set the **PUSH_TO_REGISTRY** environment variable to true.

```bash
PUSH_TO_REGISTRY=true ./build_docker.sh
```


## Using the Docker Container

```bash
docker pull devokun/hello

docker run -d --name hello -p 3000:3000 -e "HTTP_PORT=3000" devokun/hello

curl http://0.0.0.0:3000/

```













