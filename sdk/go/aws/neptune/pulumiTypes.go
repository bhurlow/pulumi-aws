// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package neptune

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClusterParameterGroupParameter struct {
	// Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
	ApplyMethod *string `pulumi:"applyMethod"`
	// The name of the neptune parameter.
	Name string `pulumi:"name"`
	// The value of the neptune parameter.
	Value string `pulumi:"value"`
}

// ClusterParameterGroupParameterInput is an input type that accepts ClusterParameterGroupParameterArgs and ClusterParameterGroupParameterOutput values.
// You can construct a concrete instance of `ClusterParameterGroupParameterInput` via:
//
//	ClusterParameterGroupParameterArgs{...}
type ClusterParameterGroupParameterInput interface {
	pulumi.Input

	ToClusterParameterGroupParameterOutput() ClusterParameterGroupParameterOutput
	ToClusterParameterGroupParameterOutputWithContext(context.Context) ClusterParameterGroupParameterOutput
}

type ClusterParameterGroupParameterArgs struct {
	// Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
	ApplyMethod pulumi.StringPtrInput `pulumi:"applyMethod"`
	// The name of the neptune parameter.
	Name pulumi.StringInput `pulumi:"name"`
	// The value of the neptune parameter.
	Value pulumi.StringInput `pulumi:"value"`
}

func (ClusterParameterGroupParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterParameterGroupParameter)(nil)).Elem()
}

func (i ClusterParameterGroupParameterArgs) ToClusterParameterGroupParameterOutput() ClusterParameterGroupParameterOutput {
	return i.ToClusterParameterGroupParameterOutputWithContext(context.Background())
}

func (i ClusterParameterGroupParameterArgs) ToClusterParameterGroupParameterOutputWithContext(ctx context.Context) ClusterParameterGroupParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterParameterGroupParameterOutput)
}

// ClusterParameterGroupParameterArrayInput is an input type that accepts ClusterParameterGroupParameterArray and ClusterParameterGroupParameterArrayOutput values.
// You can construct a concrete instance of `ClusterParameterGroupParameterArrayInput` via:
//
//	ClusterParameterGroupParameterArray{ ClusterParameterGroupParameterArgs{...} }
type ClusterParameterGroupParameterArrayInput interface {
	pulumi.Input

	ToClusterParameterGroupParameterArrayOutput() ClusterParameterGroupParameterArrayOutput
	ToClusterParameterGroupParameterArrayOutputWithContext(context.Context) ClusterParameterGroupParameterArrayOutput
}

type ClusterParameterGroupParameterArray []ClusterParameterGroupParameterInput

func (ClusterParameterGroupParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterParameterGroupParameter)(nil)).Elem()
}

func (i ClusterParameterGroupParameterArray) ToClusterParameterGroupParameterArrayOutput() ClusterParameterGroupParameterArrayOutput {
	return i.ToClusterParameterGroupParameterArrayOutputWithContext(context.Background())
}

func (i ClusterParameterGroupParameterArray) ToClusterParameterGroupParameterArrayOutputWithContext(ctx context.Context) ClusterParameterGroupParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterParameterGroupParameterArrayOutput)
}

type ClusterParameterGroupParameterOutput struct{ *pulumi.OutputState }

func (ClusterParameterGroupParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterParameterGroupParameter)(nil)).Elem()
}

func (o ClusterParameterGroupParameterOutput) ToClusterParameterGroupParameterOutput() ClusterParameterGroupParameterOutput {
	return o
}

func (o ClusterParameterGroupParameterOutput) ToClusterParameterGroupParameterOutputWithContext(ctx context.Context) ClusterParameterGroupParameterOutput {
	return o
}

// Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
func (o ClusterParameterGroupParameterOutput) ApplyMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterParameterGroupParameter) *string { return v.ApplyMethod }).(pulumi.StringPtrOutput)
}

// The name of the neptune parameter.
func (o ClusterParameterGroupParameterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterParameterGroupParameter) string { return v.Name }).(pulumi.StringOutput)
}

// The value of the neptune parameter.
func (o ClusterParameterGroupParameterOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterParameterGroupParameter) string { return v.Value }).(pulumi.StringOutput)
}

