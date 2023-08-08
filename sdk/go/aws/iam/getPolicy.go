// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can be used to fetch information about a specific
// IAM policy.
//
// ## Example Usage
// ### By ARN
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iam.LookupPolicy(ctx, &iam.LookupPolicyArgs{
//				Arn: pulumi.StringRef("arn:aws:iam::123456789012:policy/UsersManageOwnCredentials"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### By Name
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iam.LookupPolicy(ctx, &iam.LookupPolicyArgs{
//				Name: pulumi.StringRef("test_policy"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupPolicy(ctx *pulumi.Context, args *LookupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPolicyResult
	err := ctx.Invoke("aws:iam/getPolicy:getPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicy.
type LookupPolicyArgs struct {
	// ARN of the IAM policy.
	// Conflicts with `name` and `pathPrefix`.
	Arn *string `pulumi:"arn"`
	// Name of the IAM policy.
	// Conflicts with `arn`.
	Name *string `pulumi:"name"`
	// Prefix of the path to the IAM policy.
	// Defaults to a slash (`/`).
	// Conflicts with `arn`.
	PathPrefix *string `pulumi:"pathPrefix"`
	// Key-value mapping of tags for the IAM Policy.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getPolicy.
type LookupPolicyResult struct {
	// ARN of the policy.
	Arn string `pulumi:"arn"`
	// Description of the policy.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Path to the policy.
	Path       string  `pulumi:"path"`
	PathPrefix *string `pulumi:"pathPrefix"`
	// Policy document of the policy.
	Policy string `pulumi:"policy"`
	// Policy's ID.
	PolicyId string `pulumi:"policyId"`
	// Key-value mapping of tags for the IAM Policy.
	Tags map[string]string `pulumi:"tags"`
}

func LookupPolicyOutput(ctx *pulumi.Context, args LookupPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPolicyResult, error) {
			args := v.(LookupPolicyArgs)
			r, err := LookupPolicy(ctx, &args, opts...)
			var s LookupPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPolicyResultOutput)
}

// A collection of arguments for invoking getPolicy.
type LookupPolicyOutputArgs struct {
	// ARN of the IAM policy.
	// Conflicts with `name` and `pathPrefix`.
	Arn pulumi.StringPtrInput `pulumi:"arn"`
	// Name of the IAM policy.
	// Conflicts with `arn`.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Prefix of the path to the IAM policy.
	// Defaults to a slash (`/`).
	// Conflicts with `arn`.
	PathPrefix pulumi.StringPtrInput `pulumi:"pathPrefix"`
	// Key-value mapping of tags for the IAM Policy.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getPolicy.
type LookupPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicyResult)(nil)).Elem()
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutput() LookupPolicyResultOutput {
	return o
}

func (o LookupPolicyResultOutput) ToLookupPolicyResultOutputWithContext(ctx context.Context) LookupPolicyResultOutput {
	return o
}

// ARN of the policy.
func (o LookupPolicyResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Arn }).(pulumi.StringOutput)
}

// Description of the policy.
func (o LookupPolicyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Path to the policy.
func (o LookupPolicyResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Path }).(pulumi.StringOutput)
}

func (o LookupPolicyResultOutput) PathPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPolicyResult) *string { return v.PathPrefix }).(pulumi.StringPtrOutput)
}

// Policy document of the policy.
func (o LookupPolicyResultOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.Policy }).(pulumi.StringOutput)
}

// Policy's ID.
func (o LookupPolicyResultOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicyResult) string { return v.PolicyId }).(pulumi.StringOutput)
}

// Key-value mapping of tags for the IAM Policy.
func (o LookupPolicyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPolicyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicyResultOutput{})
}
