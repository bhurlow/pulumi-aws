// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CloudFront real-time log configuration resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cloudfront"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudfront.NewMonitoringSubscription(ctx, "example", &cloudfront.MonitoringSubscriptionArgs{
//				DistributionId: pulumi.Any(aws_cloudfront_distribution.Example.Id),
//				MonitoringSubscription: &cloudfront.MonitoringSubscriptionMonitoringSubscriptionArgs{
//					RealtimeMetricsSubscriptionConfig: &cloudfront.MonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigArgs{
//						RealtimeMetricsSubscriptionStatus: pulumi.String("Enabled"),
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
// CloudFront monitoring subscription can be imported using the id, e.g.,
//
// ```sh
//
//	$ pulumi import aws:cloudfront/monitoringSubscription:MonitoringSubscription example E3QYSUHO4VYRGB
//
// ```
type MonitoringSubscription struct {
	pulumi.CustomResourceState

	// The ID of the distribution that you are enabling metrics for.
	DistributionId pulumi.StringOutput `pulumi:"distributionId"`
	// A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
	MonitoringSubscription MonitoringSubscriptionMonitoringSubscriptionOutput `pulumi:"monitoringSubscription"`
}

// NewMonitoringSubscription registers a new resource with the given unique name, arguments, and options.
func NewMonitoringSubscription(ctx *pulumi.Context,
	name string, args *MonitoringSubscriptionArgs, opts ...pulumi.ResourceOption) (*MonitoringSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DistributionId == nil {
		return nil, errors.New("invalid value for required argument 'DistributionId'")
	}
	if args.MonitoringSubscription == nil {
		return nil, errors.New("invalid value for required argument 'MonitoringSubscription'")
	}
	var resource MonitoringSubscription
	err := ctx.RegisterResource("aws:cloudfront/monitoringSubscription:MonitoringSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMonitoringSubscription gets an existing MonitoringSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMonitoringSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitoringSubscriptionState, opts ...pulumi.ResourceOption) (*MonitoringSubscription, error) {
	var resource MonitoringSubscription
	err := ctx.ReadResource("aws:cloudfront/monitoringSubscription:MonitoringSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MonitoringSubscription resources.
type monitoringSubscriptionState struct {
	// The ID of the distribution that you are enabling metrics for.
	DistributionId *string `pulumi:"distributionId"`
	// A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
	MonitoringSubscription *MonitoringSubscriptionMonitoringSubscription `pulumi:"monitoringSubscription"`
}

type MonitoringSubscriptionState struct {
	// The ID of the distribution that you are enabling metrics for.
	DistributionId pulumi.StringPtrInput
	// A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
	MonitoringSubscription MonitoringSubscriptionMonitoringSubscriptionPtrInput
}

func (MonitoringSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringSubscriptionState)(nil)).Elem()
}

type monitoringSubscriptionArgs struct {
	// The ID of the distribution that you are enabling metrics for.
	DistributionId string `pulumi:"distributionId"`
	// A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
	MonitoringSubscription MonitoringSubscriptionMonitoringSubscription `pulumi:"monitoringSubscription"`
}

// The set of arguments for constructing a MonitoringSubscription resource.
type MonitoringSubscriptionArgs struct {
	// The ID of the distribution that you are enabling metrics for.
	DistributionId pulumi.StringInput
	// A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
	MonitoringSubscription MonitoringSubscriptionMonitoringSubscriptionInput
}

func (MonitoringSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringSubscriptionArgs)(nil)).Elem()
}

type MonitoringSubscriptionInput interface {
	pulumi.Input

	ToMonitoringSubscriptionOutput() MonitoringSubscriptionOutput
	ToMonitoringSubscriptionOutputWithContext(ctx context.Context) MonitoringSubscriptionOutput
}

func (*MonitoringSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringSubscription)(nil)).Elem()
}

func (i *MonitoringSubscription) ToMonitoringSubscriptionOutput() MonitoringSubscriptionOutput {
	return i.ToMonitoringSubscriptionOutputWithContext(context.Background())
}

func (i *MonitoringSubscription) ToMonitoringSubscriptionOutputWithContext(ctx context.Context) MonitoringSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringSubscriptionOutput)
}

// MonitoringSubscriptionArrayInput is an input type that accepts MonitoringSubscriptionArray and MonitoringSubscriptionArrayOutput values.
// You can construct a concrete instance of `MonitoringSubscriptionArrayInput` via:
//
//	MonitoringSubscriptionArray{ MonitoringSubscriptionArgs{...} }
type MonitoringSubscriptionArrayInput interface {
	pulumi.Input

	ToMonitoringSubscriptionArrayOutput() MonitoringSubscriptionArrayOutput
	ToMonitoringSubscriptionArrayOutputWithContext(context.Context) MonitoringSubscriptionArrayOutput
}

type MonitoringSubscriptionArray []MonitoringSubscriptionInput

