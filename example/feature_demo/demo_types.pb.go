// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/feature_demo/demo_types.proto

package example // import "github.com/infobloxopen/protoc-gen-gorm/example/feature_demo"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import user "github.com/infobloxopen/protoc-gen-gorm/example/user"
import _ "github.com/infobloxopen/protoc-gen-gorm/options"
import types "github.com/infobloxopen/protoc-gen-gorm/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// enums are mapped to the their underlying numeric value in the db.
// This is practical from an API perspective, but tougher for debugging.
// Strings with validation constraints can be used instead if desired
type TestTypesStatus int32

const (
	TestTypes_UNKNOWN TestTypesStatus = 0
	TestTypes_GOOD    TestTypesStatus = 1
	TestTypes_BAD     TestTypesStatus = 2
)

var TestTypesStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "GOOD",
	2: "BAD",
}
var TestTypesStatus_value = map[string]int32{
	"UNKNOWN": 0,
	"GOOD":    1,
	"BAD":     2,
}

func (x TestTypesStatus) String() string {
	return proto.EnumName(TestTypesStatus_name, int32(x))
}
func (TestTypesStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_demo_types_3580d965b1e3ee92, []int{0, 0}
}

// test_types is a message that includes a representative sample of the
// available types
type TestTypes struct {
	// the (gorm.field).drop option allows for setting a field to be API only
	ApiOnlyString string `protobuf:"bytes,1,opt,name=api_only_string,json=apiOnlyString" json:"api_only_string,omitempty"`
	// repeated raw types are currently unsupported, so this field will be dropped
	// at the ORM level
	Numbers []int32 `protobuf:"varint,2,rep,packed,name=numbers" json:"numbers,omitempty"`
	// a StringValue represents a Nullable string
	OptionalString *wrappers.StringValue `protobuf:"bytes,3,opt,name=optional_string,json=optionalString" json:"optional_string,omitempty"`
	BecomesInt     TestTypesStatus       `protobuf:"varint,4,opt,name=becomes_int,json=becomesInt,enum=example.TestTypesStatus" json:"becomes_int,omitempty"`
	// The Empty type serves no purpose outside of rpc calls and is dropped
	// automatically from objects
	Nothingness *empty.Empty `protobuf:"bytes,5,opt,name=nothingness" json:"nothingness,omitempty"`
	// The UUID custom type should act like a StringValue at the API level, but is
	// automatically converted to and from a uuid.UUID (github.com/satori/go.uuid)
	Uuid *types.UUID `protobuf:"bytes,6,opt,name=uuid" json:"uuid,omitempty"`
	// Timestamps convert to golang's time.Time type, and created_at and
	// updated_at values are automatically filled by GORM
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	// This represents a foreign key to the 'type_with_id' type for associations
	// This could be hidden from the API (or soon autogenerated).
	TypeWithIdId uint32 `protobuf:"varint,8,opt,name=type_with_id_id,json=typeWithIdId" json:"type_with_id_id,omitempty"`
	// This is an arbitrary JSON string that is marshalled and unmarshalled
	// specially in grpc-gateway as a JSON object
	JsonField *types.JSONValue `protobuf:"bytes,9,opt,name=json_field,json=jsonField" json:"json_field,omitempty"`
	// The UUIDValue custom type should act like a StringValue at the API level, but is
	// automatically converted to and from a *uuid.UUID (github.com/satori/go.uuid)
	NullableUuid         *types.UUIDValue `protobuf:"bytes,10,opt,name=nullable_uuid,json=nullableUuid" json:"nullable_uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TestTypes) Reset()         { *m = TestTypes{} }
func (m *TestTypes) String() string { return proto.CompactTextString(m) }
func (*TestTypes) ProtoMessage()    {}
func (*TestTypes) Descriptor() ([]byte, []int) {
	return fileDescriptor_demo_types_3580d965b1e3ee92, []int{0}
}
func (m *TestTypes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestTypes.Unmarshal(m, b)
}
func (m *TestTypes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestTypes.Marshal(b, m, deterministic)
}
func (dst *TestTypes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestTypes.Merge(dst, src)
}
func (m *TestTypes) XXX_Size() int {
	return xxx_messageInfo_TestTypes.Size(m)
}
func (m *TestTypes) XXX_DiscardUnknown() {
	xxx_messageInfo_TestTypes.DiscardUnknown(m)
}

var xxx_messageInfo_TestTypes proto.InternalMessageInfo

func (m *TestTypes) GetApiOnlyString() string {
	if m != nil {
		return m.ApiOnlyString
	}
	return ""
}

func (m *TestTypes) GetNumbers() []int32 {
	if m != nil {
		return m.Numbers
	}
	return nil
}

func (m *TestTypes) GetOptionalString() *wrappers.StringValue {
	if m != nil {
		return m.OptionalString
	}
	return nil
}

func (m *TestTypes) GetBecomesInt() TestTypesStatus {
	if m != nil {
		return m.BecomesInt
	}
	return TestTypes_UNKNOWN
}

func (m *TestTypes) GetNothingness() *empty.Empty {
	if m != nil {
		return m.Nothingness
	}
	return nil
}

func (m *TestTypes) GetUuid() *types.UUID {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func (m *TestTypes) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *TestTypes) GetTypeWithIdId() uint32 {
	if m != nil {
		return m.TypeWithIdId
	}
	return 0
}

func (m *TestTypes) GetJsonField() *types.JSONValue {
	if m != nil {
		return m.JsonField
	}
	return nil
}

func (m *TestTypes) GetNullableUuid() *types.UUIDValue {
	if m != nil {
		return m.NullableUuid
	}
	return nil
}

// TypeWithID demonstrates some basic assocation behavior
type TypeWithID struct {
	// any field named 'id' is assumed by gorm to be the primary key for the
	// object.
	Id uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// The field option also allows arbitrary tag setting, such as informing
	// gorm of a primary key, different column names or different types in the db
	Ip string `protobuf:"bytes,2,opt,name=ip" json:"ip,omitempty"`
	// A default has-many relationship, will error on generation if no FK field,
	// convention {typename}_id, is present. These FK fields will be automatically
	// populated on create and update.
	Things []*TestTypes `protobuf:"bytes,3,rep,name=things" json:"things,omitempty"`
	// A default has-one relationship, will error as above
	ANestedObject *TestTypes `protobuf:"bytes,4,opt,name=a_nested_object,json=aNestedObject" json:"a_nested_object,omitempty"`
	// An in-package and cross-package imported type (in-package can use any
	// association type, cross-package is limited to belongs_to and many_to_many)
	Point                *IntPoint        `protobuf:"bytes,5,opt,name=point" json:"point,omitempty"`
	User                 *user.User       `protobuf:"bytes,6,opt,name=user" json:"user,omitempty"`
	Address              *types.InetValue `protobuf:"bytes,7,opt,name=address" json:"address,omitempty"`
	MultiaccountTypeIds  []uint32         `protobuf:"varint,8,rep,packed,name=multiaccount_type_ids,json=multiaccountTypeIds" json:"multiaccount_type_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TypeWithID) Reset()         { *m = TypeWithID{} }
func (m *TypeWithID) String() string { return proto.CompactTextString(m) }
func (*TypeWithID) ProtoMessage()    {}
func (*TypeWithID) Descriptor() ([]byte, []int) {
	return fileDescriptor_demo_types_3580d965b1e3ee92, []int{1}
}
func (m *TypeWithID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeWithID.Unmarshal(m, b)
}
func (m *TypeWithID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeWithID.Marshal(b, m, deterministic)
}
func (dst *TypeWithID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeWithID.Merge(dst, src)
}
func (m *TypeWithID) XXX_Size() int {
	return xxx_messageInfo_TypeWithID.Size(m)
}
func (m *TypeWithID) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeWithID.DiscardUnknown(m)
}

