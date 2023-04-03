// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Route53 health check.
//
// ## Example Usage
// ### Connectivity and HTTP Status Code Check
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/route53"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := route53.NewHealthCheck(ctx, "example", &route53.HealthCheckArgs{
//				FailureThreshold: pulumi.Int(5),
//				Fqdn:             pulumi.String("example.com"),
//				Port:             pulumi.Int(80),
//				RequestInterval:  pulumi.Int(30),
//				ResourcePath:     pulumi.String("/"),
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("tf-test-health-check"),
//				},
//				Type: pulumi.String("HTTP"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Connectivity and String Matching Check
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/route53"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := route53.NewHealthCheck(ctx, "example", &route53.HealthCheckArgs{
//				FailureThreshold: pulumi.Int(5),
//				Fqdn:             pulumi.String("example.com"),
//				Port:             pulumi.Int(443),
//				RequestInterval:  pulumi.Int(30),
//				ResourcePath:     pulumi.String("/"),
//				SearchString:     pulumi.String("example"),
//				Type:             pulumi.String("HTTPS_STR_MATCH"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Aggregate Check
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/route53"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := route53.NewHealthCheck(ctx, "parent", &route53.HealthCheckArgs{
//				Type:                 pulumi.String("CALCULATED"),
//				ChildHealthThreshold: pulumi.Int(1),
//				ChildHealthchecks: pulumi.StringArray{
//					aws_route53_health_check.Child.Id,
//				},
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("tf-test-calculated-health-check"),
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
// ### CloudWatch Alarm Check
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cloudwatch"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/route53"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			foobar, err := cloudwatch.NewMetricAlarm(ctx, "foobar", &cloudwatch.MetricAlarmArgs{
//				ComparisonOperator: pulumi.String("GreaterThanOrEqualToThreshold"),
//				EvaluationPeriods:  pulumi.Int(2),
//				MetricName:         pulumi.String("CPUUtilization"),
//				Namespace:          pulumi.String("AWS/EC2"),
//				Period:             pulumi.Int(120),
//				Statistic:          pulumi.String("Average"),
//				Threshold:          pulumi.Float64(80),
//				AlarmDescription:   pulumi.String("This metric monitors ec2 cpu utilization"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = route53.NewHealthCheck(ctx, "foo", &route53.HealthCheckArgs{
//				Type:                         pulumi.String("CLOUDWATCH_METRIC"),
//				CloudwatchAlarmName:          foobar.Name,
//				CloudwatchAlarmRegion:        pulumi.String("us-west-2"),
//				InsufficientDataHealthStatus: pulumi.String("Healthy"),
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
// Route53 Health Checks can be imported using the `health check id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:route53/healthCheck:HealthCheck http_check abcdef11-2222-3333-4444-555555fedcba
//
// ```
type HealthCheck struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the Health Check.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The minimum number of child health checks that must be healthy for Route 53 to consider the parent health check to be healthy. Valid values are integers between 0 and 256, inclusive
	ChildHealthThreshold pulumi.IntPtrOutput `pulumi:"childHealthThreshold"`
	// For a specified parent health check, a list of HealthCheckId values for the associated child health checks.
	ChildHealthchecks pulumi.StringArrayOutput `pulumi:"childHealthchecks"`
	// The name of the CloudWatch alarm.
	CloudwatchAlarmName pulumi.StringPtrOutput `pulumi:"cloudwatchAlarmName"`
	// The CloudWatchRegion that the CloudWatch alarm was created in.
	CloudwatchAlarmRegion pulumi.StringPtrOutput `pulumi:"cloudwatchAlarmRegion"`
	// A boolean value that stops Route 53 from performing health checks. When set to true, Route 53 will do the following depending on the type of health check:
	// * For health checks that check the health of endpoints, Route5 53 stops submitting requests to your application, server, or other resource.
	// * For calculated health checks, Route 53 stops aggregating the status of the referenced health checks.
	// * For health checks that monitor CloudWatch alarms, Route 53 stops monitoring the corresponding CloudWatch metrics.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// A boolean value that indicates whether Route53 should send the `fqdn` to the endpoint when performing the health check. This defaults to AWS' defaults: when the `type` is "HTTPS" `enableSni` defaults to `true`, when `type` is anything else `enableSni` defaults to `false`.
	EnableSni pulumi.BoolOutput `pulumi:"enableSni"`
	// The number of consecutive health checks that an endpoint must pass or fail.
	FailureThreshold pulumi.IntOutput `pulumi:"failureThreshold"`
	// The fully qualified domain name of the endpoint to be checked.
	Fqdn pulumi.StringPtrOutput `pulumi:"fqdn"`
	// The status of the health check when CloudWatch has insufficient data about the state of associated alarm. Valid values are `Healthy` , `Unhealthy` and `LastKnownStatus`.
	InsufficientDataHealthStatus pulumi.StringPtrOutput `pulumi:"insufficientDataHealthStatus"`
	// A boolean value that indicates whether the status of health check should be inverted. For example, if a health check is healthy but Inverted is True , then Route 53 considers the health check to be unhealthy.
	InvertHealthcheck pulumi.BoolPtrOutput `pulumi:"invertHealthcheck"`
	// The IP address of the endpoint to be checked.
	IpAddress pulumi.StringPtrOutput `pulumi:"ipAddress"`
	// A Boolean value that indicates whether you want Route 53 to measure the latency between health checkers in multiple AWS regions and your endpoint and to display CloudWatch latency graphs in the Route 53 console.
	MeasureLatency pulumi.BoolPtrOutput `pulumi:"measureLatency"`
	// The port of the endpoint to be checked.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// This is a reference name used in Caller Reference
	// (helpful for identifying single healthCheck set amongst others)
	ReferenceName pulumi.StringPtrOutput `pulumi:"referenceName"`
	// A list of AWS regions that you want Amazon Route 53 health checkers to check the specified endpoint from.
	Regions pulumi.StringArrayOutput `pulumi:"regions"`
	// The number of seconds between the time that Amazon Route 53 gets a response from your endpoint and the time that it sends the next health-check request.
	RequestInterval pulumi.IntPtrOutput `pulumi:"requestInterval"`
	// The path that you want Amazon Route 53 to request when performing health checks.
	ResourcePath pulumi.StringPtrOutput `pulumi:"resourcePath"`
	// The Amazon Resource Name (ARN) for the Route 53 Application Recovery Controller routing control. This is used when health check type is `RECOVERY_CONTROL`
	RoutingControlArn pulumi.StringPtrOutput `pulumi:"routingControlArn"`
	// String searched in the first 5120 bytes of the response body for check to be considered healthy. Only valid with `HTTP_STR_MATCH` and `HTTPS_STR_MATCH`.
	SearchString pulumi.StringPtrOutput `pulumi:"searchString"`
	// A map of tags to assign to the health check. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The protocol to use when performing health checks. Valid values are `HTTP`, `HTTPS`, `HTTP_STR_MATCH`, `HTTPS_STR_MATCH`, `TCP`, `CALCULATED`, `CLOUDWATCH_METRIC` and `RECOVERY_CONTROL`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewHealthCheck registers a new resource with the given unique name, arguments, and options.
func NewHealthCheck(ctx *pulumi.Context,
	name string, args *HealthCheckArgs, opts ...pulumi.ResourceOption) (*HealthCheck, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource HealthCheck
	err := ctx.RegisterResource("aws:route53/healthCheck:HealthCheck", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHealthCheck gets an existing HealthCheck resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHealthCheck(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HealthCheckState, opts ...pulumi.ResourceOption) (*HealthCheck, error) {
	var resource HealthCheck
	err := ctx.ReadResource("aws:route53/healthCheck:HealthCheck", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HealthCheck resources.
type healthCheckState struct {
	// The Amazon Resource Name (ARN) of the Health Check.
	Arn *string `pulumi:"arn"`
	// The minimum number of child health checks that must be healthy for Route 53 to consider the parent health check to be healthy. Valid values are integers between 0 and 256, inclusive
	ChildHealthThreshold *int `pulumi:"childHealthThreshold"`
	// For a specified parent health check, a list of HealthCheckId values for the associated child health checks.
	ChildHealthchecks []string `pulumi:"childHealthchecks"`
	// The name of the CloudWatch alarm.
	CloudwatchAlarmName *string `pulumi:"cloudwatchAlarmName"`
	// The CloudWatchRegion that the CloudWatch alarm was created in.
	CloudwatchAlarmRegion *string `pulumi:"cloudwatchAlarmRegion"`
	// A boolean value that stops Route 53 from performing health checks. When set to true, Route 53 will do the following depending on the type of health check:
	// * For health checks that check the health of endpoints, Route5 53 stops submitting requests to your application, server, or other resource.
	// * For calculated health checks, Route 53 stops aggregating the status of the referenced health checks.
	// * For health checks that monitor CloudWatch alarms, Route 53 stops monitoring the corresponding CloudWatch metrics.
	Disabled *bool `pulumi:"disabled"`
	// A boolean value that indicates whether Route53 should send the `fqdn` to the endpoint when performing the health check. This defaults to AWS' defaults: when the `type` is "HTTPS" `enableSni` defaults to `true`, when `type` is anything else `enableSni` defaults to `false`.
	EnableSni *bool `pulumi:"enableSni"`
	// The number of consecutive health checks that an endpoint must pass or fail.
	FailureThreshold *int `pulumi:"failureThreshold"`
	// The fully qualified domain name of the endpoint to be checked.
	Fqdn *string `pulumi:"fqdn"`
	// The status of the health check when CloudWatch has insufficient data about the state of associated alarm. Valid values are `Healthy` , `Unhealthy` and `LastKnownStatus`.
	InsufficientDataHealthStatus *string `pulumi:"insufficientDataHealthStatus"`
	// A boolean value that indicates whether the status of health check should be inverted. For example, if a health check is healthy but Inverted is True , then Route 53 considers the health check to be unhealthy.
	InvertHealthcheck *bool `pulumi:"invertHealthcheck"`
	// The IP address of the endpoint to be checked.
	IpAddress *string `pulumi:"ipAddress"`
	// A Boolean value that indicates whether you want Route 53 to measure the latency between health checkers in multiple AWS regions and your endpoint and to display CloudWatch latency graphs in the Route 53 console.
	MeasureLatency *bool `pulumi:"measureLatency"`
	// The port of the endpoint to be checked.
	Port *int `pulumi:"port"`
	// This is a reference name used in Caller Reference
	// (helpful for identifying single healthCheck set amongst others)
	ReferenceName *string `pulumi:"referenceName"`
	// A list of AWS regions that you want Amazon Route 53 health checkers to check the specified endpoint from.
	Regions []string `pulumi:"regions"`
	// The number of seconds between the time that Amazon Route 53 gets a response from your endpoint and the time that it sends the next health-check request.
	RequestInterval *int `pulumi:"requestInterval"`
	// The path that you want Amazon Route 53 to request when performing health checks.
	ResourcePath *string `pulumi:"resourcePath"`
	// The Amazon Resource Name (ARN) for the Route 53 Application Recovery Controller routing control. This is used when health check type is `RECOVERY_CONTROL`
	RoutingControlArn *string `pulumi:"routingControlArn"`
	// String searched in the first 5120 bytes of the response body for check to be considered healthy. Only valid with `HTTP_STR_MATCH` and `HTTPS_STR_MATCH`.
	SearchString *string `pulumi:"searchString"`
	// A map of tags to assign to the health check. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The protocol to use when performing health checks. Valid values are `HTTP`, `HTTPS`, `HTTP_STR_MATCH`, `HTTPS_STR_MATCH`, `TCP`, `CALCULATED`, `CLOUDWATCH_METRIC` and `RECOVERY_CONTROL`.
	Type *string `pulumi:"type"`
}

type HealthCheckState struct {
	// The Amazon Resource Name (ARN) of the Health Check.
	Arn pulumi.StringPtrInput
	// The minimum number of child health checks that must be healthy for Route 53 to consider the parent health check to be healthy. Valid values are integers between 0 and 256, inclusive
	ChildHealthThreshold pulumi.IntPtrInput
	// For a specified parent health check, a list of HealthCheckId values for the associated child health checks.
	ChildHealthchecks pulumi.StringArrayInput
	// The name of the CloudWatch alarm.
	CloudwatchAlarmName pulumi.StringPtrInput
	// The CloudWatchRegion that the CloudWatch alarm was created in.
	CloudwatchAlarmRegion pulumi.StringPtrInput
	// A boolean value that stops Route 53 from performing health checks. When set to true, Route 53 will do the following depending on the type of health check:
	// * For health checks that check the health of endpoints, Route5 53 stops submitting requests to your application, server, or other resource.
	// * For calculated health checks, Route 53 stops aggregating the status of the referenced health checks.
	// * For health checks that monitor CloudWatch alarms, Route 53 stops monitoring the corresponding CloudWatch metrics.
	Disabled pulumi.BoolPtrInput
	// A boolean value that indicates whether Route53 should send the `fqdn` to the endpoint when performing the health check. This defaults to AWS' defaults: when the `type` is "HTTPS" `enableSni` defaults to `true`, when `type` is anything else `enableSni` defaults to `false`.
	EnableSni pulumi.BoolPtrInput
	// The number of consecutive health checks that an endpoint must pass or fail.
	FailureThreshold pulumi.IntPtrInput
	// The fully qualified domain name of the endpoint to be checked.
	Fqdn pulumi.StringPtrInput
	// The status of the health check when CloudWatch has insufficient data about the state of associated alarm. Valid values are `Healthy` , `Unhealthy` and `LastKnownStatus`.
	InsufficientDataHealthStatus pulumi.StringPtrInput
	// A boolean value that indicates whether the status of health check should be inverted. For example, if a health check is healthy but Inverted is True , then Route 53 considers the health check to be unhealthy.
	InvertHealthcheck pulumi.BoolPtrInput
	// The IP address of the endpoint to be checked.
	IpAddress pulumi.StringPtrInput
	// A Boolean value that indicates whether you want Route 53 to measure the latency between health checkers in multiple AWS regions and your endpoint and to display CloudWatch latency graphs in the Route 53 console.
	MeasureLatency pulumi.BoolPtrInput
	// The port of the endpoint to be checked.
	Port pulumi.IntPtrInput
	// This is a reference name used in Caller Reference
	// (helpful for identifying single healthCheck set amongst others)
	ReferenceName pulumi.StringPtrInput
	// A list of AWS regions that you want Amazon Route 53 health checkers to check the specified endpoint from.
	Regions pulumi.StringArrayInput
	// The number of seconds between the time that Amazon Route 53 gets a response from your endpoint and the time that it sends the next health-check request.
	RequestInterval pulumi.IntPtrInput
	// The path that you want Amazon Route 53 to request when performing health checks.
	ResourcePath pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) for the Route 53 Application Recovery Controller routing control. This is used when health check type is `RECOVERY_CONTROL`
	RoutingControlArn pulumi.StringPtrInput
	// String searched in the first 5120 bytes of the response body for check to be considered healthy. Only valid with `HTTP_STR_MATCH` and `HTTPS_STR_MATCH`.
	SearchString pulumi.StringPtrInput
	// A map of tags to assign to the health check. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// The protocol to use when performing health checks. Valid values are `HTTP`, `HTTPS`, `HTTP_STR_MATCH`, `HTTPS_STR_MATCH`, `TCP`, `CALCULATED`, `CLOUDWATCH_METRIC` and `RECOVERY_CONTROL`.
	Type pulumi.StringPtrInput
}

func (HealthCheckState) ElementType() reflect.Type {
	return reflect.TypeOf((*healthCheckState)(nil)).Elem()
}

type healthCheckArgs struct {
	// The minimum number of child health checks that must be healthy for Route 53 to consider the parent health check to be healthy. Valid values are integers between 0 and 256, inclusive
	ChildHealthThreshold *int `pulumi:"childHealthThreshold"`
	// For a specified parent health check, a list of HealthCheckId values for the associated child health checks.
	ChildHealthchecks []string `pulumi:"childHealthchecks"`
	// The name of the CloudWatch alarm.
	CloudwatchAlarmName *string `pulumi:"cloudwatchAlarmName"`
	// The CloudWatchRegion that the CloudWatch alarm was created in.
	CloudwatchAlarmRegion *string `pulumi:"cloudwatchAlarmRegion"`
	// A boolean value that stops Route 53 from performing health checks. When set to true, Route 53 will do the following depending on the type of health check:
	// * For health checks that check the health of endpoints, Route5 53 stops submitting requests to your application, server, or other resource.
	// * For calculated health checks, Route 53 stops aggregating the status of the referenced health checks.
	// * For health checks that monitor CloudWatch alarms, Route 53 stops monitoring the corresponding CloudWatch metrics.
	Disabled *bool `pulumi:"disabled"`
	// A boolean value that indicates whether Route53 should send the `fqdn` to the endpoint when performing the health check. This defaults to AWS' defaults: when the `type` is "HTTPS" `enableSni` defaults to `true`, when `type` is anything else `enableSni` defaults to `false`.
	EnableSni *bool `pulumi:"enableSni"`
	// The number of consecutive health checks that an endpoint must pass or fail.
	FailureThreshold *int `pulumi:"failureThreshold"`
	// The fully qualified domain name of the endpoint to be checked.
	Fqdn *string `pulumi:"fqdn"`
	// The status of the health check when CloudWatch has insufficient data about the state of associated alarm. Valid values are `Healthy` , `Unhealthy` and `LastKnownStatus`.
	InsufficientDataHealthStatus *string `pulumi:"insufficientDataHealthStatus"`
	// A boolean value that indicates whether the status of health check should be inverted. For example, if a health check is healthy but Inverted is True , then Route 53 considers the health check to be unhealthy.
	InvertHealthcheck *bool `pulumi:"invertHealthcheck"`
	// The IP address of the endpoint to be checked.
	IpAddress *string `pulumi:"ipAddress"`
	// A Boolean value that indicates whether you want Route 53 to measure the latency between health checkers in multiple AWS regions and your endpoint and to display CloudWatch latency graphs in the Route 53 console.
	MeasureLatency *bool `pulumi:"measureLatency"`
	// The port of the endpoint to be checked.
	Port *int `pulumi:"port"`
	// This is a reference name used in Caller Reference
	// (helpful for identifying single healthCheck set amongst others)
	ReferenceName *string `pulumi:"referenceName"`
	// A list of AWS regions that you want Amazon Route 53 health checkers to check the specified endpoint from.
	Regions []string `pulumi:"regions"`
	// The number of seconds between the time that Amazon Route 53 gets a response from your endpoint and the time that it sends the next health-check request.
	RequestInterval *int `pulumi:"requestInterval"`
	// The path that you want Amazon Route 53 to request when performing health checks.
	ResourcePath *string `pulumi:"resourcePath"`
	// The Amazon Resource Name (ARN) for the Route 53 Application Recovery Controller routing control. This is used when health check type is `RECOVERY_CONTROL`
	RoutingControlArn *string `pulumi:"routingControlArn"`
	// String searched in the first 5120 bytes of the response body for check to be considered healthy. Only valid with `HTTP_STR_MATCH` and `HTTPS_STR_MATCH`.
	SearchString *string `pulumi:"searchString"`
	// A map of tags to assign to the health check. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The protocol to use when performing health checks. Valid values are `HTTP`, `HTTPS`, `HTTP_STR_MATCH`, `HTTPS_STR_MATCH`, `TCP`, `CALCULATED`, `CLOUDWATCH_METRIC` and `RECOVERY_CONTROL`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a HealthCheck resource.
type HealthCheckArgs struct {
	// The minimum number of child health checks that must be healthy for Route 53 to consider the parent health check to be healthy. Valid values are integers between 0 and 256, inclusive
	ChildHealthThreshold pulumi.IntPtrInput
	// For a specified parent health check, a list of HealthCheckId values for the associated child health checks.
	ChildHealthchecks pulumi.StringArrayInput
	// The name of the CloudWatch alarm.
	CloudwatchAlarmName pulumi.StringPtrInput
	// The CloudWatchRegion that the CloudWatch alarm was created in.
	CloudwatchAlarmRegion pulumi.StringPtrInput
	// A boolean value that stops Route 53 from performing health checks. When set to true, Route 53 will do the following depending on the type of health check:
	// * For health checks that check the health of endpoints, Route5 53 stops submitting requests to your application, server, or other resource.
	// * For calculated health checks, Route 53 stops aggregating the status of the referenced health checks.
	// * For health checks that monitor CloudWatch alarms, Route 53 stops monitoring the corresponding CloudWatch metrics.
	Disabled pulumi.BoolPtrInput
	// A boolean value that indicates whether Route53 should send the `fqdn` to the endpoint when performing the health check. This defaults to AWS' defaults: when the `type` is "HTTPS" `enableSni` defaults to `true`, when `type` is anything else `enableSni` defaults to `false`.
	EnableSni pulumi.BoolPtrInput
	// The number of consecutive health checks that an endpoint must pass or fail.
	FailureThreshold pulumi.IntPtrInput
	// The fully qualified domain name of the endpoint to be checked.
	Fqdn pulumi.StringPtrInput
	// The status of the health check when CloudWatch has insufficient data about the state of associated alarm. Valid values are `Healthy` , `Unhealthy` and `LastKnownStatus`.
	InsufficientDataHealthStatus pulumi.StringPtrInput
	// A boolean value that indicates whether the status of health check should be inverted. For example, if a health check is healthy but Inverted is True , then Route 53 considers the health check to be unhealthy.
	InvertHealthcheck pulumi.BoolPtrInput
	// The IP address of the endpoint to be checked.
	IpAddress pulumi.StringPtrInput
	// A Boolean value that indicates whether you want Route 53 to measure the latency between health checkers in multiple AWS regions and your endpoint and to display CloudWatch latency graphs in the Route 53 console.
	MeasureLatency pulumi.BoolPtrInput
	// The port of the endpoint to be checked.
	Port pulumi.IntPtrInput
	// This is a reference name used in Caller Reference
	// (helpful for identifying single healthCheck set amongst others)
	ReferenceName pulumi.StringPtrInput
	// A list of AWS regions that you want Amazon Route 53 health checkers to check the specified endpoint from.
	Regions pulumi.StringArrayInput
	// The number of seconds between the time that Amazon Route 53 gets a response from your endpoint and the time that it sends the next health-check request.
	RequestInterval pulumi.IntPtrInput
	// The path that you want Amazon Route 53 to request when performing health checks.
	ResourcePath pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) for the Route 53 Application Recovery Controller routing control. This is used when health check type is `RECOVERY_CONTROL`
	RoutingControlArn pulumi.StringPtrInput
	// String searched in the first 5120 bytes of the response body for check to be considered healthy. Only valid with `HTTP_STR_MATCH` and `HTTPS_STR_MATCH`.
	SearchString pulumi.StringPtrInput
	// A map of tags to assign to the health check. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The protocol to use when performing health checks. Valid values are `HTTP`, `HTTPS`, `HTTP_STR_MATCH`, `HTTPS_STR_MATCH`, `TCP`, `CALCULATED`, `CLOUDWATCH_METRIC` and `RECOVERY_CONTROL`.
	Type pulumi.StringInput
}

func (HealthCheckArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*healthCheckArgs)(nil)).Elem()
}

type HealthCheckInput interface {
	pulumi.Input

	ToHealthCheckOutput() HealthCheckOutput
	ToHealthCheckOutputWithContext(ctx context.Context) HealthCheckOutput
}

func (*HealthCheck) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthCheck)(nil)).Elem()
}

