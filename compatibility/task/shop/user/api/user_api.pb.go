/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: user_api.proto

package api

import (
	reflect "reflect"
	sync "sync"
)

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"

	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *GetInfoReq) Reset() {
	*x = GetInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoReq) ProtoMessage() {}

func (x *GetInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoReq.ProtoReflect.Descriptor instead.
func (*GetInfoReq) Descriptor() ([]byte, []int) {
	return file_user_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetInfoReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_user_api_proto_rawDescGZIP(), []int{1}
}

func (x *LoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *RegisterResp) Reset() {
	*x = RegisterResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResp) ProtoMessage() {}

func (x *RegisterResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResp.ProtoReflect.Descriptor instead.
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return file_user_api_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	RealName string `protobuf:"bytes,3,opt,name=RealName,proto3" json:"RealName,omitempty"`
	Mail     string `protobuf:"bytes,4,opt,name=Mail,proto3" json:"Mail,omitempty"`
	Phone    string `protobuf:"bytes,5,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Env      string `protobuf:"bytes,6,opt,name=Env,proto3" json:"Env,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_api_proto_msgTypes[3]
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
	return file_user_api_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *User) GetMail() string {
	if x != nil {
		return x.Mail
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *User) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

var File_user_api_proto protoreflect.FileDescriptor

var file_user_api_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x28, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x64, 0x75, 0x62,
	0x62, 0x6f, 0x67, 0x6f, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x68, 0x6f,
	0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x28, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x42, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x28, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x96, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x4d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4d,
	0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x6e, 0x76,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x45, 0x6e, 0x76, 0x32, 0xd3, 0x03, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x72, 0x0a, 0x08, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x2e, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70,
	0x61, 0x63, 0x68, 0x65, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x67, 0x6f, 0x2e, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x36, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70,
	0x61, 0x63, 0x68, 0x65, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x67, 0x6f, 0x2e, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x6b, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x32, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61,
	0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x67, 0x6f, 0x2e, 0x73, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x2e, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x67,
	0x6f, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x72, 0x0a, 0x0c,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x32, 0x2e, 0x6f,
	0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x67,
	0x6f, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x1a, 0x2e, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x64, 0x75,
	0x62, 0x62, 0x6f, 0x67, 0x6f, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x6f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x64, 0x75, 0x62, 0x62, 0x6f, 0x67, 0x6f,
	0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x1a, 0x2e, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x64,
	0x75, 0x62, 0x62, 0x6f, 0x67, 0x6f, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_user_api_proto_rawDescOnce sync.Once
	file_user_api_proto_rawDescData = file_user_api_proto_rawDesc
)

func file_user_api_proto_rawDescGZIP() []byte {
	file_user_api_proto_rawDescOnce.Do(func() {
		file_user_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_api_proto_rawDescData)
	})
	return file_user_api_proto_rawDescData
}

var file_user_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_user_api_proto_goTypes = []interface{}{
	(*GetInfoReq)(nil),   // 0: org.apache.dubbogo.samples.shop.user.api.GetInfoReq
	(*LoginReq)(nil),     // 1: org.apache.dubbogo.samples.shop.user.api.LoginReq
	(*RegisterResp)(nil), // 2: org.apache.dubbogo.samples.shop.user.api.RegisterResp
	(*User)(nil),         // 3: org.apache.dubbogo.samples.shop.user.api.User
}
var file_user_api_proto_depIdxs = []int32{
	3, // 0: org.apache.dubbogo.samples.shop.user.api.UserService.Register:input_type -> org.apache.dubbogo.samples.shop.user.api.User
	1, // 1: org.apache.dubbogo.samples.shop.user.api.UserService.Login:input_type -> org.apache.dubbogo.samples.shop.user.api.LoginReq
	1, // 2: org.apache.dubbogo.samples.shop.user.api.UserService.TimeoutLogin:input_type -> org.apache.dubbogo.samples.shop.user.api.LoginReq
	0, // 3: org.apache.dubbogo.samples.shop.user.api.UserService.GetInfo:input_type -> org.apache.dubbogo.samples.shop.user.api.GetInfoReq
	2, // 4: org.apache.dubbogo.samples.shop.user.api.UserService.Register:output_type -> org.apache.dubbogo.samples.shop.user.api.RegisterResp
	3, // 5: org.apache.dubbogo.samples.shop.user.api.UserService.Login:output_type -> org.apache.dubbogo.samples.shop.user.api.User
	3, // 6: org.apache.dubbogo.samples.shop.user.api.UserService.TimeoutLogin:output_type -> org.apache.dubbogo.samples.shop.user.api.User
	3, // 7: org.apache.dubbogo.samples.shop.user.api.UserService.GetInfo:output_type -> org.apache.dubbogo.samples.shop.user.api.User
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_api_proto_init() }
func file_user_api_proto_init() {
	if File_user_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoReq); i {
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
		file_user_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_user_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResp); i {
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
		file_user_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_api_proto_goTypes,
		DependencyIndexes: file_user_api_proto_depIdxs,
		MessageInfos:      file_user_api_proto_msgTypes,
	}.Build()
	File_user_api_proto = out.File
	file_user_api_proto_rawDesc = nil
	file_user_api_proto_goTypes = nil
	file_user_api_proto_depIdxs = nil
}
