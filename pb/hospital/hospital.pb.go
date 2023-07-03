// make proto file for hospital data and make a service for it
// the proto hospital can have doctor and patient as a nested message
// the proto file can have a service for the hospital

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: hospital.proto

package hospital

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

type Patient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdentityId   string `protobuf:"bytes,1,opt,name=identityId,proto3" json:"identityId,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age          int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Disease      string `protobuf:"bytes,4,opt,name=disease,proto3" json:"disease,omitempty"`
	Hospitalcode string `protobuf:"bytes,5,opt,name=hospitalcode,proto3" json:"hospitalcode,omitempty"`
}

func (x *Patient) Reset() {
	*x = Patient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hospital_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Patient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Patient) ProtoMessage() {}

func (x *Patient) ProtoReflect() protoreflect.Message {
	mi := &file_hospital_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Patient.ProtoReflect.Descriptor instead.
func (*Patient) Descriptor() ([]byte, []int) {
	return file_hospital_proto_rawDescGZIP(), []int{0}
}

func (x *Patient) GetIdentityId() string {
	if x != nil {
		return x.IdentityId
	}
	return ""
}

func (x *Patient) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Patient) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Patient) GetDisease() string {
	if x != nil {
		return x.Disease
	}
	return ""
}

func (x *Patient) GetHospitalcode() string {
	if x != nil {
		return x.Hospitalcode
	}
	return ""
}

type Hospital struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hospitalcode string `protobuf:"bytes,1,opt,name=hospitalcode,proto3" json:"hospitalcode,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address      string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Hospital) Reset() {
	*x = Hospital{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hospital_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hospital) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hospital) ProtoMessage() {}

func (x *Hospital) ProtoReflect() protoreflect.Message {
	mi := &file_hospital_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hospital.ProtoReflect.Descriptor instead.
func (*Hospital) Descriptor() ([]byte, []int) {
	return file_hospital_proto_rawDescGZIP(), []int{1}
}

func (x *Hospital) GetHospitalcode() string {
	if x != nil {
		return x.Hospitalcode
	}
	return ""
}

func (x *Hospital) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Hospital) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type HospitalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hospitalcode string `protobuf:"bytes,1,opt,name=hospitalcode,proto3" json:"hospitalcode,omitempty"`
}

func (x *HospitalRequest) Reset() {
	*x = HospitalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hospital_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HospitalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HospitalRequest) ProtoMessage() {}

func (x *HospitalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hospital_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HospitalRequest.ProtoReflect.Descriptor instead.
func (*HospitalRequest) Descriptor() ([]byte, []int) {
	return file_hospital_proto_rawDescGZIP(), []int{2}
}

func (x *HospitalRequest) GetHospitalcode() string {
	if x != nil {
		return x.Hospitalcode
	}
	return ""
}

type HospitalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    []*Hospital `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Message string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Status  int32       `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *HospitalResponse) Reset() {
	*x = HospitalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hospital_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HospitalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HospitalResponse) ProtoMessage() {}

