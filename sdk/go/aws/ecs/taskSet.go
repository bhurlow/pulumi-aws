// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides an ECS task set - effectively a task that is expected to run until an error occurs or a user terminates it (typically a webserver or a database).
//
// See [ECS Task Set section in AWS developer guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-external.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ecs.NewTaskSet(ctx, "example", &ecs.TaskSetArgs{
//				Service:        pulumi.Any(aws_ecs_service.Example.Id),
//				Cluster:        pulumi.Any(aws_ecs_cluster.Example.Id),
//				TaskDefinition: pulumi.Any(aws_ecs_task_definition.Example.Arn),
//				LoadBalancers: ecs.TaskSetLoadBalancerArray{
//					&ecs.TaskSetLoadBalancerArgs{
//						TargetGroupArn: pulumi.Any(aws_lb_target_group.Example.Arn),
//						ContainerName:  pulumi.String("mongo"),
//						ContainerPort:  pulumi.Int(8080),
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
// Using `pulumi import`, import ECS Task Sets using the `task_set_id`, `service`, and `cluster` separated by commas (`,`). For example:
//
// ```sh
//
//	$ pulumi import aws:ecs/taskSet:TaskSet example ecs-svc/7177320696926227436,arn:aws:ecs:us-west-2:123456789101:service/example/example-1234567890,arn:aws:ecs:us-west-2:123456789101:cluster/example
//
// ```
type TaskSet struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) that identifies the task set.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The capacity provider strategy to use for the service. Can be one or more.  Defined below.
	CapacityProviderStrategies TaskSetCapacityProviderStrategyArrayOutput `pulumi:"capacityProviderStrategies"`
	// The short name or ARN of the cluster that hosts the service to create the task set in.
	Cluster pulumi.StringOutput `pulumi:"cluster"`
	// The external ID associated with the task set.
	ExternalId pulumi.StringOutput `pulumi:"externalId"`
	// Whether to allow deleting the task set without waiting for scaling down to 0. You can force a task set to delete even if it's in the process of scaling a resource. Normally, the provider drains all the tasks before deleting the task set. This bypasses that behavior and potentially leaves resources dangling.
	ForceDelete pulumi.BoolPtrOutput `pulumi:"forceDelete"`
	// The launch type on which to run your service. The valid values are `EC2`, `FARGATE`, and `EXTERNAL`. Defaults to `EC2`.
	LaunchType pulumi.StringOutput `pulumi:"launchType"`
	// Details on load balancers that are used with a task set. Detailed below.
	LoadBalancers TaskSetLoadBalancerArrayOutput `pulumi:"loadBalancers"`
	// The network configuration for the service. This parameter is required for task definitions that use the `awsvpc` network mode to receive their own Elastic Network Interface, and it is not supported for other network modes. Detailed below.
	NetworkConfiguration TaskSetNetworkConfigurationPtrOutput `pulumi:"networkConfiguration"`
	// The platform version on which to run your service. Only applicable for `launchType` set to `FARGATE`. Defaults to `LATEST`. More information about Fargate platform versions can be found in the [AWS ECS User Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
	PlatformVersion pulumi.StringOutput `pulumi:"platformVersion"`
	// A floating-point percentage of the desired number of tasks to place and keep running in the task set. Detailed below.
	Scale TaskSetScaleOutput `pulumi:"scale"`
	// The short name or ARN of the ECS service.
	Service pulumi.StringOutput `pulumi:"service"`
	// The service discovery registries for the service. The maximum number of `serviceRegistries` blocks is `1`. Detailed below.
	ServiceRegistries TaskSetServiceRegistriesPtrOutput `pulumi:"serviceRegistries"`
	// The stability status. This indicates whether the task set has reached a steady state.
	StabilityStatus pulumi.StringOutput `pulumi:"stabilityStatus"`
	// The status of the task set.
	Status pulumi.StringOutput `pulumi:"status"`
	// A map of tags to assign to the file system. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level. If you have set `copyTagsToBackups` to true, and you specify one or more tags, no existing file system tags are copied from the file system to the backup.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The family and revision (`family:revision`) or full ARN of the task definition that you want to run in your service.
	//
	// The following arguments are optional:
	TaskDefinition pulumi.StringOutput `pulumi:"taskDefinition"`
	// The ID of the task set.
	TaskSetId pulumi.StringOutput `pulumi:"taskSetId"`
	// Whether the provider should wait until the task set has reached `STEADY_STATE`.
	WaitUntilStable pulumi.BoolPtrOutput `pulumi:"waitUntilStable"`
	// Wait timeout for task set to reach `STEADY_STATE`. Valid time units include `ns`, `us` (or `µs`), `ms`, `s`, `m`, and `h`. Default `10m`.
	WaitUntilStableTimeout pulumi.StringPtrOutput `pulumi:"waitUntilStableTimeout"`
}

