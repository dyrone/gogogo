package main

import "fmt"
import "reflect"

type RPCError struct {
	Code    int64
	Message string
}

func (e *RPCError) Error() string {
	return fmt.Sprintf("%s, code=%d", e.Message, e.Code)
}

func main() {
	rpcError := &RPCError{Code: 500, Message: "Internal Server Error"}
	msg := rpcError.Error()
	fmt.Println(msg)

	//will build failed:
	//var rpcError_1 error = RPCError{Code: 400, Message: "Client Error"}
	rpcError_1 := RPCError{Code: 400, Message: "Client Error"}

	fmt.Printf("rcpError type: %v\n", reflect.TypeOf(rpcError))
	fmt.Printf("rcpError_1 type: %v\n", reflect.TypeOf(rpcError_1))

	msg = rpcError_1.Error()
	fmt.Println(msg)

}
