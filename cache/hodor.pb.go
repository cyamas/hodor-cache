// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: hodor.proto

package cache

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

type SetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Types that are assignable to Value:
	//
	//	*SetRequest_StrVal
	//	*SetRequest_IntVal
	//	*SetRequest_FloatVal
	//	*SetRequest_BoolVal
	//	*SetRequest_StrArr
	//	*SetRequest_IntArr
	//	*SetRequest_FloatArr
	//	*SetRequest_BoolArr
	Value isSetRequest_Value `protobuf_oneof:"value"`
}

func (x *SetRequest) Reset() {
	*x = SetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hodor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRequest) ProtoMessage() {}

func (x *SetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hodor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRequest.ProtoReflect.Descriptor instead.
func (*SetRequest) Descriptor() ([]byte, []int) {
	return file_hodor_proto_rawDescGZIP(), []int{0}
}

func (x *SetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (m *SetRequest) GetValue() isSetRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *SetRequest) GetStrVal() string {
	if x, ok := x.GetValue().(*SetRequest_StrVal); ok {
		return x.StrVal
	}
	return ""
}

func (x *SetRequest) GetIntVal() int64 {
	if x, ok := x.GetValue().(*SetRequest_IntVal); ok {
		return x.IntVal
	}
	return 0
}

func (x *SetRequest) GetFloatVal() float64 {
	if x, ok := x.GetValue().(*SetRequest_FloatVal); ok {
		return x.FloatVal
	}
	return 0
}

func (x *SetRequest) GetBoolVal() bool {
	if x, ok := x.GetValue().(*SetRequest_BoolVal); ok {
		return x.BoolVal
	}
	return false
}

func (x *SetRequest) GetStrArr() *StringArray {
	if x, ok := x.GetValue().(*SetRequest_StrArr); ok {
		return x.StrArr
	}
	return nil
}

func (x *SetRequest) GetIntArr() *IntArray {
	if x, ok := x.GetValue().(*SetRequest_IntArr); ok {
		return x.IntArr
	}
	return nil
}

func (x *SetRequest) GetFloatArr() *FloatArray {
	if x, ok := x.GetValue().(*SetRequest_FloatArr); ok {
		return x.FloatArr
	}
	return nil
}

func (x *SetRequest) GetBoolArr() *BoolArray {
	if x, ok := x.GetValue().(*SetRequest_BoolArr); ok {
		return x.BoolArr
	}
	return nil
}

type isSetRequest_Value interface {
	isSetRequest_Value()
}

type SetRequest_StrVal struct {
	StrVal string `protobuf:"bytes,2,opt,name=strVal,proto3,oneof"`
}

type SetRequest_IntVal struct {
	IntVal int64 `protobuf:"varint,3,opt,name=intVal,proto3,oneof"`
}

type SetRequest_FloatVal struct {
	FloatVal float64 `protobuf:"fixed64,4,opt,name=floatVal,proto3,oneof"`
}

type SetRequest_BoolVal struct {
	BoolVal bool `protobuf:"varint,5,opt,name=boolVal,proto3,oneof"`
}

type SetRequest_StrArr struct {
	StrArr *StringArray `protobuf:"bytes,6,opt,name=strArr,proto3,oneof"`
}

type SetRequest_IntArr struct {
	IntArr *IntArray `protobuf:"bytes,7,opt,name=intArr,proto3,oneof"`
}

type SetRequest_FloatArr struct {
	FloatArr *FloatArray `protobuf:"bytes,8,opt,name=floatArr,proto3,oneof"`
}

type SetRequest_BoolArr struct {
	BoolArr *BoolArray `protobuf:"bytes,9,opt,name=boolArr,proto3,oneof"`
}

func (*SetRequest_StrVal) isSetRequest_Value() {}

func (*SetRequest_IntVal) isSetRequest_Value() {}

func (*SetRequest_FloatVal) isSetRequest_Value() {}

func (*SetRequest_BoolVal) isSetRequest_Value() {}

func (*SetRequest_StrArr) isSetRequest_Value() {}

func (*SetRequest_IntArr) isSetRequest_Value() {}

func (*SetRequest_FloatArr) isSetRequest_Value() {}

func (*SetRequest_BoolArr) isSetRequest_Value() {}

type SetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stub string `protobuf:"bytes,1,opt,name=stub,proto3" json:"stub,omitempty"`
}

func (x *SetResponse) Reset() {
	*x = SetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hodor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetResponse) ProtoMessage() {}

