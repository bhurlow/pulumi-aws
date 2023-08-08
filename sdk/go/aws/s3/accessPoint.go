// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage an S3 Access Point.
//
// > **NOTE on Access Points and Access Point Policies:** This provider provides both a standalone Access Point Policy resource and an Access Point resource with a resource policy defined in-line. You cannot use an Access Point with in-line resource policy in conjunction with an Access Point Policy resource. Doing so will cause a conflict of policies and will overwrite the access point's resource policy.
//
// > Advanced usage: To use a custom API endpoint for this resource, use the `s3control` endpoint provider configuration), not the `s3` endpoint provider configuration.
//
// ## Example Usage
// ### AWS Partition Bucket
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", nil)
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewAccessPoint(ctx, "exampleAccessPoint", &s3.AccessPointArgs{
//				Bucket: exampleBucketV2.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### S3 on Outposts Bucket
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3control"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleBucket, err := s3control.NewBucket(ctx, "exampleBucket", &s3control.BucketArgs{
//				Bucket: pulumi.String("example"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
//				CidrBlock: pulumi.String("10.0.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = s3.NewAccessPoint(ctx, "exampleAccessPoint", &s3.AccessPointArgs{
//				Bucket: exampleBucket.Arn,
//				VpcConfiguration: &s3.AccessPointVpcConfigurationArgs{
//					VpcId: exampleVpc.ID(),
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
// Import using the `account_id` and `name` separated by a colon (`:`) for Access Points associated with an AWS Partition S3 Bucketterraform import {
//
//	to = aws_s3_access_point.example
//
//	id = "123456789012:example" } Import using the ARN for Access Points associated with an S3 on Outposts Bucketterraform import {
//
//	to = aws_s3_access_point.example
//
//	id = "arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-1234567890123456/accesspoint/example" } **Using `pulumi import` to import.** For exampleImport using the `account_id` and `name` separated by a colon (`:`) for Access Points associated with an AWS Partition S3 Bucketconsole % pulumi import aws_s3_access_point.example 123456789012:example Import using the ARN for Access Points associated with an S3 on Outposts Bucketconsole % pulumi import aws_s3_access_point.example arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-1234567890123456/accesspoint/example
type AccessPoint struct {
	pulumi.CustomResourceState

	// AWS account ID for the owner of the bucket for which you want to create an access point. Defaults to automatically determined account ID of the AWS provider.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// Alias of the S3 Access Point.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// ARN of the S3 Access Point.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Name of an AWS Partition S3 Bucket or the ARN of S3 on Outposts Bucket that you want to associate this access point with.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// AWS account ID associated with the S3 bucket associated with this access point.
	BucketAccountId pulumi.StringOutput `pulumi:"bucketAccountId"`
	// DNS domain name of the S3 Access Point in the format _`name`_-_`accountId`_.s3-accesspoint._region_.amazonaws.com.
	// Note: S3 access points only support secure access by HTTPS. HTTP isn't supported.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// VPC endpoints for the S3 Access Point.
	Endpoints pulumi.StringMapOutput `pulumi:"endpoints"`
	// Indicates whether this access point currently has a policy that allows public access.
	HasPublicAccessPolicy pulumi.BoolOutput `pulumi:"hasPublicAccessPolicy"`
	// Name you want to assign to this access point.
	//
	// The following arguments are optional:
	Name pulumi.StringOutput `pulumi:"name"`
	// Indicates whether this access point allows access from the public Internet. Values are `VPC` (the access point doesn't allow access from the public Internet) and `Internet` (the access point allows access from the public Internet, subject to the access point and bucket access policies).
	NetworkOrigin pulumi.StringOutput `pulumi:"networkOrigin"`
	// Valid JSON document that specifies the policy that you want to apply to this access point. Removing `policy` from your configuration or setting `policy` to null or an empty string (i.e., `policy = ""`) _will not_ delete the policy since it could have been set by `s3control.AccessPointPolicy`. To remove the `policy`, set it to `"{}"` (an empty JSON document).
	Policy pulumi.StringOutput `pulumi:"policy"`
	// Configuration block to manage the `PublicAccessBlock` configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. Detailed below.
	PublicAccessBlockConfiguration AccessPointPublicAccessBlockConfigurationPtrOutput `pulumi:"publicAccessBlockConfiguration"`
	// Configuration block to restrict access to this access point to requests from the specified Virtual Private Cloud (VPC). Required for S3 on Outposts. Detailed below.
	VpcConfiguration AccessPointVpcConfigurationPtrOutput `pulumi:"vpcConfiguration"`
}

// NewAccessPoint registers a new resource with the given unique name, arguments, and options.
func NewAccessPoint(ctx *pulumi.Context,
	name string, args *AccessPointArgs, opts ...pulumi.ResourceOption) (*AccessPoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AccessPoint
	err := ctx.RegisterResource("aws:s3/accessPoint:AccessPoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessPoint gets an existing AccessPoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessPoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessPointState, opts ...pulumi.ResourceOption) (*AccessPoint, error) {
	var resource AccessPoint
	err := ctx.ReadResource("aws:s3/accessPoint:AccessPoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessPoint resources.
type accessPointState struct {
	// AWS account ID for the owner of the bucket for which you want to create an access point. Defaults to automatically determined account ID of the AWS provider.
	AccountId *string `pulumi:"accountId"`
	// Alias of the S3 Access Point.
	Alias *string `pulumi:"alias"`
	// ARN of the S3 Access Point.
	Arn *string `pulumi:"arn"`
	// Name of an AWS Partition S3 Bucket or the ARN of S3 on Outposts Bucket that you want to associate this access point with.
	Bucket *string `pulumi:"bucket"`
	// AWS account ID associated with the S3 bucket associated with this access point.
	BucketAccountId *string `pulumi:"bucketAccountId"`
	// DNS domain name of the S3 Access Point in the format _`name`_-_`accountId`_.s3-accesspoint._region_.amazonaws.com.
	// Note: S3 access points only support secure access by HTTPS. HTTP isn't supported.
	DomainName *string `pulumi:"domainName"`
	// VPC endpoints for the S3 Access Point.
	Endpoints map[string]string `pulumi:"endpoints"`
	// Indicates whether this access point currently has a policy that allows public access.
	HasPublicAccessPolicy *bool `pulumi:"hasPublicAccessPolicy"`
	// Name you want to assign to this access point.
	//
	// The following arguments are optional:
	Name *string `pulumi:"name"`
	// Indicates whether this access point allows access from the public Internet. Values are `VPC` (the access point doesn't allow access from the public Internet) and `Internet` (the access point allows access from the public Internet, subject to the access point and bucket access policies).
	NetworkOrigin *string `pulumi:"networkOrigin"`
	// Valid JSON document that specifies the policy that you want to apply to this access point. Removing `policy` from your configuration or setting `policy` to null or an empty string (i.e., `policy = ""`) _will not_ delete the policy since it could have been set by `s3control.AccessPointPolicy`. To remove the `policy`, set it to `"{}"` (an empty JSON document).
	Policy *string `pulumi:"policy"`
	// Configuration block to manage the `PublicAccessBlock` configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. Detailed below.
	PublicAccessBlockConfiguration *AccessPointPublicAccessBlockConfiguration `pulumi:"publicAccessBlockConfiguration"`
	// Configuration block to restrict access to this access point to requests from the specified Virtual Private Cloud (VPC). Required for S3 on Outposts. Detailed below.
	VpcConfiguration *AccessPointVpcConfiguration `pulumi:"vpcConfiguration"`
}

type AccessPointState struct {
	// AWS account ID for the owner of the bucket for which you want to create an access point. Defaults to automatically determined account ID of the AWS provider.
	AccountId pulumi.StringPtrInput
	// Alias of the S3 Access Point.
	Alias pulumi.StringPtrInput
	// ARN of the S3 Access Point.
	Arn pulumi.StringPtrInput
	// Name of an AWS Partition S3 Bucket or the ARN of S3 on Outposts Bucket that you want to associate this access point with.
	Bucket pulumi.StringPtrInput
	// AWS account ID associated with the S3 bucket associated with this access point.
	BucketAccountId pulumi.StringPtrInput
	// DNS domain name of the S3 Access Point in the format _`name`_-_`accountId`_.s3-accesspoint._region_.amazonaws.com.
	// Note: S3 access points only support secure access by HTTPS. HTTP isn't supported.
	DomainName pulumi.StringPtrInput
	// VPC endpoints for the S3 Access Point.
	Endpoints pulumi.StringMapInput
	// Indicates whether this access point currently has a policy that allows public access.
	HasPublicAccessPolicy pulumi.BoolPtrInput
	// Name you want to assign to this access point.
	//
	// The following arguments are optional:
	Name pulumi.StringPtrInput
	// Indicates whether this access point allows access from the public Internet. Values are `VPC` (the access point doesn't allow access from the public Internet) and `Internet` (the access point allows access from the public Internet, subject to the access point and bucket access policies).
	NetworkOrigin pulumi.StringPtrInput
	// Valid JSON document that specifies the policy that you want to apply to this access point. Removing `policy` from your configuration or setting `policy` to null or an empty string (i.e., `policy = ""`) _will not_ delete the policy since it could have been set by `s3control.AccessPointPolicy`. To remove the `policy`, set it to `"{}"` (an empty JSON document).
	Policy pulumi.StringPtrInput
	// Configuration block to manage the `PublicAccessBlock` configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. Detailed below.
	PublicAccessBlockConfiguration AccessPointPublicAccessBlockConfigurationPtrInput
	// Configuration block to restrict access to this access point to requests from the specified Virtual Private Cloud (VPC). Required for S3 on Outposts. Detailed below.
	VpcConfiguration AccessPointVpcConfigurationPtrInput
}

func (AccessPointState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPointState)(nil)).Elem()
}

type accessPointArgs struct {
	// AWS account ID for the owner of the bucket for which you want to create an access point. Defaults to automatically determined account ID of the AWS provider.
	AccountId *string `pulumi:"accountId"`
	// Name of an AWS Partition S3 Bucket or the ARN of S3 on Outposts Bucket that you want to associate this access point with.
	Bucket string `pulumi:"bucket"`
	// AWS account ID associated with the S3 bucket associated with this access point.
	BucketAccountId *string `pulumi:"bucketAccountId"`
	// Name you want to assign to this access point.
	//
	// The following arguments are optional:
	Name *string `pulumi:"name"`
	// Valid JSON document that specifies the policy that you want to apply to this access point. Removing `policy` from your configuration or setting `policy` to null or an empty string (i.e., `policy = ""`) _will not_ delete the policy since it could have been set by `s3control.AccessPointPolicy`. To remove the `policy`, set it to `"{}"` (an empty JSON document).
	Policy *string `pulumi:"policy"`
	// Configuration block to manage the `PublicAccessBlock` configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. Detailed below.
	PublicAccessBlockConfiguration *AccessPointPublicAccessBlockConfiguration `pulumi:"publicAccessBlockConfiguration"`
	// Configuration block to restrict access to this access point to requests from the specified Virtual Private Cloud (VPC). Required for S3 on Outposts. Detailed below.
	VpcConfiguration *AccessPointVpcConfiguration `pulumi:"vpcConfiguration"`
}

// The set of arguments for constructing a AccessPoint resource.
type AccessPointArgs struct {
	// AWS account ID for the owner of the bucket for which you want to create an access point. Defaults to automatically determined account ID of the AWS provider.
	AccountId pulumi.StringPtrInput
	// Name of an AWS Partition S3 Bucket or the ARN of S3 on Outposts Bucket that you want to associate this access point with.
	Bucket pulumi.StringInput
	// AWS account ID associated with the S3 bucket associated with this access point.
	BucketAccountId pulumi.StringPtrInput
	// Name you want to assign to this access point.
	//
	// The following arguments are optional:
	Name pulumi.StringPtrInput
	// Valid JSON document that specifies the policy that you want to apply to this access point. Removing `policy` from your configuration or setting `policy` to null or an empty string (i.e., `policy = ""`) _will not_ delete the policy since it could have been set by `s3control.AccessPointPolicy`. To remove the `policy`, set it to `"{}"` (an empty JSON document).
	Policy pulumi.StringPtrInput
	// Configuration block to manage the `PublicAccessBlock` configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. Detailed below.
	PublicAccessBlockConfiguration AccessPointPublicAccessBlockConfigurationPtrInput
	// Configuration block to restrict access to this access point to requests from the specified Virtual Private Cloud (VPC). Required for S3 on Outposts. Detailed below.
	VpcConfiguration AccessPointVpcConfigurationPtrInput
}

func (AccessPointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessPointArgs)(nil)).Elem()
}

type AccessPointInput interface {
	pulumi.Input

	ToAccessPointOutput() AccessPointOutput
	ToAccessPointOutputWithContext(ctx context.Context) AccessPointOutput
}

func (*AccessPoint) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPoint)(nil)).Elem()
}

