package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"testing/cmd/config"
	"testing/cmd/services"
	pbappointment "testing/pb/appointment"
	pbdoctor "testing/pb/doctor"
	pbhospital "testing/pb/hospital"
	pbschedule "testing/pb/schedule"
	pbuser "testing/pb/user"
)

func main() {
	flag.Parse()

	db := config.ConnectDB()

	// Set up gRPC server
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userService := services.NewUserService(db)
	doctorService := services.NewDokterService(db)
	appointmentService := services.NewAppointmentService(db)
	scheduleService := services.NewScheduleService(db)
	hospitalService := services.NewHospitalService(db)
	pbuser.RegisterUserDatakuServer(grpcServer, userService)
	pbdoctor.RegisterDoctorServiceServer(grpcServer, doctorService)
	pbappointment.RegisterAppointmentServiceServer(grpcServer, appointmentService)
	pbschedule.RegisterScheduleAvailabilityDoctorServiceServer(grpcServer, scheduleService)
	pbhospital.RegisterHospitalServiceServer(grpcServer, hospitalService)

	// Get the address the listener is listening on
	addr := lis.Addr().String()
	fmt.Printf("Starting gRPC server on %s\n", addr)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
