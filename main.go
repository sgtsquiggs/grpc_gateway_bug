package grpc_gateway_bug

import (
	"fmt"

	"github.com/sgtsquiggs/grpc_gateway_bug/svc/two/proto"
)

func main() {
	f := proto.RegisterSVCTwoHandlerServer
	fmt.Println(f)
}