func (i *AccessPoint) ToAccessPointOutput() AccessPointOutput {
	return i.ToAccessPointOutputWithContext(context.Background())
}

func (i *AccessPoint) ToAccessPointOutputWithContext(ctx context.Context) AccessPointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPointOutput)
}

// AccessPointArrayInput is an input type that accepts AccessPointArray and AccessPointArrayOutput values.
// You can construct a concrete instance of `AccessPointArrayInput` via:
//
//	AccessPointArray{ AccessPointArgs{...} }
type AccessPointArrayInput interface {
	pulumi.Input

	ToAccessPointArrayOutput() AccessPointArrayOutput
	ToAccessPointArrayOutputWithContext(context.Context) AccessPointArrayOutput
}

type AccessPointArray []AccessPointInput

func (AccessPointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessPoint)(nil)).Elem()
}

func (i AccessPointArray) ToAccessPointArrayOutput() AccessPointArrayOutput {
	return i.ToAccessPointArrayOutputWithContext(context.Background())
}

func (i AccessPointArray) ToAccessPointArrayOutputWithContext(ctx context.Context) AccessPointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPointArrayOutput)
}

// AccessPointMapInput is an input type that accepts AccessPointMap and AccessPointMapOutput values.
// You can construct a concrete instance of `AccessPointMapInput` via:
//
//	AccessPointMap{ "key": AccessPointArgs{...} }
type AccessPointMapInput interface {
	pulumi.Input

	ToAccessPointMapOutput() AccessPointMapOutput
	ToAccessPointMapOutputWithContext(context.Context) AccessPointMapOutput
}

