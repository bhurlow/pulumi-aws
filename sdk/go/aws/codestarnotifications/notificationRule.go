// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codestarnotifications

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CodeStar Notifications Rule.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codecommit"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codestarnotifications"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// code, err := codecommit.NewRepository(ctx, "code", &codecommit.RepositoryArgs{
// RepositoryName: pulumi.String("example-code-repo"),
// })
// if err != nil {
// return err
// }
// notif, err := sns.NewTopic(ctx, "notif", nil)
// if err != nil {
// return err
// }
// notifAccess := notif.Arn.ApplyT(func(arn string) (iam.GetPolicyDocumentResult, error) {
// return iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
// Statements: []iam.GetPolicyDocumentStatement{
// {
// Actions: []string{
// "sns:Publish",
// },
// Principals: []iam.GetPolicyDocumentStatementPrincipal{
// {
// Type: "Service",
// Identifiers: []string{
// "codestar-notifications.amazonaws.com",
// },
// },
// },
// Resources: interface{}{
// arn,
// },
// },
// },
// }, nil), nil
// }).(iam.GetPolicyDocumentResultOutput)
// _, err = sns.NewTopicPolicy(ctx, "default", &sns.TopicPolicyArgs{
// Arn: notif.Arn,
// Policy: notifAccess.ApplyT(func(notifAccess iam.GetPolicyDocumentResult) (*string, error) {
// return &notifAccess.Json, nil
// }).(pulumi.StringPtrOutput),
// })
// if err != nil {
// return err
// }
// _, err = codestarnotifications.NewNotificationRule(ctx, "commits", &codestarnotifications.NotificationRuleArgs{
// DetailType: pulumi.String("BASIC"),
// EventTypeIds: pulumi.StringArray{
// pulumi.String("codecommit-repository-comments-on-commits"),
// },
// Resource: code.Arn,
// Targets: codestarnotifications.NotificationRuleTargetArray{
// &codestarnotifications.NotificationRuleTargetArgs{
// Address: notif.Arn,
// },
// },
// })
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
//
// ## Import
//
// terraform import {
//
//	to = aws_codestarnotifications_notification_rule.foo
//
//	id = "arn:aws:codestar-notifications:us-west-1:0123456789:notificationrule/2cdc68a3-8f7c-4893-b6a5-45b362bd4f2b" } Using `pulumi import`, import CodeStar notification rule using the ARN. For exampleconsole % pulumi import aws_codestarnotifications_notification_rule.foo arn:aws:codestar-notifications:us-west-1:0123456789:notificationrule/2cdc68a3-8f7c-4893-b6a5-45b362bd4f2b
type NotificationRule struct {
	pulumi.CustomResourceState

	// The codestar notification rule ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The level of detail to include in the notifications for this resource. Possible values are `BASIC` and `FULL`.
	DetailType pulumi.StringOutput `pulumi:"detailType"`
	// A list of event types associated with this notification rule.
	// For list of allowed events see [here](https://docs.aws.amazon.com/codestar-notifications/latest/userguide/concepts.html#concepts-api).
	EventTypeIds pulumi.StringArrayOutput `pulumi:"eventTypeIds"`
	// The name of notification rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ARN of the resource to associate with the notification rule.
	Resource pulumi.StringOutput `pulumi:"resource"`
	// The status of the notification rule. Possible values are `ENABLED` and `DISABLED`, default is `ENABLED`.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Configuration blocks containing notification target information. Can be specified multiple times. At least one target must be specified on creation.
	Targets NotificationRuleTargetArrayOutput `pulumi:"targets"`
}

