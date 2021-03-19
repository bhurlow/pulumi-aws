// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package globalaccelerator

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Global Accelerator endpoint group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/globalaccelerator"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := globalaccelerator.NewEndpointGroup(ctx, "example", &globalaccelerator.EndpointGroupArgs{
// 			ListenerArn: pulumi.Any(aws_globalaccelerator_listener.Example.Id),
// 			EndpointConfigurations: globalaccelerator.EndpointGroupEndpointConfigurationArray{
// 				&globalaccelerator.EndpointGroupEndpointConfigurationArgs{
// 					EndpointId: pulumi.Any(aws_lb.Example.Arn),
// 					Weight:     pulumi.Int(100),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Global Accelerator endpoint groups can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:globalaccelerator/endpointGroup:EndpointGroup example arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/listener/xxxxxxx/endpoint-group/xxxxxxxx
// ```
type EndpointGroup struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the endpoint group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The list of endpoint objects. Fields documented below.
	EndpointConfigurations EndpointGroupEndpointConfigurationArrayOutput `pulumi:"endpointConfigurations"`
	// The name of the AWS Region where the endpoint group is located.
	EndpointGroupRegion pulumi.StringOutput `pulumi:"endpointGroupRegion"`
	// The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
	HealthCheckIntervalSeconds pulumi.IntPtrOutput `pulumi:"healthCheckIntervalSeconds"`
	// If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (`/`). the provider will only perform drift detection of its value when present in a configuration.
	HealthCheckPath pulumi.StringOutput `pulumi:"healthCheckPath"`
	// The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
	// the provider will only perform drift detection of its value when present in a configuration.
	HealthCheckPort pulumi.IntOutput `pulumi:"healthCheckPort"`
	// The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
	HealthCheckProtocol pulumi.StringPtrOutput `pulumi:"healthCheckProtocol"`
	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn pulumi.StringOutput `pulumi:"listenerArn"`
	// Override specific listener ports used to route traffic to endpoints that are part of this endpoint group. Fields documented below.
	PortOverrides EndpointGroupPortOverrideArrayOutput `pulumi:"portOverrides"`
	// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
	ThresholdCount pulumi.IntPtrOutput `pulumi:"thresholdCount"`
	// The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
	TrafficDialPercentage pulumi.Float64PtrOutput `pulumi:"trafficDialPercentage"`
}

