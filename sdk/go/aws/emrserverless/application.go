// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package emrserverless

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an EMR Serverless Application.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/emrserverless"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := emrserverless.NewApplication(ctx, "example", &emrserverless.ApplicationArgs{
//				ReleaseLabel: pulumi.String("emr-6.6.0"),
//				Type:         pulumi.String("hive"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Initial Capacity Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/emrserverless"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := emrserverless.NewApplication(ctx, "example", &emrserverless.ApplicationArgs{
//				InitialCapacities: emrserverless.ApplicationInitialCapacityArray{
//					&emrserverless.ApplicationInitialCapacityArgs{
//						InitialCapacityConfig: &emrserverless.ApplicationInitialCapacityInitialCapacityConfigArgs{
//							WorkerConfiguration: &emrserverless.ApplicationInitialCapacityInitialCapacityConfigWorkerConfigurationArgs{
//								Cpu:    pulumi.String("2 vCPU"),
//								Memory: pulumi.String("10 GB"),
//							},
//							WorkerCount: pulumi.Int(1),
//						},
//						InitialCapacityType: pulumi.String("HiveDriver"),
//					},
//				},
//				ReleaseLabel: pulumi.String("emr-6.6.0"),
//				Type:         pulumi.String("hive"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Maximum Capacity Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/emrserverless"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := emrserverless.NewApplication(ctx, "example", &emrserverless.ApplicationArgs{
//				MaximumCapacity: &emrserverless.ApplicationMaximumCapacityArgs{
//					Cpu:    pulumi.String("2 vCPU"),
//					Memory: pulumi.String("10 GB"),
//				},
//				ReleaseLabel: pulumi.String("emr-6.6.0"),
//				Type:         pulumi.String("hive"),
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
//	to = aws_emrserverless_application.example
//
//	id = "id" } Using `pulumi import`, import EMR Severless applications using the `id`. For exampleconsole % pulumi import aws_emrserverless_application.example id
type Application struct {
	pulumi.CustomResourceState

	// The CPU architecture of an application. Valid values are `ARM64` or `X86_64`. Default value is `X86_64`.
	Architecture pulumi.StringPtrOutput `pulumi:"architecture"`
	// ARN of the cluster.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The configuration for an application to automatically start on job submission.
	AutoStartConfiguration ApplicationAutoStartConfigurationOutput `pulumi:"autoStartConfiguration"`
	// The configuration for an application to automatically stop after a certain amount of time being idle.
	AutoStopConfiguration ApplicationAutoStopConfigurationOutput `pulumi:"autoStopConfiguration"`
	// The image configuration applied to all worker types.
	ImageConfiguration ApplicationImageConfigurationOutput `pulumi:"imageConfiguration"`
	// The capacity to initialize when the application is created.
	InitialCapacities ApplicationInitialCapacityArrayOutput `pulumi:"initialCapacities"`
	// The maximum capacity to allocate when the application is created. This is cumulative across all workers at any given point in time, not just when an application is created. No new resources will be created once any one of the defined limits is hit.
	MaximumCapacity ApplicationMaximumCapacityOutput `pulumi:"maximumCapacity"`
	// The name of the application.
	Name pulumi.StringOutput `pulumi:"name"`
	// The network configuration for customer VPC connectivity.
	NetworkConfiguration ApplicationNetworkConfigurationPtrOutput `pulumi:"networkConfiguration"`
	// The EMR release version associated with the application.
	ReleaseLabel pulumi.StringOutput `pulumi:"releaseLabel"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The type of application you want to start, such as `spark` or `hive`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ReleaseLabel == nil {
		return nil, errors.New("invalid value for required argument 'ReleaseLabel'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Application
	err := ctx.RegisterResource("aws:emrserverless/application:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("aws:emrserverless/application:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
	// The CPU architecture of an application. Valid values are `ARM64` or `X86_64`. Default value is `X86_64`.
	Architecture *string `pulumi:"architecture"`
	// ARN of the cluster.
	Arn *string `pulumi:"arn"`
	// The configuration for an application to automatically start on job submission.
	AutoStartConfiguration *ApplicationAutoStartConfiguration `pulumi:"autoStartConfiguration"`
	// The configuration for an application to automatically stop after a certain amount of time being idle.
	AutoStopConfiguration *ApplicationAutoStopConfiguration `pulumi:"autoStopConfiguration"`
	// The image configuration applied to all worker types.
	ImageConfiguration *ApplicationImageConfiguration `pulumi:"imageConfiguration"`
	// The capacity to initialize when the application is created.
	InitialCapacities []ApplicationInitialCapacity `pulumi:"initialCapacities"`
	// The maximum capacity to allocate when the application is created. This is cumulative across all workers at any given point in time, not just when an application is created. No new resources will be created once any one of the defined limits is hit.
	MaximumCapacity *ApplicationMaximumCapacity `pulumi:"maximumCapacity"`
	// The name of the application.
	Name *string `pulumi:"name"`
	// The network configuration for customer VPC connectivity.
	NetworkConfiguration *ApplicationNetworkConfiguration `pulumi:"networkConfiguration"`
	// The EMR release version associated with the application.
	ReleaseLabel *string `pulumi:"releaseLabel"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The type of application you want to start, such as `spark` or `hive`.
	Type *string `pulumi:"type"`
}

