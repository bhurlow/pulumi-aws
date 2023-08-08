// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an IAM policy attached to a user.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			lbUser, err := iam.NewUser(ctx, "lbUser", &iam.UserArgs{
//				Path: pulumi.String("/system/"),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"Version": "2012-10-17",
//				"Statement": []map[string]interface{}{
//					map[string]interface{}{
//						"Action": []string{
//							"ec2:Describe*",
//						},
//						"Effect":   "Allow",
//						"Resource": "*",
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = iam.NewUserPolicy(ctx, "lbRo", &iam.UserPolicyArgs{
//				User:   lbUser.Name,
//				Policy: pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = iam.NewAccessKey(ctx, "lbAccessKey", &iam.AccessKeyArgs{
//				User: lbUser.Name,
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
//	to = aws_iam_user_policy.mypolicy
//
//	id = "user_of_mypolicy_name:mypolicy_name" } Using `pulumi import`, import IAM User Policies using the `user_name:user_policy_name`. For exampleconsole % pulumi import aws_iam_user_policy.mypolicy user_of_mypolicy_name:mypolicy_name
type UserPolicy struct {
	pulumi.CustomResourceState

	// The name of the policy. If omitted, the provider will assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// The policy document. This is a JSON formatted string.
	Policy pulumi.StringOutput `pulumi:"policy"`
	// IAM user to which to attach this policy.
	User pulumi.StringOutput `pulumi:"user"`
}

// NewUserPolicy registers a new resource with the given unique name, arguments, and options.
func NewUserPolicy(ctx *pulumi.Context,
	name string, args *UserPolicyArgs, opts ...pulumi.ResourceOption) (*UserPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	if args.User == nil {
		return nil, errors.New("invalid value for required argument 'User'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserPolicy
	err := ctx.RegisterResource("aws:iam/userPolicy:UserPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserPolicy gets an existing UserPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserPolicyState, opts ...pulumi.ResourceOption) (*UserPolicy, error) {
	var resource UserPolicy
	err := ctx.ReadResource("aws:iam/userPolicy:UserPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserPolicy resources.
type userPolicyState struct {
	// The name of the policy. If omitted, the provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The policy document. This is a JSON formatted string.
	Policy interface{} `pulumi:"policy"`
	// IAM user to which to attach this policy.
	User *string `pulumi:"user"`
}

type UserPolicyState struct {
	// The name of the policy. If omitted, the provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The policy document. This is a JSON formatted string.
	Policy pulumi.Input
	// IAM user to which to attach this policy.
	User pulumi.StringPtrInput
}

func (UserPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*userPolicyState)(nil)).Elem()
}

type userPolicyArgs struct {
	// The name of the policy. If omitted, the provider will assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The policy document. This is a JSON formatted string.
	Policy interface{} `pulumi:"policy"`
	// IAM user to which to attach this policy.
	User string `pulumi:"user"`
}

// The set of arguments for constructing a UserPolicy resource.
type UserPolicyArgs struct {
	// The name of the policy. If omitted, the provider will assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The policy document. This is a JSON formatted string.
	Policy pulumi.Input
	// IAM user to which to attach this policy.
	User pulumi.StringInput
}

func (UserPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userPolicyArgs)(nil)).Elem()
}

type UserPolicyInput interface {
	pulumi.Input

	ToUserPolicyOutput() UserPolicyOutput
	ToUserPolicyOutputWithContext(ctx context.Context) UserPolicyOutput
}

func (*UserPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPolicy)(nil)).Elem()
}

func (i *UserPolicy) ToUserPolicyOutput() UserPolicyOutput {
	return i.ToUserPolicyOutputWithContext(context.Background())
}

func (i *UserPolicy) ToUserPolicyOutputWithContext(ctx context.Context) UserPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPolicyOutput)
}

