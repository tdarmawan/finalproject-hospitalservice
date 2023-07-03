package services

import (
	"context"
	pb "testing/pb/doctor"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type dokter struct {
	pb.UnimplementedDoctorServiceServer
	db *gorm.DB
}

// make doctor struct with gorm tag and json tag

type Dokter struct {
	ID           uint      `gorm:"primaryKey"`
	DoctorId     int32     `gorm:"unique"`
	Email        string    `gorm:"unique"`
	Password     string    `gorm:"not null"`
	Name         string    `gorm:"not null"`
	Speciality   string    `gorm:"not null"`
	hospital_id  int32     `gorm:"not null"`
	phone_number string    `gorm:"embedded"`
	address      string    `gorm:"not null"`
	city         string    `gorm:"not null"`
	state        string    `gorm:"not null"`
	country      string    `gorm:"not null"`
	profil_pic   string    `gorm:"not null"`
	status       string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

func NewDokterService(db *gorm.DB) *dokter {
	server := &dokter{
		db: db,
	}
	return server
}

// make method CRUD doctor

func (s *dokter) CreateDokter(ctx context.Context, req *pb.DoctorRequest) (*pb.DoctorResponse, error) {
	// make function to create doctor

	// make variable to store data from request
	dokter := Dokter{
		DoctorId:     req.GetDoctorId(),
		Email:        req.GetEmail(),
		Password:     req.GetPassword(),
		Name:         req.GetName(),
		Speciality:   req.GetSpeciality(),
		hospital_id:  req.GetHospitalId(),
		phone_number: req.GetPhone(),
		address:      req.GetAddress(),
		city:         req.GetCity(),
		state:        req.GetState(),
		country:      req.GetCountry(),
		profil_pic:   req.GetProfilePic(),
		status:       req.GetStatus(),
	}

	// make variable to store data from database
	var dokterDB Dokter

	// make query to database
	result := s.db.Where("email = ?", dokter.Email).First(&dokterDB)

	// check if data from database is exist
	if result.RowsAffected > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "email already exists")
	}

	// make query to insert data
	if err := s.db.Create(&dokter).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	// return response
	return &pb.DoctorResponse{
		DoctorId:   dokter.DoctorId,
		Email:      dokter.Email,
		Name:       dokter.Name,
		Speciality: dokter.Speciality,
		HospitalId: int32(dokter.hospital_id),
		Phone:      dokter.phone_number,
		Address:    dokter.address,
		City:       dokter.city,
		State:      dokter.state,
		Country:    dokter.country,
		ProfilePic: dokter.profil_pic,
		Status:     dokter.status,
	}, nil
}

func (s *dokter) ReadDokter(ctx context.Context, req *pb.DoctorRequest) (*pb.DoctorResponse, error) {
	// make function to read doctor

	// make variable to store data from request
	dokter := Dokter{
		DoctorId: req.GetDoctorId(),
	}

	// make variable to store data from database
	var dokterDB Dokter

	// make query to database
	result := s.db.Where("doctor_id = ?", dokter.DoctorId).First(&dokterDB)

	// check if data from database is exist
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "doctor not found")
	}

	// return response
	return &pb.DoctorResponse{
		DoctorId:   dokterDB.DoctorId,
		Email:      dokterDB.Email,
		Name:       dokterDB.Name,
		Speciality: dokterDB.Speciality,
		HospitalId: int32(dokterDB.hospital_id),
		Phone:      dokterDB.phone_number,
		Address:    dokterDB.address,
		City:       dokterDB.city,
		State:      dokterDB.state,
		Country:    dokterDB.country,
		ProfilePic: dokterDB.profil_pic,
		Status:     dokterDB.status,
	}, nil
}

