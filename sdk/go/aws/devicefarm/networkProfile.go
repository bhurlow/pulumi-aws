// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devicefarm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a resource to manage AWS Device Farm Network Profiles.
// ∂
// > **NOTE:** AWS currently has limited regional support for Device Farm (e.g., `us-west-2`). See [AWS Device Farm endpoints and quotas](https://docs.aws.amazon.com/general/latest/gr/devicefarm.html) for information on supported regions.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/devicefarm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleProject, err := devicefarm.NewProject(ctx, "exampleProject", nil)
//			if err != nil {
//				return err
//			}
//			_, err = devicefarm.NewNetworkProfile(ctx, "exampleNetworkProfile", &devicefarm.NetworkProfileArgs{
//				ProjectArn: exampleProject.Arn,
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
// Using `pulumi import`, import DeviceFarm Network Profiles using their ARN. For example:
//
// ```sh
//
//	$ pulumi import aws:devicefarm/networkProfile:NetworkProfile example arn:aws:devicefarm:us-west-2:123456789012:networkprofile:4fa784c7-ccb4-4dbf-ba4f-02198320daa1
//
// ```
type NetworkProfile struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name of this network profile.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the network profile.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The data throughput rate in bits per second, as an integer from `0` to `104857600`. Default value is `104857600`.
	DownlinkBandwidthBits pulumi.IntPtrOutput `pulumi:"downlinkBandwidthBits"`
	// Delay time for all packets to destination in milliseconds as an integer from `0` to `2000`.
	DownlinkDelayMs pulumi.IntPtrOutput `pulumi:"downlinkDelayMs"`
	// Time variation in the delay of received packets in milliseconds as an integer from `0` to `2000`.
	DownlinkJitterMs pulumi.IntPtrOutput `pulumi:"downlinkJitterMs"`
	// Proportion of received packets that fail to arrive from `0` to `100` percent.
	DownlinkLossPercent pulumi.IntPtrOutput `pulumi:"downlinkLossPercent"`
	// The name for the network profile.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ARN of the project for the network profile.
	ProjectArn pulumi.StringOutput `pulumi:"projectArn"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The type of network profile to create. Valid values are listed are `PRIVATE` and `CURATED`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// The data throughput rate in bits per second, as an integer from `0` to `104857600`. Default value is `104857600`.
	UplinkBandwidthBits pulumi.IntPtrOutput `pulumi:"uplinkBandwidthBits"`
	// Delay time for all packets to destination in milliseconds as an integer from `0` to `2000`.
	UplinkDelayMs pulumi.IntPtrOutput `pulumi:"uplinkDelayMs"`
	// Time variation in the delay of received packets in milliseconds as an integer from `0` to `2000`.
	UplinkJitterMs pulumi.IntPtrOutput `pulumi:"uplinkJitterMs"`
	// Proportion of received packets that fail to arrive from `0` to `100` percent.
	UplinkLossPercent pulumi.IntPtrOutput `pulumi:"uplinkLossPercent"`
}

// NewNetworkProfile registers a new resource with the given unique name, arguments, and options.
func NewNetworkProfile(ctx *pulumi.Context,
	name string, args *NetworkProfileArgs, opts ...pulumi.ResourceOption) (*NetworkProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectArn == nil {
		return nil, errors.New("invalid value for required argument 'ProjectArn'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NetworkProfile
	err := ctx.RegisterResource("aws:devicefarm/networkProfile:NetworkProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkProfile gets an existing NetworkProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkProfileState, opts ...pulumi.ResourceOption) (*NetworkProfile, error) {
	var resource NetworkProfile
	err := ctx.ReadResource("aws:devicefarm/networkProfile:NetworkProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkProfile resources.
type networkProfileState struct {
	// The Amazon Resource Name of this network profile.
	Arn *string `pulumi:"arn"`
	// The description of the network profile.
	Description *string `pulumi:"description"`
	// The data throughput rate in bits per second, as an integer from `0` to `104857600`. Default value is `104857600`.
	DownlinkBandwidthBits *int `pulumi:"downlinkBandwidthBits"`
	// Delay time for all packets to destination in milliseconds as an integer from `0` to `2000`.
	DownlinkDelayMs *int `pulumi:"downlinkDelayMs"`
	// Time variation in the delay of received packets in milliseconds as an integer from `0` to `2000`.
	DownlinkJitterMs *int `pulumi:"downlinkJitterMs"`
	// Proportion of received packets that fail to arrive from `0` to `100` percent.
	DownlinkLossPercent *int `pulumi:"downlinkLossPercent"`
	// The name for the network profile.
	Name *string `pulumi:"name"`
	// The ARN of the project for the network profile.
	ProjectArn *string `pulumi:"projectArn"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The type of network profile to create. Valid values are listed are `PRIVATE` and `CURATED`.
	Type *string `pulumi:"type"`
	// The data throughput rate in bits per second, as an integer from `0` to `104857600`. Default value is `104857600`.
	UplinkBandwidthBits *int `pulumi:"uplinkBandwidthBits"`
	// Delay time for all packets to destination in milliseconds as an integer from `0` to `2000`.
	UplinkDelayMs *int `pulumi:"uplinkDelayMs"`
	// Time variation in the delay of received packets in milliseconds as an integer from `0` to `2000`.
	UplinkJitterMs *int `pulumi:"uplinkJitterMs"`
	// Proportion of received packets that fail to arrive from `0` to `100` percent.
	UplinkLossPercent *int `pulumi:"uplinkLossPercent"`
}

