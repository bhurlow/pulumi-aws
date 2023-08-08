// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connect

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about a specific Amazon Connect Instance.
//
// ## Example Usage
//
// # By instanceAlias
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := connect.LookupInstance(ctx, &connect.LookupInstanceArgs{
//				InstanceAlias: pulumi.StringRef("foo"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// # By instanceId
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := connect.LookupInstance(ctx, &connect.LookupInstanceArgs{
//				InstanceId: pulumi.StringRef("97afc98d-101a-ba98-ab97-ae114fc115ec"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInstanceResult
	err := ctx.Invoke("aws:connect/getInstance:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstance.
type LookupInstanceArgs struct {
	// Returns information on a specific connect instance by alias
	InstanceAlias *string `pulumi:"instanceAlias"`
	// Returns information on a specific connect instance by id
	InstanceId *string `pulumi:"instanceId"`
}

// A collection of values returned by getInstance.
type LookupInstanceResult struct {
	// ARN of the instance.
	Arn                          string `pulumi:"arn"`
	AutoResolveBestVoicesEnabled bool   `pulumi:"autoResolveBestVoicesEnabled"`
	// Whether contact flow logs are enabled.
	ContactFlowLogsEnabled bool `pulumi:"contactFlowLogsEnabled"`
	// Whether contact lens is enabled.
	ContactLensEnabled bool `pulumi:"contactLensEnabled"`
	// When the instance was created.
	CreatedTime string `pulumi:"createdTime"`
	// Whether early media for outbound calls is enabled .
	EarlyMediaEnabled bool `pulumi:"earlyMediaEnabled"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Specifies The identity management type attached to the instance.
	IdentityManagementType string `pulumi:"identityManagementType"`
	// Whether inbound calls are enabled.
	InboundCallsEnabled bool   `pulumi:"inboundCallsEnabled"`
	InstanceAlias       string `pulumi:"instanceAlias"`
	InstanceId          string `pulumi:"instanceId"`
	// Whether multi-party calls/conference is enabled.
	MultiPartyConferenceEnabled bool `pulumi:"multiPartyConferenceEnabled"`
	// Whether outbound calls are enabled.
	OutboundCallsEnabled bool `pulumi:"outboundCallsEnabled"`
	// Service role of the instance.
	ServiceRole string `pulumi:"serviceRole"`
	// State of the instance.
	Status string `pulumi:"status"`
}

func LookupInstanceOutput(ctx *pulumi.Context, args LookupInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceResult, error) {
			args := v.(LookupInstanceArgs)
			r, err := LookupInstance(ctx, &args, opts...)
			var s LookupInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstanceResultOutput)
}

// A collection of arguments for invoking getInstance.
type LookupInstanceOutputArgs struct {
	// Returns information on a specific connect instance by alias
	InstanceAlias pulumi.StringPtrInput `pulumi:"instanceAlias"`
	// Returns information on a specific connect instance by id
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
}

func (LookupInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getInstance.
type LookupInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceResult)(nil)).Elem()
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutput() LookupInstanceResultOutput {
	return o
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutputWithContext(ctx context.Context) LookupInstanceResultOutput {
	return o
}

// ARN of the instance.
func (o LookupInstanceResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Arn }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) AutoResolveBestVoicesEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.AutoResolveBestVoicesEnabled }).(pulumi.BoolOutput)
}

// Whether contact flow logs are enabled.
func (o LookupInstanceResultOutput) ContactFlowLogsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.ContactFlowLogsEnabled }).(pulumi.BoolOutput)
}

// Whether contact lens is enabled.
func (o LookupInstanceResultOutput) ContactLensEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.ContactLensEnabled }).(pulumi.BoolOutput)
}

// When the instance was created.
func (o LookupInstanceResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

// Whether early media for outbound calls is enabled .
func (o LookupInstanceResultOutput) EarlyMediaEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.EarlyMediaEnabled }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Specifies The identity management type attached to the instance.
func (o LookupInstanceResultOutput) IdentityManagementType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.IdentityManagementType }).(pulumi.StringOutput)
}

// Whether inbound calls are enabled.
func (o LookupInstanceResultOutput) InboundCallsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.InboundCallsEnabled }).(pulumi.BoolOutput)
}

func (o LookupInstanceResultOutput) InstanceAlias() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.InstanceAlias }).(pulumi.StringOutput)
}

func (o LookupInstanceResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// Whether multi-party calls/conference is enabled.
func (o LookupInstanceResultOutput) MultiPartyConferenceEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.MultiPartyConferenceEnabled }).(pulumi.BoolOutput)
}

// Whether outbound calls are enabled.
func (o LookupInstanceResultOutput) OutboundCallsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.OutboundCallsEnabled }).(pulumi.BoolOutput)
}

// Service role of the instance.
func (o LookupInstanceResultOutput) ServiceRole() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.ServiceRole }).(pulumi.StringOutput)
}

// State of the instance.
func (o LookupInstanceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceResultOutput{})
}
