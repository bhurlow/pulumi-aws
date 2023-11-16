// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ivs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for managing an AWS IVS (Interactive Video) Channel.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ivs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ivs.NewChannel(ctx, "example", nil)
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
// Using `pulumi import`, import IVS (Interactive Video) Channel using the ARN. For example:
//
// ```sh
//
//	$ pulumi import aws:ivs/channel:Channel example arn:aws:ivs:us-west-2:326937407773:channel/0Y1lcs4U7jk5
//
// ```
type Channel struct {
	pulumi.CustomResourceState

	// ARN of the Channel.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// If `true`, channel is private (enabled for playback authorization).
	Authorized pulumi.BoolOutput `pulumi:"authorized"`
	// Channel ingest endpoint, part of the definition of an ingest server, used when setting up streaming software.
	IngestEndpoint pulumi.StringOutput `pulumi:"ingestEndpoint"`
	// Channel latency mode. Valid values: `NORMAL`, `LOW`.
	LatencyMode pulumi.StringOutput `pulumi:"latencyMode"`
	// Channel name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Channel playback URL.
	PlaybackUrl pulumi.StringOutput `pulumi:"playbackUrl"`
	// Recording configuration ARN.
	RecordingConfigurationArn pulumi.StringOutput `pulumi:"recordingConfigurationArn"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Channel type, which determines the allowable resolution and bitrate. Valid values: `STANDARD`, `BASIC`.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewChannel registers a new resource with the given unique name, arguments, and options.
func NewChannel(ctx *pulumi.Context,
	name string, args *ChannelArgs, opts ...pulumi.ResourceOption) (*Channel, error) {
	if args == nil {
		args = &ChannelArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Channel
	err := ctx.RegisterResource("aws:ivs/channel:Channel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChannel gets an existing Channel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChannelState, opts ...pulumi.ResourceOption) (*Channel, error) {
	var resource Channel
	err := ctx.ReadResource("aws:ivs/channel:Channel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Channel resources.
type channelState struct {
	// ARN of the Channel.
	Arn *string `pulumi:"arn"`
	// If `true`, channel is private (enabled for playback authorization).
	Authorized *bool `pulumi:"authorized"`
	// Channel ingest endpoint, part of the definition of an ingest server, used when setting up streaming software.
	IngestEndpoint *string `pulumi:"ingestEndpoint"`
	// Channel latency mode. Valid values: `NORMAL`, `LOW`.
	LatencyMode *string `pulumi:"latencyMode"`
	// Channel name.
	Name *string `pulumi:"name"`
	// Channel playback URL.
	PlaybackUrl *string `pulumi:"playbackUrl"`
	// Recording configuration ARN.
	RecordingConfigurationArn *string `pulumi:"recordingConfigurationArn"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Channel type, which determines the allowable resolution and bitrate. Valid values: `STANDARD`, `BASIC`.
	Type *string `pulumi:"type"`
}

type ChannelState struct {
	// ARN of the Channel.
	Arn pulumi.StringPtrInput
	// If `true`, channel is private (enabled for playback authorization).
	Authorized pulumi.BoolPtrInput
	// Channel ingest endpoint, part of the definition of an ingest server, used when setting up streaming software.
	IngestEndpoint pulumi.StringPtrInput
	// Channel latency mode. Valid values: `NORMAL`, `LOW`.
	LatencyMode pulumi.StringPtrInput
	// Channel name.
	Name pulumi.StringPtrInput
	// Channel playback URL.
	PlaybackUrl pulumi.StringPtrInput
	// Recording configuration ARN.
	RecordingConfigurationArn pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// Channel type, which determines the allowable resolution and bitrate. Valid values: `STANDARD`, `BASIC`.
	Type pulumi.StringPtrInput
}

func (ChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*channelState)(nil)).Elem()
}

