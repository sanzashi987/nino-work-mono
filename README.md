## Get started

[Quick start & environment configuration](https://grpc.io/docs/languages/go/quickstart/)

#### For `Protocol Buffer`:
  - Go to [download](https://github.com/protocolbuffers/protobuf/releases) release for the platform. Should be an archive. Unzip and make sure the `bin/` directory from the content is added to the system `PATH` variable
    ```
     protoc --version
    ```
  - Install __GO Plugins__ for the protocol compiler:
    ```
    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    ```
  - For `go-micro` framework, additionaly download `micro-generator`
    ```
    $ go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest
    ```
  - Update the your `PATH` by adding the `GOPATH` (temporarily or permenantly)
    ```
    $ export PATH="$PATH:$(go env GOPATH)/bin"
    ```

And you should be able to use following command to generate pb.go & micro.pb.go
```
protoc --micro_out=. --go_out=. -I=./proto
```

#### Install `fresh` (golang hot reload dev server)
Run the config locates at the `apps` directories
```bash
go install github.com/gravityblast/fresh@latest
```

Start the dev server (make sure the repo directory root is the working directory)
```
fresh -c ./user/app/runner.conf
```

#### Install `etcd`
[Etcd release page](https://github.com/etcd-io/etcd/releases/)



#### Bootstrap the system
```bash
# cd to the etcd path, gerernally locates at ~/etcd/
./etcd
```


## TODOS
  
- [ ] Add makefile for quick batch generating proto files to individual directories