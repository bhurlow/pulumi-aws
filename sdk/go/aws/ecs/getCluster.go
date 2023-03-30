// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The ECS Cluster data source allows access to details of a specific
// cluster within an AWS ECS service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ecs.LookupCluster(ctx, &ecs.LookupClusterArgs{
//				ClusterName: "ecs-mongo-production",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("aws:ecs/getCluster:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCluster.
type LookupClusterArgs struct {
	// Name of the ECS Cluster
	ClusterName string `pulumi:"clusterName"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getCluster.
type LookupClusterResult struct {
	// ARN of the ECS Cluster
	Arn         string `pulumi:"arn"`
	ClusterName string `pulumi:"clusterName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Number of pending tasks for the ECS Cluster
	PendingTasksCount int `pulumi:"pendingTasksCount"`
	// The number of registered container instances for the ECS Cluster
	RegisteredContainerInstancesCount int `pulumi:"registeredContainerInstancesCount"`
	// Number of running tasks for the ECS Cluster
	RunningTasksCount int `pulumi:"runningTasksCount"`
	// The default Service Connect namespace
	ServiceConnectDefaults []GetClusterServiceConnectDefault `pulumi:"serviceConnectDefaults"`
	// Settings associated with the ECS Cluster
	Settings []GetClusterSetting `pulumi:"settings"`
	// Status of the ECS Cluster
	Status string `pulumi:"status"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

func LookupClusterOutput(ctx *pulumi.Context, args LookupClusterOutputArgs, opts ...pulumi.InvokeOption) LookupClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClusterResult, error) {
			args := v.(LookupClusterArgs)
			r, err := LookupCluster(ctx, &args, opts...)
			var s LookupClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupClusterResultOutput)
}

// A collection of arguments for invoking getCluster.
type LookupClusterOutputArgs struct {
	// Name of the ECS Cluster
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// Key-value map of resource tags
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterArgs)(nil)).Elem()
}

// A collection of values returned by getCluster.
type LookupClusterResultOutput struct{ *pulumi.OutputState }

func (LookupClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterResult)(nil)).Elem()
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutput() LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutputWithContext(ctx context.Context) LookupClusterResultOutput {
	return o
}

// ARN of the ECS Cluster
func (o LookupClusterResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Arn }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ClusterName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// Number of pending tasks for the ECS Cluster
func (o LookupClusterResultOutput) PendingTasksCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupClusterResult) int { return v.PendingTasksCount }).(pulumi.IntOutput)
}

// The number of registered container instances for the ECS Cluster
func (o LookupClusterResultOutput) RegisteredContainerInstancesCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupClusterResult) int { return v.RegisteredContainerInstancesCount }).(pulumi.IntOutput)
}

// Number of running tasks for the ECS Cluster
func (o LookupClusterResultOutput) RunningTasksCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupClusterResult) int { return v.RunningTasksCount }).(pulumi.IntOutput)
}

// The default Service Connect namespace
func (o LookupClusterResultOutput) ServiceConnectDefaults() GetClusterServiceConnectDefaultArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []GetClusterServiceConnectDefault { return v.ServiceConnectDefaults }).(GetClusterServiceConnectDefaultArrayOutput)
}

// Settings associated with the ECS Cluster
func (o LookupClusterResultOutput) Settings() GetClusterSettingArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []GetClusterSetting { return v.Settings }).(GetClusterSettingArrayOutput)
}

// Status of the ECS Cluster
func (o LookupClusterResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Status }).(pulumi.StringOutput)
}

// Key-value map of resource tags
func (o LookupClusterResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupClusterResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}
