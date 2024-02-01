// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package applicationloadbalancing

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IpAddressType string

const (
	IpAddressTypeIpv4      = IpAddressType("ipv4")
	IpAddressTypeDualstack = IpAddressType("dualstack")
)

func (IpAddressType) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddressType)(nil)).Elem()
}

func (e IpAddressType) ToIpAddressTypeOutput() IpAddressTypeOutput {
	return pulumi.ToOutput(e).(IpAddressTypeOutput)
}

func (e IpAddressType) ToIpAddressTypeOutputWithContext(ctx context.Context) IpAddressTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IpAddressTypeOutput)
}

func (e IpAddressType) ToIpAddressTypePtrOutput() IpAddressTypePtrOutput {
	return e.ToIpAddressTypePtrOutputWithContext(context.Background())
}

func (e IpAddressType) ToIpAddressTypePtrOutputWithContext(ctx context.Context) IpAddressTypePtrOutput {
	return IpAddressType(e).ToIpAddressTypeOutputWithContext(ctx).ToIpAddressTypePtrOutputWithContext(ctx)
}

func (e IpAddressType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpAddressType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpAddressType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IpAddressType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IpAddressTypeOutput struct{ *pulumi.OutputState }

func (IpAddressTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddressType)(nil)).Elem()
}

func (o IpAddressTypeOutput) ToIpAddressTypeOutput() IpAddressTypeOutput {
	return o
}

func (o IpAddressTypeOutput) ToIpAddressTypeOutputWithContext(ctx context.Context) IpAddressTypeOutput {
	return o
}

func (o IpAddressTypeOutput) ToIpAddressTypePtrOutput() IpAddressTypePtrOutput {
	return o.ToIpAddressTypePtrOutputWithContext(context.Background())
}

func (o IpAddressTypeOutput) ToIpAddressTypePtrOutputWithContext(ctx context.Context) IpAddressTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IpAddressType) *IpAddressType {
		return &v
	}).(IpAddressTypePtrOutput)
}

func (o IpAddressTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IpAddressTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpAddressType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IpAddressTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpAddressTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpAddressType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IpAddressTypePtrOutput struct{ *pulumi.OutputState }

func (IpAddressTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpAddressType)(nil)).Elem()
}

func (o IpAddressTypePtrOutput) ToIpAddressTypePtrOutput() IpAddressTypePtrOutput {
	return o
}

func (o IpAddressTypePtrOutput) ToIpAddressTypePtrOutputWithContext(ctx context.Context) IpAddressTypePtrOutput {
	return o
}

func (o IpAddressTypePtrOutput) Elem() IpAddressTypeOutput {
	return o.ApplyT(func(v *IpAddressType) IpAddressType {
		if v != nil {
			return *v
		}
		var ret IpAddressType
		return ret
	}).(IpAddressTypeOutput)
}

func (o IpAddressTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpAddressTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IpAddressType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// IpAddressTypeInput is an input type that accepts values of the IpAddressType enum
// A concrete instance of `IpAddressTypeInput` can be one of the following:
//
//	IpAddressTypeIpv4
//	IpAddressTypeDualstack
type IpAddressTypeInput interface {
	pulumi.Input

	ToIpAddressTypeOutput() IpAddressTypeOutput
	ToIpAddressTypeOutputWithContext(context.Context) IpAddressTypeOutput
}

var ipAddressTypePtrType = reflect.TypeOf((**IpAddressType)(nil)).Elem()

type IpAddressTypePtrInput interface {
	pulumi.Input

	ToIpAddressTypePtrOutput() IpAddressTypePtrOutput
	ToIpAddressTypePtrOutputWithContext(context.Context) IpAddressTypePtrOutput
}

type ipAddressTypePtr string

func IpAddressTypePtr(v string) IpAddressTypePtrInput {
	return (*ipAddressTypePtr)(&v)
}

func (*ipAddressTypePtr) ElementType() reflect.Type {
	return ipAddressTypePtrType
}

func (in *ipAddressTypePtr) ToIpAddressTypePtrOutput() IpAddressTypePtrOutput {
	return pulumi.ToOutput(in).(IpAddressTypePtrOutput)
}

func (in *ipAddressTypePtr) ToIpAddressTypePtrOutputWithContext(ctx context.Context) IpAddressTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IpAddressTypePtrOutput)
}

type LoadBalancerType string

const (
	LoadBalancerTypeApplication = LoadBalancerType("application")
	LoadBalancerTypeNetwork     = LoadBalancerType("network")
)

func (LoadBalancerType) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerType)(nil)).Elem()
}

func (e LoadBalancerType) ToLoadBalancerTypeOutput() LoadBalancerTypeOutput {
	return pulumi.ToOutput(e).(LoadBalancerTypeOutput)
}

func (e LoadBalancerType) ToLoadBalancerTypeOutputWithContext(ctx context.Context) LoadBalancerTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LoadBalancerTypeOutput)
}

