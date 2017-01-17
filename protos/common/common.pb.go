// Code generated by protoc-gen-go.
// source: common/common.proto
// DO NOT EDIT!

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common/common.proto
	common/configuration.proto
	common/msp_principal.proto

It has these top-level messages:
	LastConfiguration
	Metadata
	MetadataSignature
	Header
	ChainHeader
	SignatureHeader
	Payload
	Envelope
	Block
	BlockHeader
	BlockData
	BlockMetadata
	ConfigurationEnvelope
	ConfigurationTemplate
	SignedConfigurationItem
	ConfigurationItem
	ConfigurationSignature
	Policy
	SignaturePolicyEnvelope
	SignaturePolicy
	HashingAlgorithm
	BlockDataHashingStructure
	MSPPrincipal
	OrganizationUnit
	MSPRole
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// These status codes are intended to resemble selected HTTP status codes
type Status int32

const (
	Status_UNKNOWN                  Status = 0
	Status_SUCCESS                  Status = 200
	Status_BAD_REQUEST              Status = 400
	Status_FORBIDDEN                Status = 403
	Status_NOT_FOUND                Status = 404
	Status_REQUEST_ENTITY_TOO_LARGE Status = 413
	Status_INTERNAL_SERVER_ERROR    Status = 500
	Status_SERVICE_UNAVAILABLE      Status = 503
)

