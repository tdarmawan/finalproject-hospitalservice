// make proto file for appointment service (createappointment, readappointment, emailappointment)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: appointment.proto

package appointment

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateAppointmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AppointmentCode string `protobuf:"bytes,2,opt,name=appointmentCode,proto3" json:"appointmentCode,omitempty"`
	PatientId       string `protobuf:"bytes,3,opt,name=patientId,proto3" json:"patientId,omitempty"`
	PatientEmail    string `protobuf:"bytes,4,opt,name=patientEmail,proto3" json:"patientEmail,omitempty"`
	DoctorId        int32  `protobuf:"varint,5,opt,name=doctorId,proto3" json:"doctorId,omitempty"`
	ScheduleId      int32  `protobuf:"varint,6,opt,name=scheduleId,proto3" json:"scheduleId,omitempty"`
	Status          string `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Message         string `protobuf:"bytes,8,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateAppointmentRequest) Reset() {
	*x = CreateAppointmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appointment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAppointmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppointmentRequest) ProtoMessage() {}

func (x *CreateAppointmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_appointment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppointmentRequest.ProtoReflect.Descriptor instead.
func (*CreateAppointmentRequest) Descriptor() ([]byte, []int) {
	return file_appointment_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAppointmentRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateAppointmentRequest) GetAppointmentCode() string {
	if x != nil {
		return x.AppointmentCode
	}
	return ""
}

func (x *CreateAppointmentRequest) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

func (x *CreateAppointmentRequest) GetPatientEmail() string {
	if x != nil {
		return x.PatientEmail
	}
	return ""
}

func (x *CreateAppointmentRequest) GetDoctorId() int32 {
	if x != nil {
		return x.DoctorId
	}
	return 0
}

func (x *CreateAppointmentRequest) GetScheduleId() int32 {
	if x != nil {
		return x.ScheduleId
	}
	return 0
}

func (x *CreateAppointmentRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateAppointmentRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateAppointmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    []*CreateAppointmentRequest `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Message string                      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Status  int32                       `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateAppointmentResponse) Reset() {
	*x = CreateAppointmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appointment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAppointmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppointmentResponse) ProtoMessage() {}

func (x *CreateAppointmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_appointment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppointmentResponse.ProtoReflect.Descriptor instead.
func (*CreateAppointmentResponse) Descriptor() ([]byte, []int) {
	return file_appointment_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAppointmentResponse) GetData() []*CreateAppointmentRequest {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CreateAppointmentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateAppointmentResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type AppointmentId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppointmentCode string `protobuf:"bytes,1,opt,name=appointmentCode,proto3" json:"appointmentCode,omitempty"`
}

func (x *AppointmentId) Reset() {
	*x = AppointmentId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appointment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppointmentId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppointmentId) ProtoMessage() {}

func (x *AppointmentId) ProtoReflect() protoreflect.Message {
	mi := &file_appointment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppointmentId.ProtoReflect.Descriptor instead.
func (*AppointmentId) Descriptor() ([]byte, []int) {
	return file_appointment_proto_rawDescGZIP(), []int{2}
}

func (x *AppointmentId) GetAppointmentCode() string {
	if x != nil {
		return x.AppointmentCode
	}
	return ""
}

type ReadAppointmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    []*CreateAppointmentRequest `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Message string                      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Status  int32                       `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ReadAppointmentResponse) Reset() {
	*x = ReadAppointmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appointment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAppointmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAppointmentResponse) ProtoMessage() {}

func (x *ReadAppointmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_appointment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAppointmentResponse.ProtoReflect.Descriptor instead.
func (*ReadAppointmentResponse) Descriptor() ([]byte, []int) {
	return file_appointment_proto_rawDescGZIP(), []int{3}
}

func (x *ReadAppointmentResponse) GetData() []*CreateAppointmentRequest {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ReadAppointmentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ReadAppointmentResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type EmailAppointmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppointmentCode string `protobuf:"bytes,1,opt,name=appointmentCode,proto3" json:"appointmentCode,omitempty"`
}

func (x *EmailAppointmentRequest) Reset() {
	*x = EmailAppointmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appointment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailAppointmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailAppointmentRequest) ProtoMessage() {}

func (x *EmailAppointmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_appointment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailAppointmentRequest.ProtoReflect.Descriptor instead.
func (*EmailAppointmentRequest) Descriptor() ([]byte, []int) {
	return file_appointment_proto_rawDescGZIP(), []int{4}
}

func (x *EmailAppointmentRequest) GetAppointmentCode() string {
	if x != nil {
		return x.AppointmentCode
	}
	return ""
}

type EmailAppointmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    []*CreateAppointmentRequest `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Message string                      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Status  int32                       `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *EmailAppointmentResponse) Reset() {
	*x = EmailAppointmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appointment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailAppointmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailAppointmentResponse) ProtoMessage() {}

func (x *EmailAppointmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_appointment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailAppointmentResponse.ProtoReflect.Descriptor instead.
func (*EmailAppointmentResponse) Descriptor() ([]byte, []int) {
	return file_appointment_proto_rawDescGZIP(), []int{5}
}

func (x *EmailAppointmentResponse) GetData() []*CreateAppointmentRequest {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *EmailAppointmentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *EmailAppointmentResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appointment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_appointment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_appointment_proto_rawDescGZIP(), []int{6}
}

var File_appointment_proto protoreflect.FileDescriptor

var file_appointment_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x22, 0x84, 0x02, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a,
	0x0f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x88, 0x01, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x39, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x70,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x86, 0x01,
	0x0a, 0x17, 0x52, 0x65, 0x61, 0x64, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x43, 0x0a, 0x17, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41,
	0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x18,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x8a,
	0x03, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x13, 0x52,
	0x65, 0x61, 0x64, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x24,
	0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x10, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41,
	0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x70,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x12, 0x52, 0x65, 0x61,
	0x64, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x12,
	0x12, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e,
	0x2f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_appointment_proto_rawDescOnce sync.Once
	file_appointment_proto_rawDescData = file_appointment_proto_rawDesc
)

func file_appointment_proto_rawDescGZIP() []byte {
	file_appointment_proto_rawDescOnce.Do(func() {
		file_appointment_proto_rawDescData = protoimpl.X.CompressGZIP(file_appointment_proto_rawDescData)
	})
	return file_appointment_proto_rawDescData
}

var file_appointment_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_appointment_proto_goTypes = []interface{}{
	(*CreateAppointmentRequest)(nil),  // 0: appointment.CreateAppointmentRequest
	(*CreateAppointmentResponse)(nil), // 1: appointment.CreateAppointmentResponse
	(*AppointmentId)(nil),             // 2: appointment.AppointmentId
	(*ReadAppointmentResponse)(nil),   // 3: appointment.ReadAppointmentResponse
	(*EmailAppointmentRequest)(nil),   // 4: appointment.EmailAppointmentRequest
	(*EmailAppointmentResponse)(nil),  // 5: appointment.EmailAppointmentResponse
	(*Empty)(nil),                     // 6: appointment.Empty
}
var file_appointment_proto_depIdxs = []int32{
	0, // 0: appointment.CreateAppointmentResponse.data:type_name -> appointment.CreateAppointmentRequest
	0, // 1: appointment.ReadAppointmentResponse.data:type_name -> appointment.CreateAppointmentRequest
	0, // 2: appointment.EmailAppointmentResponse.data:type_name -> appointment.CreateAppointmentRequest
	0, // 3: appointment.AppointmentService.CreateAppointment:input_type -> appointment.CreateAppointmentRequest
	2, // 4: appointment.AppointmentService.ReadAppointmentById:input_type -> appointment.AppointmentId
	4, // 5: appointment.AppointmentService.EmailAppointment:input_type -> appointment.EmailAppointmentRequest
	6, // 6: appointment.AppointmentService.ReadAppointmentAll:input_type -> appointment.Empty
	1, // 7: appointment.AppointmentService.CreateAppointment:output_type -> appointment.CreateAppointmentResponse
	3, // 8: appointment.AppointmentService.ReadAppointmentById:output_type -> appointment.ReadAppointmentResponse
	5, // 9: appointment.AppointmentService.EmailAppointment:output_type -> appointment.EmailAppointmentResponse
	3, // 10: appointment.AppointmentService.ReadAppointmentAll:output_type -> appointment.ReadAppointmentResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_appointment_proto_init() }
func file_appointment_proto_init() {
	if File_appointment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_appointment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAppointmentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_appointment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAppointmentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_appointment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppointmentId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_appointment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAppointmentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_appointment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailAppointmentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_appointment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailAppointmentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_appointment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_appointment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_appointment_proto_goTypes,
		DependencyIndexes: file_appointment_proto_depIdxs,
		MessageInfos:      file_appointment_proto_msgTypes,
	}.Build()
	File_appointment_proto = out.File
	file_appointment_proto_rawDesc = nil
	file_appointment_proto_goTypes = nil
	file_appointment_proto_depIdxs = nil
}