// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticbeanstalk

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Elastic Beanstalk Environment Resource. Elastic Beanstalk allows
// you to deploy and manage applications in the AWS cloud without worrying about
// the infrastructure that runs those applications.
//
// Environments are often things such as `development`, `integration`, or
// `production`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/elasticbeanstalk"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tftest, err := elasticbeanstalk.NewApplication(ctx, "tftest", &elasticbeanstalk.ApplicationArgs{
//				Description: pulumi.String("tf-test-desc"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = elasticbeanstalk.NewEnvironment(ctx, "tfenvtest", &elasticbeanstalk.EnvironmentArgs{
//				Application:       tftest.Name,
//				SolutionStackName: pulumi.String("64bit Amazon Linux 2015.03 v2.0.3 running Go 1.4"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Option Settings
//
// Some options can be stack-specific, check [AWS Docs](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options-general.html)
// for supported options and examples.
//
// The `setting` and `allSettings` mappings support the following format:
//
// * `namespace` - unique namespace identifying the option's associated AWS resource
// * `name` - name of the configuration option
// * `value` - value for the configuration option
// * `resource` - (Optional) resource name for [scheduled action](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options-general.html#command-options-general-autoscalingscheduledaction)
//
// ### Example With Options
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/elasticbeanstalk"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tftest, err := elasticbeanstalk.NewApplication(ctx, "tftest", &elasticbeanstalk.ApplicationArgs{
//				Description: pulumi.String("tf-test-desc"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = elasticbeanstalk.NewEnvironment(ctx, "tfenvtest", &elasticbeanstalk.EnvironmentArgs{
//				Application:       tftest.Name,
//				SolutionStackName: pulumi.String("64bit Amazon Linux 2015.03 v2.0.3 running Go 1.4"),
//				Settings: elasticbeanstalk.EnvironmentSettingArray{
//					&elasticbeanstalk.EnvironmentSettingArgs{
//						Namespace: pulumi.String("aws:ec2:vpc"),
//						Name:      pulumi.String("VPCId"),
//						Value:     pulumi.String("vpc-xxxxxxxx"),
//					},
//					&elasticbeanstalk.EnvironmentSettingArgs{
//						Namespace: pulumi.String("aws:ec2:vpc"),
//						Name:      pulumi.String("Subnets"),
//						Value:     pulumi.String("subnet-xxxxxxxx"),
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
// Elastic Beanstalk Environments can be imported using the `id`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:elasticbeanstalk/environment:Environment prodenv e-rpqsewtp2j
//
// ```
type Environment struct {
	pulumi.CustomResourceState

	// List of all option settings configured in this Environment. These
	// are a combination of default settings and their overrides from `setting` in
	// the configuration.
	AllSettings EnvironmentAllSettingArrayOutput `pulumi:"allSettings"`
	// Name of the application that contains the version
	// to be deployed
	Application pulumi.StringOutput `pulumi:"application"`
	Arn         pulumi.StringOutput `pulumi:"arn"`
	// The autoscaling groups used by this Environment.
	AutoscalingGroups pulumi.StringArrayOutput `pulumi:"autoscalingGroups"`
	// Fully qualified DNS name for this Environment.
	Cname pulumi.StringOutput `pulumi:"cname"`
	// Prefix to use for the fully qualified DNS name of
	// the Environment.
	CnamePrefix pulumi.StringOutput `pulumi:"cnamePrefix"`
	// Short description of the Environment
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The URL to the Load Balancer for this Environment
	EndpointUrl pulumi.StringOutput `pulumi:"endpointUrl"`
	// Instances used by this Environment.
	Instances pulumi.StringArrayOutput `pulumi:"instances"`
	// Launch configurations in use by this Environment.
	LaunchConfigurations pulumi.StringArrayOutput `pulumi:"launchConfigurations"`
	// Elastic load balancers in use by this Environment.
	LoadBalancers pulumi.StringArrayOutput `pulumi:"loadBalancers"`
	// A unique name for this Environment. This name is used
	// in the application URL
	Name pulumi.StringOutput `pulumi:"name"`
	// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the Elastic Beanstalk [Platform](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-platformarn)
	// to use in deployment
	PlatformArn pulumi.StringOutput `pulumi:"platformArn"`
	// The time between polling the AWS API to
	// check if changes have been applied. Use this to adjust the rate of API calls
	// for any `create` or `update` action. Minimum `10s`, maximum `180s`. Omit this to
	// use the default behavior, which is an exponential backoff
	PollInterval pulumi.StringPtrOutput `pulumi:"pollInterval"`
	// SQS queues in use by this Environment.
	Queues pulumi.StringArrayOutput `pulumi:"queues"`
	// Option settings to configure the new Environment. These
	// override specific values that are set as defaults. The format is detailed
	// below in Option Settings
	Settings EnvironmentSettingArrayOutput `pulumi:"settings"`
	// A solution stack to base your environment
	// off of. Example stacks can be found in the [Amazon API documentation](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html)
	SolutionStackName pulumi.StringOutput `pulumi:"solutionStackName"`
	// A set of tags to apply to the Environment. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The name of the Elastic Beanstalk Configuration
	// template to use in deployment
	TemplateName pulumi.StringPtrOutput `pulumi:"templateName"`
	// Elastic Beanstalk Environment tier. Valid values are `Worker`
	// or `WebServer`. If tier is left blank `WebServer` will be used.
	Tier pulumi.StringPtrOutput `pulumi:"tier"`
	// Autoscaling triggers in use by this Environment.
	Triggers pulumi.StringArrayOutput `pulumi:"triggers"`
	// The name of the Elastic Beanstalk Application Version
	// to use in deployment.
	Version pulumi.StringOutput `pulumi:"version"`
	// The maximum
	// [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
	// wait for an Elastic Beanstalk Environment to be in a ready state before timing
	// out.
	WaitForReadyTimeout pulumi.StringPtrOutput `pulumi:"waitForReadyTimeout"`
}

