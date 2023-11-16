// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lakeformation

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Grants permissions to the principal to access metadata in the Data Catalog and data organized in underlying data storage such as Amazon S3. Permissions are granted to a principal, in a Data Catalog, relative to a Lake Formation resource, which includes the Data Catalog, databases, tables, LF-tags, and LF-tag policies. For more information, see [Security and Access Control to Metadata and Data in Lake Formation](https://docs.aws.amazon.com/lake-formation/latest/dg/security-data-access.html).
//
// !> **WARNING:** Lake Formation permissions are not in effect by default within AWS. Using this resource will not secure your data and will result in errors if you do not change the security settings for existing resources and the default security settings for new resources. See Default Behavior and `IAMAllowedPrincipals` for additional details.
//
// > **NOTE:** In general, the `principal` should _NOT_ be a Lake Formation administrator or the entity (e.g., IAM role) that is running the deployment. Administrators have implicit permissions. These should be managed by granting or not granting administrator rights using `lakeformation.DataLakeSettings`, _not_ with this resource.
//
// ## Default Behavior and `IAMAllowedPrincipals`
//
// **_Lake Formation permissions are not in effect by default within AWS._** `IAMAllowedPrincipals` (i.e., `IAM_ALLOWED_PRINCIPALS`) conflicts with individual Lake Formation permissions (i.e., non-`IAMAllowedPrincipals` permissions), will cause unexpected behavior, and may result in errors.
//
// When using Lake Formation, choose ONE of the following options as they are mutually exclusive:
//
// 1. Use this resource (`lakeformation.Permissions`), change the default security settings using `lakeformation.DataLakeSettings`, and remove existing `IAMAllowedPrincipals` permissions
// 2. Use `IAMAllowedPrincipals` without `lakeformation.Permissions`
//
// This example shows removing the `IAMAllowedPrincipals` default security settings and making the caller a Lake Formation admin. Since `createDatabaseDefaultPermissions` and `createTableDefaultPermissions` are not set in the `lakeformation.DataLakeSettings` resource, they are cleared.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lakeformation"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			currentSessionContext, err := iam.GetSessionContext(ctx, &iam.GetSessionContextArgs{
//				Arn: currentCallerIdentity.Arn,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = lakeformation.NewDataLakeSettings(ctx, "test", &lakeformation.DataLakeSettingsArgs{
//				Admins: pulumi.StringArray{
//					*pulumi.String(currentSessionContext.IssuerArn),
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
// To remove existing `IAMAllowedPrincipals` permissions, use the [AWS Lake Formation Console](https://console.aws.amazon.com/lakeformation/) or [AWS CLI](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/lakeformation/batch-revoke-permissions.html).
//
// `IAMAllowedPrincipals` is a hook to maintain backwards compatibility with AWS Glue. `IAMAllowedPrincipals` is a pseudo-entity group that acts like a Lake Formation principal. The group includes any IAM users and roles that are allowed access to your Data Catalog resources by your IAM policies.
//
// This is Lake Formation's default behavior:
//
// * Lake Formation grants `Super` permission to `IAMAllowedPrincipals` on all existing AWS Glue Data Catalog resources.
// * Lake Formation enables "Use only IAM access control" for new Data Catalog resources.
//
// For more details, see [Changing the Default Security Settings for Your Data Lake](https://docs.aws.amazon.com/lake-formation/latest/dg/change-settings.html).
//
// ### Problem Using `IAMAllowedPrincipals`
//
// AWS does not support combining `IAMAllowedPrincipals` permissions and non-`IAMAllowedPrincipals` permissions. Doing so results in unexpected permissions and behaviors. For example, this configuration grants a user `SELECT` on a column in a table.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lakeformation"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := glue.NewCatalogDatabase(ctx, "exampleCatalogDatabase", &glue.CatalogDatabaseArgs{
//				Name: pulumi.String("sadabate"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleCatalogTable, err := glue.NewCatalogTable(ctx, "exampleCatalogTable", &glue.CatalogTableArgs{
//				Name:         pulumi.String("abelt"),
//				DatabaseName: pulumi.Any(aws_glue_catalog_database.Test.Name),
//				StorageDescriptor: &glue.CatalogTableStorageDescriptorArgs{
//					Columns: glue.CatalogTableStorageDescriptorColumnArray{
//						&glue.CatalogTableStorageDescriptorColumnArgs{
//							Name: pulumi.String("event"),
//							Type: pulumi.String("string"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = lakeformation.NewPermissions(ctx, "examplePermissions", &lakeformation.PermissionsArgs{
//				Permissions: pulumi.StringArray{
//					pulumi.String("SELECT"),
//				},
//				Principal: pulumi.String("arn:aws:iam:us-east-1:123456789012:user/SanHolo"),
//				TableWithColumns: &lakeformation.PermissionsTableWithColumnsArgs{
//					DatabaseName: exampleCatalogTable.DatabaseName,
//					Name:         exampleCatalogTable.Name,
//					ColumnNames: pulumi.StringArray{
//						pulumi.String("event"),
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
//
// The resulting permissions depend on whether the table had `IAMAllowedPrincipals` (IAP) permissions or not.
//
// | Result With IAP | Result Without IAP |
// | ---- | ---- |
// | `SELECT` column wildcard (i.e., all columns) | `SELECT` on `"event"` (as expected) |
//
// ## Using Lake Formation Permissions
//
// Lake Formation grants implicit permissions to data lake administrators, database creators, and table creators. These implicit permissions cannot be revoked _per se_. If this resource reads implicit permissions, it will attempt to revoke them, which causes an error when the resource is destroyed.
//
// There are two ways to avoid these errors. First, and the way we recommend, is to avoid using this resource with principals that have implicit permissions. A second, error-prone option, is to grant explicit permissions (and `permissionsWithGrantOption`) to "overwrite" a principal's implicit permissions, which you can then revoke with this resource. For more information, see [Implicit Lake Formation Permissions](https://docs.aws.amazon.com/lake-formation/latest/dg/implicit-permissions.html).
//
// If the `principal` is also a data lake administrator, AWS grants implicit permissions that can cause errors using this resource. For example, AWS implicitly grants a `principal`/administrator `permissions` and `permissionsWithGrantOption` of `ALL`, `ALTER`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT` on a table. If you use this resource to explicitly grant the `principal`/administrator `permissions` but _not_ `permissionsWithGrantOption` of `ALL`, `ALTER`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT` on the table, this resource will read the implicit `permissionsWithGrantOption` and attempt to revoke them when the resource is destroyed. Doing so will cause an `InvalidInputException: No permissions revoked` error because you cannot revoke implicit permissions _per se_. To workaround this problem, explicitly grant the `principal`/administrator `permissions` _and_ `permissionsWithGrantOption`, which can then be revoked. Similarly, granting a `principal`/administrator permissions on a table with columns and providing `columnNames`, will result in a `InvalidInputException: Permissions modification is invalid` error because you are narrowing the implicit permissions. Instead, set `wildcard` to `true` and remove the `columnNames`.
//
// ## Example Usage
// ### Grant Permissions For A Lake Formation S3 Resource
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
//			_, err := lakeformation.NewPermissions(ctx, "example", &lakeformation.PermissionsArgs{
//				Principal: pulumi.Any(aws_iam_role.Workflow_role.Arn),
//				Permissions: pulumi.StringArray{
//					pulumi.String("ALL"),
//				},
//				DataLocation: &lakeformation.PermissionsDataLocationArgs{
//					Arn: pulumi.Any(aws_lakeformation_resource.Example.Arn),
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
// ### Grant Permissions For A Glue Catalog Database
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
//			_, err := lakeformation.NewPermissions(ctx, "example", &lakeformation.PermissionsArgs{
//				Principal: pulumi.Any(aws_iam_role.Workflow_role.Arn),
//				Permissions: pulumi.StringArray{
//					pulumi.String("CREATE_TABLE"),
//					pulumi.String("ALTER"),
//					pulumi.String("DROP"),
//				},
//				Database: &lakeformation.PermissionsDatabaseArgs{
//					Name:      pulumi.Any(aws_glue_catalog_database.Example.Name),
//					CatalogId: pulumi.String("110376042874"),
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
// ### Grant Permissions Using Tag-Based Access Control
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
//			_, err := lakeformation.NewPermissions(ctx, "test", &lakeformation.PermissionsArgs{
//				Principal: pulumi.Any(aws_iam_role.Sales_role.Arn),
//				Permissions: pulumi.StringArray{
//					pulumi.String("CREATE_TABLE"),
//					pulumi.String("ALTER"),
//					pulumi.String("DROP"),
//				},
//				LfTagPolicy: &lakeformation.PermissionsLfTagPolicyArgs{
//					ResourceType: pulumi.String("DATABASE"),
//					Expressions: lakeformation.PermissionsLfTagPolicyExpressionArray{
//						&lakeformation.PermissionsLfTagPolicyExpressionArgs{
//							Key: pulumi.String("Team"),
//							Values: pulumi.StringArray{
//								pulumi.String("Sales"),
//							},
//						},
//						&lakeformation.PermissionsLfTagPolicyExpressionArgs{
//							Key: pulumi.String("Environment"),
//							Values: pulumi.StringArray{
//								pulumi.String("Dev"),
//								pulumi.String("Production"),
//							},
//						},
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
type Permissions struct {
	pulumi.CustomResourceState

	// Identifier for the Data Catalog. By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.
	CatalogId pulumi.StringPtrOutput `pulumi:"catalogId"`
	// Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
	CatalogResource pulumi.BoolPtrOutput `pulumi:"catalogResource"`
	// Configuration block for a data location resource. Detailed below.
	DataLocation PermissionsDataLocationOutput `pulumi:"dataLocation"`
	// Configuration block for a database resource. Detailed below.
	Database PermissionsDatabaseOutput `pulumi:"database"`
	// Configuration block for an LF-tag resource. Detailed below.
	LfTag PermissionsLfTagOutput `pulumi:"lfTag"`
	// Configuration block for an LF-tag policy resource. Detailed below.
	LfTagPolicy PermissionsLfTagPolicyOutput `pulumi:"lfTagPolicy"`
	// List of permissions granted to the principal. Valid values may include `ALL`, `ALTER`, `ASSOCIATE`, `CREATE_DATABASE`, `CREATE_TABLE`, `DATA_LOCATION_ACCESS`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT`. For details on each permission, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
	Permissions pulumi.StringArrayOutput `pulumi:"permissions"`
	// Subset of `permissions` which the principal can pass.
	PermissionsWithGrantOptions pulumi.StringArrayOutput `pulumi:"permissionsWithGrantOptions"`
	// Principal to be granted the permissions on the resource. Supported principals include `IAM_ALLOWED_PRINCIPALS` (see Default Behavior and `IAMAllowedPrincipals` above), IAM roles, users, groups, SAML groups and users, QuickSight groups, OUs, and organizations as well as AWS account IDs for cross-account permissions. For more information, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
	//
	// > **NOTE:** We highly recommend that the `principal` _NOT_ be a Lake Formation administrator (granted using `lakeformation.DataLakeSettings`). The entity (e.g., IAM role) running the deployment will most likely need to be a Lake Formation administrator. As such, the entity will have implicit permissions and does not need permissions granted through this resource.
	//
	// One of the following is required:
	Principal pulumi.StringOutput `pulumi:"principal"`
	// Configuration block for a table resource. Detailed below.
	Table PermissionsTableOutput `pulumi:"table"`
	// Configuration block for a table with columns resource. Detailed below.
	//
	// The following arguments are optional:
	TableWithColumns PermissionsTableWithColumnsOutput `pulumi:"tableWithColumns"`
}

// NewPermissions registers a new resource with the given unique name, arguments, and options.
func NewPermissions(ctx *pulumi.Context,
	name string, args *PermissionsArgs, opts ...pulumi.ResourceOption) (*Permissions, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Permissions == nil {
		return nil, errors.New("invalid value for required argument 'Permissions'")
	}
	if args.Principal == nil {
		return nil, errors.New("invalid value for required argument 'Principal'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Permissions
	err := ctx.RegisterResource("aws:lakeformation/permissions:Permissions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPermissions gets an existing Permissions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPermissions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PermissionsState, opts ...pulumi.ResourceOption) (*Permissions, error) {
	var resource Permissions
	err := ctx.ReadResource("aws:lakeformation/permissions:Permissions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Permissions resources.
type permissionsState struct {
	// Identifier for the Data Catalog. By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.
	CatalogId *string `pulumi:"catalogId"`
	// Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
	CatalogResource *bool `pulumi:"catalogResource"`
	// Configuration block for a data location resource. Detailed below.
	DataLocation *PermissionsDataLocation `pulumi:"dataLocation"`
	// Configuration block for a database resource. Detailed below.
	Database *PermissionsDatabase `pulumi:"database"`
	// Configuration block for an LF-tag resource. Detailed below.
	LfTag *PermissionsLfTag `pulumi:"lfTag"`
	// Configuration block for an LF-tag policy resource. Detailed below.
	LfTagPolicy *PermissionsLfTagPolicy `pulumi:"lfTagPolicy"`
	// List of permissions granted to the principal. Valid values may include `ALL`, `ALTER`, `ASSOCIATE`, `CREATE_DATABASE`, `CREATE_TABLE`, `DATA_LOCATION_ACCESS`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT`. For details on each permission, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
	Permissions []string `pulumi:"permissions"`
	// Subset of `permissions` which the principal can pass.
	PermissionsWithGrantOptions []string `pulumi:"permissionsWithGrantOptions"`
	// Principal to be granted the permissions on the resource. Supported principals include `IAM_ALLOWED_PRINCIPALS` (see Default Behavior and `IAMAllowedPrincipals` above), IAM roles, users, groups, SAML groups and users, QuickSight groups, OUs, and organizations as well as AWS account IDs for cross-account permissions. For more information, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
	//
	// > **NOTE:** We highly recommend that the `principal` _NOT_ be a Lake Formation administrator (granted using `lakeformation.DataLakeSettings`). The entity (e.g., IAM role) running the deployment will most likely need to be a Lake Formation administrator. As such, the entity will have implicit permissions and does not need permissions granted through this resource.
	//
	// One of the following is required:
	Principal *string `pulumi:"principal"`
	// Configuration block for a table resource. Detailed below.
	Table *PermissionsTable `pulumi:"table"`
	// Configuration block for a table with columns resource. Detailed below.
	//
	// The following arguments are optional:
	TableWithColumns *PermissionsTableWithColumns `pulumi:"tableWithColumns"`
}

type PermissionsState struct {
	// Identifier for the Data Catalog. By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.
	CatalogId pulumi.StringPtrInput
	// Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
	CatalogResource pulumi.BoolPtrInput
	// Configuration block for a data location resource. Detailed below.
	DataLocation PermissionsDataLocationPtrInput
	// Configuration block for a database resource. Detailed below.
	Database PermissionsDatabasePtrInput
	// Configuration block for an LF-tag resource. Detailed below.
	LfTag PermissionsLfTagPtrInput
	// Configuration block for an LF-tag policy resource. Detailed below.
	LfTagPolicy PermissionsLfTagPolicyPtrInput
	// List of permissions granted to the principal. Valid values may include `ALL`, `ALTER`, `ASSOCIATE`, `CREATE_DATABASE`, `CREATE_TABLE`, `DATA_LOCATION_ACCESS`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT`. For details on each permission, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
	Permissions pulumi.StringArrayInput
	// Subset of `permissions` which the principal can pass.
	PermissionsWithGrantOptions pulumi.StringArrayInput
	// Principal to be granted the permissions on the resource. Supported principals include `IAM_ALLOWED_PRINCIPALS` (see Default Behavior and `IAMAllowedPrincipals` above), IAM roles, users, groups, SAML groups and users, QuickSight groups, OUs, and organizations as well as AWS account IDs for cross-account permissions. For more information, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
	//
	// > **NOTE:** We highly recommend that the `principal` _NOT_ be a Lake Formation administrator (granted using `lakeformation.DataLakeSettings`). The entity (e.g., IAM role) running the deployment will most likely need to be a Lake Formation administrator. As such, the entity will have implicit permissions and does not need permissions granted through this resource.
	//
	// One of the following is required:
	Principal pulumi.StringPtrInput
	// Configuration block for a table resource. Detailed below.
	Table PermissionsTablePtrInput
	// Configuration block for a table with columns resource. Detailed below.
	//
	// The following arguments are optional:
	TableWithColumns PermissionsTableWithColumnsPtrInput
}

func (PermissionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionsState)(nil)).Elem()
}

type permissionsArgs struct {
	// Identifier for the Data Catalog. By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.
	CatalogId *string `pulumi:"catalogId"`
	// Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
	CatalogResource *bool `pulumi:"catalogResource"`
	// Configuration block for a data location resource. Detailed below.
	DataLocation *PermissionsDataLocation `pulumi:"dataLocation"`
	// Configuration block for a database resource. Detailed below.
	Database *PermissionsDatabase `pulumi:"database"`
	// Configuration block for an LF-tag resource. Detailed below.
	LfTag *PermissionsLfTag `pulumi:"lfTag"`
	// Configuration block for an LF-tag policy resource. Detailed below.
	LfTagPolicy *PermissionsLfTagPolicy `pulumi:"lfTagPolicy"`
	// List of permissions granted to the principal. Valid values may include `ALL`, `ALTER`, `ASSOCIATE`, `CREATE_DATABASE`, `CREATE_TABLE`, `DATA_LOCATION_ACCESS`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT`. For details on each permission, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
	Permissions []string `pulumi:"permissions"`
	// Subset of `permissions` which the principal can pass.
	PermissionsWithGrantOptions []string `pulumi:"permissionsWithGrantOptions"`
	// Principal to be granted the permissions on the resource. Supported principals include `IAM_ALLOWED_PRINCIPALS` (see Default Behavior and `IAMAllowedPrincipals` above), IAM roles, users, groups, SAML groups and users, QuickSight groups, OUs, and organizations as well as AWS account IDs for cross-account permissions. For more information, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
	//
	// > **NOTE:** We highly recommend that the `principal` _NOT_ be a Lake Formation administrator (granted using `lakeformation.DataLakeSettings`). The entity (e.g., IAM role) running the deployment will most likely need to be a Lake Formation administrator. As such, the entity will have implicit permissions and does not need permissions granted through this resource.
	//
	// One of the following is required:
	Principal string `pulumi:"principal"`
	// Configuration block for a table resource. Detailed below.
	Table *PermissionsTable `pulumi:"table"`
	// Configuration block for a table with columns resource. Detailed below.
	//
	// The following arguments are optional:
	TableWithColumns *PermissionsTableWithColumns `pulumi:"tableWithColumns"`
}

// The set of arguments for constructing a Permissions resource.
type PermissionsArgs struct {
	// Identifier for the Data Catalog. By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.
	CatalogId pulumi.StringPtrInput
	// Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
	CatalogResource pulumi.BoolPtrInput
	// Configuration block for a data location resource. Detailed below.
	DataLocation PermissionsDataLocationPtrInput
	// Configuration block for a database resource. Detailed below.
	Database PermissionsDatabasePtrInput
	// Configuration block for an LF-tag resource. Detailed below.
	LfTag PermissionsLfTagPtrInput
	// Configuration block for an LF-tag policy resource. Detailed below.
	LfTagPolicy PermissionsLfTagPolicyPtrInput
	// List of permissions granted to the principal. Valid values may include `ALL`, `ALTER`, `ASSOCIATE`, `CREATE_DATABASE`, `CREATE_TABLE`, `DATA_LOCATION_ACCESS`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT`. For details on each permission, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
	Permissions pulumi.StringArrayInput
	// Subset of `permissions` which the principal can pass.
	PermissionsWithGrantOptions pulumi.StringArrayInput
	// Principal to be granted the permissions on the resource. Supported principals include `IAM_ALLOWED_PRINCIPALS` (see Default Behavior and `IAMAllowedPrincipals` above), IAM roles, users, groups, SAML groups and users, QuickSight groups, OUs, and organizations as well as AWS account IDs for cross-account permissions. For more information, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
	//
	// > **NOTE:** We highly recommend that the `principal` _NOT_ be a Lake Formation administrator (granted using `lakeformation.DataLakeSettings`). The entity (e.g., IAM role) running the deployment will most likely need to be a Lake Formation administrator. As such, the entity will have implicit permissions and does not need permissions granted through this resource.
	//
	// One of the following is required:
	Principal pulumi.StringInput
	// Configuration block for a table resource. Detailed below.
	Table PermissionsTablePtrInput
	// Configuration block for a table with columns resource. Detailed below.
	//
	// The following arguments are optional:
	TableWithColumns PermissionsTableWithColumnsPtrInput
}

func (PermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionsArgs)(nil)).Elem()
}

type PermissionsInput interface {
	pulumi.Input

	ToPermissionsOutput() PermissionsOutput
	ToPermissionsOutputWithContext(ctx context.Context) PermissionsOutput
}

func (*Permissions) ElementType() reflect.Type {
	return reflect.TypeOf((**Permissions)(nil)).Elem()
}

func (i *Permissions) ToPermissionsOutput() PermissionsOutput {
	return i.ToPermissionsOutputWithContext(context.Background())
}

func (i *Permissions) ToPermissionsOutputWithContext(ctx context.Context) PermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsOutput)
}

