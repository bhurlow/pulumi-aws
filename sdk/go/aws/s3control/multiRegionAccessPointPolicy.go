// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3control

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage an S3 Multi-Region Access Point access control policy.
//
// ## Example Usage
//
// ## Import
//
// Multi-Region Access Point Policies can be imported using the `account_id` and `name` of the Multi-Region Access Point separated by a colon (`:`), e.g.
//
// ```sh
//
//	$ pulumi import aws:s3control/multiRegionAccessPointPolicy:MultiRegionAccessPointPolicy example 123456789012:example
//
// ```
type MultiRegionAccessPointPolicy struct {
	pulumi.CustomResourceState

	// The AWS account ID for the owner of the Multi-Region Access Point. Defaults to automatically determined account ID of the AWS provider.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// A configuration block containing details about the policy for the Multi-Region Access Point. See Details Configuration Block below for more details
	Details MultiRegionAccessPointPolicyDetailsOutput `pulumi:"details"`
	// The last established policy for the Multi-Region Access Point.
	Established pulumi.StringOutput `pulumi:"established"`
	// The proposed policy for the Multi-Region Access Point.
	Proposed pulumi.StringOutput `pulumi:"proposed"`
}

// NewMultiRegionAccessPointPolicy registers a new resource with the given unique name, arguments, and options.
func NewMultiRegionAccessPointPolicy(ctx *pulumi.Context,
	name string, args *MultiRegionAccessPointPolicyArgs, opts ...pulumi.ResourceOption) (*MultiRegionAccessPointPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Details == nil {
		return nil, errors.New("invalid value for required argument 'Details'")
	}
	var resource MultiRegionAccessPointPolicy
	err := ctx.RegisterResource("aws:s3control/multiRegionAccessPointPolicy:MultiRegionAccessPointPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMultiRegionAccessPointPolicy gets an existing MultiRegionAccessPointPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMultiRegionAccessPointPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MultiRegionAccessPointPolicyState, opts ...pulumi.ResourceOption) (*MultiRegionAccessPointPolicy, error) {
	var resource MultiRegionAccessPointPolicy
	err := ctx.ReadResource("aws:s3control/multiRegionAccessPointPolicy:MultiRegionAccessPointPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MultiRegionAccessPointPolicy resources.
type multiRegionAccessPointPolicyState struct {
	// The AWS account ID for the owner of the Multi-Region Access Point. Defaults to automatically determined account ID of the AWS provider.
	AccountId *string `pulumi:"accountId"`
	// A configuration block containing details about the policy for the Multi-Region Access Point. See Details Configuration Block below for more details
	Details *MultiRegionAccessPointPolicyDetails `pulumi:"details"`
	// The last established policy for the Multi-Region Access Point.
	Established *string `pulumi:"established"`
	// The proposed policy for the Multi-Region Access Point.
	Proposed *string `pulumi:"proposed"`
}

type MultiRegionAccessPointPolicyState struct {
	// The AWS account ID for the owner of the Multi-Region Access Point. Defaults to automatically determined account ID of the AWS provider.
	AccountId pulumi.StringPtrInput
	// A configuration block containing details about the policy for the Multi-Region Access Point. See Details Configuration Block below for more details
	Details MultiRegionAccessPointPolicyDetailsPtrInput
	// The last established policy for the Multi-Region Access Point.
	Established pulumi.StringPtrInput
	// The proposed policy for the Multi-Region Access Point.
	Proposed pulumi.StringPtrInput
}

func (MultiRegionAccessPointPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*multiRegionAccessPointPolicyState)(nil)).Elem()
}

type multiRegionAccessPointPolicyArgs struct {
	// The AWS account ID for the owner of the Multi-Region Access Point. Defaults to automatically determined account ID of the AWS provider.
	AccountId *string `pulumi:"accountId"`
	// A configuration block containing details about the policy for the Multi-Region Access Point. See Details Configuration Block below for more details
	Details MultiRegionAccessPointPolicyDetails `pulumi:"details"`
}

// The set of arguments for constructing a MultiRegionAccessPointPolicy resource.
type MultiRegionAccessPointPolicyArgs struct {
	// The AWS account ID for the owner of the Multi-Region Access Point. Defaults to automatically determined account ID of the AWS provider.
	AccountId pulumi.StringPtrInput
	// A configuration block containing details about the policy for the Multi-Region Access Point. See Details Configuration Block below for more details
	Details MultiRegionAccessPointPolicyDetailsInput
}

func (MultiRegionAccessPointPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*multiRegionAccessPointPolicyArgs)(nil)).Elem()
}

