// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package imagebuilder

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ARNs and names of Image Builder Container Recipes matching the specified criteria.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/imagebuilder"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := imagebuilder.GetContainerRecipes(ctx, &imagebuilder.GetContainerRecipesArgs{
//				Filters: []imagebuilder.GetContainerRecipesFilter{
//					{
//						Name: "platform",
//						Values: []string{
//							"Linux",
//						},
//					},
//				},
//				Owner: pulumi.StringRef("Self"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetContainerRecipes(ctx *pulumi.Context, args *GetContainerRecipesArgs, opts ...pulumi.InvokeOption) (*GetContainerRecipesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetContainerRecipesResult
	err := ctx.Invoke("aws:imagebuilder/getContainerRecipes:getContainerRecipes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContainerRecipes.
type GetContainerRecipesArgs struct {
	// Configuration block(s) for filtering. Detailed below.
	Filters []GetContainerRecipesFilter `pulumi:"filters"`
	// Owner of the container recipes. Valid values are `Self`, `Shared` and `Amazon`. Defaults to `Self`.
	Owner *string `pulumi:"owner"`
}

// A collection of values returned by getContainerRecipes.
type GetContainerRecipesResult struct {
	// Set of ARNs of the matched Image Builder Container Recipes.
	Arns    []string                    `pulumi:"arns"`
	Filters []GetContainerRecipesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Set of names of the matched Image Builder Container Recipes.
	Names []string `pulumi:"names"`
	Owner *string  `pulumi:"owner"`
}

func GetContainerRecipesOutput(ctx *pulumi.Context, args GetContainerRecipesOutputArgs, opts ...pulumi.InvokeOption) GetContainerRecipesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetContainerRecipesResult, error) {
			args := v.(GetContainerRecipesArgs)
			r, err := GetContainerRecipes(ctx, &args, opts...)
			var s GetContainerRecipesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetContainerRecipesResultOutput)
}

// A collection of arguments for invoking getContainerRecipes.
type GetContainerRecipesOutputArgs struct {
	// Configuration block(s) for filtering. Detailed below.
	Filters GetContainerRecipesFilterArrayInput `pulumi:"filters"`
	// Owner of the container recipes. Valid values are `Self`, `Shared` and `Amazon`. Defaults to `Self`.
	Owner pulumi.StringPtrInput `pulumi:"owner"`
}

func (GetContainerRecipesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContainerRecipesArgs)(nil)).Elem()
}

// A collection of values returned by getContainerRecipes.
type GetContainerRecipesResultOutput struct{ *pulumi.OutputState }

func (GetContainerRecipesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContainerRecipesResult)(nil)).Elem()
}

func (o GetContainerRecipesResultOutput) ToGetContainerRecipesResultOutput() GetContainerRecipesResultOutput {
	return o
}

func (o GetContainerRecipesResultOutput) ToGetContainerRecipesResultOutputWithContext(ctx context.Context) GetContainerRecipesResultOutput {
	return o
}

// Set of ARNs of the matched Image Builder Container Recipes.
func (o GetContainerRecipesResultOutput) Arns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetContainerRecipesResult) []string { return v.Arns }).(pulumi.StringArrayOutput)
}

func (o GetContainerRecipesResultOutput) Filters() GetContainerRecipesFilterArrayOutput {
	return o.ApplyT(func(v GetContainerRecipesResult) []GetContainerRecipesFilter { return v.Filters }).(GetContainerRecipesFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetContainerRecipesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerRecipesResult) string { return v.Id }).(pulumi.StringOutput)
}

// Set of names of the matched Image Builder Container Recipes.
func (o GetContainerRecipesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetContainerRecipesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetContainerRecipesResultOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetContainerRecipesResult) *string { return v.Owner }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetContainerRecipesResultOutput{})
}
