// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rebus/mint/v1/mint.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/gogo/protobuf/gogoproto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Minter represents the minting state.
type Minter struct {
	// current annual inflation rate
	Inflation       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=inflation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation"`
	Phase           uint64                                 `protobuf:"varint,2,opt,name=phase,proto3" json:"phase,omitempty"`
	StartPhaseBlock uint64                                 `protobuf:"varint,3,opt,name=start_phase_block,json=startPhaseBlock,proto3" json:"start_phase_block,omitempty" yaml:"start_phase_block"`
	// current annual expected provisions
	AnnualProvisions github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=annual_provisions,json=annualProvisions,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"annual_provisions" yaml:"annual_provisions"`
}

func (m *Minter) Reset()      { *m = Minter{} }
func (*Minter) ProtoMessage() {}
func (*Minter) Descriptor() ([]byte, []int) {
	return fileDescriptor_d55efb74470ccb48, []int{0}
}
func (m *Minter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Minter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Minter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Minter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minter.Merge(m, src)
}
func (m *Minter) XXX_Size() int {
	return m.Size()
}
func (m *Minter) XXX_DiscardUnknown() {
	xxx_messageInfo_Minter.DiscardUnknown(m)
}

var xxx_messageInfo_Minter proto.InternalMessageInfo

func (m *Minter) GetPhase() uint64 {
	if m != nil {
		return m.Phase
	}
	return 0
}

func (m *Minter) GetStartPhaseBlock() uint64 {
	if m != nil {
		return m.StartPhaseBlock
	}
	return 0
}

// Params holds parameters for the mint module.
type Params struct {
	// type of coin to mint
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// expected blocks per year
	BlocksPerYear uint64 `protobuf:"varint,2,opt,name=blocks_per_year,json=blocksPerYear,proto3" json:"blocks_per_year,omitempty" yaml:"blocks_per_year"`
	// parameter to enable and disble the minter
	MintEnabled bool `protobuf:"varint,3,opt,name=mint_enabled,json=mintEnabled,proto3" json:"mint_enabled,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_d55efb74470ccb48, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetBlocksPerYear() uint64 {
	if m != nil {
		return m.BlocksPerYear
	}
	return 0
}

func (m *Params) GetMintEnabled() bool {
	if m != nil {
		return m.MintEnabled
	}
	return false
}

func init() {
	proto.RegisterType((*Minter)(nil), "rebus.mint.v1.Minter")
	proto.RegisterType((*Params)(nil), "rebus.mint.v1.Params")
}

func init() { proto.RegisterFile("rebus/mint/v1/mint.proto", fileDescriptor_d55efb74470ccb48) }