// UserPolicyArrayInput is an input type that accepts UserPolicyArray and UserPolicyArrayOutput values.
// You can construct a concrete instance of `UserPolicyArrayInput` via:
//
//	UserPolicyArray{ UserPolicyArgs{...} }
type UserPolicyArrayInput interface {
	pulumi.Input

	ToUserPolicyArrayOutput() UserPolicyArrayOutput
	ToUserPolicyArrayOutputWithContext(context.Context) UserPolicyArrayOutput
}

type UserPolicyArray []UserPolicyInput

func (UserPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserPolicy)(nil)).Elem()
}

func (i UserPolicyArray) ToUserPolicyArrayOutput() UserPolicyArrayOutput {
	return i.ToUserPolicyArrayOutputWithContext(context.Background())
}

func (i UserPolicyArray) ToUserPolicyArrayOutputWithContext(ctx context.Context) UserPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPolicyArrayOutput)
}

// UserPolicyMapInput is an input type that accepts UserPolicyMap and UserPolicyMapOutput values.
// You can construct a concrete instance of `UserPolicyMapInput` via:
//
//	UserPolicyMap{ "key": UserPolicyArgs{...} }
type UserPolicyMapInput interface {
	pulumi.Input

	ToUserPolicyMapOutput() UserPolicyMapOutput
	ToUserPolicyMapOutputWithContext(context.Context) UserPolicyMapOutput
}

type UserPolicyMap map[string]UserPolicyInput

func (UserPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserPolicy)(nil)).Elem()
}

func (i UserPolicyMap) ToUserPolicyMapOutput() UserPolicyMapOutput {
	return i.ToUserPolicyMapOutputWithContext(context.Background())
}

func (i UserPolicyMap) ToUserPolicyMapOutputWithContext(ctx context.Context) UserPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPolicyMapOutput)
}

type UserPolicyOutput struct{ *pulumi.OutputState }

func (UserPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPolicy)(nil)).Elem()
}

func (o UserPolicyOutput) ToUserPolicyOutput() UserPolicyOutput {
	return o
}

func (o UserPolicyOutput) ToUserPolicyOutputWithContext(ctx context.Context) UserPolicyOutput {
	return o
}

// The name of the policy. If omitted, the provider will assign a random, unique name.
func (o UserPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
func (o UserPolicyOutput) NamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserPolicy) pulumi.StringPtrOutput { return v.NamePrefix }).(pulumi.StringPtrOutput)
}

// The policy document. This is a JSON formatted string.
func (o UserPolicyOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPolicy) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

// IAM user to which to attach this policy.
func (o UserPolicyOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPolicy) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

type UserPolicyArrayOutput struct{ *pulumi.OutputState }

func (UserPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserPolicy)(nil)).Elem()
}

func (o UserPolicyArrayOutput) ToUserPolicyArrayOutput() UserPolicyArrayOutput {
	return o
}

func (o UserPolicyArrayOutput) ToUserPolicyArrayOutputWithContext(ctx context.Context) UserPolicyArrayOutput {
	return o
}

func (o UserPolicyArrayOutput) Index(i pulumi.IntInput) UserPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserPolicy {
		return vs[0].([]*UserPolicy)[vs[1].(int)]
	}).(UserPolicyOutput)
}

type UserPolicyMapOutput struct{ *pulumi.OutputState }

func (UserPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserPolicy)(nil)).Elem()
}

func (o UserPolicyMapOutput) ToUserPolicyMapOutput() UserPolicyMapOutput {
	return o
}

func (o UserPolicyMapOutput) ToUserPolicyMapOutputWithContext(ctx context.Context) UserPolicyMapOutput {
	return o
}

func (o UserPolicyMapOutput) MapIndex(k pulumi.StringInput) UserPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserPolicy {
		return vs[0].(map[string]*UserPolicy)[vs[1].(string)]
	}).(UserPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserPolicyInput)(nil)).Elem(), &UserPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPolicyArrayInput)(nil)).Elem(), UserPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPolicyMapInput)(nil)).Elem(), UserPolicyMap{})
	pulumi.RegisterOutputType(UserPolicyOutput{})
	pulumi.RegisterOutputType(UserPolicyArrayOutput{})
	pulumi.RegisterOutputType(UserPolicyMapOutput{})
}
