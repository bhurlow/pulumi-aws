// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an OpsWorks NodeJS application layer resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/opsworks"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := opsworks.NewNodejsAppLayer(ctx, "app", &opsworks.NodejsAppLayerArgs{
// 			StackId: pulumi.Any(aws_opsworks_stack.Main.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type NodejsAppLayer struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name(ARN) of the layer.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps pulumi.BoolPtrOutput `pulumi:"autoAssignElasticIps"`
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps pulumi.BoolPtrOutput `pulumi:"autoAssignPublicIps"`
	// Whether to enable auto-healing for the layer.
	AutoHealing            pulumi.BoolPtrOutput     `pulumi:"autoHealing"`
	CustomConfigureRecipes pulumi.StringArrayOutput `pulumi:"customConfigureRecipes"`
	CustomDeployRecipes    pulumi.StringArrayOutput `pulumi:"customDeployRecipes"`
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn pulumi.StringPtrOutput `pulumi:"customInstanceProfileArn"`
	// Custom JSON attributes to apply to the layer.
	CustomJson pulumi.StringPtrOutput `pulumi:"customJson"`
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds pulumi.StringArrayOutput `pulumi:"customSecurityGroupIds"`
	CustomSetupRecipes     pulumi.StringArrayOutput `pulumi:"customSetupRecipes"`
	CustomShutdownRecipes  pulumi.StringArrayOutput `pulumi:"customShutdownRecipes"`
	CustomUndeployRecipes  pulumi.StringArrayOutput `pulumi:"customUndeployRecipes"`
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown pulumi.BoolPtrOutput `pulumi:"drainElbOnShutdown"`
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes NodejsAppLayerEbsVolumeArrayOutput `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrOutput `pulumi:"elasticLoadBalancer"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrOutput `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrOutput `pulumi:"instanceShutdownTimeout"`
	// A human-readable name for the layer.
	Name pulumi.StringOutput `pulumi:"name"`
	// The version of NodeJS to use. Defaults to "0.10.38".
	NodejsVersion pulumi.StringPtrOutput `pulumi:"nodejsVersion"`
	// The id of the stack the layer will belong to.
	StackId pulumi.StringOutput `pulumi:"stackId"`
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.StringArrayOutput `pulumi:"systemPackages"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolPtrOutput `pulumi:"useEbsOptimizedInstances"`
}

