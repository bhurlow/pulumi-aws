// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource creates a WAFv2 Web ACL Logging Configuration.
//
// !> **WARNING:** When logging from a WAFv2 Web ACL to a CloudWatch Log Group, the WAFv2 service tries to create or update a generic Log Resource Policy named `AWSWAF-LOGS`. However, if there are a large number of Web ACLs or if the account frequently creates and deletes Web ACLs, this policy may exceed the maximum policy size. As a result, this resource type will fail to be created. More details about this issue can be found in this issue. To prevent this issue, you can manage a specific resource policy. Please refer to the example below for managing a CloudWatch Log Group with a managed CloudWatch Log Resource Policy.
//
// ## Example Usage
// ### With Redacted Fields
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/wafv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := wafv2.NewWebAclLoggingConfiguration(ctx, "example", &wafv2.WebAclLoggingConfigurationArgs{
//				LogDestinationConfigs: pulumi.StringArray{
//					aws_kinesis_firehose_delivery_stream.Example.Arn,
//				},
//				ResourceArn: pulumi.Any(aws_wafv2_web_acl.Example.Arn),
//				RedactedFields: wafv2.WebAclLoggingConfigurationRedactedFieldArray{
//					&wafv2.WebAclLoggingConfigurationRedactedFieldArgs{
//						SingleHeader: &wafv2.WebAclLoggingConfigurationRedactedFieldSingleHeaderArgs{
//							Name: pulumi.String("user-agent"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### With Logging Filter
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/wafv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := wafv2.NewWebAclLoggingConfiguration(ctx, "example", &wafv2.WebAclLoggingConfigurationArgs{
//				LogDestinationConfigs: pulumi.StringArray{
//					aws_kinesis_firehose_delivery_stream.Example.Arn,
//				},
//				ResourceArn: pulumi.Any(aws_wafv2_web_acl.Example.Arn),
//				LoggingFilter: &wafv2.WebAclLoggingConfigurationLoggingFilterArgs{
//					DefaultBehavior: pulumi.String("KEEP"),
//					Filters: wafv2.WebAclLoggingConfigurationLoggingFilterFilterArray{
//						&wafv2.WebAclLoggingConfigurationLoggingFilterFilterArgs{
//							Behavior: pulumi.String("DROP"),
//							Conditions: wafv2.WebAclLoggingConfigurationLoggingFilterFilterConditionArray{
//								&wafv2.WebAclLoggingConfigurationLoggingFilterFilterConditionArgs{
//									ActionCondition: &wafv2.WebAclLoggingConfigurationLoggingFilterFilterConditionActionConditionArgs{
//										Action: pulumi.String("COUNT"),
//									},
//								},
//								&wafv2.WebAclLoggingConfigurationLoggingFilterFilterConditionArgs{
//									LabelNameCondition: &wafv2.WebAclLoggingConfigurationLoggingFilterFilterConditionLabelNameConditionArgs{
//										LabelName: pulumi.String("awswaf:111122223333:rulegroup:testRules:LabelNameZ"),
//									},
//								},
//							},
//							Requirement: pulumi.String("MEETS_ALL"),
//						},
//						&wafv2.WebAclLoggingConfigurationLoggingFilterFilterArgs{
//							Behavior: pulumi.String("KEEP"),
//							Conditions: wafv2.WebAclLoggingConfigurationLoggingFilterFilterConditionArray{
//								&wafv2.WebAclLoggingConfigurationLoggingFilterFilterConditionArgs{
//									ActionCondition: &wafv2.WebAclLoggingConfigurationLoggingFilterFilterConditionActionConditionArgs{
//										Action: pulumi.String("ALLOW"),
//									},
//								},
//							},
//							Requirement: pulumi.String("MEETS_ANY"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// terraform import {
//
//	to = aws_wafv2_web_acl_logging_configuration.example
//
//	id = "arn:aws:wafv2:us-west-2:123456789012:regional/webacl/test-logs/a1b2c3d4-5678-90ab-cdef" } Using `pulumi import`, import WAFv2 Web ACL Logging Configurations using the ARN of the WAFv2 Web ACL. For exampleconsole % pulumi import aws_wafv2_web_acl_logging_configuration.example arn:aws:wafv2:us-west-2:123456789012:regional/webacl/test-logs/a1b2c3d4-5678-90ab-cdef
type WebAclLoggingConfiguration struct {
	pulumi.CustomResourceState

	// Configuration block that allows you to associate Amazon Kinesis Data Firehose, Cloudwatch Log log group, or S3 bucket Amazon Resource Names (ARNs) with the web ACL. **Note:** data firehose, log group, or bucket name **must** be prefixed with `aws-waf-logs-`, e.g. `aws-waf-logs-example-firehose`, `aws-waf-logs-example-log-group`, or `aws-waf-logs-example-bucket`.
	LogDestinationConfigs pulumi.StringArrayOutput `pulumi:"logDestinationConfigs"`
	// Configuration block that specifies which web requests are kept in the logs and which are dropped. It allows filtering based on the rule action and the web request labels applied by matching rules during web ACL evaluation. For more details, refer to the Logging Filter section below.
	LoggingFilter WebAclLoggingConfigurationLoggingFilterPtrOutput `pulumi:"loggingFilter"`
	// Configuration for parts of the request that you want to keep out of the logs. Up to 100 `redactedFields` blocks are supported. See Redacted Fields below for more details.
	RedactedFields WebAclLoggingConfigurationRedactedFieldArrayOutput `pulumi:"redactedFields"`
	// Amazon Resource Name (ARN) of the web ACL that you want to associate with `logDestinationConfigs`.
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
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// Configuration block that allows you to associate Amazon Kinesis Data Firehose, Cloudwatch Log log group, or S3 bucket Amazon Resource Names (ARNs) with the web ACL. **Note:** data firehose, log group, or bucket name **must** be prefixed with `aws-waf-logs-`, e.g. `aws-waf-logs-example-firehose`, `aws-waf-logs-example-log-group`, or `aws-waf-logs-example-bucket`.
	LogDestinationConfigs []string `pulumi:"logDestinationConfigs"`
	// Configuration block that specifies which web requests are kept in the logs and which are dropped. It allows filtering based on the rule action and the web request labels applied by matching rules during web ACL evaluation. For more details, refer to the Logging Filter section below.
	LoggingFilter *WebAclLoggingConfigurationLoggingFilter `pulumi:"loggingFilter"`
	// Configuration for parts of the request that you want to keep out of the logs. Up to 100 `redactedFields` blocks are supported. See Redacted Fields below for more details.
	RedactedFields []WebAclLoggingConfigurationRedactedField `pulumi:"redactedFields"`
	// Amazon Resource Name (ARN) of the web ACL that you want to associate with `logDestinationConfigs`.
	ResourceArn *string `pulumi:"resourceArn"`
}

type WebAclLoggingConfigurationState struct {
	// Configuration block that allows you to associate Amazon Kinesis Data Firehose, Cloudwatch Log log group, or S3 bucket Amazon Resource Names (ARNs) with the web ACL. **Note:** data firehose, log group, or bucket name **must** be prefixed with `aws-waf-logs-`, e.g. `aws-waf-logs-example-firehose`, `aws-waf-logs-example-log-group`, or `aws-waf-logs-example-bucket`.
	LogDestinationConfigs pulumi.StringArrayInput
	// Configuration block that specifies which web requests are kept in the logs and which are dropped. It allows filtering based on the rule action and the web request labels applied by matching rules during web ACL evaluation. For more details, refer to the Logging Filter section below.
	LoggingFilter WebAclLoggingConfigurationLoggingFilterPtrInput
	// Configuration for parts of the request that you want to keep out of the logs. Up to 100 `redactedFields` blocks are supported. See Redacted Fields below for more details.
	RedactedFields WebAclLoggingConfigurationRedactedFieldArrayInput
	// Amazon Resource Name (ARN) of the web ACL that you want to associate with `logDestinationConfigs`.
	ResourceArn pulumi.StringPtrInput
}

func (WebAclLoggingConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAclLoggingConfigurationState)(nil)).Elem()
}

type webAclLoggingConfigurationArgs struct {
	// Configuration block that allows you to associate Amazon Kinesis Data Firehose, Cloudwatch Log log group, or S3 bucket Amazon Resource Names (ARNs) with the web ACL. **Note:** data firehose, log group, or bucket name **must** be prefixed with `aws-waf-logs-`, e.g. `aws-waf-logs-example-firehose`, `aws-waf-logs-example-log-group`, or `aws-waf-logs-example-bucket`.
	LogDestinationConfigs []string `pulumi:"logDestinationConfigs"`
	// Configuration block that specifies which web requests are kept in the logs and which are dropped. It allows filtering based on the rule action and the web request labels applied by matching rules during web ACL evaluation. For more details, refer to the Logging Filter section below.
	LoggingFilter *WebAclLoggingConfigurationLoggingFilter `pulumi:"loggingFilter"`
	// Configuration for parts of the request that you want to keep out of the logs. Up to 100 `redactedFields` blocks are supported. See Redacted Fields below for more details.
	RedactedFields []WebAclLoggingConfigurationRedactedField `pulumi:"redactedFields"`
	// Amazon Resource Name (ARN) of the web ACL that you want to associate with `logDestinationConfigs`.
	ResourceArn string `pulumi:"resourceArn"`
}

// The set of arguments for constructing a WebAclLoggingConfiguration resource.
type WebAclLoggingConfigurationArgs struct {
	// Configuration block that allows you to associate Amazon Kinesis Data Firehose, Cloudwatch Log log group, or S3 bucket Amazon Resource Names (ARNs) with the web ACL. **Note:** data firehose, log group, or bucket name **must** be prefixed with `aws-waf-logs-`, e.g. `aws-waf-logs-example-firehose`, `aws-waf-logs-example-log-group`, or `aws-waf-logs-example-bucket`.
	LogDestinationConfigs pulumi.StringArrayInput
	// Configuration block that specifies which web requests are kept in the logs and which are dropped. It allows filtering based on the rule action and the web request labels applied by matching rules during web ACL evaluation. For more details, refer to the Logging Filter section below.
	LoggingFilter WebAclLoggingConfigurationLoggingFilterPtrInput
	// Configuration for parts of the request that you want to keep out of the logs. Up to 100 `redactedFields` blocks are supported. See Redacted Fields below for more details.
	RedactedFields WebAclLoggingConfigurationRedactedFieldArrayInput
	// Amazon Resource Name (ARN) of the web ACL that you want to associate with `logDestinationConfigs`.
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
	return reflect.TypeOf((**WebAclLoggingConfiguration)(nil)).Elem()
}

func (i *WebAclLoggingConfiguration) ToWebAclLoggingConfigurationOutput() WebAclLoggingConfigurationOutput {
	return i.ToWebAclLoggingConfigurationOutputWithContext(context.Background())
}

func (i *WebAclLoggingConfiguration) ToWebAclLoggingConfigurationOutputWithContext(ctx context.Context) WebAclLoggingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAclLoggingConfigurationOutput)
}

