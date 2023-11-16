// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an AWS Config Remediation Configuration.
//
// > **Note:** Config Remediation Configuration requires an existing Config Rule to be present.
//
// ## Example Usage
//
// AWS managed rules can be used by setting the source owner to `AWS` and the source identifier to the name of the managed rule. More information about AWS managed rules can be found in the [AWS Config Developer Guide](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html).
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cfg"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			thisRule, err := cfg.NewRule(ctx, "thisRule", &cfg.RuleArgs{
//				Source: &cfg.RuleSourceArgs{
//					Owner:            pulumi.String("AWS"),
//					SourceIdentifier: pulumi.String("S3_BUCKET_VERSIONING_ENABLED"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cfg.NewRemediationConfiguration(ctx, "thisRemediationConfiguration", &cfg.RemediationConfigurationArgs{
//				ConfigRuleName: thisRule.Name,
//				ResourceType:   pulumi.String("AWS::S3::Bucket"),
//				TargetType:     pulumi.String("SSM_DOCUMENT"),
//				TargetId:       pulumi.String("AWS-EnableS3BucketEncryption"),
//				TargetVersion:  pulumi.String("1"),
//				Parameters: cfg.RemediationConfigurationParameterArray{
//					&cfg.RemediationConfigurationParameterArgs{
//						Name:        pulumi.String("AutomationAssumeRole"),
//						StaticValue: pulumi.String("arn:aws:iam::875924563244:role/security_config"),
//					},
//					&cfg.RemediationConfigurationParameterArgs{
//						Name:          pulumi.String("BucketName"),
//						ResourceValue: pulumi.String("RESOURCE_ID"),
//					},
//					&cfg.RemediationConfigurationParameterArgs{
//						Name:        pulumi.String("SSEAlgorithm"),
//						StaticValue: pulumi.String("AES256"),
//					},
//				},
//				Automatic:                pulumi.Bool(true),
//				MaximumAutomaticAttempts: pulumi.Int(10),
//				RetryAttemptSeconds:      pulumi.Int(600),
//				ExecutionControls: &cfg.RemediationConfigurationExecutionControlsArgs{
//					SsmControls: &cfg.RemediationConfigurationExecutionControlsSsmControlsArgs{
//						ConcurrentExecutionRatePercentage: pulumi.Int(25),
//						ErrorPercentage:                   pulumi.Int(20),
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
// Using `pulumi import`, import Remediation Configurations using the name config_rule_name. For example:
//
// ```sh
//
//	$ pulumi import aws:cfg/remediationConfiguration:RemediationConfiguration this example
//
// ```
type RemediationConfiguration struct {
	pulumi.CustomResourceState

	// ARN of the Config Remediation Configuration.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Remediation is triggered automatically if `true`.
	Automatic pulumi.BoolPtrOutput `pulumi:"automatic"`
	// Name of the AWS Config rule.
	ConfigRuleName pulumi.StringOutput `pulumi:"configRuleName"`
	// Configuration block for execution controls. See below.
	ExecutionControls RemediationConfigurationExecutionControlsPtrOutput `pulumi:"executionControls"`
	// Maximum number of failed attempts for auto-remediation. If you do not select a number, the default is 5.
	MaximumAutomaticAttempts pulumi.IntPtrOutput `pulumi:"maximumAutomaticAttempts"`
	// Can be specified multiple times for each parameter. Each parameter block supports arguments below.
	Parameters RemediationConfigurationParameterArrayOutput `pulumi:"parameters"`
	// Type of resource.
	ResourceType pulumi.StringPtrOutput `pulumi:"resourceType"`
	// Maximum time in seconds that AWS Config runs auto-remediation. If you do not select a number, the default is 60 seconds.
	RetryAttemptSeconds pulumi.IntPtrOutput `pulumi:"retryAttemptSeconds"`
	// Target ID is the name of the public document.
	TargetId pulumi.StringOutput `pulumi:"targetId"`
	// Type of the target. Target executes remediation. For example, SSM document.
	//
	// The following arguments are optional:
	TargetType pulumi.StringOutput `pulumi:"targetType"`
	// Version of the target. For example, version of the SSM document
	TargetVersion pulumi.StringPtrOutput `pulumi:"targetVersion"`
}

