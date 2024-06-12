// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: pkg/pb/model/user.model.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_INACTIVE Status = 0
	Status_ACTIVE   Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "INACTIVE",
		1: "ACTIVE",
	}
	Status_value = map[string]int32{
		"INACTIVE": 0,
		"ACTIVE":   1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_pb_model_user_model_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_pkg_pb_model_user_model_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_pkg_pb_model_user_model_proto_rawDescGZIP(), []int{0}
}

// Enum for identity types
type IdentityType int32

const (
	IdentityType_NIN               IdentityType = 0
	IdentityType_VNIN              IdentityType = 1
	IdentityType_NIGERIAN_PASSPORT IdentityType = 2
	IdentityType_BVN               IdentityType = 3
	IdentityType_DRIVERS_LICENSE   IdentityType = 4
)

// Enum value maps for IdentityType.
var (
	IdentityType_name = map[int32]string{
		0: "NIN",
		1: "VNIN",
		2: "NIGERIAN_PASSPORT",
		3: "BVN",
		4: "DRIVERS_LICENSE",
	}
	IdentityType_value = map[string]int32{
		"NIN":               0,
		"VNIN":              1,
		"NIGERIAN_PASSPORT": 2,
		"BVN":               3,
		"DRIVERS_LICENSE":   4,
	}
)

func (x IdentityType) Enum() *IdentityType {
	p := new(IdentityType)
	*p = x
	return p
}

func (x IdentityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IdentityType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_pb_model_user_model_proto_enumTypes[1].Descriptor()
}

func (IdentityType) Type() protoreflect.EnumType {
	return &file_pkg_pb_model_user_model_proto_enumTypes[1]
}

