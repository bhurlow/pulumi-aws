// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package comprehend

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource for managing an AWS Comprehend Document Classifier.
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
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/comprehend"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			documents, err := s3.NewBucketObjectv2(ctx, "documents", nil)
//			if err != nil {
//				return err
//			}
//			_, err = comprehend.NewDocumentClassifier(ctx, "example", &comprehend.DocumentClassifierArgs{
//				DataAccessRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
//				LanguageCode:      pulumi.String("en"),
//				InputDataConfig: &comprehend.DocumentClassifierInputDataConfigArgs{
//					S3Uri: documents.ID().ApplyT(func(id string) (string, error) {
//						return fmt.Sprintf("s3://%v/%v", aws_s3_bucket.Test.Bucket, id), nil
//					}).(pulumi.StringOutput),
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				aws_iam_role_policy.Example,
//			}))
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketObjectv2(ctx, "entities", nil)
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
// Using `pulumi import`, import Comprehend Document Classifier using the ARN. For example:
//
// ```sh
//
//	$ pulumi import aws:comprehend/documentClassifier:DocumentClassifier example arn:aws:comprehend:us-west-2:123456789012:document_classifier/example
//
// ```
type DocumentClassifier struct {
	pulumi.CustomResourceState

	// ARN of the Document Classifier version.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ARN for an IAM Role which allows Comprehend to read the training and testing data.
	DataAccessRoleArn pulumi.StringOutput `pulumi:"dataAccessRoleArn"`
	// Configuration for the training and testing data.
	// See the `inputDataConfig` Configuration Block section below.
	InputDataConfig DocumentClassifierInputDataConfigOutput `pulumi:"inputDataConfig"`
	// Two-letter language code for the language.
	// One of `en`, `es`, `fr`, `it`, `de`, or `pt`.
	LanguageCode pulumi.StringOutput `pulumi:"languageCode"`
	// The document classification mode.
	// One of `MULTI_CLASS` or `MULTI_LABEL`.
	// `MULTI_CLASS` is also known as "Single Label" in the AWS Console.
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// KMS Key used to encrypt trained Document Classifiers.
	// Can be a KMS Key ID or a KMS Key ARN.
	ModelKmsKeyId pulumi.StringPtrOutput `pulumi:"modelKmsKeyId"`
	// Name for the Document Classifier.
	// Has a maximum length of 63 characters.
	// Can contain upper- and lower-case letters, numbers, and hypen (`-`).
	//
	// The following arguments are optional:
	Name pulumi.StringOutput `pulumi:"name"`
	// Configuration for the output results of training.
	// See the `outputDataConfig` Configuration Block section below.
	OutputDataConfig DocumentClassifierOutputDataConfigOutput `pulumi:"outputDataConfig"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` Configuration Block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Name for the version of the Document Classifier.
	// Each version must have a unique name within the Document Classifier.
	// If omitted, the provider will assign a random, unique version name.
	// If explicitly set to `""`, no version name will be set.
	// Has a maximum length of 63 characters.
	// Can contain upper- and lower-case letters, numbers, and hypen (`-`).
	// Conflicts with `versionNamePrefix`.
	VersionName pulumi.StringOutput `pulumi:"versionName"`
	// Creates a unique version name beginning with the specified prefix.
	// Has a maximum length of 37 characters.
	// Can contain upper- and lower-case letters, numbers, and hypen (`-`).
	// Conflicts with `versionName`.
	VersionNamePrefix pulumi.StringOutput `pulumi:"versionNamePrefix"`
	// KMS Key used to encrypt storage volumes during job processing.
	// Can be a KMS Key ID or a KMS Key ARN.
	VolumeKmsKeyId pulumi.StringPtrOutput `pulumi:"volumeKmsKeyId"`
	// Configuration parameters for VPC to contain Document Classifier resources.
	// See the `vpcConfig` Configuration Block section below.
	VpcConfig DocumentClassifierVpcConfigPtrOutput `pulumi:"vpcConfig"`
}

