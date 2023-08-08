// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a VPC NAT Gateway.
//
// ## Example Usage
// ### Public NAT
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.NewNatGateway(ctx, "example", &ec2.NatGatewayArgs{
//				AllocationId: pulumi.Any(aws_eip.Example.Id),
//				SubnetId:     pulumi.Any(aws_subnet.Example.Id),
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("gw NAT"),
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				aws_internet_gateway.Example,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Public NAT with Secondary Private IP Addresses
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.NewNatGateway(ctx, "example", &ec2.NatGatewayArgs{
//				AllocationId: pulumi.Any(aws_eip.Example.Id),
//				SubnetId:     pulumi.Any(aws_subnet.Example.Id),
//				SecondaryAllocationIds: pulumi.StringArray{
//					aws_eip.Secondary.Id,
//				},
//				SecondaryPrivateIpAddresses: pulumi.StringArray{
//					pulumi.String("10.0.1.5"),
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
// ### Private NAT
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.NewNatGateway(ctx, "example", &ec2.NatGatewayArgs{
//				ConnectivityType: pulumi.String("private"),
//				SubnetId:         pulumi.Any(aws_subnet.Example.Id),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Private NAT with Secondary Private IP Addresses
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.NewNatGateway(ctx, "example", &ec2.NatGatewayArgs{
//				ConnectivityType:               pulumi.String("private"),
//				SubnetId:                       pulumi.Any(aws_subnet.Example.Id),
//				SecondaryPrivateIpAddressCount: pulumi.Int(7),
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
//	to = aws_nat_gateway.private_gw
//
//	id = "nat-05dba92075d71c408" } Using `pulumi import`, import NAT Gateways using the `id`. For exampleconsole % pulumi import aws_nat_gateway.private_gw nat-05dba92075d71c408
type NatGateway struct {
	pulumi.CustomResourceState

	// The Allocation ID of the Elastic IP address for the NAT Gateway. Required for `connectivityType` of `public`.
	AllocationId pulumi.StringPtrOutput `pulumi:"allocationId"`
	// The association ID of the Elastic IP address that's associated with the NAT Gateway. Only available when `connectivityType` is `public`.
	AssociationId pulumi.StringOutput `pulumi:"associationId"`
	// Connectivity type for the NAT Gateway. Valid values are `private` and `public`. Defaults to `public`.
	ConnectivityType pulumi.StringPtrOutput `pulumi:"connectivityType"`
	// The ID of the network interface associated with the NAT Gateway.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	// The private IPv4 address to assign to the NAT Gateway. If you don't provide an address, a private IPv4 address will be automatically assigned.
	PrivateIp pulumi.StringOutput `pulumi:"privateIp"`
	// The Elastic IP address associated with the NAT Gateway.
	PublicIp pulumi.StringOutput `pulumi:"publicIp"`
	// A list of secondary allocation EIP IDs for this NAT Gateway.
	SecondaryAllocationIds pulumi.StringArrayOutput `pulumi:"secondaryAllocationIds"`
	// [Private NAT Gateway only] The number of secondary private IPv4 addresses you want to assign to the NAT Gateway.
	SecondaryPrivateIpAddressCount pulumi.IntOutput `pulumi:"secondaryPrivateIpAddressCount"`
	// A list of secondary private IPv4 addresses to assign to the NAT Gateway.
	SecondaryPrivateIpAddresses pulumi.StringArrayOutput `pulumi:"secondaryPrivateIpAddresses"`
	// The Subnet ID of the subnet in which to place the NAT Gateway.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewNatGateway registers a new resource with the given unique name, arguments, and options.
func NewNatGateway(ctx *pulumi.Context,
	name string, args *NatGatewayArgs, opts ...pulumi.ResourceOption) (*NatGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NatGateway
	err := ctx.RegisterResource("aws:ec2/natGateway:NatGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNatGateway gets an existing NatGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNatGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NatGatewayState, opts ...pulumi.ResourceOption) (*NatGateway, error) {
	var resource NatGateway
	err := ctx.ReadResource("aws:ec2/natGateway:NatGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NatGateway resources.
type natGatewayState struct {
	// The Allocation ID of the Elastic IP address for the NAT Gateway. Required for `connectivityType` of `public`.
	AllocationId *string `pulumi:"allocationId"`
	// The association ID of the Elastic IP address that's associated with the NAT Gateway. Only available when `connectivityType` is `public`.
	AssociationId *string `pulumi:"associationId"`
	// Connectivity type for the NAT Gateway. Valid values are `private` and `public`. Defaults to `public`.
	ConnectivityType *string `pulumi:"connectivityType"`
	// The ID of the network interface associated with the NAT Gateway.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The private IPv4 address to assign to the NAT Gateway. If you don't provide an address, a private IPv4 address will be automatically assigned.
	PrivateIp *string `pulumi:"privateIp"`
	// The Elastic IP address associated with the NAT Gateway.
	PublicIp *string `pulumi:"publicIp"`
	// A list of secondary allocation EIP IDs for this NAT Gateway.
	SecondaryAllocationIds []string `pulumi:"secondaryAllocationIds"`
	// [Private NAT Gateway only] The number of secondary private IPv4 addresses you want to assign to the NAT Gateway.
	SecondaryPrivateIpAddressCount *int `pulumi:"secondaryPrivateIpAddressCount"`
	// A list of secondary private IPv4 addresses to assign to the NAT Gateway.
	SecondaryPrivateIpAddresses []string `pulumi:"secondaryPrivateIpAddresses"`
	// The Subnet ID of the subnet in which to place the NAT Gateway.
	SubnetId *string `pulumi:"subnetId"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type NatGatewayState struct {
	// The Allocation ID of the Elastic IP address for the NAT Gateway. Required for `connectivityType` of `public`.
	AllocationId pulumi.StringPtrInput
	// The association ID of the Elastic IP address that's associated with the NAT Gateway. Only available when `connectivityType` is `public`.
	AssociationId pulumi.StringPtrInput
	// Connectivity type for the NAT Gateway. Valid values are `private` and `public`. Defaults to `public`.
	ConnectivityType pulumi.StringPtrInput
	// The ID of the network interface associated with the NAT Gateway.
	NetworkInterfaceId pulumi.StringPtrInput
	// The private IPv4 address to assign to the NAT Gateway. If you don't provide an address, a private IPv4 address will be automatically assigned.
	PrivateIp pulumi.StringPtrInput
	// The Elastic IP address associated with the NAT Gateway.
	PublicIp pulumi.StringPtrInput
	// A list of secondary allocation EIP IDs for this NAT Gateway.
	SecondaryAllocationIds pulumi.StringArrayInput
	// [Private NAT Gateway only] The number of secondary private IPv4 addresses you want to assign to the NAT Gateway.
	SecondaryPrivateIpAddressCount pulumi.IntPtrInput
	// A list of secondary private IPv4 addresses to assign to the NAT Gateway.
	SecondaryPrivateIpAddresses pulumi.StringArrayInput
	// The Subnet ID of the subnet in which to place the NAT Gateway.
	SubnetId pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (NatGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*natGatewayState)(nil)).Elem()
}

type natGatewayArgs struct {
	// The Allocation ID of the Elastic IP address for the NAT Gateway. Required for `connectivityType` of `public`.
	AllocationId *string `pulumi:"allocationId"`
	// Connectivity type for the NAT Gateway. Valid values are `private` and `public`. Defaults to `public`.
	ConnectivityType *string `pulumi:"connectivityType"`
	// The private IPv4 address to assign to the NAT Gateway. If you don't provide an address, a private IPv4 address will be automatically assigned.
	PrivateIp *string `pulumi:"privateIp"`
	// A list of secondary allocation EIP IDs for this NAT Gateway.
	SecondaryAllocationIds []string `pulumi:"secondaryAllocationIds"`
	// [Private NAT Gateway only] The number of secondary private IPv4 addresses you want to assign to the NAT Gateway.
	SecondaryPrivateIpAddressCount *int `pulumi:"secondaryPrivateIpAddressCount"`
	// A list of secondary private IPv4 addresses to assign to the NAT Gateway.
	SecondaryPrivateIpAddresses []string `pulumi:"secondaryPrivateIpAddresses"`
	// The Subnet ID of the subnet in which to place the NAT Gateway.
	SubnetId string `pulumi:"subnetId"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NatGateway resource.
type NatGatewayArgs struct {
	// The Allocation ID of the Elastic IP address for the NAT Gateway. Required for `connectivityType` of `public`.
	AllocationId pulumi.StringPtrInput
	// Connectivity type for the NAT Gateway. Valid values are `private` and `public`. Defaults to `public`.
	ConnectivityType pulumi.StringPtrInput
	// The private IPv4 address to assign to the NAT Gateway. If you don't provide an address, a private IPv4 address will be automatically assigned.
	PrivateIp pulumi.StringPtrInput
	// A list of secondary allocation EIP IDs for this NAT Gateway.
	SecondaryAllocationIds pulumi.StringArrayInput
	// [Private NAT Gateway only] The number of secondary private IPv4 addresses you want to assign to the NAT Gateway.
	SecondaryPrivateIpAddressCount pulumi.IntPtrInput
	// A list of secondary private IPv4 addresses to assign to the NAT Gateway.
	SecondaryPrivateIpAddresses pulumi.StringArrayInput
	// The Subnet ID of the subnet in which to place the NAT Gateway.
	SubnetId pulumi.StringInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (NatGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*natGatewayArgs)(nil)).Elem()
}

type NatGatewayInput interface {
	pulumi.Input

	ToNatGatewayOutput() NatGatewayOutput
	ToNatGatewayOutputWithContext(ctx context.Context) NatGatewayOutput
}

func (*NatGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**NatGateway)(nil)).Elem()
}

func (i *NatGateway) ToNatGatewayOutput() NatGatewayOutput {
	return i.ToNatGatewayOutputWithContext(context.Background())
}

func (i *NatGateway) ToNatGatewayOutputWithContext(ctx context.Context) NatGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatGatewayOutput)
}

// NatGatewayArrayInput is an input type that accepts NatGatewayArray and NatGatewayArrayOutput values.
// You can construct a concrete instance of `NatGatewayArrayInput` via:
//
//	NatGatewayArray{ NatGatewayArgs{...} }
type NatGatewayArrayInput interface {
	pulumi.Input

	ToNatGatewayArrayOutput() NatGatewayArrayOutput
	ToNatGatewayArrayOutputWithContext(context.Context) NatGatewayArrayOutput
}

type NatGatewayArray []NatGatewayInput

func (NatGatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NatGateway)(nil)).Elem()
}

func (i NatGatewayArray) ToNatGatewayArrayOutput() NatGatewayArrayOutput {
	return i.ToNatGatewayArrayOutputWithContext(context.Background())
}

func (i NatGatewayArray) ToNatGatewayArrayOutputWithContext(ctx context.Context) NatGatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatGatewayArrayOutput)
}

// NatGatewayMapInput is an input type that accepts NatGatewayMap and NatGatewayMapOutput values.
// You can construct a concrete instance of `NatGatewayMapInput` via:
//
//	NatGatewayMap{ "key": NatGatewayArgs{...} }
type NatGatewayMapInput interface {
	pulumi.Input

	ToNatGatewayMapOutput() NatGatewayMapOutput
	ToNatGatewayMapOutputWithContext(context.Context) NatGatewayMapOutput
}

type NatGatewayMap map[string]NatGatewayInput

func (NatGatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NatGateway)(nil)).Elem()
}

func (i NatGatewayMap) ToNatGatewayMapOutput() NatGatewayMapOutput {
	return i.ToNatGatewayMapOutputWithContext(context.Background())
}

func (i NatGatewayMap) ToNatGatewayMapOutputWithContext(ctx context.Context) NatGatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatGatewayMapOutput)
}

