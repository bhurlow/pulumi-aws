// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ivs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource for managing an AWS IVS (Interactive Video) Playback Key Pair.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"os"
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ivs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := os.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ivs.NewPlaybackKeyPair(ctx, "example", &ivs.PlaybackKeyPairArgs{
//				PublicKey: readFileOrPanic("./public-key.pem"),
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
// Using `pulumi import`, import IVS (Interactive Video) Playback Key Pair using the ARN. For example:
//
// ```sh
//
//	$ pulumi import aws:ivs/playbackKeyPair:PlaybackKeyPair example arn:aws:ivs:us-west-2:326937407773:playback-key/KDJRJNQhiQzA
//
// ```
type PlaybackKeyPair struct {
	pulumi.CustomResourceState

	// ARN of the Playback Key Pair.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Key-pair identifier.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// Playback Key Pair name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Public portion of a customer-generated key pair. Must be an ECDSA public key in PEM format.
	//
	// The following arguments are optional:
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewPlaybackKeyPair registers a new resource with the given unique name, arguments, and options.
func NewPlaybackKeyPair(ctx *pulumi.Context,
	name string, args *PlaybackKeyPairArgs, opts ...pulumi.ResourceOption) (*PlaybackKeyPair, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PublicKey == nil {
		return nil, errors.New("invalid value for required argument 'PublicKey'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PlaybackKeyPair
	err := ctx.RegisterResource("aws:ivs/playbackKeyPair:PlaybackKeyPair", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPlaybackKeyPair gets an existing PlaybackKeyPair resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPlaybackKeyPair(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PlaybackKeyPairState, opts ...pulumi.ResourceOption) (*PlaybackKeyPair, error) {
	var resource PlaybackKeyPair
	err := ctx.ReadResource("aws:ivs/playbackKeyPair:PlaybackKeyPair", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PlaybackKeyPair resources.
type playbackKeyPairState struct {
	// ARN of the Playback Key Pair.
	Arn *string `pulumi:"arn"`
	// Key-pair identifier.
	Fingerprint *string `pulumi:"fingerprint"`
	// Playback Key Pair name.
	Name *string `pulumi:"name"`
	// Public portion of a customer-generated key pair. Must be an ECDSA public key in PEM format.
	//
	// The following arguments are optional:
	PublicKey *string `pulumi:"publicKey"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type PlaybackKeyPairState struct {
	// ARN of the Playback Key Pair.
	Arn pulumi.StringPtrInput
	// Key-pair identifier.
	Fingerprint pulumi.StringPtrInput
	// Playback Key Pair name.
	Name pulumi.StringPtrInput
	// Public portion of a customer-generated key pair. Must be an ECDSA public key in PEM format.
	//
	// The following arguments are optional:
	PublicKey pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (PlaybackKeyPairState) ElementType() reflect.Type {
	return reflect.TypeOf((*playbackKeyPairState)(nil)).Elem()
}

type playbackKeyPairArgs struct {
	// Playback Key Pair name.
	Name *string `pulumi:"name"`
	// Public portion of a customer-generated key pair. Must be an ECDSA public key in PEM format.
	//
	// The following arguments are optional:
	PublicKey string `pulumi:"publicKey"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PlaybackKeyPair resource.
type PlaybackKeyPairArgs struct {
	// Playback Key Pair name.
	Name pulumi.StringPtrInput
	// Public portion of a customer-generated key pair. Must be an ECDSA public key in PEM format.
	//
	// The following arguments are optional:
	PublicKey pulumi.StringInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (PlaybackKeyPairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*playbackKeyPairArgs)(nil)).Elem()
}

type PlaybackKeyPairInput interface {
	pulumi.Input

	ToPlaybackKeyPairOutput() PlaybackKeyPairOutput
	ToPlaybackKeyPairOutputWithContext(ctx context.Context) PlaybackKeyPairOutput
}

func (*PlaybackKeyPair) ElementType() reflect.Type {
	return reflect.TypeOf((**PlaybackKeyPair)(nil)).Elem()
}

func (i *PlaybackKeyPair) ToPlaybackKeyPairOutput() PlaybackKeyPairOutput {
	return i.ToPlaybackKeyPairOutputWithContext(context.Background())
}

func (i *PlaybackKeyPair) ToPlaybackKeyPairOutputWithContext(ctx context.Context) PlaybackKeyPairOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlaybackKeyPairOutput)
}

func (i *PlaybackKeyPair) ToOutput(ctx context.Context) pulumix.Output[*PlaybackKeyPair] {
	return pulumix.Output[*PlaybackKeyPair]{
		OutputState: i.ToPlaybackKeyPairOutputWithContext(ctx).OutputState,
	}
}

// PlaybackKeyPairArrayInput is an input type that accepts PlaybackKeyPairArray and PlaybackKeyPairArrayOutput values.
// You can construct a concrete instance of `PlaybackKeyPairArrayInput` via:
//
//	PlaybackKeyPairArray{ PlaybackKeyPairArgs{...} }
type PlaybackKeyPairArrayInput interface {
	pulumi.Input

	ToPlaybackKeyPairArrayOutput() PlaybackKeyPairArrayOutput
	ToPlaybackKeyPairArrayOutputWithContext(context.Context) PlaybackKeyPairArrayOutput
}

type PlaybackKeyPairArray []PlaybackKeyPairInput

func (PlaybackKeyPairArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PlaybackKeyPair)(nil)).Elem()
}

func (i PlaybackKeyPairArray) ToPlaybackKeyPairArrayOutput() PlaybackKeyPairArrayOutput {
	return i.ToPlaybackKeyPairArrayOutputWithContext(context.Background())
}

func (i PlaybackKeyPairArray) ToPlaybackKeyPairArrayOutputWithContext(ctx context.Context) PlaybackKeyPairArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlaybackKeyPairArrayOutput)
}

func (i PlaybackKeyPairArray) ToOutput(ctx context.Context) pulumix.Output[[]*PlaybackKeyPair] {
	return pulumix.Output[[]*PlaybackKeyPair]{
		OutputState: i.ToPlaybackKeyPairArrayOutputWithContext(ctx).OutputState,
	}
}

// PlaybackKeyPairMapInput is an input type that accepts PlaybackKeyPairMap and PlaybackKeyPairMapOutput values.
// You can construct a concrete instance of `PlaybackKeyPairMapInput` via:
//
//	PlaybackKeyPairMap{ "key": PlaybackKeyPairArgs{...} }
type PlaybackKeyPairMapInput interface {
	pulumi.Input

	ToPlaybackKeyPairMapOutput() PlaybackKeyPairMapOutput
	ToPlaybackKeyPairMapOutputWithContext(context.Context) PlaybackKeyPairMapOutput
}

type PlaybackKeyPairMap map[string]PlaybackKeyPairInput

func (PlaybackKeyPairMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PlaybackKeyPair)(nil)).Elem()
}

func (i PlaybackKeyPairMap) ToPlaybackKeyPairMapOutput() PlaybackKeyPairMapOutput {
	return i.ToPlaybackKeyPairMapOutputWithContext(context.Background())
}

func (i PlaybackKeyPairMap) ToPlaybackKeyPairMapOutputWithContext(ctx context.Context) PlaybackKeyPairMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlaybackKeyPairMapOutput)
}

