// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DocDB
{
    /// <summary>
    /// Provides an DocDB Cluster Resource Instance. A Cluster Instance Resource defines
    /// attributes that are specific to a single instance in a [DocDB Cluster](https://www.terraform.io/docs/providers/aws/r/docdb_cluster.html).
    /// 
    /// You do not designate a primary and subsequent replicas. Instead, you simply add DocDB
    /// Instances and DocDB manages the replication. You can use the [count](https://www.terraform.io/docs/configuration/meta-arguments/count.html)
    /// meta-parameter to make multiple instances and join them all to the same DocDB
    /// Cluster, or you may specify different Cluster Instance resources with various
    /// `instance_class` sizes.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @default = new Aws.DocDB.Cluster("default", new Aws.DocDB.ClusterArgs
    ///         {
    ///             ClusterIdentifier = "docdb-cluster-demo",
    ///             AvailabilityZones = 
    ///             {
    ///                 "us-west-2a",
    ///                 "us-west-2b",
    ///                 "us-west-2c",
    ///             },
    ///             MasterUsername = "foo",
    ///             MasterPassword = "barbut8chars",
    ///         });
    ///         var clusterInstances = new List&lt;Aws.DocDB.ClusterInstance&gt;();
    ///         for (var rangeIndex = 0; rangeIndex &lt; 2; rangeIndex++)
    ///         {
    ///             var range = new { Value = rangeIndex };
    ///             clusterInstances.Add(new Aws.DocDB.ClusterInstance($"clusterInstances-{range.Value}", new Aws.DocDB.ClusterInstanceArgs
    ///             {
    ///                 Identifier = $"docdb-cluster-demo-{range.Value}",
    ///                 ClusterIdentifier = @default.Id,
    ///                 InstanceClass = "db.r5.large",
    ///             }));
    ///         }
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// DocDB Cluster Instances can be imported using the `identifier`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:docdb/clusterInstance:ClusterInstance prod_instance_1 aurora-cluster-instance-1
    /// ```
    /// </summary>
    [AwsResourceType("aws:docdb/clusterInstance:ClusterInstance")]
    public partial class ClusterInstance : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether any database modifications
        /// are applied immediately, or during the next maintenance window. Default is`false`.
        /// </summary>
        [Output("applyImmediately")]
        public Output<bool> ApplyImmediately { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of cluster instance
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
        /// </summary>
        [Output("autoMinorVersionUpgrade")]
        public Output<bool?> AutoMinorVersionUpgrade { get; private set; } = null!;

        /// <summary>
        /// The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// (Optional) The identifier of the CA certificate for the DB instance.
        /// </summary>
        [Output("caCertIdentifier")]
        public Output<string> CaCertIdentifier { get; private set; } = null!;

        /// <summary>
        /// The identifier of the `aws.docdb.Cluster` in which to launch this instance.
        /// </summary>
        [Output("clusterIdentifier")]
        public Output<string> ClusterIdentifier { get; private set; } = null!;

        /// <summary>
        /// The DB subnet group to associate with this DB instance.
        /// </summary>
        [Output("dbSubnetGroupName")]
        public Output<string> DbSubnetGroupName { get; private set; } = null!;

        /// <summary>
        /// The region-unique, immutable identifier for the DB instance.
        /// </summary>
        [Output("dbiResourceId")]
        public Output<string> DbiResourceId { get; private set; } = null!;

        /// <summary>
        /// The DNS address for this instance. May not be writable
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
        /// </summary>
        [Output("engine")]
        public Output<string?> Engine { get; private set; } = null!;

        /// <summary>
        /// The database engine version
        /// </summary>
        [Output("engineVersion")]
        public Output<string> EngineVersion { get; private set; } = null!;

        /// <summary>
        /// The identifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
        /// </summary>
        [Output("identifier")]
        public Output<string> Identifier { get; private set; } = null!;

        /// <summary>
        /// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
        /// </summary>
        [Output("identifierPrefix")]
        public Output<string> IdentifierPrefix { get; private set; } = null!;

        /// <summary>
        /// The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
        /// supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
        /// - db.r5.large
        /// - db.r5.xlarge
        /// - db.r5.2xlarge
        /// - db.r5.4xlarge
        /// - db.r5.12xlarge
        /// - db.r5.24xlarge
        /// - db.r4.large
        /// - db.r4.xlarge
        /// - db.r4.2xlarge
        /// - db.r4.4xlarge
        /// - db.r4.8xlarge
        /// - db.r4.16xlarge
        /// - db.t3.medium
        /// </summary>
        [Output("instanceClass")]
        public Output<string> InstanceClass { get; private set; } = null!;

        /// <summary>
        /// The ARN for the KMS encryption key if one is set to the cluster.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// The database port
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// The daily time range during which automated backups are created if automated backups are enabled.
        /// </summary>
        [Output("preferredBackupWindow")]
        public Output<string> PreferredBackupWindow { get; private set; } = null!;

        /// <summary>
        /// The window to perform maintenance in.
        /// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
        /// </summary>
        [Output("preferredMaintenanceWindow")]
        public Output<string> PreferredMaintenanceWindow { get; private set; } = null!;

        /// <summary>
        /// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
        /// </summary>
        [Output("promotionTier")]
        public Output<int?> PromotionTier { get; private set; } = null!;

        [Output("publiclyAccessible")]
        public Output<bool> PubliclyAccessible { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the DB cluster is encrypted.
        /// </summary>
        [Output("storageEncrypted")]
        public Output<bool> StorageEncrypted { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the instance. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
        /// </summary>
        [Output("writer")]
        public Output<bool> Writer { get; private set; } = null!;


        /// <summary>
        /// Create a ClusterInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClusterInstance(string name, ClusterInstanceArgs args, CustomResourceOptions? options = null)
            : base("aws:docdb/clusterInstance:ClusterInstance", name, args ?? new ClusterInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClusterInstance(string name, Input<string> id, ClusterInstanceState? state = null, CustomResourceOptions? options = null)
            : base("aws:docdb/clusterInstance:ClusterInstance", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ClusterInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClusterInstance Get(string name, Input<string> id, ClusterInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new ClusterInstance(name, id, state, options);
        }
    }

    public sealed class ClusterInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether any database modifications
        /// are applied immediately, or during the next maintenance window. Default is`false`.
        /// </summary>
        [Input("applyImmediately")]
        public Input<bool>? ApplyImmediately { get; set; }

        /// <summary>
        /// Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
        /// </summary>
        [Input("autoMinorVersionUpgrade")]
        public Input<bool>? AutoMinorVersionUpgrade { get; set; }

        /// <summary>
        /// The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// (Optional) The identifier of the CA certificate for the DB instance.
        /// </summary>
        [Input("caCertIdentifier")]
        public Input<string>? CaCertIdentifier { get; set; }

        /// <summary>
        /// The identifier of the `aws.docdb.Cluster` in which to launch this instance.
        /// </summary>
        [Input("clusterIdentifier", required: true)]
        public Input<string> ClusterIdentifier { get; set; } = null!;

        /// <summary>
        /// The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// The identifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
        /// </summary>
        [Input("identifier")]
        public Input<string>? Identifier { get; set; }

        /// <summary>
        /// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
        /// </summary>
        [Input("identifierPrefix")]
        public Input<string>? IdentifierPrefix { get; set; }

        /// <summary>
        /// The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
        /// supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
        /// - db.r5.large
        /// - db.r5.xlarge
        /// - db.r5.2xlarge
        /// - db.r5.4xlarge
        /// - db.r5.12xlarge
        /// - db.r5.24xlarge
        /// - db.r4.large
        /// - db.r4.xlarge
        /// - db.r4.2xlarge
        /// - db.r4.4xlarge
        /// - db.r4.8xlarge
        /// - db.r4.16xlarge
        /// - db.t3.medium
        /// </summary>
        [Input("instanceClass", required: true)]
        public Input<string> InstanceClass { get; set; } = null!;

        /// <summary>
        /// The window to perform maintenance in.
        /// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
        /// </summary>
        [Input("preferredMaintenanceWindow")]
        public Input<string>? PreferredMaintenanceWindow { get; set; }

        /// <summary>
        /// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
        /// </summary>
        [Input("promotionTier")]
        public Input<int>? PromotionTier { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the instance. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ClusterInstanceArgs()
        {
        }
    }

    public sealed class ClusterInstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether any database modifications
        /// are applied immediately, or during the next maintenance window. Default is`false`.
        /// </summary>
        [Input("applyImmediately")]
        public Input<bool>? ApplyImmediately { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of cluster instance
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
        /// </summary>
        [Input("autoMinorVersionUpgrade")]
        public Input<bool>? AutoMinorVersionUpgrade { get; set; }

        /// <summary>
        /// The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// (Optional) The identifier of the CA certificate for the DB instance.
        /// </summary>
        [Input("caCertIdentifier")]
        public Input<string>? CaCertIdentifier { get; set; }

        /// <summary>
        /// The identifier of the `aws.docdb.Cluster` in which to launch this instance.
        /// </summary>
        [Input("clusterIdentifier")]
        public Input<string>? ClusterIdentifier { get; set; }

        /// <summary>
        /// The DB subnet group to associate with this DB instance.
        /// </summary>
        [Input("dbSubnetGroupName")]
        public Input<string>? DbSubnetGroupName { get; set; }

        /// <summary>
        /// The region-unique, immutable identifier for the DB instance.
        /// </summary>
        [Input("dbiResourceId")]
        public Input<string>? DbiResourceId { get; set; }

        /// <summary>
        /// The DNS address for this instance. May not be writable
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// The database engine version
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// The identifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
        /// </summary>
        [Input("identifier")]
        public Input<string>? Identifier { get; set; }

        /// <summary>
        /// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
        /// </summary>
        [Input("identifierPrefix")]
        public Input<string>? IdentifierPrefix { get; set; }

        /// <summary>
        /// The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
        /// supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
        /// - db.r5.large
        /// - db.r5.xlarge
        /// - db.r5.2xlarge
        /// - db.r5.4xlarge
        /// - db.r5.12xlarge
        /// - db.r5.24xlarge
        /// - db.r4.large
        /// - db.r4.xlarge
        /// - db.r4.2xlarge
        /// - db.r4.4xlarge
        /// - db.r4.8xlarge
        /// - db.r4.16xlarge
        /// - db.t3.medium
        /// </summary>
        [Input("instanceClass")]
        public Input<string>? InstanceClass { get; set; }

        /// <summary>
        /// The ARN for the KMS encryption key if one is set to the cluster.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// The database port
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The daily time range during which automated backups are created if automated backups are enabled.
        /// </summary>
        [Input("preferredBackupWindow")]
        public Input<string>? PreferredBackupWindow { get; set; }

        /// <summary>
        /// The window to perform maintenance in.
        /// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
        /// </summary>
        [Input("preferredMaintenanceWindow")]
        public Input<string>? PreferredMaintenanceWindow { get; set; }

        /// <summary>
        /// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
        /// </summary>
        [Input("promotionTier")]
        public Input<int>? PromotionTier { get; set; }

        [Input("publiclyAccessible")]
        public Input<bool>? PubliclyAccessible { get; set; }

        /// <summary>
        /// Specifies whether the DB cluster is encrypted.
        /// </summary>
        [Input("storageEncrypted")]
        public Input<bool>? StorageEncrypted { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the instance. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// A map of tags assigned to the resource, including those inherited from the provider .
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
        /// </summary>
        [Input("writer")]
        public Input<bool>? Writer { get; set; }

        public ClusterInstanceState()
        {
        }
    }
}
