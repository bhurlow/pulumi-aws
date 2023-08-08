// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connect

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Amazon Connect User Hierarchy Structure resource. For more information see
// [Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html)
//
// ## Example Usage
// ### Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := connect.NewUserHierarchyStructure(ctx, "example", &connect.UserHierarchyStructureArgs{
//				HierarchyStructure: &connect.UserHierarchyStructureHierarchyStructureArgs{
//					LevelOne: &connect.UserHierarchyStructureHierarchyStructureLevelOneArgs{
//						Name: pulumi.String("levelone"),
//					},
//				},
//				InstanceId: pulumi.String("aaaaaaaa-bbbb-cccc-dddd-111111111111"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### With Five Levels
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := connect.NewUserHierarchyStructure(ctx, "example", &connect.UserHierarchyStructureArgs{
//				HierarchyStructure: &connect.UserHierarchyStructureHierarchyStructureArgs{
//					LevelFive: &connect.UserHierarchyStructureHierarchyStructureLevelFiveArgs{
//						Name: pulumi.String("levelfive"),
//					},
//					LevelFour: &connect.UserHierarchyStructureHierarchyStructureLevelFourArgs{
//						Name: pulumi.String("levelfour"),
//					},
//					LevelOne: &connect.UserHierarchyStructureHierarchyStructureLevelOneArgs{
//						Name: pulumi.String("levelone"),
//					},
//					LevelThree: &connect.UserHierarchyStructureHierarchyStructureLevelThreeArgs{
//						Name: pulumi.String("levelthree"),
//					},
//					LevelTwo: &connect.UserHierarchyStructureHierarchyStructureLevelTwoArgs{
//						Name: pulumi.String("leveltwo"),
//					},
//				},
//				InstanceId: pulumi.String("aaaaaaaa-bbbb-cccc-dddd-111111111111"),
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
//	to = aws_connect_user_hierarchy_structure.example
//
//	id = "f1288a1f-6193-445a-b47e-af739b2" } Using `pulumi import`, import Amazon Connect User Hierarchy Structures using the `instance_id`. For exampleconsole % pulumi import aws_connect_user_hierarchy_structure.example f1288a1f-6193-445a-b47e-af739b2
type UserHierarchyStructure struct {
	pulumi.CustomResourceState

	// A block that defines the hierarchy structure's levels. The `hierarchyStructure` block is documented below.
	HierarchyStructure UserHierarchyStructureHierarchyStructureOutput `pulumi:"hierarchyStructure"`
	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewUserHierarchyStructure registers a new resource with the given unique name, arguments, and options.
func NewUserHierarchyStructure(ctx *pulumi.Context,
	name string, args *UserHierarchyStructureArgs, opts ...pulumi.ResourceOption) (*UserHierarchyStructure, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HierarchyStructure == nil {
		return nil, errors.New("invalid value for required argument 'HierarchyStructure'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserHierarchyStructure
	err := ctx.RegisterResource("aws:connect/userHierarchyStructure:UserHierarchyStructure", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserHierarchyStructure gets an existing UserHierarchyStructure resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserHierarchyStructure(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserHierarchyStructureState, opts ...pulumi.ResourceOption) (*UserHierarchyStructure, error) {
	var resource UserHierarchyStructure
	err := ctx.ReadResource("aws:connect/userHierarchyStructure:UserHierarchyStructure", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserHierarchyStructure resources.
type userHierarchyStructureState struct {
	// A block that defines the hierarchy structure's levels. The `hierarchyStructure` block is documented below.
	HierarchyStructure *UserHierarchyStructureHierarchyStructure `pulumi:"hierarchyStructure"`
	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceId *string `pulumi:"instanceId"`
}

type UserHierarchyStructureState struct {
	// A block that defines the hierarchy structure's levels. The `hierarchyStructure` block is documented below.
	HierarchyStructure UserHierarchyStructureHierarchyStructurePtrInput
	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceId pulumi.StringPtrInput
}

func (UserHierarchyStructureState) ElementType() reflect.Type {
	return reflect.TypeOf((*userHierarchyStructureState)(nil)).Elem()
}

type userHierarchyStructureArgs struct {
	// A block that defines the hierarchy structure's levels. The `hierarchyStructure` block is documented below.
	HierarchyStructure UserHierarchyStructureHierarchyStructure `pulumi:"hierarchyStructure"`
	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a UserHierarchyStructure resource.
type UserHierarchyStructureArgs struct {
	// A block that defines the hierarchy structure's levels. The `hierarchyStructure` block is documented below.
	HierarchyStructure UserHierarchyStructureHierarchyStructureInput
	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceId pulumi.StringInput
}

func (UserHierarchyStructureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userHierarchyStructureArgs)(nil)).Elem()
}

type UserHierarchyStructureInput interface {
	pulumi.Input

	ToUserHierarchyStructureOutput() UserHierarchyStructureOutput
	ToUserHierarchyStructureOutputWithContext(ctx context.Context) UserHierarchyStructureOutput
}

func (*UserHierarchyStructure) ElementType() reflect.Type {
	return reflect.TypeOf((**UserHierarchyStructure)(nil)).Elem()
}

func (i *UserHierarchyStructure) ToUserHierarchyStructureOutput() UserHierarchyStructureOutput {
	return i.ToUserHierarchyStructureOutputWithContext(context.Background())
}

func (i *UserHierarchyStructure) ToUserHierarchyStructureOutputWithContext(ctx context.Context) UserHierarchyStructureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserHierarchyStructureOutput)
}

// UserHierarchyStructureArrayInput is an input type that accepts UserHierarchyStructureArray and UserHierarchyStructureArrayOutput values.
// You can construct a concrete instance of `UserHierarchyStructureArrayInput` via:
//
//	UserHierarchyStructureArray{ UserHierarchyStructureArgs{...} }
type UserHierarchyStructureArrayInput interface {
	pulumi.Input

	ToUserHierarchyStructureArrayOutput() UserHierarchyStructureArrayOutput
	ToUserHierarchyStructureArrayOutputWithContext(context.Context) UserHierarchyStructureArrayOutput
}

type UserHierarchyStructureArray []UserHierarchyStructureInput

func (UserHierarchyStructureArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserHierarchyStructure)(nil)).Elem()
}

func (i UserHierarchyStructureArray) ToUserHierarchyStructureArrayOutput() UserHierarchyStructureArrayOutput {
	return i.ToUserHierarchyStructureArrayOutputWithContext(context.Background())
}

func (i UserHierarchyStructureArray) ToUserHierarchyStructureArrayOutputWithContext(ctx context.Context) UserHierarchyStructureArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserHierarchyStructureArrayOutput)
}

// UserHierarchyStructureMapInput is an input type that accepts UserHierarchyStructureMap and UserHierarchyStructureMapOutput values.
// You can construct a concrete instance of `UserHierarchyStructureMapInput` via:
//
//	UserHierarchyStructureMap{ "key": UserHierarchyStructureArgs{...} }
type UserHierarchyStructureMapInput interface {
	pulumi.Input

	ToUserHierarchyStructureMapOutput() UserHierarchyStructureMapOutput
	ToUserHierarchyStructureMapOutputWithContext(context.Context) UserHierarchyStructureMapOutput
}

type UserHierarchyStructureMap map[string]UserHierarchyStructureInput

func (UserHierarchyStructureMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserHierarchyStructure)(nil)).Elem()
}

