// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package secretsmanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type SecretReplica struct {
	// ARN, Key ID, or Alias of the AWS KMS key within the region secret is replicated to. If one is not specified, then Secrets Manager defaults to using the AWS account's default KMS key (`aws/secretsmanager`) in the region or creates one for use if non-existent.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Date that you last accessed the secret in the Region.
	LastAccessedDate *string `pulumi:"lastAccessedDate"`
	// Region for replicating the secret.
	Region string `pulumi:"region"`
	// Status can be `InProgress`, `Failed`, or `InSync`.
	Status *string `pulumi:"status"`
	// Message such as `Replication succeeded` or `Secret with this name already exists in this region`.
	StatusMessage *string `pulumi:"statusMessage"`
}

// SecretReplicaInput is an input type that accepts SecretReplicaArgs and SecretReplicaOutput values.
// You can construct a concrete instance of `SecretReplicaInput` via:
//
//	SecretReplicaArgs{...}
type SecretReplicaInput interface {
	pulumi.Input

	ToSecretReplicaOutput() SecretReplicaOutput
	ToSecretReplicaOutputWithContext(context.Context) SecretReplicaOutput
}

type SecretReplicaArgs struct {
	// ARN, Key ID, or Alias of the AWS KMS key within the region secret is replicated to. If one is not specified, then Secrets Manager defaults to using the AWS account's default KMS key (`aws/secretsmanager`) in the region or creates one for use if non-existent.
	KmsKeyId pulumi.StringPtrInput `pulumi:"kmsKeyId"`
	// Date that you last accessed the secret in the Region.
	LastAccessedDate pulumi.StringPtrInput `pulumi:"lastAccessedDate"`
	// Region for replicating the secret.
	Region pulumi.StringInput `pulumi:"region"`
	// Status can be `InProgress`, `Failed`, or `InSync`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// Message such as `Replication succeeded` or `Secret with this name already exists in this region`.
	StatusMessage pulumi.StringPtrInput `pulumi:"statusMessage"`
}

func (SecretReplicaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretReplica)(nil)).Elem()
}

func (i SecretReplicaArgs) ToSecretReplicaOutput() SecretReplicaOutput {
	return i.ToSecretReplicaOutputWithContext(context.Background())
}

func (i SecretReplicaArgs) ToSecretReplicaOutputWithContext(ctx context.Context) SecretReplicaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretReplicaOutput)
}

// SecretReplicaArrayInput is an input type that accepts SecretReplicaArray and SecretReplicaArrayOutput values.
// You can construct a concrete instance of `SecretReplicaArrayInput` via:
//
//	SecretReplicaArray{ SecretReplicaArgs{...} }
type SecretReplicaArrayInput interface {
	pulumi.Input

	ToSecretReplicaArrayOutput() SecretReplicaArrayOutput
	ToSecretReplicaArrayOutputWithContext(context.Context) SecretReplicaArrayOutput
}

type SecretReplicaArray []SecretReplicaInput

func (SecretReplicaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretReplica)(nil)).Elem()
}

func (i SecretReplicaArray) ToSecretReplicaArrayOutput() SecretReplicaArrayOutput {
	return i.ToSecretReplicaArrayOutputWithContext(context.Background())
}

func (i SecretReplicaArray) ToSecretReplicaArrayOutputWithContext(ctx context.Context) SecretReplicaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretReplicaArrayOutput)
}

type SecretReplicaOutput struct{ *pulumi.OutputState }

func (SecretReplicaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretReplica)(nil)).Elem()
}

func (o SecretReplicaOutput) ToSecretReplicaOutput() SecretReplicaOutput {
	return o
}

func (o SecretReplicaOutput) ToSecretReplicaOutputWithContext(ctx context.Context) SecretReplicaOutput {
	return o
}

// ARN, Key ID, or Alias of the AWS KMS key within the region secret is replicated to. If one is not specified, then Secrets Manager defaults to using the AWS account's default KMS key (`aws/secretsmanager`) in the region or creates one for use if non-existent.
func (o SecretReplicaOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretReplica) *string { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// Date that you last accessed the secret in the Region.
func (o SecretReplicaOutput) LastAccessedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretReplica) *string { return v.LastAccessedDate }).(pulumi.StringPtrOutput)
}

