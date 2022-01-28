// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package imagebuilder

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ARNs and names of Image Builder Components matching the specified criteria.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/imagebuilder"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "Self"
// 		_, err := imagebuilder.GetComponents(ctx, &imagebuilder.GetComponentsArgs{
// 			Filters: []imagebuilder.GetComponentsFilter{
// 				imagebuilder.GetComponentsFilter{
// 					Name: "platform",
// 					Values: []string{
// 						"Linux",
// 					},
// 				},
// 			},
// 			Owner: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetComponents(ctx *pulumi.Context, args *GetComponentsArgs, opts ...pulumi.InvokeOption) (*GetComponentsResult, error) {
	var rv GetComponentsResult
	err := ctx.Invoke("aws:imagebuilder/getComponents:getComponents", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComponents.
type GetComponentsArgs struct {
	// Configuration block(s) for filtering. Detailed below.
	Filters []GetComponentsFilter `pulumi:"filters"`
	// The owner of the image recipes. Valid values are `Self`, `Shared` and `Amazon`. Defaults to `Self`.
	Owner *string `pulumi:"owner"`
}

// A collection of values returned by getComponents.
type GetComponentsResult struct {
	// Set of ARNs of the matched Image Builder Components.
	Arns    []string              `pulumi:"arns"`
	Filters []GetComponentsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Set of names of the matched Image Builder Components.
	Names []string `pulumi:"names"`
	Owner *string  `pulumi:"owner"`
}

func GetComponentsOutput(ctx *pulumi.Context, args GetComponentsOutputArgs, opts ...pulumi.InvokeOption) GetComponentsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetComponentsResult, error) {
			args := v.(GetComponentsArgs)
			r, err := GetComponents(ctx, &args, opts...)
			return *r, err
		}).(GetComponentsResultOutput)
}

// A collection of arguments for invoking getComponents.
type GetComponentsOutputArgs struct {
	// Configuration block(s) for filtering. Detailed below.
	Filters GetComponentsFilterArrayInput `pulumi:"filters"`
	// The owner of the image recipes. Valid values are `Self`, `Shared` and `Amazon`. Defaults to `Self`.
	Owner pulumi.StringPtrInput `pulumi:"owner"`
}

func (GetComponentsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetComponentsArgs)(nil)).Elem()
}

// A collection of values returned by getComponents.
type GetComponentsResultOutput struct{ *pulumi.OutputState }

func (GetComponentsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetComponentsResult)(nil)).Elem()
}

func (o GetComponentsResultOutput) ToGetComponentsResultOutput() GetComponentsResultOutput {
	return o
}

func (o GetComponentsResultOutput) ToGetComponentsResultOutputWithContext(ctx context.Context) GetComponentsResultOutput {
	return o
}

// Set of ARNs of the matched Image Builder Components.
func (o GetComponentsResultOutput) Arns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetComponentsResult) []string { return v.Arns }).(pulumi.StringArrayOutput)
}

func (o GetComponentsResultOutput) Filters() GetComponentsFilterArrayOutput {
	return o.ApplyT(func(v GetComponentsResult) []GetComponentsFilter { return v.Filters }).(GetComponentsFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetComponentsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetComponentsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Set of names of the matched Image Builder Components.
func (o GetComponentsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetComponentsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetComponentsResultOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetComponentsResult) *string { return v.Owner }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetComponentsResultOutput{})
}
