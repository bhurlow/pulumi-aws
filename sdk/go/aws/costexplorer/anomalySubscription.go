// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package costexplorer

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a CE Anomaly Subscription.
//
// ## Example Usage
//
// ### Threshold Expression Example
// ### For a Specific Dimension
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/costexplorer"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := costexplorer.NewAnomalySubscription(ctx, "test", &costexplorer.AnomalySubscriptionArgs{
//				Frequency: pulumi.String("DAILY"),
//				MonitorArnLists: pulumi.StringArray{
//					aws_ce_anomaly_monitor.Test.Arn,
//				},
//				Subscribers: costexplorer.AnomalySubscriptionSubscriberArray{
//					&costexplorer.AnomalySubscriptionSubscriberArgs{
//						Type:    pulumi.String("EMAIL"),
//						Address: pulumi.String("abc@example.com"),
//					},
//				},
//				ThresholdExpression: &costexplorer.AnomalySubscriptionThresholdExpressionArgs{
//					Dimension: &costexplorer.AnomalySubscriptionThresholdExpressionDimensionArgs{
//						Key: pulumi.String("ANOMALY_TOTAL_IMPACT_ABSOLUTE"),
//						Values: pulumi.StringArray{
//							pulumi.String("100.0"),
//						},
//						MatchOptions: pulumi.StringArray{
//							pulumi.String("GREATER_THAN_OR_EQUAL"),
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
// ### Using an `and` Expression
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/costexplorer"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := costexplorer.NewAnomalySubscription(ctx, "test", &costexplorer.AnomalySubscriptionArgs{
//				Frequency: pulumi.String("DAILY"),
//				MonitorArnLists: pulumi.StringArray{
//					aws_ce_anomaly_monitor.Test.Arn,
//				},
//				Subscribers: costexplorer.AnomalySubscriptionSubscriberArray{
//					&costexplorer.AnomalySubscriptionSubscriberArgs{
//						Type:    pulumi.String("EMAIL"),
//						Address: pulumi.String("abc@example.com"),
//					},
//				},
//				ThresholdExpression: &costexplorer.AnomalySubscriptionThresholdExpressionArgs{
//					Ands: costexplorer.AnomalySubscriptionThresholdExpressionAndArray{
//						&costexplorer.AnomalySubscriptionThresholdExpressionAndArgs{
//							Dimension: &costexplorer.AnomalySubscriptionThresholdExpressionAndDimensionArgs{
//								Key: pulumi.String("ANOMALY_TOTAL_IMPACT_ABSOLUTE"),
//								MatchOptions: pulumi.StringArray{
//									pulumi.String("GREATER_THAN_OR_EQUAL"),
//								},
//								Values: pulumi.StringArray{
//									pulumi.String("100"),
//								},
//							},
//						},
//						&costexplorer.AnomalySubscriptionThresholdExpressionAndArgs{
//							Dimension: &costexplorer.AnomalySubscriptionThresholdExpressionAndDimensionArgs{
//								Key: pulumi.String("ANOMALY_TOTAL_IMPACT_PERCENTAGE"),
//								MatchOptions: pulumi.StringArray{
//									pulumi.String("GREATER_THAN_OR_EQUAL"),
//								},
//								Values: pulumi.StringArray{
//									pulumi.String("50"),
//								},
//							},
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
// Using `pulumi import`, import `aws_ce_anomaly_subscription` using the `id`. For example:
//
// ```sh
//
//	$ pulumi import aws:costexplorer/anomalySubscription:AnomalySubscription example AnomalySubscriptionARN
//
// ```
type AnomalySubscription struct {
	pulumi.CustomResourceState

	// The unique identifier for the AWS account in which the anomaly subscription ought to be created.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// ARN of the anomaly subscription.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The frequency that anomaly reports are sent. Valid Values: `DAILY` | `IMMEDIATE` | `WEEKLY`.
	Frequency pulumi.StringOutput `pulumi:"frequency"`
	// A list of cost anomaly monitors.
	MonitorArnLists pulumi.StringArrayOutput `pulumi:"monitorArnLists"`
	// The name for the subscription.
	Name pulumi.StringOutput `pulumi:"name"`
	// A subscriber configuration. Multiple subscribers can be defined.
	Subscribers AnomalySubscriptionSubscriberArrayOutput `pulumi:"subscribers"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// An Expression object used to specify the anomalies that you want to generate alerts for. See Threshold Expression.
	ThresholdExpression AnomalySubscriptionThresholdExpressionOutput `pulumi:"thresholdExpression"`
}

// NewAnomalySubscription registers a new resource with the given unique name, arguments, and options.
func NewAnomalySubscription(ctx *pulumi.Context,
	name string, args *AnomalySubscriptionArgs, opts ...pulumi.ResourceOption) (*AnomalySubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Frequency == nil {
		return nil, errors.New("invalid value for required argument 'Frequency'")
	}
	if args.MonitorArnLists == nil {
		return nil, errors.New("invalid value for required argument 'MonitorArnLists'")
	}
	if args.Subscribers == nil {
		return nil, errors.New("invalid value for required argument 'Subscribers'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AnomalySubscription
	err := ctx.RegisterResource("aws:costexplorer/anomalySubscription:AnomalySubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnomalySubscription gets an existing AnomalySubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnomalySubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnomalySubscriptionState, opts ...pulumi.ResourceOption) (*AnomalySubscription, error) {
	var resource AnomalySubscription
	err := ctx.ReadResource("aws:costexplorer/anomalySubscription:AnomalySubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AnomalySubscription resources.
type anomalySubscriptionState struct {
	// The unique identifier for the AWS account in which the anomaly subscription ought to be created.
	AccountId *string `pulumi:"accountId"`
	// ARN of the anomaly subscription.
	Arn *string `pulumi:"arn"`
	// The frequency that anomaly reports are sent. Valid Values: `DAILY` | `IMMEDIATE` | `WEEKLY`.
	Frequency *string `pulumi:"frequency"`
	// A list of cost anomaly monitors.
	MonitorArnLists []string `pulumi:"monitorArnLists"`
	// The name for the subscription.
	Name *string `pulumi:"name"`
	// A subscriber configuration. Multiple subscribers can be defined.
	Subscribers []AnomalySubscriptionSubscriber `pulumi:"subscribers"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// An Expression object used to specify the anomalies that you want to generate alerts for. See Threshold Expression.
	ThresholdExpression *AnomalySubscriptionThresholdExpression `pulumi:"thresholdExpression"`
}

