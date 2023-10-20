// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a resource to manage the accepter's side of a VPC Peering Connection.
//
// When a cross-account (requester's AWS account differs from the accepter's AWS account) or an inter-region
// VPC Peering Connection is created, a VPC Peering Connection resource is automatically created in the
// accepter's account.
// The requester can use the `ec2.VpcPeeringConnection` resource to manage its side of the connection
// and the accepter can use the `ec2.VpcPeeringConnectionAccepter` resource to "adopt" its side of the
// connection into management.
//
// ## Example Usage
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
//			_, err := aws.NewProvider(ctx, "peer", &aws.ProviderArgs{
//				Region: pulumi.String("us-west-2"),
//			})
//			if err != nil {
//				return err
//			}
//			main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
//				CidrBlock: pulumi.String("10.0.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			peerVpc, err := ec2.NewVpc(ctx, "peerVpc", &ec2.VpcArgs{
//				CidrBlock: pulumi.String("10.1.0.0/16"),
//			}, pulumi.Provider(aws.Peer))
//			if err != nil {
//				return err
//			}
//			peerCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			peerVpcPeeringConnection, err := ec2.NewVpcPeeringConnection(ctx, "peerVpcPeeringConnection", &ec2.VpcPeeringConnectionArgs{
//				VpcId:       main.ID(),
//				PeerVpcId:   peerVpc.ID(),
//				PeerOwnerId: *pulumi.String(peerCallerIdentity.AccountId),
//				PeerRegion:  pulumi.String("us-west-2"),
//				AutoAccept:  pulumi.Bool(false),
//				Tags: pulumi.StringMap{
//					"Side": pulumi.String("Requester"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewVpcPeeringConnectionAccepter(ctx, "peerVpcPeeringConnectionAccepter", &ec2.VpcPeeringConnectionAccepterArgs{
//				VpcPeeringConnectionId: peerVpcPeeringConnection.ID(),
//				AutoAccept:             pulumi.Bool(true),
//				Tags: pulumi.StringMap{
//					"Side": pulumi.String("Accepter"),
//				},
//			}, pulumi.Provider(aws.Peer))
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
// Using `pulumi import`, import VPC Peering Connection Accepters using the Peering Connection ID. For example:
//
// ```sh
//
//	$ pulumi import aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter example pcx-12345678
//
// ```
//
//	Certain resource arguments, like `auto_accept`, do not have an EC2 API method for reading the information after peering connection creation. If the argument is set in the Pulumi program on an imported resource, Pulumi will always show a difference. To workaround this behavior, either omit the argument from the Pulumi program or use `ignore_changes` to hide the difference. For example:
type VpcPeeringConnectionAccepter struct {
	pulumi.CustomResourceState

	// The status of the VPC Peering Connection request.
	AcceptStatus pulumi.StringOutput `pulumi:"acceptStatus"`
	// A configuration block that describes [VPC Peering Connection]
	// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
	Accepter VpcPeeringConnectionAccepterAccepterOutput `pulumi:"accepter"`
	// Whether or not to accept the peering request. Defaults to `false`.
	AutoAccept pulumi.BoolPtrOutput `pulumi:"autoAccept"`
	// The AWS account ID of the owner of the requester VPC.
	PeerOwnerId pulumi.StringOutput `pulumi:"peerOwnerId"`
	// The region of the accepter VPC.
	PeerRegion pulumi.StringOutput `pulumi:"peerRegion"`
	// The ID of the requester VPC.
	PeerVpcId pulumi.StringOutput `pulumi:"peerVpcId"`
	// A configuration block that describes [VPC Peering Connection]
	// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
	Requester VpcPeeringConnectionAccepterRequesterOutput `pulumi:"requester"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The ID of the accepter VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The VPC Peering Connection ID to manage.
	VpcPeeringConnectionId pulumi.StringOutput `pulumi:"vpcPeeringConnectionId"`
}

// NewVpcPeeringConnectionAccepter registers a new resource with the given unique name, arguments, and options.
func NewVpcPeeringConnectionAccepter(ctx *pulumi.Context,
	name string, args *VpcPeeringConnectionAccepterArgs, opts ...pulumi.ResourceOption) (*VpcPeeringConnectionAccepter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcPeeringConnectionId == nil {
		return nil, errors.New("invalid value for required argument 'VpcPeeringConnectionId'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcPeeringConnectionAccepter
	err := ctx.RegisterResource("aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcPeeringConnectionAccepter gets an existing VpcPeeringConnectionAccepter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcPeeringConnectionAccepter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcPeeringConnectionAccepterState, opts ...pulumi.ResourceOption) (*VpcPeeringConnectionAccepter, error) {
	var resource VpcPeeringConnectionAccepter
	err := ctx.ReadResource("aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcPeeringConnectionAccepter resources.
type vpcPeeringConnectionAccepterState struct {
	// The status of the VPC Peering Connection request.
	AcceptStatus *string `pulumi:"acceptStatus"`
	// A configuration block that describes [VPC Peering Connection]
	// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
	Accepter *VpcPeeringConnectionAccepterAccepter `pulumi:"accepter"`
	// Whether or not to accept the peering request. Defaults to `false`.
	AutoAccept *bool `pulumi:"autoAccept"`
	// The AWS account ID of the owner of the requester VPC.
	PeerOwnerId *string `pulumi:"peerOwnerId"`
	// The region of the accepter VPC.
	PeerRegion *string `pulumi:"peerRegion"`
	// The ID of the requester VPC.
	PeerVpcId *string `pulumi:"peerVpcId"`
	// A configuration block that describes [VPC Peering Connection]
	// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
	Requester *VpcPeeringConnectionAccepterRequester `pulumi:"requester"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The ID of the accepter VPC.
	VpcId *string `pulumi:"vpcId"`
	// The VPC Peering Connection ID to manage.
	VpcPeeringConnectionId *string `pulumi:"vpcPeeringConnectionId"`
}

