// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package opensearch

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an [AWS Opensearch VPC Endpoint](https://docs.aws.amazon.com/opensearch-service/latest/APIReference/API_CreateVpcEndpoint.html). Creates an Amazon OpenSearch Service-managed VPC endpoint.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/opensearch"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := opensearch.NewVpcEndpoint(ctx, "foo", &opensearch.VpcEndpointArgs{
//				DomainArn: pulumi.Any(aws_opensearch_domain.Domain_1.Arn),
//				VpcOptions: &opensearch.VpcEndpointVpcOptionsArgs{
//					SecurityGroupIds: pulumi.StringArray{
//						aws_security_group.Test.Id,
//						aws_security_group.Test2.Id,
//					},
//					SubnetIds: pulumi.StringArray{
//						aws_subnet.Test.Id,
//						aws_subnet.Test2.Id,
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
// Using `pulumi import`, import OpenSearch VPC endpoint connections using the `id`. For example:
//
// ```sh
//
//	$ pulumi import aws:opensearch/vpcEndpoint:VpcEndpoint example endpoint-id
//
// ```
type VpcEndpoint struct {
	pulumi.CustomResourceState

	// Specifies the Amazon Resource Name (ARN) of the domain to create the endpoint for
	DomainArn pulumi.StringOutput `pulumi:"domainArn"`
	// The connection endpoint ID for connecting to the domain.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Options to specify the subnets and security groups for the endpoint.
	VpcOptions VpcEndpointVpcOptionsOutput `pulumi:"vpcOptions"`
}

// NewVpcEndpoint registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpoint(ctx *pulumi.Context,
	name string, args *VpcEndpointArgs, opts ...pulumi.ResourceOption) (*VpcEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainArn == nil {
		return nil, errors.New("invalid value for required argument 'DomainArn'")
	}
	if args.VpcOptions == nil {
		return nil, errors.New("invalid value for required argument 'VpcOptions'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcEndpoint
	err := ctx.RegisterResource("aws:opensearch/vpcEndpoint:VpcEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcEndpoint gets an existing VpcEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcEndpointState, opts ...pulumi.ResourceOption) (*VpcEndpoint, error) {
	var resource VpcEndpoint
	err := ctx.ReadResource("aws:opensearch/vpcEndpoint:VpcEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcEndpoint resources.
type vpcEndpointState struct {
	// Specifies the Amazon Resource Name (ARN) of the domain to create the endpoint for
	DomainArn *string `pulumi:"domainArn"`
	// The connection endpoint ID for connecting to the domain.
	Endpoint *string `pulumi:"endpoint"`
	// Options to specify the subnets and security groups for the endpoint.
	VpcOptions *VpcEndpointVpcOptions `pulumi:"vpcOptions"`
}

type VpcEndpointState struct {
	// Specifies the Amazon Resource Name (ARN) of the domain to create the endpoint for
	DomainArn pulumi.StringPtrInput
	// The connection endpoint ID for connecting to the domain.
	Endpoint pulumi.StringPtrInput
	// Options to specify the subnets and security groups for the endpoint.
	VpcOptions VpcEndpointVpcOptionsPtrInput
}

func (VpcEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointState)(nil)).Elem()
}

type vpcEndpointArgs struct {
	// Specifies the Amazon Resource Name (ARN) of the domain to create the endpoint for
	DomainArn string `pulumi:"domainArn"`
	// Options to specify the subnets and security groups for the endpoint.
	VpcOptions VpcEndpointVpcOptions `pulumi:"vpcOptions"`
}

// The set of arguments for constructing a VpcEndpoint resource.
type VpcEndpointArgs struct {
	// Specifies the Amazon Resource Name (ARN) of the domain to create the endpoint for
	DomainArn pulumi.StringInput
	// Options to specify the subnets and security groups for the endpoint.
	VpcOptions VpcEndpointVpcOptionsInput
}

func (VpcEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointArgs)(nil)).Elem()
}

type VpcEndpointInput interface {
	pulumi.Input

	ToVpcEndpointOutput() VpcEndpointOutput
	ToVpcEndpointOutputWithContext(ctx context.Context) VpcEndpointOutput
}

func (*VpcEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpoint)(nil)).Elem()
}

func (i *VpcEndpoint) ToVpcEndpointOutput() VpcEndpointOutput {
	return i.ToVpcEndpointOutputWithContext(context.Background())
}

func (i *VpcEndpoint) ToVpcEndpointOutputWithContext(ctx context.Context) VpcEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointOutput)
}

