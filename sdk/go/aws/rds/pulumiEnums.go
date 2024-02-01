// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EngineMode string

const (
	EngineModeProvisioned   = EngineMode("provisioned")
	EngineModeServerless    = EngineMode("serverless")
	EngineModeParallelQuery = EngineMode("parallelquery")
	EngineModeGlobal        = EngineMode("global")
)

func (EngineMode) ElementType() reflect.Type {
	return reflect.TypeOf((*EngineMode)(nil)).Elem()
}

func (e EngineMode) ToEngineModeOutput() EngineModeOutput {
	return pulumi.ToOutput(e).(EngineModeOutput)
}

func (e EngineMode) ToEngineModeOutputWithContext(ctx context.Context) EngineModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EngineModeOutput)
}

func (e EngineMode) ToEngineModePtrOutput() EngineModePtrOutput {
	return e.ToEngineModePtrOutputWithContext(context.Background())
}

func (e EngineMode) ToEngineModePtrOutputWithContext(ctx context.Context) EngineModePtrOutput {
	return EngineMode(e).ToEngineModeOutputWithContext(ctx).ToEngineModePtrOutputWithContext(ctx)
}

func (e EngineMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EngineMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EngineMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EngineMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EngineModeOutput struct{ *pulumi.OutputState }

func (EngineModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EngineMode)(nil)).Elem()
}

func (o EngineModeOutput) ToEngineModeOutput() EngineModeOutput {
	return o
}

func (o EngineModeOutput) ToEngineModeOutputWithContext(ctx context.Context) EngineModeOutput {
	return o
}

func (o EngineModeOutput) ToEngineModePtrOutput() EngineModePtrOutput {
	return o.ToEngineModePtrOutputWithContext(context.Background())
}

func (o EngineModeOutput) ToEngineModePtrOutputWithContext(ctx context.Context) EngineModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EngineMode) *EngineMode {
		return &v
	}).(EngineModePtrOutput)
}

func (o EngineModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EngineModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EngineMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EngineModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EngineModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EngineMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EngineModePtrOutput struct{ *pulumi.OutputState }

func (EngineModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EngineMode)(nil)).Elem()
}

func (o EngineModePtrOutput) ToEngineModePtrOutput() EngineModePtrOutput {
	return o
}

func (o EngineModePtrOutput) ToEngineModePtrOutputWithContext(ctx context.Context) EngineModePtrOutput {
	return o
}

func (o EngineModePtrOutput) Elem() EngineModeOutput {
	return o.ApplyT(func(v *EngineMode) EngineMode {
		if v != nil {
			return *v
		}
		var ret EngineMode
		return ret
	}).(EngineModeOutput)
}

func (o EngineModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EngineModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EngineMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// EngineModeInput is an input type that accepts values of the EngineMode enum
// A concrete instance of `EngineModeInput` can be one of the following:
//
//	EngineModeProvisioned
//	EngineModeServerless
//	EngineModeParallelQuery
//	EngineModeGlobal
type EngineModeInput interface {
	pulumi.Input

	ToEngineModeOutput() EngineModeOutput
	ToEngineModeOutputWithContext(context.Context) EngineModeOutput
}

var engineModePtrType = reflect.TypeOf((**EngineMode)(nil)).Elem()

type EngineModePtrInput interface {
	pulumi.Input

	ToEngineModePtrOutput() EngineModePtrOutput
	ToEngineModePtrOutputWithContext(context.Context) EngineModePtrOutput
}

type engineModePtr string

func EngineModePtr(v string) EngineModePtrInput {
	return (*engineModePtr)(&v)
}

func (*engineModePtr) ElementType() reflect.Type {
	return engineModePtrType
}

func (in *engineModePtr) ToEngineModePtrOutput() EngineModePtrOutput {
	return pulumi.ToOutput(in).(EngineModePtrOutput)
}

func (in *engineModePtr) ToEngineModePtrOutputWithContext(ctx context.Context) EngineModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EngineModePtrOutput)
}

type EngineType string

const (
	EngineTypeAurora           = EngineType("aurora")
	EngineTypeAuroraMysql      = EngineType("aurora-mysql")
	EngineTypeAuroraPostgresql = EngineType("aurora-postgresql")
)

func (EngineType) ElementType() reflect.Type {
	return reflect.TypeOf((*EngineType)(nil)).Elem()
}

func (e EngineType) ToEngineTypeOutput() EngineTypeOutput {
	return pulumi.ToOutput(e).(EngineTypeOutput)
}

