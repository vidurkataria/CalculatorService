package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"time"

	"log"
	// "net"

	"github.com/vidurkataria/CalculatorService/calcipb"
	"google.golang.org/grpc"
)

type calciServer struct {
	calcipb.UnimplementedCalculatorServiceServer
}

var primenums [100000]bool

func isPrime() [100000]bool {
	primenums[0] = true
	primenums[1] = true
	for i := 2; i*i < 100000; i++ {
		if !primenums[i] {
			for j := i * i; j < 100000; j += i {
				primenums[j] = true
			}
		}
	}
	return primenums
}

func main() {
	primenums = isPrime()
	fmt.Println("Server Started.....!")
	listen, err := net.Listen("tcp", "0.0.0.0:9342")
	if err != nil {
		log.Fatal("Failed to Listen: %v", err)
	}

	server := grpc.NewServer()

	calcipb.RegisterCalculatorServiceServer(server, &calciServer{})

	if err = server.Serve(listen); err != nil {
		log.Fatal("failed to serve : %v", err)
	}
}

func (*calciServer) Sum(ctx context.Context, req *calcipb.SumRequest) (*calcipb.SumResponse, error) {
	fmt.Println("Invoked function to return sum of the given two numbers")

	num1 := req.GetNum1()
	num2 := req.GetNum2()

	result := num1 + num2
	resp := &calcipb.SumResponse{
		Result: result,
	}
	return resp, nil
}

func (*calciServer) Prime(req *calcipb.PrimeRequest, res calcipb.CalculatorService_PrimeServer) error {
	fmt.Println("Invoked function to return stream of prime number smaller then the requested number")

	num := req.GetNum1()

	for i := 2; i <= int(num); i++ {
		if !primenums[i] {
			result := calcipb.PrimeResponse{
				Result: int64(i),
			}
			time.Sleep(100 * time.Millisecond)
			res.Send(&result)
		}
	}
	return nil
}

func (*calciServer) ComputeAvg(stream calcipb.CalculatorService_ComputeAvgServer) error {
	fmt.Println("Invoked ComputeAvg function to compute average of the stream of numbers")
	totalSum := 0
	totalNums := 0

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			//we have finished reading client stream
			if totalNums != 0 {
				return stream.SendAndClose(&calcipb.CompAvgResponse{
					Result: float32(totalSum / totalNums),
				})
			}
			return stream.SendAndClose(&calcipb.CompAvgResponse{
				Result: 0,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream : %v", err)
		}
		totalNums++

		totalSum += int(msg.GetNum1())
	}
}

func (*calciServer) MaxNum(stream calcipb.CalculatorService_MaxNumServer) error {
	fmt.Println("GreetEveryone Function is invoked to demonstrate Bi-directional streaming")
	LargestNum := 0
	for {

		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("error while receiving data from GreetEveryone client : %v", err)
			return err
		}

		num := req.GetNum1()
		if LargestNum < int(num) {
			LargestNum = int(num)
			sendErr := stream.Send(&calcipb.MaxNumResponse{
				Result: int64(LargestNum),
			})

			if sendErr != nil {
				log.Fatalf("error while sending response to GreetEveryone Client : %v", err)
				return err
			}
		}
	}
}
