// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package transcribe

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for managing an AWS Transcribe LanguageModel.
//
// > This resource can take a significant amount of time to provision. See Language Model [FAQ](https://aws.amazon.com/transcribe/faqs/) for more details.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//	"fmt"
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/transcribe"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			examplePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
//				Statements: []iam.GetPolicyDocumentStatement{
//					{
//						Actions: []string{
//							"sts:AssumeRole",
//						},
//						Principals: []iam.GetPolicyDocumentStatementPrincipal{
//							{
//								Type: "Service",
//								Identifiers: []string{
//									"transcribe.amazonaws.com",
//								},
//							},
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
//				AssumeRolePolicy: *pulumi.String(examplePolicyDocument.Json),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"Version": "2012-10-17",
//				"Statement": []map[string]interface{}{
//					map[string]interface{}{
//						"Action": []string{
//							"s3:GetObject",
//							"s3:ListBucket",
//						},
//						"Effect": "Allow",
//						"Resource": []string{
//							"*",
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = iam.NewRolePolicy(ctx, "testPolicy", &iam.RolePolicyArgs{
//				Role:   exampleRole.ID(),
//				Policy: pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", &s3.BucketV2Args{
//				ForceDestroy: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewBucketObjectv2(ctx, "object", &s3.BucketObjectv2Args{
//				Bucket: exampleBucketV2.ID(),
//				Key:    pulumi.String("transcribe/test1.txt"),
//				Source: pulumi.NewFileAsset("test1.txt"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = transcribe.NewLanguageModel(ctx, "exampleLanguageModel", &transcribe.LanguageModelArgs{
//				ModelName:     pulumi.String("example"),
//				BaseModelName: pulumi.String("NarrowBand"),
//				InputDataConfig: &transcribe.LanguageModelInputDataConfigArgs{
//					DataAccessRoleArn: exampleRole.Arn,
//					S3Uri: exampleBucketV2.ID().ApplyT(func(id string) (string, error) {
//						return fmt.Sprintf("s3://%v/transcribe/", id), nil
//					}).(pulumi.StringOutput),
//				},
//				LanguageCode: pulumi.String("en-US"),
//				Tags: pulumi.StringMap{
//					"ENVIRONMENT": pulumi.String("development"),
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
// Transcribe LanguageModel can be imported using the `model_name`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:transcribe/languageModel:LanguageModel example example-name
//
// ```
type LanguageModel struct {
	pulumi.CustomResourceState

	// ARN of the LanguageModel.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Name of reference base model.
	BaseModelName pulumi.StringOutput `pulumi:"baseModelName"`
	// The input data config for the LanguageModel. See Input Data Config for more details.
	InputDataConfig LanguageModelInputDataConfigOutput `pulumi:"inputDataConfig"`
	// The language code you selected for your language model. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
	LanguageCode pulumi.StringOutput `pulumi:"languageCode"`
	// The model name.
	ModelName pulumi.StringOutput `pulumi:"modelName"`
	// A map of tags to assign to the LanguageModel. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    pulumi.StringMapOutput `pulumi:"tags"`
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewLanguageModel registers a new resource with the given unique name, arguments, and options.
func NewLanguageModel(ctx *pulumi.Context,
	name string, args *LanguageModelArgs, opts ...pulumi.ResourceOption) (*LanguageModel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BaseModelName == nil {
		return nil, errors.New("invalid value for required argument 'BaseModelName'")
	}
	if args.InputDataConfig == nil {
		return nil, errors.New("invalid value for required argument 'InputDataConfig'")
	}
	if args.LanguageCode == nil {
		return nil, errors.New("invalid value for required argument 'LanguageCode'")
	}
	if args.ModelName == nil {
		return nil, errors.New("invalid value for required argument 'ModelName'")
	}
	var resource LanguageModel
	err := ctx.RegisterResource("aws:transcribe/languageModel:LanguageModel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLanguageModel gets an existing LanguageModel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLanguageModel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LanguageModelState, opts ...pulumi.ResourceOption) (*LanguageModel, error) {
	var resource LanguageModel
	err := ctx.ReadResource("aws:transcribe/languageModel:LanguageModel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LanguageModel resources.
type languageModelState struct {
	// ARN of the LanguageModel.
	Arn *string `pulumi:"arn"`
	// Name of reference base model.
	BaseModelName *string `pulumi:"baseModelName"`
	// The input data config for the LanguageModel. See Input Data Config for more details.
	InputDataConfig *LanguageModelInputDataConfig `pulumi:"inputDataConfig"`
	// The language code you selected for your language model. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
	LanguageCode *string `pulumi:"languageCode"`
	// The model name.
	ModelName *string `pulumi:"modelName"`
	// A map of tags to assign to the LanguageModel. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    map[string]string `pulumi:"tags"`
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type LanguageModelState struct {
	// ARN of the LanguageModel.
	Arn pulumi.StringPtrInput
	// Name of reference base model.
	BaseModelName pulumi.StringPtrInput
	// The input data config for the LanguageModel. See Input Data Config for more details.
	InputDataConfig LanguageModelInputDataConfigPtrInput
	// The language code you selected for your language model. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
	LanguageCode pulumi.StringPtrInput
	// The model name.
	ModelName pulumi.StringPtrInput
	// A map of tags to assign to the LanguageModel. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags    pulumi.StringMapInput
	TagsAll pulumi.StringMapInput
}

func (LanguageModelState) ElementType() reflect.Type {
	return reflect.TypeOf((*languageModelState)(nil)).Elem()
}

type languageModelArgs struct {
	// Name of reference base model.
	BaseModelName string `pulumi:"baseModelName"`
	// The input data config for the LanguageModel. See Input Data Config for more details.
	InputDataConfig LanguageModelInputDataConfig `pulumi:"inputDataConfig"`
	// The language code you selected for your language model. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
	LanguageCode string `pulumi:"languageCode"`
	// The model name.
	ModelName string `pulumi:"modelName"`
	// A map of tags to assign to the LanguageModel. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a LanguageModel resource.
type LanguageModelArgs struct {
	// Name of reference base model.
	BaseModelName pulumi.StringInput
	// The input data config for the LanguageModel. See Input Data Config for more details.
	InputDataConfig LanguageModelInputDataConfigInput
	// The language code you selected for your language model. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
	LanguageCode pulumi.StringInput
	// The model name.
	ModelName pulumi.StringInput
	// A map of tags to assign to the LanguageModel. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (LanguageModelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*languageModelArgs)(nil)).Elem()
}

type LanguageModelInput interface {
	pulumi.Input

	ToLanguageModelOutput() LanguageModelOutput
	ToLanguageModelOutputWithContext(ctx context.Context) LanguageModelOutput
}

func (*LanguageModel) ElementType() reflect.Type {
	return reflect.TypeOf((**LanguageModel)(nil)).Elem()
}

func (i *LanguageModel) ToLanguageModelOutput() LanguageModelOutput {
	return i.ToLanguageModelOutputWithContext(context.Background())
}

func (i *LanguageModel) ToLanguageModelOutputWithContext(ctx context.Context) LanguageModelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LanguageModelOutput)
}

// LanguageModelArrayInput is an input type that accepts LanguageModelArray and LanguageModelArrayOutput values.
// You can construct a concrete instance of `LanguageModelArrayInput` via:
//
//	LanguageModelArray{ LanguageModelArgs{...} }
type LanguageModelArrayInput interface {
	pulumi.Input

	ToLanguageModelArrayOutput() LanguageModelArrayOutput
	ToLanguageModelArrayOutputWithContext(context.Context) LanguageModelArrayOutput
}

type LanguageModelArray []LanguageModelInput

func (LanguageModelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LanguageModel)(nil)).Elem()
}

func (i LanguageModelArray) ToLanguageModelArrayOutput() LanguageModelArrayOutput {
	return i.ToLanguageModelArrayOutputWithContext(context.Background())
}

func (i LanguageModelArray) ToLanguageModelArrayOutputWithContext(ctx context.Context) LanguageModelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LanguageModelArrayOutput)
}

// LanguageModelMapInput is an input type that accepts LanguageModelMap and LanguageModelMapOutput values.
// You can construct a concrete instance of `LanguageModelMapInput` via:
//
//	LanguageModelMap{ "key": LanguageModelArgs{...} }
type LanguageModelMapInput interface {
	pulumi.Input

	ToLanguageModelMapOutput() LanguageModelMapOutput
	ToLanguageModelMapOutputWithContext(context.Context) LanguageModelMapOutput
}

type LanguageModelMap map[string]LanguageModelInput

func (LanguageModelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LanguageModel)(nil)).Elem()
}

