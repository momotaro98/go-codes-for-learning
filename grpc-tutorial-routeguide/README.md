# Premise

You need to install Protocbuf compiler.

You can follow [this article](https://medium.com/@erika_dike/installing-the-protobuf-compiler-on-a-mac-a0d397af46b8) to install.

# How to generate go file with `route_guide.proto` file

```
$ ls ./routeguide
route_guide.proto
$ protoc -I routeguide/ routeguide/route_guide.proto --go_out=plugins=grpc:routeguide
$ ls ./routeguide
route_guide.pb.go route_guide.proto
```

# How to run

In a terminal

```
$ go run server/server.go
```

In another terminal

```
$ go run client/client.go
```

# How to generate mock code with [gomock](https://godoc.org/github.com/golang/mock/gomock)

```
$ cd mock_routeguide
$ mockgen github.com/momotaro98/go-codes-for-learning/grpc-tutorial-routeguide/routeguide RouteGuideClient,RouteGuide_RouteChatClient | tee routeguide_mock.go
```
