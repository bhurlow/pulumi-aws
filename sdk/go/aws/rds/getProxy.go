// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about a DB Proxy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rds.LookupProxy(ctx, &rds.LookupProxyArgs{
//				Name: "my-test-db-proxy",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupProxy(ctx *pulumi.Context, args *LookupProxyArgs, opts ...pulumi.InvokeOption) (*LookupProxyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProxyResult
	err := ctx.Invoke("aws:rds/getProxy:getProxy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProxy.
type LookupProxyArgs struct {
	// Name of the DB proxy.
	Name string `pulumi:"name"`
}

// A collection of values returned by getProxy.
type LookupProxyResult struct {
	// ARN of the DB Proxy.
	Arn string `pulumi:"arn"`
	// Configuration(s) with authorization mechanisms to connect to the associated instance or cluster.
	Auths []GetProxyAuth `pulumi:"auths"`
	// Whether the proxy includes detailed information about SQL statements in its logs.
	DebugLogging bool `pulumi:"debugLogging"`
	// Endpoint that you can use to connect to the DB proxy.
	Endpoint string `pulumi:"endpoint"`
	// Kinds of databases that the proxy can connect to.
	EngineFamily string `pulumi:"engineFamily"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Number of seconds a connection to the proxy can have no activity before the proxy drops the client connection.
	IdleClientTimeout int    `pulumi:"idleClientTimeout"`
	Name              string `pulumi:"name"`
	// Whether Transport Layer Security (TLS) encryption is required for connections to the proxy.
	RequireTls bool `pulumi:"requireTls"`
	// ARN for the IAM role that the proxy uses to access Amazon Secrets Manager.
	RoleArn string `pulumi:"roleArn"`
	// Provides the VPC ID of the DB proxy.
	VpcId string `pulumi:"vpcId"`
	// Provides a list of VPC security groups that the proxy belongs to.
	VpcSecurityGroupIds []string `pulumi:"vpcSecurityGroupIds"`
	// EC2 subnet IDs for the proxy.
	VpcSubnetIds []string `pulumi:"vpcSubnetIds"`
}

func LookupProxyOutput(ctx *pulumi.Context, args LookupProxyOutputArgs, opts ...pulumi.InvokeOption) LookupProxyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProxyResult, error) {
			args := v.(LookupProxyArgs)
			r, err := LookupProxy(ctx, &args, opts...)
			var s LookupProxyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProxyResultOutput)
}

// A collection of arguments for invoking getProxy.
type LookupProxyOutputArgs struct {
	// Name of the DB proxy.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupProxyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProxyArgs)(nil)).Elem()
}

// A collection of values returned by getProxy.
type LookupProxyResultOutput struct{ *pulumi.OutputState }

func (LookupProxyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProxyResult)(nil)).Elem()
}

func (o LookupProxyResultOutput) ToLookupProxyResultOutput() LookupProxyResultOutput {
	return o
}

func (o LookupProxyResultOutput) ToLookupProxyResultOutputWithContext(ctx context.Context) LookupProxyResultOutput {
	return o
}

// ARN of the DB Proxy.
func (o LookupProxyResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyResult) string { return v.Arn }).(pulumi.StringOutput)
}

// Configuration(s) with authorization mechanisms to connect to the associated instance or cluster.
func (o LookupProxyResultOutput) Auths() GetProxyAuthArrayOutput {
	return o.ApplyT(func(v LookupProxyResult) []GetProxyAuth { return v.Auths }).(GetProxyAuthArrayOutput)
}

// Whether the proxy includes detailed information about SQL statements in its logs.
func (o LookupProxyResultOutput) DebugLogging() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProxyResult) bool { return v.DebugLogging }).(pulumi.BoolOutput)
}

// Endpoint that you can use to connect to the DB proxy.
func (o LookupProxyResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

// Kinds of databases that the proxy can connect to.
func (o LookupProxyResultOutput) EngineFamily() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyResult) string { return v.EngineFamily }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupProxyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Number of seconds a connection to the proxy can have no activity before the proxy drops the client connection.
func (o LookupProxyResultOutput) IdleClientTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v LookupProxyResult) int { return v.IdleClientTimeout }).(pulumi.IntOutput)
}

func (o LookupProxyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Whether Transport Layer Security (TLS) encryption is required for connections to the proxy.
func (o LookupProxyResultOutput) RequireTls() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupProxyResult) bool { return v.RequireTls }).(pulumi.BoolOutput)
}

// ARN for the IAM role that the proxy uses to access Amazon Secrets Manager.
func (o LookupProxyResultOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyResult) string { return v.RoleArn }).(pulumi.StringOutput)
}

// Provides the VPC ID of the DB proxy.
func (o LookupProxyResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProxyResult) string { return v.VpcId }).(pulumi.StringOutput)
}

// Provides a list of VPC security groups that the proxy belongs to.
func (o LookupProxyResultOutput) VpcSecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupProxyResult) []string { return v.VpcSecurityGroupIds }).(pulumi.StringArrayOutput)
}

// EC2 subnet IDs for the proxy.
func (o LookupProxyResultOutput) VpcSubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupProxyResult) []string { return v.VpcSubnetIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProxyResultOutput{})
}