// Region for replicating the secret.
func (o SecretReplicaOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v SecretReplica) string { return v.Region }).(pulumi.StringOutput)
}

// Status can be `InProgress`, `Failed`, or `InSync`.
func (o SecretReplicaOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretReplica) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// Message such as `Replication succeeded` or `Secret with this name already exists in this region`.
func (o SecretReplicaOutput) StatusMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretReplica) *string { return v.StatusMessage }).(pulumi.StringPtrOutput)
}

type SecretReplicaArrayOutput struct{ *pulumi.OutputState }

func (SecretReplicaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretReplica)(nil)).Elem()
}

func (o SecretReplicaArrayOutput) ToSecretReplicaArrayOutput() SecretReplicaArrayOutput {
	return o
}

func (o SecretReplicaArrayOutput) ToSecretReplicaArrayOutputWithContext(ctx context.Context) SecretReplicaArrayOutput {
	return o
}

func (o SecretReplicaArrayOutput) Index(i pulumi.IntInput) SecretReplicaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecretReplica {
		return vs[0].([]SecretReplica)[vs[1].(int)]
	}).(SecretReplicaOutput)
}

type SecretRotationRotationRules struct {
	// Specifies the number of days between automatic scheduled rotations of the secret. Either `automaticallyAfterDays` or `scheduleExpression` must be specified.
	AutomaticallyAfterDays *int `pulumi:"automaticallyAfterDays"`
	// The length of the rotation window in hours. For example, `3h` for a three hour window.
	Duration *string `pulumi:"duration"`
	// A `cron()` or `rate()` expression that defines the schedule for rotating your secret. Either `automaticallyAfterDays` or `scheduleExpression` must be specified.
	ScheduleExpression *string `pulumi:"scheduleExpression"`
}

// SecretRotationRotationRulesInput is an input type that accepts SecretRotationRotationRulesArgs and SecretRotationRotationRulesOutput values.
// You can construct a concrete instance of `SecretRotationRotationRulesInput` via:
//
//	SecretRotationRotationRulesArgs{...}
type SecretRotationRotationRulesInput interface {
	pulumi.Input

	ToSecretRotationRotationRulesOutput() SecretRotationRotationRulesOutput
	ToSecretRotationRotationRulesOutputWithContext(context.Context) SecretRotationRotationRulesOutput
}

type SecretRotationRotationRulesArgs struct {
	// Specifies the number of days between automatic scheduled rotations of the secret. Either `automaticallyAfterDays` or `scheduleExpression` must be specified.
	AutomaticallyAfterDays pulumi.IntPtrInput `pulumi:"automaticallyAfterDays"`
	// The length of the rotation window in hours. For example, `3h` for a three hour window.
	Duration pulumi.StringPtrInput `pulumi:"duration"`
	// A `cron()` or `rate()` expression that defines the schedule for rotating your secret. Either `automaticallyAfterDays` or `scheduleExpression` must be specified.
	ScheduleExpression pulumi.StringPtrInput `pulumi:"scheduleExpression"`
}

func (SecretRotationRotationRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretRotationRotationRules)(nil)).Elem()
}

func (i SecretRotationRotationRulesArgs) ToSecretRotationRotationRulesOutput() SecretRotationRotationRulesOutput {
	return i.ToSecretRotationRotationRulesOutputWithContext(context.Background())
}

func (i SecretRotationRotationRulesArgs) ToSecretRotationRotationRulesOutputWithContext(ctx context.Context) SecretRotationRotationRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretRotationRotationRulesOutput)
}

func (i SecretRotationRotationRulesArgs) ToSecretRotationRotationRulesPtrOutput() SecretRotationRotationRulesPtrOutput {
	return i.ToSecretRotationRotationRulesPtrOutputWithContext(context.Background())
}

func (i SecretRotationRotationRulesArgs) ToSecretRotationRotationRulesPtrOutputWithContext(ctx context.Context) SecretRotationRotationRulesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretRotationRotationRulesOutput).ToSecretRotationRotationRulesPtrOutputWithContext(ctx)
}