func (e EngineType) ToEngineTypeOutputWithContext(ctx context.Context) EngineTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EngineTypeOutput)
}

func (e EngineType) ToEngineTypePtrOutput() EngineTypePtrOutput {
	return e.ToEngineTypePtrOutputWithContext(context.Background())
}

func (e EngineType) ToEngineTypePtrOutputWithContext(ctx context.Context) EngineTypePtrOutput {
	return EngineType(e).ToEngineTypeOutputWithContext(ctx).ToEngineTypePtrOutputWithContext(ctx)
}

func (e EngineType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EngineType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EngineType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EngineType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EngineTypeOutput struct{ *pulumi.OutputState }

func (EngineTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EngineType)(nil)).Elem()
}

func (o EngineTypeOutput) ToEngineTypeOutput() EngineTypeOutput {
	return o
}

func (o EngineTypeOutput) ToEngineTypeOutputWithContext(ctx context.Context) EngineTypeOutput {
	return o
}

func (o EngineTypeOutput) ToEngineTypePtrOutput() EngineTypePtrOutput {
	return o.ToEngineTypePtrOutputWithContext(context.Background())
}

func (o EngineTypeOutput) ToEngineTypePtrOutputWithContext(ctx context.Context) EngineTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EngineType) *EngineType {
		return &v
	}).(EngineTypePtrOutput)
}

func (o EngineTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EngineTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EngineType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EngineTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EngineTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EngineType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EngineTypePtrOutput struct{ *pulumi.OutputState }

func (EngineTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EngineType)(nil)).Elem()
}

func (o EngineTypePtrOutput) ToEngineTypePtrOutput() EngineTypePtrOutput {
	return o
}

func (o EngineTypePtrOutput) ToEngineTypePtrOutputWithContext(ctx context.Context) EngineTypePtrOutput {
	return o
}

func (o EngineTypePtrOutput) Elem() EngineTypeOutput {
	return o.ApplyT(func(v *EngineType) EngineType {
		if v != nil {
			return *v
		}
		var ret EngineType
		return ret
	}).(EngineTypeOutput)
}

func (o EngineTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EngineTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EngineType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// EngineTypeInput is an input type that accepts values of the EngineType enum
// A concrete instance of `EngineTypeInput` can be one of the following:
//
//	EngineTypeAurora
//	EngineTypeAuroraMysql
//	EngineTypeAuroraPostgresql
type EngineTypeInput interface {
	pulumi.Input

	ToEngineTypeOutput() EngineTypeOutput
	ToEngineTypeOutputWithContext(context.Context) EngineTypeOutput
}

var engineTypePtrType = reflect.TypeOf((**EngineType)(nil)).Elem()

type EngineTypePtrInput interface {
	pulumi.Input

	ToEngineTypePtrOutput() EngineTypePtrOutput
	ToEngineTypePtrOutputWithContext(context.Context) EngineTypePtrOutput
}

type engineTypePtr string

func EngineTypePtr(v string) EngineTypePtrInput {
	return (*engineTypePtr)(&v)
}

func (*engineTypePtr) ElementType() reflect.Type {
	return engineTypePtrType
}

func (in *engineTypePtr) ToEngineTypePtrOutput() EngineTypePtrOutput {
	return pulumi.ToOutput(in).(EngineTypePtrOutput)
}

func (in *engineTypePtr) ToEngineTypePtrOutputWithContext(ctx context.Context) EngineTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EngineTypePtrOutput)
}

type InstanceType string

