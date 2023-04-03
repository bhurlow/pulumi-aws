// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Direct Connect hosted private virtual interface resource. This resource represents the allocator's side of the hosted virtual interface.
// A hosted virtual interface is a virtual interface that is owned by another AWS account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/directconnect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := directconnect.NewHostedPrivateVirtualInterface(ctx, "foo", &directconnect.HostedPrivateVirtualInterfaceArgs{
//				AddressFamily: pulumi.String("ipv4"),
//				BgpAsn:        pulumi.Int(65352),
//				ConnectionId:  pulumi.String("dxcon-zzzzzzzz"),
//				Vlan:          pulumi.Int(4094),
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
// Direct Connect hosted private virtual interfaces can be imported using the `vif id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:directconnect/hostedPrivateVirtualInterface:HostedPrivateVirtualInterface test dxvif-33cc44dd
//
// ```
type HostedPrivateVirtualInterface struct {
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
	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable pulumi.BoolOutput `pulumi:"jumboFrameCapable"`
	// The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection. The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
	Mtu pulumi.IntPtrOutput `pulumi:"mtu"`
	// The name for the virtual interface.
	Name pulumi.StringOutput `pulumi:"name"`
	// The AWS account that will own the new virtual interface.
	OwnerAccountId pulumi.StringOutput `pulumi:"ownerAccountId"`
	// The VLAN ID.
	Vlan pulumi.IntOutput `pulumi:"vlan"`
}

// NewHostedPrivateVirtualInterface registers a new resource with the given unique name, arguments, and options.
func NewHostedPrivateVirtualInterface(ctx *pulumi.Context,
	name string, args *HostedPrivateVirtualInterfaceArgs, opts ...pulumi.ResourceOption) (*HostedPrivateVirtualInterface, error) {
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
	if args.Vlan == nil {
		return nil, errors.New("invalid value for required argument 'Vlan'")
	}
	var resource HostedPrivateVirtualInterface
	err := ctx.RegisterResource("aws:directconnect/hostedPrivateVirtualInterface:HostedPrivateVirtualInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostedPrivateVirtualInterface gets an existing HostedPrivateVirtualInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostedPrivateVirtualInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostedPrivateVirtualInterfaceState, opts ...pulumi.ResourceOption) (*HostedPrivateVirtualInterface, error) {
	var resource HostedPrivateVirtualInterface
	err := ctx.ReadResource("aws:directconnect/hostedPrivateVirtualInterface:HostedPrivateVirtualInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostedPrivateVirtualInterface resources.
type hostedPrivateVirtualInterfaceState struct {
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
	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool `pulumi:"jumboFrameCapable"`
	// The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection. The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
	Mtu *int `pulumi:"mtu"`
	// The name for the virtual interface.
	Name *string `pulumi:"name"`
	// The AWS account that will own the new virtual interface.
	OwnerAccountId *string `pulumi:"ownerAccountId"`
	// The VLAN ID.
	Vlan *int `pulumi:"vlan"`
}

type HostedPrivateVirtualInterfaceState struct {
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
	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable pulumi.BoolPtrInput
	// The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection. The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
	Mtu pulumi.IntPtrInput
	// The name for the virtual interface.
	Name pulumi.StringPtrInput
	// The AWS account that will own the new virtual interface.
	OwnerAccountId pulumi.StringPtrInput
	// The VLAN ID.
	Vlan pulumi.IntPtrInput
}

func (HostedPrivateVirtualInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedPrivateVirtualInterfaceState)(nil)).Elem()
}

type hostedPrivateVirtualInterfaceArgs struct {
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
	// The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection. The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
	Mtu *int `pulumi:"mtu"`
	// The name for the virtual interface.
	Name *string `pulumi:"name"`
	// The AWS account that will own the new virtual interface.
	OwnerAccountId string `pulumi:"ownerAccountId"`
	// The VLAN ID.
	Vlan int `pulumi:"vlan"`
}

// The set of arguments for constructing a HostedPrivateVirtualInterface resource.
type HostedPrivateVirtualInterfaceArgs struct {
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
	// The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection. The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
	Mtu pulumi.IntPtrInput
	// The name for the virtual interface.
	Name pulumi.StringPtrInput
	// The AWS account that will own the new virtual interface.
	OwnerAccountId pulumi.StringInput
	// The VLAN ID.
	Vlan pulumi.IntInput
}

func (HostedPrivateVirtualInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedPrivateVirtualInterfaceArgs)(nil)).Elem()
}

type HostedPrivateVirtualInterfaceInput interface {
	pulumi.Input

	ToHostedPrivateVirtualInterfaceOutput() HostedPrivateVirtualInterfaceOutput
	ToHostedPrivateVirtualInterfaceOutputWithContext(ctx context.Context) HostedPrivateVirtualInterfaceOutput
}

func (*HostedPrivateVirtualInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**HostedPrivateVirtualInterface)(nil)).Elem()
}

