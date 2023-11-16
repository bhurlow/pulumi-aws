// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ARNs and Names of IAM Roles.
//
// ## Example Usage
// ### All roles in an account
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
//			_, err := iam.GetRoles(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Roles filtered by name regex
//
// Roles whose role-name contains `project`
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
//			_, err := iam.GetRoles(ctx, &iam.GetRolesArgs{
//				NameRegex: pulumi.StringRef(".*project.*"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Roles filtered by path prefix
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
//			_, err := iam.GetRoles(ctx, &iam.GetRolesArgs{
//				PathPrefix: pulumi.StringRef("/custom-path"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Roles provisioned by AWS SSO
//
// # Roles in the account filtered by path prefix
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
//			_, err := iam.GetRoles(ctx, &iam.GetRolesArgs{
//				PathPrefix: pulumi.StringRef("/aws-reserved/sso.amazonaws.com/"),
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
// # Specific role in the account filtered by name regex and path prefix
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
//			_, err := iam.GetRoles(ctx, &iam.GetRolesArgs{
//				NameRegex:  pulumi.StringRef("AWSReservedSSO_permission_set_name_.*"),
//				PathPrefix: pulumi.StringRef("/aws-reserved/sso.amazonaws.com/"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetRoles(ctx *pulumi.Context, args *GetRolesArgs, opts ...pulumi.InvokeOption) (*GetRolesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRolesResult
	err := ctx.Invoke("aws:iam/getRoles:getRoles", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRoles.
type GetRolesArgs struct {
	// Regex string to apply to the IAM roles list returned by AWS. This allows more advanced filtering not supported from the AWS API. This filtering is done locally on what AWS returns, and could have a performance impact if the result is large. Combine this with other options to narrow down the list AWS returns.
	NameRegex *string `pulumi:"nameRegex"`
	// Path prefix for filtering the results. For example, the prefix `/application_abc/component_xyz/` gets all roles whose path starts with `/application_abc/component_xyz/`. If it is not included, it defaults to a slash (`/`), listing all roles. For more details, check out [list-roles in the AWS CLI reference][1].
	PathPrefix *string `pulumi:"pathPrefix"`
}

// A collection of values returned by getRoles.
type GetRolesResult struct {
	// Set of ARNs of the matched IAM roles.
	Arns []string `pulumi:"arns"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	NameRegex *string `pulumi:"nameRegex"`
	// Set of Names of the matched IAM roles.
	Names      []string `pulumi:"names"`
	PathPrefix *string  `pulumi:"pathPrefix"`
}

func GetRolesOutput(ctx *pulumi.Context, args GetRolesOutputArgs, opts ...pulumi.InvokeOption) GetRolesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRolesResult, error) {
			args := v.(GetRolesArgs)
			r, err := GetRoles(ctx, &args, opts...)
			var s GetRolesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRolesResultOutput)
}

// A collection of arguments for invoking getRoles.
type GetRolesOutputArgs struct {
	// Regex string to apply to the IAM roles list returned by AWS. This allows more advanced filtering not supported from the AWS API. This filtering is done locally on what AWS returns, and could have a performance impact if the result is large. Combine this with other options to narrow down the list AWS returns.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// Path prefix for filtering the results. For example, the prefix `/application_abc/component_xyz/` gets all roles whose path starts with `/application_abc/component_xyz/`. If it is not included, it defaults to a slash (`/`), listing all roles. For more details, check out [list-roles in the AWS CLI reference][1].
	PathPrefix pulumi.StringPtrInput `pulumi:"pathPrefix"`
}

func (GetRolesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRolesArgs)(nil)).Elem()
}

// A collection of values returned by getRoles.
type GetRolesResultOutput struct{ *pulumi.OutputState }

func (GetRolesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRolesResult)(nil)).Elem()
}

func (o GetRolesResultOutput) ToGetRolesResultOutput() GetRolesResultOutput {
	return o
}

func (o GetRolesResultOutput) ToGetRolesResultOutputWithContext(ctx context.Context) GetRolesResultOutput {
	return o
}

// Set of ARNs of the matched IAM roles.
func (o GetRolesResultOutput) Arns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRolesResult) []string { return v.Arns }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRolesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRolesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRolesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRolesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// Set of Names of the matched IAM roles.
func (o GetRolesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRolesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o GetRolesResultOutput) PathPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRolesResult) *string { return v.PathPrefix }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRolesResultOutput{})
}
