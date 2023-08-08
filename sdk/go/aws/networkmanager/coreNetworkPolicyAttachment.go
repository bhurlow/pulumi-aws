// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkmanager

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Core Network Policy Attachment resource. This puts a Core Network Policy to an existing Core Network and executes the change set, which deploys changes globally based on the policy submitted (Sets the policy to `LIVE`).
//
// > **NOTE:** Deleting this resource will not delete the current policy defined in this resource. Deleting this resource will also not revert the current `LIVE` policy to the previous version.
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
//			exampleCoreNetwork, err := networkmanager.NewCoreNetwork(ctx, "exampleCoreNetwork", &networkmanager.CoreNetworkArgs{
//				GlobalNetworkId: pulumi.Any(aws_networkmanager_global_network.Example.Id),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = networkmanager.NewCoreNetworkPolicyAttachment(ctx, "exampleCoreNetworkPolicyAttachment", &networkmanager.CoreNetworkPolicyAttachmentArgs{
//				CoreNetworkId:  exampleCoreNetwork.ID(),
//				PolicyDocument: pulumi.Any(data.Aws_networkmanager_core_network_policy_document.Example.Json),
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
// The example below illustrates the scenario where your policy document has static routes pointing to VPC attachments and you want to attach your VPCs to the core network before applying the desired policy document. Set the `createBasePolicy` argument of the `networkmanager.CoreNetwork` resource to `true` if your core network does not currently have any `LIVE` policies (e.g. this is the first `pulumi up` with the core network resource), since a `LIVE` policy is required before VPCs can be attached to the core network. Otherwise, if your core network already has a `LIVE` policy, you may exclude the `createBasePolicy` argument.
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
// terraform import {
//
//	to = aws_networkmanager_core_network_policy_attachment.example
//
//	id = "core-network-0d47f6t230mz46dy4" } Using `pulumi import`, import `aws_networkmanager_core_network_policy_attachment` using the core network ID. For exampleconsole % pulumi import aws_networkmanager_core_network_policy_attachment.example core-network-0d47f6t230mz46dy4
type CoreNetworkPolicyAttachment struct {
	pulumi.CustomResourceState

	// The ID of the core network that a policy will be attached to and made `LIVE`.
	CoreNetworkId pulumi.StringOutput `pulumi:"coreNetworkId"`
	// Policy document for creating a core network. Note that updating this argument will result in the new policy document version being set as the `LATEST` and `LIVE` policy document. Refer to the [Core network policies documentation](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) for more information.
	PolicyDocument pulumi.StringOutput `pulumi:"policyDocument"`
	// Current state of a core network.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewCoreNetworkPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewCoreNetworkPolicyAttachment(ctx *pulumi.Context,
	name string, args *CoreNetworkPolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*CoreNetworkPolicyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CoreNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'CoreNetworkId'")
	}
	if args.PolicyDocument == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDocument'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CoreNetworkPolicyAttachment
	err := ctx.RegisterResource("aws:networkmanager/coreNetworkPolicyAttachment:CoreNetworkPolicyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCoreNetworkPolicyAttachment gets an existing CoreNetworkPolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCoreNetworkPolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CoreNetworkPolicyAttachmentState, opts ...pulumi.ResourceOption) (*CoreNetworkPolicyAttachment, error) {
	var resource CoreNetworkPolicyAttachment
	err := ctx.ReadResource("aws:networkmanager/coreNetworkPolicyAttachment:CoreNetworkPolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CoreNetworkPolicyAttachment resources.
type coreNetworkPolicyAttachmentState struct {
	// The ID of the core network that a policy will be attached to and made `LIVE`.
	CoreNetworkId *string `pulumi:"coreNetworkId"`
	// Policy document for creating a core network. Note that updating this argument will result in the new policy document version being set as the `LATEST` and `LIVE` policy document. Refer to the [Core network policies documentation](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) for more information.
	PolicyDocument *string `pulumi:"policyDocument"`
	// Current state of a core network.
	State *string `pulumi:"state"`
}

type CoreNetworkPolicyAttachmentState struct {
	// The ID of the core network that a policy will be attached to and made `LIVE`.
	CoreNetworkId pulumi.StringPtrInput
	// Policy document for creating a core network. Note that updating this argument will result in the new policy document version being set as the `LATEST` and `LIVE` policy document. Refer to the [Core network policies documentation](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) for more information.
	PolicyDocument pulumi.StringPtrInput
	// Current state of a core network.
	State pulumi.StringPtrInput
}

func (CoreNetworkPolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*coreNetworkPolicyAttachmentState)(nil)).Elem()
}

