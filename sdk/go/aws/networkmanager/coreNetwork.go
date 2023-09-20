// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkmanager

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Provides a core network resource.
//
// ## Example Usage
// ### Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkmanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networkmanager.NewCoreNetwork(ctx, "example", &networkmanager.CoreNetworkArgs{
//				GlobalNetworkId: pulumi.Any(aws_networkmanager_global_network.Example.Id),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### With description
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkmanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networkmanager.NewCoreNetwork(ctx, "example", &networkmanager.CoreNetworkArgs{
//				GlobalNetworkId: pulumi.Any(aws_networkmanager_global_network.Example.Id),
//				Description:     pulumi.String("example"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### With tags
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkmanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networkmanager.NewCoreNetwork(ctx, "example", &networkmanager.CoreNetworkArgs{
//				GlobalNetworkId: pulumi.Any(aws_networkmanager_global_network.Example.Id),
//				Tags: pulumi.StringMap{
//					"hello": pulumi.String("world"),
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
// ### With VPC Attachment (Single Region)
//
// The example below illustrates the scenario where your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Set the `createBasePolicy` argument to `true` if your core network does not currently have any `LIVE` policies (e.g. this is the first `pulumi up` with the core network resource), since a `LIVE` policy is required before VPCs can be attached to the core network. Otherwise, if your core network already has a `LIVE` policy, you may exclude the `createBasePolicy` argument.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkmanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleGlobalNetwork, err := networkmanager.NewGlobalNetwork(ctx, "exampleGlobalNetwork", nil)
//			if err != nil {
//				return err
//			}
//			exampleCoreNetwork, err := networkmanager.NewCoreNetwork(ctx, "exampleCoreNetwork", &networkmanager.CoreNetworkArgs{
//				GlobalNetworkId:  exampleGlobalNetwork.ID(),
//				CreateBasePolicy: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			var splat0 []interface{}
//			for _, val0 := range aws_subnet.Example {
//				splat0 = append(splat0, val0.Arn)
//			}
//			exampleVpcAttachment, err := networkmanager.NewVpcAttachment(ctx, "exampleVpcAttachment", &networkmanager.VpcAttachmentArgs{
//				CoreNetworkId: exampleCoreNetwork.ID(),
//				SubnetArns:    toPulumiAnyArray(splat0),
//				VpcArn:        pulumi.Any(aws_vpc.Example.Arn),
//			})
//			if err != nil {
//				return err
//			}
//			exampleCoreNetworkPolicyDocument := networkmanager.GetCoreNetworkPolicyDocumentOutput(ctx, networkmanager.GetCoreNetworkPolicyDocumentOutputArgs{
//				CoreNetworkConfigurations: networkmanager.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationArray{
//					&networkmanager.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationArgs{
//						AsnRanges: pulumi.StringArray{
//							pulumi.String("65022-65534"),
//						},
//						EdgeLocations: networkmanager.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocationArray{
//							&networkmanager.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocationArgs{
//								Location: pulumi.String("us-west-2"),
//							},
//						},
//					},
//				},
//				Segments: networkmanager.GetCoreNetworkPolicyDocumentSegmentArray{
//					&networkmanager.GetCoreNetworkPolicyDocumentSegmentArgs{
//						Name: pulumi.String("segment"),
//					},
//				},
//				SegmentActions: networkmanager.GetCoreNetworkPolicyDocumentSegmentActionArray{
//					&networkmanager.GetCoreNetworkPolicyDocumentSegmentActionArgs{
//						Action:  pulumi.String("create-route"),
//						Segment: pulumi.String("segment"),
//						DestinationCidrBlocks: pulumi.StringArray{
//							pulumi.String("0.0.0.0/0"),
//						},
//						Destinations: pulumi.StringArray{
//							exampleVpcAttachment.ID(),
//						},
//					},
//				},
//			}, nil)
//			_, err = networkmanager.NewCoreNetworkPolicyAttachment(ctx, "exampleCoreNetworkPolicyAttachment", &networkmanager.CoreNetworkPolicyAttachmentArgs{
//				CoreNetworkId: exampleCoreNetwork.ID(),
//				PolicyDocument: exampleCoreNetworkPolicyDocument.ApplyT(func(exampleCoreNetworkPolicyDocument networkmanager.GetCoreNetworkPolicyDocumentResult) (*string, error) {
//					return &exampleCoreNetworkPolicyDocument.Json, nil
//				}).(pulumi.StringPtrOutput),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
//	func toPulumiAnyArray(arr []Any) pulumi.AnyArray {
//		var pulumiArr pulumi.AnyArray
//		for _, v := range arr {
//			pulumiArr = append(pulumiArr, pulumi.Any(v))
//		}
//		return pulumiArr
//	}
//
// ```
// ### With VPC Attachment (Multi-Region)
//
// The example below illustrates the scenario where your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Set the `createBasePolicy` argument of the `networkmanager.CoreNetwork` resource to `true` if your core network does not currently have any `LIVE` policies (e.g. this is the first `pulumi up` with the core network resource), since a `LIVE` policy is required before VPCs can be attached to the core network. Otherwise, if your core network already has a `LIVE` policy, you may exclude the `createBasePolicy` argument. For multi-region in a core network that does not yet have a `LIVE` policy, pass a list of regions to the `networkmanager.CoreNetwork` `basePolicyRegions` argument. In the example below, `us-west-2` and `us-east-1` are specified in the base policy.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkmanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleGlobalNetwork, err := networkmanager.NewGlobalNetwork(ctx, "exampleGlobalNetwork", nil)
//			if err != nil {
//				return err
//			}
//			exampleCoreNetwork, err := networkmanager.NewCoreNetwork(ctx, "exampleCoreNetwork", &networkmanager.CoreNetworkArgs{
//				GlobalNetworkId: exampleGlobalNetwork.ID(),
//				BasePolicyRegions: pulumi.StringArray{
//					pulumi.String("us-west-2"),
//					pulumi.String("us-east-1"),
//				},
//				CreateBasePolicy: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			var splat0 []interface{}
//			for _, val0 := range aws_subnet.Example_us_west_2 {
//				splat0 = append(splat0, val0.Arn)
//			}
//			exampleUsWest2, err := networkmanager.NewVpcAttachment(ctx, "exampleUsWest2", &networkmanager.VpcAttachmentArgs{
//				CoreNetworkId: exampleCoreNetwork.ID(),
//				SubnetArns:    toPulumiAnyArray(splat0),
//				VpcArn:        pulumi.Any(aws_vpc.Example_us_west_2.Arn),
//			})
//			if err != nil {
//				return err
//			}
//			var splat1 []interface{}
//			for _, val0 := range aws_subnet.Example_us_east_1 {
//				splat1 = append(splat1, val0.Arn)
//			}
//			exampleUsEast1, err := networkmanager.NewVpcAttachment(ctx, "exampleUsEast1", &networkmanager.VpcAttachmentArgs{
//				CoreNetworkId: exampleCoreNetwork.ID(),
//				SubnetArns:    toPulumiAnyArray(splat1),
//				VpcArn:        pulumi.Any(aws_vpc.Example_us_east_1.Arn),
//			}, pulumi.Provider("alternate"))
//			if err != nil {
//				return err
//			}
//			exampleCoreNetworkPolicyDocument := networkmanager.GetCoreNetworkPolicyDocumentOutput(ctx, networkmanager.GetCoreNetworkPolicyDocumentOutputArgs{
//				CoreNetworkConfigurations: networkmanager.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationArray{
//					&networkmanager.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationArgs{
//						AsnRanges: pulumi.StringArray{
//							pulumi.String("65022-65534"),
//						},
//						EdgeLocations: networkmanager.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocationArray{
//							&networkmanager.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocationArgs{
//								Location: pulumi.String("us-west-2"),
//							},
//							&networkmanager.GetCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocationArgs{
//								Location: pulumi.String("us-east-1"),
//							},
//						},
//					},
//				},
//				Segments: networkmanager.GetCoreNetworkPolicyDocumentSegmentArray{
//					&networkmanager.GetCoreNetworkPolicyDocumentSegmentArgs{
//						Name: pulumi.String("segment"),
//					},
//					&networkmanager.GetCoreNetworkPolicyDocumentSegmentArgs{
//						Name: pulumi.String("segment2"),
//					},
//				},
//				SegmentActions: networkmanager.GetCoreNetworkPolicyDocumentSegmentActionArray{
//					&networkmanager.GetCoreNetworkPolicyDocumentSegmentActionArgs{
//						Action:  pulumi.String("create-route"),
//						Segment: pulumi.String("segment"),
//						DestinationCidrBlocks: pulumi.StringArray{
//							pulumi.String("10.0.0.0/16"),
//						},
//						Destinations: pulumi.StringArray{
//							exampleUsWest2.ID(),
//						},
//					},
//					&networkmanager.GetCoreNetworkPolicyDocumentSegmentActionArgs{
//						Action:  pulumi.String("create-route"),
//						Segment: pulumi.String("segment"),
//						DestinationCidrBlocks: pulumi.StringArray{
//							pulumi.String("10.1.0.0/16"),
//						},
//						Destinations: pulumi.StringArray{
//							exampleUsEast1.ID(),
//						},
//					},
//				},
//			}, nil)
//			_, err = networkmanager.NewCoreNetworkPolicyAttachment(ctx, "exampleCoreNetworkPolicyAttachment", &networkmanager.CoreNetworkPolicyAttachmentArgs{
//				CoreNetworkId: exampleCoreNetwork.ID(),
//				PolicyDocument: exampleCoreNetworkPolicyDocument.ApplyT(func(exampleCoreNetworkPolicyDocument networkmanager.GetCoreNetworkPolicyDocumentResult) (*string, error) {
//					return &exampleCoreNetworkPolicyDocument.Json, nil
//				}).(pulumi.StringPtrOutput),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
//	func toPulumiAnyArray(arr []Any) pulumi.AnyArray {
//		var pulumiArr pulumi.AnyArray
//		for _, v := range arr {
//			pulumiArr = append(pulumiArr, pulumi.Any(v))
//		}
//		return pulumiArr
//	}
//
// ```
//
// ## Import
//
// Using `pulumi import`, import `aws_networkmanager_core_network` using the core network ID. For example:
//
// ```sh
//
//	$ pulumi import aws:networkmanager/coreNetwork:CoreNetwork example core-network-0d47f6t230mz46dy4
//
// ```
type CoreNetwork struct {
	pulumi.CustomResourceState

	// Core Network Amazon Resource Name (ARN).
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The base policy created by setting the `createBasePolicy` argument to `true` requires a region to be set in the `edge-locations`, `location` key. If `basePolicyRegion` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
	//
	// Deprecated: Use the base_policy_regions argument instead. This argument will be removed in the next major version of the provider.
	BasePolicyRegion pulumi.StringPtrOutput `pulumi:"basePolicyRegion"`
	// A list of regions to add to the base policy. The base policy created by setting the `createBasePolicy` argument to `true` requires one or more regions to be set in the `edge-locations`, `location` key. If `basePolicyRegions` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
	BasePolicyRegions pulumi.StringArrayOutput `pulumi:"basePolicyRegions"`
	// Specifies whether to create a base policy when a core network is created or updated. A base policy is created and set to `LIVE` to allow attachments to the core network (e.g. VPC Attachments) before applying a policy document provided using the `networkmanager.CoreNetworkPolicyAttachment` resource. This base policy is needed if your core network does not have any `LIVE` policies and your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Valid values are `true` or `false`. An example of this Pulumi snippet can be found above for VPC Attachment in a single region and for VPC Attachment multi-region. An example base policy is shown below. This base policy is overridden with the policy that you specify in the `networkmanager.CoreNetworkPolicyAttachment` resource.
	//
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	CreateBasePolicy pulumi.BoolPtrOutput `pulumi:"createBasePolicy"`
	// Timestamp when a core network was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Description of the Core Network.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// One or more blocks detailing the edges within a core network. Detailed below.
	Edges CoreNetworkEdgeArrayOutput `pulumi:"edges"`
	// The ID of the global network that a core network will be a part of.
	GlobalNetworkId pulumi.StringOutput `pulumi:"globalNetworkId"`
	// One or more blocks detailing the segments within a core network. Detailed below.
	Segments CoreNetworkSegmentArrayOutput `pulumi:"segments"`
	// Current state of a core network.
	State pulumi.StringOutput `pulumi:"state"`
	// Key-value tags for the Core Network. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewCoreNetwork registers a new resource with the given unique name, arguments, and options.
func NewCoreNetwork(ctx *pulumi.Context,
	name string, args *CoreNetworkArgs, opts ...pulumi.ResourceOption) (*CoreNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GlobalNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'GlobalNetworkId'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CoreNetwork
	err := ctx.RegisterResource("aws:networkmanager/coreNetwork:CoreNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCoreNetwork gets an existing CoreNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCoreNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CoreNetworkState, opts ...pulumi.ResourceOption) (*CoreNetwork, error) {
	var resource CoreNetwork
	err := ctx.ReadResource("aws:networkmanager/coreNetwork:CoreNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CoreNetwork resources.
type coreNetworkState struct {
	// Core Network Amazon Resource Name (ARN).
	Arn *string `pulumi:"arn"`
	// The base policy created by setting the `createBasePolicy` argument to `true` requires a region to be set in the `edge-locations`, `location` key. If `basePolicyRegion` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
	//
	// Deprecated: Use the base_policy_regions argument instead. This argument will be removed in the next major version of the provider.
	BasePolicyRegion *string `pulumi:"basePolicyRegion"`
	// A list of regions to add to the base policy. The base policy created by setting the `createBasePolicy` argument to `true` requires one or more regions to be set in the `edge-locations`, `location` key. If `basePolicyRegions` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
	BasePolicyRegions []string `pulumi:"basePolicyRegions"`
	// Specifies whether to create a base policy when a core network is created or updated. A base policy is created and set to `LIVE` to allow attachments to the core network (e.g. VPC Attachments) before applying a policy document provided using the `networkmanager.CoreNetworkPolicyAttachment` resource. This base policy is needed if your core network does not have any `LIVE` policies and your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Valid values are `true` or `false`. An example of this Pulumi snippet can be found above for VPC Attachment in a single region and for VPC Attachment multi-region. An example base policy is shown below. This base policy is overridden with the policy that you specify in the `networkmanager.CoreNetworkPolicyAttachment` resource.
	//
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	CreateBasePolicy *bool `pulumi:"createBasePolicy"`
	// Timestamp when a core network was created.
	CreatedAt *string `pulumi:"createdAt"`
	// Description of the Core Network.
	Description *string `pulumi:"description"`
	// One or more blocks detailing the edges within a core network. Detailed below.
	Edges []CoreNetworkEdge `pulumi:"edges"`
	// The ID of the global network that a core network will be a part of.
	GlobalNetworkId *string `pulumi:"globalNetworkId"`
	// One or more blocks detailing the segments within a core network. Detailed below.
	Segments []CoreNetworkSegment `pulumi:"segments"`
	// Current state of a core network.
	State *string `pulumi:"state"`
	// Key-value tags for the Core Network. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type CoreNetworkState struct {
	// Core Network Amazon Resource Name (ARN).
	Arn pulumi.StringPtrInput
	// The base policy created by setting the `createBasePolicy` argument to `true` requires a region to be set in the `edge-locations`, `location` key. If `basePolicyRegion` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
	//
	// Deprecated: Use the base_policy_regions argument instead. This argument will be removed in the next major version of the provider.
	BasePolicyRegion pulumi.StringPtrInput
	// A list of regions to add to the base policy. The base policy created by setting the `createBasePolicy` argument to `true` requires one or more regions to be set in the `edge-locations`, `location` key. If `basePolicyRegions` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
	BasePolicyRegions pulumi.StringArrayInput
	// Specifies whether to create a base policy when a core network is created or updated. A base policy is created and set to `LIVE` to allow attachments to the core network (e.g. VPC Attachments) before applying a policy document provided using the `networkmanager.CoreNetworkPolicyAttachment` resource. This base policy is needed if your core network does not have any `LIVE` policies and your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Valid values are `true` or `false`. An example of this Pulumi snippet can be found above for VPC Attachment in a single region and for VPC Attachment multi-region. An example base policy is shown below. This base policy is overridden with the policy that you specify in the `networkmanager.CoreNetworkPolicyAttachment` resource.
	//
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	CreateBasePolicy pulumi.BoolPtrInput
	// Timestamp when a core network was created.
	CreatedAt pulumi.StringPtrInput
	// Description of the Core Network.
	Description pulumi.StringPtrInput
	// One or more blocks detailing the edges within a core network. Detailed below.
	Edges CoreNetworkEdgeArrayInput
	// The ID of the global network that a core network will be a part of.
	GlobalNetworkId pulumi.StringPtrInput
	// One or more blocks detailing the segments within a core network. Detailed below.
	Segments CoreNetworkSegmentArrayInput
	// Current state of a core network.
	State pulumi.StringPtrInput
	// Key-value tags for the Core Network. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (CoreNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*coreNetworkState)(nil)).Elem()
}

type coreNetworkArgs struct {
	// The base policy created by setting the `createBasePolicy` argument to `true` requires a region to be set in the `edge-locations`, `location` key. If `basePolicyRegion` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
	//
	// Deprecated: Use the base_policy_regions argument instead. This argument will be removed in the next major version of the provider.
	BasePolicyRegion *string `pulumi:"basePolicyRegion"`
	// A list of regions to add to the base policy. The base policy created by setting the `createBasePolicy` argument to `true` requires one or more regions to be set in the `edge-locations`, `location` key. If `basePolicyRegions` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
	BasePolicyRegions []string `pulumi:"basePolicyRegions"`
	// Specifies whether to create a base policy when a core network is created or updated. A base policy is created and set to `LIVE` to allow attachments to the core network (e.g. VPC Attachments) before applying a policy document provided using the `networkmanager.CoreNetworkPolicyAttachment` resource. This base policy is needed if your core network does not have any `LIVE` policies and your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Valid values are `true` or `false`. An example of this Pulumi snippet can be found above for VPC Attachment in a single region and for VPC Attachment multi-region. An example base policy is shown below. This base policy is overridden with the policy that you specify in the `networkmanager.CoreNetworkPolicyAttachment` resource.
	//
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	CreateBasePolicy *bool `pulumi:"createBasePolicy"`
	// Description of the Core Network.
	Description *string `pulumi:"description"`
	// The ID of the global network that a core network will be a part of.
	GlobalNetworkId string `pulumi:"globalNetworkId"`
	// Key-value tags for the Core Network. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CoreNetwork resource.
type CoreNetworkArgs struct {
	// The base policy created by setting the `createBasePolicy` argument to `true` requires a region to be set in the `edge-locations`, `location` key. If `basePolicyRegion` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
	//
	// Deprecated: Use the base_policy_regions argument instead. This argument will be removed in the next major version of the provider.
	BasePolicyRegion pulumi.StringPtrInput
	// A list of regions to add to the base policy. The base policy created by setting the `createBasePolicy` argument to `true` requires one or more regions to be set in the `edge-locations`, `location` key. If `basePolicyRegions` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
	BasePolicyRegions pulumi.StringArrayInput
	// Specifies whether to create a base policy when a core network is created or updated. A base policy is created and set to `LIVE` to allow attachments to the core network (e.g. VPC Attachments) before applying a policy document provided using the `networkmanager.CoreNetworkPolicyAttachment` resource. This base policy is needed if your core network does not have any `LIVE` policies and your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Valid values are `true` or `false`. An example of this Pulumi snippet can be found above for VPC Attachment in a single region and for VPC Attachment multi-region. An example base policy is shown below. This base policy is overridden with the policy that you specify in the `networkmanager.CoreNetworkPolicyAttachment` resource.
	//
	// ```go
	// package main
	//
	// import (
	// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// )
	//
	// func main() {
	// 	pulumi.Run(func(ctx *pulumi.Context) error {
	// 		return nil
	// 	})
	// }
	// ```
	CreateBasePolicy pulumi.BoolPtrInput
	// Description of the Core Network.
	Description pulumi.StringPtrInput
	// The ID of the global network that a core network will be a part of.
	GlobalNetworkId pulumi.StringInput
	// Key-value tags for the Core Network. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (CoreNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*coreNetworkArgs)(nil)).Elem()
}

type CoreNetworkInput interface {
	pulumi.Input

	ToCoreNetworkOutput() CoreNetworkOutput
	ToCoreNetworkOutputWithContext(ctx context.Context) CoreNetworkOutput
}

func (*CoreNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**CoreNetwork)(nil)).Elem()
}

func (i *CoreNetwork) ToCoreNetworkOutput() CoreNetworkOutput {
	return i.ToCoreNetworkOutputWithContext(context.Background())
}

func (i *CoreNetwork) ToCoreNetworkOutputWithContext(ctx context.Context) CoreNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreNetworkOutput)
}

func (i *CoreNetwork) ToOutput(ctx context.Context) pulumix.Output[*CoreNetwork] {
	return pulumix.Output[*CoreNetwork]{
		OutputState: i.ToCoreNetworkOutputWithContext(ctx).OutputState,
	}
}

// CoreNetworkArrayInput is an input type that accepts CoreNetworkArray and CoreNetworkArrayOutput values.
// You can construct a concrete instance of `CoreNetworkArrayInput` via:
//
//	CoreNetworkArray{ CoreNetworkArgs{...} }
type CoreNetworkArrayInput interface {
	pulumi.Input

	ToCoreNetworkArrayOutput() CoreNetworkArrayOutput
	ToCoreNetworkArrayOutputWithContext(context.Context) CoreNetworkArrayOutput
}

type CoreNetworkArray []CoreNetworkInput

func (CoreNetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CoreNetwork)(nil)).Elem()
}

func (i CoreNetworkArray) ToCoreNetworkArrayOutput() CoreNetworkArrayOutput {
	return i.ToCoreNetworkArrayOutputWithContext(context.Background())
}

func (i CoreNetworkArray) ToCoreNetworkArrayOutputWithContext(ctx context.Context) CoreNetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreNetworkArrayOutput)
}

func (i CoreNetworkArray) ToOutput(ctx context.Context) pulumix.Output[[]*CoreNetwork] {
	return pulumix.Output[[]*CoreNetwork]{
		OutputState: i.ToCoreNetworkArrayOutputWithContext(ctx).OutputState,
	}
}

// CoreNetworkMapInput is an input type that accepts CoreNetworkMap and CoreNetworkMapOutput values.
// You can construct a concrete instance of `CoreNetworkMapInput` via:
//
//	CoreNetworkMap{ "key": CoreNetworkArgs{...} }
type CoreNetworkMapInput interface {
	pulumi.Input

	ToCoreNetworkMapOutput() CoreNetworkMapOutput
	ToCoreNetworkMapOutputWithContext(context.Context) CoreNetworkMapOutput
}

type CoreNetworkMap map[string]CoreNetworkInput

func (CoreNetworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CoreNetwork)(nil)).Elem()
}

