// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicediscovery

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves information about a Service Discovery private or public DNS namespace.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicediscovery"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := servicediscovery.GetDnsNamespace(ctx, &servicediscovery.GetDnsNamespaceArgs{
//				Name: "example.service.local",
//				Type: "DNS_PRIVATE",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetDnsNamespace(ctx *pulumi.Context, args *GetDnsNamespaceArgs, opts ...pulumi.InvokeOption) (*GetDnsNamespaceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDnsNamespaceResult
	err := ctx.Invoke("aws:servicediscovery/getDnsNamespace:getDnsNamespace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDnsNamespace.
type GetDnsNamespaceArgs struct {
	// Name of the namespace.
	Name string `pulumi:"name"`
	// Map of tags for the resource.
	Tags map[string]string `pulumi:"tags"`
	// Type of the namespace. Allowed values are `DNS_PUBLIC` or `DNS_PRIVATE`.
	Type string `pulumi:"type"`
}

// A collection of values returned by getDnsNamespace.
type GetDnsNamespaceResult struct {
	// ARN of the namespace.
	Arn string `pulumi:"arn"`
	// Description of the namespace.
	Description string `pulumi:"description"`
	// ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
	HostedZone string `pulumi:"hostedZone"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Map of tags for the resource.
	Tags map[string]string `pulumi:"tags"`
	Type string            `pulumi:"type"`
}

func GetDnsNamespaceOutput(ctx *pulumi.Context, args GetDnsNamespaceOutputArgs, opts ...pulumi.InvokeOption) GetDnsNamespaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDnsNamespaceResult, error) {
			args := v.(GetDnsNamespaceArgs)
			r, err := GetDnsNamespace(ctx, &args, opts...)
			var s GetDnsNamespaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDnsNamespaceResultOutput)
}

// A collection of arguments for invoking getDnsNamespace.
type GetDnsNamespaceOutputArgs struct {
	// Name of the namespace.
	Name pulumi.StringInput `pulumi:"name"`
	// Map of tags for the resource.
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// Type of the namespace. Allowed values are `DNS_PUBLIC` or `DNS_PRIVATE`.
	Type pulumi.StringInput `pulumi:"type"`
}

func (GetDnsNamespaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDnsNamespaceArgs)(nil)).Elem()
}

// A collection of values returned by getDnsNamespace.
type GetDnsNamespaceResultOutput struct{ *pulumi.OutputState }

func (GetDnsNamespaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDnsNamespaceResult)(nil)).Elem()
}

func (o GetDnsNamespaceResultOutput) ToGetDnsNamespaceResultOutput() GetDnsNamespaceResultOutput {
	return o
}

func (o GetDnsNamespaceResultOutput) ToGetDnsNamespaceResultOutputWithContext(ctx context.Context) GetDnsNamespaceResultOutput {
	return o
}

// ARN of the namespace.
func (o GetDnsNamespaceResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v GetDnsNamespaceResult) string { return v.Arn }).(pulumi.StringOutput)
}

// Description of the namespace.
func (o GetDnsNamespaceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetDnsNamespaceResult) string { return v.Description }).(pulumi.StringOutput)
}

// ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
func (o GetDnsNamespaceResultOutput) HostedZone() pulumi.StringOutput {
	return o.ApplyT(func(v GetDnsNamespaceResult) string { return v.HostedZone }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDnsNamespaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDnsNamespaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDnsNamespaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDnsNamespaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Map of tags for the resource.
func (o GetDnsNamespaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetDnsNamespaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetDnsNamespaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetDnsNamespaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDnsNamespaceResultOutput{})
}