// NewDocumentClassifier registers a new resource with the given unique name, arguments, and options.
func NewDocumentClassifier(ctx *pulumi.Context,
	name string, args *DocumentClassifierArgs, opts ...pulumi.ResourceOption) (*DocumentClassifier, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataAccessRoleArn == nil {
		return nil, errors.New("invalid value for required argument 'DataAccessRoleArn'")
	}
	if args.InputDataConfig == nil {
		return nil, errors.New("invalid value for required argument 'InputDataConfig'")
	}
	if args.LanguageCode == nil {
		return nil, errors.New("invalid value for required argument 'LanguageCode'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DocumentClassifier
	err := ctx.RegisterResource("aws:comprehend/documentClassifier:DocumentClassifier", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDocumentClassifier gets an existing DocumentClassifier resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDocumentClassifier(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DocumentClassifierState, opts ...pulumi.ResourceOption) (*DocumentClassifier, error) {
	var resource DocumentClassifier
	err := ctx.ReadResource("aws:comprehend/documentClassifier:DocumentClassifier", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DocumentClassifier resources.
type documentClassifierState struct {
	// ARN of the Document Classifier version.
	Arn *string `pulumi:"arn"`
	// The ARN for an IAM Role which allows Comprehend to read the training and testing data.
	DataAccessRoleArn *string `pulumi:"dataAccessRoleArn"`
	// Configuration for the training and testing data.
	// See the `inputDataConfig` Configuration Block section below.
	InputDataConfig *DocumentClassifierInputDataConfig `pulumi:"inputDataConfig"`
	// Two-letter language code for the language.
	// One of `en`, `es`, `fr`, `it`, `de`, or `pt`.
	LanguageCode *string `pulumi:"languageCode"`
	// The document classification mode.
	// One of `MULTI_CLASS` or `MULTI_LABEL`.
	// `MULTI_CLASS` is also known as "Single Label" in the AWS Console.
	Mode *string `pulumi:"mode"`
	// KMS Key used to encrypt trained Document Classifiers.
	// Can be a KMS Key ID or a KMS Key ARN.
	ModelKmsKeyId *string `pulumi:"modelKmsKeyId"`
	// Name for the Document Classifier.
	// Has a maximum length of 63 characters.
	// Can contain upper- and lower-case letters, numbers, and hypen (`-`).
	//
	// The following arguments are optional:
	Name *string `pulumi:"name"`
	// Configuration for the output results of training.
	// See the `outputDataConfig` Configuration Block section below.
	OutputDataConfig *DocumentClassifierOutputDataConfig `pulumi:"outputDataConfig"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` Configuration Block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Name for the version of the Document Classifier.
	// Each version must have a unique name within the Document Classifier.
	// If omitted, the provider will assign a random, unique version name.
	// If explicitly set to `""`, no version name will be set.
	// Has a maximum length of 63 characters.
	// Can contain upper- and lower-case letters, numbers, and hypen (`-`).
	// Conflicts with `versionNamePrefix`.
	VersionName *string `pulumi:"versionName"`
	// Creates a unique version name beginning with the specified prefix.
	// Has a maximum length of 37 characters.
	// Can contain upper- and lower-case letters, numbers, and hypen (`-`).
	// Conflicts with `versionName`.
	VersionNamePrefix *string `pulumi:"versionNamePrefix"`
	// KMS Key used to encrypt storage volumes during job processing.
	// Can be a KMS Key ID or a KMS Key ARN.
	VolumeKmsKeyId *string `pulumi:"volumeKmsKeyId"`
	// Configuration parameters for VPC to contain Document Classifier resources.
	// See the `vpcConfig` Configuration Block section below.
	VpcConfig *DocumentClassifierVpcConfig `pulumi:"vpcConfig"`
}

type DocumentClassifierState struct {
	// ARN of the Document Classifier version.
	Arn pulumi.StringPtrInput
	// The ARN for an IAM Role which allows Comprehend to read the training and testing data.
	DataAccessRoleArn pulumi.StringPtrInput
	// Configuration for the training and testing data.
	// See the `inputDataConfig` Configuration Block section below.
	InputDataConfig DocumentClassifierInputDataConfigPtrInput
	// Two-letter language code for the language.
	// One of `en`, `es`, `fr`, `it`, `de`, or `pt`.
	LanguageCode pulumi.StringPtrInput
	// The document classification mode.
	// One of `MULTI_CLASS` or `MULTI_LABEL`.
	// `MULTI_CLASS` is also known as "Single Label" in the AWS Console.
	Mode pulumi.StringPtrInput
	// KMS Key used to encrypt trained Document Classifiers.
	// Can be a KMS Key ID or a KMS Key ARN.
	ModelKmsKeyId pulumi.StringPtrInput
	// Name for the Document Classifier.
	// Has a maximum length of 63 characters.
	// Can contain upper- and lower-case letters, numbers, and hypen (`-`).
	//
	// The following arguments are optional:
	Name pulumi.StringPtrInput
	// Configuration for the output results of training.
	// See the `outputDataConfig` Configuration Block section below.
	OutputDataConfig DocumentClassifierOutputDataConfigPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` Configuration Block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// Name for the version of the Document Classifier.
	// Each version must have a unique name within the Document Classifier.
	// If omitted, the provider will assign a random, unique version name.
	// If explicitly set to `""`, no version name will be set.
	// Has a maximum length of 63 characters.
	// Can contain upper- and lower-case letters, numbers, and hypen (`-`).
	// Conflicts with `versionNamePrefix`.
	VersionName pulumi.StringPtrInput
	// Creates a unique version name beginning with the specified prefix.
	// Has a maximum length of 37 characters.
	// Can contain upper- and lower-case letters, numbers, and hypen (`-`).
	// Conflicts with `versionName`.
	VersionNamePrefix pulumi.StringPtrInput
	// KMS Key used to encrypt storage volumes during job processing.
	// Can be a KMS Key ID or a KMS Key ARN.
	VolumeKmsKeyId pulumi.StringPtrInput
	// Configuration parameters for VPC to contain Document Classifier resources.
	// See the `vpcConfig` Configuration Block section below.
	VpcConfig DocumentClassifierVpcConfigPtrInput
}

func (DocumentClassifierState) ElementType() reflect.Type {
	return reflect.TypeOf((*documentClassifierState)(nil)).Elem()
}

type documentClassifierArgs struct {
	// The ARN for an IAM Role which allows Comprehend to read the training and testing data.
	DataAccessRoleArn string `pulumi:"dataAccessRoleArn"`
	// Configuration for the training and testing data.
	// See the `inputDataConfig` Configuration Block section below.
	InputDataConfig DocumentClassifierInputDataConfig `pulumi:"inputDataConfig"`
	// Two-letter language code for the language.
	// One of `en`, `es`, `fr`, `it`, `de`, or `pt`.
	LanguageCode string `pulumi:"languageCode"`
	// The document classification mode.
	// One of `MULTI_CLASS` or `MULTI_LABEL`.
	// `MULTI_CLASS` is also known as "Single Label" in the AWS Console.
	Mode *string `pulumi:"mode"`
	// KMS Key used to encrypt trained Document Classifiers.
	// Can be a KMS Key ID or a KMS Key ARN.
	ModelKmsKeyId *string `pulumi:"modelKmsKeyId"`
	// Name for the Document Classifier.
	// Has a maximum length of 63 characters.
	// Can contain upper- and lower-case letters, numbers, and hypen (`-`).
	//
	// The following arguments are optional:
	Name *string `pulumi:"name"`
	// Configuration for the output results of training.
	// See the `outputDataConfig` Configuration Block section below.
	OutputDataConfig *DocumentClassifierOutputDataConfig `pulumi:"outputDataConfig"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` Configuration Block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Name for the version of the Document Classifier.
	// Each version must have a unique name within the Document Classifier.
	// If omitted, the provider will assign a random, unique version name.
	// If explicitly set to `""`, no version name will be set.
	// Has a maximum length of 63 characters.
	// Can contain upper- and lower-case letters, numbers, and hypen (`-`).
	// Conflicts with `versionNamePrefix`.
	VersionName *string `pulumi:"versionName"`
	// Creates a unique version name beginning with the specified prefix.
	// Has a maximum length of 37 characters.
	// Can contain upper- and lower-case letters, numbers, and hypen (`-`).
	// Conflicts with `versionName`.
	VersionNamePrefix *string `pulumi:"versionNamePrefix"`
	// KMS Key used to encrypt storage volumes during job processing.
	// Can be a KMS Key ID or a KMS Key ARN.
	VolumeKmsKeyId *string `pulumi:"volumeKmsKeyId"`
	// Configuration parameters for VPC to contain Document Classifier resources.
	// See the `vpcConfig` Configuration Block section below.
	VpcConfig *DocumentClassifierVpcConfig `pulumi:"vpcConfig"`
}

// The set of arguments for constructing a DocumentClassifier resource.
type DocumentClassifierArgs struct {
	// The ARN for an IAM Role which allows Comprehend to read the training and testing data.
	DataAccessRoleArn pulumi.StringInput
	// Configuration for the training and testing data.
	// See the `inputDataConfig` Configuration Block section below.
	InputDataConfig DocumentClassifierInputDataConfigInput
	// Two-letter language code for the language.
	// One of `en`, `es`, `fr`, `it`, `de`, or `pt`.
	LanguageCode pulumi.StringInput
	// The document classification mode.
	// One of `MULTI_CLASS` or `MULTI_LABEL`.
	// `MULTI_CLASS` is also known as "Single Label" in the AWS Console.
	Mode pulumi.StringPtrInput
	// KMS Key used to encrypt trained Document Classifiers.
	// Can be a KMS Key ID or a KMS Key ARN.
	ModelKmsKeyId pulumi.StringPtrInput
	// Name for the Document Classifier.
	// Has a maximum length of 63 characters.
	// Can contain upper- and lower-case letters, numbers, and hypen (`-`).
	//
	// The following arguments are optional:
	Name pulumi.StringPtrInput
	// Configuration for the output results of training.
	// See the `outputDataConfig` Configuration Block section below.
	OutputDataConfig DocumentClassifierOutputDataConfigPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` Configuration Block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Name for the version of the Document Classifier.
	// Each version must have a unique name within the Document Classifier.
	// If omitted, the provider will assign a random, unique version name.
	// If explicitly set to `""`, no version name will be set.
	// Has a maximum length of 63 characters.
	// Can contain upper- and lower-case letters, numbers, and hypen (`-`).
	// Conflicts with `versionNamePrefix`.
	VersionName pulumi.StringPtrInput
	// Creates a unique version name beginning with the specified prefix.
	// Has a maximum length of 37 characters.
	// Can contain upper- and lower-case letters, numbers, and hypen (`-`).
	// Conflicts with `versionName`.
	VersionNamePrefix pulumi.StringPtrInput
	// KMS Key used to encrypt storage volumes during job processing.
	// Can be a KMS Key ID or a KMS Key ARN.
	VolumeKmsKeyId pulumi.StringPtrInput
	// Configuration parameters for VPC to contain Document Classifier resources.
	// See the `vpcConfig` Configuration Block section below.
	VpcConfig DocumentClassifierVpcConfigPtrInput
}

func (DocumentClassifierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*documentClassifierArgs)(nil)).Elem()
}

type DocumentClassifierInput interface {
	pulumi.Input

	ToDocumentClassifierOutput() DocumentClassifierOutput
	ToDocumentClassifierOutputWithContext(ctx context.Context) DocumentClassifierOutput
}

func (*DocumentClassifier) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentClassifier)(nil)).Elem()
}

func (i *DocumentClassifier) ToDocumentClassifierOutput() DocumentClassifierOutput {
	return i.ToDocumentClassifierOutputWithContext(context.Background())
}

func (i *DocumentClassifier) ToDocumentClassifierOutputWithContext(ctx context.Context) DocumentClassifierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentClassifierOutput)
}