type VpcPeeringConnectionAccepterState struct {
	// The status of the VPC Peering Connection request.
	AcceptStatus pulumi.StringPtrInput
	// A configuration block that describes [VPC Peering Connection]
	// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
	Accepter VpcPeeringConnectionAccepterAccepterPtrInput
	// Whether or not to accept the peering request. Defaults to `false`.
	AutoAccept pulumi.BoolPtrInput
	// The AWS account ID of the owner of the requester VPC.
	PeerOwnerId pulumi.StringPtrInput
	// The region of the accepter VPC.
	PeerRegion pulumi.StringPtrInput
	// The ID of the requester VPC.
	PeerVpcId pulumi.StringPtrInput
	// A configuration block that describes [VPC Peering Connection]
	// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
	Requester VpcPeeringConnectionAccepterRequesterPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
	// The ID of the accepter VPC.
	VpcId pulumi.StringPtrInput
	// The VPC Peering Connection ID to manage.
	VpcPeeringConnectionId pulumi.StringPtrInput
}

func (VpcPeeringConnectionAccepterState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPeeringConnectionAccepterState)(nil)).Elem()
}

type vpcPeeringConnectionAccepterArgs struct {
	// A configuration block that describes [VPC Peering Connection]
	// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
	Accepter *VpcPeeringConnectionAccepterAccepter `pulumi:"accepter"`
	// Whether or not to accept the peering request. Defaults to `false`.
	AutoAccept *bool `pulumi:"autoAccept"`
	// A configuration block that describes [VPC Peering Connection]
	// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
	Requester *VpcPeeringConnectionAccepterRequester `pulumi:"requester"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The VPC Peering Connection ID to manage.
	VpcPeeringConnectionId string `pulumi:"vpcPeeringConnectionId"`
}

// The set of arguments for constructing a VpcPeeringConnectionAccepter resource.
type VpcPeeringConnectionAccepterArgs struct {
	// A configuration block that describes [VPC Peering Connection]
	// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
	Accepter VpcPeeringConnectionAccepterAccepterPtrInput
	// Whether or not to accept the peering request. Defaults to `false`.
	AutoAccept pulumi.BoolPtrInput
	// A configuration block that describes [VPC Peering Connection]
	// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
	Requester VpcPeeringConnectionAccepterRequesterPtrInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The VPC Peering Connection ID to manage.
	VpcPeeringConnectionId pulumi.StringInput
}

func (VpcPeeringConnectionAccepterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPeeringConnectionAccepterArgs)(nil)).Elem()
}

type VpcPeeringConnectionAccepterInput interface {
	pulumi.Input

	ToVpcPeeringConnectionAccepterOutput() VpcPeeringConnectionAccepterOutput
	ToVpcPeeringConnectionAccepterOutputWithContext(ctx context.Context) VpcPeeringConnectionAccepterOutput
}

func (*VpcPeeringConnectionAccepter) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPeeringConnectionAccepter)(nil)).Elem()
}

func (i *VpcPeeringConnectionAccepter) ToVpcPeeringConnectionAccepterOutput() VpcPeeringConnectionAccepterOutput {
	return i.ToVpcPeeringConnectionAccepterOutputWithContext(context.Background())
}

func (i *VpcPeeringConnectionAccepter) ToVpcPeeringConnectionAccepterOutputWithContext(ctx context.Context) VpcPeeringConnectionAccepterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPeeringConnectionAccepterOutput)
}

func (i *VpcPeeringConnectionAccepter) ToOutput(ctx context.Context) pulumix.Output[*VpcPeeringConnectionAccepter] {
	return pulumix.Output[*VpcPeeringConnectionAccepter]{
		OutputState: i.ToVpcPeeringConnectionAccepterOutputWithContext(ctx).OutputState,
	}
}

// VpcPeeringConnectionAccepterArrayInput is an input type that accepts VpcPeeringConnectionAccepterArray and VpcPeeringConnectionAccepterArrayOutput values.
// You can construct a concrete instance of `VpcPeeringConnectionAccepterArrayInput` via:
//
//	VpcPeeringConnectionAccepterArray{ VpcPeeringConnectionAccepterArgs{...} }
type VpcPeeringConnectionAccepterArrayInput interface {
	pulumi.Input

	ToVpcPeeringConnectionAccepterArrayOutput() VpcPeeringConnectionAccepterArrayOutput
	ToVpcPeeringConnectionAccepterArrayOutputWithContext(context.Context) VpcPeeringConnectionAccepterArrayOutput
}

type VpcPeeringConnectionAccepterArray []VpcPeeringConnectionAccepterInput

func (VpcPeeringConnectionAccepterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcPeeringConnectionAccepter)(nil)).Elem()
}

func (i VpcPeeringConnectionAccepterArray) ToVpcPeeringConnectionAccepterArrayOutput() VpcPeeringConnectionAccepterArrayOutput {
	return i.ToVpcPeeringConnectionAccepterArrayOutputWithContext(context.Background())
}

func (i VpcPeeringConnectionAccepterArray) ToVpcPeeringConnectionAccepterArrayOutputWithContext(ctx context.Context) VpcPeeringConnectionAccepterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPeeringConnectionAccepterArrayOutput)
}

func (i VpcPeeringConnectionAccepterArray) ToOutput(ctx context.Context) pulumix.Output[[]*VpcPeeringConnectionAccepter] {
	return pulumix.Output[[]*VpcPeeringConnectionAccepter]{
		OutputState: i.ToVpcPeeringConnectionAccepterArrayOutputWithContext(ctx).OutputState,
	}
}

// VpcPeeringConnectionAccepterMapInput is an input type that accepts VpcPeeringConnectionAccepterMap and VpcPeeringConnectionAccepterMapOutput values.
// You can construct a concrete instance of `VpcPeeringConnectionAccepterMapInput` via:
//
//	VpcPeeringConnectionAccepterMap{ "key": VpcPeeringConnectionAccepterArgs{...} }
type VpcPeeringConnectionAccepterMapInput interface {
	pulumi.Input

	ToVpcPeeringConnectionAccepterMapOutput() VpcPeeringConnectionAccepterMapOutput
	ToVpcPeeringConnectionAccepterMapOutputWithContext(context.Context) VpcPeeringConnectionAccepterMapOutput
}

type VpcPeeringConnectionAccepterMap map[string]VpcPeeringConnectionAccepterInput

func (VpcPeeringConnectionAccepterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcPeeringConnectionAccepter)(nil)).Elem()
}

func (i VpcPeeringConnectionAccepterMap) ToVpcPeeringConnectionAccepterMapOutput() VpcPeeringConnectionAccepterMapOutput {
	return i.ToVpcPeeringConnectionAccepterMapOutputWithContext(context.Background())
}

func (i VpcPeeringConnectionAccepterMap) ToVpcPeeringConnectionAccepterMapOutputWithContext(ctx context.Context) VpcPeeringConnectionAccepterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPeeringConnectionAccepterMapOutput)
}

func (i VpcPeeringConnectionAccepterMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*VpcPeeringConnectionAccepter] {
	return pulumix.Output[map[string]*VpcPeeringConnectionAccepter]{
		OutputState: i.ToVpcPeeringConnectionAccepterMapOutputWithContext(ctx).OutputState,
	}
}

type VpcPeeringConnectionAccepterOutput struct{ *pulumi.OutputState }

func (VpcPeeringConnectionAccepterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPeeringConnectionAccepter)(nil)).Elem()
}

func (o VpcPeeringConnectionAccepterOutput) ToVpcPeeringConnectionAccepterOutput() VpcPeeringConnectionAccepterOutput {
	return o
}

func (o VpcPeeringConnectionAccepterOutput) ToVpcPeeringConnectionAccepterOutputWithContext(ctx context.Context) VpcPeeringConnectionAccepterOutput {
	return o
}

func (o VpcPeeringConnectionAccepterOutput) ToOutput(ctx context.Context) pulumix.Output[*VpcPeeringConnectionAccepter] {
	return pulumix.Output[*VpcPeeringConnectionAccepter]{
		OutputState: o.OutputState,
	}
}

// The status of the VPC Peering Connection request.
func (o VpcPeeringConnectionAccepterOutput) AcceptStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPeeringConnectionAccepter) pulumi.StringOutput { return v.AcceptStatus }).(pulumi.StringOutput)
}

// A configuration block that describes [VPC Peering Connection]
// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
func (o VpcPeeringConnectionAccepterOutput) Accepter() VpcPeeringConnectionAccepterAccepterOutput {
	return o.ApplyT(func(v *VpcPeeringConnectionAccepter) VpcPeeringConnectionAccepterAccepterOutput { return v.Accepter }).(VpcPeeringConnectionAccepterAccepterOutput)
}

// Whether or not to accept the peering request. Defaults to `false`.
func (o VpcPeeringConnectionAccepterOutput) AutoAccept() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpcPeeringConnectionAccepter) pulumi.BoolPtrOutput { return v.AutoAccept }).(pulumi.BoolPtrOutput)
}

// The AWS account ID of the owner of the requester VPC.
func (o VpcPeeringConnectionAccepterOutput) PeerOwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPeeringConnectionAccepter) pulumi.StringOutput { return v.PeerOwnerId }).(pulumi.StringOutput)
}

// The region of the accepter VPC.
func (o VpcPeeringConnectionAccepterOutput) PeerRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPeeringConnectionAccepter) pulumi.StringOutput { return v.PeerRegion }).(pulumi.StringOutput)
}

// The ID of the requester VPC.
func (o VpcPeeringConnectionAccepterOutput) PeerVpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPeeringConnectionAccepter) pulumi.StringOutput { return v.PeerVpcId }).(pulumi.StringOutput)
}

// A configuration block that describes [VPC Peering Connection]
// (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
func (o VpcPeeringConnectionAccepterOutput) Requester() VpcPeeringConnectionAccepterRequesterOutput {
	return o.ApplyT(func(v *VpcPeeringConnectionAccepter) VpcPeeringConnectionAccepterRequesterOutput { return v.Requester }).(VpcPeeringConnectionAccepterRequesterOutput)
}

// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o VpcPeeringConnectionAccepterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcPeeringConnectionAccepter) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o VpcPeeringConnectionAccepterOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcPeeringConnectionAccepter) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// The ID of the accepter VPC.
func (o VpcPeeringConnectionAccepterOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPeeringConnectionAccepter) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The VPC Peering Connection ID to manage.
func (o VpcPeeringConnectionAccepterOutput) VpcPeeringConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPeeringConnectionAccepter) pulumi.StringOutput { return v.VpcPeeringConnectionId }).(pulumi.StringOutput)
}

type VpcPeeringConnectionAccepterArrayOutput struct{ *pulumi.OutputState }

func (VpcPeeringConnectionAccepterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcPeeringConnectionAccepter)(nil)).Elem()
}

func (o VpcPeeringConnectionAccepterArrayOutput) ToVpcPeeringConnectionAccepterArrayOutput() VpcPeeringConnectionAccepterArrayOutput {
	return o
}

func (o VpcPeeringConnectionAccepterArrayOutput) ToVpcPeeringConnectionAccepterArrayOutputWithContext(ctx context.Context) VpcPeeringConnectionAccepterArrayOutput {
	return o
}

func (o VpcPeeringConnectionAccepterArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*VpcPeeringConnectionAccepter] {
	return pulumix.Output[[]*VpcPeeringConnectionAccepter]{
		OutputState: o.OutputState,
	}
}

func (o VpcPeeringConnectionAccepterArrayOutput) Index(i pulumi.IntInput) VpcPeeringConnectionAccepterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcPeeringConnectionAccepter {
		return vs[0].([]*VpcPeeringConnectionAccepter)[vs[1].(int)]
	}).(VpcPeeringConnectionAccepterOutput)
}

type VpcPeeringConnectionAccepterMapOutput struct{ *pulumi.OutputState }

func (VpcPeeringConnectionAccepterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcPeeringConnectionAccepter)(nil)).Elem()
}

func (o VpcPeeringConnectionAccepterMapOutput) ToVpcPeeringConnectionAccepterMapOutput() VpcPeeringConnectionAccepterMapOutput {
	return o
}

func (o VpcPeeringConnectionAccepterMapOutput) ToVpcPeeringConnectionAccepterMapOutputWithContext(ctx context.Context) VpcPeeringConnectionAccepterMapOutput {
	return o
}

func (o VpcPeeringConnectionAccepterMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*VpcPeeringConnectionAccepter] {
	return pulumix.Output[map[string]*VpcPeeringConnectionAccepter]{
		OutputState: o.OutputState,
	}
}

func (o VpcPeeringConnectionAccepterMapOutput) MapIndex(k pulumi.StringInput) VpcPeeringConnectionAccepterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcPeeringConnectionAccepter {
		return vs[0].(map[string]*VpcPeeringConnectionAccepter)[vs[1].(string)]
	}).(VpcPeeringConnectionAccepterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPeeringConnectionAccepterInput)(nil)).Elem(), &VpcPeeringConnectionAccepter{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPeeringConnectionAccepterArrayInput)(nil)).Elem(), VpcPeeringConnectionAccepterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPeeringConnectionAccepterMapInput)(nil)).Elem(), VpcPeeringConnectionAccepterMap{})
	pulumi.RegisterOutputType(VpcPeeringConnectionAccepterOutput{})
	pulumi.RegisterOutputType(VpcPeeringConnectionAccepterArrayOutput{})
	pulumi.RegisterOutputType(VpcPeeringConnectionAccepterMapOutput{})
}
