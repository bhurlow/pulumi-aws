// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages an RDS Global Cluster, which is an Aurora global database spread across multiple regions. The global database contains a single primary cluster with read-write capability, and a read-only secondary cluster that receives data from the primary cluster through high-speed replication performed by the Aurora storage subsystem.
//
// More information about Aurora global databases can be found in the [Aurora User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html#aurora-global-database-creating).
//
// ## Example Usage
// ### New MySQL Global Cluster
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := rds.NewGlobalCluster(ctx, "example", &rds.GlobalClusterArgs{
//				GlobalClusterIdentifier: pulumi.String("global-test"),
//				Engine:                  pulumi.String("aurora"),
//				EngineVersion:           pulumi.String("5.6.mysql_aurora.1.22.2"),
//				DatabaseName:            pulumi.String("example_db"),
//			})
//			if err != nil {
//				return err
//			}
//			primaryCluster, err := rds.NewCluster(ctx, "primaryCluster", &rds.ClusterArgs{
//				Engine:                  example.Engine,
//				EngineVersion:           example.EngineVersion,
//				ClusterIdentifier:       pulumi.String("test-primary-cluster"),
//				MasterUsername:          pulumi.String("username"),
//				MasterPassword:          pulumi.String("somepass123"),
//				DatabaseName:            pulumi.String("example_db"),
//				GlobalClusterIdentifier: example.ID(),
//				DbSubnetGroupName:       pulumi.String("default"),
//			}, pulumi.Provider(aws.Primary))
//			if err != nil {
//				return err
//			}
//			primaryClusterInstance, err := rds.NewClusterInstance(ctx, "primaryClusterInstance", &rds.ClusterInstanceArgs{
//				Engine:            example.Engine,
//				EngineVersion:     example.EngineVersion,
//				Identifier:        pulumi.String("test-primary-cluster-instance"),
//				ClusterIdentifier: primaryCluster.ID(),
//				InstanceClass:     pulumi.String("db.r4.large"),
//				DbSubnetGroupName: pulumi.String("default"),
//			}, pulumi.Provider(aws.Primary))
//			if err != nil {
//				return err
//			}
//			secondaryCluster, err := rds.NewCluster(ctx, "secondaryCluster", &rds.ClusterArgs{
//				Engine:                  example.Engine,
//				EngineVersion:           example.EngineVersion,
//				ClusterIdentifier:       pulumi.String("test-secondary-cluster"),
//				GlobalClusterIdentifier: example.ID(),
//				DbSubnetGroupName:       pulumi.String("default"),
//			}, pulumi.Provider(aws.Secondary), pulumi.DependsOn([]pulumi.Resource{
//				primaryClusterInstance,
//			}))
//			if err != nil {
//				return err
//			}
//			_, err = rds.NewClusterInstance(ctx, "secondaryClusterInstance", &rds.ClusterInstanceArgs{
//				Engine:            example.Engine,
//				EngineVersion:     example.EngineVersion,
//				Identifier:        pulumi.String("test-secondary-cluster-instance"),
//				ClusterIdentifier: secondaryCluster.ID(),
//				InstanceClass:     pulumi.String("db.r4.large"),
//				DbSubnetGroupName: pulumi.String("default"),
//			}, pulumi.Provider(aws.Secondary))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### New PostgreSQL Global Cluster
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := aws.NewProvider(ctx, "primary", &aws.ProviderArgs{
//				Region: pulumi.String("us-east-2"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = aws.NewProvider(ctx, "secondary", &aws.ProviderArgs{
//				Region: pulumi.String("us-east-1"),
//			})
//			if err != nil {
//				return err
//			}
//			example, err := rds.NewGlobalCluster(ctx, "example", &rds.GlobalClusterArgs{
//				GlobalClusterIdentifier: pulumi.String("global-test"),
//				Engine:                  pulumi.String("aurora-postgresql"),
//				EngineVersion:           pulumi.String("11.9"),
//				DatabaseName:            pulumi.String("example_db"),
//			})
//			if err != nil {
//				return err
//			}
//			primaryCluster, err := rds.NewCluster(ctx, "primaryCluster", &rds.ClusterArgs{
//				Engine:                  example.Engine,
//				EngineVersion:           example.EngineVersion,
//				ClusterIdentifier:       pulumi.String("test-primary-cluster"),
//				MasterUsername:          pulumi.String("username"),
//				MasterPassword:          pulumi.String("somepass123"),
//				DatabaseName:            pulumi.String("example_db"),
//				GlobalClusterIdentifier: example.ID(),
//				DbSubnetGroupName:       pulumi.String("default"),
//			}, pulumi.Provider(aws.Primary))
//			if err != nil {
//				return err
//			}
//			primaryClusterInstance, err := rds.NewClusterInstance(ctx, "primaryClusterInstance", &rds.ClusterInstanceArgs{
//				Engine:            example.Engine,
//				EngineVersion:     example.EngineVersion,
//				Identifier:        pulumi.String("test-primary-cluster-instance"),
//				ClusterIdentifier: primaryCluster.ID(),
//				InstanceClass:     pulumi.String("db.r4.large"),
//				DbSubnetGroupName: pulumi.String("default"),
//			}, pulumi.Provider(aws.Primary))
//			if err != nil {
//				return err
//			}
//			secondaryCluster, err := rds.NewCluster(ctx, "secondaryCluster", &rds.ClusterArgs{
//				Engine:                  example.Engine,
//				EngineVersion:           example.EngineVersion,
//				ClusterIdentifier:       pulumi.String("test-secondary-cluster"),
//				GlobalClusterIdentifier: example.ID(),
//				SkipFinalSnapshot:       pulumi.Bool(true),
//				DbSubnetGroupName:       pulumi.String("default"),
//			}, pulumi.Provider(aws.Secondary), pulumi.DependsOn([]pulumi.Resource{
//				primaryClusterInstance,
//			}))
//			if err != nil {
//				return err
//			}
//			_, err = rds.NewClusterInstance(ctx, "secondaryClusterInstance", &rds.ClusterInstanceArgs{
//				Engine:            example.Engine,
//				EngineVersion:     example.EngineVersion,
//				Identifier:        pulumi.String("test-secondary-cluster-instance"),
//				ClusterIdentifier: secondaryCluster.ID(),
//				InstanceClass:     pulumi.String("db.r4.large"),
//				DbSubnetGroupName: pulumi.String("default"),
//			}, pulumi.Provider(aws.Secondary))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### New Global Cluster From Existing DB Cluster
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleCluster, err := rds.NewCluster(ctx, "exampleCluster", nil)
//			if err != nil {
//				return err
//			}
//			_, err = rds.NewGlobalCluster(ctx, "exampleGlobalCluster", &rds.GlobalClusterArgs{
//				ForceDestroy:              pulumi.Bool(true),
//				GlobalClusterIdentifier:   pulumi.String("example"),
//				SourceDbClusterIdentifier: exampleCluster.Arn,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Upgrading Engine Versions
//
// When you upgrade the version of an `rds.GlobalCluster`, the provider will attempt to in-place upgrade the engine versions of all associated clusters. Since the `rds.Cluster` resource is being updated through the `rds.GlobalCluster`, you are likely to get an error (`Provider produced inconsistent final plan`). To avoid this, use the `lifecycle` `ignoreChanges` meta argument as shown below on the `rds.Cluster`.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := rds.NewGlobalCluster(ctx, "example", &rds.GlobalClusterArgs{
//				GlobalClusterIdentifier: pulumi.String("kyivkharkiv"),
//				Engine:                  pulumi.String("aurora-mysql"),
//				EngineVersion:           pulumi.String("5.7.mysql_aurora.2.07.5"),
//			})
//			if err != nil {
//				return err
//			}
//			primaryCluster, err := rds.NewCluster(ctx, "primaryCluster", &rds.ClusterArgs{
//				AllowMajorVersionUpgrade: pulumi.Bool(true),
//				ApplyImmediately:         pulumi.Bool(true),
//				ClusterIdentifier:        pulumi.String("odessadnipro"),
//				DatabaseName:             pulumi.String("totoro"),
//				Engine:                   example.Engine,
//				EngineVersion:            example.EngineVersion,
//				GlobalClusterIdentifier:  example.ID(),
//				MasterPassword:           pulumi.String("satsukimae"),
//				MasterUsername:           pulumi.String("maesatsuki"),
//				SkipFinalSnapshot:        pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = rds.NewClusterInstance(ctx, "primaryClusterInstance", &rds.ClusterInstanceArgs{
//				ApplyImmediately:  pulumi.Bool(true),
//				ClusterIdentifier: primaryCluster.ID(),
//				Engine:            primaryCluster.Engine,
//				EngineVersion:     primaryCluster.EngineVersion,
//				Identifier:        pulumi.String("donetsklviv"),
//				InstanceClass:     pulumi.String("db.r4.large"),
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
// Using `pulumi import`, import `aws_rds_global_cluster` using the RDS Global Cluster identifier. For example:
//
// ```sh
//
//	$ pulumi import aws:rds/globalCluster:GlobalCluster example example
//
// ```
//
//	Certain resource arguments, like `force_destroy`, only exist within this provider. If the argument is set in the the provider configuration on an imported resource, This provider will show a difference on the first plan after import to update the state value. This change is safe to apply immediately so the state matches the desired configuration.
//
// Certain resource arguments, like `source_db_cluster_identifier`, do not have an API method for reading the information after creation. If the argument is set in the Pulumi program on an imported resource, Pulumi will always show a difference. To workaround this behavior, either omit the argument from the Pulumi program or use `ignore_changes` to hide the difference. For example:
type GlobalCluster struct {
	pulumi.CustomResourceState

	// RDS Global Cluster Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Name for an automatically created database on cluster creation.
	DatabaseName pulumi.StringPtrOutput `pulumi:"databaseName"`
	// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
	DeletionProtection pulumi.BoolPtrOutput `pulumi:"deletionProtection"`
	// Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Valid values: `aurora`, `aurora-mysql`, `aurora-postgresql`. Defaults to `aurora`. Conflicts with `sourceDbClusterIdentifier`.
	Engine pulumi.StringOutput `pulumi:"engine"`
	// Engine version of the Aurora global database. The `engine`, `engineVersion`, and `instanceClass` (on the `rds.ClusterInstance`) must together support global databases. See [Using Amazon Aurora global databases](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html) for more information. By upgrading the engine version, the provider will upgrade cluster members. **NOTE:** To avoid an `inconsistent final plan` error while upgrading, use the `lifecycle` `ignoreChanges` for `engineVersion` meta argument on the associated `rds.Cluster` resource as shown above in Upgrading Engine Versions example.
	EngineVersion       pulumi.StringOutput `pulumi:"engineVersion"`
	EngineVersionActual pulumi.StringOutput `pulumi:"engineVersionActual"`
	// Enable to remove DB Cluster members from Global Cluster on destroy. Required with `sourceDbClusterIdentifier`.
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// Global cluster identifier.
	GlobalClusterIdentifier pulumi.StringOutput `pulumi:"globalClusterIdentifier"`
	// Set of objects containing Global Cluster members.
	GlobalClusterMembers GlobalClusterGlobalClusterMemberArrayOutput `pulumi:"globalClusterMembers"`
	// AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed
	GlobalClusterResourceId pulumi.StringOutput `pulumi:"globalClusterResourceId"`
	// Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
	SourceDbClusterIdentifier pulumi.StringOutput `pulumi:"sourceDbClusterIdentifier"`
	// Specifies whether the DB cluster is encrypted. The default is `false` unless `sourceDbClusterIdentifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
	StorageEncrypted pulumi.BoolOutput `pulumi:"storageEncrypted"`
}

// NewGlobalCluster registers a new resource with the given unique name, arguments, and options.
func NewGlobalCluster(ctx *pulumi.Context,
	name string, args *GlobalClusterArgs, opts ...pulumi.ResourceOption) (*GlobalCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GlobalClusterIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'GlobalClusterIdentifier'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GlobalCluster
	err := ctx.RegisterResource("aws:rds/globalCluster:GlobalCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobalCluster gets an existing GlobalCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalClusterState, opts ...pulumi.ResourceOption) (*GlobalCluster, error) {
	var resource GlobalCluster
	err := ctx.ReadResource("aws:rds/globalCluster:GlobalCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GlobalCluster resources.
type globalClusterState struct {
	// RDS Global Cluster Amazon Resource Name (ARN)
	Arn *string `pulumi:"arn"`
	// Name for an automatically created database on cluster creation.
	DatabaseName *string `pulumi:"databaseName"`
	// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Valid values: `aurora`, `aurora-mysql`, `aurora-postgresql`. Defaults to `aurora`. Conflicts with `sourceDbClusterIdentifier`.
	Engine *string `pulumi:"engine"`
	// Engine version of the Aurora global database. The `engine`, `engineVersion`, and `instanceClass` (on the `rds.ClusterInstance`) must together support global databases. See [Using Amazon Aurora global databases](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html) for more information. By upgrading the engine version, the provider will upgrade cluster members. **NOTE:** To avoid an `inconsistent final plan` error while upgrading, use the `lifecycle` `ignoreChanges` for `engineVersion` meta argument on the associated `rds.Cluster` resource as shown above in Upgrading Engine Versions example.
	EngineVersion       *string `pulumi:"engineVersion"`
	EngineVersionActual *string `pulumi:"engineVersionActual"`
	// Enable to remove DB Cluster members from Global Cluster on destroy. Required with `sourceDbClusterIdentifier`.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// Global cluster identifier.
	GlobalClusterIdentifier *string `pulumi:"globalClusterIdentifier"`
	// Set of objects containing Global Cluster members.
	GlobalClusterMembers []GlobalClusterGlobalClusterMember `pulumi:"globalClusterMembers"`
	// AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed
	GlobalClusterResourceId *string `pulumi:"globalClusterResourceId"`
	// Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
	SourceDbClusterIdentifier *string `pulumi:"sourceDbClusterIdentifier"`
	// Specifies whether the DB cluster is encrypted. The default is `false` unless `sourceDbClusterIdentifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
	StorageEncrypted *bool `pulumi:"storageEncrypted"`
}

type GlobalClusterState struct {
	// RDS Global Cluster Amazon Resource Name (ARN)
	Arn pulumi.StringPtrInput
	// Name for an automatically created database on cluster creation.
	DatabaseName pulumi.StringPtrInput
	// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
	DeletionProtection pulumi.BoolPtrInput
	// Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Valid values: `aurora`, `aurora-mysql`, `aurora-postgresql`. Defaults to `aurora`. Conflicts with `sourceDbClusterIdentifier`.
	Engine pulumi.StringPtrInput
	// Engine version of the Aurora global database. The `engine`, `engineVersion`, and `instanceClass` (on the `rds.ClusterInstance`) must together support global databases. See [Using Amazon Aurora global databases](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html) for more information. By upgrading the engine version, the provider will upgrade cluster members. **NOTE:** To avoid an `inconsistent final plan` error while upgrading, use the `lifecycle` `ignoreChanges` for `engineVersion` meta argument on the associated `rds.Cluster` resource as shown above in Upgrading Engine Versions example.
	EngineVersion       pulumi.StringPtrInput
	EngineVersionActual pulumi.StringPtrInput
	// Enable to remove DB Cluster members from Global Cluster on destroy. Required with `sourceDbClusterIdentifier`.
	ForceDestroy pulumi.BoolPtrInput
	// Global cluster identifier.
	GlobalClusterIdentifier pulumi.StringPtrInput
	// Set of objects containing Global Cluster members.
	GlobalClusterMembers GlobalClusterGlobalClusterMemberArrayInput
	// AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed
	GlobalClusterResourceId pulumi.StringPtrInput
	// Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
	SourceDbClusterIdentifier pulumi.StringPtrInput
	// Specifies whether the DB cluster is encrypted. The default is `false` unless `sourceDbClusterIdentifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
	StorageEncrypted pulumi.BoolPtrInput
}

func (GlobalClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalClusterState)(nil)).Elem()
}

type globalClusterArgs struct {
	// Name for an automatically created database on cluster creation.
	DatabaseName *string `pulumi:"databaseName"`
	// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Valid values: `aurora`, `aurora-mysql`, `aurora-postgresql`. Defaults to `aurora`. Conflicts with `sourceDbClusterIdentifier`.
	Engine *string `pulumi:"engine"`
	// Engine version of the Aurora global database. The `engine`, `engineVersion`, and `instanceClass` (on the `rds.ClusterInstance`) must together support global databases. See [Using Amazon Aurora global databases](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html) for more information. By upgrading the engine version, the provider will upgrade cluster members. **NOTE:** To avoid an `inconsistent final plan` error while upgrading, use the `lifecycle` `ignoreChanges` for `engineVersion` meta argument on the associated `rds.Cluster` resource as shown above in Upgrading Engine Versions example.
	EngineVersion *string `pulumi:"engineVersion"`
	// Enable to remove DB Cluster members from Global Cluster on destroy. Required with `sourceDbClusterIdentifier`.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// Global cluster identifier.
	GlobalClusterIdentifier string `pulumi:"globalClusterIdentifier"`
	// Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
	SourceDbClusterIdentifier *string `pulumi:"sourceDbClusterIdentifier"`
	// Specifies whether the DB cluster is encrypted. The default is `false` unless `sourceDbClusterIdentifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
	StorageEncrypted *bool `pulumi:"storageEncrypted"`
}

// The set of arguments for constructing a GlobalCluster resource.
type GlobalClusterArgs struct {
	// Name for an automatically created database on cluster creation.
	DatabaseName pulumi.StringPtrInput
	// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
	DeletionProtection pulumi.BoolPtrInput
	// Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Valid values: `aurora`, `aurora-mysql`, `aurora-postgresql`. Defaults to `aurora`. Conflicts with `sourceDbClusterIdentifier`.
	Engine pulumi.StringPtrInput
	// Engine version of the Aurora global database. The `engine`, `engineVersion`, and `instanceClass` (on the `rds.ClusterInstance`) must together support global databases. See [Using Amazon Aurora global databases](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html) for more information. By upgrading the engine version, the provider will upgrade cluster members. **NOTE:** To avoid an `inconsistent final plan` error while upgrading, use the `lifecycle` `ignoreChanges` for `engineVersion` meta argument on the associated `rds.Cluster` resource as shown above in Upgrading Engine Versions example.
	EngineVersion pulumi.StringPtrInput
	// Enable to remove DB Cluster members from Global Cluster on destroy. Required with `sourceDbClusterIdentifier`.
	ForceDestroy pulumi.BoolPtrInput
	// Global cluster identifier.
	GlobalClusterIdentifier pulumi.StringInput
	// Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
	SourceDbClusterIdentifier pulumi.StringPtrInput
	// Specifies whether the DB cluster is encrypted. The default is `false` unless `sourceDbClusterIdentifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
	StorageEncrypted pulumi.BoolPtrInput
}

func (GlobalClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalClusterArgs)(nil)).Elem()
}

type GlobalClusterInput interface {
	pulumi.Input

	ToGlobalClusterOutput() GlobalClusterOutput
	ToGlobalClusterOutputWithContext(ctx context.Context) GlobalClusterOutput
}

func (*GlobalCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalCluster)(nil)).Elem()
}

func (i *GlobalCluster) ToGlobalClusterOutput() GlobalClusterOutput {
	return i.ToGlobalClusterOutputWithContext(context.Background())
}

func (i *GlobalCluster) ToGlobalClusterOutputWithContext(ctx context.Context) GlobalClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalClusterOutput)
}

