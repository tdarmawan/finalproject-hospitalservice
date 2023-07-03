package services

import (
	"context"
	pb "testing/pb/schedule"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type schedule struct {
	pb.UnimplementedScheduleAvailabilityDoctorServiceServer
	db *gorm.DB
}

// make schedule struct with gorm tag and json tag

type Schedule struct {
	ID         uint      `gorm:"primaryKey"`
	DoctorId   int32     `gorm:"unique"`
	Start_time string    `gorm:"not null"`
	End_time   string    `gorm:"not null"`
	Day        string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

func NewScheduleService(db *gorm.DB) *schedule {
	server := &schedule{
		db: db,
	}
	return server
}

// make method CRUD schedule

func (s *schedule) CreateSchedule(ctx context.Context, req *pb.ScheduleAvailabilityDoctorRequest) (*pb.ScheduleAvailabilityDoctorResponse, error) {
	// make variable to store data from request
	schedule := Schedule{
		DoctorId:   req.GetDoctorId(),
		Start_time: req.GetStartTime(),
		End_time:   req.GetEndTime(),
		Day:        req.GetDay(),
	}

	// make variable to store data from database
	var scheduleDB Schedule

	// make query to check if doctor id and day exist
	query := s.db.Where("doctor_id = ? AND day = ?", schedule.DoctorId, schedule.Day).First(&scheduleDB)

	// check if query error
	if query.Error != nil {
		// check if query error is not found
		if query.Error == gorm.ErrRecordNotFound {
			// make query to create schedule
			query := s.db.Create(&schedule)

			// check if query error
			if query.Error != nil {
				return nil, status.Error(codes.Internal, "Error when create schedule")
			}

			// make response in Data and with data types []*ScheduleAvailabilityDoctor
			res := &pb.ScheduleAvailabilityDoctorResponse{
				Data: []*pb.ScheduleAvailabilityDoctor{
					{
						DoctorId:  schedule.DoctorId,
						StartTime: schedule.Start_time,
						EndTime:   schedule.End_time,
						Day:       schedule.Day,
					},
				},
				Message: "Success create schedule",
				Status:  200,
			}

			return res, nil
		}

		return nil, status.Error(codes.Internal, "Error when check schedule")
	}

	// make response in Data and with data types []*ScheduleAvailabilityDoctor
	res := &pb.ScheduleAvailabilityDoctorResponse{
		Data: []*pb.ScheduleAvailabilityDoctor{
			{
				DoctorId:  scheduleDB.DoctorId,
				StartTime: scheduleDB.Start_time,
				EndTime:   scheduleDB.End_time,
				Day:       scheduleDB.Day,
			},
		},
		Message: "Schedule already exist",
		Status:  200,
	}

	return res, nil
}

func (s *schedule) ReadSchedule(ctx context.Context, req *pb.ScheduleAvailabilityDoctorRequest) (*pb.ScheduleAvailabilityDoctorResponse, error) {
	// make variable to store data from request
	schedule := Schedule{
		DoctorId: req.GetDoctorId(),
		Day:      req.GetDay(),
	}

	// make variable to store data from database
	var scheduleDB Schedule

	// make query to check if doctor id and day exist
	query := s.db.Where("doctor_id = ? AND day = ?", schedule.DoctorId, schedule.Day).First(&scheduleDB)

	// check if query error
	if query.Error != nil {
		// check if query error is not found
		if query.Error == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, "Schedule not found")
		}

		return nil, status.Error(codes.Internal, "Error when check schedule")
	}

	// make response in Data and with data types []*ScheduleAvailabilityDoctor
	res := &pb.ScheduleAvailabilityDoctorResponse{
		Data: []*pb.ScheduleAvailabilityDoctor{
			{
				DoctorId:  scheduleDB.DoctorId,
				StartTime: scheduleDB.Start_time,
				EndTime:   scheduleDB.End_time,
				Day:       scheduleDB.Day,
			},
		},
		Message: "Success read schedule",
		Status:  200,
	}

	return res, nil
}

