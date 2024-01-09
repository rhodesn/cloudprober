// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.5
// source: github.com/cloudprober/cloudprober/probes/dns/proto/config.proto

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

// DNS query types from https://en.wikipedia.org/wiki/List_of_DNS_record_types
type QueryType int32

const (
	QueryType_NONE       QueryType = 0
	QueryType_A          QueryType = 1
	QueryType_NS         QueryType = 2
	QueryType_CNAME      QueryType = 5
	QueryType_SOA        QueryType = 6
	QueryType_PTR        QueryType = 12
	QueryType_MX         QueryType = 15
	QueryType_TXT        QueryType = 16
	QueryType_RP         QueryType = 17
	QueryType_AFSDB      QueryType = 18
	QueryType_SIG        QueryType = 24
	QueryType_KEY        QueryType = 25
	QueryType_AAAA       QueryType = 28
	QueryType_LOC        QueryType = 29
	QueryType_SRV        QueryType = 33
	QueryType_NAPTR      QueryType = 35
	QueryType_KX         QueryType = 36
	QueryType_CERT       QueryType = 37
	QueryType_DNAME      QueryType = 39
	QueryType_APL        QueryType = 42
	QueryType_DS         QueryType = 43
	QueryType_SSHFP      QueryType = 44
	QueryType_IPSECKEY   QueryType = 45
	QueryType_RRSIG      QueryType = 46
	QueryType_NSEC       QueryType = 47
	QueryType_DNSKEY     QueryType = 48
	QueryType_DHCID      QueryType = 49
	QueryType_NSEC3      QueryType = 50
	QueryType_NSEC3PARAM QueryType = 51
	QueryType_TLSA       QueryType = 52
	QueryType_HIP        QueryType = 55
	QueryType_CDS        QueryType = 59
	QueryType_CDNSKEY    QueryType = 60
	QueryType_OPENPGPKEY QueryType = 61
	QueryType_TKEY       QueryType = 249
	QueryType_TSIG       QueryType = 250
	QueryType_URI        QueryType = 256
	QueryType_CAA        QueryType = 257
	QueryType_TA         QueryType = 32768
	QueryType_DLV        QueryType = 32769
)

// Enum value maps for QueryType.
var (
	QueryType_name = map[int32]string{
		0:     "NONE",
		1:     "A",
		2:     "NS",
		5:     "CNAME",
		6:     "SOA",
		12:    "PTR",
		15:    "MX",
		16:    "TXT",
		17:    "RP",
		18:    "AFSDB",
		24:    "SIG",
		25:    "KEY",
		28:    "AAAA",
		29:    "LOC",
		33:    "SRV",
		35:    "NAPTR",
		36:    "KX",
		37:    "CERT",
		39:    "DNAME",
		42:    "APL",
		43:    "DS",
		44:    "SSHFP",
		45:    "IPSECKEY",
		46:    "RRSIG",
		47:    "NSEC",
		48:    "DNSKEY",
		49:    "DHCID",
		50:    "NSEC3",
		51:    "NSEC3PARAM",
		52:    "TLSA",
		55:    "HIP",
		59:    "CDS",
		60:    "CDNSKEY",
		61:    "OPENPGPKEY",
		249:   "TKEY",
		250:   "TSIG",
		256:   "URI",
		257:   "CAA",
		32768: "TA",
		32769: "DLV",
	}
	QueryType_value = map[string]int32{
		"NONE":       0,
		"A":          1,
		"NS":         2,
		"CNAME":      5,
		"SOA":        6,
		"PTR":        12,
		"MX":         15,
		"TXT":        16,
		"RP":         17,
		"AFSDB":      18,
		"SIG":        24,
		"KEY":        25,
		"AAAA":       28,
		"LOC":        29,
		"SRV":        33,
		"NAPTR":      35,
		"KX":         36,
		"CERT":       37,
		"DNAME":      39,
		"APL":        42,
		"DS":         43,
		"SSHFP":      44,
		"IPSECKEY":   45,
		"RRSIG":      46,
		"NSEC":       47,
		"DNSKEY":     48,
		"DHCID":      49,
		"NSEC3":      50,
		"NSEC3PARAM": 51,
		"TLSA":       52,
		"HIP":        55,
		"CDS":        59,
		"CDNSKEY":    60,
		"OPENPGPKEY": 61,
		"TKEY":       249,
		"TSIG":       250,
		"URI":        256,
		"CAA":        257,
		"TA":         32768,
		"DLV":        32769,
	}
)

func (x QueryType) Enum() *QueryType {
	p := new(QueryType)
	*p = x
	return p
}

func (x QueryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QueryType) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_enumTypes[0].Descriptor()
}

func (QueryType) Type() protoreflect.EnumType {
	return &file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_enumTypes[0]
}

func (x QueryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *QueryType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = QueryType(num)
	return nil
}

// Deprecated: Use QueryType.Descriptor instead.
func (QueryType) EnumDescriptor() ([]byte, []int) {
	return file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_rawDescGZIP(), []int{0}
}

type ProbeConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Domain to use when making DNS queries
	ResolvedDomain *string `protobuf:"bytes,1,opt,name=resolved_domain,json=resolvedDomain,def=www.google.com." json:"resolved_domain,omitempty"`
	// DNS Query Type
	QueryType *QueryType `protobuf:"varint,3,opt,name=query_type,json=queryType,enum=cloudprober.probes.dns.QueryType,def=15" json:"query_type,omitempty"`
	// Minimum number of answers expected. Default behavior is to return success
	// if DNS response status is NOERROR.
	MinAnswers *uint32 `protobuf:"varint,4,opt,name=min_answers,json=minAnswers,def=0" json:"min_answers,omitempty"`
	// Whether to resolve the target (target is DNS server here) before making
	// the request. If set to false, we hand over the target directly to the DNS
	// client. Otherwise, we resolve the target first to an IP address.  By
	// default we resolve first if it's a discovered resource, e.g., a k8s
	// endpoint.
	ResolveFirst *bool `protobuf:"varint,5,opt,name=resolve_first,json=resolveFirst" json:"resolve_first,omitempty"`
	// Requests per probe.
	// Number of DNS requests per probe. Requests are executed concurrently and
	// each DNS request contributes to probe results. For example, if you run two
	// requests per probe, "total" counter will be incremented by 2.
	RequestsPerProbe *int32 `protobuf:"varint,98,opt,name=requests_per_probe,json=requestsPerProbe,def=1" json:"requests_per_probe,omitempty"`
	// How long to wait between two requests to the same target. Only relevant
	// if requests_per_probe is also configured.
	//
	// This value should be less than (interval - timeout) / requests_per_probe.
	// This is to ensure that all requests are executed within one probe interval
	// and all of them get sufficient time. For example, if probe interval is 2s,
	// timeout is 1s, and requests_per_probe is 10,  requests_interval_msec
	// should be less than 10ms.
	RequestsIntervalMsec *int32 `protobuf:"varint,99,opt,name=requests_interval_msec,json=requestsIntervalMsec,def=0" json:"requests_interval_msec,omitempty"`
}

// Default values for ProbeConf fields.
const (
	Default_ProbeConf_ResolvedDomain       = string("www.google.com.")
	Default_ProbeConf_QueryType            = QueryType_MX
	Default_ProbeConf_MinAnswers           = uint32(0)
	Default_ProbeConf_RequestsPerProbe     = int32(1)
	Default_ProbeConf_RequestsIntervalMsec = int32(0)
)

func (x *ProbeConf) Reset() {
	*x = ProbeConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeConf) ProtoMessage() {}

func (x *ProbeConf) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_msgTypes[0]
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
	return file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_rawDescGZIP(), []int{0}
}

func (x *ProbeConf) GetResolvedDomain() string {
	if x != nil && x.ResolvedDomain != nil {
		return *x.ResolvedDomain
	}
	return Default_ProbeConf_ResolvedDomain
}

func (x *ProbeConf) GetQueryType() QueryType {
	if x != nil && x.QueryType != nil {
		return *x.QueryType
	}
	return Default_ProbeConf_QueryType
}

func (x *ProbeConf) GetMinAnswers() uint32 {
	if x != nil && x.MinAnswers != nil {
		return *x.MinAnswers
	}
	return Default_ProbeConf_MinAnswers
}

func (x *ProbeConf) GetResolveFirst() bool {
	if x != nil && x.ResolveFirst != nil {
		return *x.ResolveFirst
	}
	return false
}

func (x *ProbeConf) GetRequestsPerProbe() int32 {
	if x != nil && x.RequestsPerProbe != nil {
		return *x.RequestsPerProbe
	}
	return Default_ProbeConf_RequestsPerProbe
}

func (x *ProbeConf) GetRequestsIntervalMsec() int32 {
	if x != nil && x.RequestsIntervalMsec != nil {
		return *x.RequestsIntervalMsec
	}
	return Default_ProbeConf_RequestsIntervalMsec
}

var File_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto protoreflect.FileDescriptor