// NewTaskSet registers a new resource with the given unique name, arguments, and options.
func NewTaskSet(ctx *pulumi.Context,
	name string, args *TaskSetArgs, opts ...pulumi.ResourceOption) (*TaskSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cluster == nil {
		return nil, errors.New("invalid value for required argument 'Cluster'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	if args.TaskDefinition == nil {
		return nil, errors.New("invalid value for required argument 'TaskDefinition'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TaskSet
	err := ctx.RegisterResource("aws:ecs/taskSet:TaskSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTaskSet gets an existing TaskSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTaskSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaskSetState, opts ...pulumi.ResourceOption) (*TaskSet, error) {
	var resource TaskSet
	err := ctx.ReadResource("aws:ecs/taskSet:TaskSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TaskSet resources.
type taskSetState struct {
	// The Amazon Resource Name (ARN) that identifies the task set.
	Arn *string `pulumi:"arn"`
	// The capacity provider strategy to use for the service. Can be one or more.  Defined below.
	CapacityProviderStrategies []TaskSetCapacityProviderStrategy `pulumi:"capacityProviderStrategies"`
	// The short name or ARN of the cluster that hosts the service to create the task set in.
	Cluster *string `pulumi:"cluster"`
	// The external ID associated with the task set.
	ExternalId *string `pulumi:"externalId"`
	// Whether to allow deleting the task set without waiting for scaling down to 0. You can force a task set to delete even if it's in the process of scaling a resource. Normally, the provider drains all the tasks before deleting the task set. This bypasses that behavior and potentially leaves resources dangling.
	ForceDelete *bool `pulumi:"forceDelete"`
	// The launch type on which to run your service. The valid values are `EC2`, `FARGATE`, and `EXTERNAL`. Defaults to `EC2`.
	LaunchType *string `pulumi:"launchType"`
	// Details on load balancers that are used with a task set. Detailed below.
	LoadBalancers []TaskSetLoadBalancer `pulumi:"loadBalancers"`
	// The network configuration for the service. This parameter is required for task definitions that use the `awsvpc` network mode to receive their own Elastic Network Interface, and it is not supported for other network modes. Detailed below.
	NetworkConfiguration *TaskSetNetworkConfiguration `pulumi:"networkConfiguration"`
	// The platform version on which to run your service. Only applicable for `launchType` set to `FARGATE`. Defaults to `LATEST`. More information about Fargate platform versions can be found in the [AWS ECS User Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
	PlatformVersion *string `pulumi:"platformVersion"`
	// A floating-point percentage of the desired number of tasks to place and keep running in the task set. Detailed below.
	Scale *TaskSetScale `pulumi:"scale"`
	// The short name or ARN of the ECS service.
	Service *string `pulumi:"service"`
	// The service discovery registries for the service. The maximum number of `serviceRegistries` blocks is `1`. Detailed below.
	ServiceRegistries *TaskSetServiceRegistries `pulumi:"serviceRegistries"`
	// The stability status. This indicates whether the task set has reached a steady state.
	StabilityStatus *string `pulumi:"stabilityStatus"`
	// The status of the task set.
	Status *string `pulumi:"status"`
	// A map of tags to assign to the file system. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level. If you have set `copyTagsToBackups` to true, and you specify one or more tags, no existing file system tags are copied from the file system to the backup.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The family and revision (`family:revision`) or full ARN of the task definition that you want to run in your service.
	//
	// The following arguments are optional:
	TaskDefinition *string `pulumi:"taskDefinition"`
	// The ID of the task set.
	TaskSetId *string `pulumi:"taskSetId"`
	// Whether the provider should wait until the task set has reached `STEADY_STATE`.
	WaitUntilStable *bool `pulumi:"waitUntilStable"`
	// Wait timeout for task set to reach `STEADY_STATE`. Valid time units include `ns`, `us` (or `µs`), `ms`, `s`, `m`, and `h`. Default `10m`.
	WaitUntilStableTimeout *string `pulumi:"waitUntilStableTimeout"`
}

type TaskSetState struct {
	// The Amazon Resource Name (ARN) that identifies the task set.
	Arn pulumi.StringPtrInput
	// The capacity provider strategy to use for the service. Can be one or more.  Defined below.
	CapacityProviderStrategies TaskSetCapacityProviderStrategyArrayInput
	// The short name or ARN of the cluster that hosts the service to create the task set in.
	Cluster pulumi.StringPtrInput
	// The external ID associated with the task set.
	ExternalId pulumi.StringPtrInput
	// Whether to allow deleting the task set without waiting for scaling down to 0. You can force a task set to delete even if it's in the process of scaling a resource. Normally, the provider drains all the tasks before deleting the task set. This bypasses that behavior and potentially leaves resources dangling.
	ForceDelete pulumi.BoolPtrInput
	// The launch type on which to run your service. The valid values are `EC2`, `FARGATE`, and `EXTERNAL`. Defaults to `EC2`.
	LaunchType pulumi.StringPtrInput
	// Details on load balancers that are used with a task set. Detailed below.
	LoadBalancers TaskSetLoadBalancerArrayInput
	// The network configuration for the service. This parameter is required for task definitions that use the `awsvpc` network mode to receive their own Elastic Network Interface, and it is not supported for other network modes. Detailed below.
	NetworkConfiguration TaskSetNetworkConfigurationPtrInput
	// The platform version on which to run your service. Only applicable for `launchType` set to `FARGATE`. Defaults to `LATEST`. More information about Fargate platform versions can be found in the [AWS ECS User Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
	PlatformVersion pulumi.StringPtrInput
	// A floating-point percentage of the desired number of tasks to place and keep running in the task set. Detailed below.
	Scale TaskSetScalePtrInput
	// The short name or ARN of the ECS service.
	Service pulumi.StringPtrInput
	// The service discovery registries for the service. The maximum number of `serviceRegistries` blocks is `1`. Detailed below.
	ServiceRegistries TaskSetServiceRegistriesPtrInput
	// The stability status. This indicates whether the task set has reached a steady state.
	StabilityStatus pulumi.StringPtrInput
	// The status of the task set.
	Status pulumi.StringPtrInput
	// A map of tags to assign to the file system. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level. If you have set `copyTagsToBackups` to true, and you specify one or more tags, no existing file system tags are copied from the file system to the backup.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// The family and revision (`family:revision`) or full ARN of the task definition that you want to run in your service.
	//
	// The following arguments are optional:
	TaskDefinition pulumi.StringPtrInput
	// The ID of the task set.
	TaskSetId pulumi.StringPtrInput
	// Whether the provider should wait until the task set has reached `STEADY_STATE`.
	WaitUntilStable pulumi.BoolPtrInput
	// Wait timeout for task set to reach `STEADY_STATE`. Valid time units include `ns`, `us` (or `µs`), `ms`, `s`, `m`, and `h`. Default `10m`.
	WaitUntilStableTimeout pulumi.StringPtrInput
}

func (TaskSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*taskSetState)(nil)).Elem()
}

type taskSetArgs struct {
	// The capacity provider strategy to use for the service. Can be one or more.  Defined below.
	CapacityProviderStrategies []TaskSetCapacityProviderStrategy `pulumi:"capacityProviderStrategies"`
	// The short name or ARN of the cluster that hosts the service to create the task set in.
	Cluster string `pulumi:"cluster"`
	// The external ID associated with the task set.
	ExternalId *string `pulumi:"externalId"`
	// Whether to allow deleting the task set without waiting for scaling down to 0. You can force a task set to delete even if it's in the process of scaling a resource. Normally, the provider drains all the tasks before deleting the task set. This bypasses that behavior and potentially leaves resources dangling.
	ForceDelete *bool `pulumi:"forceDelete"`
	// The launch type on which to run your service. The valid values are `EC2`, `FARGATE`, and `EXTERNAL`. Defaults to `EC2`.
	LaunchType *string `pulumi:"launchType"`
	// Details on load balancers that are used with a task set. Detailed below.
	LoadBalancers []TaskSetLoadBalancer `pulumi:"loadBalancers"`
	// The network configuration for the service. This parameter is required for task definitions that use the `awsvpc` network mode to receive their own Elastic Network Interface, and it is not supported for other network modes. Detailed below.
	NetworkConfiguration *TaskSetNetworkConfiguration `pulumi:"networkConfiguration"`
	// The platform version on which to run your service. Only applicable for `launchType` set to `FARGATE`. Defaults to `LATEST`. More information about Fargate platform versions can be found in the [AWS ECS User Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
	PlatformVersion *string `pulumi:"platformVersion"`
	// A floating-point percentage of the desired number of tasks to place and keep running in the task set. Detailed below.
	Scale *TaskSetScale `pulumi:"scale"`
	// The short name or ARN of the ECS service.
	Service string `pulumi:"service"`
	// The service discovery registries for the service. The maximum number of `serviceRegistries` blocks is `1`. Detailed below.
	ServiceRegistries *TaskSetServiceRegistries `pulumi:"serviceRegistries"`
	// A map of tags to assign to the file system. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level. If you have set `copyTagsToBackups` to true, and you specify one or more tags, no existing file system tags are copied from the file system to the backup.
	Tags map[string]string `pulumi:"tags"`
	// The family and revision (`family:revision`) or full ARN of the task definition that you want to run in your service.
	//
	// The following arguments are optional:
	TaskDefinition string `pulumi:"taskDefinition"`
	// Whether the provider should wait until the task set has reached `STEADY_STATE`.
	WaitUntilStable *bool `pulumi:"waitUntilStable"`
	// Wait timeout for task set to reach `STEADY_STATE`. Valid time units include `ns`, `us` (or `µs`), `ms`, `s`, `m`, and `h`. Default `10m`.
	WaitUntilStableTimeout *string `pulumi:"waitUntilStableTimeout"`
}

// The set of arguments for constructing a TaskSet resource.
type TaskSetArgs struct {
	// The capacity provider strategy to use for the service. Can be one or more.  Defined below.
	CapacityProviderStrategies TaskSetCapacityProviderStrategyArrayInput
	// The short name or ARN of the cluster that hosts the service to create the task set in.
	Cluster pulumi.StringInput
	// The external ID associated with the task set.
	ExternalId pulumi.StringPtrInput
	// Whether to allow deleting the task set without waiting for scaling down to 0. You can force a task set to delete even if it's in the process of scaling a resource. Normally, the provider drains all the tasks before deleting the task set. This bypasses that behavior and potentially leaves resources dangling.
	ForceDelete pulumi.BoolPtrInput
	// The launch type on which to run your service. The valid values are `EC2`, `FARGATE`, and `EXTERNAL`. Defaults to `EC2`.
	LaunchType pulumi.StringPtrInput
	// Details on load balancers that are used with a task set. Detailed below.
	LoadBalancers TaskSetLoadBalancerArrayInput
	// The network configuration for the service. This parameter is required for task definitions that use the `awsvpc` network mode to receive their own Elastic Network Interface, and it is not supported for other network modes. Detailed below.
	NetworkConfiguration TaskSetNetworkConfigurationPtrInput
	// The platform version on which to run your service. Only applicable for `launchType` set to `FARGATE`. Defaults to `LATEST`. More information about Fargate platform versions can be found in the [AWS ECS User Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
	PlatformVersion pulumi.StringPtrInput
	// A floating-point percentage of the desired number of tasks to place and keep running in the task set. Detailed below.
	Scale TaskSetScalePtrInput
	// The short name or ARN of the ECS service.
	Service pulumi.StringInput
	// The service discovery registries for the service. The maximum number of `serviceRegistries` blocks is `1`. Detailed below.
	ServiceRegistries TaskSetServiceRegistriesPtrInput
	// A map of tags to assign to the file system. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level. If you have set `copyTagsToBackups` to true, and you specify one or more tags, no existing file system tags are copied from the file system to the backup.
	Tags pulumi.StringMapInput
	// The family and revision (`family:revision`) or full ARN of the task definition that you want to run in your service.
	//
	// The following arguments are optional:
	TaskDefinition pulumi.StringInput
	// Whether the provider should wait until the task set has reached `STEADY_STATE`.
	WaitUntilStable pulumi.BoolPtrInput
	// Wait timeout for task set to reach `STEADY_STATE`. Valid time units include `ns`, `us` (or `µs`), `ms`, `s`, `m`, and `h`. Default `10m`.
	WaitUntilStableTimeout pulumi.StringPtrInput
}

func (TaskSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*taskSetArgs)(nil)).Elem()
}

type TaskSetInput interface {
	pulumi.Input

	ToTaskSetOutput() TaskSetOutput
	ToTaskSetOutputWithContext(ctx context.Context) TaskSetOutput
}

func (*TaskSet) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskSet)(nil)).Elem()
}

