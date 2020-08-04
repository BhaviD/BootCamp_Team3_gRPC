pipeline {
    agent any
    tools {
        go 'go1.14'
    }
    stages {
        stage('gRPC-server docker image build') {
            steps {
                echo 'Building gRPC-server docker image'
                sh 'docker build -t img-go-grpc:1.1.1 .'
              
            }
        }

        stage('gRPC-server docker container run') {
            steps {
                echo 'Running the gRPC-server docker container'
                sh 'docker run --rm -p 50051:50051 --net=my_bridge --name=cont-go-grpc img-go-grpc:1.1.1 &'
            }
        }
    }
}