func (s *dokter) UpdateDokter(ctx context.Context, req *pb.DoctorRequest) (*pb.DoctorResponse, error) {
	// make function to update doctor

	// make variable to store data from request
	dokter := Dokter{
		DoctorId:     req.GetDoctorId(),
		Email:        req.GetEmail(),
		Password:     req.GetPassword(),
		Name:         req.GetName(),
		Speciality:   req.GetSpeciality(),
		hospital_id:  req.GetHospitalId(),
		phone_number: req.GetPhone(),
		address:      req.GetAddress(),
		city:         req.GetCity(),
		state:        req.GetState(),
		country:      req.GetCountry(),
		profil_pic:   req.GetProfilePic(),
		status:       req.GetStatus(),
	}

	// make variable to store data from database
	var dokterDB Dokter

	// make query to database
	result := s.db.Where("doctor_id = ?", dokter.DoctorId).First(&dokterDB)

	// check if data from database is exist
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "doctor not found")
	}

	// make query to update data
	if err := s.db.Model(&dokter).Updates(&dokter).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	// return response
	return &pb.DoctorResponse{
		DoctorId:   dokter.DoctorId,
		Email:      dokter.Email,
		Name:       dokter.Name,
		Speciality: dokter.Speciality,
		HospitalId: int32(dokter.hospital_id),
		Phone:      dokter.phone_number,
		Address:    dokter.address,
		City:       dokter.city,
		State:      dokter.state,
		Country:    dokter.country,
		ProfilePic: dokter.profil_pic,
		Status:     dokter.status,
	}, nil
}

func (s *dokter) DeleteDokter(ctx context.Context, req *pb.DoctorRequest) (*pb.DoctorResponse, error) {
	// make function to delete doctor

	// make variable to store data from request
	dokter := Dokter{
		DoctorId: req.GetDoctorId(),
	}

	// make variable to store data from database
	var dokterDB Dokter

	// make query to database
	result := s.db.Where("doctor_id = ?", dokter.DoctorId).First(&dokterDB)

	// check if data from database is exist
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "doctor not found")
	}

	// make query to delete data
	if err := s.db.Delete(&dokter).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	// return response
	return &pb.DoctorResponse{
		DoctorId:   dokterDB.DoctorId,
		Email:      dokterDB.Email,
		Name:       dokterDB.Name,
		Speciality: dokterDB.Speciality,
		HospitalId: int32(dokterDB.hospital_id),
		Phone:      dokterDB.phone_number,
		Address:    dokterDB.address,
		City:       dokterDB.city,
		State:      dokterDB.state,
		Country:    dokterDB.country,
		ProfilePic: dokterDB.profil_pic,
		Status:     dokterDB.status,
	}, nil
}

func (s *dokter) GetDoctorById(ctx context.Context, req *pb.DoctorId) (*pb.DoctorResponse, error) {
	// make function to get doctor by id

	// make variable to store data from request
	dokter := Dokter{
		DoctorId: req.GetDoctorId(),
	}

	// make variable to store data from database
	var dokterDB Dokter

	// make query to database
	result := s.db.Where("doctor_id = ?", dokter.DoctorId).First(&dokterDB)

	// check if data from database is exist
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "doctor not found")
	}

	// return response
	return &pb.DoctorResponse{
		DoctorId:   dokterDB.DoctorId,
		Email:      dokterDB.Email,
		Name:       dokterDB.Name,
		Speciality: dokterDB.Speciality,
		HospitalId: int32(dokterDB.hospital_id),
		Phone:      dokterDB.phone_number,
		Address:    dokterDB.address,
		City:       dokterDB.city,
		State:      dokterDB.state,
		Country:    dokterDB.country,
		ProfilePic: dokterDB.profil_pic,
		Status:     dokterDB.status,
	}, nil
}

func (s *dokter) GetDoctorByHospitalId(ctx context.Context, req *pb.HospitalId) (*pb.DoctorList, error) {

	// make function to get doctor by hospital id

	// make variable to store data from request
	dokter := Dokter{
		hospital_id: req.GetHospitalId(),
	}

	// make variable to store data from database
	var dokterDB Dokter

	// make query to database
	result := s.db.Where("hospital_id = ?", dokter.hospital_id).First(&dokterDB)

	// check if data from database is exist
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "doctor not found")
	}

	// return response
	return &pb.DoctorList{
		Data: []*pb.DoctorResponse{
			{
				DoctorId:   dokterDB.DoctorId,
				Email:      dokterDB.Email,
				Name:       dokterDB.Name,
				Speciality: dokterDB.Speciality,
				HospitalId: int32(dokterDB.hospital_id),
				Phone:      dokterDB.phone_number,
				Address:    dokterDB.address,
				City:       dokterDB.city,
				State:      dokterDB.state,
				Country:    dokterDB.country,
				ProfilePic: dokterDB.profil_pic,
				Status:     dokterDB.status,
			},
		},
		Message: "success",
		Status:  200,
	}, nil
}
