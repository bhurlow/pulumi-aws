// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apprunner

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an App Runner VPC Connector.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apprunner"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := apprunner.NewVpcConnector(ctx, "connector", &apprunner.VpcConnectorArgs{
//				SecurityGroups: pulumi.StringArray{
//					pulumi.String("sg1"),
//					pulumi.String("sg2"),
//				},
//				Subnets: pulumi.StringArray{
//					pulumi.String("subnet1"),
//					pulumi.String("subnet2"),
//				},
//				VpcConnectorName: pulumi.String("name"),
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
//	to = aws_apprunner_vpc_connector.example
//
//	id = "arn:aws:apprunner:us-east-1:1234567890:vpcconnector/example/1/0a03292a89764e5882c41d8f991c82fe" } Using `pulumi import`, import App Runner vpc connector using the `arn`. For exampleconsole % pulumi import aws_apprunner_vpc_connector.example arn:aws:apprunner:us-east-1:1234567890:vpcconnector/example/1/0a03292a89764e5882c41d8f991c82fe
type VpcConnector struct {
	pulumi.CustomResourceState

	// ARN of VPC connector.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// List of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets. If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`
	// Current state of the VPC connector. If the status of a connector revision is INACTIVE, it was deleted and can't be used. Inactive connector revisions are permanently removed some time after they are deleted.
	Status pulumi.StringOutput `pulumi:"status"`
	// List of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC. Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.
	Subnets pulumi.StringArrayOutput `pulumi:"subnets"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Name for the VPC connector.
	VpcConnectorName pulumi.StringOutput `pulumi:"vpcConnectorName"`
	// The revision of VPC connector. It's unique among all the active connectors ("Status": "ACTIVE") that share the same Name.
	VpcConnectorRevision pulumi.IntOutput `pulumi:"vpcConnectorRevision"`
}

// NewVpcConnector registers a new resource with the given unique name, arguments, and options.
func NewVpcConnector(ctx *pulumi.Context,
	name string, args *VpcConnectorArgs, opts ...pulumi.ResourceOption) (*VpcConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SecurityGroups == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroups'")
	}
	if args.Subnets == nil {
		return nil, errors.New("invalid value for required argument 'Subnets'")
	}
	if args.VpcConnectorName == nil {
		return nil, errors.New("invalid value for required argument 'VpcConnectorName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcConnector
	err := ctx.RegisterResource("aws:apprunner/vpcConnector:VpcConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcConnector gets an existing VpcConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcConnectorState, opts ...pulumi.ResourceOption) (*VpcConnector, error) {
	var resource VpcConnector
	err := ctx.ReadResource("aws:apprunner/vpcConnector:VpcConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcConnector resources.
type vpcConnectorState struct {
	// ARN of VPC connector.
	Arn *string `pulumi:"arn"`
	// List of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets. If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.
	SecurityGroups []string `pulumi:"securityGroups"`
	// Current state of the VPC connector. If the status of a connector revision is INACTIVE, it was deleted and can't be used. Inactive connector revisions are permanently removed some time after they are deleted.
	Status *string `pulumi:"status"`
	// List of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC. Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.
	Subnets []string `pulumi:"subnets"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Name for the VPC connector.
	VpcConnectorName *string `pulumi:"vpcConnectorName"`
	// The revision of VPC connector. It's unique among all the active connectors ("Status": "ACTIVE") that share the same Name.
	VpcConnectorRevision *int `pulumi:"vpcConnectorRevision"`
}

type VpcConnectorState struct {
	// ARN of VPC connector.
	Arn pulumi.StringPtrInput
	// List of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets. If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.
	SecurityGroups pulumi.StringArrayInput
	// Current state of the VPC connector. If the status of a connector revision is INACTIVE, it was deleted and can't be used. Inactive connector revisions are permanently removed some time after they are deleted.
	Status pulumi.StringPtrInput
	// List of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC. Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.
	Subnets pulumi.StringArrayInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// Name for the VPC connector.
	VpcConnectorName pulumi.StringPtrInput
	// The revision of VPC connector. It's unique among all the active connectors ("Status": "ACTIVE") that share the same Name.
	VpcConnectorRevision pulumi.IntPtrInput
}

func (VpcConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcConnectorState)(nil)).Elem()
}

type vpcConnectorArgs struct {
	// List of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets. If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.
	SecurityGroups []string `pulumi:"securityGroups"`
	// List of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC. Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.
	Subnets []string `pulumi:"subnets"`
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Name for the VPC connector.
	VpcConnectorName string `pulumi:"vpcConnectorName"`
}

// The set of arguments for constructing a VpcConnector resource.
type VpcConnectorArgs struct {
	// List of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets. If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.
	SecurityGroups pulumi.StringArrayInput
	// List of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC. Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.
	Subnets pulumi.StringArrayInput
	// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Name for the VPC connector.
	VpcConnectorName pulumi.StringInput
}

func (VpcConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcConnectorArgs)(nil)).Elem()
}

type VpcConnectorInput interface {
	pulumi.Input

	ToVpcConnectorOutput() VpcConnectorOutput
	ToVpcConnectorOutputWithContext(ctx context.Context) VpcConnectorOutput
}

func (*VpcConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcConnector)(nil)).Elem()
}

func (i *VpcConnector) ToVpcConnectorOutput() VpcConnectorOutput {
	return i.ToVpcConnectorOutputWithContext(context.Background())
}

