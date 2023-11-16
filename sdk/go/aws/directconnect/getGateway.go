// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve information about a Direct Connect Gateway.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directconnect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := directconnect.LookupGateway(ctx, &directconnect.LookupGatewayArgs{
//				Name: "example",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupGateway(ctx *pulumi.Context, args *LookupGatewayArgs, opts ...pulumi.InvokeOption) (*LookupGatewayResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGatewayResult
	err := ctx.Invoke("aws:directconnect/getGateway:getGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGateway.
type LookupGatewayArgs struct {
	// Name of the gateway to retrieve.
	Name string `pulumi:"name"`
}

// A collection of values returned by getGateway.
type LookupGatewayResult struct {
	// ASN on the Amazon side of the connection.
	AmazonSideAsn string `pulumi:"amazonSideAsn"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// AWS Account ID of the gateway.
	OwnerAccountId string `pulumi:"ownerAccountId"`
}

func LookupGatewayOutput(ctx *pulumi.Context, args LookupGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayResult, error) {
			args := v.(LookupGatewayArgs)
			r, err := LookupGateway(ctx, &args, opts...)
			var s LookupGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayResultOutput)
}

// A collection of arguments for invoking getGateway.
type LookupGatewayOutputArgs struct {
	// Name of the gateway to retrieve.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayArgs)(nil)).Elem()
}

// A collection of values returned by getGateway.
type LookupGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayResult)(nil)).Elem()
}

func (o LookupGatewayResultOutput) ToLookupGatewayResultOutput() LookupGatewayResultOutput {
	return o
}

func (o LookupGatewayResultOutput) ToLookupGatewayResultOutputWithContext(ctx context.Context) LookupGatewayResultOutput {
	return o
}

// ASN on the Amazon side of the connection.
func (o LookupGatewayResultOutput) AmazonSideAsn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.AmazonSideAsn }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

// AWS Account ID of the gateway.
func (o LookupGatewayResultOutput) OwnerAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.OwnerAccountId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayResultOutput{})
}
