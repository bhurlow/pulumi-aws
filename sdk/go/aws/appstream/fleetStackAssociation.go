// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appstream

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an AppStream Fleet Stack association.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appstream"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleFleet, err := appstream.NewFleet(ctx, "exampleFleet", &appstream.FleetArgs{
//				ImageName:    pulumi.String("Amazon-AppStream2-Sample-Image-02-04-2019"),
//				InstanceType: pulumi.String("stream.standard.small"),
//				ComputeCapacity: &appstream.FleetComputeCapacityArgs{
//					DesiredInstances: pulumi.Int(1),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleStack, err := appstream.NewStack(ctx, "exampleStack", nil)
//			if err != nil {
//				return err
//			}
//			_, err = appstream.NewFleetStackAssociation(ctx, "exampleFleetStackAssociation", &appstream.FleetStackAssociationArgs{
//				FleetName: exampleFleet.Name,
//				StackName: exampleStack.Name,
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
//	to = aws_appstream_fleet_stack_association.example
//
//	id = "fleetName/stackName" } Using `pulumi import`, import AppStream Stack Fleet Association using the `fleet_name` and `stack_name` separated by a slash (`/`). For exampleconsole % pulumi import aws_appstream_fleet_stack_association.example fleetName/stackName
type FleetStackAssociation struct {
	pulumi.CustomResourceState

	// Name of the fleet.
	FleetName pulumi.StringOutput `pulumi:"fleetName"`
	// Name of the stack.
	StackName pulumi.StringOutput `pulumi:"stackName"`
}

// NewFleetStackAssociation registers a new resource with the given unique name, arguments, and options.
func NewFleetStackAssociation(ctx *pulumi.Context,
	name string, args *FleetStackAssociationArgs, opts ...pulumi.ResourceOption) (*FleetStackAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FleetName == nil {
		return nil, errors.New("invalid value for required argument 'FleetName'")
	}
	if args.StackName == nil {
		return nil, errors.New("invalid value for required argument 'StackName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FleetStackAssociation
	err := ctx.RegisterResource("aws:appstream/fleetStackAssociation:FleetStackAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFleetStackAssociation gets an existing FleetStackAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFleetStackAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FleetStackAssociationState, opts ...pulumi.ResourceOption) (*FleetStackAssociation, error) {
	var resource FleetStackAssociation
	err := ctx.ReadResource("aws:appstream/fleetStackAssociation:FleetStackAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FleetStackAssociation resources.
type fleetStackAssociationState struct {
	// Name of the fleet.
	FleetName *string `pulumi:"fleetName"`
	// Name of the stack.
	StackName *string `pulumi:"stackName"`
}

type FleetStackAssociationState struct {
	// Name of the fleet.
	FleetName pulumi.StringPtrInput
	// Name of the stack.
	StackName pulumi.StringPtrInput
}

func (FleetStackAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetStackAssociationState)(nil)).Elem()
}

type fleetStackAssociationArgs struct {
	// Name of the fleet.
	FleetName string `pulumi:"fleetName"`
	// Name of the stack.
	StackName string `pulumi:"stackName"`
}

// The set of arguments for constructing a FleetStackAssociation resource.
type FleetStackAssociationArgs struct {
	// Name of the fleet.
	FleetName pulumi.StringInput
	// Name of the stack.
	StackName pulumi.StringInput
}

func (FleetStackAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fleetStackAssociationArgs)(nil)).Elem()
}

type FleetStackAssociationInput interface {
	pulumi.Input

	ToFleetStackAssociationOutput() FleetStackAssociationOutput
	ToFleetStackAssociationOutputWithContext(ctx context.Context) FleetStackAssociationOutput
}

func (*FleetStackAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**FleetStackAssociation)(nil)).Elem()
}

func (i *FleetStackAssociation) ToFleetStackAssociationOutput() FleetStackAssociationOutput {
	return i.ToFleetStackAssociationOutputWithContext(context.Background())
}

func (i *FleetStackAssociation) ToFleetStackAssociationOutputWithContext(ctx context.Context) FleetStackAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetStackAssociationOutput)
}

// FleetStackAssociationArrayInput is an input type that accepts FleetStackAssociationArray and FleetStackAssociationArrayOutput values.
// You can construct a concrete instance of `FleetStackAssociationArrayInput` via:
//
//	FleetStackAssociationArray{ FleetStackAssociationArgs{...} }
type FleetStackAssociationArrayInput interface {
	pulumi.Input

	ToFleetStackAssociationArrayOutput() FleetStackAssociationArrayOutput
	ToFleetStackAssociationArrayOutputWithContext(context.Context) FleetStackAssociationArrayOutput
}

