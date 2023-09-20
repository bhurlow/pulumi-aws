// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package emrcontainers

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages an EMR Containers (EMR on EKS) Virtual Cluster.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/emrcontainers"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := emrcontainers.NewVirtualCluster(ctx, "example", &emrcontainers.VirtualClusterArgs{
//				ContainerProvider: &emrcontainers.VirtualClusterContainerProviderArgs{
//					Id:   pulumi.Any(aws_eks_cluster.Example.Name),
//					Type: pulumi.String("EKS"),
//					Info: &emrcontainers.VirtualClusterContainerProviderInfoArgs{
//						EksInfo: &emrcontainers.VirtualClusterContainerProviderInfoEksInfoArgs{
//							Namespace: pulumi.String("default"),
//						},
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
// In TODO v1.5.0 and later, use an `import` block to import EKS Clusters using the `id`. For exampleterraform import {
//
//	to = aws_emrcontainers_virtual_cluster.example
//
//	id = "a1b2c3d4e5f6g7h8i9j10k11l" } Using `TODO import`, import EKS Clusters using the `id`. For exampleconsole % TODO import aws_emrcontainers_virtual_cluster.example a1b2c3d4e5f6g7h8i9j10k11l
type VirtualCluster struct {
	pulumi.CustomResourceState

	// ARN of the cluster.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Configuration block for the container provider associated with your cluster.
	ContainerProvider VirtualClusterContainerProviderOutput `pulumi:"containerProvider"`
	// Name of the virtual cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewVirtualCluster registers a new resource with the given unique name, arguments, and options.
func NewVirtualCluster(ctx *pulumi.Context,
	name string, args *VirtualClusterArgs, opts ...pulumi.ResourceOption) (*VirtualCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContainerProvider == nil {
		return nil, errors.New("invalid value for required argument 'ContainerProvider'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualCluster
	err := ctx.RegisterResource("aws:emrcontainers/virtualCluster:VirtualCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualCluster gets an existing VirtualCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualClusterState, opts ...pulumi.ResourceOption) (*VirtualCluster, error) {
	var resource VirtualCluster
	err := ctx.ReadResource("aws:emrcontainers/virtualCluster:VirtualCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualCluster resources.
type virtualClusterState struct {
	// ARN of the cluster.
	Arn *string `pulumi:"arn"`
	// Configuration block for the container provider associated with your cluster.
	ContainerProvider *VirtualClusterContainerProvider `pulumi:"containerProvider"`
	// Name of the virtual cluster.
	Name *string `pulumi:"name"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type VirtualClusterState struct {
	// ARN of the cluster.
	Arn pulumi.StringPtrInput
	// Configuration block for the container provider associated with your cluster.
	ContainerProvider VirtualClusterContainerProviderPtrInput
	// Name of the virtual cluster.
	Name pulumi.StringPtrInput
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (VirtualClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualClusterState)(nil)).Elem()
}

type virtualClusterArgs struct {
	// Configuration block for the container provider associated with your cluster.
	ContainerProvider VirtualClusterContainerProvider `pulumi:"containerProvider"`
	// Name of the virtual cluster.
	Name *string `pulumi:"name"`
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a VirtualCluster resource.
type VirtualClusterArgs struct {
	// Configuration block for the container provider associated with your cluster.
	ContainerProvider VirtualClusterContainerProviderInput
	// Name of the virtual cluster.
	Name pulumi.StringPtrInput
	// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (VirtualClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualClusterArgs)(nil)).Elem()
}

type VirtualClusterInput interface {
	pulumi.Input

	ToVirtualClusterOutput() VirtualClusterOutput
	ToVirtualClusterOutputWithContext(ctx context.Context) VirtualClusterOutput
}

func (*VirtualCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualCluster)(nil)).Elem()
}

func (i *VirtualCluster) ToVirtualClusterOutput() VirtualClusterOutput {
	return i.ToVirtualClusterOutputWithContext(context.Background())
}

func (i *VirtualCluster) ToVirtualClusterOutputWithContext(ctx context.Context) VirtualClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualClusterOutput)
}

func (i *VirtualCluster) ToOutput(ctx context.Context) pulumix.Output[*VirtualCluster] {
	return pulumix.Output[*VirtualCluster]{
		OutputState: i.ToVirtualClusterOutputWithContext(ctx).OutputState,
	}
}

// VirtualClusterArrayInput is an input type that accepts VirtualClusterArray and VirtualClusterArrayOutput values.
// You can construct a concrete instance of `VirtualClusterArrayInput` via:
//
//	VirtualClusterArray{ VirtualClusterArgs{...} }
type VirtualClusterArrayInput interface {
	pulumi.Input

	ToVirtualClusterArrayOutput() VirtualClusterArrayOutput
	ToVirtualClusterArrayOutputWithContext(context.Context) VirtualClusterArrayOutput
}

type VirtualClusterArray []VirtualClusterInput

func (VirtualClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualCluster)(nil)).Elem()
}

func (i VirtualClusterArray) ToVirtualClusterArrayOutput() VirtualClusterArrayOutput {
	return i.ToVirtualClusterArrayOutputWithContext(context.Background())
}

func (i VirtualClusterArray) ToVirtualClusterArrayOutputWithContext(ctx context.Context) VirtualClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualClusterArrayOutput)
}

func (i VirtualClusterArray) ToOutput(ctx context.Context) pulumix.Output[[]*VirtualCluster] {
	return pulumix.Output[[]*VirtualCluster]{
		OutputState: i.ToVirtualClusterArrayOutputWithContext(ctx).OutputState,
	}
}

// VirtualClusterMapInput is an input type that accepts VirtualClusterMap and VirtualClusterMapOutput values.
// You can construct a concrete instance of `VirtualClusterMapInput` via:
//
//	VirtualClusterMap{ "key": VirtualClusterArgs{...} }
type VirtualClusterMapInput interface {
	pulumi.Input

	ToVirtualClusterMapOutput() VirtualClusterMapOutput
	ToVirtualClusterMapOutputWithContext(context.Context) VirtualClusterMapOutput
}

type VirtualClusterMap map[string]VirtualClusterInput

func (VirtualClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualCluster)(nil)).Elem()
}

func (i VirtualClusterMap) ToVirtualClusterMapOutput() VirtualClusterMapOutput {
	return i.ToVirtualClusterMapOutputWithContext(context.Background())
}

func (i VirtualClusterMap) ToVirtualClusterMapOutputWithContext(ctx context.Context) VirtualClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualClusterMapOutput)
}

func (i VirtualClusterMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*VirtualCluster] {
	return pulumix.Output[map[string]*VirtualCluster]{
		OutputState: i.ToVirtualClusterMapOutputWithContext(ctx).OutputState,
	}
}

type VirtualClusterOutput struct{ *pulumi.OutputState }

func (VirtualClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualCluster)(nil)).Elem()
}

func (o VirtualClusterOutput) ToVirtualClusterOutput() VirtualClusterOutput {
	return o
}

func (o VirtualClusterOutput) ToVirtualClusterOutputWithContext(ctx context.Context) VirtualClusterOutput {
	return o
}

func (o VirtualClusterOutput) ToOutput(ctx context.Context) pulumix.Output[*VirtualCluster] {
	return pulumix.Output[*VirtualCluster]{
		OutputState: o.OutputState,
	}
}

// ARN of the cluster.
func (o VirtualClusterOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualCluster) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Configuration block for the container provider associated with your cluster.
func (o VirtualClusterOutput) ContainerProvider() VirtualClusterContainerProviderOutput {
	return o.ApplyT(func(v *VirtualCluster) VirtualClusterContainerProviderOutput { return v.ContainerProvider }).(VirtualClusterContainerProviderOutput)
}

// Name of the virtual cluster.
func (o VirtualClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o VirtualClusterOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualCluster) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o VirtualClusterOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualCluster) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type VirtualClusterArrayOutput struct{ *pulumi.OutputState }

func (VirtualClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualCluster)(nil)).Elem()
}

func (o VirtualClusterArrayOutput) ToVirtualClusterArrayOutput() VirtualClusterArrayOutput {
	return o
}

func (o VirtualClusterArrayOutput) ToVirtualClusterArrayOutputWithContext(ctx context.Context) VirtualClusterArrayOutput {
	return o
}

func (o VirtualClusterArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*VirtualCluster] {
	return pulumix.Output[[]*VirtualCluster]{
		OutputState: o.OutputState,
	}
}

func (o VirtualClusterArrayOutput) Index(i pulumi.IntInput) VirtualClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualCluster {
		return vs[0].([]*VirtualCluster)[vs[1].(int)]
	}).(VirtualClusterOutput)
}

type VirtualClusterMapOutput struct{ *pulumi.OutputState }

func (VirtualClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualCluster)(nil)).Elem()
}

func (o VirtualClusterMapOutput) ToVirtualClusterMapOutput() VirtualClusterMapOutput {
	return o
}

func (o VirtualClusterMapOutput) ToVirtualClusterMapOutputWithContext(ctx context.Context) VirtualClusterMapOutput {
	return o
}

func (o VirtualClusterMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*VirtualCluster] {
	return pulumix.Output[map[string]*VirtualCluster]{
		OutputState: o.OutputState,
	}
}

func (o VirtualClusterMapOutput) MapIndex(k pulumi.StringInput) VirtualClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualCluster {
		return vs[0].(map[string]*VirtualCluster)[vs[1].(string)]
	}).(VirtualClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualClusterInput)(nil)).Elem(), &VirtualCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualClusterArrayInput)(nil)).Elem(), VirtualClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualClusterMapInput)(nil)).Elem(), VirtualClusterMap{})
	pulumi.RegisterOutputType(VirtualClusterOutput{})
	pulumi.RegisterOutputType(VirtualClusterArrayOutput{})
	pulumi.RegisterOutputType(VirtualClusterMapOutput{})
}
