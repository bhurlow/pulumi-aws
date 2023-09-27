// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package schemas

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides an EventBridge Schema resource.
//
// > **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/schemas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testRegistry, err := schemas.NewRegistry(ctx, "testRegistry", nil)
//			if err != nil {
//				return err
//			}
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"openapi": "3.0.0",
//				"info": map[string]interface{}{
//					"version": "1.0.0",
//					"title":   "Event",
//				},
//				"paths": nil,
//				"components": map[string]interface{}{
//					"schemas": map[string]interface{}{
//						"Event": map[string]interface{}{
//							"type": "object",
//							"properties": map[string]interface{}{
//								"name": map[string]interface{}{
//									"type": "string",
//								},
//							},
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = schemas.NewSchema(ctx, "testSchema", &schemas.SchemaArgs{
//				RegistryName: testRegistry.Name,
//				Type:         pulumi.String("OpenApi3"),
//				Description:  pulumi.String("The schema definition for my event"),
//				Content:      pulumi.String(json0),
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
// Using `pulumi import`, import EventBridge schema using the `name` and `registry_name`. For example:
//
// ```sh
//
//	$ pulumi import aws:schemas/schema:Schema test name/registry
//
// ```
type Schema struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the discoverer.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The schema specification. Must be a valid Open API 3.0 spec.
	Content pulumi.StringOutput `pulumi:"content"`
	// The description of the schema. Maximum of 256 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The last modified date of the schema.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// The name of the schema. Maximum of 385 characters consisting of lower case letters, upper case letters, ., -, _, @.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the registry in which this schema belongs.
	RegistryName pulumi.StringOutput `pulumi:"registryName"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The type of the schema. Valid values: `OpenApi3` or `JSONSchemaDraft4`.
	Type pulumi.StringOutput `pulumi:"type"`
	// The version of the schema.
	Version pulumi.StringOutput `pulumi:"version"`
	// The created date of the version of the schema.
	VersionCreatedDate pulumi.StringOutput `pulumi:"versionCreatedDate"`
}

// NewSchema registers a new resource with the given unique name, arguments, and options.
func NewSchema(ctx *pulumi.Context,
	name string, args *SchemaArgs, opts ...pulumi.ResourceOption) (*Schema, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Schema
	err := ctx.RegisterResource("aws:schemas/schema:Schema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSchema gets an existing Schema resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SchemaState, opts ...pulumi.ResourceOption) (*Schema, error) {
	var resource Schema
	err := ctx.ReadResource("aws:schemas/schema:Schema", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Schema resources.
type schemaState struct {
	// The Amazon Resource Name (ARN) of the discoverer.
	Arn *string `pulumi:"arn"`
	// The schema specification. Must be a valid Open API 3.0 spec.
	Content *string `pulumi:"content"`
	// The description of the schema. Maximum of 256 characters.
	Description *string `pulumi:"description"`
	// The last modified date of the schema.
	LastModified *string `pulumi:"lastModified"`
	// The name of the schema. Maximum of 385 characters consisting of lower case letters, upper case letters, ., -, _, @.
	Name *string `pulumi:"name"`
	// The name of the registry in which this schema belongs.
	RegistryName *string `pulumi:"registryName"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The type of the schema. Valid values: `OpenApi3` or `JSONSchemaDraft4`.
	Type *string `pulumi:"type"`
	// The version of the schema.
	Version *string `pulumi:"version"`
	// The created date of the version of the schema.
	VersionCreatedDate *string `pulumi:"versionCreatedDate"`
}

type SchemaState struct {
	// The Amazon Resource Name (ARN) of the discoverer.
	Arn pulumi.StringPtrInput
	// The schema specification. Must be a valid Open API 3.0 spec.
	Content pulumi.StringPtrInput
	// The description of the schema. Maximum of 256 characters.
	Description pulumi.StringPtrInput
	// The last modified date of the schema.
	LastModified pulumi.StringPtrInput
	// The name of the schema. Maximum of 385 characters consisting of lower case letters, upper case letters, ., -, _, @.
	Name pulumi.StringPtrInput
	// The name of the registry in which this schema belongs.
	RegistryName pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// The type of the schema. Valid values: `OpenApi3` or `JSONSchemaDraft4`.
	Type pulumi.StringPtrInput
	// The version of the schema.
	Version pulumi.StringPtrInput
	// The created date of the version of the schema.
	VersionCreatedDate pulumi.StringPtrInput
}

func (SchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaState)(nil)).Elem()
}

