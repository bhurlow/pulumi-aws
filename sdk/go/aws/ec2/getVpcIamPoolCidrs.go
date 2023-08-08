// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ec2.getVpcIamPoolCidrs` provides details about an IPAM pool.
//
// This resource can prove useful when an ipam pool was shared to your account and you want to know all (or a filtered list) of the CIDRs that are provisioned into the pool.
func GetVpcIamPoolCidrs(ctx *pulumi.Context, args *GetVpcIamPoolCidrsArgs, opts ...pulumi.InvokeOption) (*GetVpcIamPoolCidrsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVpcIamPoolCidrsResult
	err := ctx.Invoke("aws:ec2/getVpcIamPoolCidrs:getVpcIamPoolCidrs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcIamPoolCidrs.
type GetVpcIamPoolCidrsArgs struct {
	// Custom filter block as described below.
	Filters []GetVpcIamPoolCidrsFilter `pulumi:"filters"`
	// ID of the IPAM pool you would like the list of provisioned CIDRs.
	IpamPoolId string `pulumi:"ipamPoolId"`
}

// A collection of values returned by getVpcIamPoolCidrs.
type GetVpcIamPoolCidrsResult struct {
	Filters []GetVpcIamPoolCidrsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The CIDRs provisioned into the IPAM pool, described below.
	IpamPoolCidrs []GetVpcIamPoolCidrsIpamPoolCidr `pulumi:"ipamPoolCidrs"`
	IpamPoolId    string                           `pulumi:"ipamPoolId"`
}

func GetVpcIamPoolCidrsOutput(ctx *pulumi.Context, args GetVpcIamPoolCidrsOutputArgs, opts ...pulumi.InvokeOption) GetVpcIamPoolCidrsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVpcIamPoolCidrsResult, error) {
			args := v.(GetVpcIamPoolCidrsArgs)
			r, err := GetVpcIamPoolCidrs(ctx, &args, opts...)
			var s GetVpcIamPoolCidrsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVpcIamPoolCidrsResultOutput)
}

// A collection of arguments for invoking getVpcIamPoolCidrs.
type GetVpcIamPoolCidrsOutputArgs struct {
	// Custom filter block as described below.
	Filters GetVpcIamPoolCidrsFilterArrayInput `pulumi:"filters"`
	// ID of the IPAM pool you would like the list of provisioned CIDRs.
	IpamPoolId pulumi.StringInput `pulumi:"ipamPoolId"`
}

func (GetVpcIamPoolCidrsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcIamPoolCidrsArgs)(nil)).Elem()
}

// A collection of values returned by getVpcIamPoolCidrs.
type GetVpcIamPoolCidrsResultOutput struct{ *pulumi.OutputState }

func (GetVpcIamPoolCidrsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcIamPoolCidrsResult)(nil)).Elem()
}

func (o GetVpcIamPoolCidrsResultOutput) ToGetVpcIamPoolCidrsResultOutput() GetVpcIamPoolCidrsResultOutput {
	return o
}

func (o GetVpcIamPoolCidrsResultOutput) ToGetVpcIamPoolCidrsResultOutputWithContext(ctx context.Context) GetVpcIamPoolCidrsResultOutput {
	return o
}

func (o GetVpcIamPoolCidrsResultOutput) Filters() GetVpcIamPoolCidrsFilterArrayOutput {
	return o.ApplyT(func(v GetVpcIamPoolCidrsResult) []GetVpcIamPoolCidrsFilter { return v.Filters }).(GetVpcIamPoolCidrsFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetVpcIamPoolCidrsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcIamPoolCidrsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The CIDRs provisioned into the IPAM pool, described below.
func (o GetVpcIamPoolCidrsResultOutput) IpamPoolCidrs() GetVpcIamPoolCidrsIpamPoolCidrArrayOutput {
	return o.ApplyT(func(v GetVpcIamPoolCidrsResult) []GetVpcIamPoolCidrsIpamPoolCidr { return v.IpamPoolCidrs }).(GetVpcIamPoolCidrsIpamPoolCidrArrayOutput)
}

func (o GetVpcIamPoolCidrsResultOutput) IpamPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcIamPoolCidrsResult) string { return v.IpamPoolId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVpcIamPoolCidrsResultOutput{})
}