func (i *DocumentClassifier) ToOutput(ctx context.Context) pulumix.Output[*DocumentClassifier] {
	return pulumix.Output[*DocumentClassifier]{
		OutputState: i.ToDocumentClassifierOutputWithContext(ctx).OutputState,
	}
}

// DocumentClassifierArrayInput is an input type that accepts DocumentClassifierArray and DocumentClassifierArrayOutput values.
// You can construct a concrete instance of `DocumentClassifierArrayInput` via:
//
//	DocumentClassifierArray{ DocumentClassifierArgs{...} }
type DocumentClassifierArrayInput interface {
	pulumi.Input

	ToDocumentClassifierArrayOutput() DocumentClassifierArrayOutput
	ToDocumentClassifierArrayOutputWithContext(context.Context) DocumentClassifierArrayOutput
}

type DocumentClassifierArray []DocumentClassifierInput

func (DocumentClassifierArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentClassifier)(nil)).Elem()
}

func (i DocumentClassifierArray) ToDocumentClassifierArrayOutput() DocumentClassifierArrayOutput {
	return i.ToDocumentClassifierArrayOutputWithContext(context.Background())
}

func (i DocumentClassifierArray) ToDocumentClassifierArrayOutputWithContext(ctx context.Context) DocumentClassifierArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentClassifierArrayOutput)
}

