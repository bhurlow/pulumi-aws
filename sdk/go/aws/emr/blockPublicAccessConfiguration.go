// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for managing an AWS EMR block public access configuration. This region level security configuration restricts the launch of EMR clusters that have associated security groups permitting public access on unspecified ports. See the [EMR Block Public Access Configuration](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-block-public-access.html) documentation for further information.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/emr"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := emr.NewBlockPublicAccessConfiguration(ctx, "example", &emr.BlockPublicAccessConfigurationArgs{
//				BlockPublicSecurityGroupRules: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Default Configuration
//
// By default, each AWS region is equipped with a block public access configuration that prevents EMR clusters from being launched if they have security group rules permitting public access on any port except for port 22. The default configuration can be managed using this resource.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/emr"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := emr.NewBlockPublicAccessConfiguration(ctx, "example", &emr.BlockPublicAccessConfigurationArgs{
//				BlockPublicSecurityGroupRules: pulumi.Bool(true),
//				PermittedPublicSecurityGroupRuleRanges: emr.BlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeArray{
//					&emr.BlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeArgs{
//						MaxRange: pulumi.Int(22),
//						MinRange: pulumi.Int(22),
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
// > **NOTE:** If an `emr.BlockPublicAccessConfiguration` resource is destroyed, the configuration will reset to this default configuration.
// ### Multiple Permitted Public Security Group Rule Ranges
//
// The resource permits specification of multiple `permittedPublicSecurityGroupRuleRange` blocks.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/emr"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := emr.NewBlockPublicAccessConfiguration(ctx, "example", &emr.BlockPublicAccessConfigurationArgs{
//				BlockPublicSecurityGroupRules: pulumi.Bool(true),
//				PermittedPublicSecurityGroupRuleRanges: emr.BlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeArray{
//					&emr.BlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeArgs{
//						MaxRange: pulumi.Int(22),
//						MinRange: pulumi.Int(22),
//					},
//					&emr.BlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeArgs{
//						MaxRange: pulumi.Int(101),
//						MinRange: pulumi.Int(100),
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
// ### Disabling Block Public Access
//
// To permit EMR clusters to be launched in the configured region regardless of associated security group rules, the Block Public Access feature can be disabled using this resource.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/emr"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := emr.NewBlockPublicAccessConfiguration(ctx, "example", &emr.BlockPublicAccessConfigurationArgs{
//				BlockPublicSecurityGroupRules: pulumi.Bool(false),
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
// The current EMR Block Public Access Configuration can be imported, e.g.,
//
// ```sh
//
//	$ pulumi import aws:emr/blockPublicAccessConfiguration:BlockPublicAccessConfiguration example current
//
// ```
type BlockPublicAccessConfiguration struct {
	pulumi.CustomResourceState

	// Enable or disable EMR Block Public Access.
	BlockPublicSecurityGroupRules pulumi.BoolOutput `pulumi:"blockPublicSecurityGroupRules"`
	// Configuration block for defining permitted public security group rule port ranges. Can be defined multiple times per resource. Only valid if `blockPublicSecurityGroupRules` is set to `true`.
	PermittedPublicSecurityGroupRuleRanges BlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeArrayOutput `pulumi:"permittedPublicSecurityGroupRuleRanges"`
}