func (i *TaskSet) ToTaskSetOutput() TaskSetOutput {
	return i.ToTaskSetOutputWithContext(context.Background())
}

func (i *TaskSet) ToTaskSetOutputWithContext(ctx context.Context) TaskSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskSetOutput)
}

func (i *TaskSet) ToOutput(ctx context.Context) pulumix.Output[*TaskSet] {
	return pulumix.Output[*TaskSet]{
		OutputState: i.ToTaskSetOutputWithContext(ctx).OutputState,
	}
}

// TaskSetArrayInput is an input type that accepts TaskSetArray and TaskSetArrayOutput values.
// You can construct a concrete instance of `TaskSetArrayInput` via:
//
//	TaskSetArray{ TaskSetArgs{...} }
type TaskSetArrayInput interface {
	pulumi.Input

	ToTaskSetArrayOutput() TaskSetArrayOutput
	ToTaskSetArrayOutputWithContext(context.Context) TaskSetArrayOutput
}

type TaskSetArray []TaskSetInput

func (TaskSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TaskSet)(nil)).Elem()
}

func (i TaskSetArray) ToTaskSetArrayOutput() TaskSetArrayOutput {
	return i.ToTaskSetArrayOutputWithContext(context.Background())
}

func (i TaskSetArray) ToTaskSetArrayOutputWithContext(ctx context.Context) TaskSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskSetArrayOutput)
}

