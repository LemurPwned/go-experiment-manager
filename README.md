# go-experiment-manager
Simple experiment manager for science projects. Written in Go.


# Development 
1. Launch MongoDB:

```bash
docker run --name mongodb --it -p 27017:27017 -e MONGODB_ROOT_PASS="test" bitnami/mongodb:latest
```

2. Launch the server

```bash 
go mod tidy 
```

