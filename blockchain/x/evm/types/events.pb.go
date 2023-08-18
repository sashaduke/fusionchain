// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ethermint/evm/v1/events.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

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

// EventEthereumTx defines the event for an Ethereum transaction
type EventEthereumTx struct {
	// amount
	Amount string `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// eth_hash is the Ethereum hash of the transaction
	EthHash string `protobuf:"bytes,2,opt,name=eth_hash,json=ethHash,proto3" json:"eth_hash,omitempty"`
	// index of the transaction in the block
	Index string `protobuf:"bytes,3,opt,name=index,proto3" json:"index,omitempty"`
	// gas_used is the amount of gas used by the transaction
	GasUsed string `protobuf:"bytes,4,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	// hash is the Tendermint hash of the transaction
	Hash string `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
	// recipient of the transaction
	Recipient string `protobuf:"bytes,6,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// eth_tx_failed contains a VM error should it occur
	EthTxFailed string `protobuf:"bytes,7,opt,name=eth_tx_failed,json=ethTxFailed,proto3" json:"eth_tx_failed,omitempty"`
}

func (m *EventEthereumTx) Reset()         { *m = EventEthereumTx{} }
func (m *EventEthereumTx) String() string { return proto.CompactTextString(m) }
func (*EventEthereumTx) ProtoMessage()    {}
func (*EventEthereumTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_432e0d592184bde3, []int{0}
}
func (m *EventEthereumTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventEthereumTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventEthereumTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventEthereumTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventEthereumTx.Merge(m, src)
}
func (m *EventEthereumTx) XXX_Size() int {
	return m.Size()
}
func (m *EventEthereumTx) XXX_DiscardUnknown() {
	xxx_messageInfo_EventEthereumTx.DiscardUnknown(m)
}

var xxx_messageInfo_EventEthereumTx proto.InternalMessageInfo

func (m *EventEthereumTx) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *EventEthereumTx) GetEthHash() string {
	if m != nil {
		return m.EthHash
	}
	return ""
}

func (m *EventEthereumTx) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *EventEthereumTx) GetGasUsed() string {
	if m != nil {
		return m.GasUsed
	}
	return ""
}

func (m *EventEthereumTx) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *EventEthereumTx) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *EventEthereumTx) GetEthTxFailed() string {
	if m != nil {
		return m.EthTxFailed
	}
	return ""
}

// EventTxLog defines the event for an Ethereum transaction log
type EventTxLog struct {
	// tx_logs is an array of transaction logs
	TxLogs []string `protobuf:"bytes,1,rep,name=tx_logs,json=txLogs,proto3" json:"tx_logs,omitempty"`
}

func (m *EventTxLog) Reset()         { *m = EventTxLog{} }
func (m *EventTxLog) String() string { return proto.CompactTextString(m) }
func (*EventTxLog) ProtoMessage()    {}
func (*EventTxLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_432e0d592184bde3, []int{1}
}
func (m *EventTxLog) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventTxLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventTxLog.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventTxLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventTxLog.Merge(m, src)
}
func (m *EventTxLog) XXX_Size() int {
	return m.Size()
}
func (m *EventTxLog) XXX_DiscardUnknown() {
	xxx_messageInfo_EventTxLog.DiscardUnknown(m)
}

var xxx_messageInfo_EventTxLog proto.InternalMessageInfo

func (m *EventTxLog) GetTxLogs() []string {
	if m != nil {
		return m.TxLogs
	}
	return nil
}

