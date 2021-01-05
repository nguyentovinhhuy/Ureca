package config_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"google.golang.org/grpc"

	pbbatch "github.com/nexledger/accelerator/protos"
)

const (
	channelId     = "mychannel"
	chaincodeName = "sample"
	numOfPings    = 50
	address       = "127.0.0.1:5050"
)

func TestAccelerator(t *testing.T) {
	set(t)
}

func set(t *testing.T) {
	client := pbbatch.NewAcceleratorServiceClient(connect(t))
	notifier_channels := make([]chan string, numOfPings)
	for i := 0; i < numOfPings; i++ {
		notifier := make(chan string)
		notifier_channels[i] = notifier
		go func(i int, notifier chan string) {
			req := &pbbatch.TxRequest{
				ChannelId:     channelId,
				ChaincodeName: chaincodeName,
				Fcn:           "set",
				Args:          [][]byte{[]byte(strconv.Itoa(i)), []byte("value of " + strconv.Itoa(i))},
			}
			resp, err := client.Execute(context.Background(), req)
			if err != nil {
				notifier <- "Failed to execute" + err.Error()
			} else {
				notifier <- strconv.Itoa(i) + ":" + resp.TxId
			}
		}(i, notifier)
	}

	for i := 0; i < numOfPings; i++ {
		fmt.Println(<-notifier_channels[i])
	}
}

func get(t *testing.T) {
	client := pbbatch.NewAcceleratorServiceClient(connect(t))
	notifier_channels := make([]chan string, numOfPings)
	for i := 0; i < numOfPings; i++ {
		notifier := make(chan string)
		notifier_channels[i] = notifier
		go func(i int, notifier chan string) {
			req := &pbbatch.TxRequest{
				ChannelId:     channelId,
				ChaincodeName: chaincodeName,
				Fcn:           "get",
				Args:          [][]byte{[]byte(strconv.Itoa(i))},
			}
			resp, err := client.Query(context.Background(), req)
			if err != nil {
				notifier <- "Failed to query" + err.Error()
			} else {
				notifier <- strconv.Itoa(i) + ":" + string(resp.Payload)
			}
		}(i, notifier)
	}

	for i := 0; i < numOfPings; i++ {
		fmt.Println(<-notifier_channels[i])
	}
}

func connect(t *testing.T) *grpc.ClientConn {
	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Failed to connect server.", err)
	}
	return cc
}
