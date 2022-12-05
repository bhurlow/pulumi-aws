// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package docdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an DocDB Cluster Resource Instance. A Cluster Instance Resource defines
// attributes that are specific to a single instance in a DocDB Cluster.
//
// You do not designate a primary and subsequent replicas. Instead, you simply add DocDB
// Instances and DocDB manages the replication. You can use the count
// meta-parameter to make multiple instances and join them all to the same DocDB
// Cluster, or you may specify different Cluster Instance resources with various
// `instanceClass` sizes.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/docdb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := docdb.NewCluster(ctx, "default", &docdb.ClusterArgs{
//				ClusterIdentifier: pulumi.String("docdb-cluster-demo"),
//				AvailabilityZones: pulumi.StringArray{
//					pulumi.String("us-west-2a"),
//					pulumi.String("us-west-2b"),
//					pulumi.String("us-west-2c"),
//				},
//				MasterUsername: pulumi.String("foo"),
//				MasterPassword: pulumi.String("barbut8chars"),
//			})
//			if err != nil {
//				return err
//			}
//			var clusterInstances []*docdb.ClusterInstance
//			for index := 0; index < 2; index++ {
//				key0 := index
//				val0 := index
//				__res, err := docdb.NewClusterInstance(ctx, fmt.Sprintf("clusterInstances-%v", key0), &docdb.ClusterInstanceArgs{
//					Identifier:        pulumi.String(fmt.Sprintf("docdb-cluster-demo-%v", val0)),
//					ClusterIdentifier: _default.ID(),
//					InstanceClass:     pulumi.String("db.r5.large"),
//				})
//				if err != nil {
//					return err
//				}
//				clusterInstances = append(clusterInstances, __res)
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// DocDB Cluster Instances can be imported using the `identifier`, e.g.,
//
// ```sh
//
//	$ pulumi import aws:docdb/clusterInstance:ClusterInstance prod_instance_1 aurora-cluster-instance-1
//
// ```
type ClusterInstance struct {
	pulumi.CustomResourceState

	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is`false`.
	ApplyImmediately pulumi.BoolOutput `pulumi:"applyImmediately"`
	// Amazon Resource Name (ARN) of cluster instance
	Arn pulumi.StringOutput `pulumi:"arn"`
	// This parameter does not apply to Amazon DocumentDB. Amazon DocumentDB does not perform minor version upgrades regardless of the value set (see [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBInstance.html)). Default `true`.
	AutoMinorVersionUpgrade pulumi.BoolPtrOutput `pulumi:"autoMinorVersionUpgrade"`
	// The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// (Optional) The identifier of the CA certificate for the DB instance.
	CaCertIdentifier pulumi.StringOutput `pulumi:"caCertIdentifier"`
	// The identifier of the `docdb.Cluster` in which to launch this instance.
	ClusterIdentifier pulumi.StringOutput `pulumi:"clusterIdentifier"`
	// The DB subnet group to associate with this DB instance.
	DbSubnetGroupName pulumi.StringOutput `pulumi:"dbSubnetGroupName"`
	// The region-unique, immutable identifier for the DB instance.
	DbiResourceId pulumi.StringOutput `pulumi:"dbiResourceId"`
	// A value that indicates whether to enable Performance Insights for the DB Instance. Default `false`. See [docs] (https://docs.aws.amazon.com/documentdb/latest/developerguide/performance-insights.html) about the details.
	EnablePerformanceInsights pulumi.BoolOutput `pulumi:"enablePerformanceInsights"`
	// The DNS address for this instance. May not be writable
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
	Engine pulumi.StringPtrOutput `pulumi:"engine"`
	// The database engine version
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// The identifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
	Identifier pulumi.StringOutput `pulumi:"identifier"`
	// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix pulumi.StringOutput `pulumi:"identifierPrefix"`
	// The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
	// supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
	// - db.r5.large
	// - db.r5.xlarge
	// - db.r5.2xlarge
	// - db.r5.4xlarge
	// - db.r5.12xlarge
	// - db.r5.24xlarge
	// - db.r4.large
	// - db.r4.xlarge
	// - db.r4.2xlarge
	// - db.r4.4xlarge
	// - db.r4.8xlarge
	// - db.r4.16xlarge
	// - db.t3.medium
	InstanceClass pulumi.StringOutput `pulumi:"instanceClass"`
	// The ARN for the KMS encryption key if one is set to the cluster.
	KmsKeyId pulumi.StringOutput `pulumi:"kmsKeyId"`
	// The KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. If you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon DocumentDB uses your default KMS key.
	PerformanceInsightsKmsKeyId pulumi.StringOutput `pulumi:"performanceInsightsKmsKeyId"`
	// The database port
	Port pulumi.IntOutput `pulumi:"port"`
	// The daily time range during which automated backups are created if automated backups are enabled.
	PreferredBackupWindow pulumi.StringOutput `pulumi:"preferredBackupWindow"`
	// The window to perform maintenance in.
	// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	PreferredMaintenanceWindow pulumi.StringOutput `pulumi:"preferredMaintenanceWindow"`
	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
	PromotionTier      pulumi.IntPtrOutput `pulumi:"promotionTier"`
	PubliclyAccessible pulumi.BoolOutput   `pulumi:"publiclyAccessible"`
	// Specifies whether the DB cluster is encrypted.
	StorageEncrypted pulumi.BoolOutput `pulumi:"storageEncrypted"`
	// A map of tags to assign to the instance. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
	Writer pulumi.BoolOutput `pulumi:"writer"`
}