// PermissionsArrayInput is an input type that accepts PermissionsArray and PermissionsArrayOutput values.
// You can construct a concrete instance of `PermissionsArrayInput` via:
//
//	PermissionsArray{ PermissionsArgs{...} }
type PermissionsArrayInput interface {
	pulumi.Input

	ToPermissionsArrayOutput() PermissionsArrayOutput
	ToPermissionsArrayOutputWithContext(context.Context) PermissionsArrayOutput
}

type PermissionsArray []PermissionsInput

func (PermissionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Permissions)(nil)).Elem()
}

func (i PermissionsArray) ToPermissionsArrayOutput() PermissionsArrayOutput {
	return i.ToPermissionsArrayOutputWithContext(context.Background())
}

func (i PermissionsArray) ToPermissionsArrayOutputWithContext(ctx context.Context) PermissionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsArrayOutput)
}

// PermissionsMapInput is an input type that accepts PermissionsMap and PermissionsMapOutput values.
// You can construct a concrete instance of `PermissionsMapInput` via:
//
//	PermissionsMap{ "key": PermissionsArgs{...} }
type PermissionsMapInput interface {
	pulumi.Input

	ToPermissionsMapOutput() PermissionsMapOutput
	ToPermissionsMapOutputWithContext(context.Context) PermissionsMapOutput
}