type AnomalySubscriptionState struct {
	// The unique identifier for the AWS account in which the anomaly subscription ought to be created.
	AccountId pulumi.StringPtrInput
	// ARN of the anomaly subscription.
	Arn pulumi.StringPtrInput
	// The frequency that anomaly reports are sent. Valid Values: `DAILY` | `IMMEDIATE` | `WEEKLY`.
	Frequency pulumi.StringPtrInput
	// A list of cost anomaly monitors.
	MonitorArnLists pulumi.StringArrayInput
	// The name for the subscription.
	Name pulumi.StringPtrInput
	// A subscriber configuration. Multiple subscribers can be defined.
	Subscribers AnomalySubscriptionSubscriberArrayInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// An Expression object used to specify the anomalies that you want to generate alerts for. See Threshold Expression.
	ThresholdExpression AnomalySubscriptionThresholdExpressionPtrInput
}

func (AnomalySubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*anomalySubscriptionState)(nil)).Elem()
}

type anomalySubscriptionArgs struct {
	// The unique identifier for the AWS account in which the anomaly subscription ought to be created.
	AccountId *string `pulumi:"accountId"`
	// The frequency that anomaly reports are sent. Valid Values: `DAILY` | `IMMEDIATE` | `WEEKLY`.
	Frequency string `pulumi:"frequency"`
	// A list of cost anomaly monitors.
	MonitorArnLists []string `pulumi:"monitorArnLists"`
	// The name for the subscription.
	Name *string `pulumi:"name"`
	// A subscriber configuration. Multiple subscribers can be defined.
	Subscribers []AnomalySubscriptionSubscriber `pulumi:"subscribers"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// An Expression object used to specify the anomalies that you want to generate alerts for. See Threshold Expression.
	ThresholdExpression *AnomalySubscriptionThresholdExpression `pulumi:"thresholdExpression"`
}

// The set of arguments for constructing a AnomalySubscription resource.
type AnomalySubscriptionArgs struct {
	// The unique identifier for the AWS account in which the anomaly subscription ought to be created.
	AccountId pulumi.StringPtrInput
	// The frequency that anomaly reports are sent. Valid Values: `DAILY` | `IMMEDIATE` | `WEEKLY`.
	Frequency pulumi.StringInput
	// A list of cost anomaly monitors.
	MonitorArnLists pulumi.StringArrayInput
	// The name for the subscription.
	Name pulumi.StringPtrInput
	// A subscriber configuration. Multiple subscribers can be defined.
	Subscribers AnomalySubscriptionSubscriberArrayInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// An Expression object used to specify the anomalies that you want to generate alerts for. See Threshold Expression.
	ThresholdExpression AnomalySubscriptionThresholdExpressionPtrInput
}

func (AnomalySubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*anomalySubscriptionArgs)(nil)).Elem()
}

type AnomalySubscriptionInput interface {
	pulumi.Input

	ToAnomalySubscriptionOutput() AnomalySubscriptionOutput
	ToAnomalySubscriptionOutputWithContext(ctx context.Context) AnomalySubscriptionOutput
}

func (*AnomalySubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**AnomalySubscription)(nil)).Elem()
}

func (i *AnomalySubscription) ToAnomalySubscriptionOutput() AnomalySubscriptionOutput {
	return i.ToAnomalySubscriptionOutputWithContext(context.Background())
}

func (i *AnomalySubscription) ToAnomalySubscriptionOutputWithContext(ctx context.Context) AnomalySubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnomalySubscriptionOutput)
}

func (i *AnomalySubscription) ToOutput(ctx context.Context) pulumix.Output[*AnomalySubscription] {
	return pulumix.Output[*AnomalySubscription]{
		OutputState: i.ToAnomalySubscriptionOutputWithContext(ctx).OutputState,
	}
}

// AnomalySubscriptionArrayInput is an input type that accepts AnomalySubscriptionArray and AnomalySubscriptionArrayOutput values.
// You can construct a concrete instance of `AnomalySubscriptionArrayInput` via:
//
//	AnomalySubscriptionArray{ AnomalySubscriptionArgs{...} }
type AnomalySubscriptionArrayInput interface {
	pulumi.Input

	ToAnomalySubscriptionArrayOutput() AnomalySubscriptionArrayOutput
	ToAnomalySubscriptionArrayOutputWithContext(context.Context) AnomalySubscriptionArrayOutput
}

type AnomalySubscriptionArray []AnomalySubscriptionInput

func (AnomalySubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AnomalySubscription)(nil)).Elem()
}

func (i AnomalySubscriptionArray) ToAnomalySubscriptionArrayOutput() AnomalySubscriptionArrayOutput {
	return i.ToAnomalySubscriptionArrayOutputWithContext(context.Background())
}

func (i AnomalySubscriptionArray) ToAnomalySubscriptionArrayOutputWithContext(ctx context.Context) AnomalySubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnomalySubscriptionArrayOutput)
}

func (i AnomalySubscriptionArray) ToOutput(ctx context.Context) pulumix.Output[[]*AnomalySubscription] {
	return pulumix.Output[[]*AnomalySubscription]{
		OutputState: i.ToAnomalySubscriptionArrayOutputWithContext(ctx).OutputState,
	}
}

// AnomalySubscriptionMapInput is an input type that accepts AnomalySubscriptionMap and AnomalySubscriptionMapOutput values.
// You can construct a concrete instance of `AnomalySubscriptionMapInput` via:
//
//	AnomalySubscriptionMap{ "key": AnomalySubscriptionArgs{...} }
type AnomalySubscriptionMapInput interface {
	pulumi.Input

	ToAnomalySubscriptionMapOutput() AnomalySubscriptionMapOutput
	ToAnomalySubscriptionMapOutputWithContext(context.Context) AnomalySubscriptionMapOutput
}

type AnomalySubscriptionMap map[string]AnomalySubscriptionInput

func (AnomalySubscriptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AnomalySubscription)(nil)).Elem()
}

func (i AnomalySubscriptionMap) ToAnomalySubscriptionMapOutput() AnomalySubscriptionMapOutput {
	return i.ToAnomalySubscriptionMapOutputWithContext(context.Background())
}

func (i AnomalySubscriptionMap) ToAnomalySubscriptionMapOutputWithContext(ctx context.Context) AnomalySubscriptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnomalySubscriptionMapOutput)
}

func (i AnomalySubscriptionMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*AnomalySubscription] {
	return pulumix.Output[map[string]*AnomalySubscription]{
		OutputState: i.ToAnomalySubscriptionMapOutputWithContext(ctx).OutputState,
	}
}

type AnomalySubscriptionOutput struct{ *pulumi.OutputState }

func (AnomalySubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnomalySubscription)(nil)).Elem()
}

func (o AnomalySubscriptionOutput) ToAnomalySubscriptionOutput() AnomalySubscriptionOutput {
	return o
}

func (o AnomalySubscriptionOutput) ToAnomalySubscriptionOutputWithContext(ctx context.Context) AnomalySubscriptionOutput {
	return o
}

func (o AnomalySubscriptionOutput) ToOutput(ctx context.Context) pulumix.Output[*AnomalySubscription] {
	return pulumix.Output[*AnomalySubscription]{
		OutputState: o.OutputState,
	}
}

// The unique identifier for the AWS account in which the anomaly subscription ought to be created.
func (o AnomalySubscriptionOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalySubscription) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// ARN of the anomaly subscription.
func (o AnomalySubscriptionOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalySubscription) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The frequency that anomaly reports are sent. Valid Values: `DAILY` | `IMMEDIATE` | `WEEKLY`.
func (o AnomalySubscriptionOutput) Frequency() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalySubscription) pulumi.StringOutput { return v.Frequency }).(pulumi.StringOutput)
}

// A list of cost anomaly monitors.
func (o AnomalySubscriptionOutput) MonitorArnLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AnomalySubscription) pulumi.StringArrayOutput { return v.MonitorArnLists }).(pulumi.StringArrayOutput)
}

// The name for the subscription.
func (o AnomalySubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AnomalySubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A subscriber configuration. Multiple subscribers can be defined.
func (o AnomalySubscriptionOutput) Subscribers() AnomalySubscriptionSubscriberArrayOutput {
	return o.ApplyT(func(v *AnomalySubscription) AnomalySubscriptionSubscriberArrayOutput { return v.Subscribers }).(AnomalySubscriptionSubscriberArrayOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o AnomalySubscriptionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AnomalySubscription) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o AnomalySubscriptionOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AnomalySubscription) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// An Expression object used to specify the anomalies that you want to generate alerts for. See Threshold Expression.
func (o AnomalySubscriptionOutput) ThresholdExpression() AnomalySubscriptionThresholdExpressionOutput {
	return o.ApplyT(func(v *AnomalySubscription) AnomalySubscriptionThresholdExpressionOutput {
		return v.ThresholdExpression
	}).(AnomalySubscriptionThresholdExpressionOutput)
}

type AnomalySubscriptionArrayOutput struct{ *pulumi.OutputState }

func (AnomalySubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AnomalySubscription)(nil)).Elem()
}

func (o AnomalySubscriptionArrayOutput) ToAnomalySubscriptionArrayOutput() AnomalySubscriptionArrayOutput {
	return o
}

func (o AnomalySubscriptionArrayOutput) ToAnomalySubscriptionArrayOutputWithContext(ctx context.Context) AnomalySubscriptionArrayOutput {
	return o
}

func (o AnomalySubscriptionArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*AnomalySubscription] {
	return pulumix.Output[[]*AnomalySubscription]{
		OutputState: o.OutputState,
	}
}

func (o AnomalySubscriptionArrayOutput) Index(i pulumi.IntInput) AnomalySubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AnomalySubscription {
		return vs[0].([]*AnomalySubscription)[vs[1].(int)]
	}).(AnomalySubscriptionOutput)
}

type AnomalySubscriptionMapOutput struct{ *pulumi.OutputState }

func (AnomalySubscriptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AnomalySubscription)(nil)).Elem()
}

func (o AnomalySubscriptionMapOutput) ToAnomalySubscriptionMapOutput() AnomalySubscriptionMapOutput {
	return o
}

func (o AnomalySubscriptionMapOutput) ToAnomalySubscriptionMapOutputWithContext(ctx context.Context) AnomalySubscriptionMapOutput {
	return o
}

func (o AnomalySubscriptionMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*AnomalySubscription] {
	return pulumix.Output[map[string]*AnomalySubscription]{
		OutputState: o.OutputState,
	}
}

func (o AnomalySubscriptionMapOutput) MapIndex(k pulumi.StringInput) AnomalySubscriptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AnomalySubscription {
		return vs[0].(map[string]*AnomalySubscription)[vs[1].(string)]
	}).(AnomalySubscriptionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AnomalySubscriptionInput)(nil)).Elem(), &AnomalySubscription{})
	pulumi.RegisterInputType(reflect.TypeOf((*AnomalySubscriptionArrayInput)(nil)).Elem(), AnomalySubscriptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AnomalySubscriptionMapInput)(nil)).Elem(), AnomalySubscriptionMap{})
	pulumi.RegisterOutputType(AnomalySubscriptionOutput{})
	pulumi.RegisterOutputType(AnomalySubscriptionArrayOutput{})
	pulumi.RegisterOutputType(AnomalySubscriptionMapOutput{})
}
