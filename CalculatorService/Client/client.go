package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"log"

	"github.com/vidurkataria/CalculatorService/calcipb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("client started")
	cc, err := grpc.Dial("localhost:9342", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	client := calcipb.NewCalculatorServiceClient(cc)

	// Sum of 2 Numbers Unary API function
	// Sum(client)

	// Sum of 2 Numbers Unary API function
	// Prime(client)

	// Sum of 2 Numbers Unary API function
	// ComputeAvg(client)

	// Sum of 2 Numbers Unary API function
	MaxNum(client)

}

func Sum(c calcipb.CalculatorServiceClient) {
	fmt.Println("Starting a Sum GRPC")

	req := calcipb.SumRequest{
		Num1: 20,
		Num2: 30,
	}
	res, err := c.Sum(context.Background(), &req)

	if err != nil {
		log.Fatalf("error while calling greet grpc unary call: %v", err)
	}

	log.Printf("Response from Greet Unary Call : %v", res.Result)
}

func Prime(c calcipb.CalculatorServiceClient) {
	fmt.Println("Starting a Prime GRPC")
	req := calcipb.PrimeRequest{
		Num1: 20,
	}
	respstream, err := c.Prime(context.Background(), &req)
	if err != nil {
		log.Fatalf("error occured while performing client-side streaming : %v", err)
	}

	for {
		msg, err := respstream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while receving server stream : %v", err)
		}

		fmt.Println("Response From GreetManyTimes Server : ", msg.Result)
	}

}

func ComputeAvg(c calcipb.CalculatorServiceClient) {
	fmt.Println("Starting Client Side Streaming over GRPC for Calculating Average ....")

	stream, err := c.ComputeAvg(context.Background())
	if err != nil {
		log.Fatalf("error occured while performing client-side streaming : %v", err)
	}

	requests := []*calcipb.CompAvgRequest{
		&calcipb.CompAvgRequest{
			Num1: 2,
		},
		&calcipb.CompAvgRequest{
			Num1: 5,
		},
		&calcipb.CompAvgRequest{
			Num1: 8,
		},
		&calcipb.CompAvgRequest{
			Num1: 11,
		},
		&calcipb.CompAvgRequest{
			Num1: 14,
		},
		&calcipb.CompAvgRequest{
			Num1: 17,
		},
	}

	for _, req := range requests {
		fmt.Println("\nSending Request.... : ", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from server : %v", err)
	}
	fmt.Println("\n****Response From Server : ", resp.GetResult())
}

func MaxNum(c calcipb.CalculatorServiceClient) {
	fmt.Println("Starting Bi-directional stream by calling GreetEveryone over GRPC......")

	requests := []*calcipb.MaxNumRequest{
		&calcipb.MaxNumRequest{
			Num1: 1,
		},
		&calcipb.MaxNumRequest{
			Num1: 3,
		},
		&calcipb.MaxNumRequest{
			Num1: 7,
		},
		&calcipb.MaxNumRequest{
			Num1: 9,
		},
		&calcipb.MaxNumRequest{
			Num1: 2,
		},
		&calcipb.MaxNumRequest{
			Num1: 5,
		},
		&calcipb.MaxNumRequest{
			Num1: 22,
		},
		&calcipb.MaxNumRequest{
			Num1: 15,
		},
		&calcipb.MaxNumRequest{
			Num1: 21,
		},
		&calcipb.MaxNumRequest{
			Num1: 19,
		},
	}

	stream, err := c.MaxNum(context.Background())
	if err != nil {
		log.Fatalf("error occured while performing client side streaming : %v", err)
	}

	//wait channel to block receiver
	waitchan := make(chan struct{})

	go func(requests []*calcipb.MaxNumRequest) {
		for _, req := range requests {

			fmt.Println("\nSending Request..... : ", req)
			err := stream.Send(req)
			if err != nil {
				log.Fatalf("error while sending request to GreetEveryone service : %v", err)
			}
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}(requests)

	go func() {
		for {

			resp, err := stream.Recv()
			if err == io.EOF {
				close(waitchan)
				return
			}

			if err != nil {
				log.Fatalf("error receiving response from server : %v", err)
			}

			fmt.Printf("\nResponse From Server : %v", resp.GetResult())
		}
	}()

	//block until everything is finished
	<-waitchan
}
