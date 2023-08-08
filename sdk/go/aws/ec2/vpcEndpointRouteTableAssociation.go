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

// Manages a VPC Endpoint Route Table Association
//
// ## Example Usage
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
//			_, err := ec2.NewVpcEndpointRouteTableAssociation(ctx, "example", &ec2.VpcEndpointRouteTableAssociationArgs{
//				RouteTableId:  pulumi.Any(aws_route_table.Example.Id),
//				VpcEndpointId: pulumi.Any(aws_vpc_endpoint.Example.Id),
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
//	to = aws_vpc_endpoint_route_table_association.example
//
//	id = "vpce-aaaaaaaa/rtb-bbbbbbbb" } Using `pulumi import`, import VPC Endpoint Route Table Associations using `vpc_endpoint_id` together with `route_table_id`. For exampleconsole % pulumi import aws_vpc_endpoint_route_table_association.example vpce-aaaaaaaa/rtb-bbbbbbbb
type VpcEndpointRouteTableAssociation struct {
	pulumi.CustomResourceState

	// Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
	RouteTableId pulumi.StringOutput `pulumi:"routeTableId"`
	// Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.
	VpcEndpointId pulumi.StringOutput `pulumi:"vpcEndpointId"`
}

// NewVpcEndpointRouteTableAssociation registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpointRouteTableAssociation(ctx *pulumi.Context,
	name string, args *VpcEndpointRouteTableAssociationArgs, opts ...pulumi.ResourceOption) (*VpcEndpointRouteTableAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RouteTableId == nil {
		return nil, errors.New("invalid value for required argument 'RouteTableId'")
	}
	if args.VpcEndpointId == nil {
		return nil, errors.New("invalid value for required argument 'VpcEndpointId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcEndpointRouteTableAssociation
	err := ctx.RegisterResource("aws:ec2/vpcEndpointRouteTableAssociation:VpcEndpointRouteTableAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcEndpointRouteTableAssociation gets an existing VpcEndpointRouteTableAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcEndpointRouteTableAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcEndpointRouteTableAssociationState, opts ...pulumi.ResourceOption) (*VpcEndpointRouteTableAssociation, error) {
	var resource VpcEndpointRouteTableAssociation
	err := ctx.ReadResource("aws:ec2/vpcEndpointRouteTableAssociation:VpcEndpointRouteTableAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcEndpointRouteTableAssociation resources.
type vpcEndpointRouteTableAssociationState struct {
	// Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
	RouteTableId *string `pulumi:"routeTableId"`
	// Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.
	VpcEndpointId *string `pulumi:"vpcEndpointId"`
}

type VpcEndpointRouteTableAssociationState struct {
	// Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
	RouteTableId pulumi.StringPtrInput
	// Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.
	VpcEndpointId pulumi.StringPtrInput
}

func (VpcEndpointRouteTableAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointRouteTableAssociationState)(nil)).Elem()
}

type vpcEndpointRouteTableAssociationArgs struct {
	// Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
	RouteTableId string `pulumi:"routeTableId"`
	// Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.
	VpcEndpointId string `pulumi:"vpcEndpointId"`
}

// The set of arguments for constructing a VpcEndpointRouteTableAssociation resource.
type VpcEndpointRouteTableAssociationArgs struct {
	// Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
	RouteTableId pulumi.StringInput
	// Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.
	VpcEndpointId pulumi.StringInput
}

func (VpcEndpointRouteTableAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointRouteTableAssociationArgs)(nil)).Elem()
}

type VpcEndpointRouteTableAssociationInput interface {
	pulumi.Input

	ToVpcEndpointRouteTableAssociationOutput() VpcEndpointRouteTableAssociationOutput
	ToVpcEndpointRouteTableAssociationOutputWithContext(ctx context.Context) VpcEndpointRouteTableAssociationOutput
}

func (*VpcEndpointRouteTableAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointRouteTableAssociation)(nil)).Elem()
}

func (i *VpcEndpointRouteTableAssociation) ToVpcEndpointRouteTableAssociationOutput() VpcEndpointRouteTableAssociationOutput {
	return i.ToVpcEndpointRouteTableAssociationOutputWithContext(context.Background())
}

func (i *VpcEndpointRouteTableAssociation) ToVpcEndpointRouteTableAssociationOutputWithContext(ctx context.Context) VpcEndpointRouteTableAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointRouteTableAssociationOutput)
}

// VpcEndpointRouteTableAssociationArrayInput is an input type that accepts VpcEndpointRouteTableAssociationArray and VpcEndpointRouteTableAssociationArrayOutput values.
// You can construct a concrete instance of `VpcEndpointRouteTableAssociationArrayInput` via:
//
//	VpcEndpointRouteTableAssociationArray{ VpcEndpointRouteTableAssociationArgs{...} }
type VpcEndpointRouteTableAssociationArrayInput interface {
	pulumi.Input

	ToVpcEndpointRouteTableAssociationArrayOutput() VpcEndpointRouteTableAssociationArrayOutput
	ToVpcEndpointRouteTableAssociationArrayOutputWithContext(context.Context) VpcEndpointRouteTableAssociationArrayOutput
}

type VpcEndpointRouteTableAssociationArray []VpcEndpointRouteTableAssociationInput

func (VpcEndpointRouteTableAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpointRouteTableAssociation)(nil)).Elem()
}

