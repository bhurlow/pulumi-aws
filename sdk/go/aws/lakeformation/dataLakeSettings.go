// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lakeformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages Lake Formation principals designated as data lake administrators and lists of principal permission entries for default create database and default create table permissions.
//
// > **NOTE:** Lake Formation introduces fine-grained access control for data in your data lake. Part of the changes include the `IAMAllowedPrincipals` principal in order to make Lake Formation backwards compatible with existing IAM and Glue permissions. For more information, see [Changing the Default Security Settings for Your Data Lake](https://docs.aws.amazon.com/lake-formation/latest/dg/change-settings.html) and [Upgrading AWS Glue Data Permissions to the AWS Lake Formation Model](https://docs.aws.amazon.com/lake-formation/latest/dg/upgrade-glue-lake-formation.html).
//
// ## Example Usage
// ### Data Lake Admins
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
//			_, err := lakeformation.NewDataLakeSettings(ctx, "example", &lakeformation.DataLakeSettingsArgs{
//				Admins: pulumi.StringArray{
//					aws_iam_user.Test.Arn,
//					aws_iam_role.Test.Arn,
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
// ### Create Default Permissions
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
//			_, err := lakeformation.NewDataLakeSettings(ctx, "example", &lakeformation.DataLakeSettingsArgs{
//				Admins: pulumi.StringArray{
//					aws_iam_user.Test.Arn,
//					aws_iam_role.Test.Arn,
//				},
//				CreateDatabaseDefaultPermissions: lakeformation.DataLakeSettingsCreateDatabaseDefaultPermissionArray{
//					&lakeformation.DataLakeSettingsCreateDatabaseDefaultPermissionArgs{
//						Permissions: pulumi.StringArray{
//							pulumi.String("SELECT"),
//							pulumi.String("ALTER"),
//							pulumi.String("DROP"),
//						},
//						Principal: pulumi.Any(aws_iam_user.Test.Arn),
//					},
//				},
//				CreateTableDefaultPermissions: lakeformation.DataLakeSettingsCreateTableDefaultPermissionArray{
//					&lakeformation.DataLakeSettingsCreateTableDefaultPermissionArgs{
//						Permissions: pulumi.StringArray{
//							pulumi.String("ALL"),
//						},
//						Principal: pulumi.Any(aws_iam_role.Test.Arn),
//					},
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
// ### Enable EMR access to LakeFormation resources
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
//			_, err := lakeformation.NewDataLakeSettings(ctx, "example", &lakeformation.DataLakeSettingsArgs{
//				Admins: pulumi.StringArray{
//					aws_iam_user.Test.Arn,
//					aws_iam_role.Test.Arn,
//				},
//				CreateDatabaseDefaultPermissions: lakeformation.DataLakeSettingsCreateDatabaseDefaultPermissionArray{
//					&lakeformation.DataLakeSettingsCreateDatabaseDefaultPermissionArgs{
//						Permissions: pulumi.StringArray{
//							pulumi.String("SELECT"),
//							pulumi.String("ALTER"),
//							pulumi.String("DROP"),
//						},
//						Principal: pulumi.Any(aws_iam_user.Test.Arn),
//					},
//				},
//				CreateTableDefaultPermissions: lakeformation.DataLakeSettingsCreateTableDefaultPermissionArray{
//					&lakeformation.DataLakeSettingsCreateTableDefaultPermissionArgs{
//						Permissions: pulumi.StringArray{
//							pulumi.String("ALL"),
//						},
//						Principal: pulumi.Any(aws_iam_role.Test.Arn),
//					},
//				},
//				AllowExternalDataFiltering: pulumi.Bool(true),
//				ExternalDataFilteringAllowLists: pulumi.StringArray{
//					data.Aws_caller_identity.Current.Account_id,
//					data.Aws_caller_identity.Third_party.Account_id,
//				},
//				AuthorizedSessionTagValueLists: pulumi.StringArray{
//					pulumi.String("Amazon EMR"),
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
type DataLakeSettings struct {
	pulumi.CustomResourceState

	// Set of ARNs of AWS Lake Formation principals (IAM users or roles).
	Admins pulumi.StringArrayOutput `pulumi:"admins"`
	// Whether to allow Amazon EMR clusters to access data managed by Lake Formation.
	AllowExternalDataFiltering pulumi.BoolPtrOutput `pulumi:"allowExternalDataFiltering"`
	// Lake Formation relies on a privileged process secured by Amazon EMR or the third party integrator to tag the user's role while assuming it.
	//
	// > **NOTE:** Although optional, not including `admins`, `createDatabaseDefaultPermissions`, `createTableDefaultPermissions`, and/or `trustedResourceOwners` results in the setting being cleared.
	AuthorizedSessionTagValueLists pulumi.StringArrayOutput `pulumi:"authorizedSessionTagValueLists"`
	// Identifier for the Data Catalog. By default, the account ID.
	CatalogId pulumi.StringPtrOutput `pulumi:"catalogId"`
	// Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
	CreateDatabaseDefaultPermissions DataLakeSettingsCreateDatabaseDefaultPermissionArrayOutput `pulumi:"createDatabaseDefaultPermissions"`
	// Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
	CreateTableDefaultPermissions DataLakeSettingsCreateTableDefaultPermissionArrayOutput `pulumi:"createTableDefaultPermissions"`
	// A list of the account IDs of Amazon Web Services accounts with Amazon EMR clusters that are to perform data filtering.
	ExternalDataFilteringAllowLists pulumi.StringArrayOutput `pulumi:"externalDataFilteringAllowLists"`
	// Set of ARNs of AWS Lake Formation principals (IAM users or roles) with only view access to the resources.
	ReadOnlyAdmins pulumi.StringArrayOutput `pulumi:"readOnlyAdmins"`
	// List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
	TrustedResourceOwners pulumi.StringArrayOutput `pulumi:"trustedResourceOwners"`
}

