// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the window IDs of SSM maintenance windows.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ssm.GetMaintenanceWindows(ctx, &ssm.GetMaintenanceWindowsArgs{
//				Filters: []ssm.GetMaintenanceWindowsFilter{
//					{
//						Name: "Enabled",
//						Values: []string{
//							"true",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetMaintenanceWindows(ctx *pulumi.Context, args *GetMaintenanceWindowsArgs, opts ...pulumi.InvokeOption) (*GetMaintenanceWindowsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMaintenanceWindowsResult
	err := ctx.Invoke("aws:ssm/getMaintenanceWindows:getMaintenanceWindows", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMaintenanceWindows.
type GetMaintenanceWindowsArgs struct {
	// Configuration block(s) for filtering. Detailed below.
	Filters []GetMaintenanceWindowsFilter `pulumi:"filters"`
}

// A collection of values returned by getMaintenanceWindows.
type GetMaintenanceWindowsResult struct {
	Filters []GetMaintenanceWindowsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of window IDs of the matched SSM maintenance windows.
	Ids []string `pulumi:"ids"`
}

func GetMaintenanceWindowsOutput(ctx *pulumi.Context, args GetMaintenanceWindowsOutputArgs, opts ...pulumi.InvokeOption) GetMaintenanceWindowsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMaintenanceWindowsResult, error) {
			args := v.(GetMaintenanceWindowsArgs)
			r, err := GetMaintenanceWindows(ctx, &args, opts...)
			var s GetMaintenanceWindowsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMaintenanceWindowsResultOutput)
}

// A collection of arguments for invoking getMaintenanceWindows.
type GetMaintenanceWindowsOutputArgs struct {
	// Configuration block(s) for filtering. Detailed below.
	Filters GetMaintenanceWindowsFilterArrayInput `pulumi:"filters"`
}

func (GetMaintenanceWindowsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMaintenanceWindowsArgs)(nil)).Elem()
}

// A collection of values returned by getMaintenanceWindows.
type GetMaintenanceWindowsResultOutput struct{ *pulumi.OutputState }

func (GetMaintenanceWindowsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMaintenanceWindowsResult)(nil)).Elem()
}

func (o GetMaintenanceWindowsResultOutput) ToGetMaintenanceWindowsResultOutput() GetMaintenanceWindowsResultOutput {
	return o
}

func (o GetMaintenanceWindowsResultOutput) ToGetMaintenanceWindowsResultOutputWithContext(ctx context.Context) GetMaintenanceWindowsResultOutput {
	return o
}

func (o GetMaintenanceWindowsResultOutput) Filters() GetMaintenanceWindowsFilterArrayOutput {
	return o.ApplyT(func(v GetMaintenanceWindowsResult) []GetMaintenanceWindowsFilter { return v.Filters }).(GetMaintenanceWindowsFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetMaintenanceWindowsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMaintenanceWindowsResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of window IDs of the matched SSM maintenance windows.
func (o GetMaintenanceWindowsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetMaintenanceWindowsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMaintenanceWindowsResultOutput{})
}
