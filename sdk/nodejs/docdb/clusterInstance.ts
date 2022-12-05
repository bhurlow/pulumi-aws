// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an DocDB Cluster Resource Instance. A Cluster Instance Resource defines
 * attributes that are specific to a single instance in a DocDB Cluster.
 *
 * You do not designate a primary and subsequent replicas. Instead, you simply add DocDB
 * Instances and DocDB manages the replication. You can use the count
 * meta-parameter to make multiple instances and join them all to the same DocDB
 * Cluster, or you may specify different Cluster Instance resources with various
 * `instanceClass` sizes.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const _default = new aws.docdb.Cluster("default", {
 *     clusterIdentifier: "docdb-cluster-demo",
 *     availabilityZones: [
 *         "us-west-2a",
 *         "us-west-2b",
 *         "us-west-2c",
 *     ],
 *     masterUsername: "foo",
 *     masterPassword: "barbut8chars",
 * });
 * const clusterInstances: aws.docdb.ClusterInstance[] = [];
 * for (const range = {value: 0}; range.value < 2; range.value++) {
 *     clusterInstances.push(new aws.docdb.ClusterInstance(`clusterInstances-${range.value}`, {
 *         identifier: `docdb-cluster-demo-${range.value}`,
 *         clusterIdentifier: _default.id,
 *         instanceClass: "db.r5.large",
 *     }));
 * }
 * ```
 *
 * ## Import
 *
 * DocDB Cluster Instances can be imported using the `identifier`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:docdb/clusterInstance:ClusterInstance prod_instance_1 aurora-cluster-instance-1
 * ```
 */
export class ClusterInstance extends pulumi.CustomResource {
    /**
     * Get an existing ClusterInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterInstanceState, opts?: pulumi.CustomResourceOptions): ClusterInstance {
        return new ClusterInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:docdb/clusterInstance:ClusterInstance';

    /**
     * Returns true if the given object is an instance of ClusterInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClusterInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClusterInstance.__pulumiType;
    }

    /**
     * Specifies whether any database modifications
     * are applied immediately, or during the next maintenance window. Default is`false`.
     */
    public readonly applyImmediately!: pulumi.Output<boolean>;
    /**
     * Amazon Resource Name (ARN) of cluster instance
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * This parameter does not apply to Amazon DocumentDB. Amazon DocumentDB does not perform minor version upgrades regardless of the value set (see [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBInstance.html)). Default `true`.
     */
    public readonly autoMinorVersionUpgrade!: pulumi.Output<boolean | undefined>;
    /**
     * The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * (Optional) The identifier of the CA certificate for the DB instance.
     */
    public readonly caCertIdentifier!: pulumi.Output<string>;
    /**
     * The identifier of the `aws.docdb.Cluster` in which to launch this instance.
     */
    public readonly clusterIdentifier!: pulumi.Output<string>;
    /**
     * The DB subnet group to associate with this DB instance.
     */
    public /*out*/ readonly dbSubnetGroupName!: pulumi.Output<string>;
    /**
     * The region-unique, immutable identifier for the DB instance.
     */
    public /*out*/ readonly dbiResourceId!: pulumi.Output<string>;
    /**
     * A value that indicates whether to enable Performance Insights for the DB Instance. Default `false`. See [docs] (https://docs.aws.amazon.com/documentdb/latest/developerguide/performance-insights.html) about the details.
     */
    public readonly enablePerformanceInsights!: pulumi.Output<boolean>;
    /**
     * The DNS address for this instance. May not be writable
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
     */
    public readonly engine!: pulumi.Output<string | undefined>;
    /**
     * The database engine version
     */
    public /*out*/ readonly engineVersion!: pulumi.Output<string>;
    /**
     * The identifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
     */
    public readonly identifier!: pulumi.Output<string>;
    /**
     * Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
     */
    public readonly identifierPrefix!: pulumi.Output<string>;
    /**
     * The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
     * supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
     * - db.r5.large
     * - db.r5.xlarge
     * - db.r5.2xlarge
     * - db.r5.4xlarge
     * - db.r5.12xlarge
     * - db.r5.24xlarge
     * - db.r4.large
     * - db.r4.xlarge
     * - db.r4.2xlarge
     * - db.r4.4xlarge
     * - db.r4.8xlarge
     * - db.r4.16xlarge
     * - db.t3.medium
     */
    public readonly instanceClass!: pulumi.Output<string>;
    /**
     * The ARN for the KMS encryption key if one is set to the cluster.
     */
    public /*out*/ readonly kmsKeyId!: pulumi.Output<string>;
    /**
     * The KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. If you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon DocumentDB uses your default KMS key.
     */
    public readonly performanceInsightsKmsKeyId!: pulumi.Output<string>;
    /**
     * The database port
     */
    public /*out*/ readonly port!: pulumi.Output<number>;
    /**
     * The daily time range during which automated backups are created if automated backups are enabled.
     */
    public /*out*/ readonly preferredBackupWindow!: pulumi.Output<string>;
    /**
     * The window to perform maintenance in.
     * Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
     */
    public readonly preferredMaintenanceWindow!: pulumi.Output<string>;
    /**
     * Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
     */
    public readonly promotionTier!: pulumi.Output<number | undefined>;
    public /*out*/ readonly publiclyAccessible!: pulumi.Output<boolean>;
    /**
     * Specifies whether the DB cluster is encrypted.
     */
    public /*out*/ readonly storageEncrypted!: pulumi.Output<boolean>;
    /**
     * A map of tags to assign to the instance. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
     */
    public /*out*/ readonly writer!: pulumi.Output<boolean>;

