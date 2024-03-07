// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.5
// source: github.com/cloudprober/cloudprober/probes/grpc/proto/config.proto

package proto

import (
	proto "github.com/cloudprober/cloudprober/internal/oauth/proto"
	proto1 "github.com/cloudprober/cloudprober/internal/tlsconfig/proto"
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

type ProbeConf_MethodType int32

const (
	ProbeConf_ECHO         ProbeConf_MethodType = 1
	ProbeConf_READ         ProbeConf_MethodType = 2
	ProbeConf_WRITE        ProbeConf_MethodType = 3
	ProbeConf_HEALTH_CHECK ProbeConf_MethodType = 4 // gRPC healthcheck service.
	ProbeConf_GENERIC      ProbeConf_MethodType = 5 // Generic gRPC request.
)

// Enum value maps for ProbeConf_MethodType.
var (
	ProbeConf_MethodType_name = map[int32]string{
		1: "ECHO",
		2: "READ",
		3: "WRITE",
		4: "HEALTH_CHECK",
		5: "GENERIC",
	}
	ProbeConf_MethodType_value = map[string]int32{
		"ECHO":         1,
		"READ":         2,
		"WRITE":        3,
		"HEALTH_CHECK": 4,
		"GENERIC":      5,
	}
)

func (x ProbeConf_MethodType) Enum() *ProbeConf_MethodType {
	p := new(ProbeConf_MethodType)
	*p = x
	return p
}

func (x ProbeConf_MethodType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProbeConf_MethodType) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_enumTypes[0].Descriptor()
}

func (ProbeConf_MethodType) Type() protoreflect.EnumType {
	return &file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_enumTypes[0]
}

func (x ProbeConf_MethodType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ProbeConf_MethodType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ProbeConf_MethodType(num)
	return nil
}

// Deprecated: Use ProbeConf_MethodType.Descriptor instead.
func (ProbeConf_MethodType) EnumDescriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_rawDescGZIP(), []int{1, 0}
}

type GenericRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Protoset contains descriptor source protos generated from the *.proto
	// files. You can use protoc to generate protoset files:
	//
	//	protoc --proto_path=. --descriptor_set_out=myservice.protoset \
	//	  --include_imports my/custom/server/service.proto
	ProtosetFile *string `protobuf:"bytes,1,opt,name=protoset_file,json=protosetFile" json:"protoset_file,omitempty"`
	// Note first 3 methods are valid only if descriptor source is not set.
	//
	// Types that are assignable to RequestType:
	//
	//	*GenericRequest_ListServices
	//	*GenericRequest_ListServiceMethods
	//	*GenericRequest_DescribeServiceMethod
	//	*GenericRequest_CallServiceMethod
	RequestType isGenericRequest_RequestType `protobuf_oneof:"request_type"`
	// Request data (in JSON format) for the call_service_method request.
	Body *string `protobuf:"bytes,6,opt,name=body" json:"body,omitempty"`
}

func (x *GenericRequest) Reset() {
	*x = GenericRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericRequest) ProtoMessage() {}

func (x *GenericRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericRequest.ProtoReflect.Descriptor instead.
func (*GenericRequest) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_rawDescGZIP(), []int{0}
}

func (x *GenericRequest) GetProtosetFile() string {
	if x != nil && x.ProtosetFile != nil {
		return *x.ProtosetFile
	}
	return ""
}

func (m *GenericRequest) GetRequestType() isGenericRequest_RequestType {
	if m != nil {
		return m.RequestType
	}
	return nil
}

func (x *GenericRequest) GetListServices() bool {
	if x, ok := x.GetRequestType().(*GenericRequest_ListServices); ok {
		return x.ListServices
	}
	return false
}

func (x *GenericRequest) GetListServiceMethods() string {
	if x, ok := x.GetRequestType().(*GenericRequest_ListServiceMethods); ok {
		return x.ListServiceMethods
	}
	return ""
}

func (x *GenericRequest) GetDescribeServiceMethod() string {
	if x, ok := x.GetRequestType().(*GenericRequest_DescribeServiceMethod); ok {
		return x.DescribeServiceMethod
	}
	return ""
}

func (x *GenericRequest) GetCallServiceMethod() string {
	if x, ok := x.GetRequestType().(*GenericRequest_CallServiceMethod); ok {
		return x.CallServiceMethod
	}
	return ""
}

func (x *GenericRequest) GetBody() string {
	if x != nil && x.Body != nil {
		return *x.Body
	}
	return ""
}

type isGenericRequest_RequestType interface {
	isGenericRequest_RequestType()
}

type GenericRequest_ListServices struct {
	// List services using reflection
	ListServices bool `protobuf:"varint,2,opt,name=list_services,json=listServices,oneof"`
}

