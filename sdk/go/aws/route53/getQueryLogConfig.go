// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `route53.ResolverQueryLogConfig` provides details about a specific Route53 Resolver Query Logging Configuration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := route53.GetQueryLogConfig(ctx, &route53.GetQueryLogConfigArgs{
//				ResolverQueryLogConfigId: pulumi.StringRef("rqlc-1abc2345ef678g91h"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := route53.GetQueryLogConfig(ctx, &route53.GetQueryLogConfigArgs{
//				Filters: []route53.GetQueryLogConfigFilter{
//					{
//						Name: "Name",
//						Values: []string{
//							"shared-query-log-config",
//						},
//					},
//					{
//						Name: "ShareStatus",
//						Values: []string{
//							"SHARED_WITH_ME",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetQueryLogConfig(ctx *pulumi.Context, args *GetQueryLogConfigArgs, opts ...pulumi.InvokeOption) (*GetQueryLogConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetQueryLogConfigResult
	err := ctx.Invoke("aws:route53/getQueryLogConfig:getQueryLogConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getQueryLogConfig.
type GetQueryLogConfigArgs struct {
	// One or more name/value pairs to use as filters. There are
	// several valid keys, for a full reference, check out
	// [Route53resolver Filter value in the AWS API reference][1].
	//
	// In addition to all arguments above, the following attributes are exported:
	Filters []GetQueryLogConfigFilter `pulumi:"filters"`
	// The name of the query logging configuration.
	Name *string `pulumi:"name"`
	// ID of the Route53 Resolver Query Logging Configuration.
	ResolverQueryLogConfigId *string `pulumi:"resolverQueryLogConfigId"`
	// Map of tags to assign to the service.
	//
	// [1]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_Filter.html
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getQueryLogConfig.
type GetQueryLogConfigResult struct {
	Arn            string                    `pulumi:"arn"`
	DestinationArn string                    `pulumi:"destinationArn"`
	Filters        []GetQueryLogConfigFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id                       string            `pulumi:"id"`
	Name                     *string           `pulumi:"name"`
	OwnerId                  string            `pulumi:"ownerId"`
	ResolverQueryLogConfigId *string           `pulumi:"resolverQueryLogConfigId"`
	ShareStatus              string            `pulumi:"shareStatus"`
	Tags                     map[string]string `pulumi:"tags"`
}

func GetQueryLogConfigOutput(ctx *pulumi.Context, args GetQueryLogConfigOutputArgs, opts ...pulumi.InvokeOption) GetQueryLogConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetQueryLogConfigResult, error) {
			args := v.(GetQueryLogConfigArgs)
			r, err := GetQueryLogConfig(ctx, &args, opts...)
			var s GetQueryLogConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetQueryLogConfigResultOutput)
}

// A collection of arguments for invoking getQueryLogConfig.
type GetQueryLogConfigOutputArgs struct {
	// One or more name/value pairs to use as filters. There are
	// several valid keys, for a full reference, check out
	// [Route53resolver Filter value in the AWS API reference][1].
	//
	// In addition to all arguments above, the following attributes are exported:
	Filters GetQueryLogConfigFilterArrayInput `pulumi:"filters"`
	// The name of the query logging configuration.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// ID of the Route53 Resolver Query Logging Configuration.
	ResolverQueryLogConfigId pulumi.StringPtrInput `pulumi:"resolverQueryLogConfigId"`
	// Map of tags to assign to the service.
	//
	// [1]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_Filter.html
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (GetQueryLogConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetQueryLogConfigArgs)(nil)).Elem()
}

// A collection of values returned by getQueryLogConfig.
type GetQueryLogConfigResultOutput struct{ *pulumi.OutputState }

func (GetQueryLogConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetQueryLogConfigResult)(nil)).Elem()
}

func (o GetQueryLogConfigResultOutput) ToGetQueryLogConfigResultOutput() GetQueryLogConfigResultOutput {
	return o
}

func (o GetQueryLogConfigResultOutput) ToGetQueryLogConfigResultOutputWithContext(ctx context.Context) GetQueryLogConfigResultOutput {
	return o
}

func (o GetQueryLogConfigResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v GetQueryLogConfigResult) string { return v.Arn }).(pulumi.StringOutput)
}

func (o GetQueryLogConfigResultOutput) DestinationArn() pulumi.StringOutput {
	return o.ApplyT(func(v GetQueryLogConfigResult) string { return v.DestinationArn }).(pulumi.StringOutput)
}

func (o GetQueryLogConfigResultOutput) Filters() GetQueryLogConfigFilterArrayOutput {
	return o.ApplyT(func(v GetQueryLogConfigResult) []GetQueryLogConfigFilter { return v.Filters }).(GetQueryLogConfigFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetQueryLogConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetQueryLogConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetQueryLogConfigResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetQueryLogConfigResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetQueryLogConfigResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v GetQueryLogConfigResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

func (o GetQueryLogConfigResultOutput) ResolverQueryLogConfigId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetQueryLogConfigResult) *string { return v.ResolverQueryLogConfigId }).(pulumi.StringPtrOutput)
}

func (o GetQueryLogConfigResultOutput) ShareStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetQueryLogConfigResult) string { return v.ShareStatus }).(pulumi.StringOutput)
}

func (o GetQueryLogConfigResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetQueryLogConfigResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetQueryLogConfigResultOutput{})
}