func (i *GlobalCluster) ToOutput(ctx context.Context) pulumix.Output[*GlobalCluster] {
	return pulumix.Output[*GlobalCluster]{
		OutputState: i.ToGlobalClusterOutputWithContext(ctx).OutputState,
	}
}

// GlobalClusterArrayInput is an input type that accepts GlobalClusterArray and GlobalClusterArrayOutput values.
// You can construct a concrete instance of `GlobalClusterArrayInput` via:
//
//	GlobalClusterArray{ GlobalClusterArgs{...} }
type GlobalClusterArrayInput interface {
	pulumi.Input

	ToGlobalClusterArrayOutput() GlobalClusterArrayOutput
	ToGlobalClusterArrayOutputWithContext(context.Context) GlobalClusterArrayOutput
}

type GlobalClusterArray []GlobalClusterInput

func (GlobalClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GlobalCluster)(nil)).Elem()
}

func (i GlobalClusterArray) ToGlobalClusterArrayOutput() GlobalClusterArrayOutput {
	return i.ToGlobalClusterArrayOutputWithContext(context.Background())
}

func (i GlobalClusterArray) ToGlobalClusterArrayOutputWithContext(ctx context.Context) GlobalClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalClusterArrayOutput)
}

func (i GlobalClusterArray) ToOutput(ctx context.Context) pulumix.Output[[]*GlobalCluster] {
	return pulumix.Output[[]*GlobalCluster]{
		OutputState: i.ToGlobalClusterArrayOutputWithContext(ctx).OutputState,
	}
}

