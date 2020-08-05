# gRPC Server

### Overview
gRPC server implementation for the corresponding gin server in [gin repo](https://github.com/varungupte/BootCamp_Team3)


### Models
```
type Address struct {
	HouseNo string
	Street  string
	City    string
	PIN     string
}

type Customer struct {
	Id           uint32
	FullName     string
	Addr         Address
	ActiveStatus bool
}

type Item struct {
	Id       uint32
	Name     string
	Cuisine  string
	Cost     float32
	Quantity uint32
}

type Order struct {
	Id           uint32
	ResId        uint32
	CustId       uint32
	Items        []Item
	Discount     float32
	DeliveryAddr Address
}

type Restaurant struct {
	Id           int64
	Name         string
	Items        []Item
	Addr         Address
	ActiveStatus bool
}
```

### Steps to run the gin server:
Open Terminal and copy-paste the following commands
```
1. docker network create -d bridge my_bridge
2. docker run -p 8000:8000 --net=my_bridge amazon/dynamodb-local
3. mkdir $HOME/GoWorkspace
4. export GOPATH=$HOME/GoWorkspace
5. mkdir $GOPATH/bin
6. export GOBIN=$GOPATH/bin
7. mkdir $GOPATH/src
8. cd $GOPATH/src
9. go get github.com/BhaviD/BootCamp_Team3_gRPC
10. docker network create -d bridge my_bridge
11. docker build -t img-go-grpc:1.1.1 .
12. docker run --rm -p 50051:50051 --net=my_bridge --name=cont-go-grpc img-go-grpc:1.1.1
```

### Project Directory Structure
```
BootCamp_Team3_gRPC
  ├── Dockerfile
  ├── Jenkinsfile
  ├── README.md
  ├── assets
  │   ├── customerData.json
  │   ├── orderData.json
  │   └── restaurantData.json
  ├── cmd
  │   └── grpc_server
  │       └── main.go
  ├── go.mod
  ├── go.sum
  └── pkg
      ├── dynamoDB
      │   ├── dbUtil
      │   │   └── dbUtil.go
      │   ├── dynamoDB.go
      │   ├── orderTableIndex.txt
      │   └── types
      │       └── types.go
      ├── errorutil
      │   └── errorutil.go
      └── services
          ├── grpcPb
          │   ├── grpc.pb.go
          │   └── grpc.proto
          ├── grpc_server
          │   ├── grpc_server.go
          │   └── grpc_server_test.go
          └── restaurantService
              └── restaurantService.go
```

#### /cmd
Main applications for this project.
The directory name for each application should match the name of the executable you want to have (e.g., cmd/order-prediction).
Don't put a lot of code in the application directory. If you think the code can be imported and used in other projects, then it should live in the /pkg directory. If the code is not reusable or if you don't want others to reuse it, 
put that code in the /internal directory.
It's common to have a small main function that imports and invokes the code from the /internal and /pkg directories and nothing else.

#### /pkg
Library code that's ok to use by external applications (e.g., pkg/orders, pkg/restaurants, pkg/users). 
