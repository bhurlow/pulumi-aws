// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can be used to fetch information about a specific
// EventBridge event bus. Use this data source to compute the ARN of
// an event bus, given the name of the bus.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudwatch.LookupEventBus(ctx, &cloudwatch.LookupEventBusArgs{
// 			Name: "example-bus-name",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupEventBus(ctx *pulumi.Context, args *LookupEventBusArgs, opts ...pulumi.InvokeOption) (*LookupEventBusResult, error) {
	var rv LookupEventBusResult
	err := ctx.Invoke("aws:cloudwatch/getEventBus:getEventBus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEventBus.
type LookupEventBusArgs struct {
	// The friendly EventBridge event bus name.
	Name string `pulumi:"name"`
}

// A collection of values returned by getEventBus.
type LookupEventBusResult struct {
	// The Amazon Resource Name (ARN) specifying the role.
	Arn string `pulumi:"arn"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}

func LookupEventBusOutput(ctx *pulumi.Context, args LookupEventBusOutputArgs, opts ...pulumi.InvokeOption) LookupEventBusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventBusResult, error) {
			args := v.(LookupEventBusArgs)
			r, err := LookupEventBus(ctx, &args, opts...)
			var s LookupEventBusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEventBusResultOutput)
}

// A collection of arguments for invoking getEventBus.
type LookupEventBusOutputArgs struct {
	// The friendly EventBridge event bus name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupEventBusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventBusArgs)(nil)).Elem()
}

// A collection of values returned by getEventBus.
type LookupEventBusResultOutput struct{ *pulumi.OutputState }

func (LookupEventBusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventBusResult)(nil)).Elem()
}

func (o LookupEventBusResultOutput) ToLookupEventBusResultOutput() LookupEventBusResultOutput {
	return o
}

func (o LookupEventBusResultOutput) ToLookupEventBusResultOutputWithContext(ctx context.Context) LookupEventBusResultOutput {
	return o
}

// The Amazon Resource Name (ARN) specifying the role.
func (o LookupEventBusResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventBusResult) string { return v.Arn }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupEventBusResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventBusResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEventBusResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEventBusResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventBusResultOutput{})
}