type NatGatewayOutput struct{ *pulumi.OutputState }

func (NatGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NatGateway)(nil)).Elem()
}

func (o NatGatewayOutput) ToNatGatewayOutput() NatGatewayOutput {
	return o
}

func (o NatGatewayOutput) ToNatGatewayOutputWithContext(ctx context.Context) NatGatewayOutput {
	return o
}

// The Allocation ID of the Elastic IP address for the NAT Gateway. Required for `connectivityType` of `public`.
func (o NatGatewayOutput) AllocationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringPtrOutput { return v.AllocationId }).(pulumi.StringPtrOutput)
}

// The association ID of the Elastic IP address that's associated with the NAT Gateway. Only available when `connectivityType` is `public`.
func (o NatGatewayOutput) AssociationId() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.AssociationId }).(pulumi.StringOutput)
}

// Connectivity type for the NAT Gateway. Valid values are `private` and `public`. Defaults to `public`.
func (o NatGatewayOutput) ConnectivityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringPtrOutput { return v.ConnectivityType }).(pulumi.StringPtrOutput)
}

// The ID of the network interface associated with the NAT Gateway.
func (o NatGatewayOutput) NetworkInterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.NetworkInterfaceId }).(pulumi.StringOutput)
}

// The private IPv4 address to assign to the NAT Gateway. If you don't provide an address, a private IPv4 address will be automatically assigned.
func (o NatGatewayOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.PrivateIp }).(pulumi.StringOutput)
}

