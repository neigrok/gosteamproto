// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: cs_usercmd.proto

package csgo

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

type CSGOInterpolationInfoPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrcTick *int32   `protobuf:"varint,1,opt,name=src_tick,json=srcTick,def=-1" json:"src_tick,omitempty"`
	DstTick *int32   `protobuf:"varint,2,opt,name=dst_tick,json=dstTick,def=-1" json:"dst_tick,omitempty"`
	Frac    *float32 `protobuf:"fixed32,3,opt,name=frac,def=0" json:"frac,omitempty"`
}

// Default values for CSGOInterpolationInfoPB fields.
const (
	Default_CSGOInterpolationInfoPB_SrcTick = int32(-1)
	Default_CSGOInterpolationInfoPB_DstTick = int32(-1)
	Default_CSGOInterpolationInfoPB_Frac    = float32(0)
)

func (x *CSGOInterpolationInfoPB) Reset() {
	*x = CSGOInterpolationInfoPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs_usercmd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSGOInterpolationInfoPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSGOInterpolationInfoPB) ProtoMessage() {}

func (x *CSGOInterpolationInfoPB) ProtoReflect() protoreflect.Message {
	mi := &file_cs_usercmd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSGOInterpolationInfoPB.ProtoReflect.Descriptor instead.
func (*CSGOInterpolationInfoPB) Descriptor() ([]byte, []int) {
	return file_cs_usercmd_proto_rawDescGZIP(), []int{0}
}

func (x *CSGOInterpolationInfoPB) GetSrcTick() int32 {
	if x != nil && x.SrcTick != nil {
		return *x.SrcTick
	}
	return Default_CSGOInterpolationInfoPB_SrcTick
}

func (x *CSGOInterpolationInfoPB) GetDstTick() int32 {
	if x != nil && x.DstTick != nil {
		return *x.DstTick
	}
	return Default_CSGOInterpolationInfoPB_DstTick
}

func (x *CSGOInterpolationInfoPB) GetFrac() float32 {
	if x != nil && x.Frac != nil {
		return *x.Frac
	}
	return Default_CSGOInterpolationInfoPB_Frac
}

type CSGOInterpolationInfoPB_CL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Frac *float32 `protobuf:"fixed32,3,opt,name=frac,def=0" json:"frac,omitempty"`
}

// Default values for CSGOInterpolationInfoPB_CL fields.
const (
	Default_CSGOInterpolationInfoPB_CL_Frac = float32(0)
)

func (x *CSGOInterpolationInfoPB_CL) Reset() {
	*x = CSGOInterpolationInfoPB_CL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs_usercmd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSGOInterpolationInfoPB_CL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSGOInterpolationInfoPB_CL) ProtoMessage() {}