// NewNotificationRule registers a new resource with the given unique name, arguments, and options.
func NewNotificationRule(ctx *pulumi.Context,
	name string, args *NotificationRuleArgs, opts ...pulumi.ResourceOption) (*NotificationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DetailType == nil {
		return nil, errors.New("invalid value for required argument 'DetailType'")
	}
	if args.EventTypeIds == nil {
		return nil, errors.New("invalid value for required argument 'EventTypeIds'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NotificationRule
	err := ctx.RegisterResource("aws:codestarnotifications/notificationRule:NotificationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotificationRule gets an existing NotificationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotificationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationRuleState, opts ...pulumi.ResourceOption) (*NotificationRule, error) {
	var resource NotificationRule
	err := ctx.ReadResource("aws:codestarnotifications/notificationRule:NotificationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotificationRule resources.
type notificationRuleState struct {
	// The codestar notification rule ARN.
	Arn *string `pulumi:"arn"`
	// The level of detail to include in the notifications for this resource. Possible values are `BASIC` and `FULL`.
	DetailType *string `pulumi:"detailType"`
	// A list of event types associated with this notification rule.
	// For list of allowed events see [here](https://docs.aws.amazon.com/codestar-notifications/latest/userguide/concepts.html#concepts-api).
	EventTypeIds []string `pulumi:"eventTypeIds"`
	// The name of notification rule.
	Name *string `pulumi:"name"`
	// The ARN of the resource to associate with the notification rule.
	Resource *string `pulumi:"resource"`
	// The status of the notification rule. Possible values are `ENABLED` and `DISABLED`, default is `ENABLED`.
	Status *string `pulumi:"status"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Configuration blocks containing notification target information. Can be specified multiple times. At least one target must be specified on creation.
	Targets []NotificationRuleTarget `pulumi:"targets"`
}

type NotificationRuleState struct {
	// The codestar notification rule ARN.
	Arn pulumi.StringPtrInput
	// The level of detail to include in the notifications for this resource. Possible values are `BASIC` and `FULL`.
	DetailType pulumi.StringPtrInput
	// A list of event types associated with this notification rule.
	// For list of allowed events see [here](https://docs.aws.amazon.com/codestar-notifications/latest/userguide/concepts.html#concepts-api).
	EventTypeIds pulumi.StringArrayInput
	// The name of notification rule.
	Name pulumi.StringPtrInput
	// The ARN of the resource to associate with the notification rule.
	Resource pulumi.StringPtrInput
	// The status of the notification rule. Possible values are `ENABLED` and `DISABLED`, default is `ENABLED`.
	Status pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// Configuration blocks containing notification target information. Can be specified multiple times. At least one target must be specified on creation.
	Targets NotificationRuleTargetArrayInput
}

func (NotificationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationRuleState)(nil)).Elem()
}

type notificationRuleArgs struct {
	// The level of detail to include in the notifications for this resource. Possible values are `BASIC` and `FULL`.
	DetailType string `pulumi:"detailType"`
	// A list of event types associated with this notification rule.
	// For list of allowed events see [here](https://docs.aws.amazon.com/codestar-notifications/latest/userguide/concepts.html#concepts-api).
	EventTypeIds []string `pulumi:"eventTypeIds"`
	// The name of notification rule.
	Name *string `pulumi:"name"`
	// The ARN of the resource to associate with the notification rule.
	Resource string `pulumi:"resource"`
	// The status of the notification rule. Possible values are `ENABLED` and `DISABLED`, default is `ENABLED`.
	Status *string `pulumi:"status"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Configuration blocks containing notification target information. Can be specified multiple times. At least one target must be specified on creation.
	Targets []NotificationRuleTarget `pulumi:"targets"`
}

// The set of arguments for constructing a NotificationRule resource.
type NotificationRuleArgs struct {
	// The level of detail to include in the notifications for this resource. Possible values are `BASIC` and `FULL`.
	DetailType pulumi.StringInput
	// A list of event types associated with this notification rule.
	// For list of allowed events see [here](https://docs.aws.amazon.com/codestar-notifications/latest/userguide/concepts.html#concepts-api).
	EventTypeIds pulumi.StringArrayInput
	// The name of notification rule.
	Name pulumi.StringPtrInput
	// The ARN of the resource to associate with the notification rule.
	Resource pulumi.StringInput
	// The status of the notification rule. Possible values are `ENABLED` and `DISABLED`, default is `ENABLED`.
	Status pulumi.StringPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Configuration blocks containing notification target information. Can be specified multiple times. At least one target must be specified on creation.
	Targets NotificationRuleTargetArrayInput
}

func (NotificationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationRuleArgs)(nil)).Elem()
}

type NotificationRuleInput interface {
	pulumi.Input

	ToNotificationRuleOutput() NotificationRuleOutput
	ToNotificationRuleOutputWithContext(ctx context.Context) NotificationRuleOutput
}

func (*NotificationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationRule)(nil)).Elem()
}

func (i *NotificationRule) ToNotificationRuleOutput() NotificationRuleOutput {
	return i.ToNotificationRuleOutputWithContext(context.Background())
}

func (i *NotificationRule) ToNotificationRuleOutputWithContext(ctx context.Context) NotificationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationRuleOutput)
}

// NotificationRuleArrayInput is an input type that accepts NotificationRuleArray and NotificationRuleArrayOutput values.
// You can construct a concrete instance of `NotificationRuleArrayInput` via:
//
//	NotificationRuleArray{ NotificationRuleArgs{...} }
type NotificationRuleArrayInput interface {
	pulumi.Input

	ToNotificationRuleArrayOutput() NotificationRuleArrayOutput
	ToNotificationRuleArrayOutputWithContext(context.Context) NotificationRuleArrayOutput
}

type NotificationRuleArray []NotificationRuleInput

func (NotificationRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NotificationRule)(nil)).Elem()
}

func (i NotificationRuleArray) ToNotificationRuleArrayOutput() NotificationRuleArrayOutput {
	return i.ToNotificationRuleArrayOutputWithContext(context.Background())
}