// NewEnvironment registers a new resource with the given unique name, arguments, and options.
func NewEnvironment(ctx *pulumi.Context,
	name string, args *EnvironmentArgs, opts ...pulumi.ResourceOption) (*Environment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Application == nil {
		return nil, errors.New("invalid value for required argument 'Application'")
	}
	var resource Environment
	err := ctx.RegisterResource("aws:elasticbeanstalk/environment:Environment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironment gets an existing Environment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentState, opts ...pulumi.ResourceOption) (*Environment, error) {
	var resource Environment
	err := ctx.ReadResource("aws:elasticbeanstalk/environment:Environment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Environment resources.
type environmentState struct {
	// List of all option settings configured in this Environment. These
	// are a combination of default settings and their overrides from `setting` in
	// the configuration.
	AllSettings []EnvironmentAllSetting `pulumi:"allSettings"`
	// Name of the application that contains the version
	// to be deployed
	Application interface{} `pulumi:"application"`
	Arn         *string     `pulumi:"arn"`
	// The autoscaling groups used by this Environment.
	AutoscalingGroups []string `pulumi:"autoscalingGroups"`
	// Fully qualified DNS name for this Environment.
	Cname *string `pulumi:"cname"`
	// Prefix to use for the fully qualified DNS name of
	// the Environment.
	CnamePrefix *string `pulumi:"cnamePrefix"`
	// Short description of the Environment
	Description *string `pulumi:"description"`
	// The URL to the Load Balancer for this Environment
	EndpointUrl *string `pulumi:"endpointUrl"`
	// Instances used by this Environment.
	Instances []string `pulumi:"instances"`
	// Launch configurations in use by this Environment.
	LaunchConfigurations []string `pulumi:"launchConfigurations"`
	// Elastic load balancers in use by this Environment.
	LoadBalancers []string `pulumi:"loadBalancers"`
	// A unique name for this Environment. This name is used
	// in the application URL
	Name *string `pulumi:"name"`
	// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the Elastic Beanstalk [Platform](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-platformarn)
	// to use in deployment
	PlatformArn *string `pulumi:"platformArn"`
	// The time between polling the AWS API to
	// check if changes have been applied. Use this to adjust the rate of API calls
	// for any `create` or `update` action. Minimum `10s`, maximum `180s`. Omit this to
	// use the default behavior, which is an exponential backoff
	PollInterval *string `pulumi:"pollInterval"`
	// SQS queues in use by this Environment.
	Queues []string `pulumi:"queues"`
	// Option settings to configure the new Environment. These
	// override specific values that are set as defaults. The format is detailed
	// below in Option Settings
	Settings []EnvironmentSetting `pulumi:"settings"`
	// A solution stack to base your environment
	// off of. Example stacks can be found in the [Amazon API documentation](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html)
	SolutionStackName *string `pulumi:"solutionStackName"`
	// A set of tags to apply to the Environment. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The name of the Elastic Beanstalk Configuration
	// template to use in deployment
	TemplateName *string `pulumi:"templateName"`
	// Elastic Beanstalk Environment tier. Valid values are `Worker`
	// or `WebServer`. If tier is left blank `WebServer` will be used.
	Tier *string `pulumi:"tier"`
	// Autoscaling triggers in use by this Environment.
	Triggers []string `pulumi:"triggers"`
	// The name of the Elastic Beanstalk Application Version
	// to use in deployment.
	Version *string `pulumi:"version"`
	// The maximum
	// [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
	// wait for an Elastic Beanstalk Environment to be in a ready state before timing
	// out.
	WaitForReadyTimeout *string `pulumi:"waitForReadyTimeout"`
}

type EnvironmentState struct {
	// List of all option settings configured in this Environment. These
	// are a combination of default settings and their overrides from `setting` in
	// the configuration.
	AllSettings EnvironmentAllSettingArrayInput
	// Name of the application that contains the version
	// to be deployed
	Application pulumi.Input
	Arn         pulumi.StringPtrInput
	// The autoscaling groups used by this Environment.
	AutoscalingGroups pulumi.StringArrayInput
	// Fully qualified DNS name for this Environment.
	Cname pulumi.StringPtrInput
	// Prefix to use for the fully qualified DNS name of
	// the Environment.
	CnamePrefix pulumi.StringPtrInput
	// Short description of the Environment
	Description pulumi.StringPtrInput
	// The URL to the Load Balancer for this Environment
	EndpointUrl pulumi.StringPtrInput
	// Instances used by this Environment.
	Instances pulumi.StringArrayInput
	// Launch configurations in use by this Environment.
	LaunchConfigurations pulumi.StringArrayInput
	// Elastic load balancers in use by this Environment.
	LoadBalancers pulumi.StringArrayInput
	// A unique name for this Environment. This name is used
	// in the application URL
	Name pulumi.StringPtrInput
	// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the Elastic Beanstalk [Platform](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-platformarn)
	// to use in deployment
	PlatformArn pulumi.StringPtrInput
	// The time between polling the AWS API to
	// check if changes have been applied. Use this to adjust the rate of API calls
	// for any `create` or `update` action. Minimum `10s`, maximum `180s`. Omit this to
	// use the default behavior, which is an exponential backoff
	PollInterval pulumi.StringPtrInput
	// SQS queues in use by this Environment.
	Queues pulumi.StringArrayInput
	// Option settings to configure the new Environment. These
	// override specific values that are set as defaults. The format is detailed
	// below in Option Settings
	Settings EnvironmentSettingArrayInput
	// A solution stack to base your environment
	// off of. Example stacks can be found in the [Amazon API documentation](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html)
	SolutionStackName pulumi.StringPtrInput
	// A set of tags to apply to the Environment. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// The name of the Elastic Beanstalk Configuration
	// template to use in deployment
	TemplateName pulumi.StringPtrInput
	// Elastic Beanstalk Environment tier. Valid values are `Worker`
	// or `WebServer`. If tier is left blank `WebServer` will be used.
	Tier pulumi.StringPtrInput
	// Autoscaling triggers in use by this Environment.
	Triggers pulumi.StringArrayInput
	// The name of the Elastic Beanstalk Application Version
	// to use in deployment.
	Version pulumi.StringPtrInput
	// The maximum
	// [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
	// wait for an Elastic Beanstalk Environment to be in a ready state before timing
	// out.
	WaitForReadyTimeout pulumi.StringPtrInput
}

func (EnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentState)(nil)).Elem()
}

type environmentArgs struct {
	// Name of the application that contains the version
	// to be deployed
	Application interface{} `pulumi:"application"`
	// Prefix to use for the fully qualified DNS name of
	// the Environment.
	CnamePrefix *string `pulumi:"cnamePrefix"`
	// Short description of the Environment
	Description *string `pulumi:"description"`
	// A unique name for this Environment. This name is used
	// in the application URL
	Name *string `pulumi:"name"`
	// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the Elastic Beanstalk [Platform](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-platformarn)
	// to use in deployment
	PlatformArn *string `pulumi:"platformArn"`
	// The time between polling the AWS API to
	// check if changes have been applied. Use this to adjust the rate of API calls
	// for any `create` or `update` action. Minimum `10s`, maximum `180s`. Omit this to
	// use the default behavior, which is an exponential backoff
	PollInterval *string `pulumi:"pollInterval"`
	// Option settings to configure the new Environment. These
	// override specific values that are set as defaults. The format is detailed
	// below in Option Settings
	Settings []EnvironmentSetting `pulumi:"settings"`
	// A solution stack to base your environment
	// off of. Example stacks can be found in the [Amazon API documentation](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html)
	SolutionStackName *string `pulumi:"solutionStackName"`
	// A set of tags to apply to the Environment. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The name of the Elastic Beanstalk Configuration
	// template to use in deployment
	TemplateName *string `pulumi:"templateName"`
	// Elastic Beanstalk Environment tier. Valid values are `Worker`
	// or `WebServer`. If tier is left blank `WebServer` will be used.
	Tier *string `pulumi:"tier"`
	// The name of the Elastic Beanstalk Application Version
	// to use in deployment.
	Version *string `pulumi:"version"`
	// The maximum
	// [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
	// wait for an Elastic Beanstalk Environment to be in a ready state before timing
	// out.
	WaitForReadyTimeout *string `pulumi:"waitForReadyTimeout"`
}

// The set of arguments for constructing a Environment resource.
type EnvironmentArgs struct {
	// Name of the application that contains the version
	// to be deployed
	Application pulumi.Input
	// Prefix to use for the fully qualified DNS name of
	// the Environment.
	CnamePrefix pulumi.StringPtrInput
	// Short description of the Environment
	Description pulumi.StringPtrInput
	// A unique name for this Environment. This name is used
	// in the application URL
	Name pulumi.StringPtrInput
	// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the Elastic Beanstalk [Platform](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-platformarn)
	// to use in deployment
	PlatformArn pulumi.StringPtrInput
	// The time between polling the AWS API to
	// check if changes have been applied. Use this to adjust the rate of API calls
	// for any `create` or `update` action. Minimum `10s`, maximum `180s`. Omit this to
	// use the default behavior, which is an exponential backoff
	PollInterval pulumi.StringPtrInput
	// Option settings to configure the new Environment. These
	// override specific values that are set as defaults. The format is detailed
	// below in Option Settings
	Settings EnvironmentSettingArrayInput
	// A solution stack to base your environment
	// off of. Example stacks can be found in the [Amazon API documentation](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html)
	SolutionStackName pulumi.StringPtrInput
	// A set of tags to apply to the Environment. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The name of the Elastic Beanstalk Configuration
	// template to use in deployment
	TemplateName pulumi.StringPtrInput
	// Elastic Beanstalk Environment tier. Valid values are `Worker`
	// or `WebServer`. If tier is left blank `WebServer` will be used.
	Tier pulumi.StringPtrInput
	// The name of the Elastic Beanstalk Application Version
	// to use in deployment.
	Version pulumi.StringPtrInput
	// The maximum
	// [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
	// wait for an Elastic Beanstalk Environment to be in a ready state before timing
	// out.
	WaitForReadyTimeout pulumi.StringPtrInput
}

func (EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentArgs)(nil)).Elem()
}

