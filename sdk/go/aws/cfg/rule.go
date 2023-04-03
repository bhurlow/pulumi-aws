// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an AWS Config Rule.
//
// > **Note:** Config Rule requires an existing Configuration Recorder to be present. Use of `dependsOn` is recommended (as shown below) to avoid race conditions.
//
// ## Example Usage
// ### AWS Managed Rules
//
// AWS managed rules can be used by setting the source owner to `AWS` and the source identifier to the name of the managed rule. More information about AWS managed rules can be found in the [AWS Config Developer Guide](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html).
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cfg"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
//				Statements: []iam.GetPolicyDocumentStatement{
//					{
//						Effect: pulumi.StringRef("Allow"),
//						Principals: []iam.GetPolicyDocumentStatementPrincipal{
//							{
//								Type: "Service",
//								Identifiers: []string{
//									"config.amazonaws.com",
//								},
//							},
//						},
//						Actions: []string{
//							"sts:AssumeRole",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			role, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
//				AssumeRolePolicy: *pulumi.String(assumeRole.Json),
//			})
//			if err != nil {
//				return err
//			}
//			foo, err := cfg.NewRecorder(ctx, "foo", &cfg.RecorderArgs{
//				RoleArn: role.Arn,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cfg.NewRule(ctx, "rule", &cfg.RuleArgs{
//				Source: &cfg.RuleSourceArgs{
//					Owner:            pulumi.String("AWS"),
//					SourceIdentifier: pulumi.String("S3_BUCKET_VERSIONING_ENABLED"),
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				foo,
//			}))
//			if err != nil {
//				return err
//			}
//			policyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
//				Statements: []iam.GetPolicyDocumentStatement{
//					{
//						Effect: pulumi.StringRef("Allow"),
//						Actions: []string{
//							"config:Put*",
//						},
//						Resources: []string{
//							"*",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = iam.NewRolePolicy(ctx, "rolePolicy", &iam.RolePolicyArgs{
//				Role:   role.ID(),
//				Policy: *pulumi.String(policyDocument.Json),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Custom Rules
//
// Custom rules can be used by setting the source owner to `CUSTOM_LAMBDA` and the source identifier to the Amazon Resource Name (ARN) of the Lambda Function. The AWS Config service must have permissions to invoke the Lambda Function, e.g., via the `lambda.Permission` resource. More information about custom rules can be found in the [AWS Config Developer Guide](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_develop-rules.html).
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cfg"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/lambda"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleRecorder, err := cfg.NewRecorder(ctx, "exampleRecorder", nil)
//			if err != nil {
//				return err
//			}
//			exampleFunction, err := lambda.NewFunction(ctx, "exampleFunction", nil)
//			if err != nil {
//				return err
//			}
//			examplePermission, err := lambda.NewPermission(ctx, "examplePermission", &lambda.PermissionArgs{
//				Action:    pulumi.String("lambda:InvokeFunction"),
//				Function:  exampleFunction.Arn,
//				Principal: pulumi.String("config.amazonaws.com"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cfg.NewRule(ctx, "exampleRule", &cfg.RuleArgs{
//				Source: &cfg.RuleSourceArgs{
//					Owner:            pulumi.String("CUSTOM_LAMBDA"),
//					SourceIdentifier: exampleFunction.Arn,
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				exampleRecorder,
//				examplePermission,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Custom Policies
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cfg"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cfg.NewRule(ctx, "example", &cfg.RuleArgs{
//				Source: &cfg.RuleSourceArgs{
//					Owner: pulumi.String("CUSTOM_POLICY"),
//					SourceDetails: cfg.RuleSourceSourceDetailArray{
//						&cfg.RuleSourceSourceDetailArgs{
//							MessageType: pulumi.String("ConfigurationItemChangeNotification"),
//						},
//					},
//					CustomPolicyDetails: &cfg.RuleSourceCustomPolicyDetailsArgs{
//						PolicyRuntime: pulumi.String("guard-2.x.x"),
//						PolicyText:    pulumi.String("	  rule tableisactive when\n		  resourceType == \"AWS::DynamoDB::Table\" {\n		  configuration.tableStatus == ['ACTIVE']\n	  }\n	  \n	  rule checkcompliance when\n		  resourceType == \"AWS::DynamoDB::Table\"\n		  tableisactive {\n			  supplementaryConfiguration.ContinuousBackupsDescription.pointInTimeRecoveryDescription.pointInTimeRecoveryStatus == \"ENABLED\"\n	  }\n"),
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
// Config Rule can be imported using the name, e.g.,
//
// ```sh
//
//	$ pulumi import aws:cfg/rule:Rule foo example
//
// ```
type Rule struct {
	pulumi.CustomResourceState

	// The ARN of the config rule
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of the rule
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A string in JSON format that is passed to the AWS Config rule Lambda function.
	InputParameters pulumi.StringPtrOutput `pulumi:"inputParameters"`
	// The maximum frequency with which AWS Config runs evaluations for a rule.
	MaximumExecutionFrequency pulumi.StringPtrOutput `pulumi:"maximumExecutionFrequency"`
	// The name of the rule
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the config rule
	RuleId pulumi.StringOutput `pulumi:"ruleId"`
	// Scope defines which resources can trigger an evaluation for the rule. See Source Below.
	Scope RuleScopePtrOutput `pulumi:"scope"`
	// Source specifies the rule owner, the rule identifier, and the notifications that cause the function to evaluate your AWS resources. See Scope Below.
	Source RuleSourceOutput `pulumi:"source"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewRule registers a new resource with the given unique name, arguments, and options.
func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOption) (*Rule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	var resource Rule
	err := ctx.RegisterResource("aws:cfg/rule:Rule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRule gets an existing Rule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleState, opts ...pulumi.ResourceOption) (*Rule, error) {
	var resource Rule
	err := ctx.ReadResource("aws:cfg/rule:Rule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rule resources.
type ruleState struct {
	// The ARN of the config rule
	Arn *string `pulumi:"arn"`
	// Description of the rule
	Description *string `pulumi:"description"`
	// A string in JSON format that is passed to the AWS Config rule Lambda function.
	InputParameters *string `pulumi:"inputParameters"`
	// The maximum frequency with which AWS Config runs evaluations for a rule.
	MaximumExecutionFrequency *string `pulumi:"maximumExecutionFrequency"`
	// The name of the rule
	Name *string `pulumi:"name"`
	// The ID of the config rule
	RuleId *string `pulumi:"ruleId"`
	// Scope defines which resources can trigger an evaluation for the rule. See Source Below.
	Scope *RuleScope `pulumi:"scope"`
	// Source specifies the rule owner, the rule identifier, and the notifications that cause the function to evaluate your AWS resources. See Scope Below.
	Source *RuleSource `pulumi:"source"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type RuleState struct {
	// The ARN of the config rule
	Arn pulumi.StringPtrInput
	// Description of the rule
	Description pulumi.StringPtrInput
	// A string in JSON format that is passed to the AWS Config rule Lambda function.
	InputParameters pulumi.StringPtrInput
	// The maximum frequency with which AWS Config runs evaluations for a rule.
	MaximumExecutionFrequency pulumi.StringPtrInput
	// The name of the rule
	Name pulumi.StringPtrInput
	// The ID of the config rule
	RuleId pulumi.StringPtrInput
	// Scope defines which resources can trigger an evaluation for the rule. See Source Below.
	Scope RuleScopePtrInput
	// Source specifies the rule owner, the rule identifier, and the notifications that cause the function to evaluate your AWS resources. See Scope Below.
	Source RuleSourcePtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleState)(nil)).Elem()
}

