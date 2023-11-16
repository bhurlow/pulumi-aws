// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Amazon Redshift endpoint authorization.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshift"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := redshift.NewEndpointAuthorization(ctx, "example", &redshift.EndpointAuthorizationArgs{
//				Account:           pulumi.String("01234567910"),
//				ClusterIdentifier: pulumi.Any(aws_redshift_cluster.Example.Cluster_identifier),
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
// Using `pulumi import`, import Redshift endpoint authorization using the `id`. For example:
//
// ```sh
//
//	$ pulumi import aws:redshift/endpointAuthorization:EndpointAuthorization example 01234567910:cluster-example-id
//
// ```
type EndpointAuthorization struct {
	pulumi.CustomResourceState

	// The Amazon Web Services account ID to grant access to.
	Account pulumi.StringOutput `pulumi:"account"`
	// Indicates whether all VPCs in the grantee account are allowed access to the cluster.
	AllowedAllVpcs pulumi.BoolOutput `pulumi:"allowedAllVpcs"`
	// The cluster identifier of the cluster to grant access to.
	ClusterIdentifier pulumi.StringOutput `pulumi:"clusterIdentifier"`
	// The number of Redshift-managed VPC endpoints created for the authorization.
	EndpointCount pulumi.IntOutput `pulumi:"endpointCount"`
	// Indicates whether to force the revoke action. If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted. Default value is `false`.
	ForceDelete pulumi.BoolPtrOutput `pulumi:"forceDelete"`
	// The Amazon Web Services account ID of the grantee of the cluster.
	Grantee pulumi.StringOutput `pulumi:"grantee"`
	// The Amazon Web Services account ID of the cluster owner.
	Grantor pulumi.StringOutput `pulumi:"grantor"`
	// The virtual private cloud (VPC) identifiers to grant access to. If none are specified all VPCs in shared account are allowed.
	VpcIds pulumi.StringArrayOutput `pulumi:"vpcIds"`
}

// NewEndpointAuthorization registers a new resource with the given unique name, arguments, and options.
func NewEndpointAuthorization(ctx *pulumi.Context,
	name string, args *EndpointAuthorizationArgs, opts ...pulumi.ResourceOption) (*EndpointAuthorization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Account == nil {
		return nil, errors.New("invalid value for required argument 'Account'")
	}
	if args.ClusterIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'ClusterIdentifier'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EndpointAuthorization
	err := ctx.RegisterResource("aws:redshift/endpointAuthorization:EndpointAuthorization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointAuthorization gets an existing EndpointAuthorization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointAuthorizationState, opts ...pulumi.ResourceOption) (*EndpointAuthorization, error) {
	var resource EndpointAuthorization
	err := ctx.ReadResource("aws:redshift/endpointAuthorization:EndpointAuthorization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointAuthorization resources.
type endpointAuthorizationState struct {
	// The Amazon Web Services account ID to grant access to.
	Account *string `pulumi:"account"`
	// Indicates whether all VPCs in the grantee account are allowed access to the cluster.
	AllowedAllVpcs *bool `pulumi:"allowedAllVpcs"`
	// The cluster identifier of the cluster to grant access to.
	ClusterIdentifier *string `pulumi:"clusterIdentifier"`
	// The number of Redshift-managed VPC endpoints created for the authorization.
	EndpointCount *int `pulumi:"endpointCount"`
	// Indicates whether to force the revoke action. If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted. Default value is `false`.
	ForceDelete *bool `pulumi:"forceDelete"`
	// The Amazon Web Services account ID of the grantee of the cluster.
	Grantee *string `pulumi:"grantee"`
	// The Amazon Web Services account ID of the cluster owner.
	Grantor *string `pulumi:"grantor"`
	// The virtual private cloud (VPC) identifiers to grant access to. If none are specified all VPCs in shared account are allowed.
	VpcIds []string `pulumi:"vpcIds"`
}

type EndpointAuthorizationState struct {
	// The Amazon Web Services account ID to grant access to.
	Account pulumi.StringPtrInput
	// Indicates whether all VPCs in the grantee account are allowed access to the cluster.
	AllowedAllVpcs pulumi.BoolPtrInput
	// The cluster identifier of the cluster to grant access to.
	ClusterIdentifier pulumi.StringPtrInput
	// The number of Redshift-managed VPC endpoints created for the authorization.
	EndpointCount pulumi.IntPtrInput
	// Indicates whether to force the revoke action. If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted. Default value is `false`.
	ForceDelete pulumi.BoolPtrInput
	// The Amazon Web Services account ID of the grantee of the cluster.
	Grantee pulumi.StringPtrInput
	// The Amazon Web Services account ID of the cluster owner.
	Grantor pulumi.StringPtrInput
	// The virtual private cloud (VPC) identifiers to grant access to. If none are specified all VPCs in shared account are allowed.
	VpcIds pulumi.StringArrayInput
}

func (EndpointAuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointAuthorizationState)(nil)).Elem()
}

type endpointAuthorizationArgs struct {
	// The Amazon Web Services account ID to grant access to.
	Account string `pulumi:"account"`
	// The cluster identifier of the cluster to grant access to.
	ClusterIdentifier string `pulumi:"clusterIdentifier"`
	// Indicates whether to force the revoke action. If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted. Default value is `false`.
	ForceDelete *bool `pulumi:"forceDelete"`
	// The virtual private cloud (VPC) identifiers to grant access to. If none are specified all VPCs in shared account are allowed.
	VpcIds []string `pulumi:"vpcIds"`
}

// The set of arguments for constructing a EndpointAuthorization resource.
type EndpointAuthorizationArgs struct {
	// The Amazon Web Services account ID to grant access to.
	Account pulumi.StringInput
	// The cluster identifier of the cluster to grant access to.
	ClusterIdentifier pulumi.StringInput
	// Indicates whether to force the revoke action. If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted. Default value is `false`.
	ForceDelete pulumi.BoolPtrInput
	// The virtual private cloud (VPC) identifiers to grant access to. If none are specified all VPCs in shared account are allowed.
	VpcIds pulumi.StringArrayInput
}

func (EndpointAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointAuthorizationArgs)(nil)).Elem()
}