func (s *schedule) UpdateSchedule(ctx context.Context, req *pb.ScheduleAvailabilityDoctorRequest) (*pb.ScheduleAvailabilityDoctorResponse, error) {
	// make variable to store data from request
	schedule := Schedule{
		DoctorId:   req.GetDoctorId(),
		Start_time: req.GetStartTime(),
		End_time:   req.GetEndTime(),
		Day:        req.GetDay(),
	}

	// make variable to store data from database
	var scheduleDB Schedule

	// make query to check if doctor id and day exist
	query := s.db.Where("doctor_id = ? AND day = ?", schedule.DoctorId, schedule.Day).First(&scheduleDB)

	// check if query error
	if query.Error != nil {
		// check if query error is not found
		if query.Error == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, "Schedule not found")
		}

		return nil, status.Error(codes.Internal, "Error when check schedule")
	}

	// make query to update schedule
	query = s.db.Model(&scheduleDB).Updates(&schedule)

	// check if query error
	if query.Error != nil {
		return nil, status.Error(codes.Internal, "Error when update schedule")
	}

	// make response in Data and with data types []*ScheduleAvailabilityDoctor
	res := &pb.ScheduleAvailabilityDoctorResponse{
		Data: []*pb.ScheduleAvailabilityDoctor{
			{
				DoctorId:  schedule.DoctorId,
				StartTime: schedule.Start_time,
				EndTime:   schedule.End_time,
				Day:       schedule.Day,
			},
		},
		Message: "Success update schedule",
		Status:  200,
	}

	return res, nil
}

func (s *schedule) DeleteSchedule(ctx context.Context, req *pb.ScheduleAvailabilityDoctorRequest) (*pb.ScheduleAvailabilityDoctorResponse, error) {
	// make variable to store data from request
	schedule := Schedule{
		DoctorId: req.GetDoctorId(),
		Day:      req.GetDay(),
	}

	// make variable to store data from database
	var scheduleDB Schedule

	// make query to check if doctor id and day exist
	query := s.db.Where("doctor_id = ? AND day = ?", schedule.DoctorId, schedule.Day).First(&scheduleDB)

	// check if query error
	if query.Error != nil {
		// check if query error is not found
		if query.Error == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, "Schedule not found")
		}

		return nil, status.Error(codes.Internal, "Error when check schedule")
	}

	// make query to delete schedule
	query = s.db.Delete(&scheduleDB)

	// check if query error
	if query.Error != nil {
		return nil, status.Error(codes.Internal, "Error when delete schedule")
	}

	// make response in Data and with data types []*ScheduleAvailabilityDoctor
	res := &pb.ScheduleAvailabilityDoctorResponse{
		Data: []*pb.ScheduleAvailabilityDoctor{
			{
				DoctorId:  schedule.DoctorId,
				StartTime: schedule.Start_time,
				EndTime:   schedule.End_time,
				Day:       schedule.Day,
			},
		},
		Message: "Success delete schedule",
		Status:  200,
	}

	return res, nil
}

func (s *schedule) ReadAllSchedule(ctx context.Context, req *pb.ScheduleAvailabilityDoctorRequest) (*pb.ScheduleAvailabilityDoctorResponse, error) {
	// make variable to store data from request
	schedule := Schedule{
		DoctorId: req.GetDoctorId(),
	}

	// make variable to store data from database
	var scheduleDB []Schedule

	// make query to check if doctor id and day exist
	query := s.db.Where("doctor_id = ?", schedule.DoctorId).Find(&scheduleDB)

	// check if query error
	if query.Error != nil {
		// check if query error is not found
		if query.Error == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, "Schedule not found")
		}

		return nil, status.Error(codes.Internal, "Error when check schedule")
	}

	// make response in Data and with data types []*ScheduleAvailabilityDoctor
	res := &pb.ScheduleAvailabilityDoctorResponse{
		Data: []*pb.ScheduleAvailabilityDoctor{},
	}

	for _, item := range scheduleDB {
		res.Data = append(res.Data, &pb.ScheduleAvailabilityDoctor{
			DoctorId:  item.DoctorId,
			StartTime: item.Start_time,
			EndTime:   item.End_time,
			Day:       item.Day,
		})
	}

	res.Message = "Success read all schedule"
	res.Status = 200

	return res, nil
}