type MultiRegionAccessPointPolicyInput interface {
	pulumi.Input

	ToMultiRegionAccessPointPolicyOutput() MultiRegionAccessPointPolicyOutput
	ToMultiRegionAccessPointPolicyOutputWithContext(ctx context.Context) MultiRegionAccessPointPolicyOutput
}

func (*MultiRegionAccessPointPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**MultiRegionAccessPointPolicy)(nil)).Elem()
}

func (i *MultiRegionAccessPointPolicy) ToMultiRegionAccessPointPolicyOutput() MultiRegionAccessPointPolicyOutput {
	return i.ToMultiRegionAccessPointPolicyOutputWithContext(context.Background())
}

func (i *MultiRegionAccessPointPolicy) ToMultiRegionAccessPointPolicyOutputWithContext(ctx context.Context) MultiRegionAccessPointPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultiRegionAccessPointPolicyOutput)
}

// MultiRegionAccessPointPolicyArrayInput is an input type that accepts MultiRegionAccessPointPolicyArray and MultiRegionAccessPointPolicyArrayOutput values.
// You can construct a concrete instance of `MultiRegionAccessPointPolicyArrayInput` via:
//
//	MultiRegionAccessPointPolicyArray{ MultiRegionAccessPointPolicyArgs{...} }
type MultiRegionAccessPointPolicyArrayInput interface {
	pulumi.Input

	ToMultiRegionAccessPointPolicyArrayOutput() MultiRegionAccessPointPolicyArrayOutput
	ToMultiRegionAccessPointPolicyArrayOutputWithContext(context.Context) MultiRegionAccessPointPolicyArrayOutput
}

type MultiRegionAccessPointPolicyArray []MultiRegionAccessPointPolicyInput

func (MultiRegionAccessPointPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MultiRegionAccessPointPolicy)(nil)).Elem()
}

func (i MultiRegionAccessPointPolicyArray) ToMultiRegionAccessPointPolicyArrayOutput() MultiRegionAccessPointPolicyArrayOutput {
	return i.ToMultiRegionAccessPointPolicyArrayOutputWithContext(context.Background())
}

func (i MultiRegionAccessPointPolicyArray) ToMultiRegionAccessPointPolicyArrayOutputWithContext(ctx context.Context) MultiRegionAccessPointPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultiRegionAccessPointPolicyArrayOutput)
}

// MultiRegionAccessPointPolicyMapInput is an input type that accepts MultiRegionAccessPointPolicyMap and MultiRegionAccessPointPolicyMapOutput values.
// You can construct a concrete instance of `MultiRegionAccessPointPolicyMapInput` via:
//
//	MultiRegionAccessPointPolicyMap{ "key": MultiRegionAccessPointPolicyArgs{...} }
type MultiRegionAccessPointPolicyMapInput interface {
	pulumi.Input

	ToMultiRegionAccessPointPolicyMapOutput() MultiRegionAccessPointPolicyMapOutput
	ToMultiRegionAccessPointPolicyMapOutputWithContext(context.Context) MultiRegionAccessPointPolicyMapOutput
}

type MultiRegionAccessPointPolicyMap map[string]MultiRegionAccessPointPolicyInput

func (MultiRegionAccessPointPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MultiRegionAccessPointPolicy)(nil)).Elem()
}

func (i MultiRegionAccessPointPolicyMap) ToMultiRegionAccessPointPolicyMapOutput() MultiRegionAccessPointPolicyMapOutput {
	return i.ToMultiRegionAccessPointPolicyMapOutputWithContext(context.Background())
}

func (i MultiRegionAccessPointPolicyMap) ToMultiRegionAccessPointPolicyMapOutputWithContext(ctx context.Context) MultiRegionAccessPointPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultiRegionAccessPointPolicyMapOutput)
}

type MultiRegionAccessPointPolicyOutput struct{ *pulumi.OutputState }