func (x *CSGOInterpolationInfoPB_CL) ProtoReflect() protoreflect.Message {
	mi := &file_cs_usercmd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSGOInterpolationInfoPB_CL.ProtoReflect.Descriptor instead.
func (*CSGOInterpolationInfoPB_CL) Descriptor() ([]byte, []int) {
	return file_cs_usercmd_proto_rawDescGZIP(), []int{1}
}

func (x *CSGOInterpolationInfoPB_CL) GetFrac() float32 {
	if x != nil && x.Frac != nil {
		return *x.Frac
	}
	return Default_CSGOInterpolationInfoPB_CL_Frac
}

type CSGOInputHistoryEntryPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ViewAngles         *CMsgQAngle                 `protobuf:"bytes,2,opt,name=view_angles,json=viewAngles" json:"view_angles,omitempty"`
	RenderTickCount    *int32                      `protobuf:"varint,4,opt,name=render_tick_count,json=renderTickCount" json:"render_tick_count,omitempty"`
	RenderTickFraction *float32                    `protobuf:"fixed32,5,opt,name=render_tick_fraction,json=renderTickFraction" json:"render_tick_fraction,omitempty"`
	PlayerTickCount    *int32                      `protobuf:"varint,6,opt,name=player_tick_count,json=playerTickCount" json:"player_tick_count,omitempty"`
	PlayerTickFraction *float32                    `protobuf:"fixed32,7,opt,name=player_tick_fraction,json=playerTickFraction" json:"player_tick_fraction,omitempty"`
	ClInterp           *CSGOInterpolationInfoPB_CL `protobuf:"bytes,12,opt,name=cl_interp,json=clInterp" json:"cl_interp,omitempty"`
	SvInterp0          *CSGOInterpolationInfoPB    `protobuf:"bytes,13,opt,name=sv_interp0,json=svInterp0" json:"sv_interp0,omitempty"`
	SvInterp1          *CSGOInterpolationInfoPB    `protobuf:"bytes,14,opt,name=sv_interp1,json=svInterp1" json:"sv_interp1,omitempty"`
	PlayerInterp       *CSGOInterpolationInfoPB    `protobuf:"bytes,15,opt,name=player_interp,json=playerInterp" json:"player_interp,omitempty"`
	FrameNumber        *int32                      `protobuf:"varint,64,opt,name=frame_number,json=frameNumber" json:"frame_number,omitempty"`
	TargetEntIndex     *int32                      `protobuf:"varint,65,opt,name=target_ent_index,json=targetEntIndex,def=-1" json:"target_ent_index,omitempty"`
	ShootPosition      *CMsgVector                 `protobuf:"bytes,66,opt,name=shoot_position,json=shootPosition" json:"shoot_position,omitempty"`
	TargetHeadPosCheck *CMsgVector                 `protobuf:"bytes,67,opt,name=target_head_pos_check,json=targetHeadPosCheck" json:"target_head_pos_check,omitempty"`
	TargetAbsPosCheck  *CMsgVector                 `protobuf:"bytes,68,opt,name=target_abs_pos_check,json=targetAbsPosCheck" json:"target_abs_pos_check,omitempty"`
	TargetAbsAngCheck  *CMsgQAngle                 `protobuf:"bytes,69,opt,name=target_abs_ang_check,json=targetAbsAngCheck" json:"target_abs_ang_check,omitempty"`
}

// Default values for CSGOInputHistoryEntryPB fields.
const (
	Default_CSGOInputHistoryEntryPB_TargetEntIndex = int32(-1)
)

func (x *CSGOInputHistoryEntryPB) Reset() {
	*x = CSGOInputHistoryEntryPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs_usercmd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSGOInputHistoryEntryPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSGOInputHistoryEntryPB) ProtoMessage() {}

