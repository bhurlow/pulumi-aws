// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Lists the paths to the specified product. A path is how the user has access to a specified product, and is necessary when provisioning a product. A path also determines the constraints put on the product.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicecatalog"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := servicecatalog.GetLaunchPaths(ctx, &servicecatalog.GetLaunchPathsArgs{
//				ProductId: "prod-yakog5pdriver",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetLaunchPaths(ctx *pulumi.Context, args *GetLaunchPathsArgs, opts ...pulumi.InvokeOption) (*GetLaunchPathsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLaunchPathsResult
	err := ctx.Invoke("aws:servicecatalog/getLaunchPaths:getLaunchPaths", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLaunchPaths.
type GetLaunchPathsArgs struct {
	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
	AcceptLanguage *string `pulumi:"acceptLanguage"`
	// Product identifier.
	//
	// The following arguments are optional:
	ProductId string `pulumi:"productId"`
}

// A collection of values returned by getLaunchPaths.
type GetLaunchPathsResult struct {
	AcceptLanguage *string `pulumi:"acceptLanguage"`
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	ProductId string `pulumi:"productId"`
	// Block with information about the launch path. See details below.
	Summaries []GetLaunchPathsSummary `pulumi:"summaries"`
}

func GetLaunchPathsOutput(ctx *pulumi.Context, args GetLaunchPathsOutputArgs, opts ...pulumi.InvokeOption) GetLaunchPathsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLaunchPathsResult, error) {
			args := v.(GetLaunchPathsArgs)
			r, err := GetLaunchPaths(ctx, &args, opts...)
			var s GetLaunchPathsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLaunchPathsResultOutput)
}

// A collection of arguments for invoking getLaunchPaths.
type GetLaunchPathsOutputArgs struct {
	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
	AcceptLanguage pulumi.StringPtrInput `pulumi:"acceptLanguage"`
	// Product identifier.
	//
	// The following arguments are optional:
	ProductId pulumi.StringInput `pulumi:"productId"`
}

func (GetLaunchPathsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLaunchPathsArgs)(nil)).Elem()
}

// A collection of values returned by getLaunchPaths.
type GetLaunchPathsResultOutput struct{ *pulumi.OutputState }

func (GetLaunchPathsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLaunchPathsResult)(nil)).Elem()
}

func (o GetLaunchPathsResultOutput) ToGetLaunchPathsResultOutput() GetLaunchPathsResultOutput {
	return o
}

func (o GetLaunchPathsResultOutput) ToGetLaunchPathsResultOutputWithContext(ctx context.Context) GetLaunchPathsResultOutput {
	return o
}

func (o GetLaunchPathsResultOutput) AcceptLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLaunchPathsResult) *string { return v.AcceptLanguage }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLaunchPathsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLaunchPathsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLaunchPathsResultOutput) ProductId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLaunchPathsResult) string { return v.ProductId }).(pulumi.StringOutput)
}

// Block with information about the launch path. See details below.
func (o GetLaunchPathsResultOutput) Summaries() GetLaunchPathsSummaryArrayOutput {
	return o.ApplyT(func(v GetLaunchPathsResult) []GetLaunchPathsSummary { return v.Summaries }).(GetLaunchPathsSummaryArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLaunchPathsResultOutput{})
}