func (i *HostedPrivateVirtualInterface) ToHostedPrivateVirtualInterfaceOutput() HostedPrivateVirtualInterfaceOutput {
	return i.ToHostedPrivateVirtualInterfaceOutputWithContext(context.Background())
}

func (i *HostedPrivateVirtualInterface) ToHostedPrivateVirtualInterfaceOutputWithContext(ctx context.Context) HostedPrivateVirtualInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedPrivateVirtualInterfaceOutput)
}

// HostedPrivateVirtualInterfaceArrayInput is an input type that accepts HostedPrivateVirtualInterfaceArray and HostedPrivateVirtualInterfaceArrayOutput values.
// You can construct a concrete instance of `HostedPrivateVirtualInterfaceArrayInput` via:
//
//	HostedPrivateVirtualInterfaceArray{ HostedPrivateVirtualInterfaceArgs{...} }
type HostedPrivateVirtualInterfaceArrayInput interface {
	pulumi.Input

	ToHostedPrivateVirtualInterfaceArrayOutput() HostedPrivateVirtualInterfaceArrayOutput
	ToHostedPrivateVirtualInterfaceArrayOutputWithContext(context.Context) HostedPrivateVirtualInterfaceArrayOutput
}

type HostedPrivateVirtualInterfaceArray []HostedPrivateVirtualInterfaceInput

func (HostedPrivateVirtualInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostedPrivateVirtualInterface)(nil)).Elem()
}

func (i HostedPrivateVirtualInterfaceArray) ToHostedPrivateVirtualInterfaceArrayOutput() HostedPrivateVirtualInterfaceArrayOutput {
	return i.ToHostedPrivateVirtualInterfaceArrayOutputWithContext(context.Background())
}

func (i HostedPrivateVirtualInterfaceArray) ToHostedPrivateVirtualInterfaceArrayOutputWithContext(ctx context.Context) HostedPrivateVirtualInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedPrivateVirtualInterfaceArrayOutput)
}

// HostedPrivateVirtualInterfaceMapInput is an input type that accepts HostedPrivateVirtualInterfaceMap and HostedPrivateVirtualInterfaceMapOutput values.
// You can construct a concrete instance of `HostedPrivateVirtualInterfaceMapInput` via:
//
//	HostedPrivateVirtualInterfaceMap{ "key": HostedPrivateVirtualInterfaceArgs{...} }
type HostedPrivateVirtualInterfaceMapInput interface {
	pulumi.Input

	ToHostedPrivateVirtualInterfaceMapOutput() HostedPrivateVirtualInterfaceMapOutput
	ToHostedPrivateVirtualInterfaceMapOutputWithContext(context.Context) HostedPrivateVirtualInterfaceMapOutput
}

type HostedPrivateVirtualInterfaceMap map[string]HostedPrivateVirtualInterfaceInput

func (HostedPrivateVirtualInterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostedPrivateVirtualInterface)(nil)).Elem()
}

func (i HostedPrivateVirtualInterfaceMap) ToHostedPrivateVirtualInterfaceMapOutput() HostedPrivateVirtualInterfaceMapOutput {
	return i.ToHostedPrivateVirtualInterfaceMapOutputWithContext(context.Background())
}

func (i HostedPrivateVirtualInterfaceMap) ToHostedPrivateVirtualInterfaceMapOutputWithContext(ctx context.Context) HostedPrivateVirtualInterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostedPrivateVirtualInterfaceMapOutput)
}

type HostedPrivateVirtualInterfaceOutput struct{ *pulumi.OutputState }

func (HostedPrivateVirtualInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostedPrivateVirtualInterface)(nil)).Elem()
}

func (o HostedPrivateVirtualInterfaceOutput) ToHostedPrivateVirtualInterfaceOutput() HostedPrivateVirtualInterfaceOutput {
	return o
}

func (o HostedPrivateVirtualInterfaceOutput) ToHostedPrivateVirtualInterfaceOutputWithContext(ctx context.Context) HostedPrivateVirtualInterfaceOutput {
	return o
}

// The address family for the BGP peer. ` ipv4  ` or `ipv6`.
func (o HostedPrivateVirtualInterfaceOutput) AddressFamily() pulumi.StringOutput {
	return o.ApplyT(func(v *HostedPrivateVirtualInterface) pulumi.StringOutput { return v.AddressFamily }).(pulumi.StringOutput)
}

// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
func (o HostedPrivateVirtualInterfaceOutput) AmazonAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *HostedPrivateVirtualInterface) pulumi.StringOutput { return v.AmazonAddress }).(pulumi.StringOutput)
}

func (o HostedPrivateVirtualInterfaceOutput) AmazonSideAsn() pulumi.StringOutput {
	return o.ApplyT(func(v *HostedPrivateVirtualInterface) pulumi.StringOutput { return v.AmazonSideAsn }).(pulumi.StringOutput)
}

// The ARN of the virtual interface.
func (o HostedPrivateVirtualInterfaceOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *HostedPrivateVirtualInterface) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The Direct Connect endpoint on which the virtual interface terminates.
func (o HostedPrivateVirtualInterfaceOutput) AwsDevice() pulumi.StringOutput {
	return o.ApplyT(func(v *HostedPrivateVirtualInterface) pulumi.StringOutput { return v.AwsDevice }).(pulumi.StringOutput)
}

// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
func (o HostedPrivateVirtualInterfaceOutput) BgpAsn() pulumi.IntOutput {
	return o.ApplyT(func(v *HostedPrivateVirtualInterface) pulumi.IntOutput { return v.BgpAsn }).(pulumi.IntOutput)
}

// The authentication key for BGP configuration.
func (o HostedPrivateVirtualInterfaceOutput) BgpAuthKey() pulumi.StringOutput {
	return o.ApplyT(func(v *HostedPrivateVirtualInterface) pulumi.StringOutput { return v.BgpAuthKey }).(pulumi.StringOutput)
}

// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
func (o HostedPrivateVirtualInterfaceOutput) ConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostedPrivateVirtualInterface) pulumi.StringOutput { return v.ConnectionId }).(pulumi.StringOutput)
}

// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
func (o HostedPrivateVirtualInterfaceOutput) CustomerAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *HostedPrivateVirtualInterface) pulumi.StringOutput { return v.CustomerAddress }).(pulumi.StringOutput)
}

// Indicates whether jumbo frames (9001 MTU) are supported.
func (o HostedPrivateVirtualInterfaceOutput) JumboFrameCapable() pulumi.BoolOutput {
	return o.ApplyT(func(v *HostedPrivateVirtualInterface) pulumi.BoolOutput { return v.JumboFrameCapable }).(pulumi.BoolOutput)
}

// The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection. The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
func (o HostedPrivateVirtualInterfaceOutput) Mtu() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HostedPrivateVirtualInterface) pulumi.IntPtrOutput { return v.Mtu }).(pulumi.IntPtrOutput)
}

// The name for the virtual interface.
func (o HostedPrivateVirtualInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HostedPrivateVirtualInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The AWS account that will own the new virtual interface.
func (o HostedPrivateVirtualInterfaceOutput) OwnerAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *HostedPrivateVirtualInterface) pulumi.StringOutput { return v.OwnerAccountId }).(pulumi.StringOutput)
}

// The VLAN ID.
func (o HostedPrivateVirtualInterfaceOutput) Vlan() pulumi.IntOutput {
	return o.ApplyT(func(v *HostedPrivateVirtualInterface) pulumi.IntOutput { return v.Vlan }).(pulumi.IntOutput)
}

type HostedPrivateVirtualInterfaceArrayOutput struct{ *pulumi.OutputState }

func (HostedPrivateVirtualInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostedPrivateVirtualInterface)(nil)).Elem()
}

func (o HostedPrivateVirtualInterfaceArrayOutput) ToHostedPrivateVirtualInterfaceArrayOutput() HostedPrivateVirtualInterfaceArrayOutput {
	return o
}

func (o HostedPrivateVirtualInterfaceArrayOutput) ToHostedPrivateVirtualInterfaceArrayOutputWithContext(ctx context.Context) HostedPrivateVirtualInterfaceArrayOutput {
	return o
}

func (o HostedPrivateVirtualInterfaceArrayOutput) Index(i pulumi.IntInput) HostedPrivateVirtualInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HostedPrivateVirtualInterface {
		return vs[0].([]*HostedPrivateVirtualInterface)[vs[1].(int)]
	}).(HostedPrivateVirtualInterfaceOutput)
}

type HostedPrivateVirtualInterfaceMapOutput struct{ *pulumi.OutputState }

func (HostedPrivateVirtualInterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostedPrivateVirtualInterface)(nil)).Elem()
}

func (o HostedPrivateVirtualInterfaceMapOutput) ToHostedPrivateVirtualInterfaceMapOutput() HostedPrivateVirtualInterfaceMapOutput {
	return o
}

func (o HostedPrivateVirtualInterfaceMapOutput) ToHostedPrivateVirtualInterfaceMapOutputWithContext(ctx context.Context) HostedPrivateVirtualInterfaceMapOutput {
	return o
}

func (o HostedPrivateVirtualInterfaceMapOutput) MapIndex(k pulumi.StringInput) HostedPrivateVirtualInterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HostedPrivateVirtualInterface {
		return vs[0].(map[string]*HostedPrivateVirtualInterface)[vs[1].(string)]
	}).(HostedPrivateVirtualInterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HostedPrivateVirtualInterfaceInput)(nil)).Elem(), &HostedPrivateVirtualInterface{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostedPrivateVirtualInterfaceArrayInput)(nil)).Elem(), HostedPrivateVirtualInterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostedPrivateVirtualInterfaceMapInput)(nil)).Elem(), HostedPrivateVirtualInterfaceMap{})
	pulumi.RegisterOutputType(HostedPrivateVirtualInterfaceOutput{})
	pulumi.RegisterOutputType(HostedPrivateVirtualInterfaceArrayOutput{})
	pulumi.RegisterOutputType(HostedPrivateVirtualInterfaceMapOutput{})
}