func (i TaskSetArray) ToOutput(ctx context.Context) pulumix.Output[[]*TaskSet] {
	return pulumix.Output[[]*TaskSet]{
		OutputState: i.ToTaskSetArrayOutputWithContext(ctx).OutputState,
	}
}

// TaskSetMapInput is an input type that accepts TaskSetMap and TaskSetMapOutput values.
// You can construct a concrete instance of `TaskSetMapInput` via:
//
//	TaskSetMap{ "key": TaskSetArgs{...} }
type TaskSetMapInput interface {
	pulumi.Input

	ToTaskSetMapOutput() TaskSetMapOutput
	ToTaskSetMapOutputWithContext(context.Context) TaskSetMapOutput
}

type TaskSetMap map[string]TaskSetInput

func (TaskSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TaskSet)(nil)).Elem()
}

func (i TaskSetMap) ToTaskSetMapOutput() TaskSetMapOutput {
	return i.ToTaskSetMapOutputWithContext(context.Background())
}

func (i TaskSetMap) ToTaskSetMapOutputWithContext(ctx context.Context) TaskSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskSetMapOutput)
}

func (i TaskSetMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*TaskSet] {
	return pulumix.Output[map[string]*TaskSet]{
		OutputState: i.ToTaskSetMapOutputWithContext(ctx).OutputState,
	}
}