type coreNetworkPolicyAttachmentArgs struct {
	// The ID of the core network that a policy will be attached to and made `LIVE`.
	CoreNetworkId string `pulumi:"coreNetworkId"`
	// Policy document for creating a core network. Note that updating this argument will result in the new policy document version being set as the `LATEST` and `LIVE` policy document. Refer to the [Core network policies documentation](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) for more information.
	PolicyDocument string `pulumi:"policyDocument"`
}

// The set of arguments for constructing a CoreNetworkPolicyAttachment resource.
type CoreNetworkPolicyAttachmentArgs struct {
	// The ID of the core network that a policy will be attached to and made `LIVE`.
	CoreNetworkId pulumi.StringInput
	// Policy document for creating a core network. Note that updating this argument will result in the new policy document version being set as the `LATEST` and `LIVE` policy document. Refer to the [Core network policies documentation](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) for more information.
	PolicyDocument pulumi.StringInput
}

func (CoreNetworkPolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*coreNetworkPolicyAttachmentArgs)(nil)).Elem()
}

type CoreNetworkPolicyAttachmentInput interface {
	pulumi.Input

	ToCoreNetworkPolicyAttachmentOutput() CoreNetworkPolicyAttachmentOutput
	ToCoreNetworkPolicyAttachmentOutputWithContext(ctx context.Context) CoreNetworkPolicyAttachmentOutput
}

func (*CoreNetworkPolicyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**CoreNetworkPolicyAttachment)(nil)).Elem()
}

func (i *CoreNetworkPolicyAttachment) ToCoreNetworkPolicyAttachmentOutput() CoreNetworkPolicyAttachmentOutput {
	return i.ToCoreNetworkPolicyAttachmentOutputWithContext(context.Background())
}

func (i *CoreNetworkPolicyAttachment) ToCoreNetworkPolicyAttachmentOutputWithContext(ctx context.Context) CoreNetworkPolicyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreNetworkPolicyAttachmentOutput)
}

// CoreNetworkPolicyAttachmentArrayInput is an input type that accepts CoreNetworkPolicyAttachmentArray and CoreNetworkPolicyAttachmentArrayOutput values.
// You can construct a concrete instance of `CoreNetworkPolicyAttachmentArrayInput` via:
//
//	CoreNetworkPolicyAttachmentArray{ CoreNetworkPolicyAttachmentArgs{...} }
type CoreNetworkPolicyAttachmentArrayInput interface {
	pulumi.Input

	ToCoreNetworkPolicyAttachmentArrayOutput() CoreNetworkPolicyAttachmentArrayOutput
	ToCoreNetworkPolicyAttachmentArrayOutputWithContext(context.Context) CoreNetworkPolicyAttachmentArrayOutput
}

type CoreNetworkPolicyAttachmentArray []CoreNetworkPolicyAttachmentInput

func (CoreNetworkPolicyAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CoreNetworkPolicyAttachment)(nil)).Elem()
}

func (i CoreNetworkPolicyAttachmentArray) ToCoreNetworkPolicyAttachmentArrayOutput() CoreNetworkPolicyAttachmentArrayOutput {
	return i.ToCoreNetworkPolicyAttachmentArrayOutputWithContext(context.Background())
}

func (i CoreNetworkPolicyAttachmentArray) ToCoreNetworkPolicyAttachmentArrayOutputWithContext(ctx context.Context) CoreNetworkPolicyAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreNetworkPolicyAttachmentArrayOutput)
}

// CoreNetworkPolicyAttachmentMapInput is an input type that accepts CoreNetworkPolicyAttachmentMap and CoreNetworkPolicyAttachmentMapOutput values.
// You can construct a concrete instance of `CoreNetworkPolicyAttachmentMapInput` via:
//
//	CoreNetworkPolicyAttachmentMap{ "key": CoreNetworkPolicyAttachmentArgs{...} }
type CoreNetworkPolicyAttachmentMapInput interface {
	pulumi.Input

	ToCoreNetworkPolicyAttachmentMapOutput() CoreNetworkPolicyAttachmentMapOutput
	ToCoreNetworkPolicyAttachmentMapOutputWithContext(context.Context) CoreNetworkPolicyAttachmentMapOutput
}

type CoreNetworkPolicyAttachmentMap map[string]CoreNetworkPolicyAttachmentInput

func (CoreNetworkPolicyAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CoreNetworkPolicyAttachment)(nil)).Elem()
}

