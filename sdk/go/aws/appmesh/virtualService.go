// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an AWS App Mesh virtual service resource.
//
// ## Example Usage
// ### Virtual Node Provider
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appmesh"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := appmesh.NewVirtualService(ctx, "servicea", &appmesh.VirtualServiceArgs{
//				MeshName: pulumi.Any(aws_appmesh_mesh.Simple.Id),
//				Spec: &appmesh.VirtualServiceSpecArgs{
//					Provider: &appmesh.VirtualServiceSpecProviderArgs{
//						VirtualNode: &appmesh.VirtualServiceSpecProviderVirtualNodeArgs{
//							VirtualNodeName: pulumi.Any(aws_appmesh_virtual_node.Serviceb1.Name),
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
// ### Virtual Router Provider
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appmesh"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := appmesh.NewVirtualService(ctx, "servicea", &appmesh.VirtualServiceArgs{
//				MeshName: pulumi.Any(aws_appmesh_mesh.Simple.Id),
//				Spec: &appmesh.VirtualServiceSpecArgs{
//					Provider: &appmesh.VirtualServiceSpecProviderArgs{
//						VirtualRouter: &appmesh.VirtualServiceSpecProviderVirtualRouterArgs{
//							VirtualRouterName: pulumi.Any(aws_appmesh_virtual_router.Serviceb.Name),
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
// Using `pulumi import`, import App Mesh virtual services using `mesh_name` together with the virtual service's `name`. For example:
//
// ```sh
//
//	$ pulumi import aws:appmesh/virtualService:VirtualService servicea simpleapp/servicea.simpleapp.local
//
// ```
type VirtualService struct {
	pulumi.CustomResourceState

	// ARN of the virtual service.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Creation date of the virtual service.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// Last update date of the virtual service.
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`
	// Name of the service mesh in which to create the virtual service. Must be between 1 and 255 characters in length.
	MeshName pulumi.StringOutput `pulumi:"meshName"`
	// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
	MeshOwner pulumi.StringOutput `pulumi:"meshOwner"`
	// Name to use for the virtual service. Must be between 1 and 255 characters in length.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource owner's AWS account ID.
	ResourceOwner pulumi.StringOutput `pulumi:"resourceOwner"`
	// Virtual service specification to apply.
	Spec VirtualServiceSpecOutput `pulumi:"spec"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewVirtualService registers a new resource with the given unique name, arguments, and options.
func NewVirtualService(ctx *pulumi.Context,
	name string, args *VirtualServiceArgs, opts ...pulumi.ResourceOption) (*VirtualService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MeshName == nil {
		return nil, errors.New("invalid value for required argument 'MeshName'")
	}
	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"tagsAll",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualService
	err := ctx.RegisterResource("aws:appmesh/virtualService:VirtualService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualService gets an existing VirtualService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualServiceState, opts ...pulumi.ResourceOption) (*VirtualService, error) {
	var resource VirtualService
	err := ctx.ReadResource("aws:appmesh/virtualService:VirtualService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualService resources.
type virtualServiceState struct {
	// ARN of the virtual service.
	Arn *string `pulumi:"arn"`
	// Creation date of the virtual service.
	CreatedDate *string `pulumi:"createdDate"`
	// Last update date of the virtual service.
	LastUpdatedDate *string `pulumi:"lastUpdatedDate"`
	// Name of the service mesh in which to create the virtual service. Must be between 1 and 255 characters in length.
	MeshName *string `pulumi:"meshName"`
	// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
	MeshOwner *string `pulumi:"meshOwner"`
	// Name to use for the virtual service. Must be between 1 and 255 characters in length.
	Name *string `pulumi:"name"`
	// Resource owner's AWS account ID.
	ResourceOwner *string `pulumi:"resourceOwner"`
	// Virtual service specification to apply.
	Spec *VirtualServiceSpec `pulumi:"spec"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type VirtualServiceState struct {
	// ARN of the virtual service.
	Arn pulumi.StringPtrInput
	// Creation date of the virtual service.
	CreatedDate pulumi.StringPtrInput
	// Last update date of the virtual service.
	LastUpdatedDate pulumi.StringPtrInput
	// Name of the service mesh in which to create the virtual service. Must be between 1 and 255 characters in length.
	MeshName pulumi.StringPtrInput
	// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
	MeshOwner pulumi.StringPtrInput
	// Name to use for the virtual service. Must be between 1 and 255 characters in length.
	Name pulumi.StringPtrInput
	// Resource owner's AWS account ID.
	ResourceOwner pulumi.StringPtrInput
	// Virtual service specification to apply.
	Spec VirtualServiceSpecPtrInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	//
	// Deprecated: Please use `tags` instead.
	TagsAll pulumi.StringMapInput
}

func (VirtualServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualServiceState)(nil)).Elem()
}

type virtualServiceArgs struct {
	// Name of the service mesh in which to create the virtual service. Must be between 1 and 255 characters in length.
	MeshName string `pulumi:"meshName"`
	// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
	MeshOwner *string `pulumi:"meshOwner"`
	// Name to use for the virtual service. Must be between 1 and 255 characters in length.
	Name *string `pulumi:"name"`
	// Virtual service specification to apply.
	Spec VirtualServiceSpec `pulumi:"spec"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a VirtualService resource.
type VirtualServiceArgs struct {
	// Name of the service mesh in which to create the virtual service. Must be between 1 and 255 characters in length.
	MeshName pulumi.StringInput
	// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
	MeshOwner pulumi.StringPtrInput
	// Name to use for the virtual service. Must be between 1 and 255 characters in length.
	Name pulumi.StringPtrInput
	// Virtual service specification to apply.
	Spec VirtualServiceSpecInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (VirtualServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualServiceArgs)(nil)).Elem()
}

type VirtualServiceInput interface {
	pulumi.Input

	ToVirtualServiceOutput() VirtualServiceOutput
	ToVirtualServiceOutputWithContext(ctx context.Context) VirtualServiceOutput
}

func (*VirtualService) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualService)(nil)).Elem()
}

func (i *VirtualService) ToVirtualServiceOutput() VirtualServiceOutput {
	return i.ToVirtualServiceOutputWithContext(context.Background())
}

func (i *VirtualService) ToVirtualServiceOutputWithContext(ctx context.Context) VirtualServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualServiceOutput)
}

