// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package location

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve information about a Location Service Place Index.
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
// 		_, err := location.LookupPlaceIndex(ctx, &location.LookupPlaceIndexArgs{
// 			IndexName: "example",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupPlaceIndex(ctx *pulumi.Context, args *LookupPlaceIndexArgs, opts ...pulumi.InvokeOption) (*LookupPlaceIndexResult, error) {
	var rv LookupPlaceIndexResult
	err := ctx.Invoke("aws:location/getPlaceIndex:getPlaceIndex", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPlaceIndex.
type LookupPlaceIndexArgs struct {
	// The name of the place index resource.
	IndexName string `pulumi:"indexName"`
	// Key-value map of resource tags for the map.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getPlaceIndex.
type LookupPlaceIndexResult struct {
	// The timestamp for when the place index resource was created in ISO 8601 format.
	CreateTime string `pulumi:"createTime"`
	// The data provider of geospatial data.
	DataSource string `pulumi:"dataSource"`
	// List of configurations that specify data storage option for requesting Places.
	DataSourceConfigurations []GetPlaceIndexDataSourceConfiguration `pulumi:"dataSourceConfigurations"`
	// The optional description for the place index resource.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Amazon Resource Name (ARN) for the place index resource.
	IndexArn  string `pulumi:"indexArn"`
	IndexName string `pulumi:"indexName"`
	// Key-value map of resource tags for the map.
	Tags map[string]string `pulumi:"tags"`
	// The timestamp for when the place index resource was last update in ISO 8601.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupPlaceIndexOutput(ctx *pulumi.Context, args LookupPlaceIndexOutputArgs, opts ...pulumi.InvokeOption) LookupPlaceIndexResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPlaceIndexResult, error) {
			args := v.(LookupPlaceIndexArgs)
			r, err := LookupPlaceIndex(ctx, &args, opts...)
			var s LookupPlaceIndexResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPlaceIndexResultOutput)
}

// A collection of arguments for invoking getPlaceIndex.
type LookupPlaceIndexOutputArgs struct {
	// The name of the place index resource.
	IndexName pulumi.StringInput `pulumi:"indexName"`
	// Key-value map of resource tags for the map.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupPlaceIndexOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlaceIndexArgs)(nil)).Elem()
}

// A collection of values returned by getPlaceIndex.
type LookupPlaceIndexResultOutput struct{ *pulumi.OutputState }

func (LookupPlaceIndexResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlaceIndexResult)(nil)).Elem()
}

func (o LookupPlaceIndexResultOutput) ToLookupPlaceIndexResultOutput() LookupPlaceIndexResultOutput {
	return o
}

func (o LookupPlaceIndexResultOutput) ToLookupPlaceIndexResultOutputWithContext(ctx context.Context) LookupPlaceIndexResultOutput {
	return o
}

// The timestamp for when the place index resource was created in ISO 8601 format.
func (o LookupPlaceIndexResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlaceIndexResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The data provider of geospatial data.
func (o LookupPlaceIndexResultOutput) DataSource() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlaceIndexResult) string { return v.DataSource }).(pulumi.StringOutput)
}

// List of configurations that specify data storage option for requesting Places.
func (o LookupPlaceIndexResultOutput) DataSourceConfigurations() GetPlaceIndexDataSourceConfigurationArrayOutput {
	return o.ApplyT(func(v LookupPlaceIndexResult) []GetPlaceIndexDataSourceConfiguration {
		return v.DataSourceConfigurations
	}).(GetPlaceIndexDataSourceConfigurationArrayOutput)
}

// The optional description for the place index resource.
func (o LookupPlaceIndexResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlaceIndexResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPlaceIndexResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlaceIndexResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) for the place index resource.
func (o LookupPlaceIndexResultOutput) IndexArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlaceIndexResult) string { return v.IndexArn }).(pulumi.StringOutput)
}

func (o LookupPlaceIndexResultOutput) IndexName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlaceIndexResult) string { return v.IndexName }).(pulumi.StringOutput)
}

// Key-value map of resource tags for the map.
func (o LookupPlaceIndexResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPlaceIndexResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The timestamp for when the place index resource was last update in ISO 8601.
func (o LookupPlaceIndexResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPlaceIndexResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPlaceIndexResultOutput{})
}