type NetworkProfileState struct {
	// The Amazon Resource Name of this network profile.
	Arn pulumi.StringPtrInput
	// The description of the network profile.
	Description pulumi.StringPtrInput
	// The data throughput rate in bits per second, as an integer from `0` to `104857600`. Default value is `104857600`.
	DownlinkBandwidthBits pulumi.IntPtrInput
	// Delay time for all packets to destination in milliseconds as an integer from `0` to `2000`.
	DownlinkDelayMs pulumi.IntPtrInput
	// Time variation in the delay of received packets in milliseconds as an integer from `0` to `2000`.
	DownlinkJitterMs pulumi.IntPtrInput
	// Proportion of received packets that fail to arrive from `0` to `100` percent.
	DownlinkLossPercent pulumi.IntPtrInput
	// The name for the network profile.
	Name pulumi.StringPtrInput
	// The ARN of the project for the network profile.
	ProjectArn pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// The type of network profile to create. Valid values are listed are `PRIVATE` and `CURATED`.
	Type pulumi.StringPtrInput
	// The data throughput rate in bits per second, as an integer from `0` to `104857600`. Default value is `104857600`.
	UplinkBandwidthBits pulumi.IntPtrInput
	// Delay time for all packets to destination in milliseconds as an integer from `0` to `2000`.
	UplinkDelayMs pulumi.IntPtrInput
	// Time variation in the delay of received packets in milliseconds as an integer from `0` to `2000`.
	UplinkJitterMs pulumi.IntPtrInput
	// Proportion of received packets that fail to arrive from `0` to `100` percent.
	UplinkLossPercent pulumi.IntPtrInput
}

func (NetworkProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkProfileState)(nil)).Elem()
}