type PermissionsMap map[string]PermissionsInput

func (PermissionsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Permissions)(nil)).Elem()
}

func (i PermissionsMap) ToPermissionsMapOutput() PermissionsMapOutput {
	return i.ToPermissionsMapOutputWithContext(context.Background())
}

func (i PermissionsMap) ToPermissionsMapOutputWithContext(ctx context.Context) PermissionsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsMapOutput)
}

type PermissionsOutput struct{ *pulumi.OutputState }

func (PermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Permissions)(nil)).Elem()
}

func (o PermissionsOutput) ToPermissionsOutput() PermissionsOutput {
	return o
}

func (o PermissionsOutput) ToPermissionsOutputWithContext(ctx context.Context) PermissionsOutput {
	return o
}

// Identifier for the Data Catalog. By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.
func (o PermissionsOutput) CatalogId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Permissions) pulumi.StringPtrOutput { return v.CatalogId }).(pulumi.StringPtrOutput)
}

// Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
func (o PermissionsOutput) CatalogResource() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Permissions) pulumi.BoolPtrOutput { return v.CatalogResource }).(pulumi.BoolPtrOutput)
}

// Configuration block for a data location resource. Detailed below.
func (o PermissionsOutput) DataLocation() PermissionsDataLocationOutput {
	return o.ApplyT(func(v *Permissions) PermissionsDataLocationOutput { return v.DataLocation }).(PermissionsDataLocationOutput)
}