type ClusterParameterGroupParameterArrayOutput struct{ *pulumi.OutputState }

func (ClusterParameterGroupParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterParameterGroupParameter)(nil)).Elem()
}

func (o ClusterParameterGroupParameterArrayOutput) ToClusterParameterGroupParameterArrayOutput() ClusterParameterGroupParameterArrayOutput {
	return o
}

func (o ClusterParameterGroupParameterArrayOutput) ToClusterParameterGroupParameterArrayOutputWithContext(ctx context.Context) ClusterParameterGroupParameterArrayOutput {
	return o
}

func (o ClusterParameterGroupParameterArrayOutput) Index(i pulumi.IntInput) ClusterParameterGroupParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterParameterGroupParameter {
		return vs[0].([]ClusterParameterGroupParameter)[vs[1].(int)]
	}).(ClusterParameterGroupParameterOutput)
}

type ClusterServerlessV2ScalingConfiguration struct {
	MaxCapacity *float64 `pulumi:"maxCapacity"`
	MinCapacity *float64 `pulumi:"minCapacity"`
}

// ClusterServerlessV2ScalingConfigurationInput is an input type that accepts ClusterServerlessV2ScalingConfigurationArgs and ClusterServerlessV2ScalingConfigurationOutput values.
// You can construct a concrete instance of `ClusterServerlessV2ScalingConfigurationInput` via:
//
//	ClusterServerlessV2ScalingConfigurationArgs{...}
type ClusterServerlessV2ScalingConfigurationInput interface {
	pulumi.Input

	ToClusterServerlessV2ScalingConfigurationOutput() ClusterServerlessV2ScalingConfigurationOutput
	ToClusterServerlessV2ScalingConfigurationOutputWithContext(context.Context) ClusterServerlessV2ScalingConfigurationOutput
}

type ClusterServerlessV2ScalingConfigurationArgs struct {
	MaxCapacity pulumi.Float64PtrInput `pulumi:"maxCapacity"`
	MinCapacity pulumi.Float64PtrInput `pulumi:"minCapacity"`
}

func (ClusterServerlessV2ScalingConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterServerlessV2ScalingConfiguration)(nil)).Elem()
}

func (i ClusterServerlessV2ScalingConfigurationArgs) ToClusterServerlessV2ScalingConfigurationOutput() ClusterServerlessV2ScalingConfigurationOutput {
	return i.ToClusterServerlessV2ScalingConfigurationOutputWithContext(context.Background())
}

func (i ClusterServerlessV2ScalingConfigurationArgs) ToClusterServerlessV2ScalingConfigurationOutputWithContext(ctx context.Context) ClusterServerlessV2ScalingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterServerlessV2ScalingConfigurationOutput)
}

func (i ClusterServerlessV2ScalingConfigurationArgs) ToClusterServerlessV2ScalingConfigurationPtrOutput() ClusterServerlessV2ScalingConfigurationPtrOutput {
	return i.ToClusterServerlessV2ScalingConfigurationPtrOutputWithContext(context.Background())
}

func (i ClusterServerlessV2ScalingConfigurationArgs) ToClusterServerlessV2ScalingConfigurationPtrOutputWithContext(ctx context.Context) ClusterServerlessV2ScalingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterServerlessV2ScalingConfigurationOutput).ToClusterServerlessV2ScalingConfigurationPtrOutputWithContext(ctx)
}

// ClusterServerlessV2ScalingConfigurationPtrInput is an input type that accepts ClusterServerlessV2ScalingConfigurationArgs, ClusterServerlessV2ScalingConfigurationPtr and ClusterServerlessV2ScalingConfigurationPtrOutput values.
// You can construct a concrete instance of `ClusterServerlessV2ScalingConfigurationPtrInput` via:
//
//	        ClusterServerlessV2ScalingConfigurationArgs{...}
//
//	or:
//
//	        nil
type ClusterServerlessV2ScalingConfigurationPtrInput interface {
	pulumi.Input

	ToClusterServerlessV2ScalingConfigurationPtrOutput() ClusterServerlessV2ScalingConfigurationPtrOutput
	ToClusterServerlessV2ScalingConfigurationPtrOutputWithContext(context.Context) ClusterServerlessV2ScalingConfigurationPtrOutput
}

