module github.com/KatsuyaAkasaka/grpc_k8s/server

go 1.12

replace (
	github.com/KatsuyaAkasaka/grpc_k8s/client => ../client
	github.com/KatsuyaAkasaka/grpc_k8s/pb => ../pb
)

require (
	github.com/KatsuyaAkasaka/grpc_k8s/pb v0.0.0-00010101000000-000000000000
	golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3 // indirect
	golang.org/x/tools v0.0.0-20190524140312-2c0ae7006135 // indirect
	google.golang.org/grpc v1.25.1
	honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc // indirect
)
