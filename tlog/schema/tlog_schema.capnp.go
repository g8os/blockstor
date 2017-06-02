// Code generated by capnpc-go. DO NOT EDIT.

package schema

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type HandshakeRequest struct{ capnp.Struct }

// HandshakeRequest_TypeID is the unique identifier for the type HandshakeRequest.
const HandshakeRequest_TypeID = 0xe0d4e6d68fa24ac0

func NewHandshakeRequest(s *capnp.Segment) (HandshakeRequest, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return HandshakeRequest{st}, err
}

func NewRootHandshakeRequest(s *capnp.Segment) (HandshakeRequest, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return HandshakeRequest{st}, err
}

func ReadRootHandshakeRequest(msg *capnp.Message) (HandshakeRequest, error) {
	root, err := msg.RootPtr()
	return HandshakeRequest{root.Struct()}, err
}

func (s HandshakeRequest) String() string {
	str, _ := text.Marshal(0xe0d4e6d68fa24ac0, s.Struct)
	return str
}

func (s HandshakeRequest) Version() uint32 {
	return s.Struct.Uint32(0)
}

func (s HandshakeRequest) SetVersion(v uint32) {
	s.Struct.SetUint32(0, v)
}

func (s HandshakeRequest) VdiskID() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s HandshakeRequest) HasVdiskID() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HandshakeRequest) VdiskIDBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s HandshakeRequest) SetVdiskID(v string) error {
	return s.Struct.SetText(0, v)
}

func (s HandshakeRequest) FirstSequence() uint64 {
	return s.Struct.Uint64(8)
}

func (s HandshakeRequest) SetFirstSequence(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s HandshakeRequest) ResetFirstSequence() bool {
	return s.Struct.Bit(32)
}

func (s HandshakeRequest) SetResetFirstSequence(v bool) {
	s.Struct.SetBit(32, v)
}

// HandshakeRequest_List is a list of HandshakeRequest.
type HandshakeRequest_List struct{ capnp.List }

// NewHandshakeRequest creates a new list of HandshakeRequest.
func NewHandshakeRequest_List(s *capnp.Segment, sz int32) (HandshakeRequest_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	return HandshakeRequest_List{l}, err
}

func (s HandshakeRequest_List) At(i int) HandshakeRequest { return HandshakeRequest{s.List.Struct(i)} }

func (s HandshakeRequest_List) Set(i int, v HandshakeRequest) error {
	return s.List.SetStruct(i, v.Struct)
}

// HandshakeRequest_Promise is a wrapper for a HandshakeRequest promised by a client call.
type HandshakeRequest_Promise struct{ *capnp.Pipeline }

func (p HandshakeRequest_Promise) Struct() (HandshakeRequest, error) {
	s, err := p.Pipeline.Struct()
	return HandshakeRequest{s}, err
}

type HandshakeResponse struct{ capnp.Struct }

// HandshakeResponse_TypeID is the unique identifier for the type HandshakeResponse.
const HandshakeResponse_TypeID = 0xee959a7d96c96641

func NewHandshakeResponse(s *capnp.Segment) (HandshakeResponse, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return HandshakeResponse{st}, err
}

func NewRootHandshakeResponse(s *capnp.Segment) (HandshakeResponse, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return HandshakeResponse{st}, err
}

func ReadRootHandshakeResponse(msg *capnp.Message) (HandshakeResponse, error) {
	root, err := msg.RootPtr()
	return HandshakeResponse{root.Struct()}, err
}

func (s HandshakeResponse) String() string {
	str, _ := text.Marshal(0xee959a7d96c96641, s.Struct)
	return str
}

func (s HandshakeResponse) Version() uint32 {
	return s.Struct.Uint32(0)
}

func (s HandshakeResponse) SetVersion(v uint32) {
	s.Struct.SetUint32(0, v)
}

func (s HandshakeResponse) Status() int8 {
	return int8(s.Struct.Uint8(4))
}

func (s HandshakeResponse) SetStatus(v int8) {
	s.Struct.SetUint8(4, uint8(v))
}

// HandshakeResponse_List is a list of HandshakeResponse.
type HandshakeResponse_List struct{ capnp.List }

// NewHandshakeResponse creates a new list of HandshakeResponse.
func NewHandshakeResponse_List(s *capnp.Segment, sz int32) (HandshakeResponse_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return HandshakeResponse_List{l}, err
}