func (i UserHierarchyStructureMap) ToUserHierarchyStructureMapOutput() UserHierarchyStructureMapOutput {
	return i.ToUserHierarchyStructureMapOutputWithContext(context.Background())
}

func (i UserHierarchyStructureMap) ToUserHierarchyStructureMapOutputWithContext(ctx context.Context) UserHierarchyStructureMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserHierarchyStructureMapOutput)
}

type UserHierarchyStructureOutput struct{ *pulumi.OutputState }

func (UserHierarchyStructureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserHierarchyStructure)(nil)).Elem()
}

func (o UserHierarchyStructureOutput) ToUserHierarchyStructureOutput() UserHierarchyStructureOutput {
	return o
}

func (o UserHierarchyStructureOutput) ToUserHierarchyStructureOutputWithContext(ctx context.Context) UserHierarchyStructureOutput {
	return o
}

// A block that defines the hierarchy structure's levels. The `hierarchyStructure` block is documented below.
func (o UserHierarchyStructureOutput) HierarchyStructure() UserHierarchyStructureHierarchyStructureOutput {
	return o.ApplyT(func(v *UserHierarchyStructure) UserHierarchyStructureHierarchyStructureOutput {
		return v.HierarchyStructure
	}).(UserHierarchyStructureHierarchyStructureOutput)
}

// Specifies the identifier of the hosting Amazon Connect Instance.
func (o UserHierarchyStructureOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *UserHierarchyStructure) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type UserHierarchyStructureArrayOutput struct{ *pulumi.OutputState }

func (UserHierarchyStructureArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserHierarchyStructure)(nil)).Elem()
}

func (o UserHierarchyStructureArrayOutput) ToUserHierarchyStructureArrayOutput() UserHierarchyStructureArrayOutput {
	return o
}

func (o UserHierarchyStructureArrayOutput) ToUserHierarchyStructureArrayOutputWithContext(ctx context.Context) UserHierarchyStructureArrayOutput {
	return o
}

func (o UserHierarchyStructureArrayOutput) Index(i pulumi.IntInput) UserHierarchyStructureOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserHierarchyStructure {
		return vs[0].([]*UserHierarchyStructure)[vs[1].(int)]
	}).(UserHierarchyStructureOutput)
}

type UserHierarchyStructureMapOutput struct{ *pulumi.OutputState }

func (UserHierarchyStructureMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserHierarchyStructure)(nil)).Elem()
}

func (o UserHierarchyStructureMapOutput) ToUserHierarchyStructureMapOutput() UserHierarchyStructureMapOutput {
	return o
}

func (o UserHierarchyStructureMapOutput) ToUserHierarchyStructureMapOutputWithContext(ctx context.Context) UserHierarchyStructureMapOutput {
	return o
}

func (o UserHierarchyStructureMapOutput) MapIndex(k pulumi.StringInput) UserHierarchyStructureOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserHierarchyStructure {
		return vs[0].(map[string]*UserHierarchyStructure)[vs[1].(string)]
	}).(UserHierarchyStructureOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserHierarchyStructureInput)(nil)).Elem(), &UserHierarchyStructure{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserHierarchyStructureArrayInput)(nil)).Elem(), UserHierarchyStructureArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserHierarchyStructureMapInput)(nil)).Elem(), UserHierarchyStructureMap{})
	pulumi.RegisterOutputType(UserHierarchyStructureOutput{})
	pulumi.RegisterOutputType(UserHierarchyStructureArrayOutput{})
	pulumi.RegisterOutputType(UserHierarchyStructureMapOutput{})
}
