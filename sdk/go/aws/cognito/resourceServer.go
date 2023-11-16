// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cognito Resource Server.
//
// ## Example Usage
// ### Create a basic resource server
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			pool, err := cognito.NewUserPool(ctx, "pool", nil)
//			if err != nil {
//				return err
//			}
//			_, err = cognito.NewResourceServer(ctx, "resource", &cognito.ResourceServerArgs{
//				Identifier: pulumi.String("https://example.com"),
//				UserPoolId: pool.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Create a resource server with sample-scope
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			pool, err := cognito.NewUserPool(ctx, "pool", nil)
//			if err != nil {
//				return err
//			}
//			_, err = cognito.NewResourceServer(ctx, "resource", &cognito.ResourceServerArgs{
//				Identifier: pulumi.String("https://example.com"),
//				Scopes: cognito.ResourceServerScopeArray{
//					&cognito.ResourceServerScopeArgs{
//						ScopeName:        pulumi.String("sample-scope"),
//						ScopeDescription: pulumi.String("a Sample Scope Description"),
//					},
//				},
//				UserPoolId: pool.ID(),
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
// Using `pulumi import`, import `aws_cognito_resource_server` using their User Pool ID and Identifier. For example:
//
// ```sh
//
//	$ pulumi import aws:cognito/resourceServer:ResourceServer example "us-west-2_abc123|https://example.com"
//
// ```
type ResourceServer struct {
	pulumi.CustomResourceState

	// An identifier for the resource server.
	Identifier pulumi.StringOutput `pulumi:"identifier"`
	// A name for the resource server.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of all scopes configured for this resource server in the format identifier/scope_name.
	ScopeIdentifiers pulumi.StringArrayOutput `pulumi:"scopeIdentifiers"`
	// A list of Authorization Scope.
	Scopes     ResourceServerScopeArrayOutput `pulumi:"scopes"`
	UserPoolId pulumi.StringOutput            `pulumi:"userPoolId"`
}

// NewResourceServer registers a new resource with the given unique name, arguments, and options.
func NewResourceServer(ctx *pulumi.Context,
	name string, args *ResourceServerArgs, opts ...pulumi.ResourceOption) (*ResourceServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Identifier == nil {
		return nil, errors.New("invalid value for required argument 'Identifier'")
	}
	if args.UserPoolId == nil {
		return nil, errors.New("invalid value for required argument 'UserPoolId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResourceServer
	err := ctx.RegisterResource("aws:cognito/resourceServer:ResourceServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceServer gets an existing ResourceServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceServerState, opts ...pulumi.ResourceOption) (*ResourceServer, error) {
	var resource ResourceServer
	err := ctx.ReadResource("aws:cognito/resourceServer:ResourceServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceServer resources.
type resourceServerState struct {
	// An identifier for the resource server.
	Identifier *string `pulumi:"identifier"`
	// A name for the resource server.
	Name *string `pulumi:"name"`
	// A list of all scopes configured for this resource server in the format identifier/scope_name.
	ScopeIdentifiers []string `pulumi:"scopeIdentifiers"`
	// A list of Authorization Scope.
	Scopes     []ResourceServerScope `pulumi:"scopes"`
	UserPoolId *string               `pulumi:"userPoolId"`
}

type ResourceServerState struct {
	// An identifier for the resource server.
	Identifier pulumi.StringPtrInput
	// A name for the resource server.
	Name pulumi.StringPtrInput
	// A list of all scopes configured for this resource server in the format identifier/scope_name.
	ScopeIdentifiers pulumi.StringArrayInput
	// A list of Authorization Scope.
	Scopes     ResourceServerScopeArrayInput
	UserPoolId pulumi.StringPtrInput
}

func (ResourceServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceServerState)(nil)).Elem()
}

type resourceServerArgs struct {
	// An identifier for the resource server.
	Identifier string `pulumi:"identifier"`
	// A name for the resource server.
	Name *string `pulumi:"name"`
	// A list of Authorization Scope.
	Scopes     []ResourceServerScope `pulumi:"scopes"`
	UserPoolId string                `pulumi:"userPoolId"`
}

// The set of arguments for constructing a ResourceServer resource.
type ResourceServerArgs struct {
	// An identifier for the resource server.
	Identifier pulumi.StringInput
	// A name for the resource server.
	Name pulumi.StringPtrInput
	// A list of Authorization Scope.
	Scopes     ResourceServerScopeArrayInput
	UserPoolId pulumi.StringInput
}

func (ResourceServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceServerArgs)(nil)).Elem()
}

type ResourceServerInput interface {
	pulumi.Input

	ToResourceServerOutput() ResourceServerOutput
	ToResourceServerOutputWithContext(ctx context.Context) ResourceServerOutput
}

func (*ResourceServer) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceServer)(nil)).Elem()
}

func (i *ResourceServer) ToResourceServerOutput() ResourceServerOutput {
	return i.ToResourceServerOutputWithContext(context.Background())
}