const (
	InstanceType_T4G_Micro    = InstanceType("db.t4g.micro")
	InstanceType_T4G_Small    = InstanceType("db.t4g.small")
	InstanceType_T4G_Medium   = InstanceType("db.t4g.medium")
	InstanceType_T4G_Large    = InstanceType("db.t4g.large")
	InstanceType_T4G_XLarge   = InstanceType("db.t4g.xlarge")
	InstanceType_T4G_2XLarge  = InstanceType("db.t4g.2xlarge")
	InstanceType_T3_Micro     = InstanceType("db.t3.micro")
	InstanceType_T3_Small     = InstanceType("db.t3.small")
	InstanceType_T3_Medium    = InstanceType("db.t3.medium")
	InstanceType_T3_Large     = InstanceType("db.t3.large")
	InstanceType_T3_XLarge    = InstanceType("db.t3.xlarge")
	InstanceType_T3_2XLarge   = InstanceType("db.t3.2xlarge")
	InstanceType_T2_Micro     = InstanceType("db.t2.micro")
	InstanceType_T2_Small     = InstanceType("db.t2.small")
	InstanceType_T2_Medium    = InstanceType("db.t2.medium")
	InstanceType_T2_Large     = InstanceType("db.t2.large")
	InstanceType_T2_XLarge    = InstanceType("db.t2.xlarge")
	InstanceType_T2_2XLarge   = InstanceType("db.t2.2xlarge")
	InstanceType_M1_Small     = InstanceType("db.m1.small")
	InstanceType_M1_Medium    = InstanceType("db.m1.medium")
	InstanceType_M1_Large     = InstanceType("db.m1.large")
	InstanceType_M1_XLarge    = InstanceType("db.m1.xlarge")
	InstanceType_M2_XLarge    = InstanceType("db.m2.xlarge")
	InstanceType_M2_2XLarge   = InstanceType("db.m2.2xlarge")
	InstanceType_M2_4XLarge   = InstanceType("db.m2.4xlarge")
	InstanceType_M3_Medium    = InstanceType("db.m3.medium")
	InstanceType_M3_Large     = InstanceType("db.m3.large")
	InstanceType_M3_XLarge    = InstanceType("db.m3.xlarge")
	InstanceType_M3_2XLarge   = InstanceType("db.m3.2xlarge")
	InstanceType_M4_Large     = InstanceType("db.m4.large")
	InstanceType_M4_XLarge    = InstanceType("db.m4.xlarge")
	InstanceType_M4_2XLarge   = InstanceType("db.m4.2xlarge")
	InstanceType_M4_4XLarge   = InstanceType("db.m4.4xlarge")
	InstanceType_M4_10XLarge  = InstanceType("db.m4.10xlarge")
	InstanceType_M4_16XLarge  = InstanceType("db.m4.10xlarge")
	InstanceType_M5_Large     = InstanceType("db.m5.large")
	InstanceType_M5_XLarge    = InstanceType("db.m5.xlarge")
	InstanceType_M5_2XLarge   = InstanceType("db.m5.2xlarge")
	InstanceType_M5_4XLarge   = InstanceType("db.m5.4xlarge")
	InstanceType_M5_12XLarge  = InstanceType("db.m5.12xlarge")
	InstanceType_M5_24XLarge  = InstanceType("db.m5.24xlarge")
	InstanceType_M6G_Large    = InstanceType("db.m6g.large")
	InstanceType_M6G_XLarge   = InstanceType("db.m6g.xlarge")
	InstanceType_M6G_2XLarge  = InstanceType("db.m6g.2xlarge")
	InstanceType_M6G_4XLarge  = InstanceType("db.m6g.4xlarge")
	InstanceType_M6G_8XLarge  = InstanceType("db.m6g.8xlarge")
	InstanceType_M6G_12XLarge = InstanceType("db.m6g.12xlarge")
	InstanceType_M6G_16XLarge = InstanceType("db.m6g.16xlarge")
	InstanceType_R3_Large     = InstanceType("db.r3.large")
	InstanceType_R3_XLarge    = InstanceType("db.r3.xlarge")
	InstanceType_R3_2XLarge   = InstanceType("db.r3.2xlarge")
	InstanceType_R3_4XLarge   = InstanceType("db.r3.4xlarge")
	InstanceType_R3_8XLarge   = InstanceType("db.r3.8xlarge")
	InstanceType_R4_Large     = InstanceType("db.r4.large")
	InstanceType_R4_XLarge    = InstanceType("db.r4.xlarge")
	InstanceType_R4_2XLarge   = InstanceType("db.r4.2xlarge")
	InstanceType_R4_4XLarge   = InstanceType("db.r4.4xlarge")
	InstanceType_R4_8XLarge   = InstanceType("db.r4.8xlarge")
	InstanceType_R4_16XLarge  = InstanceType("db.r4.16xlarge")
	InstanceType_R5_Large     = InstanceType("db.r5.large")
	InstanceType_R5_XLarge    = InstanceType("db.r5.xlarge")
	InstanceType_R5_2XLarge   = InstanceType("db.r5.2xlarge")
	InstanceType_R5_4XLarge   = InstanceType("db.r5.4xlarge")
	InstanceType_R5_12XLarge  = InstanceType("db.r5.12xlarge")
	InstanceType_R5_24XLarge  = InstanceType("db.r5.24xlarge")
	InstanceType_R6G_Large    = InstanceType("db.r6g.large")
	InstanceType_R6G_XLarge   = InstanceType("db.r6g.xlarge")
	InstanceType_R6G_2XLarge  = InstanceType("db.r6g.2xlarge")
	InstanceType_R6G_4XLarge  = InstanceType("db.r6g.4xlarge")
	InstanceType_R6G_8XLarge  = InstanceType("db.r6g.8xlarge")
	InstanceType_R6G_12XLarge = InstanceType("db.r6g.12xlarge")
	InstanceType_R6G_16XLarge = InstanceType("db.r6g.16xlarge")
	InstanceType_X1_16XLarge  = InstanceType("db.x1.16xlarge")
	InstanceType_X1_32XLarge  = InstanceType("db.x1.32xlarge")
	InstanceType_X1E_XLarge   = InstanceType("db.x1e.xlarge")
	InstanceType_X1E_2XLarge  = InstanceType("db.x1e.2xlarge")
	InstanceType_X1E_4XLarge  = InstanceType("db.x1e.4xlarge")
	InstanceType_X1E_8XLarge  = InstanceType("db.x1e.8xlarge")
	InstanceType_X1E_32XLarge = InstanceType("db.x1e.32xlarge")
)

