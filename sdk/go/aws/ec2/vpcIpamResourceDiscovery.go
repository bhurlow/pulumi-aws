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

// Provides an IPAM Resource Discovery resource. IPAM Resource Discoveries are resources meant for multi-organization customers. If you wish to use a single IPAM across multiple orgs, a resource discovery can be created and shared from a subordinate organization to the management organizations IPAM delegated admin account. For a full deployment example, see `ec2.VpcIpamResourceDiscoveryAssociation` resource.
//
// ## Example Usage
//
// Basic usage:
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
//			current, err := aws.GetRegion(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ec2.NewVpcIpamResourceDiscovery(ctx, "main", &ec2.VpcIpamResourceDiscoveryArgs{
//				Description: pulumi.String("My IPAM Resource Discovery"),
//				OperatingRegions: ec2.VpcIpamResourceDiscoveryOperatingRegionArray{
//					&ec2.VpcIpamResourceDiscoveryOperatingRegionArgs{
//						RegionName: *pulumi.String(current.Name),
//					},
//				},
//				Tags: pulumi.StringMap{
//					"Test": pulumi.String("Main"),
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
//	to = aws_vpc_ipam_resource_discovery.example
//
//	id = "ipam-res-disco-0178368ad2146a492" } Using `pulumi import`, import IPAMs using the IPAM resource discovery `id`. For exampleconsole % pulumi import aws_vpc_ipam_resource_discovery.example ipam-res-disco-0178368ad2146a492
type VpcIpamResourceDiscovery struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of IPAM Resource Discovery
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A description for the IPAM Resource Discovery.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The home region of the Resource Discovery
	IpamResourceDiscoveryRegion pulumi.StringOutput `pulumi:"ipamResourceDiscoveryRegion"`
	// A boolean to identify if the Resource Discovery is the accounts default resource discovery
	IsDefault pulumi.BoolOutput `pulumi:"isDefault"`
	// Determines which regions the Resource Discovery will enable IPAM features for usage and monitoring. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM Resource Discovery. You can only create VPCs from a pool whose locale matches the VPC's Region. You specify a region using the regionName parameter. **You must set your provider block region as an operating_region.**
	OperatingRegions VpcIpamResourceDiscoveryOperatingRegionArrayOutput `pulumi:"operatingRegions"`
	// The account ID for the account that manages the Resource Discovery
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewVpcIpamResourceDiscovery registers a new resource with the given unique name, arguments, and options.
func NewVpcIpamResourceDiscovery(ctx *pulumi.Context,
	name string, args *VpcIpamResourceDiscoveryArgs, opts ...pulumi.ResourceOption) (*VpcIpamResourceDiscovery, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OperatingRegions == nil {
		return nil, errors.New("invalid value for required argument 'OperatingRegions'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcIpamResourceDiscovery
	err := ctx.RegisterResource("aws:ec2/vpcIpamResourceDiscovery:VpcIpamResourceDiscovery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcIpamResourceDiscovery gets an existing VpcIpamResourceDiscovery resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcIpamResourceDiscovery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcIpamResourceDiscoveryState, opts ...pulumi.ResourceOption) (*VpcIpamResourceDiscovery, error) {
	var resource VpcIpamResourceDiscovery
	err := ctx.ReadResource("aws:ec2/vpcIpamResourceDiscovery:VpcIpamResourceDiscovery", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcIpamResourceDiscovery resources.
type vpcIpamResourceDiscoveryState struct {
	// Amazon Resource Name (ARN) of IPAM Resource Discovery
	Arn *string `pulumi:"arn"`
	// A description for the IPAM Resource Discovery.
	Description *string `pulumi:"description"`
	// The home region of the Resource Discovery
	IpamResourceDiscoveryRegion *string `pulumi:"ipamResourceDiscoveryRegion"`
	// A boolean to identify if the Resource Discovery is the accounts default resource discovery
	IsDefault *bool `pulumi:"isDefault"`
	// Determines which regions the Resource Discovery will enable IPAM features for usage and monitoring. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM Resource Discovery. You can only create VPCs from a pool whose locale matches the VPC's Region. You specify a region using the regionName parameter. **You must set your provider block region as an operating_region.**
	OperatingRegions []VpcIpamResourceDiscoveryOperatingRegion `pulumi:"operatingRegions"`
	// The account ID for the account that manages the Resource Discovery
	OwnerId *string `pulumi:"ownerId"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type VpcIpamResourceDiscoveryState struct {
	// Amazon Resource Name (ARN) of IPAM Resource Discovery
	Arn pulumi.StringPtrInput
	// A description for the IPAM Resource Discovery.
	Description pulumi.StringPtrInput
	// The home region of the Resource Discovery
	IpamResourceDiscoveryRegion pulumi.StringPtrInput
	// A boolean to identify if the Resource Discovery is the accounts default resource discovery
	IsDefault pulumi.BoolPtrInput
	// Determines which regions the Resource Discovery will enable IPAM features for usage and monitoring. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM Resource Discovery. You can only create VPCs from a pool whose locale matches the VPC's Region. You specify a region using the regionName parameter. **You must set your provider block region as an operating_region.**
	OperatingRegions VpcIpamResourceDiscoveryOperatingRegionArrayInput
	// The account ID for the account that manages the Resource Discovery
	OwnerId pulumi.StringPtrInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (VpcIpamResourceDiscoveryState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcIpamResourceDiscoveryState)(nil)).Elem()
}

type vpcIpamResourceDiscoveryArgs struct {
	// A description for the IPAM Resource Discovery.
	Description *string `pulumi:"description"`
	// Determines which regions the Resource Discovery will enable IPAM features for usage and monitoring. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM Resource Discovery. You can only create VPCs from a pool whose locale matches the VPC's Region. You specify a region using the regionName parameter. **You must set your provider block region as an operating_region.**
	OperatingRegions []VpcIpamResourceDiscoveryOperatingRegion `pulumi:"operatingRegions"`
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a VpcIpamResourceDiscovery resource.
type VpcIpamResourceDiscoveryArgs struct {
	// A description for the IPAM Resource Discovery.
	Description pulumi.StringPtrInput
	// Determines which regions the Resource Discovery will enable IPAM features for usage and monitoring. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM Resource Discovery. You can only create VPCs from a pool whose locale matches the VPC's Region. You specify a region using the regionName parameter. **You must set your provider block region as an operating_region.**
	OperatingRegions VpcIpamResourceDiscoveryOperatingRegionArrayInput
	// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (VpcIpamResourceDiscoveryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcIpamResourceDiscoveryArgs)(nil)).Elem()
}

type VpcIpamResourceDiscoveryInput interface {
	pulumi.Input

	ToVpcIpamResourceDiscoveryOutput() VpcIpamResourceDiscoveryOutput
	ToVpcIpamResourceDiscoveryOutputWithContext(ctx context.Context) VpcIpamResourceDiscoveryOutput
}

func (*VpcIpamResourceDiscovery) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcIpamResourceDiscovery)(nil)).Elem()
}

func (i *VpcIpamResourceDiscovery) ToVpcIpamResourceDiscoveryOutput() VpcIpamResourceDiscoveryOutput {
	return i.ToVpcIpamResourceDiscoveryOutputWithContext(context.Background())
}

func (i *VpcIpamResourceDiscovery) ToVpcIpamResourceDiscoveryOutputWithContext(ctx context.Context) VpcIpamResourceDiscoveryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpamResourceDiscoveryOutput)
}

// VpcIpamResourceDiscoveryArrayInput is an input type that accepts VpcIpamResourceDiscoveryArray and VpcIpamResourceDiscoveryArrayOutput values.
// You can construct a concrete instance of `VpcIpamResourceDiscoveryArrayInput` via:
//
//	VpcIpamResourceDiscoveryArray{ VpcIpamResourceDiscoveryArgs{...} }
type VpcIpamResourceDiscoveryArrayInput interface {
	pulumi.Input

	ToVpcIpamResourceDiscoveryArrayOutput() VpcIpamResourceDiscoveryArrayOutput
	ToVpcIpamResourceDiscoveryArrayOutputWithContext(context.Context) VpcIpamResourceDiscoveryArrayOutput
}

type VpcIpamResourceDiscoveryArray []VpcIpamResourceDiscoveryInput

func (VpcIpamResourceDiscoveryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcIpamResourceDiscovery)(nil)).Elem()
}

func (i VpcIpamResourceDiscoveryArray) ToVpcIpamResourceDiscoveryArrayOutput() VpcIpamResourceDiscoveryArrayOutput {
	return i.ToVpcIpamResourceDiscoveryArrayOutputWithContext(context.Background())
}

func (i VpcIpamResourceDiscoveryArray) ToVpcIpamResourceDiscoveryArrayOutputWithContext(ctx context.Context) VpcIpamResourceDiscoveryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpamResourceDiscoveryArrayOutput)
}

// VpcIpamResourceDiscoveryMapInput is an input type that accepts VpcIpamResourceDiscoveryMap and VpcIpamResourceDiscoveryMapOutput values.
// You can construct a concrete instance of `VpcIpamResourceDiscoveryMapInput` via:
//
//	VpcIpamResourceDiscoveryMap{ "key": VpcIpamResourceDiscoveryArgs{...} }
type VpcIpamResourceDiscoveryMapInput interface {
	pulumi.Input

	ToVpcIpamResourceDiscoveryMapOutput() VpcIpamResourceDiscoveryMapOutput
	ToVpcIpamResourceDiscoveryMapOutputWithContext(context.Context) VpcIpamResourceDiscoveryMapOutput
}

type VpcIpamResourceDiscoveryMap map[string]VpcIpamResourceDiscoveryInput

func (VpcIpamResourceDiscoveryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcIpamResourceDiscovery)(nil)).Elem()
}

func (i VpcIpamResourceDiscoveryMap) ToVpcIpamResourceDiscoveryMapOutput() VpcIpamResourceDiscoveryMapOutput {
	return i.ToVpcIpamResourceDiscoveryMapOutputWithContext(context.Background())
}

func (i VpcIpamResourceDiscoveryMap) ToVpcIpamResourceDiscoveryMapOutputWithContext(ctx context.Context) VpcIpamResourceDiscoveryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcIpamResourceDiscoveryMapOutput)
}

type VpcIpamResourceDiscoveryOutput struct{ *pulumi.OutputState }

func (VpcIpamResourceDiscoveryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcIpamResourceDiscovery)(nil)).Elem()
}

func (o VpcIpamResourceDiscoveryOutput) ToVpcIpamResourceDiscoveryOutput() VpcIpamResourceDiscoveryOutput {
	return o
}

func (o VpcIpamResourceDiscoveryOutput) ToVpcIpamResourceDiscoveryOutputWithContext(ctx context.Context) VpcIpamResourceDiscoveryOutput {
	return o
}

// Amazon Resource Name (ARN) of IPAM Resource Discovery
func (o VpcIpamResourceDiscoveryOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamResourceDiscovery) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// A description for the IPAM Resource Discovery.
func (o VpcIpamResourceDiscoveryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcIpamResourceDiscovery) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The home region of the Resource Discovery
func (o VpcIpamResourceDiscoveryOutput) IpamResourceDiscoveryRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamResourceDiscovery) pulumi.StringOutput { return v.IpamResourceDiscoveryRegion }).(pulumi.StringOutput)
}

// A boolean to identify if the Resource Discovery is the accounts default resource discovery
func (o VpcIpamResourceDiscoveryOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v *VpcIpamResourceDiscovery) pulumi.BoolOutput { return v.IsDefault }).(pulumi.BoolOutput)
}

// Determines which regions the Resource Discovery will enable IPAM features for usage and monitoring. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM Resource Discovery. You can only create VPCs from a pool whose locale matches the VPC's Region. You specify a region using the regionName parameter. **You must set your provider block region as an operating_region.**
func (o VpcIpamResourceDiscoveryOutput) OperatingRegions() VpcIpamResourceDiscoveryOperatingRegionArrayOutput {
	return o.ApplyT(func(v *VpcIpamResourceDiscovery) VpcIpamResourceDiscoveryOperatingRegionArrayOutput {
		return v.OperatingRegions
	}).(VpcIpamResourceDiscoveryOperatingRegionArrayOutput)
}

// The account ID for the account that manages the Resource Discovery
func (o VpcIpamResourceDiscoveryOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcIpamResourceDiscovery) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o VpcIpamResourceDiscoveryOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcIpamResourceDiscovery) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o VpcIpamResourceDiscoveryOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VpcIpamResourceDiscovery) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type VpcIpamResourceDiscoveryArrayOutput struct{ *pulumi.OutputState }

func (VpcIpamResourceDiscoveryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcIpamResourceDiscovery)(nil)).Elem()
}

func (o VpcIpamResourceDiscoveryArrayOutput) ToVpcIpamResourceDiscoveryArrayOutput() VpcIpamResourceDiscoveryArrayOutput {
	return o
}

func (o VpcIpamResourceDiscoveryArrayOutput) ToVpcIpamResourceDiscoveryArrayOutputWithContext(ctx context.Context) VpcIpamResourceDiscoveryArrayOutput {
	return o
}

func (o VpcIpamResourceDiscoveryArrayOutput) Index(i pulumi.IntInput) VpcIpamResourceDiscoveryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcIpamResourceDiscovery {
		return vs[0].([]*VpcIpamResourceDiscovery)[vs[1].(int)]
	}).(VpcIpamResourceDiscoveryOutput)
}

type VpcIpamResourceDiscoveryMapOutput struct{ *pulumi.OutputState }

func (VpcIpamResourceDiscoveryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcIpamResourceDiscovery)(nil)).Elem()
}

func (o VpcIpamResourceDiscoveryMapOutput) ToVpcIpamResourceDiscoveryMapOutput() VpcIpamResourceDiscoveryMapOutput {
	return o
}

func (o VpcIpamResourceDiscoveryMapOutput) ToVpcIpamResourceDiscoveryMapOutputWithContext(ctx context.Context) VpcIpamResourceDiscoveryMapOutput {
	return o
}

func (o VpcIpamResourceDiscoveryMapOutput) MapIndex(k pulumi.StringInput) VpcIpamResourceDiscoveryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcIpamResourceDiscovery {
		return vs[0].(map[string]*VpcIpamResourceDiscovery)[vs[1].(string)]
	}).(VpcIpamResourceDiscoveryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpamResourceDiscoveryInput)(nil)).Elem(), &VpcIpamResourceDiscovery{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpamResourceDiscoveryArrayInput)(nil)).Elem(), VpcIpamResourceDiscoveryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcIpamResourceDiscoveryMapInput)(nil)).Elem(), VpcIpamResourceDiscoveryMap{})
	pulumi.RegisterOutputType(VpcIpamResourceDiscoveryOutput{})
	pulumi.RegisterOutputType(VpcIpamResourceDiscoveryArrayOutput{})
	pulumi.RegisterOutputType(VpcIpamResourceDiscoveryMapOutput{})
}
