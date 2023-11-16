// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package transcribe

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for managing an AWS Transcribe MedicalVocabulary.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/transcribe"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", &s3.BucketV2Args{
//				ForceDestroy: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			object, err := s3.NewBucketObjectv2(ctx, "object", &s3.BucketObjectv2Args{
//				Bucket: exampleBucketV2.ID(),
//				Key:    pulumi.String("transcribe/test1.txt"),
//				Source: pulumi.NewFileAsset("test.txt"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = transcribe.NewMedicalVocabulary(ctx, "exampleMedicalVocabulary", &transcribe.MedicalVocabularyArgs{
//				VocabularyName: pulumi.String("example"),
//				LanguageCode:   pulumi.String("en-US"),
//				VocabularyFileUri: pulumi.All(exampleBucketV2.ID(), object.Key).ApplyT(func(_args []interface{}) (string, error) {
//					id := _args[0].(string)
//					key := _args[1].(string)
//					return fmt.Sprintf("s3://%v/%v", id, key), nil
//				}).(pulumi.StringOutput),
//				Tags: pulumi.StringMap{
//					"tag1": pulumi.String("value1"),
//					"tag2": pulumi.String("value3"),
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				object,
//			}))
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
// Using `pulumi import`, import Transcribe MedicalVocabulary using the `vocabulary_name`. For example:
//
// ```sh
//
//	$ pulumi import aws:transcribe/medicalVocabulary:MedicalVocabulary example example-name
//
// ```
type MedicalVocabulary struct {
	pulumi.CustomResourceState

	// ARN of the MedicalVocabulary.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Generated download URI.
	DownloadUri pulumi.StringOutput `pulumi:"downloadUri"`
	// The language code you selected for your medical vocabulary. US English (en-US) is the only language supported with Amazon Transcribe Medical.
	LanguageCode pulumi.StringOutput `pulumi:"languageCode"`
	// A map of tags to assign to the MedicalVocabulary. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The Amazon S3 location (URI) of the text file that contains your custom medical vocabulary.
	VocabularyFileUri pulumi.StringOutput `pulumi:"vocabularyFileUri"`
	// The name of the Medical Vocabulary.
	//
	// The following arguments are optional:
	VocabularyName pulumi.StringOutput `pulumi:"vocabularyName"`
}

// NewMedicalVocabulary registers a new resource with the given unique name, arguments, and options.
func NewMedicalVocabulary(ctx *pulumi.Context,
	name string, args *MedicalVocabularyArgs, opts ...pulumi.ResourceOption) (*MedicalVocabulary, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LanguageCode == nil {
		return nil, errors.New("invalid value for required argument 'LanguageCode'")
	}
	if args.VocabularyFileUri == nil {
		return nil, errors.New("invalid value for required argument 'VocabularyFileUri'")
	}
	if args.VocabularyName == nil {
		return nil, errors.New("invalid value for required argument 'VocabularyName'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MedicalVocabulary
	err := ctx.RegisterResource("aws:transcribe/medicalVocabulary:MedicalVocabulary", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMedicalVocabulary gets an existing MedicalVocabulary resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMedicalVocabulary(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MedicalVocabularyState, opts ...pulumi.ResourceOption) (*MedicalVocabulary, error) {
	var resource MedicalVocabulary
	err := ctx.ReadResource("aws:transcribe/medicalVocabulary:MedicalVocabulary", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MedicalVocabulary resources.
type medicalVocabularyState struct {
	// ARN of the MedicalVocabulary.
	Arn *string `pulumi:"arn"`
	// Generated download URI.
	DownloadUri *string `pulumi:"downloadUri"`
	// The language code you selected for your medical vocabulary. US English (en-US) is the only language supported with Amazon Transcribe Medical.
	LanguageCode *string `pulumi:"languageCode"`
	// A map of tags to assign to the MedicalVocabulary. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The Amazon S3 location (URI) of the text file that contains your custom medical vocabulary.
	VocabularyFileUri *string `pulumi:"vocabularyFileUri"`
	// The name of the Medical Vocabulary.
	//
	// The following arguments are optional:
	VocabularyName *string `pulumi:"vocabularyName"`
}

type MedicalVocabularyState struct {
	// ARN of the MedicalVocabulary.
	Arn pulumi.StringPtrInput
	// Generated download URI.
	DownloadUri pulumi.StringPtrInput
	// The language code you selected for your medical vocabulary. US English (en-US) is the only language supported with Amazon Transcribe Medical.
	LanguageCode pulumi.StringPtrInput
	// A map of tags to assign to the MedicalVocabulary. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// The Amazon S3 location (URI) of the text file that contains your custom medical vocabulary.
	VocabularyFileUri pulumi.StringPtrInput
	// The name of the Medical Vocabulary.
	//
	// The following arguments are optional:
	VocabularyName pulumi.StringPtrInput
}

func (MedicalVocabularyState) ElementType() reflect.Type {
	return reflect.TypeOf((*medicalVocabularyState)(nil)).Elem()
}

type medicalVocabularyArgs struct {
	// The language code you selected for your medical vocabulary. US English (en-US) is the only language supported with Amazon Transcribe Medical.
	LanguageCode string `pulumi:"languageCode"`
	// A map of tags to assign to the MedicalVocabulary. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The Amazon S3 location (URI) of the text file that contains your custom medical vocabulary.
	VocabularyFileUri string `pulumi:"vocabularyFileUri"`
	// The name of the Medical Vocabulary.
	//
	// The following arguments are optional:
	VocabularyName string `pulumi:"vocabularyName"`
}

// The set of arguments for constructing a MedicalVocabulary resource.
type MedicalVocabularyArgs struct {
	// The language code you selected for your medical vocabulary. US English (en-US) is the only language supported with Amazon Transcribe Medical.
	LanguageCode pulumi.StringInput
	// A map of tags to assign to the MedicalVocabulary. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The Amazon S3 location (URI) of the text file that contains your custom medical vocabulary.
	VocabularyFileUri pulumi.StringInput
	// The name of the Medical Vocabulary.
	//
	// The following arguments are optional:
	VocabularyName pulumi.StringInput
}

func (MedicalVocabularyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*medicalVocabularyArgs)(nil)).Elem()
}

type MedicalVocabularyInput interface {
	pulumi.Input

	ToMedicalVocabularyOutput() MedicalVocabularyOutput
	ToMedicalVocabularyOutputWithContext(ctx context.Context) MedicalVocabularyOutput
}

func (*MedicalVocabulary) ElementType() reflect.Type {
	return reflect.TypeOf((**MedicalVocabulary)(nil)).Elem()
}

func (i *MedicalVocabulary) ToMedicalVocabularyOutput() MedicalVocabularyOutput {
	return i.ToMedicalVocabularyOutputWithContext(context.Background())
}

func (i *MedicalVocabulary) ToMedicalVocabularyOutputWithContext(ctx context.Context) MedicalVocabularyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MedicalVocabularyOutput)
}

// MedicalVocabularyArrayInput is an input type that accepts MedicalVocabularyArray and MedicalVocabularyArrayOutput values.
// You can construct a concrete instance of `MedicalVocabularyArrayInput` via:
//
//	MedicalVocabularyArray{ MedicalVocabularyArgs{...} }
type MedicalVocabularyArrayInput interface {
	pulumi.Input

	ToMedicalVocabularyArrayOutput() MedicalVocabularyArrayOutput
	ToMedicalVocabularyArrayOutputWithContext(context.Context) MedicalVocabularyArrayOutput
}

type MedicalVocabularyArray []MedicalVocabularyInput

func (MedicalVocabularyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MedicalVocabulary)(nil)).Elem()
}

func (i MedicalVocabularyArray) ToMedicalVocabularyArrayOutput() MedicalVocabularyArrayOutput {
	return i.ToMedicalVocabularyArrayOutputWithContext(context.Background())
}

func (i MedicalVocabularyArray) ToMedicalVocabularyArrayOutputWithContext(ctx context.Context) MedicalVocabularyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MedicalVocabularyArrayOutput)
}