// GlobalClusterMapInput is an input type that accepts GlobalClusterMap and GlobalClusterMapOutput values.
// You can construct a concrete instance of `GlobalClusterMapInput` via:
//
//	GlobalClusterMap{ "key": GlobalClusterArgs{...} }
type GlobalClusterMapInput interface {
	pulumi.Input

	ToGlobalClusterMapOutput() GlobalClusterMapOutput
	ToGlobalClusterMapOutputWithContext(context.Context) GlobalClusterMapOutput
}

type GlobalClusterMap map[string]GlobalClusterInput

func (GlobalClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GlobalCluster)(nil)).Elem()
}

func (i GlobalClusterMap) ToGlobalClusterMapOutput() GlobalClusterMapOutput {
	return i.ToGlobalClusterMapOutputWithContext(context.Background())
}

func (i GlobalClusterMap) ToGlobalClusterMapOutputWithContext(ctx context.Context) GlobalClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalClusterMapOutput)
}

func (i GlobalClusterMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*GlobalCluster] {
	return pulumix.Output[map[string]*GlobalCluster]{
		OutputState: i.ToGlobalClusterMapOutputWithContext(ctx).OutputState,
	}
}

type GlobalClusterOutput struct{ *pulumi.OutputState }

func (GlobalClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalCluster)(nil)).Elem()
}