type TaskSetOutput struct{ *pulumi.OutputState }

func (TaskSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskSet)(nil)).Elem()
}

func (o TaskSetOutput) ToTaskSetOutput() TaskSetOutput {
	return o
}

func (o TaskSetOutput) ToTaskSetOutputWithContext(ctx context.Context) TaskSetOutput {
	return o
}

func (o TaskSetOutput) ToOutput(ctx context.Context) pulumix.Output[*TaskSet] {
	return pulumix.Output[*TaskSet]{
		OutputState: o.OutputState,
	}
}

// The Amazon Resource Name (ARN) that identifies the task set.
func (o TaskSetOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The capacity provider strategy to use for the service. Can be one or more.  Defined below.
func (o TaskSetOutput) CapacityProviderStrategies() TaskSetCapacityProviderStrategyArrayOutput {
	return o.ApplyT(func(v *TaskSet) TaskSetCapacityProviderStrategyArrayOutput { return v.CapacityProviderStrategies }).(TaskSetCapacityProviderStrategyArrayOutput)
}

// The short name or ARN of the cluster that hosts the service to create the task set in.
func (o TaskSetOutput) Cluster() pulumi.StringOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.StringOutput { return v.Cluster }).(pulumi.StringOutput)
}

// The external ID associated with the task set.
func (o TaskSetOutput) ExternalId() pulumi.StringOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.StringOutput { return v.ExternalId }).(pulumi.StringOutput)
}

