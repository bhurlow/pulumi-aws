// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Canonical User ID data source allows access to the [canonical user ID](http://docs.aws.amazon.com/general/latest/gr/acct-identifiers.html)
// for the effective account in which this provider is working.
//
// > **NOTE:** To use this data source, you must have the `s3:ListAllMyBuckets` permission.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			current, err := s3.GetCanonicalUserId(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("canonicalUserId", current.Id)
//			return nil
//		})
//	}
//
// ```
func GetCanonicalUserId(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetCanonicalUserIdResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCanonicalUserIdResult
	err := ctx.Invoke("aws:s3/getCanonicalUserId:getCanonicalUserId", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getCanonicalUserId.
type GetCanonicalUserIdResult struct {
	// Human-friendly name linked to the canonical user ID. The bucket owner's display name. **NOTE:** [This value](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTServiceGET.html) is only included in the response in the US East (N. Virginia), US West (N. California), US West (Oregon), Asia Pacific (Singapore), Asia Pacific (Sydney), Asia Pacific (Tokyo), EU (Ireland), and South America (São Paulo) regions.
	DisplayName string `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
