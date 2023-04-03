// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an AWS App Mesh gateway route resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/appmesh"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := appmesh.NewGatewayRoute(ctx, "example", &appmesh.GatewayRouteArgs{
//				MeshName:           pulumi.String("example-service-mesh"),
//				VirtualGatewayName: pulumi.Any(aws_appmesh_virtual_gateway.Example.Name),
//				Spec: &appmesh.GatewayRouteSpecArgs{
//					HttpRoute: &appmesh.GatewayRouteSpecHttpRouteArgs{
//						Action: &appmesh.GatewayRouteSpecHttpRouteActionArgs{
//							Target: &appmesh.GatewayRouteSpecHttpRouteActionTargetArgs{
//								VirtualService: &appmesh.GatewayRouteSpecHttpRouteActionTargetVirtualServiceArgs{
//									VirtualServiceName: pulumi.Any(aws_appmesh_virtual_service.Example.Name),
//								},
//							},
//						},
//						Match: &appmesh.GatewayRouteSpecHttpRouteMatchArgs{
//							Prefix: pulumi.String("/"),
//						},
//					},
//				},
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
//
// ## Import
//
// App Mesh gateway routes can be imported using `mesh_name` and `virtual_gateway_name` together with the gateway route's `name`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:appmesh/gatewayRoute:GatewayRoute example mesh/gw1/example-gateway-route
//
// ```
type GatewayRoute struct {
	pulumi.CustomResourceState

	// ARN of the gateway route.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Creation date of the gateway route.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// Last update date of the gateway route.
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`
	// Name of the service mesh in which to create the gateway route. Must be between 1 and 255 characters in length.
	MeshName pulumi.StringOutput `pulumi:"meshName"`
	// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
	MeshOwner pulumi.StringOutput `pulumi:"meshOwner"`
	// Name to use for the gateway route. Must be between 1 and 255 characters in length.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource owner's AWS account ID.
	ResourceOwner pulumi.StringOutput `pulumi:"resourceOwner"`
	// Gateway route specification to apply.
	Spec GatewayRouteSpecOutput `pulumi:"spec"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Name of the virtual gateway to associate the gateway route with. Must be between 1 and 255 characters in length.
	VirtualGatewayName pulumi.StringOutput `pulumi:"virtualGatewayName"`
}