func (i DocumentClassifierArray) ToOutput(ctx context.Context) pulumix.Output[[]*DocumentClassifier] {
	return pulumix.Output[[]*DocumentClassifier]{
		OutputState: i.ToDocumentClassifierArrayOutputWithContext(ctx).OutputState,
	}
}

// DocumentClassifierMapInput is an input type that accepts DocumentClassifierMap and DocumentClassifierMapOutput values.
// You can construct a concrete instance of `DocumentClassifierMapInput` via:
//
//	DocumentClassifierMap{ "key": DocumentClassifierArgs{...} }
type DocumentClassifierMapInput interface {
	pulumi.Input

	ToDocumentClassifierMapOutput() DocumentClassifierMapOutput
	ToDocumentClassifierMapOutputWithContext(context.Context) DocumentClassifierMapOutput
}

type DocumentClassifierMap map[string]DocumentClassifierInput

func (DocumentClassifierMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentClassifier)(nil)).Elem()
}

func (i DocumentClassifierMap) ToDocumentClassifierMapOutput() DocumentClassifierMapOutput {
	return i.ToDocumentClassifierMapOutputWithContext(context.Background())
}

func (i DocumentClassifierMap) ToDocumentClassifierMapOutputWithContext(ctx context.Context) DocumentClassifierMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentClassifierMapOutput)
}