func (InstanceType) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceType)(nil)).Elem()
}

func (e InstanceType) ToInstanceTypeOutput() InstanceTypeOutput {
	return pulumi.ToOutput(e).(InstanceTypeOutput)
}

func (e InstanceType) ToInstanceTypeOutputWithContext(ctx context.Context) InstanceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InstanceTypeOutput)
}

func (e InstanceType) ToInstanceTypePtrOutput() InstanceTypePtrOutput {
	return e.ToInstanceTypePtrOutputWithContext(context.Background())
}

func (e InstanceType) ToInstanceTypePtrOutputWithContext(ctx context.Context) InstanceTypePtrOutput {
	return InstanceType(e).ToInstanceTypeOutputWithContext(ctx).ToInstanceTypePtrOutputWithContext(ctx)
}

func (e InstanceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InstanceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InstanceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InstanceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InstanceTypeOutput struct{ *pulumi.OutputState }

func (InstanceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceType)(nil)).Elem()
}

func (o InstanceTypeOutput) ToInstanceTypeOutput() InstanceTypeOutput {
	return o
}

func (o InstanceTypeOutput) ToInstanceTypeOutputWithContext(ctx context.Context) InstanceTypeOutput {
	return o
}

func (o InstanceTypeOutput) ToInstanceTypePtrOutput() InstanceTypePtrOutput {
	return o.ToInstanceTypePtrOutputWithContext(context.Background())
}

func (o InstanceTypeOutput) ToInstanceTypePtrOutputWithContext(ctx context.Context) InstanceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InstanceType) *InstanceType {
		return &v
	}).(InstanceTypePtrOutput)
}

func (o InstanceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InstanceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InstanceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InstanceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InstanceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InstanceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InstanceTypePtrOutput struct{ *pulumi.OutputState }

func (InstanceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceType)(nil)).Elem()
}

func (o InstanceTypePtrOutput) ToInstanceTypePtrOutput() InstanceTypePtrOutput {
	return o
}

func (o InstanceTypePtrOutput) ToInstanceTypePtrOutputWithContext(ctx context.Context) InstanceTypePtrOutput {
	return o
}

func (o InstanceTypePtrOutput) Elem() InstanceTypeOutput {
	return o.ApplyT(func(v *InstanceType) InstanceType {
		if v != nil {
			return *v
		}
		var ret InstanceType
		return ret
	}).(InstanceTypeOutput)
}

