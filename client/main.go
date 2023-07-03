package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"testing/cmd/middleware"

	"testing/api"
	pbappointment "testing/pb/appointment"
	pbdoctor "testing/pb/doctor"
	pbhospital "testing/pb/hospital"
	pbschedule "testing/pb/schedule"
	pbuser "testing/pb/user"
)

func main() {
	// Initialize gRPC client
	grpcAddr := ":50052" // address of the gRPC server
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial gRPC server: %v", err)
	}
	defer conn.Close()

	grpcClientUser := pbuser.NewUserDatakuClient(conn)
	grpcClientDoctor := pbdoctor.NewDoctorServiceClient(conn)
	grpcClientAppointment := pbappointment.NewAppointmentServiceClient(conn)
	grpcClientSchedule := pbschedule.NewScheduleAvailabilityDoctorServiceClient(conn)
	grpcClientHospital := pbhospital.NewHospitalServiceClient(conn)

	// Initialize Gin web server
	router := gin.Default()

	// Initialize handlers
	userHandler := api.NewUserHandler(grpcClientUser)
	doctorHandler := api.NewDoctorHandler(grpcClientDoctor)
	appointmentHandler := api.NewAppointmentHandler(grpcClientAppointment)
	scheduleHandler := api.NewScheduleHandler(grpcClientSchedule)
	hospitalHandler := api.NewHospitalHandler(grpcClientHospital)

	// fmt.Println("dssd")
	// Setup routes
	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		auth.POST("/login", userHandler.AuthUser)

		user := api.Group("/users")
		user.Use(middleware.AuthMiddleware(grpcClientUser))
		user.GET("/:id", userHandler.GetUserData)
		user.GET("/", userHandler.GetAllUserData)
		user.POST("/", userHandler.AddUserData)
		user.PUT("/:id", userHandler.UpdateUserData)
		// api.DELETE("/users/:id", userHandler.Dele)

		doctor := api.Group("/doctors")
		doctor.Use(middleware.AuthMiddleware(grpcClientUser))
		doctor.GET("/:id", doctorHandler.GetDoctorById)
		doctor.GET("/", doctorHandler.GetAllDoctors)
		doctor.POST("/", doctorHandler.CreateDoctor)
		doctor.PUT("/:id", doctorHandler.UpdateDoctor)
		doctor.DELETE("/:id", doctorHandler.DeleteDoctor)
		doctor.GET("/hospital/:id", doctorHandler.GetDoctorByHospitalId)

		schedule := api.Group("/schedules")
		schedule.Use(middleware.AuthMiddleware(grpcClientUser))
		schedule.GET("/doctorid/:id", scheduleHandler.GetAllScheduleByDoctorId)
		schedule.GET("/day/:day", scheduleHandler.GetScheduleByDay)
		schedule.GET("/", scheduleHandler.GetAllSchedule)
		schedule.POST("/", scheduleHandler.CreateSchedule)
		schedule.PUT("/:id", scheduleHandler.UpdateSchedule)
		schedule.DELETE("/:id", scheduleHandler.DeleteSchedule)

		appointment := api.Group("/appointments")
		appointment.Use(middleware.AuthMiddleware(grpcClientUser))
		appointment.GET("/:id", appointmentHandler.GetAppointmentById)
		appointment.GET("/", appointmentHandler.GetAllAppointment)
		appointment.POST("/", appointmentHandler.CreateAppointment)
		appointment.POST("/emailappointment", appointmentHandler.EmailAppointment)

		hospital := api.Group("/hospitals")
		hospital.Use(middleware.AuthMiddleware(grpcClientUser))
		hospital.GET("/", hospitalHandler.GetHospital)
		hospital.POST("/", hospitalHandler.CreateHospital)
		hospital.POST("/patient", hospitalHandler.AddPatient)
		hospital.GET("/patient/:id", hospitalHandler.GetPatientById)
	}

	// Start server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start web server: %v", err)
		}
	}()

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to gracefully shutdown web server: %v", err)
	}

	log.Println("Server shutdown complete.")
}

// Note: Make sure to build the protobuf files by running:
// protoc --proto_path=$GOPATH/src:. --grpc-gateway_out=logtostderr=true:. helloworld.proto
// protoc --proto_path=/users/hacktiv8/documents/go/src/grpcgateway --go-grpc_out=/users/hacktiv8/documents/go/src/grpcgateway --grpc-gateway_out=logtostderr=true:/users/hacktiv8/documents/go/src/grpcgateway /users/hacktiv8/documents/go/src/grpcgateway/pb/helloworld.proto