type EnvironmentInput interface {
	pulumi.Input

	ToEnvironmentOutput() EnvironmentOutput
	ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput
}

func (*Environment) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil)).Elem()
}

func (i *Environment) ToEnvironmentOutput() EnvironmentOutput {
	return i.ToEnvironmentOutputWithContext(context.Background())
}

func (i *Environment) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentOutput)
}

// EnvironmentArrayInput is an input type that accepts EnvironmentArray and EnvironmentArrayOutput values.
// You can construct a concrete instance of `EnvironmentArrayInput` via:
//
//	EnvironmentArray{ EnvironmentArgs{...} }
type EnvironmentArrayInput interface {
	pulumi.Input

	ToEnvironmentArrayOutput() EnvironmentArrayOutput
	ToEnvironmentArrayOutputWithContext(context.Context) EnvironmentArrayOutput
}

type EnvironmentArray []EnvironmentInput

func (EnvironmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Environment)(nil)).Elem()
}

func (i EnvironmentArray) ToEnvironmentArrayOutput() EnvironmentArrayOutput {
	return i.ToEnvironmentArrayOutputWithContext(context.Background())
}

func (i EnvironmentArray) ToEnvironmentArrayOutputWithContext(ctx context.Context) EnvironmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentArrayOutput)
}

// EnvironmentMapInput is an input type that accepts EnvironmentMap and EnvironmentMapOutput values.
// You can construct a concrete instance of `EnvironmentMapInput` via:
//
//	EnvironmentMap{ "key": EnvironmentArgs{...} }
type EnvironmentMapInput interface {
	pulumi.Input

	ToEnvironmentMapOutput() EnvironmentMapOutput
	ToEnvironmentMapOutputWithContext(context.Context) EnvironmentMapOutput
}

