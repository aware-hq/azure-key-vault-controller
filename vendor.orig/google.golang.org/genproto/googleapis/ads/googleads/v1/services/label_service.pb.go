// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/label_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [LabelService.GetLabel][google.ads.googleads.v1.services.LabelService.GetLabel].
type GetLabelRequest struct {
	// The resource name of the label to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLabelRequest) Reset()         { *m = GetLabelRequest{} }
func (m *GetLabelRequest) String() string { return proto.CompactTextString(m) }
func (*GetLabelRequest) ProtoMessage()    {}
func (*GetLabelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_43c34accccca4b63, []int{0}
}

func (m *GetLabelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLabelRequest.Unmarshal(m, b)
}
func (m *GetLabelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLabelRequest.Marshal(b, m, deterministic)
}
func (m *GetLabelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLabelRequest.Merge(m, src)
}
func (m *GetLabelRequest) XXX_Size() int {
	return xxx_messageInfo_GetLabelRequest.Size(m)
}
func (m *GetLabelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLabelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLabelRequest proto.InternalMessageInfo

func (m *GetLabelRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [LabelService.MutateLabels][google.ads.googleads.v1.services.LabelService.MutateLabels].
type MutateLabelsRequest struct {
	// ID of the customer whose labels are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on labels.
	Operations []*LabelOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateLabelsRequest) Reset()         { *m = MutateLabelsRequest{} }
func (m *MutateLabelsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateLabelsRequest) ProtoMessage()    {}
func (*MutateLabelsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_43c34accccca4b63, []int{1}
}

func (m *MutateLabelsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateLabelsRequest.Unmarshal(m, b)
}
func (m *MutateLabelsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateLabelsRequest.Marshal(b, m, deterministic)
}
func (m *MutateLabelsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateLabelsRequest.Merge(m, src)
}
func (m *MutateLabelsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateLabelsRequest.Size(m)
}
func (m *MutateLabelsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateLabelsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateLabelsRequest proto.InternalMessageInfo

func (m *MutateLabelsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateLabelsRequest) GetOperations() []*LabelOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateLabelsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateLabelsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, remove, update) on a label.
type LabelOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*LabelOperation_Create
	//	*LabelOperation_Update
	//	*LabelOperation_Remove
	Operation            isLabelOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *LabelOperation) Reset()         { *m = LabelOperation{} }
func (m *LabelOperation) String() string { return proto.CompactTextString(m) }
func (*LabelOperation) ProtoMessage()    {}
func (*LabelOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_43c34accccca4b63, []int{2}
}

func (m *LabelOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelOperation.Unmarshal(m, b)
}
func (m *LabelOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelOperation.Marshal(b, m, deterministic)
}
func (m *LabelOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelOperation.Merge(m, src)
}
func (m *LabelOperation) XXX_Size() int {
	return xxx_messageInfo_LabelOperation.Size(m)
}
func (m *LabelOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelOperation.DiscardUnknown(m)
}

var xxx_messageInfo_LabelOperation proto.InternalMessageInfo

func (m *LabelOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isLabelOperation_Operation interface {
	isLabelOperation_Operation()
}

type LabelOperation_Create struct {
	Create *resources.Label `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type LabelOperation_Update struct {
	Update *resources.Label `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type LabelOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*LabelOperation_Create) isLabelOperation_Operation() {}

func (*LabelOperation_Update) isLabelOperation_Operation() {}

func (*LabelOperation_Remove) isLabelOperation_Operation() {}

func (m *LabelOperation) GetOperation() isLabelOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *LabelOperation) GetCreate() *resources.Label {
	if x, ok := m.GetOperation().(*LabelOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *LabelOperation) GetUpdate() *resources.Label {
	if x, ok := m.GetOperation().(*LabelOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *LabelOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*LabelOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LabelOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LabelOperation_Create)(nil),
		(*LabelOperation_Update)(nil),
		(*LabelOperation_Remove)(nil),
	}
}

// Response message for a labels mutate.
type MutateLabelsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateLabelResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MutateLabelsResponse) Reset()         { *m = MutateLabelsResponse{} }
func (m *MutateLabelsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateLabelsResponse) ProtoMessage()    {}
func (*MutateLabelsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_43c34accccca4b63, []int{3}
}

func (m *MutateLabelsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateLabelsResponse.Unmarshal(m, b)
}
func (m *MutateLabelsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateLabelsResponse.Marshal(b, m, deterministic)
}
func (m *MutateLabelsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateLabelsResponse.Merge(m, src)
}
func (m *MutateLabelsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateLabelsResponse.Size(m)
}
func (m *MutateLabelsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateLabelsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateLabelsResponse proto.InternalMessageInfo

func (m *MutateLabelsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateLabelsResponse) GetResults() []*MutateLabelResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for a label mutate.
type MutateLabelResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateLabelResult) Reset()         { *m = MutateLabelResult{} }
func (m *MutateLabelResult) String() string { return proto.CompactTextString(m) }
func (*MutateLabelResult) ProtoMessage()    {}
func (*MutateLabelResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_43c34accccca4b63, []int{4}
}

func (m *MutateLabelResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateLabelResult.Unmarshal(m, b)
}
func (m *MutateLabelResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateLabelResult.Marshal(b, m, deterministic)
}
func (m *MutateLabelResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateLabelResult.Merge(m, src)
}
func (m *MutateLabelResult) XXX_Size() int {
	return xxx_messageInfo_MutateLabelResult.Size(m)
}
func (m *MutateLabelResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateLabelResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateLabelResult proto.InternalMessageInfo

func (m *MutateLabelResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetLabelRequest)(nil), "google.ads.googleads.v1.services.GetLabelRequest")
	proto.RegisterType((*MutateLabelsRequest)(nil), "google.ads.googleads.v1.services.MutateLabelsRequest")
	proto.RegisterType((*LabelOperation)(nil), "google.ads.googleads.v1.services.LabelOperation")
	proto.RegisterType((*MutateLabelsResponse)(nil), "google.ads.googleads.v1.services.MutateLabelsResponse")
	proto.RegisterType((*MutateLabelResult)(nil), "google.ads.googleads.v1.services.MutateLabelResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/label_service.proto", fileDescriptor_43c34accccca4b63)
}