func (i *HealthCheck) ToHealthCheckOutput() HealthCheckOutput {
	return i.ToHealthCheckOutputWithContext(context.Background())
}

func (i *HealthCheck) ToHealthCheckOutputWithContext(ctx context.Context) HealthCheckOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthCheckOutput)
}

// HealthCheckArrayInput is an input type that accepts HealthCheckArray and HealthCheckArrayOutput values.
// You can construct a concrete instance of `HealthCheckArrayInput` via:
//
//	HealthCheckArray{ HealthCheckArgs{...} }
type HealthCheckArrayInput interface {
	pulumi.Input

	ToHealthCheckArrayOutput() HealthCheckArrayOutput
	ToHealthCheckArrayOutputWithContext(context.Context) HealthCheckArrayOutput
}

type HealthCheckArray []HealthCheckInput

func (HealthCheckArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HealthCheck)(nil)).Elem()
}

func (i HealthCheckArray) ToHealthCheckArrayOutput() HealthCheckArrayOutput {
	return i.ToHealthCheckArrayOutputWithContext(context.Background())
}

func (i HealthCheckArray) ToHealthCheckArrayOutputWithContext(ctx context.Context) HealthCheckArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthCheckArrayOutput)
}

// HealthCheckMapInput is an input type that accepts HealthCheckMap and HealthCheckMapOutput values.
// You can construct a concrete instance of `HealthCheckMapInput` via:
//
//	HealthCheckMap{ "key": HealthCheckArgs{...} }
type HealthCheckMapInput interface {
	pulumi.Input

	ToHealthCheckMapOutput() HealthCheckMapOutput
	ToHealthCheckMapOutputWithContext(context.Context) HealthCheckMapOutput
}