func (s *schedule) ReadAllScheduleByDay(ctx context.Context, req *pb.ScheduleAvailabilityDoctorRequest) (*pb.ScheduleAvailabilityDoctorResponse, error) {
	// make variable to store data from request
	schedule := Schedule{
		Day: req.GetDay(),
	}

	// make variable to store data from database
	var scheduleDB []Schedule

	// make query to check if doctor id and day exist
	query := s.db.Where("day = ?", schedule.Day).Find(&scheduleDB)

	// check if query error
	if query.Error != nil {
		// check if query error is not found
		if query.Error == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, "Schedule not found")
		}

		return nil, status.Error(codes.Internal, "Error when check schedule")
	}

	// make response in Data and with data types []*ScheduleAvailabilityDoctor
	res := &pb.ScheduleAvailabilityDoctorResponse{
		Data: []*pb.ScheduleAvailabilityDoctor{},
	}

	for _, item := range scheduleDB {
		res.Data = append(res.Data, &pb.ScheduleAvailabilityDoctor{
			DoctorId:  item.DoctorId,
			StartTime: item.Start_time,
			EndTime:   item.End_time,
			Day:       item.Day,
		})
	}

	res.Message = "Success read all schedule"
	res.Status = 200

	return res, nil
}

func (s *schedule) ReadAllScheduleByDoctorId(ctx context.Context, req *pb.ScheduleAvailabilityDoctorRequest) (*pb.ScheduleAvailabilityDoctorResponse, error) {
	// make variable to store data from request
	schedule := Schedule{
		DoctorId: req.GetDoctorId(),
	}

	// make variable to store data from database
	var scheduleDB []Schedule

	// make query to check if doctor id and day exist
	query := s.db.Where("doctor_id = ?", schedule.DoctorId).Find(&scheduleDB)

	// check if query error
	if query.Error != nil {
		// check if query error is not found
		if query.Error == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, "Schedule not found")
		}

		return nil, status.Error(codes.Internal, "Error when check schedule")
	}

	// make response in Data and with data types []*ScheduleAvailabilityDoctor
	res := &pb.ScheduleAvailabilityDoctorResponse{
		Data: []*pb.ScheduleAvailabilityDoctor{},
	}

	for _, item := range scheduleDB {
		res.Data = append(res.Data, &pb.ScheduleAvailabilityDoctor{
			DoctorId:  item.DoctorId,
			StartTime: item.Start_time,
			EndTime:   item.End_time,
			Day:       item.Day,
		})
	}

	res.Message = "Success read all schedule"
	res.Status = 200

	return res, nil
}

func (s *schedule) ReadAllScheduleByDoctorIdAndDay(ctx context.Context, req *pb.ScheduleAvailabilityDoctorRequest) (*pb.ScheduleAvailabilityDoctorResponse, error) {
	// make variable to store data from request
	schedule := Schedule{
		DoctorId: req.GetDoctorId(),
		Day:      req.GetDay(),
	}

	// make variable to store data from database
	var scheduleDB Schedule

	// make query to check if doctor id and day exist
	query := s.db.Where("doctor_id = ? AND day = ?", schedule.DoctorId, schedule.Day).First(&scheduleDB)

	// check if query error
	if query.Error != nil {
		// check if query error is not found
		if query.Error == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, "Schedule not found")
		}

		return nil, status.Error(codes.Internal, "Error when check schedule")
	}

	// make response in Data and with data types []*ScheduleAvailabilityDoctor
	res := &pb.ScheduleAvailabilityDoctorResponse{
		Data: []*pb.ScheduleAvailabilityDoctor{
			{
				DoctorId:  int32(scheduleDB.DoctorId),
				StartTime: scheduleDB.Start_time,
				EndTime:   scheduleDB.End_time,
				Day:       scheduleDB.Day,
			},
		},
		Message: "Success read all schedule",
		Status:  200,
	}

	return res, nil
}
