// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for managing a QuickSight Template.
//
// ## Example Usage
// ### From Source Template
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := quicksight.NewTemplate(ctx, "example", &quicksight.TemplateArgs{
//				TemplateId:         pulumi.String("example-id"),
//				VersionDescription: pulumi.String("version"),
//				SourceEntity: &quicksight.TemplateSourceEntityArgs{
//					SourceTemplate: &quicksight.TemplateSourceEntitySourceTemplateArgs{
//						Arn: pulumi.Any(aws_quicksight_template.Source.Arn),
//					},
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
// Using `pulumi import`, import a QuickSight Template using the AWS account ID and template ID separated by a comma (`,`). For example:
//
// ```sh
//
//	$ pulumi import aws:quicksight/template:Template example 123456789012,example-id
//
// ```
type Template struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the resource.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// AWS account ID.
	AwsAccountId pulumi.StringOutput `pulumi:"awsAccountId"`
	// The time that the template was created.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The time that the template was last updated.
	LastUpdatedTime pulumi.StringOutput `pulumi:"lastUpdatedTime"`
	// Display name for the template.
	Name pulumi.StringOutput `pulumi:"name"`
	// A set of resource permissions on the template. Maximum of 64 items. See permissions.
	Permissions TemplatePermissionArrayOutput `pulumi:"permissions"`
	// The entity that you are using as a source when you create the template (analysis or template). Only one of `definition` or `sourceEntity` should be configured. See source_entity.
	SourceEntity TemplateSourceEntityPtrOutput `pulumi:"sourceEntity"`
	// Amazon Resource Name (ARN) of an analysis or template that was used to create this template.
	SourceEntityArn pulumi.StringOutput `pulumi:"sourceEntityArn"`
	// The template creation status.
	Status pulumi.StringOutput `pulumi:"status"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Identifier for the template.
	TemplateId pulumi.StringOutput `pulumi:"templateId"`
	// A description of the current template version being created/updated.
	//
	// The following arguments are optional:
	VersionDescription pulumi.StringOutput `pulumi:"versionDescription"`
	// The version number of the template version.
	VersionNumber pulumi.IntOutput `pulumi:"versionNumber"`
}

// NewTemplate registers a new resource with the given unique name, arguments, and options.
func NewTemplate(ctx *pulumi.Context,
	name string, args *TemplateArgs, opts ...pulumi.ResourceOption) (*Template, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TemplateId == nil {
		return nil, errors.New("invalid value for required argument 'TemplateId'")
	}
	if args.VersionDescription == nil {
		return nil, errors.New("invalid value for required argument 'VersionDescription'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Template
	err := ctx.RegisterResource("aws:quicksight/template:Template", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTemplate gets an existing Template resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateState, opts ...pulumi.ResourceOption) (*Template, error) {
	var resource Template
	err := ctx.ReadResource("aws:quicksight/template:Template", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Template resources.
type templateState struct {
	// The Amazon Resource Name (ARN) of the resource.
	Arn *string `pulumi:"arn"`
	// AWS account ID.
	AwsAccountId *string `pulumi:"awsAccountId"`
	// The time that the template was created.
	CreatedTime *string `pulumi:"createdTime"`
	// The time that the template was last updated.
	LastUpdatedTime *string `pulumi:"lastUpdatedTime"`
	// Display name for the template.
	Name *string `pulumi:"name"`
	// A set of resource permissions on the template. Maximum of 64 items. See permissions.
	Permissions []TemplatePermission `pulumi:"permissions"`
	// The entity that you are using as a source when you create the template (analysis or template). Only one of `definition` or `sourceEntity` should be configured. See source_entity.
	SourceEntity *TemplateSourceEntity `pulumi:"sourceEntity"`
	// Amazon Resource Name (ARN) of an analysis or template that was used to create this template.
	SourceEntityArn *string `pulumi:"sourceEntityArn"`
	// The template creation status.
	Status *string `pulumi:"status"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Identifier for the template.
	TemplateId *string `pulumi:"templateId"`
	// A description of the current template version being created/updated.
	//
	// The following arguments are optional:
	VersionDescription *string `pulumi:"versionDescription"`
	// The version number of the template version.
	VersionNumber *int `pulumi:"versionNumber"`
}

type TemplateState struct {
	// The Amazon Resource Name (ARN) of the resource.
	Arn pulumi.StringPtrInput
	// AWS account ID.
	AwsAccountId pulumi.StringPtrInput
	// The time that the template was created.
	CreatedTime pulumi.StringPtrInput
	// The time that the template was last updated.
	LastUpdatedTime pulumi.StringPtrInput
	// Display name for the template.
	Name pulumi.StringPtrInput
	// A set of resource permissions on the template. Maximum of 64 items. See permissions.
	Permissions TemplatePermissionArrayInput
	// The entity that you are using as a source when you create the template (analysis or template). Only one of `definition` or `sourceEntity` should be configured. See source_entity.
	SourceEntity TemplateSourceEntityPtrInput
	// Amazon Resource Name (ARN) of an analysis or template that was used to create this template.
	SourceEntityArn pulumi.StringPtrInput
	// The template creation status.
	Status pulumi.StringPtrInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// Identifier for the template.
	TemplateId pulumi.StringPtrInput
	// A description of the current template version being created/updated.
	//
	// The following arguments are optional:
	VersionDescription pulumi.StringPtrInput
	// The version number of the template version.
	VersionNumber pulumi.IntPtrInput
}

func (TemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateState)(nil)).Elem()
}