// The Elastic IP address associated with the NAT Gateway.
func (o NatGatewayOutput) PublicIp() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.PublicIp }).(pulumi.StringOutput)
}

// A list of secondary allocation EIP IDs for this NAT Gateway.
func (o NatGatewayOutput) SecondaryAllocationIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringArrayOutput { return v.SecondaryAllocationIds }).(pulumi.StringArrayOutput)
}

// [Private NAT Gateway only] The number of secondary private IPv4 addresses you want to assign to the NAT Gateway.
func (o NatGatewayOutput) SecondaryPrivateIpAddressCount() pulumi.IntOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.IntOutput { return v.SecondaryPrivateIpAddressCount }).(pulumi.IntOutput)
}

// A list of secondary private IPv4 addresses to assign to the NAT Gateway.
func (o NatGatewayOutput) SecondaryPrivateIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringArrayOutput { return v.SecondaryPrivateIpAddresses }).(pulumi.StringArrayOutput)
}

// The Subnet ID of the subnet in which to place the NAT Gateway.
func (o NatGatewayOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o NatGatewayOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o NatGatewayOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type NatGatewayArrayOutput struct{ *pulumi.OutputState }

func (NatGatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NatGateway)(nil)).Elem()
}

func (o NatGatewayArrayOutput) ToNatGatewayArrayOutput() NatGatewayArrayOutput {
	return o
}