func (i *ResourceServer) ToResourceServerOutputWithContext(ctx context.Context) ResourceServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceServerOutput)
}

// ResourceServerArrayInput is an input type that accepts ResourceServerArray and ResourceServerArrayOutput values.
// You can construct a concrete instance of `ResourceServerArrayInput` via:
//
//	ResourceServerArray{ ResourceServerArgs{...} }
type ResourceServerArrayInput interface {
	pulumi.Input

	ToResourceServerArrayOutput() ResourceServerArrayOutput
	ToResourceServerArrayOutputWithContext(context.Context) ResourceServerArrayOutput
}

type ResourceServerArray []ResourceServerInput

func (ResourceServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceServer)(nil)).Elem()
}

func (i ResourceServerArray) ToResourceServerArrayOutput() ResourceServerArrayOutput {
	return i.ToResourceServerArrayOutputWithContext(context.Background())
}

func (i ResourceServerArray) ToResourceServerArrayOutputWithContext(ctx context.Context) ResourceServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceServerArrayOutput)
}

// ResourceServerMapInput is an input type that accepts ResourceServerMap and ResourceServerMapOutput values.
// You can construct a concrete instance of `ResourceServerMapInput` via:
//
//	ResourceServerMap{ "key": ResourceServerArgs{...} }
type ResourceServerMapInput interface {
	pulumi.Input

	ToResourceServerMapOutput() ResourceServerMapOutput
	ToResourceServerMapOutputWithContext(context.Context) ResourceServerMapOutput
}

type ResourceServerMap map[string]ResourceServerInput

func (ResourceServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceServer)(nil)).Elem()
}

func (i ResourceServerMap) ToResourceServerMapOutput() ResourceServerMapOutput {
	return i.ToResourceServerMapOutputWithContext(context.Background())
}

func (i ResourceServerMap) ToResourceServerMapOutputWithContext(ctx context.Context) ResourceServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceServerMapOutput)
}

type ResourceServerOutput struct{ *pulumi.OutputState }

func (ResourceServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceServer)(nil)).Elem()
}

func (o ResourceServerOutput) ToResourceServerOutput() ResourceServerOutput {
	return o
}

func (o ResourceServerOutput) ToResourceServerOutputWithContext(ctx context.Context) ResourceServerOutput {
	return o
}

// An identifier for the resource server.
func (o ResourceServerOutput) Identifier() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceServer) pulumi.StringOutput { return v.Identifier }).(pulumi.StringOutput)
}

// A name for the resource server.
func (o ResourceServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of all scopes configured for this resource server in the format identifier/scope_name.
func (o ResourceServerOutput) ScopeIdentifiers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResourceServer) pulumi.StringArrayOutput { return v.ScopeIdentifiers }).(pulumi.StringArrayOutput)
}

// A list of Authorization Scope.
func (o ResourceServerOutput) Scopes() ResourceServerScopeArrayOutput {
	return o.ApplyT(func(v *ResourceServer) ResourceServerScopeArrayOutput { return v.Scopes }).(ResourceServerScopeArrayOutput)
}

func (o ResourceServerOutput) UserPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourceServer) pulumi.StringOutput { return v.UserPoolId }).(pulumi.StringOutput)
}

type ResourceServerArrayOutput struct{ *pulumi.OutputState }

func (ResourceServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourceServer)(nil)).Elem()
}

func (o ResourceServerArrayOutput) ToResourceServerArrayOutput() ResourceServerArrayOutput {
	return o
}

func (o ResourceServerArrayOutput) ToResourceServerArrayOutputWithContext(ctx context.Context) ResourceServerArrayOutput {
	return o
}

func (o ResourceServerArrayOutput) Index(i pulumi.IntInput) ResourceServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourceServer {
		return vs[0].([]*ResourceServer)[vs[1].(int)]
	}).(ResourceServerOutput)
}

type ResourceServerMapOutput struct{ *pulumi.OutputState }

func (ResourceServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourceServer)(nil)).Elem()
}

func (o ResourceServerMapOutput) ToResourceServerMapOutput() ResourceServerMapOutput {
	return o
}

func (o ResourceServerMapOutput) ToResourceServerMapOutputWithContext(ctx context.Context) ResourceServerMapOutput {
	return o
}

func (o ResourceServerMapOutput) MapIndex(k pulumi.StringInput) ResourceServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourceServer {
		return vs[0].(map[string]*ResourceServer)[vs[1].(string)]
	}).(ResourceServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceServerInput)(nil)).Elem(), &ResourceServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceServerArrayInput)(nil)).Elem(), ResourceServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceServerMapInput)(nil)).Elem(), ResourceServerMap{})
	pulumi.RegisterOutputType(ResourceServerOutput{})
	pulumi.RegisterOutputType(ResourceServerArrayOutput{})
	pulumi.RegisterOutputType(ResourceServerMapOutput{})
}
