// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: proto/auth.proto

package auth

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AuthProvider int32

const (
	AuthProvider_EMAIL  AuthProvider = 0
	AuthProvider_GOOGLE AuthProvider = 1
)

// Enum value maps for AuthProvider.
var (
	AuthProvider_name = map[int32]string{
		0: "EMAIL",
		1: "GOOGLE",
	}
	AuthProvider_value = map[string]int32{
		"EMAIL":  0,
		"GOOGLE": 1,
	}
)

func (x AuthProvider) Enum() *AuthProvider {
	p := new(AuthProvider)
	*p = x
	return p
}

func (x AuthProvider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthProvider) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_auth_proto_enumTypes[0].Descriptor()
}

func (AuthProvider) Type() protoreflect.EnumType {
	return &file_proto_auth_proto_enumTypes[0]
}

func (x AuthProvider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthProvider.Descriptor instead.
func (AuthProvider) EnumDescriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{0}
}

type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider   AuthProvider `protobuf:"varint,1,opt,name=provider,proto3,enum=auth.AuthProvider" json:"provider,omitempty"`
	ProviderId string       `protobuf:"bytes,2,opt,name=providerId,proto3" json:"providerId,omitempty"`
	Password   string       `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthTokenRequest) Reset() {
	*x = AuthTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthTokenRequest) ProtoMessage() {}

func (x *AuthTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthTokenRequest.ProtoReflect.Descriptor instead.
func (*AuthTokenRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthTokenRequest) GetProvider() AuthProvider {
	if x != nil {
		return x.Provider
	}
	return AuthProvider_EMAIL
}

func (x *AuthTokenRequest) GetProviderId() string {
	if x != nil {
		return x.ProviderId
	}
	return ""
}

func (x *AuthTokenRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authenticated bool `protobuf:"varint,1,opt,name=authenticated,proto3" json:"authenticated,omitempty"`
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{2}
}

func (x *AuthResponse) GetAuthenticated() bool {
	if x != nil {
		return x.Authenticated
	}
	return false
}

type AuthToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthToken) Reset() {
	*x = AuthToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthToken) ProtoMessage() {}

func (x *AuthToken) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthToken.ProtoReflect.Descriptor instead.
func (*AuthToken) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{3}
}

func (x *AuthToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Okay bool `protobuf:"varint,1,opt,name=okay,proto3" json:"okay,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{4}
}

func (x *Result) GetOkay() bool {
	if x != nil {
		return x.Okay
	}
	return false
}

var File_proto_auth_proto protoreflect.FileDescriptor

var file_proto_auth_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x23, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x7e, 0x0a,
	0x10, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2e, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x34, 0x0a,
	0x0c, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x22, 0x21, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x1c, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6f, 0x6b, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x6f, 0x6b, 0x61, 0x79, 0x2a, 0x25, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x10, 0x01, 0x32, 0xae, 0x01, 0x0a, 0x0b,
	0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x49,
	0x73, 0x41, 0x75, 0x74, 0x68, 0x12, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a,
	0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x06, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x12, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06,
	0x2e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_auth_proto_rawDescOnce sync.Once
	file_proto_auth_proto_rawDescData = file_proto_auth_proto_rawDesc
)

func file_proto_auth_proto_rawDescGZIP() []byte {
	file_proto_auth_proto_rawDescOnce.Do(func() {
		file_proto_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_auth_proto_rawDescData)
	})
	return file_proto_auth_proto_rawDescData
}

var file_proto_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_auth_proto_goTypes = []interface{}{
	(AuthProvider)(0),        // 0: auth.AuthProvider
	(*AuthRequest)(nil),      // 1: auth.AuthRequest
	(*AuthTokenRequest)(nil), // 2: auth.AuthTokenRequest
	(*AuthResponse)(nil),     // 3: auth.AuthResponse
	(*AuthToken)(nil),        // 4: auth.AuthToken
	(*Result)(nil),           // 5: auth.Result
}
var file_proto_auth_proto_depIdxs = []int32{
	0, // 0: auth.AuthTokenRequest.provider:type_name -> auth.AuthProvider
	1, // 1: auth.AuthService.IsAuth:input_type -> auth.AuthRequest
	2, // 2: auth.AuthService.GenerateToken:input_type -> auth.AuthTokenRequest
	2, // 3: auth.AuthService.SignUp:input_type -> auth.AuthTokenRequest
	3, // 4: auth.AuthService.IsAuth:output_type -> auth.AuthResponse
	4, // 5: auth.AuthService.GenerateToken:output_type -> auth.AuthToken
	5, // 6: auth.AuthService.SignUp:output_type -> auth.Result
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_auth_proto_init() }
func file_proto_auth_proto_init() {
	if File_proto_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest); i {
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
		file_proto_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthTokenRequest); i {
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
		file_proto_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResponse); i {
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
		file_proto_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthToken); i {
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
		file_proto_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_auth_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_auth_proto_goTypes,
		DependencyIndexes: file_proto_auth_proto_depIdxs,
		EnumInfos:         file_proto_auth_proto_enumTypes,
		MessageInfos:      file_proto_auth_proto_msgTypes,
	}.Build()
	File_proto_auth_proto = out.File
	file_proto_auth_proto_rawDesc = nil
	file_proto_auth_proto_goTypes = nil
	file_proto_auth_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	IsAuth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	GenerateToken(ctx context.Context, in *AuthTokenRequest, opts ...grpc.CallOption) (*AuthToken, error)
	SignUp(ctx context.Context, in *AuthTokenRequest, opts ...grpc.CallOption) (*Result, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) IsAuth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/IsAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GenerateToken(ctx context.Context, in *AuthTokenRequest, opts ...grpc.CallOption) (*AuthToken, error) {
	out := new(AuthToken)
	err := c.cc.Invoke(ctx, "/auth.AuthService/GenerateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SignUp(ctx context.Context, in *AuthTokenRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/auth.AuthService/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	IsAuth(context.Context, *AuthRequest) (*AuthResponse, error)
	GenerateToken(context.Context, *AuthTokenRequest) (*AuthToken, error)
	SignUp(context.Context, *AuthTokenRequest) (*Result, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) IsAuth(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAuth not implemented")
}
func (*UnimplementedAuthServiceServer) GenerateToken(context.Context, *AuthTokenRequest) (*AuthToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateToken not implemented")
}
func (*UnimplementedAuthServiceServer) SignUp(context.Context, *AuthTokenRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_IsAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).IsAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/IsAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).IsAuth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/GenerateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GenerateToken(ctx, req.(*AuthTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SignUp(ctx, req.(*AuthTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsAuth",
			Handler:    _AuthService_IsAuth_Handler,
		},
		{
			MethodName: "GenerateToken",
			Handler:    _AuthService_GenerateToken_Handler,
		},
		{
			MethodName: "SignUp",
			Handler:    _AuthService_SignUp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/auth.proto",
}
