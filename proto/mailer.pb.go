// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: mailer.proto

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

type SendSMTPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateName string            `protobuf:"bytes,10101,opt,name=templateName,proto3" json:"templateName,omitempty"`
	Subject      string            `protobuf:"bytes,10102,opt,name=subject,proto3" json:"subject,omitempty"`
	To           []string          `protobuf:"bytes,10103,rep,name=to,proto3" json:"to,omitempty"`
	Data         map[string]string `protobuf:"bytes,10104,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SendSMTPRequest) Reset() {
	*x = SendSMTPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mailer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSMTPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSMTPRequest) ProtoMessage() {}

func (x *SendSMTPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mailer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSMTPRequest.ProtoReflect.Descriptor instead.
func (*SendSMTPRequest) Descriptor() ([]byte, []int) {
	return file_mailer_proto_rawDescGZIP(), []int{0}
}

func (x *SendSMTPRequest) GetTemplateName() string {
	if x != nil {
		return x.TemplateName
	}
	return ""
}

func (x *SendSMTPRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SendSMTPRequest) GetTo() []string {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *SendSMTPRequest) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

type SendSMTPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,20101,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,20102,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendSMTPResponse) Reset() {
	*x = SendSMTPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mailer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSMTPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSMTPResponse) ProtoMessage() {}

func (x *SendSMTPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mailer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSMTPResponse.ProtoReflect.Descriptor instead.
func (*SendSMTPResponse) Descriptor() ([]byte, []int) {
	return file_mailer_proto_rawDescGZIP(), []int{1}
}

func (x *SendSMTPResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SendSMTPResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_mailer_proto protoreflect.FileDescriptor

var file_mailer_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x4d,
	0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0c, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0xf5, 0x4e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19,
	0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0xf6, 0x4e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0f, 0x0a, 0x02, 0x74, 0x6f, 0x18,
	0xf7, 0x4e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0xf8, 0x4e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x4d, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x44, 0x0a, 0x10, 0x53, 0x65,
	0x6e, 0x64, 0x53, 0x4d, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x85, 0x9d, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x86, 0x9d, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0x45, 0x0a, 0x06, 0x4d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x08, 0x53, 0x65,
	0x6e, 0x64, 0x53, 0x4d, 0x54, 0x50, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x53, 0x4d, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x4d, 0x54, 0x50, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mailer_proto_rawDescOnce sync.Once
	file_mailer_proto_rawDescData = file_mailer_proto_rawDesc
)

func file_mailer_proto_rawDescGZIP() []byte {
	file_mailer_proto_rawDescOnce.Do(func() {
		file_mailer_proto_rawDescData = protoimpl.X.CompressGZIP(file_mailer_proto_rawDescData)
	})
	return file_mailer_proto_rawDescData
}

var file_mailer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mailer_proto_goTypes = []any{
	(*SendSMTPRequest)(nil),  // 0: proto.SendSMTPRequest
	(*SendSMTPResponse)(nil), // 1: proto.SendSMTPResponse
	nil,                      // 2: proto.SendSMTPRequest.DataEntry
}
var file_mailer_proto_depIdxs = []int32{
	2, // 0: proto.SendSMTPRequest.data:type_name -> proto.SendSMTPRequest.DataEntry
	0, // 1: proto.Mailer.SendSMTP:input_type -> proto.SendSMTPRequest
	1, // 2: proto.Mailer.SendSMTP:output_type -> proto.SendSMTPResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mailer_proto_init() }
func file_mailer_proto_init() {
	if File_mailer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mailer_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SendSMTPRequest); i {
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
		file_mailer_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SendSMTPResponse); i {
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
			RawDescriptor: file_mailer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mailer_proto_goTypes,
		DependencyIndexes: file_mailer_proto_depIdxs,
		MessageInfos:      file_mailer_proto_msgTypes,
	}.Build()
	File_mailer_proto = out.File
	file_mailer_proto_rawDesc = nil
	file_mailer_proto_goTypes = nil
	file_mailer_proto_depIdxs = nil
}