func (o InstanceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InstanceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InstanceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// InstanceTypeInput is an input type that accepts values of the InstanceType enum
// A concrete instance of `InstanceTypeInput` can be one of the following:
//
//	InstanceType_T4G_Micro
//	InstanceType_T4G_Small
//	InstanceType_T4G_Medium
//	InstanceType_T4G_Large
//	InstanceType_T4G_XLarge
//	InstanceType_T4G_2XLarge
//	InstanceType_T3_Micro
//	InstanceType_T3_Small
//	InstanceType_T3_Medium
//	InstanceType_T3_Large
//	InstanceType_T3_XLarge
//	InstanceType_T3_2XLarge
//	InstanceType_T2_Micro
//	InstanceType_T2_Small
//	InstanceType_T2_Medium
//	InstanceType_T2_Large
//	InstanceType_T2_XLarge
//	InstanceType_T2_2XLarge
//	InstanceType_M1_Small
//	InstanceType_M1_Medium
//	InstanceType_M1_Large
//	InstanceType_M1_XLarge
//	InstanceType_M2_XLarge
//	InstanceType_M2_2XLarge
//	InstanceType_M2_4XLarge
//	InstanceType_M3_Medium
//	InstanceType_M3_Large
//	InstanceType_M3_XLarge
//	InstanceType_M3_2XLarge
//	InstanceType_M4_Large
//	InstanceType_M4_XLarge
//	InstanceType_M4_2XLarge
//	InstanceType_M4_4XLarge
//	InstanceType_M4_10XLarge
//	InstanceType_M4_16XLarge
//	InstanceType_M5_Large
//	InstanceType_M5_XLarge
//	InstanceType_M5_2XLarge
//	InstanceType_M5_4XLarge
//	InstanceType_M5_12XLarge
//	InstanceType_M5_24XLarge
//	InstanceType_M6G_Large
//	InstanceType_M6G_XLarge
//	InstanceType_M6G_2XLarge
//	InstanceType_M6G_4XLarge
//	InstanceType_M6G_8XLarge
//	InstanceType_M6G_12XLarge
//	InstanceType_M6G_16XLarge
//	InstanceType_R3_Large
//	InstanceType_R3_XLarge
//	InstanceType_R3_2XLarge
//	InstanceType_R3_4XLarge
//	InstanceType_R3_8XLarge
//	InstanceType_R4_Large
//	InstanceType_R4_XLarge
//	InstanceType_R4_2XLarge
//	InstanceType_R4_4XLarge
//	InstanceType_R4_8XLarge
//	InstanceType_R4_16XLarge
//	InstanceType_R5_Large
//	InstanceType_R5_XLarge
//	InstanceType_R5_2XLarge
//	InstanceType_R5_4XLarge
//	InstanceType_R5_12XLarge
//	InstanceType_R5_24XLarge
//	InstanceType_R6G_Large
//	InstanceType_R6G_XLarge
//	InstanceType_R6G_2XLarge
//	InstanceType_R6G_4XLarge
//	InstanceType_R6G_8XLarge
//	InstanceType_R6G_12XLarge
//	InstanceType_R6G_16XLarge
//	InstanceType_X1_16XLarge
//	InstanceType_X1_32XLarge
//	InstanceType_X1E_XLarge
//	InstanceType_X1E_2XLarge
//	InstanceType_X1E_4XLarge
//	InstanceType_X1E_8XLarge
//	InstanceType_X1E_32XLarge
type InstanceTypeInput interface {
	pulumi.Input

	ToInstanceTypeOutput() InstanceTypeOutput
	ToInstanceTypeOutputWithContext(context.Context) InstanceTypeOutput
}

var instanceTypePtrType = reflect.TypeOf((**InstanceType)(nil)).Elem()

type InstanceTypePtrInput interface {
	pulumi.Input

	ToInstanceTypePtrOutput() InstanceTypePtrOutput
	ToInstanceTypePtrOutputWithContext(context.Context) InstanceTypePtrOutput
}

type instanceTypePtr string

func InstanceTypePtr(v string) InstanceTypePtrInput {
	return (*instanceTypePtr)(&v)
}

func (*instanceTypePtr) ElementType() reflect.Type {
	return instanceTypePtrType
}

func (in *instanceTypePtr) ToInstanceTypePtrOutput() InstanceTypePtrOutput {
	return pulumi.ToOutput(in).(InstanceTypePtrOutput)
}

func (in *instanceTypePtr) ToInstanceTypePtrOutputWithContext(ctx context.Context) InstanceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InstanceTypePtrOutput)
}

type StorageType string

const (
	StorageTypeStandard = StorageType("standard")
	StorageTypeGP2      = StorageType("gp2")
	StorageTypeGP3      = StorageType("gp3")
	StorageTypeIO1      = StorageType("io1")
)

func (StorageType) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageType)(nil)).Elem()
}

func (e StorageType) ToStorageTypeOutput() StorageTypeOutput {
	return pulumi.ToOutput(e).(StorageTypeOutput)
}

func (e StorageType) ToStorageTypeOutputWithContext(ctx context.Context) StorageTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StorageTypeOutput)
}

func (e StorageType) ToStorageTypePtrOutput() StorageTypePtrOutput {
	return e.ToStorageTypePtrOutputWithContext(context.Background())
}