var fileDescriptor_d55efb74470ccb48 = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0x6f, 0xd4, 0x30,
	0x14, 0xc7, 0xed, 0x72, 0x9c, 0x38, 0x43, 0x55, 0x6a, 0x55, 0x28, 0xaa, 0xc0, 0x29, 0x19, 0x50,
	0x17, 0x12, 0x2a, 0xb6, 0x8e, 0x51, 0x91, 0x10, 0x02, 0x29, 0xca, 0x06, 0x4b, 0xe4, 0xe4, 0xdc,
	0x3b, 0xab, 0x89, 0x1d, 0xd9, 0xbe, 0x83, 0xdb, 0xf8, 0x08, 0x0c, 0x0c, 0x8c, 0x8c, 0x6c, 0x7c,
	0x8d, 0x8e, 0x37, 0x56, 0x0c, 0x11, 0x97, 0x5b, 0x18, 0xab, 0x7e, 0x02, 0x64, 0x1b, 0xe9, 0x10,
	0x37, 0x31, 0xd9, 0xef, 0xf7, 0xfe, 0x7a, 0xef, 0xe9, 0xaf, 0x3f, 0x0a, 0x14, 0x2b, 0x67, 0x3a,
	0x69, 0xb8, 0x30, 0xc9, 0xfc, 0xc4, 0xbd, 0x71, 0xab, 0xa4, 0x91, 0x78, 0xd7, 0x75, 0x62, 0x47,
	0xe6, 0x27, 0x87, 0x07, 0x13, 0x39, 0x91, 0xae, 0x93, 0xd8, 0x9f, 0x17, 0x45, 0xdf, 0x77, 0xd0,
	0xf0, 0x0d, 0x17, 0x86, 0x29, 0xfc, 0x1a, 0x8d, 0xb8, 0x38, 0xaf, 0xa9, 0xe1, 0x52, 0x04, 0xf0,
	0x08, 0x1e, 0x8f, 0xd2, 0xf8, 0xb2, 0x0b, 0xc1, 0x8f, 0x2e, 0x7c, 0x32, 0xe1, 0x66, 0x3a, 0x2b,
	0xe3, 0x4a, 0x36, 0x49, 0x25, 0x75, 0x23, 0xf5, 0x9f, 0xe7, 0xa9, 0x1e, 0x5f, 0x24, 0x66, 0xd1,
	0x32, 0x1d, 0x9f, 0xb1, 0x2a, 0xdf, 0x0c, 0xc0, 0x07, 0xe8, 0x76, 0x3b, 0xa5, 0x9a, 0x05, 0x3b,
	0x47, 0xf0, 0x78, 0x90, 0xfb, 0x02, 0xbf, 0x44, 0xfb, 0xda, 0x50, 0x65, 0x0a, 0x57, 0x16, 0x65,
	0x2d, 0xab, 0x8b, 0xe0, 0x96, 0x55, 0xa4, 0x0f, 0x6f, 0xba, 0x30, 0x58, 0xd0, 0xa6, 0x3e, 0x8d,
	0xb6, 0x24, 0x51, 0xbe, 0xe7, 0x58, 0x66, 0x51, 0x6a, 0x09, 0x7e, 0x8f, 0xf6, 0xa9, 0x10, 0x33,
	0x5a, 0x17, 0xad, 0x92, 0x73, 0xae, 0xb9, 0x14, 0x3a, 0x18, 0xb8, 0xab, 0x5f, 0xfd, 0xdf, 0xd5,
	0x9b, 0xbd, 0x5b, 0x03, 0xa3, 0xfc, 0xbe, 0x67, 0xd9, 0x06, 0x7d, 0x86, 0x68, 0x98, 0x51, 0x45,
	0x1b, 0x8d, 0x1f, 0x21, 0x64, 0xdd, 0x2d, 0xc6, 0x4c, 0xc8, 0xc6, 0x5b, 0x96, 0x8f, 0x2c, 0x39,
	0xb3, 0x00, 0xa7, 0x68, 0xcf, 0x5d, 0xaf, 0x8b, 0x96, 0xa9, 0x62, 0xc1, 0xa8, 0xf2, 0x66, 0xa4,
	0x87, 0x37, 0x5d, 0xf8, 0xc0, 0xaf, 0xfc, 0x47, 0x10, 0xe5, 0xbb, 0x9e, 0x64, 0x4c, 0xbd, 0x65,
	0x54, 0xe1, 0xc7, 0xe8, 0x9e, 0x5b, 0xc1, 0x04, 0x2d, 0x6b, 0x36, 0x76, 0x5e, 0xdd, 0xc9, 0xef,
	0x5a, 0xf6, 0xc2, 0xa3, 0xd3, 0xc1, 0x97, 0xaf, 0x21, 0x48, 0xcf, 0x97, 0x2b, 0x02, 0xae, 0x56,
	0x04, 0x5c, 0xaf, 0x08, 0xfc, 0xd8, 0x13, 0xf8, 0xad, 0x27, 0xf0, 0xb2, 0x27, 0x70, 0xd9, 0x13,
	0xf8, 0xb3, 0x27, 0xf0, 0x57, 0x4f, 0xc0, 0x75, 0x4f, 0xe0, 0xa7, 0x35, 0x01, 0xcb, 0x35, 0x01,
	0x57, 0x6b, 0x02, 0xde, 0x3d, 0xfb, 0xcb, 0x22, 0x17, 0x97, 0x6a, 0x4a, 0xb9, 0xf0, 0xdf, 0xb8,
	0x92, 0x8a, 0xd9, 0x4c, 0x7d, 0xf0, 0xe9, 0x72, 0x86, 0x95, 0x43, 0x97, 0x9b, 0xe7, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x8f, 0x18, 0x2e, 0x87, 0x78, 0x02, 0x00, 0x00,
}

