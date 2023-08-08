// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Direct Connect hosted public virtual interface resource. This resource represents the allocator's side of the hosted virtual interface.
// A hosted virtual interface is a virtual interface that is owned by another AWS account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directconnect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := directconnect.NewHostedPublicVirtualInterface(ctx, "foo", &directconnect.HostedPublicVirtualInterfaceArgs{
//				AddressFamily:   pulumi.String("ipv4"),
//				AmazonAddress:   pulumi.String("175.45.176.2/30"),
//				BgpAsn:          pulumi.Int(65352),
//				ConnectionId:    pulumi.String("dxcon-zzzzzzzz"),
//				CustomerAddress: pulumi.String("175.45.176.1/30"),
//				RouteFilterPrefixes: pulumi.StringArray{
//					pulumi.String("210.52.109.0/24"),
//					pulumi.String("175.45.176.0/22"),
//				},
//				Vlan: pulumi.Int(4094),
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
//	to = aws_dx_hosted_public_virtual_interface.test
//
//	id = "dxvif-33cc44dd" } Using `pulumi import`, import Direct Connect hosted public virtual interfaces using the VIF `id`. For exampleconsole % pulumi import aws_dx_hosted_public_virtual_interface.test dxvif-33cc44dd
type HostedPublicVirtualInterface struct {
	pulumi.CustomResourceState

	// The address family for the BGP peer. ` ipv4  ` or `ipv6`.
	AddressFamily pulumi.StringOutput `pulumi:"addressFamily"`
	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress pulumi.StringOutput `pulumi:"amazonAddress"`
	AmazonSideAsn pulumi.StringOutput `pulumi:"amazonSideAsn"`
	// The ARN of the virtual interface.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The Direct Connect endpoint on which the virtual interface terminates.
	AwsDevice pulumi.StringOutput `pulumi:"awsDevice"`
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn pulumi.IntOutput `pulumi:"bgpAsn"`
	// The authentication key for BGP configuration.
	BgpAuthKey pulumi.StringOutput `pulumi:"bgpAuthKey"`
	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionId pulumi.StringOutput `pulumi:"connectionId"`
	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress pulumi.StringOutput `pulumi:"customerAddress"`
	// The name for the virtual interface.
	Name pulumi.StringOutput `pulumi:"name"`
	// The AWS account that will own the new virtual interface.
	OwnerAccountId pulumi.StringOutput `pulumi:"ownerAccountId"`
	// A list of routes to be advertised to the AWS network in this region.
	RouteFilterPrefixes pulumi.StringArrayOutput `pulumi:"routeFilterPrefixes"`
	// The VLAN ID.
	Vlan pulumi.IntOutput `pulumi:"vlan"`
}

