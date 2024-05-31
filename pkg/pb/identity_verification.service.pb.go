// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: pkg/pb/identity_verification.service.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// VerifyNINRequest is the request message for verifying NIN
type VerifyNINRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdNumber   string `protobuf:"bytes,1,opt,name=idNumber,proto3" json:"idNumber,omitempty"`
	Firstname  string `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname   string `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Middlename string `protobuf:"bytes,4,opt,name=middlename,proto3" json:"middlename,omitempty"`
	Dob        string `protobuf:"bytes,5,opt,name=dob,proto3" json:"dob,omitempty"`
	Phone      string `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	Email      string `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	Gender     string `protobuf:"bytes,8,opt,name=gender,proto3" json:"gender,omitempty"`
}

func (x *VerifyNINRequest) Reset() {
	*x = VerifyNINRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_identity_verification_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyNINRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyNINRequest) ProtoMessage() {}

func (x *VerifyNINRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_identity_verification_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyNINRequest.ProtoReflect.Descriptor instead.
func (*VerifyNINRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_identity_verification_service_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyNINRequest) GetIdNumber() string {
	if x != nil {
		return x.IdNumber
	}
	return ""
}

func (x *VerifyNINRequest) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *VerifyNINRequest) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *VerifyNINRequest) GetMiddlename() string {
	if x != nil {
		return x.Middlename
	}
	return ""
}

func (x *VerifyNINRequest) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

func (x *VerifyNINRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *VerifyNINRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *VerifyNINRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

type Applicant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Firstname string `protobuf:"bytes,1,opt,name=Firstname,proto3" json:"Firstname,omitempty"`
	Lastname  string `protobuf:"bytes,2,opt,name=Lastname,proto3" json:"Lastname,omitempty"`
}

func (x *Applicant) Reset() {
	*x = Applicant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_identity_verification_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Applicant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Applicant) ProtoMessage() {}

func (x *Applicant) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_identity_verification_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Applicant.ProtoReflect.Descriptor instead.
func (*Applicant) Descriptor() ([]byte, []int) {
	return file_pkg_pb_identity_verification_service_proto_rawDescGZIP(), []int{1}
}

func (x *Applicant) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *Applicant) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

type FieldMatches struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Firstname bool `protobuf:"varint,1,opt,name=Firstname,proto3" json:"Firstname,omitempty"`
	Lastname  bool `protobuf:"varint,2,opt,name=Lastname,proto3" json:"Lastname,omitempty"`
}

func (x *FieldMatches) Reset() {
	*x = FieldMatches{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_identity_verification_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldMatches) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldMatches) ProtoMessage() {}

func (x *FieldMatches) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_identity_verification_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldMatches.ProtoReflect.Descriptor instead.
func (*FieldMatches) Descriptor() ([]byte, []int) {
	return file_pkg_pb_identity_verification_service_proto_rawDescGZIP(), []int{2}
}

func (x *FieldMatches) GetFirstname() bool {
	if x != nil {
		return x.Firstname
	}
	return false
}

func (x *FieldMatches) GetLastname() bool {
	if x != nil {
		return x.Lastname
	}
	return false
}

type NinCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       string        `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	FieldMatches *FieldMatches `protobuf:"bytes,2,opt,name=FieldMatches,proto3" json:"FieldMatches,omitempty"`
}

func (x *NinCheck) Reset() {
	*x = NinCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_identity_verification_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NinCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NinCheck) ProtoMessage() {}

func (x *NinCheck) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_identity_verification_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NinCheck.ProtoReflect.Descriptor instead.
func (*NinCheck) Descriptor() ([]byte, []int) {
	return file_pkg_pb_identity_verification_service_proto_rawDescGZIP(), []int{3}
}

func (x *NinCheck) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *NinCheck) GetFieldMatches() *FieldMatches {
	if x != nil {
		return x.FieldMatches
	}
	return nil
}

type Summary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NinCheck *NinCheck `protobuf:"bytes,1,opt,name=NinCheck,proto3" json:"NinCheck,omitempty"`
}

func (x *Summary) Reset() {
	*x = Summary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_identity_verification_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Summary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Summary) ProtoMessage() {}

func (x *Summary) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_identity_verification_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Summary.ProtoReflect.Descriptor instead.
func (*Summary) Descriptor() ([]byte, []int) {
	return file_pkg_pb_identity_verification_service_proto_rawDescGZIP(), []int{4}
}

func (x *Summary) GetNinCheck() *NinCheck {
	if x != nil {
		return x.NinCheck
	}
	return nil
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State  string `protobuf:"bytes,1,opt,name=State,proto3" json:"State,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_identity_verification_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_identity_verification_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_pkg_pb_identity_verification_service_proto_rawDescGZIP(), []int{5}
}