func (i CoreNetworkMap) ToCoreNetworkMapOutput() CoreNetworkMapOutput {
	return i.ToCoreNetworkMapOutputWithContext(context.Background())
}

func (i CoreNetworkMap) ToCoreNetworkMapOutputWithContext(ctx context.Context) CoreNetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreNetworkMapOutput)
}

func (i CoreNetworkMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*CoreNetwork] {
	return pulumix.Output[map[string]*CoreNetwork]{
		OutputState: i.ToCoreNetworkMapOutputWithContext(ctx).OutputState,
	}
}

type CoreNetworkOutput struct{ *pulumi.OutputState }

func (CoreNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CoreNetwork)(nil)).Elem()
}

func (o CoreNetworkOutput) ToCoreNetworkOutput() CoreNetworkOutput {
	return o
}

func (o CoreNetworkOutput) ToCoreNetworkOutputWithContext(ctx context.Context) CoreNetworkOutput {
	return o
}

func (o CoreNetworkOutput) ToOutput(ctx context.Context) pulumix.Output[*CoreNetwork] {
	return pulumix.Output[*CoreNetwork]{
		OutputState: o.OutputState,
	}
}

// Core Network Amazon Resource Name (ARN).
func (o CoreNetworkOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *CoreNetwork) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The base policy created by setting the `createBasePolicy` argument to `true` requires a region to be set in the `edge-locations`, `location` key. If `basePolicyRegion` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
//
// Deprecated: Use the base_policy_regions argument instead. This argument will be removed in the next major version of the provider.
func (o CoreNetworkOutput) BasePolicyRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CoreNetwork) pulumi.StringPtrOutput { return v.BasePolicyRegion }).(pulumi.StringPtrOutput)
}