// Whether to allow deleting the task set without waiting for scaling down to 0. You can force a task set to delete even if it's in the process of scaling a resource. Normally, the provider drains all the tasks before deleting the task set. This bypasses that behavior and potentially leaves resources dangling.
func (o TaskSetOutput) ForceDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.BoolPtrOutput { return v.ForceDelete }).(pulumi.BoolPtrOutput)
}

// The launch type on which to run your service. The valid values are `EC2`, `FARGATE`, and `EXTERNAL`. Defaults to `EC2`.
func (o TaskSetOutput) LaunchType() pulumi.StringOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.StringOutput { return v.LaunchType }).(pulumi.StringOutput)
}

// Details on load balancers that are used with a task set. Detailed below.
func (o TaskSetOutput) LoadBalancers() TaskSetLoadBalancerArrayOutput {
	return o.ApplyT(func(v *TaskSet) TaskSetLoadBalancerArrayOutput { return v.LoadBalancers }).(TaskSetLoadBalancerArrayOutput)
}

// The network configuration for the service. This parameter is required for task definitions that use the `awsvpc` network mode to receive their own Elastic Network Interface, and it is not supported for other network modes. Detailed below.
func (o TaskSetOutput) NetworkConfiguration() TaskSetNetworkConfigurationPtrOutput {
	return o.ApplyT(func(v *TaskSet) TaskSetNetworkConfigurationPtrOutput { return v.NetworkConfiguration }).(TaskSetNetworkConfigurationPtrOutput)
}

// The platform version on which to run your service. Only applicable for `launchType` set to `FARGATE`. Defaults to `LATEST`. More information about Fargate platform versions can be found in the [AWS ECS User Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
func (o TaskSetOutput) PlatformVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.StringOutput { return v.PlatformVersion }).(pulumi.StringOutput)
}

// A floating-point percentage of the desired number of tasks to place and keep running in the task set. Detailed below.
func (o TaskSetOutput) Scale() TaskSetScaleOutput {
	return o.ApplyT(func(v *TaskSet) TaskSetScaleOutput { return v.Scale }).(TaskSetScaleOutput)
}