func (x *Status) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Status) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Nin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nin        string `protobuf:"bytes,1,opt,name=Nin,proto3" json:"Nin,omitempty"`
	Firstname  string `protobuf:"bytes,2,opt,name=Firstname,proto3" json:"Firstname,omitempty"`
	Lastname   string `protobuf:"bytes,3,opt,name=Lastname,proto3" json:"Lastname,omitempty"`
	Middlename string `protobuf:"bytes,4,opt,name=Middlename,proto3" json:"Middlename,omitempty"`
	Phone      string `protobuf:"bytes,5,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Gender     string `protobuf:"bytes,6,opt,name=Gender,proto3" json:"Gender,omitempty"`
	Birthdate  string `protobuf:"bytes,7,opt,name=Birthdate,proto3" json:"Birthdate,omitempty"`
	Photo      string `protobuf:"bytes,8,opt,name=Photo,proto3" json:"Photo,omitempty"`
	Address    string `protobuf:"bytes,9,opt,name=Address,proto3" json:"Address,omitempty"`
}

func (x *Nin) Reset() {
	*x = Nin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_identity_verification_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nin) ProtoMessage() {}

func (x *Nin) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_identity_verification_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nin.ProtoReflect.Descriptor instead.
func (*Nin) Descriptor() ([]byte, []int) {
	return file_pkg_pb_identity_verification_service_proto_rawDescGZIP(), []int{6}
}

func (x *Nin) GetNin() string {
	if x != nil {
		return x.Nin
	}
	return ""
}

func (x *Nin) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *Nin) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *Nin) GetMiddlename() string {
	if x != nil {
		return x.Middlename
	}
	return ""
}

func (x *Nin) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Nin) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Nin) GetBirthdate() string {
	if x != nil {
		return x.Birthdate
	}
	return ""
}

func (x *Nin) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *Nin) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// VerifyNINResponse is the response message for verifying NIN
type VerifyNINResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32      `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Applicant *Applicant `protobuf:"bytes,2,opt,name=Applicant,proto3" json:"Applicant,omitempty"`
	Summary   *Summary   `protobuf:"bytes,3,opt,name=Summary,proto3" json:"Summary,omitempty"`
	Status    *Status    `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`
	Nin       *Nin       `protobuf:"bytes,5,opt,name=Nin,proto3" json:"Nin,omitempty"`
}

func (x *VerifyNINResponse) Reset() {
	*x = VerifyNINResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_identity_verification_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyNINResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyNINResponse) ProtoMessage() {}