// VirtualServiceArrayInput is an input type that accepts VirtualServiceArray and VirtualServiceArrayOutput values.
// You can construct a concrete instance of `VirtualServiceArrayInput` via:
//
//	VirtualServiceArray{ VirtualServiceArgs{...} }
type VirtualServiceArrayInput interface {
	pulumi.Input

	ToVirtualServiceArrayOutput() VirtualServiceArrayOutput
	ToVirtualServiceArrayOutputWithContext(context.Context) VirtualServiceArrayOutput
}

type VirtualServiceArray []VirtualServiceInput

func (VirtualServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualService)(nil)).Elem()
}

func (i VirtualServiceArray) ToVirtualServiceArrayOutput() VirtualServiceArrayOutput {
	return i.ToVirtualServiceArrayOutputWithContext(context.Background())
}

func (i VirtualServiceArray) ToVirtualServiceArrayOutputWithContext(ctx context.Context) VirtualServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualServiceArrayOutput)
}

// VirtualServiceMapInput is an input type that accepts VirtualServiceMap and VirtualServiceMapOutput values.
// You can construct a concrete instance of `VirtualServiceMapInput` via:
//
//	VirtualServiceMap{ "key": VirtualServiceArgs{...} }
type VirtualServiceMapInput interface {
	pulumi.Input

	ToVirtualServiceMapOutput() VirtualServiceMapOutput
	ToVirtualServiceMapOutputWithContext(context.Context) VirtualServiceMapOutput
}

type VirtualServiceMap map[string]VirtualServiceInput