// MedicalVocabularyMapInput is an input type that accepts MedicalVocabularyMap and MedicalVocabularyMapOutput values.
// You can construct a concrete instance of `MedicalVocabularyMapInput` via:
//
//	MedicalVocabularyMap{ "key": MedicalVocabularyArgs{...} }
type MedicalVocabularyMapInput interface {
	pulumi.Input

	ToMedicalVocabularyMapOutput() MedicalVocabularyMapOutput
	ToMedicalVocabularyMapOutputWithContext(context.Context) MedicalVocabularyMapOutput
}

type MedicalVocabularyMap map[string]MedicalVocabularyInput

func (MedicalVocabularyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MedicalVocabulary)(nil)).Elem()
}

func (i MedicalVocabularyMap) ToMedicalVocabularyMapOutput() MedicalVocabularyMapOutput {
	return i.ToMedicalVocabularyMapOutputWithContext(context.Background())
}

func (i MedicalVocabularyMap) ToMedicalVocabularyMapOutputWithContext(ctx context.Context) MedicalVocabularyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MedicalVocabularyMapOutput)
}

type MedicalVocabularyOutput struct{ *pulumi.OutputState }

func (MedicalVocabularyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MedicalVocabulary)(nil)).Elem()
}

func (o MedicalVocabularyOutput) ToMedicalVocabularyOutput() MedicalVocabularyOutput {
	return o
}