// NewDataLakeSettings registers a new resource with the given unique name, arguments, and options.
func NewDataLakeSettings(ctx *pulumi.Context,
	name string, args *DataLakeSettingsArgs, opts ...pulumi.ResourceOption) (*DataLakeSettings, error) {
	if args == nil {
		args = &DataLakeSettingsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DataLakeSettings
	err := ctx.RegisterResource("aws:lakeformation/dataLakeSettings:DataLakeSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataLakeSettings gets an existing DataLakeSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataLakeSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataLakeSettingsState, opts ...pulumi.ResourceOption) (*DataLakeSettings, error) {
	var resource DataLakeSettings
	err := ctx.ReadResource("aws:lakeformation/dataLakeSettings:DataLakeSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataLakeSettings resources.
type dataLakeSettingsState struct {
	// Set of ARNs of AWS Lake Formation principals (IAM users or roles).
	Admins []string `pulumi:"admins"`
	// Whether to allow Amazon EMR clusters to access data managed by Lake Formation.
	AllowExternalDataFiltering *bool `pulumi:"allowExternalDataFiltering"`
	// Lake Formation relies on a privileged process secured by Amazon EMR or the third party integrator to tag the user's role while assuming it.
	//
	// > **NOTE:** Although optional, not including `admins`, `createDatabaseDefaultPermissions`, `createTableDefaultPermissions`, and/or `trustedResourceOwners` results in the setting being cleared.
	AuthorizedSessionTagValueLists []string `pulumi:"authorizedSessionTagValueLists"`
	// Identifier for the Data Catalog. By default, the account ID.
	CatalogId *string `pulumi:"catalogId"`
	// Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
	CreateDatabaseDefaultPermissions []DataLakeSettingsCreateDatabaseDefaultPermission `pulumi:"createDatabaseDefaultPermissions"`
	// Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
	CreateTableDefaultPermissions []DataLakeSettingsCreateTableDefaultPermission `pulumi:"createTableDefaultPermissions"`
	// A list of the account IDs of Amazon Web Services accounts with Amazon EMR clusters that are to perform data filtering.
	ExternalDataFilteringAllowLists []string `pulumi:"externalDataFilteringAllowLists"`
	// Set of ARNs of AWS Lake Formation principals (IAM users or roles) with only view access to the resources.
	ReadOnlyAdmins []string `pulumi:"readOnlyAdmins"`
	// List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
	TrustedResourceOwners []string `pulumi:"trustedResourceOwners"`
}

type DataLakeSettingsState struct {
	// Set of ARNs of AWS Lake Formation principals (IAM users or roles).
	Admins pulumi.StringArrayInput
	// Whether to allow Amazon EMR clusters to access data managed by Lake Formation.
	AllowExternalDataFiltering pulumi.BoolPtrInput
	// Lake Formation relies on a privileged process secured by Amazon EMR or the third party integrator to tag the user's role while assuming it.
	//
	// > **NOTE:** Although optional, not including `admins`, `createDatabaseDefaultPermissions`, `createTableDefaultPermissions`, and/or `trustedResourceOwners` results in the setting being cleared.
	AuthorizedSessionTagValueLists pulumi.StringArrayInput
	// Identifier for the Data Catalog. By default, the account ID.
	CatalogId pulumi.StringPtrInput
	// Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
	CreateDatabaseDefaultPermissions DataLakeSettingsCreateDatabaseDefaultPermissionArrayInput
	// Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
	CreateTableDefaultPermissions DataLakeSettingsCreateTableDefaultPermissionArrayInput
	// A list of the account IDs of Amazon Web Services accounts with Amazon EMR clusters that are to perform data filtering.
	ExternalDataFilteringAllowLists pulumi.StringArrayInput
	// Set of ARNs of AWS Lake Formation principals (IAM users or roles) with only view access to the resources.
	ReadOnlyAdmins pulumi.StringArrayInput
	// List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
	TrustedResourceOwners pulumi.StringArrayInput
}

func (DataLakeSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataLakeSettingsState)(nil)).Elem()
}

type dataLakeSettingsArgs struct {
	// Set of ARNs of AWS Lake Formation principals (IAM users or roles).
	Admins []string `pulumi:"admins"`
	// Whether to allow Amazon EMR clusters to access data managed by Lake Formation.
	AllowExternalDataFiltering *bool `pulumi:"allowExternalDataFiltering"`
	// Lake Formation relies on a privileged process secured by Amazon EMR or the third party integrator to tag the user's role while assuming it.
	//
	// > **NOTE:** Although optional, not including `admins`, `createDatabaseDefaultPermissions`, `createTableDefaultPermissions`, and/or `trustedResourceOwners` results in the setting being cleared.
	AuthorizedSessionTagValueLists []string `pulumi:"authorizedSessionTagValueLists"`
	// Identifier for the Data Catalog. By default, the account ID.
	CatalogId *string `pulumi:"catalogId"`
	// Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
	CreateDatabaseDefaultPermissions []DataLakeSettingsCreateDatabaseDefaultPermission `pulumi:"createDatabaseDefaultPermissions"`
	// Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
	CreateTableDefaultPermissions []DataLakeSettingsCreateTableDefaultPermission `pulumi:"createTableDefaultPermissions"`
	// A list of the account IDs of Amazon Web Services accounts with Amazon EMR clusters that are to perform data filtering.
	ExternalDataFilteringAllowLists []string `pulumi:"externalDataFilteringAllowLists"`
	// Set of ARNs of AWS Lake Formation principals (IAM users or roles) with only view access to the resources.
	ReadOnlyAdmins []string `pulumi:"readOnlyAdmins"`
	// List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
	TrustedResourceOwners []string `pulumi:"trustedResourceOwners"`
}

// The set of arguments for constructing a DataLakeSettings resource.
type DataLakeSettingsArgs struct {
	// Set of ARNs of AWS Lake Formation principals (IAM users or roles).
	Admins pulumi.StringArrayInput
	// Whether to allow Amazon EMR clusters to access data managed by Lake Formation.
	AllowExternalDataFiltering pulumi.BoolPtrInput
	// Lake Formation relies on a privileged process secured by Amazon EMR or the third party integrator to tag the user's role while assuming it.
	//
	// > **NOTE:** Although optional, not including `admins`, `createDatabaseDefaultPermissions`, `createTableDefaultPermissions`, and/or `trustedResourceOwners` results in the setting being cleared.
	AuthorizedSessionTagValueLists pulumi.StringArrayInput
	// Identifier for the Data Catalog. By default, the account ID.
	CatalogId pulumi.StringPtrInput
	// Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
	CreateDatabaseDefaultPermissions DataLakeSettingsCreateDatabaseDefaultPermissionArrayInput
	// Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
	CreateTableDefaultPermissions DataLakeSettingsCreateTableDefaultPermissionArrayInput
	// A list of the account IDs of Amazon Web Services accounts with Amazon EMR clusters that are to perform data filtering.
	ExternalDataFilteringAllowLists pulumi.StringArrayInput
	// Set of ARNs of AWS Lake Formation principals (IAM users or roles) with only view access to the resources.
	ReadOnlyAdmins pulumi.StringArrayInput
	// List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
	TrustedResourceOwners pulumi.StringArrayInput
}

func (DataLakeSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataLakeSettingsArgs)(nil)).Elem()
}