type GenericRequest_ListServiceMethods struct {
	// List service methods using reflection.
	ListServiceMethods string `protobuf:"bytes,3,opt,name=list_service_methods,json=listServiceMethods,oneof"`
}

type GenericRequest_DescribeServiceMethod struct {
	// Describe service method using reflection.
	DescribeServiceMethod string `protobuf:"bytes,4,opt,name=describe_service_method,json=describeServiceMethod,oneof"`
}

type GenericRequest_CallServiceMethod struct {
	// Call service method. For this to succeed, you should either provide the
	// protoset file or the server should support gRPC reflection.
	// https://github.com/grpc/grpc/blob/master/doc/server-reflection.md
	CallServiceMethod string `protobuf:"bytes,5,opt,name=call_service_method,json=callServiceMethod,oneof"`
}

func (*GenericRequest_ListServices) isGenericRequest_RequestType() {}

func (*GenericRequest_ListServiceMethods) isGenericRequest_RequestType() {}

func (*GenericRequest_DescribeServiceMethod) isGenericRequest_RequestType() {}

func (*GenericRequest_CallServiceMethod) isGenericRequest_RequestType() {}

// Next tag: 14
type ProbeConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional oauth config. For GOOGLE_DEFAULT_CREDENTIALS, use:
	// oauth_config: { bearer_token { gce_service_account: "default" } }
	OauthConfig *proto.Config `protobuf:"bytes,1,opt,name=oauth_config,json=oauthConfig" json:"oauth_config,omitempty"`
	// If alts_config is provided, gRPC client uses ALTS for authentication and
	// encryption. For default alts configs, use:
	// alts_config: {}
	AltsConfig *ProbeConf_ALTSConfig `protobuf:"bytes,2,opt,name=alts_config,json=altsConfig" json:"alts_config,omitempty"`
	// If TLSConfig is specified, it's used for authentication.
	// Note that only one of ALTSConfig and TLSConfig can be enabled at a time.
	TlsConfig *proto1.TLSConfig `protobuf:"bytes,9,opt,name=tls_config,json=tlsConfig" json:"tls_config,omitempty"`
	// if insecure_transport is set to true, TLS will not be used.
	InsecureTransport *bool                 `protobuf:"varint,12,opt,name=insecure_transport,json=insecureTransport" json:"insecure_transport,omitempty"`
	Method            *ProbeConf_MethodType `protobuf:"varint,3,opt,name=method,enum=cloudprober.probes.grpc.ProbeConf_MethodType,def=1" json:"method,omitempty"`
	// Blob size for ECHO, READ, and WRITE methods.
	BlobSize *int32 `protobuf:"varint,4,opt,name=blob_size,json=blobSize,def=1024" json:"blob_size,omitempty"`
	// For HEALTH_CHECK, name of the service to health check.
	HealthCheckService *string `protobuf:"bytes,10,opt,name=health_check_service,json=healthCheckService" json:"health_check_service,omitempty"`
	// For HEALTH_CHECK, ignore status. By default, HEALTH_CHECK test passes
	// only if response-status is SERVING. Setting the following option makes
	// HEALTH_CHECK pass regardless of the response-status.
	HealthCheckIgnoreStatus *bool `protobuf:"varint,11,opt,name=health_check_ignore_status,json=healthCheckIgnoreStatus" json:"health_check_ignore_status,omitempty"`
	// Request definition for the GENERIC method.
	Request   *GenericRequest `protobuf:"bytes,14,opt,name=request" json:"request,omitempty"`
	NumConns  *int32          `protobuf:"varint,5,opt,name=num_conns,json=numConns,def=2" json:"num_conns,omitempty"`
	KeepAlive *bool           `protobuf:"varint,6,opt,name=keep_alive,json=keepAlive,def=1" json:"keep_alive,omitempty"`
	// If connect_timeout is not specified, reuse probe timeout.
	ConnectTimeoutMsec *int32 `protobuf:"varint,7,opt,name=connect_timeout_msec,json=connectTimeoutMsec" json:"connect_timeout_msec,omitempty"`
	// URI scheme allows gRPC to use different resolvers
	// Example URI scheme: "google-c2p:///"
	// See https://github.com/grpc/grpc/blob/master/doc/naming.md for more details
	UriScheme *string             `protobuf:"bytes,8,opt,name=uri_scheme,json=uriScheme" json:"uri_scheme,omitempty"`
	Headers   []*ProbeConf_Header `protobuf:"bytes,13,rep,name=headers" json:"headers,omitempty"`
}

// Default values for ProbeConf fields.
const (
	Default_ProbeConf_Method    = ProbeConf_ECHO
	Default_ProbeConf_BlobSize  = int32(1024)
	Default_ProbeConf_NumConns  = int32(2)
	Default_ProbeConf_KeepAlive = bool(true)
)

func (x *ProbeConf) Reset() {
	*x = ProbeConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeConf) ProtoMessage() {}