// WebAclLoggingConfigurationArrayInput is an input type that accepts WebAclLoggingConfigurationArray and WebAclLoggingConfigurationArrayOutput values.
// You can construct a concrete instance of `WebAclLoggingConfigurationArrayInput` via:
//
//	WebAclLoggingConfigurationArray{ WebAclLoggingConfigurationArgs{...} }
type WebAclLoggingConfigurationArrayInput interface {
	pulumi.Input

	ToWebAclLoggingConfigurationArrayOutput() WebAclLoggingConfigurationArrayOutput
	ToWebAclLoggingConfigurationArrayOutputWithContext(context.Context) WebAclLoggingConfigurationArrayOutput
}

type WebAclLoggingConfigurationArray []WebAclLoggingConfigurationInput

func (WebAclLoggingConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebAclLoggingConfiguration)(nil)).Elem()
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
//	WebAclLoggingConfigurationMap{ "key": WebAclLoggingConfigurationArgs{...} }
type WebAclLoggingConfigurationMapInput interface {
	pulumi.Input

	ToWebAclLoggingConfigurationMapOutput() WebAclLoggingConfigurationMapOutput
	ToWebAclLoggingConfigurationMapOutputWithContext(context.Context) WebAclLoggingConfigurationMapOutput
}

type WebAclLoggingConfigurationMap map[string]WebAclLoggingConfigurationInput