func (MonitoringSubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MonitoringSubscription)(nil)).Elem()
}

func (i MonitoringSubscriptionArray) ToMonitoringSubscriptionArrayOutput() MonitoringSubscriptionArrayOutput {
	return i.ToMonitoringSubscriptionArrayOutputWithContext(context.Background())
}

func (i MonitoringSubscriptionArray) ToMonitoringSubscriptionArrayOutputWithContext(ctx context.Context) MonitoringSubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringSubscriptionArrayOutput)
}

// MonitoringSubscriptionMapInput is an input type that accepts MonitoringSubscriptionMap and MonitoringSubscriptionMapOutput values.
// You can construct a concrete instance of `MonitoringSubscriptionMapInput` via:
//
//	MonitoringSubscriptionMap{ "key": MonitoringSubscriptionArgs{...} }
type MonitoringSubscriptionMapInput interface {
	pulumi.Input

	ToMonitoringSubscriptionMapOutput() MonitoringSubscriptionMapOutput
	ToMonitoringSubscriptionMapOutputWithContext(context.Context) MonitoringSubscriptionMapOutput
}

type MonitoringSubscriptionMap map[string]MonitoringSubscriptionInput

func (MonitoringSubscriptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MonitoringSubscription)(nil)).Elem()
}

func (i MonitoringSubscriptionMap) ToMonitoringSubscriptionMapOutput() MonitoringSubscriptionMapOutput {
	return i.ToMonitoringSubscriptionMapOutputWithContext(context.Background())
}

func (i MonitoringSubscriptionMap) ToMonitoringSubscriptionMapOutputWithContext(ctx context.Context) MonitoringSubscriptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringSubscriptionMapOutput)
}

type MonitoringSubscriptionOutput struct{ *pulumi.OutputState }

func (MonitoringSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringSubscription)(nil)).Elem()
}

func (o MonitoringSubscriptionOutput) ToMonitoringSubscriptionOutput() MonitoringSubscriptionOutput {
	return o
}

func (o MonitoringSubscriptionOutput) ToMonitoringSubscriptionOutputWithContext(ctx context.Context) MonitoringSubscriptionOutput {
	return o
}

// The ID of the distribution that you are enabling metrics for.
func (o MonitoringSubscriptionOutput) DistributionId() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringSubscription) pulumi.StringOutput { return v.DistributionId }).(pulumi.StringOutput)
}

// A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
func (o MonitoringSubscriptionOutput) MonitoringSubscription() MonitoringSubscriptionMonitoringSubscriptionOutput {
	return o.ApplyT(func(v *MonitoringSubscription) MonitoringSubscriptionMonitoringSubscriptionOutput {
		return v.MonitoringSubscription
	}).(MonitoringSubscriptionMonitoringSubscriptionOutput)
}

type MonitoringSubscriptionArrayOutput struct{ *pulumi.OutputState }

func (MonitoringSubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MonitoringSubscription)(nil)).Elem()
}

func (o MonitoringSubscriptionArrayOutput) ToMonitoringSubscriptionArrayOutput() MonitoringSubscriptionArrayOutput {
	return o
}

func (o MonitoringSubscriptionArrayOutput) ToMonitoringSubscriptionArrayOutputWithContext(ctx context.Context) MonitoringSubscriptionArrayOutput {
	return o
}

func (o MonitoringSubscriptionArrayOutput) Index(i pulumi.IntInput) MonitoringSubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MonitoringSubscription {
		return vs[0].([]*MonitoringSubscription)[vs[1].(int)]
	}).(MonitoringSubscriptionOutput)
}

type MonitoringSubscriptionMapOutput struct{ *pulumi.OutputState }

func (MonitoringSubscriptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MonitoringSubscription)(nil)).Elem()
}

func (o MonitoringSubscriptionMapOutput) ToMonitoringSubscriptionMapOutput() MonitoringSubscriptionMapOutput {
	return o
}

func (o MonitoringSubscriptionMapOutput) ToMonitoringSubscriptionMapOutputWithContext(ctx context.Context) MonitoringSubscriptionMapOutput {
	return o
}

func (o MonitoringSubscriptionMapOutput) MapIndex(k pulumi.StringInput) MonitoringSubscriptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MonitoringSubscription {
		return vs[0].(map[string]*MonitoringSubscription)[vs[1].(string)]
	}).(MonitoringSubscriptionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MonitoringSubscriptionInput)(nil)).Elem(), &MonitoringSubscription{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitoringSubscriptionArrayInput)(nil)).Elem(), MonitoringSubscriptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitoringSubscriptionMapInput)(nil)).Elem(), MonitoringSubscriptionMap{})
	pulumi.RegisterOutputType(MonitoringSubscriptionOutput{})
	pulumi.RegisterOutputType(MonitoringSubscriptionArrayOutput{})
	pulumi.RegisterOutputType(MonitoringSubscriptionMapOutput{})
}
