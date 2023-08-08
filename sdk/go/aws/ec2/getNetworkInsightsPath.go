// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ec2.NetworkInsightsPath` provides details about a specific Network Insights Path.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ec2.LookupNetworkInsightsPath(ctx, &ec2.LookupNetworkInsightsPathArgs{
//				NetworkInsightsPathId: pulumi.StringRef(aws_ec2_network_insights_path.Example.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupNetworkInsightsPath(ctx *pulumi.Context, args *LookupNetworkInsightsPathArgs, opts ...pulumi.InvokeOption) (*LookupNetworkInsightsPathResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkInsightsPathResult
	err := ctx.Invoke("aws:ec2/getNetworkInsightsPath:getNetworkInsightsPath", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetworkInsightsPath.
type LookupNetworkInsightsPathArgs struct {
	// Configuration block(s) for filtering. Detailed below.
	Filters []GetNetworkInsightsPathFilter `pulumi:"filters"`
	// ID of the Network Insights Path to select.
	NetworkInsightsPathId *string `pulumi:"networkInsightsPathId"`
	// Map of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getNetworkInsightsPath.
type LookupNetworkInsightsPathResult struct {
	// ARN of the selected Network Insights Path.
	Arn string `pulumi:"arn"`
	// AWS resource that is the destination of the path.
	Destination string `pulumi:"destination"`
	// IP address of the AWS resource that is the destination of the path.
	DestinationIp string `pulumi:"destinationIp"`
	// Destination port.
	DestinationPort int                            `pulumi:"destinationPort"`
	Filters         []GetNetworkInsightsPathFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id                    string `pulumi:"id"`
	NetworkInsightsPathId string `pulumi:"networkInsightsPathId"`
	// Protocol.
	Protocol string `pulumi:"protocol"`
	// AWS resource that is the source of the path.
	Source string `pulumi:"source"`
	// IP address of the AWS resource that is the source of the path.
	SourceIp string `pulumi:"sourceIp"`
	// Map of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
}

func LookupNetworkInsightsPathOutput(ctx *pulumi.Context, args LookupNetworkInsightsPathOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkInsightsPathResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkInsightsPathResult, error) {
			args := v.(LookupNetworkInsightsPathArgs)
			r, err := LookupNetworkInsightsPath(ctx, &args, opts...)
			var s LookupNetworkInsightsPathResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkInsightsPathResultOutput)
}

// A collection of arguments for invoking getNetworkInsightsPath.
type LookupNetworkInsightsPathOutputArgs struct {
	// Configuration block(s) for filtering. Detailed below.
	Filters GetNetworkInsightsPathFilterArrayInput `pulumi:"filters"`
	// ID of the Network Insights Path to select.
	NetworkInsightsPathId pulumi.StringPtrInput `pulumi:"networkInsightsPathId"`
	// Map of tags assigned to the resource.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupNetworkInsightsPathOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkInsightsPathArgs)(nil)).Elem()
}

// A collection of values returned by getNetworkInsightsPath.
type LookupNetworkInsightsPathResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkInsightsPathResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkInsightsPathResult)(nil)).Elem()
}

func (o LookupNetworkInsightsPathResultOutput) ToLookupNetworkInsightsPathResultOutput() LookupNetworkInsightsPathResultOutput {
	return o
}

func (o LookupNetworkInsightsPathResultOutput) ToLookupNetworkInsightsPathResultOutputWithContext(ctx context.Context) LookupNetworkInsightsPathResultOutput {
	return o
}

// ARN of the selected Network Insights Path.
func (o LookupNetworkInsightsPathResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInsightsPathResult) string { return v.Arn }).(pulumi.StringOutput)
}

// AWS resource that is the destination of the path.
func (o LookupNetworkInsightsPathResultOutput) Destination() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInsightsPathResult) string { return v.Destination }).(pulumi.StringOutput)
}

// IP address of the AWS resource that is the destination of the path.
func (o LookupNetworkInsightsPathResultOutput) DestinationIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInsightsPathResult) string { return v.DestinationIp }).(pulumi.StringOutput)
}

// Destination port.
func (o LookupNetworkInsightsPathResultOutput) DestinationPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNetworkInsightsPathResult) int { return v.DestinationPort }).(pulumi.IntOutput)
}

func (o LookupNetworkInsightsPathResultOutput) Filters() GetNetworkInsightsPathFilterArrayOutput {
	return o.ApplyT(func(v LookupNetworkInsightsPathResult) []GetNetworkInsightsPathFilter { return v.Filters }).(GetNetworkInsightsPathFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNetworkInsightsPathResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInsightsPathResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNetworkInsightsPathResultOutput) NetworkInsightsPathId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInsightsPathResult) string { return v.NetworkInsightsPathId }).(pulumi.StringOutput)
}

// Protocol.
func (o LookupNetworkInsightsPathResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInsightsPathResult) string { return v.Protocol }).(pulumi.StringOutput)
}

// AWS resource that is the source of the path.
func (o LookupNetworkInsightsPathResultOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInsightsPathResult) string { return v.Source }).(pulumi.StringOutput)
}

// IP address of the AWS resource that is the source of the path.
func (o LookupNetworkInsightsPathResultOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkInsightsPathResult) string { return v.SourceIp }).(pulumi.StringOutput)
}

// Map of tags assigned to the resource.
func (o LookupNetworkInsightsPathResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkInsightsPathResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkInsightsPathResultOutput{})
}