func (WebAclLoggingConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebAclLoggingConfiguration)(nil)).Elem()
}

func (i WebAclLoggingConfigurationMap) ToWebAclLoggingConfigurationMapOutput() WebAclLoggingConfigurationMapOutput {
	return i.ToWebAclLoggingConfigurationMapOutputWithContext(context.Background())
}

func (i WebAclLoggingConfigurationMap) ToWebAclLoggingConfigurationMapOutputWithContext(ctx context.Context) WebAclLoggingConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAclLoggingConfigurationMapOutput)
}

type WebAclLoggingConfigurationOutput struct{ *pulumi.OutputState }

func (WebAclLoggingConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAclLoggingConfiguration)(nil)).Elem()
}

func (o WebAclLoggingConfigurationOutput) ToWebAclLoggingConfigurationOutput() WebAclLoggingConfigurationOutput {
	return o
}

func (o WebAclLoggingConfigurationOutput) ToWebAclLoggingConfigurationOutputWithContext(ctx context.Context) WebAclLoggingConfigurationOutput {
	return o
}

// Configuration block that allows you to associate Amazon Kinesis Data Firehose, Cloudwatch Log log group, or S3 bucket Amazon Resource Names (ARNs) with the web ACL. **Note:** data firehose, log group, or bucket name **must** be prefixed with `aws-waf-logs-`, e.g. `aws-waf-logs-example-firehose`, `aws-waf-logs-example-log-group`, or `aws-waf-logs-example-bucket`.
func (o WebAclLoggingConfigurationOutput) LogDestinationConfigs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAclLoggingConfiguration) pulumi.StringArrayOutput { return v.LogDestinationConfigs }).(pulumi.StringArrayOutput)
}