type networkProfileArgs struct {
	// The description of the network profile.
	Description *string `pulumi:"description"`
	// The data throughput rate in bits per second, as an integer from `0` to `104857600`. Default value is `104857600`.
	DownlinkBandwidthBits *int `pulumi:"downlinkBandwidthBits"`
	// Delay time for all packets to destination in milliseconds as an integer from `0` to `2000`.
	DownlinkDelayMs *int `pulumi:"downlinkDelayMs"`
	// Time variation in the delay of received packets in milliseconds as an integer from `0` to `2000`.
	DownlinkJitterMs *int `pulumi:"downlinkJitterMs"`
	// Proportion of received packets that fail to arrive from `0` to `100` percent.
	DownlinkLossPercent *int `pulumi:"downlinkLossPercent"`
	// The name for the network profile.
	Name *string `pulumi:"name"`
	// The ARN of the project for the network profile.
	ProjectArn string `pulumi:"projectArn"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The type of network profile to create. Valid values are listed are `PRIVATE` and `CURATED`.
	Type *string `pulumi:"type"`
	// The data throughput rate in bits per second, as an integer from `0` to `104857600`. Default value is `104857600`.
	UplinkBandwidthBits *int `pulumi:"uplinkBandwidthBits"`
	// Delay time for all packets to destination in milliseconds as an integer from `0` to `2000`.
	UplinkDelayMs *int `pulumi:"uplinkDelayMs"`
	// Time variation in the delay of received packets in milliseconds as an integer from `0` to `2000`.
	UplinkJitterMs *int `pulumi:"uplinkJitterMs"`
	// Proportion of received packets that fail to arrive from `0` to `100` percent.
	UplinkLossPercent *int `pulumi:"uplinkLossPercent"`
}

// The set of arguments for constructing a NetworkProfile resource.
type NetworkProfileArgs struct {
	// The description of the network profile.
	Description pulumi.StringPtrInput
	// The data throughput rate in bits per second, as an integer from `0` to `104857600`. Default value is `104857600`.
	DownlinkBandwidthBits pulumi.IntPtrInput
	// Delay time for all packets to destination in milliseconds as an integer from `0` to `2000`.
	DownlinkDelayMs pulumi.IntPtrInput
	// Time variation in the delay of received packets in milliseconds as an integer from `0` to `2000`.
	DownlinkJitterMs pulumi.IntPtrInput
	// Proportion of received packets that fail to arrive from `0` to `100` percent.
	DownlinkLossPercent pulumi.IntPtrInput
	// The name for the network profile.
	Name pulumi.StringPtrInput
	// The ARN of the project for the network profile.
	ProjectArn pulumi.StringInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The type of network profile to create. Valid values are listed are `PRIVATE` and `CURATED`.
	Type pulumi.StringPtrInput
	// The data throughput rate in bits per second, as an integer from `0` to `104857600`. Default value is `104857600`.
	UplinkBandwidthBits pulumi.IntPtrInput
	// Delay time for all packets to destination in milliseconds as an integer from `0` to `2000`.
	UplinkDelayMs pulumi.IntPtrInput
	// Time variation in the delay of received packets in milliseconds as an integer from `0` to `2000`.
	UplinkJitterMs pulumi.IntPtrInput
	// Proportion of received packets that fail to arrive from `0` to `100` percent.
	UplinkLossPercent pulumi.IntPtrInput
}

func (NetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkProfileArgs)(nil)).Elem()
}

type NetworkProfileInput interface {
	pulumi.Input

	ToNetworkProfileOutput() NetworkProfileOutput
	ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput
}

func (*NetworkProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (i *NetworkProfile) ToNetworkProfileOutput() NetworkProfileOutput {
	return i.ToNetworkProfileOutputWithContext(context.Background())
}

func (i *NetworkProfile) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput)
}

func (i *NetworkProfile) ToOutput(ctx context.Context) pulumix.Output[*NetworkProfile] {
	return pulumix.Output[*NetworkProfile]{
		OutputState: i.ToNetworkProfileOutputWithContext(ctx).OutputState,
	}
}

// NetworkProfileArrayInput is an input type that accepts NetworkProfileArray and NetworkProfileArrayOutput values.
// You can construct a concrete instance of `NetworkProfileArrayInput` via:
//
//	NetworkProfileArray{ NetworkProfileArgs{...} }
type NetworkProfileArrayInput interface {
	pulumi.Input

	ToNetworkProfileArrayOutput() NetworkProfileArrayOutput
	ToNetworkProfileArrayOutputWithContext(context.Context) NetworkProfileArrayOutput
}

type NetworkProfileArray []NetworkProfileInput

func (NetworkProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkProfile)(nil)).Elem()
}

func (i NetworkProfileArray) ToNetworkProfileArrayOutput() NetworkProfileArrayOutput {
	return i.ToNetworkProfileArrayOutputWithContext(context.Background())
}

func (i NetworkProfileArray) ToNetworkProfileArrayOutputWithContext(ctx context.Context) NetworkProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileArrayOutput)
}

func (i NetworkProfileArray) ToOutput(ctx context.Context) pulumix.Output[[]*NetworkProfile] {
	return pulumix.Output[[]*NetworkProfile]{
		OutputState: i.ToNetworkProfileArrayOutputWithContext(ctx).OutputState,
	}
}

// NetworkProfileMapInput is an input type that accepts NetworkProfileMap and NetworkProfileMapOutput values.
// You can construct a concrete instance of `NetworkProfileMapInput` via:
//
//	NetworkProfileMap{ "key": NetworkProfileArgs{...} }
type NetworkProfileMapInput interface {
	pulumi.Input

	ToNetworkProfileMapOutput() NetworkProfileMapOutput
	ToNetworkProfileMapOutputWithContext(context.Context) NetworkProfileMapOutput
}

type NetworkProfileMap map[string]NetworkProfileInput

func (NetworkProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkProfile)(nil)).Elem()
}

func (i NetworkProfileMap) ToNetworkProfileMapOutput() NetworkProfileMapOutput {
	return i.ToNetworkProfileMapOutputWithContext(context.Background())
}

func (i NetworkProfileMap) ToNetworkProfileMapOutputWithContext(ctx context.Context) NetworkProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileMapOutput)
}

func (i NetworkProfileMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*NetworkProfile] {
	return pulumix.Output[map[string]*NetworkProfile]{
		OutputState: i.ToNetworkProfileMapOutputWithContext(ctx).OutputState,
	}
}

type NetworkProfileOutput struct{ *pulumi.OutputState }

func (NetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (o NetworkProfileOutput) ToNetworkProfileOutput() NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToOutput(ctx context.Context) pulumix.Output[*NetworkProfile] {
	return pulumix.Output[*NetworkProfile]{
		OutputState: o.OutputState,
	}
}

// The Amazon Resource Name of this network profile.
func (o NetworkProfileOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The description of the network profile.
func (o NetworkProfileOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The data throughput rate in bits per second, as an integer from `0` to `104857600`. Default value is `104857600`.
func (o NetworkProfileOutput) DownlinkBandwidthBits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.IntPtrOutput { return v.DownlinkBandwidthBits }).(pulumi.IntPtrOutput)
}

// Delay time for all packets to destination in milliseconds as an integer from `0` to `2000`.
func (o NetworkProfileOutput) DownlinkDelayMs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.IntPtrOutput { return v.DownlinkDelayMs }).(pulumi.IntPtrOutput)
}

// Time variation in the delay of received packets in milliseconds as an integer from `0` to `2000`.
func (o NetworkProfileOutput) DownlinkJitterMs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.IntPtrOutput { return v.DownlinkJitterMs }).(pulumi.IntPtrOutput)
}

// Proportion of received packets that fail to arrive from `0` to `100` percent.
func (o NetworkProfileOutput) DownlinkLossPercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.IntPtrOutput { return v.DownlinkLossPercent }).(pulumi.IntPtrOutput)
}

// The name for the network profile.
func (o NetworkProfileOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ARN of the project for the network profile.
func (o NetworkProfileOutput) ProjectArn() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringOutput { return v.ProjectArn }).(pulumi.StringOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o NetworkProfileOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o NetworkProfileOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The type of network profile to create. Valid values are listed are `PRIVATE` and `CURATED`.
func (o NetworkProfileOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

// The data throughput rate in bits per second, as an integer from `0` to `104857600`. Default value is `104857600`.
func (o NetworkProfileOutput) UplinkBandwidthBits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.IntPtrOutput { return v.UplinkBandwidthBits }).(pulumi.IntPtrOutput)
}

// Delay time for all packets to destination in milliseconds as an integer from `0` to `2000`.
func (o NetworkProfileOutput) UplinkDelayMs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.IntPtrOutput { return v.UplinkDelayMs }).(pulumi.IntPtrOutput)
}

// Time variation in the delay of received packets in milliseconds as an integer from `0` to `2000`.
func (o NetworkProfileOutput) UplinkJitterMs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.IntPtrOutput { return v.UplinkJitterMs }).(pulumi.IntPtrOutput)
}

// Proportion of received packets that fail to arrive from `0` to `100` percent.
func (o NetworkProfileOutput) UplinkLossPercent() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NetworkProfile) pulumi.IntPtrOutput { return v.UplinkLossPercent }).(pulumi.IntPtrOutput)
}

type NetworkProfileArrayOutput struct{ *pulumi.OutputState }

func (NetworkProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkProfile)(nil)).Elem()
}

func (o NetworkProfileArrayOutput) ToNetworkProfileArrayOutput() NetworkProfileArrayOutput {
	return o
}

func (o NetworkProfileArrayOutput) ToNetworkProfileArrayOutputWithContext(ctx context.Context) NetworkProfileArrayOutput {
	return o
}

func (o NetworkProfileArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*NetworkProfile] {
	return pulumix.Output[[]*NetworkProfile]{
		OutputState: o.OutputState,
	}
}

func (o NetworkProfileArrayOutput) Index(i pulumi.IntInput) NetworkProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkProfile {
		return vs[0].([]*NetworkProfile)[vs[1].(int)]
	}).(NetworkProfileOutput)
}

type NetworkProfileMapOutput struct{ *pulumi.OutputState }

func (NetworkProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkProfile)(nil)).Elem()
}

func (o NetworkProfileMapOutput) ToNetworkProfileMapOutput() NetworkProfileMapOutput {
	return o
}

func (o NetworkProfileMapOutput) ToNetworkProfileMapOutputWithContext(ctx context.Context) NetworkProfileMapOutput {
	return o
}

func (o NetworkProfileMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*NetworkProfile] {
	return pulumix.Output[map[string]*NetworkProfile]{
		OutputState: o.OutputState,
	}
}

func (o NetworkProfileMapOutput) MapIndex(k pulumi.StringInput) NetworkProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkProfile {
		return vs[0].(map[string]*NetworkProfile)[vs[1].(string)]
	}).(NetworkProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkProfileInput)(nil)).Elem(), &NetworkProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkProfileArrayInput)(nil)).Elem(), NetworkProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkProfileMapInput)(nil)).Elem(), NetworkProfileMap{})
	pulumi.RegisterOutputType(NetworkProfileOutput{})
	pulumi.RegisterOutputType(NetworkProfileArrayOutput{})
	pulumi.RegisterOutputType(NetworkProfileMapOutput{})
}