type AccessPointMap map[string]AccessPointInput

func (AccessPointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessPoint)(nil)).Elem()
}

func (i AccessPointMap) ToAccessPointMapOutput() AccessPointMapOutput {
	return i.ToAccessPointMapOutputWithContext(context.Background())
}

func (i AccessPointMap) ToAccessPointMapOutputWithContext(ctx context.Context) AccessPointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPointMapOutput)
}

type AccessPointOutput struct{ *pulumi.OutputState }

func (AccessPointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPoint)(nil)).Elem()
}

func (o AccessPointOutput) ToAccessPointOutput() AccessPointOutput {
	return o
}

func (o AccessPointOutput) ToAccessPointOutputWithContext(ctx context.Context) AccessPointOutput {
	return o
}

// AWS account ID for the owner of the bucket for which you want to create an access point. Defaults to automatically determined account ID of the AWS provider.
func (o AccessPointOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPoint) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// Alias of the S3 Access Point.
func (o AccessPointOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPoint) pulumi.StringOutput { return v.Alias }).(pulumi.StringOutput)
}

// ARN of the S3 Access Point.
func (o AccessPointOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPoint) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Name of an AWS Partition S3 Bucket or the ARN of S3 on Outposts Bucket that you want to associate this access point with.
func (o AccessPointOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPoint) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// AWS account ID associated with the S3 bucket associated with this access point.
func (o AccessPointOutput) BucketAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPoint) pulumi.StringOutput { return v.BucketAccountId }).(pulumi.StringOutput)
}