func (x *HospitalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hospital_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HospitalResponse.ProtoReflect.Descriptor instead.
func (*HospitalResponse) Descriptor() ([]byte, []int) {
	return file_hospital_proto_rawDescGZIP(), []int{3}
}

func (x *HospitalResponse) GetData() []*Hospital {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *HospitalResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *HospitalResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type PatientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdentityId string `protobuf:"bytes,1,opt,name=identityId,proto3" json:"identityId,omitempty"`
}

func (x *PatientRequest) Reset() {
	*x = PatientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hospital_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientRequest) ProtoMessage() {}

func (x *PatientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hospital_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientRequest.ProtoReflect.Descriptor instead.
func (*PatientRequest) Descriptor() ([]byte, []int) {
	return file_hospital_proto_rawDescGZIP(), []int{4}
}

func (x *PatientRequest) GetIdentityId() string {
	if x != nil {
		return x.IdentityId
	}
	return ""
}

type PatientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data    []*Patient `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Status  int32      `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *PatientResponse) Reset() {
	*x = PatientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hospital_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientResponse) ProtoMessage() {}

func (x *PatientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hospital_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientResponse.ProtoReflect.Descriptor instead.
func (*PatientResponse) Descriptor() ([]byte, []int) {
	return file_hospital_proto_rawDescGZIP(), []int{5}
}

func (x *PatientResponse) GetData() []*Patient {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PatientResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PatientResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type PatientId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdentityId string `protobuf:"bytes,1,opt,name=identityId,proto3" json:"identityId,omitempty"`
}

func (x *PatientId) Reset() {
	*x = PatientId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hospital_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatientId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientId) ProtoMessage() {}

func (x *PatientId) ProtoReflect() protoreflect.Message {
	mi := &file_hospital_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientId.ProtoReflect.Descriptor instead.
func (*PatientId) Descriptor() ([]byte, []int) {
	return file_hospital_proto_rawDescGZIP(), []int{6}
}

func (x *PatientId) GetIdentityId() string {
	if x != nil {
		return x.IdentityId
	}
	return ""
}

var File_hospital_proto protoreflect.FileDescriptor

var file_hospital_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x22, 0x8d, 0x01, 0x0a, 0x07, 0x50,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x69, 0x73, 0x65, 0x61, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64,
	0x69, 0x73, 0x65, 0x61, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74,
	0x61, 0x6c, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x6f,
	0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x5c, 0x0a, 0x08, 0x48, 0x6f,
	0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74,
	0x61, 0x6c, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x68, 0x6f,
	0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x35, 0x0a, 0x0f, 0x48, 0x6f, 0x73, 0x70,
	0x69, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x68,
	0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x63, 0x6f, 0x64, 0x65, 0x22,
	0x6c, 0x0a, 0x10, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x48, 0x6f, 0x73,
	0x70, 0x69, 0x74, 0x61, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x30, 0x0a,
	0x0e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22,
	0x6a, 0x0a, 0x0f, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2b, 0x0a, 0x09, 0x50,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x32, 0xe5, 0x02, 0x0a, 0x0f, 0x48, 0x6f, 0x73,
	0x70, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x12, 0x12,
	0x2e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74,
	0x61, 0x6c, 0x1a, 0x1a, 0x2e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x48, 0x6f,
	0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x46, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x12,
	0x19, 0x2e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x48, 0x6f, 0x73, 0x70, 0x69,
	0x74, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x68, 0x6f, 0x73,
	0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x48, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74,
	0x61, 0x6c, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x50, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c,
	0x0a, 0x0a, 0x41, 0x64, 0x64, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x11, 0x2e, 0x68,
	0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x1a,
	0x19, 0x2e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x13,
	0x2e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x1a, 0x19, 0x2e, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x2e, 0x50,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x68, 0x6f, 0x73, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hospital_proto_rawDescOnce sync.Once
	file_hospital_proto_rawDescData = file_hospital_proto_rawDesc
)

func file_hospital_proto_rawDescGZIP() []byte {
	file_hospital_proto_rawDescOnce.Do(func() {
		file_hospital_proto_rawDescData = protoimpl.X.CompressGZIP(file_hospital_proto_rawDescData)
	})
	return file_hospital_proto_rawDescData
}

var file_hospital_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_hospital_proto_goTypes = []interface{}{
	(*Patient)(nil),          // 0: hospital.Patient
	(*Hospital)(nil),         // 1: hospital.Hospital
	(*HospitalRequest)(nil),  // 2: hospital.HospitalRequest
	(*HospitalResponse)(nil), // 3: hospital.HospitalResponse
	(*PatientRequest)(nil),   // 4: hospital.PatientRequest
	(*PatientResponse)(nil),  // 5: hospital.PatientResponse
	(*PatientId)(nil),        // 6: hospital.PatientId
}
var file_hospital_proto_depIdxs = []int32{
	1, // 0: hospital.HospitalResponse.data:type_name -> hospital.Hospital
	0, // 1: hospital.PatientResponse.data:type_name -> hospital.Patient
	1, // 2: hospital.HospitalService.CreateHospital:input_type -> hospital.Hospital
	2, // 3: hospital.HospitalService.GetHospital:input_type -> hospital.HospitalRequest
	4, // 4: hospital.HospitalService.GetPatients:input_type -> hospital.PatientRequest
	0, // 5: hospital.HospitalService.AddPatient:input_type -> hospital.Patient
	6, // 6: hospital.HospitalService.GetPatientById:input_type -> hospital.PatientId
	3, // 7: hospital.HospitalService.CreateHospital:output_type -> hospital.HospitalResponse
	3, // 8: hospital.HospitalService.GetHospital:output_type -> hospital.HospitalResponse
	5, // 9: hospital.HospitalService.GetPatients:output_type -> hospital.PatientResponse
	5, // 10: hospital.HospitalService.AddPatient:output_type -> hospital.PatientResponse
	5, // 11: hospital.HospitalService.GetPatientById:output_type -> hospital.PatientResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_hospital_proto_init() }
func file_hospital_proto_init() {
	if File_hospital_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hospital_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Patient); i {
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
		file_hospital_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hospital); i {
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
		file_hospital_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HospitalRequest); i {
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
		file_hospital_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HospitalResponse); i {
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
		file_hospital_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatientRequest); i {
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
		file_hospital_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatientResponse); i {
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
		file_hospital_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatientId); i {
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
			RawDescriptor: file_hospital_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hospital_proto_goTypes,
		DependencyIndexes: file_hospital_proto_depIdxs,
		MessageInfos:      file_hospital_proto_msgTypes,
	}.Build()
	File_hospital_proto = out.File
	file_hospital_proto_rawDesc = nil
	file_hospital_proto_goTypes = nil
	file_hospital_proto_depIdxs = nil
}