func (o MedicalVocabularyOutput) ToMedicalVocabularyOutputWithContext(ctx context.Context) MedicalVocabularyOutput {
	return o
}

// ARN of the MedicalVocabulary.
func (o MedicalVocabularyOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *MedicalVocabulary) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Generated download URI.
func (o MedicalVocabularyOutput) DownloadUri() pulumi.StringOutput {
	return o.ApplyT(func(v *MedicalVocabulary) pulumi.StringOutput { return v.DownloadUri }).(pulumi.StringOutput)
}

// The language code you selected for your medical vocabulary. US English (en-US) is the only language supported with Amazon Transcribe Medical.
func (o MedicalVocabularyOutput) LanguageCode() pulumi.StringOutput {
	return o.ApplyT(func(v *MedicalVocabulary) pulumi.StringOutput { return v.LanguageCode }).(pulumi.StringOutput)
}

// A map of tags to assign to the MedicalVocabulary. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o MedicalVocabularyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MedicalVocabulary) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Deprecated: Please use `tags` instead.
func (o MedicalVocabularyOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MedicalVocabulary) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The Amazon S3 location (URI) of the text file that contains your custom medical vocabulary.
func (o MedicalVocabularyOutput) VocabularyFileUri() pulumi.StringOutput {
	return o.ApplyT(func(v *MedicalVocabulary) pulumi.StringOutput { return v.VocabularyFileUri }).(pulumi.StringOutput)
}

// The name of the Medical Vocabulary.
//
// The following arguments are optional:
func (o MedicalVocabularyOutput) VocabularyName() pulumi.StringOutput {
	return o.ApplyT(func(v *MedicalVocabulary) pulumi.StringOutput { return v.VocabularyName }).(pulumi.StringOutput)
}

type MedicalVocabularyArrayOutput struct{ *pulumi.OutputState }

func (MedicalVocabularyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MedicalVocabulary)(nil)).Elem()
}

func (o MedicalVocabularyArrayOutput) ToMedicalVocabularyArrayOutput() MedicalVocabularyArrayOutput {
	return o
}

func (o MedicalVocabularyArrayOutput) ToMedicalVocabularyArrayOutputWithContext(ctx context.Context) MedicalVocabularyArrayOutput {
	return o
}

func (o MedicalVocabularyArrayOutput) Index(i pulumi.IntInput) MedicalVocabularyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MedicalVocabulary {
		return vs[0].([]*MedicalVocabulary)[vs[1].(int)]
	}).(MedicalVocabularyOutput)
}

type MedicalVocabularyMapOutput struct{ *pulumi.OutputState }

func (MedicalVocabularyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MedicalVocabulary)(nil)).Elem()
}

func (o MedicalVocabularyMapOutput) ToMedicalVocabularyMapOutput() MedicalVocabularyMapOutput {
	return o
}

func (o MedicalVocabularyMapOutput) ToMedicalVocabularyMapOutputWithContext(ctx context.Context) MedicalVocabularyMapOutput {
	return o
}

func (o MedicalVocabularyMapOutput) MapIndex(k pulumi.StringInput) MedicalVocabularyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MedicalVocabulary {
		return vs[0].(map[string]*MedicalVocabulary)[vs[1].(string)]
	}).(MedicalVocabularyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MedicalVocabularyInput)(nil)).Elem(), &MedicalVocabulary{})
	pulumi.RegisterInputType(reflect.TypeOf((*MedicalVocabularyArrayInput)(nil)).Elem(), MedicalVocabularyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MedicalVocabularyMapInput)(nil)).Elem(), MedicalVocabularyMap{})
	pulumi.RegisterOutputType(MedicalVocabularyOutput{})
	pulumi.RegisterOutputType(MedicalVocabularyArrayOutput{})
	pulumi.RegisterOutputType(MedicalVocabularyMapOutput{})
}