func (e LoadBalancerType) ToLoadBalancerTypePtrOutput() LoadBalancerTypePtrOutput {
	return e.ToLoadBalancerTypePtrOutputWithContext(context.Background())
}

func (e LoadBalancerType) ToLoadBalancerTypePtrOutputWithContext(ctx context.Context) LoadBalancerTypePtrOutput {
	return LoadBalancerType(e).ToLoadBalancerTypeOutputWithContext(ctx).ToLoadBalancerTypePtrOutputWithContext(ctx)
}

func (e LoadBalancerType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LoadBalancerType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LoadBalancerType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LoadBalancerType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LoadBalancerTypeOutput struct{ *pulumi.OutputState }

func (LoadBalancerTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerType)(nil)).Elem()
}

func (o LoadBalancerTypeOutput) ToLoadBalancerTypeOutput() LoadBalancerTypeOutput {
	return o
}

func (o LoadBalancerTypeOutput) ToLoadBalancerTypeOutputWithContext(ctx context.Context) LoadBalancerTypeOutput {
	return o
}

func (o LoadBalancerTypeOutput) ToLoadBalancerTypePtrOutput() LoadBalancerTypePtrOutput {
	return o.ToLoadBalancerTypePtrOutputWithContext(context.Background())
}

func (o LoadBalancerTypeOutput) ToLoadBalancerTypePtrOutputWithContext(ctx context.Context) LoadBalancerTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoadBalancerType) *LoadBalancerType {
		return &v
	}).(LoadBalancerTypePtrOutput)
}

func (o LoadBalancerTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LoadBalancerTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LoadBalancerType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LoadBalancerTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LoadBalancerTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LoadBalancerType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LoadBalancerTypePtrOutput struct{ *pulumi.OutputState }

func (LoadBalancerTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerType)(nil)).Elem()
}

func (o LoadBalancerTypePtrOutput) ToLoadBalancerTypePtrOutput() LoadBalancerTypePtrOutput {
	return o
}

func (o LoadBalancerTypePtrOutput) ToLoadBalancerTypePtrOutputWithContext(ctx context.Context) LoadBalancerTypePtrOutput {
	return o
}

func (o LoadBalancerTypePtrOutput) Elem() LoadBalancerTypeOutput {
	return o.ApplyT(func(v *LoadBalancerType) LoadBalancerType {
		if v != nil {
			return *v
		}
		var ret LoadBalancerType
		return ret
	}).(LoadBalancerTypeOutput)
}

func (o LoadBalancerTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LoadBalancerTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LoadBalancerType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// LoadBalancerTypeInput is an input type that accepts values of the LoadBalancerType enum
// A concrete instance of `LoadBalancerTypeInput` can be one of the following:
//
//	LoadBalancerTypeApplication
//	LoadBalancerTypeNetwork
type LoadBalancerTypeInput interface {
	pulumi.Input

	ToLoadBalancerTypeOutput() LoadBalancerTypeOutput
	ToLoadBalancerTypeOutputWithContext(context.Context) LoadBalancerTypeOutput
}

var loadBalancerTypePtrType = reflect.TypeOf((**LoadBalancerType)(nil)).Elem()

type LoadBalancerTypePtrInput interface {
	pulumi.Input

	ToLoadBalancerTypePtrOutput() LoadBalancerTypePtrOutput
	ToLoadBalancerTypePtrOutputWithContext(context.Context) LoadBalancerTypePtrOutput
}

type loadBalancerTypePtr string

func LoadBalancerTypePtr(v string) LoadBalancerTypePtrInput {
	return (*loadBalancerTypePtr)(&v)
}

func (*loadBalancerTypePtr) ElementType() reflect.Type {
	return loadBalancerTypePtrType
}

func (in *loadBalancerTypePtr) ToLoadBalancerTypePtrOutput() LoadBalancerTypePtrOutput {
	return pulumi.ToOutput(in).(LoadBalancerTypePtrOutput)
}

func (in *loadBalancerTypePtr) ToLoadBalancerTypePtrOutputWithContext(ctx context.Context) LoadBalancerTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LoadBalancerTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpAddressTypeInput)(nil)).Elem(), IpAddressType("ipv4"))
	pulumi.RegisterInputType(reflect.TypeOf((*IpAddressTypePtrInput)(nil)).Elem(), IpAddressType("ipv4"))
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerTypeInput)(nil)).Elem(), LoadBalancerType("application"))
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerTypePtrInput)(nil)).Elem(), LoadBalancerType("application"))
	pulumi.RegisterOutputType(IpAddressTypeOutput{})
	pulumi.RegisterOutputType(IpAddressTypePtrOutput{})
	pulumi.RegisterOutputType(LoadBalancerTypeOutput{})
	pulumi.RegisterOutputType(LoadBalancerTypePtrOutput{})
}