// A list of regions to add to the base policy. The base policy created by setting the `createBasePolicy` argument to `true` requires one or more regions to be set in the `edge-locations`, `location` key. If `basePolicyRegions` is not specified, the region used in the base policy defaults to the region specified in the `provider` block.
func (o CoreNetworkOutput) BasePolicyRegions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CoreNetwork) pulumi.StringArrayOutput { return v.BasePolicyRegions }).(pulumi.StringArrayOutput)
}

// Specifies whether to create a base policy when a core network is created or updated. A base policy is created and set to `LIVE` to allow attachments to the core network (e.g. VPC Attachments) before applying a policy document provided using the `networkmanager.CoreNetworkPolicyAttachment` resource. This base policy is needed if your core network does not have any `LIVE` policies and your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Valid values are `true` or `false`. An example of this Pulumi snippet can be found above for VPC Attachment in a single region and for VPC Attachment multi-region. An example base policy is shown below. This base policy is overridden with the policy that you specify in the `networkmanager.CoreNetworkPolicyAttachment` resource.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			return nil
//		})
//	}
//
// ```
func (o CoreNetworkOutput) CreateBasePolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CoreNetwork) pulumi.BoolPtrOutput { return v.CreateBasePolicy }).(pulumi.BoolPtrOutput)
}