func (i NotificationRuleArray) ToNotificationRuleArrayOutputWithContext(ctx context.Context) NotificationRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationRuleArrayOutput)
}

// NotificationRuleMapInput is an input type that accepts NotificationRuleMap and NotificationRuleMapOutput values.
// You can construct a concrete instance of `NotificationRuleMapInput` via:
//
//	NotificationRuleMap{ "key": NotificationRuleArgs{...} }
type NotificationRuleMapInput interface {
	pulumi.Input

	ToNotificationRuleMapOutput() NotificationRuleMapOutput
	ToNotificationRuleMapOutputWithContext(context.Context) NotificationRuleMapOutput
}

type NotificationRuleMap map[string]NotificationRuleInput

func (NotificationRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NotificationRule)(nil)).Elem()
}

func (i NotificationRuleMap) ToNotificationRuleMapOutput() NotificationRuleMapOutput {
	return i.ToNotificationRuleMapOutputWithContext(context.Background())
}

func (i NotificationRuleMap) ToNotificationRuleMapOutputWithContext(ctx context.Context) NotificationRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationRuleMapOutput)
}

type NotificationRuleOutput struct{ *pulumi.OutputState }

func (NotificationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationRule)(nil)).Elem()
}

func (o NotificationRuleOutput) ToNotificationRuleOutput() NotificationRuleOutput {
	return o
}

func (o NotificationRuleOutput) ToNotificationRuleOutputWithContext(ctx context.Context) NotificationRuleOutput {
	return o
}

// The codestar notification rule ARN.
func (o NotificationRuleOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationRule) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The level of detail to include in the notifications for this resource. Possible values are `BASIC` and `FULL`.
func (o NotificationRuleOutput) DetailType() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationRule) pulumi.StringOutput { return v.DetailType }).(pulumi.StringOutput)
}

// A list of event types associated with this notification rule.
// For list of allowed events see [here](https://docs.aws.amazon.com/codestar-notifications/latest/userguide/concepts.html#concepts-api).
func (o NotificationRuleOutput) EventTypeIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NotificationRule) pulumi.StringArrayOutput { return v.EventTypeIds }).(pulumi.StringArrayOutput)
}

// The name of notification rule.
func (o NotificationRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ARN of the resource to associate with the notification rule.
func (o NotificationRuleOutput) Resource() pulumi.StringOutput {
	return o.ApplyT(func(v *NotificationRule) pulumi.StringOutput { return v.Resource }).(pulumi.StringOutput)
}

// The status of the notification rule. Possible values are `ENABLED` and `DISABLED`, default is `ENABLED`.
func (o NotificationRuleOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotificationRule) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o NotificationRuleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NotificationRule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o NotificationRuleOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NotificationRule) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Configuration blocks containing notification target information. Can be specified multiple times. At least one target must be specified on creation.
func (o NotificationRuleOutput) Targets() NotificationRuleTargetArrayOutput {
	return o.ApplyT(func(v *NotificationRule) NotificationRuleTargetArrayOutput { return v.Targets }).(NotificationRuleTargetArrayOutput)
}

type NotificationRuleArrayOutput struct{ *pulumi.OutputState }

func (NotificationRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NotificationRule)(nil)).Elem()
}

func (o NotificationRuleArrayOutput) ToNotificationRuleArrayOutput() NotificationRuleArrayOutput {
	return o
}

func (o NotificationRuleArrayOutput) ToNotificationRuleArrayOutputWithContext(ctx context.Context) NotificationRuleArrayOutput {
	return o
}

func (o NotificationRuleArrayOutput) Index(i pulumi.IntInput) NotificationRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NotificationRule {
		return vs[0].([]*NotificationRule)[vs[1].(int)]
	}).(NotificationRuleOutput)
}

type NotificationRuleMapOutput struct{ *pulumi.OutputState }

func (NotificationRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NotificationRule)(nil)).Elem()
}

func (o NotificationRuleMapOutput) ToNotificationRuleMapOutput() NotificationRuleMapOutput {
	return o
}

func (o NotificationRuleMapOutput) ToNotificationRuleMapOutputWithContext(ctx context.Context) NotificationRuleMapOutput {
	return o
}

func (o NotificationRuleMapOutput) MapIndex(k pulumi.StringInput) NotificationRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NotificationRule {
		return vs[0].(map[string]*NotificationRule)[vs[1].(string)]
	}).(NotificationRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationRuleInput)(nil)).Elem(), &NotificationRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationRuleArrayInput)(nil)).Elem(), NotificationRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationRuleMapInput)(nil)).Elem(), NotificationRuleMap{})
	pulumi.RegisterOutputType(NotificationRuleOutput{})
	pulumi.RegisterOutputType(NotificationRuleArrayOutput{})
	pulumi.RegisterOutputType(NotificationRuleMapOutput{})
}
