// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmwasm/wasm/v1/ibc.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgIBCSend
type MsgIBCSend struct {
	// the channel by which the packet will be sent
	Channel string `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty" yaml:"source_channel"`
	// Timeout height relative to the current block height.
	// The timeout is disabled when set to 0.
	TimeoutHeight uint64 `protobuf:"varint,4,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height,omitempty" yaml:"timeout_height"`
	// Timeout timestamp (in nanoseconds) relative to the current block timestamp.
	// The timeout is disabled when set to 0.
	TimeoutTimestamp uint64 `protobuf:"varint,5,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty" yaml:"timeout_timestamp"`
	// Data is the payload to transfer. We must not make assumption what format or
	// content is in here.
	Data []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *MsgIBCSend) Reset()         { *m = MsgIBCSend{} }
func (m *MsgIBCSend) String() string { return proto.CompactTextString(m) }
func (*MsgIBCSend) ProtoMessage()    {}
func (*MsgIBCSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_af0d1c43ea53c4b9, []int{0}
}
func (m *MsgIBCSend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgIBCSend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgIBCSend.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgIBCSend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgIBCSend.Merge(m, src)
}
func (m *MsgIBCSend) XXX_Size() int {
	return m.Size()
}
func (m *MsgIBCSend) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgIBCSend.DiscardUnknown(m)
}

var xxx_messageInfo_MsgIBCSend proto.InternalMessageInfo

// MsgIBCSendResponse
type MsgIBCSendResponse struct {
	// Sequence number of the IBC packet sent
	Sequence uint64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *MsgIBCSendResponse) Reset()         { *m = MsgIBCSendResponse{} }
func (m *MsgIBCSendResponse) String() string { return proto.CompactTextString(m) }
func (*MsgIBCSendResponse) ProtoMessage()    {}
func (*MsgIBCSendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af0d1c43ea53c4b9, []int{1}
}
func (m *MsgIBCSendResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgIBCSendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgIBCSendResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgIBCSendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgIBCSendResponse.Merge(m, src)
}
func (m *MsgIBCSendResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgIBCSendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgIBCSendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgIBCSendResponse proto.InternalMessageInfo

// MsgIBCCloseChannel port and channel need to be owned by the contract
type MsgIBCCloseChannel struct {
	Channel string `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty" yaml:"source_channel"`
}

func (m *MsgIBCCloseChannel) Reset()         { *m = MsgIBCCloseChannel{} }
func (m *MsgIBCCloseChannel) String() string { return proto.CompactTextString(m) }
func (*MsgIBCCloseChannel) ProtoMessage()    {}
func (*MsgIBCCloseChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_af0d1c43ea53c4b9, []int{2}
}
func (m *MsgIBCCloseChannel) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgIBCCloseChannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgIBCCloseChannel.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgIBCCloseChannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgIBCCloseChannel.Merge(m, src)
}
func (m *MsgIBCCloseChannel) XXX_Size() int {
	return m.Size()
}
func (m *MsgIBCCloseChannel) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgIBCCloseChannel.DiscardUnknown(m)
}

var xxx_messageInfo_MsgIBCCloseChannel proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgIBCSend)(nil), "cosmwasm.wasm.v1.MsgIBCSend")
	proto.RegisterType((*MsgIBCSendResponse)(nil), "cosmwasm.wasm.v1.MsgIBCSendResponse")
	proto.RegisterType((*MsgIBCCloseChannel)(nil), "cosmwasm.wasm.v1.MsgIBCCloseChannel")
}

func init() { proto.RegisterFile("cosmwasm/wasm/v1/ibc.proto", fileDescriptor_af0d1c43ea53c4b9) }

