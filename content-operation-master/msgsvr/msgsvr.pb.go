// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: msgsvr.proto

package msgsvr

import (
	encoding_xml "encoding/xml"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// CheckSignatureReq 签名校验接口请求
type CheckSignatureReq struct {
	Signature            string   `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Nonce                int64    `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Echostr              string   `protobuf:"bytes,4,opt,name=echostr,proto3" json:"echostr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckSignatureReq) Reset()         { *m = CheckSignatureReq{} }
func (m *CheckSignatureReq) String() string { return proto.CompactTextString(m) }
func (*CheckSignatureReq) ProtoMessage()    {}
func (*CheckSignatureReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_48ecc092994564ed, []int{0}
}
func (m *CheckSignatureReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckSignatureReq.Unmarshal(m, b)
}
func (m *CheckSignatureReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckSignatureReq.Marshal(b, m, deterministic)
}
func (m *CheckSignatureReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckSignatureReq.Merge(m, src)
}
func (m *CheckSignatureReq) XXX_Size() int {
	return xxx_messageInfo_CheckSignatureReq.Size(m)
}
func (m *CheckSignatureReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckSignatureReq.DiscardUnknown(m)
}

var xxx_messageInfo_CheckSignatureReq proto.InternalMessageInfo

func (m *CheckSignatureReq) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *CheckSignatureReq) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *CheckSignatureReq) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *CheckSignatureReq) GetEchostr() string {
	if m != nil {
		return m.Echostr
	}
	return ""
}

// CheckSignatureReq 签名校验接口回报
type CheckSignatureRsp struct {
	Echostr              string   `protobuf:"bytes,1,opt,name=echostr,proto3" json:"echostr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckSignatureRsp) Reset()         { *m = CheckSignatureRsp{} }
func (m *CheckSignatureRsp) String() string { return proto.CompactTextString(m) }
func (*CheckSignatureRsp) ProtoMessage()    {}
func (*CheckSignatureRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_48ecc092994564ed, []int{1}
}
func (m *CheckSignatureRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckSignatureRsp.Unmarshal(m, b)
}
func (m *CheckSignatureRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckSignatureRsp.Marshal(b, m, deterministic)
}
func (m *CheckSignatureRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckSignatureRsp.Merge(m, src)
}
func (m *CheckSignatureRsp) XXX_Size() int {
	return xxx_messageInfo_CheckSignatureRsp.Size(m)
}
func (m *CheckSignatureRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckSignatureRsp.DiscardUnknown(m)
}

var xxx_messageInfo_CheckSignatureRsp proto.InternalMessageInfo

func (m *CheckSignatureRsp) GetEchostr() string {
	if m != nil {
		return m.Echostr
	}
	return ""
}

// MessageReq 接受普通消息接口请求
type MessageReq struct {
	ToUserName           string   `protobuf:"bytes,1,opt,name=to_user_name,json=toUserName,proto3" json:"to_user_name,omitempty"`
	FromUserName         string   `protobuf:"bytes,2,opt,name=from_user_name,json=fromUserName,proto3" json:"from_user_name,omitempty"`
	CreateTime           int64    `protobuf:"varint,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	MsgType              string   `protobuf:"bytes,4,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`
	Content              string   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	Recognition          string   `protobuf:"bytes,6,opt,name=recognition,proto3" json:"recognition,omitempty"`
	MsgId                int64    `protobuf:"varint,7,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	MsgDataId            int64    `protobuf:"varint,8,opt,name=msg_data_id,json=msgDataId,proto3" json:"msg_data_id,omitempty"`
	Idx                  int64    `protobuf:"varint,9,opt,name=idx,proto3" json:"idx,omitempty"`
	Event                string   `protobuf:"bytes,10,opt,name=event,proto3" json:"event,omitempty"`
	EventKey             string   `protobuf:"bytes,11,opt,name=event_key,json=eventKey,proto3" json:"event_key,omitempty"`
	Ticket               string   `protobuf:"bytes,12,opt,name=ticket,proto3" json:"ticket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageReq) Reset()         { *m = MessageReq{} }
func (m *MessageReq) String() string { return proto.CompactTextString(m) }
func (*MessageReq) ProtoMessage()    {}
func (*MessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_48ecc092994564ed, []int{2}
}
func (m *MessageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageReq.Unmarshal(m, b)
}
func (m *MessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageReq.Marshal(b, m, deterministic)
}
func (m *MessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageReq.Merge(m, src)
}
func (m *MessageReq) XXX_Size() int {
	return xxx_messageInfo_MessageReq.Size(m)
}
func (m *MessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_MessageReq proto.InternalMessageInfo

func (m *MessageReq) GetToUserName() string {
	if m != nil {
		return m.ToUserName
	}
	return ""
}

func (m *MessageReq) GetFromUserName() string {
	if m != nil {
		return m.FromUserName
	}
	return ""
}

func (m *MessageReq) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *MessageReq) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *MessageReq) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *MessageReq) GetRecognition() string {
	if m != nil {
		return m.Recognition
	}
	return ""
}

func (m *MessageReq) GetMsgId() int64 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *MessageReq) GetMsgDataId() int64 {
	if m != nil {
		return m.MsgDataId
	}
	return 0
}