type HealthCheckMap map[string]HealthCheckInput

func (HealthCheckMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HealthCheck)(nil)).Elem()
}

func (i HealthCheckMap) ToHealthCheckMapOutput() HealthCheckMapOutput {
	return i.ToHealthCheckMapOutputWithContext(context.Background())
}

func (i HealthCheckMap) ToHealthCheckMapOutputWithContext(ctx context.Context) HealthCheckMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthCheckMapOutput)
}

type HealthCheckOutput struct{ *pulumi.OutputState }

func (HealthCheckOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthCheck)(nil)).Elem()
}

func (o HealthCheckOutput) ToHealthCheckOutput() HealthCheckOutput {
	return o
}

func (o HealthCheckOutput) ToHealthCheckOutputWithContext(ctx context.Context) HealthCheckOutput {
	return o
}

// The Amazon Resource Name (ARN) of the Health Check.
func (o HealthCheckOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthCheck) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The minimum number of child health checks that must be healthy for Route 53 to consider the parent health check to be healthy. Valid values are integers between 0 and 256, inclusive
func (o HealthCheckOutput) ChildHealthThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HealthCheck) pulumi.IntPtrOutput { return v.ChildHealthThreshold }).(pulumi.IntPtrOutput)
}

// For a specified parent health check, a list of HealthCheckId values for the associated child health checks.
func (o HealthCheckOutput) ChildHealthchecks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HealthCheck) pulumi.StringArrayOutput { return v.ChildHealthchecks }).(pulumi.StringArrayOutput)
}

