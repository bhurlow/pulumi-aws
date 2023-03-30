// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Information about a single RDS Reserved Instance Offering.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rds.GetReservedInstanceOffering(ctx, &rds.GetReservedInstanceOfferingArgs{
//				DbInstanceClass:    "db.t2.micro",
//				Duration:           31536000,
//				MultiAz:            false,
//				OfferingType:       "All Upfront",
//				ProductDescription: "mysql",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetReservedInstanceOffering(ctx *pulumi.Context, args *GetReservedInstanceOfferingArgs, opts ...pulumi.InvokeOption) (*GetReservedInstanceOfferingResult, error) {
	var rv GetReservedInstanceOfferingResult
	err := ctx.Invoke("aws:rds/getReservedInstanceOffering:getReservedInstanceOffering", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getReservedInstanceOffering.
type GetReservedInstanceOfferingArgs struct {
	// DB instance class for the reserved DB instance.
	DbInstanceClass string `pulumi:"dbInstanceClass"`
	// Duration of the reservation in years or seconds. Valid values are `1`, `3`, `31536000`, `94608000`
	Duration int `pulumi:"duration"`
	// Whether the reservation applies to Multi-AZ deployments.
	MultiAz bool `pulumi:"multiAz"`
	// Offering type of this reserved DB instance. Valid values are `No Upfront`, `Partial Upfront`, `All Upfront`.
	OfferingType string `pulumi:"offeringType"`
	// Description of the reserved DB instance.
	ProductDescription string `pulumi:"productDescription"`
}

// A collection of values returned by getReservedInstanceOffering.
type GetReservedInstanceOfferingResult struct {
	// Currency code for the reserved DB instance.
	CurrencyCode    string `pulumi:"currencyCode"`
	DbInstanceClass string `pulumi:"dbInstanceClass"`
	Duration        int    `pulumi:"duration"`
	// Fixed price charged for this reserved DB instance.
	FixedPrice float64 `pulumi:"fixedPrice"`
	// The provider-assigned unique ID for this managed resource.
	Id      string `pulumi:"id"`
	MultiAz bool   `pulumi:"multiAz"`
	// Unique identifier for the reservation.
	OfferingId         string `pulumi:"offeringId"`
	OfferingType       string `pulumi:"offeringType"`
	ProductDescription string `pulumi:"productDescription"`
}

func GetReservedInstanceOfferingOutput(ctx *pulumi.Context, args GetReservedInstanceOfferingOutputArgs, opts ...pulumi.InvokeOption) GetReservedInstanceOfferingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetReservedInstanceOfferingResult, error) {
			args := v.(GetReservedInstanceOfferingArgs)
			r, err := GetReservedInstanceOffering(ctx, &args, opts...)
			var s GetReservedInstanceOfferingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetReservedInstanceOfferingResultOutput)
}

// A collection of arguments for invoking getReservedInstanceOffering.
type GetReservedInstanceOfferingOutputArgs struct {
	// DB instance class for the reserved DB instance.
	DbInstanceClass pulumi.StringInput `pulumi:"dbInstanceClass"`
	// Duration of the reservation in years or seconds. Valid values are `1`, `3`, `31536000`, `94608000`
	Duration pulumi.IntInput `pulumi:"duration"`
	// Whether the reservation applies to Multi-AZ deployments.
	MultiAz pulumi.BoolInput `pulumi:"multiAz"`
	// Offering type of this reserved DB instance. Valid values are `No Upfront`, `Partial Upfront`, `All Upfront`.
	OfferingType pulumi.StringInput `pulumi:"offeringType"`
	// Description of the reserved DB instance.
	ProductDescription pulumi.StringInput `pulumi:"productDescription"`
}

func (GetReservedInstanceOfferingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetReservedInstanceOfferingArgs)(nil)).Elem()
}

// A collection of values returned by getReservedInstanceOffering.
type GetReservedInstanceOfferingResultOutput struct{ *pulumi.OutputState }

func (GetReservedInstanceOfferingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetReservedInstanceOfferingResult)(nil)).Elem()
}

func (o GetReservedInstanceOfferingResultOutput) ToGetReservedInstanceOfferingResultOutput() GetReservedInstanceOfferingResultOutput {
	return o
}

func (o GetReservedInstanceOfferingResultOutput) ToGetReservedInstanceOfferingResultOutputWithContext(ctx context.Context) GetReservedInstanceOfferingResultOutput {
	return o
}

// Currency code for the reserved DB instance.
func (o GetReservedInstanceOfferingResultOutput) CurrencyCode() pulumi.StringOutput {
	return o.ApplyT(func(v GetReservedInstanceOfferingResult) string { return v.CurrencyCode }).(pulumi.StringOutput)
}

func (o GetReservedInstanceOfferingResultOutput) DbInstanceClass() pulumi.StringOutput {
	return o.ApplyT(func(v GetReservedInstanceOfferingResult) string { return v.DbInstanceClass }).(pulumi.StringOutput)
}

func (o GetReservedInstanceOfferingResultOutput) Duration() pulumi.IntOutput {
	return o.ApplyT(func(v GetReservedInstanceOfferingResult) int { return v.Duration }).(pulumi.IntOutput)
}

// Fixed price charged for this reserved DB instance.
func (o GetReservedInstanceOfferingResultOutput) FixedPrice() pulumi.Float64Output {
	return o.ApplyT(func(v GetReservedInstanceOfferingResult) float64 { return v.FixedPrice }).(pulumi.Float64Output)
}

// The provider-assigned unique ID for this managed resource.
func (o GetReservedInstanceOfferingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetReservedInstanceOfferingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetReservedInstanceOfferingResultOutput) MultiAz() pulumi.BoolOutput {
	return o.ApplyT(func(v GetReservedInstanceOfferingResult) bool { return v.MultiAz }).(pulumi.BoolOutput)
}

// Unique identifier for the reservation.
func (o GetReservedInstanceOfferingResultOutput) OfferingId() pulumi.StringOutput {
	return o.ApplyT(func(v GetReservedInstanceOfferingResult) string { return v.OfferingId }).(pulumi.StringOutput)
}

func (o GetReservedInstanceOfferingResultOutput) OfferingType() pulumi.StringOutput {
	return o.ApplyT(func(v GetReservedInstanceOfferingResult) string { return v.OfferingType }).(pulumi.StringOutput)
}

func (o GetReservedInstanceOfferingResultOutput) ProductDescription() pulumi.StringOutput {
	return o.ApplyT(func(v GetReservedInstanceOfferingResult) string { return v.ProductDescription }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetReservedInstanceOfferingResultOutput{})
}