func (this *Minter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Minter)
	if !ok {
		that2, ok := that.(Minter)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Inflation.Equal(that1.Inflation) {
		return false
	}
	if this.Phase != that1.Phase {
		return false
	}
	if this.StartPhaseBlock != that1.StartPhaseBlock {
		return false
	}
	if !this.AnnualProvisions.Equal(that1.AnnualProvisions) {
		return false
	}
	return true
}
func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.MintDenom != that1.MintDenom {
		return false
	}
	if this.BlocksPerYear != that1.BlocksPerYear {
		return false
	}
	if this.MintEnabled != that1.MintEnabled {
		return false
	}
	return true
}
func (this *Minter) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&types.Minter{")
	s = append(s, "Inflation: "+fmt.Sprintf("%#v", this.Inflation)+",\n")
	s = append(s, "Phase: "+fmt.Sprintf("%#v", this.Phase)+",\n")
	s = append(s, "StartPhaseBlock: "+fmt.Sprintf("%#v", this.StartPhaseBlock)+",\n")
	s = append(s, "AnnualProvisions: "+fmt.Sprintf("%#v", this.AnnualProvisions)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Params) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&types.Params{")
	s = append(s, "MintDenom: "+fmt.Sprintf("%#v", this.MintDenom)+",\n")
	s = append(s, "BlocksPerYear: "+fmt.Sprintf("%#v", this.BlocksPerYear)+",\n")
	s = append(s, "MintEnabled: "+fmt.Sprintf("%#v", this.MintEnabled)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMint(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Minter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Minter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Minter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.AnnualProvisions.Size()
		i -= size
		if _, err := m.AnnualProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.StartPhaseBlock != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.StartPhaseBlock))
		i--
		dAtA[i] = 0x18
	}
	if m.Phase != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.Phase))
		i--
		dAtA[i] = 0x10
	}
	{
		size := m.Inflation.Size()
		i -= size
		if _, err := m.Inflation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MintEnabled {
		i--
		if m.MintEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.BlocksPerYear != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.BlocksPerYear))
		i--
		dAtA[i] = 0x10
	}
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintMint(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMint(dAtA []byte, offset int, v uint64) int {
	offset -= sovMint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Minter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Inflation.Size()
	n += 1 + l + sovMint(uint64(l))
	if m.Phase != 0 {
		n += 1 + sovMint(uint64(m.Phase))
	}
	if m.StartPhaseBlock != 0 {
		n += 1 + sovMint(uint64(m.StartPhaseBlock))
	}
	l = m.AnnualProvisions.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	if m.BlocksPerYear != 0 {
		n += 1 + sovMint(uint64(m.BlocksPerYear))
	}
	if m.MintEnabled {
		n += 2
	}
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Minter) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Minter{`,
		`Inflation:` + fmt.Sprintf("%v", this.Inflation) + `,`,
		`Phase:` + fmt.Sprintf("%v", this.Phase) + `,`,
		`StartPhaseBlock:` + fmt.Sprintf("%v", this.StartPhaseBlock) + `,`,
		`AnnualProvisions:` + fmt.Sprintf("%v", this.AnnualProvisions) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Params) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Params{`,
		`MintDenom:` + fmt.Sprintf("%v", this.MintDenom) + `,`,
		`BlocksPerYear:` + fmt.Sprintf("%v", this.BlocksPerYear) + `,`,
		`MintEnabled:` + fmt.Sprintf("%v", this.MintEnabled) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMint(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Minter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Minter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Minter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Inflation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Phase", wireType)
			}
			m.Phase = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Phase |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartPhaseBlock", wireType)
			}
			m.StartPhaseBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartPhaseBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnnualProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AnnualProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksPerYear", wireType)
			}
			m.BlocksPerYear = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlocksPerYear |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.MintEnabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMint
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMint
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMint
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMint = fmt.Errorf("proto: unexpected end of group")
)
