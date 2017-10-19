// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata/metadata.proto

/*
Package metadata is a generated protocol buffer package.

It is generated from these files:
	metadata/metadata.proto

It has these top-level messages:
	HashingCosts
	WrappedKeyData
	ProtectorData
	EncryptionOptions
	WrappedPolicyKey
	PolicyData
	Config
*/
package metadata

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Specifies the method in which an outside secret is obtained for a Protector
type SourceType int32

const (
	SourceType_none              SourceType = 0
	SourceType_default           SourceType = 0
	SourceType_pam_passphrase    SourceType = 1
	SourceType_custom_passphrase SourceType = 2
	SourceType_raw_key           SourceType = 3
)

var SourceType_name = map[int32]string{
	0: "none",
	// Duplicate value: 0: "default",
	1: "pam_passphrase",
	2: "custom_passphrase",
	3: "raw_key",
}
var SourceType_value = map[string]int32{
	"none":              0,
	"default":           0,
	"pam_passphrase":    1,
	"custom_passphrase": 2,
	"raw_key":           3,
}

func (x SourceType) String() string {
	return proto.EnumName(SourceType_name, int32(x))
}
func (SourceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Type of encryption; should match declarations of unix.FS_ENCRYPTION_MODE
type EncryptionOptions_Mode int32

const (
	EncryptionOptions_default     EncryptionOptions_Mode = 0
	EncryptionOptions_AES_256_XTS EncryptionOptions_Mode = 1
	EncryptionOptions_AES_256_GCM EncryptionOptions_Mode = 2
	EncryptionOptions_AES_256_CBC EncryptionOptions_Mode = 3
	EncryptionOptions_AES_256_CTS EncryptionOptions_Mode = 4
	EncryptionOptions_AES_128_CBC EncryptionOptions_Mode = 5
	EncryptionOptions_AES_128_CTS EncryptionOptions_Mode = 6
)

var EncryptionOptions_Mode_name = map[int32]string{
	0: "default",
	1: "AES_256_XTS",
	2: "AES_256_GCM",
	3: "AES_256_CBC",
	4: "AES_256_CTS",
	5: "AES_128_CBC",
	6: "AES_128_CTS",
}
var EncryptionOptions_Mode_value = map[string]int32{
	"default":     0,
	"AES_256_XTS": 1,
	"AES_256_GCM": 2,
	"AES_256_CBC": 3,
	"AES_256_CTS": 4,
	"AES_128_CBC": 5,
	"AES_128_CTS": 6,
}

func (x EncryptionOptions_Mode) String() string {
	return proto.EnumName(EncryptionOptions_Mode_name, int32(x))
}
func (EncryptionOptions_Mode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

// Cost parameters to be used in our hashing functions.
type HashingCosts struct {
	Time        int64 `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	Memory      int64 `protobuf:"varint,3,opt,name=memory" json:"memory,omitempty"`
	Parallelism int64 `protobuf:"varint,4,opt,name=parallelism" json:"parallelism,omitempty"`
}

func (m *HashingCosts) Reset()                    { *m = HashingCosts{} }
func (m *HashingCosts) String() string            { return proto.CompactTextString(m) }
func (*HashingCosts) ProtoMessage()               {}
func (*HashingCosts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HashingCosts) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *HashingCosts) GetMemory() int64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *HashingCosts) GetParallelism() int64 {
	if m != nil {
		return m.Parallelism
	}
	return 0
}

// This structure is used for our authenticated wrapping/unwrapping of keys.
type WrappedKeyData struct {
	IV           []byte `protobuf:"bytes,1,opt,name=IV,json=iV,proto3" json:"IV,omitempty"`
	EncryptedKey []byte `protobuf:"bytes,2,opt,name=encrypted_key,json=encryptedKey,proto3" json:"encrypted_key,omitempty"`
	Hmac         []byte `protobuf:"bytes,3,opt,name=hmac,proto3" json:"hmac,omitempty"`
}

func (m *WrappedKeyData) Reset()                    { *m = WrappedKeyData{} }
func (m *WrappedKeyData) String() string            { return proto.CompactTextString(m) }
func (*WrappedKeyData) ProtoMessage()               {}
func (*WrappedKeyData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *WrappedKeyData) GetIV() []byte {
	if m != nil {
		return m.IV
	}
	return nil
}

func (m *WrappedKeyData) GetEncryptedKey() []byte {
	if m != nil {
		return m.EncryptedKey
	}
	return nil
}

func (m *WrappedKeyData) GetHmac() []byte {
	if m != nil {
		return m.Hmac
	}
	return nil
}

// The associated data for each protector
type ProtectorData struct {
	ProtectorDescriptor string     `protobuf:"bytes,1,opt,name=protector_descriptor,json=protectorDescriptor" json:"protector_descriptor,omitempty"`
	Source              SourceType `protobuf:"varint,2,opt,name=source,enum=metadata.SourceType" json:"source,omitempty"`
	// These are only used by some of the protector types
	Name       string          `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Costs      *HashingCosts   `protobuf:"bytes,4,opt,name=costs" json:"costs,omitempty"`
	Salt       []byte          `protobuf:"bytes,5,opt,name=salt,proto3" json:"salt,omitempty"`
	Uid        int64           `protobuf:"varint,6,opt,name=uid" json:"uid,omitempty"`
	WrappedKey *WrappedKeyData `protobuf:"bytes,7,opt,name=wrapped_key,json=wrappedKey" json:"wrapped_key,omitempty"`
}

func (m *ProtectorData) Reset()                    { *m = ProtectorData{} }
func (m *ProtectorData) String() string            { return proto.CompactTextString(m) }
func (*ProtectorData) ProtoMessage()               {}
func (*ProtectorData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ProtectorData) GetProtectorDescriptor() string {
	if m != nil {
		return m.ProtectorDescriptor
	}
	return ""
}

func (m *ProtectorData) GetSource() SourceType {
	if m != nil {
		return m.Source
	}
	return SourceType_none
}

func (m *ProtectorData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProtectorData) GetCosts() *HashingCosts {
	if m != nil {
		return m.Costs
	}
	return nil
}

func (m *ProtectorData) GetSalt() []byte {
	if m != nil {
		return m.Salt
	}
	return nil
}

func (m *ProtectorData) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *ProtectorData) GetWrappedKey() *WrappedKeyData {
	if m != nil {
		return m.WrappedKey
	}
	return nil
}

// Encryption policy specifics, corresponds to the fscrypt_policy struct
type EncryptionOptions struct {
	Padding   int64                  `protobuf:"varint,1,opt,name=padding" json:"padding,omitempty"`
	Contents  EncryptionOptions_Mode `protobuf:"varint,2,opt,name=contents,enum=metadata.EncryptionOptions_Mode" json:"contents,omitempty"`
	Filenames EncryptionOptions_Mode `protobuf:"varint,3,opt,name=filenames,enum=metadata.EncryptionOptions_Mode" json:"filenames,omitempty"`
}

func (m *EncryptionOptions) Reset()                    { *m = EncryptionOptions{} }
func (m *EncryptionOptions) String() string            { return proto.CompactTextString(m) }
func (*EncryptionOptions) ProtoMessage()               {}
func (*EncryptionOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *EncryptionOptions) GetPadding() int64 {
	if m != nil {
		return m.Padding
	}
	return 0
}

func (m *EncryptionOptions) GetContents() EncryptionOptions_Mode {
	if m != nil {
		return m.Contents
	}
	return EncryptionOptions_default
}

func (m *EncryptionOptions) GetFilenames() EncryptionOptions_Mode {
	if m != nil {
		return m.Filenames
	}
	return EncryptionOptions_default
}

type WrappedPolicyKey struct {
	ProtectorDescriptor string          `protobuf:"bytes,1,opt,name=protector_descriptor,json=protectorDescriptor" json:"protector_descriptor,omitempty"`
	WrappedKey          *WrappedKeyData `protobuf:"bytes,2,opt,name=wrapped_key,json=wrappedKey" json:"wrapped_key,omitempty"`
}

func (m *WrappedPolicyKey) Reset()                    { *m = WrappedPolicyKey{} }
func (m *WrappedPolicyKey) String() string            { return proto.CompactTextString(m) }
func (*WrappedPolicyKey) ProtoMessage()               {}
func (*WrappedPolicyKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *WrappedPolicyKey) GetProtectorDescriptor() string {
	if m != nil {
		return m.ProtectorDescriptor
	}
	return ""
}

func (m *WrappedPolicyKey) GetWrappedKey() *WrappedKeyData {
	if m != nil {
		return m.WrappedKey
	}
	return nil
}

// The associated data for each policy
type PolicyData struct {
	KeyDescriptor     string              `protobuf:"bytes,1,opt,name=key_descriptor,json=keyDescriptor" json:"key_descriptor,omitempty"`
	Options           *EncryptionOptions  `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
	WrappedPolicyKeys []*WrappedPolicyKey `protobuf:"bytes,3,rep,name=wrapped_policy_keys,json=wrappedPolicyKeys" json:"wrapped_policy_keys,omitempty"`
}

func (m *PolicyData) Reset()                    { *m = PolicyData{} }
func (m *PolicyData) String() string            { return proto.CompactTextString(m) }
func (*PolicyData) ProtoMessage()               {}
func (*PolicyData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PolicyData) GetKeyDescriptor() string {
	if m != nil {
		return m.KeyDescriptor
	}
	return ""
}

func (m *PolicyData) GetOptions() *EncryptionOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *PolicyData) GetWrappedPolicyKeys() []*WrappedPolicyKey {
	if m != nil {
		return m.WrappedPolicyKeys
	}
	return nil
}

// Data stored in the config file
type Config struct {
	Source        SourceType         `protobuf:"varint,1,opt,name=source,enum=metadata.SourceType" json:"source,omitempty"`
	HashCosts     *HashingCosts      `protobuf:"bytes,2,opt,name=hash_costs,json=hashCosts" json:"hash_costs,omitempty"`
	Compatibility string             `protobuf:"bytes,3,opt,name=compatibility" json:"compatibility,omitempty"`
	Options       *EncryptionOptions `protobuf:"bytes,4,opt,name=options" json:"options,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Config) GetSource() SourceType {
	if m != nil {
		return m.Source
	}
	return SourceType_none
}

func (m *Config) GetHashCosts() *HashingCosts {
	if m != nil {
		return m.HashCosts
	}
	return nil
}

func (m *Config) GetCompatibility() string {
	if m != nil {
		return m.Compatibility
	}
	return ""
}

func (m *Config) GetOptions() *EncryptionOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func init() {
	proto.RegisterType((*HashingCosts)(nil), "metadata.HashingCosts")
	proto.RegisterType((*WrappedKeyData)(nil), "metadata.WrappedKeyData")
	proto.RegisterType((*ProtectorData)(nil), "metadata.ProtectorData")
	proto.RegisterType((*EncryptionOptions)(nil), "metadata.EncryptionOptions")
	proto.RegisterType((*WrappedPolicyKey)(nil), "metadata.WrappedPolicyKey")
	proto.RegisterType((*PolicyData)(nil), "metadata.PolicyData")
	proto.RegisterType((*Config)(nil), "metadata.Config")
	proto.RegisterEnum("metadata.SourceType", SourceType_name, SourceType_value)
	proto.RegisterEnum("metadata.EncryptionOptions_Mode", EncryptionOptions_Mode_name, EncryptionOptions_Mode_value)
}

func init() { proto.RegisterFile("metadata/metadata.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 651 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xdd, 0x6a, 0xdb, 0x30,
	0x14, 0xc7, 0x6b, 0x27, 0x4d, 0x9a, 0x93, 0x8f, 0xb9, 0x6a, 0xd7, 0x99, 0xee, 0x26, 0x78, 0x1b,
	0x94, 0x51, 0x3a, 0x9a, 0xd1, 0xb1, 0xc1, 0x18, 0x6c, 0x69, 0xd9, 0x47, 0x29, 0x2b, 0x4a, 0xe8,
	0x36, 0x18, 0x18, 0xd5, 0x56, 0x1b, 0x51, 0xdb, 0x12, 0x96, 0x42, 0xf0, 0xdd, 0xde, 0x61, 0xef,
	0xb0, 0x47, 0xd8, 0x43, 0xec, 0xa9, 0x86, 0xe4, 0xd8, 0x71, 0x5a, 0x28, 0xdd, 0x6e, 0xcc, 0xd1,
	0x5f, 0xd2, 0xf9, 0x1f, 0xfd, 0xa4, 0x63, 0x78, 0x10, 0x53, 0x45, 0x42, 0xa2, 0xc8, 0xb3, 0x22,
	0xd8, 0x13, 0x29, 0x57, 0x1c, 0xad, 0x15, 0x63, 0xef, 0x3b, 0x74, 0x3e, 0x10, 0x39, 0x61, 0xc9,
	0xe5, 0x90, 0x4b, 0x25, 0x11, 0x82, 0xba, 0x62, 0x31, 0x75, 0xed, 0xbe, 0xb5, 0x53, 0xc3, 0x26,
	0x46, 0x5b, 0xd0, 0x88, 0x69, 0xcc, 0xd3, 0xcc, 0xad, 0x19, 0x75, 0x3e, 0x42, 0x7d, 0x68, 0x0b,
	0x92, 0x92, 0x28, 0xa2, 0x11, 0x93, 0xb1, 0x5b, 0x37, 0x93, 0x55, 0xc9, 0xfb, 0x06, 0xbd, 0x2f,
	0x29, 0x11, 0x82, 0x86, 0xc7, 0x34, 0x3b, 0x24, 0x8a, 0xa0, 0x1e, 0xd8, 0x1f, 0xcf, 0x5c, 0xab,
	0x6f, 0xed, 0x74, 0xb0, 0xcd, 0xce, 0xd0, 0x23, 0xe8, 0xd2, 0x24, 0x48, 0x33, 0xa1, 0x68, 0xe8,
	0x5f, 0xd1, 0xcc, 0x18, 0x77, 0x70, 0xa7, 0x14, 0x8f, 0x69, 0xa6, 0x8b, 0x9a, 0xc4, 0x24, 0x30,
	0xf6, 0x1d, 0x6c, 0x62, 0xef, 0xa7, 0x0d, 0xdd, 0xd3, 0x94, 0x2b, 0x1a, 0x28, 0x9e, 0x9a, 0xd4,
	0xfb, 0xb0, 0x29, 0x0a, 0xc1, 0x0f, 0xa9, 0x0c, 0x52, 0x26, 0x14, 0x4f, 0x8d, 0x59, 0x0b, 0x6f,
	0x94, 0x73, 0x87, 0xe5, 0x14, 0xda, 0x85, 0x86, 0xe4, 0xd3, 0x34, 0xc8, 0xcf, 0xdb, 0x1b, 0x6c,
	0xee, 0x95, 0xa0, 0x46, 0x46, 0x1f, 0x67, 0x82, 0xe2, 0xf9, 0x1a, 0x5d, 0x46, 0x42, 0x62, 0x6a,
	0xca, 0x68, 0x61, 0x13, 0xa3, 0x5d, 0x58, 0x0d, 0x34, 0x38, 0x73, 0xfa, 0xf6, 0x60, 0x6b, 0x91,
	0xa0, 0x8a, 0x15, 0xe7, 0x8b, 0x74, 0x06, 0x49, 0x22, 0xe5, 0xae, 0xe6, 0x07, 0xd1, 0x31, 0x72,
	0xa0, 0x36, 0x65, 0xa1, 0xdb, 0x30, 0xf4, 0x74, 0x88, 0x5e, 0x41, 0x7b, 0x96, 0x53, 0x33, 0x44,
	0x9a, 0x26, 0xb3, 0xbb, 0xc8, 0xbc, 0x8c, 0x14, 0xc3, 0xac, 0x1c, 0x7b, 0xbf, 0x6c, 0x58, 0x3f,
	0xca, 0xd1, 0x31, 0x9e, 0x7c, 0x36, 0x5f, 0x89, 0x5c, 0x68, 0x0a, 0x12, 0x86, 0x2c, 0xb9, 0x34,
	0x30, 0x6a, 0xb8, 0x18, 0xa2, 0xd7, 0xb0, 0x16, 0xf0, 0x44, 0xd1, 0x44, 0xc9, 0x39, 0x82, 0xfe,
	0xc2, 0xe7, 0x46, 0xa2, 0xbd, 0x13, 0x1e, 0x52, 0x5c, 0xee, 0x40, 0x6f, 0xa0, 0x75, 0xc1, 0x22,
	0xaa, 0x41, 0x48, 0x43, 0xe5, 0x2e, 0xdb, 0x17, 0x5b, 0xbc, 0x0c, 0xea, 0x5a, 0x42, 0x6d, 0x68,
	0x86, 0xf4, 0x82, 0x4c, 0x23, 0xe5, 0xac, 0xa0, 0x7b, 0xd0, 0x7e, 0x7b, 0x34, 0xf2, 0x07, 0x07,
	0x2f, 0xfc, 0xaf, 0xe3, 0x91, 0x63, 0x55, 0x85, 0xf7, 0xc3, 0x13, 0xc7, 0xae, 0x0a, 0xc3, 0x77,
	0x43, 0xa7, 0xb6, 0x24, 0x8c, 0x47, 0x4e, 0xbd, 0x10, 0xf6, 0x07, 0x2f, 0xcd, 0x8a, 0xd5, 0x25,
	0x61, 0x3c, 0x72, 0x1a, 0xde, 0x0f, 0x0b, 0x9c, 0x39, 0xc7, 0x53, 0x1e, 0xb1, 0x20, 0xd3, 0xef,
	0xec, 0x3f, 0x5e, 0xd0, 0xb5, 0xbb, 0xb2, 0xff, 0xe1, 0xae, 0x7e, 0x5b, 0x00, 0xb9, 0xb7, 0x79,
	0xbe, 0x4f, 0xa0, 0x77, 0x45, 0xb3, 0x9b, 0xb6, 0xdd, 0x2b, 0x9a, 0x55, 0x0c, 0x0f, 0xa0, 0xc9,
	0x73, 0x9c, 0x73, 0xb3, 0x87, 0xb7, 0x10, 0xc7, 0xc5, 0x5a, 0xf4, 0x09, 0x36, 0x8a, 0x3a, 0x85,
	0xf1, 0xd4, 0xe5, 0xea, 0x4b, 0xab, 0xed, 0xb4, 0x07, 0xdb, 0x37, 0xea, 0x2d, 0x99, 0xe0, 0xf5,
	0xd9, 0x35, 0x45, 0x7a, 0x7f, 0x2c, 0x68, 0x0c, 0x79, 0x72, 0xc1, 0x2e, 0x2b, 0x0d, 0x64, 0xdd,
	0xa1, 0x81, 0x0e, 0x00, 0x26, 0x44, 0x4e, 0xfc, 0xbc, 0x63, 0xec, 0x5b, 0x3b, 0xa6, 0xa5, 0x57,
	0xe6, 0xff, 0xa4, 0xc7, 0xd0, 0x0d, 0x78, 0x2c, 0x88, 0x62, 0xe7, 0x2c, 0x62, 0x2a, 0x9b, 0x37,
	0xe0, 0xb2, 0x58, 0x05, 0x53, 0xbf, 0x3b, 0x98, 0xa7, 0x3e, 0xc0, 0xa2, 0x52, 0xb4, 0x06, 0xf5,
	0x84, 0x27, 0xd4, 0x59, 0x59, 0x7e, 0x93, 0x08, 0x7a, 0x82, 0xc4, 0xbe, 0x20, 0x52, 0x8a, 0x49,
	0x4a, 0x24, 0x75, 0x2c, 0x74, 0x1f, 0xd6, 0x83, 0xa9, 0x54, 0x7c, 0x49, 0xb6, 0xf5, 0xbe, 0x94,
	0xcc, 0x34, 0x5d, 0xa7, 0xb6, 0x6d, 0x3b, 0xd6, 0x79, 0xc3, 0xfc, 0x72, 0x9f, 0xff, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0xf8, 0xd5, 0x1e, 0x01, 0x8d, 0x05, 0x00, 0x00,
}
