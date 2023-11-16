// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resourceexplorer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage a Resource Explorer view.
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
//			exampleIndex, err := resourceexplorer.NewIndex(ctx, "exampleIndex", &resourceexplorer.IndexArgs{
//				Type: pulumi.String("LOCAL"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = resourceexplorer.NewView(ctx, "exampleView", &resourceexplorer.ViewArgs{
//				Filters: &resourceexplorer.ViewFiltersArgs{
//					FilterString: pulumi.String("resourcetype:ec2:instance"),
//				},
//				IncludedProperties: resourceexplorer.ViewIncludedPropertyArray{
//					&resourceexplorer.ViewIncludedPropertyArgs{
//						Name: pulumi.String("tags"),
//					},
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				exampleIndex,
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
// Using `pulumi import`, import Resource Explorer views using the `arn`. For example:
//
// ```sh
//
//	$ pulumi import aws:resourceexplorer/view:View example arn:aws:resource-explorer-2:us-west-2:123456789012:view/exampleview/e0914f6c-6c27-4b47-b5d4-6b28381a2421
//
// ```
type View struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the Resource Explorer view.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specifies whether the view is the [_default view_](https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-views-about.html#manage-views-about-default) for the AWS Region. Default: `false`.
	DefaultView pulumi.BoolOutput `pulumi:"defaultView"`
	// Specifies which resources are included in the results of queries made using this view. See Filters below for more details.
	Filters ViewFiltersPtrOutput `pulumi:"filters"`
	// Optional fields to be included in search results from this view. See Included Properties below for more details.
	IncludedProperties ViewIncludedPropertyArrayOutput `pulumi:"includedProperties"`
	// The name of the view. The name must be no more than 64 characters long, and can include letters, digits, and the dash (-) character. The name must be unique within its AWS Region.
	Name pulumi.StringOutput `pulumi:"name"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewView registers a new resource with the given unique name, arguments, and options.
func NewView(ctx *pulumi.Context,
	name string, args *ViewArgs, opts ...pulumi.ResourceOption) (*View, error) {
	if args == nil {
		args = &ViewArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource View
	err := ctx.RegisterResource("aws:resourceexplorer/view:View", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetView gets an existing View resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetView(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ViewState, opts ...pulumi.ResourceOption) (*View, error) {
	var resource View
	err := ctx.ReadResource("aws:resourceexplorer/view:View", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering View resources.
type viewState struct {
	// Amazon Resource Name (ARN) of the Resource Explorer view.
	Arn *string `pulumi:"arn"`
	// Specifies whether the view is the [_default view_](https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-views-about.html#manage-views-about-default) for the AWS Region. Default: `false`.
	DefaultView *bool `pulumi:"defaultView"`
	// Specifies which resources are included in the results of queries made using this view. See Filters below for more details.
	Filters *ViewFilters `pulumi:"filters"`
	// Optional fields to be included in search results from this view. See Included Properties below for more details.
	IncludedProperties []ViewIncludedProperty `pulumi:"includedProperties"`
	// The name of the view. The name must be no more than 64 characters long, and can include letters, digits, and the dash (-) character. The name must be unique within its AWS Region.
	Name *string `pulumi:"name"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type ViewState struct {
	// Amazon Resource Name (ARN) of the Resource Explorer view.
	Arn pulumi.StringPtrInput
	// Specifies whether the view is the [_default view_](https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-views-about.html#manage-views-about-default) for the AWS Region. Default: `false`.
	DefaultView pulumi.BoolPtrInput
	// Specifies which resources are included in the results of queries made using this view. See Filters below for more details.
	Filters ViewFiltersPtrInput
	// Optional fields to be included in search results from this view. See Included Properties below for more details.
	IncludedProperties ViewIncludedPropertyArrayInput
	// The name of the view. The name must be no more than 64 characters long, and can include letters, digits, and the dash (-) character. The name must be unique within its AWS Region.
	Name pulumi.StringPtrInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (ViewState) ElementType() reflect.Type {
	return reflect.TypeOf((*viewState)(nil)).Elem()
}

type viewArgs struct {
	// Specifies whether the view is the [_default view_](https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-views-about.html#manage-views-about-default) for the AWS Region. Default: `false`.
	DefaultView *bool `pulumi:"defaultView"`
	// Specifies which resources are included in the results of queries made using this view. See Filters below for more details.
	Filters *ViewFilters `pulumi:"filters"`
	// Optional fields to be included in search results from this view. See Included Properties below for more details.
	IncludedProperties []ViewIncludedProperty `pulumi:"includedProperties"`
	// The name of the view. The name must be no more than 64 characters long, and can include letters, digits, and the dash (-) character. The name must be unique within its AWS Region.
	Name *string `pulumi:"name"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a View resource.
type ViewArgs struct {
	// Specifies whether the view is the [_default view_](https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-views-about.html#manage-views-about-default) for the AWS Region. Default: `false`.
	DefaultView pulumi.BoolPtrInput
	// Specifies which resources are included in the results of queries made using this view. See Filters below for more details.
	Filters ViewFiltersPtrInput
	// Optional fields to be included in search results from this view. See Included Properties below for more details.
	IncludedProperties ViewIncludedPropertyArrayInput
	// The name of the view. The name must be no more than 64 characters long, and can include letters, digits, and the dash (-) character. The name must be unique within its AWS Region.
	Name pulumi.StringPtrInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (ViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*viewArgs)(nil)).Elem()
}

type ViewInput interface {
	pulumi.Input

	ToViewOutput() ViewOutput
	ToViewOutputWithContext(ctx context.Context) ViewOutput
}

func (*View) ElementType() reflect.Type {
	return reflect.TypeOf((**View)(nil)).Elem()
}

func (i *View) ToViewOutput() ViewOutput {
	return i.ToViewOutputWithContext(context.Background())
}

func (i *View) ToViewOutputWithContext(ctx context.Context) ViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViewOutput)
}

// ViewArrayInput is an input type that accepts ViewArray and ViewArrayOutput values.
// You can construct a concrete instance of `ViewArrayInput` via:
//
//	ViewArray{ ViewArgs{...} }
type ViewArrayInput interface {
	pulumi.Input

	ToViewArrayOutput() ViewArrayOutput
	ToViewArrayOutputWithContext(context.Context) ViewArrayOutput
}

type ViewArray []ViewInput

func (ViewArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*View)(nil)).Elem()
}

func (i ViewArray) ToViewArrayOutput() ViewArrayOutput {
	return i.ToViewArrayOutputWithContext(context.Background())
}

func (i ViewArray) ToViewArrayOutputWithContext(ctx context.Context) ViewArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViewArrayOutput)
}

// ViewMapInput is an input type that accepts ViewMap and ViewMapOutput values.
// You can construct a concrete instance of `ViewMapInput` via:
//
//	ViewMap{ "key": ViewArgs{...} }
type ViewMapInput interface {
	pulumi.Input

	ToViewMapOutput() ViewMapOutput
	ToViewMapOutputWithContext(context.Context) ViewMapOutput
}

type ViewMap map[string]ViewInput

func (ViewMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*View)(nil)).Elem()
}

func (i ViewMap) ToViewMapOutput() ViewMapOutput {
	return i.ToViewMapOutputWithContext(context.Background())
}

func (i ViewMap) ToViewMapOutputWithContext(ctx context.Context) ViewMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViewMapOutput)
}

type ViewOutput struct{ *pulumi.OutputState }

func (ViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**View)(nil)).Elem()
}

func (o ViewOutput) ToViewOutput() ViewOutput {
	return o
}

func (o ViewOutput) ToViewOutputWithContext(ctx context.Context) ViewOutput {
	return o
}

// Amazon Resource Name (ARN) of the Resource Explorer view.
func (o ViewOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Specifies whether the view is the [_default view_](https://docs.aws.amazon.com/resource-explorer/latest/userguide/manage-views-about.html#manage-views-about-default) for the AWS Region. Default: `false`.
func (o ViewOutput) DefaultView() pulumi.BoolOutput {
	return o.ApplyT(func(v *View) pulumi.BoolOutput { return v.DefaultView }).(pulumi.BoolOutput)
}

// Specifies which resources are included in the results of queries made using this view. See Filters below for more details.
func (o ViewOutput) Filters() ViewFiltersPtrOutput {
	return o.ApplyT(func(v *View) ViewFiltersPtrOutput { return v.Filters }).(ViewFiltersPtrOutput)
}

// Optional fields to be included in search results from this view. See Included Properties below for more details.
func (o ViewOutput) IncludedProperties() ViewIncludedPropertyArrayOutput {
	return o.ApplyT(func(v *View) ViewIncludedPropertyArrayOutput { return v.IncludedProperties }).(ViewIncludedPropertyArrayOutput)
}

// The name of the view. The name must be no more than 64 characters long, and can include letters, digits, and the dash (-) character. The name must be unique within its AWS Region.
func (o ViewOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *View) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ViewOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *View) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o ViewOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *View) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type ViewArrayOutput struct{ *pulumi.OutputState }

func (ViewArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*View)(nil)).Elem()
}

func (o ViewArrayOutput) ToViewArrayOutput() ViewArrayOutput {
	return o
}

func (o ViewArrayOutput) ToViewArrayOutputWithContext(ctx context.Context) ViewArrayOutput {
	return o
}

func (o ViewArrayOutput) Index(i pulumi.IntInput) ViewOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *View {
		return vs[0].([]*View)[vs[1].(int)]
	}).(ViewOutput)
}

type ViewMapOutput struct{ *pulumi.OutputState }

func (ViewMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*View)(nil)).Elem()
}

func (o ViewMapOutput) ToViewMapOutput() ViewMapOutput {
	return o
}

func (o ViewMapOutput) ToViewMapOutputWithContext(ctx context.Context) ViewMapOutput {
	return o
}

func (o ViewMapOutput) MapIndex(k pulumi.StringInput) ViewOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *View {
		return vs[0].(map[string]*View)[vs[1].(string)]
	}).(ViewOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ViewInput)(nil)).Elem(), &View{})
	pulumi.RegisterInputType(reflect.TypeOf((*ViewArrayInput)(nil)).Elem(), ViewArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ViewMapInput)(nil)).Elem(), ViewMap{})
	pulumi.RegisterOutputType(ViewOutput{})
	pulumi.RegisterOutputType(ViewArrayOutput{})
	pulumi.RegisterOutputType(ViewMapOutput{})
}
