// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ram.ResourceShare` Retrieve information about a RAM Resource Share.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ram"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ram.LookupResourceShare(ctx, &ram.LookupResourceShareArgs{
//				Name:          "example",
//				ResourceOwner: "SELF",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Search by filters
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ram"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ram.LookupResourceShare(ctx, &ram.LookupResourceShareArgs{
//				Filters: []ram.GetResourceShareFilter{
//					{
//						Name: "NameOfTag",
//						Values: []string{
//							"exampleNameTagValue",
//						},
//					},
//				},
//				Name:          "MyResourceName",
//				ResourceOwner: "SELF",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupResourceShare(ctx *pulumi.Context, args *LookupResourceShareArgs, opts ...pulumi.InvokeOption) (*LookupResourceShareResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupResourceShareResult
	err := ctx.Invoke("aws:ram/getResourceShare:getResourceShare", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResourceShare.
type LookupResourceShareArgs struct {
	// Filter used to scope the list e.g., by tags. See [related docs] (https://docs.aws.amazon.com/ram/latest/APIReference/API_TagFilter.html).
	Filters []GetResourceShareFilter `pulumi:"filters"`
	// Name of the tag key to filter on.
	Name string `pulumi:"name"`
	// Owner of the resource share. Valid values are `SELF` or `OTHER-ACCOUNTS`.
	ResourceOwner string `pulumi:"resourceOwner"`
	// Specifies that you want to retrieve details of only those resource shares that have this status. Valid values are `PENDING`, `ACTIVE`, `FAILED`, `DELETING`, and `DELETED`.
	ResourceShareStatus *string `pulumi:"resourceShareStatus"`
	// Tags attached to the RAM share
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getResourceShare.
type LookupResourceShareResult struct {
	// ARN of the resource share.
	Arn     string                   `pulumi:"arn"`
	Filters []GetResourceShareFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// ID of the AWS account that owns the resource share.
	OwningAccountId     string  `pulumi:"owningAccountId"`
	ResourceOwner       string  `pulumi:"resourceOwner"`
	ResourceShareStatus *string `pulumi:"resourceShareStatus"`
	// Status of the RAM share.
	Status string `pulumi:"status"`
	// Tags attached to the RAM share
	Tags map[string]string `pulumi:"tags"`
}

func LookupResourceShareOutput(ctx *pulumi.Context, args LookupResourceShareOutputArgs, opts ...pulumi.InvokeOption) LookupResourceShareResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourceShareResult, error) {
			args := v.(LookupResourceShareArgs)
			r, err := LookupResourceShare(ctx, &args, opts...)
			var s LookupResourceShareResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourceShareResultOutput)
}

// A collection of arguments for invoking getResourceShare.
type LookupResourceShareOutputArgs struct {
	// Filter used to scope the list e.g., by tags. See [related docs] (https://docs.aws.amazon.com/ram/latest/APIReference/API_TagFilter.html).
	Filters GetResourceShareFilterArrayInput `pulumi:"filters"`
	// Name of the tag key to filter on.
	Name pulumi.StringInput `pulumi:"name"`
	// Owner of the resource share. Valid values are `SELF` or `OTHER-ACCOUNTS`.
	ResourceOwner pulumi.StringInput `pulumi:"resourceOwner"`
	// Specifies that you want to retrieve details of only those resource shares that have this status. Valid values are `PENDING`, `ACTIVE`, `FAILED`, `DELETING`, and `DELETED`.
	ResourceShareStatus pulumi.StringPtrInput `pulumi:"resourceShareStatus"`
	// Tags attached to the RAM share
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupResourceShareOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceShareArgs)(nil)).Elem()
}

// A collection of values returned by getResourceShare.
type LookupResourceShareResultOutput struct{ *pulumi.OutputState }

func (LookupResourceShareResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceShareResult)(nil)).Elem()
}

func (o LookupResourceShareResultOutput) ToLookupResourceShareResultOutput() LookupResourceShareResultOutput {
	return o
}

func (o LookupResourceShareResultOutput) ToLookupResourceShareResultOutputWithContext(ctx context.Context) LookupResourceShareResultOutput {
	return o
}

// ARN of the resource share.
func (o LookupResourceShareResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceShareResult) string { return v.Arn }).(pulumi.StringOutput)
}

func (o LookupResourceShareResultOutput) Filters() GetResourceShareFilterArrayOutput {
	return o.ApplyT(func(v LookupResourceShareResult) []GetResourceShareFilter { return v.Filters }).(GetResourceShareFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupResourceShareResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceShareResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupResourceShareResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceShareResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID of the AWS account that owns the resource share.
func (o LookupResourceShareResultOutput) OwningAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceShareResult) string { return v.OwningAccountId }).(pulumi.StringOutput)
}

func (o LookupResourceShareResultOutput) ResourceOwner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceShareResult) string { return v.ResourceOwner }).(pulumi.StringOutput)
}

func (o LookupResourceShareResultOutput) ResourceShareStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceShareResult) *string { return v.ResourceShareStatus }).(pulumi.StringPtrOutput)
}

// Status of the RAM share.
func (o LookupResourceShareResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceShareResult) string { return v.Status }).(pulumi.StringOutput)
}

// Tags attached to the RAM share
func (o LookupResourceShareResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupResourceShareResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceShareResultOutput{})
}