// NewRemediationConfiguration registers a new resource with the given unique name, arguments, and options.
func NewRemediationConfiguration(ctx *pulumi.Context,
	name string, args *RemediationConfigurationArgs, opts ...pulumi.ResourceOption) (*RemediationConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigRuleName == nil {
		return nil, errors.New("invalid value for required argument 'ConfigRuleName'")
	}
	if args.TargetId == nil {
		return nil, errors.New("invalid value for required argument 'TargetId'")
	}
	if args.TargetType == nil {
		return nil, errors.New("invalid value for required argument 'TargetType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RemediationConfiguration
	err := ctx.RegisterResource("aws:cfg/remediationConfiguration:RemediationConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRemediationConfiguration gets an existing RemediationConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRemediationConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemediationConfigurationState, opts ...pulumi.ResourceOption) (*RemediationConfiguration, error) {
	var resource RemediationConfiguration
	err := ctx.ReadResource("aws:cfg/remediationConfiguration:RemediationConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RemediationConfiguration resources.
type remediationConfigurationState struct {
	// ARN of the Config Remediation Configuration.
	Arn *string `pulumi:"arn"`
	// Remediation is triggered automatically if `true`.
	Automatic *bool `pulumi:"automatic"`
	// Name of the AWS Config rule.
	ConfigRuleName *string `pulumi:"configRuleName"`
	// Configuration block for execution controls. See below.
	ExecutionControls *RemediationConfigurationExecutionControls `pulumi:"executionControls"`
	// Maximum number of failed attempts for auto-remediation. If you do not select a number, the default is 5.
	MaximumAutomaticAttempts *int `pulumi:"maximumAutomaticAttempts"`
	// Can be specified multiple times for each parameter. Each parameter block supports arguments below.
	Parameters []RemediationConfigurationParameter `pulumi:"parameters"`
	// Type of resource.
	ResourceType *string `pulumi:"resourceType"`
	// Maximum time in seconds that AWS Config runs auto-remediation. If you do not select a number, the default is 60 seconds.
	RetryAttemptSeconds *int `pulumi:"retryAttemptSeconds"`
	// Target ID is the name of the public document.
	TargetId *string `pulumi:"targetId"`
	// Type of the target. Target executes remediation. For example, SSM document.
	//
	// The following arguments are optional:
	TargetType *string `pulumi:"targetType"`
	// Version of the target. For example, version of the SSM document
	TargetVersion *string `pulumi:"targetVersion"`
}

type RemediationConfigurationState struct {
	// ARN of the Config Remediation Configuration.
	Arn pulumi.StringPtrInput
	// Remediation is triggered automatically if `true`.
	Automatic pulumi.BoolPtrInput
	// Name of the AWS Config rule.
	ConfigRuleName pulumi.StringPtrInput
	// Configuration block for execution controls. See below.
	ExecutionControls RemediationConfigurationExecutionControlsPtrInput
	// Maximum number of failed attempts for auto-remediation. If you do not select a number, the default is 5.
	MaximumAutomaticAttempts pulumi.IntPtrInput
	// Can be specified multiple times for each parameter. Each parameter block supports arguments below.
	Parameters RemediationConfigurationParameterArrayInput
	// Type of resource.
	ResourceType pulumi.StringPtrInput
	// Maximum time in seconds that AWS Config runs auto-remediation. If you do not select a number, the default is 60 seconds.
	RetryAttemptSeconds pulumi.IntPtrInput
	// Target ID is the name of the public document.
	TargetId pulumi.StringPtrInput
	// Type of the target. Target executes remediation. For example, SSM document.
	//
	// The following arguments are optional:
	TargetType pulumi.StringPtrInput
	// Version of the target. For example, version of the SSM document
	TargetVersion pulumi.StringPtrInput
}

func (RemediationConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationConfigurationState)(nil)).Elem()
}

type remediationConfigurationArgs struct {
	// Remediation is triggered automatically if `true`.
	Automatic *bool `pulumi:"automatic"`
	// Name of the AWS Config rule.
	ConfigRuleName string `pulumi:"configRuleName"`
	// Configuration block for execution controls. See below.
	ExecutionControls *RemediationConfigurationExecutionControls `pulumi:"executionControls"`
	// Maximum number of failed attempts for auto-remediation. If you do not select a number, the default is 5.
	MaximumAutomaticAttempts *int `pulumi:"maximumAutomaticAttempts"`
	// Can be specified multiple times for each parameter. Each parameter block supports arguments below.
	Parameters []RemediationConfigurationParameter `pulumi:"parameters"`
	// Type of resource.
	ResourceType *string `pulumi:"resourceType"`
	// Maximum time in seconds that AWS Config runs auto-remediation. If you do not select a number, the default is 60 seconds.
	RetryAttemptSeconds *int `pulumi:"retryAttemptSeconds"`
	// Target ID is the name of the public document.
	TargetId string `pulumi:"targetId"`
	// Type of the target. Target executes remediation. For example, SSM document.
	//
	// The following arguments are optional:
	TargetType string `pulumi:"targetType"`
	// Version of the target. For example, version of the SSM document
	TargetVersion *string `pulumi:"targetVersion"`
}

// The set of arguments for constructing a RemediationConfiguration resource.
type RemediationConfigurationArgs struct {
	// Remediation is triggered automatically if `true`.
	Automatic pulumi.BoolPtrInput
	// Name of the AWS Config rule.
	ConfigRuleName pulumi.StringInput
	// Configuration block for execution controls. See below.
	ExecutionControls RemediationConfigurationExecutionControlsPtrInput
	// Maximum number of failed attempts for auto-remediation. If you do not select a number, the default is 5.
	MaximumAutomaticAttempts pulumi.IntPtrInput
	// Can be specified multiple times for each parameter. Each parameter block supports arguments below.
	Parameters RemediationConfigurationParameterArrayInput
	// Type of resource.
	ResourceType pulumi.StringPtrInput
	// Maximum time in seconds that AWS Config runs auto-remediation. If you do not select a number, the default is 60 seconds.
	RetryAttemptSeconds pulumi.IntPtrInput
	// Target ID is the name of the public document.
	TargetId pulumi.StringInput
	// Type of the target. Target executes remediation. For example, SSM document.
	//
	// The following arguments are optional:
	TargetType pulumi.StringInput
	// Version of the target. For example, version of the SSM document
	TargetVersion pulumi.StringPtrInput
}

func (RemediationConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationConfigurationArgs)(nil)).Elem()
}

type RemediationConfigurationInput interface {
	pulumi.Input

	ToRemediationConfigurationOutput() RemediationConfigurationOutput
	ToRemediationConfigurationOutputWithContext(ctx context.Context) RemediationConfigurationOutput
}

func (*RemediationConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationConfiguration)(nil)).Elem()
}

