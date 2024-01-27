// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: plugin_api.proto

package protoapi

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Source struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepoUri string `protobuf:"bytes,1,opt,name=repoUri,proto3" json:"repoUri,omitempty"`
}

func (x *Source) Reset() {
	*x = Source{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source) ProtoMessage() {}

func (x *Source) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Source.ProtoReflect.Descriptor instead.
func (*Source) Descriptor() ([]byte, []int) {
	return file_plugin_api_proto_rawDescGZIP(), []int{0}
}

func (x *Source) GetRepoUri() string {
	if x != nil {
		return x.RepoUri
	}
	return ""
}

type Release struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	GithubReleaseTag string               `protobuf:"bytes,2,opt,name=githubReleaseTag,proto3" json:"githubReleaseTag,omitempty"`
	PublishedAt      *timestamp.Timestamp `protobuf:"bytes,3,opt,name=publishedAt,proto3" json:"publishedAt,omitempty"`
	IsDeprecated     bool                 `protobuf:"varint,4,opt,name=isDeprecated,proto3" json:"isDeprecated,omitempty"`
	DownloadURI      string               `protobuf:"bytes,5,opt,name=DownloadURI,proto3" json:"DownloadURI,omitempty"`
}

func (x *Release) Reset() {
	*x = Release{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Release) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Release) ProtoMessage() {}

func (x *Release) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Release.ProtoReflect.Descriptor instead.
func (*Release) Descriptor() ([]byte, []int) {
	return file_plugin_api_proto_rawDescGZIP(), []int{1}
}

func (x *Release) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Release) GetGithubReleaseTag() string {
	if x != nil {
		return x.GithubReleaseTag
	}
	return ""
}

func (x *Release) GetPublishedAt() *timestamp.Timestamp {
	if x != nil {
		return x.PublishedAt
	}
	return nil
}

func (x *Release) GetIsDeprecated() bool {
	if x != nil {
		return x.IsDeprecated
	}
	return false
}

func (x *Release) GetDownloadURI() string {
	if x != nil {
		return x.DownloadURI
	}
	return ""
}

type PluginCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	RepoUri string `protobuf:"bytes,2,opt,name=repoUri,proto3" json:"repoUri,omitempty"`
}

func (x *PluginCreateRequest) Reset() {
	*x = PluginCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginCreateRequest) ProtoMessage() {}