func (i LanguageModelMap) ToLanguageModelMapOutput() LanguageModelMapOutput {
	return i.ToLanguageModelMapOutputWithContext(context.Background())
}

func (i LanguageModelMap) ToLanguageModelMapOutputWithContext(ctx context.Context) LanguageModelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LanguageModelMapOutput)
}

type LanguageModelOutput struct{ *pulumi.OutputState }

func (LanguageModelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LanguageModel)(nil)).Elem()
}

func (o LanguageModelOutput) ToLanguageModelOutput() LanguageModelOutput {
	return o
}

func (o LanguageModelOutput) ToLanguageModelOutputWithContext(ctx context.Context) LanguageModelOutput {
	return o
}

// ARN of the LanguageModel.
func (o LanguageModelOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *LanguageModel) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Name of reference base model.
func (o LanguageModelOutput) BaseModelName() pulumi.StringOutput {
	return o.ApplyT(func(v *LanguageModel) pulumi.StringOutput { return v.BaseModelName }).(pulumi.StringOutput)
}

// The input data config for the LanguageModel. See Input Data Config for more details.
func (o LanguageModelOutput) InputDataConfig() LanguageModelInputDataConfigOutput {
	return o.ApplyT(func(v *LanguageModel) LanguageModelInputDataConfigOutput { return v.InputDataConfig }).(LanguageModelInputDataConfigOutput)
}

