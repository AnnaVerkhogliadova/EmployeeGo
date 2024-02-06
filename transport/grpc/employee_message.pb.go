// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: transport/grpc/employee_message.proto

package __

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

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Sex    string `protobuf:"bytes,3,opt,name=sex,proto3" json:"sex,omitempty"`
	Age    int64  `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	Salary int64  `protobuf:"varint,5,opt,name=salary,proto3" json:"salary,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_employee_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_employee_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_transport_grpc_employee_message_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRequest) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *CreateRequest) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *CreateRequest) GetSalary() int64 {
	if x != nil {
		return x.Salary
	}
	return 0
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Sex    string `protobuf:"bytes,3,opt,name=sex,proto3" json:"sex,omitempty"`
	Age    int64  `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	Salary int64  `protobuf:"varint,5,opt,name=salary,proto3" json:"salary,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_employee_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_employee_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_transport_grpc_employee_message_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateResponse) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *CreateResponse) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *CreateResponse) GetSalary() int64 {
	if x != nil {
		return x.Salary
	}
	return 0
}

var File_transport_grpc_employee_message_proto protoreflect.FileDescriptor

var file_transport_grpc_employee_message_proto_rawDesc = []byte{
	0x0a, 0x25, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x22, 0x6f, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61,
	0x6c, 0x61, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61,
	0x72, 0x79, 0x22, 0x70, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x61,
	0x6c, 0x61, 0x72, 0x79, 0x32, 0x4e, 0x0a, 0x0f, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x17, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_transport_grpc_employee_message_proto_rawDescOnce sync.Once
	file_transport_grpc_employee_message_proto_rawDescData = file_transport_grpc_employee_message_proto_rawDesc
)

func file_transport_grpc_employee_message_proto_rawDescGZIP() []byte {
	file_transport_grpc_employee_message_proto_rawDescOnce.Do(func() {
		file_transport_grpc_employee_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_grpc_employee_message_proto_rawDescData)
	})
	return file_transport_grpc_employee_message_proto_rawDescData
}

var file_transport_grpc_employee_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transport_grpc_employee_message_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),  // 0: employee.CreateRequest
	(*CreateResponse)(nil), // 1: employee.CreateResponse
}
var file_transport_grpc_employee_message_proto_depIdxs = []int32{
	0, // 0: employee.EmployeeService.Create:input_type -> employee.CreateRequest
	1, // 1: employee.EmployeeService.Create:output_type -> employee.CreateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transport_grpc_employee_message_proto_init() }
func file_transport_grpc_employee_message_proto_init() {
	if File_transport_grpc_employee_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_grpc_employee_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_transport_grpc_employee_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
			RawDescriptor: file_transport_grpc_employee_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transport_grpc_employee_message_proto_goTypes,
		DependencyIndexes: file_transport_grpc_employee_message_proto_depIdxs,
		MessageInfos:      file_transport_grpc_employee_message_proto_msgTypes,
	}.Build()
	File_transport_grpc_employee_message_proto = out.File
	file_transport_grpc_employee_message_proto_rawDesc = nil
	file_transport_grpc_employee_message_proto_goTypes = nil
	file_transport_grpc_employee_message_proto_depIdxs = nil
}