func (x *PluginCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginCreateRequest.ProtoReflect.Descriptor instead.
func (*PluginCreateRequest) Descriptor() ([]byte, []int) {
	return file_plugin_api_proto_rawDescGZIP(), []int{2}
}

func (x *PluginCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PluginCreateRequest) GetRepoUri() string {
	if x != nil {
		return x.RepoUri
	}
	return ""
}

type Plugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	OwnerName string     `protobuf:"bytes,3,opt,name=ownerName,proto3" json:"ownerName,omitempty"`
	Source    *Source    `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Releases  []*Release `protobuf:"bytes,5,rep,name=releases,proto3" json:"releases,omitempty"`
}

func (x *Plugin) Reset() {
	*x = Plugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plugin) ProtoMessage() {}

func (x *Plugin) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plugin.ProtoReflect.Descriptor instead.
func (*Plugin) Descriptor() ([]byte, []int) {
	return file_plugin_api_proto_rawDescGZIP(), []int{3}
}

func (x *Plugin) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Plugin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Plugin) GetOwnerName() string {
	if x != nil {
		return x.OwnerName
	}
	return ""
}

func (x *Plugin) GetSource() *Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *Plugin) GetReleases() []*Release {
	if x != nil {
		return x.Releases
	}
	return nil
}

type PluginPublishReleaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PluginId         string `protobuf:"bytes,1,opt,name=pluginId,proto3" json:"pluginId,omitempty"`
	GithubReleaseTag string `protobuf:"bytes,2,opt,name=githubReleaseTag,proto3" json:"githubReleaseTag,omitempty"`
	Name             string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PluginPublishReleaseRequest) Reset() {
	*x = PluginPublishReleaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginPublishReleaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginPublishReleaseRequest) ProtoMessage() {}

func (x *PluginPublishReleaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginPublishReleaseRequest.ProtoReflect.Descriptor instead.
func (*PluginPublishReleaseRequest) Descriptor() ([]byte, []int) {
	return file_plugin_api_proto_rawDescGZIP(), []int{4}
}

func (x *PluginPublishReleaseRequest) GetPluginId() string {
	if x != nil {
		return x.PluginId
	}
	return ""
}

func (x *PluginPublishReleaseRequest) GetGithubReleaseTag() string {
	if x != nil {
		return x.GithubReleaseTag
	}
	return ""
}

func (x *PluginPublishReleaseRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PluginDeprecateReleaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PluginId    string `protobuf:"bytes,1,opt,name=pluginId,proto3" json:"pluginId,omitempty"`
	ReleaseName string `protobuf:"bytes,2,opt,name=releaseName,proto3" json:"releaseName,omitempty"`
}

func (x *PluginDeprecateReleaseRequest) Reset() {
	*x = PluginDeprecateReleaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginDeprecateReleaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginDeprecateReleaseRequest) ProtoMessage() {}

func (x *PluginDeprecateReleaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginDeprecateReleaseRequest.ProtoReflect.Descriptor instead.
func (*PluginDeprecateReleaseRequest) Descriptor() ([]byte, []int) {
	return file_plugin_api_proto_rawDescGZIP(), []int{5}
}

func (x *PluginDeprecateReleaseRequest) GetPluginId() string {
	if x != nil {
		return x.PluginId
	}
	return ""
}

func (x *PluginDeprecateReleaseRequest) GetReleaseName() string {
	if x != nil {
		return x.ReleaseName
	}
	return ""
}

type PluginGetReadmeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PluginId string `protobuf:"bytes,1,opt,name=pluginId,proto3" json:"pluginId,omitempty"`
}

func (x *PluginGetReadmeRequest) Reset() {
	*x = PluginGetReadmeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginGetReadmeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginGetReadmeRequest) ProtoMessage() {}

func (x *PluginGetReadmeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginGetReadmeRequest.ProtoReflect.Descriptor instead.
func (*PluginGetReadmeRequest) Descriptor() ([]byte, []int) {
	return file_plugin_api_proto_rawDescGZIP(), []int{6}
}

func (x *PluginGetReadmeRequest) GetPluginId() string {
	if x != nil {
		return x.PluginId
	}
	return ""
}

type PluginGetReadmeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Readme string `protobuf:"bytes,1,opt,name=readme,proto3" json:"readme,omitempty"`
}

func (x *PluginGetReadmeResponse) Reset() {
	*x = PluginGetReadmeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginGetReadmeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginGetReadmeResponse) ProtoMessage() {}

func (x *PluginGetReadmeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginGetReadmeResponse.ProtoReflect.Descriptor instead.
func (*PluginGetReadmeResponse) Descriptor() ([]byte, []int) {
	return file_plugin_api_proto_rawDescGZIP(), []int{7}
}

func (x *PluginGetReadmeResponse) GetReadme() string {
	if x != nil {
		return x.Readme
	}
	return ""
}

type PluginSearchByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plugins []*Plugin `protobuf:"bytes,1,rep,name=plugins,proto3" json:"plugins,omitempty"`
	Pages   int32     `protobuf:"varint,2,opt,name=pages,proto3" json:"pages,omitempty"`
	Page    int32     `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *PluginSearchByNameResponse) Reset() {
	*x = PluginSearchByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginSearchByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginSearchByNameResponse) ProtoMessage() {}

func (x *PluginSearchByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginSearchByNameResponse.ProtoReflect.Descriptor instead.
func (*PluginSearchByNameResponse) Descriptor() ([]byte, []int) {
	return file_plugin_api_proto_rawDescGZIP(), []int{8}
}

func (x *PluginSearchByNameResponse) GetPlugins() []*Plugin {
	if x != nil {
		return x.Plugins
	}
	return nil
}

func (x *PluginSearchByNameResponse) GetPages() int32 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *PluginSearchByNameResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type PluginSearchByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Page           int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PluginsPerPage int32  `protobuf:"varint,3,opt,name=pluginsPerPage,proto3" json:"pluginsPerPage,omitempty"`
	Limit          int32  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *PluginSearchByNameRequest) Reset() {
	*x = PluginSearchByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginSearchByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginSearchByNameRequest) ProtoMessage() {}