// The name of the CloudWatch alarm.
func (o HealthCheckOutput) CloudwatchAlarmName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HealthCheck) pulumi.StringPtrOutput { return v.CloudwatchAlarmName }).(pulumi.StringPtrOutput)
}

// The CloudWatchRegion that the CloudWatch alarm was created in.
func (o HealthCheckOutput) CloudwatchAlarmRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HealthCheck) pulumi.StringPtrOutput { return v.CloudwatchAlarmRegion }).(pulumi.StringPtrOutput)
}

// A boolean value that stops Route 53 from performing health checks. When set to true, Route 53 will do the following depending on the type of health check:
// * For health checks that check the health of endpoints, Route5 53 stops submitting requests to your application, server, or other resource.
// * For calculated health checks, Route 53 stops aggregating the status of the referenced health checks.
// * For health checks that monitor CloudWatch alarms, Route 53 stops monitoring the corresponding CloudWatch metrics.
func (o HealthCheckOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HealthCheck) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// A boolean value that indicates whether Route53 should send the `fqdn` to the endpoint when performing the health check. This defaults to AWS' defaults: when the `type` is "HTTPS" `enableSni` defaults to `true`, when `type` is anything else `enableSni` defaults to `false`.
func (o HealthCheckOutput) EnableSni() pulumi.BoolOutput {
	return o.ApplyT(func(v *HealthCheck) pulumi.BoolOutput { return v.EnableSni }).(pulumi.BoolOutput)
}

