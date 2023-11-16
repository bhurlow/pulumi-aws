// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package macie

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage an [AWS Macie Custom Data Identifier](https://docs.aws.amazon.com/macie/latest/APIReference/custom-data-identifiers-id.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/macie"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/macie2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := macie2.NewAccount(ctx, "exampleAccount", nil)
//			if err != nil {
//				return err
//			}
//			_, err = macie.NewCustomDataIdentifier(ctx, "exampleCustomDataIdentifier", &macie.CustomDataIdentifierArgs{
//				Regex:                pulumi.String("[0-9]{3}-[0-9]{2}-[0-9]{4}"),
//				Description:          pulumi.String("DESCRIPTION"),
//				MaximumMatchDistance: pulumi.Int(10),
//				Keywords: pulumi.StringArray{
//					pulumi.String("keyword"),
//				},
//				IgnoreWords: pulumi.StringArray{
//					pulumi.String("ignore"),
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				aws_macie2_account.Test,
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
// Using `pulumi import`, import `aws_macie2_custom_data_identifier` using the id. For example:
//
// ```sh
//
//	$ pulumi import aws:macie/customDataIdentifier:CustomDataIdentifier example abcd1
//
// ```
type CustomDataIdentifier struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the custom data identifier.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The date and time, in UTC and extended RFC 3339 format, when the Amazon Macie account was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// A custom description of the custom data identifier. The description can contain as many as 512 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// An array that lists specific character sequences (ignore words) to exclude from the results. If the text matched by the regular expression is the same as any string in this array, Amazon Macie ignores it. The array can contain as many as 10 ignore words. Each ignore word can contain 4 - 90 characters. Ignore words are case sensitive.
	IgnoreWords pulumi.StringArrayOutput `pulumi:"ignoreWords"`
	// An array that lists specific character sequences (keywords), one of which must be within proximity (`maximumMatchDistance`) of the regular expression to match. The array can contain as many as 50 keywords. Each keyword can contain 3 - 90 characters. Keywords aren't case sensitive.
	Keywords pulumi.StringArrayOutput `pulumi:"keywords"`
	// The maximum number of characters that can exist between text that matches the regex pattern and the character sequences specified by the keywords array. Macie includes or excludes a result based on the proximity of a keyword to text that matches the regex pattern. The distance can be 1 - 300 characters. The default value is 50.
	MaximumMatchDistance pulumi.IntOutput `pulumi:"maximumMatchDistance"`
	// A custom name for the custom data identifier. The name can contain as many as 128 characters. If omitted, the provider will assign a random, unique name. Conflicts with `namePrefix`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// The regular expression (regex) that defines the pattern to match. The expression can contain as many as 512 characters.
	Regex pulumi.StringPtrOutput `pulumi:"regex"`
	// A map of key-value pairs that specifies the tags to associate with the custom data identifier.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewCustomDataIdentifier registers a new resource with the given unique name, arguments, and options.
func NewCustomDataIdentifier(ctx *pulumi.Context,
	name string, args *CustomDataIdentifierArgs, opts ...pulumi.ResourceOption) (*CustomDataIdentifier, error) {
	if args == nil {
		args = &CustomDataIdentifierArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CustomDataIdentifier
	err := ctx.RegisterResource("aws:macie/customDataIdentifier:CustomDataIdentifier", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomDataIdentifier gets an existing CustomDataIdentifier resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomDataIdentifier(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomDataIdentifierState, opts ...pulumi.ResourceOption) (*CustomDataIdentifier, error) {
	var resource CustomDataIdentifier
	err := ctx.ReadResource("aws:macie/customDataIdentifier:CustomDataIdentifier", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomDataIdentifier resources.
type customDataIdentifierState struct {
	// The Amazon Resource Name (ARN) of the custom data identifier.
	Arn *string `pulumi:"arn"`
	// The date and time, in UTC and extended RFC 3339 format, when the Amazon Macie account was created.
	CreatedAt *string `pulumi:"createdAt"`
	// A custom description of the custom data identifier. The description can contain as many as 512 characters.
	Description *string `pulumi:"description"`
	// An array that lists specific character sequences (ignore words) to exclude from the results. If the text matched by the regular expression is the same as any string in this array, Amazon Macie ignores it. The array can contain as many as 10 ignore words. Each ignore word can contain 4 - 90 characters. Ignore words are case sensitive.
	IgnoreWords []string `pulumi:"ignoreWords"`
	// An array that lists specific character sequences (keywords), one of which must be within proximity (`maximumMatchDistance`) of the regular expression to match. The array can contain as many as 50 keywords. Each keyword can contain 3 - 90 characters. Keywords aren't case sensitive.
	Keywords []string `pulumi:"keywords"`
	// The maximum number of characters that can exist between text that matches the regex pattern and the character sequences specified by the keywords array. Macie includes or excludes a result based on the proximity of a keyword to text that matches the regex pattern. The distance can be 1 - 300 characters. The default value is 50.
	MaximumMatchDistance *int `pulumi:"maximumMatchDistance"`
	// A custom name for the custom data identifier. The name can contain as many as 128 characters. If omitted, the provider will assign a random, unique name. Conflicts with `namePrefix`.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The regular expression (regex) that defines the pattern to match. The expression can contain as many as 512 characters.
	Regex *string `pulumi:"regex"`
	// A map of key-value pairs that specifies the tags to associate with the custom data identifier.
	Tags map[string]string `pulumi:"tags"`
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type CustomDataIdentifierState struct {
	// The Amazon Resource Name (ARN) of the custom data identifier.
	Arn pulumi.StringPtrInput
	// The date and time, in UTC and extended RFC 3339 format, when the Amazon Macie account was created.
	CreatedAt pulumi.StringPtrInput
	// A custom description of the custom data identifier. The description can contain as many as 512 characters.
	Description pulumi.StringPtrInput
	// An array that lists specific character sequences (ignore words) to exclude from the results. If the text matched by the regular expression is the same as any string in this array, Amazon Macie ignores it. The array can contain as many as 10 ignore words. Each ignore word can contain 4 - 90 characters. Ignore words are case sensitive.
	IgnoreWords pulumi.StringArrayInput
	// An array that lists specific character sequences (keywords), one of which must be within proximity (`maximumMatchDistance`) of the regular expression to match. The array can contain as many as 50 keywords. Each keyword can contain 3 - 90 characters. Keywords aren't case sensitive.
	Keywords pulumi.StringArrayInput
	// The maximum number of characters that can exist between text that matches the regex pattern and the character sequences specified by the keywords array. Macie includes or excludes a result based on the proximity of a keyword to text that matches the regex pattern. The distance can be 1 - 300 characters. The default value is 50.
	MaximumMatchDistance pulumi.IntPtrInput
	// A custom name for the custom data identifier. The name can contain as many as 128 characters. If omitted, the provider will assign a random, unique name. Conflicts with `namePrefix`.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The regular expression (regex) that defines the pattern to match. The expression can contain as many as 512 characters.
	Regex pulumi.StringPtrInput
	// A map of key-value pairs that specifies the tags to associate with the custom data identifier.
	Tags pulumi.StringMapInput
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (CustomDataIdentifierState) ElementType() reflect.Type {
	return reflect.TypeOf((*customDataIdentifierState)(nil)).Elem()
}

type customDataIdentifierArgs struct {
	// A custom description of the custom data identifier. The description can contain as many as 512 characters.
	Description *string `pulumi:"description"`
	// An array that lists specific character sequences (ignore words) to exclude from the results. If the text matched by the regular expression is the same as any string in this array, Amazon Macie ignores it. The array can contain as many as 10 ignore words. Each ignore word can contain 4 - 90 characters. Ignore words are case sensitive.
	IgnoreWords []string `pulumi:"ignoreWords"`
	// An array that lists specific character sequences (keywords), one of which must be within proximity (`maximumMatchDistance`) of the regular expression to match. The array can contain as many as 50 keywords. Each keyword can contain 3 - 90 characters. Keywords aren't case sensitive.
	Keywords []string `pulumi:"keywords"`
	// The maximum number of characters that can exist between text that matches the regex pattern and the character sequences specified by the keywords array. Macie includes or excludes a result based on the proximity of a keyword to text that matches the regex pattern. The distance can be 1 - 300 characters. The default value is 50.
	MaximumMatchDistance *int `pulumi:"maximumMatchDistance"`
	// A custom name for the custom data identifier. The name can contain as many as 128 characters. If omitted, the provider will assign a random, unique name. Conflicts with `namePrefix`.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The regular expression (regex) that defines the pattern to match. The expression can contain as many as 512 characters.
	Regex *string `pulumi:"regex"`
	// A map of key-value pairs that specifies the tags to associate with the custom data identifier.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CustomDataIdentifier resource.
type CustomDataIdentifierArgs struct {
	// A custom description of the custom data identifier. The description can contain as many as 512 characters.
	Description pulumi.StringPtrInput
	// An array that lists specific character sequences (ignore words) to exclude from the results. If the text matched by the regular expression is the same as any string in this array, Amazon Macie ignores it. The array can contain as many as 10 ignore words. Each ignore word can contain 4 - 90 characters. Ignore words are case sensitive.
	IgnoreWords pulumi.StringArrayInput
	// An array that lists specific character sequences (keywords), one of which must be within proximity (`maximumMatchDistance`) of the regular expression to match. The array can contain as many as 50 keywords. Each keyword can contain 3 - 90 characters. Keywords aren't case sensitive.
	Keywords pulumi.StringArrayInput
	// The maximum number of characters that can exist between text that matches the regex pattern and the character sequences specified by the keywords array. Macie includes or excludes a result based on the proximity of a keyword to text that matches the regex pattern. The distance can be 1 - 300 characters. The default value is 50.
	MaximumMatchDistance pulumi.IntPtrInput
	// A custom name for the custom data identifier. The name can contain as many as 128 characters. If omitted, the provider will assign a random, unique name. Conflicts with `namePrefix`.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The regular expression (regex) that defines the pattern to match. The expression can contain as many as 512 characters.
	Regex pulumi.StringPtrInput
	// A map of key-value pairs that specifies the tags to associate with the custom data identifier.
	Tags pulumi.StringMapInput
}

func (CustomDataIdentifierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customDataIdentifierArgs)(nil)).Elem()
}

type CustomDataIdentifierInput interface {
	pulumi.Input

	ToCustomDataIdentifierOutput() CustomDataIdentifierOutput
	ToCustomDataIdentifierOutputWithContext(ctx context.Context) CustomDataIdentifierOutput
}

func (*CustomDataIdentifier) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDataIdentifier)(nil)).Elem()
}

func (i *CustomDataIdentifier) ToCustomDataIdentifierOutput() CustomDataIdentifierOutput {
	return i.ToCustomDataIdentifierOutputWithContext(context.Background())
}

func (i *CustomDataIdentifier) ToCustomDataIdentifierOutputWithContext(ctx context.Context) CustomDataIdentifierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDataIdentifierOutput)
}

// CustomDataIdentifierArrayInput is an input type that accepts CustomDataIdentifierArray and CustomDataIdentifierArrayOutput values.
// You can construct a concrete instance of `CustomDataIdentifierArrayInput` via:
//
//	CustomDataIdentifierArray{ CustomDataIdentifierArgs{...} }
type CustomDataIdentifierArrayInput interface {
	pulumi.Input

	ToCustomDataIdentifierArrayOutput() CustomDataIdentifierArrayOutput
	ToCustomDataIdentifierArrayOutputWithContext(context.Context) CustomDataIdentifierArrayOutput
}

type CustomDataIdentifierArray []CustomDataIdentifierInput

func (CustomDataIdentifierArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomDataIdentifier)(nil)).Elem()
}

