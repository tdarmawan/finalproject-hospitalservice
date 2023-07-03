package services

import (
	"context"
	_ "net"
	pb "testing/pb/hospital"
	"time"

	"gorm.io/gorm"
)

type hospital struct {
	pb.UnimplementedHospitalServiceServer
	db *gorm.DB
}

// make hospital struct with gorm tag and json tag

type Hospital struct {
	ID           uint32    `gorm:"primaryKey"`
	HospitalCode string    `gorm:"unique"`
	Name         string    `gorm:"not null"`
	Address      string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

// make conjuction table for hospital have patient
type HospitalPatient struct {
	ID           uint32 `gorm:"primaryKey"`
	HospitalCode string `gorm:"primaryKey"`
	PatientId    uint32 `gorm:"primaryKey"`
}

//make patient struct with gorm tah and json tag

type Patient struct {
	ID         uint32    `gorm:"primaryKey"`
	IdentityId string    `gorm:"unique"`
	Name       string    `gorm:"not null"`
	Age        int32     `gorm:"not null"`
	Disease    string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

func NewHospitalService(db *gorm.DB) *hospital {
	server := &hospital{
		db: db,
	}
	return server
}

// make method CRUD hospital

func (s *hospital) CreateHospital(ctx context.Context, req *pb.Hospital) (*pb.HospitalResponse, error) {
	// make function to create hospital

	// make variable to store data from request
	hospital := Hospital{
		HospitalCode: req.GetHospitalcode(),
		Name:         req.GetName(),
		Address:      req.GetAddress(),
	}

	// create hospital
	s.db.Create(&hospital)

	// return response
	return &pb.HospitalResponse{
		Data: []*pb.Hospital{
			{
				Hospitalcode: hospital.HospitalCode,
				Name:         hospital.Name,
				Address:      hospital.Address,
			},
		},
		Message: "Success",
		Status:  200,
	}, nil

}

func (s *hospital) GetHospital(ctx context.Context, req *pb.HospitalRequest) (*pb.HospitalResponse, error) {
	// make function to get hospital

	// make variable to store data from request
	hospital := Hospital{
		HospitalCode: req.GetHospitalcode(),
	}

	// get hospital
	s.db.Find(&hospital)

	// return response
	return &pb.HospitalResponse{
		Data: []*pb.Hospital{
			{
				Hospitalcode: hospital.HospitalCode,
				Name:         hospital.Name,
				Address:      hospital.Address,
			},
		},
		Message: "Success",
		Status:  200,
	}, nil

}

func (s *hospital) AddPatient(ctx context.Context, req *pb.Patient) (*pb.PatientResponse, error) {
	// make function to add patient

	// make variable to store data from request
	patient := Patient{
		IdentityId: req.GetIdentityId(),
		Name:       req.GetName(),
		Age:        req.GetAge(),
		Disease:    req.GetDisease(),
	}

	// create patient
	s.db.Create(&patient)

	// add patient to hospital in conjuction table
	s.db.Create(&HospitalPatient{
		HospitalCode: req.GetHospitalcode(),
		PatientId:    patient.ID,
	})

	// return response
	return &pb.PatientResponse{
		Data: []*pb.Patient{
			{
				IdentityId: patient.IdentityId,
				Name:       patient.Name,
				Age:        patient.Age,
				Disease:    patient.Disease,
			},
		},
		Message: "Success",
		Status:  200,
	}, nil

}

// make method get patient by id
func (s *hospital) GetPatientById(ctx context.Context, req *pb.PatientId) (*pb.PatientResponse, error) {
	// make variable to store data from request
	patient := Patient{
		IdentityId: req.GetIdentityId(),
	}

	// get patient
	s.db.Where("identity_id = ?", patient.IdentityId).Find(&patient)

	// return error handling if patient not found
	if patient.IdentityId == "" {
		return &pb.PatientResponse{
			Data: []*pb.Patient{
				{
					IdentityId: patient.IdentityId,
					Name:       patient.Name,
					Age:        patient.Age,
					Disease:    patient.Disease,
				},
			},
			Message: "Patient not found",
			Status:  404,
		}, nil
	}

	// return response
	return &pb.PatientResponse{
		Data: []*pb.Patient{
			{
				IdentityId: patient.IdentityId,
				Name:       patient.Name,
				Age:        patient.Age,
				Disease:    patient.Disease,
			},
		},
		Message: "Success",
		Status:  200,
	}, nil
}

func (s *hospital) GetPatients(ctx context.Context, req *pb.PatientRequest) (*pb.PatientResponse, error) {
	// make function to get patients

	// make variable to store data from request
	patient := Patient{
		IdentityId: req.GetIdentityId(),
	}

	// get patients
	s.db.Find(&patient)

	// return response
	return &pb.PatientResponse{
		Data: []*pb.Patient{
			{
				IdentityId: patient.IdentityId,
				Name:       patient.Name,
				Age:        patient.Age,
				Disease:    patient.Disease,
			},
		},
		Message: "Success",
		Status:  200,
	}, nil

}