    /**
     * Create a ClusterInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterInstanceArgs | ClusterInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterInstanceState | undefined;
            resourceInputs["applyImmediately"] = state ? state.applyImmediately : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["autoMinorVersionUpgrade"] = state ? state.autoMinorVersionUpgrade : undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["caCertIdentifier"] = state ? state.caCertIdentifier : undefined;
            resourceInputs["clusterIdentifier"] = state ? state.clusterIdentifier : undefined;
            resourceInputs["dbSubnetGroupName"] = state ? state.dbSubnetGroupName : undefined;
            resourceInputs["dbiResourceId"] = state ? state.dbiResourceId : undefined;
            resourceInputs["enablePerformanceInsights"] = state ? state.enablePerformanceInsights : undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["engine"] = state ? state.engine : undefined;
            resourceInputs["engineVersion"] = state ? state.engineVersion : undefined;
            resourceInputs["identifier"] = state ? state.identifier : undefined;
            resourceInputs["identifierPrefix"] = state ? state.identifierPrefix : undefined;
            resourceInputs["instanceClass"] = state ? state.instanceClass : undefined;
            resourceInputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            resourceInputs["performanceInsightsKmsKeyId"] = state ? state.performanceInsightsKmsKeyId : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["preferredBackupWindow"] = state ? state.preferredBackupWindow : undefined;
            resourceInputs["preferredMaintenanceWindow"] = state ? state.preferredMaintenanceWindow : undefined;
            resourceInputs["promotionTier"] = state ? state.promotionTier : undefined;
            resourceInputs["publiclyAccessible"] = state ? state.publiclyAccessible : undefined;
            resourceInputs["storageEncrypted"] = state ? state.storageEncrypted : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["writer"] = state ? state.writer : undefined;
        } else {
            const args = argsOrState as ClusterInstanceArgs | undefined;
            if ((!args || args.clusterIdentifier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterIdentifier'");
            }
            if ((!args || args.instanceClass === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceClass'");
            }
            resourceInputs["applyImmediately"] = args ? args.applyImmediately : undefined;
            resourceInputs["autoMinorVersionUpgrade"] = args ? args.autoMinorVersionUpgrade : undefined;
            resourceInputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            resourceInputs["caCertIdentifier"] = args ? args.caCertIdentifier : undefined;
            resourceInputs["clusterIdentifier"] = args ? args.clusterIdentifier : undefined;
            resourceInputs["enablePerformanceInsights"] = args ? args.enablePerformanceInsights : undefined;
            resourceInputs["engine"] = args ? args.engine : undefined;
            resourceInputs["identifier"] = args ? args.identifier : undefined;
            resourceInputs["identifierPrefix"] = args ? args.identifierPrefix : undefined;
            resourceInputs["instanceClass"] = args ? args.instanceClass : undefined;
            resourceInputs["performanceInsightsKmsKeyId"] = args ? args.performanceInsightsKmsKeyId : undefined;
            resourceInputs["preferredMaintenanceWindow"] = args ? args.preferredMaintenanceWindow : undefined;
            resourceInputs["promotionTier"] = args ? args.promotionTier : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["dbSubnetGroupName"] = undefined /*out*/;
            resourceInputs["dbiResourceId"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["engineVersion"] = undefined /*out*/;
            resourceInputs["kmsKeyId"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["preferredBackupWindow"] = undefined /*out*/;
            resourceInputs["publiclyAccessible"] = undefined /*out*/;
            resourceInputs["storageEncrypted"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["writer"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ClusterInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClusterInstance resources.
 */
export interface ClusterInstanceState {
    /**
     * Specifies whether any database modifications
     * are applied immediately, or during the next maintenance window. Default is`false`.
     */
    applyImmediately?: pulumi.Input<boolean>;
    /**
     * Amazon Resource Name (ARN) of cluster instance
     */
    arn?: pulumi.Input<string>;
    /**
     * This parameter does not apply to Amazon DocumentDB. Amazon DocumentDB does not perform minor version upgrades regardless of the value set (see [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBInstance.html)). Default `true`.
     */
    autoMinorVersionUpgrade?: pulumi.Input<boolean>;
    /**
     * The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * (Optional) The identifier of the CA certificate for the DB instance.
     */
    caCertIdentifier?: pulumi.Input<string>;
    /**
     * The identifier of the `aws.docdb.Cluster` in which to launch this instance.
     */
    clusterIdentifier?: pulumi.Input<string>;
    /**
     * The DB subnet group to associate with this DB instance.
     */
    dbSubnetGroupName?: pulumi.Input<string>;
    /**
     * The region-unique, immutable identifier for the DB instance.
     */
    dbiResourceId?: pulumi.Input<string>;
    /**
     * A value that indicates whether to enable Performance Insights for the DB Instance. Default `false`. See [docs] (https://docs.aws.amazon.com/documentdb/latest/developerguide/performance-insights.html) about the details.
     */
    enablePerformanceInsights?: pulumi.Input<boolean>;
    /**
     * The DNS address for this instance. May not be writable
     */
    endpoint?: pulumi.Input<string>;
    /**
     * The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
     */
    engine?: pulumi.Input<string>;
    /**
     * The database engine version
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * The identifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
     */
    identifier?: pulumi.Input<string>;
    /**
     * Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
     */
    identifierPrefix?: pulumi.Input<string>;
    /**
     * The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
     * supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
     * - db.r5.large
     * - db.r5.xlarge
     * - db.r5.2xlarge
     * - db.r5.4xlarge
     * - db.r5.12xlarge
     * - db.r5.24xlarge
     * - db.r4.large
     * - db.r4.xlarge
     * - db.r4.2xlarge
     * - db.r4.4xlarge
     * - db.r4.8xlarge
     * - db.r4.16xlarge
     * - db.t3.medium
     */
    instanceClass?: pulumi.Input<string>;
    /**
     * The ARN for the KMS encryption key if one is set to the cluster.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * The KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. If you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon DocumentDB uses your default KMS key.
     */
    performanceInsightsKmsKeyId?: pulumi.Input<string>;
    /**
     * The database port
     */
    port?: pulumi.Input<number>;
    /**
     * The daily time range during which automated backups are created if automated backups are enabled.
     */
    preferredBackupWindow?: pulumi.Input<string>;
    /**
     * The window to perform maintenance in.
     * Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
     */
    preferredMaintenanceWindow?: pulumi.Input<string>;
    /**
     * Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
     */
    promotionTier?: pulumi.Input<number>;
    publiclyAccessible?: pulumi.Input<boolean>;
    /**
     * Specifies whether the DB cluster is encrypted.
     */
    storageEncrypted?: pulumi.Input<boolean>;
    /**
     * A map of tags to assign to the instance. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
     */
    writer?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ClusterInstance resource.
 */
export interface ClusterInstanceArgs {
    /**
     * Specifies whether any database modifications
     * are applied immediately, or during the next maintenance window. Default is`false`.
     */
    applyImmediately?: pulumi.Input<boolean>;
    /**
     * This parameter does not apply to Amazon DocumentDB. Amazon DocumentDB does not perform minor version upgrades regardless of the value set (see [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBInstance.html)). Default `true`.
     */
    autoMinorVersionUpgrade?: pulumi.Input<boolean>;
    /**
     * The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * (Optional) The identifier of the CA certificate for the DB instance.
     */
    caCertIdentifier?: pulumi.Input<string>;
    /**
     * The identifier of the `aws.docdb.Cluster` in which to launch this instance.
     */
    clusterIdentifier: pulumi.Input<string>;
    /**
     * A value that indicates whether to enable Performance Insights for the DB Instance. Default `false`. See [docs] (https://docs.aws.amazon.com/documentdb/latest/developerguide/performance-insights.html) about the details.
     */
    enablePerformanceInsights?: pulumi.Input<boolean>;
    /**
     * The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
     */
    engine?: pulumi.Input<string>;
    /**
     * The identifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
     */
    identifier?: pulumi.Input<string>;
    /**
     * Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
     */
    identifierPrefix?: pulumi.Input<string>;
    /**
     * The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
     * supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
     * - db.r5.large
     * - db.r5.xlarge
     * - db.r5.2xlarge
     * - db.r5.4xlarge
     * - db.r5.12xlarge
     * - db.r5.24xlarge
     * - db.r4.large
     * - db.r4.xlarge
     * - db.r4.2xlarge
     * - db.r4.4xlarge
     * - db.r4.8xlarge
     * - db.r4.16xlarge
     * - db.t3.medium
     */
    instanceClass: pulumi.Input<string>;
    /**
     * The KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. If you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon DocumentDB uses your default KMS key.
     */
    performanceInsightsKmsKeyId?: pulumi.Input<string>;
    /**
     * The window to perform maintenance in.
     * Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
     */
    preferredMaintenanceWindow?: pulumi.Input<string>;
    /**
     * Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
     */
    promotionTier?: pulumi.Input<number>;
    /**
     * A map of tags to assign to the instance. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