func (i CustomDataIdentifierArray) ToCustomDataIdentifierArrayOutput() CustomDataIdentifierArrayOutput {
	return i.ToCustomDataIdentifierArrayOutputWithContext(context.Background())
}

func (i CustomDataIdentifierArray) ToCustomDataIdentifierArrayOutputWithContext(ctx context.Context) CustomDataIdentifierArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDataIdentifierArrayOutput)
}

// CustomDataIdentifierMapInput is an input type that accepts CustomDataIdentifierMap and CustomDataIdentifierMapOutput values.
// You can construct a concrete instance of `CustomDataIdentifierMapInput` via:
//
//	CustomDataIdentifierMap{ "key": CustomDataIdentifierArgs{...} }
type CustomDataIdentifierMapInput interface {
	pulumi.Input

	ToCustomDataIdentifierMapOutput() CustomDataIdentifierMapOutput
	ToCustomDataIdentifierMapOutputWithContext(context.Context) CustomDataIdentifierMapOutput
}

type CustomDataIdentifierMap map[string]CustomDataIdentifierInput

func (CustomDataIdentifierMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomDataIdentifier)(nil)).Elem()
}

func (i CustomDataIdentifierMap) ToCustomDataIdentifierMapOutput() CustomDataIdentifierMapOutput {
	return i.ToCustomDataIdentifierMapOutputWithContext(context.Background())
}

