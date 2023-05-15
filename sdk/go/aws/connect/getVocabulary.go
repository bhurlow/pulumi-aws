// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connect

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about a specific Amazon Connect Vocabulary.
//
// ## Example Usage
//
// By `name`
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/connect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := connect.LookupVocabulary(ctx, &connect.LookupVocabularyArgs{
//				InstanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
//				Name:       pulumi.StringRef("Example"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// By `vocabularyId`
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/connect"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := connect.LookupVocabulary(ctx, &connect.LookupVocabularyArgs{
//				InstanceId:   "aaaaaaaa-bbbb-cccc-dddd-111111111111",
//				VocabularyId: pulumi.StringRef("cccccccc-bbbb-cccc-dddd-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupVocabulary(ctx *pulumi.Context, args *LookupVocabularyArgs, opts ...pulumi.InvokeOption) (*LookupVocabularyResult, error) {
	var rv LookupVocabularyResult
	err := ctx.Invoke("aws:connect/getVocabulary:getVocabulary", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVocabulary.
type LookupVocabularyArgs struct {
	// Reference to the hosting Amazon Connect Instance
	InstanceId string `pulumi:"instanceId"`
	// Returns information on a specific Vocabulary by name
	Name *string `pulumi:"name"`
	// A map of tags to assign to the Vocabulary.
	Tags map[string]string `pulumi:"tags"`
	// Returns information on a specific Vocabulary by Vocabulary id
	VocabularyId *string `pulumi:"vocabularyId"`
}

// A collection of values returned by getVocabulary.
type LookupVocabularyResult struct {
	// The Amazon Resource Name (ARN) of the Vocabulary.
	Arn string `pulumi:"arn"`
	// The content of the custom vocabulary in plain-text format with a table of values. Each row in the table represents a word or a phrase, described with Phrase, IPA, SoundsLike, and DisplayAs fields. Separate the fields with TAB characters. For more information, see [Create a custom vocabulary using a table](https://docs.aws.amazon.com/transcribe/latest/dg/custom-vocabulary.html#create-vocabulary-table).
	Content string `pulumi:"content"`
	// The reason why the custom vocabulary was not created.
	FailureReason string `pulumi:"failureReason"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// The language code of the vocabulary entries. For a list of languages and their corresponding language codes, see [What is Amazon Transcribe?](https://docs.aws.amazon.com/transcribe/latest/dg/transcribe-whatis.html). Valid Values are `ar-AE`, `de-CH`, `de-DE`, `en-AB`, `en-AU`, `en-GB`, `en-IE`, `en-IN`, `en-US`, `en-WL`, `es-ES`, `es-US`, `fr-CA`, `fr-FR`, `hi-IN`, `it-IT`, `ja-JP`, `ko-KR`, `pt-BR`, `pt-PT`, `zh-CN`.
	LanguageCode string `pulumi:"languageCode"`
	// The timestamp when the custom vocabulary was last modified.
	LastModifiedTime string `pulumi:"lastModifiedTime"`
	Name             string `pulumi:"name"`
	// The current state of the custom vocabulary. Valid values are `CREATION_IN_PROGRESS`, `ACTIVE`, `CREATION_FAILED`, `DELETE_IN_PROGRESS`.
	State string `pulumi:"state"`
	// A map of tags to assign to the Vocabulary.
	Tags map[string]string `pulumi:"tags"`
	// The identifier of the custom vocabulary.
	VocabularyId string `pulumi:"vocabularyId"`
}

func LookupVocabularyOutput(ctx *pulumi.Context, args LookupVocabularyOutputArgs, opts ...pulumi.InvokeOption) LookupVocabularyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVocabularyResult, error) {
			args := v.(LookupVocabularyArgs)
			r, err := LookupVocabulary(ctx, &args, opts...)
			var s LookupVocabularyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVocabularyResultOutput)
}

// A collection of arguments for invoking getVocabulary.
type LookupVocabularyOutputArgs struct {
	// Reference to the hosting Amazon Connect Instance
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Returns information on a specific Vocabulary by name
	Name pulumi.StringPtrInput `pulumi:"name"`
	// A map of tags to assign to the Vocabulary.
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// Returns information on a specific Vocabulary by Vocabulary id
	VocabularyId pulumi.StringPtrInput `pulumi:"vocabularyId"`
}

func (LookupVocabularyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVocabularyArgs)(nil)).Elem()
}

// A collection of values returned by getVocabulary.
type LookupVocabularyResultOutput struct{ *pulumi.OutputState }

func (LookupVocabularyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVocabularyResult)(nil)).Elem()
}

func (o LookupVocabularyResultOutput) ToLookupVocabularyResultOutput() LookupVocabularyResultOutput {
	return o
}

func (o LookupVocabularyResultOutput) ToLookupVocabularyResultOutputWithContext(ctx context.Context) LookupVocabularyResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the Vocabulary.
func (o LookupVocabularyResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVocabularyResult) string { return v.Arn }).(pulumi.StringOutput)
}

// The content of the custom vocabulary in plain-text format with a table of values. Each row in the table represents a word or a phrase, described with Phrase, IPA, SoundsLike, and DisplayAs fields. Separate the fields with TAB characters. For more information, see [Create a custom vocabulary using a table](https://docs.aws.amazon.com/transcribe/latest/dg/custom-vocabulary.html#create-vocabulary-table).
func (o LookupVocabularyResultOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVocabularyResult) string { return v.Content }).(pulumi.StringOutput)
}

// The reason why the custom vocabulary was not created.
func (o LookupVocabularyResultOutput) FailureReason() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVocabularyResult) string { return v.FailureReason }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVocabularyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVocabularyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVocabularyResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVocabularyResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// The language code of the vocabulary entries. For a list of languages and their corresponding language codes, see [What is Amazon Transcribe?](https://docs.aws.amazon.com/transcribe/latest/dg/transcribe-whatis.html). Valid Values are `ar-AE`, `de-CH`, `de-DE`, `en-AB`, `en-AU`, `en-GB`, `en-IE`, `en-IN`, `en-US`, `en-WL`, `es-ES`, `es-US`, `fr-CA`, `fr-FR`, `hi-IN`, `it-IT`, `ja-JP`, `ko-KR`, `pt-BR`, `pt-PT`, `zh-CN`.
func (o LookupVocabularyResultOutput) LanguageCode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVocabularyResult) string { return v.LanguageCode }).(pulumi.StringOutput)
}

// The timestamp when the custom vocabulary was last modified.
func (o LookupVocabularyResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVocabularyResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o LookupVocabularyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVocabularyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The current state of the custom vocabulary. Valid values are `CREATION_IN_PROGRESS`, `ACTIVE`, `CREATION_FAILED`, `DELETE_IN_PROGRESS`.
func (o LookupVocabularyResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVocabularyResult) string { return v.State }).(pulumi.StringOutput)
}

// A map of tags to assign to the Vocabulary.
func (o LookupVocabularyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVocabularyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The identifier of the custom vocabulary.
func (o LookupVocabularyResultOutput) VocabularyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVocabularyResult) string { return v.VocabularyId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVocabularyResultOutput{})
}
