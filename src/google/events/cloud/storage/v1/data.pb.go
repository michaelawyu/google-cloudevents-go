// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: google/events/cloud/storage/v1/data.proto

package google_events_cloud_storage_v1

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// An object within Google Cloud Storage.
type StorageObjectData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Content-Encoding of the object data, matching
	// [https://tools.ietf.org/html/rfc7231#section-3.1.2.2][RFC 7231 §3.1.2.2]
	ContentEncoding string `protobuf:"bytes,1,opt,name=content_encoding,json=contentEncoding,proto3" json:"content_encoding,omitempty"`
	// Content-Disposition of the object data, matching
	// [https://tools.ietf.org/html/rfc6266][RFC 6266].
	ContentDisposition string `protobuf:"bytes,2,opt,name=content_disposition,json=contentDisposition,proto3" json:"content_disposition,omitempty"`
	// Cache-Control directive for the object data, matching
	// [https://tools.ietf.org/html/rfc7234#section-5.2"][RFC 7234 §5.2].
	CacheControl string `protobuf:"bytes,3,opt,name=cache_control,json=cacheControl,proto3" json:"cache_control,omitempty"`
	// Content-Language of the object data, matching
	// [https://tools.ietf.org/html/rfc7231#section-3.1.3.2][RFC 7231 §3.1.3.2].
	ContentLanguage string `protobuf:"bytes,5,opt,name=content_language,json=contentLanguage,proto3" json:"content_language,omitempty"`
	// The version of the metadata for this object at this generation. Used for
	// preconditions and for detecting changes in metadata. A metageneration
	// number is only meaningful in the context of a particular generation of a
	// particular object.
	Metageneration int64 `protobuf:"varint,6,opt,name=metageneration,proto3" json:"metageneration,omitempty"`
	// The deletion time of the object. Will be returned if and only if this
	// version of the object has been deleted.
	TimeDeleted *timestamp.Timestamp `protobuf:"bytes,7,opt,name=time_deleted,json=timeDeleted,proto3" json:"time_deleted,omitempty"`
	// Content-Type of the object data, matching
	// [https://tools.ietf.org/html/rfc7231#section-3.1.1.5][RFC 7231 §3.1.1.5].
	// If an object is stored without a Content-Type, it is served as
	// `application/octet-stream`.
	ContentType string `protobuf:"bytes,8,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// Content-Length of the object data in bytes, matching
	// [https://tools.ietf.org/html/rfc7230#section-3.3.2][RFC 7230 §3.3.2].
	Size int64 `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty"`
	// The creation time of the object.
	// Attempting to set this field will result in an error.
	TimeCreated *timestamp.Timestamp `protobuf:"bytes,10,opt,name=time_created,json=timeCreated,proto3" json:"time_created,omitempty"`
	// CRC32c checksum. For more information about using the CRC32c
	// checksum, see
	// [https://cloud.google.com/storage/docs/hashes-etags#_JSONAPI][Hashes and
	// ETags: Best Practices].
	Crc32C string `protobuf:"bytes,11,opt,name=crc32c,proto3" json:"crc32c,omitempty"`
	// Number of underlying components that make up this object. Components are
	// accumulated by compose operations.
	// Attempting to set this field will result in an error.
	ComponentCount int32 `protobuf:"varint,12,opt,name=component_count,json=componentCount,proto3" json:"component_count,omitempty"`
	// MD5 hash of the data; encoded using base64 as per
	// [https://tools.ietf.org/html/rfc4648#section-4][RFC 4648 §4]. For more
	// information about using the MD5 hash, see
	// [https://cloud.google.com/storage/docs/hashes-etags#_JSONAPI][Hashes and
	// ETags: Best Practices].
	Md5Hash string `protobuf:"bytes,13,opt,name=md5_hash,json=md5Hash,proto3" json:"md5_hash,omitempty"`
	// HTTP 1.1 Entity tag for the object. See
	// [https://tools.ietf.org/html/rfc7232#section-2.3][RFC 7232 §2.3].
	Etag string `protobuf:"bytes,14,opt,name=etag,proto3" json:"etag,omitempty"`
	// The modification time of the object metadata.
	Updated *timestamp.Timestamp `protobuf:"bytes,15,opt,name=updated,proto3" json:"updated,omitempty"`
	// Storage class of the object.
	StorageClass string `protobuf:"bytes,16,opt,name=storage_class,json=storageClass,proto3" json:"storage_class,omitempty"`
	// Cloud KMS Key used to encrypt this object, if the object is encrypted by
	// such a key.
	KmsKeyName string `protobuf:"bytes,17,opt,name=kms_key_name,json=kmsKeyName,proto3" json:"kms_key_name,omitempty"`
	// The time at which the object's storage class was last changed.
	TimeStorageClassUpdated *timestamp.Timestamp `protobuf:"bytes,18,opt,name=time_storage_class_updated,json=timeStorageClassUpdated,proto3" json:"time_storage_class_updated,omitempty"`
	// Whether an object is under temporary hold.
	TemporaryHold bool `protobuf:"varint,19,opt,name=temporary_hold,json=temporaryHold,proto3" json:"temporary_hold,omitempty"`
	// A server-determined value that specifies the earliest time that the
	// object's retention period expires.
	RetentionExpirationTime *timestamp.Timestamp `protobuf:"bytes,20,opt,name=retention_expiration_time,json=retentionExpirationTime,proto3" json:"retention_expiration_time,omitempty"`
	// User-provided metadata, in key/value pairs.
	Metadata map[string]string `protobuf:"bytes,21,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Whether an object is under event-based hold.
	EventBasedHold bool `protobuf:"varint,29,opt,name=event_based_hold,json=eventBasedHold,proto3" json:"event_based_hold,omitempty"`
	// The name of the object.
	Name string `protobuf:"bytes,23,opt,name=name,proto3" json:"name,omitempty"`
	// The ID of the object, including the bucket name, object name, and
	// generation number.
	Id string `protobuf:"bytes,24,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the bucket containing this object.
	Bucket string `protobuf:"bytes,25,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// The content generation of this object. Used for object versioning.
	// Attempting to set this field will result in an error.
	Generation int64 `protobuf:"varint,26,opt,name=generation,proto3" json:"generation,omitempty"`
	// Metadata of customer-supplied encryption key, if the object is encrypted by
	// such a key.
	CustomerEncryption *StorageObjectData_CustomerEncryption `protobuf:"bytes,28,opt,name=customer_encryption,json=customerEncryption,proto3" json:"customer_encryption,omitempty"`
	// The link to this object.
	MediaLink string `protobuf:"bytes,100,opt,name=media_link,json=mediaLink,proto3" json:"media_link,omitempty"`
	// Media download link.
	SelfLink string `protobuf:"bytes,101,opt,name=self_link,json=selfLink,proto3" json:"self_link,omitempty"`
}

func (x *StorageObjectData) Reset() {
	*x = StorageObjectData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_events_cloud_storage_v1_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageObjectData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageObjectData) ProtoMessage() {}

func (x *StorageObjectData) ProtoReflect() protoreflect.Message {
	mi := &file_google_events_cloud_storage_v1_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageObjectData.ProtoReflect.Descriptor instead.
func (*StorageObjectData) Descriptor() ([]byte, []int) {
	return file_google_events_cloud_storage_v1_data_proto_rawDescGZIP(), []int{0}
}

func (x *StorageObjectData) GetContentEncoding() string {
	if x != nil {
		return x.ContentEncoding
	}
	return ""
}

func (x *StorageObjectData) GetContentDisposition() string {
	if x != nil {
		return x.ContentDisposition
	}
	return ""
}

func (x *StorageObjectData) GetCacheControl() string {
	if x != nil {
		return x.CacheControl
	}
	return ""
}

func (x *StorageObjectData) GetContentLanguage() string {
	if x != nil {
		return x.ContentLanguage
	}
	return ""
}

func (x *StorageObjectData) GetMetageneration() int64 {
	if x != nil {
		return x.Metageneration
	}
	return 0
}

func (x *StorageObjectData) GetTimeDeleted() *timestamp.Timestamp {
	if x != nil {
		return x.TimeDeleted
	}
	return nil
}

func (x *StorageObjectData) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *StorageObjectData) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *StorageObjectData) GetTimeCreated() *timestamp.Timestamp {
	if x != nil {
		return x.TimeCreated
	}
	return nil
}

func (x *StorageObjectData) GetCrc32C() string {
	if x != nil {
		return x.Crc32C
	}
	return ""
}

func (x *StorageObjectData) GetComponentCount() int32 {
	if x != nil {
		return x.ComponentCount
	}
	return 0
}

func (x *StorageObjectData) GetMd5Hash() string {
	if x != nil {
		return x.Md5Hash
	}
	return ""
}

func (x *StorageObjectData) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *StorageObjectData) GetUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (x *StorageObjectData) GetStorageClass() string {
	if x != nil {
		return x.StorageClass
	}
	return ""
}

func (x *StorageObjectData) GetKmsKeyName() string {
	if x != nil {
		return x.KmsKeyName
	}
	return ""
}

func (x *StorageObjectData) GetTimeStorageClassUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.TimeStorageClassUpdated
	}
	return nil
}

func (x *StorageObjectData) GetTemporaryHold() bool {
	if x != nil {
		return x.TemporaryHold
	}
	return false
}

func (x *StorageObjectData) GetRetentionExpirationTime() *timestamp.Timestamp {
	if x != nil {
		return x.RetentionExpirationTime
	}
	return nil
}

func (x *StorageObjectData) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *StorageObjectData) GetEventBasedHold() bool {
	if x != nil {
		return x.EventBasedHold
	}
	return false
}

func (x *StorageObjectData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StorageObjectData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StorageObjectData) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *StorageObjectData) GetGeneration() int64 {
	if x != nil {
		return x.Generation
	}
	return 0
}

func (x *StorageObjectData) GetCustomerEncryption() *StorageObjectData_CustomerEncryption {
	if x != nil {
		return x.CustomerEncryption
	}
	return nil
}

func (x *StorageObjectData) GetMediaLink() string {
	if x != nil {
		return x.MediaLink
	}
	return ""
}

func (x *StorageObjectData) GetSelfLink() string {
	if x != nil {
		return x.SelfLink
	}
	return ""
}

// Describes the customer-specified mechanism used to store the data at rest.
type StorageObjectData_CustomerEncryption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The encryption algorithm.
	EncryptionAlgorithm string `protobuf:"bytes,1,opt,name=encryption_algorithm,json=encryptionAlgorithm,proto3" json:"encryption_algorithm,omitempty"`
	// SHA256 hash value of the encryption key.
	KeySha256 string `protobuf:"bytes,2,opt,name=key_sha256,json=keySha256,proto3" json:"key_sha256,omitempty"`
}

func (x *StorageObjectData_CustomerEncryption) Reset() {
	*x = StorageObjectData_CustomerEncryption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_events_cloud_storage_v1_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageObjectData_CustomerEncryption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageObjectData_CustomerEncryption) ProtoMessage() {}

func (x *StorageObjectData_CustomerEncryption) ProtoReflect() protoreflect.Message {
	mi := &file_google_events_cloud_storage_v1_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageObjectData_CustomerEncryption.ProtoReflect.Descriptor instead.
func (*StorageObjectData_CustomerEncryption) Descriptor() ([]byte, []int) {
	return file_google_events_cloud_storage_v1_data_proto_rawDescGZIP(), []int{0, 0}
}

func (x *StorageObjectData_CustomerEncryption) GetEncryptionAlgorithm() string {
	if x != nil {
		return x.EncryptionAlgorithm
	}
	return ""
}

func (x *StorageObjectData_CustomerEncryption) GetKeySha256() string {
	if x != nil {
		return x.KeySha256
	}
	return ""
}

var File_google_events_cloud_storage_v1_data_proto protoreflect.FileDescriptor

var file_google_events_cloud_storage_v1_data_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x0b, 0x0a,
	0x11, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x6e,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x2f, 0x0a,
	0x13, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x26,
	0x0a, 0x0e, 0x6d, 0x65, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6d, 0x65, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x3d, 0x0a, 0x0c,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b,
	0x74, 0x69, 0x6d, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x72, 0x63, 0x33, 0x32, 0x63, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x72, 0x63,
	0x33, 0x32, 0x63, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x6d, 0x64, 0x35, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x64, 0x35, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x12, 0x34, 0x0a, 0x07, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x6b, 0x6d, 0x73, 0x5f, 0x6b, 0x65,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6b, 0x6d,
	0x73, 0x4b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x57, 0x0a, 0x1a, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x17, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x68,
	0x6f, 0x6c, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x72, 0x79, 0x48, 0x6f, 0x6c, 0x64, 0x12, 0x56, 0x0a, 0x19, 0x72, 0x65, 0x74, 0x65,
	0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x17, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x5b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x15, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a,
	0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x64, 0x5f, 0x68, 0x6f, 0x6c,
	0x64, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x61,
	0x73, 0x65, 0x64, 0x48, 0x6f, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x75, 0x0a, 0x13, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x44, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6c,
	0x66, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65,
	0x6c, 0x66, 0x4c, 0x69, 0x6e, 0x6b, 0x1a, 0x66, 0x0a, 0x12, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x14,
	0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x65, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12,
	0x1d, 0x0a, 0x0a, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6b, 0x65, 0x79, 0x53, 0x68, 0x61, 0x32, 0x35, 0x36, 0x1a, 0x3b,
	0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2a, 0xaa, 0x02, 0x27,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_events_cloud_storage_v1_data_proto_rawDescOnce sync.Once
	file_google_events_cloud_storage_v1_data_proto_rawDescData = file_google_events_cloud_storage_v1_data_proto_rawDesc
)

func file_google_events_cloud_storage_v1_data_proto_rawDescGZIP() []byte {
	file_google_events_cloud_storage_v1_data_proto_rawDescOnce.Do(func() {
		file_google_events_cloud_storage_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_events_cloud_storage_v1_data_proto_rawDescData)
	})
	return file_google_events_cloud_storage_v1_data_proto_rawDescData
}

var file_google_events_cloud_storage_v1_data_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_events_cloud_storage_v1_data_proto_goTypes = []interface{}{
	(*StorageObjectData)(nil),                    // 0: google.events.cloud.storage.v1.StorageObjectData
	(*StorageObjectData_CustomerEncryption)(nil), // 1: google.events.cloud.storage.v1.StorageObjectData.CustomerEncryption
	nil,                         // 2: google.events.cloud.storage.v1.StorageObjectData.MetadataEntry
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_google_events_cloud_storage_v1_data_proto_depIdxs = []int32{
	3, // 0: google.events.cloud.storage.v1.StorageObjectData.time_deleted:type_name -> google.protobuf.Timestamp
	3, // 1: google.events.cloud.storage.v1.StorageObjectData.time_created:type_name -> google.protobuf.Timestamp
	3, // 2: google.events.cloud.storage.v1.StorageObjectData.updated:type_name -> google.protobuf.Timestamp
	3, // 3: google.events.cloud.storage.v1.StorageObjectData.time_storage_class_updated:type_name -> google.protobuf.Timestamp
	3, // 4: google.events.cloud.storage.v1.StorageObjectData.retention_expiration_time:type_name -> google.protobuf.Timestamp
	2, // 5: google.events.cloud.storage.v1.StorageObjectData.metadata:type_name -> google.events.cloud.storage.v1.StorageObjectData.MetadataEntry
	1, // 6: google.events.cloud.storage.v1.StorageObjectData.customer_encryption:type_name -> google.events.cloud.storage.v1.StorageObjectData.CustomerEncryption
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_events_cloud_storage_v1_data_proto_init() }
func file_google_events_cloud_storage_v1_data_proto_init() {
	if File_google_events_cloud_storage_v1_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_events_cloud_storage_v1_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageObjectData); i {
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
		file_google_events_cloud_storage_v1_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageObjectData_CustomerEncryption); i {
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
			RawDescriptor: file_google_events_cloud_storage_v1_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_events_cloud_storage_v1_data_proto_goTypes,
		DependencyIndexes: file_google_events_cloud_storage_v1_data_proto_depIdxs,
		MessageInfos:      file_google_events_cloud_storage_v1_data_proto_msgTypes,
	}.Build()
	File_google_events_cloud_storage_v1_data_proto = out.File
	file_google_events_cloud_storage_v1_data_proto_rawDesc = nil
	file_google_events_cloud_storage_v1_data_proto_goTypes = nil
	file_google_events_cloud_storage_v1_data_proto_depIdxs = nil
}