// NewClusterInstance registers a new resource with the given unique name, arguments, and options.
func NewClusterInstance(ctx *pulumi.Context,
	name string, args *ClusterInstanceArgs, opts ...pulumi.ResourceOption) (*ClusterInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'ClusterIdentifier'")
	}
	if args.InstanceClass == nil {
		return nil, errors.New("invalid value for required argument 'InstanceClass'")
	}
	var resource ClusterInstance
	err := ctx.RegisterResource("aws:docdb/clusterInstance:ClusterInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterInstance gets an existing ClusterInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterInstanceState, opts ...pulumi.ResourceOption) (*ClusterInstance, error) {
	var resource ClusterInstance
	err := ctx.ReadResource("aws:docdb/clusterInstance:ClusterInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterInstance resources.
type clusterInstanceState struct {
	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is`false`.
	ApplyImmediately *bool `pulumi:"applyImmediately"`
	// Amazon Resource Name (ARN) of cluster instance
	Arn *string `pulumi:"arn"`
	// This parameter does not apply to Amazon DocumentDB. Amazon DocumentDB does not perform minor version upgrades regardless of the value set (see [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBInstance.html)). Default `true`.
	AutoMinorVersionUpgrade *bool `pulumi:"autoMinorVersionUpgrade"`
	// The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// (Optional) The identifier of the CA certificate for the DB instance.
	CaCertIdentifier *string `pulumi:"caCertIdentifier"`
	// The identifier of the `docdb.Cluster` in which to launch this instance.
	ClusterIdentifier *string `pulumi:"clusterIdentifier"`
	// The DB subnet group to associate with this DB instance.
	DbSubnetGroupName *string `pulumi:"dbSubnetGroupName"`
	// The region-unique, immutable identifier for the DB instance.
	DbiResourceId *string `pulumi:"dbiResourceId"`
	// A value that indicates whether to enable Performance Insights for the DB Instance. Default `false`. See [docs] (https://docs.aws.amazon.com/documentdb/latest/developerguide/performance-insights.html) about the details.
	EnablePerformanceInsights *bool `pulumi:"enablePerformanceInsights"`
	// The DNS address for this instance. May not be writable
	Endpoint *string `pulumi:"endpoint"`
	// The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
	Engine *string `pulumi:"engine"`
	// The database engine version
	EngineVersion *string `pulumi:"engineVersion"`
	// The identifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
	Identifier *string `pulumi:"identifier"`
	// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix *string `pulumi:"identifierPrefix"`
	// The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
	// supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
	// - db.r5.large
	// - db.r5.xlarge
	// - db.r5.2xlarge
	// - db.r5.4xlarge
	// - db.r5.12xlarge
	// - db.r5.24xlarge
	// - db.r4.large
	// - db.r4.xlarge
	// - db.r4.2xlarge
	// - db.r4.4xlarge
	// - db.r4.8xlarge
	// - db.r4.16xlarge
	// - db.t3.medium
	InstanceClass *string `pulumi:"instanceClass"`
	// The ARN for the KMS encryption key if one is set to the cluster.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. If you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon DocumentDB uses your default KMS key.
	PerformanceInsightsKmsKeyId *string `pulumi:"performanceInsightsKmsKeyId"`
	// The database port
	Port *int `pulumi:"port"`
	// The daily time range during which automated backups are created if automated backups are enabled.
	PreferredBackupWindow *string `pulumi:"preferredBackupWindow"`
	// The window to perform maintenance in.
	// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	PreferredMaintenanceWindow *string `pulumi:"preferredMaintenanceWindow"`
	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
	PromotionTier      *int  `pulumi:"promotionTier"`
	PubliclyAccessible *bool `pulumi:"publiclyAccessible"`
	// Specifies whether the DB cluster is encrypted.
	StorageEncrypted *bool `pulumi:"storageEncrypted"`
	// A map of tags to assign to the instance. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
	Writer *bool `pulumi:"writer"`
}

type ClusterInstanceState struct {
	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is`false`.
	ApplyImmediately pulumi.BoolPtrInput
	// Amazon Resource Name (ARN) of cluster instance
	Arn pulumi.StringPtrInput
	// This parameter does not apply to Amazon DocumentDB. Amazon DocumentDB does not perform minor version upgrades regardless of the value set (see [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBInstance.html)). Default `true`.
	AutoMinorVersionUpgrade pulumi.BoolPtrInput
	// The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
	AvailabilityZone pulumi.StringPtrInput
	// (Optional) The identifier of the CA certificate for the DB instance.
	CaCertIdentifier pulumi.StringPtrInput
	// The identifier of the `docdb.Cluster` in which to launch this instance.
	ClusterIdentifier pulumi.StringPtrInput
	// The DB subnet group to associate with this DB instance.
	DbSubnetGroupName pulumi.StringPtrInput
	// The region-unique, immutable identifier for the DB instance.
	DbiResourceId pulumi.StringPtrInput
	// A value that indicates whether to enable Performance Insights for the DB Instance. Default `false`. See [docs] (https://docs.aws.amazon.com/documentdb/latest/developerguide/performance-insights.html) about the details.
	EnablePerformanceInsights pulumi.BoolPtrInput
	// The DNS address for this instance. May not be writable
	Endpoint pulumi.StringPtrInput
	// The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
	Engine pulumi.StringPtrInput
	// The database engine version
	EngineVersion pulumi.StringPtrInput
	// The identifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
	Identifier pulumi.StringPtrInput
	// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix pulumi.StringPtrInput
	// The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
	// supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
	// - db.r5.large
	// - db.r5.xlarge
	// - db.r5.2xlarge
	// - db.r5.4xlarge
	// - db.r5.12xlarge
	// - db.r5.24xlarge
	// - db.r4.large
	// - db.r4.xlarge
	// - db.r4.2xlarge
	// - db.r4.4xlarge
	// - db.r4.8xlarge
	// - db.r4.16xlarge
	// - db.t3.medium
	InstanceClass pulumi.StringPtrInput
	// The ARN for the KMS encryption key if one is set to the cluster.
	KmsKeyId pulumi.StringPtrInput
	// The KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. If you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon DocumentDB uses your default KMS key.
	PerformanceInsightsKmsKeyId pulumi.StringPtrInput
	// The database port
	Port pulumi.IntPtrInput
	// The daily time range during which automated backups are created if automated backups are enabled.
	PreferredBackupWindow pulumi.StringPtrInput
	// The window to perform maintenance in.
	// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	PreferredMaintenanceWindow pulumi.StringPtrInput
	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
	PromotionTier      pulumi.IntPtrInput
	PubliclyAccessible pulumi.BoolPtrInput
	// Specifies whether the DB cluster is encrypted.
	StorageEncrypted pulumi.BoolPtrInput
	// A map of tags to assign to the instance. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
	TagsAll pulumi.StringMapInput
	// Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
	Writer pulumi.BoolPtrInput
}

func (ClusterInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterInstanceState)(nil)).Elem()
}