func (o GlobalClusterOutput) ToGlobalClusterOutput() GlobalClusterOutput {
	return o
}

func (o GlobalClusterOutput) ToGlobalClusterOutputWithContext(ctx context.Context) GlobalClusterOutput {
	return o
}

func (o GlobalClusterOutput) ToOutput(ctx context.Context) pulumix.Output[*GlobalCluster] {
	return pulumix.Output[*GlobalCluster]{
		OutputState: o.OutputState,
	}
}

// RDS Global Cluster Amazon Resource Name (ARN)
func (o GlobalClusterOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalCluster) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Name for an automatically created database on cluster creation.
func (o GlobalClusterOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GlobalCluster) pulumi.StringPtrOutput { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
func (o GlobalClusterOutput) DeletionProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GlobalCluster) pulumi.BoolPtrOutput { return v.DeletionProtection }).(pulumi.BoolPtrOutput)
}

// Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Valid values: `aurora`, `aurora-mysql`, `aurora-postgresql`. Defaults to `aurora`. Conflicts with `sourceDbClusterIdentifier`.
func (o GlobalClusterOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalCluster) pulumi.StringOutput { return v.Engine }).(pulumi.StringOutput)
}

// Engine version of the Aurora global database. The `engine`, `engineVersion`, and `instanceClass` (on the `rds.ClusterInstance`) must together support global databases. See [Using Amazon Aurora global databases](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html) for more information. By upgrading the engine version, the provider will upgrade cluster members. **NOTE:** To avoid an `inconsistent final plan` error while upgrading, use the `lifecycle` `ignoreChanges` for `engineVersion` meta argument on the associated `rds.Cluster` resource as shown above in Upgrading Engine Versions example.
func (o GlobalClusterOutput) EngineVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalCluster) pulumi.StringOutput { return v.EngineVersion }).(pulumi.StringOutput)
}