func (x *SetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hodor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetResponse.ProtoReflect.Descriptor instead.
func (*SetResponse) Descriptor() ([]byte, []int) {
	return file_hodor_proto_rawDescGZIP(), []int{1}
}

func (x *SetResponse) GetStub() string {
	if x != nil {
		return x.Stub
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hodor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hodor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_hodor_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*GetResponse_StrVal
	//	*GetResponse_IntVal
	//	*GetResponse_FloatVal
	//	*GetResponse_BoolVal
	//	*GetResponse_StrArr
	//	*GetResponse_IntArr
	//	*GetResponse_FloatArr
	//	*GetResponse_BoolArr
	Value isGetResponse_Value `protobuf_oneof:"value"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hodor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hodor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_hodor_proto_rawDescGZIP(), []int{3}
}

func (m *GetResponse) GetValue() isGetResponse_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *GetResponse) GetStrVal() string {
	if x, ok := x.GetValue().(*GetResponse_StrVal); ok {
		return x.StrVal
	}
	return ""
}

func (x *GetResponse) GetIntVal() int64 {
	if x, ok := x.GetValue().(*GetResponse_IntVal); ok {
		return x.IntVal
	}
	return 0
}

func (x *GetResponse) GetFloatVal() float64 {
	if x, ok := x.GetValue().(*GetResponse_FloatVal); ok {
		return x.FloatVal
	}
	return 0
}

func (x *GetResponse) GetBoolVal() bool {
	if x, ok := x.GetValue().(*GetResponse_BoolVal); ok {
		return x.BoolVal
	}
	return false
}

func (x *GetResponse) GetStrArr() *StringArray {
	if x, ok := x.GetValue().(*GetResponse_StrArr); ok {
		return x.StrArr
	}
	return nil
}

func (x *GetResponse) GetIntArr() *IntArray {
	if x, ok := x.GetValue().(*GetResponse_IntArr); ok {
		return x.IntArr
	}
	return nil
}

func (x *GetResponse) GetFloatArr() *FloatArray {
	if x, ok := x.GetValue().(*GetResponse_FloatArr); ok {
		return x.FloatArr
	}
	return nil
}

func (x *GetResponse) GetBoolArr() *BoolArray {
	if x, ok := x.GetValue().(*GetResponse_BoolArr); ok {
		return x.BoolArr
	}
	return nil
}

type isGetResponse_Value interface {
	isGetResponse_Value()
}

type GetResponse_StrVal struct {
	StrVal string `protobuf:"bytes,1,opt,name=strVal,proto3,oneof"`
}

type GetResponse_IntVal struct {
	IntVal int64 `protobuf:"varint,2,opt,name=intVal,proto3,oneof"`
}

type GetResponse_FloatVal struct {
	FloatVal float64 `protobuf:"fixed64,3,opt,name=floatVal,proto3,oneof"`
}

type GetResponse_BoolVal struct {
	BoolVal bool `protobuf:"varint,4,opt,name=boolVal,proto3,oneof"`
}

type GetResponse_StrArr struct {
	StrArr *StringArray `protobuf:"bytes,5,opt,name=strArr,proto3,oneof"`
}

type GetResponse_IntArr struct {
	IntArr *IntArray `protobuf:"bytes,6,opt,name=intArr,proto3,oneof"`
}

type GetResponse_FloatArr struct {
	FloatArr *FloatArray `protobuf:"bytes,7,opt,name=floatArr,proto3,oneof"`
}

type GetResponse_BoolArr struct {
	BoolArr *BoolArray `protobuf:"bytes,8,opt,name=boolArr,proto3,oneof"`
}

func (*GetResponse_StrVal) isGetResponse_Value() {}

func (*GetResponse_IntVal) isGetResponse_Value() {}

func (*GetResponse_FloatVal) isGetResponse_Value() {}

func (*GetResponse_BoolVal) isGetResponse_Value() {}

func (*GetResponse_StrArr) isGetResponse_Value() {}

func (*GetResponse_IntArr) isGetResponse_Value() {}

func (*GetResponse_FloatArr) isGetResponse_Value() {}

func (*GetResponse_BoolArr) isGetResponse_Value() {}

type DelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DelRequest) Reset() {
	*x = DelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hodor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelRequest) ProtoMessage() {}

func (x *DelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hodor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelRequest.ProtoReflect.Descriptor instead.
func (*DelRequest) Descriptor() ([]byte, []int) {
	return file_hodor_proto_rawDescGZIP(), []int{4}
}

func (x *DelRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stub string `protobuf:"bytes,1,opt,name=stub,proto3" json:"stub,omitempty"`
}

func (x *DelResponse) Reset() {
	*x = DelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hodor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelResponse) ProtoMessage() {}

func (x *DelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hodor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelResponse.ProtoReflect.Descriptor instead.
func (*DelResponse) Descriptor() ([]byte, []int) {
	return file_hodor_proto_rawDescGZIP(), []int{5}
}

func (x *DelResponse) GetStub() string {
	if x != nil {
		return x.Stub
	}
	return ""
}

type StringArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val []string `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty"`
}

func (x *StringArray) Reset() {
	*x = StringArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hodor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringArray) ProtoMessage() {}

func (x *StringArray) ProtoReflect() protoreflect.Message {
	mi := &file_hodor_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringArray.ProtoReflect.Descriptor instead.
func (*StringArray) Descriptor() ([]byte, []int) {
	return file_hodor_proto_rawDescGZIP(), []int{6}
}

func (x *StringArray) GetVal() []string {
	if x != nil {
		return x.Val
	}
	return nil
}

type IntArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val []int64 `protobuf:"varint,1,rep,packed,name=val,proto3" json:"val,omitempty"`
}

func (x *IntArray) Reset() {
	*x = IntArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hodor_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntArray) ProtoMessage() {}

func (x *IntArray) ProtoReflect() protoreflect.Message {
	mi := &file_hodor_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntArray.ProtoReflect.Descriptor instead.
func (*IntArray) Descriptor() ([]byte, []int) {
	return file_hodor_proto_rawDescGZIP(), []int{7}
}

func (x *IntArray) GetVal() []int64 {
	if x != nil {
		return x.Val
	}
	return nil
}

type FloatArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val []float64 `protobuf:"fixed64,1,rep,packed,name=val,proto3" json:"val,omitempty"`
}

func (x *FloatArray) Reset() {
	*x = FloatArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hodor_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FloatArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FloatArray) ProtoMessage() {}

func (x *FloatArray) ProtoReflect() protoreflect.Message {
	mi := &file_hodor_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FloatArray.ProtoReflect.Descriptor instead.
func (*FloatArray) Descriptor() ([]byte, []int) {
	return file_hodor_proto_rawDescGZIP(), []int{8}
}

func (x *FloatArray) GetVal() []float64 {
	if x != nil {
		return x.Val
	}
	return nil
}

type BoolArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val []bool `protobuf:"varint,1,rep,packed,name=val,proto3" json:"val,omitempty"`
}

func (x *BoolArray) Reset() {
	*x = BoolArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hodor_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolArray) ProtoMessage() {}

func (x *BoolArray) ProtoReflect() protoreflect.Message {
	mi := &file_hodor_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolArray.ProtoReflect.Descriptor instead.
func (*BoolArray) Descriptor() ([]byte, []int) {
	return file_hodor_proto_rawDescGZIP(), []int{9}
}

func (x *BoolArray) GetVal() []bool {
	if x != nil {
		return x.Val
	}
	return nil
}

var File_hodor_proto protoreflect.FileDescriptor

var file_hodor_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x68, 0x6f, 0x64, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x02,
	0x0a, 0x0a, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x18,
	0x0a, 0x06, 0x73, 0x74, 0x72, 0x56, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x06, 0x73, 0x74, 0x72, 0x56, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x56,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x56,
	0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c,
	0x12, 0x1a, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x00, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x12, 0x26, 0x0a, 0x06,
	0x73, 0x74, 0x72, 0x41, 0x72, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x41, 0x72, 0x72, 0x61, 0x79, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74,
	0x72, 0x41, 0x72, 0x72, 0x12, 0x23, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x41, 0x72, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x6e, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x48,
	0x00, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x41, 0x72, 0x72, 0x12, 0x29, 0x0a, 0x08, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x41, 0x72, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x46, 0x6c,
	0x6f, 0x61, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x48, 0x00, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x61,
	0x74, 0x41, 0x72, 0x72, 0x12, 0x26, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x41, 0x72, 0x72, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x41, 0x72, 0x72, 0x61,
	0x79, 0x48, 0x00, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x41, 0x72, 0x72, 0x42, 0x07, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x21, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x75, 0x62, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x74, 0x75, 0x62, 0x22, 0x1e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xa4, 0x02, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x56,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x56,
	0x61, 0x6c, 0x12, 0x18, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x00, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x08,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00,
	0x52, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x07, 0x62, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x62,
	0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x41, 0x72, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x41,
	0x72, 0x72, 0x61, 0x79, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x41, 0x72, 0x72, 0x12, 0x23,
	0x0a, 0x06, 0x69, 0x6e, 0x74, 0x41, 0x72, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x49, 0x6e, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x48, 0x00, 0x52, 0x06, 0x69, 0x6e, 0x74,
	0x41, 0x72, 0x72, 0x12, 0x29, 0x0a, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x41, 0x72, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x41, 0x72, 0x72,
	0x61, 0x79, 0x48, 0x00, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x41, 0x72, 0x72, 0x12, 0x26,
	0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x41, 0x72, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x41, 0x72, 0x72, 0x61, 0x79, 0x48, 0x00, 0x52, 0x07, 0x62,
	0x6f, 0x6f, 0x6c, 0x41, 0x72, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0x21, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x74, 0x75, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x74,
	0x75, 0x62, 0x22, 0x1f, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x41, 0x72, 0x72, 0x61,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03,
	0x76, 0x61, 0x6c, 0x22, 0x1c, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x76, 0x61,
	0x6c, 0x22, 0x1e, 0x0a, 0x0a, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x01, 0x52, 0x03, 0x76, 0x61,
	0x6c, 0x22, 0x1d, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x6c, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x08, 0x52, 0x03, 0x76, 0x61, 0x6c,
	0x32, 0x6d, 0x0a, 0x05, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x20, 0x0a, 0x03, 0x53, 0x65, 0x74,
	0x12, 0x0b, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e,
	0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x0b, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0c, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a,
	0x03, 0x44, 0x65, 0x6c, 0x12, 0x0b, 0x2e, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0c, 0x2e, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x79,
	0x61, 0x6d, 0x61, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x68,
	0x6f, 0x64, 0x6f, 0x72, 0x2d, 0x63, 0x61, 0x63, 0x68, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_hodor_proto_rawDescOnce sync.Once
	file_hodor_proto_rawDescData = file_hodor_proto_rawDesc
)

func file_hodor_proto_rawDescGZIP() []byte {
	file_hodor_proto_rawDescOnce.Do(func() {
		file_hodor_proto_rawDescData = protoimpl.X.CompressGZIP(file_hodor_proto_rawDescData)
	})
	return file_hodor_proto_rawDescData
}

var file_hodor_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_hodor_proto_goTypes = []any{
	(*SetRequest)(nil),  // 0: SetRequest
	(*SetResponse)(nil), // 1: SetResponse
	(*GetRequest)(nil),  // 2: GetRequest
	(*GetResponse)(nil), // 3: GetResponse
	(*DelRequest)(nil),  // 4: DelRequest
	(*DelResponse)(nil), // 5: DelResponse
	(*StringArray)(nil), // 6: StringArray
	(*IntArray)(nil),    // 7: IntArray
	(*FloatArray)(nil),  // 8: FloatArray
	(*BoolArray)(nil),   // 9: BoolArray
}
var file_hodor_proto_depIdxs = []int32{
	6,  // 0: SetRequest.strArr:type_name -> StringArray
	7,  // 1: SetRequest.intArr:type_name -> IntArray
	8,  // 2: SetRequest.floatArr:type_name -> FloatArray
	9,  // 3: SetRequest.boolArr:type_name -> BoolArray
	6,  // 4: GetResponse.strArr:type_name -> StringArray
	7,  // 5: GetResponse.intArr:type_name -> IntArray
	8,  // 6: GetResponse.floatArr:type_name -> FloatArray
	9,  // 7: GetResponse.boolArr:type_name -> BoolArray
	0,  // 8: Cache.Set:input_type -> SetRequest
	2,  // 9: Cache.Get:input_type -> GetRequest
	4,  // 10: Cache.Del:input_type -> DelRequest
	1,  // 11: Cache.Set:output_type -> SetResponse
	3,  // 12: Cache.Get:output_type -> GetResponse
	5,  // 13: Cache.Del:output_type -> DelResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_hodor_proto_init() }
func file_hodor_proto_init() {
	if File_hodor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hodor_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SetRequest); i {
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
		file_hodor_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SetResponse); i {
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
		file_hodor_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetRequest); i {
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
		file_hodor_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetResponse); i {
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
		file_hodor_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DelRequest); i {
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
		file_hodor_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DelResponse); i {
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
		file_hodor_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*StringArray); i {
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
		file_hodor_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*IntArray); i {
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
		file_hodor_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*FloatArray); i {
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
		file_hodor_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*BoolArray); i {
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
	file_hodor_proto_msgTypes[0].OneofWrappers = []any{
		(*SetRequest_StrVal)(nil),
		(*SetRequest_IntVal)(nil),
		(*SetRequest_FloatVal)(nil),
		(*SetRequest_BoolVal)(nil),
		(*SetRequest_StrArr)(nil),
		(*SetRequest_IntArr)(nil),
		(*SetRequest_FloatArr)(nil),
		(*SetRequest_BoolArr)(nil),
	}
	file_hodor_proto_msgTypes[3].OneofWrappers = []any{
		(*GetResponse_StrVal)(nil),
		(*GetResponse_IntVal)(nil),
		(*GetResponse_FloatVal)(nil),
		(*GetResponse_BoolVal)(nil),
		(*GetResponse_StrArr)(nil),
		(*GetResponse_IntArr)(nil),
		(*GetResponse_FloatArr)(nil),
		(*GetResponse_BoolArr)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hodor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hodor_proto_goTypes,
		DependencyIndexes: file_hodor_proto_depIdxs,
		MessageInfos:      file_hodor_proto_msgTypes,
	}.Build()
	File_hodor_proto = out.File
	file_hodor_proto_rawDesc = nil
	file_hodor_proto_goTypes = nil
	file_hodor_proto_depIdxs = nil
}
