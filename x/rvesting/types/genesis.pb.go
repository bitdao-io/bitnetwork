// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitnetwork/rvesting/v1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// GenesisState defines the module's genesis state
type GenesisState struct {
	// module parameters invariant
	Params     Params                                   `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	From       string                                   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	InitReward github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=init_reward,json=initReward,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"init_reward"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2aa648627cd7256, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *GenesisState) GetInitReward() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.InitReward
	}
	return nil
}

// Params defines the rvesting module params
type Params struct {
	// parameter to enable the vesting of staking reward
	EnableVesting  bool                                     `protobuf:"varint,1,opt,name=enable_vesting,json=enableVesting,proto3" json:"enable_vesting,omitempty"`
	PerBlockReward github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=per_block_reward,json=perBlockReward,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"per_block_reward"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2aa648627cd7256, []int{1}
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

func (m *Params) GetEnableVesting() bool {
	if m != nil {
		return m.EnableVesting
	}
	return false
}

func (m *Params) GetPerBlockReward() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.PerBlockReward
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "bitnetwork.rvesting.v1.GenesisState")
	proto.RegisterType((*Params)(nil), "bitnetwork.rvesting.v1.Params")
}

func init() {
	proto.RegisterFile("bitnetwork/rvesting/v1/genesis.proto", fileDescriptor_b2aa648627cd7256)
}

var fileDescriptor_b2aa648627cd7256 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xb1, 0x4e, 0xf3, 0x30,
	0x10, 0xc7, 0xe3, 0xb6, 0xaa, 0xbe, 0xcf, 0xfd, 0xbe, 0x0a, 0x45, 0x1d, 0x4a, 0x85, 0xd2, 0xaa,
	0x12, 0x52, 0x96, 0xda, 0xa4, 0x6c, 0x8c, 0x61, 0x80, 0x11, 0x05, 0x89, 0x81, 0xa5, 0xb2, 0x53,
	0x93, 0x5a, 0x6d, 0xe2, 0xc8, 0x76, 0x03, 0xbc, 0x05, 0x4f, 0xc1, 0xc0, 0x93, 0x74, 0x60, 0xe8,
	0xc8, 0x04, 0xa8, 0x7d, 0x11, 0x14, 0x3b, 0x05, 0x06, 0x46, 0x26, 0x9f, 0xce, 0xff, 0xfb, 0xdf,
	0xef, 0x74, 0x07, 0x87, 0x94, 0xeb, 0x78, 0x46, 0x78, 0x86, 0x65, 0xc1, 0x94, 0xe6, 0x59, 0x82,
	0x8b, 0x00, 0x27, 0x2c, 0x63, 0x8a, 0x2b, 0x94, 0x4b, 0xa1, 0x85, 0xdb, 0xd9, 0x69, 0xd0, 0x4e,
	0x83, 0x8a, 0xa0, 0xd7, 0x49, 0x44, 0x22, 0x8c, 0x00, 0x97, 0x91, 0xd5, 0xf6, 0xbc, 0x58, 0xa8,
	0x54, 0x28, 0x4c, 0x89, 0x62, 0xb8, 0x08, 0x28, 0xd3, 0x24, 0xc0, 0xb1, 0xe0, 0x99, 0xfd, 0x1f,
	0x3e, 0x03, 0xf8, 0xef, 0xcc, 0xba, 0x5f, 0x6a, 0xa2, 0x99, 0x7b, 0x02, 0x9b, 0x39, 0x91, 0x24,
	0x55, 0x5d, 0x30, 0x00, 0x7e, 0x6b, 0x7c, 0x80, 0x7e, 0xea, 0x86, 0x2e, 0x8c, 0x26, 0x6c, 0xac,
	0x5e, 0xfb, 0x4e, 0x54, 0x55, 0xb8, 0x2e, 0x6c, 0xdc, 0x48, 0x91, 0x76, 0x6b, 0x03, 0xe0, 0xff,
	0x8d, 0x4c, 0xec, 0x2e, 0x60, 0x8b, 0x67, 0x5c, 0x4f, 0x24, 0xbb, 0x25, 0x72, 0xda, 0xad, 0x0f,
	0xea, 0x7e, 0x6b, 0xbc, 0x8f, 0x2c, 0x16, 0x2a, 0xb1, 0x50, 0x85, 0x85, 0x4e, 0x05, 0xcf, 0xc2,
	0xa3, 0xd2, 0xf1, 0xe9, 0xad, 0xef, 0x27, 0x5c, 0xcf, 0x96, 0x14, 0xc5, 0x22, 0xc5, 0xd5, 0x0c,
	0xf6, 0x19, 0xa9, 0xe9, 0x1c, 0xeb, 0xfb, 0x9c, 0x29, 0x53, 0xa0, 0x22, 0x58, 0xfa, 0x47, 0xc6,
	0x7e, 0xf8, 0x08, 0x60, 0xd3, 0xa2, 0xb9, 0x87, 0xb0, 0xcd, 0x32, 0x42, 0x17, 0x6c, 0x52, 0x61,
	0x9b, 0x81, 0xfe, 0x44, 0xff, 0x6d, 0xf6, 0xca, 0x26, 0xdd, 0x25, 0xdc, 0xcb, 0x99, 0x9c, 0xd0,
	0x85, 0x88, 0xe7, 0x3b, 0xc8, 0xda, 0xef, 0x43, 0xb6, 0x73, 0x26, 0xc3, 0xb2, 0x87, 0x05, 0x0d,
	0xcf, 0x57, 0x1b, 0x0f, 0xac, 0x37, 0x1e, 0x78, 0xdf, 0x78, 0xe0, 0x61, 0xeb, 0x39, 0xeb, 0xad,
	0xe7, 0xbc, 0x6c, 0x3d, 0xe7, 0x1a, 0x7d, 0xf3, 0xa4, 0x5c, 0x4f, 0x89, 0x18, 0x71, 0x81, 0x3f,
	0xcf, 0xe2, 0xee, 0xeb, 0x30, 0x8c, 0x3f, 0x6d, 0x9a, 0x45, 0x1e, 0x7f, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x8d, 0xfb, 0x81, 0x98, 0x3a, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InitReward) > 0 {
		for iNdEx := len(m.InitReward) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitReward[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
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
	if len(m.PerBlockReward) > 0 {
		for iNdEx := len(m.PerBlockReward) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PerBlockReward[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.EnableVesting {
		i--
		if m.EnableVesting {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.InitReward) > 0 {
		for _, e := range m.InitReward {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EnableVesting {
		n += 2
	}
	if len(m.PerBlockReward) > 0 {
		for _, e := range m.PerBlockReward {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitReward", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitReward = append(m.InitReward, types.Coin{})
			if err := m.InitReward[len(m.InitReward)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
				return ErrIntOverflowGenesis
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableVesting", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
			m.EnableVesting = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PerBlockReward", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PerBlockReward = append(m.PerBlockReward, types.Coin{})
			if err := m.PerBlockReward[len(m.PerBlockReward)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