// The language code you selected for your language model. Refer to the [supported languages](https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) page for accepted codes.
func (o LanguageModelOutput) LanguageCode() pulumi.StringOutput {
	return o.ApplyT(func(v *LanguageModel) pulumi.StringOutput { return v.LanguageCode }).(pulumi.StringOutput)
}

// The model name.
func (o LanguageModelOutput) ModelName() pulumi.StringOutput {
	return o.ApplyT(func(v *LanguageModel) pulumi.StringOutput { return v.ModelName }).(pulumi.StringOutput)
}

// A map of tags to assign to the LanguageModel. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o LanguageModelOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LanguageModel) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LanguageModelOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LanguageModel) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type LanguageModelArrayOutput struct{ *pulumi.OutputState }

func (LanguageModelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LanguageModel)(nil)).Elem()
}

func (o LanguageModelArrayOutput) ToLanguageModelArrayOutput() LanguageModelArrayOutput {
	return o
}

func (o LanguageModelArrayOutput) ToLanguageModelArrayOutputWithContext(ctx context.Context) LanguageModelArrayOutput {
	return o
}

func (o LanguageModelArrayOutput) Index(i pulumi.IntInput) LanguageModelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LanguageModel {
		return vs[0].([]*LanguageModel)[vs[1].(int)]
	}).(LanguageModelOutput)
}

type LanguageModelMapOutput struct{ *pulumi.OutputState }

func (LanguageModelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LanguageModel)(nil)).Elem()
}

func (o LanguageModelMapOutput) ToLanguageModelMapOutput() LanguageModelMapOutput {
	return o
}

func (o LanguageModelMapOutput) ToLanguageModelMapOutputWithContext(ctx context.Context) LanguageModelMapOutput {
	return o
}

func (o LanguageModelMapOutput) MapIndex(k pulumi.StringInput) LanguageModelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LanguageModel {
		return vs[0].(map[string]*LanguageModel)[vs[1].(string)]
	}).(LanguageModelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LanguageModelInput)(nil)).Elem(), &LanguageModel{})
	pulumi.RegisterInputType(reflect.TypeOf((*LanguageModelArrayInput)(nil)).Elem(), LanguageModelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LanguageModelMapInput)(nil)).Elem(), LanguageModelMap{})
	pulumi.RegisterOutputType(LanguageModelOutput{})
	pulumi.RegisterOutputType(LanguageModelArrayOutput{})
	pulumi.RegisterOutputType(LanguageModelMapOutput{})
}
