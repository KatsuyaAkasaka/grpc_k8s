module github.com/KatsuyaAkasaka/grpc_k8s/server

go 1.12

replace (
	github.com/KatsuyaAkasaka/grpc_k8s/client => ../client
	github.com/KatsuyaAkasaka/grpc_k8s/pb => ../pb
)

require (
	github.com/KatsuyaAkasaka/grpc_k8s/pb v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.25.1
)
