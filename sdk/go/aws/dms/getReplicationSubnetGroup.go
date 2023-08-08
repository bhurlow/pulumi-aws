// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Data source for managing an AWS DMS (Database Migration) Replication Subnet Group.
//
// ## Example Usage
// ### Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dms.LookupReplicationSubnetGroup(ctx, &dms.LookupReplicationSubnetGroupArgs{
//				ReplicationSubnetGroupId: aws_dms_replication_subnet_group.Test.Replication_subnet_group_id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupReplicationSubnetGroup(ctx *pulumi.Context, args *LookupReplicationSubnetGroupArgs, opts ...pulumi.InvokeOption) (*LookupReplicationSubnetGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupReplicationSubnetGroupResult
	err := ctx.Invoke("aws:dms/getReplicationSubnetGroup:getReplicationSubnetGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getReplicationSubnetGroup.
type LookupReplicationSubnetGroupArgs struct {
	// Name for the replication subnet group. This value is stored as a lowercase string. It must contain no more than 255 alphanumeric characters, periods, spaces, underscores, or hyphens and cannot be `default`.
	ReplicationSubnetGroupId string            `pulumi:"replicationSubnetGroupId"`
	Tags                     map[string]string `pulumi:"tags"`
}

// A collection of values returned by getReplicationSubnetGroup.
type LookupReplicationSubnetGroupResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                        string `pulumi:"id"`
	ReplicationSubnetGroupArn string `pulumi:"replicationSubnetGroupArn"`
	// Description for the subnet group.
	ReplicationSubnetGroupDescription string `pulumi:"replicationSubnetGroupDescription"`
	ReplicationSubnetGroupId          string `pulumi:"replicationSubnetGroupId"`
	SubnetGroupStatus                 string `pulumi:"subnetGroupStatus"`
	// List of at least 2 EC2 subnet IDs for the subnet group. The subnets must cover at least 2 availability zones.
	SubnetIds []string          `pulumi:"subnetIds"`
	Tags      map[string]string `pulumi:"tags"`
	// The ID of the VPC the subnet group is in.
	VpcId string `pulumi:"vpcId"`
}

func LookupReplicationSubnetGroupOutput(ctx *pulumi.Context, args LookupReplicationSubnetGroupOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationSubnetGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReplicationSubnetGroupResult, error) {
			args := v.(LookupReplicationSubnetGroupArgs)
			r, err := LookupReplicationSubnetGroup(ctx, &args, opts...)
			var s LookupReplicationSubnetGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReplicationSubnetGroupResultOutput)
}

// A collection of arguments for invoking getReplicationSubnetGroup.
type LookupReplicationSubnetGroupOutputArgs struct {
	// Name for the replication subnet group. This value is stored as a lowercase string. It must contain no more than 255 alphanumeric characters, periods, spaces, underscores, or hyphens and cannot be `default`.
	ReplicationSubnetGroupId pulumi.StringInput    `pulumi:"replicationSubnetGroupId"`
	Tags                     pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupReplicationSubnetGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationSubnetGroupArgs)(nil)).Elem()
}

// A collection of values returned by getReplicationSubnetGroup.
type LookupReplicationSubnetGroupResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationSubnetGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationSubnetGroupResult)(nil)).Elem()
}

func (o LookupReplicationSubnetGroupResultOutput) ToLookupReplicationSubnetGroupResultOutput() LookupReplicationSubnetGroupResultOutput {
	return o
}

func (o LookupReplicationSubnetGroupResultOutput) ToLookupReplicationSubnetGroupResultOutputWithContext(ctx context.Context) LookupReplicationSubnetGroupResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupReplicationSubnetGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationSubnetGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupReplicationSubnetGroupResultOutput) ReplicationSubnetGroupArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationSubnetGroupResult) string { return v.ReplicationSubnetGroupArn }).(pulumi.StringOutput)
}

// Description for the subnet group.
func (o LookupReplicationSubnetGroupResultOutput) ReplicationSubnetGroupDescription() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationSubnetGroupResult) string { return v.ReplicationSubnetGroupDescription }).(pulumi.StringOutput)
}

func (o LookupReplicationSubnetGroupResultOutput) ReplicationSubnetGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationSubnetGroupResult) string { return v.ReplicationSubnetGroupId }).(pulumi.StringOutput)
}

func (o LookupReplicationSubnetGroupResultOutput) SubnetGroupStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationSubnetGroupResult) string { return v.SubnetGroupStatus }).(pulumi.StringOutput)
}

// List of at least 2 EC2 subnet IDs for the subnet group. The subnets must cover at least 2 availability zones.
func (o LookupReplicationSubnetGroupResultOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupReplicationSubnetGroupResult) []string { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

func (o LookupReplicationSubnetGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupReplicationSubnetGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The ID of the VPC the subnet group is in.
func (o LookupReplicationSubnetGroupResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReplicationSubnetGroupResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationSubnetGroupResultOutput{})
}