func (x *CSGOInputHistoryEntryPB) ProtoReflect() protoreflect.Message {
	mi := &file_cs_usercmd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSGOInputHistoryEntryPB.ProtoReflect.Descriptor instead.
func (*CSGOInputHistoryEntryPB) Descriptor() ([]byte, []int) {
	return file_cs_usercmd_proto_rawDescGZIP(), []int{2}
}

func (x *CSGOInputHistoryEntryPB) GetViewAngles() *CMsgQAngle {
	if x != nil {
		return x.ViewAngles
	}
	return nil
}

func (x *CSGOInputHistoryEntryPB) GetRenderTickCount() int32 {
	if x != nil && x.RenderTickCount != nil {
		return *x.RenderTickCount
	}
	return 0
}

func (x *CSGOInputHistoryEntryPB) GetRenderTickFraction() float32 {
	if x != nil && x.RenderTickFraction != nil {
		return *x.RenderTickFraction
	}
	return 0
}

func (x *CSGOInputHistoryEntryPB) GetPlayerTickCount() int32 {
	if x != nil && x.PlayerTickCount != nil {
		return *x.PlayerTickCount
	}
	return 0
}

func (x *CSGOInputHistoryEntryPB) GetPlayerTickFraction() float32 {
	if x != nil && x.PlayerTickFraction != nil {
		return *x.PlayerTickFraction
	}
	return 0
}

func (x *CSGOInputHistoryEntryPB) GetClInterp() *CSGOInterpolationInfoPB_CL {
	if x != nil {
		return x.ClInterp
	}
	return nil
}

func (x *CSGOInputHistoryEntryPB) GetSvInterp0() *CSGOInterpolationInfoPB {
	if x != nil {
		return x.SvInterp0
	}
	return nil
}

func (x *CSGOInputHistoryEntryPB) GetSvInterp1() *CSGOInterpolationInfoPB {
	if x != nil {
		return x.SvInterp1
	}
	return nil
}

func (x *CSGOInputHistoryEntryPB) GetPlayerInterp() *CSGOInterpolationInfoPB {
	if x != nil {
		return x.PlayerInterp
	}
	return nil
}

func (x *CSGOInputHistoryEntryPB) GetFrameNumber() int32 {
	if x != nil && x.FrameNumber != nil {
		return *x.FrameNumber
	}
	return 0
}

func (x *CSGOInputHistoryEntryPB) GetTargetEntIndex() int32 {
	if x != nil && x.TargetEntIndex != nil {
		return *x.TargetEntIndex
	}
	return Default_CSGOInputHistoryEntryPB_TargetEntIndex
}

func (x *CSGOInputHistoryEntryPB) GetShootPosition() *CMsgVector {
	if x != nil {
		return x.ShootPosition
	}
	return nil
}

func (x *CSGOInputHistoryEntryPB) GetTargetHeadPosCheck() *CMsgVector {
	if x != nil {
		return x.TargetHeadPosCheck
	}
	return nil
}

func (x *CSGOInputHistoryEntryPB) GetTargetAbsPosCheck() *CMsgVector {
	if x != nil {
		return x.TargetAbsPosCheck
	}
	return nil
}

func (x *CSGOInputHistoryEntryPB) GetTargetAbsAngCheck() *CMsgQAngle {
	if x != nil {
		return x.TargetAbsAngCheck
	}
	return nil
}

type CSGOUserCmdPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base                     *CBaseUserCmdPB            `protobuf:"bytes,1,opt,name=base" json:"base,omitempty"`
	InputHistory             []*CSGOInputHistoryEntryPB `protobuf:"bytes,2,rep,name=input_history,json=inputHistory" json:"input_history,omitempty"`
	Attack1StartHistoryIndex *int32                     `protobuf:"varint,6,opt,name=attack1_start_history_index,json=attack1StartHistoryIndex,def=-1" json:"attack1_start_history_index,omitempty"`
	Attack2StartHistoryIndex *int32                     `protobuf:"varint,7,opt,name=attack2_start_history_index,json=attack2StartHistoryIndex,def=-1" json:"attack2_start_history_index,omitempty"`
	Attack3StartHistoryIndex *int32                     `protobuf:"varint,8,opt,name=attack3_start_history_index,json=attack3StartHistoryIndex,def=-1" json:"attack3_start_history_index,omitempty"`
	LeftHandDesired          *bool                      `protobuf:"varint,9,opt,name=left_hand_desired,json=leftHandDesired,def=0" json:"left_hand_desired,omitempty"`
	IsPredictingBodyShotFx   *bool                      `protobuf:"varint,11,opt,name=is_predicting_body_shot_fx,json=isPredictingBodyShotFx,def=0" json:"is_predicting_body_shot_fx,omitempty"`
	IsPredictingHeadShotFx   *bool                      `protobuf:"varint,12,opt,name=is_predicting_head_shot_fx,json=isPredictingHeadShotFx,def=0" json:"is_predicting_head_shot_fx,omitempty"`
	IsPredictingKillRagdolls *bool                      `protobuf:"varint,13,opt,name=is_predicting_kill_ragdolls,json=isPredictingKillRagdolls,def=0" json:"is_predicting_kill_ragdolls,omitempty"`
}

// Default values for CSGOUserCmdPB fields.
const (
	Default_CSGOUserCmdPB_Attack1StartHistoryIndex = int32(-1)
	Default_CSGOUserCmdPB_Attack2StartHistoryIndex = int32(-1)
	Default_CSGOUserCmdPB_Attack3StartHistoryIndex = int32(-1)
	Default_CSGOUserCmdPB_LeftHandDesired          = bool(false)
	Default_CSGOUserCmdPB_IsPredictingBodyShotFx   = bool(false)
	Default_CSGOUserCmdPB_IsPredictingHeadShotFx   = bool(false)
	Default_CSGOUserCmdPB_IsPredictingKillRagdolls = bool(false)
)