// The number of consecutive health checks that an endpoint must pass or fail.
func (o HealthCheckOutput) FailureThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *HealthCheck) pulumi.IntOutput { return v.FailureThreshold }).(pulumi.IntOutput)
}

// The fully qualified domain name of the endpoint to be checked.
func (o HealthCheckOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HealthCheck) pulumi.StringPtrOutput { return v.Fqdn }).(pulumi.StringPtrOutput)
}

// The status of the health check when CloudWatch has insufficient data about the state of associated alarm. Valid values are `Healthy` , `Unhealthy` and `LastKnownStatus`.
func (o HealthCheckOutput) InsufficientDataHealthStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HealthCheck) pulumi.StringPtrOutput { return v.InsufficientDataHealthStatus }).(pulumi.StringPtrOutput)
}

// A boolean value that indicates whether the status of health check should be inverted. For example, if a health check is healthy but Inverted is True , then Route 53 considers the health check to be unhealthy.
func (o HealthCheckOutput) InvertHealthcheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HealthCheck) pulumi.BoolPtrOutput { return v.InvertHealthcheck }).(pulumi.BoolPtrOutput)
}

// The IP address of the endpoint to be checked.
func (o HealthCheckOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HealthCheck) pulumi.StringPtrOutput { return v.IpAddress }).(pulumi.StringPtrOutput)
}

