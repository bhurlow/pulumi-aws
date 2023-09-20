// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kendra

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource for managing an AWS Kendra Thesaurus.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kendra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := kendra.NewThesaurus(ctx, "example", &kendra.ThesaurusArgs{
//				IndexId: pulumi.Any(aws_kendra_index.Example.Id),
//				RoleArn: pulumi.Any(aws_iam_role.Example.Arn),
//				SourceS3Path: &kendra.ThesaurusSourceS3PathArgs{
//					Bucket: pulumi.Any(aws_s3_bucket.Example.Id),
//					Key:    pulumi.Any(aws_s3_object.Example.Key),
//				},
//				Tags: pulumi.StringMap{
//					"Name": pulumi.String("Example Kendra Thesaurus"),
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
// Using `pulumi import`, import `aws_kendra_thesaurus` using the unique identifiers of the thesaurus and index separated by a slash (`/`). For example:
//
// ```sh
//
//	$ pulumi import aws:kendra/thesaurus:Thesaurus example thesaurus-123456780/idx-8012925589
//
// ```
type Thesaurus struct {
	pulumi.CustomResourceState

	// ARN of the thesaurus.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description for a thesaurus.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The identifier of the index for a thesaurus.
	IndexId pulumi.StringOutput `pulumi:"indexId"`
	// The name for the thesaurus.
	Name pulumi.StringOutput `pulumi:"name"`
	// The IAM (Identity and Access Management) role used to access the thesaurus file in S3.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// The S3 path where your thesaurus file sits in S3. Detailed below.
	//
	// The `sourceS3Path` configuration block supports the following arguments:
	SourceS3Path ThesaurusSourceS3PathOutput `pulumi:"sourceS3Path"`
	// The current status of the thesaurus.
	Status pulumi.StringOutput `pulumi:"status"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll     pulumi.StringMapOutput `pulumi:"tagsAll"`
	ThesaurusId pulumi.StringOutput    `pulumi:"thesaurusId"`
}

// NewThesaurus registers a new resource with the given unique name, arguments, and options.
func NewThesaurus(ctx *pulumi.Context,
	name string, args *ThesaurusArgs, opts ...pulumi.ResourceOption) (*Thesaurus, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IndexId == nil {
		return nil, errors.New("invalid value for required argument 'IndexId'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	if args.SourceS3Path == nil {
		return nil, errors.New("invalid value for required argument 'SourceS3Path'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Thesaurus
	err := ctx.RegisterResource("aws:kendra/thesaurus:Thesaurus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetThesaurus gets an existing Thesaurus resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetThesaurus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ThesaurusState, opts ...pulumi.ResourceOption) (*Thesaurus, error) {
	var resource Thesaurus
	err := ctx.ReadResource("aws:kendra/thesaurus:Thesaurus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Thesaurus resources.
type thesaurusState struct {
	// ARN of the thesaurus.
	Arn *string `pulumi:"arn"`
	// The description for a thesaurus.
	Description *string `pulumi:"description"`
	// The identifier of the index for a thesaurus.
	IndexId *string `pulumi:"indexId"`
	// The name for the thesaurus.
	Name *string `pulumi:"name"`
	// The IAM (Identity and Access Management) role used to access the thesaurus file in S3.
	RoleArn *string `pulumi:"roleArn"`
	// The S3 path where your thesaurus file sits in S3. Detailed below.
	//
	// The `sourceS3Path` configuration block supports the following arguments:
	SourceS3Path *ThesaurusSourceS3Path `pulumi:"sourceS3Path"`
	// The current status of the thesaurus.
	Status *string `pulumi:"status"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll     map[string]string `pulumi:"tagsAll"`
	ThesaurusId *string           `pulumi:"thesaurusId"`
}

type ThesaurusState struct {
	// ARN of the thesaurus.
	Arn pulumi.StringPtrInput
	// The description for a thesaurus.
	Description pulumi.StringPtrInput
	// The identifier of the index for a thesaurus.
	IndexId pulumi.StringPtrInput
	// The name for the thesaurus.
	Name pulumi.StringPtrInput
	// The IAM (Identity and Access Management) role used to access the thesaurus file in S3.
	RoleArn pulumi.StringPtrInput
	// The S3 path where your thesaurus file sits in S3. Detailed below.
	//
	// The `sourceS3Path` configuration block supports the following arguments:
	SourceS3Path ThesaurusSourceS3PathPtrInput
	// The current status of the thesaurus.
	Status pulumi.StringPtrInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll     pulumi.StringMapInput
	ThesaurusId pulumi.StringPtrInput
}

func (ThesaurusState) ElementType() reflect.Type {
	return reflect.TypeOf((*thesaurusState)(nil)).Elem()
}

type thesaurusArgs struct {
	// The description for a thesaurus.
	Description *string `pulumi:"description"`
	// The identifier of the index for a thesaurus.
	IndexId string `pulumi:"indexId"`
	// The name for the thesaurus.
	Name *string `pulumi:"name"`
	// The IAM (Identity and Access Management) role used to access the thesaurus file in S3.
	RoleArn string `pulumi:"roleArn"`
	// The S3 path where your thesaurus file sits in S3. Detailed below.
	//
	// The `sourceS3Path` configuration block supports the following arguments:
	SourceS3Path ThesaurusSourceS3Path `pulumi:"sourceS3Path"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Thesaurus resource.
type ThesaurusArgs struct {
	// The description for a thesaurus.
	Description pulumi.StringPtrInput
	// The identifier of the index for a thesaurus.
	IndexId pulumi.StringInput
	// The name for the thesaurus.
	Name pulumi.StringPtrInput
	// The IAM (Identity and Access Management) role used to access the thesaurus file in S3.
	RoleArn pulumi.StringInput
	// The S3 path where your thesaurus file sits in S3. Detailed below.
	//
	// The `sourceS3Path` configuration block supports the following arguments:
	SourceS3Path ThesaurusSourceS3PathInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (ThesaurusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*thesaurusArgs)(nil)).Elem()
}

type ThesaurusInput interface {
	pulumi.Input

	ToThesaurusOutput() ThesaurusOutput
	ToThesaurusOutputWithContext(ctx context.Context) ThesaurusOutput
}

func (*Thesaurus) ElementType() reflect.Type {
	return reflect.TypeOf((**Thesaurus)(nil)).Elem()
}

func (i *Thesaurus) ToThesaurusOutput() ThesaurusOutput {
	return i.ToThesaurusOutputWithContext(context.Background())
}

func (i *Thesaurus) ToThesaurusOutputWithContext(ctx context.Context) ThesaurusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThesaurusOutput)
}

func (i *Thesaurus) ToOutput(ctx context.Context) pulumix.Output[*Thesaurus] {
	return pulumix.Output[*Thesaurus]{
		OutputState: i.ToThesaurusOutputWithContext(ctx).OutputState,
	}
}

// ThesaurusArrayInput is an input type that accepts ThesaurusArray and ThesaurusArrayOutput values.
// You can construct a concrete instance of `ThesaurusArrayInput` via:
//
//	ThesaurusArray{ ThesaurusArgs{...} }
type ThesaurusArrayInput interface {
	pulumi.Input

	ToThesaurusArrayOutput() ThesaurusArrayOutput
	ToThesaurusArrayOutputWithContext(context.Context) ThesaurusArrayOutput
}

type ThesaurusArray []ThesaurusInput

func (ThesaurusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Thesaurus)(nil)).Elem()
}

func (i ThesaurusArray) ToThesaurusArrayOutput() ThesaurusArrayOutput {
	return i.ToThesaurusArrayOutputWithContext(context.Background())
}

func (i ThesaurusArray) ToThesaurusArrayOutputWithContext(ctx context.Context) ThesaurusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThesaurusArrayOutput)
}

func (i ThesaurusArray) ToOutput(ctx context.Context) pulumix.Output[[]*Thesaurus] {
	return pulumix.Output[[]*Thesaurus]{
		OutputState: i.ToThesaurusArrayOutputWithContext(ctx).OutputState,
	}
}

// ThesaurusMapInput is an input type that accepts ThesaurusMap and ThesaurusMapOutput values.
// You can construct a concrete instance of `ThesaurusMapInput` via:
//
//	ThesaurusMap{ "key": ThesaurusArgs{...} }
type ThesaurusMapInput interface {
	pulumi.Input

	ToThesaurusMapOutput() ThesaurusMapOutput
	ToThesaurusMapOutputWithContext(context.Context) ThesaurusMapOutput
}

type ThesaurusMap map[string]ThesaurusInput

func (ThesaurusMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Thesaurus)(nil)).Elem()
}

func (i ThesaurusMap) ToThesaurusMapOutput() ThesaurusMapOutput {
	return i.ToThesaurusMapOutputWithContext(context.Background())
}

func (i ThesaurusMap) ToThesaurusMapOutputWithContext(ctx context.Context) ThesaurusMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThesaurusMapOutput)
}

func (i ThesaurusMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Thesaurus] {
	return pulumix.Output[map[string]*Thesaurus]{
		OutputState: i.ToThesaurusMapOutputWithContext(ctx).OutputState,
	}
}

type ThesaurusOutput struct{ *pulumi.OutputState }

func (ThesaurusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Thesaurus)(nil)).Elem()
}

func (o ThesaurusOutput) ToThesaurusOutput() ThesaurusOutput {
	return o
}

func (o ThesaurusOutput) ToThesaurusOutputWithContext(ctx context.Context) ThesaurusOutput {
	return o
}

func (o ThesaurusOutput) ToOutput(ctx context.Context) pulumix.Output[*Thesaurus] {
	return pulumix.Output[*Thesaurus]{
		OutputState: o.OutputState,
	}
}

// ARN of the thesaurus.
func (o ThesaurusOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Thesaurus) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The description for a thesaurus.
func (o ThesaurusOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Thesaurus) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The identifier of the index for a thesaurus.
func (o ThesaurusOutput) IndexId() pulumi.StringOutput {
	return o.ApplyT(func(v *Thesaurus) pulumi.StringOutput { return v.IndexId }).(pulumi.StringOutput)
}

// The name for the thesaurus.
func (o ThesaurusOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Thesaurus) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The IAM (Identity and Access Management) role used to access the thesaurus file in S3.
func (o ThesaurusOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Thesaurus) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

// The S3 path where your thesaurus file sits in S3. Detailed below.
//
// The `sourceS3Path` configuration block supports the following arguments:
func (o ThesaurusOutput) SourceS3Path() ThesaurusSourceS3PathOutput {
	return o.ApplyT(func(v *Thesaurus) ThesaurusSourceS3PathOutput { return v.SourceS3Path }).(ThesaurusSourceS3PathOutput)
}

// The current status of the thesaurus.
func (o ThesaurusOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Thesaurus) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ThesaurusOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Thesaurus) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o ThesaurusOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Thesaurus) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

func (o ThesaurusOutput) ThesaurusId() pulumi.StringOutput {
	return o.ApplyT(func(v *Thesaurus) pulumi.StringOutput { return v.ThesaurusId }).(pulumi.StringOutput)
}

type ThesaurusArrayOutput struct{ *pulumi.OutputState }

func (ThesaurusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Thesaurus)(nil)).Elem()
}

func (o ThesaurusArrayOutput) ToThesaurusArrayOutput() ThesaurusArrayOutput {
	return o
}

func (o ThesaurusArrayOutput) ToThesaurusArrayOutputWithContext(ctx context.Context) ThesaurusArrayOutput {
	return o
}

func (o ThesaurusArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Thesaurus] {
	return pulumix.Output[[]*Thesaurus]{
		OutputState: o.OutputState,
	}
}

func (o ThesaurusArrayOutput) Index(i pulumi.IntInput) ThesaurusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Thesaurus {
		return vs[0].([]*Thesaurus)[vs[1].(int)]
	}).(ThesaurusOutput)
}

type ThesaurusMapOutput struct{ *pulumi.OutputState }

func (ThesaurusMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Thesaurus)(nil)).Elem()
}

func (o ThesaurusMapOutput) ToThesaurusMapOutput() ThesaurusMapOutput {
	return o
}

func (o ThesaurusMapOutput) ToThesaurusMapOutputWithContext(ctx context.Context) ThesaurusMapOutput {
	return o
}

func (o ThesaurusMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Thesaurus] {
	return pulumix.Output[map[string]*Thesaurus]{
		OutputState: o.OutputState,
	}
}

func (o ThesaurusMapOutput) MapIndex(k pulumi.StringInput) ThesaurusOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Thesaurus {
		return vs[0].(map[string]*Thesaurus)[vs[1].(string)]
	}).(ThesaurusOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ThesaurusInput)(nil)).Elem(), &Thesaurus{})
	pulumi.RegisterInputType(reflect.TypeOf((*ThesaurusArrayInput)(nil)).Elem(), ThesaurusArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ThesaurusMapInput)(nil)).Elem(), ThesaurusMap{})
	pulumi.RegisterOutputType(ThesaurusOutput{})
	pulumi.RegisterOutputType(ThesaurusArrayOutput{})
	pulumi.RegisterOutputType(ThesaurusMapOutput{})
}