type clusterInstanceArgs struct {
	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is`false`.
	ApplyImmediately *bool `pulumi:"applyImmediately"`
	// This parameter does not apply to Amazon DocumentDB. Amazon DocumentDB does not perform minor version upgrades regardless of the value set (see [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBInstance.html)). Default `true`.
	AutoMinorVersionUpgrade *bool `pulumi:"autoMinorVersionUpgrade"`
	// The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// (Optional) The identifier of the CA certificate for the DB instance.
	CaCertIdentifier *string `pulumi:"caCertIdentifier"`
	// The identifier of the `docdb.Cluster` in which to launch this instance.
	ClusterIdentifier string `pulumi:"clusterIdentifier"`
	// A value that indicates whether to enable Performance Insights for the DB Instance. Default `false`. See [docs] (https://docs.aws.amazon.com/documentdb/latest/developerguide/performance-insights.html) about the details.
	EnablePerformanceInsights *bool `pulumi:"enablePerformanceInsights"`
	// The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
	Engine *string `pulumi:"engine"`
	// The identifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
	Identifier *string `pulumi:"identifier"`
	// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix *string `pulumi:"identifierPrefix"`
	// The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
	// supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
	// - db.r5.large
	// - db.r5.xlarge
	// - db.r5.2xlarge
	// - db.r5.4xlarge
	// - db.r5.12xlarge
	// - db.r5.24xlarge
	// - db.r4.large
	// - db.r4.xlarge
	// - db.r4.2xlarge
	// - db.r4.4xlarge
	// - db.r4.8xlarge
	// - db.r4.16xlarge
	// - db.t3.medium
	InstanceClass string `pulumi:"instanceClass"`
	// The KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. If you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon DocumentDB uses your default KMS key.
	PerformanceInsightsKmsKeyId *string `pulumi:"performanceInsightsKmsKeyId"`
	// The window to perform maintenance in.
	// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	PreferredMaintenanceWindow *string `pulumi:"preferredMaintenanceWindow"`
	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
	PromotionTier *int `pulumi:"promotionTier"`
	// A map of tags to assign to the instance. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ClusterInstance resource.
type ClusterInstanceArgs struct {
	// Specifies whether any database modifications
	// are applied immediately, or during the next maintenance window. Default is`false`.
	ApplyImmediately pulumi.BoolPtrInput
	// This parameter does not apply to Amazon DocumentDB. Amazon DocumentDB does not perform minor version upgrades regardless of the value set (see [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBInstance.html)). Default `true`.
	AutoMinorVersionUpgrade pulumi.BoolPtrInput
	// The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
	AvailabilityZone pulumi.StringPtrInput
	// (Optional) The identifier of the CA certificate for the DB instance.
	CaCertIdentifier pulumi.StringPtrInput
	// The identifier of the `docdb.Cluster` in which to launch this instance.
	ClusterIdentifier pulumi.StringInput
	// A value that indicates whether to enable Performance Insights for the DB Instance. Default `false`. See [docs] (https://docs.aws.amazon.com/documentdb/latest/developerguide/performance-insights.html) about the details.
	EnablePerformanceInsights pulumi.BoolPtrInput
	// The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
	Engine pulumi.StringPtrInput
	// The identifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
	Identifier pulumi.StringPtrInput
	// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix pulumi.StringPtrInput
	// The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
	// supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
	// - db.r5.large
	// - db.r5.xlarge
	// - db.r5.2xlarge
	// - db.r5.4xlarge
	// - db.r5.12xlarge
	// - db.r5.24xlarge
	// - db.r4.large
	// - db.r4.xlarge
	// - db.r4.2xlarge
	// - db.r4.4xlarge
	// - db.r4.8xlarge
	// - db.r4.16xlarge
	// - db.t3.medium
	InstanceClass pulumi.StringInput
	// The KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. If you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon DocumentDB uses your default KMS key.
	PerformanceInsightsKmsKeyId pulumi.StringPtrInput
	// The window to perform maintenance in.
	// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	PreferredMaintenanceWindow pulumi.StringPtrInput
	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
	PromotionTier pulumi.IntPtrInput
	// A map of tags to assign to the instance. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (ClusterInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterInstanceArgs)(nil)).Elem()
}

type ClusterInstanceInput interface {
	pulumi.Input

	ToClusterInstanceOutput() ClusterInstanceOutput
	ToClusterInstanceOutputWithContext(ctx context.Context) ClusterInstanceOutput
}

func (*ClusterInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterInstance)(nil)).Elem()
}