type EndpointAuthorizationInput interface {
	pulumi.Input

	ToEndpointAuthorizationOutput() EndpointAuthorizationOutput
	ToEndpointAuthorizationOutputWithContext(ctx context.Context) EndpointAuthorizationOutput
}

func (*EndpointAuthorization) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointAuthorization)(nil)).Elem()
}

func (i *EndpointAuthorization) ToEndpointAuthorizationOutput() EndpointAuthorizationOutput {
	return i.ToEndpointAuthorizationOutputWithContext(context.Background())
}

func (i *EndpointAuthorization) ToEndpointAuthorizationOutputWithContext(ctx context.Context) EndpointAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointAuthorizationOutput)
}

// EndpointAuthorizationArrayInput is an input type that accepts EndpointAuthorizationArray and EndpointAuthorizationArrayOutput values.
// You can construct a concrete instance of `EndpointAuthorizationArrayInput` via:
//
//	EndpointAuthorizationArray{ EndpointAuthorizationArgs{...} }
type EndpointAuthorizationArrayInput interface {
	pulumi.Input

	ToEndpointAuthorizationArrayOutput() EndpointAuthorizationArrayOutput
	ToEndpointAuthorizationArrayOutputWithContext(context.Context) EndpointAuthorizationArrayOutput
}

type EndpointAuthorizationArray []EndpointAuthorizationInput

func (EndpointAuthorizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointAuthorization)(nil)).Elem()
}

func (i EndpointAuthorizationArray) ToEndpointAuthorizationArrayOutput() EndpointAuthorizationArrayOutput {
	return i.ToEndpointAuthorizationArrayOutputWithContext(context.Background())
}

func (i EndpointAuthorizationArray) ToEndpointAuthorizationArrayOutputWithContext(ctx context.Context) EndpointAuthorizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointAuthorizationArrayOutput)
}

// EndpointAuthorizationMapInput is an input type that accepts EndpointAuthorizationMap and EndpointAuthorizationMapOutput values.
// You can construct a concrete instance of `EndpointAuthorizationMapInput` via:
//
//	EndpointAuthorizationMap{ "key": EndpointAuthorizationArgs{...} }
type EndpointAuthorizationMapInput interface {
	pulumi.Input

	ToEndpointAuthorizationMapOutput() EndpointAuthorizationMapOutput
	ToEndpointAuthorizationMapOutputWithContext(context.Context) EndpointAuthorizationMapOutput
}

type EndpointAuthorizationMap map[string]EndpointAuthorizationInput

func (EndpointAuthorizationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointAuthorization)(nil)).Elem()
}

func (i EndpointAuthorizationMap) ToEndpointAuthorizationMapOutput() EndpointAuthorizationMapOutput {
	return i.ToEndpointAuthorizationMapOutputWithContext(context.Background())
}

func (i EndpointAuthorizationMap) ToEndpointAuthorizationMapOutputWithContext(ctx context.Context) EndpointAuthorizationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointAuthorizationMapOutput)
}