// NewEndpointGroup registers a new resource with the given unique name, arguments, and options.
func NewEndpointGroup(ctx *pulumi.Context,
	name string, args *EndpointGroupArgs, opts ...pulumi.ResourceOption) (*EndpointGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ListenerArn == nil {
		return nil, errors.New("invalid value for required argument 'ListenerArn'")
	}
	var resource EndpointGroup
	err := ctx.RegisterResource("aws:globalaccelerator/endpointGroup:EndpointGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointGroup gets an existing EndpointGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointGroupState, opts ...pulumi.ResourceOption) (*EndpointGroup, error) {
	var resource EndpointGroup
	err := ctx.ReadResource("aws:globalaccelerator/endpointGroup:EndpointGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointGroup resources.
type endpointGroupState struct {
	// The Amazon Resource Name (ARN) of the endpoint group.
	Arn *string `pulumi:"arn"`
	// The list of endpoint objects. Fields documented below.
	EndpointConfigurations []EndpointGroupEndpointConfiguration `pulumi:"endpointConfigurations"`
	// The name of the AWS Region where the endpoint group is located.
	EndpointGroupRegion *string `pulumi:"endpointGroupRegion"`
	// The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
	HealthCheckIntervalSeconds *int `pulumi:"healthCheckIntervalSeconds"`
	// If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (`/`). the provider will only perform drift detection of its value when present in a configuration.
	HealthCheckPath *string `pulumi:"healthCheckPath"`
	// The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
	// the provider will only perform drift detection of its value when present in a configuration.
	HealthCheckPort *int `pulumi:"healthCheckPort"`
	// The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
	HealthCheckProtocol *string `pulumi:"healthCheckProtocol"`
	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn *string `pulumi:"listenerArn"`
	// Override specific listener ports used to route traffic to endpoints that are part of this endpoint group. Fields documented below.
	PortOverrides []EndpointGroupPortOverride `pulumi:"portOverrides"`
	// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
	ThresholdCount *int `pulumi:"thresholdCount"`
	// The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
	TrafficDialPercentage *float64 `pulumi:"trafficDialPercentage"`
}

type EndpointGroupState struct {
	// The Amazon Resource Name (ARN) of the endpoint group.
	Arn pulumi.StringPtrInput
	// The list of endpoint objects. Fields documented below.
	EndpointConfigurations EndpointGroupEndpointConfigurationArrayInput
	// The name of the AWS Region where the endpoint group is located.
	EndpointGroupRegion pulumi.StringPtrInput
	// The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
	HealthCheckIntervalSeconds pulumi.IntPtrInput
	// If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (`/`). the provider will only perform drift detection of its value when present in a configuration.
	HealthCheckPath pulumi.StringPtrInput
	// The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
	// the provider will only perform drift detection of its value when present in a configuration.
	HealthCheckPort pulumi.IntPtrInput
	// The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
	HealthCheckProtocol pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn pulumi.StringPtrInput
	// Override specific listener ports used to route traffic to endpoints that are part of this endpoint group. Fields documented below.
	PortOverrides EndpointGroupPortOverrideArrayInput
	// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
	ThresholdCount pulumi.IntPtrInput
	// The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
	TrafficDialPercentage pulumi.Float64PtrInput
}

func (EndpointGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointGroupState)(nil)).Elem()
}

type endpointGroupArgs struct {
	// The list of endpoint objects. Fields documented below.
	EndpointConfigurations []EndpointGroupEndpointConfiguration `pulumi:"endpointConfigurations"`
	// The name of the AWS Region where the endpoint group is located.
	EndpointGroupRegion *string `pulumi:"endpointGroupRegion"`
	// The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
	HealthCheckIntervalSeconds *int `pulumi:"healthCheckIntervalSeconds"`
	// If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (`/`). the provider will only perform drift detection of its value when present in a configuration.
	HealthCheckPath *string `pulumi:"healthCheckPath"`
	// The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
	// the provider will only perform drift detection of its value when present in a configuration.
	HealthCheckPort *int `pulumi:"healthCheckPort"`
	// The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
	HealthCheckProtocol *string `pulumi:"healthCheckProtocol"`
	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn string `pulumi:"listenerArn"`
	// Override specific listener ports used to route traffic to endpoints that are part of this endpoint group. Fields documented below.
	PortOverrides []EndpointGroupPortOverride `pulumi:"portOverrides"`
	// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
	ThresholdCount *int `pulumi:"thresholdCount"`
	// The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
	TrafficDialPercentage *float64 `pulumi:"trafficDialPercentage"`
}

// The set of arguments for constructing a EndpointGroup resource.
type EndpointGroupArgs struct {
	// The list of endpoint objects. Fields documented below.
	EndpointConfigurations EndpointGroupEndpointConfigurationArrayInput
	// The name of the AWS Region where the endpoint group is located.
	EndpointGroupRegion pulumi.StringPtrInput
	// The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
	HealthCheckIntervalSeconds pulumi.IntPtrInput
	// If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (`/`). the provider will only perform drift detection of its value when present in a configuration.
	HealthCheckPath pulumi.StringPtrInput
	// The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
	// the provider will only perform drift detection of its value when present in a configuration.
	HealthCheckPort pulumi.IntPtrInput
	// The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
	HealthCheckProtocol pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn pulumi.StringInput
	// Override specific listener ports used to route traffic to endpoints that are part of this endpoint group. Fields documented below.
	PortOverrides EndpointGroupPortOverrideArrayInput
	// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
	ThresholdCount pulumi.IntPtrInput
	// The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
	TrafficDialPercentage pulumi.Float64PtrInput
}

func (EndpointGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointGroupArgs)(nil)).Elem()
}

type EndpointGroupInput interface {
	pulumi.Input

	ToEndpointGroupOutput() EndpointGroupOutput
	ToEndpointGroupOutputWithContext(ctx context.Context) EndpointGroupOutput
}

func (*EndpointGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointGroup)(nil))
}

func (i *EndpointGroup) ToEndpointGroupOutput() EndpointGroupOutput {
	return i.ToEndpointGroupOutputWithContext(context.Background())
}

func (i *EndpointGroup) ToEndpointGroupOutputWithContext(ctx context.Context) EndpointGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointGroupOutput)
}

func (i *EndpointGroup) ToEndpointGroupPtrOutput() EndpointGroupPtrOutput {
	return i.ToEndpointGroupPtrOutputWithContext(context.Background())
}

func (i *EndpointGroup) ToEndpointGroupPtrOutputWithContext(ctx context.Context) EndpointGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointGroupPtrOutput)
}

type EndpointGroupPtrInput interface {
	pulumi.Input

	ToEndpointGroupPtrOutput() EndpointGroupPtrOutput
	ToEndpointGroupPtrOutputWithContext(ctx context.Context) EndpointGroupPtrOutput
}

type endpointGroupPtrType EndpointGroupArgs

func (*endpointGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointGroup)(nil))
}

func (i *endpointGroupPtrType) ToEndpointGroupPtrOutput() EndpointGroupPtrOutput {
	return i.ToEndpointGroupPtrOutputWithContext(context.Background())
}

func (i *endpointGroupPtrType) ToEndpointGroupPtrOutputWithContext(ctx context.Context) EndpointGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointGroupPtrOutput)
}

