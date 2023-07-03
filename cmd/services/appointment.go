package services

import (
	"context"
	"fmt"
	"log"
	_ "net"
	"net/mail"
	"net/smtp"
	pb "testing/pb/appointment"
	"time"

	"gorm.io/gorm"
)

type appointment struct {
	pb.UnimplementedAppointmentServiceServer
	db *gorm.DB
}

// make appointment struct with gorm tag and json tag

type Appointment struct {
	ID              uint32    `gorm:"primaryKey"`
	AppointmentCode string    `gorm:"unique"`
	DoctorId        int32     `gorm:"not null"`
	PatientId       string    `gorm:"not null"`
	ScheduleId      int32     `gorm:"not null"`
	message         string    `gorm:"not null"`
	status          string    `gorm:"not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}

func NewAppointmentService(db *gorm.DB) *appointment {
	server := &appointment{
		db: db,
	}
	return server
}

// make method CRUD appointment

func (s *appointment) CreateAppointment(ctx context.Context, req *pb.CreateAppointmentRequest) (*pb.CreateAppointmentResponse, error) {
	// make function to create appointment

	// make variable to store data from request
	appointment := Appointment{
		AppointmentCode: req.GetAppointmentCode(),
		DoctorId:        req.GetDoctorId(),
		PatientId:       req.GetPatientId(),
		ScheduleId:      req.GetScheduleId(),
		message:         req.GetMessage(),
		status:          req.GetStatus(),
	}

	// create appointment
	s.db.Create(&appointment)

	// return response
	return &pb.CreateAppointmentResponse{
		Data: []*pb.CreateAppointmentRequest{
			{
				AppointmentCode: appointment.AppointmentCode,
				DoctorId:        appointment.DoctorId,
				PatientId:       appointment.PatientId,
				ScheduleId:      appointment.ScheduleId,
				Message:         appointment.message,
				Status:          appointment.status,
			},
		},
		Message: "Success",
		Status:  200,
	}, nil

}

func (s *appointment) ReadAppointmentAll(ctx context.Context, req *pb.Empty) (*pb.ReadAppointmentResponse, error) {
	// make function to read all appointment

	// make variable to store data from request
	appointments := []Appointment{}

	// read all appointment
	s.db.Find(&appointments)

	// create a slice of pb.CreateAppointmentRequest to store the response data
	data := make([]*pb.CreateAppointmentRequest, len(appointments))

	// loop through the appointments slice and populate the data slice
	for i, appointment := range appointments {
		data[i] = &pb.CreateAppointmentRequest{
			Id:              appointment.ID,
			AppointmentCode: appointment.AppointmentCode,
			DoctorId:        appointment.DoctorId,
			PatientId:       appointment.PatientId,
			ScheduleId:      appointment.ScheduleId,
			Message:         appointment.message,
			Status:          appointment.status,
		}
	}

	// return response
	return &pb.ReadAppointmentResponse{
		Data:    data,
		Message: "Success",
		Status:  200,
	}, nil
}

func (s *appointment) GetAppointmentById(ctx context.Context, req *pb.AppointmentId) (*pb.ReadAppointmentResponse, error) {
	// make function to get appointment by id

	// make variable to store data from request
	appointment := Appointment{}

	// get appointment by id
	s.db.First(&appointment, req)

	// return response
	return &pb.ReadAppointmentResponse{
		Data: []*pb.CreateAppointmentRequest{
			{
				Id:              appointment.ID,
				AppointmentCode: appointment.AppointmentCode,
				DoctorId:        appointment.DoctorId,
				PatientId:       appointment.PatientId,
				ScheduleId:      appointment.ScheduleId,
				Message:         appointment.message,
				Status:          appointment.status,
			},
		},
		Message: "Success",
		Status:  200,
	}, nil
}

func (s *appointment) EmailAppointment(ctx context.Context, req *pb.EmailAppointmentRequest) (*pb.EmailAppointmentResponse, error) {
	// make function to email appointment

	// make variable to store data from request
	appointment := Appointment{}

	// make proccess for sending email to patient about the appointment
	// send email to patient using smtp email
	// make func sendEmail return the json message

	go sendEmail(appointment.PatientId, appointment.message) // make it as goroutine

	// return response
	return &pb.EmailAppointmentResponse{
		Data: []*pb.CreateAppointmentRequest{
			{
				Id:              appointment.ID,
				AppointmentCode: appointment.AppointmentCode,
				DoctorId:        appointment.DoctorId,
				PatientId:       appointment.PatientId,
				ScheduleId:      appointment.ScheduleId,
				Message:         appointment.message,
				Status:          appointment.status,
			},
		},
		Message: "Success",
		Status:  200,
	}, nil
}

func sendEmail(email string, message string) {
	// make function to send email
	// make proccess to send email to patient
	// send email to patient using smtp email

	from := mail.Address{Name: "Sender Name", Address: "tdarmawan@hacktiv8.com"}
	to := mail.Address{Name: "Receiver Name", Address: email}
	subj := "Appointment Confirmation"
	body := message

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message = ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	auth := smtp.PlainAuth("", "tdarmawan@hacktiv8.com", "ceorfiqfumnxuayc", "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", auth, from.Address, []string{to.Address}, []byte(message))
	if err != nil {
		fmt.Println(err)
		log.Println("failed")
		return
	}

	log.Println("success")
}
