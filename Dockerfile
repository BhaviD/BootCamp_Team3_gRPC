FROM golang:1.14
RUN mkdir /usr/go-grpc-server
COPY . /usr/go-grpc-server
WORKDIR /usr/go-grpc-server
RUN go build -o grpc_server ./cmd/grpc_server
EXPOSE 50051
CMD ["/usr/go-grpc-server/grpc_server"]