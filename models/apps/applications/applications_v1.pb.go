/// This file has messages for describing applications

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.7
// source: apps/applications/applications_v1.proto

package applications

import (
	contours "github.com/badhouseplants/envspotting-go-proto/models/apps/contours"
	common "github.com/badhouseplants/envspotting-go-proto/models/common"
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

//*
// Represents an application name only
type AppNameAndDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`               // Application name
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"` // Application brief description
}

func (x *AppNameAndDescription) Reset() {
	*x = AppNameAndDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_applications_applications_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppNameAndDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppNameAndDescription) ProtoMessage() {}

func (x *AppNameAndDescription) ProtoReflect() protoreflect.Message {
	mi := &file_apps_applications_applications_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppNameAndDescription.ProtoReflect.Descriptor instead.
func (*AppNameAndDescription) Descriptor() ([]byte, []int) {
	return file_apps_applications_applications_v1_proto_rawDescGZIP(), []int{0}
}

func (x *AppNameAndDescription) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppNameAndDescription) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

//*
// Represents an application UUID only
type AppId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // Application id: UUID
}

func (x *AppId) Reset() {
	*x = AppId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_applications_applications_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppId) ProtoMessage() {}

func (x *AppId) ProtoReflect() protoreflect.Message {
	mi := &file_apps_applications_applications_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppId.ProtoReflect.Descriptor instead.
func (*AppId) Descriptor() ([]byte, []int) {
	return file_apps_applications_applications_v1_proto_rawDescGZIP(), []int{1}
}

func (x *AppId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

//*
// Represents an application name and id
type AppIdAndName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`     // UUID
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // Application name: Unique string
}

func (x *AppIdAndName) Reset() {
	*x = AppIdAndName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_applications_applications_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppIdAndName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppIdAndName) ProtoMessage() {}

func (x *AppIdAndName) ProtoReflect() protoreflect.Message {
	mi := &file_apps_applications_applications_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppIdAndName.ProtoReflect.Descriptor instead.
func (*AppIdAndName) Descriptor() ([]byte, []int) {
	return file_apps_applications_applications_v1_proto_rawDescGZIP(), []int{2}
}

func (x *AppIdAndName) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AppIdAndName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//*
// Represents a application
type AppWithoutContours struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                   // UUID
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`               // Application name: Unique string
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"` // Application brief description
}

func (x *AppWithoutContours) Reset() {
	*x = AppWithoutContours{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_applications_applications_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppWithoutContours) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppWithoutContours) ProtoMessage() {}

func (x *AppWithoutContours) ProtoReflect() protoreflect.Message {
	mi := &file_apps_applications_applications_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppWithoutContours.ProtoReflect.Descriptor instead.
func (*AppWithoutContours) Descriptor() ([]byte, []int) {
	return file_apps_applications_applications_v1_proto_rawDescGZIP(), []int{3}
}

func (x *AppWithoutContours) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AppWithoutContours) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppWithoutContours) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

//*
// Represents an application with contours
type AppFullInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                   // UUID
	Name        string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`               // Application name: Unique string
	Description string                  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"` // Application brief description
	Contours    []*contours.ContourInfo `protobuf:"bytes,4,rep,name=contours,proto3" json:"contours,omitempty"`       // Map: <contour name: [services]>
}

func (x *AppFullInfo) Reset() {
	*x = AppFullInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_applications_applications_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppFullInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppFullInfo) ProtoMessage() {}

func (x *AppFullInfo) ProtoReflect() protoreflect.Message {
	mi := &file_apps_applications_applications_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppFullInfo.ProtoReflect.Descriptor instead.
func (*AppFullInfo) Descriptor() ([]byte, []int) {
	return file_apps_applications_applications_v1_proto_rawDescGZIP(), []int{4}
}

func (x *AppFullInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AppFullInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppFullInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AppFullInfo) GetContours() []*contours.ContourInfo {
	if x != nil {
		return x.Contours
	}
	return nil
}

type ListOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Added bool `protobuf:"varint,1,opt,name=added,proto3" json:"added,omitempty"`
}

func (x *ListOptions) Reset() {
	*x = ListOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_applications_applications_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOptions) ProtoMessage() {}