type templateArgs struct {
	// AWS account ID.
	AwsAccountId *string `pulumi:"awsAccountId"`
	// Display name for the template.
	Name *string `pulumi:"name"`
	// A set of resource permissions on the template. Maximum of 64 items. See permissions.
	Permissions []TemplatePermission `pulumi:"permissions"`
	// The entity that you are using as a source when you create the template (analysis or template). Only one of `definition` or `sourceEntity` should be configured. See source_entity.
	SourceEntity *TemplateSourceEntity `pulumi:"sourceEntity"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Identifier for the template.
	TemplateId string `pulumi:"templateId"`
	// A description of the current template version being created/updated.
	//
	// The following arguments are optional:
	VersionDescription string `pulumi:"versionDescription"`
}

// The set of arguments for constructing a Template resource.
type TemplateArgs struct {
	// AWS account ID.
	AwsAccountId pulumi.StringPtrInput
	// Display name for the template.
	Name pulumi.StringPtrInput
	// A set of resource permissions on the template. Maximum of 64 items. See permissions.
	Permissions TemplatePermissionArrayInput
	// The entity that you are using as a source when you create the template (analysis or template). Only one of `definition` or `sourceEntity` should be configured. See source_entity.
	SourceEntity TemplateSourceEntityPtrInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Identifier for the template.
	TemplateId pulumi.StringInput
	// A description of the current template version being created/updated.
	//
	// The following arguments are optional:
	VersionDescription pulumi.StringInput
}

func (TemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateArgs)(nil)).Elem()
}

type TemplateInput interface {
	pulumi.Input

	ToTemplateOutput() TemplateOutput
	ToTemplateOutputWithContext(ctx context.Context) TemplateOutput
}

func (*Template) ElementType() reflect.Type {
	return reflect.TypeOf((**Template)(nil)).Elem()
}

func (i *Template) ToTemplateOutput() TemplateOutput {
	return i.ToTemplateOutputWithContext(context.Background())
}

func (i *Template) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateOutput)
}

// TemplateArrayInput is an input type that accepts TemplateArray and TemplateArrayOutput values.
// You can construct a concrete instance of `TemplateArrayInput` via:
//
//	TemplateArray{ TemplateArgs{...} }
type TemplateArrayInput interface {
	pulumi.Input

	ToTemplateArrayOutput() TemplateArrayOutput
	ToTemplateArrayOutputWithContext(context.Context) TemplateArrayOutput
}

type TemplateArray []TemplateInput

func (TemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Template)(nil)).Elem()
}

func (i TemplateArray) ToTemplateArrayOutput() TemplateArrayOutput {
	return i.ToTemplateArrayOutputWithContext(context.Background())
}

func (i TemplateArray) ToTemplateArrayOutputWithContext(ctx context.Context) TemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateArrayOutput)
}

// TemplateMapInput is an input type that accepts TemplateMap and TemplateMapOutput values.
// You can construct a concrete instance of `TemplateMapInput` via:
//
//	TemplateMap{ "key": TemplateArgs{...} }
type TemplateMapInput interface {
	pulumi.Input

	ToTemplateMapOutput() TemplateMapOutput
	ToTemplateMapOutputWithContext(context.Context) TemplateMapOutput
}

type TemplateMap map[string]TemplateInput

func (TemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Template)(nil)).Elem()
}

func (i TemplateMap) ToTemplateMapOutput() TemplateMapOutput {
	return i.ToTemplateMapOutputWithContext(context.Background())
}

func (i TemplateMap) ToTemplateMapOutputWithContext(ctx context.Context) TemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateMapOutput)
}

type TemplateOutput struct{ *pulumi.OutputState }

func (TemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Template)(nil)).Elem()
}

func (o TemplateOutput) ToTemplateOutput() TemplateOutput {
	return o
}

func (o TemplateOutput) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return o
}

// The Amazon Resource Name (ARN) of the resource.
func (o TemplateOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// AWS account ID.
func (o TemplateOutput) AwsAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.AwsAccountId }).(pulumi.StringOutput)
}

// The time that the template was created.
func (o TemplateOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// The time that the template was last updated.
func (o TemplateOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

// Display name for the template.
func (o TemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A set of resource permissions on the template. Maximum of 64 items. See permissions.
func (o TemplateOutput) Permissions() TemplatePermissionArrayOutput {
	return o.ApplyT(func(v *Template) TemplatePermissionArrayOutput { return v.Permissions }).(TemplatePermissionArrayOutput)
}

// The entity that you are using as a source when you create the template (analysis or template). Only one of `definition` or `sourceEntity` should be configured. See source_entity.
func (o TemplateOutput) SourceEntity() TemplateSourceEntityPtrOutput {
	return o.ApplyT(func(v *Template) TemplateSourceEntityPtrOutput { return v.SourceEntity }).(TemplateSourceEntityPtrOutput)
}

// Amazon Resource Name (ARN) of an analysis or template that was used to create this template.
func (o TemplateOutput) SourceEntityArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.SourceEntityArn }).(pulumi.StringOutput)
}

// The template creation status.
func (o TemplateOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o TemplateOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Template) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o TemplateOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Template) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Identifier for the template.
func (o TemplateOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.TemplateId }).(pulumi.StringOutput)
}

// A description of the current template version being created/updated.
//
// The following arguments are optional:
func (o TemplateOutput) VersionDescription() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.VersionDescription }).(pulumi.StringOutput)
}

// The version number of the template version.
func (o TemplateOutput) VersionNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *Template) pulumi.IntOutput { return v.VersionNumber }).(pulumi.IntOutput)
}

type TemplateArrayOutput struct{ *pulumi.OutputState }

func (TemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Template)(nil)).Elem()
}

func (o TemplateArrayOutput) ToTemplateArrayOutput() TemplateArrayOutput {
	return o
}

func (o TemplateArrayOutput) ToTemplateArrayOutputWithContext(ctx context.Context) TemplateArrayOutput {
	return o
}

func (o TemplateArrayOutput) Index(i pulumi.IntInput) TemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Template {
		return vs[0].([]*Template)[vs[1].(int)]
	}).(TemplateOutput)
}

type TemplateMapOutput struct{ *pulumi.OutputState }

func (TemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Template)(nil)).Elem()
}

func (o TemplateMapOutput) ToTemplateMapOutput() TemplateMapOutput {
	return o
}

func (o TemplateMapOutput) ToTemplateMapOutputWithContext(ctx context.Context) TemplateMapOutput {
	return o
}

func (o TemplateMapOutput) MapIndex(k pulumi.StringInput) TemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Template {
		return vs[0].(map[string]*Template)[vs[1].(string)]
	}).(TemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateInput)(nil)).Elem(), &Template{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateArrayInput)(nil)).Elem(), TemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateMapInput)(nil)).Elem(), TemplateMap{})
	pulumi.RegisterOutputType(TemplateOutput{})
	pulumi.RegisterOutputType(TemplateArrayOutput{})
	pulumi.RegisterOutputType(TemplateMapOutput{})
}