func (s HandshakeResponse_List) At(i int) HandshakeResponse {
	return HandshakeResponse{s.List.Struct(i)}
}

func (s HandshakeResponse_List) Set(i int, v HandshakeResponse) error {
	return s.List.SetStruct(i, v.Struct)
}

// HandshakeResponse_Promise is a wrapper for a HandshakeResponse promised by a client call.
type HandshakeResponse_Promise struct{ *capnp.Pipeline }

func (p HandshakeResponse_Promise) Struct() (HandshakeResponse, error) {
	s, err := p.Pipeline.Struct()
	return HandshakeResponse{s}, err
}

type TlogResponse struct{ capnp.Struct }

// TlogResponse_TypeID is the unique identifier for the type TlogResponse.
const TlogResponse_TypeID = 0x98d11ae1c78a24d9

func NewTlogResponse(s *capnp.Segment) (TlogResponse, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return TlogResponse{st}, err
}

func NewRootTlogResponse(s *capnp.Segment) (TlogResponse, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return TlogResponse{st}, err
}

func ReadRootTlogResponse(msg *capnp.Message) (TlogResponse, error) {
	root, err := msg.RootPtr()
	return TlogResponse{root.Struct()}, err
}

func (s TlogResponse) String() string {
	str, _ := text.Marshal(0x98d11ae1c78a24d9, s.Struct)
	return str
}

func (s TlogResponse) Status() int8 {
	return int8(s.Struct.Uint8(0))
}

func (s TlogResponse) SetStatus(v int8) {
	s.Struct.SetUint8(0, uint8(v))
}

func (s TlogResponse) Sequences() (capnp.UInt64List, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.UInt64List{List: p.List()}, err
}

func (s TlogResponse) HasSequences() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s TlogResponse) SetSequences(v capnp.UInt64List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewSequences sets the sequences field to a newly
// allocated capnp.UInt64List, preferring placement in s's segment.
func (s TlogResponse) NewSequences(n int32) (capnp.UInt64List, error) {
	l, err := capnp.NewUInt64List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.UInt64List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// TlogResponse_List is a list of TlogResponse.
type TlogResponse_List struct{ capnp.List }

// NewTlogResponse creates a new list of TlogResponse.
func NewTlogResponse_List(s *capnp.Segment, sz int32) (TlogResponse_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return TlogResponse_List{l}, err
}

func (s TlogResponse_List) At(i int) TlogResponse { return TlogResponse{s.List.Struct(i)} }

func (s TlogResponse_List) Set(i int, v TlogResponse) error { return s.List.SetStruct(i, v.Struct) }

// TlogResponse_Promise is a wrapper for a TlogResponse promised by a client call.
type TlogResponse_Promise struct{ *capnp.Pipeline }

func (p TlogResponse_Promise) Struct() (TlogResponse, error) {
	s, err := p.Pipeline.Struct()
	return TlogResponse{s}, err
}

type TlogBlock struct{ capnp.Struct }

// TlogBlock_TypeID is the unique identifier for the type TlogBlock.
const TlogBlock_TypeID = 0x8cf178de3c82d431

func NewTlogBlock(s *capnp.Segment) (TlogBlock, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 40, PointerCount: 2})
	return TlogBlock{st}, err
}

func NewRootTlogBlock(s *capnp.Segment) (TlogBlock, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 40, PointerCount: 2})
	return TlogBlock{st}, err
}

func ReadRootTlogBlock(msg *capnp.Message) (TlogBlock, error) {
	root, err := msg.RootPtr()
	return TlogBlock{root.Struct()}, err
}

func (s TlogBlock) String() string {
	str, _ := text.Marshal(0x8cf178de3c82d431, s.Struct)
	return str
}

func (s TlogBlock) Sequence() uint64 {
	return s.Struct.Uint64(0)
}

func (s TlogBlock) SetSequence(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s TlogBlock) Offset() uint64 {
	return s.Struct.Uint64(8)
}

func (s TlogBlock) SetOffset(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s TlogBlock) Size() uint64 {
	return s.Struct.Uint64(16)
}

func (s TlogBlock) SetSize(v uint64) {
	s.Struct.SetUint64(16, v)
}