func (x *VerifyNINResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_identity_verification_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyNINResponse.ProtoReflect.Descriptor instead.
func (*VerifyNINResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_identity_verification_service_proto_rawDescGZIP(), []int{7}
}

func (x *VerifyNINResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VerifyNINResponse) GetApplicant() *Applicant {
	if x != nil {
		return x.Applicant
	}
	return nil
}

func (x *VerifyNINResponse) GetSummary() *Summary {
	if x != nil {
		return x.Summary
	}
	return nil
}

func (x *VerifyNINResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *VerifyNINResponse) GetNin() *Nin {
	if x != nil {
		return x.Nin
	}
	return nil
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId  string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	SecretKey string `protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_identity_verification_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_identity_verification_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_identity_verification_service_proto_rawDescGZIP(), []int{8}
}

func (x *LoginRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *LoginRequest) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_identity_verification_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_identity_verification_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_identity_verification_service_proto_rawDescGZIP(), []int{9}
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var file_pkg_pb_identity_verification_service_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50001,
		Name:          "identity_verification.json_name",
		Tag:           "bytes,50001,opt,name=json_name",
		Filename:      "pkg/pb/identity_verification.service.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional string json_name = 50001;
	E_JsonName = &file_pkg_pb_identity_verification_service_proto_extTypes[0]
)

var File_pkg_pb_identity_verification_service_proto protoreflect.FileDescriptor

var file_pkg_pb_identity_verification_service_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x01, 0x0a, 0x10, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x4e, 0x49, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64,
	0x6f, 0x62, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x62, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0x8a, 0xb5, 0x18, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x28, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0c, 0x8a, 0xb5, 0x18, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x65, 0x0a, 0x0c, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x09, 0x46, 0x69,
	0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0d, 0x8a,
	0xb5, 0x18, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x09, 0x46, 0x69,
	0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0c, 0x8a, 0xb5, 0x18, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x89, 0x01, 0x0a, 0x08, 0x4e, 0x69, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x22,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0x8a, 0xb5, 0x18, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x59, 0x0a, 0x0c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x42, 0x10, 0x8a,
	0xb5, 0x18, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52,
	0x0c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x22, 0x55, 0x0a,
	0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x4a, 0x0a, 0x08, 0x4e, 0x69, 0x6e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x4e, 0x69, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x0d, 0x8a, 0xb5, 0x18,
	0x09, 0x6e, 0x69, 0x6e, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x08, 0x4e, 0x69, 0x6e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x22, 0x4d, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x8a,
	0xb5, 0x18, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x22, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0a, 0x8a, 0xb5, 0x18, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0xe1, 0x02, 0x0a, 0x03, 0x4e, 0x69, 0x6e, 0x12, 0x19, 0x0a, 0x03, 0x4e,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0x8a, 0xb5, 0x18, 0x03, 0x6e, 0x69,
	0x6e, 0x52, 0x03, 0x4e, 0x69, 0x6e, 0x12, 0x2b, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0x8a, 0xb5, 0x18, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0x8a, 0xb5, 0x18, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a,
	0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0e, 0x8a, 0xb5, 0x18, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x52, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x8a, 0xb5,
	0x18, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x22,
	0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0x8a, 0xb5, 0x18, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x2b, 0x0a, 0x09, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0x8a, 0xb5, 0x18, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x09, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1f, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09,
	0x8a, 0xb5, 0x18, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x74, 0x6f,
	0x12, 0x25, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0b, 0x8a, 0xb5, 0x18, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xbb, 0x02, 0x0a, 0x11, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x4e, 0x49, 0x4e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x69,
	0x64, 0x52, 0x02, 0x49, 0x64, 0x12, 0x4d, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x6e, 0x74, 0x42, 0x0d, 0x8a, 0xb5, 0x18, 0x09,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x6e, 0x74, 0x52, 0x09, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x6e, 0x74, 0x12, 0x45, 0x0a, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x42, 0x0b, 0x8a, 0xb5, 0x18, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x52, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x41, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0a, 0x8a, 0xb5, 0x18, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x35,
	0x0a, 0x03, 0x4e, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x69, 0x6e, 0x42, 0x07, 0x8a, 0xb5, 0x18, 0x03, 0x6e, 0x69, 0x6e,
	0x52, 0x03, 0x4e, 0x69, 0x6e, 0x22, 0x4a, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65,
	0x79, 0x22, 0x25, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xaa, 0x02, 0x0a, 0x13, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x5e, 0x0a, 0x09, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4e, 0x49, 0x4e, 0x12, 0x27, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4e, 0x49, 0x4e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x4e, 0x49, 0x4e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5f, 0x0a, 0x0a, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x56, 0x4e, 0x49, 0x4e, 0x12, 0x27,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4e, 0x49, 0x4e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4e, 0x49, 0x4e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x52, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x23, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x3a, 0x3c, 0x0a, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xd1, 0x86, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_identity_verification_service_proto_rawDescOnce sync.Once
	file_pkg_pb_identity_verification_service_proto_rawDescData = file_pkg_pb_identity_verification_service_proto_rawDesc
)

func file_pkg_pb_identity_verification_service_proto_rawDescGZIP() []byte {
	file_pkg_pb_identity_verification_service_proto_rawDescOnce.Do(func() {
		file_pkg_pb_identity_verification_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_identity_verification_service_proto_rawDescData)
	})
	return file_pkg_pb_identity_verification_service_proto_rawDescData
}

var file_pkg_pb_identity_verification_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkg_pb_identity_verification_service_proto_goTypes = []interface{}{
	(*VerifyNINRequest)(nil),          // 0: identity_verification.VerifyNINRequest
	(*Applicant)(nil),                 // 1: identity_verification.Applicant
	(*FieldMatches)(nil),              // 2: identity_verification.FieldMatches
	(*NinCheck)(nil),                  // 3: identity_verification.NinCheck
	(*Summary)(nil),                   // 4: identity_verification.Summary
	(*Status)(nil),                    // 5: identity_verification.Status
	(*Nin)(nil),                       // 6: identity_verification.Nin
	(*VerifyNINResponse)(nil),         // 7: identity_verification.VerifyNINResponse
	(*LoginRequest)(nil),              // 8: identity_verification.LoginRequest
	(*LoginResponse)(nil),             // 9: identity_verification.LoginResponse
	(*descriptorpb.FieldOptions)(nil), // 10: google.protobuf.FieldOptions
}
var file_pkg_pb_identity_verification_service_proto_depIdxs = []int32{
	2,  // 0: identity_verification.NinCheck.FieldMatches:type_name -> identity_verification.FieldMatches
	3,  // 1: identity_verification.Summary.NinCheck:type_name -> identity_verification.NinCheck
	1,  // 2: identity_verification.VerifyNINResponse.Applicant:type_name -> identity_verification.Applicant
	4,  // 3: identity_verification.VerifyNINResponse.Summary:type_name -> identity_verification.Summary
	5,  // 4: identity_verification.VerifyNINResponse.Status:type_name -> identity_verification.Status
	6,  // 5: identity_verification.VerifyNINResponse.Nin:type_name -> identity_verification.Nin
	10, // 6: identity_verification.json_name:extendee -> google.protobuf.FieldOptions
	0,  // 7: identity_verification.VerificationService.VerifyNIN:input_type -> identity_verification.VerifyNINRequest
	0,  // 8: identity_verification.VerificationService.VerifyVNIN:input_type -> identity_verification.VerifyNINRequest
	8,  // 9: identity_verification.VerificationService.Login:input_type -> identity_verification.LoginRequest
	7,  // 10: identity_verification.VerificationService.VerifyNIN:output_type -> identity_verification.VerifyNINResponse
	7,  // 11: identity_verification.VerificationService.VerifyVNIN:output_type -> identity_verification.VerifyNINResponse
	9,  // 12: identity_verification.VerificationService.Login:output_type -> identity_verification.LoginResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	6,  // [6:7] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_pb_identity_verification_service_proto_init() }
func file_pkg_pb_identity_verification_service_proto_init() {
	if File_pkg_pb_identity_verification_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_identity_verification_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyNINRequest); i {
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
		file_pkg_pb_identity_verification_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Applicant); i {
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
		file_pkg_pb_identity_verification_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldMatches); i {
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
		file_pkg_pb_identity_verification_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NinCheck); i {
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
		file_pkg_pb_identity_verification_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Summary); i {
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
		file_pkg_pb_identity_verification_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_pkg_pb_identity_verification_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nin); i {
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
		file_pkg_pb_identity_verification_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyNINResponse); i {
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
		file_pkg_pb_identity_verification_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_pkg_pb_identity_verification_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
			RawDescriptor: file_pkg_pb_identity_verification_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 1,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_identity_verification_service_proto_goTypes,
		DependencyIndexes: file_pkg_pb_identity_verification_service_proto_depIdxs,
		MessageInfos:      file_pkg_pb_identity_verification_service_proto_msgTypes,
		ExtensionInfos:    file_pkg_pb_identity_verification_service_proto_extTypes,
	}.Build()
	File_pkg_pb_identity_verification_service_proto = out.File
	file_pkg_pb_identity_verification_service_proto_rawDesc = nil
	file_pkg_pb_identity_verification_service_proto_goTypes = nil
	file_pkg_pb_identity_verification_service_proto_depIdxs = nil
}