// Configuration block for a database resource. Detailed below.
func (o PermissionsOutput) Database() PermissionsDatabaseOutput {
	return o.ApplyT(func(v *Permissions) PermissionsDatabaseOutput { return v.Database }).(PermissionsDatabaseOutput)
}

// Configuration block for an LF-tag resource. Detailed below.
func (o PermissionsOutput) LfTag() PermissionsLfTagOutput {
	return o.ApplyT(func(v *Permissions) PermissionsLfTagOutput { return v.LfTag }).(PermissionsLfTagOutput)
}

// Configuration block for an LF-tag policy resource. Detailed below.
func (o PermissionsOutput) LfTagPolicy() PermissionsLfTagPolicyOutput {
	return o.ApplyT(func(v *Permissions) PermissionsLfTagPolicyOutput { return v.LfTagPolicy }).(PermissionsLfTagPolicyOutput)
}

// List of permissions granted to the principal. Valid values may include `ALL`, `ALTER`, `ASSOCIATE`, `CREATE_DATABASE`, `CREATE_TABLE`, `DATA_LOCATION_ACCESS`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT`. For details on each permission, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
func (o PermissionsOutput) Permissions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Permissions) pulumi.StringArrayOutput { return v.Permissions }).(pulumi.StringArrayOutput)
}

