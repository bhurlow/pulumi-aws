// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an AWS App Mesh route resource.
//
// ## Example Usage
// ### HTTP Routing
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
//			_, err := appmesh.NewRoute(ctx, "serviceb", &appmesh.RouteArgs{
//				MeshName:          pulumi.Any(aws_appmesh_mesh.Simple.Id),
//				VirtualRouterName: pulumi.Any(aws_appmesh_virtual_router.Serviceb.Name),
//				Spec: &appmesh.RouteSpecArgs{
//					HttpRoute: &appmesh.RouteSpecHttpRouteArgs{
//						Match: &appmesh.RouteSpecHttpRouteMatchArgs{
//							Prefix: pulumi.String("/"),
//						},
//						Action: &appmesh.RouteSpecHttpRouteActionArgs{
//							WeightedTargets: appmesh.RouteSpecHttpRouteActionWeightedTargetArray{
//								&appmesh.RouteSpecHttpRouteActionWeightedTargetArgs{
//									VirtualNode: pulumi.Any(aws_appmesh_virtual_node.Serviceb1.Name),
//									Weight:      pulumi.Int(90),
//								},
//								&appmesh.RouteSpecHttpRouteActionWeightedTargetArgs{
//									VirtualNode: pulumi.Any(aws_appmesh_virtual_node.Serviceb2.Name),
//									Weight:      pulumi.Int(10),
//								},
//							},
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
// ### HTTP Header Routing
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
//			_, err := appmesh.NewRoute(ctx, "serviceb", &appmesh.RouteArgs{
//				MeshName:          pulumi.Any(aws_appmesh_mesh.Simple.Id),
//				VirtualRouterName: pulumi.Any(aws_appmesh_virtual_router.Serviceb.Name),
//				Spec: &appmesh.RouteSpecArgs{
//					HttpRoute: &appmesh.RouteSpecHttpRouteArgs{
//						Match: &appmesh.RouteSpecHttpRouteMatchArgs{
//							Method: pulumi.String("POST"),
//							Prefix: pulumi.String("/"),
//							Scheme: pulumi.String("https"),
//							Headers: appmesh.RouteSpecHttpRouteMatchHeaderArray{
//								&appmesh.RouteSpecHttpRouteMatchHeaderArgs{
//									Name: pulumi.String("clientRequestId"),
//									Match: &appmesh.RouteSpecHttpRouteMatchHeaderMatchArgs{
//										Prefix: pulumi.String("123"),
//									},
//								},
//							},
//						},
//						Action: &appmesh.RouteSpecHttpRouteActionArgs{
//							WeightedTargets: appmesh.RouteSpecHttpRouteActionWeightedTargetArray{
//								&appmesh.RouteSpecHttpRouteActionWeightedTargetArgs{
//									VirtualNode: pulumi.Any(aws_appmesh_virtual_node.Serviceb.Name),
//									Weight:      pulumi.Int(100),
//								},
//							},
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
// ### Retry Policy
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
//			_, err := appmesh.NewRoute(ctx, "serviceb", &appmesh.RouteArgs{
//				MeshName:          pulumi.Any(aws_appmesh_mesh.Simple.Id),
//				VirtualRouterName: pulumi.Any(aws_appmesh_virtual_router.Serviceb.Name),
//				Spec: &appmesh.RouteSpecArgs{
//					HttpRoute: &appmesh.RouteSpecHttpRouteArgs{
//						Match: &appmesh.RouteSpecHttpRouteMatchArgs{
//							Prefix: pulumi.String("/"),
//						},
//						RetryPolicy: &appmesh.RouteSpecHttpRouteRetryPolicyArgs{
//							HttpRetryEvents: pulumi.StringArray{
//								pulumi.String("server-error"),
//							},
//							MaxRetries: pulumi.Int(1),
//							PerRetryTimeout: &appmesh.RouteSpecHttpRouteRetryPolicyPerRetryTimeoutArgs{
//								Unit:  pulumi.String("s"),
//								Value: pulumi.Int(15),
//							},
//						},
//						Action: &appmesh.RouteSpecHttpRouteActionArgs{
//							WeightedTargets: appmesh.RouteSpecHttpRouteActionWeightedTargetArray{
//								&appmesh.RouteSpecHttpRouteActionWeightedTargetArgs{
//									VirtualNode: pulumi.Any(aws_appmesh_virtual_node.Serviceb.Name),
//									Weight:      pulumi.Int(100),
//								},
//							},
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
// ### TCP Routing
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
//			_, err := appmesh.NewRoute(ctx, "serviceb", &appmesh.RouteArgs{
//				MeshName:          pulumi.Any(aws_appmesh_mesh.Simple.Id),
//				VirtualRouterName: pulumi.Any(aws_appmesh_virtual_router.Serviceb.Name),
//				Spec: &appmesh.RouteSpecArgs{
//					TcpRoute: &appmesh.RouteSpecTcpRouteArgs{
//						Action: &appmesh.RouteSpecTcpRouteActionArgs{
//							WeightedTargets: appmesh.RouteSpecTcpRouteActionWeightedTargetArray{
//								&appmesh.RouteSpecTcpRouteActionWeightedTargetArgs{
//									VirtualNode: pulumi.Any(aws_appmesh_virtual_node.Serviceb1.Name),
//									Weight:      pulumi.Int(100),
//								},
//							},
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
// App Mesh virtual routes can be imported using `mesh_name` and `virtual_router_name` together with the route's `name`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:appmesh/route:Route serviceb simpleapp/serviceB/serviceB-route
//
// ```
type Route struct {
	pulumi.CustomResourceState

	// ARN of the route.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Creation date of the route.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// Last update date of the route.
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`
	// Name of the service mesh in which to create the route. Must be between 1 and 255 characters in length.
	MeshName pulumi.StringOutput `pulumi:"meshName"`
	// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
	MeshOwner pulumi.StringOutput `pulumi:"meshOwner"`
	// Name to use for the route. Must be between 1 and 255 characters in length.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource owner's AWS account ID.
	ResourceOwner pulumi.StringOutput `pulumi:"resourceOwner"`
	// Route specification to apply.
	Spec RouteSpecOutput `pulumi:"spec"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Name of the virtual router in which to create the route. Must be between 1 and 255 characters in length.
	VirtualRouterName pulumi.StringOutput `pulumi:"virtualRouterName"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MeshName == nil {
		return nil, errors.New("invalid value for required argument 'MeshName'")
	}
	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	if args.VirtualRouterName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualRouterName'")
	}
	var resource Route
	err := ctx.RegisterResource("aws:appmesh/route:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("aws:appmesh/route:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type routeState struct {
	// ARN of the route.
	Arn *string `pulumi:"arn"`
	// Creation date of the route.
	CreatedDate *string `pulumi:"createdDate"`
	// Last update date of the route.
	LastUpdatedDate *string `pulumi:"lastUpdatedDate"`
	// Name of the service mesh in which to create the route. Must be between 1 and 255 characters in length.
	MeshName *string `pulumi:"meshName"`
	// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
	MeshOwner *string `pulumi:"meshOwner"`
	// Name to use for the route. Must be between 1 and 255 characters in length.
	Name *string `pulumi:"name"`
	// Resource owner's AWS account ID.
	ResourceOwner *string `pulumi:"resourceOwner"`
	// Route specification to apply.
	Spec *RouteSpec `pulumi:"spec"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Name of the virtual router in which to create the route. Must be between 1 and 255 characters in length.
	VirtualRouterName *string `pulumi:"virtualRouterName"`
}