func (o NatGatewayArrayOutput) ToNatGatewayArrayOutputWithContext(ctx context.Context) NatGatewayArrayOutput {
	return o
}

func (o NatGatewayArrayOutput) Index(i pulumi.IntInput) NatGatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NatGateway {
		return vs[0].([]*NatGateway)[vs[1].(int)]
	}).(NatGatewayOutput)
}

type NatGatewayMapOutput struct{ *pulumi.OutputState }

func (NatGatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NatGateway)(nil)).Elem()
}

func (o NatGatewayMapOutput) ToNatGatewayMapOutput() NatGatewayMapOutput {
	return o
}

func (o NatGatewayMapOutput) ToNatGatewayMapOutputWithContext(ctx context.Context) NatGatewayMapOutput {
	return o
}

func (o NatGatewayMapOutput) MapIndex(k pulumi.StringInput) NatGatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NatGateway {
		return vs[0].(map[string]*NatGateway)[vs[1].(string)]
	}).(NatGatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NatGatewayInput)(nil)).Elem(), &NatGateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*NatGatewayArrayInput)(nil)).Elem(), NatGatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NatGatewayMapInput)(nil)).Elem(), NatGatewayMap{})
	pulumi.RegisterOutputType(NatGatewayOutput{})
	pulumi.RegisterOutputType(NatGatewayArrayOutput{})
	pulumi.RegisterOutputType(NatGatewayMapOutput{})
}