var fileDescriptor_43c34accccca4b63 = []byte{
	// 699 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0x36, 0xa9, 0xd4, 0x76, 0xb2, 0xb6, 0x74, 0xaa, 0xb8, 0x2c, 0xa2, 0x4b, 0x2c, 0xb8, 0xac,
	0x98, 0x74, 0xb7, 0x5a, 0x24, 0xa5, 0x87, 0x5d, 0xb0, 0xad, 0x60, 0x6d, 0x49, 0xa1, 0x07, 0x59,
	0x08, 0xd3, 0xcd, 0x74, 0x09, 0x4d, 0x32, 0x71, 0x66, 0xb2, 0x52, 0x4a, 0x2f, 0xfe, 0x05, 0x4f,
	0x5e, 0x3d, 0x0a, 0x1e, 0xfd, 0x13, 0x5e, 0x05, 0x7f, 0x80, 0x78, 0xf2, 0x17, 0x78, 0xf0, 0x20,
	0x33, 0x93, 0x89, 0xbb, 0x95, 0xb2, 0xdd, 0xdb, 0xcb, 0x9b, 0xef, 0xfb, 0xde, 0x37, 0xef, 0xe5,
	0x0d, 0x78, 0x32, 0x20, 0x64, 0x10, 0x63, 0x17, 0x85, 0xcc, 0x55, 0xa1, 0x88, 0x86, 0x2d, 0x97,
	0x61, 0x3a, 0x8c, 0xfa, 0x98, 0xb9, 0x31, 0x3a, 0xc2, 0x71, 0x50, 0x7c, 0x3a, 0x19, 0x25, 0x9c,
	0xc0, 0xba, 0x82, 0x3a, 0x28, 0x64, 0x4e, 0xc9, 0x72, 0x86, 0x2d, 0x47, 0xb3, 0x6a, 0x8f, 0x2f,
	0xd3, 0xa5, 0x98, 0x91, 0x9c, 0x96, 0xc2, 0x4a, 0xb0, 0x76, 0x57, 0xc3, 0xb3, 0xc8, 0x45, 0x69,
	0x4a, 0x38, 0xe2, 0x11, 0x49, 0x59, 0x71, 0x5a, 0x94, 0x73, 0xe5, 0xd7, 0x51, 0x7e, 0xec, 0x1e,
	0x47, 0x38, 0x0e, 0x83, 0x04, 0xb1, 0x93, 0x02, 0x71, 0xef, 0x22, 0xe2, 0x2d, 0x45, 0x59, 0x86,
	0xa9, 0x56, 0xb8, 0x53, 0x9c, 0xd3, 0xac, 0xef, 0x32, 0x8e, 0x78, 0x5e, 0x1c, 0xd8, 0xeb, 0x60,
	0x71, 0x1b, 0xf3, 0x97, 0xc2, 0x8a, 0x8f, 0xdf, 0xe4, 0x98, 0x71, 0xf8, 0x00, 0xdc, 0xd4, 0x26,
	0x83, 0x14, 0x25, 0xb8, 0x6a, 0xd4, 0x8d, 0xc6, 0xbc, 0x5f, 0xd1, 0xc9, 0x57, 0x28, 0xc1, 0xf6,
	0x77, 0x03, 0x2c, 0xef, 0xe6, 0x1c, 0x71, 0x2c, 0xb9, 0x4c, 0x93, 0xef, 0x03, 0xab, 0x9f, 0x33,
	0x4e, 0x12, 0x4c, 0x83, 0x28, 0x2c, 0xa8, 0x40, 0xa7, 0x5e, 0x84, 0x70, 0x1f, 0x00, 0x92, 0x61,
	0xaa, 0xee, 0x57, 0x35, 0xeb, 0x33, 0x0d, 0xab, 0xbd, 0xea, 0x4c, 0xea, 0xa7, 0x23, 0xab, 0xec,
	0x69, 0xa2, 0x3f, 0xa2, 0x01, 0x1f, 0x82, 0xc5, 0x0c, 0x51, 0x1e, 0xa1, 0x38, 0x38, 0x46, 0x51,
	0x9c, 0x53, 0x5c, 0x9d, 0xa9, 0x1b, 0x8d, 0x39, 0x7f, 0xa1, 0x48, 0x6f, 0xa9, 0xac, 0xb8, 0xd8,
	0x10, 0xc5, 0x51, 0x88, 0x38, 0x0e, 0x48, 0x1a, 0x9f, 0x56, 0xaf, 0x4b, 0x58, 0x45, 0x27, 0xf7,
	0xd2, 0xf8, 0xd4, 0xfe, 0x63, 0x80, 0x85, 0xf1, 0x62, 0x70, 0x03, 0x58, 0x79, 0x26, 0x59, 0xa2,
	0xe3, 0x92, 0x65, 0xb5, 0x6b, 0xda, 0xb3, 0x6e, 0xb9, 0xb3, 0x25, 0x86, 0xb2, 0x8b, 0xd8, 0x89,
	0x0f, 0x14, 0x5c, 0xc4, 0xb0, 0x0b, 0x66, 0xfb, 0x14, 0x23, 0xae, 0xda, 0x68, 0xb5, 0x1b, 0x97,
	0xde, 0xb5, 0xfc, 0x33, 0xd4, 0x65, 0x77, 0xae, 0xf9, 0x05, 0x53, 0x68, 0x28, 0xc5, 0xaa, 0x39,
	0xbd, 0x86, 0x62, 0xc2, 0x2a, 0x98, 0xa5, 0x38, 0x21, 0x43, 0xd5, 0x9c, 0x79, 0x71, 0xa2, 0xbe,
	0xbb, 0x16, 0x98, 0x2f, 0xbb, 0x69, 0x7f, 0x36, 0xc0, 0xad, 0xf1, 0xb9, 0xb2, 0x8c, 0xa4, 0x0c,
	0xc3, 0x2d, 0x70, 0xfb, 0x42, 0x97, 0x03, 0x4c, 0x29, 0xa1, 0x52, 0xce, 0x6a, 0x43, 0x6d, 0x89,
	0x66, 0x7d, 0xe7, 0x40, 0xfe, 0x61, 0xfe, 0xf2, 0x78, 0xff, 0x9f, 0x0b, 0x38, 0xdc, 0x05, 0x37,
	0x28, 0x66, 0x79, 0xcc, 0xf5, 0xf0, 0xd7, 0x26, 0x0f, 0x7f, 0xc4, 0x90, 0x2f, 0xb9, 0xbe, 0xd6,
	0xb0, 0x9f, 0x81, 0xa5, 0xff, 0x4e, 0xaf, 0xf4, 0x07, 0xb7, 0x7f, 0x98, 0xa0, 0x22, 0x49, 0x07,
	0xaa, 0x0c, 0xfc, 0x60, 0x80, 0x39, 0xbd, 0x0b, 0xb0, 0x35, 0xd9, 0xd5, 0x85, 0xbd, 0xa9, 0x5d,
	0x79, 0x2a, 0xf6, 0xea, 0xbb, 0x6f, 0x3f, 0xdf, 0x9b, 0x4d, 0xd8, 0x10, 0x0f, 0xc2, 0xd9, 0x98,
	0xd5, 0x4d, 0xbd, 0x2a, 0xcc, 0x6d, 0xaa, 0x17, 0x82, 0xb9, 0xcd, 0x73, 0xf8, 0xc5, 0x00, 0x95,
	0xd1, 0xb1, 0xc0, 0xa7, 0x53, 0x75, 0x4d, 0xaf, 0x67, 0x6d, 0x7d, 0x5a, 0x9a, 0x9a, 0xbe, 0xbd,
	0x2e, 0x1d, 0xaf, 0xda, 0x8f, 0x84, 0xe3, 0x7f, 0x16, 0xcf, 0x46, 0x76, 0x7d, 0xb3, 0x79, 0x5e,
	0x18, 0xf6, 0x12, 0x29, 0xe1, 0x19, 0xcd, 0xee, 0x6f, 0x03, 0xac, 0xf4, 0x49, 0x32, 0xb1, 0x6a,
	0x77, 0x69, 0x74, 0x14, 0xfb, 0x62, 0xa5, 0xf6, 0x8d, 0xd7, 0x3b, 0x05, 0x6d, 0x40, 0x62, 0x94,
	0x0e, 0x1c, 0x42, 0x07, 0xee, 0x00, 0xa7, 0x72, 0xe1, 0xf4, 0xa3, 0x9a, 0x45, 0xec, 0xf2, 0xb7,
	0x7b, 0x43, 0x07, 0x1f, 0xcd, 0x99, 0xed, 0x4e, 0xe7, 0x93, 0x59, 0xdf, 0x56, 0x82, 0x9d, 0x90,
	0x39, 0x2a, 0x14, 0xd1, 0x61, 0xcb, 0x29, 0x0a, 0xb3, 0xaf, 0x1a, 0xd2, 0xeb, 0x84, 0xac, 0x57,
	0x42, 0x7a, 0x87, 0xad, 0x9e, 0x86, 0xfc, 0x32, 0x57, 0x54, 0xde, 0xf3, 0x3a, 0x21, 0xf3, 0xbc,
	0x12, 0xe4, 0x79, 0x87, 0x2d, 0xcf, 0xd3, 0xb0, 0xa3, 0x59, 0xe9, 0x73, 0xed, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x38, 0xc2, 0xb6, 0xf8, 0x62, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LabelServiceClient is the client API for LabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LabelServiceClient interface {
	// Returns the requested label in full detail.
	GetLabel(ctx context.Context, in *GetLabelRequest, opts ...grpc.CallOption) (*resources.Label, error)
	// Creates, updates, or removes labels. Operation statuses are returned.
	MutateLabels(ctx context.Context, in *MutateLabelsRequest, opts ...grpc.CallOption) (*MutateLabelsResponse, error)
}

type labelServiceClient struct {
	cc *grpc.ClientConn
}

func NewLabelServiceClient(cc *grpc.ClientConn) LabelServiceClient {
	return &labelServiceClient{cc}
}

func (c *labelServiceClient) GetLabel(ctx context.Context, in *GetLabelRequest, opts ...grpc.CallOption) (*resources.Label, error) {
	out := new(resources.Label)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.LabelService/GetLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelServiceClient) MutateLabels(ctx context.Context, in *MutateLabelsRequest, opts ...grpc.CallOption) (*MutateLabelsResponse, error) {
	out := new(MutateLabelsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.LabelService/MutateLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LabelServiceServer is the server API for LabelService service.
type LabelServiceServer interface {
	// Returns the requested label in full detail.
	GetLabel(context.Context, *GetLabelRequest) (*resources.Label, error)
	// Creates, updates, or removes labels. Operation statuses are returned.
	MutateLabels(context.Context, *MutateLabelsRequest) (*MutateLabelsResponse, error)
}

func RegisterLabelServiceServer(s *grpc.Server, srv LabelServiceServer) {
	s.RegisterService(&_LabelService_serviceDesc, srv)
}

func _LabelService_GetLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServiceServer).GetLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.LabelService/GetLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServiceServer).GetLabel(ctx, req.(*GetLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelService_MutateLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServiceServer).MutateLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.LabelService/MutateLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServiceServer).MutateLabels(ctx, req.(*MutateLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LabelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.LabelService",
	HandlerType: (*LabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLabel",
			Handler:    _LabelService_GetLabel_Handler,
		},
		{
			MethodName: "MutateLabels",
			Handler:    _LabelService_MutateLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/label_service.proto",
}