var Status_name = map[int32]string{
	0:   "UNKNOWN",
	200: "SUCCESS",
	400: "BAD_REQUEST",
	403: "FORBIDDEN",
	404: "NOT_FOUND",
	413: "REQUEST_ENTITY_TOO_LARGE",
	500: "INTERNAL_SERVER_ERROR",
	503: "SERVICE_UNAVAILABLE",
}
var Status_value = map[string]int32{
	"UNKNOWN":                  0,
	"SUCCESS":                  200,
	"BAD_REQUEST":              400,
	"FORBIDDEN":                403,
	"NOT_FOUND":                404,
	"REQUEST_ENTITY_TOO_LARGE": 413,
	"INTERNAL_SERVER_ERROR":    500,
	"SERVICE_UNAVAILABLE":      503,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HeaderType int32

const (
	HeaderType_MESSAGE                   HeaderType = 0
	HeaderType_CONFIGURATION_TRANSACTION HeaderType = 1
	HeaderType_CONFIGURATION_ITEM        HeaderType = 2
	HeaderType_ENDORSER_TRANSACTION      HeaderType = 3
	HeaderType_ORDERER_TRANSACTION       HeaderType = 4
	HeaderType_DELIVER_SEEK_INFO         HeaderType = 5
)

var HeaderType_name = map[int32]string{
	0: "MESSAGE",
	1: "CONFIGURATION_TRANSACTION",
	2: "CONFIGURATION_ITEM",
	3: "ENDORSER_TRANSACTION",
	4: "ORDERER_TRANSACTION",
	5: "DELIVER_SEEK_INFO",
}
var HeaderType_value = map[string]int32{
	"MESSAGE":                   0,
	"CONFIGURATION_TRANSACTION": 1,
	"CONFIGURATION_ITEM":        2,
	"ENDORSER_TRANSACTION":      3,
	"ORDERER_TRANSACTION":       4,
	"DELIVER_SEEK_INFO":         5,
}

func (x HeaderType) String() string {
	return proto.EnumName(HeaderType_name, int32(x))
}
func (HeaderType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// This enum enlist indexes of the block metadata array
type BlockMetadataIndex int32

const (
	BlockMetadataIndex_SIGNATURES          BlockMetadataIndex = 0
	BlockMetadataIndex_LAST_CONFIGURATION  BlockMetadataIndex = 1
	BlockMetadataIndex_TRANSACTIONS_FILTER BlockMetadataIndex = 2
)

var BlockMetadataIndex_name = map[int32]string{
	0: "SIGNATURES",
	1: "LAST_CONFIGURATION",
	2: "TRANSACTIONS_FILTER",
}
var BlockMetadataIndex_value = map[string]int32{
	"SIGNATURES":          0,
	"LAST_CONFIGURATION":  1,
	"TRANSACTIONS_FILTER": 2,
}

func (x BlockMetadataIndex) String() string {
	return proto.EnumName(BlockMetadataIndex_name, int32(x))
}
func (BlockMetadataIndex) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// LastConfiguration is the encoded value for the Metadata message which is encoded in the LAST_CONFIGURATION block metadata index
type LastConfiguration struct {
	Index uint64 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
}

func (m *LastConfiguration) Reset()                    { *m = LastConfiguration{} }
func (m *LastConfiguration) String() string            { return proto.CompactTextString(m) }
func (*LastConfiguration) ProtoMessage()               {}
func (*LastConfiguration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Metadata is a common structure to be used to encode block metadata
type Metadata struct {
	Value      []byte               `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Signatures []*MetadataSignature `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Metadata) GetSignatures() []*MetadataSignature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

type MetadataSignature struct {
	SignatureHeader []byte `protobuf:"bytes,1,opt,name=signatureHeader,proto3" json:"signatureHeader,omitempty"`
	Signature       []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *MetadataSignature) Reset()                    { *m = MetadataSignature{} }
func (m *MetadataSignature) String() string            { return proto.CompactTextString(m) }
func (*MetadataSignature) ProtoMessage()               {}
func (*MetadataSignature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Header struct {
	ChainHeader     *ChainHeader     `protobuf:"bytes,1,opt,name=chainHeader" json:"chainHeader,omitempty"`
	SignatureHeader *SignatureHeader `protobuf:"bytes,2,opt,name=signatureHeader" json:"signatureHeader,omitempty"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Header) GetChainHeader() *ChainHeader {
	if m != nil {
		return m.ChainHeader
	}
	return nil
}

func (m *Header) GetSignatureHeader() *SignatureHeader {
	if m != nil {
		return m.SignatureHeader
	}
	return nil
}

// Header is a generic replay prevention and identity message to include in a signed payload
type ChainHeader struct {
	Type int32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	// Version indicates message protocol version
	Version int32 `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	// Timestamp is the local time when the message was created
	// by the sender
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	// Identifier of the chain this message is bound for
	ChainID string `protobuf:"bytes,4,opt,name=chainID" json:"chainID,omitempty"`
	// An unique identifier that is used end-to-end.
	//  -  set by higher layers such as end user or SDK
	//  -  passed to the endorser (which will check for uniqueness)
	//  -  as the header is passed along unchanged, it will be
	//     be retrieved by the committer (uniqueness check here as well)
	//  -  to be stored in the ledger
	TxID string `protobuf:"bytes,5,opt,name=txID" json:"txID,omitempty"`
	// The epoch in which this header was generated, where epoch is defined based on block height
	// Epoch in which the response has been generated. This field identifies a
	// logical window of time. A proposal response is accepted by a peer only if
	// two conditions hold:
	// 1. the epoch specified in the message is the current epoch
	// 2. this message has been only seen once during this epoch (i.e. it hasn't
	//    been replayed)
	Epoch uint64 `protobuf:"varint,6,opt,name=epoch" json:"epoch,omitempty"`
	// Extension that may be attached based on the header type
	Extension []byte `protobuf:"bytes,7,opt,name=extension,proto3" json:"extension,omitempty"`
}

func (m *ChainHeader) Reset()                    { *m = ChainHeader{} }
func (m *ChainHeader) String() string            { return proto.CompactTextString(m) }
func (*ChainHeader) ProtoMessage()               {}
func (*ChainHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ChainHeader) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type SignatureHeader struct {
	// Creator of the message, specified as a certificate chain
	Creator []byte `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// Arbitrary number that may only be used once. Can be used to detect replay attacks.
	Nonce []byte `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *SignatureHeader) Reset()                    { *m = SignatureHeader{} }
func (m *SignatureHeader) String() string            { return proto.CompactTextString(m) }
func (*SignatureHeader) ProtoMessage()               {}
func (*SignatureHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// Payload is the message contents (and header to allow for signing)
type Payload struct {
	// Header is included to provide identity and prevent replay
	Header *Header `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// Data, the encoding of which is defined by the type in the header
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Payload) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

// Envelope wraps a Payload with a signature so that the message may be authenticated
type Envelope struct {
	// A marshaled Payload
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// A signature by the creator specified in the Payload header
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// This is finalized block structure to be shared among the orderer and peer
// Note that the BlockHeader chains to the previous BlockHeader, and the BlockData hash is embedded
// in the BlockHeader.  This makes it natural and obvious that the Data is included in the hash, but
// the Metadata is not.
type Block struct {
	Header   *BlockHeader   `protobuf:"bytes,1,opt,name=Header" json:"Header,omitempty"`
	Data     *BlockData     `protobuf:"bytes,2,opt,name=Data" json:"Data,omitempty"`
	Metadata *BlockMetadata `protobuf:"bytes,3,opt,name=Metadata" json:"Metadata,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetData() *BlockData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Block) GetMetadata() *BlockMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type BlockHeader struct {
	Number       uint64 `protobuf:"varint,1,opt,name=Number" json:"Number,omitempty"`
	PreviousHash []byte `protobuf:"bytes,2,opt,name=PreviousHash,proto3" json:"PreviousHash,omitempty"`
	DataHash     []byte `protobuf:"bytes,3,opt,name=DataHash,proto3" json:"DataHash,omitempty"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type BlockData struct {
	Data [][]byte `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (m *BlockData) Reset()                    { *m = BlockData{} }
func (m *BlockData) String() string            { return proto.CompactTextString(m) }
func (*BlockData) ProtoMessage()               {}
func (*BlockData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type BlockMetadata struct {
	Metadata [][]byte `protobuf:"bytes,1,rep,name=Metadata,proto3" json:"Metadata,omitempty"`
}

func (m *BlockMetadata) Reset()                    { *m = BlockMetadata{} }
func (m *BlockMetadata) String() string            { return proto.CompactTextString(m) }
func (*BlockMetadata) ProtoMessage()               {}
func (*BlockMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func init() {
	proto.RegisterType((*LastConfiguration)(nil), "common.LastConfiguration")
	proto.RegisterType((*Metadata)(nil), "common.Metadata")
	proto.RegisterType((*MetadataSignature)(nil), "common.MetadataSignature")
	proto.RegisterType((*Header)(nil), "common.Header")
	proto.RegisterType((*ChainHeader)(nil), "common.ChainHeader")
	proto.RegisterType((*SignatureHeader)(nil), "common.SignatureHeader")
	proto.RegisterType((*Payload)(nil), "common.Payload")
	proto.RegisterType((*Envelope)(nil), "common.Envelope")
	proto.RegisterType((*Block)(nil), "common.Block")
	proto.RegisterType((*BlockHeader)(nil), "common.BlockHeader")
	proto.RegisterType((*BlockData)(nil), "common.BlockData")
	proto.RegisterType((*BlockMetadata)(nil), "common.BlockMetadata")
	proto.RegisterEnum("common.Status", Status_name, Status_value)
	proto.RegisterEnum("common.HeaderType", HeaderType_name, HeaderType_value)
	proto.RegisterEnum("common.BlockMetadataIndex", BlockMetadataIndex_name, BlockMetadataIndex_value)
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 866 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xae, 0xe3, 0xfc, 0x34, 0x27, 0xa5, 0xeb, 0x4e, 0xb7, 0xbb, 0xde, 0x8a, 0xd5, 0x56, 0x96,
	0x40, 0xa5, 0x15, 0x89, 0x28, 0x42, 0x82, 0x4b, 0x27, 0x9e, 0x74, 0xad, 0x4d, 0xed, 0x65, 0xc6,
	0x59, 0x04, 0xbd, 0xb0, 0x9c, 0x64, 0x9a, 0x58, 0x24, 0x76, 0x64, 0x3b, 0x55, 0x7b, 0xcb, 0x03,
	0x20, 0x24, 0xb8, 0x42, 0xe2, 0x05, 0x78, 0x12, 0xde, 0x80, 0x17, 0x41, 0xe2, 0x16, 0xcd, 0xd8,
	0xe3, 0xc6, 0xed, 0x05, 0x57, 0x9d, 0xef, 0x7c, 0xdf, 0x9c, 0xf3, 0x9d, 0x73, 0x26, 0x35, 0x1c,
	0x4e, 0xe3, 0xd5, 0x2a, 0x8e, 0x7a, 0xf9, 0x9f, 0xee, 0x3a, 0x89, 0xb3, 0x18, 0x35, 0x73, 0x74,
	0xfc, 0x66, 0x1e, 0xc7, 0xf3, 0x25, 0xeb, 0x89, 0xe8, 0x64, 0x73, 0xd3, 0xcb, 0xc2, 0x15, 0x4b,
	0xb3, 0x60, 0xb5, 0xce, 0x85, 0xc6, 0x67, 0x70, 0x30, 0x0a, 0xd2, 0x6c, 0x10, 0x47, 0x37, 0xe1,
	0x7c, 0x93, 0x04, 0x59, 0x18, 0x47, 0xe8, 0x39, 0x34, 0xc2, 0x68, 0xc6, 0xee, 0x74, 0xe5, 0x44,
	0x39, 0xad, 0x93, 0x1c, 0x18, 0xd7, 0xb0, 0x7b, 0xc5, 0xb2, 0x60, 0x16, 0x64, 0x01, 0x57, 0xdc,
	0x06, 0xcb, 0x0d, 0x13, 0x8a, 0x3d, 0x92, 0x03, 0xf4, 0x0d, 0x40, 0x1a, 0xce, 0xa3, 0x20, 0xdb,
	0x24, 0x2c, 0xd5, 0x6b, 0x27, 0xea, 0x69, 0xe7, 0xe2, 0x55, 0xb7, 0x30, 0x26, 0xef, 0x52, 0xa9,
	0x20, 0x5b, 0x62, 0xe3, 0x1a, 0x0e, 0x9e, 0x08, 0xd0, 0x29, 0x3c, 0x2b, 0x25, 0x6f, 0x59, 0x30,
	0x63, 0x49, 0x51, 0xef, 0x71, 0x18, 0x7d, 0x0c, 0xed, 0x32, 0xa4, 0xd7, 0x84, 0xe6, 0x21, 0x60,
	0xfc, 0xa4, 0x40, 0xb3, 0x10, 0x7e, 0x05, 0x9d, 0xe9, 0x22, 0x08, 0xa3, 0xad, 0x74, 0x9d, 0x8b,
	0x43, 0xe9, 0x71, 0xf0, 0x40, 0x91, 0x6d, 0x1d, 0x32, 0x9f, 0x3a, 0xa9, 0x89, 0xab, 0x2f, 0xe5,
	0x55, 0x5a, 0xa5, 0x9f, 0x58, 0x34, 0xfe, 0x56, 0xa0, 0xb3, 0x95, 0x1f, 0x21, 0xa8, 0x67, 0xf7,
	0xeb, 0x7c, 0x82, 0x0d, 0x22, 0xce, 0x48, 0x87, 0xd6, 0x2d, 0x4b, 0xd2, 0x30, 0x8e, 0x44, 0xfa,
	0x06, 0x91, 0x10, 0x7d, 0x0d, 0xed, 0x72, 0x75, 0xba, 0x2a, 0x4a, 0x1f, 0x77, 0xf3, 0xe5, 0x76,
	0xe5, 0x72, 0xbb, 0x9e, 0x54, 0x90, 0x07, 0x31, 0xcf, 0x29, 0x3a, 0xb1, 0x2d, 0xbd, 0x7e, 0xa2,
	0x9c, 0xb6, 0x89, 0x84, 0xc2, 0xc1, 0x9d, 0x6d, 0xe9, 0x0d, 0x11, 0x16, 0x67, 0xbe, 0x58, 0xb6,
	0x8e, 0xa7, 0x0b, 0xbd, 0x99, 0xaf, 0x5e, 0x00, 0x3e, 0x5e, 0x76, 0x97, 0xb1, 0x48, 0x38, 0x6b,
	0xe5, 0xe3, 0x2d, 0x03, 0x86, 0x09, 0xcf, 0x1e, 0x75, 0x2f, 0x8a, 0x26, 0x2c, 0xc8, 0x62, 0xb9,
	0x31, 0x09, 0x79, 0x81, 0x28, 0x8e, 0xa6, 0x72, 0x4b, 0x39, 0x30, 0x30, 0xb4, 0xde, 0x07, 0xf7,
	0xcb, 0x38, 0x98, 0xa1, 0x4f, 0xa1, 0xb9, 0xd8, 0x5e, 0xce, 0xbe, 0x9c, 0x70, 0x31, 0xd8, 0x82,
	0xe5, 0xee, 0xf9, 0x6b, 0x29, 0xf2, 0x88, 0xb3, 0xd1, 0x87, 0x5d, 0x1c, 0xdd, 0xb2, 0x65, 0x9c,
	0xcf, 0x72, 0x9d, 0xa7, 0x94, 0x16, 0x0a, 0xf8, 0x3f, 0x8f, 0xe5, 0x67, 0x05, 0x1a, 0xfd, 0x65,
	0x3c, 0xfd, 0x11, 0x9d, 0xcb, 0x57, 0xf3, 0xf8, 0x99, 0x08, 0x5a, 0xda, 0x29, 0x3a, 0xfe, 0x04,
	0xea, 0x96, 0xb4, 0xd3, 0xb9, 0x38, 0xa8, 0x48, 0x39, 0x41, 0x04, 0x8d, 0xbe, 0x78, 0xf8, 0x11,
	0x15, 0x6b, 0x3c, 0xaa, 0x48, 0x25, 0x49, 0x4a, 0x99, 0xc1, 0xa0, 0xb3, 0x55, 0x10, 0xbd, 0x80,
	0xa6, 0xb3, 0x59, 0x4d, 0x0a, 0x57, 0x75, 0x52, 0x20, 0x64, 0xc0, 0xde, 0xfb, 0x84, 0xdd, 0x86,
	0xf1, 0x26, 0x7d, 0x1b, 0xa4, 0x8b, 0xa2, 0xb1, 0x4a, 0x0c, 0x1d, 0xc3, 0x2e, 0x77, 0x21, 0x78,
	0x55, 0xf0, 0x25, 0x36, 0xde, 0x40, 0xbb, 0x34, 0xcb, 0x87, 0x2b, 0xba, 0x51, 0x4e, 0x54, 0x3e,
	0x5c, 0x7e, 0x36, 0xce, 0xe1, 0xa3, 0x8a, 0x45, 0x9e, 0xad, 0xec, 0x25, 0x17, 0x96, 0xf8, 0xec,
	0x4f, 0x05, 0x9a, 0x34, 0x0b, 0xb2, 0x4d, 0x8a, 0x3a, 0xd0, 0x1a, 0x3b, 0xef, 0x1c, 0xf7, 0x3b,
	0x47, 0xdb, 0x41, 0x7b, 0xd0, 0xa2, 0xe3, 0xc1, 0x00, 0x53, 0xaa, 0xfd, 0xa5, 0x20, 0x0d, 0x3a,
	0x7d, 0xd3, 0xf2, 0x09, 0xfe, 0x76, 0x8c, 0xa9, 0xa7, 0xfd, 0xa2, 0xa2, 0x7d, 0x68, 0x0f, 0x5d,
	0xd2, 0xb7, 0x2d, 0x0b, 0x3b, 0xda, 0xaf, 0x02, 0x3b, 0xae, 0xe7, 0x0f, 0xdd, 0xb1, 0x63, 0x69,
	0xbf, 0xa9, 0xe8, 0x35, 0xe8, 0x85, 0xda, 0xc7, 0x8e, 0x67, 0x7b, 0xdf, 0xfb, 0x9e, 0xeb, 0xfa,
	0x23, 0x93, 0x5c, 0x62, 0xed, 0x0f, 0x15, 0x1d, 0xc3, 0x91, 0xed, 0x78, 0x98, 0x38, 0xe6, 0xc8,
	0xa7, 0x98, 0x7c, 0xc0, 0xc4, 0xc7, 0x84, 0xb8, 0x44, 0xfb, 0x47, 0x45, 0x3a, 0x1c, 0xf2, 0x90,
	0x3d, 0xc0, 0xfe, 0xd8, 0x31, 0x3f, 0x98, 0xf6, 0xc8, 0xec, 0x8f, 0xb0, 0xf6, 0xaf, 0x7a, 0xf6,
	0xbb, 0x02, 0x90, 0x4f, 0xd7, 0xe3, 0xbf, 0xc2, 0x0e, 0xb4, 0xae, 0x30, 0xa5, 0xe6, 0x25, 0xd6,
	0x76, 0xd0, 0x6b, 0x78, 0x35, 0x70, 0x9d, 0xa1, 0x7d, 0x39, 0x26, 0xa6, 0x67, 0xbb, 0x8e, 0xef,
	0x11, 0xd3, 0xa1, 0xe6, 0x80, 0x9f, 0x35, 0x05, 0xbd, 0x00, 0x54, 0xa5, 0x6d, 0x0f, 0x5f, 0x69,
	0x35, 0xa4, 0xc3, 0x73, 0xec, 0x58, 0x2e, 0xa1, 0x98, 0x54, 0x6e, 0xa8, 0xe8, 0x25, 0x1c, 0xba,
	0xc4, 0xc2, 0xe4, 0x11, 0x51, 0x47, 0x47, 0x70, 0x60, 0xe1, 0x91, 0xcd, 0x3d, 0x53, 0x8c, 0xdf,
	0xf9, 0xb6, 0x33, 0x74, 0xb5, 0xc6, 0xd9, 0x18, 0x50, 0x65, 0xec, 0x36, 0xff, 0x67, 0x8c, 0xf6,
	0x01, 0xa8, 0x7d, 0xe9, 0x98, 0xde, 0x98, 0x60, 0xaa, 0xed, 0x70, 0x1f, 0x23, 0x93, 0x7a, 0x7e,
	0xc5, 0x8c, 0xa6, 0xf0, 0x6a, 0x5b, 0x55, 0xa8, 0x3f, 0xb4, 0x47, 0x1e, 0x26, 0x5a, 0xad, 0xff,
	0xf9, 0x0f, 0xe7, 0xf3, 0x30, 0x5b, 0x6c, 0x26, 0xfc, 0xf9, 0xf5, 0x16, 0xf7, 0x6b, 0x96, 0x2c,
	0xd9, 0x6c, 0xce, 0x92, 0xde, 0x4d, 0x30, 0x49, 0xc2, 0x69, 0xfe, 0xc9, 0x48, 0x8b, 0xcf, 0xca,
	0xa4, 0x29, 0xe0, 0x97, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x24, 0x97, 0x92, 0x02, 0x6e, 0x06,
	0x00, 0x00,
}
