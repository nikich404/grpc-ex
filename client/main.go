package main

import (
	"context"
	"log"
	"time"

	pb "grpcEx/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("не удалось подключиться: %v", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorServiceClient(conn)

	examples := []struct {
		a, b float64
		op   pb.Operation
	}{
		{10, 5, pb.Operation_ADD},
		{10, 5, pb.Operation_SUBTRACT},
		{10, 5, pb.Operation_MULTIPLY},
		{10, 5, pb.Operation_DIVIDE},
		{10, 0, pb.Operation_DIVIDE},
	}

	for _, ex := range examples {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		resp, err := client.Calculate(ctx, &pb.CalculateRequest{
			FirstNumber:  ex.a,
			SecondNumber: ex.b,
			Operation:    ex.op,
		})
		cancel()

		if err != nil {
			log.Printf("%v %v %v -> ошибка: %v", ex.a, ex.op, ex.b, err)
			continue
		}
		log.Printf("%v %v %v = %v", ex.a, ex.op, ex.b, resp.GetResult())
	}
}
