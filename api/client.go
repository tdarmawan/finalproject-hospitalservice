package api

import (
	"context"

	"google.golang.org/grpc"

	pbappointment "testing/pb/appointment"
	pbdoctor "testing/pb/doctor"
	pbhospital "testing/pb/hospital"
	pbschedule "testing/pb/schedule"
	pbuser "testing/pb/user" // import the protobuf package containing the service definition
)

type Client struct {
	conn        *grpc.ClientConn
	client      pbuser.UserDatakuClient // the gRPC client generated from the protobuf service definition
	doctor      pbdoctor.DoctorServiceClient
	appointment pbappointment.AppointmentServiceClient
	schedule    pbschedule.ScheduleAvailabilityDoctorServiceClient
	hospital    pbhospital.HospitalServiceClient
}

func NewClient(addr string) (*Client, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pbuser.NewUserDatakuClient(conn)
	doctor := pbdoctor.NewDoctorServiceClient(conn)
	hospital := pbhospital.NewHospitalServiceClient(conn)

	return &Client{
		conn:     conn,
		client:   client,
		doctor:   doctor,
		hospital: hospital,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) GetUserData(ctx context.Context, req *pbuser.GetUserRequest) (*pbuser.User, error) {
	return c.client.GetUserData(ctx, req)
}

func (c *Client) GetAllData(ctx context.Context, req *pbuser.GetAllUserRequest) (*pbuser.GetAllUserResponse, error) {
	return c.client.GetAllData(ctx, req)
}

func (c *Client) AddUser(ctx context.Context, req *pbuser.AddUserRequest) (*pbuser.AddUserResponse, error) {
	return c.client.AddUser(ctx, req)
}

func (c *Client) UpdateUser(ctx context.Context, req *pbuser.UpdateUserRequest) (*pbuser.UpdateUserResponse, error) {
	return c.client.UpdateUser(ctx, req)
}

func (c *Client) AuthUser(ctx context.Context, req *pbuser.AuthRequest) (*pbuser.AuthResponse, error) {
	return c.client.AuthUser(ctx, req)
}

func (c *Client) GenerateToken(ctx context.Context, req *pbuser.GenerateTokenRequest) (*pbuser.GenerateTokenResponse, error) {
	return c.client.GenerateToken(ctx, req)
}

func (c *Client) GetDoctorById(ctx context.Context, req *pbdoctor.DoctorId) (*pbdoctor.DoctorResponse, error) {
	return c.doctor.GetDoctorById(ctx, req)
}

func (c *Client) GetAllDoctors(ctx context.Context, req *pbdoctor.Empty) (*pbdoctor.DoctorList, error) {
	return c.doctor.GetAllDoctors(ctx, req)
}

func (c *Client) CreateDoctor(ctx context.Context, req *pbdoctor.DoctorRequest) (*pbdoctor.DoctorResponse, error) {
	return c.doctor.CreateDoctor(ctx, req)
}

func (c *Client) UpdateDoctor(ctx context.Context, req *pbdoctor.DoctorRequest) (*pbdoctor.DoctorResponse, error) {
	return c.doctor.UpdateDoctor(ctx, req)
}

func (c *Client) DeleteDoctor(ctx context.Context, req *pbdoctor.DoctorId) (*pbdoctor.DoctorResponse, error) {
	return c.doctor.DeleteDoctor(ctx, req)
}

func (c *Client) CreateAppointment(ctx context.Context, req *pbappointment.CreateAppointmentRequest) (*pbappointment.CreateAppointmentResponse, error) {
	return c.appointment.CreateAppointment(ctx, req)
}

func (c *Client) GetAllAppointment(ctx context.Context, req *pbappointment.Empty) (*pbappointment.ReadAppointmentResponse, error) {
	return c.appointment.ReadAppointmentAll(ctx, req)
}

func (c *Client) GetAppointmentById(ctx context.Context, req *pbappointment.AppointmentId) (*pbappointment.ReadAppointmentResponse, error) {
	return c.appointment.ReadAppointmentById(ctx, req)
}

func (c *Client) EmailAppointment(ctx context.Context, req *pbappointment.EmailAppointmentRequest) (*pbappointment.EmailAppointmentResponse, error) {
	return c.appointment.EmailAppointment(ctx, req)
}

func (c *Client) CreateSchedule(ctx context.Context, req *pbschedule.ScheduleAvailabilityDoctorRequest) (*pbschedule.ScheduleAvailabilityDoctorResponse, error) {
	return c.schedule.Create(ctx, req)
}

func (c *Client) GetAllSchedule(ctx context.Context, req *pbschedule.ScheduleAvailabilityDoctorRequest) (*pbschedule.ScheduleAvailabilityDoctorResponse, error) {
	return c.schedule.Read(ctx, req)
}

func (c *Client) UpdateSchedule(ctx context.Context, req *pbschedule.ScheduleAvailabilityDoctorRequest) (*pbschedule.ScheduleAvailabilityDoctorResponse, error) {
	return c.schedule.Update(ctx, req)
}

func (c *Client) DeleteSchedule(ctx context.Context, req *pbschedule.ScheduleAvailabilityDoctorRequest) (*pbschedule.ScheduleAvailabilityDoctorResponse, error) {
	return c.schedule.Delete(ctx, req)
}

func (c *Client) GetAllScheduleByDoctorId(ctx context.Context, req *pbschedule.ScheduleAvailabilityDoctorRequestByDoctorId) (*pbschedule.ScheduleAvailabilityDoctorResponse, error) {
	return c.schedule.ReadScheduleByDoctorId(ctx, req)
}

func (c *Client) GetAllScheduleByDay(ctx context.Context, req *pbschedule.ScheduleAvailabilityDoctorRequestByDay) (*pbschedule.ScheduleAvailabilityDoctorResponse, error) {
	return c.schedule.ReadScheduleByDay(ctx, req)
}

func (c *Client) CreateHospital(ctx context.Context, req *pbhospital.Hospital) (*pbhospital.HospitalResponse, error) {
	return c.hospital.CreateHospital(ctx, req)
}

func (c *Client) GetHospital(ctx context.Context, req *pbhospital.HospitalRequest) (*pbhospital.HospitalResponse, error) {
	return c.hospital.GetHospital(ctx, req)
}

func (c *Client) GetPatients(ctx context.Context, req *pbhospital.PatientRequest) (*pbhospital.PatientResponse, error) {
	return c.hospital.GetPatients(ctx, req)
}

func (c *Client) AddPatient(ctx context.Context, req *pbhospital.Patient) (*pbhospital.PatientResponse, error) {
	return c.hospital.AddPatient(ctx, req)
}

func (c *Client) GetPatientById(ctx context.Context, req *pbhospital.PatientId) (*pbhospital.PatientResponse, error) {
	return c.hospital.GetPatientById(ctx, req)
}
