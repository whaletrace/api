package main

import (
	"context"
	"fmt"
	"io"
	"time"

	types ".."

	"google.golang.org/grpc"
)

const Server = "grpc.whaletrace.com:30000" //replace with server connection
const TOKEN = "YOUR_TOKEN"                 //replace with your api token
const ASSET = "BTC"                        //replace with different asset if interested

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

	//create crypto request
	cryptoRequest := &types.CryptoSubscribeRequest{
		Type: ASSET,
	}

	//subscribe to Transactions on server
	sub, err := client.SubscribeTransactions(context.Background(), cryptoRequest)
	if err != nil {
		panic(err)
	}

	//infinite stream
	for {
		transaction, err := sub.Recv()
		if err == io.EOF {
			panic("connection closed!")
		}
		if err != nil {
			panic(err)
		}

		fmt.Println(transaction)
	}
}
