// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2transitgateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides informations for routes of a specific transit gateway, such as state, type, cidr
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2transitgateway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2transitgateway.GetRouteTableRoutes(ctx, &ec2transitgateway.GetRouteTableRoutesArgs{
//				Filters: []ec2transitgateway.GetRouteTableRoutesFilter{
//					{
//						Name: "type",
//						Values: []string{
//							"propagated",
//						},
//					},
//				},
//				TransitGatewayRouteTableId: aws_ec2_transit_gateway_route_table.Example.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetRouteTableRoutes(ctx *pulumi.Context, args *GetRouteTableRoutesArgs, opts ...pulumi.InvokeOption) (*GetRouteTableRoutesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRouteTableRoutesResult
	err := ctx.Invoke("aws:ec2transitgateway/getRouteTableRoutes:getRouteTableRoutes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouteTableRoutes.
type GetRouteTableRoutesArgs struct {
	// Custom filter block as described below.
	Filters []GetRouteTableRoutesFilter `pulumi:"filters"`
	// Identifier of EC2 Transit Gateway Route Table.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	TransitGatewayRouteTableId string `pulumi:"transitGatewayRouteTableId"`
}

// A collection of values returned by getRouteTableRoutes.
type GetRouteTableRoutesResult struct {
	Filters []GetRouteTableRoutesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of Transit Gateway Routes.
	Routes                     []GetRouteTableRoutesRoute `pulumi:"routes"`
	TransitGatewayRouteTableId string                     `pulumi:"transitGatewayRouteTableId"`
}

func GetRouteTableRoutesOutput(ctx *pulumi.Context, args GetRouteTableRoutesOutputArgs, opts ...pulumi.InvokeOption) GetRouteTableRoutesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRouteTableRoutesResult, error) {
			args := v.(GetRouteTableRoutesArgs)
			r, err := GetRouteTableRoutes(ctx, &args, opts...)
			var s GetRouteTableRoutesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRouteTableRoutesResultOutput)
}

// A collection of arguments for invoking getRouteTableRoutes.
type GetRouteTableRoutesOutputArgs struct {
	// Custom filter block as described below.
	Filters GetRouteTableRoutesFilterArrayInput `pulumi:"filters"`
	// Identifier of EC2 Transit Gateway Route Table.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	TransitGatewayRouteTableId pulumi.StringInput `pulumi:"transitGatewayRouteTableId"`
}

func (GetRouteTableRoutesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouteTableRoutesArgs)(nil)).Elem()
}

// A collection of values returned by getRouteTableRoutes.
type GetRouteTableRoutesResultOutput struct{ *pulumi.OutputState }

func (GetRouteTableRoutesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouteTableRoutesResult)(nil)).Elem()
}

func (o GetRouteTableRoutesResultOutput) ToGetRouteTableRoutesResultOutput() GetRouteTableRoutesResultOutput {
	return o
}

func (o GetRouteTableRoutesResultOutput) ToGetRouteTableRoutesResultOutputWithContext(ctx context.Context) GetRouteTableRoutesResultOutput {
	return o
}

func (o GetRouteTableRoutesResultOutput) Filters() GetRouteTableRoutesFilterArrayOutput {
	return o.ApplyT(func(v GetRouteTableRoutesResult) []GetRouteTableRoutesFilter { return v.Filters }).(GetRouteTableRoutesFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRouteTableRoutesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRouteTableRoutesResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of Transit Gateway Routes.
func (o GetRouteTableRoutesResultOutput) Routes() GetRouteTableRoutesRouteArrayOutput {
	return o.ApplyT(func(v GetRouteTableRoutesResult) []GetRouteTableRoutesRoute { return v.Routes }).(GetRouteTableRoutesRouteArrayOutput)
}

func (o GetRouteTableRoutesResultOutput) TransitGatewayRouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRouteTableRoutesResult) string { return v.TransitGatewayRouteTableId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRouteTableRoutesResultOutput{})
}
