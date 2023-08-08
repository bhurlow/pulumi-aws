// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package outposts

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about an Outposts Site.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/outposts"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := outposts.GetSite(ctx, &outposts.GetSiteArgs{
//				Name: pulumi.StringRef("example"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetSite(ctx *pulumi.Context, args *GetSiteArgs, opts ...pulumi.InvokeOption) (*GetSiteResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSiteResult
	err := ctx.Invoke("aws:outposts/getSite:getSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSite.
type GetSiteArgs struct {
	// Identifier of the Site.
	Id *string `pulumi:"id"`
	// Name of the Site.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getSite.
type GetSiteResult struct {
	// AWS Account identifier.
	AccountId string `pulumi:"accountId"`
	// Description.
	Description string `pulumi:"description"`
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
}

func GetSiteOutput(ctx *pulumi.Context, args GetSiteOutputArgs, opts ...pulumi.InvokeOption) GetSiteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSiteResult, error) {
			args := v.(GetSiteArgs)
			r, err := GetSite(ctx, &args, opts...)
			var s GetSiteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSiteResultOutput)
}

// A collection of arguments for invoking getSite.
type GetSiteOutputArgs struct {
	// Identifier of the Site.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Name of the Site.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetSiteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSiteArgs)(nil)).Elem()
}

// A collection of values returned by getSite.
type GetSiteResultOutput struct{ *pulumi.OutputState }

func (GetSiteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSiteResult)(nil)).Elem()
}

func (o GetSiteResultOutput) ToGetSiteResultOutput() GetSiteResultOutput {
	return o
}

func (o GetSiteResultOutput) ToGetSiteResultOutputWithContext(ctx context.Context) GetSiteResultOutput {
	return o
}

// AWS Account identifier.
func (o GetSiteResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSiteResult) string { return v.AccountId }).(pulumi.StringOutput)
}

// Description.
func (o GetSiteResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetSiteResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetSiteResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSiteResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSiteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetSiteResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSiteResultOutput{})
}
