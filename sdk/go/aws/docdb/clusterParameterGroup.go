// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package docdb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a DocumentDB Cluster Parameter Group
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/docdb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := docdb.NewClusterParameterGroup(ctx, "example", &docdb.ClusterParameterGroupArgs{
//				Description: pulumi.String("docdb cluster parameter group"),
//				Family:      pulumi.String("docdb3.6"),
//				Parameters: docdb.ClusterParameterGroupParameterArray{
//					&docdb.ClusterParameterGroupParameterArgs{
//						Name:  pulumi.String("tls"),
//						Value: pulumi.String("enabled"),
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
// terraform import {
//
//	to = aws_docdb_cluster_parameter_group.cluster_pg
//
//	id = "production-pg-1" } Using `pulumi import`, import DocumentDB Cluster Parameter Groups using the `name`. For exampleconsole % pulumi import aws_docdb_cluster_parameter_group.cluster_pg production-pg-1
type ClusterParameterGroup struct {
	pulumi.CustomResourceState

	// The ARN of the DocumentDB cluster parameter group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the DocumentDB cluster parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The family of the DocumentDB cluster parameter group.
	Family pulumi.StringOutput `pulumi:"family"`
	// The name of the DocumentDB parameter.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// A list of DocumentDB parameters to apply. Setting parameters to system default values may show a difference on imported resources.
	Parameters ClusterParameterGroupParameterArrayOutput `pulumi:"parameters"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewClusterParameterGroup registers a new resource with the given unique name, arguments, and options.
func NewClusterParameterGroup(ctx *pulumi.Context,
	name string, args *ClusterParameterGroupArgs, opts ...pulumi.ResourceOption) (*ClusterParameterGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Family == nil {
		return nil, errors.New("invalid value for required argument 'Family'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClusterParameterGroup
	err := ctx.RegisterResource("aws:docdb/clusterParameterGroup:ClusterParameterGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterParameterGroup gets an existing ClusterParameterGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterParameterGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterParameterGroupState, opts ...pulumi.ResourceOption) (*ClusterParameterGroup, error) {
	var resource ClusterParameterGroup
	err := ctx.ReadResource("aws:docdb/clusterParameterGroup:ClusterParameterGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterParameterGroup resources.
type clusterParameterGroupState struct {
	// The ARN of the DocumentDB cluster parameter group.
	Arn *string `pulumi:"arn"`
	// The description of the DocumentDB cluster parameter group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// The family of the DocumentDB cluster parameter group.
	Family *string `pulumi:"family"`
	// The name of the DocumentDB parameter.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// A list of DocumentDB parameters to apply. Setting parameters to system default values may show a difference on imported resources.
	Parameters []ClusterParameterGroupParameter `pulumi:"parameters"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type ClusterParameterGroupState struct {
	// The ARN of the DocumentDB cluster parameter group.
	Arn pulumi.StringPtrInput
	// The description of the DocumentDB cluster parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The family of the DocumentDB cluster parameter group.
	Family pulumi.StringPtrInput
	// The name of the DocumentDB parameter.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// A list of DocumentDB parameters to apply. Setting parameters to system default values may show a difference on imported resources.
	Parameters ClusterParameterGroupParameterArrayInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (ClusterParameterGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterParameterGroupState)(nil)).Elem()
}

type clusterParameterGroupArgs struct {
	// The description of the DocumentDB cluster parameter group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// The family of the DocumentDB cluster parameter group.
	Family string `pulumi:"family"`
	// The name of the DocumentDB parameter.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// A list of DocumentDB parameters to apply. Setting parameters to system default values may show a difference on imported resources.
	Parameters []ClusterParameterGroupParameter `pulumi:"parameters"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ClusterParameterGroup resource.
type ClusterParameterGroupArgs struct {
	// The description of the DocumentDB cluster parameter group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// The family of the DocumentDB cluster parameter group.
	Family pulumi.StringInput
	// The name of the DocumentDB parameter.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// A list of DocumentDB parameters to apply. Setting parameters to system default values may show a difference on imported resources.
	Parameters ClusterParameterGroupParameterArrayInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (ClusterParameterGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterParameterGroupArgs)(nil)).Elem()
}

type ClusterParameterGroupInput interface {
	pulumi.Input

	ToClusterParameterGroupOutput() ClusterParameterGroupOutput
	ToClusterParameterGroupOutputWithContext(ctx context.Context) ClusterParameterGroupOutput
}

func (*ClusterParameterGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterParameterGroup)(nil)).Elem()
}

func (i *ClusterParameterGroup) ToClusterParameterGroupOutput() ClusterParameterGroupOutput {
	return i.ToClusterParameterGroupOutputWithContext(context.Background())
}

func (i *ClusterParameterGroup) ToClusterParameterGroupOutputWithContext(ctx context.Context) ClusterParameterGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterParameterGroupOutput)
}

// ClusterParameterGroupArrayInput is an input type that accepts ClusterParameterGroupArray and ClusterParameterGroupArrayOutput values.
// You can construct a concrete instance of `ClusterParameterGroupArrayInput` via:
//
//	ClusterParameterGroupArray{ ClusterParameterGroupArgs{...} }
type ClusterParameterGroupArrayInput interface {
	pulumi.Input

	ToClusterParameterGroupArrayOutput() ClusterParameterGroupArrayOutput
	ToClusterParameterGroupArrayOutputWithContext(context.Context) ClusterParameterGroupArrayOutput
}

type ClusterParameterGroupArray []ClusterParameterGroupInput

func (ClusterParameterGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterParameterGroup)(nil)).Elem()
}

func (i ClusterParameterGroupArray) ToClusterParameterGroupArrayOutput() ClusterParameterGroupArrayOutput {
	return i.ToClusterParameterGroupArrayOutputWithContext(context.Background())
}

func (i ClusterParameterGroupArray) ToClusterParameterGroupArrayOutputWithContext(ctx context.Context) ClusterParameterGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterParameterGroupArrayOutput)
}

// ClusterParameterGroupMapInput is an input type that accepts ClusterParameterGroupMap and ClusterParameterGroupMapOutput values.
// You can construct a concrete instance of `ClusterParameterGroupMapInput` via:
//
//	ClusterParameterGroupMap{ "key": ClusterParameterGroupArgs{...} }
type ClusterParameterGroupMapInput interface {
	pulumi.Input

	ToClusterParameterGroupMapOutput() ClusterParameterGroupMapOutput
	ToClusterParameterGroupMapOutputWithContext(context.Context) ClusterParameterGroupMapOutput
}

type ClusterParameterGroupMap map[string]ClusterParameterGroupInput

func (ClusterParameterGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterParameterGroup)(nil)).Elem()
}

func (i ClusterParameterGroupMap) ToClusterParameterGroupMapOutput() ClusterParameterGroupMapOutput {
	return i.ToClusterParameterGroupMapOutputWithContext(context.Background())
}

func (i ClusterParameterGroupMap) ToClusterParameterGroupMapOutputWithContext(ctx context.Context) ClusterParameterGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterParameterGroupMapOutput)
}

type ClusterParameterGroupOutput struct{ *pulumi.OutputState }

func (ClusterParameterGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterParameterGroup)(nil)).Elem()
}

func (o ClusterParameterGroupOutput) ToClusterParameterGroupOutput() ClusterParameterGroupOutput {
	return o
}

func (o ClusterParameterGroupOutput) ToClusterParameterGroupOutputWithContext(ctx context.Context) ClusterParameterGroupOutput {
	return o
}

// The ARN of the DocumentDB cluster parameter group.
func (o ClusterParameterGroupOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterParameterGroup) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The description of the DocumentDB cluster parameter group. Defaults to "Managed by Pulumi".
func (o ClusterParameterGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterParameterGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The family of the DocumentDB cluster parameter group.
func (o ClusterParameterGroupOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterParameterGroup) pulumi.StringOutput { return v.Family }).(pulumi.StringOutput)
}

// The name of the DocumentDB parameter.
func (o ClusterParameterGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterParameterGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
func (o ClusterParameterGroupOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterParameterGroup) pulumi.StringOutput { return v.NamePrefix }).(pulumi.StringOutput)
}

// A list of DocumentDB parameters to apply. Setting parameters to system default values may show a difference on imported resources.
func (o ClusterParameterGroupOutput) Parameters() ClusterParameterGroupParameterArrayOutput {
	return o.ApplyT(func(v *ClusterParameterGroup) ClusterParameterGroupParameterArrayOutput { return v.Parameters }).(ClusterParameterGroupParameterArrayOutput)
}

// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ClusterParameterGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ClusterParameterGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o ClusterParameterGroupOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ClusterParameterGroup) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type ClusterParameterGroupArrayOutput struct{ *pulumi.OutputState }

func (ClusterParameterGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterParameterGroup)(nil)).Elem()
}

func (o ClusterParameterGroupArrayOutput) ToClusterParameterGroupArrayOutput() ClusterParameterGroupArrayOutput {
	return o
}

func (o ClusterParameterGroupArrayOutput) ToClusterParameterGroupArrayOutputWithContext(ctx context.Context) ClusterParameterGroupArrayOutput {
	return o
}

func (o ClusterParameterGroupArrayOutput) Index(i pulumi.IntInput) ClusterParameterGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterParameterGroup {
		return vs[0].([]*ClusterParameterGroup)[vs[1].(int)]
	}).(ClusterParameterGroupOutput)
}

type ClusterParameterGroupMapOutput struct{ *pulumi.OutputState }

func (ClusterParameterGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterParameterGroup)(nil)).Elem()
}

func (o ClusterParameterGroupMapOutput) ToClusterParameterGroupMapOutput() ClusterParameterGroupMapOutput {
	return o
}

func (o ClusterParameterGroupMapOutput) ToClusterParameterGroupMapOutputWithContext(ctx context.Context) ClusterParameterGroupMapOutput {
	return o
}

func (o ClusterParameterGroupMapOutput) MapIndex(k pulumi.StringInput) ClusterParameterGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterParameterGroup {
		return vs[0].(map[string]*ClusterParameterGroup)[vs[1].(string)]
	}).(ClusterParameterGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterParameterGroupInput)(nil)).Elem(), &ClusterParameterGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterParameterGroupArrayInput)(nil)).Elem(), ClusterParameterGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterParameterGroupMapInput)(nil)).Elem(), ClusterParameterGroupMap{})
	pulumi.RegisterOutputType(ClusterParameterGroupOutput{})
	pulumi.RegisterOutputType(ClusterParameterGroupArrayOutput{})
	pulumi.RegisterOutputType(ClusterParameterGroupMapOutput{})
}