func (x IdentityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IdentityType.Descriptor instead.
func (IdentityType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_pb_model_user_model_proto_rawDescGZIP(), []int{1}
}

type IdType int32

const (
	IdType_DRIVERS_LICENCE IdType = 0
	IdType_PASSPORT        IdType = 1
	IdType_IDENTITY_CARD   IdType = 2
)

// Enum value maps for IdType.
var (
	IdType_name = map[int32]string{
		0: "DRIVERS_LICENCE",
		1: "PASSPORT",
		2: "IDENTITY_CARD",
	}
	IdType_value = map[string]int32{
		"DRIVERS_LICENCE": 0,
		"PASSPORT":        1,
		"IDENTITY_CARD":   2,
	}
)

func (x IdType) Enum() *IdType {
	p := new(IdType)
	*p = x
	return p
}

func (x IdType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IdType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_pb_model_user_model_proto_enumTypes[2].Descriptor()
}

func (IdType) Type() protoreflect.EnumType {
	return &file_pkg_pb_model_user_model_proto_enumTypes[2]
}

func (x IdType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IdType.Descriptor instead.
func (IdType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_pb_model_user_model_proto_rawDescGZIP(), []int{2}
}

type VerificationStatus int32

const (
	VerificationStatus_PENDING    VerificationStatus = 0
	VerificationStatus_PROCESSING VerificationStatus = 1
	VerificationStatus_VERIFIED   VerificationStatus = 2
	VerificationStatus_FAILED     VerificationStatus = 3
	VerificationStatus_PARTIAL    VerificationStatus = 4
)

// Enum value maps for VerificationStatus.
var (
	VerificationStatus_name = map[int32]string{
		0: "PENDING",
		1: "PROCESSING",
		2: "VERIFIED",
		3: "FAILED",
		4: "PARTIAL",
	}
	VerificationStatus_value = map[string]int32{
		"PENDING":    0,
		"PROCESSING": 1,
		"VERIFIED":   2,
		"FAILED":     3,
		"PARTIAL":    4,
	}
)

func (x VerificationStatus) Enum() *VerificationStatus {
	p := new(VerificationStatus)
	*p = x
	return p
}

func (x VerificationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VerificationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_pb_model_user_model_proto_enumTypes[3].Descriptor()
}

func (VerificationStatus) Type() protoreflect.EnumType {
	return &file_pkg_pb_model_user_model_proto_enumTypes[3]
}

func (x VerificationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VerificationStatus.Descriptor instead.
func (VerificationStatus) EnumDescriptor() ([]byte, []int) {
	return file_pkg_pb_model_user_model_proto_rawDescGZIP(), []int{3}
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32                  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Address   string                 `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
	City      string                 `protobuf:"bytes,3,opt,name=City,proto3" json:"City,omitempty"`
	State     string                 `protobuf:"bytes,4,opt,name=State,proto3" json:"State,omitempty"`
	Zip       string                 `protobuf:"bytes,5,opt,name=Zip,proto3" json:"Zip,omitempty"`
	Country   string                 `protobuf:"bytes,6,opt,name=Country,proto3" json:"Country,omitempty"`
	Type      string                 `protobuf:"bytes,7,opt,name=Type,proto3" json:"Type,omitempty"`
	User      *User                  `protobuf:"bytes,8,opt,name=User,proto3" json:"User,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_model_user_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_model_user_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_pkg_pb_model_user_model_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Address) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetZip() string {
	if x != nil {
		return x.Zip
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Address) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Address) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Address) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type UserPermission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User       *User                  `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	Permission string                 `protobuf:"bytes,2,opt,name=Permission,proto3" json:"Permission,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	Status     Status                 `protobuf:"varint,5,opt,name=Status,proto3,enum=Status" json:"Status,omitempty"`
	Id         int32                  `protobuf:"varint,6,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *UserPermission) Reset() {
	*x = UserPermission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_model_user_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPermission) ProtoMessage() {}

func (x *UserPermission) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_model_user_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPermission.ProtoReflect.Descriptor instead.
func (*UserPermission) Descriptor() ([]byte, []int) {
	return file_pkg_pb_model_user_model_proto_rawDescGZIP(), []int{1}
}

func (x *UserPermission) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserPermission) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *UserPermission) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserPermission) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *UserPermission) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_INACTIVE
}

func (x *UserPermission) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ! This should be defined elsewere
	Id                 string                 `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Email              string                 `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Username           string                 `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	Password           string                 `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty"`
	Firstname          string                 `protobuf:"bytes,5,opt,name=Firstname,proto3" json:"Firstname,omitempty"`
	Lastname           string                 `protobuf:"bytes,6,opt,name=Lastname,proto3" json:"Lastname,omitempty"`
	Telephone          string                 `protobuf:"bytes,7,opt,name=Telephone,proto3" json:"Telephone,omitempty"`
	Token              string                 `protobuf:"bytes,8,opt,name=Token,proto3" json:"Token,omitempty"`
	Role               string                 `protobuf:"bytes,9,opt,name=Role,proto3" json:"Role,omitempty"`
	ImageUrl           string                 `protobuf:"bytes,10,opt,name=ImageUrl,json=imageurl,proto3" json:"ImageUrl,omitempty"`
	Bio                string                 `protobuf:"bytes,11,opt,name=Bio,json=bio,proto3" json:"Bio,omitempty"`
	CreatedAt          *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=CreatedAt,json=created_at,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt          *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=UpdatedAt,json=updated_at,proto3" json:"UpdatedAt,omitempty"`
	VerificationStatus VerificationStatus     `protobuf:"varint,14,opt,name=VerificationStatus,json=verificationstatus,proto3,enum=VerificationStatus" json:"VerificationStatus,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_model_user_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_model_user_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_pkg_pb_model_user_model_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *User) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *User) GetTelephone() string {
	if x != nil {
		return x.Telephone
	}
	return ""
}

func (x *User) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *User) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *User) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *User) GetVerificationStatus() VerificationStatus {
	if x != nil {
		return x.VerificationStatus
	}
	return VerificationStatus_PENDING
}

type UserVerification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountryCode string `protobuf:"bytes,1,opt,name=CountryCode,proto3" json:"CountryCode,omitempty"`
	IdType      IdType `protobuf:"varint,2,opt,name=IdType,proto3,enum=IdType" json:"IdType,omitempty"`
	Selfie      string `protobuf:"bytes,3,opt,name=Selfie,proto3" json:"Selfie,omitempty"`
	IdFilePath  string `protobuf:"bytes,4,opt,name=IdFilePath,proto3" json:"IdFilePath,omitempty"`
	IdNumber    string `protobuf:"bytes,5,opt,name=IdNumber,proto3" json:"IdNumber,omitempty"`
	User        *User  `protobuf:"bytes,6,opt,name=User,proto3" json:"User,omitempty"`
	FirstName   string `protobuf:"bytes,7,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName    string `protobuf:"bytes,8,opt,name=LastName,proto3" json:"LastName,omitempty"`
}

func (x *UserVerification) Reset() {
	*x = UserVerification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_model_user_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserVerification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserVerification) ProtoMessage() {}

func (x *UserVerification) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_model_user_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserVerification.ProtoReflect.Descriptor instead.
func (*UserVerification) Descriptor() ([]byte, []int) {
	return file_pkg_pb_model_user_model_proto_rawDescGZIP(), []int{3}
}

func (x *UserVerification) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *UserVerification) GetIdType() IdType {
	if x != nil {
		return x.IdType
	}
	return IdType_DRIVERS_LICENCE
}