// NewGatewayRoute registers a new resource with the given unique name, arguments, and options.
func NewGatewayRoute(ctx *pulumi.Context,
	name string, args *GatewayRouteArgs, opts ...pulumi.ResourceOption) (*GatewayRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MeshName == nil {
		return nil, errors.New("invalid value for required argument 'MeshName'")
	}
	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	if args.VirtualGatewayName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualGatewayName'")
	}
	var resource GatewayRoute
	err := ctx.RegisterResource("aws:appmesh/gatewayRoute:GatewayRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayRoute gets an existing GatewayRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayRouteState, opts ...pulumi.ResourceOption) (*GatewayRoute, error) {
	var resource GatewayRoute
	err := ctx.ReadResource("aws:appmesh/gatewayRoute:GatewayRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayRoute resources.
type gatewayRouteState struct {
	// ARN of the gateway route.
	Arn *string `pulumi:"arn"`
	// Creation date of the gateway route.
	CreatedDate *string `pulumi:"createdDate"`
	// Last update date of the gateway route.
	LastUpdatedDate *string `pulumi:"lastUpdatedDate"`
	// Name of the service mesh in which to create the gateway route. Must be between 1 and 255 characters in length.
	MeshName *string `pulumi:"meshName"`
	// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
	MeshOwner *string `pulumi:"meshOwner"`
	// Name to use for the gateway route. Must be between 1 and 255 characters in length.
	Name *string `pulumi:"name"`
	// Resource owner's AWS account ID.
	ResourceOwner *string `pulumi:"resourceOwner"`
	// Gateway route specification to apply.
	Spec *GatewayRouteSpec `pulumi:"spec"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Name of the virtual gateway to associate the gateway route with. Must be between 1 and 255 characters in length.
	VirtualGatewayName *string `pulumi:"virtualGatewayName"`
}

type GatewayRouteState struct {
	// ARN of the gateway route.
	Arn pulumi.StringPtrInput
	// Creation date of the gateway route.
	CreatedDate pulumi.StringPtrInput
	// Last update date of the gateway route.
	LastUpdatedDate pulumi.StringPtrInput
	// Name of the service mesh in which to create the gateway route. Must be between 1 and 255 characters in length.
	MeshName pulumi.StringPtrInput
	// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
	MeshOwner pulumi.StringPtrInput
	// Name to use for the gateway route. Must be between 1 and 255 characters in length.
	Name pulumi.StringPtrInput
	// Resource owner's AWS account ID.
	ResourceOwner pulumi.StringPtrInput
	// Gateway route specification to apply.
	Spec GatewayRouteSpecPtrInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// Name of the virtual gateway to associate the gateway route with. Must be between 1 and 255 characters in length.
	VirtualGatewayName pulumi.StringPtrInput
}

func (GatewayRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayRouteState)(nil)).Elem()
}

type gatewayRouteArgs struct {
	// Name of the service mesh in which to create the gateway route. Must be between 1 and 255 characters in length.
	MeshName string `pulumi:"meshName"`
	// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
	MeshOwner *string `pulumi:"meshOwner"`
	// Name to use for the gateway route. Must be between 1 and 255 characters in length.
	Name *string `pulumi:"name"`
	// Gateway route specification to apply.
	Spec GatewayRouteSpec `pulumi:"spec"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Name of the virtual gateway to associate the gateway route with. Must be between 1 and 255 characters in length.
	VirtualGatewayName string `pulumi:"virtualGatewayName"`
}

// The set of arguments for constructing a GatewayRoute resource.
type GatewayRouteArgs struct {
	// Name of the service mesh in which to create the gateway route. Must be between 1 and 255 characters in length.
	MeshName pulumi.StringInput
	// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
	MeshOwner pulumi.StringPtrInput
	// Name to use for the gateway route. Must be between 1 and 255 characters in length.
	Name pulumi.StringPtrInput
	// Gateway route specification to apply.
	Spec GatewayRouteSpecInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Name of the virtual gateway to associate the gateway route with. Must be between 1 and 255 characters in length.
	VirtualGatewayName pulumi.StringInput
}

func (GatewayRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayRouteArgs)(nil)).Elem()
}

type GatewayRouteInput interface {
	pulumi.Input

	ToGatewayRouteOutput() GatewayRouteOutput
	ToGatewayRouteOutputWithContext(ctx context.Context) GatewayRouteOutput
}

func (*GatewayRoute) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayRoute)(nil)).Elem()
}

func (i *GatewayRoute) ToGatewayRouteOutput() GatewayRouteOutput {
	return i.ToGatewayRouteOutputWithContext(context.Background())
}

func (i *GatewayRoute) ToGatewayRouteOutputWithContext(ctx context.Context) GatewayRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayRouteOutput)
}

// GatewayRouteArrayInput is an input type that accepts GatewayRouteArray and GatewayRouteArrayOutput values.
// You can construct a concrete instance of `GatewayRouteArrayInput` via:
//
//	GatewayRouteArray{ GatewayRouteArgs{...} }
type GatewayRouteArrayInput interface {
	pulumi.Input

	ToGatewayRouteArrayOutput() GatewayRouteArrayOutput
	ToGatewayRouteArrayOutputWithContext(context.Context) GatewayRouteArrayOutput
}

type GatewayRouteArray []GatewayRouteInput

func (GatewayRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayRoute)(nil)).Elem()
}

func (i GatewayRouteArray) ToGatewayRouteArrayOutput() GatewayRouteArrayOutput {
	return i.ToGatewayRouteArrayOutputWithContext(context.Background())
}

func (i GatewayRouteArray) ToGatewayRouteArrayOutputWithContext(ctx context.Context) GatewayRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayRouteArrayOutput)
}

// GatewayRouteMapInput is an input type that accepts GatewayRouteMap and GatewayRouteMapOutput values.
// You can construct a concrete instance of `GatewayRouteMapInput` via:
//
//	GatewayRouteMap{ "key": GatewayRouteArgs{...} }
type GatewayRouteMapInput interface {
	pulumi.Input

	ToGatewayRouteMapOutput() GatewayRouteMapOutput
	ToGatewayRouteMapOutputWithContext(context.Context) GatewayRouteMapOutput
}