func (i *ClusterInstance) ToClusterInstanceOutput() ClusterInstanceOutput {
	return i.ToClusterInstanceOutputWithContext(context.Background())
}

func (i *ClusterInstance) ToClusterInstanceOutputWithContext(ctx context.Context) ClusterInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterInstanceOutput)
}

// ClusterInstanceArrayInput is an input type that accepts ClusterInstanceArray and ClusterInstanceArrayOutput values.
// You can construct a concrete instance of `ClusterInstanceArrayInput` via:
//
//	ClusterInstanceArray{ ClusterInstanceArgs{...} }
type ClusterInstanceArrayInput interface {
	pulumi.Input

	ToClusterInstanceArrayOutput() ClusterInstanceArrayOutput
	ToClusterInstanceArrayOutputWithContext(context.Context) ClusterInstanceArrayOutput
}

type ClusterInstanceArray []ClusterInstanceInput

func (ClusterInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterInstance)(nil)).Elem()
}

func (i ClusterInstanceArray) ToClusterInstanceArrayOutput() ClusterInstanceArrayOutput {
	return i.ToClusterInstanceArrayOutputWithContext(context.Background())
}

func (i ClusterInstanceArray) ToClusterInstanceArrayOutputWithContext(ctx context.Context) ClusterInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterInstanceArrayOutput)
}

