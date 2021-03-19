// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an IAM role inline policy.
//
// > **NOTE:** For a given role, this resource is incompatible with using the `iam.Role` resource `inlinePolicy` argument. When using that argument and this resource, both will attempt to manage the role's inline policies and the provider will show a permanent difference.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"encoding/json"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		tmpJSON0, err := json.Marshal(map[string]interface{}{
// 			"Version": "2012-10-17",
// 			"Statement": []map[string]interface{}{
// 				map[string]interface{}{
// 					"Action": "sts:AssumeRole",
// 					"Effect": "Allow",
// 					"Sid":    "",
// 					"Principal": map[string]interface{}{
// 						"Service": "ec2.amazonaws.com",
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		json0 := string(tmpJSON0)
// 		testRole, err := iam.NewRole(ctx, "testRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(json0),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		tmpJSON1, err := json.Marshal(map[string]interface{}{
// 			"Version": "2012-10-17",
// 			"Statement": []map[string]interface{}{
// 				map[string]interface{}{
// 					"Action": []string{
// 						"ec2:Describe*",
// 					},
// 					"Effect":   "Allow",
// 					"Resource": "*",
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		json1 := string(tmpJSON1)
// 		_, err = iam.NewRolePolicy(ctx, "testPolicy", &iam.RolePolicyArgs{
// 			Role:   testRole.ID(),
// 			Policy: pulumi.String(json1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// IAM Role Policies can be imported using the `role_name:role_policy_name`, e.g.
//
// ```sh
//  $ pulumi import aws:iam/rolePolicy:RolePolicy mypolicy role_of_mypolicy_name:mypolicy_name
// ```
type RolePolicy struct {
	pulumi.CustomResourceState

	// The name of the role policy. If omitted, this provider will
	// assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// The policy document. This is a JSON formatted string.
	Policy pulumi.StringOutput `pulumi:"policy"`
	// The IAM role to attach to the policy.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewRolePolicy registers a new resource with the given unique name, arguments, and options.
func NewRolePolicy(ctx *pulumi.Context,
	name string, args *RolePolicyArgs, opts ...pulumi.ResourceOption) (*RolePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource RolePolicy
	err := ctx.RegisterResource("aws:iam/rolePolicy:RolePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRolePolicy gets an existing RolePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRolePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RolePolicyState, opts ...pulumi.ResourceOption) (*RolePolicy, error) {
	var resource RolePolicy
	err := ctx.ReadResource("aws:iam/rolePolicy:RolePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RolePolicy resources.
type rolePolicyState struct {
	// The name of the role policy. If omitted, this provider will
	// assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The policy document. This is a JSON formatted string.
	Policy *string `pulumi:"policy"`
	// The IAM role to attach to the policy.
	Role *string `pulumi:"role"`
}

type RolePolicyState struct {
	// The name of the role policy. If omitted, this provider will
	// assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The policy document. This is a JSON formatted string.
	Policy pulumi.StringPtrInput
	// The IAM role to attach to the policy.
	Role pulumi.StringPtrInput
}

func (RolePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*rolePolicyState)(nil)).Elem()
}

type rolePolicyArgs struct {
	// The name of the role policy. If omitted, this provider will
	// assign a random, unique name.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The policy document. This is a JSON formatted string.
	Policy interface{} `pulumi:"policy"`
	// The IAM role to attach to the policy.
	Role interface{} `pulumi:"role"`
}

// The set of arguments for constructing a RolePolicy resource.
type RolePolicyArgs struct {
	// The name of the role policy. If omitted, this provider will
	// assign a random, unique name.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The policy document. This is a JSON formatted string.
	Policy pulumi.Input
	// The IAM role to attach to the policy.
	Role pulumi.Input
}

func (RolePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rolePolicyArgs)(nil)).Elem()
}

type RolePolicyInput interface {
	pulumi.Input

	ToRolePolicyOutput() RolePolicyOutput
	ToRolePolicyOutputWithContext(ctx context.Context) RolePolicyOutput
}

func (*RolePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*RolePolicy)(nil))
}

func (i *RolePolicy) ToRolePolicyOutput() RolePolicyOutput {
	return i.ToRolePolicyOutputWithContext(context.Background())
}

func (i *RolePolicy) ToRolePolicyOutputWithContext(ctx context.Context) RolePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RolePolicyOutput)
}

func (i *RolePolicy) ToRolePolicyPtrOutput() RolePolicyPtrOutput {
	return i.ToRolePolicyPtrOutputWithContext(context.Background())
}

func (i *RolePolicy) ToRolePolicyPtrOutputWithContext(ctx context.Context) RolePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RolePolicyPtrOutput)
}

type RolePolicyPtrInput interface {
	pulumi.Input

	ToRolePolicyPtrOutput() RolePolicyPtrOutput
	ToRolePolicyPtrOutputWithContext(ctx context.Context) RolePolicyPtrOutput
}

type rolePolicyPtrType RolePolicyArgs