func (s TlogBlock) Hash() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s TlogBlock) HasHash() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s TlogBlock) SetHash(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s TlogBlock) Data() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s TlogBlock) HasData() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s TlogBlock) SetData(v []byte) error {
	return s.Struct.SetData(1, v)
}

func (s TlogBlock) Timestamp() uint64 {
	return s.Struct.Uint64(24)
}

func (s TlogBlock) SetTimestamp(v uint64) {
	s.Struct.SetUint64(24, v)
}

func (s TlogBlock) Operation() uint8 {
	return s.Struct.Uint8(32)
}

func (s TlogBlock) SetOperation(v uint8) {
	s.Struct.SetUint8(32, v)
}

// TlogBlock_List is a list of TlogBlock.
type TlogBlock_List struct{ capnp.List }

// NewTlogBlock creates a new list of TlogBlock.
func NewTlogBlock_List(s *capnp.Segment, sz int32) (TlogBlock_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 40, PointerCount: 2}, sz)
	return TlogBlock_List{l}, err
}

func (s TlogBlock_List) At(i int) TlogBlock { return TlogBlock{s.List.Struct(i)} }

func (s TlogBlock_List) Set(i int, v TlogBlock) error { return s.List.SetStruct(i, v.Struct) }

// TlogBlock_Promise is a wrapper for a TlogBlock promised by a client call.
type TlogBlock_Promise struct{ *capnp.Pipeline }

func (p TlogBlock_Promise) Struct() (TlogBlock, error) {
	s, err := p.Pipeline.Struct()
	return TlogBlock{s}, err
}

type TlogAggregation struct{ capnp.Struct }

// TlogAggregation_TypeID is the unique identifier for the type TlogAggregation.
const TlogAggregation_TypeID = 0xe46ab5b4b619e094

func NewTlogAggregation(s *capnp.Segment) (TlogAggregation, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 4})
	return TlogAggregation{st}, err
}

func NewRootTlogAggregation(s *capnp.Segment) (TlogAggregation, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 4})
	return TlogAggregation{st}, err
}

func ReadRootTlogAggregation(msg *capnp.Message) (TlogAggregation, error) {
	root, err := msg.RootPtr()
	return TlogAggregation{root.Struct()}, err
}

func (s TlogAggregation) String() string {
	str, _ := text.Marshal(0xe46ab5b4b619e094, s.Struct)
	return str
}

func (s TlogAggregation) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s TlogAggregation) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s TlogAggregation) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s TlogAggregation) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s TlogAggregation) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s TlogAggregation) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s TlogAggregation) Timestamp() uint64 {
	return s.Struct.Uint64(8)
}

func (s TlogAggregation) SetTimestamp(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s TlogAggregation) VdiskID() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s TlogAggregation) HasVdiskID() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s TlogAggregation) VdiskIDBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s TlogAggregation) SetVdiskID(v string) error {
	return s.Struct.SetText(1, v)
}

func (s TlogAggregation) Blocks() (TlogBlock_List, error) {
	p, err := s.Struct.Ptr(2)
	return TlogBlock_List{List: p.List()}, err
}

func (s TlogAggregation) HasBlocks() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s TlogAggregation) SetBlocks(v TlogBlock_List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewBlocks sets the blocks field to a newly
// allocated TlogBlock_List, preferring placement in s's segment.
func (s TlogAggregation) NewBlocks(n int32) (TlogBlock_List, error) {
	l, err := NewTlogBlock_List(s.Struct.Segment(), n)
	if err != nil {
		return TlogBlock_List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s TlogAggregation) Prev() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return []byte(p.Data()), err
}

func (s TlogAggregation) HasPrev() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s TlogAggregation) SetPrev(v []byte) error {
	return s.Struct.SetData(3, v)
}

// TlogAggregation_List is a list of TlogAggregation.
type TlogAggregation_List struct{ capnp.List }

// NewTlogAggregation creates a new list of TlogAggregation.
func NewTlogAggregation_List(s *capnp.Segment, sz int32) (TlogAggregation_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 4}, sz)
	return TlogAggregation_List{l}, err
}

func (s TlogAggregation_List) At(i int) TlogAggregation { return TlogAggregation{s.List.Struct(i)} }

func (s TlogAggregation_List) Set(i int, v TlogAggregation) error {
	return s.List.SetStruct(i, v.Struct)
}