func (i *RemediationConfiguration) ToRemediationConfigurationOutput() RemediationConfigurationOutput {
	return i.ToRemediationConfigurationOutputWithContext(context.Background())
}

func (i *RemediationConfiguration) ToRemediationConfigurationOutputWithContext(ctx context.Context) RemediationConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationConfigurationOutput)
}

// RemediationConfigurationArrayInput is an input type that accepts RemediationConfigurationArray and RemediationConfigurationArrayOutput values.
// You can construct a concrete instance of `RemediationConfigurationArrayInput` via:
//
//	RemediationConfigurationArray{ RemediationConfigurationArgs{...} }
type RemediationConfigurationArrayInput interface {
	pulumi.Input

	ToRemediationConfigurationArrayOutput() RemediationConfigurationArrayOutput
	ToRemediationConfigurationArrayOutputWithContext(context.Context) RemediationConfigurationArrayOutput
}

type RemediationConfigurationArray []RemediationConfigurationInput

func (RemediationConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RemediationConfiguration)(nil)).Elem()
}

func (i RemediationConfigurationArray) ToRemediationConfigurationArrayOutput() RemediationConfigurationArrayOutput {
	return i.ToRemediationConfigurationArrayOutputWithContext(context.Background())
}

func (i RemediationConfigurationArray) ToRemediationConfigurationArrayOutputWithContext(ctx context.Context) RemediationConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationConfigurationArrayOutput)
}

// RemediationConfigurationMapInput is an input type that accepts RemediationConfigurationMap and RemediationConfigurationMapOutput values.
// You can construct a concrete instance of `RemediationConfigurationMapInput` via:
//
//	RemediationConfigurationMap{ "key": RemediationConfigurationArgs{...} }
type RemediationConfigurationMapInput interface {
	pulumi.Input

	ToRemediationConfigurationMapOutput() RemediationConfigurationMapOutput
	ToRemediationConfigurationMapOutputWithContext(context.Context) RemediationConfigurationMapOutput
}

type RemediationConfigurationMap map[string]RemediationConfigurationInput

func (RemediationConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RemediationConfiguration)(nil)).Elem()
}

func (i RemediationConfigurationMap) ToRemediationConfigurationMapOutput() RemediationConfigurationMapOutput {
	return i.ToRemediationConfigurationMapOutputWithContext(context.Background())
}

func (i RemediationConfigurationMap) ToRemediationConfigurationMapOutputWithContext(ctx context.Context) RemediationConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemediationConfigurationMapOutput)
}

type RemediationConfigurationOutput struct{ *pulumi.OutputState }

func (RemediationConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemediationConfiguration)(nil)).Elem()
}

func (o RemediationConfigurationOutput) ToRemediationConfigurationOutput() RemediationConfigurationOutput {
	return o
}

func (o RemediationConfigurationOutput) ToRemediationConfigurationOutputWithContext(ctx context.Context) RemediationConfigurationOutput {
	return o
}