// DNS domain name of the S3 Access Point in the format _`name`_-_`accountId`_.s3-accesspoint._region_.amazonaws.com.
// Note: S3 access points only support secure access by HTTPS. HTTP isn't supported.
func (o AccessPointOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPoint) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// VPC endpoints for the S3 Access Point.
func (o AccessPointOutput) Endpoints() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AccessPoint) pulumi.StringMapOutput { return v.Endpoints }).(pulumi.StringMapOutput)
}

// Indicates whether this access point currently has a policy that allows public access.
func (o AccessPointOutput) HasPublicAccessPolicy() pulumi.BoolOutput {
	return o.ApplyT(func(v *AccessPoint) pulumi.BoolOutput { return v.HasPublicAccessPolicy }).(pulumi.BoolOutput)
}

// Name you want to assign to this access point.
//
// The following arguments are optional:
func (o AccessPointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Indicates whether this access point allows access from the public Internet. Values are `VPC` (the access point doesn't allow access from the public Internet) and `Internet` (the access point allows access from the public Internet, subject to the access point and bucket access policies).
func (o AccessPointOutput) NetworkOrigin() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPoint) pulumi.StringOutput { return v.NetworkOrigin }).(pulumi.StringOutput)
}

// Valid JSON document that specifies the policy that you want to apply to this access point. Removing `policy` from your configuration or setting `policy` to null or an empty string (i.e., `policy = ""`) _will not_ delete the policy since it could have been set by `s3control.AccessPointPolicy`. To remove the `policy`, set it to `"{}"` (an empty JSON document).
func (o AccessPointOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessPoint) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

