// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the summary of a WAFv2 Web ACL.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/wafv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := wafv2.LookupWebAcl(ctx, &wafv2.LookupWebAclArgs{
//				Name:  "some-web-acl",
//				Scope: "REGIONAL",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupWebAcl(ctx *pulumi.Context, args *LookupWebAclArgs, opts ...pulumi.InvokeOption) (*LookupWebAclResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAclResult
	err := ctx.Invoke("aws:wafv2/getWebAcl:getWebAcl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWebAcl.
type LookupWebAclArgs struct {
	// Name of the WAFv2 Web ACL.
	Name string `pulumi:"name"`
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
	Scope string `pulumi:"scope"`
}

// A collection of values returned by getWebAcl.
type LookupWebAclResult struct {
	// ARN of the entity.
	Arn string `pulumi:"arn"`
	// Description of the WebACL that helps with identification.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id    string `pulumi:"id"`
	Name  string `pulumi:"name"`
	Scope string `pulumi:"scope"`
}

func LookupWebAclOutput(ctx *pulumi.Context, args LookupWebAclOutputArgs, opts ...pulumi.InvokeOption) LookupWebAclResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAclResult, error) {
			args := v.(LookupWebAclArgs)
			r, err := LookupWebAcl(ctx, &args, opts...)
			var s LookupWebAclResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAclResultOutput)
}

// A collection of arguments for invoking getWebAcl.
type LookupWebAclOutputArgs struct {
	// Name of the WAFv2 Web ACL.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
	Scope pulumi.StringInput `pulumi:"scope"`
}

func (LookupWebAclOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAclArgs)(nil)).Elem()
}

// A collection of values returned by getWebAcl.
type LookupWebAclResultOutput struct{ *pulumi.OutputState }

func (LookupWebAclResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAclResult)(nil)).Elem()
}

func (o LookupWebAclResultOutput) ToLookupWebAclResultOutput() LookupWebAclResultOutput {
	return o
}

func (o LookupWebAclResultOutput) ToLookupWebAclResultOutputWithContext(ctx context.Context) LookupWebAclResultOutput {
	return o
}

// ARN of the entity.
func (o LookupWebAclResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAclResult) string { return v.Arn }).(pulumi.StringOutput)
}

// Description of the WebACL that helps with identification.
func (o LookupWebAclResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAclResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupWebAclResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAclResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAclResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAclResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAclResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAclResult) string { return v.Scope }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAclResultOutput{})
}
