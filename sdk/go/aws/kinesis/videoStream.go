// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kinesis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Kinesis Video Stream resource. Amazon Kinesis Video Streams makes it easy to securely stream video from connected devices to AWS for analytics, machine learning (ML), playback, and other processing.
//
// For more details, see the [Amazon Kinesis Documentation](https://aws.amazon.com/documentation/kinesis/).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := kinesis.NewVideoStream(ctx, "default", &kinesis.VideoStreamArgs{
//				DataRetentionInHours: pulumi.Int(1),
//				DeviceName:           pulumi.String("kinesis-video-device-name"),
//				MediaType:            pulumi.String("video/h264"),
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("kinesis-video-stream"),
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
// Using `pulumi import`, import Kinesis Streams using the `arn`. For example:
//
// ```sh
//
//	$ pulumi import aws:kinesis/videoStream:VideoStream test_stream arn:aws:kinesisvideo:us-west-2:123456789012:stream/pulumi-kinesis-test/1554978910975
//
// ```
type VideoStream struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A time stamp that indicates when the stream was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The number of hours that you want to retain the data in the stream. Kinesis Video Streams retains the data in a data store that is associated with the stream. The default value is `0`, indicating that the stream does not persist data.
	DataRetentionInHours pulumi.IntPtrOutput `pulumi:"dataRetentionInHours"`
	// The name of the device that is writing to the stream. **In the current implementation, Kinesis Video Streams does not use this name.**
	DeviceName pulumi.StringPtrOutput `pulumi:"deviceName"`
	// The ID of the AWS Key Management Service (AWS KMS) key that you want Kinesis Video Streams to use to encrypt stream data. If no key ID is specified, the default, Kinesis Video-managed key (`aws/kinesisvideo`) is used.
	KmsKeyId pulumi.StringOutput `pulumi:"kmsKeyId"`
	// The media type of the stream. Consumers of the stream can use this information when processing the stream. For more information about media types, see [Media Types](http://www.iana.org/assignments/media-types/media-types.xhtml). If you choose to specify the MediaType, see [Naming Requirements](https://tools.ietf.org/html/rfc6838#section-4.2) for guidelines.
	MediaType pulumi.StringPtrOutput `pulumi:"mediaType"`
	// A name to identify the stream. This is unique to the
	// AWS account and region the Stream is created in.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The version of the stream.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewVideoStream registers a new resource with the given unique name, arguments, and options.