type EndpointAuthorizationOutput struct{ *pulumi.OutputState }

func (EndpointAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointAuthorization)(nil)).Elem()
}

func (o EndpointAuthorizationOutput) ToEndpointAuthorizationOutput() EndpointAuthorizationOutput {
	return o
}

func (o EndpointAuthorizationOutput) ToEndpointAuthorizationOutputWithContext(ctx context.Context) EndpointAuthorizationOutput {
	return o
}

// The Amazon Web Services account ID to grant access to.
func (o EndpointAuthorizationOutput) Account() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointAuthorization) pulumi.StringOutput { return v.Account }).(pulumi.StringOutput)
}

// Indicates whether all VPCs in the grantee account are allowed access to the cluster.
func (o EndpointAuthorizationOutput) AllowedAllVpcs() pulumi.BoolOutput {
	return o.ApplyT(func(v *EndpointAuthorization) pulumi.BoolOutput { return v.AllowedAllVpcs }).(pulumi.BoolOutput)
}

// The cluster identifier of the cluster to grant access to.
func (o EndpointAuthorizationOutput) ClusterIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointAuthorization) pulumi.StringOutput { return v.ClusterIdentifier }).(pulumi.StringOutput)
}

// The number of Redshift-managed VPC endpoints created for the authorization.
func (o EndpointAuthorizationOutput) EndpointCount() pulumi.IntOutput {
	return o.ApplyT(func(v *EndpointAuthorization) pulumi.IntOutput { return v.EndpointCount }).(pulumi.IntOutput)
}

// Indicates whether to force the revoke action. If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted. Default value is `false`.
func (o EndpointAuthorizationOutput) ForceDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EndpointAuthorization) pulumi.BoolPtrOutput { return v.ForceDelete }).(pulumi.BoolPtrOutput)
}

// The Amazon Web Services account ID of the grantee of the cluster.
func (o EndpointAuthorizationOutput) Grantee() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointAuthorization) pulumi.StringOutput { return v.Grantee }).(pulumi.StringOutput)
}

// The Amazon Web Services account ID of the cluster owner.
func (o EndpointAuthorizationOutput) Grantor() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointAuthorization) pulumi.StringOutput { return v.Grantor }).(pulumi.StringOutput)
}

// The virtual private cloud (VPC) identifiers to grant access to. If none are specified all VPCs in shared account are allowed.
func (o EndpointAuthorizationOutput) VpcIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EndpointAuthorization) pulumi.StringArrayOutput { return v.VpcIds }).(pulumi.StringArrayOutput)
}

type EndpointAuthorizationArrayOutput struct{ *pulumi.OutputState }

func (EndpointAuthorizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointAuthorization)(nil)).Elem()
}

func (o EndpointAuthorizationArrayOutput) ToEndpointAuthorizationArrayOutput() EndpointAuthorizationArrayOutput {
	return o
}

func (o EndpointAuthorizationArrayOutput) ToEndpointAuthorizationArrayOutputWithContext(ctx context.Context) EndpointAuthorizationArrayOutput {
	return o
}

func (o EndpointAuthorizationArrayOutput) Index(i pulumi.IntInput) EndpointAuthorizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EndpointAuthorization {
		return vs[0].([]*EndpointAuthorization)[vs[1].(int)]
	}).(EndpointAuthorizationOutput)
}

type EndpointAuthorizationMapOutput struct{ *pulumi.OutputState }

func (EndpointAuthorizationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointAuthorization)(nil)).Elem()
}

func (o EndpointAuthorizationMapOutput) ToEndpointAuthorizationMapOutput() EndpointAuthorizationMapOutput {
	return o
}

func (o EndpointAuthorizationMapOutput) ToEndpointAuthorizationMapOutputWithContext(ctx context.Context) EndpointAuthorizationMapOutput {
	return o
}

func (o EndpointAuthorizationMapOutput) MapIndex(k pulumi.StringInput) EndpointAuthorizationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EndpointAuthorization {
		return vs[0].(map[string]*EndpointAuthorization)[vs[1].(string)]
	}).(EndpointAuthorizationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointAuthorizationInput)(nil)).Elem(), &EndpointAuthorization{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointAuthorizationArrayInput)(nil)).Elem(), EndpointAuthorizationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointAuthorizationMapInput)(nil)).Elem(), EndpointAuthorizationMap{})
	pulumi.RegisterOutputType(EndpointAuthorizationOutput{})
	pulumi.RegisterOutputType(EndpointAuthorizationArrayOutput{})
	pulumi.RegisterOutputType(EndpointAuthorizationMapOutput{})
}
