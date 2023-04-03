// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Glue resource policy. Only one can exist per region.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/glue"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			currentPartition, err := aws.GetPartition(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			currentRegion, err := aws.GetRegion(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			glue_example_policy, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
//				Statements: []iam.GetPolicyDocumentStatement{
//					{
//						Actions: []string{
//							"glue:CreateTable",
//						},
//						Resources: []string{
//							fmt.Sprintf("arn:%v:glue:%v:%v:*", currentPartition.Partition, currentRegion.Name, currentCallerIdentity.AccountId),
//						},
//						Principals: []iam.GetPolicyDocumentStatementPrincipal{
//							{
//								Identifiers: []string{
//									"*",
//								},
//								Type: "AWS",
//							},
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = glue.NewResourcePolicy(ctx, "example", &glue.ResourcePolicyArgs{
//				Policy: *pulumi.String(glue_example_policy.Json),
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
// # Glue Resource Policy can be imported using the account ID
//
// ```sh
//
//	$ pulumi import aws:glue/resourcePolicy:ResourcePolicy Test 12356789012
//
// ```
type ResourcePolicy struct {
	pulumi.CustomResourceState

	// Indicates that you are using both methods to grant cross-account. Valid values are `TRUE` and `FALSE`. Note the provider will not perform drift detetction on this field as its not return on read.
	EnableHybrid pulumi.StringPtrOutput `pulumi:"enableHybrid"`
	// The policy to be applied to the aws glue data catalog.
	Policy pulumi.StringOutput `pulumi:"policy"`
}

// NewResourcePolicy registers a new resource with the given unique name, arguments, and options.
func NewResourcePolicy(ctx *pulumi.Context,
	name string, args *ResourcePolicyArgs, opts ...pulumi.ResourceOption) (*ResourcePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	var resource ResourcePolicy
	err := ctx.RegisterResource("aws:glue/resourcePolicy:ResourcePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourcePolicy gets an existing ResourcePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourcePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourcePolicyState, opts ...pulumi.ResourceOption) (*ResourcePolicy, error) {
	var resource ResourcePolicy
	err := ctx.ReadResource("aws:glue/resourcePolicy:ResourcePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourcePolicy resources.
type resourcePolicyState struct {
	// Indicates that you are using both methods to grant cross-account. Valid values are `TRUE` and `FALSE`. Note the provider will not perform drift detetction on this field as its not return on read.
	EnableHybrid *string `pulumi:"enableHybrid"`
	// The policy to be applied to the aws glue data catalog.
	Policy *string `pulumi:"policy"`
}

type ResourcePolicyState struct {
	// Indicates that you are using both methods to grant cross-account. Valid values are `TRUE` and `FALSE`. Note the provider will not perform drift detetction on this field as its not return on read.
	EnableHybrid pulumi.StringPtrInput
	// The policy to be applied to the aws glue data catalog.
	Policy pulumi.StringPtrInput
}

func (ResourcePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePolicyState)(nil)).Elem()
}

type resourcePolicyArgs struct {
	// Indicates that you are using both methods to grant cross-account. Valid values are `TRUE` and `FALSE`. Note the provider will not perform drift detetction on this field as its not return on read.
	EnableHybrid *string `pulumi:"enableHybrid"`
	// The policy to be applied to the aws glue data catalog.
	Policy string `pulumi:"policy"`
}

// The set of arguments for constructing a ResourcePolicy resource.
type ResourcePolicyArgs struct {
	// Indicates that you are using both methods to grant cross-account. Valid values are `TRUE` and `FALSE`. Note the provider will not perform drift detetction on this field as its not return on read.
	EnableHybrid pulumi.StringPtrInput
	// The policy to be applied to the aws glue data catalog.
	Policy pulumi.StringInput
}

func (ResourcePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePolicyArgs)(nil)).Elem()
}

type ResourcePolicyInput interface {
	pulumi.Input

	ToResourcePolicyOutput() ResourcePolicyOutput
	ToResourcePolicyOutputWithContext(ctx context.Context) ResourcePolicyOutput
}

func (*ResourcePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePolicy)(nil)).Elem()
}

func (i *ResourcePolicy) ToResourcePolicyOutput() ResourcePolicyOutput {
	return i.ToResourcePolicyOutputWithContext(context.Background())
}

func (i *ResourcePolicy) ToResourcePolicyOutputWithContext(ctx context.Context) ResourcePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePolicyOutput)
}