func (m *MessageReq) GetIdx() int64 {
	if m != nil {
		return m.Idx
	}
	return 0
}

func (m *MessageReq) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *MessageReq) GetEventKey() string {
	if m != nil {
		return m.EventKey
	}
	return ""
}

func (m *MessageReq) GetTicket() string {
	if m != nil {
		return m.Ticket
	}
	return ""
}

// MessageRsp 接受普通消息接口回包
type MessageRsp struct {
	XMLName              *encoding_xml.Name `protobuf:"bytes,1,opt,name=XMLName,proto3,customtype=encoding/xml.Name" json:"XMLName,omitempty" xml:"xml"`
	ToUserName           *ToUserName        `protobuf:"bytes,2,opt,name=to_user_name,json=toUserName,proto3" json:"to_user_name,omitempty"`
	FromUserName         *FromUserName      `protobuf:"bytes,3,opt,name=from_user_name,json=fromUserName,proto3" json:"from_user_name,omitempty"`
	CreateTime           int64              `protobuf:"varint,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	MsgType              string             `protobuf:"bytes,5,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`
	Content              *Content           `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Video                *Video             `protobuf:"bytes,7,opt,name=video,proto3" json:"video,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MessageRsp) Reset()         { *m = MessageRsp{} }
func (m *MessageRsp) String() string { return proto.CompactTextString(m) }
func (*MessageRsp) ProtoMessage()    {}
func (*MessageRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_48ecc092994564ed, []int{3}
}
func (m *MessageRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageRsp.Unmarshal(m, b)
}
func (m *MessageRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageRsp.Marshal(b, m, deterministic)
}
func (m *MessageRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageRsp.Merge(m, src)
}
func (m *MessageRsp) XXX_Size() int {
	return xxx_messageInfo_MessageRsp.Size(m)
}
func (m *MessageRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageRsp.DiscardUnknown(m)
}

var xxx_messageInfo_MessageRsp proto.InternalMessageInfo

func (m *MessageRsp) GetToUserName() *ToUserName {
	if m != nil {
		return m.ToUserName
	}
	return nil
}

func (m *MessageRsp) GetFromUserName() *FromUserName {
	if m != nil {
		return m.FromUserName
	}
	return nil
}

func (m *MessageRsp) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *MessageRsp) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *MessageRsp) GetContent() *Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *MessageRsp) GetVideo() *Video {
	if m != nil {
		return m.Video
	}
	return nil
}

type Video struct {
	MediaId              *MediaId `protobuf:"bytes,1,opt,name=media_id,json=mediaId,proto3" json:"media_id,omitempty"`
	Title                *Title   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Video) Reset()         { *m = Video{} }
func (m *Video) String() string { return proto.CompactTextString(m) }
func (*Video) ProtoMessage()    {}
func (*Video) Descriptor() ([]byte, []int) {
	return fileDescriptor_48ecc092994564ed, []int{4}
}
func (m *Video) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Video.Unmarshal(m, b)
}
func (m *Video) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Video.Marshal(b, m, deterministic)
}
func (m *Video) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Video.Merge(m, src)
}
func (m *Video) XXX_Size() int {
	return xxx_messageInfo_Video.Size(m)
}
func (m *Video) XXX_DiscardUnknown() {
	xxx_messageInfo_Video.DiscardUnknown(m)
}