// SecretRotationRotationRulesPtrInput is an input type that accepts SecretRotationRotationRulesArgs, SecretRotationRotationRulesPtr and SecretRotationRotationRulesPtrOutput values.
// You can construct a concrete instance of `SecretRotationRotationRulesPtrInput` via:
//
//	        SecretRotationRotationRulesArgs{...}
//
//	or:
//
//	        nil
type SecretRotationRotationRulesPtrInput interface {
	pulumi.Input

	ToSecretRotationRotationRulesPtrOutput() SecretRotationRotationRulesPtrOutput
	ToSecretRotationRotationRulesPtrOutputWithContext(context.Context) SecretRotationRotationRulesPtrOutput
}

type secretRotationRotationRulesPtrType SecretRotationRotationRulesArgs

func SecretRotationRotationRulesPtr(v *SecretRotationRotationRulesArgs) SecretRotationRotationRulesPtrInput {
	return (*secretRotationRotationRulesPtrType)(v)
}

func (*secretRotationRotationRulesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretRotationRotationRules)(nil)).Elem()
}

func (i *secretRotationRotationRulesPtrType) ToSecretRotationRotationRulesPtrOutput() SecretRotationRotationRulesPtrOutput {
	return i.ToSecretRotationRotationRulesPtrOutputWithContext(context.Background())
}

func (i *secretRotationRotationRulesPtrType) ToSecretRotationRotationRulesPtrOutputWithContext(ctx context.Context) SecretRotationRotationRulesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretRotationRotationRulesPtrOutput)
}

type SecretRotationRotationRulesOutput struct{ *pulumi.OutputState }

func (SecretRotationRotationRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretRotationRotationRules)(nil)).Elem()
}

func (o SecretRotationRotationRulesOutput) ToSecretRotationRotationRulesOutput() SecretRotationRotationRulesOutput {
	return o
}

func (o SecretRotationRotationRulesOutput) ToSecretRotationRotationRulesOutputWithContext(ctx context.Context) SecretRotationRotationRulesOutput {
	return o
}

func (o SecretRotationRotationRulesOutput) ToSecretRotationRotationRulesPtrOutput() SecretRotationRotationRulesPtrOutput {
	return o.ToSecretRotationRotationRulesPtrOutputWithContext(context.Background())
}

func (o SecretRotationRotationRulesOutput) ToSecretRotationRotationRulesPtrOutputWithContext(ctx context.Context) SecretRotationRotationRulesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecretRotationRotationRules) *SecretRotationRotationRules {
		return &v
	}).(SecretRotationRotationRulesPtrOutput)
}

// Specifies the number of days between automatic scheduled rotations of the secret. Either `automaticallyAfterDays` or `scheduleExpression` must be specified.
func (o SecretRotationRotationRulesOutput) AutomaticallyAfterDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SecretRotationRotationRules) *int { return v.AutomaticallyAfterDays }).(pulumi.IntPtrOutput)
}

// The length of the rotation window in hours. For example, `3h` for a three hour window.
func (o SecretRotationRotationRulesOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretRotationRotationRules) *string { return v.Duration }).(pulumi.StringPtrOutput)
}

// A `cron()` or `rate()` expression that defines the schedule for rotating your secret. Either `automaticallyAfterDays` or `scheduleExpression` must be specified.
func (o SecretRotationRotationRulesOutput) ScheduleExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretRotationRotationRules) *string { return v.ScheduleExpression }).(pulumi.StringPtrOutput)
}

type SecretRotationRotationRulesPtrOutput struct{ *pulumi.OutputState }

func (SecretRotationRotationRulesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretRotationRotationRules)(nil)).Elem()
}

func (o SecretRotationRotationRulesPtrOutput) ToSecretRotationRotationRulesPtrOutput() SecretRotationRotationRulesPtrOutput {
	return o
}

func (o SecretRotationRotationRulesPtrOutput) ToSecretRotationRotationRulesPtrOutputWithContext(ctx context.Context) SecretRotationRotationRulesPtrOutput {
	return o
}

