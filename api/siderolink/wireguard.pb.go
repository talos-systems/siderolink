// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.24.4
// source: siderolink/wireguard.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PeerPacket is a packet sent between "server" and client.
type PeerPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *PeerPacket) Reset() {
	*x = PeerPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_siderolink_wireguard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerPacket) ProtoMessage() {}

func (x *PeerPacket) ProtoReflect() protoreflect.Message {
	mi := &file_siderolink_wireguard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerPacket.ProtoReflect.Descriptor instead.
func (*PeerPacket) Descriptor() ([]byte, []int) {
	return file_siderolink_wireguard_proto_rawDescGZIP(), []int{0}
}

func (x *PeerPacket) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_siderolink_wireguard_proto protoreflect.FileDescriptor

var file_siderolink_wireguard_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x77, 0x69, 0x72,
	0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x69,
	0x64, 0x65, 0x72, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x20, 0x0a, 0x0a, 0x50, 0x65, 0x65,
	0x72, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x32, 0x60, 0x0a, 0x18, 0x57,
	0x69, 0x72, 0x65, 0x47, 0x75, 0x61, 0x72, 0x64, 0x4f, 0x76, 0x65, 0x72, 0x47, 0x52, 0x50, 0x43,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x17, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f,
	0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x1a, 0x17, 0x2e, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x50,
	0x65, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x28, 0x01, 0x30, 0x01, 0x42, 0x34, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x64, 0x65,
	0x72, 0x6f, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x69, 0x6e,
	0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x69, 0x6e, 0x6b,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_siderolink_wireguard_proto_rawDescOnce sync.Once
	file_siderolink_wireguard_proto_rawDescData = file_siderolink_wireguard_proto_rawDesc
)

func file_siderolink_wireguard_proto_rawDescGZIP() []byte {
	file_siderolink_wireguard_proto_rawDescOnce.Do(func() {
		file_siderolink_wireguard_proto_rawDescData = protoimpl.X.CompressGZIP(file_siderolink_wireguard_proto_rawDescData)
	})
	return file_siderolink_wireguard_proto_rawDescData
}

var file_siderolink_wireguard_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_siderolink_wireguard_proto_goTypes = []interface{}{
	(*PeerPacket)(nil), // 0: sidero.link.PeerPacket
}
var file_siderolink_wireguard_proto_depIdxs = []int32{
	0, // 0: sidero.link.WireGuardOverGRPCService.CreateStream:input_type -> sidero.link.PeerPacket
	0, // 1: sidero.link.WireGuardOverGRPCService.CreateStream:output_type -> sidero.link.PeerPacket
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_siderolink_wireguard_proto_init() }
func file_siderolink_wireguard_proto_init() {
	if File_siderolink_wireguard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_siderolink_wireguard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerPacket); i {
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
			RawDescriptor: file_siderolink_wireguard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_siderolink_wireguard_proto_goTypes,
		DependencyIndexes: file_siderolink_wireguard_proto_depIdxs,
		MessageInfos:      file_siderolink_wireguard_proto_msgTypes,
	}.Build()
	File_siderolink_wireguard_proto = out.File
	file_siderolink_wireguard_proto_rawDesc = nil
	file_siderolink_wireguard_proto_goTypes = nil
	file_siderolink_wireguard_proto_depIdxs = nil
}