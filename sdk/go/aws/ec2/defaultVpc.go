// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage the [default AWS VPC](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/default-vpc.html)
// in the current region.
//
// For AWS accounts created after 2013-12-04, each region comes with a Default VPC.
// **This is an advanced resource**, and has special caveats to be aware of when
// using it. Please read this document in its entirety before using this resource.
//
// The `ec2.DefaultVpc` behaves differently from normal resources, in that
// this provider does not _create_ this resource, but instead "adopts" it
// into management.
//
// ## Example Usage
//
// Basic usage with tags:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2.NewDefaultVpc(ctx, "_default", &ec2.DefaultVpcArgs{
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("Default VPC"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Default VPCs can be imported using the `vpc id`, e.g.
//
// ```sh
//  $ pulumi import aws:ec2/defaultVpc:DefaultVpc default vpc-a01106c2
// ```
type DefaultVpc struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of VPC
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Whether or not an Amazon-provided IPv6 CIDR
	// block with a /56 prefix length for the VPC was assigned
	AssignGeneratedIpv6CidrBlock pulumi.BoolOutput `pulumi:"assignGeneratedIpv6CidrBlock"`
	// The CIDR block of the VPC
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// The ID of the network ACL created by default on VPC creation
	DefaultNetworkAclId pulumi.StringOutput `pulumi:"defaultNetworkAclId"`
	// The ID of the route table created by default on VPC creation
	DefaultRouteTableId pulumi.StringOutput `pulumi:"defaultRouteTableId"`
	// The ID of the security group created by default on VPC creation
	DefaultSecurityGroupId pulumi.StringOutput `pulumi:"defaultSecurityGroupId"`
	DhcpOptionsId          pulumi.StringOutput `pulumi:"dhcpOptionsId"`
	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
	EnableClassiclink           pulumi.BoolOutput `pulumi:"enableClassiclink"`
	EnableClassiclinkDnsSupport pulumi.BoolOutput `pulumi:"enableClassiclinkDnsSupport"`
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames pulumi.BoolOutput `pulumi:"enableDnsHostnames"`
	// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
	EnableDnsSupport pulumi.BoolPtrOutput `pulumi:"enableDnsSupport"`
	// Tenancy of instances spin up within VPC.
	InstanceTenancy pulumi.StringOutput `pulumi:"instanceTenancy"`
	// The association ID for the IPv6 CIDR block of the VPC
	Ipv6AssociationId pulumi.StringOutput `pulumi:"ipv6AssociationId"`
	// The IPv6 CIDR block of the VPC
	Ipv6CidrBlock pulumi.StringOutput `pulumi:"ipv6CidrBlock"`
	// The ID of the main route table associated with
	// this VPC. Note that you can change a VPC's main route table by using an
	// `ec2.MainRouteTableAssociation`
	MainRouteTableId pulumi.StringOutput `pulumi:"mainRouteTableId"`
	// The ID of the AWS account that owns the VPC.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// A map of tags to assign to the resource.
	Tags    pulumi.StringMapOutput `pulumi:"tags"`
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewDefaultVpc registers a new resource with the given unique name, arguments, and options.
func NewDefaultVpc(ctx *pulumi.Context,
	name string, args *DefaultVpcArgs, opts ...pulumi.ResourceOption) (*DefaultVpc, error) {
	if args == nil {
		args = &DefaultVpcArgs{}
	}

	var resource DefaultVpc
	err := ctx.RegisterResource("aws:ec2/defaultVpc:DefaultVpc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultVpc gets an existing DefaultVpc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultVpc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultVpcState, opts ...pulumi.ResourceOption) (*DefaultVpc, error) {
	var resource DefaultVpc
	err := ctx.ReadResource("aws:ec2/defaultVpc:DefaultVpc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultVpc resources.
type defaultVpcState struct {
	// Amazon Resource Name (ARN) of VPC
	Arn *string `pulumi:"arn"`
	// Whether or not an Amazon-provided IPv6 CIDR
	// block with a /56 prefix length for the VPC was assigned
	AssignGeneratedIpv6CidrBlock *bool `pulumi:"assignGeneratedIpv6CidrBlock"`
	// The CIDR block of the VPC
	CidrBlock *string `pulumi:"cidrBlock"`
	// The ID of the network ACL created by default on VPC creation
	DefaultNetworkAclId *string `pulumi:"defaultNetworkAclId"`
	// The ID of the route table created by default on VPC creation
	DefaultRouteTableId *string `pulumi:"defaultRouteTableId"`
	// The ID of the security group created by default on VPC creation
	DefaultSecurityGroupId *string `pulumi:"defaultSecurityGroupId"`
	DhcpOptionsId          *string `pulumi:"dhcpOptionsId"`
	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
	EnableClassiclink           *bool `pulumi:"enableClassiclink"`
	EnableClassiclinkDnsSupport *bool `pulumi:"enableClassiclinkDnsSupport"`
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames *bool `pulumi:"enableDnsHostnames"`
	// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
	EnableDnsSupport *bool `pulumi:"enableDnsSupport"`
	// Tenancy of instances spin up within VPC.
	InstanceTenancy *string `pulumi:"instanceTenancy"`
	// The association ID for the IPv6 CIDR block of the VPC
	Ipv6AssociationId *string `pulumi:"ipv6AssociationId"`
	// The IPv6 CIDR block of the VPC
	Ipv6CidrBlock *string `pulumi:"ipv6CidrBlock"`
	// The ID of the main route table associated with
	// this VPC. Note that you can change a VPC's main route table by using an
	// `ec2.MainRouteTableAssociation`
	MainRouteTableId *string `pulumi:"mainRouteTableId"`
	// The ID of the AWS account that owns the VPC.
	OwnerId *string `pulumi:"ownerId"`
	// A map of tags to assign to the resource.
	Tags    map[string]string `pulumi:"tags"`
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type DefaultVpcState struct {
	// Amazon Resource Name (ARN) of VPC
	Arn pulumi.StringPtrInput
	// Whether or not an Amazon-provided IPv6 CIDR
	// block with a /56 prefix length for the VPC was assigned
	AssignGeneratedIpv6CidrBlock pulumi.BoolPtrInput
	// The CIDR block of the VPC
	CidrBlock pulumi.StringPtrInput
	// The ID of the network ACL created by default on VPC creation
	DefaultNetworkAclId pulumi.StringPtrInput
	// The ID of the route table created by default on VPC creation
	DefaultRouteTableId pulumi.StringPtrInput
	// The ID of the security group created by default on VPC creation
	DefaultSecurityGroupId pulumi.StringPtrInput
	DhcpOptionsId          pulumi.StringPtrInput
	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
	EnableClassiclink           pulumi.BoolPtrInput
	EnableClassiclinkDnsSupport pulumi.BoolPtrInput
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames pulumi.BoolPtrInput
	// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
	EnableDnsSupport pulumi.BoolPtrInput
	// Tenancy of instances spin up within VPC.
	InstanceTenancy pulumi.StringPtrInput
	// The association ID for the IPv6 CIDR block of the VPC
	Ipv6AssociationId pulumi.StringPtrInput
	// The IPv6 CIDR block of the VPC
	Ipv6CidrBlock pulumi.StringPtrInput
	// The ID of the main route table associated with
	// this VPC. Note that you can change a VPC's main route table by using an
	// `ec2.MainRouteTableAssociation`
	MainRouteTableId pulumi.StringPtrInput
	// The ID of the AWS account that owns the VPC.
	OwnerId pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags    pulumi.StringMapInput
	TagsAll pulumi.StringMapInput
}

func (DefaultVpcState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultVpcState)(nil)).Elem()
}

type defaultVpcArgs struct {
	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
	EnableClassiclink           *bool `pulumi:"enableClassiclink"`
	EnableClassiclinkDnsSupport *bool `pulumi:"enableClassiclinkDnsSupport"`
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames *bool `pulumi:"enableDnsHostnames"`
	// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
	EnableDnsSupport *bool `pulumi:"enableDnsSupport"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DefaultVpc resource.
type DefaultVpcArgs struct {
	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
	EnableClassiclink           pulumi.BoolPtrInput
	EnableClassiclinkDnsSupport pulumi.BoolPtrInput
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames pulumi.BoolPtrInput
	// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
	EnableDnsSupport pulumi.BoolPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (DefaultVpcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultVpcArgs)(nil)).Elem()
}

type DefaultVpcInput interface {
	pulumi.Input

	ToDefaultVpcOutput() DefaultVpcOutput
	ToDefaultVpcOutputWithContext(ctx context.Context) DefaultVpcOutput
}

func (*DefaultVpc) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultVpc)(nil))
}

func (i *DefaultVpc) ToDefaultVpcOutput() DefaultVpcOutput {
	return i.ToDefaultVpcOutputWithContext(context.Background())
}

func (i *DefaultVpc) ToDefaultVpcOutputWithContext(ctx context.Context) DefaultVpcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultVpcOutput)
}

func (i *DefaultVpc) ToDefaultVpcPtrOutput() DefaultVpcPtrOutput {
	return i.ToDefaultVpcPtrOutputWithContext(context.Background())
}

func (i *DefaultVpc) ToDefaultVpcPtrOutputWithContext(ctx context.Context) DefaultVpcPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultVpcPtrOutput)
}

type DefaultVpcPtrInput interface {
	pulumi.Input

	ToDefaultVpcPtrOutput() DefaultVpcPtrOutput
	ToDefaultVpcPtrOutputWithContext(ctx context.Context) DefaultVpcPtrOutput
}

type defaultVpcPtrType DefaultVpcArgs

func (*defaultVpcPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultVpc)(nil))
}

func (i *defaultVpcPtrType) ToDefaultVpcPtrOutput() DefaultVpcPtrOutput {
	return i.ToDefaultVpcPtrOutputWithContext(context.Background())
}

func (i *defaultVpcPtrType) ToDefaultVpcPtrOutputWithContext(ctx context.Context) DefaultVpcPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultVpcPtrOutput)
}

// DefaultVpcArrayInput is an input type that accepts DefaultVpcArray and DefaultVpcArrayOutput values.
// You can construct a concrete instance of `DefaultVpcArrayInput` via:
//
//          DefaultVpcArray{ DefaultVpcArgs{...} }
type DefaultVpcArrayInput interface {
	pulumi.Input

	ToDefaultVpcArrayOutput() DefaultVpcArrayOutput
	ToDefaultVpcArrayOutputWithContext(context.Context) DefaultVpcArrayOutput
}

type DefaultVpcArray []DefaultVpcInput

func (DefaultVpcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultVpc)(nil)).Elem()
}

func (i DefaultVpcArray) ToDefaultVpcArrayOutput() DefaultVpcArrayOutput {
	return i.ToDefaultVpcArrayOutputWithContext(context.Background())
}

func (i DefaultVpcArray) ToDefaultVpcArrayOutputWithContext(ctx context.Context) DefaultVpcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultVpcArrayOutput)
}

// DefaultVpcMapInput is an input type that accepts DefaultVpcMap and DefaultVpcMapOutput values.
// You can construct a concrete instance of `DefaultVpcMapInput` via:
//
//          DefaultVpcMap{ "key": DefaultVpcArgs{...} }
type DefaultVpcMapInput interface {
	pulumi.Input

	ToDefaultVpcMapOutput() DefaultVpcMapOutput
	ToDefaultVpcMapOutputWithContext(context.Context) DefaultVpcMapOutput
}

type DefaultVpcMap map[string]DefaultVpcInput

func (DefaultVpcMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultVpc)(nil)).Elem()
}

func (i DefaultVpcMap) ToDefaultVpcMapOutput() DefaultVpcMapOutput {
	return i.ToDefaultVpcMapOutputWithContext(context.Background())
}

func (i DefaultVpcMap) ToDefaultVpcMapOutputWithContext(ctx context.Context) DefaultVpcMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultVpcMapOutput)
}

type DefaultVpcOutput struct{ *pulumi.OutputState }

func (DefaultVpcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultVpc)(nil))
}

func (o DefaultVpcOutput) ToDefaultVpcOutput() DefaultVpcOutput {
	return o
}

func (o DefaultVpcOutput) ToDefaultVpcOutputWithContext(ctx context.Context) DefaultVpcOutput {
	return o
}

func (o DefaultVpcOutput) ToDefaultVpcPtrOutput() DefaultVpcPtrOutput {
	return o.ToDefaultVpcPtrOutputWithContext(context.Background())
}

func (o DefaultVpcOutput) ToDefaultVpcPtrOutputWithContext(ctx context.Context) DefaultVpcPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefaultVpc) *DefaultVpc {
		return &v
	}).(DefaultVpcPtrOutput)
}

type DefaultVpcPtrOutput struct{ *pulumi.OutputState }

func (DefaultVpcPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultVpc)(nil))
}

func (o DefaultVpcPtrOutput) ToDefaultVpcPtrOutput() DefaultVpcPtrOutput {
	return o
}

func (o DefaultVpcPtrOutput) ToDefaultVpcPtrOutputWithContext(ctx context.Context) DefaultVpcPtrOutput {
	return o
}

func (o DefaultVpcPtrOutput) Elem() DefaultVpcOutput {
	return o.ApplyT(func(v *DefaultVpc) DefaultVpc {
		if v != nil {
			return *v
		}
		var ret DefaultVpc
		return ret
	}).(DefaultVpcOutput)
}

type DefaultVpcArrayOutput struct{ *pulumi.OutputState }

func (DefaultVpcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DefaultVpc)(nil))
}

func (o DefaultVpcArrayOutput) ToDefaultVpcArrayOutput() DefaultVpcArrayOutput {
	return o
}

func (o DefaultVpcArrayOutput) ToDefaultVpcArrayOutputWithContext(ctx context.Context) DefaultVpcArrayOutput {
	return o
}

func (o DefaultVpcArrayOutput) Index(i pulumi.IntInput) DefaultVpcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DefaultVpc {
		return vs[0].([]DefaultVpc)[vs[1].(int)]
	}).(DefaultVpcOutput)
}

type DefaultVpcMapOutput struct{ *pulumi.OutputState }

func (DefaultVpcMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DefaultVpc)(nil))
}

func (o DefaultVpcMapOutput) ToDefaultVpcMapOutput() DefaultVpcMapOutput {
	return o
}

func (o DefaultVpcMapOutput) ToDefaultVpcMapOutputWithContext(ctx context.Context) DefaultVpcMapOutput {
	return o
}

func (o DefaultVpcMapOutput) MapIndex(k pulumi.StringInput) DefaultVpcOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DefaultVpc {
		return vs[0].(map[string]DefaultVpc)[vs[1].(string)]
	}).(DefaultVpcOutput)
}

func init() {
	pulumi.RegisterOutputType(DefaultVpcOutput{})
	pulumi.RegisterOutputType(DefaultVpcPtrOutput{})
	pulumi.RegisterOutputType(DefaultVpcArrayOutput{})
	pulumi.RegisterOutputType(DefaultVpcMapOutput{})
}
