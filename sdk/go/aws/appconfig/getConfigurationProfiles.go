// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appconfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides access to all Configuration Properties for an AppConfig Application. This will allow you to pass Configuration
// Profile IDs to another resource.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appconfig"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleConfigurationProfiles, err := appconfig.GetConfigurationProfiles(ctx, &appconfig.GetConfigurationProfilesArgs{
//				ApplicationId: "a1d3rpe",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_ := "TODO: For expression"
//			return nil
//		})
//	}
//
// ```
func GetConfigurationProfiles(ctx *pulumi.Context, args *GetConfigurationProfilesArgs, opts ...pulumi.InvokeOption) (*GetConfigurationProfilesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetConfigurationProfilesResult
	err := ctx.Invoke("aws:appconfig/getConfigurationProfiles:getConfigurationProfiles", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getConfigurationProfiles.
type GetConfigurationProfilesArgs struct {
	// ID of the AppConfig Application.
	ApplicationId string `pulumi:"applicationId"`
}

// A collection of values returned by getConfigurationProfiles.
type GetConfigurationProfilesResult struct {
	ApplicationId string `pulumi:"applicationId"`
	// Set of Configuration Profile IDs associated with the AppConfig Application.
	ConfigurationProfileIds []string `pulumi:"configurationProfileIds"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}

func GetConfigurationProfilesOutput(ctx *pulumi.Context, args GetConfigurationProfilesOutputArgs, opts ...pulumi.InvokeOption) GetConfigurationProfilesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetConfigurationProfilesResult, error) {
			args := v.(GetConfigurationProfilesArgs)
			r, err := GetConfigurationProfiles(ctx, &args, opts...)
			var s GetConfigurationProfilesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetConfigurationProfilesResultOutput)
}

// A collection of arguments for invoking getConfigurationProfiles.
type GetConfigurationProfilesOutputArgs struct {
	// ID of the AppConfig Application.
	ApplicationId pulumi.StringInput `pulumi:"applicationId"`
}

func (GetConfigurationProfilesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetConfigurationProfilesArgs)(nil)).Elem()
}

// A collection of values returned by getConfigurationProfiles.
type GetConfigurationProfilesResultOutput struct{ *pulumi.OutputState }

func (GetConfigurationProfilesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetConfigurationProfilesResult)(nil)).Elem()
}

func (o GetConfigurationProfilesResultOutput) ToGetConfigurationProfilesResultOutput() GetConfigurationProfilesResultOutput {
	return o
}

func (o GetConfigurationProfilesResultOutput) ToGetConfigurationProfilesResultOutputWithContext(ctx context.Context) GetConfigurationProfilesResultOutput {
	return o
}

func (o GetConfigurationProfilesResultOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetConfigurationProfilesResult) string { return v.ApplicationId }).(pulumi.StringOutput)
}

// Set of Configuration Profile IDs associated with the AppConfig Application.
func (o GetConfigurationProfilesResultOutput) ConfigurationProfileIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetConfigurationProfilesResult) []string { return v.ConfigurationProfileIds }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetConfigurationProfilesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetConfigurationProfilesResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetConfigurationProfilesResultOutput{})
}