func (o SecretRotationRotationRulesPtrOutput) Elem() SecretRotationRotationRulesOutput {
	return o.ApplyT(func(v *SecretRotationRotationRules) SecretRotationRotationRules {
		if v != nil {
			return *v
		}
		var ret SecretRotationRotationRules
		return ret
	}).(SecretRotationRotationRulesOutput)
}

// Specifies the number of days between automatic scheduled rotations of the secret. Either `automaticallyAfterDays` or `scheduleExpression` must be specified.
func (o SecretRotationRotationRulesPtrOutput) AutomaticallyAfterDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretRotationRotationRules) *int {
		if v == nil {
			return nil
		}
		return v.AutomaticallyAfterDays
	}).(pulumi.IntPtrOutput)
}

// The length of the rotation window in hours. For example, `3h` for a three hour window.
func (o SecretRotationRotationRulesPtrOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretRotationRotationRules) *string {
		if v == nil {
			return nil
		}
		return v.Duration
	}).(pulumi.StringPtrOutput)
}

// A `cron()` or `rate()` expression that defines the schedule for rotating your secret. Either `automaticallyAfterDays` or `scheduleExpression` must be specified.
func (o SecretRotationRotationRulesPtrOutput) ScheduleExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretRotationRotationRules) *string {
		if v == nil {
			return nil
		}
		return v.ScheduleExpression
	}).(pulumi.StringPtrOutput)
}

type GetSecretRotationRotationRule struct {
	AutomaticallyAfterDays int    `pulumi:"automaticallyAfterDays"`
	Duration               string `pulumi:"duration"`
	ScheduleExpression     string `pulumi:"scheduleExpression"`
}

// GetSecretRotationRotationRuleInput is an input type that accepts GetSecretRotationRotationRuleArgs and GetSecretRotationRotationRuleOutput values.
// You can construct a concrete instance of `GetSecretRotationRotationRuleInput` via:
//
//	GetSecretRotationRotationRuleArgs{...}
type GetSecretRotationRotationRuleInput interface {
	pulumi.Input

	ToGetSecretRotationRotationRuleOutput() GetSecretRotationRotationRuleOutput
	ToGetSecretRotationRotationRuleOutputWithContext(context.Context) GetSecretRotationRotationRuleOutput
}

type GetSecretRotationRotationRuleArgs struct {
	AutomaticallyAfterDays pulumi.IntInput    `pulumi:"automaticallyAfterDays"`
	Duration               pulumi.StringInput `pulumi:"duration"`
	ScheduleExpression     pulumi.StringInput `pulumi:"scheduleExpression"`
}

func (GetSecretRotationRotationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretRotationRotationRule)(nil)).Elem()
}

func (i GetSecretRotationRotationRuleArgs) ToGetSecretRotationRotationRuleOutput() GetSecretRotationRotationRuleOutput {
	return i.ToGetSecretRotationRotationRuleOutputWithContext(context.Background())
}

func (i GetSecretRotationRotationRuleArgs) ToGetSecretRotationRotationRuleOutputWithContext(ctx context.Context) GetSecretRotationRotationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSecretRotationRotationRuleOutput)
}

// GetSecretRotationRotationRuleArrayInput is an input type that accepts GetSecretRotationRotationRuleArray and GetSecretRotationRotationRuleArrayOutput values.
// You can construct a concrete instance of `GetSecretRotationRotationRuleArrayInput` via:
//
//	GetSecretRotationRotationRuleArray{ GetSecretRotationRotationRuleArgs{...} }
type GetSecretRotationRotationRuleArrayInput interface {
	pulumi.Input

	ToGetSecretRotationRotationRuleArrayOutput() GetSecretRotationRotationRuleArrayOutput
	ToGetSecretRotationRotationRuleArrayOutputWithContext(context.Context) GetSecretRotationRotationRuleArrayOutput
}

type GetSecretRotationRotationRuleArray []GetSecretRotationRotationRuleInput

func (GetSecretRotationRotationRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSecretRotationRotationRule)(nil)).Elem()
}

