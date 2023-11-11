# GoSentinel

English | [Chinese](README_zh.md)

GoSentinel: A comprehensive Go-based tool for Linux process management, monitoring, and system performance analysis. It offers cluster management, server handling, advanced process control, security auditing, user management, and real-time performance insights.


## Benefits
- Powerful performance:
    - Golang-based development
- Features
    - Manage servers in clusters
    - Process guarding
    - Performance visualization
    - Privilege management
- Rich usage scenarios

## Deployment and Usage

### Binary deployment
```shell
# get https://github.com/WingYat/GoSentinel/releases/new
$ . /gosentinel --config=./config/config.yaml 
$ nohup . /gosentinel --config=./config/config.yaml > /dev/null 2>&1 &
```

### Deploying the docker way
```shell
docker run -d \
  --name gosentinel \
  --privileged=true \
  --net=host \
  --p 8080:8080 \
  wingyat/gosentinel:latest
```

### Deploying the source code way
```shell
$ git clone git@github.com:WingYat/GoSentinel.git
$ cd GoSentinel
$ make
$ . /gosentinel --config=./config/config.yaml 
$ nohup . /gosentinel --config=./config/config.yaml  > /dev/null 2>&1 &
```

## Documentation
Please visit [GoSentinel documentation](docs/main.md).

## Author.
* [WingYat Cheung](https://github.com/WingYat)

## License
GoSentinel is licensed under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0).

# GoSentinel
Linux process management tools based on golang development.