// EndpointGroupArrayInput is an input type that accepts EndpointGroupArray and EndpointGroupArrayOutput values.
// You can construct a concrete instance of `EndpointGroupArrayInput` via:
//
//          EndpointGroupArray{ EndpointGroupArgs{...} }
type EndpointGroupArrayInput interface {
	pulumi.Input

	ToEndpointGroupArrayOutput() EndpointGroupArrayOutput
	ToEndpointGroupArrayOutputWithContext(context.Context) EndpointGroupArrayOutput
}

type EndpointGroupArray []EndpointGroupInput

func (EndpointGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*EndpointGroup)(nil))
}

func (i EndpointGroupArray) ToEndpointGroupArrayOutput() EndpointGroupArrayOutput {
	return i.ToEndpointGroupArrayOutputWithContext(context.Background())
}

func (i EndpointGroupArray) ToEndpointGroupArrayOutputWithContext(ctx context.Context) EndpointGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointGroupArrayOutput)
}

// EndpointGroupMapInput is an input type that accepts EndpointGroupMap and EndpointGroupMapOutput values.
// You can construct a concrete instance of `EndpointGroupMapInput` via:
//
//          EndpointGroupMap{ "key": EndpointGroupArgs{...} }
type EndpointGroupMapInput interface {
	pulumi.Input

	ToEndpointGroupMapOutput() EndpointGroupMapOutput
	ToEndpointGroupMapOutputWithContext(context.Context) EndpointGroupMapOutput
}

type EndpointGroupMap map[string]EndpointGroupInput

func (EndpointGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*EndpointGroup)(nil))
}

func (i EndpointGroupMap) ToEndpointGroupMapOutput() EndpointGroupMapOutput {
	return i.ToEndpointGroupMapOutputWithContext(context.Background())
}

func (i EndpointGroupMap) ToEndpointGroupMapOutputWithContext(ctx context.Context) EndpointGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointGroupMapOutput)
}

type EndpointGroupOutput struct {
	*pulumi.OutputState
}

func (EndpointGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointGroup)(nil))
}

func (o EndpointGroupOutput) ToEndpointGroupOutput() EndpointGroupOutput {
	return o
}

func (o EndpointGroupOutput) ToEndpointGroupOutputWithContext(ctx context.Context) EndpointGroupOutput {
	return o
}

func (o EndpointGroupOutput) ToEndpointGroupPtrOutput() EndpointGroupPtrOutput {
	return o.ToEndpointGroupPtrOutputWithContext(context.Background())
}

func (o EndpointGroupOutput) ToEndpointGroupPtrOutputWithContext(ctx context.Context) EndpointGroupPtrOutput {
	return o.ApplyT(func(v EndpointGroup) *EndpointGroup {
		return &v
	}).(EndpointGroupPtrOutput)
}

type EndpointGroupPtrOutput struct {
	*pulumi.OutputState
}

func (EndpointGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointGroup)(nil))
}

func (o EndpointGroupPtrOutput) ToEndpointGroupPtrOutput() EndpointGroupPtrOutput {
	return o
}

func (o EndpointGroupPtrOutput) ToEndpointGroupPtrOutputWithContext(ctx context.Context) EndpointGroupPtrOutput {
	return o
}

type EndpointGroupArrayOutput struct{ *pulumi.OutputState }

func (EndpointGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointGroup)(nil))
}

func (o EndpointGroupArrayOutput) ToEndpointGroupArrayOutput() EndpointGroupArrayOutput {
	return o
}

func (o EndpointGroupArrayOutput) ToEndpointGroupArrayOutputWithContext(ctx context.Context) EndpointGroupArrayOutput {
	return o
}

func (o EndpointGroupArrayOutput) Index(i pulumi.IntInput) EndpointGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointGroup {
		return vs[0].([]EndpointGroup)[vs[1].(int)]
	}).(EndpointGroupOutput)
}

type EndpointGroupMapOutput struct{ *pulumi.OutputState }

func (EndpointGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EndpointGroup)(nil))
}

func (o EndpointGroupMapOutput) ToEndpointGroupMapOutput() EndpointGroupMapOutput {
	return o
}

func (o EndpointGroupMapOutput) ToEndpointGroupMapOutputWithContext(ctx context.Context) EndpointGroupMapOutput {
	return o
}

func (o EndpointGroupMapOutput) MapIndex(k pulumi.StringInput) EndpointGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EndpointGroup {
		return vs[0].(map[string]EndpointGroup)[vs[1].(string)]
	}).(EndpointGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(EndpointGroupOutput{})
	pulumi.RegisterOutputType(EndpointGroupPtrOutput{})
	pulumi.RegisterOutputType(EndpointGroupArrayOutput{})
	pulumi.RegisterOutputType(EndpointGroupMapOutput{})
}