// NewHostedPublicVirtualInterface registers a new resource with the given unique name, arguments, and options.
func NewHostedPublicVirtualInterface(ctx *pulumi.Context,
	name string, args *HostedPublicVirtualInterfaceArgs, opts ...pulumi.ResourceOption) (*HostedPublicVirtualInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddressFamily == nil {
		return nil, errors.New("invalid value for required argument 'AddressFamily'")
	}
	if args.BgpAsn == nil {
		return nil, errors.New("invalid value for required argument 'BgpAsn'")
	}
	if args.ConnectionId == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionId'")
	}
	if args.OwnerAccountId == nil {
		return nil, errors.New("invalid value for required argument 'OwnerAccountId'")
	}
	if args.RouteFilterPrefixes == nil {
		return nil, errors.New("invalid value for required argument 'RouteFilterPrefixes'")
	}
	if args.Vlan == nil {
		return nil, errors.New("invalid value for required argument 'Vlan'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource HostedPublicVirtualInterface
	err := ctx.RegisterResource("aws:directconnect/hostedPublicVirtualInterface:HostedPublicVirtualInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostedPublicVirtualInterface gets an existing HostedPublicVirtualInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostedPublicVirtualInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostedPublicVirtualInterfaceState, opts ...pulumi.ResourceOption) (*HostedPublicVirtualInterface, error) {
	var resource HostedPublicVirtualInterface
	err := ctx.ReadResource("aws:directconnect/hostedPublicVirtualInterface:HostedPublicVirtualInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostedPublicVirtualInterface resources.
type hostedPublicVirtualInterfaceState struct {
	// The address family for the BGP peer. ` ipv4  ` or `ipv6`.
	AddressFamily *string `pulumi:"addressFamily"`
	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress *string `pulumi:"amazonAddress"`
	AmazonSideAsn *string `pulumi:"amazonSideAsn"`
	// The ARN of the virtual interface.
	Arn *string `pulumi:"arn"`
	// The Direct Connect endpoint on which the virtual interface terminates.
	AwsDevice *string `pulumi:"awsDevice"`
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn *int `pulumi:"bgpAsn"`
	// The authentication key for BGP configuration.
	BgpAuthKey *string `pulumi:"bgpAuthKey"`
	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionId *string `pulumi:"connectionId"`
	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress *string `pulumi:"customerAddress"`
	// The name for the virtual interface.
	Name *string `pulumi:"name"`
	// The AWS account that will own the new virtual interface.
	OwnerAccountId *string `pulumi:"ownerAccountId"`
	// A list of routes to be advertised to the AWS network in this region.
	RouteFilterPrefixes []string `pulumi:"routeFilterPrefixes"`
	// The VLAN ID.
	Vlan *int `pulumi:"vlan"`
}

type HostedPublicVirtualInterfaceState struct {
	// The address family for the BGP peer. ` ipv4  ` or `ipv6`.
	AddressFamily pulumi.StringPtrInput
	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress pulumi.StringPtrInput
	AmazonSideAsn pulumi.StringPtrInput
	// The ARN of the virtual interface.
	Arn pulumi.StringPtrInput
	// The Direct Connect endpoint on which the virtual interface terminates.
	AwsDevice pulumi.StringPtrInput
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn pulumi.IntPtrInput
	// The authentication key for BGP configuration.
	BgpAuthKey pulumi.StringPtrInput
	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionId pulumi.StringPtrInput
	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress pulumi.StringPtrInput
	// The name for the virtual interface.
	Name pulumi.StringPtrInput
	// The AWS account that will own the new virtual interface.
	OwnerAccountId pulumi.StringPtrInput
	// A list of routes to be advertised to the AWS network in this region.
	RouteFilterPrefixes pulumi.StringArrayInput
	// The VLAN ID.
	Vlan pulumi.IntPtrInput
}

func (HostedPublicVirtualInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedPublicVirtualInterfaceState)(nil)).Elem()
}

type hostedPublicVirtualInterfaceArgs struct {
	// The address family for the BGP peer. ` ipv4  ` or `ipv6`.
	AddressFamily string `pulumi:"addressFamily"`
	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress *string `pulumi:"amazonAddress"`
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn int `pulumi:"bgpAsn"`
	// The authentication key for BGP configuration.
	BgpAuthKey *string `pulumi:"bgpAuthKey"`
	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionId string `pulumi:"connectionId"`
	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress *string `pulumi:"customerAddress"`
	// The name for the virtual interface.
	Name *string `pulumi:"name"`
	// The AWS account that will own the new virtual interface.
	OwnerAccountId string `pulumi:"ownerAccountId"`
	// A list of routes to be advertised to the AWS network in this region.
	RouteFilterPrefixes []string `pulumi:"routeFilterPrefixes"`
	// The VLAN ID.
	Vlan int `pulumi:"vlan"`
}

// The set of arguments for constructing a HostedPublicVirtualInterface resource.
type HostedPublicVirtualInterfaceArgs struct {
	// The address family for the BGP peer. ` ipv4  ` or `ipv6`.
	AddressFamily pulumi.StringInput
	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress pulumi.StringPtrInput
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn pulumi.IntInput
	// The authentication key for BGP configuration.
	BgpAuthKey pulumi.StringPtrInput
	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionId pulumi.StringInput
	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress pulumi.StringPtrInput
	// The name for the virtual interface.
	Name pulumi.StringPtrInput
	// The AWS account that will own the new virtual interface.
	OwnerAccountId pulumi.StringInput
	// A list of routes to be advertised to the AWS network in this region.
	RouteFilterPrefixes pulumi.StringArrayInput
	// The VLAN ID.
	Vlan pulumi.IntInput
}

func (HostedPublicVirtualInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedPublicVirtualInterfaceArgs)(nil)).Elem()
}

type HostedPublicVirtualInterfaceInput interface {
	pulumi.Input

	ToHostedPublicVirtualInterfaceOutput() HostedPublicVirtualInterfaceOutput
	ToHostedPublicVirtualInterfaceOutputWithContext(ctx context.Context) HostedPublicVirtualInterfaceOutput
}

func (*HostedPublicVirtualInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**HostedPublicVirtualInterface)(nil)).Elem()
}

func (i *HostedPublicVirtualInterface) ToHostedPublicVirtualInterfaceOutput() HostedPublicVirtualInterfaceOutput {
	return i.ToHostedPublicVirtualInterfaceOutputWithContext(context.Background())
}

