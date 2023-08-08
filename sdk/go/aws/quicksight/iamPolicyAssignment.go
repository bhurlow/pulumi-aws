// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for managing an AWS QuickSight IAM Policy Assignment.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := quicksight.NewIamPolicyAssignment(ctx, "example", &quicksight.IamPolicyAssignmentArgs{
//				AssignmentName:   pulumi.String("example"),
//				AssignmentStatus: pulumi.String("ENABLED"),
//				PolicyArn:        pulumi.Any(aws_iam_policy.Example.Arn),
//				Identities: &quicksight.IamPolicyAssignmentIdentitiesArgs{
//					Users: pulumi.StringArray{
//						aws_quicksight_user.Example.User_name,
//					},
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
//	to = aws_quicksight_iam_policy_assignment.example
//
//	id = "123456789012,default,example" } Using `pulumi import`, import QuickSight IAM Policy Assignment using the AWS account ID, namespace, and assignment name separated by commas (`,`). For exampleconsole % pulumi import aws_quicksight_iam_policy_assignment.example 123456789012,default,example
type IamPolicyAssignment struct {
	pulumi.CustomResourceState

	// Assignment ID.
	AssignmentId pulumi.StringOutput `pulumi:"assignmentId"`
	// Name of the assignment.
	AssignmentName pulumi.StringOutput `pulumi:"assignmentName"`
	// Status of the assignment. Valid values are `ENABLED`, `DISABLED`, and `DRAFT`.
	//
	// The following arguments are optional:
	AssignmentStatus pulumi.StringOutput `pulumi:"assignmentStatus"`
	// AWS account ID.
	AwsAccountId pulumi.StringOutput `pulumi:"awsAccountId"`
	// Amazon QuickSight users, groups, or both to assign the policy to. See `identities`.
	Identities IamPolicyAssignmentIdentitiesPtrOutput `pulumi:"identities"`
	// Namespace that contains the assignment. Defaults to `default`.
	Namespace pulumi.StringOutput `pulumi:"namespace"`
	// ARN of the IAM policy to apply to the Amazon QuickSight users and groups specified in this assignment.
	PolicyArn pulumi.StringPtrOutput `pulumi:"policyArn"`
}

