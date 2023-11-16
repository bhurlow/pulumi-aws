// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package codebuild

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CodeBuild Resource Policy Resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//	"fmt"
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codebuild"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleReportGroup, err := codebuild.NewReportGroup(ctx, "exampleReportGroup", &codebuild.ReportGroupArgs{
//				Type: pulumi.String("TEST"),
//				ExportConfig: &codebuild.ReportGroupExportConfigArgs{
//					Type: pulumi.String("NO_EXPORT"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			currentPartition, err := aws.GetPartition(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = codebuild.NewResourcePolicy(ctx, "exampleResourcePolicy", &codebuild.ResourcePolicyArgs{
//				ResourceArn: exampleReportGroup.Arn,
//				Policy: exampleReportGroup.Arn.ApplyT(func(arn string) (pulumi.String, error) {
//					var _zero pulumi.String
//					tmpJSON0, err := json.Marshal(map[string]interface{}{
//						"Version": "2012-10-17",
//						"Id":      "default",
//						"Statement": []map[string]interface{}{
//							map[string]interface{}{
//								"Sid":    "default",
//								"Effect": "Allow",
//								"Principal": map[string]interface{}{
//									"AWS": fmt.Sprintf("arn:%v:iam::%v:root", currentPartition.Partition, currentCallerIdentity.AccountId),
//								},
//								"Action": []string{
//									"codebuild:BatchGetReportGroups",
//									"codebuild:BatchGetReports",
//									"codebuild:ListReportsForReportGroup",
//									"codebuild:DescribeTestCases",
//								},
//								"Resource": arn,
//							},
//						},
//					})
//					if err != nil {
//						return _zero, err
//					}
//					json0 := string(tmpJSON0)
//					return pulumi.String(json0), nil
//				}).(pulumi.StringOutput),
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
// Using `pulumi import`, import CodeBuild Resource Policy using the CodeBuild Resource Policy arn. For example:
//
// ```sh
//
//	$ pulumi import aws:codebuild/resourcePolicy:ResourcePolicy example arn:aws:codebuild:us-west-2:123456789:report-group/report-group-name
//
// ```
type ResourcePolicy struct {
	pulumi.CustomResourceState

	// A JSON-formatted resource policy. For more information, see [Sharing a Projec](https://docs.aws.amazon.com/codebuild/latest/userguide/project-sharing.html#project-sharing-share) and [Sharing a Report Group](https://docs.aws.amazon.com/codebuild/latest/userguide/report-groups-sharing.html#report-groups-sharing-share).
	Policy pulumi.StringOutput `pulumi:"policy"`
	// The ARN of the Project or ReportGroup resource you want to associate with a resource policy.
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`
}

// NewResourcePolicy registers a new resource with the given unique name, arguments, and options.
func NewResourcePolicy(ctx *pulumi.Context,
	name string, args *ResourcePolicyArgs, opts ...pulumi.ResourceOption) (*ResourcePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	if args.ResourceArn == nil {
		return nil, errors.New("invalid value for required argument 'ResourceArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResourcePolicy
	err := ctx.RegisterResource("aws:codebuild/resourcePolicy:ResourcePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourcePolicy gets an existing ResourcePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourcePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourcePolicyState, opts ...pulumi.ResourceOption) (*ResourcePolicy, error) {
	var resource ResourcePolicy
	err := ctx.ReadResource("aws:codebuild/resourcePolicy:ResourcePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourcePolicy resources.
type resourcePolicyState struct {
	// A JSON-formatted resource policy. For more information, see [Sharing a Projec](https://docs.aws.amazon.com/codebuild/latest/userguide/project-sharing.html#project-sharing-share) and [Sharing a Report Group](https://docs.aws.amazon.com/codebuild/latest/userguide/report-groups-sharing.html#report-groups-sharing-share).
	Policy *string `pulumi:"policy"`
	// The ARN of the Project or ReportGroup resource you want to associate with a resource policy.
	ResourceArn *string `pulumi:"resourceArn"`
}

type ResourcePolicyState struct {
	// A JSON-formatted resource policy. For more information, see [Sharing a Projec](https://docs.aws.amazon.com/codebuild/latest/userguide/project-sharing.html#project-sharing-share) and [Sharing a Report Group](https://docs.aws.amazon.com/codebuild/latest/userguide/report-groups-sharing.html#report-groups-sharing-share).
	Policy pulumi.StringPtrInput
	// The ARN of the Project or ReportGroup resource you want to associate with a resource policy.
	ResourceArn pulumi.StringPtrInput
}

func (ResourcePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePolicyState)(nil)).Elem()
}

type resourcePolicyArgs struct {
	// A JSON-formatted resource policy. For more information, see [Sharing a Projec](https://docs.aws.amazon.com/codebuild/latest/userguide/project-sharing.html#project-sharing-share) and [Sharing a Report Group](https://docs.aws.amazon.com/codebuild/latest/userguide/report-groups-sharing.html#report-groups-sharing-share).
	Policy string `pulumi:"policy"`
	// The ARN of the Project or ReportGroup resource you want to associate with a resource policy.
	ResourceArn string `pulumi:"resourceArn"`
}

// The set of arguments for constructing a ResourcePolicy resource.
type ResourcePolicyArgs struct {
	// A JSON-formatted resource policy. For more information, see [Sharing a Projec](https://docs.aws.amazon.com/codebuild/latest/userguide/project-sharing.html#project-sharing-share) and [Sharing a Report Group](https://docs.aws.amazon.com/codebuild/latest/userguide/report-groups-sharing.html#report-groups-sharing-share).
	Policy pulumi.StringInput
	// The ARN of the Project or ReportGroup resource you want to associate with a resource policy.
	ResourceArn pulumi.StringInput
}

func (ResourcePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourcePolicyArgs)(nil)).Elem()
}

type ResourcePolicyInput interface {
	pulumi.Input

	ToResourcePolicyOutput() ResourcePolicyOutput
	ToResourcePolicyOutputWithContext(ctx context.Context) ResourcePolicyOutput
}

func (*ResourcePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePolicy)(nil)).Elem()
}

func (i *ResourcePolicy) ToResourcePolicyOutput() ResourcePolicyOutput {
	return i.ToResourcePolicyOutputWithContext(context.Background())
}

func (i *ResourcePolicy) ToResourcePolicyOutputWithContext(ctx context.Context) ResourcePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePolicyOutput)
}

// ResourcePolicyArrayInput is an input type that accepts ResourcePolicyArray and ResourcePolicyArrayOutput values.
// You can construct a concrete instance of `ResourcePolicyArrayInput` via:
//
//	ResourcePolicyArray{ ResourcePolicyArgs{...} }
type ResourcePolicyArrayInput interface {
	pulumi.Input

	ToResourcePolicyArrayOutput() ResourcePolicyArrayOutput
	ToResourcePolicyArrayOutputWithContext(context.Context) ResourcePolicyArrayOutput
}

type ResourcePolicyArray []ResourcePolicyInput

func (ResourcePolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourcePolicy)(nil)).Elem()
}

func (i ResourcePolicyArray) ToResourcePolicyArrayOutput() ResourcePolicyArrayOutput {
	return i.ToResourcePolicyArrayOutputWithContext(context.Background())
}

func (i ResourcePolicyArray) ToResourcePolicyArrayOutputWithContext(ctx context.Context) ResourcePolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePolicyArrayOutput)
}

// ResourcePolicyMapInput is an input type that accepts ResourcePolicyMap and ResourcePolicyMapOutput values.
// You can construct a concrete instance of `ResourcePolicyMapInput` via:
//
//	ResourcePolicyMap{ "key": ResourcePolicyArgs{...} }
type ResourcePolicyMapInput interface {
	pulumi.Input

	ToResourcePolicyMapOutput() ResourcePolicyMapOutput
	ToResourcePolicyMapOutputWithContext(context.Context) ResourcePolicyMapOutput
}

type ResourcePolicyMap map[string]ResourcePolicyInput

func (ResourcePolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourcePolicy)(nil)).Elem()
}

func (i ResourcePolicyMap) ToResourcePolicyMapOutput() ResourcePolicyMapOutput {
	return i.ToResourcePolicyMapOutputWithContext(context.Background())
}

func (i ResourcePolicyMap) ToResourcePolicyMapOutputWithContext(ctx context.Context) ResourcePolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePolicyMapOutput)
}

type ResourcePolicyOutput struct{ *pulumi.OutputState }

func (ResourcePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourcePolicy)(nil)).Elem()
}

func (o ResourcePolicyOutput) ToResourcePolicyOutput() ResourcePolicyOutput {
	return o
}

func (o ResourcePolicyOutput) ToResourcePolicyOutputWithContext(ctx context.Context) ResourcePolicyOutput {
	return o
}

// A JSON-formatted resource policy. For more information, see [Sharing a Projec](https://docs.aws.amazon.com/codebuild/latest/userguide/project-sharing.html#project-sharing-share) and [Sharing a Report Group](https://docs.aws.amazon.com/codebuild/latest/userguide/report-groups-sharing.html#report-groups-sharing-share).
func (o ResourcePolicyOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePolicy) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

// The ARN of the Project or ReportGroup resource you want to associate with a resource policy.
func (o ResourcePolicyOutput) ResourceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ResourcePolicy) pulumi.StringOutput { return v.ResourceArn }).(pulumi.StringOutput)
}

type ResourcePolicyArrayOutput struct{ *pulumi.OutputState }

func (ResourcePolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResourcePolicy)(nil)).Elem()
}

func (o ResourcePolicyArrayOutput) ToResourcePolicyArrayOutput() ResourcePolicyArrayOutput {
	return o
}

func (o ResourcePolicyArrayOutput) ToResourcePolicyArrayOutputWithContext(ctx context.Context) ResourcePolicyArrayOutput {
	return o
}

func (o ResourcePolicyArrayOutput) Index(i pulumi.IntInput) ResourcePolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResourcePolicy {
		return vs[0].([]*ResourcePolicy)[vs[1].(int)]
	}).(ResourcePolicyOutput)
}

type ResourcePolicyMapOutput struct{ *pulumi.OutputState }

func (ResourcePolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResourcePolicy)(nil)).Elem()
}

func (o ResourcePolicyMapOutput) ToResourcePolicyMapOutput() ResourcePolicyMapOutput {
	return o
}

func (o ResourcePolicyMapOutput) ToResourcePolicyMapOutputWithContext(ctx context.Context) ResourcePolicyMapOutput {
	return o
}

func (o ResourcePolicyMapOutput) MapIndex(k pulumi.StringInput) ResourcePolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResourcePolicy {
		return vs[0].(map[string]*ResourcePolicy)[vs[1].(string)]
	}).(ResourcePolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcePolicyInput)(nil)).Elem(), &ResourcePolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcePolicyArrayInput)(nil)).Elem(), ResourcePolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourcePolicyMapInput)(nil)).Elem(), ResourcePolicyMap{})
	pulumi.RegisterOutputType(ResourcePolicyOutput{})
	pulumi.RegisterOutputType(ResourcePolicyArrayOutput{})
	pulumi.RegisterOutputType(ResourcePolicyMapOutput{})
}