// Timestamp when a core network was created.
func (o CoreNetworkOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CoreNetwork) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Description of the Core Network.
func (o CoreNetworkOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CoreNetwork) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// One or more blocks detailing the edges within a core network. Detailed below.
func (o CoreNetworkOutput) Edges() CoreNetworkEdgeArrayOutput {
	return o.ApplyT(func(v *CoreNetwork) CoreNetworkEdgeArrayOutput { return v.Edges }).(CoreNetworkEdgeArrayOutput)
}

// The ID of the global network that a core network will be a part of.
func (o CoreNetworkOutput) GlobalNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *CoreNetwork) pulumi.StringOutput { return v.GlobalNetworkId }).(pulumi.StringOutput)
}

// One or more blocks detailing the segments within a core network. Detailed below.
func (o CoreNetworkOutput) Segments() CoreNetworkSegmentArrayOutput {
	return o.ApplyT(func(v *CoreNetwork) CoreNetworkSegmentArrayOutput { return v.Segments }).(CoreNetworkSegmentArrayOutput)
}

// Current state of a core network.
func (o CoreNetworkOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *CoreNetwork) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Key-value tags for the Core Network. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o CoreNetworkOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CoreNetwork) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o CoreNetworkOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CoreNetwork) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type CoreNetworkArrayOutput struct{ *pulumi.OutputState }