func (i GetSecretRotationRotationRuleArray) ToGetSecretRotationRotationRuleArrayOutput() GetSecretRotationRotationRuleArrayOutput {
	return i.ToGetSecretRotationRotationRuleArrayOutputWithContext(context.Background())
}

func (i GetSecretRotationRotationRuleArray) ToGetSecretRotationRotationRuleArrayOutputWithContext(ctx context.Context) GetSecretRotationRotationRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSecretRotationRotationRuleArrayOutput)
}

type GetSecretRotationRotationRuleOutput struct{ *pulumi.OutputState }

func (GetSecretRotationRotationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretRotationRotationRule)(nil)).Elem()
}

func (o GetSecretRotationRotationRuleOutput) ToGetSecretRotationRotationRuleOutput() GetSecretRotationRotationRuleOutput {
	return o
}

func (o GetSecretRotationRotationRuleOutput) ToGetSecretRotationRotationRuleOutputWithContext(ctx context.Context) GetSecretRotationRotationRuleOutput {
	return o
}

func (o GetSecretRotationRotationRuleOutput) AutomaticallyAfterDays() pulumi.IntOutput {
	return o.ApplyT(func(v GetSecretRotationRotationRule) int { return v.AutomaticallyAfterDays }).(pulumi.IntOutput)
}

func (o GetSecretRotationRotationRuleOutput) Duration() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretRotationRotationRule) string { return v.Duration }).(pulumi.StringOutput)
}

func (o GetSecretRotationRotationRuleOutput) ScheduleExpression() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretRotationRotationRule) string { return v.ScheduleExpression }).(pulumi.StringOutput)
}

type GetSecretRotationRotationRuleArrayOutput struct{ *pulumi.OutputState }

func (GetSecretRotationRotationRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSecretRotationRotationRule)(nil)).Elem()
}

func (o GetSecretRotationRotationRuleArrayOutput) ToGetSecretRotationRotationRuleArrayOutput() GetSecretRotationRotationRuleArrayOutput {
	return o
}

func (o GetSecretRotationRotationRuleArrayOutput) ToGetSecretRotationRotationRuleArrayOutputWithContext(ctx context.Context) GetSecretRotationRotationRuleArrayOutput {
	return o
}

func (o GetSecretRotationRotationRuleArrayOutput) Index(i pulumi.IntInput) GetSecretRotationRotationRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetSecretRotationRotationRule {
		return vs[0].([]GetSecretRotationRotationRule)[vs[1].(int)]
	}).(GetSecretRotationRotationRuleOutput)
}

type GetSecretsFilter struct {
	// Name of the filter field. Valid values can be found in the [Secrets Manager ListSecrets API Reference](https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_ListSecrets.html).
	Name string `pulumi:"name"`
	// Set of values that are accepted for the given filter field. Results will be selected if any given value matches.
	Values []string `pulumi:"values"`
}

// GetSecretsFilterInput is an input type that accepts GetSecretsFilterArgs and GetSecretsFilterOutput values.
// You can construct a concrete instance of `GetSecretsFilterInput` via:
//
//	GetSecretsFilterArgs{...}
type GetSecretsFilterInput interface {
	pulumi.Input

	ToGetSecretsFilterOutput() GetSecretsFilterOutput
	ToGetSecretsFilterOutputWithContext(context.Context) GetSecretsFilterOutput
}

type GetSecretsFilterArgs struct {
	// Name of the filter field. Valid values can be found in the [Secrets Manager ListSecrets API Reference](https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_ListSecrets.html).
	Name pulumi.StringInput `pulumi:"name"`
	// Set of values that are accepted for the given filter field. Results will be selected if any given value matches.
	Values pulumi.StringArrayInput `pulumi:"values"`
}

func (GetSecretsFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretsFilter)(nil)).Elem()
}

func (i GetSecretsFilterArgs) ToGetSecretsFilterOutput() GetSecretsFilterOutput {
	return i.ToGetSecretsFilterOutputWithContext(context.Background())
}