// NewBlockPublicAccessConfiguration registers a new resource with the given unique name, arguments, and options.
func NewBlockPublicAccessConfiguration(ctx *pulumi.Context,
	name string, args *BlockPublicAccessConfigurationArgs, opts ...pulumi.ResourceOption) (*BlockPublicAccessConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BlockPublicSecurityGroupRules == nil {
		return nil, errors.New("invalid value for required argument 'BlockPublicSecurityGroupRules'")
	}
	var resource BlockPublicAccessConfiguration
	err := ctx.RegisterResource("aws:emr/blockPublicAccessConfiguration:BlockPublicAccessConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBlockPublicAccessConfiguration gets an existing BlockPublicAccessConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBlockPublicAccessConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlockPublicAccessConfigurationState, opts ...pulumi.ResourceOption) (*BlockPublicAccessConfiguration, error) {
	var resource BlockPublicAccessConfiguration
	err := ctx.ReadResource("aws:emr/blockPublicAccessConfiguration:BlockPublicAccessConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BlockPublicAccessConfiguration resources.
type blockPublicAccessConfigurationState struct {
	// Enable or disable EMR Block Public Access.
	BlockPublicSecurityGroupRules *bool `pulumi:"blockPublicSecurityGroupRules"`
	// Configuration block for defining permitted public security group rule port ranges. Can be defined multiple times per resource. Only valid if `blockPublicSecurityGroupRules` is set to `true`.
	PermittedPublicSecurityGroupRuleRanges []BlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRange `pulumi:"permittedPublicSecurityGroupRuleRanges"`
}

type BlockPublicAccessConfigurationState struct {
	// Enable or disable EMR Block Public Access.
	BlockPublicSecurityGroupRules pulumi.BoolPtrInput
	// Configuration block for defining permitted public security group rule port ranges. Can be defined multiple times per resource. Only valid if `blockPublicSecurityGroupRules` is set to `true`.
	PermittedPublicSecurityGroupRuleRanges BlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeArrayInput
}

func (BlockPublicAccessConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*blockPublicAccessConfigurationState)(nil)).Elem()
}

type blockPublicAccessConfigurationArgs struct {
	// Enable or disable EMR Block Public Access.
	BlockPublicSecurityGroupRules bool `pulumi:"blockPublicSecurityGroupRules"`
	// Configuration block for defining permitted public security group rule port ranges. Can be defined multiple times per resource. Only valid if `blockPublicSecurityGroupRules` is set to `true`.
	PermittedPublicSecurityGroupRuleRanges []BlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRange `pulumi:"permittedPublicSecurityGroupRuleRanges"`
}

// The set of arguments for constructing a BlockPublicAccessConfiguration resource.
type BlockPublicAccessConfigurationArgs struct {
	// Enable or disable EMR Block Public Access.
	BlockPublicSecurityGroupRules pulumi.BoolInput
	// Configuration block for defining permitted public security group rule port ranges. Can be defined multiple times per resource. Only valid if `blockPublicSecurityGroupRules` is set to `true`.
	PermittedPublicSecurityGroupRuleRanges BlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeArrayInput
}

func (BlockPublicAccessConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blockPublicAccessConfigurationArgs)(nil)).Elem()
}

type BlockPublicAccessConfigurationInput interface {
	pulumi.Input

	ToBlockPublicAccessConfigurationOutput() BlockPublicAccessConfigurationOutput
	ToBlockPublicAccessConfigurationOutputWithContext(ctx context.Context) BlockPublicAccessConfigurationOutput
}

func (*BlockPublicAccessConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**BlockPublicAccessConfiguration)(nil)).Elem()
}

func (i *BlockPublicAccessConfiguration) ToBlockPublicAccessConfigurationOutput() BlockPublicAccessConfigurationOutput {
	return i.ToBlockPublicAccessConfigurationOutputWithContext(context.Background())
}

func (i *BlockPublicAccessConfiguration) ToBlockPublicAccessConfigurationOutputWithContext(ctx context.Context) BlockPublicAccessConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlockPublicAccessConfigurationOutput)
}

// BlockPublicAccessConfigurationArrayInput is an input type that accepts BlockPublicAccessConfigurationArray and BlockPublicAccessConfigurationArrayOutput values.
// You can construct a concrete instance of `BlockPublicAccessConfigurationArrayInput` via:
//
//	BlockPublicAccessConfigurationArray{ BlockPublicAccessConfigurationArgs{...} }
type BlockPublicAccessConfigurationArrayInput interface {
	pulumi.Input

	ToBlockPublicAccessConfigurationArrayOutput() BlockPublicAccessConfigurationArrayOutput
	ToBlockPublicAccessConfigurationArrayOutputWithContext(context.Context) BlockPublicAccessConfigurationArrayOutput
}

type BlockPublicAccessConfigurationArray []BlockPublicAccessConfigurationInput

func (BlockPublicAccessConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BlockPublicAccessConfiguration)(nil)).Elem()
}

func (i BlockPublicAccessConfigurationArray) ToBlockPublicAccessConfigurationArrayOutput() BlockPublicAccessConfigurationArrayOutput {
	return i.ToBlockPublicAccessConfigurationArrayOutputWithContext(context.Background())
}

func (i BlockPublicAccessConfigurationArray) ToBlockPublicAccessConfigurationArrayOutputWithContext(ctx context.Context) BlockPublicAccessConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlockPublicAccessConfigurationArrayOutput)
}

// BlockPublicAccessConfigurationMapInput is an input type that accepts BlockPublicAccessConfigurationMap and BlockPublicAccessConfigurationMapOutput values.
// You can construct a concrete instance of `BlockPublicAccessConfigurationMapInput` via:
//
//	BlockPublicAccessConfigurationMap{ "key": BlockPublicAccessConfigurationArgs{...} }
type BlockPublicAccessConfigurationMapInput interface {
	pulumi.Input

	ToBlockPublicAccessConfigurationMapOutput() BlockPublicAccessConfigurationMapOutput
	ToBlockPublicAccessConfigurationMapOutputWithContext(context.Context) BlockPublicAccessConfigurationMapOutput
}