func (x *ProbeConf) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeConf.ProtoReflect.Descriptor instead.
func (*ProbeConf) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_rawDescGZIP(), []int{1}
}

func (x *ProbeConf) GetOauthConfig() *proto.Config {
	if x != nil {
		return x.OauthConfig
	}
	return nil
}

func (x *ProbeConf) GetAltsConfig() *ProbeConf_ALTSConfig {
	if x != nil {
		return x.AltsConfig
	}
	return nil
}

func (x *ProbeConf) GetTlsConfig() *proto1.TLSConfig {
	if x != nil {
		return x.TlsConfig
	}
	return nil
}

func (x *ProbeConf) GetInsecureTransport() bool {
	if x != nil && x.InsecureTransport != nil {
		return *x.InsecureTransport
	}
	return false
}

func (x *ProbeConf) GetMethod() ProbeConf_MethodType {
	if x != nil && x.Method != nil {
		return *x.Method
	}
	return Default_ProbeConf_Method
}

func (x *ProbeConf) GetBlobSize() int32 {
	if x != nil && x.BlobSize != nil {
		return *x.BlobSize
	}
	return Default_ProbeConf_BlobSize
}

func (x *ProbeConf) GetHealthCheckService() string {
	if x != nil && x.HealthCheckService != nil {
		return *x.HealthCheckService
	}
	return ""
}

func (x *ProbeConf) GetHealthCheckIgnoreStatus() bool {
	if x != nil && x.HealthCheckIgnoreStatus != nil {
		return *x.HealthCheckIgnoreStatus
	}
	return false
}

func (x *ProbeConf) GetRequest() *GenericRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *ProbeConf) GetNumConns() int32 {
	if x != nil && x.NumConns != nil {
		return *x.NumConns
	}
	return Default_ProbeConf_NumConns
}

func (x *ProbeConf) GetKeepAlive() bool {
	if x != nil && x.KeepAlive != nil {
		return *x.KeepAlive
	}
	return Default_ProbeConf_KeepAlive
}

func (x *ProbeConf) GetConnectTimeoutMsec() int32 {
	if x != nil && x.ConnectTimeoutMsec != nil {
		return *x.ConnectTimeoutMsec
	}
	return 0
}

func (x *ProbeConf) GetUriScheme() string {
	if x != nil && x.UriScheme != nil {
		return *x.UriScheme
	}
	return ""
}

func (x *ProbeConf) GetHeaders() []*ProbeConf_Header {
	if x != nil {
		return x.Headers
	}
	return nil
}

// ALTS is a gRPC security method supported by some Google services.
// If enabled, peers, with the help of a handshaker service (e.g. metadata
// server of GCE instances), use credentials attached to the service accounts
// to authenticate each other. See
// https://cloud.google.com/security/encryption-in-transit/#service_integrity_encryption
// for more details.
type ProbeConf_ALTSConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If provided, ALTS verifies that peer is using one of the given service
	// accounts.
	TargetServiceAccount []string `protobuf:"bytes,1,rep,name=target_service_account,json=targetServiceAccount" json:"target_service_account,omitempty"`
	// Handshaker service address. Default is to use the local metadata server.
	// For most of the ALTS use cases, default address should be okay.
	HandshakerServiceAddress *string `protobuf:"bytes,2,opt,name=handshaker_service_address,json=handshakerServiceAddress" json:"handshaker_service_address,omitempty"`
}

func (x *ProbeConf_ALTSConfig) Reset() {
	*x = ProbeConf_ALTSConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeConf_ALTSConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeConf_ALTSConfig) ProtoMessage() {}

func (x *ProbeConf_ALTSConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeConf_ALTSConfig.ProtoReflect.Descriptor instead.
func (*ProbeConf_ALTSConfig) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ProbeConf_ALTSConfig) GetTargetServiceAccount() []string {
	if x != nil {
		return x.TargetServiceAccount
	}
	return nil
}

func (x *ProbeConf_ALTSConfig) GetHandshakerServiceAddress() string {
	if x != nil && x.HandshakerServiceAddress != nil {
		return *x.HandshakerServiceAddress
	}
	return ""
}

type ProbeConf_Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (x *ProbeConf_Header) Reset() {
	*x = ProbeConf_Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeConf_Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeConf_Header) ProtoMessage() {}

func (x *ProbeConf_Header) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeConf_Header.ProtoReflect.Descriptor instead.
func (*ProbeConf_Header) Descriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_rawDescGZIP(), []int{1, 1}
}

func (x *ProbeConf_Header) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ProbeConf_Header) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

var File_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto protoreflect.FileDescriptor

