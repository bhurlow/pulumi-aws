// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a WAFv2 Web ACL Logging Configuration resource.
//
// > **Note:** To start logging from a WAFv2 Web ACL, an Amazon Kinesis Data Firehose (e.g. `kinesis.FirehoseDeliveryStream` resourc must also be created with a PUT source (not a stream) and in the region that you are operating.
// If you are capturing logs for Amazon CloudFront, always create the firehose in US East (N. Virginia).
// Be sure to give the data firehose a name that starts with the prefix `aws-waf-logs-`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/wafv2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := wafv2.NewWebAclLoggingConfiguration(ctx, "example", &wafv2.WebAclLoggingConfigurationArgs{
// 			LogDestinationConfigs: pulumi.StringArray{
// 				pulumi.Any(aws_kinesis_firehose_delivery_stream.Example.Arn),
// 			},
// 			ResourceArn: pulumi.Any(aws_wafv2_web_acl.Example.Arn),
// 			RedactedFields: wafv2.WebAclLoggingConfigurationRedactedFieldArray{
// 				&wafv2.WebAclLoggingConfigurationRedactedFieldArgs{
// 					SingleHeader: &wafv2.WebAclLoggingConfigurationRedactedFieldSingleHeaderArgs{
// 						Name: pulumi.String("user-agent"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// WAFv2 Web ACL Logging Configurations can be imported using the WAFv2 Web ACL ARN e.g.
//
// ```sh
//  $ pulumi import aws:wafv2/webAclLoggingConfiguration:WebAclLoggingConfiguration example arn:aws:wafv2:us-west-2:123456789012:regional/webacl/test-logs/a1b2c3d4-5678-90ab-cdef
// ```
type WebAclLoggingConfiguration struct {
	pulumi.CustomResourceState

	// The Amazon Kinesis Data Firehose Amazon Resource Name (ARNs) that you want to associate with the web ACL. Currently, only 1 ARN is supported.
	LogDestinationConfigs pulumi.StringArrayOutput `pulumi:"logDestinationConfigs"`
	// The parts of the request that you want to keep out of the logs. Up to 100 `redactedFields` blocks are supported.
	RedactedFields WebAclLoggingConfigurationRedactedFieldArrayOutput `pulumi:"redactedFields"`
	// The Amazon Resource Name (ARN) of the web ACL that you want to associate with `logDestinationConfigs`.
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
}

// NewWebAclLoggingConfiguration registers a new resource with the given unique name, arguments, and options.
func NewWebAclLoggingConfiguration(ctx *pulumi.Context,
	name string, args *WebAclLoggingConfigurationArgs, opts ...pulumi.ResourceOption) (*WebAclLoggingConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LogDestinationConfigs == nil {
		return nil, errors.New("invalid value for required argument 'LogDestinationConfigs'")
	}
	if args.ResourceArn == nil {
		return nil, errors.New("invalid value for required argument 'ResourceArn'")
	}
	var resource WebAclLoggingConfiguration
	err := ctx.RegisterResource("aws:wafv2/webAclLoggingConfiguration:WebAclLoggingConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAclLoggingConfiguration gets an existing WebAclLoggingConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAclLoggingConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAclLoggingConfigurationState, opts ...pulumi.ResourceOption) (*WebAclLoggingConfiguration, error) {
	var resource WebAclLoggingConfiguration
	err := ctx.ReadResource("aws:wafv2/webAclLoggingConfiguration:WebAclLoggingConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAclLoggingConfiguration resources.
type webAclLoggingConfigurationState struct {
	// The Amazon Kinesis Data Firehose Amazon Resource Name (ARNs) that you want to associate with the web ACL. Currently, only 1 ARN is supported.
	LogDestinationConfigs []string `pulumi:"logDestinationConfigs"`
	// The parts of the request that you want to keep out of the logs. Up to 100 `redactedFields` blocks are supported.
	RedactedFields []WebAclLoggingConfigurationRedactedField `pulumi:"redactedFields"`
	// The Amazon Resource Name (ARN) of the web ACL that you want to associate with `logDestinationConfigs`.
	ResourceArn *string `pulumi:"resourceArn"`
}

type WebAclLoggingConfigurationState struct {
	// The Amazon Kinesis Data Firehose Amazon Resource Name (ARNs) that you want to associate with the web ACL. Currently, only 1 ARN is supported.
	LogDestinationConfigs pulumi.StringArrayInput
	// The parts of the request that you want to keep out of the logs. Up to 100 `redactedFields` blocks are supported.
	RedactedFields WebAclLoggingConfigurationRedactedFieldArrayInput
	// The Amazon Resource Name (ARN) of the web ACL that you want to associate with `logDestinationConfigs`.
	ResourceArn pulumi.StringPtrInput
}

func (WebAclLoggingConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclLoggingConfigurationState)(nil)).Elem()
}

type webAclLoggingConfigurationArgs struct {
	// The Amazon Kinesis Data Firehose Amazon Resource Name (ARNs) that you want to associate with the web ACL. Currently, only 1 ARN is supported.
	LogDestinationConfigs []string `pulumi:"logDestinationConfigs"`
	// The parts of the request that you want to keep out of the logs. Up to 100 `redactedFields` blocks are supported.
	RedactedFields []WebAclLoggingConfigurationRedactedField `pulumi:"redactedFields"`
	// The Amazon Resource Name (ARN) of the web ACL that you want to associate with `logDestinationConfigs`.
	ResourceArn string `pulumi:"resourceArn"`
}

// The set of arguments for constructing a WebAclLoggingConfiguration resource.
type WebAclLoggingConfigurationArgs struct {
	// The Amazon Kinesis Data Firehose Amazon Resource Name (ARNs) that you want to associate with the web ACL. Currently, only 1 ARN is supported.
	LogDestinationConfigs pulumi.StringArrayInput
	// The parts of the request that you want to keep out of the logs. Up to 100 `redactedFields` blocks are supported.
	RedactedFields WebAclLoggingConfigurationRedactedFieldArrayInput
	// The Amazon Resource Name (ARN) of the web ACL that you want to associate with `logDestinationConfigs`.
	ResourceArn pulumi.StringInput
}

func (WebAclLoggingConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclLoggingConfigurationArgs)(nil)).Elem()
}