// NewNodejsAppLayer registers a new resource with the given unique name, arguments, and options.
func NewNodejsAppLayer(ctx *pulumi.Context,
	name string, args *NodejsAppLayerArgs, opts ...pulumi.ResourceOption) (*NodejsAppLayer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StackId == nil {
		return nil, errors.New("invalid value for required argument 'StackId'")
	}
	var resource NodejsAppLayer
	err := ctx.RegisterResource("aws:opsworks/nodejsAppLayer:NodejsAppLayer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNodejsAppLayer gets an existing NodejsAppLayer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodejsAppLayer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodejsAppLayerState, opts ...pulumi.ResourceOption) (*NodejsAppLayer, error) {
	var resource NodejsAppLayer
	err := ctx.ReadResource("aws:opsworks/nodejsAppLayer:NodejsAppLayer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NodejsAppLayer resources.
type nodejsAppLayerState struct {
	// The Amazon Resource Name(ARN) of the layer.
	Arn *string `pulumi:"arn"`
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps *bool `pulumi:"autoAssignElasticIps"`
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps *bool `pulumi:"autoAssignPublicIps"`
	// Whether to enable auto-healing for the layer.
	AutoHealing            *bool    `pulumi:"autoHealing"`
	CustomConfigureRecipes []string `pulumi:"customConfigureRecipes"`
	CustomDeployRecipes    []string `pulumi:"customDeployRecipes"`
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn *string `pulumi:"customInstanceProfileArn"`
	// Custom JSON attributes to apply to the layer.
	CustomJson *string `pulumi:"customJson"`
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds []string `pulumi:"customSecurityGroupIds"`
	CustomSetupRecipes     []string `pulumi:"customSetupRecipes"`
	CustomShutdownRecipes  []string `pulumi:"customShutdownRecipes"`
	CustomUndeployRecipes  []string `pulumi:"customUndeployRecipes"`
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown *bool `pulumi:"drainElbOnShutdown"`
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes []NodejsAppLayerEbsVolume `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer *string `pulumi:"elasticLoadBalancer"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot *bool `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout *int `pulumi:"instanceShutdownTimeout"`
	// A human-readable name for the layer.
	Name *string `pulumi:"name"`
	// The version of NodeJS to use. Defaults to "0.10.38".
	NodejsVersion *string `pulumi:"nodejsVersion"`
	// The id of the stack the layer will belong to.
	StackId *string `pulumi:"stackId"`
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages []string `pulumi:"systemPackages"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances *bool `pulumi:"useEbsOptimizedInstances"`
}

type NodejsAppLayerState struct {
	// The Amazon Resource Name(ARN) of the layer.
	Arn pulumi.StringPtrInput
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps pulumi.BoolPtrInput
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps pulumi.BoolPtrInput
	// Whether to enable auto-healing for the layer.
	AutoHealing            pulumi.BoolPtrInput
	CustomConfigureRecipes pulumi.StringArrayInput
	CustomDeployRecipes    pulumi.StringArrayInput
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn pulumi.StringPtrInput
	// Custom JSON attributes to apply to the layer.
	CustomJson pulumi.StringPtrInput
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds pulumi.StringArrayInput
	CustomSetupRecipes     pulumi.StringArrayInput
	CustomShutdownRecipes  pulumi.StringArrayInput
	CustomUndeployRecipes  pulumi.StringArrayInput
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown pulumi.BoolPtrInput
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes NodejsAppLayerEbsVolumeArrayInput
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrInput
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrInput
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrInput
	// A human-readable name for the layer.
	Name pulumi.StringPtrInput
	// The version of NodeJS to use. Defaults to "0.10.38".
	NodejsVersion pulumi.StringPtrInput
	// The id of the stack the layer will belong to.
	StackId pulumi.StringPtrInput
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.StringArrayInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolPtrInput
}

func (NodejsAppLayerState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodejsAppLayerState)(nil)).Elem()
}

type nodejsAppLayerArgs struct {
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps *bool `pulumi:"autoAssignElasticIps"`
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps *bool `pulumi:"autoAssignPublicIps"`
	// Whether to enable auto-healing for the layer.
	AutoHealing            *bool    `pulumi:"autoHealing"`
	CustomConfigureRecipes []string `pulumi:"customConfigureRecipes"`
	CustomDeployRecipes    []string `pulumi:"customDeployRecipes"`
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn *string `pulumi:"customInstanceProfileArn"`
	// Custom JSON attributes to apply to the layer.
	CustomJson *string `pulumi:"customJson"`
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds []string `pulumi:"customSecurityGroupIds"`
	CustomSetupRecipes     []string `pulumi:"customSetupRecipes"`
	CustomShutdownRecipes  []string `pulumi:"customShutdownRecipes"`
	CustomUndeployRecipes  []string `pulumi:"customUndeployRecipes"`
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown *bool `pulumi:"drainElbOnShutdown"`
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes []NodejsAppLayerEbsVolume `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer *string `pulumi:"elasticLoadBalancer"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot *bool `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout *int `pulumi:"instanceShutdownTimeout"`
	// A human-readable name for the layer.
	Name *string `pulumi:"name"`
	// The version of NodeJS to use. Defaults to "0.10.38".
	NodejsVersion *string `pulumi:"nodejsVersion"`
	// The id of the stack the layer will belong to.
	StackId string `pulumi:"stackId"`
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages []string `pulumi:"systemPackages"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances *bool `pulumi:"useEbsOptimizedInstances"`
}

// The set of arguments for constructing a NodejsAppLayer resource.
type NodejsAppLayerArgs struct {
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps pulumi.BoolPtrInput
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps pulumi.BoolPtrInput
	// Whether to enable auto-healing for the layer.
	AutoHealing            pulumi.BoolPtrInput
	CustomConfigureRecipes pulumi.StringArrayInput
	CustomDeployRecipes    pulumi.StringArrayInput
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn pulumi.StringPtrInput
	// Custom JSON attributes to apply to the layer.
	CustomJson pulumi.StringPtrInput
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds pulumi.StringArrayInput
	CustomSetupRecipes     pulumi.StringArrayInput
	CustomShutdownRecipes  pulumi.StringArrayInput
	CustomUndeployRecipes  pulumi.StringArrayInput
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown pulumi.BoolPtrInput
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes NodejsAppLayerEbsVolumeArrayInput
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrInput
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrInput
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrInput
	// A human-readable name for the layer.
	Name pulumi.StringPtrInput
	// The version of NodeJS to use. Defaults to "0.10.38".
	NodejsVersion pulumi.StringPtrInput
	// The id of the stack the layer will belong to.
	StackId pulumi.StringInput
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.StringArrayInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolPtrInput
}

func (NodejsAppLayerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodejsAppLayerArgs)(nil)).Elem()
}

