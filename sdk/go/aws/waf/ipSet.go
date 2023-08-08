// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a WAF IPSet Resource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/waf"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := waf.NewIpSet(ctx, "ipset", &waf.IpSetArgs{
//				IpSetDescriptors: waf.IpSetIpSetDescriptorArray{
//					&waf.IpSetIpSetDescriptorArgs{
//						Type:  pulumi.String("IPV4"),
//						Value: pulumi.String("192.0.7.0/24"),
//					},
//					&waf.IpSetIpSetDescriptorArgs{
//						Type:  pulumi.String("IPV4"),
//						Value: pulumi.String("10.16.16.0/16"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// terraform import {
//
//	to = aws_waf_ipset.example
//
//	id = "a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc" } Using `pulumi import`, import WAF IPSets using their ID. For exampleconsole % pulumi import aws_waf_ipset.example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
type IpSet struct {
	pulumi.CustomResourceState

	// The ARN of the WAF IPSet.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR format) from which web requests originate.
	IpSetDescriptors IpSetIpSetDescriptorArrayOutput `pulumi:"ipSetDescriptors"`
	// The name or description of the IPSet.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewIpSet registers a new resource with the given unique name, arguments, and options.
func NewIpSet(ctx *pulumi.Context,
	name string, args *IpSetArgs, opts ...pulumi.ResourceOption) (*IpSet, error) {
	if args == nil {
		args = &IpSetArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IpSet
	err := ctx.RegisterResource("aws:waf/ipSet:IpSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpSet gets an existing IpSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpSetState, opts ...pulumi.ResourceOption) (*IpSet, error) {
	var resource IpSet
	err := ctx.ReadResource("aws:waf/ipSet:IpSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpSet resources.
type ipSetState struct {
	// The ARN of the WAF IPSet.
	Arn *string `pulumi:"arn"`
	// One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR format) from which web requests originate.
	IpSetDescriptors []IpSetIpSetDescriptor `pulumi:"ipSetDescriptors"`
	// The name or description of the IPSet.
	Name *string `pulumi:"name"`
}

type IpSetState struct {
	// The ARN of the WAF IPSet.
	Arn pulumi.StringPtrInput
	// One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR format) from which web requests originate.
	IpSetDescriptors IpSetIpSetDescriptorArrayInput
	// The name or description of the IPSet.
	Name pulumi.StringPtrInput
}

func (IpSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipSetState)(nil)).Elem()
}

type ipSetArgs struct {
	// One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR format) from which web requests originate.
	IpSetDescriptors []IpSetIpSetDescriptor `pulumi:"ipSetDescriptors"`
	// The name or description of the IPSet.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a IpSet resource.
type IpSetArgs struct {
	// One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR format) from which web requests originate.
	IpSetDescriptors IpSetIpSetDescriptorArrayInput
	// The name or description of the IPSet.
	Name pulumi.StringPtrInput
}

func (IpSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipSetArgs)(nil)).Elem()
}

type IpSetInput interface {
	pulumi.Input

	ToIpSetOutput() IpSetOutput
	ToIpSetOutputWithContext(ctx context.Context) IpSetOutput
}

func (*IpSet) ElementType() reflect.Type {
	return reflect.TypeOf((**IpSet)(nil)).Elem()
}

func (i *IpSet) ToIpSetOutput() IpSetOutput {
	return i.ToIpSetOutputWithContext(context.Background())
}

func (i *IpSet) ToIpSetOutputWithContext(ctx context.Context) IpSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSetOutput)
}

// IpSetArrayInput is an input type that accepts IpSetArray and IpSetArrayOutput values.
// You can construct a concrete instance of `IpSetArrayInput` via:
//
//	IpSetArray{ IpSetArgs{...} }
type IpSetArrayInput interface {
	pulumi.Input

	ToIpSetArrayOutput() IpSetArrayOutput
	ToIpSetArrayOutputWithContext(context.Context) IpSetArrayOutput
}

type IpSetArray []IpSetInput

func (IpSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpSet)(nil)).Elem()
}

func (i IpSetArray) ToIpSetArrayOutput() IpSetArrayOutput {
	return i.ToIpSetArrayOutputWithContext(context.Background())
}

func (i IpSetArray) ToIpSetArrayOutputWithContext(ctx context.Context) IpSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSetArrayOutput)
}

// IpSetMapInput is an input type that accepts IpSetMap and IpSetMapOutput values.
// You can construct a concrete instance of `IpSetMapInput` via:
//
//	IpSetMap{ "key": IpSetArgs{...} }
type IpSetMapInput interface {
	pulumi.Input

	ToIpSetMapOutput() IpSetMapOutput
	ToIpSetMapOutputWithContext(context.Context) IpSetMapOutput
}

type IpSetMap map[string]IpSetInput

func (IpSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpSet)(nil)).Elem()
}

func (i IpSetMap) ToIpSetMapOutput() IpSetMapOutput {
	return i.ToIpSetMapOutputWithContext(context.Background())
}

func (i IpSetMap) ToIpSetMapOutputWithContext(ctx context.Context) IpSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSetMapOutput)
}

type IpSetOutput struct{ *pulumi.OutputState }

func (IpSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpSet)(nil)).Elem()
}

func (o IpSetOutput) ToIpSetOutput() IpSetOutput {
	return o
}

func (o IpSetOutput) ToIpSetOutputWithContext(ctx context.Context) IpSetOutput {
	return o
}

// The ARN of the WAF IPSet.
func (o IpSetOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSet) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR format) from which web requests originate.
func (o IpSetOutput) IpSetDescriptors() IpSetIpSetDescriptorArrayOutput {
	return o.ApplyT(func(v *IpSet) IpSetIpSetDescriptorArrayOutput { return v.IpSetDescriptors }).(IpSetIpSetDescriptorArrayOutput)
}

// The name or description of the IPSet.
func (o IpSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type IpSetArrayOutput struct{ *pulumi.OutputState }

func (IpSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpSet)(nil)).Elem()
}

func (o IpSetArrayOutput) ToIpSetArrayOutput() IpSetArrayOutput {
	return o
}

func (o IpSetArrayOutput) ToIpSetArrayOutputWithContext(ctx context.Context) IpSetArrayOutput {
	return o
}

func (o IpSetArrayOutput) Index(i pulumi.IntInput) IpSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpSet {
		return vs[0].([]*IpSet)[vs[1].(int)]
	}).(IpSetOutput)
}

type IpSetMapOutput struct{ *pulumi.OutputState }

func (IpSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpSet)(nil)).Elem()
}

func (o IpSetMapOutput) ToIpSetMapOutput() IpSetMapOutput {
	return o
}

func (o IpSetMapOutput) ToIpSetMapOutputWithContext(ctx context.Context) IpSetMapOutput {
	return o
}

func (o IpSetMapOutput) MapIndex(k pulumi.StringInput) IpSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpSet {
		return vs[0].(map[string]*IpSet)[vs[1].(string)]
	}).(IpSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpSetInput)(nil)).Elem(), &IpSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpSetArrayInput)(nil)).Elem(), IpSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpSetMapInput)(nil)).Elem(), IpSetMap{})
	pulumi.RegisterOutputType(IpSetOutput{})
	pulumi.RegisterOutputType(IpSetArrayOutput{})
	pulumi.RegisterOutputType(IpSetMapOutput{})
}