func (i CustomDataIdentifierMap) ToCustomDataIdentifierMapOutputWithContext(ctx context.Context) CustomDataIdentifierMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDataIdentifierMapOutput)
}

type CustomDataIdentifierOutput struct{ *pulumi.OutputState }

func (CustomDataIdentifierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDataIdentifier)(nil)).Elem()
}

func (o CustomDataIdentifierOutput) ToCustomDataIdentifierOutput() CustomDataIdentifierOutput {
	return o
}

func (o CustomDataIdentifierOutput) ToCustomDataIdentifierOutputWithContext(ctx context.Context) CustomDataIdentifierOutput {
	return o
}

// The Amazon Resource Name (ARN) of the custom data identifier.
func (o CustomDataIdentifierOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDataIdentifier) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The date and time, in UTC and extended RFC 3339 format, when the Amazon Macie account was created.
func (o CustomDataIdentifierOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDataIdentifier) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// A custom description of the custom data identifier. The description can contain as many as 512 characters.
func (o CustomDataIdentifierOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDataIdentifier) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// An array that lists specific character sequences (ignore words) to exclude from the results. If the text matched by the regular expression is the same as any string in this array, Amazon Macie ignores it. The array can contain as many as 10 ignore words. Each ignore word can contain 4 - 90 characters. Ignore words are case sensitive.
func (o CustomDataIdentifierOutput) IgnoreWords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CustomDataIdentifier) pulumi.StringArrayOutput { return v.IgnoreWords }).(pulumi.StringArrayOutput)
}

