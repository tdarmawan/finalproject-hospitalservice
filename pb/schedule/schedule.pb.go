// make file proto for schedule availability doctor from doctor.proto file

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: schedule.proto

package schedule

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

type ScheduleAvailabilityDoctor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DoctorId  int32  `protobuf:"varint,2,opt,name=doctor_id,json=doctorId,proto3" json:"doctor_id,omitempty"`
	Day       string `protobuf:"bytes,3,opt,name=day,proto3" json:"day,omitempty"`
	StartTime string `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   string `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *ScheduleAvailabilityDoctor) Reset() {
	*x = ScheduleAvailabilityDoctor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleAvailabilityDoctor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleAvailabilityDoctor) ProtoMessage() {}

func (x *ScheduleAvailabilityDoctor) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleAvailabilityDoctor.ProtoReflect.Descriptor instead.
func (*ScheduleAvailabilityDoctor) Descriptor() ([]byte, []int) {
	return file_schedule_proto_rawDescGZIP(), []int{0}
}

func (x *ScheduleAvailabilityDoctor) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ScheduleAvailabilityDoctor) GetDoctorId() int32 {
	if x != nil {
		return x.DoctorId
	}
	return 0
}

func (x *ScheduleAvailabilityDoctor) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *ScheduleAvailabilityDoctor) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *ScheduleAvailabilityDoctor) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

type ScheduleAvailabilityDoctorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    []*ScheduleAvailabilityDoctor `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Message string                        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Status  int32                         `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ScheduleAvailabilityDoctorResponse) Reset() {
	*x = ScheduleAvailabilityDoctorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleAvailabilityDoctorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleAvailabilityDoctorResponse) ProtoMessage() {}

func (x *ScheduleAvailabilityDoctorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleAvailabilityDoctorResponse.ProtoReflect.Descriptor instead.
func (*ScheduleAvailabilityDoctorResponse) Descriptor() ([]byte, []int) {
	return file_schedule_proto_rawDescGZIP(), []int{1}
}

func (x *ScheduleAvailabilityDoctorResponse) GetData() []*ScheduleAvailabilityDoctor {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ScheduleAvailabilityDoctorResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ScheduleAvailabilityDoctorResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type ScheduleAvailabilityDoctorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DoctorId  int32  `protobuf:"varint,2,opt,name=doctor_id,json=doctorId,proto3" json:"doctor_id,omitempty"`
	Day       string `protobuf:"bytes,3,opt,name=day,proto3" json:"day,omitempty"`
	StartTime string `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   string `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *ScheduleAvailabilityDoctorRequest) Reset() {
	*x = ScheduleAvailabilityDoctorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleAvailabilityDoctorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleAvailabilityDoctorRequest) ProtoMessage() {}

func (x *ScheduleAvailabilityDoctorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleAvailabilityDoctorRequest.ProtoReflect.Descriptor instead.
func (*ScheduleAvailabilityDoctorRequest) Descriptor() ([]byte, []int) {
	return file_schedule_proto_rawDescGZIP(), []int{2}
}

func (x *ScheduleAvailabilityDoctorRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ScheduleAvailabilityDoctorRequest) GetDoctorId() int32 {
	if x != nil {
		return x.DoctorId
	}
	return 0
}

func (x *ScheduleAvailabilityDoctorRequest) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *ScheduleAvailabilityDoctorRequest) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *ScheduleAvailabilityDoctorRequest) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

type ScheduleAvailabilityDoctorRequestByDoctorId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DoctorId int32 `protobuf:"varint,1,opt,name=doctor_id,json=doctorId,proto3" json:"doctor_id,omitempty"`
}

func (x *ScheduleAvailabilityDoctorRequestByDoctorId) Reset() {
	*x = ScheduleAvailabilityDoctorRequestByDoctorId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedule_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleAvailabilityDoctorRequestByDoctorId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleAvailabilityDoctorRequestByDoctorId) ProtoMessage() {}

func (x *ScheduleAvailabilityDoctorRequestByDoctorId) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleAvailabilityDoctorRequestByDoctorId.ProtoReflect.Descriptor instead.
func (*ScheduleAvailabilityDoctorRequestByDoctorId) Descriptor() ([]byte, []int) {
	return file_schedule_proto_rawDescGZIP(), []int{3}
}

func (x *ScheduleAvailabilityDoctorRequestByDoctorId) GetDoctorId() int32 {
	if x != nil {
		return x.DoctorId
	}
	return 0
}

type ScheduleAvailabilityDoctorRequestByDay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Day string `protobuf:"bytes,1,opt,name=day,proto3" json:"day,omitempty"`
}

func (x *ScheduleAvailabilityDoctorRequestByDay) Reset() {
	*x = ScheduleAvailabilityDoctorRequestByDay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedule_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleAvailabilityDoctorRequestByDay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleAvailabilityDoctorRequestByDay) ProtoMessage() {}

func (x *ScheduleAvailabilityDoctorRequestByDay) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleAvailabilityDoctorRequestByDay.ProtoReflect.Descriptor instead.
func (*ScheduleAvailabilityDoctorRequestByDay) Descriptor() ([]byte, []int) {
	return file_schedule_proto_rawDescGZIP(), []int{4}
}