var fileDescriptor_af0d1c43ea53c4b9 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x51, 0xb1, 0x4e, 0xf3, 0x30,
	0x10, 0x8e, 0x7f, 0xf5, 0x2f, 0x60, 0x01, 0x2a, 0x16, 0x48, 0xa1, 0x42, 0x6e, 0xe5, 0xa9, 0x53,
	0x42, 0xd5, 0x8d, 0x09, 0xb5, 0x0b, 0x1d, 0x18, 0x08, 0x4c, 0x2c, 0x95, 0xeb, 0x98, 0x24, 0x52,
	0x62, 0xa7, 0xb1, 0x53, 0xe8, 0x5b, 0xf0, 0x58, 0x1d, 0x3b, 0x32, 0x55, 0x90, 0xbe, 0x41, 0x9f,
	0x00, 0xd5, 0x4d, 0x0a, 0x5d, 0x59, 0xce, 0xdf, 0x7d, 0xdf, 0x77, 0x67, 0xdd, 0x1d, 0x6c, 0x32,
	0xa9, 0x92, 0x57, 0xaa, 0x12, 0xd7, 0x84, 0x69, 0xd7, 0x8d, 0xc6, 0xcc, 0x49, 0x33, 0xa9, 0x25,
	0x6a, 0x54, 0x9a, 0x63, 0xc2, 0xb4, 0xdb, 0x3c, 0x0f, 0x64, 0x20, 0x8d, 0xe8, 0x6e, 0xd0, 0xd6,
	0x47, 0x0a, 0x00, 0xe1, 0xbd, 0x0a, 0x86, 0xfd, 0xc1, 0x23, 0x17, 0x3e, 0xea, 0xc1, 0x03, 0x16,
	0x52, 0x21, 0x78, 0x6c, 0xff, 0x6b, 0x83, 0xce, 0x51, 0xff, 0x72, 0xbd, 0x6c, 0x5d, 0xcc, 0x68,
	0x12, 0xdf, 0x10, 0x25, 0xf3, 0x8c, 0xf1, 0x51, 0xa9, 0x13, 0xaf, 0x72, 0xa2, 0x5b, 0x78, 0xaa,
	0xa3, 0x84, 0xcb, 0x5c, 0x8f, 0x42, 0x1e, 0x05, 0xa1, 0xb6, 0x6b, 0x6d, 0xd0, 0xa9, 0xfd, 0xae,
	0xdd, 0xd7, 0x89, 0x77, 0x52, 0x12, 0x77, 0x26, 0x47, 0x43, 0x78, 0x56, 0x39, 0x36, 0xaf, 0xd2,
	0x34, 0x49, 0xed, 0xff, 0xa6, 0xc9, 0xd5, 0x7a, 0xd9, 0xb2, 0xf7, 0x9b, 0xec, 0x2c, 0xc4, 0x6b,
	0x94, 0xdc, 0x53, 0x45, 0x21, 0x04, 0x6b, 0x3e, 0xd5, 0xd4, 0xae, 0xb7, 0x41, 0xe7, 0xd8, 0x33,
	0x98, 0x5c, 0x43, 0xf4, 0x33, 0xa3, 0xc7, 0x55, 0x2a, 0x85, 0xe2, 0xa8, 0x09, 0x0f, 0x15, 0x9f,
	0xe4, 0x5c, 0x30, 0x6e, 0x83, 0xcd, 0x5f, 0xde, 0x2e, 0x27, 0xc3, 0xaa, 0x62, 0x10, 0x4b, 0xc5,
	0x07, 0xe5, 0xa0, 0x7f, 0xd9, 0x4e, 0xff, 0x61, 0xfe, 0x85, 0xad, 0x79, 0x81, 0xc1, 0xa2, 0xc0,
	0xe0, 0xb3, 0xc0, 0xe0, 0x7d, 0x85, 0xad, 0xc5, 0x0a, 0x5b, 0x1f, 0x2b, 0x6c, 0x3d, 0xf7, 0x82,
	0x48, 0xc7, 0x74, 0xec, 0x4c, 0x32, 0xee, 0x4b, 0x87, 0xc9, 0xc4, 0x9d, 0x64, 0xbe, 0x64, 0x21,
	0x8d, 0x84, 0xfb, 0x92, 0xab, 0x48, 0x8a, 0x2d, 0x7e, 0xdb, 0xde, 0x58, 0xcf, 0x52, 0xae, 0xc6,
	0x75, 0x73, 0xbb, 0xde, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x36, 0xf3, 0x41, 0x6a, 0x01, 0x02,
	0x00, 0x00,
}

func (m *MsgIBCSend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgIBCSend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgIBCSend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintIbc(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x32
	}
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintIbc(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x28
	}
	if m.TimeoutHeight != 0 {
		i = encodeVarintIbc(dAtA, i, uint64(m.TimeoutHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Channel) > 0 {
		i -= len(m.Channel)
		copy(dAtA[i:], m.Channel)
		i = encodeVarintIbc(dAtA, i, uint64(len(m.Channel)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *MsgIBCSendResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgIBCSendResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgIBCSendResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintIbc(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgIBCCloseChannel) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgIBCCloseChannel) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgIBCCloseChannel) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Channel) > 0 {
		i -= len(m.Channel)
		copy(dAtA[i:], m.Channel)
		i = encodeVarintIbc(dAtA, i, uint64(len(m.Channel)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func encodeVarintIbc(dAtA []byte, offset int, v uint64) int {
	offset -= sovIbc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgIBCSend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovIbc(uint64(l))
	}
	if m.TimeoutHeight != 0 {
		n += 1 + sovIbc(uint64(m.TimeoutHeight))
	}
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovIbc(uint64(m.TimeoutTimestamp))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovIbc(uint64(l))
	}
	return n
}

func (m *MsgIBCSendResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sequence != 0 {
		n += 1 + sovIbc(uint64(m.Sequence))
	}
	return n
}

func (m *MsgIBCCloseChannel) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovIbc(uint64(l))
	}
	return n
}

func sovIbc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIbc(x uint64) (n int) {
	return sovIbc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgIBCSend) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIbc
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
			return fmt.Errorf("proto: MsgIBCSend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgIBCSend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbc
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
				return ErrInvalidLengthIbc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIbc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			m.TimeoutHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthIbc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIbc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIbc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIbc
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
func (m *MsgIBCSendResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIbc
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
			return fmt.Errorf("proto: MsgIBCSendResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgIBCSendResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIbc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIbc
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
func (m *MsgIBCCloseChannel) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIbc
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
			return fmt.Errorf("proto: MsgIBCCloseChannel: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgIBCCloseChannel: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbc
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
				return ErrInvalidLengthIbc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIbc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIbc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIbc
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
func skipIbc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIbc
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
					return 0, ErrIntOverflowIbc
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
					return 0, ErrIntOverflowIbc
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
				return 0, ErrInvalidLengthIbc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIbc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIbc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIbc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIbc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIbc = fmt.Errorf("proto: unexpected end of group")
)