func (*rolePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RolePolicy)(nil))
}

func (i *rolePolicyPtrType) ToRolePolicyPtrOutput() RolePolicyPtrOutput {
	return i.ToRolePolicyPtrOutputWithContext(context.Background())
}

func (i *rolePolicyPtrType) ToRolePolicyPtrOutputWithContext(ctx context.Context) RolePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RolePolicyPtrOutput)
}

// RolePolicyArrayInput is an input type that accepts RolePolicyArray and RolePolicyArrayOutput values.
// You can construct a concrete instance of `RolePolicyArrayInput` via:
//
//          RolePolicyArray{ RolePolicyArgs{...} }
type RolePolicyArrayInput interface {
	pulumi.Input

	ToRolePolicyArrayOutput() RolePolicyArrayOutput
	ToRolePolicyArrayOutputWithContext(context.Context) RolePolicyArrayOutput
}

type RolePolicyArray []RolePolicyInput

func (RolePolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*RolePolicy)(nil))
}

func (i RolePolicyArray) ToRolePolicyArrayOutput() RolePolicyArrayOutput {
	return i.ToRolePolicyArrayOutputWithContext(context.Background())
}

func (i RolePolicyArray) ToRolePolicyArrayOutputWithContext(ctx context.Context) RolePolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RolePolicyArrayOutput)
}

// RolePolicyMapInput is an input type that accepts RolePolicyMap and RolePolicyMapOutput values.
// You can construct a concrete instance of `RolePolicyMapInput` via:
//
//          RolePolicyMap{ "key": RolePolicyArgs{...} }
type RolePolicyMapInput interface {
	pulumi.Input

	ToRolePolicyMapOutput() RolePolicyMapOutput
	ToRolePolicyMapOutputWithContext(context.Context) RolePolicyMapOutput
}

type RolePolicyMap map[string]RolePolicyInput

func (RolePolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*RolePolicy)(nil))
}

func (i RolePolicyMap) ToRolePolicyMapOutput() RolePolicyMapOutput {
	return i.ToRolePolicyMapOutputWithContext(context.Background())
}

func (i RolePolicyMap) ToRolePolicyMapOutputWithContext(ctx context.Context) RolePolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RolePolicyMapOutput)
}

type RolePolicyOutput struct {
	*pulumi.OutputState
}

func (RolePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RolePolicy)(nil))
}

func (o RolePolicyOutput) ToRolePolicyOutput() RolePolicyOutput {
	return o
}

func (o RolePolicyOutput) ToRolePolicyOutputWithContext(ctx context.Context) RolePolicyOutput {
	return o
}

func (o RolePolicyOutput) ToRolePolicyPtrOutput() RolePolicyPtrOutput {
	return o.ToRolePolicyPtrOutputWithContext(context.Background())
}

func (o RolePolicyOutput) ToRolePolicyPtrOutputWithContext(ctx context.Context) RolePolicyPtrOutput {
	return o.ApplyT(func(v RolePolicy) *RolePolicy {
		return &v
	}).(RolePolicyPtrOutput)
}

type RolePolicyPtrOutput struct {
	*pulumi.OutputState
}

func (RolePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RolePolicy)(nil))
}

func (o RolePolicyPtrOutput) ToRolePolicyPtrOutput() RolePolicyPtrOutput {
	return o
}

func (o RolePolicyPtrOutput) ToRolePolicyPtrOutputWithContext(ctx context.Context) RolePolicyPtrOutput {
	return o
}

type RolePolicyArrayOutput struct{ *pulumi.OutputState }

func (RolePolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RolePolicy)(nil))
}

func (o RolePolicyArrayOutput) ToRolePolicyArrayOutput() RolePolicyArrayOutput {
	return o
}

func (o RolePolicyArrayOutput) ToRolePolicyArrayOutputWithContext(ctx context.Context) RolePolicyArrayOutput {
	return o
}

func (o RolePolicyArrayOutput) Index(i pulumi.IntInput) RolePolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RolePolicy {
		return vs[0].([]RolePolicy)[vs[1].(int)]
	}).(RolePolicyOutput)
}

type RolePolicyMapOutput struct{ *pulumi.OutputState }

func (RolePolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RolePolicy)(nil))
}

func (o RolePolicyMapOutput) ToRolePolicyMapOutput() RolePolicyMapOutput {
	return o
}

func (o RolePolicyMapOutput) ToRolePolicyMapOutputWithContext(ctx context.Context) RolePolicyMapOutput {
	return o
}

func (o RolePolicyMapOutput) MapIndex(k pulumi.StringInput) RolePolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RolePolicy {
		return vs[0].(map[string]RolePolicy)[vs[1].(string)]
	}).(RolePolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(RolePolicyOutput{})
	pulumi.RegisterOutputType(RolePolicyPtrOutput{})
	pulumi.RegisterOutputType(RolePolicyArrayOutput{})
	pulumi.RegisterOutputType(RolePolicyMapOutput{})
}
