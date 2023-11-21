// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/wireguard/v1/config.proto

package wg_pb

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

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrivateKey *PrivateKey `protobuf:"bytes,1,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	Cidr       string      `protobuf:"bytes,2,opt,name=cidr,proto3" json:"cidr,omitempty"`
	Dns        []string    `protobuf:"bytes,3,rep,name=dns,proto3" json:"dns,omitempty"`
	ListenPort int64       `protobuf:"varint,4,opt,name=listen_port,json=listenPort,proto3" json:"listen_port,omitempty"`
	Routes     *Routes     `protobuf:"bytes,5,opt,name=routes,proto3" json:"routes,omitempty"`
	Users      []*User     `protobuf:"bytes,6,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_wireguard_v1_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_proto_wireguard_v1_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_proto_wireguard_v1_config_proto_rawDescGZIP(), []int{0}
}

func (x *Server) GetPrivateKey() *PrivateKey {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *Server) GetCidr() string {
	if x != nil {
		return x.Cidr
	}
	return ""
}

func (x *Server) GetDns() []string {
	if x != nil {
		return x.Dns
	}
	return nil
}

func (x *Server) GetListenPort() int64 {
	if x != nil {
		return x.ListenPort
	}
	return 0
}

func (x *Server) GetRoutes() *Routes {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *Server) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type PrivateKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Store:
	//
	//	*PrivateKey_EnvVar
	Store isPrivateKey_Store `protobuf_oneof:"store"`
}

func (x *PrivateKey) Reset() {
	*x = PrivateKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_wireguard_v1_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateKey) ProtoMessage() {}

func (x *PrivateKey) ProtoReflect() protoreflect.Message {
	mi := &file_proto_wireguard_v1_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateKey.ProtoReflect.Descriptor instead.
func (*PrivateKey) Descriptor() ([]byte, []int) {
	return file_proto_wireguard_v1_config_proto_rawDescGZIP(), []int{1}
}

func (m *PrivateKey) GetStore() isPrivateKey_Store {
	if m != nil {
		return m.Store
	}
	return nil
}

func (x *PrivateKey) GetEnvVar() string {
	if x, ok := x.GetStore().(*PrivateKey_EnvVar); ok {
		return x.EnvVar
	}
	return ""
}

type isPrivateKey_Store interface {
	isPrivateKey_Store()
}

type PrivateKey_EnvVar struct {
	EnvVar string `protobuf:"bytes,1,opt,name=env_var,json=envVar,proto3,oneof"`
}

func (*PrivateKey_EnvVar) isPrivateKey_Store() {}

type Routes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accept []string `protobuf:"bytes,1,rep,name=accept,proto3" json:"accept,omitempty"`
}

func (x *Routes) Reset() {
	*x = Routes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_wireguard_v1_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Routes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Routes) ProtoMessage() {}