// ResourcePolicyArrayInput is an input type that accepts ResourcePolicyArray and ResourcePolicyArrayOutput values.
// You can construct a concrete instance of `ResourcePolicyArrayInput` via:
//
//	ResourcePolicyArray{ ResourcePolicyArgs{...} }
type ResourcePolicyArrayInput interface {
	pulumi.Input

	ToResourcePolicyArrayOutput() ResourcePolicyArrayOutput
	ToResourcePolicyArrayOutputWithContext(context.Context) ResourcePolicyArrayOutput
}

type ResourcePolicyArray []ResourcePolicyInput

func (ResourcePolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourcePolicy)(nil)).Elem()
}

func (i ResourcePolicyArray) ToResourcePolicyArrayOutput() ResourcePolicyArrayOutput {
	return i.ToResourcePolicyArrayOutputWithContext(context.Background())
}

func (i ResourcePolicyArray) ToResourcePolicyArrayOutputWithContext(ctx context.Context) ResourcePolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePolicyArrayOutput)
}

// ResourcePolicyMapInput is an input type that accepts ResourcePolicyMap and ResourcePolicyMapOutput values.
// You can construct a concrete instance of `ResourcePolicyMapInput` via:
//
//	ResourcePolicyMap{ "key": ResourcePolicyArgs{...} }
type ResourcePolicyMapInput interface {
	pulumi.Input

	ToResourcePolicyMapOutput() ResourcePolicyMapOutput
	ToResourcePolicyMapOutputWithContext(context.Context) ResourcePolicyMapOutput
}

type ResourcePolicyMap map[string]ResourcePolicyInput

func (ResourcePolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourcePolicy)(nil)).Elem()
}

func (i ResourcePolicyMap) ToResourcePolicyMapOutput() ResourcePolicyMapOutput {
	return i.ToResourcePolicyMapOutputWithContext(context.Background())
}

func (i ResourcePolicyMap) ToResourcePolicyMapOutputWithContext(ctx context.Context) ResourcePolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePolicyMapOutput)
}

type ResourcePolicyOutput struct{ *pulumi.OutputState }

func (ResourcePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePolicy)(nil)).Elem()
}

func (o ResourcePolicyOutput) ToResourcePolicyOutput() ResourcePolicyOutput {
	return o
}

func (o ResourcePolicyOutput) ToResourcePolicyOutputWithContext(ctx context.Context) ResourcePolicyOutput {
	return o
}

// Indicates that you are using both methods to grant cross-account. Valid values are `TRUE` and `FALSE`. Note the provider will not perform drift detetction on this field as its not return on read.
func (o ResourcePolicyOutput) EnableHybrid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourcePolicy) pulumi.StringPtrOutput { return v.EnableHybrid }).(pulumi.StringPtrOutput)
}

// The policy to be applied to the aws glue data catalog.
func (o ResourcePolicyOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePolicy) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

type ResourcePolicyArrayOutput struct{ *pulumi.OutputState }

func (ResourcePolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourcePolicy)(nil)).Elem()
}

func (o ResourcePolicyArrayOutput) ToResourcePolicyArrayOutput() ResourcePolicyArrayOutput {
	return o
}

func (o ResourcePolicyArrayOutput) ToResourcePolicyArrayOutputWithContext(ctx context.Context) ResourcePolicyArrayOutput {
	return o
}

func (o ResourcePolicyArrayOutput) Index(i pulumi.IntInput) ResourcePolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourcePolicy {
		return vs[0].([]*ResourcePolicy)[vs[1].(int)]
	}).(ResourcePolicyOutput)
}

type ResourcePolicyMapOutput struct{ *pulumi.OutputState }

func (ResourcePolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourcePolicy)(nil)).Elem()
}

func (o ResourcePolicyMapOutput) ToResourcePolicyMapOutput() ResourcePolicyMapOutput {
	return o
}

func (o ResourcePolicyMapOutput) ToResourcePolicyMapOutputWithContext(ctx context.Context) ResourcePolicyMapOutput {
	return o
}

func (o ResourcePolicyMapOutput) MapIndex(k pulumi.StringInput) ResourcePolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourcePolicy {
		return vs[0].(map[string]*ResourcePolicy)[vs[1].(string)]
	}).(ResourcePolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcePolicyInput)(nil)).Elem(), &ResourcePolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcePolicyArrayInput)(nil)).Elem(), ResourcePolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcePolicyMapInput)(nil)).Elem(), ResourcePolicyMap{})
	pulumi.RegisterOutputType(ResourcePolicyOutput{})
	pulumi.RegisterOutputType(ResourcePolicyArrayOutput{})
	pulumi.RegisterOutputType(ResourcePolicyMapOutput{})
}
