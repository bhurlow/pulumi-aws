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

// Provides a subnet CIDR reservation resource.
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
//			_, err := ec2.NewSubnetCidrReservation(ctx, "example", &ec2.SubnetCidrReservationArgs{
//				CidrBlock:       pulumi.String("10.0.0.16/28"),
//				ReservationType: pulumi.String("prefix"),
//				SubnetId:        pulumi.Any(aws_subnet.Example.Id),
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
//	to = aws_ec2_subnet_cidr_reservation.example
//
//	id = "subnet-01llsxvsxabqiymcz:scr-4mnvz6wb7otksjcs9" } Using `pulumi import`, import Existing CIDR reservations using `SUBNET_ID:RESERVATION_ID`. For exampleconsole % pulumi import aws_ec2_subnet_cidr_reservation.example subnet-01llsxvsxabqiymcz:scr-4mnvz6wb7otksjcs9
type SubnetCidrReservation struct {
	pulumi.CustomResourceState

	// The CIDR block for the reservation.
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// A brief description of the reservation.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// ID of the AWS account that owns this CIDR reservation.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// The type of reservation to create. Valid values: `explicit`, `prefix`
	ReservationType pulumi.StringOutput `pulumi:"reservationType"`
	// The ID of the subnet to create the reservation for.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
}

// NewSubnetCidrReservation registers a new resource with the given unique name, arguments, and options.
func NewSubnetCidrReservation(ctx *pulumi.Context,
	name string, args *SubnetCidrReservationArgs, opts ...pulumi.ResourceOption) (*SubnetCidrReservation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'CidrBlock'")
	}
	if args.ReservationType == nil {
		return nil, errors.New("invalid value for required argument 'ReservationType'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SubnetCidrReservation
	err := ctx.RegisterResource("aws:ec2/subnetCidrReservation:SubnetCidrReservation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnetCidrReservation gets an existing SubnetCidrReservation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetCidrReservation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetCidrReservationState, opts ...pulumi.ResourceOption) (*SubnetCidrReservation, error) {
	var resource SubnetCidrReservation
	err := ctx.ReadResource("aws:ec2/subnetCidrReservation:SubnetCidrReservation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubnetCidrReservation resources.
type subnetCidrReservationState struct {
	// The CIDR block for the reservation.
	CidrBlock *string `pulumi:"cidrBlock"`
	// A brief description of the reservation.
	Description *string `pulumi:"description"`
	// ID of the AWS account that owns this CIDR reservation.
	OwnerId *string `pulumi:"ownerId"`
	// The type of reservation to create. Valid values: `explicit`, `prefix`
	ReservationType *string `pulumi:"reservationType"`
	// The ID of the subnet to create the reservation for.
	SubnetId *string `pulumi:"subnetId"`
}

type SubnetCidrReservationState struct {
	// The CIDR block for the reservation.
	CidrBlock pulumi.StringPtrInput
	// A brief description of the reservation.
	Description pulumi.StringPtrInput
	// ID of the AWS account that owns this CIDR reservation.
	OwnerId pulumi.StringPtrInput
	// The type of reservation to create. Valid values: `explicit`, `prefix`
	ReservationType pulumi.StringPtrInput
	// The ID of the subnet to create the reservation for.
	SubnetId pulumi.StringPtrInput
}

func (SubnetCidrReservationState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetCidrReservationState)(nil)).Elem()
}

type subnetCidrReservationArgs struct {
	// The CIDR block for the reservation.
	CidrBlock string `pulumi:"cidrBlock"`
	// A brief description of the reservation.
	Description *string `pulumi:"description"`
	// The type of reservation to create. Valid values: `explicit`, `prefix`
	ReservationType string `pulumi:"reservationType"`
	// The ID of the subnet to create the reservation for.
	SubnetId string `pulumi:"subnetId"`
}

// The set of arguments for constructing a SubnetCidrReservation resource.
type SubnetCidrReservationArgs struct {
	// The CIDR block for the reservation.
	CidrBlock pulumi.StringInput
	// A brief description of the reservation.
	Description pulumi.StringPtrInput
	// The type of reservation to create. Valid values: `explicit`, `prefix`
	ReservationType pulumi.StringInput
	// The ID of the subnet to create the reservation for.
	SubnetId pulumi.StringInput
}

func (SubnetCidrReservationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetCidrReservationArgs)(nil)).Elem()
}

type SubnetCidrReservationInput interface {
	pulumi.Input

	ToSubnetCidrReservationOutput() SubnetCidrReservationOutput
	ToSubnetCidrReservationOutputWithContext(ctx context.Context) SubnetCidrReservationOutput
}

func (*SubnetCidrReservation) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetCidrReservation)(nil)).Elem()
}

func (i *SubnetCidrReservation) ToSubnetCidrReservationOutput() SubnetCidrReservationOutput {
	return i.ToSubnetCidrReservationOutputWithContext(context.Background())
}

func (i *SubnetCidrReservation) ToSubnetCidrReservationOutputWithContext(ctx context.Context) SubnetCidrReservationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetCidrReservationOutput)
}

// SubnetCidrReservationArrayInput is an input type that accepts SubnetCidrReservationArray and SubnetCidrReservationArrayOutput values.
// You can construct a concrete instance of `SubnetCidrReservationArrayInput` via:
//
//	SubnetCidrReservationArray{ SubnetCidrReservationArgs{...} }
type SubnetCidrReservationArrayInput interface {
	pulumi.Input

	ToSubnetCidrReservationArrayOutput() SubnetCidrReservationArrayOutput
	ToSubnetCidrReservationArrayOutputWithContext(context.Context) SubnetCidrReservationArrayOutput
}