type GatewayRouteMap map[string]GatewayRouteInput

func (GatewayRouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayRoute)(nil)).Elem()
}

func (i GatewayRouteMap) ToGatewayRouteMapOutput() GatewayRouteMapOutput {
	return i.ToGatewayRouteMapOutputWithContext(context.Background())
}

func (i GatewayRouteMap) ToGatewayRouteMapOutputWithContext(ctx context.Context) GatewayRouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayRouteMapOutput)
}

type GatewayRouteOutput struct{ *pulumi.OutputState }

func (GatewayRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayRoute)(nil)).Elem()
}

func (o GatewayRouteOutput) ToGatewayRouteOutput() GatewayRouteOutput {
	return o
}

func (o GatewayRouteOutput) ToGatewayRouteOutputWithContext(ctx context.Context) GatewayRouteOutput {
	return o
}

// ARN of the gateway route.
func (o GatewayRouteOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Creation date of the gateway route.
func (o GatewayRouteOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

// Last update date of the gateway route.
func (o GatewayRouteOutput) LastUpdatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringOutput { return v.LastUpdatedDate }).(pulumi.StringOutput)
}

// Name of the service mesh in which to create the gateway route. Must be between 1 and 255 characters in length.
func (o GatewayRouteOutput) MeshName() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringOutput { return v.MeshName }).(pulumi.StringOutput)
}

// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
func (o GatewayRouteOutput) MeshOwner() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringOutput { return v.MeshOwner }).(pulumi.StringOutput)
}

// Name to use for the gateway route. Must be between 1 and 255 characters in length.
func (o GatewayRouteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource owner's AWS account ID.
func (o GatewayRouteOutput) ResourceOwner() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringOutput { return v.ResourceOwner }).(pulumi.StringOutput)
}

// Gateway route specification to apply.
func (o GatewayRouteOutput) Spec() GatewayRouteSpecOutput {
	return o.ApplyT(func(v *GatewayRoute) GatewayRouteSpecOutput { return v.Spec }).(GatewayRouteSpecOutput)
}

// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o GatewayRouteOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o GatewayRouteOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Name of the virtual gateway to associate the gateway route with. Must be between 1 and 255 characters in length.
func (o GatewayRouteOutput) VirtualGatewayName() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRoute) pulumi.StringOutput { return v.VirtualGatewayName }).(pulumi.StringOutput)
}

type GatewayRouteArrayOutput struct{ *pulumi.OutputState }

func (GatewayRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GatewayRoute)(nil)).Elem()
}

func (o GatewayRouteArrayOutput) ToGatewayRouteArrayOutput() GatewayRouteArrayOutput {
	return o
}

func (o GatewayRouteArrayOutput) ToGatewayRouteArrayOutputWithContext(ctx context.Context) GatewayRouteArrayOutput {
	return o
}

func (o GatewayRouteArrayOutput) Index(i pulumi.IntInput) GatewayRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GatewayRoute {
		return vs[0].([]*GatewayRoute)[vs[1].(int)]
	}).(GatewayRouteOutput)
}

type GatewayRouteMapOutput struct{ *pulumi.OutputState }

func (GatewayRouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GatewayRoute)(nil)).Elem()
}

func (o GatewayRouteMapOutput) ToGatewayRouteMapOutput() GatewayRouteMapOutput {
	return o
}

func (o GatewayRouteMapOutput) ToGatewayRouteMapOutputWithContext(ctx context.Context) GatewayRouteMapOutput {
	return o
}

func (o GatewayRouteMapOutput) MapIndex(k pulumi.StringInput) GatewayRouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GatewayRoute {
		return vs[0].(map[string]*GatewayRoute)[vs[1].(string)]
	}).(GatewayRouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayRouteInput)(nil)).Elem(), &GatewayRoute{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayRouteArrayInput)(nil)).Elem(), GatewayRouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayRouteMapInput)(nil)).Elem(), GatewayRouteMap{})
	pulumi.RegisterOutputType(GatewayRouteOutput{})
	pulumi.RegisterOutputType(GatewayRouteArrayOutput{})
	pulumi.RegisterOutputType(GatewayRouteMapOutput{})
}