var xxx_messageInfo_Video proto.InternalMessageInfo

func (m *Video) GetMediaId() *MediaId {
	if m != nil {
		return m.MediaId
	}
	return nil
}

func (m *Video) GetTitle() *Title {
	if m != nil {
		return m.Title
	}
	return nil
}

type ToUserName struct {
	ToUserName           string   `protobuf:"bytes,1,opt,name=to_user_name,json=toUserName,proto3" json:"to_user_name,omitempty" xml:",cdata"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToUserName) Reset()         { *m = ToUserName{} }
func (m *ToUserName) String() string { return proto.CompactTextString(m) }
func (*ToUserName) ProtoMessage()    {}
func (*ToUserName) Descriptor() ([]byte, []int) {
	return fileDescriptor_48ecc092994564ed, []int{5}
}
func (m *ToUserName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToUserName.Unmarshal(m, b)
}
func (m *ToUserName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToUserName.Marshal(b, m, deterministic)
}
func (m *ToUserName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToUserName.Merge(m, src)
}
func (m *ToUserName) XXX_Size() int {
	return xxx_messageInfo_ToUserName.Size(m)
}
func (m *ToUserName) XXX_DiscardUnknown() {
	xxx_messageInfo_ToUserName.DiscardUnknown(m)
}

var xxx_messageInfo_ToUserName proto.InternalMessageInfo

func (m *ToUserName) GetToUserName() string {
	if m != nil {
		return m.ToUserName
	}
	return ""
}

type FromUserName struct {
	FromUserName         string   `protobuf:"bytes,1,opt,name=from_user_name,json=fromUserName,proto3" json:"from_user_name,omitempty" xml:",cdata"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FromUserName) Reset()         { *m = FromUserName{} }
func (m *FromUserName) String() string { return proto.CompactTextString(m) }
func (*FromUserName) ProtoMessage()    {}
func (*FromUserName) Descriptor() ([]byte, []int) {
	return fileDescriptor_48ecc092994564ed, []int{6}
}
func (m *FromUserName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FromUserName.Unmarshal(m, b)
}
func (m *FromUserName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FromUserName.Marshal(b, m, deterministic)
}
func (m *FromUserName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FromUserName.Merge(m, src)
}
func (m *FromUserName) XXX_Size() int {
	return xxx_messageInfo_FromUserName.Size(m)
}
func (m *FromUserName) XXX_DiscardUnknown() {
	xxx_messageInfo_FromUserName.DiscardUnknown(m)
}

var xxx_messageInfo_FromUserName proto.InternalMessageInfo

func (m *FromUserName) GetFromUserName() string {
	if m != nil {
		return m.FromUserName
	}
	return ""
}

type Content struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty" xml:",cdata"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Content) Reset()         { *m = Content{} }
func (m *Content) String() string { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()    {}
func (*Content) Descriptor() ([]byte, []int) {
	return fileDescriptor_48ecc092994564ed, []int{7}
}
func (m *Content) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Content.Unmarshal(m, b)
}
func (m *Content) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Content.Marshal(b, m, deterministic)
}
func (m *Content) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Content.Merge(m, src)
}
func (m *Content) XXX_Size() int {
	return xxx_messageInfo_Content.Size(m)
}
func (m *Content) XXX_DiscardUnknown() {
	xxx_messageInfo_Content.DiscardUnknown(m)
}

var xxx_messageInfo_Content proto.InternalMessageInfo