type schemaArgs struct {
	// The schema specification. Must be a valid Open API 3.0 spec.
	Content string `pulumi:"content"`
	// The description of the schema. Maximum of 256 characters.
	Description *string `pulumi:"description"`
	// The name of the schema. Maximum of 385 characters consisting of lower case letters, upper case letters, ., -, _, @.
	Name *string `pulumi:"name"`
	// The name of the registry in which this schema belongs.
	RegistryName string `pulumi:"registryName"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The type of the schema. Valid values: `OpenApi3` or `JSONSchemaDraft4`.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Schema resource.
type SchemaArgs struct {
	// The schema specification. Must be a valid Open API 3.0 spec.
	Content pulumi.StringInput
	// The description of the schema. Maximum of 256 characters.
	Description pulumi.StringPtrInput
	// The name of the schema. Maximum of 385 characters consisting of lower case letters, upper case letters, ., -, _, @.
	Name pulumi.StringPtrInput
	// The name of the registry in which this schema belongs.
	RegistryName pulumi.StringInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The type of the schema. Valid values: `OpenApi3` or `JSONSchemaDraft4`.
	Type pulumi.StringInput
}

func (SchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaArgs)(nil)).Elem()
}

type SchemaInput interface {
	pulumi.Input

	ToSchemaOutput() SchemaOutput
	ToSchemaOutputWithContext(ctx context.Context) SchemaOutput
}

func (*Schema) ElementType() reflect.Type {
	return reflect.TypeOf((**Schema)(nil)).Elem()
}

func (i *Schema) ToSchemaOutput() SchemaOutput {
	return i.ToSchemaOutputWithContext(context.Background())
}

func (i *Schema) ToSchemaOutputWithContext(ctx context.Context) SchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaOutput)
}

func (i *Schema) ToOutput(ctx context.Context) pulumix.Output[*Schema] {
	return pulumix.Output[*Schema]{
		OutputState: i.ToSchemaOutputWithContext(ctx).OutputState,
	}
}

// SchemaArrayInput is an input type that accepts SchemaArray and SchemaArrayOutput values.
// You can construct a concrete instance of `SchemaArrayInput` via:
//
//	SchemaArray{ SchemaArgs{...} }
type SchemaArrayInput interface {
	pulumi.Input

	ToSchemaArrayOutput() SchemaArrayOutput
	ToSchemaArrayOutputWithContext(context.Context) SchemaArrayOutput
}

type SchemaArray []SchemaInput

func (SchemaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Schema)(nil)).Elem()
}

func (i SchemaArray) ToSchemaArrayOutput() SchemaArrayOutput {
	return i.ToSchemaArrayOutputWithContext(context.Background())
}

func (i SchemaArray) ToSchemaArrayOutputWithContext(ctx context.Context) SchemaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaArrayOutput)
}

func (i SchemaArray) ToOutput(ctx context.Context) pulumix.Output[[]*Schema] {
	return pulumix.Output[[]*Schema]{
		OutputState: i.ToSchemaArrayOutputWithContext(ctx).OutputState,
	}
}

// SchemaMapInput is an input type that accepts SchemaMap and SchemaMapOutput values.
// You can construct a concrete instance of `SchemaMapInput` via:
//
//	SchemaMap{ "key": SchemaArgs{...} }
type SchemaMapInput interface {
	pulumi.Input

	ToSchemaMapOutput() SchemaMapOutput
	ToSchemaMapOutputWithContext(context.Context) SchemaMapOutput
}

type SchemaMap map[string]SchemaInput

func (SchemaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Schema)(nil)).Elem()
}

func (i SchemaMap) ToSchemaMapOutput() SchemaMapOutput {
	return i.ToSchemaMapOutputWithContext(context.Background())
}

func (i SchemaMap) ToSchemaMapOutputWithContext(ctx context.Context) SchemaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaMapOutput)
}

func (i SchemaMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Schema] {
	return pulumix.Output[map[string]*Schema]{
		OutputState: i.ToSchemaMapOutputWithContext(ctx).OutputState,
	}
}

type SchemaOutput struct{ *pulumi.OutputState }

func (SchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Schema)(nil)).Elem()
}

func (o SchemaOutput) ToSchemaOutput() SchemaOutput {
	return o
}

func (o SchemaOutput) ToSchemaOutputWithContext(ctx context.Context) SchemaOutput {
	return o
}

func (o SchemaOutput) ToOutput(ctx context.Context) pulumix.Output[*Schema] {
	return pulumix.Output[*Schema]{
		OutputState: o.OutputState,
	}
}

// The Amazon Resource Name (ARN) of the discoverer.
func (o SchemaOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The schema specification. Must be a valid Open API 3.0 spec.
func (o SchemaOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

// The description of the schema. Maximum of 256 characters.
func (o SchemaOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The last modified date of the schema.
func (o SchemaOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.LastModified }).(pulumi.StringOutput)
}

// The name of the schema. Maximum of 385 characters consisting of lower case letters, upper case letters, ., -, _, @.
func (o SchemaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the registry in which this schema belongs.
func (o SchemaOutput) RegistryName() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.RegistryName }).(pulumi.StringOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o SchemaOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o SchemaOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The type of the schema. Valid values: `OpenApi3` or `JSONSchemaDraft4`.
func (o SchemaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The version of the schema.
func (o SchemaOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

// The created date of the version of the schema.
func (o SchemaOutput) VersionCreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.VersionCreatedDate }).(pulumi.StringOutput)
}

type SchemaArrayOutput struct{ *pulumi.OutputState }

func (SchemaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Schema)(nil)).Elem()
}

func (o SchemaArrayOutput) ToSchemaArrayOutput() SchemaArrayOutput {
	return o
}

func (o SchemaArrayOutput) ToSchemaArrayOutputWithContext(ctx context.Context) SchemaArrayOutput {
	return o
}

func (o SchemaArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Schema] {
	return pulumix.Output[[]*Schema]{
		OutputState: o.OutputState,
	}
}

func (o SchemaArrayOutput) Index(i pulumi.IntInput) SchemaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Schema {
		return vs[0].([]*Schema)[vs[1].(int)]
	}).(SchemaOutput)
}

type SchemaMapOutput struct{ *pulumi.OutputState }

func (SchemaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Schema)(nil)).Elem()
}

func (o SchemaMapOutput) ToSchemaMapOutput() SchemaMapOutput {
	return o
}

func (o SchemaMapOutput) ToSchemaMapOutputWithContext(ctx context.Context) SchemaMapOutput {
	return o
}

func (o SchemaMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Schema] {
	return pulumix.Output[map[string]*Schema]{
		OutputState: o.OutputState,
	}
}

func (o SchemaMapOutput) MapIndex(k pulumi.StringInput) SchemaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Schema {
		return vs[0].(map[string]*Schema)[vs[1].(string)]
	}).(SchemaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaInput)(nil)).Elem(), &Schema{})
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaArrayInput)(nil)).Elem(), SchemaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaMapInput)(nil)).Elem(), SchemaMap{})
	pulumi.RegisterOutputType(SchemaOutput{})
	pulumi.RegisterOutputType(SchemaArrayOutput{})
	pulumi.RegisterOutputType(SchemaMapOutput{})
}
