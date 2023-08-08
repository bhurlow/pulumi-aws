// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rbin

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for managing an AWS RBin Rule.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rbin"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rbin.NewRule(ctx, "example", &rbin.RuleArgs{
//				Description: pulumi.String("example_rule"),
//				ResourceTags: rbin.RuleResourceTagArray{
//					&rbin.RuleResourceTagArgs{
//						ResourceTagKey:   pulumi.String("tag_key"),
//						ResourceTagValue: pulumi.String("tag_value"),
//					},
//				},
//				ResourceType: pulumi.String("EBS_SNAPSHOT"),
//				RetentionPeriod: &rbin.RuleRetentionPeriodArgs{
//					RetentionPeriodUnit:  pulumi.String("DAYS"),
//					RetentionPeriodValue: pulumi.Int(10),
//				},
//				Tags: pulumi.StringMap{
//					"test_tag_key": pulumi.String("test_tag_value"),
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
//	to = aws_rbin_rule.example
//
//	id = "examplerule" } Using `pulumi import`, import RBin Rule using the `id`. For exampleconsole % pulumi import aws_rbin_rule.example examplerule
type Rule struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// The retention rule description.
	Description pulumi.StringOutput `pulumi:"description"`
	// Information about the retention rule lock configuration. See `lockConfiguration` below.
	LockConfiguration RuleLockConfigurationPtrOutput `pulumi:"lockConfiguration"`
	// (Timestamp) The date and time at which the unlock delay is set to expire. Only returned for retention rules that have been unlocked and that are still within the unlock delay period.
	LockEndTime pulumi.StringOutput `pulumi:"lockEndTime"`
	// (Optional) The lock state of the retention rules to list. Only retention rules with the specified lock state are returned. Valid values are `locked`, `pendingUnlock`, `unlocked`.
	LockState pulumi.StringOutput `pulumi:"lockState"`
	// Specifies the resource tags to use to identify resources that are to be retained by a tag-level retention rule. See `resourceTags` below.
	ResourceTags RuleResourceTagArrayOutput `pulumi:"resourceTags"`
	// The resource type to be retained by the retention rule. Valid values are `EBS_SNAPSHOT` and `EC2_IMAGE`.
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	// Information about the retention period for which the retention rule is to retain resources. See `retentionPeriod` below.
	//
	// The following arguments are optional:
	RetentionPeriod RuleRetentionPeriodOutput `pulumi:"retentionPeriod"`
	// (String) The state of the retention rule. Only retention rules that are in the `available` state retain resources. Valid values include `pending` and `available`.
	Status  pulumi.StringOutput    `pulumi:"status"`
	Tags    pulumi.StringMapOutput `pulumi:"tags"`
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewRule registers a new resource with the given unique name, arguments, and options.
func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOption) (*Rule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	if args.RetentionPeriod == nil {
		return nil, errors.New("invalid value for required argument 'RetentionPeriod'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Rule
	err := ctx.RegisterResource("aws:rbin/rule:Rule", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:rbin/rule:Rule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rule resources.
type ruleState struct {
	Arn *string `pulumi:"arn"`
	// The retention rule description.
	Description *string `pulumi:"description"`
	// Information about the retention rule lock configuration. See `lockConfiguration` below.
	LockConfiguration *RuleLockConfiguration `pulumi:"lockConfiguration"`
	// (Timestamp) The date and time at which the unlock delay is set to expire. Only returned for retention rules that have been unlocked and that are still within the unlock delay period.
	LockEndTime *string `pulumi:"lockEndTime"`
	// (Optional) The lock state of the retention rules to list. Only retention rules with the specified lock state are returned. Valid values are `locked`, `pendingUnlock`, `unlocked`.
	LockState *string `pulumi:"lockState"`
	// Specifies the resource tags to use to identify resources that are to be retained by a tag-level retention rule. See `resourceTags` below.
	ResourceTags []RuleResourceTag `pulumi:"resourceTags"`
	// The resource type to be retained by the retention rule. Valid values are `EBS_SNAPSHOT` and `EC2_IMAGE`.
	ResourceType *string `pulumi:"resourceType"`
	// Information about the retention period for which the retention rule is to retain resources. See `retentionPeriod` below.
	//
	// The following arguments are optional:
	RetentionPeriod *RuleRetentionPeriod `pulumi:"retentionPeriod"`
	// (String) The state of the retention rule. Only retention rules that are in the `available` state retain resources. Valid values include `pending` and `available`.
	Status  *string           `pulumi:"status"`
	Tags    map[string]string `pulumi:"tags"`
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type RuleState struct {
	Arn pulumi.StringPtrInput
	// The retention rule description.
	Description pulumi.StringPtrInput
	// Information about the retention rule lock configuration. See `lockConfiguration` below.
	LockConfiguration RuleLockConfigurationPtrInput
	// (Timestamp) The date and time at which the unlock delay is set to expire. Only returned for retention rules that have been unlocked and that are still within the unlock delay period.
	LockEndTime pulumi.StringPtrInput
	// (Optional) The lock state of the retention rules to list. Only retention rules with the specified lock state are returned. Valid values are `locked`, `pendingUnlock`, `unlocked`.
	LockState pulumi.StringPtrInput
	// Specifies the resource tags to use to identify resources that are to be retained by a tag-level retention rule. See `resourceTags` below.
	ResourceTags RuleResourceTagArrayInput
	// The resource type to be retained by the retention rule. Valid values are `EBS_SNAPSHOT` and `EC2_IMAGE`.
	ResourceType pulumi.StringPtrInput
	// Information about the retention period for which the retention rule is to retain resources. See `retentionPeriod` below.
	//
	// The following arguments are optional:
	RetentionPeriod RuleRetentionPeriodPtrInput
	// (String) The state of the retention rule. Only retention rules that are in the `available` state retain resources. Valid values include `pending` and `available`.
	Status  pulumi.StringPtrInput
	Tags    pulumi.StringMapInput
	TagsAll pulumi.StringMapInput
}

func (RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleState)(nil)).Elem()
}

type ruleArgs struct {
	// The retention rule description.
	Description *string `pulumi:"description"`
	// Information about the retention rule lock configuration. See `lockConfiguration` below.
	LockConfiguration *RuleLockConfiguration `pulumi:"lockConfiguration"`
	// Specifies the resource tags to use to identify resources that are to be retained by a tag-level retention rule. See `resourceTags` below.
	ResourceTags []RuleResourceTag `pulumi:"resourceTags"`
	// The resource type to be retained by the retention rule. Valid values are `EBS_SNAPSHOT` and `EC2_IMAGE`.
	ResourceType string `pulumi:"resourceType"`
	// Information about the retention period for which the retention rule is to retain resources. See `retentionPeriod` below.
	//
	// The following arguments are optional:
	RetentionPeriod RuleRetentionPeriod `pulumi:"retentionPeriod"`
	Tags            map[string]string   `pulumi:"tags"`
}

// The set of arguments for constructing a Rule resource.
type RuleArgs struct {
	// The retention rule description.
	Description pulumi.StringPtrInput
	// Information about the retention rule lock configuration. See `lockConfiguration` below.
	LockConfiguration RuleLockConfigurationPtrInput
	// Specifies the resource tags to use to identify resources that are to be retained by a tag-level retention rule. See `resourceTags` below.
	ResourceTags RuleResourceTagArrayInput
	// The resource type to be retained by the retention rule. Valid values are `EBS_SNAPSHOT` and `EC2_IMAGE`.
	ResourceType pulumi.StringInput
	// Information about the retention period for which the retention rule is to retain resources. See `retentionPeriod` below.
	//
	// The following arguments are optional:
	RetentionPeriod RuleRetentionPeriodInput
	Tags            pulumi.StringMapInput
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

func (o RuleOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The retention rule description.
func (o RuleOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Information about the retention rule lock configuration. See `lockConfiguration` below.
func (o RuleOutput) LockConfiguration() RuleLockConfigurationPtrOutput {
	return o.ApplyT(func(v *Rule) RuleLockConfigurationPtrOutput { return v.LockConfiguration }).(RuleLockConfigurationPtrOutput)
}

// (Timestamp) The date and time at which the unlock delay is set to expire. Only returned for retention rules that have been unlocked and that are still within the unlock delay period.
func (o RuleOutput) LockEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.LockEndTime }).(pulumi.StringOutput)
}

// (Optional) The lock state of the retention rules to list. Only retention rules with the specified lock state are returned. Valid values are `locked`, `pendingUnlock`, `unlocked`.
func (o RuleOutput) LockState() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.LockState }).(pulumi.StringOutput)
}

// Specifies the resource tags to use to identify resources that are to be retained by a tag-level retention rule. See `resourceTags` below.
func (o RuleOutput) ResourceTags() RuleResourceTagArrayOutput {
	return o.ApplyT(func(v *Rule) RuleResourceTagArrayOutput { return v.ResourceTags }).(RuleResourceTagArrayOutput)
}

// The resource type to be retained by the retention rule. Valid values are `EBS_SNAPSHOT` and `EC2_IMAGE`.
func (o RuleOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

// Information about the retention period for which the retention rule is to retain resources. See `retentionPeriod` below.
//
// The following arguments are optional:
func (o RuleOutput) RetentionPeriod() RuleRetentionPeriodOutput {
	return o.ApplyT(func(v *Rule) RuleRetentionPeriodOutput { return v.RetentionPeriod }).(RuleRetentionPeriodOutput)
}

// (String) The state of the retention rule. Only retention rules that are in the `available` state retain resources. Valid values include `pending` and `available`.
func (o RuleOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o RuleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

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
