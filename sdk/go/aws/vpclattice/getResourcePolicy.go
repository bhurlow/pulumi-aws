// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpclattice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Data source for managing an AWS VPC Lattice Resource Policy.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/vpclattice"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpclattice.LookupResourcePolicy(ctx, &vpclattice.LookupResourcePolicyArgs{
//				ResourceArn: aws_vpclattice_service_network.Example.Arn,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupResourcePolicy(ctx *pulumi.Context, args *LookupResourcePolicyArgs, opts ...pulumi.InvokeOption) (*LookupResourcePolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupResourcePolicyResult
	err := ctx.Invoke("aws:vpclattice/getResourcePolicy:getResourcePolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResourcePolicy.
type LookupResourcePolicyArgs struct {
	// Resource ARN of the resource for which a policy is retrieved.
	ResourceArn string `pulumi:"resourceArn"`
}

// A collection of values returned by getResourcePolicy.
type LookupResourcePolicyResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// JSON-encoded string representation of the applied resource policy.
	Policy      string `pulumi:"policy"`
	ResourceArn string `pulumi:"resourceArn"`
}

func LookupResourcePolicyOutput(ctx *pulumi.Context, args LookupResourcePolicyOutputArgs, opts ...pulumi.InvokeOption) LookupResourcePolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourcePolicyResult, error) {
			args := v.(LookupResourcePolicyArgs)
			r, err := LookupResourcePolicy(ctx, &args, opts...)
			var s LookupResourcePolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourcePolicyResultOutput)
}

// A collection of arguments for invoking getResourcePolicy.
type LookupResourcePolicyOutputArgs struct {
	// Resource ARN of the resource for which a policy is retrieved.
	ResourceArn pulumi.StringInput `pulumi:"resourceArn"`
}

func (LookupResourcePolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourcePolicyArgs)(nil)).Elem()
}

// A collection of values returned by getResourcePolicy.
type LookupResourcePolicyResultOutput struct{ *pulumi.OutputState }

func (LookupResourcePolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourcePolicyResult)(nil)).Elem()
}

func (o LookupResourcePolicyResultOutput) ToLookupResourcePolicyResultOutput() LookupResourcePolicyResultOutput {
	return o
}

func (o LookupResourcePolicyResultOutput) ToLookupResourcePolicyResultOutputWithContext(ctx context.Context) LookupResourcePolicyResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupResourcePolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// JSON-encoded string representation of the applied resource policy.
func (o LookupResourcePolicyResultOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePolicyResult) string { return v.Policy }).(pulumi.StringOutput)
}

func (o LookupResourcePolicyResultOutput) ResourceArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePolicyResult) string { return v.ResourceArn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourcePolicyResultOutput{})
}
