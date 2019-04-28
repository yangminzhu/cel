export CEL_OUT:=${GOPATH}/out/yangminzhu

.PHONY: build
build: cel attributes_proto

.PHONY: cel
cel ${CEL_OUT}/cel: attributes_proto
	go build -o ${CEL_OUT}/cel ./main

.PHONY: output
output:
	@echo ${CEL_OUT}/cel

.PHONY: attributes_proto
attributes_proto attributes/attributes.proto:
	protoc --go_out=. attributes/*.proto

.PHONY: clean
clean:
	rm -fr attributes/*.pb.go ${CEL_OUT}/cel