// A Boolean value that indicates whether you want Route 53 to measure the latency between health checkers in multiple AWS regions and your endpoint and to display CloudWatch latency graphs in the Route 53 console.
func (o HealthCheckOutput) MeasureLatency() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HealthCheck) pulumi.BoolPtrOutput { return v.MeasureLatency }).(pulumi.BoolPtrOutput)
}

// The port of the endpoint to be checked.
func (o HealthCheckOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HealthCheck) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

// This is a reference name used in Caller Reference
// (helpful for identifying single healthCheck set amongst others)
func (o HealthCheckOutput) ReferenceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HealthCheck) pulumi.StringPtrOutput { return v.ReferenceName }).(pulumi.StringPtrOutput)
}

// A list of AWS regions that you want Amazon Route 53 health checkers to check the specified endpoint from.
func (o HealthCheckOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HealthCheck) pulumi.StringArrayOutput { return v.Regions }).(pulumi.StringArrayOutput)
}

// The number of seconds between the time that Amazon Route 53 gets a response from your endpoint and the time that it sends the next health-check request.
func (o HealthCheckOutput) RequestInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HealthCheck) pulumi.IntPtrOutput { return v.RequestInterval }).(pulumi.IntPtrOutput)
}

// The path that you want Amazon Route 53 to request when performing health checks.
func (o HealthCheckOutput) ResourcePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HealthCheck) pulumi.StringPtrOutput { return v.ResourcePath }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) for the Route 53 Application Recovery Controller routing control. This is used when health check type is `RECOVERY_CONTROL`
func (o HealthCheckOutput) RoutingControlArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HealthCheck) pulumi.StringPtrOutput { return v.RoutingControlArn }).(pulumi.StringPtrOutput)
}