// The short name or ARN of the ECS service.
func (o TaskSetOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

// The service discovery registries for the service. The maximum number of `serviceRegistries` blocks is `1`. Detailed below.
func (o TaskSetOutput) ServiceRegistries() TaskSetServiceRegistriesPtrOutput {
	return o.ApplyT(func(v *TaskSet) TaskSetServiceRegistriesPtrOutput { return v.ServiceRegistries }).(TaskSetServiceRegistriesPtrOutput)
}

// The stability status. This indicates whether the task set has reached a steady state.
func (o TaskSetOutput) StabilityStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.StringOutput { return v.StabilityStatus }).(pulumi.StringOutput)
}

// The status of the task set.
func (o TaskSetOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A map of tags to assign to the file system. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level. If you have set `copyTagsToBackups` to true, and you specify one or more tags, no existing file system tags are copied from the file system to the backup.
func (o TaskSetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o TaskSetOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The family and revision (`family:revision`) or full ARN of the task definition that you want to run in your service.
//
// The following arguments are optional:
func (o TaskSetOutput) TaskDefinition() pulumi.StringOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.StringOutput { return v.TaskDefinition }).(pulumi.StringOutput)
}

// The ID of the task set.
func (o TaskSetOutput) TaskSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.StringOutput { return v.TaskSetId }).(pulumi.StringOutput)
}

// Whether the provider should wait until the task set has reached `STEADY_STATE`.
func (o TaskSetOutput) WaitUntilStable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.BoolPtrOutput { return v.WaitUntilStable }).(pulumi.BoolPtrOutput)
}

// Wait timeout for task set to reach `STEADY_STATE`. Valid time units include `ns`, `us` (or `µs`), `ms`, `s`, `m`, and `h`. Default `10m`.
func (o TaskSetOutput) WaitUntilStableTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaskSet) pulumi.StringPtrOutput { return v.WaitUntilStableTimeout }).(pulumi.StringPtrOutput)
}

type TaskSetArrayOutput struct{ *pulumi.OutputState }

func (TaskSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TaskSet)(nil)).Elem()
}

func (o TaskSetArrayOutput) ToTaskSetArrayOutput() TaskSetArrayOutput {
	return o
}

func (o TaskSetArrayOutput) ToTaskSetArrayOutputWithContext(ctx context.Context) TaskSetArrayOutput {
	return o
}

func (o TaskSetArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*TaskSet] {
	return pulumix.Output[[]*TaskSet]{
		OutputState: o.OutputState,
	}
}

func (o TaskSetArrayOutput) Index(i pulumi.IntInput) TaskSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TaskSet {
		return vs[0].([]*TaskSet)[vs[1].(int)]
	}).(TaskSetOutput)
}

type TaskSetMapOutput struct{ *pulumi.OutputState }

func (TaskSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TaskSet)(nil)).Elem()
}

func (o TaskSetMapOutput) ToTaskSetMapOutput() TaskSetMapOutput {
	return o
}

func (o TaskSetMapOutput) ToTaskSetMapOutputWithContext(ctx context.Context) TaskSetMapOutput {
	return o
}

func (o TaskSetMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*TaskSet] {
	return pulumix.Output[map[string]*TaskSet]{
		OutputState: o.OutputState,
	}
}

func (o TaskSetMapOutput) MapIndex(k pulumi.StringInput) TaskSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TaskSet {
		return vs[0].(map[string]*TaskSet)[vs[1].(string)]
	}).(TaskSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TaskSetInput)(nil)).Elem(), &TaskSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaskSetArrayInput)(nil)).Elem(), TaskSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TaskSetMapInput)(nil)).Elem(), TaskSetMap{})
	pulumi.RegisterOutputType(TaskSetOutput{})
	pulumi.RegisterOutputType(TaskSetArrayOutput{})
	pulumi.RegisterOutputType(TaskSetMapOutput{})
}