var file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2f, 0x64, 0x6e, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x16, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2e, 0x64, 0x6e, 0x73, 0x22, 0xbe, 0x02, 0x0a, 0x09, 0x50,
	0x72, 0x6f, 0x62, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x38, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x64, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x3a, 0x0f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x44, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2e, 0x64, 0x6e, 0x73, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x02, 0x4d, 0x58, 0x52, 0x09, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0b, 0x6d, 0x69, 0x6e, 0x5f,
	0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x01, 0x30,
	0x52, 0x0a, 0x6d, 0x69, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x46, 0x69, 0x72, 0x73,
	0x74, 0x12, 0x2f, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x70, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x18, 0x62, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x31,
	0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x50, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x62, 0x65, 0x12, 0x37, 0x0a, 0x16, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x6d, 0x73, 0x65, 0x63, 0x18, 0x63, 0x20, 0x01,
	0x28, 0x05, 0x3a, 0x01, 0x30, 0x52, 0x14, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x4d, 0x73, 0x65, 0x63, 0x2a, 0xa4, 0x03, 0x0a, 0x09,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x41, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x4e, 0x53,
	0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x05, 0x12, 0x07, 0x0a,
	0x03, 0x53, 0x4f, 0x41, 0x10, 0x06, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x54, 0x52, 0x10, 0x0c, 0x12,
	0x06, 0x0a, 0x02, 0x4d, 0x58, 0x10, 0x0f, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x58, 0x54, 0x10, 0x10,
	0x12, 0x06, 0x0a, 0x02, 0x52, 0x50, 0x10, 0x11, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x46, 0x53, 0x44,
	0x42, 0x10, 0x12, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x49, 0x47, 0x10, 0x18, 0x12, 0x07, 0x0a, 0x03,
	0x4b, 0x45, 0x59, 0x10, 0x19, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x41, 0x41, 0x41, 0x10, 0x1c, 0x12,
	0x07, 0x0a, 0x03, 0x4c, 0x4f, 0x43, 0x10, 0x1d, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x52, 0x56, 0x10,
	0x21, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x41, 0x50, 0x54, 0x52, 0x10, 0x23, 0x12, 0x06, 0x0a, 0x02,
	0x4b, 0x58, 0x10, 0x24, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x45, 0x52, 0x54, 0x10, 0x25, 0x12, 0x09,
	0x0a, 0x05, 0x44, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x27, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x50, 0x4c,
	0x10, 0x2a, 0x12, 0x06, 0x0a, 0x02, 0x44, 0x53, 0x10, 0x2b, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x53,
	0x48, 0x46, 0x50, 0x10, 0x2c, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x50, 0x53, 0x45, 0x43, 0x4b, 0x45,
	0x59, 0x10, 0x2d, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x52, 0x53, 0x49, 0x47, 0x10, 0x2e, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x53, 0x45, 0x43, 0x10, 0x2f, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x4e, 0x53, 0x4b,
	0x45, 0x59, 0x10, 0x30, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x48, 0x43, 0x49, 0x44, 0x10, 0x31, 0x12,
	0x09, 0x0a, 0x05, 0x4e, 0x53, 0x45, 0x43, 0x33, 0x10, 0x32, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x53,
	0x45, 0x43, 0x33, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x10, 0x33, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x4c,
	0x53, 0x41, 0x10, 0x34, 0x12, 0x07, 0x0a, 0x03, 0x48, 0x49, 0x50, 0x10, 0x37, 0x12, 0x07, 0x0a,
	0x03, 0x43, 0x44, 0x53, 0x10, 0x3b, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x44, 0x4e, 0x53, 0x4b, 0x45,
	0x59, 0x10, 0x3c, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x50, 0x45, 0x4e, 0x50, 0x47, 0x50, 0x4b, 0x45,
	0x59, 0x10, 0x3d, 0x12, 0x09, 0x0a, 0x04, 0x54, 0x4b, 0x45, 0x59, 0x10, 0xf9, 0x01, 0x12, 0x09,
	0x0a, 0x04, 0x54, 0x53, 0x49, 0x47, 0x10, 0xfa, 0x01, 0x12, 0x08, 0x0a, 0x03, 0x55, 0x52, 0x49,
	0x10, 0x80, 0x02, 0x12, 0x08, 0x0a, 0x03, 0x43, 0x41, 0x41, 0x10, 0x81, 0x02, 0x12, 0x08, 0x0a,
	0x02, 0x54, 0x41, 0x10, 0x80, 0x80, 0x02, 0x12, 0x09, 0x0a, 0x03, 0x44, 0x4c, 0x56, 0x10, 0x81,
	0x80, 0x02, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2f,
	0x64, 0x6e, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_rawDescOnce sync.Once
	file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_rawDescData = file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_rawDesc
)

func file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_rawDescGZIP() []byte {
	file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_rawDescOnce.Do(func() {
		file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_rawDescData)
	})
	return file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_rawDescData
}

var file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_goTypes = []interface{}{
	(QueryType)(0),    // 0: cloudprober.probes.dns.QueryType
	(*ProbeConf)(nil), // 1: cloudprober.probes.dns.ProbeConf
}
var file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_depIdxs = []int32{
	0, // 0: cloudprober.probes.dns.ProbeConf.query_type:type_name -> cloudprober.probes.dns.QueryType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_init() }
func file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_init() {
	if File_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_goTypes,
		DependencyIndexes: file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_depIdxs,
		EnumInfos:         file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_enumTypes,
		MessageInfos:      file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_msgTypes,
	}.Build()
	File_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto = out.File
	file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_rawDesc = nil
	file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_goTypes = nil
	file_github_com_cloudprober_cloudprober_probes_dns_proto_config_proto_depIdxs = nil
}