type ApplicationState struct {
	// The CPU architecture of an application. Valid values are `ARM64` or `X86_64`. Default value is `X86_64`.
	Architecture pulumi.StringPtrInput
	// ARN of the cluster.
	Arn pulumi.StringPtrInput
	// The configuration for an application to automatically start on job submission.
	AutoStartConfiguration ApplicationAutoStartConfigurationPtrInput
	// The configuration for an application to automatically stop after a certain amount of time being idle.
	AutoStopConfiguration ApplicationAutoStopConfigurationPtrInput
	// The image configuration applied to all worker types.
	ImageConfiguration ApplicationImageConfigurationPtrInput
	// The capacity to initialize when the application is created.
	InitialCapacities ApplicationInitialCapacityArrayInput
	// The maximum capacity to allocate when the application is created. This is cumulative across all workers at any given point in time, not just when an application is created. No new resources will be created once any one of the defined limits is hit.
	MaximumCapacity ApplicationMaximumCapacityPtrInput
	// The name of the application.
	Name pulumi.StringPtrInput
	// The network configuration for customer VPC connectivity.
	NetworkConfiguration ApplicationNetworkConfigurationPtrInput
	// The EMR release version associated with the application.
	ReleaseLabel pulumi.StringPtrInput
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// The type of application you want to start, such as `spark` or `hive`.
	Type pulumi.StringPtrInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// The CPU architecture of an application. Valid values are `ARM64` or `X86_64`. Default value is `X86_64`.
	Architecture *string `pulumi:"architecture"`
	// The configuration for an application to automatically start on job submission.
	AutoStartConfiguration *ApplicationAutoStartConfiguration `pulumi:"autoStartConfiguration"`
	// The configuration for an application to automatically stop after a certain amount of time being idle.
	AutoStopConfiguration *ApplicationAutoStopConfiguration `pulumi:"autoStopConfiguration"`
	// The image configuration applied to all worker types.
	ImageConfiguration *ApplicationImageConfiguration `pulumi:"imageConfiguration"`
	// The capacity to initialize when the application is created.
	InitialCapacities []ApplicationInitialCapacity `pulumi:"initialCapacities"`
	// The maximum capacity to allocate when the application is created. This is cumulative across all workers at any given point in time, not just when an application is created. No new resources will be created once any one of the defined limits is hit.
	MaximumCapacity *ApplicationMaximumCapacity `pulumi:"maximumCapacity"`
	// The name of the application.
	Name *string `pulumi:"name"`
	// The network configuration for customer VPC connectivity.
	NetworkConfiguration *ApplicationNetworkConfiguration `pulumi:"networkConfiguration"`
	// The EMR release version associated with the application.
	ReleaseLabel string `pulumi:"releaseLabel"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
	Tags map[string]string `pulumi:"tags"`
	// The type of application you want to start, such as `spark` or `hive`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// The CPU architecture of an application. Valid values are `ARM64` or `X86_64`. Default value is `X86_64`.
	Architecture pulumi.StringPtrInput
	// The configuration for an application to automatically start on job submission.
	AutoStartConfiguration ApplicationAutoStartConfigurationPtrInput
	// The configuration for an application to automatically stop after a certain amount of time being idle.
	AutoStopConfiguration ApplicationAutoStopConfigurationPtrInput
	// The image configuration applied to all worker types.
	ImageConfiguration ApplicationImageConfigurationPtrInput
	// The capacity to initialize when the application is created.
	InitialCapacities ApplicationInitialCapacityArrayInput
	// The maximum capacity to allocate when the application is created. This is cumulative across all workers at any given point in time, not just when an application is created. No new resources will be created once any one of the defined limits is hit.
	MaximumCapacity ApplicationMaximumCapacityPtrInput
	// The name of the application.
	Name pulumi.StringPtrInput
	// The network configuration for customer VPC connectivity.
	NetworkConfiguration ApplicationNetworkConfigurationPtrInput
	// The EMR release version associated with the application.
	ReleaseLabel pulumi.StringInput
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
	Tags pulumi.StringMapInput
	// The type of application you want to start, such as `spark` or `hive`.
	Type pulumi.StringInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

type ApplicationInput interface {
	pulumi.Input

	ToApplicationOutput() ApplicationOutput
	ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput
}

func (*Application) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (i *Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i *Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

// ApplicationArrayInput is an input type that accepts ApplicationArray and ApplicationArrayOutput values.
// You can construct a concrete instance of `ApplicationArrayInput` via:
//
//	ApplicationArray{ ApplicationArgs{...} }
type ApplicationArrayInput interface {
	pulumi.Input

	ToApplicationArrayOutput() ApplicationArrayOutput
	ToApplicationArrayOutputWithContext(context.Context) ApplicationArrayOutput
}

type ApplicationArray []ApplicationInput

func (ApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Application)(nil)).Elem()
}

func (i ApplicationArray) ToApplicationArrayOutput() ApplicationArrayOutput {
	return i.ToApplicationArrayOutputWithContext(context.Background())
}

func (i ApplicationArray) ToApplicationArrayOutputWithContext(ctx context.Context) ApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationArrayOutput)
}

// ApplicationMapInput is an input type that accepts ApplicationMap and ApplicationMapOutput values.
// You can construct a concrete instance of `ApplicationMapInput` via:
//
//	ApplicationMap{ "key": ApplicationArgs{...} }
type ApplicationMapInput interface {
	pulumi.Input

	ToApplicationMapOutput() ApplicationMapOutput
	ToApplicationMapOutputWithContext(context.Context) ApplicationMapOutput
}

type ApplicationMap map[string]ApplicationInput

func (ApplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Application)(nil)).Elem()
}

func (i ApplicationMap) ToApplicationMapOutput() ApplicationMapOutput {
	return i.ToApplicationMapOutputWithContext(context.Background())
}

func (i ApplicationMap) ToApplicationMapOutputWithContext(ctx context.Context) ApplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationMapOutput)
}

type ApplicationOutput struct{ *pulumi.OutputState }

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

// The CPU architecture of an application. Valid values are `ARM64` or `X86_64`. Default value is `X86_64`.
func (o ApplicationOutput) Architecture() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.Architecture }).(pulumi.StringPtrOutput)
}

// ARN of the cluster.
func (o ApplicationOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The configuration for an application to automatically start on job submission.
func (o ApplicationOutput) AutoStartConfiguration() ApplicationAutoStartConfigurationOutput {
	return o.ApplyT(func(v *Application) ApplicationAutoStartConfigurationOutput { return v.AutoStartConfiguration }).(ApplicationAutoStartConfigurationOutput)
}

// The configuration for an application to automatically stop after a certain amount of time being idle.
func (o ApplicationOutput) AutoStopConfiguration() ApplicationAutoStopConfigurationOutput {
	return o.ApplyT(func(v *Application) ApplicationAutoStopConfigurationOutput { return v.AutoStopConfiguration }).(ApplicationAutoStopConfigurationOutput)
}

// The image configuration applied to all worker types.
func (o ApplicationOutput) ImageConfiguration() ApplicationImageConfigurationOutput {
	return o.ApplyT(func(v *Application) ApplicationImageConfigurationOutput { return v.ImageConfiguration }).(ApplicationImageConfigurationOutput)
}

// The capacity to initialize when the application is created.
func (o ApplicationOutput) InitialCapacities() ApplicationInitialCapacityArrayOutput {
	return o.ApplyT(func(v *Application) ApplicationInitialCapacityArrayOutput { return v.InitialCapacities }).(ApplicationInitialCapacityArrayOutput)
}

// The maximum capacity to allocate when the application is created. This is cumulative across all workers at any given point in time, not just when an application is created. No new resources will be created once any one of the defined limits is hit.
func (o ApplicationOutput) MaximumCapacity() ApplicationMaximumCapacityOutput {
	return o.ApplyT(func(v *Application) ApplicationMaximumCapacityOutput { return v.MaximumCapacity }).(ApplicationMaximumCapacityOutput)
}

// The name of the application.
func (o ApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The network configuration for customer VPC connectivity.
func (o ApplicationOutput) NetworkConfiguration() ApplicationNetworkConfigurationPtrOutput {
	return o.ApplyT(func(v *Application) ApplicationNetworkConfigurationPtrOutput { return v.NetworkConfiguration }).(ApplicationNetworkConfigurationPtrOutput)
}

// The EMR release version associated with the application.
func (o ApplicationOutput) ReleaseLabel() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.ReleaseLabel }).(pulumi.StringOutput)
}

// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
func (o ApplicationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Application) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o ApplicationOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Application) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The type of application you want to start, such as `spark` or `hive`.
func (o ApplicationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type ApplicationArrayOutput struct{ *pulumi.OutputState }

func (ApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Application)(nil)).Elem()
}

func (o ApplicationArrayOutput) ToApplicationArrayOutput() ApplicationArrayOutput {
	return o
}

func (o ApplicationArrayOutput) ToApplicationArrayOutputWithContext(ctx context.Context) ApplicationArrayOutput {
	return o
}

func (o ApplicationArrayOutput) Index(i pulumi.IntInput) ApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Application {
		return vs[0].([]*Application)[vs[1].(int)]
	}).(ApplicationOutput)
}

type ApplicationMapOutput struct{ *pulumi.OutputState }

func (ApplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Application)(nil)).Elem()
}

func (o ApplicationMapOutput) ToApplicationMapOutput() ApplicationMapOutput {
	return o
}

func (o ApplicationMapOutput) ToApplicationMapOutputWithContext(ctx context.Context) ApplicationMapOutput {
	return o
}

func (o ApplicationMapOutput) MapIndex(k pulumi.StringInput) ApplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Application {
		return vs[0].(map[string]*Application)[vs[1].(string)]
	}).(ApplicationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationInput)(nil)).Elem(), &Application{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationArrayInput)(nil)).Elem(), ApplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationMapInput)(nil)).Elem(), ApplicationMap{})
	pulumi.RegisterOutputType(ApplicationOutput{})
	pulumi.RegisterOutputType(ApplicationArrayOutput{})
	pulumi.RegisterOutputType(ApplicationMapOutput{})
}
