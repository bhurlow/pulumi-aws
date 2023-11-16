// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an API Gateway VPC Link.
//
// > **Note:** Amazon API Gateway Version 1 VPC Links enable private integrations that connect REST APIs to private resources in a VPC.
// To enable private integration for HTTP APIs, use the Amazon API Gateway Version 2 VPC Link resource.
//
// ## Import
//
// Using `pulumi import`, import API Gateway VPC Link using the `id`. For example:
//
// ```sh
//
//	$ pulumi import aws:apigateway/vpcLink:VpcLink example 12345abcde
//
// ```
type VpcLink struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of the VPC link.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name used to label and identify the VPC link.
	Name pulumi.StringOutput `pulumi:"name"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// List of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
	TargetArn pulumi.StringOutput `pulumi:"targetArn"`
}

// NewVpcLink registers a new resource with the given unique name, arguments, and options.
func NewVpcLink(ctx *pulumi.Context,
	name string, args *VpcLinkArgs, opts ...pulumi.ResourceOption) (*VpcLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TargetArn == nil {
		return nil, errors.New("invalid value for required argument 'TargetArn'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcLink
	err := ctx.RegisterResource("aws:apigateway/vpcLink:VpcLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcLink gets an existing VpcLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcLinkState, opts ...pulumi.ResourceOption) (*VpcLink, error) {
	var resource VpcLink
	err := ctx.ReadResource("aws:apigateway/vpcLink:VpcLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcLink resources.
type vpcLinkState struct {
	Arn *string `pulumi:"arn"`
	// Description of the VPC link.
	Description *string `pulumi:"description"`
	// Name used to label and identify the VPC link.
	Name *string `pulumi:"name"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// List of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
	TargetArn *string `pulumi:"targetArn"`
}

type VpcLinkState struct {
	Arn pulumi.StringPtrInput
	// Description of the VPC link.
	Description pulumi.StringPtrInput
	// Name used to label and identify the VPC link.
	Name pulumi.StringPtrInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// List of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
	TargetArn pulumi.StringPtrInput
}

func (VpcLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcLinkState)(nil)).Elem()
}

type vpcLinkArgs struct {
	// Description of the VPC link.
	Description *string `pulumi:"description"`
	// Name used to label and identify the VPC link.
	Name *string `pulumi:"name"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// List of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
	TargetArn string `pulumi:"targetArn"`
}

// The set of arguments for constructing a VpcLink resource.
type VpcLinkArgs struct {
	// Description of the VPC link.
	Description pulumi.StringPtrInput
	// Name used to label and identify the VPC link.
	Name pulumi.StringPtrInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// List of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
	TargetArn pulumi.StringInput
}

func (VpcLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcLinkArgs)(nil)).Elem()
}

type VpcLinkInput interface {
	pulumi.Input

	ToVpcLinkOutput() VpcLinkOutput
	ToVpcLinkOutputWithContext(ctx context.Context) VpcLinkOutput
}

func (*VpcLink) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcLink)(nil)).Elem()
}

func (i *VpcLink) ToVpcLinkOutput() VpcLinkOutput {
	return i.ToVpcLinkOutputWithContext(context.Background())
}

func (i *VpcLink) ToVpcLinkOutputWithContext(ctx context.Context) VpcLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcLinkOutput)
}

// VpcLinkArrayInput is an input type that accepts VpcLinkArray and VpcLinkArrayOutput values.
// You can construct a concrete instance of `VpcLinkArrayInput` via:
//
//	VpcLinkArray{ VpcLinkArgs{...} }
type VpcLinkArrayInput interface {
	pulumi.Input

	ToVpcLinkArrayOutput() VpcLinkArrayOutput
	ToVpcLinkArrayOutputWithContext(context.Context) VpcLinkArrayOutput
}

type VpcLinkArray []VpcLinkInput

func (VpcLinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcLink)(nil)).Elem()
}

func (i VpcLinkArray) ToVpcLinkArrayOutput() VpcLinkArrayOutput {
	return i.ToVpcLinkArrayOutputWithContext(context.Background())
}

func (i VpcLinkArray) ToVpcLinkArrayOutputWithContext(ctx context.Context) VpcLinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcLinkArrayOutput)
}

// VpcLinkMapInput is an input type that accepts VpcLinkMap and VpcLinkMapOutput values.
// You can construct a concrete instance of `VpcLinkMapInput` via:
//
//	VpcLinkMap{ "key": VpcLinkArgs{...} }
type VpcLinkMapInput interface {
	pulumi.Input

	ToVpcLinkMapOutput() VpcLinkMapOutput
	ToVpcLinkMapOutputWithContext(context.Context) VpcLinkMapOutput
}

type VpcLinkMap map[string]VpcLinkInput

func (VpcLinkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcLink)(nil)).Elem()
}

func (i VpcLinkMap) ToVpcLinkMapOutput() VpcLinkMapOutput {
	return i.ToVpcLinkMapOutputWithContext(context.Background())
}

func (i VpcLinkMap) ToVpcLinkMapOutputWithContext(ctx context.Context) VpcLinkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcLinkMapOutput)
}

type VpcLinkOutput struct{ *pulumi.OutputState }

func (VpcLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcLink)(nil)).Elem()
}

func (o VpcLinkOutput) ToVpcLinkOutput() VpcLinkOutput {
	return o
}

func (o VpcLinkOutput) ToVpcLinkOutputWithContext(ctx context.Context) VpcLinkOutput {
	return o
}

func (o VpcLinkOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcLink) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Description of the VPC link.
func (o VpcLinkOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcLink) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Name used to label and identify the VPC link.
func (o VpcLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o VpcLinkOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcLink) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o VpcLinkOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcLink) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// List of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
func (o VpcLinkOutput) TargetArn() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcLink) pulumi.StringOutput { return v.TargetArn }).(pulumi.StringOutput)
}

type VpcLinkArrayOutput struct{ *pulumi.OutputState }

func (VpcLinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcLink)(nil)).Elem()
}

func (o VpcLinkArrayOutput) ToVpcLinkArrayOutput() VpcLinkArrayOutput {
	return o
}

func (o VpcLinkArrayOutput) ToVpcLinkArrayOutputWithContext(ctx context.Context) VpcLinkArrayOutput {
	return o
}

func (o VpcLinkArrayOutput) Index(i pulumi.IntInput) VpcLinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcLink {
		return vs[0].([]*VpcLink)[vs[1].(int)]
	}).(VpcLinkOutput)
}

type VpcLinkMapOutput struct{ *pulumi.OutputState }

func (VpcLinkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcLink)(nil)).Elem()
}

func (o VpcLinkMapOutput) ToVpcLinkMapOutput() VpcLinkMapOutput {
	return o
}

func (o VpcLinkMapOutput) ToVpcLinkMapOutputWithContext(ctx context.Context) VpcLinkMapOutput {
	return o
}

func (o VpcLinkMapOutput) MapIndex(k pulumi.StringInput) VpcLinkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcLink {
		return vs[0].(map[string]*VpcLink)[vs[1].(string)]
	}).(VpcLinkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcLinkInput)(nil)).Elem(), &VpcLink{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcLinkArrayInput)(nil)).Elem(), VpcLinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcLinkMapInput)(nil)).Elem(), VpcLinkMap{})
	pulumi.RegisterOutputType(VpcLinkOutput{})
	pulumi.RegisterOutputType(VpcLinkArrayOutput{})
	pulumi.RegisterOutputType(VpcLinkMapOutput{})
}