// Configuration block to manage the `PublicAccessBlock` configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. Detailed below.
func (o AccessPointOutput) PublicAccessBlockConfiguration() AccessPointPublicAccessBlockConfigurationPtrOutput {
	return o.ApplyT(func(v *AccessPoint) AccessPointPublicAccessBlockConfigurationPtrOutput {
		return v.PublicAccessBlockConfiguration
	}).(AccessPointPublicAccessBlockConfigurationPtrOutput)
}

// Configuration block to restrict access to this access point to requests from the specified Virtual Private Cloud (VPC). Required for S3 on Outposts. Detailed below.
func (o AccessPointOutput) VpcConfiguration() AccessPointVpcConfigurationPtrOutput {
	return o.ApplyT(func(v *AccessPoint) AccessPointVpcConfigurationPtrOutput { return v.VpcConfiguration }).(AccessPointVpcConfigurationPtrOutput)
}

type AccessPointArrayOutput struct{ *pulumi.OutputState }

func (AccessPointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessPoint)(nil)).Elem()
}

func (o AccessPointArrayOutput) ToAccessPointArrayOutput() AccessPointArrayOutput {
	return o
}

func (o AccessPointArrayOutput) ToAccessPointArrayOutputWithContext(ctx context.Context) AccessPointArrayOutput {
	return o
}

func (o AccessPointArrayOutput) Index(i pulumi.IntInput) AccessPointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccessPoint {
		return vs[0].([]*AccessPoint)[vs[1].(int)]
	}).(AccessPointOutput)
}

type AccessPointMapOutput struct{ *pulumi.OutputState }

func (AccessPointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessPoint)(nil)).Elem()
}

func (o AccessPointMapOutput) ToAccessPointMapOutput() AccessPointMapOutput {
	return o
}

func (o AccessPointMapOutput) ToAccessPointMapOutputWithContext(ctx context.Context) AccessPointMapOutput {
	return o
}

func (o AccessPointMapOutput) MapIndex(k pulumi.StringInput) AccessPointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccessPoint {
		return vs[0].(map[string]*AccessPoint)[vs[1].(string)]
	}).(AccessPointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPointInput)(nil)).Elem(), &AccessPoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPointArrayInput)(nil)).Elem(), AccessPointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessPointMapInput)(nil)).Elem(), AccessPointMap{})
	pulumi.RegisterOutputType(AccessPointOutput{})
	pulumi.RegisterOutputType(AccessPointArrayOutput{})
	pulumi.RegisterOutputType(AccessPointMapOutput{})
}