type BlockPublicAccessConfigurationMap map[string]BlockPublicAccessConfigurationInput

func (BlockPublicAccessConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BlockPublicAccessConfiguration)(nil)).Elem()
}

func (i BlockPublicAccessConfigurationMap) ToBlockPublicAccessConfigurationMapOutput() BlockPublicAccessConfigurationMapOutput {
	return i.ToBlockPublicAccessConfigurationMapOutputWithContext(context.Background())
}

func (i BlockPublicAccessConfigurationMap) ToBlockPublicAccessConfigurationMapOutputWithContext(ctx context.Context) BlockPublicAccessConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlockPublicAccessConfigurationMapOutput)
}

type BlockPublicAccessConfigurationOutput struct{ *pulumi.OutputState }

func (BlockPublicAccessConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlockPublicAccessConfiguration)(nil)).Elem()
}

func (o BlockPublicAccessConfigurationOutput) ToBlockPublicAccessConfigurationOutput() BlockPublicAccessConfigurationOutput {
	return o
}

func (o BlockPublicAccessConfigurationOutput) ToBlockPublicAccessConfigurationOutputWithContext(ctx context.Context) BlockPublicAccessConfigurationOutput {
	return o
}

// Enable or disable EMR Block Public Access.
func (o BlockPublicAccessConfigurationOutput) BlockPublicSecurityGroupRules() pulumi.BoolOutput {
	return o.ApplyT(func(v *BlockPublicAccessConfiguration) pulumi.BoolOutput { return v.BlockPublicSecurityGroupRules }).(pulumi.BoolOutput)
}

// Configuration block for defining permitted public security group rule port ranges. Can be defined multiple times per resource. Only valid if `blockPublicSecurityGroupRules` is set to `true`.
func (o BlockPublicAccessConfigurationOutput) PermittedPublicSecurityGroupRuleRanges() BlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeArrayOutput {
	return o.ApplyT(func(v *BlockPublicAccessConfiguration) BlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeArrayOutput {
		return v.PermittedPublicSecurityGroupRuleRanges
	}).(BlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRangeArrayOutput)
}

type BlockPublicAccessConfigurationArrayOutput struct{ *pulumi.OutputState }

func (BlockPublicAccessConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BlockPublicAccessConfiguration)(nil)).Elem()
}

func (o BlockPublicAccessConfigurationArrayOutput) ToBlockPublicAccessConfigurationArrayOutput() BlockPublicAccessConfigurationArrayOutput {
	return o
}

func (o BlockPublicAccessConfigurationArrayOutput) ToBlockPublicAccessConfigurationArrayOutputWithContext(ctx context.Context) BlockPublicAccessConfigurationArrayOutput {
	return o
}

func (o BlockPublicAccessConfigurationArrayOutput) Index(i pulumi.IntInput) BlockPublicAccessConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BlockPublicAccessConfiguration {
		return vs[0].([]*BlockPublicAccessConfiguration)[vs[1].(int)]
	}).(BlockPublicAccessConfigurationOutput)
}

type BlockPublicAccessConfigurationMapOutput struct{ *pulumi.OutputState }

func (BlockPublicAccessConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BlockPublicAccessConfiguration)(nil)).Elem()
}

func (o BlockPublicAccessConfigurationMapOutput) ToBlockPublicAccessConfigurationMapOutput() BlockPublicAccessConfigurationMapOutput {
	return o
}

func (o BlockPublicAccessConfigurationMapOutput) ToBlockPublicAccessConfigurationMapOutputWithContext(ctx context.Context) BlockPublicAccessConfigurationMapOutput {
	return o
}

func (o BlockPublicAccessConfigurationMapOutput) MapIndex(k pulumi.StringInput) BlockPublicAccessConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BlockPublicAccessConfiguration {
		return vs[0].(map[string]*BlockPublicAccessConfiguration)[vs[1].(string)]
	}).(BlockPublicAccessConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BlockPublicAccessConfigurationInput)(nil)).Elem(), &BlockPublicAccessConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*BlockPublicAccessConfigurationArrayInput)(nil)).Elem(), BlockPublicAccessConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BlockPublicAccessConfigurationMapInput)(nil)).Elem(), BlockPublicAccessConfigurationMap{})
	pulumi.RegisterOutputType(BlockPublicAccessConfigurationOutput{})
	pulumi.RegisterOutputType(BlockPublicAccessConfigurationArrayOutput{})
	pulumi.RegisterOutputType(BlockPublicAccessConfigurationMapOutput{})
}