func (x *UserVerification) GetSelfie() string {
	if x != nil {
		return x.Selfie
	}
	return ""
}

func (x *UserVerification) GetIdFilePath() string {
	if x != nil {
		return x.IdFilePath
	}
	return ""
}

func (x *UserVerification) GetIdNumber() string {
	if x != nil {
		return x.IdNumber
	}
	return ""
}

func (x *UserVerification) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserVerification) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserVerification) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalPages  int32 `protobuf:"varint,1,opt,name=TotalPages,proto3" json:"TotalPages,omitempty"`
	Limit       int32 `protobuf:"varint,2,opt,name=Limit,proto3" json:"Limit,omitempty"`
	CurrentPage int32 `protobuf:"varint,3,opt,name=CurrentPage,proto3" json:"CurrentPage,omitempty"`
	NextPage    int32 `protobuf:"varint,4,opt,name=NextPage,proto3" json:"NextPage,omitempty"`
	TotalCount  int32 `protobuf:"varint,5,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_model_user_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_model_user_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_pkg_pb_model_user_model_proto_rawDescGZIP(), []int{4}
}

func (x *Meta) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *Meta) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Meta) GetCurrentPage() int32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *Meta) GetNextPage() int32 {
	if x != nil {
		return x.NextPage
	}
	return 0
}

func (x *Meta) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_pkg_pb_model_user_model_proto protoreflect.FileDescriptor

var file_pkg_pb_model_user_model_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x67,
	0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x02, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x11, 0xba, 0xb9, 0x19, 0x0d, 0x0a, 0x0b, 0x12, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x65, 0x72, 0x28, 0x01, 0x52, 0x02, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x5a, 0x69, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x5a, 0x69, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x22, 0x00, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x22, 0x93, 0x02, 0x0a, 0x0e,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x22, 0x00, 0x52, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x11, 0xba, 0xb9, 0x19, 0x0d, 0x0a, 0x0b, 0x12, 0x07, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x65, 0x72, 0x28, 0x01, 0x52, 0x02, 0x49, 0x64, 0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08,
	0x01, 0x22, 0xe7, 0x03, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xba, 0xb9, 0x19, 0x0a, 0x0a, 0x08, 0x12, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x28, 0x01, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x75, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x42, 0x69, 0x6f, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x12, 0x39, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x12, 0x39, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x12, 0x43,
	0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x12, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x22, 0x92, 0x02, 0x0a, 0x10,
	0x55, 0x73, 0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x49, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x07, 0x2e, 0x49, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x49, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x6c, 0x66, 0x69, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x65, 0x6c, 0x66, 0x69, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x49,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x49, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x49,
	0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49,
	0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x0a, 0xba, 0xb9,
	0x19, 0x06, 0x22, 0x04, 0x1a, 0x02, 0x69, 0x64, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01,
	0x22, 0x9a, 0x01, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x2a, 0x22, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x41, 0x43, 0x54,
	0x49, 0x56, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10,
	0x01, 0x2a, 0x56, 0x0a, 0x0c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x4e,
	0x49, 0x4e, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x4e, 0x49, 0x47, 0x45, 0x52, 0x49, 0x41, 0x4e,
	0x5f, 0x50, 0x41, 0x53, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x42,
	0x56, 0x4e, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x52, 0x49, 0x56, 0x45, 0x52, 0x53, 0x5f,
	0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x10, 0x04, 0x2a, 0x3e, 0x0a, 0x06, 0x49, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x52, 0x49, 0x56, 0x45, 0x52, 0x53, 0x5f, 0x4c,
	0x49, 0x43, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x41, 0x53, 0x53,
	0x50, 0x4f, 0x52, 0x54, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49,
	0x54, 0x59, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x02, 0x2a, 0x58, 0x0a, 0x12, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x41, 0x52, 0x54, 0x49, 0x41,
	0x4c, 0x10, 0x04, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_model_user_model_proto_rawDescOnce sync.Once
	file_pkg_pb_model_user_model_proto_rawDescData = file_pkg_pb_model_user_model_proto_rawDesc
)

func file_pkg_pb_model_user_model_proto_rawDescGZIP() []byte {
	file_pkg_pb_model_user_model_proto_rawDescOnce.Do(func() {
		file_pkg_pb_model_user_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_model_user_model_proto_rawDescData)
	})
	return file_pkg_pb_model_user_model_proto_rawDescData
}

var file_pkg_pb_model_user_model_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_pkg_pb_model_user_model_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_pb_model_user_model_proto_goTypes = []interface{}{
	(Status)(0),                   // 0: Status
	(IdentityType)(0),             // 1: IdentityType
	(IdType)(0),                   // 2: IdType
	(VerificationStatus)(0),       // 3: VerificationStatus
	(*Address)(nil),               // 4: Address
	(*UserPermission)(nil),        // 5: UserPermission
	(*User)(nil),                  // 6: User
	(*UserVerification)(nil),      // 7: UserVerification
	(*Meta)(nil),                  // 8: Meta
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_pkg_pb_model_user_model_proto_depIdxs = []int32{
	6,  // 0: Address.User:type_name -> User
	9,  // 1: Address.CreatedAt:type_name -> google.protobuf.Timestamp
	9,  // 2: Address.UpdatedAt:type_name -> google.protobuf.Timestamp
	6,  // 3: UserPermission.User:type_name -> User
	9,  // 4: UserPermission.CreatedAt:type_name -> google.protobuf.Timestamp
	9,  // 5: UserPermission.UpdatedAt:type_name -> google.protobuf.Timestamp
	0,  // 6: UserPermission.Status:type_name -> Status
	9,  // 7: User.CreatedAt:type_name -> google.protobuf.Timestamp
	9,  // 8: User.UpdatedAt:type_name -> google.protobuf.Timestamp
	3,  // 9: User.VerificationStatus:type_name -> VerificationStatus
	2,  // 10: UserVerification.IdType:type_name -> IdType
	6,  // 11: UserVerification.User:type_name -> User
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_pkg_pb_model_user_model_proto_init() }
func file_pkg_pb_model_user_model_proto_init() {
	if File_pkg_pb_model_user_model_proto != nil {
		return
	}
	file_pkg_pb_model_gorm_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_model_user_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_pkg_pb_model_user_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPermission); i {
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
		file_pkg_pb_model_user_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_pkg_pb_model_user_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserVerification); i {
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
		file_pkg_pb_model_user_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
			RawDescriptor: file_pkg_pb_model_user_model_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_pb_model_user_model_proto_goTypes,
		DependencyIndexes: file_pkg_pb_model_user_model_proto_depIdxs,
		EnumInfos:         file_pkg_pb_model_user_model_proto_enumTypes,
		MessageInfos:      file_pkg_pb_model_user_model_proto_msgTypes,
	}.Build()
	File_pkg_pb_model_user_model_proto = out.File
	file_pkg_pb_model_user_model_proto_rawDesc = nil
	file_pkg_pb_model_user_model_proto_goTypes = nil
	file_pkg_pb_model_user_model_proto_depIdxs = nil
}