func (x *CSGOUserCmdPB) Reset() {
	*x = CSGOUserCmdPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cs_usercmd_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSGOUserCmdPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSGOUserCmdPB) ProtoMessage() {}

func (x *CSGOUserCmdPB) ProtoReflect() protoreflect.Message {
	mi := &file_cs_usercmd_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSGOUserCmdPB.ProtoReflect.Descriptor instead.
func (*CSGOUserCmdPB) Descriptor() ([]byte, []int) {
	return file_cs_usercmd_proto_rawDescGZIP(), []int{3}
}

func (x *CSGOUserCmdPB) GetBase() *CBaseUserCmdPB {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *CSGOUserCmdPB) GetInputHistory() []*CSGOInputHistoryEntryPB {
	if x != nil {
		return x.InputHistory
	}
	return nil
}

func (x *CSGOUserCmdPB) GetAttack1StartHistoryIndex() int32 {
	if x != nil && x.Attack1StartHistoryIndex != nil {
		return *x.Attack1StartHistoryIndex
	}
	return Default_CSGOUserCmdPB_Attack1StartHistoryIndex
}

func (x *CSGOUserCmdPB) GetAttack2StartHistoryIndex() int32 {
	if x != nil && x.Attack2StartHistoryIndex != nil {
		return *x.Attack2StartHistoryIndex
	}
	return Default_CSGOUserCmdPB_Attack2StartHistoryIndex
}

func (x *CSGOUserCmdPB) GetAttack3StartHistoryIndex() int32 {
	if x != nil && x.Attack3StartHistoryIndex != nil {
		return *x.Attack3StartHistoryIndex
	}
	return Default_CSGOUserCmdPB_Attack3StartHistoryIndex
}

func (x *CSGOUserCmdPB) GetLeftHandDesired() bool {
	if x != nil && x.LeftHandDesired != nil {
		return *x.LeftHandDesired
	}
	return Default_CSGOUserCmdPB_LeftHandDesired
}

func (x *CSGOUserCmdPB) GetIsPredictingBodyShotFx() bool {
	if x != nil && x.IsPredictingBodyShotFx != nil {
		return *x.IsPredictingBodyShotFx
	}
	return Default_CSGOUserCmdPB_IsPredictingBodyShotFx
}

func (x *CSGOUserCmdPB) GetIsPredictingHeadShotFx() bool {
	if x != nil && x.IsPredictingHeadShotFx != nil {
		return *x.IsPredictingHeadShotFx
	}
	return Default_CSGOUserCmdPB_IsPredictingHeadShotFx
}

func (x *CSGOUserCmdPB) GetIsPredictingKillRagdolls() bool {
	if x != nil && x.IsPredictingKillRagdolls != nil {
		return *x.IsPredictingKillRagdolls
	}
	return Default_CSGOUserCmdPB_IsPredictingKillRagdolls
}

var File_cs_usercmd_proto protoreflect.FileDescriptor

var file_cs_usercmd_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6d, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x75, 0x73, 0x65, 0x72,
	0x63, 0x6d, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x17, 0x43, 0x53, 0x47,
	0x4f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x50, 0x42, 0x12, 0x1d, 0x0a, 0x08, 0x73, 0x72, 0x63, 0x5f, 0x74, 0x69, 0x63, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x02, 0x2d, 0x31, 0x52, 0x07, 0x73, 0x72, 0x63, 0x54,
	0x69, 0x63, 0x6b, 0x12, 0x1d, 0x0a, 0x08, 0x64, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x02, 0x2d, 0x31, 0x52, 0x07, 0x64, 0x73, 0x74, 0x54, 0x69,
	0x63, 0x6b, 0x12, 0x15, 0x0a, 0x04, 0x66, 0x72, 0x61, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x3a, 0x01, 0x30, 0x52, 0x04, 0x66, 0x72, 0x61, 0x63, 0x22, 0x33, 0x0a, 0x1a, 0x43, 0x53, 0x47,
	0x4f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x50, 0x42, 0x5f, 0x43, 0x4c, 0x12, 0x15, 0x0a, 0x04, 0x66, 0x72, 0x61, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x01, 0x30, 0x52, 0x04, 0x66, 0x72, 0x61, 0x63, 0x22, 0xaf,
	0x06, 0x0a, 0x17, 0x43, 0x53, 0x47, 0x4f, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x42, 0x12, 0x2c, 0x0a, 0x0b, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x51, 0x41, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x0a, 0x76, 0x69,
	0x65, 0x77, 0x41, 0x6e, 0x67, 0x6c, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0f, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x74,
	0x69, 0x63, 0x6b, 0x5f, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x12, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x46, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x74, 0x69, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x63,
	0x6b, 0x5f, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x12, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x46, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x70, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x43, 0x53, 0x47, 0x4f, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x50,
	0x42, 0x5f, 0x43, 0x4c, 0x52, 0x08, 0x63, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x12, 0x37,
	0x0a, 0x0a, 0x73, 0x76, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x30, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x43, 0x53, 0x47, 0x4f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x6f,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x42, 0x52, 0x09, 0x73, 0x76,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x30, 0x12, 0x37, 0x0a, 0x0a, 0x73, 0x76, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x31, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x43, 0x53,
	0x47, 0x4f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x50, 0x42, 0x52, 0x09, 0x73, 0x76, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x31,
	0x12, 0x3d, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x70, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x43, 0x53, 0x47, 0x4f, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x50,
	0x42, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x12,
	0x21, 0x0a, 0x0c, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x40, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x2c, 0x0a, 0x10, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x41, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x02, 0x2d, 0x31,
	0x52, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x32, 0x0a, 0x0e, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x42, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0d, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x15, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x68,
	0x65, 0x61, 0x64, 0x5f, 0x70, 0x6f, 0x73, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x43, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x12, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x48, 0x65, 0x61, 0x64, 0x50, 0x6f, 0x73, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x12, 0x3c, 0x0a, 0x14, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x61,
	0x62, 0x73, 0x5f, 0x70, 0x6f, 0x73, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x44, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x11, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x62, 0x73, 0x50, 0x6f, 0x73, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x12, 0x3c, 0x0a, 0x14, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x62, 0x73,
	0x5f, 0x61, 0x6e, 0x67, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x45, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x51, 0x41, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x11, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x62, 0x73, 0x41, 0x6e, 0x67, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x22, 0xbb, 0x04, 0x0a, 0x0d, 0x43, 0x53, 0x47, 0x4f, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6d, 0x64,
	0x50, 0x42, 0x12, 0x23, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x43, 0x42, 0x61, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6d, 0x64, 0x50,
	0x42, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0d, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x43, 0x53, 0x47, 0x4f, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x42, 0x52, 0x0c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x41, 0x0a, 0x1b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b,
	0x31, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x02, 0x2d, 0x31, 0x52,
	0x18, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x31, 0x53, 0x74, 0x61, 0x72, 0x74, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x41, 0x0a, 0x1b, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x6b, 0x32, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x02,
	0x2d, 0x31, 0x52, 0x18, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x32, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x41, 0x0a, 0x1b,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x33, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x3a, 0x02, 0x2d, 0x31, 0x52, 0x18, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x33, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x31, 0x0a, 0x11, 0x6c, 0x65, 0x66, 0x74, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x5f, 0x64, 0x65, 0x73,
	0x69, 0x72, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73,
	0x65, 0x52, 0x0f, 0x6c, 0x65, 0x66, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x44, 0x65, 0x73, 0x69, 0x72,
	0x65, 0x64, 0x12, 0x41, 0x0a, 0x1a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x66, 0x78,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x16, 0x69,
	0x73, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x64, 0x79, 0x53,
	0x68, 0x6f, 0x74, 0x46, 0x78, 0x12, 0x41, 0x0a, 0x1a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x73, 0x68, 0x6f, 0x74,
	0x5f, 0x66, 0x78, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65,
	0x52, 0x16, 0x69, 0x73, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x48, 0x65,
	0x61, 0x64, 0x53, 0x68, 0x6f, 0x74, 0x46, 0x78, 0x12, 0x44, 0x0a, 0x1b, 0x69, 0x73, 0x5f, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x72,
	0x61, 0x67, 0x64, 0x6f, 0x6c, 0x6c, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66,
	0x61, 0x6c, 0x73, 0x65, 0x52, 0x18, 0x69, 0x73, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69,
	0x6e, 0x67, 0x4b, 0x69, 0x6c, 0x6c, 0x52, 0x61, 0x67, 0x64, 0x6f, 0x6c, 0x6c, 0x73,
}