// TlogAggregation_Promise is a wrapper for a TlogAggregation promised by a client call.
type TlogAggregation_Promise struct{ *capnp.Pipeline }

func (p TlogAggregation_Promise) Struct() (TlogAggregation, error) {
	s, err := p.Pipeline.Struct()
	return TlogAggregation{s}, err
}

type Command struct{ capnp.Struct }

// Command_TypeID is the unique identifier for the type Command.
const Command_TypeID = 0xdbe14b5e7e7c6009

func NewCommand(s *capnp.Segment) (Command, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Command{st}, err
}

func NewRootCommand(s *capnp.Segment) (Command, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Command{st}, err
}

func ReadRootCommand(msg *capnp.Message) (Command, error) {
	root, err := msg.RootPtr()
	return Command{root.Struct()}, err
}

func (s Command) String() string {
	str, _ := text.Marshal(0xdbe14b5e7e7c6009, s.Struct)
	return str
}

func (s Command) Type() uint8 {
	return s.Struct.Uint8(0)
}

func (s Command) SetType(v uint8) {
	s.Struct.SetUint8(0, v)
}

func (s Command) Sequences() (capnp.UInt64List, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.UInt64List{List: p.List()}, err
}

func (s Command) HasSequences() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Command) SetSequences(v capnp.UInt64List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewSequences sets the sequences field to a newly
// allocated capnp.UInt64List, preferring placement in s's segment.
func (s Command) NewSequences(n int32) (capnp.UInt64List, error) {
	l, err := capnp.NewUInt64List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.UInt64List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Command_List is a list of Command.
type Command_List struct{ capnp.List }

// NewCommand creates a new list of Command.
func NewCommand_List(s *capnp.Segment, sz int32) (Command_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Command_List{l}, err
}

func (s Command_List) At(i int) Command { return Command{s.List.Struct(i)} }

func (s Command_List) Set(i int, v Command) error { return s.List.SetStruct(i, v.Struct) }

// Command_Promise is a wrapper for a Command promised by a client call.
type Command_Promise struct{ *capnp.Pipeline }

func (p Command_Promise) Struct() (Command, error) {
	s, err := p.Pipeline.Struct()
	return Command{s}, err
}

const schema_f4533cbae6e08506 = "x\xda\x8cT]h\x1cU\x14>\xdf=3\xd3D\x9a" +
	"n\x87\xdd\x07[\x84\xf6\xa1\x0fm\xb1\xdaZ_,\x95" +
	"\x98(\xd2F\x0a\xbd[A\x10\xd1\x8e\xbbww\xc7\xdd" +
	"\x9d\x99\xcc\x9d$*jE\xf0E\x84\xbeTi\x0bB" +
	"+}0`@!\x11\x11\x03\x11\"\x1aH@!\x8a" +
	"\"\xc8\x06$\xbe\x09\x82\xef#w6\xfb\xe3&$\xbe" +
	"\xed9\xfb\xdd=\xe7\xfb\xce\xf7\xed\xe9\x11\xf1\x848c" +
	"\xcf0\x91<n;\xe9\x99\xf5w\xce\xff\xfe\xea\xdf\xef" +
	"\x93<\x04;u\xdemm~u\xfe\xca?d\x8b}" +
	"Dg\xef\xe20\xf2\x9f\xc3|\x9c\xc3u\x10\xd2_\x8f" +
	"\xbd\xf7\xdd\xc6\xe1\x1fo\x1a8\xfa\xe0\x19f\x8c\x1fA" +
	"^\xf2>\xa2\xfc%\x9e!\xa4\xc3W\xdfx\xeb\xc5g" +
	"6~\xdb\x11\xbd\xc8\xf7!\xbf\x96\xa1W2\xf4\xd2\xc4" +
	"\xc7\xd7\x7f\xde\\o\x19\xb4\x18D?f\x15\x91\xbfd" +
	"\x19\xf4E\xebOBz\xa3u\xe8\x8b\xf9\x85W\xfe\x18" +
	"D\x1b\xc8\xd9\x13\xf6\x04\xf2\x8f\xdb\xd9C\xfb9\xb3\xf8" +
	"Xe\xe5\xc37o\x7f\xf0\xd7\xc0*\x19\xfa\x96\xf3<" +
	"\xf2s\x8e\xf9\xedYg\x86N\xa5\xbaTSM\xef\xe1" +
	"\x84\x1ba\xf5\xa5v\xf1P\xc9\x8b\x82\xe8\xdc\xb3\x8d\xb0" +
	":\xde\x08\xb9T\xbf\x0c\xc8\x07\xd8\"\xb2@\xe4.L" +
	"\x10\xc9y\x86\\\x12p\x81\x02Ls\xf1\x1c\x91\xfc\x92" +
	"!\x97\x05\\!\x0a\x10D\xee7'\x89\xe4\xd7\x0c\xf9" +
	"\xbd\x00\xb8\x00&r\xbf5\xbd%\x86\\\x15p-\x14" +
	"`\x11\xb9+\xa6\xb9\xcc\x90?\x08\xb86\x17`\x13\xb9" +
	"kE\"\xb9\xca\x90\xbf\x08\xb8\xce\xd1\x02\x1c\"\xf7'" +
	"\xd3\\g\xc8\x96@\xaa\xd5\xe4\x94\x0aJ\x8a\x880L" +
	"\x02\xc3\x84\xd1\xb0R\xd1*\xe9\x949\xed\xbf\xae\xbaE" +
	"\xcd\xd35\x8c\x90\xc0\x08!W\xf6\x12\xafS\xa4\x89\xdf" +
	"T:\xf1\x9a\x84\xa8\x83N\xc3H\xc5^\xe2\x87\x84\x00" +
	"\x0e\x098\x84=\xd4**}$\x0a\x03\xad\x8c`C" +
	"]\xc1N\x18m\x8e1\xe4i\x81\x8e^\xa7\x0c\x8f\x07" +
	"\x19\xf2\x82\xc0\xa8N\xbcdJC\x90\x80\xa0>Z\xd0" +
	"8@\xb8\xcc\xc8v:\xb0\xeb\xfc'\xc3f\xd3\x0b\xca" +
	"D\x03\xb3O\xee>;\x97\xbc\x16\xa9\x1e\xbd\xbd'[" +
	"\xdb&_\xf0\x82\xb2\xaeyuU4\xaf5\x12\xb3\xc1" +
	"\xc1\xee\x06\xde8\x91|\x81!k\xbd\x0d\x94\xe9]e" +
	"\xc8\x86q\x0b\xdan\xf1c\"Yc\xc8D\xc0\xe5\xa3" +
	"m\xbbL\xde&\x92\x09C\xbe-pmZ\xc5\xda\x0f" +
	"\x03\x0c\x91\xc0\x10\xe1\xdat\xd9\xd7\xf5\x8bOa?\x09" +
	"\xec'\xa4\x15?\xd6\xc9\x155IG2\x16\xddK\xc6" +
	"J\xab\xe4i?F\xf6\xed\x94\x0a\xb8\xa4\x00\x12\xc0\xae" +
	"\xc4\xccI\xc7\xaa\xd5XU\x8d\x0d\x82\xb6\xb4\xf7w\x89" +
	"\xdd2\xd2\xde`\xc8;=b\x1f\x99\xdeM\x86\xbc\xd7" +
	"G\xec\xae\xd1\xfb\x0eC~j\x88\xa1Ml\xd6Hp" +
	"\x8f!?39\x10\xed\x1c\xcc\x19\xa7|\xb2\x95\xa2N" +
	"\x0e\xfaS\x94\x0b\xbc\xa6\xea\xf0\xfd\x8f\xb9w\xb2\xf0\xa0" +
	"@\xa3/7\xc2R\xbd{\xda\x83\xbd\x7fD\x82i\xe6" +
	"\xa2XMw3\xf1\xbf.\xae\xa30\xe0m\x86\x1f\xef" +
	"\x99\xce\x85\xb5\xe5:\xc3\xed8C>\xba\xfd\x92\x03\x09" +
	"\xf87\x00\x00\xff\xff3\x1aRH"

func init() {
	schemas.Register(schema_f4533cbae6e08506,
		0x8cf178de3c82d431,
		0x98d11ae1c78a24d9,
		0xdbe14b5e7e7c6009,
		0xe0d4e6d68fa24ac0,
		0xe46ab5b4b619e094,
		0xee959a7d96c96641)
}