type clusterServerlessV2ScalingConfigurationPtrType ClusterServerlessV2ScalingConfigurationArgs

func ClusterServerlessV2ScalingConfigurationPtr(v *ClusterServerlessV2ScalingConfigurationArgs) ClusterServerlessV2ScalingConfigurationPtrInput {
	return (*clusterServerlessV2ScalingConfigurationPtrType)(v)
}

func (*clusterServerlessV2ScalingConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterServerlessV2ScalingConfiguration)(nil)).Elem()
}

func (i *clusterServerlessV2ScalingConfigurationPtrType) ToClusterServerlessV2ScalingConfigurationPtrOutput() ClusterServerlessV2ScalingConfigurationPtrOutput {
	return i.ToClusterServerlessV2ScalingConfigurationPtrOutputWithContext(context.Background())
}

func (i *clusterServerlessV2ScalingConfigurationPtrType) ToClusterServerlessV2ScalingConfigurationPtrOutputWithContext(ctx context.Context) ClusterServerlessV2ScalingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterServerlessV2ScalingConfigurationPtrOutput)
}

type ClusterServerlessV2ScalingConfigurationOutput struct{ *pulumi.OutputState }

func (ClusterServerlessV2ScalingConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterServerlessV2ScalingConfiguration)(nil)).Elem()
}

func (o ClusterServerlessV2ScalingConfigurationOutput) ToClusterServerlessV2ScalingConfigurationOutput() ClusterServerlessV2ScalingConfigurationOutput {
	return o
}

func (o ClusterServerlessV2ScalingConfigurationOutput) ToClusterServerlessV2ScalingConfigurationOutputWithContext(ctx context.Context) ClusterServerlessV2ScalingConfigurationOutput {
	return o
}

func (o ClusterServerlessV2ScalingConfigurationOutput) ToClusterServerlessV2ScalingConfigurationPtrOutput() ClusterServerlessV2ScalingConfigurationPtrOutput {
	return o.ToClusterServerlessV2ScalingConfigurationPtrOutputWithContext(context.Background())
}

func (o ClusterServerlessV2ScalingConfigurationOutput) ToClusterServerlessV2ScalingConfigurationPtrOutputWithContext(ctx context.Context) ClusterServerlessV2ScalingConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterServerlessV2ScalingConfiguration) *ClusterServerlessV2ScalingConfiguration {
		return &v
	}).(ClusterServerlessV2ScalingConfigurationPtrOutput)
}

func (o ClusterServerlessV2ScalingConfigurationOutput) MaxCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ClusterServerlessV2ScalingConfiguration) *float64 { return v.MaxCapacity }).(pulumi.Float64PtrOutput)
}

func (o ClusterServerlessV2ScalingConfigurationOutput) MinCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ClusterServerlessV2ScalingConfiguration) *float64 { return v.MinCapacity }).(pulumi.Float64PtrOutput)
}

type ClusterServerlessV2ScalingConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ClusterServerlessV2ScalingConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterServerlessV2ScalingConfiguration)(nil)).Elem()
}

func (o ClusterServerlessV2ScalingConfigurationPtrOutput) ToClusterServerlessV2ScalingConfigurationPtrOutput() ClusterServerlessV2ScalingConfigurationPtrOutput {
	return o
}

func (o ClusterServerlessV2ScalingConfigurationPtrOutput) ToClusterServerlessV2ScalingConfigurationPtrOutputWithContext(ctx context.Context) ClusterServerlessV2ScalingConfigurationPtrOutput {
	return o
}

func (o ClusterServerlessV2ScalingConfigurationPtrOutput) Elem() ClusterServerlessV2ScalingConfigurationOutput {
	return o.ApplyT(func(v *ClusterServerlessV2ScalingConfiguration) ClusterServerlessV2ScalingConfiguration {
		if v != nil {
			return *v
		}
		var ret ClusterServerlessV2ScalingConfiguration
		return ret
	}).(ClusterServerlessV2ScalingConfigurationOutput)
}