// EventMessage
type EventMessage struct {
	// module which emits the event
	Module string `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	// sender of the message
	Sender string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	// tx_type is the type of the message
	TxType string `protobuf:"bytes,3,opt,name=tx_type,json=txType,proto3" json:"tx_type,omitempty"`
}

func (m *EventMessage) Reset()         { *m = EventMessage{} }
func (m *EventMessage) String() string { return proto.CompactTextString(m) }
func (*EventMessage) ProtoMessage()    {}
func (*EventMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_432e0d592184bde3, []int{2}
}
func (m *EventMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventMessage.Merge(m, src)
}
func (m *EventMessage) XXX_Size() int {
	return m.Size()
}
func (m *EventMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EventMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EventMessage proto.InternalMessageInfo

func (m *EventMessage) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *EventMessage) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *EventMessage) GetTxType() string {
	if m != nil {
		return m.TxType
	}
	return ""
}

// EventBlockBloom defines an Ethereum block bloom filter event
type EventBlockBloom struct {
	// bloom is the bloom filter of the block
	Bloom string `protobuf:"bytes,1,opt,name=bloom,proto3" json:"bloom,omitempty"`
}

func (m *EventBlockBloom) Reset()         { *m = EventBlockBloom{} }
func (m *EventBlockBloom) String() string { return proto.CompactTextString(m) }
func (*EventBlockBloom) ProtoMessage()    {}
func (*EventBlockBloom) Descriptor() ([]byte, []int) {
	return fileDescriptor_432e0d592184bde3, []int{3}
}
func (m *EventBlockBloom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventBlockBloom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventBlockBloom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventBlockBloom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBlockBloom.Merge(m, src)
}
func (m *EventBlockBloom) XXX_Size() int {
	return m.Size()
}
func (m *EventBlockBloom) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBlockBloom.DiscardUnknown(m)
}

var xxx_messageInfo_EventBlockBloom proto.InternalMessageInfo

func (m *EventBlockBloom) GetBloom() string {
	if m != nil {
		return m.Bloom
	}
	return ""
}

func init() {
	proto.RegisterType((*EventEthereumTx)(nil), "ethermint.evm.v1.EventEthereumTx")
	proto.RegisterType((*EventTxLog)(nil), "ethermint.evm.v1.EventTxLog")
	proto.RegisterType((*EventMessage)(nil), "ethermint.evm.v1.EventMessage")
	proto.RegisterType((*EventBlockBloom)(nil), "ethermint.evm.v1.EventBlockBloom")
}

func init() { proto.RegisterFile("ethermint/evm/v1/events.proto", fileDescriptor_432e0d592184bde3) }

var fileDescriptor_432e0d592184bde3 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x51, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x6d, 0x68, 0x9b, 0x52, 0x03, 0x02, 0x59, 0x08, 0x82, 0x04, 0x51, 0x15, 0x09, 0xc1, 0x29,
	0x51, 0xe1, 0x0f, 0x2a, 0x15, 0x71, 0x28, 0x17, 0x14, 0x84, 0xc4, 0x25, 0x72, 0xe3, 0x69, 0x6c,
	0x11, 0xdb, 0x69, 0xec, 0x44, 0xee, 0x5f, 0xf0, 0x59, 0x48, 0x5c, 0x7a, 0xe4, 0xb8, 0x6a, 0x7f,
	0x64, 0x65, 0x27, 0xbb, 0x7b, 0x9b, 0xf7, 0xde, 0xcc, 0xb3, 0xdf, 0x0c, 0x7a, 0x07, 0x86, 0x41,
	0x2b, 0xb8, 0x34, 0x19, 0xf4, 0x22, 0xeb, 0xd7, 0x19, 0xf4, 0x20, 0x8d, 0x4e, 0x9b, 0x56, 0x19,
	0x85, 0x5f, 0xdc, 0xcb, 0x29, 0xf4, 0x22, 0xed, 0xd7, 0xc9, 0xbf, 0x00, 0x3d, 0xdf, 0xba, 0x96,
	0xad, 0x53, 0xa0, 0x13, 0xb9, 0xc5, 0xaf, 0x50, 0x48, 0x84, 0xea, 0xa4, 0x89, 0x82, 0x55, 0xf0,
	0x71, 0xf9, 0x7d, 0x44, 0xf8, 0x0d, 0x7a, 0x0c, 0x86, 0x15, 0x8c, 0x68, 0x16, 0x3d, 0xf2, 0xca,
	0x02, 0x0c, 0xfb, 0x4a, 0x34, 0xc3, 0x2f, 0xd1, 0x9c, 0x4b, 0x0a, 0x36, 0x9a, 0x7a, 0x7e, 0x00,
	0x6e, 0xa0, 0x22, 0xba, 0xe8, 0x34, 0xd0, 0x68, 0x36, 0x0c, 0x54, 0x44, 0xff, 0xd0, 0x40, 0x31,
	0x46, 0x33, 0xef, 0x33, 0xf7, 0xb4, 0xaf, 0xf1, 0x5b, 0xb4, 0x6c, 0xa1, 0xe4, 0x0d, 0x07, 0x69,
	0xa2, 0xd0, 0x0b, 0x0f, 0x04, 0x4e, 0xd0, 0x33, 0xf7, 0xba, 0xb1, 0xc5, 0x81, 0xf0, 0x1a, 0x68,
	0xb4, 0xf0, 0x1d, 0x4f, 0xc0, 0xb0, 0xdc, 0x7e, 0xf1, 0x54, 0xf2, 0x1e, 0x21, 0x1f, 0x26, 0xb7,
	0x3b, 0x55, 0xe1, 0xd7, 0x68, 0x61, 0x6c, 0x51, 0xab, 0x4a, 0x47, 0xc1, 0x6a, 0xea, 0x82, 0x18,
	0xc7, 0xeb, 0xe4, 0x27, 0x7a, 0xea, 0xdb, 0xbe, 0x81, 0xd6, 0xa4, 0x02, 0x17, 0x58, 0x28, 0xda,
	0xd5, 0x70, 0x17, 0x78, 0x40, 0x8e, 0xd7, 0x20, 0x29, 0xb4, 0x63, 0xdc, 0x11, 0x8d, 0xc6, 0xe6,
	0xd4, 0xc0, 0x98, 0x37, 0x34, 0x36, 0x3f, 0x35, 0x90, 0x7c, 0x18, 0x97, 0xb9, 0xa9, 0x55, 0xf9,
	0x7b, 0x53, 0x2b, 0x25, 0xdc, 0x66, 0xf6, 0xae, 0x18, 0xad, 0x07, 0xb0, 0xd9, 0xfd, 0xbd, 0xc4,
	0xc1, 0xf9, 0x12, 0x07, 0x37, 0x97, 0x38, 0xf8, 0x73, 0x8d, 0x27, 0xe7, 0x6b, 0x3c, 0xf9, 0x7f,
	0x8d, 0x27, 0xbf, 0x3e, 0x55, 0xdc, 0xd4, 0x64, 0x9f, 0x1e, 0x5b, 0xa0, 0x2a, 0x2d, 0x95, 0xc8,
	0x8e, 0x2d, 0x55, 0x25, 0x23, 0x5c, 0x66, 0x87, 0x4e, 0x73, 0x25, 0x87, 0xda, 0xfa, 0x03, 0xbb,
	0x3f, 0xe8, 0x7d, 0xe8, 0xaf, 0xfb, 0xf9, 0x36, 0x00, 0x00, 0xff, 0xff, 0x55, 0x9a, 0xad, 0x7a,
	0xfe, 0x01, 0x00, 0x00,
}

func (m *EventEthereumTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventEthereumTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventEthereumTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EthTxFailed) > 0 {
		i -= len(m.EthTxFailed)
		copy(dAtA[i:], m.EthTxFailed)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.EthTxFailed)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.GasUsed) > 0 {
		i -= len(m.GasUsed)
		copy(dAtA[i:], m.GasUsed)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.GasUsed)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.EthHash) > 0 {
		i -= len(m.EthHash)
		copy(dAtA[i:], m.EthHash)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.EthHash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventTxLog) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventTxLog) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventTxLog) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxLogs) > 0 {
		for iNdEx := len(m.TxLogs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.TxLogs[iNdEx])
			copy(dAtA[i:], m.TxLogs[iNdEx])
			i = encodeVarintEvents(dAtA, i, uint64(len(m.TxLogs[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *EventMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxType) > 0 {
		i -= len(m.TxType)
		copy(dAtA[i:], m.TxType)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.TxType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Module) > 0 {
		i -= len(m.Module)
		copy(dAtA[i:], m.Module)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Module)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventBlockBloom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventBlockBloom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventBlockBloom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bloom) > 0 {
		i -= len(m.Bloom)
		copy(dAtA[i:], m.Bloom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Bloom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventEthereumTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.EthHash)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.GasUsed)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.EthTxFailed)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventTxLog) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TxLogs) > 0 {
		for _, s := range m.TxLogs {
			l = len(s)
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func (m *EventMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Module)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.TxType)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventBlockBloom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Bloom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventEthereumTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventEthereumTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventEthereumTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasUsed = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthTxFailed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthTxFailed = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventTxLog) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventTxLog: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventTxLog: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxLogs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxLogs = append(m.TxLogs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Module", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Module = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventBlockBloom) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventBlockBloom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventBlockBloom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bloom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bloom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
