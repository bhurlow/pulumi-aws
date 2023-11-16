// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resourceexplorer

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage a Resource Explorer index in the current AWS Region.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/resourceexplorer"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := resourceexplorer.NewIndex(ctx, "example", &resourceexplorer.IndexArgs{
//				Type: pulumi.String("LOCAL"),
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
// Using `pulumi import`, import Resource Explorer indexes using the `arn`. For example:
//
// ```sh
//
//	$ pulumi import aws:resourceexplorer/index:Index example arn:aws:resource-explorer-2:us-east-1:123456789012:index/6047ac4e-207e-4487-9bcf-cb53bb0ff5cc
//
// ```
type Index struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the Resource Explorer index.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll  pulumi.StringMapOutput `pulumi:"tagsAll"`
	Timeouts IndexTimeoutsPtrOutput `pulumi:"timeouts"`
	// The type of the index. Valid values: `AGGREGATOR`, `LOCAL`. To understand the difference between `LOCAL` and `AGGREGATOR`, see the [_AWS Resource Explorer User Guide_](https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-aggregator-region.html).
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIndex registers a new resource with the given unique name, arguments, and options.
func NewIndex(ctx *pulumi.Context,
	name string, args *IndexArgs, opts ...pulumi.ResourceOption) (*Index, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Index
	err := ctx.RegisterResource("aws:resourceexplorer/index:Index", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIndex gets an existing Index resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIndex(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IndexState, opts ...pulumi.ResourceOption) (*Index, error) {
	var resource Index
	err := ctx.ReadResource("aws:resourceexplorer/index:Index", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Index resources.
type indexState struct {
	// Amazon Resource Name (ARN) of the Resource Explorer index.
	Arn *string `pulumi:"arn"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll  map[string]string `pulumi:"tagsAll"`
	Timeouts *IndexTimeouts    `pulumi:"timeouts"`
	// The type of the index. Valid values: `AGGREGATOR`, `LOCAL`. To understand the difference between `LOCAL` and `AGGREGATOR`, see the [_AWS Resource Explorer User Guide_](https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-aggregator-region.html).
	Type *string `pulumi:"type"`
}

type IndexState struct {
	// Amazon Resource Name (ARN) of the Resource Explorer index.
	Arn pulumi.StringPtrInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll  pulumi.StringMapInput
	Timeouts IndexTimeoutsPtrInput
	// The type of the index. Valid values: `AGGREGATOR`, `LOCAL`. To understand the difference between `LOCAL` and `AGGREGATOR`, see the [_AWS Resource Explorer User Guide_](https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-aggregator-region.html).
	Type pulumi.StringPtrInput
}

func (IndexState) ElementType() reflect.Type {
	return reflect.TypeOf((*indexState)(nil)).Elem()
}

type indexArgs struct {
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags     map[string]string `pulumi:"tags"`
	Timeouts *IndexTimeouts    `pulumi:"timeouts"`
	// The type of the index. Valid values: `AGGREGATOR`, `LOCAL`. To understand the difference between `LOCAL` and `AGGREGATOR`, see the [_AWS Resource Explorer User Guide_](https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-aggregator-region.html).
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Index resource.
type IndexArgs struct {
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags     pulumi.StringMapInput
	Timeouts IndexTimeoutsPtrInput
	// The type of the index. Valid values: `AGGREGATOR`, `LOCAL`. To understand the difference between `LOCAL` and `AGGREGATOR`, see the [_AWS Resource Explorer User Guide_](https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-aggregator-region.html).
	Type pulumi.StringInput
}

func (IndexArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*indexArgs)(nil)).Elem()
}

type IndexInput interface {
	pulumi.Input

	ToIndexOutput() IndexOutput
	ToIndexOutputWithContext(ctx context.Context) IndexOutput
}

func (*Index) ElementType() reflect.Type {
	return reflect.TypeOf((**Index)(nil)).Elem()
}

func (i *Index) ToIndexOutput() IndexOutput {
	return i.ToIndexOutputWithContext(context.Background())
}

func (i *Index) ToIndexOutputWithContext(ctx context.Context) IndexOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexOutput)
}

// IndexArrayInput is an input type that accepts IndexArray and IndexArrayOutput values.
// You can construct a concrete instance of `IndexArrayInput` via:
//
//	IndexArray{ IndexArgs{...} }
type IndexArrayInput interface {
	pulumi.Input

	ToIndexArrayOutput() IndexArrayOutput
	ToIndexArrayOutputWithContext(context.Context) IndexArrayOutput
}

type IndexArray []IndexInput

func (IndexArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Index)(nil)).Elem()
}

func (i IndexArray) ToIndexArrayOutput() IndexArrayOutput {
	return i.ToIndexArrayOutputWithContext(context.Background())
}

func (i IndexArray) ToIndexArrayOutputWithContext(ctx context.Context) IndexArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexArrayOutput)
}

// IndexMapInput is an input type that accepts IndexMap and IndexMapOutput values.
// You can construct a concrete instance of `IndexMapInput` via:
//
//	IndexMap{ "key": IndexArgs{...} }
type IndexMapInput interface {
	pulumi.Input

	ToIndexMapOutput() IndexMapOutput
	ToIndexMapOutputWithContext(context.Context) IndexMapOutput
}

type IndexMap map[string]IndexInput

func (IndexMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Index)(nil)).Elem()
}

func (i IndexMap) ToIndexMapOutput() IndexMapOutput {
	return i.ToIndexMapOutputWithContext(context.Background())
}

func (i IndexMap) ToIndexMapOutputWithContext(ctx context.Context) IndexMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexMapOutput)
}

type IndexOutput struct{ *pulumi.OutputState }

func (IndexOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Index)(nil)).Elem()
}

func (o IndexOutput) ToIndexOutput() IndexOutput {
	return o
}

func (o IndexOutput) ToIndexOutputWithContext(ctx context.Context) IndexOutput {
	return o
}

// Amazon Resource Name (ARN) of the Resource Explorer index.
func (o IndexOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Index) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o IndexOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Index) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o IndexOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Index) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

func (o IndexOutput) Timeouts() IndexTimeoutsPtrOutput {
	return o.ApplyT(func(v *Index) IndexTimeoutsPtrOutput { return v.Timeouts }).(IndexTimeoutsPtrOutput)
}

// The type of the index. Valid values: `AGGREGATOR`, `LOCAL`. To understand the difference between `LOCAL` and `AGGREGATOR`, see the [_AWS Resource Explorer User Guide_](https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-aggregator-region.html).
func (o IndexOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Index) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type IndexArrayOutput struct{ *pulumi.OutputState }

func (IndexArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Index)(nil)).Elem()
}

func (o IndexArrayOutput) ToIndexArrayOutput() IndexArrayOutput {
	return o
}

func (o IndexArrayOutput) ToIndexArrayOutputWithContext(ctx context.Context) IndexArrayOutput {
	return o
}

func (o IndexArrayOutput) Index(i pulumi.IntInput) IndexOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Index {
		return vs[0].([]*Index)[vs[1].(int)]
	}).(IndexOutput)
}

type IndexMapOutput struct{ *pulumi.OutputState }

func (IndexMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Index)(nil)).Elem()
}

func (o IndexMapOutput) ToIndexMapOutput() IndexMapOutput {
	return o
}

func (o IndexMapOutput) ToIndexMapOutputWithContext(ctx context.Context) IndexMapOutput {
	return o
}

func (o IndexMapOutput) MapIndex(k pulumi.StringInput) IndexOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Index {
		return vs[0].(map[string]*Index)[vs[1].(string)]
	}).(IndexOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IndexInput)(nil)).Elem(), &Index{})
	pulumi.RegisterInputType(reflect.TypeOf((*IndexArrayInput)(nil)).Elem(), IndexArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IndexMapInput)(nil)).Elem(), IndexMap{})
	pulumi.RegisterOutputType(IndexOutput{})
	pulumi.RegisterOutputType(IndexArrayOutput{})
	pulumi.RegisterOutputType(IndexMapOutput{})
}