func (i VpcEndpointRouteTableAssociationArray) ToVpcEndpointRouteTableAssociationArrayOutput() VpcEndpointRouteTableAssociationArrayOutput {
	return i.ToVpcEndpointRouteTableAssociationArrayOutputWithContext(context.Background())
}

func (i VpcEndpointRouteTableAssociationArray) ToVpcEndpointRouteTableAssociationArrayOutputWithContext(ctx context.Context) VpcEndpointRouteTableAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointRouteTableAssociationArrayOutput)
}

// VpcEndpointRouteTableAssociationMapInput is an input type that accepts VpcEndpointRouteTableAssociationMap and VpcEndpointRouteTableAssociationMapOutput values.
// You can construct a concrete instance of `VpcEndpointRouteTableAssociationMapInput` via:
//
//	VpcEndpointRouteTableAssociationMap{ "key": VpcEndpointRouteTableAssociationArgs{...} }
type VpcEndpointRouteTableAssociationMapInput interface {
	pulumi.Input

	ToVpcEndpointRouteTableAssociationMapOutput() VpcEndpointRouteTableAssociationMapOutput
	ToVpcEndpointRouteTableAssociationMapOutputWithContext(context.Context) VpcEndpointRouteTableAssociationMapOutput
}

type VpcEndpointRouteTableAssociationMap map[string]VpcEndpointRouteTableAssociationInput

func (VpcEndpointRouteTableAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpointRouteTableAssociation)(nil)).Elem()
}

func (i VpcEndpointRouteTableAssociationMap) ToVpcEndpointRouteTableAssociationMapOutput() VpcEndpointRouteTableAssociationMapOutput {
	return i.ToVpcEndpointRouteTableAssociationMapOutputWithContext(context.Background())
}

func (i VpcEndpointRouteTableAssociationMap) ToVpcEndpointRouteTableAssociationMapOutputWithContext(ctx context.Context) VpcEndpointRouteTableAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointRouteTableAssociationMapOutput)
}

type VpcEndpointRouteTableAssociationOutput struct{ *pulumi.OutputState }

func (VpcEndpointRouteTableAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointRouteTableAssociation)(nil)).Elem()
}

func (o VpcEndpointRouteTableAssociationOutput) ToVpcEndpointRouteTableAssociationOutput() VpcEndpointRouteTableAssociationOutput {
	return o
}

func (o VpcEndpointRouteTableAssociationOutput) ToVpcEndpointRouteTableAssociationOutputWithContext(ctx context.Context) VpcEndpointRouteTableAssociationOutput {
	return o
}

// Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
func (o VpcEndpointRouteTableAssociationOutput) RouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointRouteTableAssociation) pulumi.StringOutput { return v.RouteTableId }).(pulumi.StringOutput)
}

// Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.
func (o VpcEndpointRouteTableAssociationOutput) VpcEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointRouteTableAssociation) pulumi.StringOutput { return v.VpcEndpointId }).(pulumi.StringOutput)
}

type VpcEndpointRouteTableAssociationArrayOutput struct{ *pulumi.OutputState }

func (VpcEndpointRouteTableAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpointRouteTableAssociation)(nil)).Elem()
}

func (o VpcEndpointRouteTableAssociationArrayOutput) ToVpcEndpointRouteTableAssociationArrayOutput() VpcEndpointRouteTableAssociationArrayOutput {
	return o
}

func (o VpcEndpointRouteTableAssociationArrayOutput) ToVpcEndpointRouteTableAssociationArrayOutputWithContext(ctx context.Context) VpcEndpointRouteTableAssociationArrayOutput {
	return o
}

func (o VpcEndpointRouteTableAssociationArrayOutput) Index(i pulumi.IntInput) VpcEndpointRouteTableAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcEndpointRouteTableAssociation {
		return vs[0].([]*VpcEndpointRouteTableAssociation)[vs[1].(int)]
	}).(VpcEndpointRouteTableAssociationOutput)
}

type VpcEndpointRouteTableAssociationMapOutput struct{ *pulumi.OutputState }

func (VpcEndpointRouteTableAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpointRouteTableAssociation)(nil)).Elem()
}

func (o VpcEndpointRouteTableAssociationMapOutput) ToVpcEndpointRouteTableAssociationMapOutput() VpcEndpointRouteTableAssociationMapOutput {
	return o
}

func (o VpcEndpointRouteTableAssociationMapOutput) ToVpcEndpointRouteTableAssociationMapOutputWithContext(ctx context.Context) VpcEndpointRouteTableAssociationMapOutput {
	return o
}

func (o VpcEndpointRouteTableAssociationMapOutput) MapIndex(k pulumi.StringInput) VpcEndpointRouteTableAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcEndpointRouteTableAssociation {
		return vs[0].(map[string]*VpcEndpointRouteTableAssociation)[vs[1].(string)]
	}).(VpcEndpointRouteTableAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointRouteTableAssociationInput)(nil)).Elem(), &VpcEndpointRouteTableAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointRouteTableAssociationArrayInput)(nil)).Elem(), VpcEndpointRouteTableAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointRouteTableAssociationMapInput)(nil)).Elem(), VpcEndpointRouteTableAssociationMap{})
	pulumi.RegisterOutputType(VpcEndpointRouteTableAssociationOutput{})
	pulumi.RegisterOutputType(VpcEndpointRouteTableAssociationArrayOutput{})
	pulumi.RegisterOutputType(VpcEndpointRouteTableAssociationMapOutput{})
}