func (i DocumentClassifierMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DocumentClassifier] {
	return pulumix.Output[map[string]*DocumentClassifier]{
		OutputState: i.ToDocumentClassifierMapOutputWithContext(ctx).OutputState,
	}
}

type DocumentClassifierOutput struct{ *pulumi.OutputState }

func (DocumentClassifierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentClassifier)(nil)).Elem()
}

func (o DocumentClassifierOutput) ToDocumentClassifierOutput() DocumentClassifierOutput {
	return o
}

func (o DocumentClassifierOutput) ToDocumentClassifierOutputWithContext(ctx context.Context) DocumentClassifierOutput {
	return o
}

func (o DocumentClassifierOutput) ToOutput(ctx context.Context) pulumix.Output[*DocumentClassifier] {
	return pulumix.Output[*DocumentClassifier]{
		OutputState: o.OutputState,
	}
}

// ARN of the Document Classifier version.
func (o DocumentClassifierOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentClassifier) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The ARN for an IAM Role which allows Comprehend to read the training and testing data.
func (o DocumentClassifierOutput) DataAccessRoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentClassifier) pulumi.StringOutput { return v.DataAccessRoleArn }).(pulumi.StringOutput)
}

// Configuration for the training and testing data.
// See the `inputDataConfig` Configuration Block section below.
func (o DocumentClassifierOutput) InputDataConfig() DocumentClassifierInputDataConfigOutput {
	return o.ApplyT(func(v *DocumentClassifier) DocumentClassifierInputDataConfigOutput { return v.InputDataConfig }).(DocumentClassifierInputDataConfigOutput)
}

// Two-letter language code for the language.
// One of `en`, `es`, `fr`, `it`, `de`, or `pt`.
func (o DocumentClassifierOutput) LanguageCode() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentClassifier) pulumi.StringOutput { return v.LanguageCode }).(pulumi.StringOutput)
}

// The document classification mode.
// One of `MULTI_CLASS` or `MULTI_LABEL`.
// `MULTI_CLASS` is also known as "Single Label" in the AWS Console.
func (o DocumentClassifierOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DocumentClassifier) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

// KMS Key used to encrypt trained Document Classifiers.
// Can be a KMS Key ID or a KMS Key ARN.
func (o DocumentClassifierOutput) ModelKmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DocumentClassifier) pulumi.StringPtrOutput { return v.ModelKmsKeyId }).(pulumi.StringPtrOutput)
}