type DataLakeSettingsInput interface {
	pulumi.Input

	ToDataLakeSettingsOutput() DataLakeSettingsOutput
	ToDataLakeSettingsOutputWithContext(ctx context.Context) DataLakeSettingsOutput
}

func (*DataLakeSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**DataLakeSettings)(nil)).Elem()
}

func (i *DataLakeSettings) ToDataLakeSettingsOutput() DataLakeSettingsOutput {
	return i.ToDataLakeSettingsOutputWithContext(context.Background())
}

func (i *DataLakeSettings) ToDataLakeSettingsOutputWithContext(ctx context.Context) DataLakeSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeSettingsOutput)
}

// DataLakeSettingsArrayInput is an input type that accepts DataLakeSettingsArray and DataLakeSettingsArrayOutput values.
// You can construct a concrete instance of `DataLakeSettingsArrayInput` via:
//
//	DataLakeSettingsArray{ DataLakeSettingsArgs{...} }
type DataLakeSettingsArrayInput interface {
	pulumi.Input

	ToDataLakeSettingsArrayOutput() DataLakeSettingsArrayOutput
	ToDataLakeSettingsArrayOutputWithContext(context.Context) DataLakeSettingsArrayOutput
}

type DataLakeSettingsArray []DataLakeSettingsInput

