APP := echo
GO := GOOS=linux go
IMAGE := microservicedemo/${APP}

generate:
	protoc -I. \
	 -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	 --go_out=plugins=grpc:. \
	 --grpc-gateway_out=logtostderr=true:. \
	 service/*.proto

build:
	${GO} build -o bundles/${APP} .

image: build
	docker build -t ${IMAGE} .

run:
	-docker service rm ${APP}
	docker service create --name ${APP} --network demo -p 9090:9090 ${IMAGE}
