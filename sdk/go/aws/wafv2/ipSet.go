// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a WAFv2 IP Set Resource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/wafv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := wafv2.NewIpSet(ctx, "example", &wafv2.IpSetArgs{
//				Addresses: pulumi.StringArray{
//					pulumi.String("1.2.3.4/32"),
//					pulumi.String("5.6.7.8/32"),
//				},
//				Description:      pulumi.String("Example IP set"),
//				IpAddressVersion: pulumi.String("IPV4"),
//				Scope:            pulumi.String("REGIONAL"),
//				Tags: pulumi.StringMap{
//					"Tag1": pulumi.String("Value1"),
//					"Tag2": pulumi.String("Value2"),
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
//	to = aws_wafv2_ip_set.example
//
//	id = "a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc/example/REGIONAL" } Using `pulumi import`, import WAFv2 IP Sets using `ID/name/scope`. For exampleconsole % pulumi import aws_wafv2_ip_set.example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc/example/REGIONAL
type IpSet struct {
	pulumi.CustomResourceState

	// Contains an array of strings that specifies zero or more IP addresses or blocks of IP addresses. All addresses must be specified using Classless Inter-Domain Routing (CIDR) notation. WAF supports all IPv4 and IPv6 CIDR ranges except for `/0`.
	Addresses pulumi.StringArrayOutput `pulumi:"addresses"`
	// The Amazon Resource Name (ARN) of the IP set.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A friendly description of the IP set.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specify IPV4 or IPV6. Valid values are `IPV4` or `IPV6`.
	IpAddressVersion pulumi.StringOutput `pulumi:"ipAddressVersion"`
	LockToken        pulumi.StringOutput `pulumi:"lockToken"`
	// A friendly name of the IP set.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the Region US East (N. Virginia).
	Scope pulumi.StringOutput `pulumi:"scope"`
	// An array of key:value pairs to associate with the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewIpSet registers a new resource with the given unique name, arguments, and options.
func NewIpSet(ctx *pulumi.Context,
	name string, args *IpSetArgs, opts ...pulumi.ResourceOption) (*IpSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpAddressVersion == nil {
		return nil, errors.New("invalid value for required argument 'IpAddressVersion'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IpSet
	err := ctx.RegisterResource("aws:wafv2/ipSet:IpSet", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:wafv2/ipSet:IpSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpSet resources.
type ipSetState struct {
	// Contains an array of strings that specifies zero or more IP addresses or blocks of IP addresses. All addresses must be specified using Classless Inter-Domain Routing (CIDR) notation. WAF supports all IPv4 and IPv6 CIDR ranges except for `/0`.
	Addresses []string `pulumi:"addresses"`
	// The Amazon Resource Name (ARN) of the IP set.
	Arn *string `pulumi:"arn"`
	// A friendly description of the IP set.
	Description *string `pulumi:"description"`
	// Specify IPV4 or IPV6. Valid values are `IPV4` or `IPV6`.
	IpAddressVersion *string `pulumi:"ipAddressVersion"`
	LockToken        *string `pulumi:"lockToken"`
	// A friendly name of the IP set.
	Name *string `pulumi:"name"`
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the Region US East (N. Virginia).
	Scope *string `pulumi:"scope"`
	// An array of key:value pairs to associate with the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type IpSetState struct {
	// Contains an array of strings that specifies zero or more IP addresses or blocks of IP addresses. All addresses must be specified using Classless Inter-Domain Routing (CIDR) notation. WAF supports all IPv4 and IPv6 CIDR ranges except for `/0`.
	Addresses pulumi.StringArrayInput
	// The Amazon Resource Name (ARN) of the IP set.
	Arn pulumi.StringPtrInput
	// A friendly description of the IP set.
	Description pulumi.StringPtrInput
	// Specify IPV4 or IPV6. Valid values are `IPV4` or `IPV6`.
	IpAddressVersion pulumi.StringPtrInput
	LockToken        pulumi.StringPtrInput
	// A friendly name of the IP set.
	Name pulumi.StringPtrInput
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the Region US East (N. Virginia).
	Scope pulumi.StringPtrInput
	// An array of key:value pairs to associate with the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (IpSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipSetState)(nil)).Elem()
}

type ipSetArgs struct {
	// Contains an array of strings that specifies zero or more IP addresses or blocks of IP addresses. All addresses must be specified using Classless Inter-Domain Routing (CIDR) notation. WAF supports all IPv4 and IPv6 CIDR ranges except for `/0`.
	Addresses []string `pulumi:"addresses"`
	// A friendly description of the IP set.
	Description *string `pulumi:"description"`
	// Specify IPV4 or IPV6. Valid values are `IPV4` or `IPV6`.
	IpAddressVersion string `pulumi:"ipAddressVersion"`
	// A friendly name of the IP set.
	Name *string `pulumi:"name"`
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the Region US East (N. Virginia).
	Scope string `pulumi:"scope"`
	// An array of key:value pairs to associate with the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IpSet resource.
type IpSetArgs struct {
	// Contains an array of strings that specifies zero or more IP addresses or blocks of IP addresses. All addresses must be specified using Classless Inter-Domain Routing (CIDR) notation. WAF supports all IPv4 and IPv6 CIDR ranges except for `/0`.
	Addresses pulumi.StringArrayInput
	// A friendly description of the IP set.
	Description pulumi.StringPtrInput
	// Specify IPV4 or IPV6. Valid values are `IPV4` or `IPV6`.
	IpAddressVersion pulumi.StringInput
	// A friendly name of the IP set.
	Name pulumi.StringPtrInput
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the Region US East (N. Virginia).
	Scope pulumi.StringInput
	// An array of key:value pairs to associate with the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
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

// Contains an array of strings that specifies zero or more IP addresses or blocks of IP addresses. All addresses must be specified using Classless Inter-Domain Routing (CIDR) notation. WAF supports all IPv4 and IPv6 CIDR ranges except for `/0`.
func (o IpSetOutput) Addresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IpSet) pulumi.StringArrayOutput { return v.Addresses }).(pulumi.StringArrayOutput)
}

// The Amazon Resource Name (ARN) of the IP set.
func (o IpSetOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSet) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// A friendly description of the IP set.
func (o IpSetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpSet) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specify IPV4 or IPV6. Valid values are `IPV4` or `IPV6`.
func (o IpSetOutput) IpAddressVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSet) pulumi.StringOutput { return v.IpAddressVersion }).(pulumi.StringOutput)
}

func (o IpSetOutput) LockToken() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSet) pulumi.StringOutput { return v.LockToken }).(pulumi.StringOutput)
}

// A friendly name of the IP set.
func (o IpSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the Region US East (N. Virginia).
func (o IpSetOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSet) pulumi.StringOutput { return v.Scope }).(pulumi.StringOutput)
}

// An array of key:value pairs to associate with the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o IpSetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IpSet) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o IpSetOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IpSet) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
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