type WebAclLoggingConfigurationInput interface {
	pulumi.Input

	ToWebAclLoggingConfigurationOutput() WebAclLoggingConfigurationOutput
	ToWebAclLoggingConfigurationOutputWithContext(ctx context.Context) WebAclLoggingConfigurationOutput
}

func (*WebAclLoggingConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAclLoggingConfiguration)(nil))
}

func (i *WebAclLoggingConfiguration) ToWebAclLoggingConfigurationOutput() WebAclLoggingConfigurationOutput {
	return i.ToWebAclLoggingConfigurationOutputWithContext(context.Background())
}

func (i *WebAclLoggingConfiguration) ToWebAclLoggingConfigurationOutputWithContext(ctx context.Context) WebAclLoggingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAclLoggingConfigurationOutput)
}

func (i *WebAclLoggingConfiguration) ToWebAclLoggingConfigurationPtrOutput() WebAclLoggingConfigurationPtrOutput {
	return i.ToWebAclLoggingConfigurationPtrOutputWithContext(context.Background())
}

func (i *WebAclLoggingConfiguration) ToWebAclLoggingConfigurationPtrOutputWithContext(ctx context.Context) WebAclLoggingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAclLoggingConfigurationPtrOutput)
}

type WebAclLoggingConfigurationPtrInput interface {
	pulumi.Input

	ToWebAclLoggingConfigurationPtrOutput() WebAclLoggingConfigurationPtrOutput
	ToWebAclLoggingConfigurationPtrOutputWithContext(ctx context.Context) WebAclLoggingConfigurationPtrOutput
}

type webAclLoggingConfigurationPtrType WebAclLoggingConfigurationArgs

func (*webAclLoggingConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAclLoggingConfiguration)(nil))
}

func (i *webAclLoggingConfigurationPtrType) ToWebAclLoggingConfigurationPtrOutput() WebAclLoggingConfigurationPtrOutput {
	return i.ToWebAclLoggingConfigurationPtrOutputWithContext(context.Background())
}

func (i *webAclLoggingConfigurationPtrType) ToWebAclLoggingConfigurationPtrOutputWithContext(ctx context.Context) WebAclLoggingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAclLoggingConfigurationPtrOutput)
}

// WebAclLoggingConfigurationArrayInput is an input type that accepts WebAclLoggingConfigurationArray and WebAclLoggingConfigurationArrayOutput values.
// You can construct a concrete instance of `WebAclLoggingConfigurationArrayInput` via:
//
//          WebAclLoggingConfigurationArray{ WebAclLoggingConfigurationArgs{...} }
type WebAclLoggingConfigurationArrayInput interface {
	pulumi.Input

	ToWebAclLoggingConfigurationArrayOutput() WebAclLoggingConfigurationArrayOutput
	ToWebAclLoggingConfigurationArrayOutputWithContext(context.Context) WebAclLoggingConfigurationArrayOutput
}

type WebAclLoggingConfigurationArray []WebAclLoggingConfigurationInput

func (WebAclLoggingConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*WebAclLoggingConfiguration)(nil))
}

func (i WebAclLoggingConfigurationArray) ToWebAclLoggingConfigurationArrayOutput() WebAclLoggingConfigurationArrayOutput {
	return i.ToWebAclLoggingConfigurationArrayOutputWithContext(context.Background())
}

func (i WebAclLoggingConfigurationArray) ToWebAclLoggingConfigurationArrayOutputWithContext(ctx context.Context) WebAclLoggingConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAclLoggingConfigurationArrayOutput)
}