func (x *ScheduleAvailabilityDoctorRequestByDay) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

var File_schedule_proto protoreflect.FileDescriptor

var file_schedule_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x1a, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x6f,
	0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x22, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x9c, 0x01, 0x0a, 0x21, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x44, 0x6f,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x6f, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x4a, 0x0a, 0x2b, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x44, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x22, 0x3a, 0x0a, 0x26, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x44, 0x61, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79, 0x32, 0xa9, 0x05, 0x0a,
	0x21, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x63, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12,
	0x2b, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x44,
	0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7d, 0x0a, 0x16, 0x52, 0x65,
	0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x42, 0x79, 0x44, 0x6f, 0x63, 0x74,
	0x6f, 0x72, 0x49, 0x64, 0x12, 0x35, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x42, 0x79, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x1a, 0x2c, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x73, 0x0a, 0x11, 0x52, 0x65, 0x61,
	0x64, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x42, 0x79, 0x44, 0x61, 0x79, 0x12, 0x30,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x44, 0x6f,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x44, 0x61, 0x79,
	0x1a, 0x2c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2b, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schedule_proto_rawDescOnce sync.Once
	file_schedule_proto_rawDescData = file_schedule_proto_rawDesc
)

func file_schedule_proto_rawDescGZIP() []byte {
	file_schedule_proto_rawDescOnce.Do(func() {
		file_schedule_proto_rawDescData = protoimpl.X.CompressGZIP(file_schedule_proto_rawDescData)
	})
	return file_schedule_proto_rawDescData
}

var file_schedule_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_schedule_proto_goTypes = []interface{}{
	(*ScheduleAvailabilityDoctor)(nil),                  // 0: schedule.ScheduleAvailabilityDoctor
	(*ScheduleAvailabilityDoctorResponse)(nil),          // 1: schedule.ScheduleAvailabilityDoctorResponse
	(*ScheduleAvailabilityDoctorRequest)(nil),           // 2: schedule.ScheduleAvailabilityDoctorRequest
	(*ScheduleAvailabilityDoctorRequestByDoctorId)(nil), // 3: schedule.ScheduleAvailabilityDoctorRequestByDoctorId
	(*ScheduleAvailabilityDoctorRequestByDay)(nil),      // 4: schedule.ScheduleAvailabilityDoctorRequestByDay
}
var file_schedule_proto_depIdxs = []int32{
	0, // 0: schedule.ScheduleAvailabilityDoctorResponse.data:type_name -> schedule.ScheduleAvailabilityDoctor
	2, // 1: schedule.ScheduleAvailabilityDoctorService.Create:input_type -> schedule.ScheduleAvailabilityDoctorRequest
	2, // 2: schedule.ScheduleAvailabilityDoctorService.Read:input_type -> schedule.ScheduleAvailabilityDoctorRequest
	3, // 3: schedule.ScheduleAvailabilityDoctorService.ReadScheduleByDoctorId:input_type -> schedule.ScheduleAvailabilityDoctorRequestByDoctorId
	4, // 4: schedule.ScheduleAvailabilityDoctorService.ReadScheduleByDay:input_type -> schedule.ScheduleAvailabilityDoctorRequestByDay
	2, // 5: schedule.ScheduleAvailabilityDoctorService.Update:input_type -> schedule.ScheduleAvailabilityDoctorRequest
	2, // 6: schedule.ScheduleAvailabilityDoctorService.Delete:input_type -> schedule.ScheduleAvailabilityDoctorRequest
	1, // 7: schedule.ScheduleAvailabilityDoctorService.Create:output_type -> schedule.ScheduleAvailabilityDoctorResponse
	1, // 8: schedule.ScheduleAvailabilityDoctorService.Read:output_type -> schedule.ScheduleAvailabilityDoctorResponse
	1, // 9: schedule.ScheduleAvailabilityDoctorService.ReadScheduleByDoctorId:output_type -> schedule.ScheduleAvailabilityDoctorResponse
	1, // 10: schedule.ScheduleAvailabilityDoctorService.ReadScheduleByDay:output_type -> schedule.ScheduleAvailabilityDoctorResponse
	1, // 11: schedule.ScheduleAvailabilityDoctorService.Update:output_type -> schedule.ScheduleAvailabilityDoctorResponse
	1, // 12: schedule.ScheduleAvailabilityDoctorService.Delete:output_type -> schedule.ScheduleAvailabilityDoctorResponse
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_schedule_proto_init() }
func file_schedule_proto_init() {
	if File_schedule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schedule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleAvailabilityDoctor); i {
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
		file_schedule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleAvailabilityDoctorResponse); i {
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
		file_schedule_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleAvailabilityDoctorRequest); i {
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
		file_schedule_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleAvailabilityDoctorRequestByDoctorId); i {
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
		file_schedule_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleAvailabilityDoctorRequestByDay); i {
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
			RawDescriptor: file_schedule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_schedule_proto_goTypes,
		DependencyIndexes: file_schedule_proto_depIdxs,
		MessageInfos:      file_schedule_proto_msgTypes,
	}.Build()
	File_schedule_proto = out.File
	file_schedule_proto_rawDesc = nil
	file_schedule_proto_goTypes = nil
	file_schedule_proto_depIdxs = nil
}