func (i *HostedPublicVirtualInterface) ToHostedPublicVirtualInterfaceOutputWithContext(ctx context.Context) HostedPublicVirtualInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedPublicVirtualInterfaceOutput)
}

// HostedPublicVirtualInterfaceArrayInput is an input type that accepts HostedPublicVirtualInterfaceArray and HostedPublicVirtualInterfaceArrayOutput values.
// You can construct a concrete instance of `HostedPublicVirtualInterfaceArrayInput` via:
//
//	HostedPublicVirtualInterfaceArray{ HostedPublicVirtualInterfaceArgs{...} }
type HostedPublicVirtualInterfaceArrayInput interface {
	pulumi.Input

	ToHostedPublicVirtualInterfaceArrayOutput() HostedPublicVirtualInterfaceArrayOutput
	ToHostedPublicVirtualInterfaceArrayOutputWithContext(context.Context) HostedPublicVirtualInterfaceArrayOutput
}

type HostedPublicVirtualInterfaceArray []HostedPublicVirtualInterfaceInput

func (HostedPublicVirtualInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostedPublicVirtualInterface)(nil)).Elem()
}

func (i HostedPublicVirtualInterfaceArray) ToHostedPublicVirtualInterfaceArrayOutput() HostedPublicVirtualInterfaceArrayOutput {
	return i.ToHostedPublicVirtualInterfaceArrayOutputWithContext(context.Background())
}

func (i HostedPublicVirtualInterfaceArray) ToHostedPublicVirtualInterfaceArrayOutputWithContext(ctx context.Context) HostedPublicVirtualInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedPublicVirtualInterfaceArrayOutput)
}

// HostedPublicVirtualInterfaceMapInput is an input type that accepts HostedPublicVirtualInterfaceMap and HostedPublicVirtualInterfaceMapOutput values.
// You can construct a concrete instance of `HostedPublicVirtualInterfaceMapInput` via:
//
//	HostedPublicVirtualInterfaceMap{ "key": HostedPublicVirtualInterfaceArgs{...} }
type HostedPublicVirtualInterfaceMapInput interface {
	pulumi.Input

	ToHostedPublicVirtualInterfaceMapOutput() HostedPublicVirtualInterfaceMapOutput
	ToHostedPublicVirtualInterfaceMapOutputWithContext(context.Context) HostedPublicVirtualInterfaceMapOutput
}

type HostedPublicVirtualInterfaceMap map[string]HostedPublicVirtualInterfaceInput

func (HostedPublicVirtualInterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostedPublicVirtualInterface)(nil)).Elem()
}

func (i HostedPublicVirtualInterfaceMap) ToHostedPublicVirtualInterfaceMapOutput() HostedPublicVirtualInterfaceMapOutput {
	return i.ToHostedPublicVirtualInterfaceMapOutputWithContext(context.Background())
}

func (i HostedPublicVirtualInterfaceMap) ToHostedPublicVirtualInterfaceMapOutputWithContext(ctx context.Context) HostedPublicVirtualInterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedPublicVirtualInterfaceMapOutput)
}

type HostedPublicVirtualInterfaceOutput struct{ *pulumi.OutputState }

func (HostedPublicVirtualInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostedPublicVirtualInterface)(nil)).Elem()
}

func (o HostedPublicVirtualInterfaceOutput) ToHostedPublicVirtualInterfaceOutput() HostedPublicVirtualInterfaceOutput {
	return o
}

func (o HostedPublicVirtualInterfaceOutput) ToHostedPublicVirtualInterfaceOutputWithContext(ctx context.Context) HostedPublicVirtualInterfaceOutput {
	return o
}

// The address family for the BGP peer. ` ipv4  ` or `ipv6`.
func (o HostedPublicVirtualInterfaceOutput) AddressFamily() pulumi.StringOutput {
	return o.ApplyT(func(v *HostedPublicVirtualInterface) pulumi.StringOutput { return v.AddressFamily }).(pulumi.StringOutput)
}

// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
func (o HostedPublicVirtualInterfaceOutput) AmazonAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *HostedPublicVirtualInterface) pulumi.StringOutput { return v.AmazonAddress }).(pulumi.StringOutput)
}

func (o HostedPublicVirtualInterfaceOutput) AmazonSideAsn() pulumi.StringOutput {
	return o.ApplyT(func(v *HostedPublicVirtualInterface) pulumi.StringOutput { return v.AmazonSideAsn }).(pulumi.StringOutput)
}

// The ARN of the virtual interface.
func (o HostedPublicVirtualInterfaceOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *HostedPublicVirtualInterface) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The Direct Connect endpoint on which the virtual interface terminates.
func (o HostedPublicVirtualInterfaceOutput) AwsDevice() pulumi.StringOutput {
	return o.ApplyT(func(v *HostedPublicVirtualInterface) pulumi.StringOutput { return v.AwsDevice }).(pulumi.StringOutput)
}

// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
func (o HostedPublicVirtualInterfaceOutput) BgpAsn() pulumi.IntOutput {
	return o.ApplyT(func(v *HostedPublicVirtualInterface) pulumi.IntOutput { return v.BgpAsn }).(pulumi.IntOutput)
}

// The authentication key for BGP configuration.
func (o HostedPublicVirtualInterfaceOutput) BgpAuthKey() pulumi.StringOutput {
	return o.ApplyT(func(v *HostedPublicVirtualInterface) pulumi.StringOutput { return v.BgpAuthKey }).(pulumi.StringOutput)
}

// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
func (o HostedPublicVirtualInterfaceOutput) ConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostedPublicVirtualInterface) pulumi.StringOutput { return v.ConnectionId }).(pulumi.StringOutput)
}

// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
func (o HostedPublicVirtualInterfaceOutput) CustomerAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *HostedPublicVirtualInterface) pulumi.StringOutput { return v.CustomerAddress }).(pulumi.StringOutput)
}

// The name for the virtual interface.
func (o HostedPublicVirtualInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HostedPublicVirtualInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The AWS account that will own the new virtual interface.
func (o HostedPublicVirtualInterfaceOutput) OwnerAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostedPublicVirtualInterface) pulumi.StringOutput { return v.OwnerAccountId }).(pulumi.StringOutput)
}

// A list of routes to be advertised to the AWS network in this region.
func (o HostedPublicVirtualInterfaceOutput) RouteFilterPrefixes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HostedPublicVirtualInterface) pulumi.StringArrayOutput { return v.RouteFilterPrefixes }).(pulumi.StringArrayOutput)
}

// The VLAN ID.
func (o HostedPublicVirtualInterfaceOutput) Vlan() pulumi.IntOutput {
	return o.ApplyT(func(v *HostedPublicVirtualInterface) pulumi.IntOutput { return v.Vlan }).(pulumi.IntOutput)
}

type HostedPublicVirtualInterfaceArrayOutput struct{ *pulumi.OutputState }

func (HostedPublicVirtualInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostedPublicVirtualInterface)(nil)).Elem()
}

func (o HostedPublicVirtualInterfaceArrayOutput) ToHostedPublicVirtualInterfaceArrayOutput() HostedPublicVirtualInterfaceArrayOutput {
	return o
}

func (o HostedPublicVirtualInterfaceArrayOutput) ToHostedPublicVirtualInterfaceArrayOutputWithContext(ctx context.Context) HostedPublicVirtualInterfaceArrayOutput {
	return o
}

func (o HostedPublicVirtualInterfaceArrayOutput) Index(i pulumi.IntInput) HostedPublicVirtualInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HostedPublicVirtualInterface {
		return vs[0].([]*HostedPublicVirtualInterface)[vs[1].(int)]
	}).(HostedPublicVirtualInterfaceOutput)
}

type HostedPublicVirtualInterfaceMapOutput struct{ *pulumi.OutputState }

func (HostedPublicVirtualInterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostedPublicVirtualInterface)(nil)).Elem()
}

func (o HostedPublicVirtualInterfaceMapOutput) ToHostedPublicVirtualInterfaceMapOutput() HostedPublicVirtualInterfaceMapOutput {
	return o
}

func (o HostedPublicVirtualInterfaceMapOutput) ToHostedPublicVirtualInterfaceMapOutputWithContext(ctx context.Context) HostedPublicVirtualInterfaceMapOutput {
	return o
}

func (o HostedPublicVirtualInterfaceMapOutput) MapIndex(k pulumi.StringInput) HostedPublicVirtualInterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HostedPublicVirtualInterface {
		return vs[0].(map[string]*HostedPublicVirtualInterface)[vs[1].(string)]
	}).(HostedPublicVirtualInterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HostedPublicVirtualInterfaceInput)(nil)).Elem(), &HostedPublicVirtualInterface{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostedPublicVirtualInterfaceArrayInput)(nil)).Elem(), HostedPublicVirtualInterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostedPublicVirtualInterfaceMapInput)(nil)).Elem(), HostedPublicVirtualInterfaceMap{})
	pulumi.RegisterOutputType(HostedPublicVirtualInterfaceOutput{})
	pulumi.RegisterOutputType(HostedPublicVirtualInterfaceArrayOutput{})
	pulumi.RegisterOutputType(HostedPublicVirtualInterfaceMapOutput{})
}
