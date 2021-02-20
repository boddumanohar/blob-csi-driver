// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.1
// source: azure_blob_mount.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MountAzureBlobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetPath    string `protobuf:"bytes,1,opt,name=targetPath,proto3" json:"targetPath,omitempty"`
	AccountName   string `protobuf:"bytes,2,opt,name=accountName,proto3" json:"accountName,omitempty"`
	ContainerName string `protobuf:"bytes,3,opt,name=containerName,proto3" json:"containerName,omitempty"`
	AccountKey    string `protobuf:"bytes,4,opt,name=accountKey,proto3" json:"accountKey,omitempty"`
	TmpPath       string `protobuf:"bytes,5,opt,name=tmpPath,proto3" json:"tmpPath,omitempty"`
}

func (x *MountAzureBlobRequest) Reset() {
	*x = MountAzureBlobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_azure_blob_mount_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountAzureBlobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountAzureBlobRequest) ProtoMessage() {}

func (x *MountAzureBlobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_azure_blob_mount_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MountAzureBlobRequest.ProtoReflect.Descriptor instead.
func (*MountAzureBlobRequest) Descriptor() ([]byte, []int) {
	return file_azure_blob_mount_proto_rawDescGZIP(), []int{0}
}

func (x *MountAzureBlobRequest) GetTargetPath() string {
	if x != nil {
		return x.TargetPath
	}
	return ""
}

func (x *MountAzureBlobRequest) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *MountAzureBlobRequest) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *MountAzureBlobRequest) GetAccountKey() string {
	if x != nil {
		return x.AccountKey
	}
	return ""
}

func (x *MountAzureBlobRequest) GetTmpPath() string {
	if x != nil {
		return x.TmpPath
	}
	return ""
}

type MountAzureBlobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err string `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *MountAzureBlobResponse) Reset() {
	*x = MountAzureBlobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_azure_blob_mount_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MountAzureBlobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MountAzureBlobResponse) ProtoMessage() {}

func (x *MountAzureBlobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_azure_blob_mount_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MountAzureBlobResponse.ProtoReflect.Descriptor instead.
func (*MountAzureBlobResponse) Descriptor() ([]byte, []int) {
	return file_azure_blob_mount_proto_rawDescGZIP(), []int{1}
}

func (x *MountAzureBlobResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_azure_blob_mount_proto protoreflect.FileDescriptor

var file_azure_blob_mount_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x5f, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x15, 0x4d, 0x6f, 0x75,
	0x6e, 0x74, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x6d,
	0x70, 0x50, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x6d, 0x70,
	0x50, 0x61, 0x74, 0x68, 0x22, 0x2a, 0x0a, 0x16, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x7a, 0x75,
	0x72, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72,
	0x32, 0x53, 0x0a, 0x0c, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x43, 0x0a, 0x0e, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x42, 0x6c,
	0x6f, 0x62, 0x12, 0x16, 0x2e, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x42,
	0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x4d, 0x6f, 0x75,
	0x6e, 0x74, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_azure_blob_mount_proto_rawDescOnce sync.Once
	file_azure_blob_mount_proto_rawDescData = file_azure_blob_mount_proto_rawDesc
)

func file_azure_blob_mount_proto_rawDescGZIP() []byte {
	file_azure_blob_mount_proto_rawDescOnce.Do(func() {
		file_azure_blob_mount_proto_rawDescData = protoimpl.X.CompressGZIP(file_azure_blob_mount_proto_rawDescData)
	})
	return file_azure_blob_mount_proto_rawDescData
}

var file_azure_blob_mount_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_azure_blob_mount_proto_goTypes = []interface{}{
	(*MountAzureBlobRequest)(nil),  // 0: MountAzureBlobRequest
	(*MountAzureBlobResponse)(nil), // 1: MountAzureBlobResponse
}
var file_azure_blob_mount_proto_depIdxs = []int32{
	0, // 0: MountService.MountAzureBlob:input_type -> MountAzureBlobRequest
	1, // 1: MountService.MountAzureBlob:output_type -> MountAzureBlobResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_azure_blob_mount_proto_init() }
func file_azure_blob_mount_proto_init() {
	if File_azure_blob_mount_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_azure_blob_mount_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MountAzureBlobRequest); i {
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
		file_azure_blob_mount_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MountAzureBlobResponse); i {
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
			RawDescriptor: file_azure_blob_mount_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_azure_blob_mount_proto_goTypes,
		DependencyIndexes: file_azure_blob_mount_proto_depIdxs,
		MessageInfos:      file_azure_blob_mount_proto_msgTypes,
	}.Build()
	File_azure_blob_mount_proto = out.File
	file_azure_blob_mount_proto_rawDesc = nil
	file_azure_blob_mount_proto_goTypes = nil
	file_azure_blob_mount_proto_depIdxs = nil
}
