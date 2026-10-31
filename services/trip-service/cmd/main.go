package main

import (
	"context"
	"log"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"
)

func main() {
	ctx := context.Background()
	inmem := repository.NewInMemRepository()
	svc := service.NewService(inmem)
	fare := &domain.RideFareModel{
		UserId: "abc123",
	}
	trip, err := svc.CreateTrip(ctx, fare)
	if err != nil {
		log.Println(err)
	}
	log.Print(trip)

	for {
		time.Sleep(time.Second)
	}
}