// NewIamPolicyAssignment registers a new resource with the given unique name, arguments, and options.
func NewIamPolicyAssignment(ctx *pulumi.Context,
	name string, args *IamPolicyAssignmentArgs, opts ...pulumi.ResourceOption) (*IamPolicyAssignment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssignmentName == nil {
		return nil, errors.New("invalid value for required argument 'AssignmentName'")
	}
	if args.AssignmentStatus == nil {
		return nil, errors.New("invalid value for required argument 'AssignmentStatus'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IamPolicyAssignment
	err := ctx.RegisterResource("aws:quicksight/iamPolicyAssignment:IamPolicyAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamPolicyAssignment gets an existing IamPolicyAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamPolicyAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamPolicyAssignmentState, opts ...pulumi.ResourceOption) (*IamPolicyAssignment, error) {
	var resource IamPolicyAssignment
	err := ctx.ReadResource("aws:quicksight/iamPolicyAssignment:IamPolicyAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamPolicyAssignment resources.
type iamPolicyAssignmentState struct {
	// Assignment ID.
	AssignmentId *string `pulumi:"assignmentId"`
	// Name of the assignment.
	AssignmentName *string `pulumi:"assignmentName"`
	// Status of the assignment. Valid values are `ENABLED`, `DISABLED`, and `DRAFT`.
	//
	// The following arguments are optional:
	AssignmentStatus *string `pulumi:"assignmentStatus"`
	// AWS account ID.
	AwsAccountId *string `pulumi:"awsAccountId"`
	// Amazon QuickSight users, groups, or both to assign the policy to. See `identities`.
	Identities *IamPolicyAssignmentIdentities `pulumi:"identities"`
	// Namespace that contains the assignment. Defaults to `default`.
	Namespace *string `pulumi:"namespace"`
	// ARN of the IAM policy to apply to the Amazon QuickSight users and groups specified in this assignment.
	PolicyArn *string `pulumi:"policyArn"`
}

type IamPolicyAssignmentState struct {
	// Assignment ID.
	AssignmentId pulumi.StringPtrInput
	// Name of the assignment.
	AssignmentName pulumi.StringPtrInput
	// Status of the assignment. Valid values are `ENABLED`, `DISABLED`, and `DRAFT`.
	//
	// The following arguments are optional:
	AssignmentStatus pulumi.StringPtrInput
	// AWS account ID.
	AwsAccountId pulumi.StringPtrInput
	// Amazon QuickSight users, groups, or both to assign the policy to. See `identities`.
	Identities IamPolicyAssignmentIdentitiesPtrInput
	// Namespace that contains the assignment. Defaults to `default`.
	Namespace pulumi.StringPtrInput
	// ARN of the IAM policy to apply to the Amazon QuickSight users and groups specified in this assignment.
	PolicyArn pulumi.StringPtrInput
}

func (IamPolicyAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamPolicyAssignmentState)(nil)).Elem()
}

type iamPolicyAssignmentArgs struct {
	// Name of the assignment.
	AssignmentName string `pulumi:"assignmentName"`
	// Status of the assignment. Valid values are `ENABLED`, `DISABLED`, and `DRAFT`.
	//
	// The following arguments are optional:
	AssignmentStatus string `pulumi:"assignmentStatus"`
	// AWS account ID.
	AwsAccountId *string `pulumi:"awsAccountId"`
	// Amazon QuickSight users, groups, or both to assign the policy to. See `identities`.
	Identities *IamPolicyAssignmentIdentities `pulumi:"identities"`
	// Namespace that contains the assignment. Defaults to `default`.
	Namespace *string `pulumi:"namespace"`
	// ARN of the IAM policy to apply to the Amazon QuickSight users and groups specified in this assignment.
	PolicyArn *string `pulumi:"policyArn"`
}

// The set of arguments for constructing a IamPolicyAssignment resource.
type IamPolicyAssignmentArgs struct {
	// Name of the assignment.
	AssignmentName pulumi.StringInput
	// Status of the assignment. Valid values are `ENABLED`, `DISABLED`, and `DRAFT`.
	//
	// The following arguments are optional:
	AssignmentStatus pulumi.StringInput
	// AWS account ID.
	AwsAccountId pulumi.StringPtrInput
	// Amazon QuickSight users, groups, or both to assign the policy to. See `identities`.
	Identities IamPolicyAssignmentIdentitiesPtrInput
	// Namespace that contains the assignment. Defaults to `default`.
	Namespace pulumi.StringPtrInput
	// ARN of the IAM policy to apply to the Amazon QuickSight users and groups specified in this assignment.
	PolicyArn pulumi.StringPtrInput
}

func (IamPolicyAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamPolicyAssignmentArgs)(nil)).Elem()
}

type IamPolicyAssignmentInput interface {
	pulumi.Input

	ToIamPolicyAssignmentOutput() IamPolicyAssignmentOutput
	ToIamPolicyAssignmentOutputWithContext(ctx context.Context) IamPolicyAssignmentOutput
}

func (*IamPolicyAssignment) ElementType() reflect.Type {
	return reflect.TypeOf((**IamPolicyAssignment)(nil)).Elem()
}

func (i *IamPolicyAssignment) ToIamPolicyAssignmentOutput() IamPolicyAssignmentOutput {
	return i.ToIamPolicyAssignmentOutputWithContext(context.Background())
}

func (i *IamPolicyAssignment) ToIamPolicyAssignmentOutputWithContext(ctx context.Context) IamPolicyAssignmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamPolicyAssignmentOutput)
}

// IamPolicyAssignmentArrayInput is an input type that accepts IamPolicyAssignmentArray and IamPolicyAssignmentArrayOutput values.
// You can construct a concrete instance of `IamPolicyAssignmentArrayInput` via:
//
//	IamPolicyAssignmentArray{ IamPolicyAssignmentArgs{...} }
type IamPolicyAssignmentArrayInput interface {
	pulumi.Input

	ToIamPolicyAssignmentArrayOutput() IamPolicyAssignmentArrayOutput
	ToIamPolicyAssignmentArrayOutputWithContext(context.Context) IamPolicyAssignmentArrayOutput
}

type IamPolicyAssignmentArray []IamPolicyAssignmentInput

func (IamPolicyAssignmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamPolicyAssignment)(nil)).Elem()
}

func (i IamPolicyAssignmentArray) ToIamPolicyAssignmentArrayOutput() IamPolicyAssignmentArrayOutput {
	return i.ToIamPolicyAssignmentArrayOutputWithContext(context.Background())
}

func (i IamPolicyAssignmentArray) ToIamPolicyAssignmentArrayOutputWithContext(ctx context.Context) IamPolicyAssignmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamPolicyAssignmentArrayOutput)
}

// IamPolicyAssignmentMapInput is an input type that accepts IamPolicyAssignmentMap and IamPolicyAssignmentMapOutput values.
// You can construct a concrete instance of `IamPolicyAssignmentMapInput` via:
//
//	IamPolicyAssignmentMap{ "key": IamPolicyAssignmentArgs{...} }
type IamPolicyAssignmentMapInput interface {
	pulumi.Input

	ToIamPolicyAssignmentMapOutput() IamPolicyAssignmentMapOutput
	ToIamPolicyAssignmentMapOutputWithContext(context.Context) IamPolicyAssignmentMapOutput
}

type IamPolicyAssignmentMap map[string]IamPolicyAssignmentInput

func (IamPolicyAssignmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamPolicyAssignment)(nil)).Elem()
}