func (o ClusterServerlessV2ScalingConfigurationPtrOutput) MaxCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ClusterServerlessV2ScalingConfiguration) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxCapacity
	}).(pulumi.Float64PtrOutput)
}

func (o ClusterServerlessV2ScalingConfigurationPtrOutput) MinCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ClusterServerlessV2ScalingConfiguration) *float64 {
		if v == nil {
			return nil
		}
		return v.MinCapacity
	}).(pulumi.Float64PtrOutput)
}

type GlobalClusterGlobalClusterMember struct {
	// Amazon Resource Name (ARN) of member DB Cluster.
	DbClusterArn *string `pulumi:"dbClusterArn"`
	// Whether the member is the primary DB Cluster.
	IsWriter *bool `pulumi:"isWriter"`
}

// GlobalClusterGlobalClusterMemberInput is an input type that accepts GlobalClusterGlobalClusterMemberArgs and GlobalClusterGlobalClusterMemberOutput values.
// You can construct a concrete instance of `GlobalClusterGlobalClusterMemberInput` via:
//
//	GlobalClusterGlobalClusterMemberArgs{...}
type GlobalClusterGlobalClusterMemberInput interface {
	pulumi.Input

	ToGlobalClusterGlobalClusterMemberOutput() GlobalClusterGlobalClusterMemberOutput
	ToGlobalClusterGlobalClusterMemberOutputWithContext(context.Context) GlobalClusterGlobalClusterMemberOutput
}

type GlobalClusterGlobalClusterMemberArgs struct {
	// Amazon Resource Name (ARN) of member DB Cluster.
	DbClusterArn pulumi.StringPtrInput `pulumi:"dbClusterArn"`
	// Whether the member is the primary DB Cluster.
	IsWriter pulumi.BoolPtrInput `pulumi:"isWriter"`
}

func (GlobalClusterGlobalClusterMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalClusterGlobalClusterMember)(nil)).Elem()
}

func (i GlobalClusterGlobalClusterMemberArgs) ToGlobalClusterGlobalClusterMemberOutput() GlobalClusterGlobalClusterMemberOutput {
	return i.ToGlobalClusterGlobalClusterMemberOutputWithContext(context.Background())
}

func (i GlobalClusterGlobalClusterMemberArgs) ToGlobalClusterGlobalClusterMemberOutputWithContext(ctx context.Context) GlobalClusterGlobalClusterMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalClusterGlobalClusterMemberOutput)
}

// GlobalClusterGlobalClusterMemberArrayInput is an input type that accepts GlobalClusterGlobalClusterMemberArray and GlobalClusterGlobalClusterMemberArrayOutput values.
// You can construct a concrete instance of `GlobalClusterGlobalClusterMemberArrayInput` via:
//
//	GlobalClusterGlobalClusterMemberArray{ GlobalClusterGlobalClusterMemberArgs{...} }
type GlobalClusterGlobalClusterMemberArrayInput interface {
	pulumi.Input

	ToGlobalClusterGlobalClusterMemberArrayOutput() GlobalClusterGlobalClusterMemberArrayOutput
	ToGlobalClusterGlobalClusterMemberArrayOutputWithContext(context.Context) GlobalClusterGlobalClusterMemberArrayOutput
}

type GlobalClusterGlobalClusterMemberArray []GlobalClusterGlobalClusterMemberInput

func (GlobalClusterGlobalClusterMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GlobalClusterGlobalClusterMember)(nil)).Elem()
}

func (i GlobalClusterGlobalClusterMemberArray) ToGlobalClusterGlobalClusterMemberArrayOutput() GlobalClusterGlobalClusterMemberArrayOutput {
	return i.ToGlobalClusterGlobalClusterMemberArrayOutputWithContext(context.Background())
}

func (i GlobalClusterGlobalClusterMemberArray) ToGlobalClusterGlobalClusterMemberArrayOutputWithContext(ctx context.Context) GlobalClusterGlobalClusterMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalClusterGlobalClusterMemberArrayOutput)
}

type GlobalClusterGlobalClusterMemberOutput struct{ *pulumi.OutputState }

func (GlobalClusterGlobalClusterMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalClusterGlobalClusterMember)(nil)).Elem()
}

func (o GlobalClusterGlobalClusterMemberOutput) ToGlobalClusterGlobalClusterMemberOutput() GlobalClusterGlobalClusterMemberOutput {
	return o
}

func (o GlobalClusterGlobalClusterMemberOutput) ToGlobalClusterGlobalClusterMemberOutputWithContext(ctx context.Context) GlobalClusterGlobalClusterMemberOutput {
	return o
}

// Amazon Resource Name (ARN) of member DB Cluster.
func (o GlobalClusterGlobalClusterMemberOutput) DbClusterArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GlobalClusterGlobalClusterMember) *string { return v.DbClusterArn }).(pulumi.StringPtrOutput)
}

// Whether the member is the primary DB Cluster.
func (o GlobalClusterGlobalClusterMemberOutput) IsWriter() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GlobalClusterGlobalClusterMember) *bool { return v.IsWriter }).(pulumi.BoolPtrOutput)
}

type GlobalClusterGlobalClusterMemberArrayOutput struct{ *pulumi.OutputState }

func (GlobalClusterGlobalClusterMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GlobalClusterGlobalClusterMember)(nil)).Elem()
}

func (o GlobalClusterGlobalClusterMemberArrayOutput) ToGlobalClusterGlobalClusterMemberArrayOutput() GlobalClusterGlobalClusterMemberArrayOutput {
	return o
}

func (o GlobalClusterGlobalClusterMemberArrayOutput) ToGlobalClusterGlobalClusterMemberArrayOutputWithContext(ctx context.Context) GlobalClusterGlobalClusterMemberArrayOutput {
	return o
}

func (o GlobalClusterGlobalClusterMemberArrayOutput) Index(i pulumi.IntInput) GlobalClusterGlobalClusterMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GlobalClusterGlobalClusterMember {
		return vs[0].([]GlobalClusterGlobalClusterMember)[vs[1].(int)]
	}).(GlobalClusterGlobalClusterMemberOutput)
}

type ParameterGroupParameter struct {
	// The apply method of the Neptune parameter. Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
	ApplyMethod *string `pulumi:"applyMethod"`
	// The name of the Neptune parameter.
	Name string `pulumi:"name"`
	// The value of the Neptune parameter.
	Value string `pulumi:"value"`
}

// ParameterGroupParameterInput is an input type that accepts ParameterGroupParameterArgs and ParameterGroupParameterOutput values.
// You can construct a concrete instance of `ParameterGroupParameterInput` via:
//
//	ParameterGroupParameterArgs{...}
type ParameterGroupParameterInput interface {
	pulumi.Input

	ToParameterGroupParameterOutput() ParameterGroupParameterOutput
	ToParameterGroupParameterOutputWithContext(context.Context) ParameterGroupParameterOutput
}

type ParameterGroupParameterArgs struct {
	// The apply method of the Neptune parameter. Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
	ApplyMethod pulumi.StringPtrInput `pulumi:"applyMethod"`
	// The name of the Neptune parameter.
	Name pulumi.StringInput `pulumi:"name"`
	// The value of the Neptune parameter.
	Value pulumi.StringInput `pulumi:"value"`
}

func (ParameterGroupParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterGroupParameter)(nil)).Elem()
}

func (i ParameterGroupParameterArgs) ToParameterGroupParameterOutput() ParameterGroupParameterOutput {
	return i.ToParameterGroupParameterOutputWithContext(context.Background())
}

func (i ParameterGroupParameterArgs) ToParameterGroupParameterOutputWithContext(ctx context.Context) ParameterGroupParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupParameterOutput)
}

// ParameterGroupParameterArrayInput is an input type that accepts ParameterGroupParameterArray and ParameterGroupParameterArrayOutput values.
// You can construct a concrete instance of `ParameterGroupParameterArrayInput` via:
//
//	ParameterGroupParameterArray{ ParameterGroupParameterArgs{...} }
type ParameterGroupParameterArrayInput interface {
	pulumi.Input

	ToParameterGroupParameterArrayOutput() ParameterGroupParameterArrayOutput
	ToParameterGroupParameterArrayOutputWithContext(context.Context) ParameterGroupParameterArrayOutput
}

