// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kendra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about a specific Amazon Kendra block list used for query suggestions for an index.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kendra"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := kendra.LookupQuerySuggestionsBlockList(ctx, &kendra.LookupQuerySuggestionsBlockListArgs{
//				IndexId:                     "12345678-1234-1234-1234-123456789123",
//				QuerySuggestionsBlockListId: "87654321-1234-4321-4321-321987654321",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupQuerySuggestionsBlockList(ctx *pulumi.Context, args *LookupQuerySuggestionsBlockListArgs, opts ...pulumi.InvokeOption) (*LookupQuerySuggestionsBlockListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupQuerySuggestionsBlockListResult
	err := ctx.Invoke("aws:kendra/getQuerySuggestionsBlockList:getQuerySuggestionsBlockList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getQuerySuggestionsBlockList.
type LookupQuerySuggestionsBlockListArgs struct {
	// Identifier of the index that contains the block list.
	IndexId string `pulumi:"indexId"`
	// Identifier of the block list.
	QuerySuggestionsBlockListId string `pulumi:"querySuggestionsBlockListId"`
	// Metadata that helps organize the block list you create.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getQuerySuggestionsBlockList.
type LookupQuerySuggestionsBlockListResult struct {
	// ARN of the block list.
	Arn string `pulumi:"arn"`
	// Date-time a block list was created.
	CreatedAt string `pulumi:"createdAt"`
	// Description for the block list.
	Description string `pulumi:"description"`
	// Error message containing details if there are issues processing the block list.
	ErrorMessage string `pulumi:"errorMessage"`
	// Current size of the block list text file in S3.
	FileSizeBytes int `pulumi:"fileSizeBytes"`
	// The provider-assigned unique ID for this managed resource.
	Id      string `pulumi:"id"`
	IndexId string `pulumi:"indexId"`
	// Current number of valid, non-empty words or phrases in the block list text file.
	ItemCount int `pulumi:"itemCount"`
	// Name of the block list.
	Name                        string `pulumi:"name"`
	QuerySuggestionsBlockListId string `pulumi:"querySuggestionsBlockListId"`
	// ARN of a role with permission to access the S3 bucket that contains the block list. For more information, see [IAM Roles for Amazon Kendra](https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html).
	RoleArn string `pulumi:"roleArn"`
	// S3 location of the block list input data. Detailed below.
	SourceS3Paths []GetQuerySuggestionsBlockListSourceS3Path `pulumi:"sourceS3Paths"`
	// Current status of the block list. When the value is `ACTIVE`, the block list is ready for use.
	Status string `pulumi:"status"`
	// Metadata that helps organize the block list you create.
	Tags map[string]string `pulumi:"tags"`
	// Date and time that the block list was last updated.
	UpdatedAt string `pulumi:"updatedAt"`
}

func LookupQuerySuggestionsBlockListOutput(ctx *pulumi.Context, args LookupQuerySuggestionsBlockListOutputArgs, opts ...pulumi.InvokeOption) LookupQuerySuggestionsBlockListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupQuerySuggestionsBlockListResult, error) {
			args := v.(LookupQuerySuggestionsBlockListArgs)
			r, err := LookupQuerySuggestionsBlockList(ctx, &args, opts...)
			var s LookupQuerySuggestionsBlockListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupQuerySuggestionsBlockListResultOutput)
}

// A collection of arguments for invoking getQuerySuggestionsBlockList.
type LookupQuerySuggestionsBlockListOutputArgs struct {
	// Identifier of the index that contains the block list.
	IndexId pulumi.StringInput `pulumi:"indexId"`
	// Identifier of the block list.
	QuerySuggestionsBlockListId pulumi.StringInput `pulumi:"querySuggestionsBlockListId"`
	// Metadata that helps organize the block list you create.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupQuerySuggestionsBlockListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQuerySuggestionsBlockListArgs)(nil)).Elem()
}

// A collection of values returned by getQuerySuggestionsBlockList.
type LookupQuerySuggestionsBlockListResultOutput struct{ *pulumi.OutputState }

func (LookupQuerySuggestionsBlockListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQuerySuggestionsBlockListResult)(nil)).Elem()
}

func (o LookupQuerySuggestionsBlockListResultOutput) ToLookupQuerySuggestionsBlockListResultOutput() LookupQuerySuggestionsBlockListResultOutput {
	return o
}

func (o LookupQuerySuggestionsBlockListResultOutput) ToLookupQuerySuggestionsBlockListResultOutputWithContext(ctx context.Context) LookupQuerySuggestionsBlockListResultOutput {
	return o
}

// ARN of the block list.
func (o LookupQuerySuggestionsBlockListResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQuerySuggestionsBlockListResult) string { return v.Arn }).(pulumi.StringOutput)
}

// Date-time a block list was created.
func (o LookupQuerySuggestionsBlockListResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQuerySuggestionsBlockListResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Description for the block list.
func (o LookupQuerySuggestionsBlockListResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQuerySuggestionsBlockListResult) string { return v.Description }).(pulumi.StringOutput)
}

// Error message containing details if there are issues processing the block list.
func (o LookupQuerySuggestionsBlockListResultOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQuerySuggestionsBlockListResult) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

// Current size of the block list text file in S3.
func (o LookupQuerySuggestionsBlockListResultOutput) FileSizeBytes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQuerySuggestionsBlockListResult) int { return v.FileSizeBytes }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupQuerySuggestionsBlockListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQuerySuggestionsBlockListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupQuerySuggestionsBlockListResultOutput) IndexId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQuerySuggestionsBlockListResult) string { return v.IndexId }).(pulumi.StringOutput)
}

// Current number of valid, non-empty words or phrases in the block list text file.
func (o LookupQuerySuggestionsBlockListResultOutput) ItemCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQuerySuggestionsBlockListResult) int { return v.ItemCount }).(pulumi.IntOutput)
}

// Name of the block list.
func (o LookupQuerySuggestionsBlockListResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQuerySuggestionsBlockListResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupQuerySuggestionsBlockListResultOutput) QuerySuggestionsBlockListId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQuerySuggestionsBlockListResult) string { return v.QuerySuggestionsBlockListId }).(pulumi.StringOutput)
}

// ARN of a role with permission to access the S3 bucket that contains the block list. For more information, see [IAM Roles for Amazon Kendra](https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html).
func (o LookupQuerySuggestionsBlockListResultOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQuerySuggestionsBlockListResult) string { return v.RoleArn }).(pulumi.StringOutput)
}

// S3 location of the block list input data. Detailed below.
func (o LookupQuerySuggestionsBlockListResultOutput) SourceS3Paths() GetQuerySuggestionsBlockListSourceS3PathArrayOutput {
	return o.ApplyT(func(v LookupQuerySuggestionsBlockListResult) []GetQuerySuggestionsBlockListSourceS3Path {
		return v.SourceS3Paths
	}).(GetQuerySuggestionsBlockListSourceS3PathArrayOutput)
}

// Current status of the block list. When the value is `ACTIVE`, the block list is ready for use.
func (o LookupQuerySuggestionsBlockListResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQuerySuggestionsBlockListResult) string { return v.Status }).(pulumi.StringOutput)
}

// Metadata that helps organize the block list you create.
func (o LookupQuerySuggestionsBlockListResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupQuerySuggestionsBlockListResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Date and time that the block list was last updated.
func (o LookupQuerySuggestionsBlockListResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQuerySuggestionsBlockListResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQuerySuggestionsBlockListResultOutput{})
}