func (DataLakeSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataLakeSettings)(nil)).Elem()
}

func (i DataLakeSettingsArray) ToDataLakeSettingsArrayOutput() DataLakeSettingsArrayOutput {
	return i.ToDataLakeSettingsArrayOutputWithContext(context.Background())
}

func (i DataLakeSettingsArray) ToDataLakeSettingsArrayOutputWithContext(ctx context.Context) DataLakeSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeSettingsArrayOutput)
}

// DataLakeSettingsMapInput is an input type that accepts DataLakeSettingsMap and DataLakeSettingsMapOutput values.
// You can construct a concrete instance of `DataLakeSettingsMapInput` via:
//
//	DataLakeSettingsMap{ "key": DataLakeSettingsArgs{...} }
type DataLakeSettingsMapInput interface {
	pulumi.Input

	ToDataLakeSettingsMapOutput() DataLakeSettingsMapOutput
	ToDataLakeSettingsMapOutputWithContext(context.Context) DataLakeSettingsMapOutput
}

type DataLakeSettingsMap map[string]DataLakeSettingsInput

func (DataLakeSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataLakeSettings)(nil)).Elem()
}

func (i DataLakeSettingsMap) ToDataLakeSettingsMapOutput() DataLakeSettingsMapOutput {
	return i.ToDataLakeSettingsMapOutputWithContext(context.Background())
}

func (i DataLakeSettingsMap) ToDataLakeSettingsMapOutputWithContext(ctx context.Context) DataLakeSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataLakeSettingsMapOutput)
}

type DataLakeSettingsOutput struct{ *pulumi.OutputState }

func (DataLakeSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataLakeSettings)(nil)).Elem()
}

func (o DataLakeSettingsOutput) ToDataLakeSettingsOutput() DataLakeSettingsOutput {
	return o
}

func (o DataLakeSettingsOutput) ToDataLakeSettingsOutputWithContext(ctx context.Context) DataLakeSettingsOutput {
	return o
}

// Set of ARNs of AWS Lake Formation principals (IAM users or roles).
func (o DataLakeSettingsOutput) Admins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DataLakeSettings) pulumi.StringArrayOutput { return v.Admins }).(pulumi.StringArrayOutput)
}

// Whether to allow Amazon EMR clusters to access data managed by Lake Formation.
func (o DataLakeSettingsOutput) AllowExternalDataFiltering() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DataLakeSettings) pulumi.BoolPtrOutput { return v.AllowExternalDataFiltering }).(pulumi.BoolPtrOutput)
}