type NodejsAppLayerInput interface {
	pulumi.Input

	ToNodejsAppLayerOutput() NodejsAppLayerOutput
	ToNodejsAppLayerOutputWithContext(ctx context.Context) NodejsAppLayerOutput
}

func (*NodejsAppLayer) ElementType() reflect.Type {
	return reflect.TypeOf((*NodejsAppLayer)(nil))
}

func (i *NodejsAppLayer) ToNodejsAppLayerOutput() NodejsAppLayerOutput {
	return i.ToNodejsAppLayerOutputWithContext(context.Background())
}

func (i *NodejsAppLayer) ToNodejsAppLayerOutputWithContext(ctx context.Context) NodejsAppLayerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodejsAppLayerOutput)
}

func (i *NodejsAppLayer) ToNodejsAppLayerPtrOutput() NodejsAppLayerPtrOutput {
	return i.ToNodejsAppLayerPtrOutputWithContext(context.Background())
}

func (i *NodejsAppLayer) ToNodejsAppLayerPtrOutputWithContext(ctx context.Context) NodejsAppLayerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodejsAppLayerPtrOutput)
}

type NodejsAppLayerPtrInput interface {
	pulumi.Input

	ToNodejsAppLayerPtrOutput() NodejsAppLayerPtrOutput
	ToNodejsAppLayerPtrOutputWithContext(ctx context.Context) NodejsAppLayerPtrOutput
}

type nodejsAppLayerPtrType NodejsAppLayerArgs

func (*nodejsAppLayerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NodejsAppLayer)(nil))
}

func (i *nodejsAppLayerPtrType) ToNodejsAppLayerPtrOutput() NodejsAppLayerPtrOutput {
	return i.ToNodejsAppLayerPtrOutputWithContext(context.Background())
}

func (i *nodejsAppLayerPtrType) ToNodejsAppLayerPtrOutputWithContext(ctx context.Context) NodejsAppLayerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodejsAppLayerPtrOutput)
}

// NodejsAppLayerArrayInput is an input type that accepts NodejsAppLayerArray and NodejsAppLayerArrayOutput values.
// You can construct a concrete instance of `NodejsAppLayerArrayInput` via:
//
//          NodejsAppLayerArray{ NodejsAppLayerArgs{...} }
type NodejsAppLayerArrayInput interface {
	pulumi.Input

	ToNodejsAppLayerArrayOutput() NodejsAppLayerArrayOutput
	ToNodejsAppLayerArrayOutputWithContext(context.Context) NodejsAppLayerArrayOutput
}

type NodejsAppLayerArray []NodejsAppLayerInput

func (NodejsAppLayerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NodejsAppLayer)(nil)).Elem()
}

func (i NodejsAppLayerArray) ToNodejsAppLayerArrayOutput() NodejsAppLayerArrayOutput {
	return i.ToNodejsAppLayerArrayOutputWithContext(context.Background())
}

func (i NodejsAppLayerArray) ToNodejsAppLayerArrayOutputWithContext(ctx context.Context) NodejsAppLayerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodejsAppLayerArrayOutput)
}

// NodejsAppLayerMapInput is an input type that accepts NodejsAppLayerMap and NodejsAppLayerMapOutput values.
// You can construct a concrete instance of `NodejsAppLayerMapInput` via:
//
//          NodejsAppLayerMap{ "key": NodejsAppLayerArgs{...} }
type NodejsAppLayerMapInput interface {
	pulumi.Input

	ToNodejsAppLayerMapOutput() NodejsAppLayerMapOutput
	ToNodejsAppLayerMapOutputWithContext(context.Context) NodejsAppLayerMapOutput
}

