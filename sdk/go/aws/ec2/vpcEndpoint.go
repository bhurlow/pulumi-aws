// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VPC Endpoint resource.
//
// > **NOTE on VPC Endpoints and VPC Endpoint Associations:** The provider provides both standalone VPC Endpoint Associations for
// Route Tables - (an association between a VPC endpoint and a single `routeTableId`),
// Security Groups - (an association between a VPC endpoint and a single `securityGroupId`),
// and Subnets - (an association between a VPC endpoint and a single `subnetId`) and
// a VPC Endpoint resource with `routeTableIds` and `subnetIds` attributes.
// Do not use the same resource ID in both a VPC Endpoint resource and a VPC Endpoint Association resource.
// Doing so will cause a conflict of associations and will overwrite the association.
//
// ## Example Usage
// ### Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.NewVpcEndpoint(ctx, "s3", &ec2.VpcEndpointArgs{
//				VpcId:       pulumi.Any(aws_vpc.Main.Id),
//				ServiceName: pulumi.String("com.amazonaws.us-west-2.s3"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Basic w/ Tags
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.NewVpcEndpoint(ctx, "s3", &ec2.VpcEndpointArgs{
//				VpcId:       pulumi.Any(aws_vpc.Main.Id),
//				ServiceName: pulumi.String("com.amazonaws.us-west-2.s3"),
//				Tags: pulumi.StringMap{
//					"Environment": pulumi.String("test"),
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
// ### Interface Endpoint Type
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.NewVpcEndpoint(ctx, "ec2", &ec2.VpcEndpointArgs{
//				VpcId:           pulumi.Any(aws_vpc.Main.Id),
//				ServiceName:     pulumi.String("com.amazonaws.us-west-2.ec2"),
//				VpcEndpointType: pulumi.String("Interface"),
//				SecurityGroupIds: pulumi.StringArray{
//					aws_security_group.Sg1.Id,
//				},
//				PrivateDnsEnabled: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Gateway Load Balancer Endpoint Type
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			current, err := aws.GetCallerIdentity(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			exampleVpcEndpointService, err := ec2.NewVpcEndpointService(ctx, "exampleVpcEndpointService", &ec2.VpcEndpointServiceArgs{
//				AcceptanceRequired: pulumi.Bool(false),
//				AllowedPrincipals: pulumi.StringArray{
//					*pulumi.String(current.Arn),
//				},
//				GatewayLoadBalancerArns: pulumi.StringArray{
//					aws_lb.Example.Arn,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewVpcEndpoint(ctx, "exampleVpcEndpoint", &ec2.VpcEndpointArgs{
//				ServiceName: exampleVpcEndpointService.ServiceName,
//				SubnetIds: pulumi.StringArray{
//					aws_subnet.Example.Id,
//				},
//				VpcEndpointType: exampleVpcEndpointService.ServiceType,
//				VpcId:           pulumi.Any(aws_vpc.Example.Id),
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
//	to = aws_vpc_endpoint.endpoint1
//
//	id = "vpce-3ecf2a57" } Using `pulumi import`, import VPC Endpoints using the VPC endpoint `id`. For exampleconsole % pulumi import aws_vpc_endpoint.endpoint1 vpce-3ecf2a57
type VpcEndpoint struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the VPC endpoint.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
	AutoAccept pulumi.BoolPtrOutput `pulumi:"autoAccept"`
	// The list of CIDR blocks for the exposed AWS service. Applicable for endpoints of type `Gateway`.
	CidrBlocks pulumi.StringArrayOutput `pulumi:"cidrBlocks"`
	// The DNS entries for the VPC Endpoint. Applicable for endpoints of type `Interface`. DNS blocks are documented below.
	DnsEntries VpcEndpointDnsEntryArrayOutput `pulumi:"dnsEntries"`
	// The DNS options for the endpoint. See dnsOptions below.
	DnsOptions VpcEndpointDnsOptionsOutput `pulumi:"dnsOptions"`
	// The IP address type for the endpoint. Valid values are `ipv4`, `dualstack`, and `ipv6`.
	IpAddressType pulumi.StringOutput `pulumi:"ipAddressType"`
	// One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type `Interface`.
	NetworkInterfaceIds pulumi.StringArrayOutput `pulumi:"networkInterfaceIds"`
	// The ID of the AWS account that owns the VPC endpoint.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// A policy to attach to the endpoint that controls access to the service. This is a JSON formatted string. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
	Policy pulumi.StringOutput `pulumi:"policy"`
	// The prefix list ID of the exposed AWS service. Applicable for endpoints of type `Gateway`.
	PrefixListId pulumi.StringOutput `pulumi:"prefixListId"`
	// Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
	// Defaults to `false`.
	PrivateDnsEnabled pulumi.BoolPtrOutput `pulumi:"privateDnsEnabled"`
	// Whether or not the VPC Endpoint is being managed by its service - `true` or `false`.
	RequesterManaged pulumi.BoolOutput `pulumi:"requesterManaged"`
	// One or more route table IDs. Applicable for endpoints of type `Gateway`.
	RouteTableIds pulumi.StringArrayOutput `pulumi:"routeTableIds"`
	// The ID of one or more security groups to associate with the network interface. Applicable for endpoints of type `Interface`.
	// If no security groups are specified, the VPC's [default security group](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html#DefaultSecurityGroup) is associated with the endpoint.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// The service name. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// The state of the VPC endpoint.
	State pulumi.StringOutput `pulumi:"state"`
	// The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `GatewayLoadBalancer` and `Interface`.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The VPC endpoint type, `Gateway`, `GatewayLoadBalancer`, or `Interface`. Defaults to `Gateway`.
	VpcEndpointType pulumi.StringPtrOutput `pulumi:"vpcEndpointType"`
	// The ID of the VPC in which the endpoint will be used.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewVpcEndpoint registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpoint(ctx *pulumi.Context,
	name string, args *VpcEndpointArgs, opts ...pulumi.ResourceOption) (*VpcEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcEndpoint
	err := ctx.RegisterResource("aws:ec2/vpcEndpoint:VpcEndpoint", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:ec2/vpcEndpoint:VpcEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcEndpoint resources.
type vpcEndpointState struct {
	// The Amazon Resource Name (ARN) of the VPC endpoint.
	Arn *string `pulumi:"arn"`
	// Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
	AutoAccept *bool `pulumi:"autoAccept"`
	// The list of CIDR blocks for the exposed AWS service. Applicable for endpoints of type `Gateway`.
	CidrBlocks []string `pulumi:"cidrBlocks"`
	// The DNS entries for the VPC Endpoint. Applicable for endpoints of type `Interface`. DNS blocks are documented below.
	DnsEntries []VpcEndpointDnsEntry `pulumi:"dnsEntries"`
	// The DNS options for the endpoint. See dnsOptions below.
	DnsOptions *VpcEndpointDnsOptions `pulumi:"dnsOptions"`
	// The IP address type for the endpoint. Valid values are `ipv4`, `dualstack`, and `ipv6`.
	IpAddressType *string `pulumi:"ipAddressType"`
	// One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type `Interface`.
	NetworkInterfaceIds []string `pulumi:"networkInterfaceIds"`
	// The ID of the AWS account that owns the VPC endpoint.
	OwnerId *string `pulumi:"ownerId"`
	// A policy to attach to the endpoint that controls access to the service. This is a JSON formatted string. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
	Policy *string `pulumi:"policy"`
	// The prefix list ID of the exposed AWS service. Applicable for endpoints of type `Gateway`.
	PrefixListId *string `pulumi:"prefixListId"`
	// Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
	// Defaults to `false`.
	PrivateDnsEnabled *bool `pulumi:"privateDnsEnabled"`
	// Whether or not the VPC Endpoint is being managed by its service - `true` or `false`.
	RequesterManaged *bool `pulumi:"requesterManaged"`
	// One or more route table IDs. Applicable for endpoints of type `Gateway`.
	RouteTableIds []string `pulumi:"routeTableIds"`
	// The ID of one or more security groups to associate with the network interface. Applicable for endpoints of type `Interface`.
	// If no security groups are specified, the VPC's [default security group](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html#DefaultSecurityGroup) is associated with the endpoint.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The service name. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
	ServiceName *string `pulumi:"serviceName"`
	// The state of the VPC endpoint.
	State *string `pulumi:"state"`
	// The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `GatewayLoadBalancer` and `Interface`.
	SubnetIds []string `pulumi:"subnetIds"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The VPC endpoint type, `Gateway`, `GatewayLoadBalancer`, or `Interface`. Defaults to `Gateway`.
	VpcEndpointType *string `pulumi:"vpcEndpointType"`
	// The ID of the VPC in which the endpoint will be used.
	VpcId *string `pulumi:"vpcId"`
}

type VpcEndpointState struct {
	// The Amazon Resource Name (ARN) of the VPC endpoint.
	Arn pulumi.StringPtrInput
	// Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
	AutoAccept pulumi.BoolPtrInput
	// The list of CIDR blocks for the exposed AWS service. Applicable for endpoints of type `Gateway`.
	CidrBlocks pulumi.StringArrayInput
	// The DNS entries for the VPC Endpoint. Applicable for endpoints of type `Interface`. DNS blocks are documented below.
	DnsEntries VpcEndpointDnsEntryArrayInput
	// The DNS options for the endpoint. See dnsOptions below.
	DnsOptions VpcEndpointDnsOptionsPtrInput
	// The IP address type for the endpoint. Valid values are `ipv4`, `dualstack`, and `ipv6`.
	IpAddressType pulumi.StringPtrInput
	// One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type `Interface`.
	NetworkInterfaceIds pulumi.StringArrayInput
	// The ID of the AWS account that owns the VPC endpoint.
	OwnerId pulumi.StringPtrInput
	// A policy to attach to the endpoint that controls access to the service. This is a JSON formatted string. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
	Policy pulumi.StringPtrInput
	// The prefix list ID of the exposed AWS service. Applicable for endpoints of type `Gateway`.
	PrefixListId pulumi.StringPtrInput
	// Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
	// Defaults to `false`.
	PrivateDnsEnabled pulumi.BoolPtrInput
	// Whether or not the VPC Endpoint is being managed by its service - `true` or `false`.
	RequesterManaged pulumi.BoolPtrInput
	// One or more route table IDs. Applicable for endpoints of type `Gateway`.
	RouteTableIds pulumi.StringArrayInput
	// The ID of one or more security groups to associate with the network interface. Applicable for endpoints of type `Interface`.
	// If no security groups are specified, the VPC's [default security group](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html#DefaultSecurityGroup) is associated with the endpoint.
	SecurityGroupIds pulumi.StringArrayInput
	// The service name. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
	ServiceName pulumi.StringPtrInput
	// The state of the VPC endpoint.
	State pulumi.StringPtrInput
	// The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `GatewayLoadBalancer` and `Interface`.
	SubnetIds pulumi.StringArrayInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// The VPC endpoint type, `Gateway`, `GatewayLoadBalancer`, or `Interface`. Defaults to `Gateway`.
	VpcEndpointType pulumi.StringPtrInput
	// The ID of the VPC in which the endpoint will be used.
	VpcId pulumi.StringPtrInput
}

func (VpcEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointState)(nil)).Elem()
}

type vpcEndpointArgs struct {
	// Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
	AutoAccept *bool `pulumi:"autoAccept"`
	// The DNS options for the endpoint. See dnsOptions below.
	DnsOptions *VpcEndpointDnsOptions `pulumi:"dnsOptions"`
	// The IP address type for the endpoint. Valid values are `ipv4`, `dualstack`, and `ipv6`.
	IpAddressType *string `pulumi:"ipAddressType"`
	// A policy to attach to the endpoint that controls access to the service. This is a JSON formatted string. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
	Policy *string `pulumi:"policy"`
	// Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
	// Defaults to `false`.
	PrivateDnsEnabled *bool `pulumi:"privateDnsEnabled"`
	// One or more route table IDs. Applicable for endpoints of type `Gateway`.
	RouteTableIds []string `pulumi:"routeTableIds"`
	// The ID of one or more security groups to associate with the network interface. Applicable for endpoints of type `Interface`.
	// If no security groups are specified, the VPC's [default security group](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html#DefaultSecurityGroup) is associated with the endpoint.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The service name. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
	ServiceName string `pulumi:"serviceName"`
	// The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `GatewayLoadBalancer` and `Interface`.
	SubnetIds []string `pulumi:"subnetIds"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The VPC endpoint type, `Gateway`, `GatewayLoadBalancer`, or `Interface`. Defaults to `Gateway`.
	VpcEndpointType *string `pulumi:"vpcEndpointType"`
	// The ID of the VPC in which the endpoint will be used.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a VpcEndpoint resource.
type VpcEndpointArgs struct {
	// Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
	AutoAccept pulumi.BoolPtrInput
	// The DNS options for the endpoint. See dnsOptions below.
	DnsOptions VpcEndpointDnsOptionsPtrInput
	// The IP address type for the endpoint. Valid values are `ipv4`, `dualstack`, and `ipv6`.
	IpAddressType pulumi.StringPtrInput
	// A policy to attach to the endpoint that controls access to the service. This is a JSON formatted string. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
	Policy pulumi.StringPtrInput
	// Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
	// Defaults to `false`.
	PrivateDnsEnabled pulumi.BoolPtrInput
	// One or more route table IDs. Applicable for endpoints of type `Gateway`.
	RouteTableIds pulumi.StringArrayInput
	// The ID of one or more security groups to associate with the network interface. Applicable for endpoints of type `Interface`.
	// If no security groups are specified, the VPC's [default security group](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html#DefaultSecurityGroup) is associated with the endpoint.
	SecurityGroupIds pulumi.StringArrayInput
	// The service name. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
	ServiceName pulumi.StringInput
	// The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `GatewayLoadBalancer` and `Interface`.
	SubnetIds pulumi.StringArrayInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The VPC endpoint type, `Gateway`, `GatewayLoadBalancer`, or `Interface`. Defaults to `Gateway`.
	VpcEndpointType pulumi.StringPtrInput
	// The ID of the VPC in which the endpoint will be used.
	VpcId pulumi.StringInput
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

// The Amazon Resource Name (ARN) of the VPC endpoint.
func (o VpcEndpointOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
func (o VpcEndpointOutput) AutoAccept() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.BoolPtrOutput { return v.AutoAccept }).(pulumi.BoolPtrOutput)
}

// The list of CIDR blocks for the exposed AWS service. Applicable for endpoints of type `Gateway`.
func (o VpcEndpointOutput) CidrBlocks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringArrayOutput { return v.CidrBlocks }).(pulumi.StringArrayOutput)
}

// The DNS entries for the VPC Endpoint. Applicable for endpoints of type `Interface`. DNS blocks are documented below.
func (o VpcEndpointOutput) DnsEntries() VpcEndpointDnsEntryArrayOutput {
	return o.ApplyT(func(v *VpcEndpoint) VpcEndpointDnsEntryArrayOutput { return v.DnsEntries }).(VpcEndpointDnsEntryArrayOutput)
}

// The DNS options for the endpoint. See dnsOptions below.
func (o VpcEndpointOutput) DnsOptions() VpcEndpointDnsOptionsOutput {
	return o.ApplyT(func(v *VpcEndpoint) VpcEndpointDnsOptionsOutput { return v.DnsOptions }).(VpcEndpointDnsOptionsOutput)
}

// The IP address type for the endpoint. Valid values are `ipv4`, `dualstack`, and `ipv6`.
func (o VpcEndpointOutput) IpAddressType() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringOutput { return v.IpAddressType }).(pulumi.StringOutput)
}

// One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type `Interface`.
func (o VpcEndpointOutput) NetworkInterfaceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringArrayOutput { return v.NetworkInterfaceIds }).(pulumi.StringArrayOutput)
}

// The ID of the AWS account that owns the VPC endpoint.
func (o VpcEndpointOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// A policy to attach to the endpoint that controls access to the service. This is a JSON formatted string. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
func (o VpcEndpointOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

// The prefix list ID of the exposed AWS service. Applicable for endpoints of type `Gateway`.
func (o VpcEndpointOutput) PrefixListId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringOutput { return v.PrefixListId }).(pulumi.StringOutput)
}

// Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
// Defaults to `false`.
func (o VpcEndpointOutput) PrivateDnsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.BoolPtrOutput { return v.PrivateDnsEnabled }).(pulumi.BoolPtrOutput)
}

// Whether or not the VPC Endpoint is being managed by its service - `true` or `false`.
func (o VpcEndpointOutput) RequesterManaged() pulumi.BoolOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.BoolOutput { return v.RequesterManaged }).(pulumi.BoolOutput)
}

// One or more route table IDs. Applicable for endpoints of type `Gateway`.
func (o VpcEndpointOutput) RouteTableIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringArrayOutput { return v.RouteTableIds }).(pulumi.StringArrayOutput)
}

// The ID of one or more security groups to associate with the network interface. Applicable for endpoints of type `Interface`.
// If no security groups are specified, the VPC's [default security group](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html#DefaultSecurityGroup) is associated with the endpoint.
func (o VpcEndpointOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The service name. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
func (o VpcEndpointOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// The state of the VPC endpoint.
func (o VpcEndpointOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `GatewayLoadBalancer` and `Interface`.
func (o VpcEndpointOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o VpcEndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o VpcEndpointOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The VPC endpoint type, `Gateway`, `GatewayLoadBalancer`, or `Interface`. Defaults to `Gateway`.
func (o VpcEndpointOutput) VpcEndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringPtrOutput { return v.VpcEndpointType }).(pulumi.StringPtrOutput)
}

// The ID of the VPC in which the endpoint will be used.
func (o VpcEndpointOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpoint) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
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
