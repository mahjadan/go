// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: proto/tasks.proto

package grpc

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

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tasks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tasks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_proto_tasks_proto_rawDescGZIP(), []int{0}
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Done bool   `protobuf:"varint,2,opt,name=done,proto3" json:"done,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tasks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tasks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_proto_tasks_proto_rawDescGZIP(), []int{1}
}

func (x *Task) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Task) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type TaskList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *TaskList) Reset() {
	*x = TaskList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tasks_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskList) ProtoMessage() {}

func (x *TaskList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tasks_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskList.ProtoReflect.Descriptor instead.
func (*TaskList) Descriptor() ([]byte, []int) {
	return file_proto_tasks_proto_rawDescGZIP(), []int{2}
}

func (x *TaskList) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type NewTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NewTask) Reset() {
	*x = NewTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tasks_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTask) ProtoMessage() {}

func (x *NewTask) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tasks_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTask.ProtoReflect.Descriptor instead.
func (*NewTask) Descriptor() ([]byte, []int) {
	return file_proto_tasks_proto_rawDescGZIP(), []int{3}
}

func (x *NewTask) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proto_tasks_proto protoreflect.FileDescriptor

var file_proto_tasks_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69,
	0x64, 0x22, 0x2e, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e,
	0x65, 0x22, 0x2c, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22,
	0x1d, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x51,
	0x0a, 0x05, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x0a, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x0e, 0x2e, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x22, 0x0a,
	0x03, 0x41, 0x64, 0x64, 0x12, 0x0d, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x4e, 0x65, 0x77, 0x54,
	0x61, 0x73, 0x6b, 0x1a, 0x0a, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x22,
	0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x61, 0x68, 0x6a, 0x61, 0x64, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_tasks_proto_rawDescOnce sync.Once
	file_proto_tasks_proto_rawDescData = file_proto_tasks_proto_rawDesc
)

func file_proto_tasks_proto_rawDescGZIP() []byte {
	file_proto_tasks_proto_rawDescOnce.Do(func() {
		file_proto_tasks_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_tasks_proto_rawDescData)
	})
	return file_proto_tasks_proto_rawDescData
}

var file_proto_tasks_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_tasks_proto_goTypes = []interface{}{
	(*Void)(nil),     // 0: list.Void
	(*Task)(nil),     // 1: list.Task
	(*TaskList)(nil), // 2: list.TaskList
	(*NewTask)(nil),  // 3: list.NewTask
}
var file_proto_tasks_proto_depIdxs = []int32{
	1, // 0: list.TaskList.tasks:type_name -> list.Task
	0, // 1: list.Tasks.List:input_type -> list.Void
	3, // 2: list.Tasks.Add:input_type -> list.NewTask
	2, // 3: list.Tasks.List:output_type -> list.TaskList
	1, // 4: list.Tasks.Add:output_type -> list.Task
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_tasks_proto_init() }
func file_proto_tasks_proto_init() {
	if File_proto_tasks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_tasks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_proto_tasks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_proto_tasks_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskList); i {
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
		file_proto_tasks_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewTask); i {
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
			RawDescriptor: file_proto_tasks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_tasks_proto_goTypes,
		DependencyIndexes: file_proto_tasks_proto_depIdxs,
		MessageInfos:      file_proto_tasks_proto_msgTypes,
	}.Build()
	File_proto_tasks_proto = out.File
	file_proto_tasks_proto_rawDesc = nil
	file_proto_tasks_proto_goTypes = nil
	file_proto_tasks_proto_depIdxs = nil
}
