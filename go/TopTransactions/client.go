package main

import (
	"context"
	"fmt"
	"io"
	"time"

	types ".."

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

const Server = "grpc.whaletrace.com:30000" //replace with server connection
const TOKEN = "YOUR_TOKEN"  //replace with your api token
const ASSET = "BTC"                        //replace with different asset if interested
const AMOUNT = 10                          //number of returned transactions

type tokenAuth struct {
	token string
}

func (t tokenAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"Authorization": "Bearer " + t.token,
	}, nil
}

func (tokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	conn, err := grpc.Dial(Server, grpc.WithInsecure(), grpc.WithPerRPCCredentials(tokenAuth{TOKEN}), grpc.WithTimeout(5*time.Second), grpc.WithBlock())

	if err != nil {
		panic(fmt.Sprintf("did not connect: %v", err))
	}
	fmt.Println("Connected")

	client := types.NewTransactionServerClient(conn)

	// create time frame
	hourAgo, _ := ptypes.TimestampProto(time.Now().Add(-10 * time.Hour))
	now, _ := ptypes.TimestampProto(time.Now())

	//create crypto request
	cryptoRequest := &types.CryptoTransactionRequest{
		Type:  ASSET,
		From:  hourAgo,
		To:    now,
		Count: AMOUNT,
	}

	//call TopTransactions metohd on server
	sub, err := client.TopTransactions(context.Background(), cryptoRequest)
	if err != nil {
		panic(err)
	}

	for {
		transaction, err := sub.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		fmt.Println(transaction)
	}
}
