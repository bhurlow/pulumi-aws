// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicequotas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Data source for managing an AWS Service Quotas Templates.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicequotas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := servicequotas.GetTemplates(ctx, &servicequotas.GetTemplatesArgs{
//				Region: "us-east-1",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetTemplates(ctx *pulumi.Context, args *GetTemplatesArgs, opts ...pulumi.InvokeOption) (*GetTemplatesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetTemplatesResult
	err := ctx.Invoke("aws:servicequotas/getTemplates:getTemplates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTemplates.
type GetTemplatesArgs struct {
	// AWS Region to which the quota increases apply.
	Region string `pulumi:"region"`
	// A list of quota increase templates for specified region. See `templates`.
	Templates []GetTemplatesTemplate `pulumi:"templates"`
}

// A collection of values returned by getTemplates.
type GetTemplatesResult struct {
	Id string `pulumi:"id"`
	// AWS Region to which the template applies.
	Region string `pulumi:"region"`
	// A list of quota increase templates for specified region. See `templates`.
	Templates []GetTemplatesTemplate `pulumi:"templates"`
}

func GetTemplatesOutput(ctx *pulumi.Context, args GetTemplatesOutputArgs, opts ...pulumi.InvokeOption) GetTemplatesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTemplatesResult, error) {
			args := v.(GetTemplatesArgs)
			r, err := GetTemplates(ctx, &args, opts...)
			var s GetTemplatesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTemplatesResultOutput)
}

// A collection of arguments for invoking getTemplates.
type GetTemplatesOutputArgs struct {
	// AWS Region to which the quota increases apply.
	Region pulumi.StringInput `pulumi:"region"`
	// A list of quota increase templates for specified region. See `templates`.
	Templates GetTemplatesTemplateArrayInput `pulumi:"templates"`
}

func (GetTemplatesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTemplatesArgs)(nil)).Elem()
}

// A collection of values returned by getTemplates.
type GetTemplatesResultOutput struct{ *pulumi.OutputState }

func (GetTemplatesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTemplatesResult)(nil)).Elem()
}

func (o GetTemplatesResultOutput) ToGetTemplatesResultOutput() GetTemplatesResultOutput {
	return o
}

func (o GetTemplatesResultOutput) ToGetTemplatesResultOutputWithContext(ctx context.Context) GetTemplatesResultOutput {
	return o
}

func (o GetTemplatesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTemplatesResult) string { return v.Id }).(pulumi.StringOutput)
}

// AWS Region to which the template applies.
func (o GetTemplatesResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetTemplatesResult) string { return v.Region }).(pulumi.StringOutput)
}

// A list of quota increase templates for specified region. See `templates`.
func (o GetTemplatesResultOutput) Templates() GetTemplatesTemplateArrayOutput {
	return o.ApplyT(func(v GetTemplatesResult) []GetTemplatesTemplate { return v.Templates }).(GetTemplatesTemplateArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTemplatesResultOutput{})
}
