// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: checkers/stored_game.proto

package types

import (
	fmt "fmt"
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

type StoredGame struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index     string `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	Game      string `protobuf:"bytes,3,opt,name=game,proto3" json:"game,omitempty"`
	Turn      string `protobuf:"bytes,4,opt,name=turn,proto3" json:"turn,omitempty"`
	Red       string `protobuf:"bytes,5,opt,name=red,proto3" json:"red,omitempty"`
	Black     string `protobuf:"bytes,6,opt,name=black,proto3" json:"black,omitempty"`
	MoveCount uint64 `protobuf:"varint,7,opt,name=moveCount,proto3" json:"moveCount,omitempty"`
	BeforeId  string `protobuf:"bytes,8,opt,name=beforeId,proto3" json:"beforeId,omitempty"`
	AfterId   string `protobuf:"bytes,9,opt,name=afterId,proto3" json:"afterId,omitempty"`
	Deadline  string `protobuf:"bytes,10,opt,name=deadline,proto3" json:"deadline,omitempty"`
}

func (m *StoredGame) Reset()         { *m = StoredGame{} }
func (m *StoredGame) String() string { return proto.CompactTextString(m) }
func (*StoredGame) ProtoMessage()    {}
func (*StoredGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_8439c9c90688ff75, []int{0}
}
func (m *StoredGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoredGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoredGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoredGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoredGame.Merge(m, src)
}
func (m *StoredGame) XXX_Size() int {
	return m.Size()
}
func (m *StoredGame) XXX_DiscardUnknown() {
	xxx_messageInfo_StoredGame.DiscardUnknown(m)
}

var xxx_messageInfo_StoredGame proto.InternalMessageInfo

func (m *StoredGame) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *StoredGame) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *StoredGame) GetGame() string {
	if m != nil {
		return m.Game
	}
	return ""
}

func (m *StoredGame) GetTurn() string {
	if m != nil {
		return m.Turn
	}
	return ""
}

func (m *StoredGame) GetRed() string {
	if m != nil {
		return m.Red
	}
	return ""
}

func (m *StoredGame) GetBlack() string {
	if m != nil {
		return m.Black
	}
	return ""
}

func (m *StoredGame) GetMoveCount() uint64 {
	if m != nil {
		return m.MoveCount
	}
	return 0
}

func (m *StoredGame) GetBeforeId() string {
	if m != nil {
		return m.BeforeId
	}
	return ""
}

func (m *StoredGame) GetAfterId() string {
	if m != nil {
		return m.AfterId
	}
	return ""
}

func (m *StoredGame) GetDeadline() string {
	if m != nil {
		return m.Deadline
	}
	return ""
}

func init() {
	proto.RegisterType((*StoredGame)(nil), "alice.checkers.checkers.StoredGame")
}

func init() { proto.RegisterFile("checkers/stored_game.proto", fileDescriptor_8439c9c90688ff75) }

var fileDescriptor_8439c9c90688ff75 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x63, 0x9a, 0xfe, 0x89, 0x27, 0x64, 0x21, 0x71, 0xaa, 0x90, 0x55, 0x31, 0x55, 0x0c,
	0xc9, 0xc0, 0x1b, 0x00, 0x12, 0xea, 0x5a, 0x36, 0x16, 0xe4, 0xd8, 0xd7, 0x36, 0x6a, 0x12, 0x57,
	0x8e, 0x83, 0xca, 0x5b, 0xf0, 0x58, 0x8c, 0x1d, 0x19, 0x51, 0xf2, 0x1a, 0x0c, 0xc8, 0x0e, 0x49,
	0xb7, 0xef, 0xf7, 0x93, 0xef, 0x7c, 0xfa, 0xe8, 0x5c, 0xee, 0x50, 0xee, 0xd1, 0x54, 0x49, 0x65,
	0xb5, 0x41, 0xf5, 0xb6, 0x15, 0x05, 0xc6, 0x07, 0xa3, 0xad, 0x66, 0xd7, 0x22, 0xcf, 0x24, 0xc6,
	0xfd, 0x8b, 0x21, 0xdc, 0xfe, 0x12, 0x4a, 0x5f, 0xfc, 0xf3, 0x67, 0x51, 0x20, 0x03, 0x3a, 0x95,
	0x06, 0x85, 0xd5, 0x06, 0xc8, 0x82, 0x2c, 0xa3, 0x75, 0x8f, 0xec, 0x8a, 0x8e, 0xb3, 0x52, 0xe1,
	0x11, 0x2e, 0xbc, 0xef, 0x80, 0x31, 0x1a, 0xba, 0x5f, 0x60, 0xe4, 0xa5, 0xcf, 0xce, 0xd9, 0xda,
	0x94, 0x10, 0x76, 0xce, 0x65, 0x76, 0x49, 0x47, 0x06, 0x15, 0x8c, 0xbd, 0x72, 0xd1, 0xed, 0x4b,
	0x73, 0x21, 0xf7, 0x30, 0xe9, 0xf6, 0x79, 0x60, 0x37, 0x34, 0x2a, 0xf4, 0x3b, 0x3e, 0xea, 0xba,
	0xb4, 0x30, 0x5d, 0x90, 0x65, 0xb8, 0x3e, 0x0b, 0x36, 0xa7, 0xb3, 0x14, 0x37, 0xda, 0xe0, 0x4a,
	0xc1, 0xcc, 0x8f, 0x0d, 0xec, 0x2e, 0x17, 0x1b, 0x8b, 0x66, 0xa5, 0x20, 0xea, 0x2e, 0xff, 0x47,
	0x37, 0xa5, 0x50, 0xa8, 0x3c, 0x2b, 0x11, 0x68, 0x37, 0xd5, 0xf3, 0xc3, 0xd3, 0x57, 0xc3, 0xc9,
	0xa9, 0xe1, 0xe4, 0xa7, 0xe1, 0xe4, 0xb3, 0xe5, 0xc1, 0xa9, 0xe5, 0xc1, 0x77, 0xcb, 0x83, 0xd7,
	0xbb, 0x6d, 0x66, 0x77, 0x75, 0x1a, 0x4b, 0x5d, 0x24, 0xbe, 0xbc, 0x64, 0xa8, 0xf7, 0x78, 0x8e,
	0xf6, 0xe3, 0x80, 0x55, 0x3a, 0xf1, 0x25, 0xdf, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x31,
	0x82, 0xbd, 0x82, 0x01, 0x00, 0x00,
}

func (m *StoredGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoredGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoredGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Deadline) > 0 {
		i -= len(m.Deadline)
		copy(dAtA[i:], m.Deadline)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Deadline)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.AfterId) > 0 {
		i -= len(m.AfterId)
		copy(dAtA[i:], m.AfterId)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.AfterId)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.BeforeId) > 0 {
		i -= len(m.BeforeId)
		copy(dAtA[i:], m.BeforeId)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.BeforeId)))
		i--
		dAtA[i] = 0x42
	}
	if m.MoveCount != 0 {
		i = encodeVarintStoredGame(dAtA, i, uint64(m.MoveCount))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Black) > 0 {
		i -= len(m.Black)
		copy(dAtA[i:], m.Black)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Black)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Red) > 0 {
		i -= len(m.Red)
		copy(dAtA[i:], m.Red)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Red)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Turn) > 0 {
		i -= len(m.Turn)
		copy(dAtA[i:], m.Turn)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Turn)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Game) > 0 {
		i -= len(m.Game)
		copy(dAtA[i:], m.Game)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Game)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStoredGame(dAtA []byte, offset int, v uint64) int {
	offset -= sovStoredGame(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StoredGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.Game)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.Turn)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.Red)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.Black)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	if m.MoveCount != 0 {
		n += 1 + sovStoredGame(uint64(m.MoveCount))
	}
	l = len(m.BeforeId)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.AfterId)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.Deadline)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	return n
}

func sovStoredGame(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStoredGame(x uint64) (n int) {
	return sovStoredGame(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoredGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStoredGame
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
			return fmt.Errorf("proto: StoredGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoredGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Game", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Game = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Turn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Turn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Red", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Red = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Black", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Black = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MoveCount", wireType)
			}
			m.MoveCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MoveCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeforeId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BeforeId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AfterId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AfterId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deadline", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deadline = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStoredGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStoredGame
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
func skipStoredGame(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStoredGame
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
					return 0, ErrIntOverflowStoredGame
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
					return 0, ErrIntOverflowStoredGame
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
				return 0, ErrInvalidLengthStoredGame
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStoredGame
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStoredGame
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStoredGame        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStoredGame          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStoredGame = fmt.Errorf("proto: unexpected end of group")
)