func (i GetSecretsFilterArgs) ToGetSecretsFilterOutputWithContext(ctx context.Context) GetSecretsFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSecretsFilterOutput)
}

// GetSecretsFilterArrayInput is an input type that accepts GetSecretsFilterArray and GetSecretsFilterArrayOutput values.
// You can construct a concrete instance of `GetSecretsFilterArrayInput` via:
//
//	GetSecretsFilterArray{ GetSecretsFilterArgs{...} }
type GetSecretsFilterArrayInput interface {
	pulumi.Input

	ToGetSecretsFilterArrayOutput() GetSecretsFilterArrayOutput
	ToGetSecretsFilterArrayOutputWithContext(context.Context) GetSecretsFilterArrayOutput
}

type GetSecretsFilterArray []GetSecretsFilterInput

func (GetSecretsFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSecretsFilter)(nil)).Elem()
}

func (i GetSecretsFilterArray) ToGetSecretsFilterArrayOutput() GetSecretsFilterArrayOutput {
	return i.ToGetSecretsFilterArrayOutputWithContext(context.Background())
}

func (i GetSecretsFilterArray) ToGetSecretsFilterArrayOutputWithContext(ctx context.Context) GetSecretsFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSecretsFilterArrayOutput)
}

type GetSecretsFilterOutput struct{ *pulumi.OutputState }

func (GetSecretsFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretsFilter)(nil)).Elem()
}

func (o GetSecretsFilterOutput) ToGetSecretsFilterOutput() GetSecretsFilterOutput {
	return o
}

func (o GetSecretsFilterOutput) ToGetSecretsFilterOutputWithContext(ctx context.Context) GetSecretsFilterOutput {
	return o
}

// Name of the filter field. Valid values can be found in the [Secrets Manager ListSecrets API Reference](https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_ListSecrets.html).
func (o GetSecretsFilterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretsFilter) string { return v.Name }).(pulumi.StringOutput)
}

// Set of values that are accepted for the given filter field. Results will be selected if any given value matches.
func (o GetSecretsFilterOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSecretsFilter) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type GetSecretsFilterArrayOutput struct{ *pulumi.OutputState }

func (GetSecretsFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSecretsFilter)(nil)).Elem()
}

func (o GetSecretsFilterArrayOutput) ToGetSecretsFilterArrayOutput() GetSecretsFilterArrayOutput {
	return o
}

func (o GetSecretsFilterArrayOutput) ToGetSecretsFilterArrayOutputWithContext(ctx context.Context) GetSecretsFilterArrayOutput {
	return o
}

func (o GetSecretsFilterArrayOutput) Index(i pulumi.IntInput) GetSecretsFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetSecretsFilter {
		return vs[0].([]GetSecretsFilter)[vs[1].(int)]
	}).(GetSecretsFilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretReplicaInput)(nil)).Elem(), SecretReplicaArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretReplicaArrayInput)(nil)).Elem(), SecretReplicaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretRotationRotationRulesInput)(nil)).Elem(), SecretRotationRotationRulesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretRotationRotationRulesPtrInput)(nil)).Elem(), SecretRotationRotationRulesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetSecretRotationRotationRuleInput)(nil)).Elem(), GetSecretRotationRotationRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetSecretRotationRotationRuleArrayInput)(nil)).Elem(), GetSecretRotationRotationRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetSecretsFilterInput)(nil)).Elem(), GetSecretsFilterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetSecretsFilterArrayInput)(nil)).Elem(), GetSecretsFilterArray{})
	pulumi.RegisterOutputType(SecretReplicaOutput{})
	pulumi.RegisterOutputType(SecretReplicaArrayOutput{})
	pulumi.RegisterOutputType(SecretRotationRotationRulesOutput{})
	pulumi.RegisterOutputType(SecretRotationRotationRulesPtrOutput{})
	pulumi.RegisterOutputType(GetSecretRotationRotationRuleOutput{})
	pulumi.RegisterOutputType(GetSecretRotationRotationRuleArrayOutput{})
	pulumi.RegisterOutputType(GetSecretsFilterOutput{})
	pulumi.RegisterOutputType(GetSecretsFilterArrayOutput{})
}
