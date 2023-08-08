// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lakeformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get Lake Formation principals designated as data lake administrators and lists of principal permission entries for default create database and default create table permissions.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lakeformation"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := lakeformation.LookupDataLakeSettings(ctx, &lakeformation.LookupDataLakeSettingsArgs{
//				CatalogId: pulumi.StringRef("14916253649"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDataLakeSettings(ctx *pulumi.Context, args *LookupDataLakeSettingsArgs, opts ...pulumi.InvokeOption) (*LookupDataLakeSettingsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDataLakeSettingsResult
	err := ctx.Invoke("aws:lakeformation/getDataLakeSettings:getDataLakeSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDataLakeSettings.
type LookupDataLakeSettingsArgs struct {
	// Identifier for the Data Catalog. By default, the account ID.
	CatalogId *string `pulumi:"catalogId"`
}

// A collection of values returned by getDataLakeSettings.
type LookupDataLakeSettingsResult struct {
	// List of ARNs of AWS Lake Formation principals (IAM users or roles).
	Admins []string `pulumi:"admins"`
	// Whether to allow Amazon EMR clusters to access data managed by Lake Formation.
	AllowExternalDataFiltering bool `pulumi:"allowExternalDataFiltering"`
	// Lake Formation relies on a privileged process secured by Amazon EMR or the third party integrator to tag the user's role while assuming it.
	AuthorizedSessionTagValueLists []string `pulumi:"authorizedSessionTagValueLists"`
	CatalogId                      *string  `pulumi:"catalogId"`
	// Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
	CreateDatabaseDefaultPermissions []GetDataLakeSettingsCreateDatabaseDefaultPermission `pulumi:"createDatabaseDefaultPermissions"`
	// Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
	CreateTableDefaultPermissions []GetDataLakeSettingsCreateTableDefaultPermission `pulumi:"createTableDefaultPermissions"`
	// A list of the account IDs of Amazon Web Services accounts with Amazon EMR clusters that are to perform data filtering.
	ExternalDataFilteringAllowLists []string `pulumi:"externalDataFilteringAllowLists"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
	TrustedResourceOwners []string `pulumi:"trustedResourceOwners"`
}

func LookupDataLakeSettingsOutput(ctx *pulumi.Context, args LookupDataLakeSettingsOutputArgs, opts ...pulumi.InvokeOption) LookupDataLakeSettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataLakeSettingsResult, error) {
			args := v.(LookupDataLakeSettingsArgs)
			r, err := LookupDataLakeSettings(ctx, &args, opts...)
			var s LookupDataLakeSettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataLakeSettingsResultOutput)
}

// A collection of arguments for invoking getDataLakeSettings.
type LookupDataLakeSettingsOutputArgs struct {
	// Identifier for the Data Catalog. By default, the account ID.
	CatalogId pulumi.StringPtrInput `pulumi:"catalogId"`
}

func (LookupDataLakeSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataLakeSettingsArgs)(nil)).Elem()
}

// A collection of values returned by getDataLakeSettings.
type LookupDataLakeSettingsResultOutput struct{ *pulumi.OutputState }

func (LookupDataLakeSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataLakeSettingsResult)(nil)).Elem()
}

func (o LookupDataLakeSettingsResultOutput) ToLookupDataLakeSettingsResultOutput() LookupDataLakeSettingsResultOutput {
	return o
}

func (o LookupDataLakeSettingsResultOutput) ToLookupDataLakeSettingsResultOutputWithContext(ctx context.Context) LookupDataLakeSettingsResultOutput {
	return o
}

// List of ARNs of AWS Lake Formation principals (IAM users or roles).
func (o LookupDataLakeSettingsResultOutput) Admins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDataLakeSettingsResult) []string { return v.Admins }).(pulumi.StringArrayOutput)
}

// Whether to allow Amazon EMR clusters to access data managed by Lake Formation.
func (o LookupDataLakeSettingsResultOutput) AllowExternalDataFiltering() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDataLakeSettingsResult) bool { return v.AllowExternalDataFiltering }).(pulumi.BoolOutput)
}

// Lake Formation relies on a privileged process secured by Amazon EMR or the third party integrator to tag the user's role while assuming it.
func (o LookupDataLakeSettingsResultOutput) AuthorizedSessionTagValueLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDataLakeSettingsResult) []string { return v.AuthorizedSessionTagValueLists }).(pulumi.StringArrayOutput)
}

func (o LookupDataLakeSettingsResultOutput) CatalogId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataLakeSettingsResult) *string { return v.CatalogId }).(pulumi.StringPtrOutput)
}

// Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
func (o LookupDataLakeSettingsResultOutput) CreateDatabaseDefaultPermissions() GetDataLakeSettingsCreateDatabaseDefaultPermissionArrayOutput {
	return o.ApplyT(func(v LookupDataLakeSettingsResult) []GetDataLakeSettingsCreateDatabaseDefaultPermission {
		return v.CreateDatabaseDefaultPermissions
	}).(GetDataLakeSettingsCreateDatabaseDefaultPermissionArrayOutput)
}

// Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
func (o LookupDataLakeSettingsResultOutput) CreateTableDefaultPermissions() GetDataLakeSettingsCreateTableDefaultPermissionArrayOutput {
	return o.ApplyT(func(v LookupDataLakeSettingsResult) []GetDataLakeSettingsCreateTableDefaultPermission {
		return v.CreateTableDefaultPermissions
	}).(GetDataLakeSettingsCreateTableDefaultPermissionArrayOutput)
}

// A list of the account IDs of Amazon Web Services accounts with Amazon EMR clusters that are to perform data filtering.
func (o LookupDataLakeSettingsResultOutput) ExternalDataFilteringAllowLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDataLakeSettingsResult) []string { return v.ExternalDataFilteringAllowLists }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDataLakeSettingsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataLakeSettingsResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
func (o LookupDataLakeSettingsResultOutput) TrustedResourceOwners() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDataLakeSettingsResult) []string { return v.TrustedResourceOwners }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataLakeSettingsResultOutput{})
}