func (CoreNetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CoreNetwork)(nil)).Elem()
}

func (o CoreNetworkArrayOutput) ToCoreNetworkArrayOutput() CoreNetworkArrayOutput {
	return o
}

func (o CoreNetworkArrayOutput) ToCoreNetworkArrayOutputWithContext(ctx context.Context) CoreNetworkArrayOutput {
	return o
}

func (o CoreNetworkArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*CoreNetwork] {
	return pulumix.Output[[]*CoreNetwork]{
		OutputState: o.OutputState,
	}
}

func (o CoreNetworkArrayOutput) Index(i pulumi.IntInput) CoreNetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CoreNetwork {
		return vs[0].([]*CoreNetwork)[vs[1].(int)]
	}).(CoreNetworkOutput)
}

type CoreNetworkMapOutput struct{ *pulumi.OutputState }

func (CoreNetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CoreNetwork)(nil)).Elem()
}

func (o CoreNetworkMapOutput) ToCoreNetworkMapOutput() CoreNetworkMapOutput {
	return o
}

func (o CoreNetworkMapOutput) ToCoreNetworkMapOutputWithContext(ctx context.Context) CoreNetworkMapOutput {
	return o
}

func (o CoreNetworkMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*CoreNetwork] {
	return pulumix.Output[map[string]*CoreNetwork]{
		OutputState: o.OutputState,
	}
}

func (o CoreNetworkMapOutput) MapIndex(k pulumi.StringInput) CoreNetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CoreNetwork {
		return vs[0].(map[string]*CoreNetwork)[vs[1].(string)]
	}).(CoreNetworkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CoreNetworkInput)(nil)).Elem(), &CoreNetwork{})
	pulumi.RegisterInputType(reflect.TypeOf((*CoreNetworkArrayInput)(nil)).Elem(), CoreNetworkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CoreNetworkMapInput)(nil)).Elem(), CoreNetworkMap{})
	pulumi.RegisterOutputType(CoreNetworkOutput{})
	pulumi.RegisterOutputType(CoreNetworkArrayOutput{})
	pulumi.RegisterOutputType(CoreNetworkMapOutput{})
}