// An array that lists specific character sequences (keywords), one of which must be within proximity (`maximumMatchDistance`) of the regular expression to match. The array can contain as many as 50 keywords. Each keyword can contain 3 - 90 characters. Keywords aren't case sensitive.
func (o CustomDataIdentifierOutput) Keywords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CustomDataIdentifier) pulumi.StringArrayOutput { return v.Keywords }).(pulumi.StringArrayOutput)
}

// The maximum number of characters that can exist between text that matches the regex pattern and the character sequences specified by the keywords array. Macie includes or excludes a result based on the proximity of a keyword to text that matches the regex pattern. The distance can be 1 - 300 characters. The default value is 50.
func (o CustomDataIdentifierOutput) MaximumMatchDistance() pulumi.IntOutput {
	return o.ApplyT(func(v *CustomDataIdentifier) pulumi.IntOutput { return v.MaximumMatchDistance }).(pulumi.IntOutput)
}

// A custom name for the custom data identifier. The name can contain as many as 128 characters. If omitted, the provider will assign a random, unique name. Conflicts with `namePrefix`.
func (o CustomDataIdentifierOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDataIdentifier) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
func (o CustomDataIdentifierOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDataIdentifier) pulumi.StringOutput { return v.NamePrefix }).(pulumi.StringOutput)
}

// The regular expression (regex) that defines the pattern to match. The expression can contain as many as 512 characters.
func (o CustomDataIdentifierOutput) Regex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDataIdentifier) pulumi.StringPtrOutput { return v.Regex }).(pulumi.StringPtrOutput)
}

// A map of key-value pairs that specifies the tags to associate with the custom data identifier.
func (o CustomDataIdentifierOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CustomDataIdentifier) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Deprecated: Please use `tags` instead.
func (o CustomDataIdentifierOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CustomDataIdentifier) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type CustomDataIdentifierArrayOutput struct{ *pulumi.OutputState }

func (CustomDataIdentifierArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomDataIdentifier)(nil)).Elem()
}

func (o CustomDataIdentifierArrayOutput) ToCustomDataIdentifierArrayOutput() CustomDataIdentifierArrayOutput {
	return o
}

func (o CustomDataIdentifierArrayOutput) ToCustomDataIdentifierArrayOutputWithContext(ctx context.Context) CustomDataIdentifierArrayOutput {
	return o
}

func (o CustomDataIdentifierArrayOutput) Index(i pulumi.IntInput) CustomDataIdentifierOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomDataIdentifier {
		return vs[0].([]*CustomDataIdentifier)[vs[1].(int)]
	}).(CustomDataIdentifierOutput)
}

type CustomDataIdentifierMapOutput struct{ *pulumi.OutputState }

func (CustomDataIdentifierMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomDataIdentifier)(nil)).Elem()
}

func (o CustomDataIdentifierMapOutput) ToCustomDataIdentifierMapOutput() CustomDataIdentifierMapOutput {
	return o
}

func (o CustomDataIdentifierMapOutput) ToCustomDataIdentifierMapOutputWithContext(ctx context.Context) CustomDataIdentifierMapOutput {
	return o
}

func (o CustomDataIdentifierMapOutput) MapIndex(k pulumi.StringInput) CustomDataIdentifierOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomDataIdentifier {
		return vs[0].(map[string]*CustomDataIdentifier)[vs[1].(string)]
	}).(CustomDataIdentifierOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomDataIdentifierInput)(nil)).Elem(), &CustomDataIdentifier{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomDataIdentifierArrayInput)(nil)).Elem(), CustomDataIdentifierArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomDataIdentifierMapInput)(nil)).Elem(), CustomDataIdentifierMap{})
	pulumi.RegisterOutputType(CustomDataIdentifierOutput{})
	pulumi.RegisterOutputType(CustomDataIdentifierArrayOutput{})
	pulumi.RegisterOutputType(CustomDataIdentifierMapOutput{})
}