// Subset of `permissions` which the principal can pass.
func (o PermissionsOutput) PermissionsWithGrantOptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Permissions) pulumi.StringArrayOutput { return v.PermissionsWithGrantOptions }).(pulumi.StringArrayOutput)
}

// Principal to be granted the permissions on the resource. Supported principals include `IAM_ALLOWED_PRINCIPALS` (see Default Behavior and `IAMAllowedPrincipals` above), IAM roles, users, groups, SAML groups and users, QuickSight groups, OUs, and organizations as well as AWS account IDs for cross-account permissions. For more information, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
//
// > **NOTE:** We highly recommend that the `principal` _NOT_ be a Lake Formation administrator (granted using `lakeformation.DataLakeSettings`). The entity (e.g., IAM role) running the deployment will most likely need to be a Lake Formation administrator. As such, the entity will have implicit permissions and does not need permissions granted through this resource.
//
// One of the following is required:
func (o PermissionsOutput) Principal() pulumi.StringOutput {
	return o.ApplyT(func(v *Permissions) pulumi.StringOutput { return v.Principal }).(pulumi.StringOutput)
}

// Configuration block for a table resource. Detailed below.
func (o PermissionsOutput) Table() PermissionsTableOutput {
	return o.ApplyT(func(v *Permissions) PermissionsTableOutput { return v.Table }).(PermissionsTableOutput)
}