type EnvironmentMap map[string]EnvironmentInput

func (EnvironmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Environment)(nil)).Elem()
}

func (i EnvironmentMap) ToEnvironmentMapOutput() EnvironmentMapOutput {
	return i.ToEnvironmentMapOutputWithContext(context.Background())
}

func (i EnvironmentMap) ToEnvironmentMapOutputWithContext(ctx context.Context) EnvironmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentMapOutput)
}

type EnvironmentOutput struct{ *pulumi.OutputState }

func (EnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil)).Elem()
}

func (o EnvironmentOutput) ToEnvironmentOutput() EnvironmentOutput {
	return o
}

func (o EnvironmentOutput) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return o
}

// List of all option settings configured in this Environment. These
// are a combination of default settings and their overrides from `setting` in
// the configuration.
func (o EnvironmentOutput) AllSettings() EnvironmentAllSettingArrayOutput {
	return o.ApplyT(func(v *Environment) EnvironmentAllSettingArrayOutput { return v.AllSettings }).(EnvironmentAllSettingArrayOutput)
}

// Name of the application that contains the version
// to be deployed
func (o EnvironmentOutput) Application() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.Application }).(pulumi.StringOutput)
}

func (o EnvironmentOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The autoscaling groups used by this Environment.
func (o EnvironmentOutput) AutoscalingGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringArrayOutput { return v.AutoscalingGroups }).(pulumi.StringArrayOutput)
}

