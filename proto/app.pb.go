// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: proto/app.proto

package proto

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User    string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Segmant string `protobuf:"bytes,2,opt,name=segmant,proto3" json:"segmant,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_app_proto_msgTypes[0]
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
	return file_proto_app_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *User) GetSegmant() string {
	if x != nil {
		return x.Segmant
	}
	return ""
}

type UserSegmantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User    string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Segmant string `protobuf:"bytes,2,opt,name=segmant,proto3" json:"segmant,omitempty"`
}

func (x *UserSegmantRequest) Reset() {
	*x = UserSegmantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSegmantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSegmantRequest) ProtoMessage() {}

func (x *UserSegmantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSegmantRequest.ProtoReflect.Descriptor instead.
func (*UserSegmantRequest) Descriptor() ([]byte, []int) {
	return file_proto_app_proto_rawDescGZIP(), []int{1}
}

func (x *UserSegmantRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *UserSegmantRequest) GetSegmant() string {
	if x != nil {
		return x.Segmant
	}
	return ""
}

type UserSegmantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success string `protobuf:"bytes,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UserSegmantResponse) Reset() {
	*x = UserSegmantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSegmantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSegmantResponse) ProtoMessage() {}

func (x *UserSegmantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSegmantResponse.ProtoReflect.Descriptor instead.
func (*UserSegmantResponse) Descriptor() ([]byte, []int) {
	return file_proto_app_proto_rawDescGZIP(), []int{2}
}

func (x *UserSegmantResponse) GetSuccess() string {
	if x != nil {
		return x.Success
	}
	return ""
}

func (x *UserSegmantResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SegmantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Segmant string `protobuf:"bytes,1,opt,name=segmant,proto3" json:"segmant,omitempty"`
}

func (x *SegmantRequest) Reset() {
	*x = SegmantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SegmantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SegmantRequest) ProtoMessage() {}

func (x *SegmantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SegmantRequest.ProtoReflect.Descriptor instead.
func (*SegmantRequest) Descriptor() ([]byte, []int) {
	return file_proto_app_proto_rawDescGZIP(), []int{3}
}

func (x *SegmantRequest) GetSegmant() string {
	if x != nil {
		return x.Segmant
	}
	return ""
}

type SegmantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success string  `protobuf:"bytes,1,opt,name=success,proto3" json:"success,omitempty"`
	Users   []*User `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *SegmantResponse) Reset() {
	*x = SegmantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_app_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SegmantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SegmantResponse) ProtoMessage() {}

func (x *SegmantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_app_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SegmantResponse.ProtoReflect.Descriptor instead.
func (*SegmantResponse) Descriptor() ([]byte, []int) {
	return file_proto_app_proto_rawDescGZIP(), []int{4}
}

func (x *SegmantResponse) GetSuccess() string {
	if x != nil {
		return x.Success
	}
	return ""
}

func (x *SegmantResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_proto_app_proto protoreflect.FileDescriptor

var file_proto_app_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x67, 0x6d, 0x61, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x67, 0x6d, 0x61, 0x6e, 0x74, 0x22, 0x42,
	0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x67, 0x6d, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x67, 0x6d,
	0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x67, 0x6d, 0x61,
	0x6e, 0x74, 0x22, 0x49, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x67, 0x6d, 0x61, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a,
	0x0e, 0x53, 0x65, 0x67, 0x6d, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x65, 0x67, 0x6d, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x65, 0x67, 0x6d, 0x61, 0x6e, 0x74, 0x22, 0x4e, 0x0a, 0x0f, 0x53, 0x65, 0x67,
	0x6d, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0xa8, 0x01, 0x0a, 0x0e, 0x53, 0x65,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x15,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x67, 0x6d, 0x61, 0x6e, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x67, 0x6d, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x67,
	0x6d, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44,
	0x0a, 0x11, 0x53, 0x68, 0x6f, 0x77, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x53, 0x65, 0x67, 0x6d,
	0x61, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x67, 0x6d,
	0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x65, 0x67, 0x6d, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_app_proto_rawDescOnce sync.Once
	file_proto_app_proto_rawDescData = file_proto_app_proto_rawDesc
)

func file_proto_app_proto_rawDescGZIP() []byte {
	file_proto_app_proto_rawDescOnce.Do(func() {
		file_proto_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_app_proto_rawDescData)
	})
	return file_proto_app_proto_rawDescData
}

var file_proto_app_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_app_proto_goTypes = []interface{}{
	(*User)(nil),                // 0: proto.User
	(*UserSegmantRequest)(nil),  // 1: proto.UserSegmantRequest
	(*UserSegmantResponse)(nil), // 2: proto.UserSegmantResponse
	(*SegmantRequest)(nil),      // 3: proto.SegmantRequest
	(*SegmantResponse)(nil),     // 4: proto.SegmantResponse
}
var file_proto_app_proto_depIdxs = []int32{
	0, // 0: proto.SegmantResponse.users:type_name -> proto.User
	1, // 1: proto.SegmentService.StoreUserSegmantation:input_type -> proto.UserSegmantRequest
	3, // 2: proto.SegmentService.ShowUserInSegmant:input_type -> proto.SegmantRequest
	2, // 3: proto.SegmentService.StoreUserSegmantation:output_type -> proto.UserSegmantResponse
	4, // 4: proto.SegmentService.ShowUserInSegmant:output_type -> proto.SegmantResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_app_proto_init() }
func file_proto_app_proto_init() {
	if File_proto_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSegmantRequest); i {
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
		file_proto_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSegmantResponse); i {
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
		file_proto_app_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SegmantRequest); i {
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
		file_proto_app_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SegmantResponse); i {
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
			RawDescriptor: file_proto_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_app_proto_goTypes,
		DependencyIndexes: file_proto_app_proto_depIdxs,
		MessageInfos:      file_proto_app_proto_msgTypes,
	}.Build()
	File_proto_app_proto = out.File
	file_proto_app_proto_rawDesc = nil
	file_proto_app_proto_goTypes = nil
	file_proto_app_proto_depIdxs = nil
}