var xxx_messageInfo_TypeWithID proto.InternalMessageInfo

func (m *TypeWithID) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TypeWithID) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *TypeWithID) GetThings() []*TestTypes {
	if m != nil {
		return m.Things
	}
	return nil
}

func (m *TypeWithID) GetANestedObject() *TestTypes {
	if m != nil {
		return m.ANestedObject
	}
	return nil
}

func (m *TypeWithID) GetPoint() *IntPoint {
	if m != nil {
		return m.Point
	}
	return nil
}

func (m *TypeWithID) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *TypeWithID) GetAddress() *types.InetValue {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *TypeWithID) GetMultiaccountTypeIds() []uint32 {
	if m != nil {
		return m.MultiaccountTypeIds
	}
	return nil
}

// MultiaccountTypeWithID demonstrates the generated multi-account support
type MultiaccountTypeWithID struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	SomeField            string   `protobuf:"bytes,2,opt,name=some_field,json=someField" json:"some_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiaccountTypeWithID) Reset()         { *m = MultiaccountTypeWithID{} }
func (m *MultiaccountTypeWithID) String() string { return proto.CompactTextString(m) }
func (*MultiaccountTypeWithID) ProtoMessage()    {}
func (*MultiaccountTypeWithID) Descriptor() ([]byte, []int) {
	return fileDescriptor_demo_types_3580d965b1e3ee92, []int{2}
}
func (m *MultiaccountTypeWithID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiaccountTypeWithID.Unmarshal(m, b)
}
func (m *MultiaccountTypeWithID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiaccountTypeWithID.Marshal(b, m, deterministic)
}
func (dst *MultiaccountTypeWithID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiaccountTypeWithID.Merge(dst, src)
}
func (m *MultiaccountTypeWithID) XXX_Size() int {
	return xxx_messageInfo_MultiaccountTypeWithID.Size(m)
}
func (m *MultiaccountTypeWithID) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiaccountTypeWithID.DiscardUnknown(m)
}

var xxx_messageInfo_MultiaccountTypeWithID proto.InternalMessageInfo

func (m *MultiaccountTypeWithID) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MultiaccountTypeWithID) GetSomeField() string {
	if m != nil {
		return m.SomeField
	}
	return ""
}

type MultiaccountTypeWithoutID struct {
	SomeField            string   `protobuf:"bytes,1,opt,name=some_field,json=someField" json:"some_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiaccountTypeWithoutID) Reset()         { *m = MultiaccountTypeWithoutID{} }