// ClusterInstanceMapInput is an input type that accepts ClusterInstanceMap and ClusterInstanceMapOutput values.
// You can construct a concrete instance of `ClusterInstanceMapInput` via:
//
//	ClusterInstanceMap{ "key": ClusterInstanceArgs{...} }
type ClusterInstanceMapInput interface {
	pulumi.Input

	ToClusterInstanceMapOutput() ClusterInstanceMapOutput
	ToClusterInstanceMapOutputWithContext(context.Context) ClusterInstanceMapOutput
}

type ClusterInstanceMap map[string]ClusterInstanceInput

func (ClusterInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterInstance)(nil)).Elem()
}

func (i ClusterInstanceMap) ToClusterInstanceMapOutput() ClusterInstanceMapOutput {
	return i.ToClusterInstanceMapOutputWithContext(context.Background())
}

func (i ClusterInstanceMap) ToClusterInstanceMapOutputWithContext(ctx context.Context) ClusterInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterInstanceMapOutput)
}

type ClusterInstanceOutput struct{ *pulumi.OutputState }

func (ClusterInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterInstance)(nil)).Elem()
}

func (o ClusterInstanceOutput) ToClusterInstanceOutput() ClusterInstanceOutput {
	return o
}

func (o ClusterInstanceOutput) ToClusterInstanceOutputWithContext(ctx context.Context) ClusterInstanceOutput {
	return o
}

// Specifies whether any database modifications
// are applied immediately, or during the next maintenance window. Default is`false`.
func (o ClusterInstanceOutput) ApplyImmediately() pulumi.BoolOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.BoolOutput { return v.ApplyImmediately }).(pulumi.BoolOutput)
}

// Amazon Resource Name (ARN) of cluster instance
func (o ClusterInstanceOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// This parameter does not apply to Amazon DocumentDB. Amazon DocumentDB does not perform minor version upgrades regardless of the value set (see [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBInstance.html)). Default `true`.
func (o ClusterInstanceOutput) AutoMinorVersionUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.BoolPtrOutput { return v.AutoMinorVersionUpgrade }).(pulumi.BoolPtrOutput)
}

// The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
func (o ClusterInstanceOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// (Optional) The identifier of the CA certificate for the DB instance.
func (o ClusterInstanceOutput) CaCertIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.StringOutput { return v.CaCertIdentifier }).(pulumi.StringOutput)
}

// The identifier of the `docdb.Cluster` in which to launch this instance.
func (o ClusterInstanceOutput) ClusterIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.StringOutput { return v.ClusterIdentifier }).(pulumi.StringOutput)
}

// The DB subnet group to associate with this DB instance.
func (o ClusterInstanceOutput) DbSubnetGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.StringOutput { return v.DbSubnetGroupName }).(pulumi.StringOutput)
}

// The region-unique, immutable identifier for the DB instance.
func (o ClusterInstanceOutput) DbiResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.StringOutput { return v.DbiResourceId }).(pulumi.StringOutput)
}

// A value that indicates whether to enable Performance Insights for the DB Instance. Default `false`. See [docs] (https://docs.aws.amazon.com/documentdb/latest/developerguide/performance-insights.html) about the details.
func (o ClusterInstanceOutput) EnablePerformanceInsights() pulumi.BoolOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.BoolOutput { return v.EnablePerformanceInsights }).(pulumi.BoolOutput)
}

// The DNS address for this instance. May not be writable
func (o ClusterInstanceOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
func (o ClusterInstanceOutput) Engine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.StringPtrOutput { return v.Engine }).(pulumi.StringPtrOutput)
}

// The database engine version
func (o ClusterInstanceOutput) EngineVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.StringOutput { return v.EngineVersion }).(pulumi.StringOutput)
}

// The identifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
func (o ClusterInstanceOutput) Identifier() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.StringOutput { return v.Identifier }).(pulumi.StringOutput)
}

// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
func (o ClusterInstanceOutput) IdentifierPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.StringOutput { return v.IdentifierPrefix }).(pulumi.StringOutput)
}

// The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
// supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
// - db.r5.large
// - db.r5.xlarge
// - db.r5.2xlarge
// - db.r5.4xlarge
// - db.r5.12xlarge
// - db.r5.24xlarge
// - db.r4.large
// - db.r4.xlarge
// - db.r4.2xlarge
// - db.r4.4xlarge
// - db.r4.8xlarge
// - db.r4.16xlarge
// - db.t3.medium
func (o ClusterInstanceOutput) InstanceClass() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.StringOutput { return v.InstanceClass }).(pulumi.StringOutput)
}

// The ARN for the KMS encryption key if one is set to the cluster.
func (o ClusterInstanceOutput) KmsKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.StringOutput { return v.KmsKeyId }).(pulumi.StringOutput)
}

// The KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. If you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon DocumentDB uses your default KMS key.
func (o ClusterInstanceOutput) PerformanceInsightsKmsKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.StringOutput { return v.PerformanceInsightsKmsKeyId }).(pulumi.StringOutput)
}

// The database port
func (o ClusterInstanceOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// The daily time range during which automated backups are created if automated backups are enabled.
func (o ClusterInstanceOutput) PreferredBackupWindow() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.StringOutput { return v.PreferredBackupWindow }).(pulumi.StringOutput)
}

// The window to perform maintenance in.
// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
func (o ClusterInstanceOutput) PreferredMaintenanceWindow() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.StringOutput { return v.PreferredMaintenanceWindow }).(pulumi.StringOutput)
}

// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
func (o ClusterInstanceOutput) PromotionTier() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.IntPtrOutput { return v.PromotionTier }).(pulumi.IntPtrOutput)
}

func (o ClusterInstanceOutput) PubliclyAccessible() pulumi.BoolOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.BoolOutput { return v.PubliclyAccessible }).(pulumi.BoolOutput)
}

// Specifies whether the DB cluster is encrypted.
func (o ClusterInstanceOutput) StorageEncrypted() pulumi.BoolOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.BoolOutput { return v.StorageEncrypted }).(pulumi.BoolOutput)
}

// A map of tags to assign to the instance. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o ClusterInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
func (o ClusterInstanceOutput) TagsAll() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.StringMapOutput { return v.TagsAll }).(pulumi.StringMapOutput)
}

// Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
func (o ClusterInstanceOutput) Writer() pulumi.BoolOutput {
	return o.ApplyT(func(v *ClusterInstance) pulumi.BoolOutput { return v.Writer }).(pulumi.BoolOutput)
}

type ClusterInstanceArrayOutput struct{ *pulumi.OutputState }

func (ClusterInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterInstance)(nil)).Elem()
}

func (o ClusterInstanceArrayOutput) ToClusterInstanceArrayOutput() ClusterInstanceArrayOutput {
	return o
}

func (o ClusterInstanceArrayOutput) ToClusterInstanceArrayOutputWithContext(ctx context.Context) ClusterInstanceArrayOutput {
	return o
}

func (o ClusterInstanceArrayOutput) Index(i pulumi.IntInput) ClusterInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterInstance {
		return vs[0].([]*ClusterInstance)[vs[1].(int)]
	}).(ClusterInstanceOutput)
}

type ClusterInstanceMapOutput struct{ *pulumi.OutputState }

func (ClusterInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterInstance)(nil)).Elem()
}

func (o ClusterInstanceMapOutput) ToClusterInstanceMapOutput() ClusterInstanceMapOutput {
	return o
}

func (o ClusterInstanceMapOutput) ToClusterInstanceMapOutputWithContext(ctx context.Context) ClusterInstanceMapOutput {
	return o
}

func (o ClusterInstanceMapOutput) MapIndex(k pulumi.StringInput) ClusterInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterInstance {
		return vs[0].(map[string]*ClusterInstance)[vs[1].(string)]
	}).(ClusterInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInstanceInput)(nil)).Elem(), &ClusterInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInstanceArrayInput)(nil)).Elem(), ClusterInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInstanceMapInput)(nil)).Elem(), ClusterInstanceMap{})
	pulumi.RegisterOutputType(ClusterInstanceOutput{})
	pulumi.RegisterOutputType(ClusterInstanceArrayOutput{})
	pulumi.RegisterOutputType(ClusterInstanceMapOutput{})
}
