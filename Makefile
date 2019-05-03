export CEL_OUT:=${GOPATH}/out/yangminzhu

.PHONY: build
build: ${CEL_OUT}/cel attributes_proto

.PHONY: cel
${CEL_OUT}/cel: attributes_proto
	go build -o ${CEL_OUT}/cel ./main

.PHONY: attributes_proto
attributes_proto attributes/attributes.proto:
	protoc --go_out=. attributes/*.proto

.PHONY: out
out:
	@echo ${CEL_OUT}/cel

.PHONY: clean
clean:
	rm -fr attributes/*.pb.go ${CEL_OUT}/cel