func (x *Routes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_wireguard_v1_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Routes.ProtoReflect.Descriptor instead.
func (*Routes) Descriptor() ([]byte, []int) {
	return file_proto_wireguard_v1_config_proto_rawDescGZIP(), []int{2}
}

func (x *Routes) GetAccept() []string {
	if x != nil {
		return x.Accept
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PublicKey string `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Revoked   bool   `protobuf:"varint,3,opt,name=revoked,proto3" json:"revoked,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_wireguard_v1_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_wireguard_v1_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_wireguard_v1_config_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *User) GetRevoked() bool {
	if x != nil {
		return x.Revoked
	}
	return false
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interface *Interface `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
	Peers     []*Peer    `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_wireguard_v1_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_proto_wireguard_v1_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_proto_wireguard_v1_config_proto_rawDescGZIP(), []int{4}
}

func (x *Node) GetInterface() *Interface {
	if x != nil {
		return x.Interface
	}
	return nil
}

func (x *Node) GetPeers() []*Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

type Interface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address    string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	ListenPort int64  `protobuf:"varint,2,opt,name=listen_port,json=listenPort,proto3" json:"listen_port,omitempty"`
	PrivateKey string `protobuf:"bytes,3,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	PreUp      string `protobuf:"bytes,4,opt,name=pre_up,json=preUp,proto3" json:"pre_up,omitempty"`
	PostUp     string `protobuf:"bytes,5,opt,name=post_up,json=postUp,proto3" json:"post_up,omitempty"`
	PreDown    string `protobuf:"bytes,6,opt,name=pre_down,json=preDown,proto3" json:"pre_down,omitempty"`
	PostDown   string `protobuf:"bytes,7,opt,name=post_down,json=postDown,proto3" json:"post_down,omitempty"`
}

func (x *Interface) Reset() {
	*x = Interface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_wireguard_v1_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interface) ProtoMessage() {}

func (x *Interface) ProtoReflect() protoreflect.Message {
	mi := &file_proto_wireguard_v1_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interface.ProtoReflect.Descriptor instead.
func (*Interface) Descriptor() ([]byte, []int) {
	return file_proto_wireguard_v1_config_proto_rawDescGZIP(), []int{5}
}

func (x *Interface) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Interface) GetListenPort() int64 {
	if x != nil {
		return x.ListenPort
	}
	return 0
}

func (x *Interface) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *Interface) GetPreUp() string {
	if x != nil {
		return x.PreUp
	}
	return ""
}

func (x *Interface) GetPostUp() string {
	if x != nil {
		return x.PostUp
	}
	return ""
}

func (x *Interface) GetPreDown() string {
	if x != nil {
		return x.PreDown
	}
	return ""
}

func (x *Interface) GetPostDown() string {
	if x != nil {
		return x.PostDown
	}
	return ""
}

type Peer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey    string `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	PresharedKey string `protobuf:"bytes,2,opt,name=preshared_key,json=presharedKey,proto3" json:"preshared_key,omitempty"`
	AllowedIps   string `protobuf:"bytes,3,opt,name=allowed_ips,json=allowedIps,proto3" json:"allowed_ips,omitempty"`
}

func (x *Peer) Reset() {
	*x = Peer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_wireguard_v1_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peer) ProtoMessage() {}

func (x *Peer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_wireguard_v1_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peer.ProtoReflect.Descriptor instead.
func (*Peer) Descriptor() ([]byte, []int) {
	return file_proto_wireguard_v1_config_proto_rawDescGZIP(), []int{6}
}

func (x *Peer) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *Peer) GetPresharedKey() string {
	if x != nil {
		return x.PresharedKey
	}
	return ""
}

func (x *Peer) GetAllowedIps() string {
	if x != nil {
		return x.AllowedIps
	}
	return ""
}

var File_proto_wireguard_v1_config_proto protoreflect.FileDescriptor

var file_proto_wireguard_v1_config_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72,
	0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x77, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x22,
	0xe2, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0b, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x77, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x64, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x64, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x06,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x77,
	0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x77, 0x69, 0x72, 0x65,
	0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x22, 0x30, 0x0a, 0x0a, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b,
	0x65, 0x79, 0x12, 0x19, 0x0a, 0x07, 0x65, 0x6e, 0x76, 0x5f, 0x76, 0x61, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x65, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x42, 0x07, 0x0a,
	0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x20, 0x0a, 0x06, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x22, 0x53, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x22, 0x67, 0x0a,
	0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x77, 0x69, 0x72, 0x65, 0x67,
	0x75, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x05,
	0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x77, 0x69,
	0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52,
	0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x22, 0xcf, 0x01, 0x0a, 0x09, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x12, 0x15, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x5f, 0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x72, 0x65, 0x55, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f,
	0x75, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x55, 0x70,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x6f, 0x73, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x22, 0x6b, 0x0a, 0x04, 0x50, 0x65, 0x65, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12,
	0x23, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f,
	0x69, 0x70, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x49, 0x70, 0x73, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x73, 0x2f, 0x77, 0x69, 0x72, 0x65,
	0x67, 0x75, 0x61, 0x72, 0x64, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2d, 0x67, 0x6f, 0x2f,
	0x77, 0x69, 0x72, 0x65, 0x67, 0x75, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x67, 0x5f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_wireguard_v1_config_proto_rawDescOnce sync.Once
	file_proto_wireguard_v1_config_proto_rawDescData = file_proto_wireguard_v1_config_proto_rawDesc
)

func file_proto_wireguard_v1_config_proto_rawDescGZIP() []byte {
	file_proto_wireguard_v1_config_proto_rawDescOnce.Do(func() {
		file_proto_wireguard_v1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_wireguard_v1_config_proto_rawDescData)
	})
	return file_proto_wireguard_v1_config_proto_rawDescData
}

var file_proto_wireguard_v1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_wireguard_v1_config_proto_goTypes = []interface{}{
	(*Server)(nil),     // 0: wireguard.v1.Server
	(*PrivateKey)(nil), // 1: wireguard.v1.PrivateKey
	(*Routes)(nil),     // 2: wireguard.v1.Routes
	(*User)(nil),       // 3: wireguard.v1.User
	(*Node)(nil),       // 4: wireguard.v1.Node
	(*Interface)(nil),  // 5: wireguard.v1.Interface
	(*Peer)(nil),       // 6: wireguard.v1.Peer
}
var file_proto_wireguard_v1_config_proto_depIdxs = []int32{
	1, // 0: wireguard.v1.Server.private_key:type_name -> wireguard.v1.PrivateKey
	2, // 1: wireguard.v1.Server.routes:type_name -> wireguard.v1.Routes
	3, // 2: wireguard.v1.Server.users:type_name -> wireguard.v1.User
	5, // 3: wireguard.v1.Node.interface:type_name -> wireguard.v1.Interface
	6, // 4: wireguard.v1.Node.peers:type_name -> wireguard.v1.Peer
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_wireguard_v1_config_proto_init() }
func file_proto_wireguard_v1_config_proto_init() {
	if File_proto_wireguard_v1_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_wireguard_v1_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_proto_wireguard_v1_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateKey); i {
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
		file_proto_wireguard_v1_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Routes); i {
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
		file_proto_wireguard_v1_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_proto_wireguard_v1_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_proto_wireguard_v1_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interface); i {
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
		file_proto_wireguard_v1_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Peer); i {
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
	file_proto_wireguard_v1_config_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*PrivateKey_EnvVar)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_wireguard_v1_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_wireguard_v1_config_proto_goTypes,
		DependencyIndexes: file_proto_wireguard_v1_config_proto_depIdxs,
		MessageInfos:      file_proto_wireguard_v1_config_proto_msgTypes,
	}.Build()
	File_proto_wireguard_v1_config_proto = out.File
	file_proto_wireguard_v1_config_proto_rawDesc = nil
	file_proto_wireguard_v1_config_proto_goTypes = nil
	file_proto_wireguard_v1_config_proto_depIdxs = nil
}