func (i *VpcConnector) ToVpcConnectorOutputWithContext(ctx context.Context) VpcConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcConnectorOutput)
}

// VpcConnectorArrayInput is an input type that accepts VpcConnectorArray and VpcConnectorArrayOutput values.
// You can construct a concrete instance of `VpcConnectorArrayInput` via:
//
//	VpcConnectorArray{ VpcConnectorArgs{...} }
type VpcConnectorArrayInput interface {
	pulumi.Input

	ToVpcConnectorArrayOutput() VpcConnectorArrayOutput
	ToVpcConnectorArrayOutputWithContext(context.Context) VpcConnectorArrayOutput
}

type VpcConnectorArray []VpcConnectorInput

func (VpcConnectorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcConnector)(nil)).Elem()
}

func (i VpcConnectorArray) ToVpcConnectorArrayOutput() VpcConnectorArrayOutput {
	return i.ToVpcConnectorArrayOutputWithContext(context.Background())
}

func (i VpcConnectorArray) ToVpcConnectorArrayOutputWithContext(ctx context.Context) VpcConnectorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcConnectorArrayOutput)
}

// VpcConnectorMapInput is an input type that accepts VpcConnectorMap and VpcConnectorMapOutput values.
// You can construct a concrete instance of `VpcConnectorMapInput` via:
//
//	VpcConnectorMap{ "key": VpcConnectorArgs{...} }
type VpcConnectorMapInput interface {
	pulumi.Input

	ToVpcConnectorMapOutput() VpcConnectorMapOutput
	ToVpcConnectorMapOutputWithContext(context.Context) VpcConnectorMapOutput
}

type VpcConnectorMap map[string]VpcConnectorInput

func (VpcConnectorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcConnector)(nil)).Elem()
}

func (i VpcConnectorMap) ToVpcConnectorMapOutput() VpcConnectorMapOutput {
	return i.ToVpcConnectorMapOutputWithContext(context.Background())
}

func (i VpcConnectorMap) ToVpcConnectorMapOutputWithContext(ctx context.Context) VpcConnectorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcConnectorMapOutput)
}

type VpcConnectorOutput struct{ *pulumi.OutputState }

func (VpcConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcConnector)(nil)).Elem()
}

func (o VpcConnectorOutput) ToVpcConnectorOutput() VpcConnectorOutput {
	return o
}

func (o VpcConnectorOutput) ToVpcConnectorOutputWithContext(ctx context.Context) VpcConnectorOutput {
	return o
}

// ARN of VPC connector.
func (o VpcConnectorOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcConnector) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// List of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets. If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.
func (o VpcConnectorOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcConnector) pulumi.StringArrayOutput { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

// Current state of the VPC connector. If the status of a connector revision is INACTIVE, it was deleted and can't be used. Inactive connector revisions are permanently removed some time after they are deleted.
func (o VpcConnectorOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcConnector) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// List of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC. Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.
func (o VpcConnectorOutput) Subnets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcConnector) pulumi.StringArrayOutput { return v.Subnets }).(pulumi.StringArrayOutput)
}

// Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o VpcConnectorOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcConnector) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o VpcConnectorOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcConnector) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Name for the VPC connector.
func (o VpcConnectorOutput) VpcConnectorName() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcConnector) pulumi.StringOutput { return v.VpcConnectorName }).(pulumi.StringOutput)
}

// The revision of VPC connector. It's unique among all the active connectors ("Status": "ACTIVE") that share the same Name.
func (o VpcConnectorOutput) VpcConnectorRevision() pulumi.IntOutput {
	return o.ApplyT(func(v *VpcConnector) pulumi.IntOutput { return v.VpcConnectorRevision }).(pulumi.IntOutput)
}

type VpcConnectorArrayOutput struct{ *pulumi.OutputState }

func (VpcConnectorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcConnector)(nil)).Elem()
}

func (o VpcConnectorArrayOutput) ToVpcConnectorArrayOutput() VpcConnectorArrayOutput {
	return o
}

func (o VpcConnectorArrayOutput) ToVpcConnectorArrayOutputWithContext(ctx context.Context) VpcConnectorArrayOutput {
	return o
}

func (o VpcConnectorArrayOutput) Index(i pulumi.IntInput) VpcConnectorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcConnector {
		return vs[0].([]*VpcConnector)[vs[1].(int)]
	}).(VpcConnectorOutput)
}

type VpcConnectorMapOutput struct{ *pulumi.OutputState }

func (VpcConnectorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcConnector)(nil)).Elem()
}

func (o VpcConnectorMapOutput) ToVpcConnectorMapOutput() VpcConnectorMapOutput {
	return o
}

func (o VpcConnectorMapOutput) ToVpcConnectorMapOutputWithContext(ctx context.Context) VpcConnectorMapOutput {
	return o
}

func (o VpcConnectorMapOutput) MapIndex(k pulumi.StringInput) VpcConnectorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcConnector {
		return vs[0].(map[string]*VpcConnector)[vs[1].(string)]
	}).(VpcConnectorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcConnectorInput)(nil)).Elem(), &VpcConnector{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcConnectorArrayInput)(nil)).Elem(), VpcConnectorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcConnectorMapInput)(nil)).Elem(), VpcConnectorMap{})
	pulumi.RegisterOutputType(VpcConnectorOutput{})
	pulumi.RegisterOutputType(VpcConnectorArrayOutput{})
	pulumi.RegisterOutputType(VpcConnectorMapOutput{})
}
