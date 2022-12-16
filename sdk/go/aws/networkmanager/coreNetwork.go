// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkmanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/networkmanager"
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
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/networkmanager"
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
// ### With policy document
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/networkmanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networkmanager.NewCoreNetwork(ctx, "example", &networkmanager.CoreNetworkArgs{
//				GlobalNetworkId: pulumi.Any(aws_networkmanager_global_network.Example.Id),
//				PolicyDocument:  pulumi.Any(data.Aws_networkmanager_core_network_policy_document.Example.Json),
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
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/networkmanager"
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
//
// ## Import
//
// `aws_networkmanager_core_network` can be imported using the core network ID, e.g.
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
	// Timestamp when a core network was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Description of the Core Network.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// One or more blocks detailing the edges within a core network. Detailed below.
	Edges CoreNetworkEdgeArrayOutput `pulumi:"edges"`
	// The ID of the global network that a core network will be a part of.
	GlobalNetworkId pulumi.StringOutput `pulumi:"globalNetworkId"`
	// Policy document for creating a core network. Note that updating this argument will result in the new policy document version being set as the `LATEST` and `LIVE` policy document. Refer to the [Core network policies documentation](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) for more information.
	PolicyDocument pulumi.StringPtrOutput `pulumi:"policyDocument"`
	// One or more blocks detailing the segments within a core network. Detailed below.
	Segments CoreNetworkSegmentArrayOutput `pulumi:"segments"`
	// Current state of a core network.
	State pulumi.StringOutput `pulumi:"state"`
	// Key-value tags for the Core Network. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
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
	// Timestamp when a core network was created.
	CreatedAt *string `pulumi:"createdAt"`
	// Description of the Core Network.
	Description *string `pulumi:"description"`
	// One or more blocks detailing the edges within a core network. Detailed below.
	Edges []CoreNetworkEdge `pulumi:"edges"`
	// The ID of the global network that a core network will be a part of.
	GlobalNetworkId *string `pulumi:"globalNetworkId"`
	// Policy document for creating a core network. Note that updating this argument will result in the new policy document version being set as the `LATEST` and `LIVE` policy document. Refer to the [Core network policies documentation](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) for more information.
	PolicyDocument *string `pulumi:"policyDocument"`
	// One or more blocks detailing the segments within a core network. Detailed below.
	Segments []CoreNetworkSegment `pulumi:"segments"`
	// Current state of a core network.
	State *string `pulumi:"state"`
	// Key-value tags for the Core Network. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type CoreNetworkState struct {
	// Core Network Amazon Resource Name (ARN).
	Arn pulumi.StringPtrInput
	// Timestamp when a core network was created.
	CreatedAt pulumi.StringPtrInput
	// Description of the Core Network.
	Description pulumi.StringPtrInput
	// One or more blocks detailing the edges within a core network. Detailed below.
	Edges CoreNetworkEdgeArrayInput
	// The ID of the global network that a core network will be a part of.
	GlobalNetworkId pulumi.StringPtrInput
	// Policy document for creating a core network. Note that updating this argument will result in the new policy document version being set as the `LATEST` and `LIVE` policy document. Refer to the [Core network policies documentation](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) for more information.
	PolicyDocument pulumi.StringPtrInput
	// One or more blocks detailing the segments within a core network. Detailed below.
	Segments CoreNetworkSegmentArrayInput
	// Current state of a core network.
	State pulumi.StringPtrInput
	// Key-value tags for the Core Network. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
}

func (CoreNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*coreNetworkState)(nil)).Elem()
}

type coreNetworkArgs struct {
	// Description of the Core Network.
	Description *string `pulumi:"description"`
	// The ID of the global network that a core network will be a part of.
	GlobalNetworkId string `pulumi:"globalNetworkId"`
	// Policy document for creating a core network. Note that updating this argument will result in the new policy document version being set as the `LATEST` and `LIVE` policy document. Refer to the [Core network policies documentation](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) for more information.
	PolicyDocument *string `pulumi:"policyDocument"`
	// Key-value tags for the Core Network. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CoreNetwork resource.
type CoreNetworkArgs struct {
	// Description of the Core Network.
	Description pulumi.StringPtrInput
	// The ID of the global network that a core network will be a part of.
	GlobalNetworkId pulumi.StringInput
	// Policy document for creating a core network. Note that updating this argument will result in the new policy document version being set as the `LATEST` and `LIVE` policy document. Refer to the [Core network policies documentation](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) for more information.
	PolicyDocument pulumi.StringPtrInput
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

// Core Network Amazon Resource Name (ARN).
func (o CoreNetworkOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *CoreNetwork) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
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

// Policy document for creating a core network. Note that updating this argument will result in the new policy document version being set as the `LATEST` and `LIVE` policy document. Refer to the [Core network policies documentation](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html) for more information.
func (o CoreNetworkOutput) PolicyDocument() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CoreNetwork) pulumi.StringPtrOutput { return v.PolicyDocument }).(pulumi.StringPtrOutput)
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