// Configuration block that specifies which web requests are kept in the logs and which are dropped. It allows filtering based on the rule action and the web request labels applied by matching rules during web ACL evaluation. For more details, refer to the Logging Filter section below.
func (o WebAclLoggingConfigurationOutput) LoggingFilter() WebAclLoggingConfigurationLoggingFilterPtrOutput {
	return o.ApplyT(func(v *WebAclLoggingConfiguration) WebAclLoggingConfigurationLoggingFilterPtrOutput {
		return v.LoggingFilter
	}).(WebAclLoggingConfigurationLoggingFilterPtrOutput)
}

// Configuration for parts of the request that you want to keep out of the logs. Up to 100 `redactedFields` blocks are supported. See Redacted Fields below for more details.
func (o WebAclLoggingConfigurationOutput) RedactedFields() WebAclLoggingConfigurationRedactedFieldArrayOutput {
	return o.ApplyT(func(v *WebAclLoggingConfiguration) WebAclLoggingConfigurationRedactedFieldArrayOutput {
		return v.RedactedFields
	}).(WebAclLoggingConfigurationRedactedFieldArrayOutput)
}

// Amazon Resource Name (ARN) of the web ACL that you want to associate with `logDestinationConfigs`.
func (o WebAclLoggingConfigurationOutput) ResourceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAclLoggingConfiguration) pulumi.StringOutput { return v.ResourceArn }).(pulumi.StringOutput)
}

type WebAclLoggingConfigurationArrayOutput struct{ *pulumi.OutputState }

func (WebAclLoggingConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebAclLoggingConfiguration)(nil)).Elem()
}

func (o WebAclLoggingConfigurationArrayOutput) ToWebAclLoggingConfigurationArrayOutput() WebAclLoggingConfigurationArrayOutput {
	return o
}

func (o WebAclLoggingConfigurationArrayOutput) ToWebAclLoggingConfigurationArrayOutputWithContext(ctx context.Context) WebAclLoggingConfigurationArrayOutput {
	return o
}

func (o WebAclLoggingConfigurationArrayOutput) Index(i pulumi.IntInput) WebAclLoggingConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebAclLoggingConfiguration {
		return vs[0].([]*WebAclLoggingConfiguration)[vs[1].(int)]
	}).(WebAclLoggingConfigurationOutput)
}

type WebAclLoggingConfigurationMapOutput struct{ *pulumi.OutputState }

func (WebAclLoggingConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebAclLoggingConfiguration)(nil)).Elem()
}

func (o WebAclLoggingConfigurationMapOutput) ToWebAclLoggingConfigurationMapOutput() WebAclLoggingConfigurationMapOutput {
	return o
}

func (o WebAclLoggingConfigurationMapOutput) ToWebAclLoggingConfigurationMapOutputWithContext(ctx context.Context) WebAclLoggingConfigurationMapOutput {
	return o
}

func (o WebAclLoggingConfigurationMapOutput) MapIndex(k pulumi.StringInput) WebAclLoggingConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebAclLoggingConfiguration {
		return vs[0].(map[string]*WebAclLoggingConfiguration)[vs[1].(string)]
	}).(WebAclLoggingConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebAclLoggingConfigurationInput)(nil)).Elem(), &WebAclLoggingConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebAclLoggingConfigurationArrayInput)(nil)).Elem(), WebAclLoggingConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebAclLoggingConfigurationMapInput)(nil)).Elem(), WebAclLoggingConfigurationMap{})
	pulumi.RegisterOutputType(WebAclLoggingConfigurationOutput{})
	pulumi.RegisterOutputType(WebAclLoggingConfigurationArrayOutput{})
	pulumi.RegisterOutputType(WebAclLoggingConfigurationMapOutput{})
}
