pipeline {
    agent any
    tools {
        go 'go1.14'
    }
    stages {
        stage('Build gRPC docker image') {
            steps {
                echo 'Stopping existing gRPC container'
                sh 'docker stop cont-go-grpc'

                echo 'Remove existing gRPC image'
                sh 'docker rmi img-go-grpc:1.1.1'

                echo 'Building new gRPC image'
                sh 'docker build -t img-go-grpc:1.1.1 .'
            }
        }

        stage('Run gRPC docker container') {
            steps {
                echo 'Running new gRPC container'
                sh 'docker run --rm -d -p 50051:50051 --net=my_bridge --name=cont-go-grpc img-go-grpc:1.1.1'
            }
        }
    }
}
