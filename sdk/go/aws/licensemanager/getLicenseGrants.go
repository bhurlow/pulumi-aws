// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package licensemanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource can be used to get a set of license grant ARNs matching a filter.
//
// ## Example Usage
//
// The following shows getting all license grant ARNs granted to your account.
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/licensemanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			current, err := aws.GetCallerIdentity(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = licensemanager.GetLicenseGrants(ctx, &licensemanager.GetLicenseGrantsArgs{
//				Filters: []licensemanager.GetLicenseGrantsFilter{
//					{
//						Name: "GranteePrincipalARN",
//						Values: []string{
//							fmt.Sprintf("arn:aws:iam::%v:root", current.AccountId),
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
func GetLicenseGrants(ctx *pulumi.Context, args *GetLicenseGrantsArgs, opts ...pulumi.InvokeOption) (*GetLicenseGrantsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLicenseGrantsResult
	err := ctx.Invoke("aws:licensemanager/getLicenseGrants:getLicenseGrants", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLicenseGrants.
type GetLicenseGrantsArgs struct {
	// Custom filter block as described below.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	Filters []GetLicenseGrantsFilter `pulumi:"filters"`
}

// A collection of values returned by getLicenseGrants.
type GetLicenseGrantsResult struct {
	// List of all the license grant ARNs found.
	Arns    []string                 `pulumi:"arns"`
	Filters []GetLicenseGrantsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}

func GetLicenseGrantsOutput(ctx *pulumi.Context, args GetLicenseGrantsOutputArgs, opts ...pulumi.InvokeOption) GetLicenseGrantsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLicenseGrantsResult, error) {
			args := v.(GetLicenseGrantsArgs)
			r, err := GetLicenseGrants(ctx, &args, opts...)
			var s GetLicenseGrantsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLicenseGrantsResultOutput)
}

// A collection of arguments for invoking getLicenseGrants.
type GetLicenseGrantsOutputArgs struct {
	// Custom filter block as described below.
	//
	// More complex filters can be expressed using one or more `filter` sub-blocks,
	// which take the following arguments:
	Filters GetLicenseGrantsFilterArrayInput `pulumi:"filters"`
}

func (GetLicenseGrantsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLicenseGrantsArgs)(nil)).Elem()
}

// A collection of values returned by getLicenseGrants.
type GetLicenseGrantsResultOutput struct{ *pulumi.OutputState }

func (GetLicenseGrantsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLicenseGrantsResult)(nil)).Elem()
}

func (o GetLicenseGrantsResultOutput) ToGetLicenseGrantsResultOutput() GetLicenseGrantsResultOutput {
	return o
}

func (o GetLicenseGrantsResultOutput) ToGetLicenseGrantsResultOutputWithContext(ctx context.Context) GetLicenseGrantsResultOutput {
	return o
}

// List of all the license grant ARNs found.
func (o GetLicenseGrantsResultOutput) Arns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLicenseGrantsResult) []string { return v.Arns }).(pulumi.StringArrayOutput)
}

func (o GetLicenseGrantsResultOutput) Filters() GetLicenseGrantsFilterArrayOutput {
	return o.ApplyT(func(v GetLicenseGrantsResult) []GetLicenseGrantsFilter { return v.Filters }).(GetLicenseGrantsFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLicenseGrantsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLicenseGrantsResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLicenseGrantsResultOutput{})
}