func (o GlobalClusterOutput) EngineVersionActual() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalCluster) pulumi.StringOutput { return v.EngineVersionActual }).(pulumi.StringOutput)
}

// Enable to remove DB Cluster members from Global Cluster on destroy. Required with `sourceDbClusterIdentifier`.
func (o GlobalClusterOutput) ForceDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GlobalCluster) pulumi.BoolPtrOutput { return v.ForceDestroy }).(pulumi.BoolPtrOutput)
}

// Global cluster identifier.
func (o GlobalClusterOutput) GlobalClusterIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalCluster) pulumi.StringOutput { return v.GlobalClusterIdentifier }).(pulumi.StringOutput)
}

// Set of objects containing Global Cluster members.
func (o GlobalClusterOutput) GlobalClusterMembers() GlobalClusterGlobalClusterMemberArrayOutput {
	return o.ApplyT(func(v *GlobalCluster) GlobalClusterGlobalClusterMemberArrayOutput { return v.GlobalClusterMembers }).(GlobalClusterGlobalClusterMemberArrayOutput)
}

// AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed
func (o GlobalClusterOutput) GlobalClusterResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalCluster) pulumi.StringOutput { return v.GlobalClusterResourceId }).(pulumi.StringOutput)
}

// Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
func (o GlobalClusterOutput) SourceDbClusterIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *GlobalCluster) pulumi.StringOutput { return v.SourceDbClusterIdentifier }).(pulumi.StringOutput)
}

