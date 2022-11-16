// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Amazon API Gateway Version 2 VPC Link.
//
// > **Note:** Amazon API Gateway Version 2 VPC Links enable private integrations that connect HTTP APIs to private resources in a VPC.
// To enable private integration for REST APIs, use the Amazon API Gateway Version 1 VPC Link resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/apigatewayv2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := apigatewayv2.NewVpcLink(ctx, "example", &apigatewayv2.VpcLinkArgs{
//				SecurityGroupIds: pulumi.StringArray{
//					pulumi.Any(data.Aws_security_group.Example.Id),
//				},
//				SubnetIds: pulumi.Any(data.Aws_subnet_ids.Example.Ids),
//				Tags: pulumi.StringMap{
//					"Usage": pulumi.String("example"),
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
// `aws_apigatewayv2_vpc_link` can be imported by using the VPC Link identifier, e.g.,
//
// ```sh
//
//	$ pulumi import aws:apigatewayv2/vpcLink:VpcLink example aabbccddee
//
// ```
type VpcLink struct {
	pulumi.CustomResourceState

	// VPC Link ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Name of the VPC Link. Must be between 1 and 128 characters in length.
	Name pulumi.StringOutput `pulumi:"name"`
	// Security group IDs for the VPC Link.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// Subnet IDs for the VPC Link.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// Map of tags to assign to the VPC Link. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewVpcLink registers a new resource with the given unique name, arguments, and options.
func NewVpcLink(ctx *pulumi.Context,
	name string, args *VpcLinkArgs, opts ...pulumi.ResourceOption) (*VpcLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SecurityGroupIds == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupIds'")
	}
	if args.SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIds'")
	}
	var resource VpcLink
	err := ctx.RegisterResource("aws:apigatewayv2/vpcLink:VpcLink", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:apigatewayv2/vpcLink:VpcLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcLink resources.
type vpcLinkState struct {
	// VPC Link ARN.
	Arn *string `pulumi:"arn"`
	// Name of the VPC Link. Must be between 1 and 128 characters in length.
	Name *string `pulumi:"name"`
	// Security group IDs for the VPC Link.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Subnet IDs for the VPC Link.
	SubnetIds []string `pulumi:"subnetIds"`
	// Map of tags to assign to the VPC Link. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type VpcLinkState struct {
	// VPC Link ARN.
	Arn pulumi.StringPtrInput
	// Name of the VPC Link. Must be between 1 and 128 characters in length.
	Name pulumi.StringPtrInput
	// Security group IDs for the VPC Link.
	SecurityGroupIds pulumi.StringArrayInput
	// Subnet IDs for the VPC Link.
	SubnetIds pulumi.StringArrayInput
	// Map of tags to assign to the VPC Link. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (VpcLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcLinkState)(nil)).Elem()
}

type vpcLinkArgs struct {
	// Name of the VPC Link. Must be between 1 and 128 characters in length.
	Name *string `pulumi:"name"`
	// Security group IDs for the VPC Link.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Subnet IDs for the VPC Link.
	SubnetIds []string `pulumi:"subnetIds"`
	// Map of tags to assign to the VPC Link. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a VpcLink resource.
type VpcLinkArgs struct {
	// Name of the VPC Link. Must be between 1 and 128 characters in length.
	Name pulumi.StringPtrInput
	// Security group IDs for the VPC Link.
	SecurityGroupIds pulumi.StringArrayInput
	// Subnet IDs for the VPC Link.
	SubnetIds pulumi.StringArrayInput
	// Map of tags to assign to the VPC Link. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
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

// VPC Link ARN.
func (o VpcLinkOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcLink) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Name of the VPC Link. Must be between 1 and 128 characters in length.
func (o VpcLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Security group IDs for the VPC Link.
func (o VpcLinkOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcLink) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// Subnet IDs for the VPC Link.
func (o VpcLinkOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcLink) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// Map of tags to assign to the VPC Link. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o VpcLinkOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcLink) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o VpcLinkOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcLink) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
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