var (
	file_cs_usercmd_proto_rawDescOnce sync.Once
	file_cs_usercmd_proto_rawDescData = file_cs_usercmd_proto_rawDesc
)

func file_cs_usercmd_proto_rawDescGZIP() []byte {
	file_cs_usercmd_proto_rawDescOnce.Do(func() {
		file_cs_usercmd_proto_rawDescData = protoimpl.X.CompressGZIP(file_cs_usercmd_proto_rawDescData)
	})
	return file_cs_usercmd_proto_rawDescData
}

var file_cs_usercmd_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cs_usercmd_proto_goTypes = []any{
	(*CSGOInterpolationInfoPB)(nil),    // 0: CSGOInterpolationInfoPB
	(*CSGOInterpolationInfoPB_CL)(nil), // 1: CSGOInterpolationInfoPB_CL
	(*CSGOInputHistoryEntryPB)(nil),    // 2: CSGOInputHistoryEntryPB
	(*CSGOUserCmdPB)(nil),              // 3: CSGOUserCmdPB
	(*CMsgQAngle)(nil),                 // 4: CMsgQAngle
	(*CMsgVector)(nil),                 // 5: CMsgVector
	(*CBaseUserCmdPB)(nil),             // 6: CBaseUserCmdPB
}
var file_cs_usercmd_proto_depIdxs = []int32{
	4,  // 0: CSGOInputHistoryEntryPB.view_angles:type_name -> CMsgQAngle
	1,  // 1: CSGOInputHistoryEntryPB.cl_interp:type_name -> CSGOInterpolationInfoPB_CL
	0,  // 2: CSGOInputHistoryEntryPB.sv_interp0:type_name -> CSGOInterpolationInfoPB
	0,  // 3: CSGOInputHistoryEntryPB.sv_interp1:type_name -> CSGOInterpolationInfoPB
	0,  // 4: CSGOInputHistoryEntryPB.player_interp:type_name -> CSGOInterpolationInfoPB
	5,  // 5: CSGOInputHistoryEntryPB.shoot_position:type_name -> CMsgVector
	5,  // 6: CSGOInputHistoryEntryPB.target_head_pos_check:type_name -> CMsgVector
	5,  // 7: CSGOInputHistoryEntryPB.target_abs_pos_check:type_name -> CMsgVector
	4,  // 8: CSGOInputHistoryEntryPB.target_abs_ang_check:type_name -> CMsgQAngle
	6,  // 9: CSGOUserCmdPB.base:type_name -> CBaseUserCmdPB
	2,  // 10: CSGOUserCmdPB.input_history:type_name -> CSGOInputHistoryEntryPB
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_cs_usercmd_proto_init() }
func file_cs_usercmd_proto_init() {
	if File_cs_usercmd_proto != nil {
		return
	}
	file_networkbasetypes_proto_init()
	file_usercmd_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cs_usercmd_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CSGOInterpolationInfoPB); i {
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
		file_cs_usercmd_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CSGOInterpolationInfoPB_CL); i {
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
		file_cs_usercmd_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CSGOInputHistoryEntryPB); i {
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
		file_cs_usercmd_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CSGOUserCmdPB); i {
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
			RawDescriptor: file_cs_usercmd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cs_usercmd_proto_goTypes,
		DependencyIndexes: file_cs_usercmd_proto_depIdxs,
		MessageInfos:      file_cs_usercmd_proto_msgTypes,
	}.Build()
	File_cs_usercmd_proto = out.File
	file_cs_usercmd_proto_rawDesc = nil
	file_cs_usercmd_proto_goTypes = nil
	file_cs_usercmd_proto_depIdxs = nil
}