func (e StorageType) ToStorageTypePtrOutputWithContext(ctx context.Context) StorageTypePtrOutput {
	return StorageType(e).ToStorageTypeOutputWithContext(ctx).ToStorageTypePtrOutputWithContext(ctx)
}

func (e StorageType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StorageType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StorageTypeOutput struct{ *pulumi.OutputState }

func (StorageTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageType)(nil)).Elem()
}

func (o StorageTypeOutput) ToStorageTypeOutput() StorageTypeOutput {
	return o
}

func (o StorageTypeOutput) ToStorageTypeOutputWithContext(ctx context.Context) StorageTypeOutput {
	return o
}

func (o StorageTypeOutput) ToStorageTypePtrOutput() StorageTypePtrOutput {
	return o.ToStorageTypePtrOutputWithContext(context.Background())
}

func (o StorageTypeOutput) ToStorageTypePtrOutputWithContext(ctx context.Context) StorageTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageType) *StorageType {
		return &v
	}).(StorageTypePtrOutput)
}

func (o StorageTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StorageTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StorageTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StorageTypePtrOutput struct{ *pulumi.OutputState }

func (StorageTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageType)(nil)).Elem()
}

func (o StorageTypePtrOutput) ToStorageTypePtrOutput() StorageTypePtrOutput {
	return o
}

func (o StorageTypePtrOutput) ToStorageTypePtrOutputWithContext(ctx context.Context) StorageTypePtrOutput {
	return o
}

func (o StorageTypePtrOutput) Elem() StorageTypeOutput {
	return o.ApplyT(func(v *StorageType) StorageType {
		if v != nil {
			return *v
		}
		var ret StorageType
		return ret
	}).(StorageTypeOutput)
}

func (o StorageTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StorageType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// StorageTypeInput is an input type that accepts values of the StorageType enum
// A concrete instance of `StorageTypeInput` can be one of the following:
//
//	StorageTypeStandard
//	StorageTypeGP2
//	StorageTypeGP3
//	StorageTypeIO1
type StorageTypeInput interface {
	pulumi.Input

	ToStorageTypeOutput() StorageTypeOutput
	ToStorageTypeOutputWithContext(context.Context) StorageTypeOutput
}

var storageTypePtrType = reflect.TypeOf((**StorageType)(nil)).Elem()

type StorageTypePtrInput interface {
	pulumi.Input

	ToStorageTypePtrOutput() StorageTypePtrOutput
	ToStorageTypePtrOutputWithContext(context.Context) StorageTypePtrOutput
}

type storageTypePtr string

func StorageTypePtr(v string) StorageTypePtrInput {
	return (*storageTypePtr)(&v)
}

func (*storageTypePtr) ElementType() reflect.Type {
	return storageTypePtrType
}

func (in *storageTypePtr) ToStorageTypePtrOutput() StorageTypePtrOutput {
	return pulumi.ToOutput(in).(StorageTypePtrOutput)
}

func (in *storageTypePtr) ToStorageTypePtrOutputWithContext(ctx context.Context) StorageTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StorageTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EngineModeInput)(nil)).Elem(), EngineMode("provisioned"))
	pulumi.RegisterInputType(reflect.TypeOf((*EngineModePtrInput)(nil)).Elem(), EngineMode("provisioned"))
	pulumi.RegisterInputType(reflect.TypeOf((*EngineTypeInput)(nil)).Elem(), EngineType("aurora"))
	pulumi.RegisterInputType(reflect.TypeOf((*EngineTypePtrInput)(nil)).Elem(), EngineType("aurora"))
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceTypeInput)(nil)).Elem(), InstanceType("db.t4g.micro"))
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceTypePtrInput)(nil)).Elem(), InstanceType("db.t4g.micro"))
	pulumi.RegisterInputType(reflect.TypeOf((*StorageTypeInput)(nil)).Elem(), StorageType("standard"))
	pulumi.RegisterInputType(reflect.TypeOf((*StorageTypePtrInput)(nil)).Elem(), StorageType("standard"))
	pulumi.RegisterOutputType(EngineModeOutput{})
	pulumi.RegisterOutputType(EngineModePtrOutput{})
	pulumi.RegisterOutputType(EngineTypeOutput{})
	pulumi.RegisterOutputType(EngineTypePtrOutput{})
	pulumi.RegisterOutputType(InstanceTypeOutput{})
	pulumi.RegisterOutputType(InstanceTypePtrOutput{})
	pulumi.RegisterOutputType(StorageTypeOutput{})
	pulumi.RegisterOutputType(StorageTypePtrOutput{})
}
