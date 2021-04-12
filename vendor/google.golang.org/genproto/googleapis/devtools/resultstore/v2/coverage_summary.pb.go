// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/resultstore/v2/coverage_summary.proto

package resultstore

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Summary of line coverage
type LineCoverageSummary struct {
	// Number of lines instrumented for coverage.
	InstrumentedLineCount int32 `protobuf:"varint,1,opt,name=instrumented_line_count,json=instrumentedLineCount,proto3" json:"instrumented_line_count,omitempty"`
	// Number of instrumented lines that were executed by the test.
	ExecutedLineCount    int32    `protobuf:"varint,2,opt,name=executed_line_count,json=executedLineCount,proto3" json:"executed_line_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LineCoverageSummary) Reset()         { *m = LineCoverageSummary{} }
func (m *LineCoverageSummary) String() string { return proto.CompactTextString(m) }
func (*LineCoverageSummary) ProtoMessage()    {}
func (*LineCoverageSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_a504af212dd04847, []int{0}
}

func (m *LineCoverageSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LineCoverageSummary.Unmarshal(m, b)
}
func (m *LineCoverageSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LineCoverageSummary.Marshal(b, m, deterministic)
}
func (m *LineCoverageSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LineCoverageSummary.Merge(m, src)
}
func (m *LineCoverageSummary) XXX_Size() int {
	return xxx_messageInfo_LineCoverageSummary.Size(m)
}
func (m *LineCoverageSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_LineCoverageSummary.DiscardUnknown(m)
}

var xxx_messageInfo_LineCoverageSummary proto.InternalMessageInfo

func (m *LineCoverageSummary) GetInstrumentedLineCount() int32 {
	if m != nil {
		return m.InstrumentedLineCount
	}
	return 0
}

func (m *LineCoverageSummary) GetExecutedLineCount() int32 {
	if m != nil {
		return m.ExecutedLineCount
	}
	return 0
}

// Summary of branch coverage
// A branch may be:
//  * not executed.  Counted only in total.
//  * executed but not taken.  Appears in total and executed.
//  * executed and taken.  Appears in all three fields.
type BranchCoverageSummary struct {
	// The number of branches present in the file.
	TotalBranchCount int32 `protobuf:"varint,1,opt,name=total_branch_count,json=totalBranchCount,proto3" json:"total_branch_count,omitempty"`
	// The number of branches executed out of the total branches present.
	// A branch is executed when its condition is evaluated.
	// This is <= total_branch_count as not all branches are executed.
	ExecutedBranchCount int32 `protobuf:"varint,2,opt,name=executed_branch_count,json=executedBranchCount,proto3" json:"executed_branch_count,omitempty"`
	// The number of branches taken out of the total branches executed.
	// A branch is taken when its condition is satisfied.
	// This is <= executed_branch_count as not all executed branches are taken.
	TakenBranchCount     int32    `protobuf:"varint,3,opt,name=taken_branch_count,json=takenBranchCount,proto3" json:"taken_branch_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BranchCoverageSummary) Reset()         { *m = BranchCoverageSummary{} }
func (m *BranchCoverageSummary) String() string { return proto.CompactTextString(m) }
func (*BranchCoverageSummary) ProtoMessage()    {}
func (*BranchCoverageSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_a504af212dd04847, []int{1}
}

func (m *BranchCoverageSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BranchCoverageSummary.Unmarshal(m, b)
}
func (m *BranchCoverageSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BranchCoverageSummary.Marshal(b, m, deterministic)
}
func (m *BranchCoverageSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BranchCoverageSummary.Merge(m, src)
}
func (m *BranchCoverageSummary) XXX_Size() int {
	return xxx_messageInfo_BranchCoverageSummary.Size(m)
}
func (m *BranchCoverageSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_BranchCoverageSummary.DiscardUnknown(m)
}

var xxx_messageInfo_BranchCoverageSummary proto.InternalMessageInfo

func (m *BranchCoverageSummary) GetTotalBranchCount() int32 {
	if m != nil {
		return m.TotalBranchCount
	}
	return 0
}

func (m *BranchCoverageSummary) GetExecutedBranchCount() int32 {
	if m != nil {
		return m.ExecutedBranchCount
	}
	return 0
}

func (m *BranchCoverageSummary) GetTakenBranchCount() int32 {
	if m != nil {
		return m.TakenBranchCount
	}
	return 0
}

// Summary of coverage in each language
type LanguageCoverageSummary struct {
	// This summary is for all files written in this programming language.
	Language Language `protobuf:"varint,1,opt,name=language,proto3,enum=google.devtools.resultstore.v2.Language" json:"language,omitempty"`
	// Summary of lines covered vs instrumented.
	LineSummary *LineCoverageSummary `protobuf:"bytes,2,opt,name=line_summary,json=lineSummary,proto3" json:"line_summary,omitempty"`
	// Summary of branch coverage.
	BranchSummary        *BranchCoverageSummary `protobuf:"bytes,3,opt,name=branch_summary,json=branchSummary,proto3" json:"branch_summary,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *LanguageCoverageSummary) Reset()         { *m = LanguageCoverageSummary{} }
func (m *LanguageCoverageSummary) String() string { return proto.CompactTextString(m) }
func (*LanguageCoverageSummary) ProtoMessage()    {}
func (*LanguageCoverageSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_a504af212dd04847, []int{2}
}

func (m *LanguageCoverageSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LanguageCoverageSummary.Unmarshal(m, b)
}
func (m *LanguageCoverageSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LanguageCoverageSummary.Marshal(b, m, deterministic)
}
func (m *LanguageCoverageSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LanguageCoverageSummary.Merge(m, src)
}
func (m *LanguageCoverageSummary) XXX_Size() int {
	return xxx_messageInfo_LanguageCoverageSummary.Size(m)
}
func (m *LanguageCoverageSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_LanguageCoverageSummary.DiscardUnknown(m)
}

var xxx_messageInfo_LanguageCoverageSummary proto.InternalMessageInfo

func (m *LanguageCoverageSummary) GetLanguage() Language {
	if m != nil {
		return m.Language
	}
	return Language_LANGUAGE_UNSPECIFIED
}

func (m *LanguageCoverageSummary) GetLineSummary() *LineCoverageSummary {
	if m != nil {
		return m.LineSummary
	}
	return nil
}

func (m *LanguageCoverageSummary) GetBranchSummary() *BranchCoverageSummary {
	if m != nil {
		return m.BranchSummary
	}
	return nil
}

func init() {
	proto.RegisterType((*LineCoverageSummary)(nil), "google.devtools.resultstore.v2.LineCoverageSummary")
	proto.RegisterType((*BranchCoverageSummary)(nil), "google.devtools.resultstore.v2.BranchCoverageSummary")
	proto.RegisterType((*LanguageCoverageSummary)(nil), "google.devtools.resultstore.v2.LanguageCoverageSummary")
}

func init() {
	proto.RegisterFile("google/devtools/resultstore/v2/coverage_summary.proto", fileDescriptor_a504af212dd04847)
}

var fileDescriptor_a504af212dd04847 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x4b, 0xeb, 0x40,
	0x14, 0x85, 0x49, 0xcb, 0x7b, 0x3c, 0xa6, 0xef, 0x95, 0x67, 0x4a, 0xa9, 0xb8, 0x10, 0xc9, 0xaa,
	0xa0, 0x4c, 0x20, 0xa5, 0x6e, 0xdc, 0x55, 0x37, 0x42, 0x17, 0x12, 0xc1, 0x85, 0x08, 0x61, 0x9a,
	0x5e, 0xc6, 0x60, 0x32, 0xb7, 0x4e, 0x26, 0x41, 0x17, 0xfe, 0x0c, 0x7f, 0x83, 0x7f, 0x53, 0x32,
	0x33, 0x09, 0x69, 0xab, 0x66, 0x39, 0x9c, 0x7b, 0xbe, 0x73, 0x66, 0xee, 0x90, 0x39, 0x47, 0xe4,
	0x29, 0xf8, 0x6b, 0x28, 0x15, 0x62, 0x9a, 0xfb, 0x12, 0xf2, 0x22, 0x55, 0xb9, 0x42, 0x09, 0x7e,
	0x19, 0xf8, 0x31, 0x96, 0x20, 0x19, 0x87, 0x28, 0x2f, 0xb2, 0x8c, 0xc9, 0x57, 0xba, 0x91, 0xa8,
	0xd0, 0x3d, 0x36, 0x36, 0x5a, 0xdb, 0x68, 0xcb, 0x46, 0xcb, 0xe0, 0xe8, 0xb4, 0x13, 0x9b, 0x65,
	0x28, 0x0c, 0xcc, 0x7b, 0x23, 0xa3, 0x65, 0x22, 0xe0, 0xd2, 0x46, 0xdd, 0x9a, 0x24, 0xf7, 0x9c,
	0x4c, 0x12, 0x91, 0x2b, 0x59, 0x64, 0x20, 0x14, 0xac, 0xa3, 0x34, 0x11, 0x10, 0xc5, 0x58, 0x08,
	0x75, 0xe8, 0x9c, 0x38, 0xd3, 0x5f, 0xe1, 0xb8, 0x2d, 0x1b, 0x42, 0x21, 0x94, 0x4b, 0xc9, 0x08,
	0x5e, 0x20, 0x2e, 0x76, 0x3c, 0x3d, 0xed, 0x39, 0xa8, 0xa5, 0x66, 0xde, 0xfb, 0x70, 0xc8, 0x78,
	0x21, 0x99, 0x88, 0x1f, 0x77, 0x1b, 0x9c, 0x11, 0x57, 0xa1, 0x62, 0x69, 0xb4, 0xd2, 0xf2, 0x56,
	0xf8, 0x7f, 0xad, 0xd4, 0xbe, 0x2a, 0x37, 0x20, 0xe3, 0x26, 0x77, 0xcb, 0x60, 0x92, 0x9b, 0x52,
	0x6d, 0x4f, 0x95, 0xc0, 0x9e, 0x40, 0x6c, 0x1b, 0xfa, 0x36, 0xa1, 0x52, 0x5a, 0xd3, 0xde, 0x7b,
	0x8f, 0x4c, 0x96, 0x4c, 0xf0, 0x82, 0xf1, 0xbd, 0xd7, 0xba, 0x22, 0x7f, 0x52, 0x2b, 0xe9, 0x86,
	0xc3, 0x60, 0x4a, 0x7f, 0x5e, 0x12, 0xad, 0x51, 0x61, 0xe3, 0x74, 0xef, 0xc8, 0x5f, 0xfd, 0x64,
	0x76, 0xdb, 0xba, 0xfa, 0x20, 0x98, 0x75, 0x92, 0xf6, 0xd7, 0x17, 0x0e, 0x2a, 0x50, 0xdd, 0xee,
	0x81, 0x0c, 0xed, 0x0d, 0x6b, 0x72, 0x5f, 0x93, 0xe7, 0x5d, 0xe4, 0x2f, 0x17, 0x13, 0xfe, 0x33,
	0x30, 0x7b, 0x5c, 0x3c, 0x13, 0x2f, 0xc6, 0xac, 0x03, 0x75, 0xe3, 0xdc, 0x5f, 0xdb, 0x09, 0x8e,
	0xd5, 0x85, 0x29, 0x4a, 0xee, 0x73, 0x10, 0xfa, 0x13, 0xfa, 0x46, 0x62, 0x9b, 0x24, 0xff, 0xee,
	0xd3, 0x5e, 0xb4, 0x8e, 0xab, 0xdf, 0xda, 0x35, 0xfb, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x65,
	0x24, 0xb5, 0x40, 0x03, 0x00, 0x00,
}
