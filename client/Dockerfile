FROM golang:1.12 as builder

ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64

RUN cd / && mkdir k8s-grpc
WORKDIR /k8s-grpc
COPY ./client ./client
COPY ./pb ./pb
RUN cd client && go build -o image client.go
RUN chmod 777 client/image

FROM alpine
ENV GOOS=linux
ENV GOARCH=amd64
EXPOSE 8080
COPY --from=builder /k8s-grpc/client/image /image
RUN mkdir /lib64
RUN ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
CMD ["/image"]