func (VirtualServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualService)(nil)).Elem()
}

func (i VirtualServiceMap) ToVirtualServiceMapOutput() VirtualServiceMapOutput {
	return i.ToVirtualServiceMapOutputWithContext(context.Background())
}

func (i VirtualServiceMap) ToVirtualServiceMapOutputWithContext(ctx context.Context) VirtualServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualServiceMapOutput)
}

type VirtualServiceOutput struct{ *pulumi.OutputState }

func (VirtualServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualService)(nil)).Elem()
}

func (o VirtualServiceOutput) ToVirtualServiceOutput() VirtualServiceOutput {
	return o
}

func (o VirtualServiceOutput) ToVirtualServiceOutputWithContext(ctx context.Context) VirtualServiceOutput {
	return o
}

// ARN of the virtual service.
func (o VirtualServiceOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualService) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Creation date of the virtual service.
func (o VirtualServiceOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualService) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

// Last update date of the virtual service.
func (o VirtualServiceOutput) LastUpdatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualService) pulumi.StringOutput { return v.LastUpdatedDate }).(pulumi.StringOutput)
}

// Name of the service mesh in which to create the virtual service. Must be between 1 and 255 characters in length.
func (o VirtualServiceOutput) MeshName() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualService) pulumi.StringOutput { return v.MeshName }).(pulumi.StringOutput)
}

// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
func (o VirtualServiceOutput) MeshOwner() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualService) pulumi.StringOutput { return v.MeshOwner }).(pulumi.StringOutput)
}

// Name to use for the virtual service. Must be between 1 and 255 characters in length.
func (o VirtualServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource owner's AWS account ID.
func (o VirtualServiceOutput) ResourceOwner() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualService) pulumi.StringOutput { return v.ResourceOwner }).(pulumi.StringOutput)
}

// Virtual service specification to apply.
func (o VirtualServiceOutput) Spec() VirtualServiceSpecOutput {
	return o.ApplyT(func(v *VirtualService) VirtualServiceSpecOutput { return v.Spec }).(VirtualServiceSpecOutput)
}

// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o VirtualServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
//
// Deprecated: Please use `tags` instead.
func (o VirtualServiceOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualService) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

type VirtualServiceArrayOutput struct{ *pulumi.OutputState }

func (VirtualServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualService)(nil)).Elem()
}

func (o VirtualServiceArrayOutput) ToVirtualServiceArrayOutput() VirtualServiceArrayOutput {
	return o
}

func (o VirtualServiceArrayOutput) ToVirtualServiceArrayOutputWithContext(ctx context.Context) VirtualServiceArrayOutput {
	return o
}

func (o VirtualServiceArrayOutput) Index(i pulumi.IntInput) VirtualServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualService {
		return vs[0].([]*VirtualService)[vs[1].(int)]
	}).(VirtualServiceOutput)
}

type VirtualServiceMapOutput struct{ *pulumi.OutputState }

func (VirtualServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualService)(nil)).Elem()
}

func (o VirtualServiceMapOutput) ToVirtualServiceMapOutput() VirtualServiceMapOutput {
	return o
}

func (o VirtualServiceMapOutput) ToVirtualServiceMapOutputWithContext(ctx context.Context) VirtualServiceMapOutput {
	return o
}

func (o VirtualServiceMapOutput) MapIndex(k pulumi.StringInput) VirtualServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualService {
		return vs[0].(map[string]*VirtualService)[vs[1].(string)]
	}).(VirtualServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualServiceInput)(nil)).Elem(), &VirtualService{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualServiceArrayInput)(nil)).Elem(), VirtualServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualServiceMapInput)(nil)).Elem(), VirtualServiceMap{})
	pulumi.RegisterOutputType(VirtualServiceOutput{})
	pulumi.RegisterOutputType(VirtualServiceArrayOutput{})
	pulumi.RegisterOutputType(VirtualServiceMapOutput{})
}