func (m *MultiaccountTypeWithoutID) String() string { return proto.CompactTextString(m) }
func (*MultiaccountTypeWithoutID) ProtoMessage()    {}
func (*MultiaccountTypeWithoutID) Descriptor() ([]byte, []int) {
	return fileDescriptor_demo_types_3580d965b1e3ee92, []int{3}
}
func (m *MultiaccountTypeWithoutID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiaccountTypeWithoutID.Unmarshal(m, b)
}
func (m *MultiaccountTypeWithoutID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiaccountTypeWithoutID.Marshal(b, m, deterministic)
}
func (dst *MultiaccountTypeWithoutID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiaccountTypeWithoutID.Merge(dst, src)
}
func (m *MultiaccountTypeWithoutID) XXX_Size() int {
	return xxx_messageInfo_MultiaccountTypeWithoutID.Size(m)
}
func (m *MultiaccountTypeWithoutID) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiaccountTypeWithoutID.DiscardUnknown(m)
}

var xxx_messageInfo_MultiaccountTypeWithoutID proto.InternalMessageInfo

func (m *MultiaccountTypeWithoutID) GetSomeField() string {
	if m != nil {
		return m.SomeField
	}
	return ""
}

type APIOnlyType struct {
	// here the ormable flag is not used, so nothing will be generated for this
	// object at the ORM level, and when this type is used as a field or
	// repeated field in another message that field will be dropped in the Orm
	// model, and would have to be set by hook
	Contents             string   `protobuf:"bytes,1,opt,name=contents" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *APIOnlyType) Reset()         { *m = APIOnlyType{} }
func (m *APIOnlyType) String() string { return proto.CompactTextString(m) }
func (*APIOnlyType) ProtoMessage()    {}
func (*APIOnlyType) Descriptor() ([]byte, []int) {
	return fileDescriptor_demo_types_3580d965b1e3ee92, []int{4}
}
func (m *APIOnlyType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APIOnlyType.Unmarshal(m, b)
}
func (m *APIOnlyType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APIOnlyType.Marshal(b, m, deterministic)
}
func (dst *APIOnlyType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APIOnlyType.Merge(dst, src)
}
func (m *APIOnlyType) XXX_Size() int {
	return xxx_messageInfo_APIOnlyType.Size(m)
}
func (m *APIOnlyType) XXX_DiscardUnknown() {
	xxx_messageInfo_APIOnlyType.DiscardUnknown(m)
}

var xxx_messageInfo_APIOnlyType proto.InternalMessageInfo

func (m *APIOnlyType) GetContents() string {
	if m != nil {
		return m.Contents
	}
	return ""
}

func init() {
	proto.RegisterType((*TestTypes)(nil), "example.TestTypes")
	proto.RegisterType((*TypeWithID)(nil), "example.TypeWithID")
	proto.RegisterType((*MultiaccountTypeWithID)(nil), "example.MultiaccountTypeWithID")
	proto.RegisterType((*MultiaccountTypeWithoutID)(nil), "example.MultiaccountTypeWithoutID")
	proto.RegisterType((*APIOnlyType)(nil), "example.APIOnlyType")
	proto.RegisterEnum("example.TestTypesStatus", TestTypesStatus_name, TestTypesStatus_value)
}

func init() {
	proto.RegisterFile("example/feature_demo/demo_types.proto", fileDescriptor_demo_types_3580d965b1e3ee92)
}

var fileDescriptor_demo_types_3580d965b1e3ee92 = []byte{
	// 896 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0xc7, 0xeb, 0x7c, 0xe7, 0x64, 0xb3, 0xbb, 0x1d, 0xd4, 0xe2, 0x84, 0x8f, 0x46, 0x11, 0xa5,
	0xa1, 0x52, 0x6d, 0x29, 0xe5, 0x82, 0x86, 0x0b, 0xb4, 0xab, 0xb4, 0xc8, 0x45, 0x24, 0x8b, 0x9b,
	0x50, 0x09, 0x84, 0xac, 0xb1, 0x3d, 0xc9, 0x4e, 0x65, 0xcf, 0x18, 0xcf, 0x98, 0x36, 0x3c, 0x0c,
	0xef, 0x91, 0xdc, 0xf1, 0x66, 0x68, 0xc6, 0x76, 0x08, 0x9b, 0xad, 0x54, 0x6e, 0x12, 0xcd, 0x9c,
	0xdf, 0x39, 0xc7, 0x67, 0xfe, 0xff, 0x19, 0x78, 0x48, 0xde, 0xe1, 0x38, 0x89, 0x88, 0xbd, 0x22,
	0x58, 0x66, 0x29, 0xf1, 0x42, 0x12, 0x73, 0x5b, 0xfd, 0x78, 0x72, 0x93, 0x10, 0x61, 0x25, 0x29,
	0x97, 0x1c, 0x35, 0x0b, 0xac, 0x3f, 0x59, 0x53, 0x79, 0x9d, 0xf9, 0x56, 0xc0, 0x63, 0x9b, 0xb2,
	0x15, 0xf7, 0x23, 0xfe, 0x8e, 0x27, 0x84, 0xd9, 0x9a, 0x0b, 0x9e, 0xac, 0x09, 0x7b, 0xb2, 0xe6,
	0x69, 0x6c, 0xf3, 0x44, 0x52, 0xce, 0x84, 0xad, 0x16, 0x79, 0x91, 0xfe, 0xb3, 0x0f, 0xcd, 0xd5,
	0x9d, 0xed, 0x83, 0xfe, 0xfd, 0xcf, 0xd7, 0x9c, 0xaf, 0x23, 0x92, 0x93, 0x7e, 0xb6, 0xb2, 0xdf,
	0xa6, 0x38, 0x49, 0x48, 0x5a, 0xc6, 0x3f, 0xb9, 0x19, 0x27, 0x71, 0x22, 0x37, 0x45, 0xf0, 0xc1,
	0xcd, 0xa0, 0xa4, 0x31, 0x11, 0x12, 0xc7, 0x49, 0x01, 0x3c, 0x7a, 0xff, 0x21, 0x08, 0x92, 0xfe,
	0x41, 0x03, 0x52, 0x80, 0xdf, 0x7d, 0xe8, 0x04, 0x65, 0xc1, 0x4c, 0x90, 0x54, 0xff, 0xe4, 0x05,
	0x86, 0x7f, 0xd5, 0xa1, 0xbd, 0x20, 0x42, 0x2e, 0xd4, 0x6c, 0xc8, 0x82, 0x33, 0x9c, 0x50, 0x8f,
	0xb3, 0x68, 0xe3, 0x09, 0x99, 0x52, 0xb6, 0x36, 0x8d, 0x81, 0x31, 0x6a, 0x5f, 0x36, 0x76, 0xdb,
	0x5e, 0xe5, 0xdc, 0x70, 0xbb, 0x38, 0xa1, 0x73, 0x16, 0x6d, 0x5e, 0xe9, 0x20, 0x32, 0xa1, 0xc9,
	0xb2, 0xd8, 0x27, 0xa9, 0x30, 0x2b, 0x83, 0xea, 0xa8, 0xee, 0x96, 0x4b, 0xf4, 0x1c, 0xce, 0xf2,
	0x03, 0xc7, 0x51, 0x59, 0xa9, 0x3a, 0x30, 0x46, 0x9d, 0xf1, 0xa7, 0x56, 0x3e, 0xbc, 0x55, 0x0e,
	0x6f, 0xe5, 0xb5, 0x7e, 0xc6, 0x51, 0x46, 0xdc, 0xd3, 0x32, 0xa9, 0x68, 0x30, 0x81, 0x8e, 0x4f,
	0x02, 0x1e, 0x13, 0xe1, 0x51, 0x26, 0xcd, 0xda, 0xc0, 0x18, 0x9d, 0x8e, 0x7b, 0x56, 0x31, 0x8d,
	0xb5, 0xff, 0x72, 0x4b, 0x48, 0x2c, 0x33, 0xe1, 0x42, 0x41, 0x3b, 0x4c, 0xa2, 0x6f, 0xa0, 0xc3,
	0xb8, 0xbc, 0xa6, 0x6c, 0xcd, 0x88, 0x10, 0x66, 0x5d, 0xb7, 0xbf, 0x7f, 0xd4, 0xfe, 0xb9, 0x12,
	0xc6, 0x3d, 0x44, 0xd1, 0x17, 0x50, 0xcb, 0x32, 0x1a, 0x9a, 0x0d, 0x9d, 0x72, 0x6e, 0x69, 0xcb,
	0xe4, 0xea, 0x2f, 0x97, 0xce, 0xd4, 0xd5, 0x51, 0xf4, 0x0c, 0x20, 0x48, 0x09, 0x96, 0x24, 0xf4,
	0xb0, 0x34, 0x9b, 0x9a, 0xed, 0x1f, 0x95, 0x5f, 0x94, 0xd2, 0xba, 0xed, 0x82, 0xbe, 0x90, 0xe8,
	0x21, 0x9c, 0xa9, 0x72, 0xde, 0x5b, 0x2a, 0xaf, 0x3d, 0x1a, 0x7a, 0x34, 0x34, 0x5b, 0x03, 0x63,
	0xd4, 0x75, 0x4f, 0xd4, 0xf6, 0x6b, 0x2a, 0xaf, 0x9d, 0xd0, 0x09, 0xd1, 0xd7, 0x00, 0x6f, 0x04,
	0x67, 0xde, 0x8a, 0x92, 0x28, 0x34, 0xdb, 0xba, 0xc3, 0xbd, 0xc3, 0xaf, 0x79, 0xf9, 0x6a, 0x3e,
	0xcb, 0x0f, 0xae, 0xad, 0xc0, 0x17, 0x8a, 0x43, 0x13, 0xe8, 0xb2, 0x2c, 0x8a, 0xb0, 0x1f, 0x11,
	0x4f, 0x8f, 0x01, 0xc7, 0x89, 0x6a, 0x8c, 0x3c, 0xf1, 0xa4, 0x64, 0x97, 0x19, 0x0d, 0x87, 0x23,
	0x68, 0xe4, 0x27, 0x89, 0x3a, 0xd0, 0x5c, 0xce, 0x7e, 0x98, 0xcd, 0x5f, 0xcf, 0xce, 0xef, 0xa0,
	0x16, 0xd4, 0xbe, 0x9f, 0xcf, 0xa7, 0xe7, 0x06, 0x6a, 0x42, 0xf5, 0xf2, 0x62, 0x7a, 0x5e, 0x99,
	0xac, 0x76, 0xdb, 0x9e, 0xdf, 0x32, 0xd0, 0x23, 0xe8, 0xe4, 0x5a, 0x5d, 0xa4, 0x29, 0xde, 0xa0,
	0x3a, 0x56, 0x7f, 0xc3, 0xbb, 0x07, 0xbe, 0x8c, 0xa8, 0x6f, 0x27, 0xbf, 0xa3, 0xd1, 0x7f, 0xc1,
	0x86, 0x06, 0xc7, 0xb7, 0x90, 0xfd, 0x8e, 0x88, 0x79, 0xba, 0xc6, 0xc2, 0xe7, 0x69, 0x38, 0xfc,
	0xbb, 0x0a, 0xb0, 0x28, 0x0f, 0x65, 0x8a, 0x4e, 0xa1, 0x42, 0x43, 0x6d, 0xca, 0xae, 0x5b, 0xa1,
	0x21, 0x7a, 0x00, 0x15, 0x9a, 0x98, 0x15, 0x6d, 0xd2, 0xb3, 0xdd, 0xb6, 0xd7, 0x81, 0x36, 0x34,
	0x69, 0xe2, 0xe1, 0x30, 0x4c, 0xdd, 0x0a, 0x4d, 0xd0, 0x63, 0x68, 0x68, 0x61, 0x85, 0x59, 0x1d,
	0x54, 0x47, 0x9d, 0x31, 0x3a, 0x36, 0x8f, 0x5b, 0x10, 0x68, 0x02, 0x67, 0xd8, 0x63, 0x44, 0x28,
	0x49, 0xb9, 0xff, 0x86, 0x04, 0xb9, 0xe3, 0x6e, 0x4f, 0xea, 0xe2, 0x99, 0x26, 0xe7, 0x1a, 0x44,
	0x36, 0xd4, 0x13, 0xae, 0x3c, 0x9a, 0xfb, 0xec, 0xee, 0x3e, 0xc3, 0x61, 0xf2, 0x4a, 0x05, 0xf2,
	0x3b, 0x34, 0xbc, 0xe3, 0xe6, 0x1c, 0xfa, 0x12, 0x6a, 0xea, 0x1e, 0x16, 0x26, 0x03, 0x4b, 0x5f,
	0xca, 0xa5, 0x20, 0xe9, 0x1e, 0xd4, 0x71, 0x64, 0x43, 0x53, 0x0d, 0xa3, 0x2c, 0xdc, 0x3c, 0x16,
	0xd2, 0x61, 0x44, 0xe6, 0x42, 0x96, 0x14, 0x9a, 0xc0, 0xbd, 0x38, 0x8b, 0x24, 0xc5, 0x41, 0xc0,
	0x33, 0x26, 0xf5, 0xb3, 0xe9, 0xd1, 0x50, 0x98, 0xad, 0x41, 0x75, 0xd4, 0xdd, 0x5f, 0xe5, 0x8f,
	0x0e, 0x21, 0x35, 0x92, 0x13, 0x8a, 0xc9, 0x62, 0xb7, 0xed, 0x5d, 0xb5, 0x0c, 0xf4, 0x31, 0xd4,
	0x29, 0x93, 0x4f, 0xc7, 0x08, 0x04, 0x09, 0x52, 0x22, 0xd5, 0xed, 0xeb, 0x57, 0x12, 0x03, 0x3d,
	0x85, 0x93, 0x5f, 0x7f, 0x7b, 0xfc, 0x92, 0x53, 0xb6, 0x50, 0xa6, 0x41, 0x79, 0x15, 0xef, 0xb0,
	0x97, 0xe8, 0x9f, 0xfe, 0x79, 0xf2, 0xaf, 0x64, 0xce, 0x74, 0xf8, 0x13, 0xdc, 0xff, 0xf1, 0x46,
	0xb3, 0x23, 0x39, 0x6b, 0x5a, 0xce, 0xcf, 0x00, 0x04, 0x8f, 0x49, 0xe1, 0x78, 0x2d, 0xab, 0xdb,
	0x56, 0x3b, 0xda, 0xda, 0x93, 0xd6, 0x6e, 0xdb, 0xab, 0xb5, 0x8c, 0x81, 0x31, 0x9c, 0x42, 0xef,
	0xb6, 0x92, 0x3c, 0x93, 0xce, 0xf4, 0x46, 0x15, 0xe3, 0xfd, 0x55, 0xbe, 0x82, 0xce, 0xc5, 0x95,
	0xa3, 0x1e, 0x34, 0x55, 0x00, 0xf5, 0xa1, 0x15, 0x70, 0x26, 0x09, 0x93, 0xa2, 0xc8, 0xda, 0xaf,
	0x2f, 0x5f, 0xfc, 0x32, 0xfd, 0xbf, 0x6f, 0xed, 0xe1, 0xe3, 0xfd, 0x6d, 0xb1, 0xe9, 0x37, 0x34,
	0xfd, 0xf4, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x82, 0x2a, 0xa7, 0xb6, 0xe8, 0x06, 0x00, 0x00,
}