func (x *PluginSearchByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginSearchByNameRequest.ProtoReflect.Descriptor instead.
func (*PluginSearchByNameRequest) Descriptor() ([]byte, []int) {
	return file_plugin_api_proto_rawDescGZIP(), []int{9}
}

func (x *PluginSearchByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PluginSearchByNameRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PluginSearchByNameRequest) GetPluginsPerPage() int32 {
	if x != nil {
		return x.PluginsPerPage
	}
	return 0
}

func (x *PluginSearchByNameRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

var File_plugin_api_proto protoreflect.FileDescriptor

var file_plugin_api_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x70, 0x6f, 0x55, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x72, 0x65, 0x70, 0x6f, 0x55, 0x72, 0x69, 0x22, 0xcd, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x54, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x54, 0x61, 0x67, 0x12, 0x3c, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x44, 0x65, 0x70, 0x72, 0x65,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x55, 0x52, 0x49, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x55, 0x52, 0x49, 0x22, 0x43, 0x0a, 0x13, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x55, 0x72, 0x69, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x55, 0x72, 0x69, 0x22, 0x91, 0x01, 0x0a,
	0x06, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x08, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x73,
	0x22, 0x79, 0x0a, 0x1b, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x54, 0x61, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x54, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5d, 0x0a, 0x1d, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x44, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x16, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64,
	0x22, 0x31, 0x0a, 0x17, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61,
	0x64, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x61, 0x64, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61,
	0x64, 0x6d, 0x65, 0x22, 0x69, 0x0a, 0x1a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x21, 0x0a, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x07, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x81,
	0x01, 0x0a, 0x19, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x50,
	0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x50, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x67, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugin_api_proto_rawDescOnce sync.Once
	file_plugin_api_proto_rawDescData = file_plugin_api_proto_rawDesc
)

func file_plugin_api_proto_rawDescGZIP() []byte {
	file_plugin_api_proto_rawDescOnce.Do(func() {
		file_plugin_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugin_api_proto_rawDescData)
	})
	return file_plugin_api_proto_rawDescData
}

var file_plugin_api_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_plugin_api_proto_goTypes = []interface{}{
	(*Source)(nil),                        // 0: Source
	(*Release)(nil),                       // 1: Release
	(*PluginCreateRequest)(nil),           // 2: PluginCreateRequest
	(*Plugin)(nil),                        // 3: Plugin
	(*PluginPublishReleaseRequest)(nil),   // 4: PluginPublishReleaseRequest
	(*PluginDeprecateReleaseRequest)(nil), // 5: PluginDeprecateReleaseRequest
	(*PluginGetReadmeRequest)(nil),        // 6: PluginGetReadmeRequest
	(*PluginGetReadmeResponse)(nil),       // 7: PluginGetReadmeResponse
	(*PluginSearchByNameResponse)(nil),    // 8: PluginSearchByNameResponse
	(*PluginSearchByNameRequest)(nil),     // 9: PluginSearchByNameRequest
	(*timestamp.Timestamp)(nil),           // 10: google.protobuf.Timestamp
}
var file_plugin_api_proto_depIdxs = []int32{
	10, // 0: Release.publishedAt:type_name -> google.protobuf.Timestamp
	0,  // 1: Plugin.source:type_name -> Source
	1,  // 2: Plugin.releases:type_name -> Release
	3,  // 3: PluginSearchByNameResponse.plugins:type_name -> Plugin
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_plugin_api_proto_init() }
func file_plugin_api_proto_init() {
	if File_plugin_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugin_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Source); i {
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
		file_plugin_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Release); i {
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
		file_plugin_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginCreateRequest); i {
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
		file_plugin_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plugin); i {
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
		file_plugin_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginPublishReleaseRequest); i {
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
		file_plugin_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginDeprecateReleaseRequest); i {
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
		file_plugin_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginGetReadmeRequest); i {
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
		file_plugin_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginGetReadmeResponse); i {
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
		file_plugin_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginSearchByNameResponse); i {
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
		file_plugin_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginSearchByNameRequest); i {
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
			RawDescriptor: file_plugin_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugin_api_proto_goTypes,
		DependencyIndexes: file_plugin_api_proto_depIdxs,
		MessageInfos:      file_plugin_api_proto_msgTypes,
	}.Build()
	File_plugin_api_proto = out.File
	file_plugin_api_proto_rawDesc = nil
	file_plugin_api_proto_goTypes = nil
	file_plugin_api_proto_depIdxs = nil
}
