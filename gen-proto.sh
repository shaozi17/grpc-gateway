# change work dir to project dir
# cd $GOPATH/pkg/mod/github.com/shaozi17/grpc-gateway/protocol

# Generate gRPC stub
protoc -I=./protocol \
    -I$GOPATH/pkg/mod/github.com/shaozi17/grpc-gateway/third_party/googleapis \
    --go_out=plugins=grpc:../../../ \
    ./protocol/core/*.proto


protoc -I=./protocol \
    -I$GOPATH/pkg/mod/github.com/shaozi17/grpc-gateway/third_party/googleapis \
    --go_out=plugins=grpc:../../../ \
    ./protocol/api/*.proto


# Generate reverse-proxy
protoc -I=./protocol \
    -I$GOPATH/pkg/mod/github.com/shaozi17/grpc-gateway/third_party/googleapis \
    --grpc-gateway_out=logtostderr=true:../../../ \
    ./protocol/api/*.proto


# (Optional) Generate swagger definitions
protoc -I./protocol   \
    -I$GOPATH/pkg/mod \
    -I$GOPATH/pkg/mod/github.com/shaozi17/grpc-gateway/third_party/googleapis \
    --swagger_out=logtostderr=true:../../../ \
    ./protocol/api/*.proto
