package subscription

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SubscriptionServer struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	ServerMetadata map[string]string      `protobuf:"bytes,2,rep,name=serverMetadata,proto3" json:"serverMetadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Tag            string                 `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *SubscriptionServer) Reset() {
	*x = SubscriptionServer{}
	mi := &file_app_subscription_subscription_rpc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscriptionServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionServer) ProtoMessage() {}

func (x *SubscriptionServer) ProtoReflect() protoreflect.Message {
	mi := &file_app_subscription_subscription_rpc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionServer.ProtoReflect.Descriptor instead.
func (*SubscriptionServer) Descriptor() ([]byte, []int) {
	return file_app_subscription_subscription_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *SubscriptionServer) GetServerMetadata() map[string]string {
	if x != nil {
		return x.ServerMetadata
	}
	return nil
}

func (x *SubscriptionServer) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type TrackedSubscriptionStatus struct {
	state            protoimpl.MessageState         `protogen:"open.v1"`
	Servers          map[string]*SubscriptionServer `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	DocumentMetadata map[string]string              `protobuf:"bytes,2,rep,name=documentMetadata,proto3" json:"documentMetadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ImportSource     *ImportSource                  `protobuf:"bytes,3,opt,name=importSource,proto3" json:"importSource,omitempty"`
	AddedByApi       bool                           `protobuf:"varint,4,opt,name=added_by_api,json=addedByApi,proto3" json:"added_by_api,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *TrackedSubscriptionStatus) Reset() {
	*x = TrackedSubscriptionStatus{}
	mi := &file_app_subscription_subscription_rpc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrackedSubscriptionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackedSubscriptionStatus) ProtoMessage() {}

func (x *TrackedSubscriptionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_app_subscription_subscription_rpc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackedSubscriptionStatus.ProtoReflect.Descriptor instead.
func (*TrackedSubscriptionStatus) Descriptor() ([]byte, []int) {
	return file_app_subscription_subscription_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *TrackedSubscriptionStatus) GetServers() map[string]*SubscriptionServer {
	if x != nil {
		return x.Servers
	}
	return nil
}

func (x *TrackedSubscriptionStatus) GetDocumentMetadata() map[string]string {
	if x != nil {
		return x.DocumentMetadata
	}
	return nil
}

func (x *TrackedSubscriptionStatus) GetImportSource() *ImportSource {
	if x != nil {
		return x.ImportSource
	}
	return nil
}

func (x *TrackedSubscriptionStatus) GetAddedByApi() bool {
	if x != nil {
		return x.AddedByApi
	}
	return false
}

var File_app_subscription_subscription_rpc_proto protoreflect.FileDescriptor

var file_app_subscription_subscription_rpc_proto_rawDesc = string([]byte{
	0x0a, 0x27, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x76, 0x32, 0x72, 0x61, 0x79,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1d, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x01, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x6b, 0x0a, 0x0e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x1a, 0x41, 0x0a, 0x13, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x97,
	0x04, 0x0a, 0x19, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x5d, 0x0a, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e,
	0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x64, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x78, 0x0a, 0x10, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4c, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x10, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x4d, 0x0a, 0x0c, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x76, 0x32,
	0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0c, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x65, 0x64, 0x5f, 0x62, 0x79,
	0x5f, 0x61, 0x70, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x64, 0x64, 0x65,
	0x64, 0x42, 0x79, 0x41, 0x70, 0x69, 0x1a, 0x6b, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x45, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x43, 0x0a, 0x15, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x72, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e,
	0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x01, 0x5a, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x32, 0x66, 0x6c, 0x79, 0x2f,
	0x76, 0x32, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x35, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0xaa, 0x02,
	0x1b, 0x56, 0x32, 0x52, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_app_subscription_subscription_rpc_proto_rawDescOnce sync.Once
	file_app_subscription_subscription_rpc_proto_rawDescData []byte
)

func file_app_subscription_subscription_rpc_proto_rawDescGZIP() []byte {
	file_app_subscription_subscription_rpc_proto_rawDescOnce.Do(func() {
		file_app_subscription_subscription_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_app_subscription_subscription_rpc_proto_rawDesc), len(file_app_subscription_subscription_rpc_proto_rawDesc)))
	})
	return file_app_subscription_subscription_rpc_proto_rawDescData
}

var file_app_subscription_subscription_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_app_subscription_subscription_rpc_proto_goTypes = []any{
	(*SubscriptionServer)(nil),        // 0: v2ray.core.app.subscription.SubscriptionServer
	(*TrackedSubscriptionStatus)(nil), // 1: v2ray.core.app.subscription.TrackedSubscriptionStatus
	nil,                               // 2: v2ray.core.app.subscription.SubscriptionServer.ServerMetadataEntry
	nil,                               // 3: v2ray.core.app.subscription.TrackedSubscriptionStatus.ServersEntry
	nil,                               // 4: v2ray.core.app.subscription.TrackedSubscriptionStatus.DocumentMetadataEntry
	(*ImportSource)(nil),              // 5: v2ray.core.app.subscription.ImportSource
}
var file_app_subscription_subscription_rpc_proto_depIdxs = []int32{
	2, // 0: v2ray.core.app.subscription.SubscriptionServer.serverMetadata:type_name -> v2ray.core.app.subscription.SubscriptionServer.ServerMetadataEntry
	3, // 1: v2ray.core.app.subscription.TrackedSubscriptionStatus.servers:type_name -> v2ray.core.app.subscription.TrackedSubscriptionStatus.ServersEntry
	4, // 2: v2ray.core.app.subscription.TrackedSubscriptionStatus.documentMetadata:type_name -> v2ray.core.app.subscription.TrackedSubscriptionStatus.DocumentMetadataEntry
	5, // 3: v2ray.core.app.subscription.TrackedSubscriptionStatus.importSource:type_name -> v2ray.core.app.subscription.ImportSource
	0, // 4: v2ray.core.app.subscription.TrackedSubscriptionStatus.ServersEntry.value:type_name -> v2ray.core.app.subscription.SubscriptionServer
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_app_subscription_subscription_rpc_proto_init() }
func file_app_subscription_subscription_rpc_proto_init() {
	if File_app_subscription_subscription_rpc_proto != nil {
		return
	}
	file_app_subscription_config_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_app_subscription_subscription_rpc_proto_rawDesc), len(file_app_subscription_subscription_rpc_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_subscription_subscription_rpc_proto_goTypes,
		DependencyIndexes: file_app_subscription_subscription_rpc_proto_depIdxs,
		MessageInfos:      file_app_subscription_subscription_rpc_proto_msgTypes,
	}.Build()
	File_app_subscription_subscription_rpc_proto = out.File
	file_app_subscription_subscription_rpc_proto_goTypes = nil
	file_app_subscription_subscription_rpc_proto_depIdxs = nil
}