// VpcEndpointArrayInput is an input type that accepts VpcEndpointArray and VpcEndpointArrayOutput values.
// You can construct a concrete instance of `VpcEndpointArrayInput` via:
//
//	VpcEndpointArray{ VpcEndpointArgs{...} }
type VpcEndpointArrayInput interface {
	pulumi.Input

	ToVpcEndpointArrayOutput() VpcEndpointArrayOutput
	ToVpcEndpointArrayOutputWithContext(context.Context) VpcEndpointArrayOutput
}

type VpcEndpointArray []VpcEndpointInput

func (VpcEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpoint)(nil)).Elem()
}

func (i VpcEndpointArray) ToVpcEndpointArrayOutput() VpcEndpointArrayOutput {
	return i.ToVpcEndpointArrayOutputWithContext(context.Background())
}

func (i VpcEndpointArray) ToVpcEndpointArrayOutputWithContext(ctx context.Context) VpcEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointArrayOutput)
}

// VpcEndpointMapInput is an input type that accepts VpcEndpointMap and VpcEndpointMapOutput values.
// You can construct a concrete instance of `VpcEndpointMapInput` via:
//
//	VpcEndpointMap{ "key": VpcEndpointArgs{...} }
type VpcEndpointMapInput interface {
	pulumi.Input

	ToVpcEndpointMapOutput() VpcEndpointMapOutput
	ToVpcEndpointMapOutputWithContext(context.Context) VpcEndpointMapOutput
}

type VpcEndpointMap map[string]VpcEndpointInput

func (VpcEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpoint)(nil)).Elem()
}

func (i VpcEndpointMap) ToVpcEndpointMapOutput() VpcEndpointMapOutput {
	return i.ToVpcEndpointMapOutputWithContext(context.Background())
}

func (i VpcEndpointMap) ToVpcEndpointMapOutputWithContext(ctx context.Context) VpcEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointMapOutput)
}

type VpcEndpointOutput struct{ *pulumi.OutputState }

func (VpcEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpoint)(nil)).Elem()
}

func (o VpcEndpointOutput) ToVpcEndpointOutput() VpcEndpointOutput {
	return o
}

func (o VpcEndpointOutput) ToVpcEndpointOutputWithContext(ctx context.Context) VpcEndpointOutput {
	return o
}

// Specifies the Amazon Resource Name (ARN) of the domain to create the endpoint for
func (o VpcEndpointOutput) DomainArn() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringOutput { return v.DomainArn }).(pulumi.StringOutput)
}

// The connection endpoint ID for connecting to the domain.
func (o VpcEndpointOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// Options to specify the subnets and security groups for the endpoint.
func (o VpcEndpointOutput) VpcOptions() VpcEndpointVpcOptionsOutput {
	return o.ApplyT(func(v *VpcEndpoint) VpcEndpointVpcOptionsOutput { return v.VpcOptions }).(VpcEndpointVpcOptionsOutput)
}

type VpcEndpointArrayOutput struct{ *pulumi.OutputState }

func (VpcEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpoint)(nil)).Elem()
}

func (o VpcEndpointArrayOutput) ToVpcEndpointArrayOutput() VpcEndpointArrayOutput {
	return o
}

func (o VpcEndpointArrayOutput) ToVpcEndpointArrayOutputWithContext(ctx context.Context) VpcEndpointArrayOutput {
	return o
}

func (o VpcEndpointArrayOutput) Index(i pulumi.IntInput) VpcEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcEndpoint {
		return vs[0].([]*VpcEndpoint)[vs[1].(int)]
	}).(VpcEndpointOutput)
}

type VpcEndpointMapOutput struct{ *pulumi.OutputState }

func (VpcEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpoint)(nil)).Elem()
}

func (o VpcEndpointMapOutput) ToVpcEndpointMapOutput() VpcEndpointMapOutput {
	return o
}

func (o VpcEndpointMapOutput) ToVpcEndpointMapOutputWithContext(ctx context.Context) VpcEndpointMapOutput {
	return o
}

func (o VpcEndpointMapOutput) MapIndex(k pulumi.StringInput) VpcEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcEndpoint {
		return vs[0].(map[string]*VpcEndpoint)[vs[1].(string)]
	}).(VpcEndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointInput)(nil)).Elem(), &VpcEndpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointArrayInput)(nil)).Elem(), VpcEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointMapInput)(nil)).Elem(), VpcEndpointMap{})
	pulumi.RegisterOutputType(VpcEndpointOutput{})
	pulumi.RegisterOutputType(VpcEndpointArrayOutput{})
	pulumi.RegisterOutputType(VpcEndpointMapOutput{})
}