type SubnetCidrReservationArray []SubnetCidrReservationInput

func (SubnetCidrReservationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubnetCidrReservation)(nil)).Elem()
}

func (i SubnetCidrReservationArray) ToSubnetCidrReservationArrayOutput() SubnetCidrReservationArrayOutput {
	return i.ToSubnetCidrReservationArrayOutputWithContext(context.Background())
}

func (i SubnetCidrReservationArray) ToSubnetCidrReservationArrayOutputWithContext(ctx context.Context) SubnetCidrReservationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetCidrReservationArrayOutput)
}

// SubnetCidrReservationMapInput is an input type that accepts SubnetCidrReservationMap and SubnetCidrReservationMapOutput values.
// You can construct a concrete instance of `SubnetCidrReservationMapInput` via:
//
//	SubnetCidrReservationMap{ "key": SubnetCidrReservationArgs{...} }
type SubnetCidrReservationMapInput interface {
	pulumi.Input

	ToSubnetCidrReservationMapOutput() SubnetCidrReservationMapOutput
	ToSubnetCidrReservationMapOutputWithContext(context.Context) SubnetCidrReservationMapOutput
}

type SubnetCidrReservationMap map[string]SubnetCidrReservationInput

func (SubnetCidrReservationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubnetCidrReservation)(nil)).Elem()
}

func (i SubnetCidrReservationMap) ToSubnetCidrReservationMapOutput() SubnetCidrReservationMapOutput {
	return i.ToSubnetCidrReservationMapOutputWithContext(context.Background())
}

func (i SubnetCidrReservationMap) ToSubnetCidrReservationMapOutputWithContext(ctx context.Context) SubnetCidrReservationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetCidrReservationMapOutput)
}

type SubnetCidrReservationOutput struct{ *pulumi.OutputState }

func (SubnetCidrReservationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetCidrReservation)(nil)).Elem()
}

func (o SubnetCidrReservationOutput) ToSubnetCidrReservationOutput() SubnetCidrReservationOutput {
	return o
}

func (o SubnetCidrReservationOutput) ToSubnetCidrReservationOutputWithContext(ctx context.Context) SubnetCidrReservationOutput {
	return o
}

// The CIDR block for the reservation.
func (o SubnetCidrReservationOutput) CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *SubnetCidrReservation) pulumi.StringOutput { return v.CidrBlock }).(pulumi.StringOutput)
}

// A brief description of the reservation.
func (o SubnetCidrReservationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetCidrReservation) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// ID of the AWS account that owns this CIDR reservation.
func (o SubnetCidrReservationOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *SubnetCidrReservation) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// The type of reservation to create. Valid values: `explicit`, `prefix`
func (o SubnetCidrReservationOutput) ReservationType() pulumi.StringOutput {
	return o.ApplyT(func(v *SubnetCidrReservation) pulumi.StringOutput { return v.ReservationType }).(pulumi.StringOutput)
}

// The ID of the subnet to create the reservation for.
func (o SubnetCidrReservationOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *SubnetCidrReservation) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

type SubnetCidrReservationArrayOutput struct{ *pulumi.OutputState }

func (SubnetCidrReservationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubnetCidrReservation)(nil)).Elem()
}

func (o SubnetCidrReservationArrayOutput) ToSubnetCidrReservationArrayOutput() SubnetCidrReservationArrayOutput {
	return o
}

func (o SubnetCidrReservationArrayOutput) ToSubnetCidrReservationArrayOutputWithContext(ctx context.Context) SubnetCidrReservationArrayOutput {
	return o
}

func (o SubnetCidrReservationArrayOutput) Index(i pulumi.IntInput) SubnetCidrReservationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SubnetCidrReservation {
		return vs[0].([]*SubnetCidrReservation)[vs[1].(int)]
	}).(SubnetCidrReservationOutput)
}

type SubnetCidrReservationMapOutput struct{ *pulumi.OutputState }

func (SubnetCidrReservationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubnetCidrReservation)(nil)).Elem()
}

func (o SubnetCidrReservationMapOutput) ToSubnetCidrReservationMapOutput() SubnetCidrReservationMapOutput {
	return o
}

func (o SubnetCidrReservationMapOutput) ToSubnetCidrReservationMapOutputWithContext(ctx context.Context) SubnetCidrReservationMapOutput {
	return o
}

func (o SubnetCidrReservationMapOutput) MapIndex(k pulumi.StringInput) SubnetCidrReservationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SubnetCidrReservation {
		return vs[0].(map[string]*SubnetCidrReservation)[vs[1].(string)]
	}).(SubnetCidrReservationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetCidrReservationInput)(nil)).Elem(), &SubnetCidrReservation{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetCidrReservationArrayInput)(nil)).Elem(), SubnetCidrReservationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetCidrReservationMapInput)(nil)).Elem(), SubnetCidrReservationMap{})
	pulumi.RegisterOutputType(SubnetCidrReservationOutput{})
	pulumi.RegisterOutputType(SubnetCidrReservationArrayOutput{})
	pulumi.RegisterOutputType(SubnetCidrReservationMapOutput{})
}
