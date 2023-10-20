// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This data source provides information on the IAM source role of an STS assumed role. For non-role ARNs, this data source simply passes the ARN through in `issuerArn`.
//
// For some AWS resources, multiple types of principals are allowed in the same argument (e.g., IAM users and IAM roles). However, these arguments often do not allow assumed-role (i.e., STS, temporary credential) principals. Given an STS ARN, this data source provides the ARN for the source IAM role.
//
// ## Example Usage
// ### Basic Example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iam.GetSessionContext(ctx, &iam.GetSessionContextArgs{
//				Arn: "arn:aws:sts::123456789012:assumed-role/Audien-Heaven/MatyNoyes",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Find the Runner's Source Role
//
// Combined with `getCallerIdentity`, you can get the current user's source IAM role ARN (`issuerArn`) if you're using an assumed role. If you're not using an assumed role, the caller's (e.g., an IAM user's) ARN will simply be passed through. In environments where both IAM users and individuals using assumed roles need to apply the same configurations, this data source enables seamless use.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			current, err := aws.GetCallerIdentity(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = iam.GetSessionContext(ctx, &iam.GetSessionContextArgs{
//				Arn: current.Arn,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetSessionContext(ctx *pulumi.Context, args *GetSessionContextArgs, opts ...pulumi.InvokeOption) (*GetSessionContextResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSessionContextResult
	err := ctx.Invoke("aws:iam/getSessionContext:getSessionContext", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSessionContext.
type GetSessionContextArgs struct {
	// ARN for an assumed role.
	//
	// > If `arn` is a non-role ARN, Pulumi gives no error and `issuerArn` will be equal to the `arn` value. For STS assumed-role ARNs, Pulumi gives an error if the identified IAM role does not exist.
	Arn string `pulumi:"arn"`
}

// A collection of values returned by getSessionContext.
type GetSessionContextResult struct {
	Arn string `pulumi:"arn"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IAM source role ARN if `arn` corresponds to an STS assumed role. Otherwise, `issuerArn` is equal to `arn`.
	IssuerArn string `pulumi:"issuerArn"`
	// Unique identifier of the IAM role that issues the STS assumed role.
	IssuerId string `pulumi:"issuerId"`
	// Name of the source role. Only available if `arn` corresponds to an STS assumed role.
	IssuerName string `pulumi:"issuerName"`
	// Name of the STS session. Only available if `arn` corresponds to an STS assumed role.
	SessionName string `pulumi:"sessionName"`
}

func GetSessionContextOutput(ctx *pulumi.Context, args GetSessionContextOutputArgs, opts ...pulumi.InvokeOption) GetSessionContextResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSessionContextResult, error) {
			args := v.(GetSessionContextArgs)
			r, err := GetSessionContext(ctx, &args, opts...)
			var s GetSessionContextResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSessionContextResultOutput)
}

// A collection of arguments for invoking getSessionContext.
type GetSessionContextOutputArgs struct {
	// ARN for an assumed role.
	//
	// > If `arn` is a non-role ARN, Pulumi gives no error and `issuerArn` will be equal to the `arn` value. For STS assumed-role ARNs, Pulumi gives an error if the identified IAM role does not exist.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (GetSessionContextOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSessionContextArgs)(nil)).Elem()
}

// A collection of values returned by getSessionContext.
type GetSessionContextResultOutput struct{ *pulumi.OutputState }

func (GetSessionContextResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSessionContextResult)(nil)).Elem()
}

func (o GetSessionContextResultOutput) ToGetSessionContextResultOutput() GetSessionContextResultOutput {
	return o
}

func (o GetSessionContextResultOutput) ToGetSessionContextResultOutputWithContext(ctx context.Context) GetSessionContextResultOutput {
	return o
}

func (o GetSessionContextResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetSessionContextResult] {
	return pulumix.Output[GetSessionContextResult]{
		OutputState: o.OutputState,
	}
}

func (o GetSessionContextResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v GetSessionContextResult) string { return v.Arn }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSessionContextResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSessionContextResult) string { return v.Id }).(pulumi.StringOutput)
}

// IAM source role ARN if `arn` corresponds to an STS assumed role. Otherwise, `issuerArn` is equal to `arn`.
func (o GetSessionContextResultOutput) IssuerArn() pulumi.StringOutput {
	return o.ApplyT(func(v GetSessionContextResult) string { return v.IssuerArn }).(pulumi.StringOutput)
}

// Unique identifier of the IAM role that issues the STS assumed role.
func (o GetSessionContextResultOutput) IssuerId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSessionContextResult) string { return v.IssuerId }).(pulumi.StringOutput)
}

// Name of the source role. Only available if `arn` corresponds to an STS assumed role.
func (o GetSessionContextResultOutput) IssuerName() pulumi.StringOutput {
	return o.ApplyT(func(v GetSessionContextResult) string { return v.IssuerName }).(pulumi.StringOutput)
}

// Name of the STS session. Only available if `arn` corresponds to an STS assumed role.
func (o GetSessionContextResultOutput) SessionName() pulumi.StringOutput {
	return o.ApplyT(func(v GetSessionContextResult) string { return v.SessionName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSessionContextResultOutput{})
}