func (i CoreNetworkPolicyAttachmentMap) ToCoreNetworkPolicyAttachmentMapOutput() CoreNetworkPolicyAttachmentMapOutput {
	return i.ToCoreNetworkPolicyAttachmentMapOutputWithContext(context.Background())
}

func (i CoreNetworkPolicyAttachmentMap) ToCoreNetworkPolicyAttachmentMapOutputWithContext(ctx context.Context) CoreNetworkPolicyAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreNetworkPolicyAttachmentMapOutput)
}

type CoreNetworkPolicyAttachmentOutput struct{ *pulumi.OutputState }

func (CoreNetworkPolicyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CoreNetworkPolicyAttachment)(nil)).Elem()
}

func (o CoreNetworkPolicyAttachmentOutput) ToCoreNetworkPolicyAttachmentOutput() CoreNetworkPolicyAttachmentOutput {
	return o
}

func (o CoreNetworkPolicyAttachmentOutput) ToCoreNetworkPolicyAttachmentOutputWithContext(ctx context.Context) CoreNetworkPolicyAttachmentOutput {
	return o
}

// The ID of the core network that a policy will be attached to and made `LIVE`.
func (o CoreNetworkPolicyAttachmentOutput) CoreNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *CoreNetworkPolicyAttachment) pulumi.StringOutput { return v.CoreNetworkId }).(pulumi.StringOutput)
}

// Policy document for creating a core network. Note that updating this argument will result in the new policy document version being set as the `LATEST` and `LIVE` policy document. Refer to the [Core network policies documentation](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) for more information.
func (o CoreNetworkPolicyAttachmentOutput) PolicyDocument() pulumi.StringOutput {
	return o.ApplyT(func(v *CoreNetworkPolicyAttachment) pulumi.StringOutput { return v.PolicyDocument }).(pulumi.StringOutput)
}

// Current state of a core network.
func (o CoreNetworkPolicyAttachmentOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *CoreNetworkPolicyAttachment) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

type CoreNetworkPolicyAttachmentArrayOutput struct{ *pulumi.OutputState }

func (CoreNetworkPolicyAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CoreNetworkPolicyAttachment)(nil)).Elem()
}

func (o CoreNetworkPolicyAttachmentArrayOutput) ToCoreNetworkPolicyAttachmentArrayOutput() CoreNetworkPolicyAttachmentArrayOutput {
	return o
}

func (o CoreNetworkPolicyAttachmentArrayOutput) ToCoreNetworkPolicyAttachmentArrayOutputWithContext(ctx context.Context) CoreNetworkPolicyAttachmentArrayOutput {
	return o
}

func (o CoreNetworkPolicyAttachmentArrayOutput) Index(i pulumi.IntInput) CoreNetworkPolicyAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CoreNetworkPolicyAttachment {
		return vs[0].([]*CoreNetworkPolicyAttachment)[vs[1].(int)]
	}).(CoreNetworkPolicyAttachmentOutput)
}

type CoreNetworkPolicyAttachmentMapOutput struct{ *pulumi.OutputState }

func (CoreNetworkPolicyAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CoreNetworkPolicyAttachment)(nil)).Elem()
}

func (o CoreNetworkPolicyAttachmentMapOutput) ToCoreNetworkPolicyAttachmentMapOutput() CoreNetworkPolicyAttachmentMapOutput {
	return o
}

func (o CoreNetworkPolicyAttachmentMapOutput) ToCoreNetworkPolicyAttachmentMapOutputWithContext(ctx context.Context) CoreNetworkPolicyAttachmentMapOutput {
	return o
}

func (o CoreNetworkPolicyAttachmentMapOutput) MapIndex(k pulumi.StringInput) CoreNetworkPolicyAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CoreNetworkPolicyAttachment {
		return vs[0].(map[string]*CoreNetworkPolicyAttachment)[vs[1].(string)]
	}).(CoreNetworkPolicyAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CoreNetworkPolicyAttachmentInput)(nil)).Elem(), &CoreNetworkPolicyAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*CoreNetworkPolicyAttachmentArrayInput)(nil)).Elem(), CoreNetworkPolicyAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CoreNetworkPolicyAttachmentMapInput)(nil)).Elem(), CoreNetworkPolicyAttachmentMap{})
	pulumi.RegisterOutputType(CoreNetworkPolicyAttachmentOutput{})
	pulumi.RegisterOutputType(CoreNetworkPolicyAttachmentArrayOutput{})
	pulumi.RegisterOutputType(CoreNetworkPolicyAttachmentMapOutput{})
}