type NodejsAppLayerMap map[string]NodejsAppLayerInput

func (NodejsAppLayerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NodejsAppLayer)(nil)).Elem()
}

func (i NodejsAppLayerMap) ToNodejsAppLayerMapOutput() NodejsAppLayerMapOutput {
	return i.ToNodejsAppLayerMapOutputWithContext(context.Background())
}

func (i NodejsAppLayerMap) ToNodejsAppLayerMapOutputWithContext(ctx context.Context) NodejsAppLayerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodejsAppLayerMapOutput)
}

type NodejsAppLayerOutput struct{ *pulumi.OutputState }

func (NodejsAppLayerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodejsAppLayer)(nil))
}

func (o NodejsAppLayerOutput) ToNodejsAppLayerOutput() NodejsAppLayerOutput {
	return o
}

func (o NodejsAppLayerOutput) ToNodejsAppLayerOutputWithContext(ctx context.Context) NodejsAppLayerOutput {
	return o
}

func (o NodejsAppLayerOutput) ToNodejsAppLayerPtrOutput() NodejsAppLayerPtrOutput {
	return o.ToNodejsAppLayerPtrOutputWithContext(context.Background())
}

func (o NodejsAppLayerOutput) ToNodejsAppLayerPtrOutputWithContext(ctx context.Context) NodejsAppLayerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NodejsAppLayer) *NodejsAppLayer {
		return &v
	}).(NodejsAppLayerPtrOutput)
}

type NodejsAppLayerPtrOutput struct{ *pulumi.OutputState }

func (NodejsAppLayerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NodejsAppLayer)(nil))
}

func (o NodejsAppLayerPtrOutput) ToNodejsAppLayerPtrOutput() NodejsAppLayerPtrOutput {
	return o
}

func (o NodejsAppLayerPtrOutput) ToNodejsAppLayerPtrOutputWithContext(ctx context.Context) NodejsAppLayerPtrOutput {
	return o
}

func (o NodejsAppLayerPtrOutput) Elem() NodejsAppLayerOutput {
	return o.ApplyT(func(v *NodejsAppLayer) NodejsAppLayer {
		if v != nil {
			return *v
		}
		var ret NodejsAppLayer
		return ret
	}).(NodejsAppLayerOutput)
}

type NodejsAppLayerArrayOutput struct{ *pulumi.OutputState }

func (NodejsAppLayerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodejsAppLayer)(nil))
}

func (o NodejsAppLayerArrayOutput) ToNodejsAppLayerArrayOutput() NodejsAppLayerArrayOutput {
	return o
}

func (o NodejsAppLayerArrayOutput) ToNodejsAppLayerArrayOutputWithContext(ctx context.Context) NodejsAppLayerArrayOutput {
	return o
}

func (o NodejsAppLayerArrayOutput) Index(i pulumi.IntInput) NodejsAppLayerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NodejsAppLayer {
		return vs[0].([]NodejsAppLayer)[vs[1].(int)]
	}).(NodejsAppLayerOutput)
}

type NodejsAppLayerMapOutput struct{ *pulumi.OutputState }

func (NodejsAppLayerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NodejsAppLayer)(nil))
}

func (o NodejsAppLayerMapOutput) ToNodejsAppLayerMapOutput() NodejsAppLayerMapOutput {
	return o
}

func (o NodejsAppLayerMapOutput) ToNodejsAppLayerMapOutputWithContext(ctx context.Context) NodejsAppLayerMapOutput {
	return o
}

func (o NodejsAppLayerMapOutput) MapIndex(k pulumi.StringInput) NodejsAppLayerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NodejsAppLayer {
		return vs[0].(map[string]NodejsAppLayer)[vs[1].(string)]
	}).(NodejsAppLayerOutput)
}

func init() {
	pulumi.RegisterOutputType(NodejsAppLayerOutput{})
	pulumi.RegisterOutputType(NodejsAppLayerPtrOutput{})
	pulumi.RegisterOutputType(NodejsAppLayerArrayOutput{})
	pulumi.RegisterOutputType(NodejsAppLayerMapOutput{})
}