func NewVideoStream(ctx *pulumi.Context,
	name string, args *VideoStreamArgs, opts ...pulumi.ResourceOption) (*VideoStream, error) {
	if args == nil {
		args = &VideoStreamArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VideoStream
	err := ctx.RegisterResource("aws:kinesis/videoStream:VideoStream", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVideoStream gets an existing VideoStream resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVideoStream(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VideoStreamState, opts ...pulumi.ResourceOption) (*VideoStream, error) {
	var resource VideoStream
	err := ctx.ReadResource("aws:kinesis/videoStream:VideoStream", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VideoStream resources.
type videoStreamState struct {
	// The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
	Arn *string `pulumi:"arn"`
	// A time stamp that indicates when the stream was created.
	CreationTime *string `pulumi:"creationTime"`
	// The number of hours that you want to retain the data in the stream. Kinesis Video Streams retains the data in a data store that is associated with the stream. The default value is `0`, indicating that the stream does not persist data.
	DataRetentionInHours *int `pulumi:"dataRetentionInHours"`
	// The name of the device that is writing to the stream. **In the current implementation, Kinesis Video Streams does not use this name.**
	DeviceName *string `pulumi:"deviceName"`
	// The ID of the AWS Key Management Service (AWS KMS) key that you want Kinesis Video Streams to use to encrypt stream data. If no key ID is specified, the default, Kinesis Video-managed key (`aws/kinesisvideo`) is used.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The media type of the stream. Consumers of the stream can use this information when processing the stream. For more information about media types, see [Media Types](http://www.iana.org/assignments/media-types/media-types.xhtml). If you choose to specify the MediaType, see [Naming Requirements](https://tools.ietf.org/html/rfc6838#section-4.2) for guidelines.
	MediaType *string `pulumi:"mediaType"`
	// A name to identify the stream. This is unique to the
	// AWS account and region the Stream is created in.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The version of the stream.
	Version *string `pulumi:"version"`
}

type VideoStreamState struct {
	// The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
	Arn pulumi.StringPtrInput
	// A time stamp that indicates when the stream was created.
	CreationTime pulumi.StringPtrInput
	// The number of hours that you want to retain the data in the stream. Kinesis Video Streams retains the data in a data store that is associated with the stream. The default value is `0`, indicating that the stream does not persist data.
	DataRetentionInHours pulumi.IntPtrInput
	// The name of the device that is writing to the stream. **In the current implementation, Kinesis Video Streams does not use this name.**
	DeviceName pulumi.StringPtrInput
	// The ID of the AWS Key Management Service (AWS KMS) key that you want Kinesis Video Streams to use to encrypt stream data. If no key ID is specified, the default, Kinesis Video-managed key (`aws/kinesisvideo`) is used.
	KmsKeyId pulumi.StringPtrInput
	// The media type of the stream. Consumers of the stream can use this information when processing the stream. For more information about media types, see [Media Types](http://www.iana.org/assignments/media-types/media-types.xhtml). If you choose to specify the MediaType, see [Naming Requirements](https://tools.ietf.org/html/rfc6838#section-4.2) for guidelines.
	MediaType pulumi.StringPtrInput
	// A name to identify the stream. This is unique to the
	// AWS account and region the Stream is created in.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// The version of the stream.
	Version pulumi.StringPtrInput
}

func (VideoStreamState) ElementType() reflect.Type {
	return reflect.TypeOf((*videoStreamState)(nil)).Elem()
}

type videoStreamArgs struct {
	// The number of hours that you want to retain the data in the stream. Kinesis Video Streams retains the data in a data store that is associated with the stream. The default value is `0`, indicating that the stream does not persist data.
	DataRetentionInHours *int `pulumi:"dataRetentionInHours"`
	// The name of the device that is writing to the stream. **In the current implementation, Kinesis Video Streams does not use this name.**
	DeviceName *string `pulumi:"deviceName"`
	// The ID of the AWS Key Management Service (AWS KMS) key that you want Kinesis Video Streams to use to encrypt stream data. If no key ID is specified, the default, Kinesis Video-managed key (`aws/kinesisvideo`) is used.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The media type of the stream. Consumers of the stream can use this information when processing the stream. For more information about media types, see [Media Types](http://www.iana.org/assignments/media-types/media-types.xhtml). If you choose to specify the MediaType, see [Naming Requirements](https://tools.ietf.org/html/rfc6838#section-4.2) for guidelines.
	MediaType *string `pulumi:"mediaType"`
	// A name to identify the stream. This is unique to the
	// AWS account and region the Stream is created in.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a VideoStream resource.
type VideoStreamArgs struct {
	// The number of hours that you want to retain the data in the stream. Kinesis Video Streams retains the data in a data store that is associated with the stream. The default value is `0`, indicating that the stream does not persist data.
	DataRetentionInHours pulumi.IntPtrInput
	// The name of the device that is writing to the stream. **In the current implementation, Kinesis Video Streams does not use this name.**
	DeviceName pulumi.StringPtrInput
	// The ID of the AWS Key Management Service (AWS KMS) key that you want Kinesis Video Streams to use to encrypt stream data. If no key ID is specified, the default, Kinesis Video-managed key (`aws/kinesisvideo`) is used.
	KmsKeyId pulumi.StringPtrInput
	// The media type of the stream. Consumers of the stream can use this information when processing the stream. For more information about media types, see [Media Types](http://www.iana.org/assignments/media-types/media-types.xhtml). If you choose to specify the MediaType, see [Naming Requirements](https://tools.ietf.org/html/rfc6838#section-4.2) for guidelines.
	MediaType pulumi.StringPtrInput
	// A name to identify the stream. This is unique to the
	// AWS account and region the Stream is created in.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (VideoStreamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*videoStreamArgs)(nil)).Elem()
}

type VideoStreamInput interface {
	pulumi.Input

	ToVideoStreamOutput() VideoStreamOutput
	ToVideoStreamOutputWithContext(ctx context.Context) VideoStreamOutput
}

func (*VideoStream) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoStream)(nil)).Elem()
}

func (i *VideoStream) ToVideoStreamOutput() VideoStreamOutput {
	return i.ToVideoStreamOutputWithContext(context.Background())
}

func (i *VideoStream) ToVideoStreamOutputWithContext(ctx context.Context) VideoStreamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoStreamOutput)
}

// VideoStreamArrayInput is an input type that accepts VideoStreamArray and VideoStreamArrayOutput values.
// You can construct a concrete instance of `VideoStreamArrayInput` via:
//
//	VideoStreamArray{ VideoStreamArgs{...} }
type VideoStreamArrayInput interface {
	pulumi.Input

	ToVideoStreamArrayOutput() VideoStreamArrayOutput
	ToVideoStreamArrayOutputWithContext(context.Context) VideoStreamArrayOutput
}

type VideoStreamArray []VideoStreamInput

func (VideoStreamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VideoStream)(nil)).Elem()
}

func (i VideoStreamArray) ToVideoStreamArrayOutput() VideoStreamArrayOutput {
	return i.ToVideoStreamArrayOutputWithContext(context.Background())
}

func (i VideoStreamArray) ToVideoStreamArrayOutputWithContext(ctx context.Context) VideoStreamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoStreamArrayOutput)
}