type ParameterGroupParameterArray []ParameterGroupParameterInput

func (ParameterGroupParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterGroupParameter)(nil)).Elem()
}

func (i ParameterGroupParameterArray) ToParameterGroupParameterArrayOutput() ParameterGroupParameterArrayOutput {
	return i.ToParameterGroupParameterArrayOutputWithContext(context.Background())
}

func (i ParameterGroupParameterArray) ToParameterGroupParameterArrayOutputWithContext(ctx context.Context) ParameterGroupParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupParameterArrayOutput)
}

type ParameterGroupParameterOutput struct{ *pulumi.OutputState }

func (ParameterGroupParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterGroupParameter)(nil)).Elem()
}

func (o ParameterGroupParameterOutput) ToParameterGroupParameterOutput() ParameterGroupParameterOutput {
	return o
}

func (o ParameterGroupParameterOutput) ToParameterGroupParameterOutputWithContext(ctx context.Context) ParameterGroupParameterOutput {
	return o
}

// The apply method of the Neptune parameter. Valid values are `immediate` and `pending-reboot`. Defaults to `pending-reboot`.
func (o ParameterGroupParameterOutput) ApplyMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParameterGroupParameter) *string { return v.ApplyMethod }).(pulumi.StringPtrOutput)
}

// The name of the Neptune parameter.
func (o ParameterGroupParameterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterGroupParameter) string { return v.Name }).(pulumi.StringOutput)
}

// The value of the Neptune parameter.
func (o ParameterGroupParameterOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterGroupParameter) string { return v.Value }).(pulumi.StringOutput)
}

type ParameterGroupParameterArrayOutput struct{ *pulumi.OutputState }

func (ParameterGroupParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterGroupParameter)(nil)).Elem()
}

func (o ParameterGroupParameterArrayOutput) ToParameterGroupParameterArrayOutput() ParameterGroupParameterArrayOutput {
	return o
}

func (o ParameterGroupParameterArrayOutput) ToParameterGroupParameterArrayOutputWithContext(ctx context.Context) ParameterGroupParameterArrayOutput {
	return o
}

func (o ParameterGroupParameterArrayOutput) Index(i pulumi.IntInput) ParameterGroupParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ParameterGroupParameter {
		return vs[0].([]ParameterGroupParameter)[vs[1].(int)]
	}).(ParameterGroupParameterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterParameterGroupParameterInput)(nil)).Elem(), ClusterParameterGroupParameterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterParameterGroupParameterArrayInput)(nil)).Elem(), ClusterParameterGroupParameterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterServerlessV2ScalingConfigurationInput)(nil)).Elem(), ClusterServerlessV2ScalingConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterServerlessV2ScalingConfigurationPtrInput)(nil)).Elem(), ClusterServerlessV2ScalingConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalClusterGlobalClusterMemberInput)(nil)).Elem(), GlobalClusterGlobalClusterMemberArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalClusterGlobalClusterMemberArrayInput)(nil)).Elem(), GlobalClusterGlobalClusterMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterGroupParameterInput)(nil)).Elem(), ParameterGroupParameterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterGroupParameterArrayInput)(nil)).Elem(), ParameterGroupParameterArray{})
	pulumi.RegisterOutputType(ClusterParameterGroupParameterOutput{})
	pulumi.RegisterOutputType(ClusterParameterGroupParameterArrayOutput{})
	pulumi.RegisterOutputType(ClusterServerlessV2ScalingConfigurationOutput{})
	pulumi.RegisterOutputType(ClusterServerlessV2ScalingConfigurationPtrOutput{})
	pulumi.RegisterOutputType(GlobalClusterGlobalClusterMemberOutput{})
	pulumi.RegisterOutputType(GlobalClusterGlobalClusterMemberArrayOutput{})
	pulumi.RegisterOutputType(ParameterGroupParameterOutput{})
	pulumi.RegisterOutputType(ParameterGroupParameterArrayOutput{})
}