// String searched in the first 5120 bytes of the response body for check to be considered healthy. Only valid with `HTTP_STR_MATCH` and `HTTPS_STR_MATCH`.
func (o HealthCheckOutput) SearchString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HealthCheck) pulumi.StringPtrOutput { return v.SearchString }).(pulumi.StringPtrOutput)
}

// A map of tags to assign to the health check. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o HealthCheckOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *HealthCheck) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o HealthCheckOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *HealthCheck) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The protocol to use when performing health checks. Valid values are `HTTP`, `HTTPS`, `HTTP_STR_MATCH`, `HTTPS_STR_MATCH`, `TCP`, `CALCULATED`, `CLOUDWATCH_METRIC` and `RECOVERY_CONTROL`.
func (o HealthCheckOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HealthCheck) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type HealthCheckArrayOutput struct{ *pulumi.OutputState }

func (HealthCheckArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HealthCheck)(nil)).Elem()
}

func (o HealthCheckArrayOutput) ToHealthCheckArrayOutput() HealthCheckArrayOutput {
	return o
}

func (o HealthCheckArrayOutput) ToHealthCheckArrayOutputWithContext(ctx context.Context) HealthCheckArrayOutput {
	return o
}

func (o HealthCheckArrayOutput) Index(i pulumi.IntInput) HealthCheckOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HealthCheck {
		return vs[0].([]*HealthCheck)[vs[1].(int)]
	}).(HealthCheckOutput)
}

type HealthCheckMapOutput struct{ *pulumi.OutputState }

func (HealthCheckMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HealthCheck)(nil)).Elem()
}

func (o HealthCheckMapOutput) ToHealthCheckMapOutput() HealthCheckMapOutput {
	return o
}

func (o HealthCheckMapOutput) ToHealthCheckMapOutputWithContext(ctx context.Context) HealthCheckMapOutput {
	return o
}

func (o HealthCheckMapOutput) MapIndex(k pulumi.StringInput) HealthCheckOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HealthCheck {
		return vs[0].(map[string]*HealthCheck)[vs[1].(string)]
	}).(HealthCheckOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HealthCheckInput)(nil)).Elem(), &HealthCheck{})
	pulumi.RegisterInputType(reflect.TypeOf((*HealthCheckArrayInput)(nil)).Elem(), HealthCheckArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HealthCheckMapInput)(nil)).Elem(), HealthCheckMap{})
	pulumi.RegisterOutputType(HealthCheckOutput{})
	pulumi.RegisterOutputType(HealthCheckArrayOutput{})
	pulumi.RegisterOutputType(HealthCheckMapOutput{})
}