// Configuration block for a table with columns resource. Detailed below.
//
// The following arguments are optional:
func (o PermissionsOutput) TableWithColumns() PermissionsTableWithColumnsOutput {
	return o.ApplyT(func(v *Permissions) PermissionsTableWithColumnsOutput { return v.TableWithColumns }).(PermissionsTableWithColumnsOutput)
}

type PermissionsArrayOutput struct{ *pulumi.OutputState }

func (PermissionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Permissions)(nil)).Elem()
}

func (o PermissionsArrayOutput) ToPermissionsArrayOutput() PermissionsArrayOutput {
	return o
}

func (o PermissionsArrayOutput) ToPermissionsArrayOutputWithContext(ctx context.Context) PermissionsArrayOutput {
	return o
}

func (o PermissionsArrayOutput) Index(i pulumi.IntInput) PermissionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Permissions {
		return vs[0].([]*Permissions)[vs[1].(int)]
	}).(PermissionsOutput)
}

type PermissionsMapOutput struct{ *pulumi.OutputState }

func (PermissionsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Permissions)(nil)).Elem()
}

func (o PermissionsMapOutput) ToPermissionsMapOutput() PermissionsMapOutput {
	return o
}

func (o PermissionsMapOutput) ToPermissionsMapOutputWithContext(ctx context.Context) PermissionsMapOutput {
	return o
}

func (o PermissionsMapOutput) MapIndex(k pulumi.StringInput) PermissionsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Permissions {
		return vs[0].(map[string]*Permissions)[vs[1].(string)]
	}).(PermissionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionsInput)(nil)).Elem(), &Permissions{})
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionsArrayInput)(nil)).Elem(), PermissionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionsMapInput)(nil)).Elem(), PermissionsMap{})
	pulumi.RegisterOutputType(PermissionsOutput{})
	pulumi.RegisterOutputType(PermissionsArrayOutput{})
	pulumi.RegisterOutputType(PermissionsMapOutput{})
}