// Lake Formation relies on a privileged process secured by Amazon EMR or the third party integrator to tag the user's role while assuming it.
//
// > **NOTE:** Although optional, not including `admins`, `createDatabaseDefaultPermissions`, `createTableDefaultPermissions`, and/or `trustedResourceOwners` results in the setting being cleared.
func (o DataLakeSettingsOutput) AuthorizedSessionTagValueLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DataLakeSettings) pulumi.StringArrayOutput { return v.AuthorizedSessionTagValueLists }).(pulumi.StringArrayOutput)
}

// Identifier for the Data Catalog. By default, the account ID.
func (o DataLakeSettingsOutput) CatalogId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataLakeSettings) pulumi.StringPtrOutput { return v.CatalogId }).(pulumi.StringPtrOutput)
}

// Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
func (o DataLakeSettingsOutput) CreateDatabaseDefaultPermissions() DataLakeSettingsCreateDatabaseDefaultPermissionArrayOutput {
	return o.ApplyT(func(v *DataLakeSettings) DataLakeSettingsCreateDatabaseDefaultPermissionArrayOutput {
		return v.CreateDatabaseDefaultPermissions
	}).(DataLakeSettingsCreateDatabaseDefaultPermissionArrayOutput)
}

// Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
func (o DataLakeSettingsOutput) CreateTableDefaultPermissions() DataLakeSettingsCreateTableDefaultPermissionArrayOutput {
	return o.ApplyT(func(v *DataLakeSettings) DataLakeSettingsCreateTableDefaultPermissionArrayOutput {
		return v.CreateTableDefaultPermissions
	}).(DataLakeSettingsCreateTableDefaultPermissionArrayOutput)
}

// A list of the account IDs of Amazon Web Services accounts with Amazon EMR clusters that are to perform data filtering.
func (o DataLakeSettingsOutput) ExternalDataFilteringAllowLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DataLakeSettings) pulumi.StringArrayOutput { return v.ExternalDataFilteringAllowLists }).(pulumi.StringArrayOutput)
}

// Set of ARNs of AWS Lake Formation principals (IAM users or roles) with only view access to the resources.
func (o DataLakeSettingsOutput) ReadOnlyAdmins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DataLakeSettings) pulumi.StringArrayOutput { return v.ReadOnlyAdmins }).(pulumi.StringArrayOutput)
}

// List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
func (o DataLakeSettingsOutput) TrustedResourceOwners() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DataLakeSettings) pulumi.StringArrayOutput { return v.TrustedResourceOwners }).(pulumi.StringArrayOutput)
}

type DataLakeSettingsArrayOutput struct{ *pulumi.OutputState }

func (DataLakeSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataLakeSettings)(nil)).Elem()
}

func (o DataLakeSettingsArrayOutput) ToDataLakeSettingsArrayOutput() DataLakeSettingsArrayOutput {
	return o
}

func (o DataLakeSettingsArrayOutput) ToDataLakeSettingsArrayOutputWithContext(ctx context.Context) DataLakeSettingsArrayOutput {
	return o
}

func (o DataLakeSettingsArrayOutput) Index(i pulumi.IntInput) DataLakeSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DataLakeSettings {
		return vs[0].([]*DataLakeSettings)[vs[1].(int)]
	}).(DataLakeSettingsOutput)
}

type DataLakeSettingsMapOutput struct{ *pulumi.OutputState }

func (DataLakeSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataLakeSettings)(nil)).Elem()
}

func (o DataLakeSettingsMapOutput) ToDataLakeSettingsMapOutput() DataLakeSettingsMapOutput {
	return o
}

func (o DataLakeSettingsMapOutput) ToDataLakeSettingsMapOutputWithContext(ctx context.Context) DataLakeSettingsMapOutput {
	return o
}

func (o DataLakeSettingsMapOutput) MapIndex(k pulumi.StringInput) DataLakeSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DataLakeSettings {
		return vs[0].(map[string]*DataLakeSettings)[vs[1].(string)]
	}).(DataLakeSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataLakeSettingsInput)(nil)).Elem(), &DataLakeSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataLakeSettingsArrayInput)(nil)).Elem(), DataLakeSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataLakeSettingsMapInput)(nil)).Elem(), DataLakeSettingsMap{})
	pulumi.RegisterOutputType(DataLakeSettingsOutput{})
	pulumi.RegisterOutputType(DataLakeSettingsArrayOutput{})
	pulumi.RegisterOutputType(DataLakeSettingsMapOutput{})
}