// Specifies whether the DB cluster is encrypted. The default is `false` unless `sourceDbClusterIdentifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
func (o GlobalClusterOutput) StorageEncrypted() pulumi.BoolOutput {
	return o.ApplyT(func(v *GlobalCluster) pulumi.BoolOutput { return v.StorageEncrypted }).(pulumi.BoolOutput)
}

type GlobalClusterArrayOutput struct{ *pulumi.OutputState }

func (GlobalClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GlobalCluster)(nil)).Elem()
}

func (o GlobalClusterArrayOutput) ToGlobalClusterArrayOutput() GlobalClusterArrayOutput {
	return o
}

func (o GlobalClusterArrayOutput) ToGlobalClusterArrayOutputWithContext(ctx context.Context) GlobalClusterArrayOutput {
	return o
}

func (o GlobalClusterArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*GlobalCluster] {
	return pulumix.Output[[]*GlobalCluster]{
		OutputState: o.OutputState,
	}
}

func (o GlobalClusterArrayOutput) Index(i pulumi.IntInput) GlobalClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GlobalCluster {
		return vs[0].([]*GlobalCluster)[vs[1].(int)]
	}).(GlobalClusterOutput)
}

type GlobalClusterMapOutput struct{ *pulumi.OutputState }

func (GlobalClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GlobalCluster)(nil)).Elem()
}

func (o GlobalClusterMapOutput) ToGlobalClusterMapOutput() GlobalClusterMapOutput {
	return o
}

func (o GlobalClusterMapOutput) ToGlobalClusterMapOutputWithContext(ctx context.Context) GlobalClusterMapOutput {
	return o
}

func (o GlobalClusterMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*GlobalCluster] {
	return pulumix.Output[map[string]*GlobalCluster]{
		OutputState: o.OutputState,
	}
}

func (o GlobalClusterMapOutput) MapIndex(k pulumi.StringInput) GlobalClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GlobalCluster {
		return vs[0].(map[string]*GlobalCluster)[vs[1].(string)]
	}).(GlobalClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalClusterInput)(nil)).Elem(), &GlobalCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalClusterArrayInput)(nil)).Elem(), GlobalClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalClusterMapInput)(nil)).Elem(), GlobalClusterMap{})
	pulumi.RegisterOutputType(GlobalClusterOutput{})
	pulumi.RegisterOutputType(GlobalClusterArrayOutput{})
	pulumi.RegisterOutputType(GlobalClusterMapOutput{})
}
