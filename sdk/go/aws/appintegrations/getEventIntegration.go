// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appintegrations

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information on an existing AppIntegrations Event Integration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/appintegrations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := appintegrations.GetEventIntegration(ctx, &appintegrations.GetEventIntegrationArgs{
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
func GetEventIntegration(ctx *pulumi.Context, args *GetEventIntegrationArgs, opts ...pulumi.InvokeOption) (*GetEventIntegrationResult, error) {
	var rv GetEventIntegrationResult
	err := ctx.Invoke("aws:appintegrations/getEventIntegration:getEventIntegration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEventIntegration.
type GetEventIntegrationArgs struct {
	// The AppIntegrations Event Integration name.
	Name string `pulumi:"name"`
	// Metadata that you can assign to help organize the report plans you create.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getEventIntegration.
type GetEventIntegrationResult struct {
	// The ARN of the AppIntegrations Event Integration.
	Arn string `pulumi:"arn"`
	// The description of the Event Integration.
	Description string `pulumi:"description"`
	// A block that defines the configuration information for the event filter. The Event Filter block is documented below.
	EventFilters []GetEventIntegrationEventFilter `pulumi:"eventFilters"`
	// The EventBridge bus.
	EventbridgeBus string `pulumi:"eventbridgeBus"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Metadata that you can assign to help organize the report plans you create.
	Tags map[string]string `pulumi:"tags"`
}

func GetEventIntegrationOutput(ctx *pulumi.Context, args GetEventIntegrationOutputArgs, opts ...pulumi.InvokeOption) GetEventIntegrationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEventIntegrationResult, error) {
			args := v.(GetEventIntegrationArgs)
			r, err := GetEventIntegration(ctx, &args, opts...)
			var s GetEventIntegrationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEventIntegrationResultOutput)
}

// A collection of arguments for invoking getEventIntegration.
type GetEventIntegrationOutputArgs struct {
	// The AppIntegrations Event Integration name.
	Name pulumi.StringInput `pulumi:"name"`
	// Metadata that you can assign to help organize the report plans you create.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (GetEventIntegrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEventIntegrationArgs)(nil)).Elem()
}

// A collection of values returned by getEventIntegration.
type GetEventIntegrationResultOutput struct{ *pulumi.OutputState }

func (GetEventIntegrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEventIntegrationResult)(nil)).Elem()
}

func (o GetEventIntegrationResultOutput) ToGetEventIntegrationResultOutput() GetEventIntegrationResultOutput {
	return o
}

func (o GetEventIntegrationResultOutput) ToGetEventIntegrationResultOutputWithContext(ctx context.Context) GetEventIntegrationResultOutput {
	return o
}

// The ARN of the AppIntegrations Event Integration.
func (o GetEventIntegrationResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v GetEventIntegrationResult) string { return v.Arn }).(pulumi.StringOutput)
}

// The description of the Event Integration.
func (o GetEventIntegrationResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetEventIntegrationResult) string { return v.Description }).(pulumi.StringOutput)
}

// A block that defines the configuration information for the event filter. The Event Filter block is documented below.
func (o GetEventIntegrationResultOutput) EventFilters() GetEventIntegrationEventFilterArrayOutput {
	return o.ApplyT(func(v GetEventIntegrationResult) []GetEventIntegrationEventFilter { return v.EventFilters }).(GetEventIntegrationEventFilterArrayOutput)
}

// The EventBridge bus.
func (o GetEventIntegrationResultOutput) EventbridgeBus() pulumi.StringOutput {
	return o.ApplyT(func(v GetEventIntegrationResult) string { return v.EventbridgeBus }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEventIntegrationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEventIntegrationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetEventIntegrationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetEventIntegrationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Metadata that you can assign to help organize the report plans you create.
func (o GetEventIntegrationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetEventIntegrationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEventIntegrationResultOutput{})
}
