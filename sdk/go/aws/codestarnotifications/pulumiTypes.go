// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codestarnotifications

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type NotificationRuleTarget struct {
	// The ARN of notification rule target. For example, a SNS Topic ARN.
	Address string `pulumi:"address"`
	// The status of the notification rule. Possible values are `ENABLED` and `DISABLED`, default is `ENABLED`.
	Status *string `pulumi:"status"`
	// The type of the notification target. Default value is `SNS`.
	Type *string `pulumi:"type"`
}

// NotificationRuleTargetInput is an input type that accepts NotificationRuleTargetArgs and NotificationRuleTargetOutput values.
// You can construct a concrete instance of `NotificationRuleTargetInput` via:
//
//	NotificationRuleTargetArgs{...}
type NotificationRuleTargetInput interface {
	pulumi.Input

	ToNotificationRuleTargetOutput() NotificationRuleTargetOutput
	ToNotificationRuleTargetOutputWithContext(context.Context) NotificationRuleTargetOutput
}

type NotificationRuleTargetArgs struct {
	// The ARN of notification rule target. For example, a SNS Topic ARN.
	Address pulumi.StringInput `pulumi:"address"`
	// The status of the notification rule. Possible values are `ENABLED` and `DISABLED`, default is `ENABLED`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The type of the notification target. Default value is `SNS`.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (NotificationRuleTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationRuleTarget)(nil)).Elem()
}

func (i NotificationRuleTargetArgs) ToNotificationRuleTargetOutput() NotificationRuleTargetOutput {
	return i.ToNotificationRuleTargetOutputWithContext(context.Background())
}

func (i NotificationRuleTargetArgs) ToNotificationRuleTargetOutputWithContext(ctx context.Context) NotificationRuleTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationRuleTargetOutput)
}

// NotificationRuleTargetArrayInput is an input type that accepts NotificationRuleTargetArray and NotificationRuleTargetArrayOutput values.
// You can construct a concrete instance of `NotificationRuleTargetArrayInput` via:
//
//	NotificationRuleTargetArray{ NotificationRuleTargetArgs{...} }
type NotificationRuleTargetArrayInput interface {
	pulumi.Input

	ToNotificationRuleTargetArrayOutput() NotificationRuleTargetArrayOutput
	ToNotificationRuleTargetArrayOutputWithContext(context.Context) NotificationRuleTargetArrayOutput
}

type NotificationRuleTargetArray []NotificationRuleTargetInput

func (NotificationRuleTargetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NotificationRuleTarget)(nil)).Elem()
}

func (i NotificationRuleTargetArray) ToNotificationRuleTargetArrayOutput() NotificationRuleTargetArrayOutput {
	return i.ToNotificationRuleTargetArrayOutputWithContext(context.Background())
}

func (i NotificationRuleTargetArray) ToNotificationRuleTargetArrayOutputWithContext(ctx context.Context) NotificationRuleTargetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationRuleTargetArrayOutput)
}

type NotificationRuleTargetOutput struct{ *pulumi.OutputState }

func (NotificationRuleTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationRuleTarget)(nil)).Elem()
}

func (o NotificationRuleTargetOutput) ToNotificationRuleTargetOutput() NotificationRuleTargetOutput {
	return o
}

func (o NotificationRuleTargetOutput) ToNotificationRuleTargetOutputWithContext(ctx context.Context) NotificationRuleTargetOutput {
	return o
}

// The ARN of notification rule target. For example, a SNS Topic ARN.
func (o NotificationRuleTargetOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v NotificationRuleTarget) string { return v.Address }).(pulumi.StringOutput)
}

// The status of the notification rule. Possible values are `ENABLED` and `DISABLED`, default is `ENABLED`.
func (o NotificationRuleTargetOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationRuleTarget) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// The type of the notification target. Default value is `SNS`.
func (o NotificationRuleTargetOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationRuleTarget) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type NotificationRuleTargetArrayOutput struct{ *pulumi.OutputState }

func (NotificationRuleTargetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NotificationRuleTarget)(nil)).Elem()
}

func (o NotificationRuleTargetArrayOutput) ToNotificationRuleTargetArrayOutput() NotificationRuleTargetArrayOutput {
	return o
}

func (o NotificationRuleTargetArrayOutput) ToNotificationRuleTargetArrayOutputWithContext(ctx context.Context) NotificationRuleTargetArrayOutput {
	return o
}

func (o NotificationRuleTargetArrayOutput) Index(i pulumi.IntInput) NotificationRuleTargetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NotificationRuleTarget {
		return vs[0].([]NotificationRuleTarget)[vs[1].(int)]
	}).(NotificationRuleTargetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationRuleTargetInput)(nil)).Elem(), NotificationRuleTargetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationRuleTargetArrayInput)(nil)).Elem(), NotificationRuleTargetArray{})
	pulumi.RegisterOutputType(NotificationRuleTargetOutput{})
	pulumi.RegisterOutputType(NotificationRuleTargetArrayOutput{})
}
