dep/install:
	curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-osx-x86_64.zip
	mkdir -p third_party
	unzip protoc-3.19.4-osx-x86_64.zip -d third_party/protoc-3.19.4-osx-x86_64/
	rm -rf protoc-3.19.4-osx-x86_64.zip
	go install github.com/golang/protobuf/protoc-gen-go@v1.4.0

protoc/gen/go:
	./third_party/protoc-3.19.4-osx-x86_64/bin/protoc \
		--proto_path=./third_party \
		-I=./proto \
		--go_out=plugins=grpc:./proto/gen \
		./proto/test1.proto ./proto/test2.proto

# needs `eval $(minikube docker-env)`
minikube/setup: minikube/start \
	minikube/wait \
	docker/build \
	minikube/apply

minikube/destroy:
	minikube delete --all

minikube/start:
	minikube start

minikube/wait:
	sleep 30

minikube/apply:
	kustomize build manifests/dev/test1 | minikube kubectl -- apply -f -
	kustomize build manifests/dev/test2 | minikube kubectl -- apply -f -
	kustomize build manifests/dev/test3 | minikube kubectl -- apply -f -

docker/build:
	docker build \
		-t k8s-metadata-example-test1:latest \
		-f build/test1/Dockerfile .
	docker build \
		-t k8s-metadata-example-test2:latest \
		-f build/test2/Dockerfile .
	docker build \
		-t k8s-metadata-example-test3:latest \
		-f build/test3/Dockerfile .

