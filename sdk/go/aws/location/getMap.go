// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package location

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve information about a Location Service Map.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/location"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := location.LookupMap(ctx, &location.LookupMapArgs{
// 			MapName: "example",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupMap(ctx *pulumi.Context, args *LookupMapArgs, opts ...pulumi.InvokeOption) (*LookupMapResult, error) {
	var rv LookupMapResult
	err := ctx.Invoke("aws:location/getMap:getMap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMap.
type LookupMapArgs struct {
	// The name of the map resource.
	MapName string `pulumi:"mapName"`
	// Key-value map of resource tags for the map.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getMap.
type LookupMapResult struct {
	// List of configurations that specify the map tile style selected from a partner data provider.
	Configurations []GetMapConfiguration `pulumi:"configurations"`
	// The timestamp for when the map resource was created in ISO 8601 format.
	CreateTime string `pulumi:"createTime"`
	// The optional description for the map resource.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Amazon Resource Name (ARN) for the map resource.
	MapArn  string `pulumi:"mapArn"`
	MapName string `pulumi:"mapName"`
	// Key-value map of resource tags for the map.
	Tags map[string]string `pulumi:"tags"`
	// The timestamp for when the map resource was last updated in ISO 8601.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupMapOutput(ctx *pulumi.Context, args LookupMapOutputArgs, opts ...pulumi.InvokeOption) LookupMapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMapResult, error) {
			args := v.(LookupMapArgs)
			r, err := LookupMap(ctx, &args, opts...)
			var s LookupMapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMapResultOutput)
}

// A collection of arguments for invoking getMap.
type LookupMapOutputArgs struct {
	// The name of the map resource.
	MapName pulumi.StringInput `pulumi:"mapName"`
	// Key-value map of resource tags for the map.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupMapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMapArgs)(nil)).Elem()
}

// A collection of values returned by getMap.
type LookupMapResultOutput struct{ *pulumi.OutputState }

func (LookupMapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMapResult)(nil)).Elem()
}

func (o LookupMapResultOutput) ToLookupMapResultOutput() LookupMapResultOutput {
	return o
}

func (o LookupMapResultOutput) ToLookupMapResultOutputWithContext(ctx context.Context) LookupMapResultOutput {
	return o
}

// List of configurations that specify the map tile style selected from a partner data provider.
func (o LookupMapResultOutput) Configurations() GetMapConfigurationArrayOutput {
	return o.ApplyT(func(v LookupMapResult) []GetMapConfiguration { return v.Configurations }).(GetMapConfigurationArrayOutput)
}

// The timestamp for when the map resource was created in ISO 8601 format.
func (o LookupMapResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMapResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The optional description for the map resource.
func (o LookupMapResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMapResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMapResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) for the map resource.
func (o LookupMapResultOutput) MapArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMapResult) string { return v.MapArn }).(pulumi.StringOutput)
}

func (o LookupMapResultOutput) MapName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMapResult) string { return v.MapName }).(pulumi.StringOutput)
}

// Key-value map of resource tags for the map.
func (o LookupMapResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMapResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The timestamp for when the map resource was last updated in ISO 8601.
func (o LookupMapResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMapResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMapResultOutput{})
}