// Name for the Document Classifier.
// Has a maximum length of 63 characters.
// Can contain upper- and lower-case letters, numbers, and hypen (`-`).
//
// The following arguments are optional:
func (o DocumentClassifierOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentClassifier) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Configuration for the output results of training.
// See the `outputDataConfig` Configuration Block section below.
func (o DocumentClassifierOutput) OutputDataConfig() DocumentClassifierOutputDataConfigOutput {
	return o.ApplyT(func(v *DocumentClassifier) DocumentClassifierOutputDataConfigOutput { return v.OutputDataConfig }).(DocumentClassifierOutputDataConfigOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` Configuration Block present, tags with matching keys will overwrite those defined at the provider-level.
func (o DocumentClassifierOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DocumentClassifier) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o DocumentClassifierOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DocumentClassifier) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Name for the version of the Document Classifier.
// Each version must have a unique name within the Document Classifier.
// If omitted, the provider will assign a random, unique version name.
// If explicitly set to `""`, no version name will be set.
// Has a maximum length of 63 characters.
// Can contain upper- and lower-case letters, numbers, and hypen (`-`).
// Conflicts with `versionNamePrefix`.
func (o DocumentClassifierOutput) VersionName() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentClassifier) pulumi.StringOutput { return v.VersionName }).(pulumi.StringOutput)
}

// Creates a unique version name beginning with the specified prefix.
// Has a maximum length of 37 characters.
// Can contain upper- and lower-case letters, numbers, and hypen (`-`).
// Conflicts with `versionName`.
func (o DocumentClassifierOutput) VersionNamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentClassifier) pulumi.StringOutput { return v.VersionNamePrefix }).(pulumi.StringOutput)
}

// KMS Key used to encrypt storage volumes during job processing.
// Can be a KMS Key ID or a KMS Key ARN.
func (o DocumentClassifierOutput) VolumeKmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DocumentClassifier) pulumi.StringPtrOutput { return v.VolumeKmsKeyId }).(pulumi.StringPtrOutput)
}

// Configuration parameters for VPC to contain Document Classifier resources.
// See the `vpcConfig` Configuration Block section below.
func (o DocumentClassifierOutput) VpcConfig() DocumentClassifierVpcConfigPtrOutput {
	return o.ApplyT(func(v *DocumentClassifier) DocumentClassifierVpcConfigPtrOutput { return v.VpcConfig }).(DocumentClassifierVpcConfigPtrOutput)
}

type DocumentClassifierArrayOutput struct{ *pulumi.OutputState }

func (DocumentClassifierArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentClassifier)(nil)).Elem()
}

func (o DocumentClassifierArrayOutput) ToDocumentClassifierArrayOutput() DocumentClassifierArrayOutput {
	return o
}

func (o DocumentClassifierArrayOutput) ToDocumentClassifierArrayOutputWithContext(ctx context.Context) DocumentClassifierArrayOutput {
	return o
}

func (o DocumentClassifierArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DocumentClassifier] {
	return pulumix.Output[[]*DocumentClassifier]{
		OutputState: o.OutputState,
	}
}

func (o DocumentClassifierArrayOutput) Index(i pulumi.IntInput) DocumentClassifierOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DocumentClassifier {
		return vs[0].([]*DocumentClassifier)[vs[1].(int)]
	}).(DocumentClassifierOutput)
}

type DocumentClassifierMapOutput struct{ *pulumi.OutputState }

func (DocumentClassifierMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentClassifier)(nil)).Elem()
}

func (o DocumentClassifierMapOutput) ToDocumentClassifierMapOutput() DocumentClassifierMapOutput {
	return o
}

func (o DocumentClassifierMapOutput) ToDocumentClassifierMapOutputWithContext(ctx context.Context) DocumentClassifierMapOutput {
	return o
}

func (o DocumentClassifierMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DocumentClassifier] {
	return pulumix.Output[map[string]*DocumentClassifier]{
		OutputState: o.OutputState,
	}
}

func (o DocumentClassifierMapOutput) MapIndex(k pulumi.StringInput) DocumentClassifierOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DocumentClassifier {
		return vs[0].(map[string]*DocumentClassifier)[vs[1].(string)]
	}).(DocumentClassifierOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentClassifierInput)(nil)).Elem(), &DocumentClassifier{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentClassifierArrayInput)(nil)).Elem(), DocumentClassifierArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentClassifierMapInput)(nil)).Elem(), DocumentClassifierMap{})
	pulumi.RegisterOutputType(DocumentClassifierOutput{})
	pulumi.RegisterOutputType(DocumentClassifierArrayOutput{})
	pulumi.RegisterOutputType(DocumentClassifierMapOutput{})
}