func (MultiRegionAccessPointPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MultiRegionAccessPointPolicy)(nil)).Elem()
}

func (o MultiRegionAccessPointPolicyOutput) ToMultiRegionAccessPointPolicyOutput() MultiRegionAccessPointPolicyOutput {
	return o
}

func (o MultiRegionAccessPointPolicyOutput) ToMultiRegionAccessPointPolicyOutputWithContext(ctx context.Context) MultiRegionAccessPointPolicyOutput {
	return o
}

// The AWS account ID for the owner of the Multi-Region Access Point. Defaults to automatically determined account ID of the AWS provider.
func (o MultiRegionAccessPointPolicyOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *MultiRegionAccessPointPolicy) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// A configuration block containing details about the policy for the Multi-Region Access Point. See Details Configuration Block below for more details
func (o MultiRegionAccessPointPolicyOutput) Details() MultiRegionAccessPointPolicyDetailsOutput {
	return o.ApplyT(func(v *MultiRegionAccessPointPolicy) MultiRegionAccessPointPolicyDetailsOutput { return v.Details }).(MultiRegionAccessPointPolicyDetailsOutput)
}

// The last established policy for the Multi-Region Access Point.
func (o MultiRegionAccessPointPolicyOutput) Established() pulumi.StringOutput {
	return o.ApplyT(func(v *MultiRegionAccessPointPolicy) pulumi.StringOutput { return v.Established }).(pulumi.StringOutput)
}

// The proposed policy for the Multi-Region Access Point.
func (o MultiRegionAccessPointPolicyOutput) Proposed() pulumi.StringOutput {
	return o.ApplyT(func(v *MultiRegionAccessPointPolicy) pulumi.StringOutput { return v.Proposed }).(pulumi.StringOutput)
}

type MultiRegionAccessPointPolicyArrayOutput struct{ *pulumi.OutputState }

func (MultiRegionAccessPointPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MultiRegionAccessPointPolicy)(nil)).Elem()
}

func (o MultiRegionAccessPointPolicyArrayOutput) ToMultiRegionAccessPointPolicyArrayOutput() MultiRegionAccessPointPolicyArrayOutput {
	return o
}

func (o MultiRegionAccessPointPolicyArrayOutput) ToMultiRegionAccessPointPolicyArrayOutputWithContext(ctx context.Context) MultiRegionAccessPointPolicyArrayOutput {
	return o
}

func (o MultiRegionAccessPointPolicyArrayOutput) Index(i pulumi.IntInput) MultiRegionAccessPointPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MultiRegionAccessPointPolicy {
		return vs[0].([]*MultiRegionAccessPointPolicy)[vs[1].(int)]
	}).(MultiRegionAccessPointPolicyOutput)
}

type MultiRegionAccessPointPolicyMapOutput struct{ *pulumi.OutputState }

func (MultiRegionAccessPointPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MultiRegionAccessPointPolicy)(nil)).Elem()
}

func (o MultiRegionAccessPointPolicyMapOutput) ToMultiRegionAccessPointPolicyMapOutput() MultiRegionAccessPointPolicyMapOutput {
	return o
}

func (o MultiRegionAccessPointPolicyMapOutput) ToMultiRegionAccessPointPolicyMapOutputWithContext(ctx context.Context) MultiRegionAccessPointPolicyMapOutput {
	return o
}

func (o MultiRegionAccessPointPolicyMapOutput) MapIndex(k pulumi.StringInput) MultiRegionAccessPointPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MultiRegionAccessPointPolicy {
		return vs[0].(map[string]*MultiRegionAccessPointPolicy)[vs[1].(string)]
	}).(MultiRegionAccessPointPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MultiRegionAccessPointPolicyInput)(nil)).Elem(), &MultiRegionAccessPointPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*MultiRegionAccessPointPolicyArrayInput)(nil)).Elem(), MultiRegionAccessPointPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MultiRegionAccessPointPolicyMapInput)(nil)).Elem(), MultiRegionAccessPointPolicyMap{})
	pulumi.RegisterOutputType(MultiRegionAccessPointPolicyOutput{})
	pulumi.RegisterOutputType(MultiRegionAccessPointPolicyArrayOutput{})
	pulumi.RegisterOutputType(MultiRegionAccessPointPolicyMapOutput{})
}