var file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x44, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x74, 0x6c, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x02, 0x0a,
	0x0e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x65, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x65, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x25, 0x0a, 0x0d, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0c, 0x6c,
	0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x14, 0x6c,
	0x69, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x12, 0x6c, 0x69, 0x73,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12,
	0x38, 0x0a, 0x17, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x15, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x63, 0x61, 0x6c,
	0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x11, 0x63, 0x61, 0x6c, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x42,
	0x0e, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x89, 0x08, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x3c, 0x0a,
	0x0c, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65,
	0x72, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4e, 0x0a, 0x0b, 0x61,
	0x6c, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x62, 0x65, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x2e, 0x41, 0x4c, 0x54, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x0a, 0x61, 0x6c, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3f, 0x0a, 0x0a, 0x74,
	0x6c, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x74, 0x6c,
	0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x4c, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x09, 0x74, 0x6c, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2d, 0x0a, 0x12,
	0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x4b, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x04, 0x45, 0x43, 0x48, 0x4f,
	0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x21, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x62,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x04, 0x31, 0x30, 0x32,
	0x34, 0x52, 0x08, 0x62, 0x6c, 0x6f, 0x62, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a,
	0x1a, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x67,
	0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x17, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x67,
	0x6e, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x41, 0x0a, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x3a, 0x01, 0x32, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x43, 0x6f, 0x6e, 0x6e, 0x73, 0x12, 0x23, 0x0a,
	0x0a, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x3a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69,
	0x76, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x73, 0x65, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x12, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x4d, 0x73, 0x65, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x72, 0x69, 0x5f, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x72, 0x69, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x0d,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50,
	0x72, 0x6f, 0x62, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x80, 0x01, 0x0a, 0x0a, 0x41, 0x4c, 0x54,
	0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x34, 0x0a, 0x16, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x14, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3c, 0x0a,
	0x1a, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x18, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x32, 0x0a, 0x06, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x4a, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a,
	0x04, 0x45, 0x43, 0x48, 0x4f, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x41, 0x44, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c,
	0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x10, 0x04, 0x12, 0x0b,
	0x0a, 0x07, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x49, 0x43, 0x10, 0x05, 0x42, 0x36, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70,
	0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f,
}

var (
	file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_rawDescOnce sync.Once
	file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_rawDescData = file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_rawDesc
)

func file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_rawDescGZIP() []byte {
	file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_rawDescOnce.Do(func() {
		file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_rawDescData)
	})
	return file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_rawDescData
}

var file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_goTypes = []interface{}{
	(ProbeConf_MethodType)(0),    // 0: cloudprober.probes.grpc.ProbeConf.MethodType
	(*GenericRequest)(nil),       // 1: cloudprober.probes.grpc.GenericRequest
	(*ProbeConf)(nil),            // 2: cloudprober.probes.grpc.ProbeConf
	(*ProbeConf_ALTSConfig)(nil), // 3: cloudprober.probes.grpc.ProbeConf.ALTSConfig
	(*ProbeConf_Header)(nil),     // 4: cloudprober.probes.grpc.ProbeConf.Header
	(*proto.Config)(nil),         // 5: cloudprober.oauth.Config
	(*proto1.TLSConfig)(nil),     // 6: cloudprober.tlsconfig.TLSConfig
}
var file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_depIdxs = []int32{
	5, // 0: cloudprober.probes.grpc.ProbeConf.oauth_config:type_name -> cloudprober.oauth.Config
	3, // 1: cloudprober.probes.grpc.ProbeConf.alts_config:type_name -> cloudprober.probes.grpc.ProbeConf.ALTSConfig
	6, // 2: cloudprober.probes.grpc.ProbeConf.tls_config:type_name -> cloudprober.tlsconfig.TLSConfig
	0, // 3: cloudprober.probes.grpc.ProbeConf.method:type_name -> cloudprober.probes.grpc.ProbeConf.MethodType
	1, // 4: cloudprober.probes.grpc.ProbeConf.request:type_name -> cloudprober.probes.grpc.GenericRequest
	4, // 5: cloudprober.probes.grpc.ProbeConf.headers:type_name -> cloudprober.probes.grpc.ProbeConf.Header
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_init() }
func file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_init() {
	if File_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericRequest); i {
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
		file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeConf); i {
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
		file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeConf_ALTSConfig); i {
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
		file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeConf_Header); i {
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
	file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*GenericRequest_ListServices)(nil),
		(*GenericRequest_ListServiceMethods)(nil),
		(*GenericRequest_DescribeServiceMethod)(nil),
		(*GenericRequest_CallServiceMethod)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_goTypes,
		DependencyIndexes: file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_depIdxs,
		EnumInfos:         file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_enumTypes,
		MessageInfos:      file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_msgTypes,
	}.Build()
	File_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto = out.File
	file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_rawDesc = nil
	file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_goTypes = nil
	file_github_com_cloudprober_cloudprober_probes_grpc_proto_config_proto_depIdxs = nil
}