type RouteState struct {
	// ARN of the route.
	Arn pulumi.StringPtrInput
	// Creation date of the route.
	CreatedDate pulumi.StringPtrInput
	// Last update date of the route.
	LastUpdatedDate pulumi.StringPtrInput
	// Name of the service mesh in which to create the route. Must be between 1 and 255 characters in length.
	MeshName pulumi.StringPtrInput
	// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
	MeshOwner pulumi.StringPtrInput
	// Name to use for the route. Must be between 1 and 255 characters in length.
	Name pulumi.StringPtrInput
	// Resource owner's AWS account ID.
	ResourceOwner pulumi.StringPtrInput
	// Route specification to apply.
	Spec RouteSpecPtrInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// Name of the virtual router in which to create the route. Must be between 1 and 255 characters in length.
	VirtualRouterName pulumi.StringPtrInput
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	// Name of the service mesh in which to create the route. Must be between 1 and 255 characters in length.
	MeshName string `pulumi:"meshName"`
	// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
	MeshOwner *string `pulumi:"meshOwner"`
	// Name to use for the route. Must be between 1 and 255 characters in length.
	Name *string `pulumi:"name"`
	// Route specification to apply.
	Spec RouteSpec `pulumi:"spec"`
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Name of the virtual router in which to create the route. Must be between 1 and 255 characters in length.
	VirtualRouterName string `pulumi:"virtualRouterName"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// Name of the service mesh in which to create the route. Must be between 1 and 255 characters in length.
	MeshName pulumi.StringInput
	// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
	MeshOwner pulumi.StringPtrInput
	// Name to use for the route. Must be between 1 and 255 characters in length.
	Name pulumi.StringPtrInput
	// Route specification to apply.
	Spec RouteSpecInput
	// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// Name of the virtual router in which to create the route. Must be between 1 and 255 characters in length.
	VirtualRouterName pulumi.StringInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}

type RouteInput interface {
	pulumi.Input

	ToRouteOutput() RouteOutput
	ToRouteOutputWithContext(ctx context.Context) RouteOutput
}

func (*Route) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (i *Route) ToRouteOutput() RouteOutput {
	return i.ToRouteOutputWithContext(context.Background())
}

func (i *Route) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteOutput)
}

// RouteArrayInput is an input type that accepts RouteArray and RouteArrayOutput values.
// You can construct a concrete instance of `RouteArrayInput` via:
//
//	RouteArray{ RouteArgs{...} }
type RouteArrayInput interface {
	pulumi.Input

	ToRouteArrayOutput() RouteArrayOutput
	ToRouteArrayOutputWithContext(context.Context) RouteArrayOutput
}

type RouteArray []RouteInput

func (RouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Route)(nil)).Elem()
}

func (i RouteArray) ToRouteArrayOutput() RouteArrayOutput {
	return i.ToRouteArrayOutputWithContext(context.Background())
}

func (i RouteArray) ToRouteArrayOutputWithContext(ctx context.Context) RouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteArrayOutput)
}

// RouteMapInput is an input type that accepts RouteMap and RouteMapOutput values.
// You can construct a concrete instance of `RouteMapInput` via:
//
//	RouteMap{ "key": RouteArgs{...} }
type RouteMapInput interface {
	pulumi.Input

	ToRouteMapOutput() RouteMapOutput
	ToRouteMapOutputWithContext(context.Context) RouteMapOutput
}

type RouteMap map[string]RouteInput

func (RouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Route)(nil)).Elem()
}

func (i RouteMap) ToRouteMapOutput() RouteMapOutput {
	return i.ToRouteMapOutputWithContext(context.Background())
}

func (i RouteMap) ToRouteMapOutputWithContext(ctx context.Context) RouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteMapOutput)
}

type RouteOutput struct{ *pulumi.OutputState }

func (RouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (o RouteOutput) ToRouteOutput() RouteOutput {
	return o
}

func (o RouteOutput) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return o
}

// ARN of the route.
func (o RouteOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Creation date of the route.
func (o RouteOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

// Last update date of the route.
func (o RouteOutput) LastUpdatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.LastUpdatedDate }).(pulumi.StringOutput)
}

// Name of the service mesh in which to create the route. Must be between 1 and 255 characters in length.
func (o RouteOutput) MeshName() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.MeshName }).(pulumi.StringOutput)
}

// AWS account ID of the service mesh's owner. Defaults to the account ID the AWS provider is currently connected to.
func (o RouteOutput) MeshOwner() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.MeshOwner }).(pulumi.StringOutput)
}

// Name to use for the route. Must be between 1 and 255 characters in length.
func (o RouteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Resource owner's AWS account ID.
func (o RouteOutput) ResourceOwner() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.ResourceOwner }).(pulumi.StringOutput)
}

// Route specification to apply.
func (o RouteOutput) Spec() RouteSpecOutput {
	return o.ApplyT(func(v *Route) RouteSpecOutput { return v.Spec }).(RouteSpecOutput)
}

// Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o RouteOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Route) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o RouteOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Route) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Name of the virtual router in which to create the route. Must be between 1 and 255 characters in length.
func (o RouteOutput) VirtualRouterName() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.VirtualRouterName }).(pulumi.StringOutput)
}

type RouteArrayOutput struct{ *pulumi.OutputState }

func (RouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Route)(nil)).Elem()
}

func (o RouteArrayOutput) ToRouteArrayOutput() RouteArrayOutput {
	return o
}

func (o RouteArrayOutput) ToRouteArrayOutputWithContext(ctx context.Context) RouteArrayOutput {
	return o
}

func (o RouteArrayOutput) Index(i pulumi.IntInput) RouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Route {
		return vs[0].([]*Route)[vs[1].(int)]
	}).(RouteOutput)
}

type RouteMapOutput struct{ *pulumi.OutputState }

func (RouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Route)(nil)).Elem()
}

func (o RouteMapOutput) ToRouteMapOutput() RouteMapOutput {
	return o
}

func (o RouteMapOutput) ToRouteMapOutputWithContext(ctx context.Context) RouteMapOutput {
	return o
}

func (o RouteMapOutput) MapIndex(k pulumi.StringInput) RouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Route {
		return vs[0].(map[string]*Route)[vs[1].(string)]
	}).(RouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteInput)(nil)).Elem(), &Route{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteArrayInput)(nil)).Elem(), RouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteMapInput)(nil)).Elem(), RouteMap{})
	pulumi.RegisterOutputType(RouteOutput{})
	pulumi.RegisterOutputType(RouteArrayOutput{})
	pulumi.RegisterOutputType(RouteMapOutput{})
}