// ARN of the Config Remediation Configuration.
func (o RemediationConfigurationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationConfiguration) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Remediation is triggered automatically if `true`.
func (o RemediationConfigurationOutput) Automatic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RemediationConfiguration) pulumi.BoolPtrOutput { return v.Automatic }).(pulumi.BoolPtrOutput)
}

// Name of the AWS Config rule.
func (o RemediationConfigurationOutput) ConfigRuleName() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationConfiguration) pulumi.StringOutput { return v.ConfigRuleName }).(pulumi.StringOutput)
}

// Configuration block for execution controls. See below.
func (o RemediationConfigurationOutput) ExecutionControls() RemediationConfigurationExecutionControlsPtrOutput {
	return o.ApplyT(func(v *RemediationConfiguration) RemediationConfigurationExecutionControlsPtrOutput {
		return v.ExecutionControls
	}).(RemediationConfigurationExecutionControlsPtrOutput)
}

// Maximum number of failed attempts for auto-remediation. If you do not select a number, the default is 5.
func (o RemediationConfigurationOutput) MaximumAutomaticAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RemediationConfiguration) pulumi.IntPtrOutput { return v.MaximumAutomaticAttempts }).(pulumi.IntPtrOutput)
}

// Can be specified multiple times for each parameter. Each parameter block supports arguments below.
func (o RemediationConfigurationOutput) Parameters() RemediationConfigurationParameterArrayOutput {
	return o.ApplyT(func(v *RemediationConfiguration) RemediationConfigurationParameterArrayOutput { return v.Parameters }).(RemediationConfigurationParameterArrayOutput)
}

// Type of resource.
func (o RemediationConfigurationOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemediationConfiguration) pulumi.StringPtrOutput { return v.ResourceType }).(pulumi.StringPtrOutput)
}

// Maximum time in seconds that AWS Config runs auto-remediation. If you do not select a number, the default is 60 seconds.
func (o RemediationConfigurationOutput) RetryAttemptSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RemediationConfiguration) pulumi.IntPtrOutput { return v.RetryAttemptSeconds }).(pulumi.IntPtrOutput)
}

// Target ID is the name of the public document.
func (o RemediationConfigurationOutput) TargetId() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationConfiguration) pulumi.StringOutput { return v.TargetId }).(pulumi.StringOutput)
}

// Type of the target. Target executes remediation. For example, SSM document.
//
// The following arguments are optional:
func (o RemediationConfigurationOutput) TargetType() pulumi.StringOutput {
	return o.ApplyT(func(v *RemediationConfiguration) pulumi.StringOutput { return v.TargetType }).(pulumi.StringOutput)
}

// Version of the target. For example, version of the SSM document
func (o RemediationConfigurationOutput) TargetVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemediationConfiguration) pulumi.StringPtrOutput { return v.TargetVersion }).(pulumi.StringPtrOutput)
}

type RemediationConfigurationArrayOutput struct{ *pulumi.OutputState }

func (RemediationConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RemediationConfiguration)(nil)).Elem()
}

func (o RemediationConfigurationArrayOutput) ToRemediationConfigurationArrayOutput() RemediationConfigurationArrayOutput {
	return o
}

func (o RemediationConfigurationArrayOutput) ToRemediationConfigurationArrayOutputWithContext(ctx context.Context) RemediationConfigurationArrayOutput {
	return o
}

func (o RemediationConfigurationArrayOutput) Index(i pulumi.IntInput) RemediationConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RemediationConfiguration {
		return vs[0].([]*RemediationConfiguration)[vs[1].(int)]
	}).(RemediationConfigurationOutput)
}

type RemediationConfigurationMapOutput struct{ *pulumi.OutputState }

func (RemediationConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RemediationConfiguration)(nil)).Elem()
}

func (o RemediationConfigurationMapOutput) ToRemediationConfigurationMapOutput() RemediationConfigurationMapOutput {
	return o
}

func (o RemediationConfigurationMapOutput) ToRemediationConfigurationMapOutputWithContext(ctx context.Context) RemediationConfigurationMapOutput {
	return o
}

func (o RemediationConfigurationMapOutput) MapIndex(k pulumi.StringInput) RemediationConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RemediationConfiguration {
		return vs[0].(map[string]*RemediationConfiguration)[vs[1].(string)]
	}).(RemediationConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RemediationConfigurationInput)(nil)).Elem(), &RemediationConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*RemediationConfigurationArrayInput)(nil)).Elem(), RemediationConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RemediationConfigurationMapInput)(nil)).Elem(), RemediationConfigurationMap{})
	pulumi.RegisterOutputType(RemediationConfigurationOutput{})
	pulumi.RegisterOutputType(RemediationConfigurationArrayOutput{})
	pulumi.RegisterOutputType(RemediationConfigurationMapOutput{})
}