type ruleArgs struct {
	// Description of the rule
	Description *string `pulumi:"description"`
	// A string in JSON format that is passed to the AWS Config rule Lambda function.
	InputParameters *string `pulumi:"inputParameters"`
	// The maximum frequency with which AWS Config runs evaluations for a rule.
	MaximumExecutionFrequency *string `pulumi:"maximumExecutionFrequency"`
	// The name of the rule
	Name *string `pulumi:"name"`
	// Scope defines which resources can trigger an evaluation for the rule. See Source Below.
	Scope *RuleScope `pulumi:"scope"`
	// Source specifies the rule owner, the rule identifier, and the notifications that cause the function to evaluate your AWS resources. See Scope Below.
	Source RuleSource `pulumi:"source"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Rule resource.
type RuleArgs struct {
	// Description of the rule
	Description pulumi.StringPtrInput
	// A string in JSON format that is passed to the AWS Config rule Lambda function.
	InputParameters pulumi.StringPtrInput
	// The maximum frequency with which AWS Config runs evaluations for a rule.
	MaximumExecutionFrequency pulumi.StringPtrInput
	// The name of the rule
	Name pulumi.StringPtrInput
	// Scope defines which resources can trigger an evaluation for the rule. See Source Below.
	Scope RuleScopePtrInput
	// Source specifies the rule owner, the rule identifier, and the notifications that cause the function to evaluate your AWS resources. See Scope Below.
	Source RuleSourceInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (RuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleArgs)(nil)).Elem()
}

type RuleInput interface {
	pulumi.Input

	ToRuleOutput() RuleOutput
	ToRuleOutputWithContext(ctx context.Context) RuleOutput
}

func (*Rule) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil)).Elem()
}

func (i *Rule) ToRuleOutput() RuleOutput {
	return i.ToRuleOutputWithContext(context.Background())
}

func (i *Rule) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleOutput)
}

// RuleArrayInput is an input type that accepts RuleArray and RuleArrayOutput values.
// You can construct a concrete instance of `RuleArrayInput` via:
//
//	RuleArray{ RuleArgs{...} }
type RuleArrayInput interface {
	pulumi.Input

	ToRuleArrayOutput() RuleArrayOutput
	ToRuleArrayOutputWithContext(context.Context) RuleArrayOutput
}

type RuleArray []RuleInput

func (RuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rule)(nil)).Elem()
}

func (i RuleArray) ToRuleArrayOutput() RuleArrayOutput {
	return i.ToRuleArrayOutputWithContext(context.Background())
}

func (i RuleArray) ToRuleArrayOutputWithContext(ctx context.Context) RuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleArrayOutput)
}

// RuleMapInput is an input type that accepts RuleMap and RuleMapOutput values.
// You can construct a concrete instance of `RuleMapInput` via:
//
//	RuleMap{ "key": RuleArgs{...} }
type RuleMapInput interface {
	pulumi.Input

	ToRuleMapOutput() RuleMapOutput
	ToRuleMapOutputWithContext(context.Context) RuleMapOutput
}

type RuleMap map[string]RuleInput

func (RuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rule)(nil)).Elem()
}

func (i RuleMap) ToRuleMapOutput() RuleMapOutput {
	return i.ToRuleMapOutputWithContext(context.Background())
}

func (i RuleMap) ToRuleMapOutputWithContext(ctx context.Context) RuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleMapOutput)
}

type RuleOutput struct{ *pulumi.OutputState }

func (RuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil)).Elem()
}

func (o RuleOutput) ToRuleOutput() RuleOutput {
	return o
}

func (o RuleOutput) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return o
}

// The ARN of the config rule
func (o RuleOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Description of the rule
func (o RuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// A string in JSON format that is passed to the AWS Config rule Lambda function.
func (o RuleOutput) InputParameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.InputParameters }).(pulumi.StringPtrOutput)
}

// The maximum frequency with which AWS Config runs evaluations for a rule.
func (o RuleOutput) MaximumExecutionFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.MaximumExecutionFrequency }).(pulumi.StringPtrOutput)
}

// The name of the rule
func (o RuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the config rule
func (o RuleOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.RuleId }).(pulumi.StringOutput)
}

// Scope defines which resources can trigger an evaluation for the rule. See Source Below.
func (o RuleOutput) Scope() RuleScopePtrOutput {
	return o.ApplyT(func(v *Rule) RuleScopePtrOutput { return v.Scope }).(RuleScopePtrOutput)
}

// Source specifies the rule owner, the rule identifier, and the notifications that cause the function to evaluate your AWS resources. See Scope Below.
func (o RuleOutput) Source() RuleSourceOutput {
	return o.ApplyT(func(v *Rule) RuleSourceOutput { return v.Source }).(RuleSourceOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o RuleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o RuleOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type RuleArrayOutput struct{ *pulumi.OutputState }

func (RuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rule)(nil)).Elem()
}

func (o RuleArrayOutput) ToRuleArrayOutput() RuleArrayOutput {
	return o
}

func (o RuleArrayOutput) ToRuleArrayOutputWithContext(ctx context.Context) RuleArrayOutput {
	return o
}

func (o RuleArrayOutput) Index(i pulumi.IntInput) RuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Rule {
		return vs[0].([]*Rule)[vs[1].(int)]
	}).(RuleOutput)
}

type RuleMapOutput struct{ *pulumi.OutputState }

func (RuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rule)(nil)).Elem()
}

func (o RuleMapOutput) ToRuleMapOutput() RuleMapOutput {
	return o
}

func (o RuleMapOutput) ToRuleMapOutputWithContext(ctx context.Context) RuleMapOutput {
	return o
}

func (o RuleMapOutput) MapIndex(k pulumi.StringInput) RuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Rule {
		return vs[0].(map[string]*Rule)[vs[1].(string)]
	}).(RuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuleInput)(nil)).Elem(), &Rule{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleArrayInput)(nil)).Elem(), RuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleMapInput)(nil)).Elem(), RuleMap{})
	pulumi.RegisterOutputType(RuleOutput{})
	pulumi.RegisterOutputType(RuleArrayOutput{})
	pulumi.RegisterOutputType(RuleMapOutput{})
}