// WebAclLoggingConfigurationMapInput is an input type that accepts WebAclLoggingConfigurationMap and WebAclLoggingConfigurationMapOutput values.
// You can construct a concrete instance of `WebAclLoggingConfigurationMapInput` via:
//
//          WebAclLoggingConfigurationMap{ "key": WebAclLoggingConfigurationArgs{...} }
type WebAclLoggingConfigurationMapInput interface {
	pulumi.Input

	ToWebAclLoggingConfigurationMapOutput() WebAclLoggingConfigurationMapOutput
	ToWebAclLoggingConfigurationMapOutputWithContext(context.Context) WebAclLoggingConfigurationMapOutput
}

type WebAclLoggingConfigurationMap map[string]WebAclLoggingConfigurationInput

func (WebAclLoggingConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*WebAclLoggingConfiguration)(nil))
}

func (i WebAclLoggingConfigurationMap) ToWebAclLoggingConfigurationMapOutput() WebAclLoggingConfigurationMapOutput {
	return i.ToWebAclLoggingConfigurationMapOutputWithContext(context.Background())
}

func (i WebAclLoggingConfigurationMap) ToWebAclLoggingConfigurationMapOutputWithContext(ctx context.Context) WebAclLoggingConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAclLoggingConfigurationMapOutput)
}

type WebAclLoggingConfigurationOutput struct {
	*pulumi.OutputState
}

func (WebAclLoggingConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAclLoggingConfiguration)(nil))
}

func (o WebAclLoggingConfigurationOutput) ToWebAclLoggingConfigurationOutput() WebAclLoggingConfigurationOutput {
	return o
}

func (o WebAclLoggingConfigurationOutput) ToWebAclLoggingConfigurationOutputWithContext(ctx context.Context) WebAclLoggingConfigurationOutput {
	return o
}

func (o WebAclLoggingConfigurationOutput) ToWebAclLoggingConfigurationPtrOutput() WebAclLoggingConfigurationPtrOutput {
	return o.ToWebAclLoggingConfigurationPtrOutputWithContext(context.Background())
}

func (o WebAclLoggingConfigurationOutput) ToWebAclLoggingConfigurationPtrOutputWithContext(ctx context.Context) WebAclLoggingConfigurationPtrOutput {
	return o.ApplyT(func(v WebAclLoggingConfiguration) *WebAclLoggingConfiguration {
		return &v
	}).(WebAclLoggingConfigurationPtrOutput)
}

type WebAclLoggingConfigurationPtrOutput struct {
	*pulumi.OutputState
}

func (WebAclLoggingConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAclLoggingConfiguration)(nil))
}

func (o WebAclLoggingConfigurationPtrOutput) ToWebAclLoggingConfigurationPtrOutput() WebAclLoggingConfigurationPtrOutput {
	return o
}

func (o WebAclLoggingConfigurationPtrOutput) ToWebAclLoggingConfigurationPtrOutputWithContext(ctx context.Context) WebAclLoggingConfigurationPtrOutput {
	return o
}

type WebAclLoggingConfigurationArrayOutput struct{ *pulumi.OutputState }

func (WebAclLoggingConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebAclLoggingConfiguration)(nil))
}

func (o WebAclLoggingConfigurationArrayOutput) ToWebAclLoggingConfigurationArrayOutput() WebAclLoggingConfigurationArrayOutput {
	return o
}

func (o WebAclLoggingConfigurationArrayOutput) ToWebAclLoggingConfigurationArrayOutputWithContext(ctx context.Context) WebAclLoggingConfigurationArrayOutput {
	return o
}

func (o WebAclLoggingConfigurationArrayOutput) Index(i pulumi.IntInput) WebAclLoggingConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebAclLoggingConfiguration {
		return vs[0].([]WebAclLoggingConfiguration)[vs[1].(int)]
	}).(WebAclLoggingConfigurationOutput)
}

type WebAclLoggingConfigurationMapOutput struct{ *pulumi.OutputState }

func (WebAclLoggingConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WebAclLoggingConfiguration)(nil))
}

func (o WebAclLoggingConfigurationMapOutput) ToWebAclLoggingConfigurationMapOutput() WebAclLoggingConfigurationMapOutput {
	return o
}

func (o WebAclLoggingConfigurationMapOutput) ToWebAclLoggingConfigurationMapOutputWithContext(ctx context.Context) WebAclLoggingConfigurationMapOutput {
	return o
}

func (o WebAclLoggingConfigurationMapOutput) MapIndex(k pulumi.StringInput) WebAclLoggingConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WebAclLoggingConfiguration {
		return vs[0].(map[string]WebAclLoggingConfiguration)[vs[1].(string)]
	}).(WebAclLoggingConfigurationOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAclLoggingConfigurationOutput{})
	pulumi.RegisterOutputType(WebAclLoggingConfigurationPtrOutput{})
	pulumi.RegisterOutputType(WebAclLoggingConfigurationArrayOutput{})
	pulumi.RegisterOutputType(WebAclLoggingConfigurationMapOutput{})
}