// VideoStreamMapInput is an input type that accepts VideoStreamMap and VideoStreamMapOutput values.
// You can construct a concrete instance of `VideoStreamMapInput` via:
//
//	VideoStreamMap{ "key": VideoStreamArgs{...} }
type VideoStreamMapInput interface {
	pulumi.Input

	ToVideoStreamMapOutput() VideoStreamMapOutput
	ToVideoStreamMapOutputWithContext(context.Context) VideoStreamMapOutput
}

type VideoStreamMap map[string]VideoStreamInput

func (VideoStreamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VideoStream)(nil)).Elem()
}

func (i VideoStreamMap) ToVideoStreamMapOutput() VideoStreamMapOutput {
	return i.ToVideoStreamMapOutputWithContext(context.Background())
}

func (i VideoStreamMap) ToVideoStreamMapOutputWithContext(ctx context.Context) VideoStreamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VideoStreamMapOutput)
}

type VideoStreamOutput struct{ *pulumi.OutputState }

func (VideoStreamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VideoStream)(nil)).Elem()
}

func (o VideoStreamOutput) ToVideoStreamOutput() VideoStreamOutput {
	return o
}

func (o VideoStreamOutput) ToVideoStreamOutputWithContext(ctx context.Context) VideoStreamOutput {
	return o
}

// The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
func (o VideoStreamOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *VideoStream) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// A time stamp that indicates when the stream was created.
func (o VideoStreamOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VideoStream) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// The number of hours that you want to retain the data in the stream. Kinesis Video Streams retains the data in a data store that is associated with the stream. The default value is `0`, indicating that the stream does not persist data.
func (o VideoStreamOutput) DataRetentionInHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VideoStream) pulumi.IntPtrOutput { return v.DataRetentionInHours }).(pulumi.IntPtrOutput)
}

// The name of the device that is writing to the stream. **In the current implementation, Kinesis Video Streams does not use this name.**
func (o VideoStreamOutput) DeviceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoStream) pulumi.StringPtrOutput { return v.DeviceName }).(pulumi.StringPtrOutput)
}

// The ID of the AWS Key Management Service (AWS KMS) key that you want Kinesis Video Streams to use to encrypt stream data. If no key ID is specified, the default, Kinesis Video-managed key (`aws/kinesisvideo`) is used.
func (o VideoStreamOutput) KmsKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *VideoStream) pulumi.StringOutput { return v.KmsKeyId }).(pulumi.StringOutput)
}

// The media type of the stream. Consumers of the stream can use this information when processing the stream. For more information about media types, see [Media Types](http://www.iana.org/assignments/media-types/media-types.xhtml). If you choose to specify the MediaType, see [Naming Requirements](https://tools.ietf.org/html/rfc6838#section-4.2) for guidelines.
func (o VideoStreamOutput) MediaType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VideoStream) pulumi.StringPtrOutput { return v.MediaType }).(pulumi.StringPtrOutput)
}

// A name to identify the stream. This is unique to the
// AWS account and region the Stream is created in.
func (o VideoStreamOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VideoStream) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o VideoStreamOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VideoStream) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o VideoStreamOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VideoStream) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The version of the stream.
func (o VideoStreamOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *VideoStream) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type VideoStreamArrayOutput struct{ *pulumi.OutputState }

func (VideoStreamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VideoStream)(nil)).Elem()
}

func (o VideoStreamArrayOutput) ToVideoStreamArrayOutput() VideoStreamArrayOutput {
	return o
}

func (o VideoStreamArrayOutput) ToVideoStreamArrayOutputWithContext(ctx context.Context) VideoStreamArrayOutput {
	return o
}

func (o VideoStreamArrayOutput) Index(i pulumi.IntInput) VideoStreamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VideoStream {
		return vs[0].([]*VideoStream)[vs[1].(int)]
	}).(VideoStreamOutput)
}

type VideoStreamMapOutput struct{ *pulumi.OutputState }

func (VideoStreamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VideoStream)(nil)).Elem()
}

func (o VideoStreamMapOutput) ToVideoStreamMapOutput() VideoStreamMapOutput {
	return o
}

func (o VideoStreamMapOutput) ToVideoStreamMapOutputWithContext(ctx context.Context) VideoStreamMapOutput {
	return o
}

func (o VideoStreamMapOutput) MapIndex(k pulumi.StringInput) VideoStreamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VideoStream {
		return vs[0].(map[string]*VideoStream)[vs[1].(string)]
	}).(VideoStreamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VideoStreamInput)(nil)).Elem(), &VideoStream{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoStreamArrayInput)(nil)).Elem(), VideoStreamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VideoStreamMapInput)(nil)).Elem(), VideoStreamMap{})
	pulumi.RegisterOutputType(VideoStreamOutput{})
	pulumi.RegisterOutputType(VideoStreamArrayOutput{})
	pulumi.RegisterOutputType(VideoStreamMapOutput{})
}