type FleetStackAssociationArray []FleetStackAssociationInput

func (FleetStackAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FleetStackAssociation)(nil)).Elem()
}

func (i FleetStackAssociationArray) ToFleetStackAssociationArrayOutput() FleetStackAssociationArrayOutput {
	return i.ToFleetStackAssociationArrayOutputWithContext(context.Background())
}

func (i FleetStackAssociationArray) ToFleetStackAssociationArrayOutputWithContext(ctx context.Context) FleetStackAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetStackAssociationArrayOutput)
}

// FleetStackAssociationMapInput is an input type that accepts FleetStackAssociationMap and FleetStackAssociationMapOutput values.
// You can construct a concrete instance of `FleetStackAssociationMapInput` via:
//
//	FleetStackAssociationMap{ "key": FleetStackAssociationArgs{...} }
type FleetStackAssociationMapInput interface {
	pulumi.Input

	ToFleetStackAssociationMapOutput() FleetStackAssociationMapOutput
	ToFleetStackAssociationMapOutputWithContext(context.Context) FleetStackAssociationMapOutput
}

type FleetStackAssociationMap map[string]FleetStackAssociationInput

func (FleetStackAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FleetStackAssociation)(nil)).Elem()
}

func (i FleetStackAssociationMap) ToFleetStackAssociationMapOutput() FleetStackAssociationMapOutput {
	return i.ToFleetStackAssociationMapOutputWithContext(context.Background())
}

func (i FleetStackAssociationMap) ToFleetStackAssociationMapOutputWithContext(ctx context.Context) FleetStackAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FleetStackAssociationMapOutput)
}

type FleetStackAssociationOutput struct{ *pulumi.OutputState }

func (FleetStackAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FleetStackAssociation)(nil)).Elem()
}

func (o FleetStackAssociationOutput) ToFleetStackAssociationOutput() FleetStackAssociationOutput {
	return o
}

func (o FleetStackAssociationOutput) ToFleetStackAssociationOutputWithContext(ctx context.Context) FleetStackAssociationOutput {
	return o
}

// Name of the fleet.
func (o FleetStackAssociationOutput) FleetName() pulumi.StringOutput {
	return o.ApplyT(func(v *FleetStackAssociation) pulumi.StringOutput { return v.FleetName }).(pulumi.StringOutput)
}

// Name of the stack.
func (o FleetStackAssociationOutput) StackName() pulumi.StringOutput {
	return o.ApplyT(func(v *FleetStackAssociation) pulumi.StringOutput { return v.StackName }).(pulumi.StringOutput)
}

type FleetStackAssociationArrayOutput struct{ *pulumi.OutputState }

func (FleetStackAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FleetStackAssociation)(nil)).Elem()
}

func (o FleetStackAssociationArrayOutput) ToFleetStackAssociationArrayOutput() FleetStackAssociationArrayOutput {
	return o
}

func (o FleetStackAssociationArrayOutput) ToFleetStackAssociationArrayOutputWithContext(ctx context.Context) FleetStackAssociationArrayOutput {
	return o
}

func (o FleetStackAssociationArrayOutput) Index(i pulumi.IntInput) FleetStackAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FleetStackAssociation {
		return vs[0].([]*FleetStackAssociation)[vs[1].(int)]
	}).(FleetStackAssociationOutput)
}

type FleetStackAssociationMapOutput struct{ *pulumi.OutputState }

func (FleetStackAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FleetStackAssociation)(nil)).Elem()
}

func (o FleetStackAssociationMapOutput) ToFleetStackAssociationMapOutput() FleetStackAssociationMapOutput {
	return o
}

func (o FleetStackAssociationMapOutput) ToFleetStackAssociationMapOutputWithContext(ctx context.Context) FleetStackAssociationMapOutput {
	return o
}

func (o FleetStackAssociationMapOutput) MapIndex(k pulumi.StringInput) FleetStackAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FleetStackAssociation {
		return vs[0].(map[string]*FleetStackAssociation)[vs[1].(string)]
	}).(FleetStackAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FleetStackAssociationInput)(nil)).Elem(), &FleetStackAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*FleetStackAssociationArrayInput)(nil)).Elem(), FleetStackAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FleetStackAssociationMapInput)(nil)).Elem(), FleetStackAssociationMap{})
	pulumi.RegisterOutputType(FleetStackAssociationOutput{})
	pulumi.RegisterOutputType(FleetStackAssociationArrayOutput{})
	pulumi.RegisterOutputType(FleetStackAssociationMapOutput{})
}