func (x *ListOptions) ProtoReflect() protoreflect.Message {
	mi := &file_apps_applications_applications_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOptions.ProtoReflect.Descriptor instead.
func (*ListOptions) Descriptor() ([]byte, []int) {
	return file_apps_applications_applications_v1_proto_rawDescGZIP(), []int{5}
}

func (x *ListOptions) GetAdded() bool {
	if x != nil {
		return x.Added
	}
	return false
}

var File_apps_applications_applications_v1_proto protoreflect.FileDescriptor

var file_apps_applications_applications_v1_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x70, 0x70, 0x73, 0x1a,
	0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x6f, 0x75, 0x72, 0x73, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x6f, 0x75, 0x72, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x15, 0x41, 0x70, 0x70, 0x4e,
	0x61, 0x6d, 0x65, 0x41, 0x6e, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x17, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x32, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5a, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x57, 0x69, 0x74, 0x68, 0x6f,
	0x75, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x82, 0x01, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x6f, 0x75,
	0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x6f, 0x75, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x6f, 0x6e,
	0x74, 0x6f, 0x75, 0x72, 0x73, 0x22, 0x23, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x32, 0xa9, 0x02, 0x0a, 0x0c, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x41, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x41, 0x70, 0x70,
	0x4e, 0x61, 0x6d, 0x65, 0x41, 0x6e, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x57, 0x69, 0x74,
	0x68, 0x6f, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x6f, 0x75, 0x72, 0x73, 0x22, 0x00, 0x12, 0x27,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0b, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x41, 0x70, 0x70,
	0x49, 0x64, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x46, 0x75, 0x6c,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x11, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x57, 0x69, 0x74,
	0x68, 0x6f, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x6f, 0x75, 0x72, 0x73, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x3e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x70,
	0x73, 0x2e, 0x41, 0x70, 0x70, 0x57, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x6f, 0x75, 0x72, 0x73, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x57,
	0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x6f, 0x75, 0x72, 0x73, 0x22, 0x00,
	0x12, 0x34, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x70,
	0x73, 0x2e, 0x41, 0x70, 0x70, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x14,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x64, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x73, 0x2f, 0x65, 0x6e, 0x76, 0x73, 0x70, 0x6f, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2d,
	0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_applications_applications_v1_proto_rawDescOnce sync.Once
	file_apps_applications_applications_v1_proto_rawDescData = file_apps_applications_applications_v1_proto_rawDesc
)

func file_apps_applications_applications_v1_proto_rawDescGZIP() []byte {
	file_apps_applications_applications_v1_proto_rawDescOnce.Do(func() {
		file_apps_applications_applications_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_applications_applications_v1_proto_rawDescData)
	})
	return file_apps_applications_applications_v1_proto_rawDescData
}

var file_apps_applications_applications_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_apps_applications_applications_v1_proto_goTypes = []interface{}{
	(*AppNameAndDescription)(nil), // 0: apps.AppNameAndDescription
	(*AppId)(nil),                 // 1: apps.AppId
	(*AppIdAndName)(nil),          // 2: apps.AppIdAndName
	(*AppWithoutContours)(nil),    // 3: apps.AppWithoutContours
	(*AppFullInfo)(nil),           // 4: apps.AppFullInfo
	(*ListOptions)(nil),           // 5: apps.ListOptions
	(*contours.ContourInfo)(nil),  // 6: apps.ContourInfo
	(*common.EmptyMessage)(nil),   // 7: common.EmptyMessage
}
var file_apps_applications_applications_v1_proto_depIdxs = []int32{
	6, // 0: apps.AppFullInfo.contours:type_name -> apps.ContourInfo
	0, // 1: apps.Applications.Create:input_type -> apps.AppNameAndDescription
	1, // 2: apps.Applications.Get:input_type -> apps.AppId
	5, // 3: apps.Applications.List:input_type -> apps.ListOptions
	3, // 4: apps.Applications.Update:input_type -> apps.AppWithoutContours
	2, // 5: apps.Applications.Delete:input_type -> apps.AppIdAndName
	3, // 6: apps.Applications.Create:output_type -> apps.AppWithoutContours
	4, // 7: apps.Applications.Get:output_type -> apps.AppFullInfo
	3, // 8: apps.Applications.List:output_type -> apps.AppWithoutContours
	3, // 9: apps.Applications.Update:output_type -> apps.AppWithoutContours
	7, // 10: apps.Applications.Delete:output_type -> common.EmptyMessage
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apps_applications_applications_v1_proto_init() }
func file_apps_applications_applications_v1_proto_init() {
	if File_apps_applications_applications_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_applications_applications_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppNameAndDescription); i {
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
		file_apps_applications_applications_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppId); i {
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
		file_apps_applications_applications_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppIdAndName); i {
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
		file_apps_applications_applications_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppWithoutContours); i {
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
		file_apps_applications_applications_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppFullInfo); i {
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
		file_apps_applications_applications_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOptions); i {
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
			RawDescriptor: file_apps_applications_applications_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_applications_applications_v1_proto_goTypes,
		DependencyIndexes: file_apps_applications_applications_v1_proto_depIdxs,
		MessageInfos:      file_apps_applications_applications_v1_proto_msgTypes,
	}.Build()
	File_apps_applications_applications_v1_proto = out.File
	file_apps_applications_applications_v1_proto_rawDesc = nil
	file_apps_applications_applications_v1_proto_goTypes = nil
	file_apps_applications_applications_v1_proto_depIdxs = nil
}