type channelArgs struct {
	// If `true`, channel is private (enabled for playback authorization).
	Authorized *bool `pulumi:"authorized"`
	// Channel latency mode. Valid values: `NORMAL`, `LOW`.
	LatencyMode *string `pulumi:"latencyMode"`
	// Channel name.
	Name *string `pulumi:"name"`
	// Recording configuration ARN.
	RecordingConfigurationArn *string `pulumi:"recordingConfigurationArn"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Channel type, which determines the allowable resolution and bitrate. Valid values: `STANDARD`, `BASIC`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Channel resource.
type ChannelArgs struct {
	// If `true`, channel is private (enabled for playback authorization).
	Authorized pulumi.BoolPtrInput
	// Channel latency mode. Valid values: `NORMAL`, `LOW`.
	LatencyMode pulumi.StringPtrInput
	// Channel name.
	Name pulumi.StringPtrInput
	// Recording configuration ARN.
	RecordingConfigurationArn pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Channel type, which determines the allowable resolution and bitrate. Valid values: `STANDARD`, `BASIC`.
	Type pulumi.StringPtrInput
}

func (ChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*channelArgs)(nil)).Elem()
}

type ChannelInput interface {
	pulumi.Input

	ToChannelOutput() ChannelOutput
	ToChannelOutputWithContext(ctx context.Context) ChannelOutput
}

func (*Channel) ElementType() reflect.Type {
	return reflect.TypeOf((**Channel)(nil)).Elem()
}

func (i *Channel) ToChannelOutput() ChannelOutput {
	return i.ToChannelOutputWithContext(context.Background())
}

func (i *Channel) ToChannelOutputWithContext(ctx context.Context) ChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelOutput)
}

// ChannelArrayInput is an input type that accepts ChannelArray and ChannelArrayOutput values.
// You can construct a concrete instance of `ChannelArrayInput` via:
//
//	ChannelArray{ ChannelArgs{...} }
type ChannelArrayInput interface {
	pulumi.Input

	ToChannelArrayOutput() ChannelArrayOutput
	ToChannelArrayOutputWithContext(context.Context) ChannelArrayOutput
}

type ChannelArray []ChannelInput

func (ChannelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Channel)(nil)).Elem()
}

func (i ChannelArray) ToChannelArrayOutput() ChannelArrayOutput {
	return i.ToChannelArrayOutputWithContext(context.Background())
}

func (i ChannelArray) ToChannelArrayOutputWithContext(ctx context.Context) ChannelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelArrayOutput)
}

// ChannelMapInput is an input type that accepts ChannelMap and ChannelMapOutput values.
// You can construct a concrete instance of `ChannelMapInput` via:
//
//	ChannelMap{ "key": ChannelArgs{...} }
type ChannelMapInput interface {
	pulumi.Input

	ToChannelMapOutput() ChannelMapOutput
	ToChannelMapOutputWithContext(context.Context) ChannelMapOutput
}

type ChannelMap map[string]ChannelInput

func (ChannelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Channel)(nil)).Elem()
}

func (i ChannelMap) ToChannelMapOutput() ChannelMapOutput {
	return i.ToChannelMapOutputWithContext(context.Background())
}

func (i ChannelMap) ToChannelMapOutputWithContext(ctx context.Context) ChannelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelMapOutput)
}

type ChannelOutput struct{ *pulumi.OutputState }

func (ChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Channel)(nil)).Elem()
}

func (o ChannelOutput) ToChannelOutput() ChannelOutput {
	return o
}

func (o ChannelOutput) ToChannelOutputWithContext(ctx context.Context) ChannelOutput {
	return o
}

// ARN of the Channel.
func (o ChannelOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// If `true`, channel is private (enabled for playback authorization).
func (o ChannelOutput) Authorized() pulumi.BoolOutput {
	return o.ApplyT(func(v *Channel) pulumi.BoolOutput { return v.Authorized }).(pulumi.BoolOutput)
}

// Channel ingest endpoint, part of the definition of an ingest server, used when setting up streaming software.
func (o ChannelOutput) IngestEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.IngestEndpoint }).(pulumi.StringOutput)
}

// Channel latency mode. Valid values: `NORMAL`, `LOW`.
func (o ChannelOutput) LatencyMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.LatencyMode }).(pulumi.StringOutput)
}

// Channel name.
func (o ChannelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Channel playback URL.
func (o ChannelOutput) PlaybackUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.PlaybackUrl }).(pulumi.StringOutput)
}

// Recording configuration ARN.
func (o ChannelOutput) RecordingConfigurationArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.RecordingConfigurationArn }).(pulumi.StringOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ChannelOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o ChannelOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Channel type, which determines the allowable resolution and bitrate. Valid values: `STANDARD`, `BASIC`.
func (o ChannelOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type ChannelArrayOutput struct{ *pulumi.OutputState }

func (ChannelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Channel)(nil)).Elem()
}

func (o ChannelArrayOutput) ToChannelArrayOutput() ChannelArrayOutput {
	return o
}

func (o ChannelArrayOutput) ToChannelArrayOutputWithContext(ctx context.Context) ChannelArrayOutput {
	return o
}

func (o ChannelArrayOutput) Index(i pulumi.IntInput) ChannelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Channel {
		return vs[0].([]*Channel)[vs[1].(int)]
	}).(ChannelOutput)
}

type ChannelMapOutput struct{ *pulumi.OutputState }

func (ChannelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Channel)(nil)).Elem()
}

func (o ChannelMapOutput) ToChannelMapOutput() ChannelMapOutput {
	return o
}

func (o ChannelMapOutput) ToChannelMapOutputWithContext(ctx context.Context) ChannelMapOutput {
	return o
}

func (o ChannelMapOutput) MapIndex(k pulumi.StringInput) ChannelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Channel {
		return vs[0].(map[string]*Channel)[vs[1].(string)]
	}).(ChannelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ChannelInput)(nil)).Elem(), &Channel{})
	pulumi.RegisterInputType(reflect.TypeOf((*ChannelArrayInput)(nil)).Elem(), ChannelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ChannelMapInput)(nil)).Elem(), ChannelMap{})
	pulumi.RegisterOutputType(ChannelOutput{})
	pulumi.RegisterOutputType(ChannelArrayOutput{})
	pulumi.RegisterOutputType(ChannelMapOutput{})
}
