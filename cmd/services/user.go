package services

import (
	"context"
	"fmt"
	pb "testing/pb/user"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/streadway/amqp"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type server struct {
	pb.UnimplementedUserDatakuServer
	db *gorm.DB
}

type CustomClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"unique"`
	Password  string `gorm:"not null"`
	Name      string
	Age       int32
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func NewUserService(db *gorm.DB) *server {
	server := &server{
		db: db,
	}
	// if err := server.StartSchedulerJob(); err != nil {
	// 	log.Fatalf("failed to start scheduled job: %v", err)
	// }
	return server
}

func ConnectToMessageBroker() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (s *server) AuthUser(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	//make function to login
	//use bcrypt to compare password
	//if password match, generate token
	//if password not match, return error
	//return token
	// Find user in PostgreSQL
	var user User
	err := s.db.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	// Verify password using bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Email,

		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})

	// Sign token with secret key
	// TODO: use asymmetric signing when in production
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not generate token")
	}

	return &pb.AuthResponse{
		Token: tokenString,
	}, nil
}

// grpcurl -plaintext -d '{"email":"jamal@gmail.com", "password":"12345"}' -import-path ./proto -proto user.proto localhost:50053 testing.Userku/GenerateToken
func (s *server) GenerateToken(ctx context.Context, req *pb.GenerateTokenRequest) (*pb.GenerateTokenResponse, error) {
	// TODO: validate request

	// Find user in PostgreSQL
	var user User
	err := s.db.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	// Verify password using bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err :=

		token.SignedString([]byte("mysecret"))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate token")
	}

	// Return response
	return &pb.GenerateTokenResponse{
		Token: tokenString,
	}, nil
}

// grpcurl -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODE1NDM4ODgsInN1YiI6IiJ9.ETfQhTZgPduKWLZSQbLW26g6HB3HK-_NVxL4yRwZmjc' -plaintext -d ' ' -import-path ./proto -proto user.proto localhost:50053 testing.Userku/GetAllData
func (s *server) GetAllData(ctx context.Context, req *pb.GetAllUserRequest) (*pb.GetAllUserResponse, error) {
	// Retrieve the user data from PostgreSQL
	var users []*pb.User
	err := s.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &pb.GetAllUserResponse{Data: users}, nil

}

// grpcurl -plaintext -d '{"user_id": "1"}' -import-path ./proto -proto user.proto localhost:50053 testing.Userku/GetUser
func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	// Retrieve the user data from PostgreSQL
	var user pb.User
	err := s.db.First(&user, req.UserId).Error
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	return &user, nil

}

// grpcurl -plaintext -d '{"name": "farid", "age":"21", "email":"tes@gmail.com","password":"12345"}' -import-path ./proto -proto user.proto localhost:50053 testing.Userku/AddUser
func (s *server) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	// Hash the password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password")
	}

	// Create a new user
	user := User{
		Name:     req.Name,
		Age:      req.Age,
		Email:    req.Email,
		Password: string(passwordHash),
	}

	// Insert the user into PostgreSQL
	err = s.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	// Return the response
	res := &pb.AddUserResponse{
		Id: fmt.Sprint(user.ID),
	}

	return res, nil

}

func (s *server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	// Update the user in PostgreSQL
	var user User
	err := s.db.First(&user, req.Id).Error
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	user.Name = req.Name
	user.Age = req.Age

	err = s.db.Save(&user).Error
	if err != nil {
		return nil, err
	}

	res := &pb.UpdateUserResponse{
		Message: fmt.Sprintf("User with ID %s successfully updated", req.Id),
	}

	return res, nil
}
