#Whaletrace b2b client - golang

## Client examples

You can eventually transpile proto file to `*.go` file with `protoc -I .. -I $GOPATH/src --go_out=plugins=grpc:./ ../types.proto`

This will generate transpiled file into current directory. To adhere golang conventions and also to make these examples work, compiler should be able to find this file. Otherwise put it into your GOPATH directory (`$GOPATH/src/types`) and then just import it like import types.