func (i IamPolicyAssignmentMap) ToIamPolicyAssignmentMapOutput() IamPolicyAssignmentMapOutput {
	return i.ToIamPolicyAssignmentMapOutputWithContext(context.Background())
}

func (i IamPolicyAssignmentMap) ToIamPolicyAssignmentMapOutputWithContext(ctx context.Context) IamPolicyAssignmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamPolicyAssignmentMapOutput)
}

type IamPolicyAssignmentOutput struct{ *pulumi.OutputState }

func (IamPolicyAssignmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamPolicyAssignment)(nil)).Elem()
}

func (o IamPolicyAssignmentOutput) ToIamPolicyAssignmentOutput() IamPolicyAssignmentOutput {
	return o
}

func (o IamPolicyAssignmentOutput) ToIamPolicyAssignmentOutputWithContext(ctx context.Context) IamPolicyAssignmentOutput {
	return o
}

// Assignment ID.
func (o IamPolicyAssignmentOutput) AssignmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *IamPolicyAssignment) pulumi.StringOutput { return v.AssignmentId }).(pulumi.StringOutput)
}

// Name of the assignment.
func (o IamPolicyAssignmentOutput) AssignmentName() pulumi.StringOutput {
	return o.ApplyT(func(v *IamPolicyAssignment) pulumi.StringOutput { return v.AssignmentName }).(pulumi.StringOutput)
}

// Status of the assignment. Valid values are `ENABLED`, `DISABLED`, and `DRAFT`.
//
// The following arguments are optional:
func (o IamPolicyAssignmentOutput) AssignmentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *IamPolicyAssignment) pulumi.StringOutput { return v.AssignmentStatus }).(pulumi.StringOutput)
}

// AWS account ID.
func (o IamPolicyAssignmentOutput) AwsAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *IamPolicyAssignment) pulumi.StringOutput { return v.AwsAccountId }).(pulumi.StringOutput)
}

// Amazon QuickSight users, groups, or both to assign the policy to. See `identities`.
func (o IamPolicyAssignmentOutput) Identities() IamPolicyAssignmentIdentitiesPtrOutput {
	return o.ApplyT(func(v *IamPolicyAssignment) IamPolicyAssignmentIdentitiesPtrOutput { return v.Identities }).(IamPolicyAssignmentIdentitiesPtrOutput)
}

// Namespace that contains the assignment. Defaults to `default`.
func (o IamPolicyAssignmentOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v *IamPolicyAssignment) pulumi.StringOutput { return v.Namespace }).(pulumi.StringOutput)
}

// ARN of the IAM policy to apply to the Amazon QuickSight users and groups specified in this assignment.
func (o IamPolicyAssignmentOutput) PolicyArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IamPolicyAssignment) pulumi.StringPtrOutput { return v.PolicyArn }).(pulumi.StringPtrOutput)
}

type IamPolicyAssignmentArrayOutput struct{ *pulumi.OutputState }

func (IamPolicyAssignmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamPolicyAssignment)(nil)).Elem()
}

func (o IamPolicyAssignmentArrayOutput) ToIamPolicyAssignmentArrayOutput() IamPolicyAssignmentArrayOutput {
	return o
}

func (o IamPolicyAssignmentArrayOutput) ToIamPolicyAssignmentArrayOutputWithContext(ctx context.Context) IamPolicyAssignmentArrayOutput {
	return o
}

func (o IamPolicyAssignmentArrayOutput) Index(i pulumi.IntInput) IamPolicyAssignmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IamPolicyAssignment {
		return vs[0].([]*IamPolicyAssignment)[vs[1].(int)]
	}).(IamPolicyAssignmentOutput)
}

type IamPolicyAssignmentMapOutput struct{ *pulumi.OutputState }

func (IamPolicyAssignmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamPolicyAssignment)(nil)).Elem()
}

func (o IamPolicyAssignmentMapOutput) ToIamPolicyAssignmentMapOutput() IamPolicyAssignmentMapOutput {
	return o
}

func (o IamPolicyAssignmentMapOutput) ToIamPolicyAssignmentMapOutputWithContext(ctx context.Context) IamPolicyAssignmentMapOutput {
	return o
}

func (o IamPolicyAssignmentMapOutput) MapIndex(k pulumi.StringInput) IamPolicyAssignmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IamPolicyAssignment {
		return vs[0].(map[string]*IamPolicyAssignment)[vs[1].(string)]
	}).(IamPolicyAssignmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamPolicyAssignmentInput)(nil)).Elem(), &IamPolicyAssignment{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamPolicyAssignmentArrayInput)(nil)).Elem(), IamPolicyAssignmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamPolicyAssignmentMapInput)(nil)).Elem(), IamPolicyAssignmentMap{})
	pulumi.RegisterOutputType(IamPolicyAssignmentOutput{})
	pulumi.RegisterOutputType(IamPolicyAssignmentArrayOutput{})
	pulumi.RegisterOutputType(IamPolicyAssignmentMapOutput{})
}
