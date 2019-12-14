FROM golang:latest

ENV GO111MODULE=on

WORKDIR /usr/local/go/apps/k8s-grpc
ADD . ./

CMD [ "go", "run", "client.go" ]