// Fully qualified DNS name for this Environment.
func (o EnvironmentOutput) Cname() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.Cname }).(pulumi.StringOutput)
}

// Prefix to use for the fully qualified DNS name of
// the Environment.
func (o EnvironmentOutput) CnamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.CnamePrefix }).(pulumi.StringOutput)
}

// Short description of the Environment
func (o EnvironmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The URL to the Load Balancer for this Environment
func (o EnvironmentOutput) EndpointUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.EndpointUrl }).(pulumi.StringOutput)
}

// Instances used by this Environment.
func (o EnvironmentOutput) Instances() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringArrayOutput { return v.Instances }).(pulumi.StringArrayOutput)
}

// Launch configurations in use by this Environment.
func (o EnvironmentOutput) LaunchConfigurations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringArrayOutput { return v.LaunchConfigurations }).(pulumi.StringArrayOutput)
}

// Elastic load balancers in use by this Environment.
func (o EnvironmentOutput) LoadBalancers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringArrayOutput { return v.LoadBalancers }).(pulumi.StringArrayOutput)
}

// A unique name for this Environment. This name is used
// in the application URL
func (o EnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the Elastic Beanstalk [Platform](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-platformarn)
// to use in deployment
func (o EnvironmentOutput) PlatformArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.PlatformArn }).(pulumi.StringOutput)
}

