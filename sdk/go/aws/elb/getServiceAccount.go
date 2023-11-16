// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the Account ID of the [AWS Elastic Load Balancing Service Account](http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-access-logs.html#attach-bucket-policy)
// in a given region for the purpose of permitting in S3 bucket policy.
//
// > **Note:** For AWS Regions opened since Jakarta (`ap-southeast-3`) in December 2021, AWS [documents that](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-access-logs.html#attach-bucket-policy) a [service principal name](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html#principal-services) should be used instead of an AWS account ID in any relevant IAM policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elb"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// main, err := elb.GetServiceAccount(ctx, nil, nil);
// if err != nil {
// return err
// }
// elbLogs, err := s3.NewBucketV2(ctx, "elbLogs", nil)
// if err != nil {
// return err
// }
// _, err = s3.NewBucketAclV2(ctx, "elbLogsAcl", &s3.BucketAclV2Args{
// Bucket: elbLogs.ID(),
// Acl: pulumi.String("private"),
// })
// if err != nil {
// return err
// }
// allowElbLoggingPolicyDocument := elbLogs.Arn.ApplyT(func(arn string) (iam.GetPolicyDocumentResult, error) {
// return iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
// Statements: []iam.GetPolicyDocumentStatement{
// {
// Effect: "Allow",
// Principals: []iam.GetPolicyDocumentStatementPrincipal{
// {
// Type: "AWS",
// Identifiers: interface{}{
// main.Arn,
// },
// },
// },
// Actions: []string{
// "s3:PutObject",
// },
// Resources: []string{
// fmt.Sprintf("%v/AWSLogs/*", arn),
// },
// },
// },
// }, nil), nil
// }).(iam.GetPolicyDocumentResultOutput)
// _, err = s3.NewBucketPolicy(ctx, "allowElbLoggingBucketPolicy", &s3.BucketPolicyArgs{
// Bucket: elbLogs.ID(),
// Policy: allowElbLoggingPolicyDocument.ApplyT(func(allowElbLoggingPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
// return &allowElbLoggingPolicyDocument.Json, nil
// }).(pulumi.StringPtrOutput),
// })
// if err != nil {
// return err
// }
// _, err = elb.NewLoadBalancer(ctx, "bar", &elb.LoadBalancerArgs{
// AvailabilityZones: pulumi.StringArray{
// pulumi.String("us-west-2a"),
// },
// AccessLogs: &elb.LoadBalancerAccessLogsArgs{
// Bucket: elbLogs.ID(),
// Interval: pulumi.Int(5),
// },
// Listeners: elb.LoadBalancerListenerArray{
// &elb.LoadBalancerListenerArgs{
// InstancePort: pulumi.Int(8000),
// InstanceProtocol: pulumi.String("http"),
// LbPort: pulumi.Int(80),
// LbProtocol: pulumi.String("http"),
// },
// },
// })
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
func GetServiceAccount(ctx *pulumi.Context, args *GetServiceAccountArgs, opts ...pulumi.InvokeOption) (*GetServiceAccountResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServiceAccountResult
	err := ctx.Invoke("aws:elb/getServiceAccount:getServiceAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceAccount.
type GetServiceAccountArgs struct {
	// Name of the region whose AWS ELB account ID is desired.
	// Defaults to the region from the AWS provider configuration.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getServiceAccount.
type GetServiceAccountResult struct {
	// ARN of the AWS ELB service account in the selected region.
	Arn string `pulumi:"arn"`
	// The provider-assigned unique ID for this managed resource.
	Id     string  `pulumi:"id"`
	Region *string `pulumi:"region"`
}

func GetServiceAccountOutput(ctx *pulumi.Context, args GetServiceAccountOutputArgs, opts ...pulumi.InvokeOption) GetServiceAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServiceAccountResult, error) {
			args := v.(GetServiceAccountArgs)
			r, err := GetServiceAccount(ctx, &args, opts...)
			var s GetServiceAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetServiceAccountResultOutput)
}

// A collection of arguments for invoking getServiceAccount.
type GetServiceAccountOutputArgs struct {
	// Name of the region whose AWS ELB account ID is desired.
	// Defaults to the region from the AWS provider configuration.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetServiceAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceAccountArgs)(nil)).Elem()
}

// A collection of values returned by getServiceAccount.
type GetServiceAccountResultOutput struct{ *pulumi.OutputState }

func (GetServiceAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceAccountResult)(nil)).Elem()
}

func (o GetServiceAccountResultOutput) ToGetServiceAccountResultOutput() GetServiceAccountResultOutput {
	return o
}

func (o GetServiceAccountResultOutput) ToGetServiceAccountResultOutputWithContext(ctx context.Context) GetServiceAccountResultOutput {
	return o
}

// ARN of the AWS ELB service account in the selected region.
func (o GetServiceAccountResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceAccountResult) string { return v.Arn }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServiceAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetServiceAccountResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServiceAccountResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServiceAccountResultOutput{})
}