func (m *Content) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type MediaId struct {
	MediaId              string   `protobuf:"bytes,1,opt,name=media_id,json=mediaId,proto3" json:"media_id,omitempty" xml:",cdata"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MediaId) Reset()         { *m = MediaId{} }
func (m *MediaId) String() string { return proto.CompactTextString(m) }
func (*MediaId) ProtoMessage()    {}
func (*MediaId) Descriptor() ([]byte, []int) {
	return fileDescriptor_48ecc092994564ed, []int{8}
}
func (m *MediaId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaId.Unmarshal(m, b)
}
func (m *MediaId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaId.Marshal(b, m, deterministic)
}
func (m *MediaId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaId.Merge(m, src)
}
func (m *MediaId) XXX_Size() int {
	return xxx_messageInfo_MediaId.Size(m)
}
func (m *MediaId) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaId.DiscardUnknown(m)
}

var xxx_messageInfo_MediaId proto.InternalMessageInfo

func (m *MediaId) GetMediaId() string {
	if m != nil {
		return m.MediaId
	}
	return ""
}

type Title struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" xml:",cdata"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Title) Reset()         { *m = Title{} }
func (m *Title) String() string { return proto.CompactTextString(m) }
func (*Title) ProtoMessage()    {}
func (*Title) Descriptor() ([]byte, []int) {
	return fileDescriptor_48ecc092994564ed, []int{9}
}
func (m *Title) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Title.Unmarshal(m, b)
}
func (m *Title) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Title.Marshal(b, m, deterministic)
}
func (m *Title) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Title.Merge(m, src)
}
func (m *Title) XXX_Size() int {
	return xxx_messageInfo_Title.Size(m)
}
func (m *Title) XXX_DiscardUnknown() {
	xxx_messageInfo_Title.DiscardUnknown(m)
}

var xxx_messageInfo_Title proto.InternalMessageInfo

func (m *Title) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func init() {
	proto.RegisterType((*CheckSignatureReq)(nil), "msgsvr.CheckSignatureReq")
	proto.RegisterType((*CheckSignatureRsp)(nil), "msgsvr.CheckSignatureRsp")
	proto.RegisterType((*MessageReq)(nil), "msgsvr.MessageReq")
	proto.RegisterType((*MessageRsp)(nil), "msgsvr.MessageRsp")
	proto.RegisterType((*Video)(nil), "msgsvr.Video")
	proto.RegisterType((*ToUserName)(nil), "msgsvr.ToUserName")
	proto.RegisterType((*FromUserName)(nil), "msgsvr.FromUserName")
	proto.RegisterType((*Content)(nil), "msgsvr.Content")
	proto.RegisterType((*MediaId)(nil), "msgsvr.MediaId")
	proto.RegisterType((*Title)(nil), "msgsvr.Title")
}

func init() { proto.RegisterFile("msgsvr.proto", fileDescriptor_48ecc092994564ed) }

var fileDescriptor_48ecc092994564ed = []byte{
	// 763 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0xd6, 0xd8, 0xf1, 0x5f, 0xd9, 0x84, 0xa4, 0x15, 0x60, 0x62, 0x22, 0x6c, 0x35, 0x08, 0x42,
	0x48, 0x3c, 0x92, 0x81, 0x08, 0x59, 0x39, 0xa0, 0x04, 0x45, 0x8a, 0xc0, 0x39, 0x38, 0x06, 0x45,
	0x5c, 0xac, 0xce, 0x4c, 0xd3, 0x19, 0xe2, 0x9e, 0x1e, 0xa6, 0xdb, 0x96, 0x7d, 0x62, 0xb5, 0xaf,
	0xb0, 0x2f, 0xb5, 0xf7, 0xbd, 0xed, 0x21, 0xda, 0xc3, 0x3e, 0x41, 0x9e, 0x60, 0xd5, 0xdd, 0x33,
	0xfe, 0x89, 0xe3, 0x68, 0x6f, 0x5d, 0xdf, 0x57, 0x55, 0xdd, 0xf5, 0x7d, 0x33, 0x05, 0x35, 0x2e,
	0x99, 0x1c, 0x27, 0xad, 0x38, 0x11, 0x4a, 0xa0, 0xa2, 0x8d, 0xea, 0x7b, 0x4c, 0x08, 0x36, 0xa4,
	0x1e, 0x89, 0x43, 0x8f, 0x44, 0x91, 0x50, 0x44, 0x85, 0x22, 0x92, 0x36, 0xab, 0xfe, 0xc5, 0x98,
	0x0c, 0xc3, 0x80, 0x28, 0xea, 0x65, 0x87, 0x94, 0xd8, 0x61, 0x82, 0x09, 0x73, 0xf4, 0xf4, 0xc9,
	0xa2, 0xf8, 0x7f, 0xd8, 0x3e, 0xbb, 0xa5, 0xfe, 0xdd, 0x55, 0xc8, 0x22, 0xa2, 0x46, 0x09, 0xed,
	0xd1, 0xff, 0xd0, 0x1e, 0x54, 0x64, 0x16, 0xbb, 0x4e, 0xd3, 0xd9, 0xaf, 0xf4, 0xe6, 0x80, 0x66,
	0x55, 0xc8, 0xa9, 0x54, 0x84, 0xc7, 0x6e, 0xae, 0xe9, 0xec, 0xe7, 0x7b, 0x73, 0x00, 0xed, 0x40,
	0x21, 0x12, 0x91, 0x4f, 0xdd, 0xbc, 0x61, 0x6c, 0x80, 0x5c, 0x28, 0x51, 0xff, 0x56, 0x48, 0x95,
	0xb8, 0x1b, 0xa6, 0x5f, 0x16, 0xe2, 0xa3, 0x95, 0x07, 0xc8, 0x78, 0x31, 0xdd, 0x59, 0x4e, 0x7f,
	0x97, 0x03, 0xe8, 0x52, 0x29, 0x09, 0x33, 0x2f, 0x6d, 0x42, 0x4d, 0x89, 0xc1, 0x48, 0xd2, 0x64,
	0x10, 0x11, 0x9e, 0x3d, 0x16, 0x94, 0xf8, 0x53, 0xd2, 0xe4, 0x92, 0x70, 0x8a, 0xbe, 0x81, 0xcd,
	0x7f, 0x12, 0xc1, 0x17, 0x72, 0x72, 0x26, 0xa7, 0xa6, 0xd1, 0x59, 0x56, 0x03, 0xaa, 0x7e, 0x42,
	0x89, 0xa2, 0x03, 0x3d, 0x49, 0xfa, 0x76, 0xb0, 0x50, 0x3f, 0xe4, 0x14, 0xed, 0x42, 0x99, 0x4b,
	0x36, 0x50, 0xd3, 0x98, 0x66, 0x13, 0x70, 0xc9, 0xfa, 0xd3, 0xd8, 0xcc, 0xe6, 0x8b, 0x48, 0xd1,
	0x48, 0xb9, 0x05, 0xcb, 0xa4, 0x21, 0x6a, 0x42, 0x35, 0xa1, 0xbe, 0x60, 0x51, 0xa8, 0x1d, 0x72,
	0x8b, 0x86, 0x5d, 0x84, 0xd0, 0x67, 0xa0, 0x5d, 0x1d, 0x84, 0x81, 0x5b, 0xb2, 0x72, 0x71, 0xc9,
	0x2e, 0x02, 0xf4, 0x15, 0x54, 0x35, 0x1c, 0x10, 0x45, 0x34, 0x57, 0xb6, 0x22, 0x73, 0xc9, 0x7e,
	0x23, 0x8a, 0x5c, 0x04, 0x68, 0x0b, 0xf2, 0x61, 0x30, 0x71, 0x2b, 0x06, 0xd7, 0x47, 0x2d, 0x3b,
	0x1d, 0xeb, 0x27, 0x80, 0xb9, 0xc4, 0x06, 0xe8, 0x4b, 0xa8, 0x98, 0xc3, 0xe0, 0x8e, 0x4e, 0xdd,
	0xaa, 0x61, 0xca, 0x06, 0xf8, 0x9d, 0x4e, 0xd1, 0xe7, 0x50, 0x54, 0xa1, 0x7f, 0x47, 0x95, 0x5b,
	0x33, 0x4c, 0x1a, 0xe1, 0xd7, 0x0b, 0x12, 0xcb, 0x18, 0x9d, 0x40, 0xe9, 0xba, 0xfb, 0xc7, 0x65,
	0xa6, 0x6e, 0xed, 0x14, 0xbf, 0xbd, 0x6f, 0x6c, 0xd3, 0xc8, 0x17, 0x41, 0x18, 0x31, 0x6f, 0xc2,
	0x87, 0x2d, 0x4d, 0x3e, 0xdc, 0x37, 0x2a, 0x13, 0x3e, 0xec, 0xe0, 0x09, 0x1f, 0xe2, 0x5e, 0x56,
	0x82, 0x7e, 0x7a, 0x64, 0x90, 0x16, 0xbf, 0xda, 0x46, 0xad, 0xf4, 0xcb, 0xee, 0xcf, 0x8c, 0x5a,
	0x32, 0xad, 0xb3, 0x62, 0x5a, 0xde, 0xd4, 0xed, 0x64, 0x75, 0xe7, 0x0b, 0xe6, 0x3d, 0x6f, 0xe5,
	0xc6, 0xb3, 0x56, 0x16, 0x96, 0xad, 0xfc, 0x7e, 0x6e, 0x65, 0xd1, 0x5c, 0xf8, 0x69, 0x76, 0xe1,
	0x99, 0x85, 0xe7, 0xde, 0x7e, 0x0d, 0x85, 0x71, 0x18, 0x50, 0x61, 0x8c, 0xab, 0xb6, 0x3f, 0xc9,
	0x12, 0xff, 0xd2, 0x60, 0xcf, 0x72, 0xf8, 0x1a, 0x0a, 0x26, 0x46, 0x07, 0x50, 0xe6, 0x34, 0x08,
	0x8d, 0x9b, 0xce, 0x72, 0xe7, 0xae, 0xc6, 0x2f, 0x82, 0x5e, 0x89, 0xdb, 0x83, 0xee, 0xac, 0x42,
	0x35, 0xcc, 0xb4, 0x9a, 0x75, 0xee, 0x6b, 0xb0, 0x67, 0x39, 0xfc, 0x2b, 0xc0, 0x5c, 0x3b, 0xd4,
	0x7e, 0xea, 0x37, 0x38, 0xdd, 0x7a, 0xb8, 0x6f, 0xd4, 0x8c, 0x27, 0x87, 0xbe, 0xfe, 0x94, 0xf0,
	0xa2, 0xc6, 0xf8, 0x1c, 0x6a, 0x8b, 0x2a, 0xa2, 0xe3, 0x15, 0xcd, 0xd7, 0x75, 0x59, 0xd2, 0x1b,
	0xff, 0x0c, 0xa5, 0x54, 0x1c, 0x74, 0x30, 0x97, 0x6f, 0x5d, 0x6d, 0x96, 0x80, 0x8f, 0xa1, 0x94,
	0x4e, 0x8e, 0x7e, 0x78, 0x24, 0xce, 0x93, 0x75, 0xa9, 0x3a, 0xd8, 0x83, 0x82, 0x11, 0x02, 0x7d,
	0x9b, 0xc9, 0xb4, 0xae, 0xc4, 0xd2, 0xed, 0x07, 0x07, 0x8a, 0x5d, 0xc9, 0xae, 0xc6, 0x09, 0x7a,
	0xe1, 0xc0, 0xe6, 0xf2, 0xb2, 0x41, 0xbb, 0x33, 0x83, 0x1f, 0x6f, 0xc1, 0xfa, 0x3a, 0x4a, 0xc6,
	0xf8, 0x97, 0x97, 0x6f, 0xde, 0xbf, 0xca, 0xb5, 0xd1, 0x77, 0x5e, 0x3a, 0xce, 0x91, 0x88, 0x69,
	0x62, 0xf6, 0xf0, 0x51, 0x9c, 0x88, 0xc9, 0xd4, 0xb3, 0xa5, 0x1e, 0xb7, 0xff, 0xd1, 0x4d, 0xb6,
	0xbf, 0x10, 0xd5, 0x63, 0x1b, 0x0c, 0xa1, 0xf9, 0x17, 0x90, 0xed, 0xb3, 0xfa, 0x0a, 0x26, 0x63,
	0xdc, 0x36, 0x97, 0x1d, 0xe2, 0x8f, 0xbd, 0xac, 0xe3, 0x1c, 0x9c, 0x9e, 0xfc, 0xdd, 0x61, 0xa1,
	0xba, 0x1d, 0xdd, 0xb4, 0x7c, 0xc1, 0xbd, 0x7f, 0x69, 0x24, 0x69, 0xc4, 0x46, 0xc2, 0xb3, 0xfb,
	0x7f, 0xb5, 0x0f, 0x27, 0x52, 0xd1, 0x24, 0x6d, 0x74, 0x53, 0x34, 0x69, 0x3f, 0x7e, 0x08, 0x00,
	0x00, 0xff, 0xff, 0xc6, 0xf6, 0xaa, 0xa9, 0x80, 0x06, 0x00, 0x00,
}