// The time between polling the AWS API to
// check if changes have been applied. Use this to adjust the rate of API calls
// for any `create` or `update` action. Minimum `10s`, maximum `180s`. Omit this to
// use the default behavior, which is an exponential backoff
func (o EnvironmentOutput) PollInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.PollInterval }).(pulumi.StringPtrOutput)
}

// SQS queues in use by this Environment.
func (o EnvironmentOutput) Queues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringArrayOutput { return v.Queues }).(pulumi.StringArrayOutput)
}

// Option settings to configure the new Environment. These
// override specific values that are set as defaults. The format is detailed
// below in Option Settings
func (o EnvironmentOutput) Settings() EnvironmentSettingArrayOutput {
	return o.ApplyT(func(v *Environment) EnvironmentSettingArrayOutput { return v.Settings }).(EnvironmentSettingArrayOutput)
}

// A solution stack to base your environment
// off of. Example stacks can be found in the [Amazon API documentation](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html)
func (o EnvironmentOutput) SolutionStackName() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.SolutionStackName }).(pulumi.StringOutput)
}

// A set of tags to apply to the Environment. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o EnvironmentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o EnvironmentOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The name of the Elastic Beanstalk Configuration
// template to use in deployment
func (o EnvironmentOutput) TemplateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.TemplateName }).(pulumi.StringPtrOutput)
}

// Elastic Beanstalk Environment tier. Valid values are `Worker`
// or `WebServer`. If tier is left blank `WebServer` will be used.
func (o EnvironmentOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.Tier }).(pulumi.StringPtrOutput)
}

// Autoscaling triggers in use by this Environment.
func (o EnvironmentOutput) Triggers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringArrayOutput { return v.Triggers }).(pulumi.StringArrayOutput)
}

// The name of the Elastic Beanstalk Application Version
// to use in deployment.
func (o EnvironmentOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

// The maximum
// [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
// wait for an Elastic Beanstalk Environment to be in a ready state before timing
// out.
func (o EnvironmentOutput) WaitForReadyTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Environment) pulumi.StringPtrOutput { return v.WaitForReadyTimeout }).(pulumi.StringPtrOutput)
}

type EnvironmentArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Environment)(nil)).Elem()
}

func (o EnvironmentArrayOutput) ToEnvironmentArrayOutput() EnvironmentArrayOutput {
	return o
}

func (o EnvironmentArrayOutput) ToEnvironmentArrayOutputWithContext(ctx context.Context) EnvironmentArrayOutput {
	return o
}

func (o EnvironmentArrayOutput) Index(i pulumi.IntInput) EnvironmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Environment {
		return vs[0].([]*Environment)[vs[1].(int)]
	}).(EnvironmentOutput)
}

type EnvironmentMapOutput struct{ *pulumi.OutputState }

func (EnvironmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Environment)(nil)).Elem()
}

func (o EnvironmentMapOutput) ToEnvironmentMapOutput() EnvironmentMapOutput {
	return o
}

func (o EnvironmentMapOutput) ToEnvironmentMapOutputWithContext(ctx context.Context) EnvironmentMapOutput {
	return o
}

func (o EnvironmentMapOutput) MapIndex(k pulumi.StringInput) EnvironmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Environment {
		return vs[0].(map[string]*Environment)[vs[1].(string)]
	}).(EnvironmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentInput)(nil)).Elem(), &Environment{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentArrayInput)(nil)).Elem(), EnvironmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentMapInput)(nil)).Elem(), EnvironmentMap{})
	pulumi.RegisterOutputType(EnvironmentOutput{})
	pulumi.RegisterOutputType(EnvironmentArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentMapOutput{})
}
