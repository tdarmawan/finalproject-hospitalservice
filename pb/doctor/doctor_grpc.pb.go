//make proto file for doctor data with field doctor_id, name, speciality, email, password, phone, address, city, state, country, profile_pic, status, created_at, updated_at

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: doctor.proto

package doctor

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DoctorService_CreateDoctor_FullMethodName          = "/doctor.DoctorService/CreateDoctor"
	DoctorService_GetDoctorById_FullMethodName         = "/doctor.DoctorService/GetDoctorById"
	DoctorService_GetDoctorByHospitalId_FullMethodName = "/doctor.DoctorService/GetDoctorByHospitalId"
	DoctorService_UpdateDoctor_FullMethodName          = "/doctor.DoctorService/UpdateDoctor"
	DoctorService_DeleteDoctor_FullMethodName          = "/doctor.DoctorService/DeleteDoctor"
	DoctorService_GetAllDoctors_FullMethodName         = "/doctor.DoctorService/GetAllDoctors"
)

// DoctorServiceClient is the client API for DoctorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DoctorServiceClient interface {
	CreateDoctor(ctx context.Context, in *DoctorRequest, opts ...grpc.CallOption) (*DoctorResponse, error)
	GetDoctorById(ctx context.Context, in *DoctorId, opts ...grpc.CallOption) (*DoctorResponse, error)
	GetDoctorByHospitalId(ctx context.Context, in *HospitalId, opts ...grpc.CallOption) (*DoctorList, error)
	UpdateDoctor(ctx context.Context, in *DoctorRequest, opts ...grpc.CallOption) (*DoctorResponse, error)
	DeleteDoctor(ctx context.Context, in *DoctorId, opts ...grpc.CallOption) (*DoctorResponse, error)
	GetAllDoctors(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DoctorList, error)
}

type doctorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDoctorServiceClient(cc grpc.ClientConnInterface) DoctorServiceClient {
	return &doctorServiceClient{cc}
}

func (c *doctorServiceClient) CreateDoctor(ctx context.Context, in *DoctorRequest, opts ...grpc.CallOption) (*DoctorResponse, error) {
	out := new(DoctorResponse)
	err := c.cc.Invoke(ctx, DoctorService_CreateDoctor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) GetDoctorById(ctx context.Context, in *DoctorId, opts ...grpc.CallOption) (*DoctorResponse, error) {
	out := new(DoctorResponse)
	err := c.cc.Invoke(ctx, DoctorService_GetDoctorById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) GetDoctorByHospitalId(ctx context.Context, in *HospitalId, opts ...grpc.CallOption) (*DoctorList, error) {
	out := new(DoctorList)
	err := c.cc.Invoke(ctx, DoctorService_GetDoctorByHospitalId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) UpdateDoctor(ctx context.Context, in *DoctorRequest, opts ...grpc.CallOption) (*DoctorResponse, error) {
	out := new(DoctorResponse)
	err := c.cc.Invoke(ctx, DoctorService_UpdateDoctor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) DeleteDoctor(ctx context.Context, in *DoctorId, opts ...grpc.CallOption) (*DoctorResponse, error) {
	out := new(DoctorResponse)
	err := c.cc.Invoke(ctx, DoctorService_DeleteDoctor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doctorServiceClient) GetAllDoctors(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DoctorList, error) {
	out := new(DoctorList)
	err := c.cc.Invoke(ctx, DoctorService_GetAllDoctors_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DoctorServiceServer is the server API for DoctorService service.
// All implementations must embed UnimplementedDoctorServiceServer
// for forward compatibility
type DoctorServiceServer interface {
	CreateDoctor(context.Context, *DoctorRequest) (*DoctorResponse, error)
	GetDoctorById(context.Context, *DoctorId) (*DoctorResponse, error)
	GetDoctorByHospitalId(context.Context, *HospitalId) (*DoctorList, error)
	UpdateDoctor(context.Context, *DoctorRequest) (*DoctorResponse, error)
	DeleteDoctor(context.Context, *DoctorId) (*DoctorResponse, error)
	GetAllDoctors(context.Context, *Empty) (*DoctorList, error)
	mustEmbedUnimplementedDoctorServiceServer()
}

// UnimplementedDoctorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDoctorServiceServer struct {
}

func (UnimplementedDoctorServiceServer) CreateDoctor(context.Context, *DoctorRequest) (*DoctorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDoctor not implemented")
}
func (UnimplementedDoctorServiceServer) GetDoctorById(context.Context, *DoctorId) (*DoctorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDoctorById not implemented")
}
func (UnimplementedDoctorServiceServer) GetDoctorByHospitalId(context.Context, *HospitalId) (*DoctorList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDoctorByHospitalId not implemented")
}
func (UnimplementedDoctorServiceServer) UpdateDoctor(context.Context, *DoctorRequest) (*DoctorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDoctor not implemented")
}
func (UnimplementedDoctorServiceServer) DeleteDoctor(context.Context, *DoctorId) (*DoctorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDoctor not implemented")
}
func (UnimplementedDoctorServiceServer) GetAllDoctors(context.Context, *Empty) (*DoctorList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllDoctors not implemented")
}
func (UnimplementedDoctorServiceServer) mustEmbedUnimplementedDoctorServiceServer() {}

// UnsafeDoctorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DoctorServiceServer will
// result in compilation errors.
type UnsafeDoctorServiceServer interface {
	mustEmbedUnimplementedDoctorServiceServer()
}

func RegisterDoctorServiceServer(s grpc.ServiceRegistrar, srv DoctorServiceServer) {
	s.RegisterService(&DoctorService_ServiceDesc, srv)
}

func _DoctorService_CreateDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoctorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).CreateDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_CreateDoctor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).CreateDoctor(ctx, req.(*DoctorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_GetDoctorById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoctorId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).GetDoctorById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_GetDoctorById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).GetDoctorById(ctx, req.(*DoctorId))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_GetDoctorByHospitalId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HospitalId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).GetDoctorByHospitalId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_GetDoctorByHospitalId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).GetDoctorByHospitalId(ctx, req.(*HospitalId))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_UpdateDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoctorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).UpdateDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_UpdateDoctor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).UpdateDoctor(ctx, req.(*DoctorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_DeleteDoctor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoctorId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).DeleteDoctor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_DeleteDoctor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).DeleteDoctor(ctx, req.(*DoctorId))
	}
	return interceptor(ctx, in, info, handler)
}

func _DoctorService_GetAllDoctors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoctorServiceServer).GetAllDoctors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DoctorService_GetAllDoctors_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoctorServiceServer).GetAllDoctors(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// DoctorService_ServiceDesc is the grpc.ServiceDesc for DoctorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DoctorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "doctor.DoctorService",
	HandlerType: (*DoctorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDoctor",
			Handler:    _DoctorService_CreateDoctor_Handler,
		},
		{
			MethodName: "GetDoctorById",
			Handler:    _DoctorService_GetDoctorById_Handler,
		},
		{
			MethodName: "GetDoctorByHospitalId",
			Handler:    _DoctorService_GetDoctorByHospitalId_Handler,
		},
		{
			MethodName: "UpdateDoctor",
			Handler:    _DoctorService_UpdateDoctor_Handler,
		},
		{
			MethodName: "DeleteDoctor",
			Handler:    _DoctorService_DeleteDoctor_Handler,
		},
		{
			MethodName: "GetAllDoctors",
			Handler:    _DoctorService_GetAllDoctors_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "doctor.proto",
}