func (i PlaybackKeyPairMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*PlaybackKeyPair] {
	return pulumix.Output[map[string]*PlaybackKeyPair]{
		OutputState: i.ToPlaybackKeyPairMapOutputWithContext(ctx).OutputState,
	}
}

type PlaybackKeyPairOutput struct{ *pulumi.OutputState }

func (PlaybackKeyPairOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlaybackKeyPair)(nil)).Elem()
}

func (o PlaybackKeyPairOutput) ToPlaybackKeyPairOutput() PlaybackKeyPairOutput {
	return o
}

func (o PlaybackKeyPairOutput) ToPlaybackKeyPairOutputWithContext(ctx context.Context) PlaybackKeyPairOutput {
	return o
}

func (o PlaybackKeyPairOutput) ToOutput(ctx context.Context) pulumix.Output[*PlaybackKeyPair] {
	return pulumix.Output[*PlaybackKeyPair]{
		OutputState: o.OutputState,
	}
}

// ARN of the Playback Key Pair.
func (o PlaybackKeyPairOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *PlaybackKeyPair) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Key-pair identifier.
func (o PlaybackKeyPairOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *PlaybackKeyPair) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// Playback Key Pair name.
func (o PlaybackKeyPairOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PlaybackKeyPair) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Public portion of a customer-generated key pair. Must be an ECDSA public key in PEM format.
//
// The following arguments are optional:
func (o PlaybackKeyPairOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *PlaybackKeyPair) pulumi.StringOutput { return v.PublicKey }).(pulumi.StringOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o PlaybackKeyPairOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PlaybackKeyPair) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o PlaybackKeyPairOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PlaybackKeyPair) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type PlaybackKeyPairArrayOutput struct{ *pulumi.OutputState }

func (PlaybackKeyPairArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PlaybackKeyPair)(nil)).Elem()
}

func (o PlaybackKeyPairArrayOutput) ToPlaybackKeyPairArrayOutput() PlaybackKeyPairArrayOutput {
	return o
}

func (o PlaybackKeyPairArrayOutput) ToPlaybackKeyPairArrayOutputWithContext(ctx context.Context) PlaybackKeyPairArrayOutput {
	return o
}

func (o PlaybackKeyPairArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*PlaybackKeyPair] {
	return pulumix.Output[[]*PlaybackKeyPair]{
		OutputState: o.OutputState,
	}
}

func (o PlaybackKeyPairArrayOutput) Index(i pulumi.IntInput) PlaybackKeyPairOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PlaybackKeyPair {
		return vs[0].([]*PlaybackKeyPair)[vs[1].(int)]
	}).(PlaybackKeyPairOutput)
}

type PlaybackKeyPairMapOutput struct{ *pulumi.OutputState }

func (PlaybackKeyPairMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PlaybackKeyPair)(nil)).Elem()
}

func (o PlaybackKeyPairMapOutput) ToPlaybackKeyPairMapOutput() PlaybackKeyPairMapOutput {
	return o
}

func (o PlaybackKeyPairMapOutput) ToPlaybackKeyPairMapOutputWithContext(ctx context.Context) PlaybackKeyPairMapOutput {
	return o
}

func (o PlaybackKeyPairMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*PlaybackKeyPair] {
	return pulumix.Output[map[string]*PlaybackKeyPair]{
		OutputState: o.OutputState,
	}
}

func (o PlaybackKeyPairMapOutput) MapIndex(k pulumi.StringInput) PlaybackKeyPairOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PlaybackKeyPair {
		return vs[0].(map[string]*PlaybackKeyPair)[vs[1].(string)]
	}).(PlaybackKeyPairOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PlaybackKeyPairInput)(nil)).Elem(), &PlaybackKeyPair{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlaybackKeyPairArrayInput)(nil)).Elem(), PlaybackKeyPairArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlaybackKeyPairMapInput)(nil)).Elem(), PlaybackKeyPairMap{})
	pulumi.RegisterOutputType(PlaybackKeyPairOutput{})
	pulumi.RegisterOutputType(PlaybackKeyPairArrayOutput{})
	pulumi.RegisterOutputType(PlaybackKeyPairMapOutput{})
}
