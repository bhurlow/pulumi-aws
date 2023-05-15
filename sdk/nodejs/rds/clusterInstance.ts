// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

import {EngineType} from "./index";

/**
 * Provides an RDS Cluster Instance Resource. A Cluster Instance Resource defines
 * attributes that are specific to a single instance in a RDS Cluster,
 * specifically running Amazon Aurora.
 *
 * Unlike other RDS resources that support replication, with Amazon Aurora you do
 * not designate a primary and subsequent replicas. Instead, you simply add RDS
 * Instances and Aurora manages the replication. You can use the [count][5]
 * meta-parameter to make multiple instances and join them all to the same RDS
 * Cluster, or you may specify different Cluster Instance resources with various
 * `instanceClass` sizes.
 *
 * For more information on Amazon Aurora, see [Aurora on Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Aurora.html) in the Amazon RDS User Guide.
 *
 * > **NOTE:** Deletion Protection from the RDS service can only be enabled at the cluster level, not for individual cluster instances. You can still add the [`protect` CustomResourceOption](https://www.pulumi.com/docs/intro/concepts/programming-model/#protect) to this resource configuration if you desire protection from accidental deletion.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const _default = new aws.rds.Cluster("default", {
 *     clusterIdentifier: "aurora-cluster-demo",
 *     availabilityZones: [
 *         "us-west-2a",
 *         "us-west-2b",
 *         "us-west-2c",
 *     ],
 *     databaseName: "mydb",
 *     masterUsername: "foo",
 *     masterPassword: "barbut8chars",
 * });
 * const clusterInstances: aws.rds.ClusterInstance[] = [];
 * for (const range = {value: 0}; range.value < 2; range.value++) {
 *     clusterInstances.push(new aws.rds.ClusterInstance(`clusterInstances-${range.value}`, {
 *         identifier: `aurora-cluster-demo-${range.value}`,
 *         clusterIdentifier: _default.id,
 *         instanceClass: "db.r4.large",
 *         engine: _default.engine,
 *         engineVersion: _default.engineVersion,
 *     }));
 * }
 * ```
 *
 * ## Import
 *
 * RDS Cluster Instances can be imported using the `identifier`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:rds/clusterInstance:ClusterInstance prod_instance_1 aurora-cluster-instance-1
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
    public static readonly __pulumiType = 'aws:rds/clusterInstance:ClusterInstance';

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
     * Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
     */
    public readonly autoMinorVersionUpgrade!: pulumi.Output<boolean | undefined>;
    /**
     * The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) about the details.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * The identifier of the CA certificate for the DB instance.
     */
    public readonly caCertIdentifier!: pulumi.Output<string>;
    /**
     * The identifier of the `aws.rds.Cluster` in which to launch this instance.
     */
    public readonly clusterIdentifier!: pulumi.Output<string>;
    /**
     * Indicates whether to copy all of the user-defined tags from the DB instance to snapshots of the DB instance. Default `false`.
     */
    public readonly copyTagsToSnapshot!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the DB parameter group to associate with this instance.
     */
    public readonly dbParameterGroupName!: pulumi.Output<string>;
    /**
     * A DB subnet group to associate with this DB instance. **NOTE:** This must match the `dbSubnetGroupName` of the attached `aws.rds.Cluster`.
     */
    public readonly dbSubnetGroupName!: pulumi.Output<string>;
    /**
     * The region-unique, immutable identifier for the DB instance.
     */
    public /*out*/ readonly dbiResourceId!: pulumi.Output<string>;
    /**
     * The DNS address for this instance. May not be writable
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * The name of the database engine to be used for the RDS instance. Defaults to `aurora`. Valid Values: `aurora`, `aurora-mysql`, `aurora-postgresql`.
     * For information on the difference between the available Aurora MySQL engines
     * see [Comparison between Aurora MySQL 1 and Aurora MySQL 2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraMySQL.Updates.20180206.html)
     * in the Amazon RDS User Guide.
     */
    public readonly engine!: pulumi.Output<EngineType | undefined>;
    /**
     * The database engine version.
     */
    public readonly engineVersion!: pulumi.Output<string>;
    /**
     * The database engine version
     */
    public /*out*/ readonly engineVersionActual!: pulumi.Output<string>;
    /**
     * The indentifier for the RDS instance, if omitted, this provider will assign a random, unique identifier.
     */
    public readonly identifier!: pulumi.Output<string>;
    /**
     * Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
     */
    public readonly identifierPrefix!: pulumi.Output<string>;
    /**
     * The instance class to use. For details on CPU
     * and memory, see [Scaling Aurora DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Aurora.Managing.html). Aurora uses `db.*` instance classes/types. Please see [AWS Documentation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) for currently available instance classes and complete details.
     */
    public readonly instanceClass!: pulumi.Output<string>;
    /**
     * The ARN for the KMS encryption key if one is set to the cluster.
     */
    public /*out*/ readonly kmsKeyId!: pulumi.Output<string>;
    /**
     * The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collecting Enhanced Monitoring metrics, specify 0. The default is 0. Valid Values: 0, 1, 5, 10, 15, 30, 60.
     */
    public readonly monitoringInterval!: pulumi.Output<number | undefined>;
    /**
     * The ARN for the IAM role that permits RDS to send
     * enhanced monitoring metrics to CloudWatch Logs. You can find more information on the [AWS Documentation](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.html)
     * what IAM permissions are needed to allow Enhanced Monitoring for RDS Instances.
     */
    public readonly monitoringRoleArn!: pulumi.Output<string>;
    /**
     * The network type of the DB instance.
     */
    public /*out*/ readonly networkType!: pulumi.Output<string>;
    /**
     * Specifies whether Performance Insights is enabled or not.
     */
    public readonly performanceInsightsEnabled!: pulumi.Output<boolean>;
    /**
     * ARN for the KMS key to encrypt Performance Insights data. When specifying `performanceInsightsKmsKeyId`, `performanceInsightsEnabled` needs to be set to true.
     */
    public readonly performanceInsightsKmsKeyId!: pulumi.Output<string>;
    /**
     * Amount of time in days to retain Performance Insights data. Valid values are `7`, `731` (2 years) or a multiple of `31`. When specifying `performanceInsightsRetentionPeriod`, `performanceInsightsEnabled` needs to be set to true. Defaults to '7'.
     */
    public readonly performanceInsightsRetentionPeriod!: pulumi.Output<number>;
    /**
     * The database port
     */
    public /*out*/ readonly port!: pulumi.Output<number>;
    /**
     * The daily time range during which automated backups are created if automated backups are enabled. Eg: "04:00-09:00". **NOTE:** If `preferredBackupWindow` is set at the cluster level, this argument **must** be omitted.
     */
    public readonly preferredBackupWindow!: pulumi.Output<string>;
    /**
     * The window to perform maintenance in.
     * Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
     */
    public readonly preferredMaintenanceWindow!: pulumi.Output<string>;
    /**
     * Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoted to writer.
     */
    public readonly promotionTier!: pulumi.Output<number | undefined>;
    /**
     * Bool to control if instance is publicly accessible.
     * Default `false`. See the documentation on [Creating DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) for more
     * details on controlling this property.
     */
    public readonly publiclyAccessible!: pulumi.Output<boolean | undefined>;
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
            resourceInputs["copyTagsToSnapshot"] = state ? state.copyTagsToSnapshot : undefined;
            resourceInputs["dbParameterGroupName"] = state ? state.dbParameterGroupName : undefined;
            resourceInputs["dbSubnetGroupName"] = state ? state.dbSubnetGroupName : undefined;
            resourceInputs["dbiResourceId"] = state ? state.dbiResourceId : undefined;
            resourceInputs["endpoint"] = state ? state.endpoint : undefined;
            resourceInputs["engine"] = state ? state.engine : undefined;
            resourceInputs["engineVersion"] = state ? state.engineVersion : undefined;
            resourceInputs["engineVersionActual"] = state ? state.engineVersionActual : undefined;
            resourceInputs["identifier"] = state ? state.identifier : undefined;
            resourceInputs["identifierPrefix"] = state ? state.identifierPrefix : undefined;
            resourceInputs["instanceClass"] = state ? state.instanceClass : undefined;
            resourceInputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            resourceInputs["monitoringInterval"] = state ? state.monitoringInterval : undefined;
            resourceInputs["monitoringRoleArn"] = state ? state.monitoringRoleArn : undefined;
            resourceInputs["networkType"] = state ? state.networkType : undefined;
            resourceInputs["performanceInsightsEnabled"] = state ? state.performanceInsightsEnabled : undefined;
            resourceInputs["performanceInsightsKmsKeyId"] = state ? state.performanceInsightsKmsKeyId : undefined;
            resourceInputs["performanceInsightsRetentionPeriod"] = state ? state.performanceInsightsRetentionPeriod : undefined;
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
            resourceInputs["copyTagsToSnapshot"] = args ? args.copyTagsToSnapshot : undefined;
            resourceInputs["dbParameterGroupName"] = args ? args.dbParameterGroupName : undefined;
            resourceInputs["dbSubnetGroupName"] = args ? args.dbSubnetGroupName : undefined;
            resourceInputs["engine"] = args ? args.engine : undefined;
            resourceInputs["engineVersion"] = args ? args.engineVersion : undefined;
            resourceInputs["identifier"] = args ? args.identifier : undefined;
            resourceInputs["identifierPrefix"] = args ? args.identifierPrefix : undefined;
            resourceInputs["instanceClass"] = args ? args.instanceClass : undefined;
            resourceInputs["monitoringInterval"] = args ? args.monitoringInterval : undefined;
            resourceInputs["monitoringRoleArn"] = args ? args.monitoringRoleArn : undefined;
            resourceInputs["performanceInsightsEnabled"] = args ? args.performanceInsightsEnabled : undefined;
            resourceInputs["performanceInsightsKmsKeyId"] = args ? args.performanceInsightsKmsKeyId : undefined;
            resourceInputs["performanceInsightsRetentionPeriod"] = args ? args.performanceInsightsRetentionPeriod : undefined;
            resourceInputs["preferredBackupWindow"] = args ? args.preferredBackupWindow : undefined;
            resourceInputs["preferredMaintenanceWindow"] = args ? args.preferredMaintenanceWindow : undefined;
            resourceInputs["promotionTier"] = args ? args.promotionTier : undefined;
            resourceInputs["publiclyAccessible"] = args ? args.publiclyAccessible : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["dbiResourceId"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["engineVersionActual"] = undefined /*out*/;
            resourceInputs["kmsKeyId"] = undefined /*out*/;
            resourceInputs["networkType"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
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
     * Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
     */
    autoMinorVersionUpgrade?: pulumi.Input<boolean>;
    /**
     * The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) about the details.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The identifier of the CA certificate for the DB instance.
     */
    caCertIdentifier?: pulumi.Input<string>;
    /**
     * The identifier of the `aws.rds.Cluster` in which to launch this instance.
     */
    clusterIdentifier?: pulumi.Input<string>;
    /**
     * Indicates whether to copy all of the user-defined tags from the DB instance to snapshots of the DB instance. Default `false`.
     */
    copyTagsToSnapshot?: pulumi.Input<boolean>;
    /**
     * The name of the DB parameter group to associate with this instance.
     */
    dbParameterGroupName?: pulumi.Input<string>;
    /**
     * A DB subnet group to associate with this DB instance. **NOTE:** This must match the `dbSubnetGroupName` of the attached `aws.rds.Cluster`.
     */
    dbSubnetGroupName?: pulumi.Input<string>;
    /**
     * The region-unique, immutable identifier for the DB instance.
     */
    dbiResourceId?: pulumi.Input<string>;
    /**
     * The DNS address for this instance. May not be writable
     */
    endpoint?: pulumi.Input<string>;
    /**
     * The name of the database engine to be used for the RDS instance. Defaults to `aurora`. Valid Values: `aurora`, `aurora-mysql`, `aurora-postgresql`.
     * For information on the difference between the available Aurora MySQL engines
     * see [Comparison between Aurora MySQL 1 and Aurora MySQL 2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraMySQL.Updates.20180206.html)
     * in the Amazon RDS User Guide.
     */
    engine?: pulumi.Input<EngineType>;
    /**
     * The database engine version.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * The database engine version
     */
    engineVersionActual?: pulumi.Input<string>;
    /**
     * The indentifier for the RDS instance, if omitted, this provider will assign a random, unique identifier.
     */
    identifier?: pulumi.Input<string>;
    /**
     * Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
     */
    identifierPrefix?: pulumi.Input<string>;
    /**
     * The instance class to use. For details on CPU
     * and memory, see [Scaling Aurora DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Aurora.Managing.html). Aurora uses `db.*` instance classes/types. Please see [AWS Documentation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) for currently available instance classes and complete details.
     */
    instanceClass?: pulumi.Input<string | enums.rds.InstanceType>;
    /**
     * The ARN for the KMS encryption key if one is set to the cluster.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collecting Enhanced Monitoring metrics, specify 0. The default is 0. Valid Values: 0, 1, 5, 10, 15, 30, 60.
     */
    monitoringInterval?: pulumi.Input<number>;
    /**
     * The ARN for the IAM role that permits RDS to send
     * enhanced monitoring metrics to CloudWatch Logs. You can find more information on the [AWS Documentation](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.html)
     * what IAM permissions are needed to allow Enhanced Monitoring for RDS Instances.
     */
    monitoringRoleArn?: pulumi.Input<string>;
    /**
     * The network type of the DB instance.
     */
    networkType?: pulumi.Input<string>;
    /**
     * Specifies whether Performance Insights is enabled or not.
     */
    performanceInsightsEnabled?: pulumi.Input<boolean>;
    /**
     * ARN for the KMS key to encrypt Performance Insights data. When specifying `performanceInsightsKmsKeyId`, `performanceInsightsEnabled` needs to be set to true.
     */
    performanceInsightsKmsKeyId?: pulumi.Input<string>;
    /**
     * Amount of time in days to retain Performance Insights data. Valid values are `7`, `731` (2 years) or a multiple of `31`. When specifying `performanceInsightsRetentionPeriod`, `performanceInsightsEnabled` needs to be set to true. Defaults to '7'.
     */
    performanceInsightsRetentionPeriod?: pulumi.Input<number>;
    /**
     * The database port
     */
    port?: pulumi.Input<number>;
    /**
     * The daily time range during which automated backups are created if automated backups are enabled. Eg: "04:00-09:00". **NOTE:** If `preferredBackupWindow` is set at the cluster level, this argument **must** be omitted.
     */
    preferredBackupWindow?: pulumi.Input<string>;
    /**
     * The window to perform maintenance in.
     * Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
     */
    preferredMaintenanceWindow?: pulumi.Input<string>;
    /**
     * Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoted to writer.
     */
    promotionTier?: pulumi.Input<number>;
    /**
     * Bool to control if instance is publicly accessible.
     * Default `false`. See the documentation on [Creating DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) for more
     * details on controlling this property.
     */
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
     * Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
     */
    autoMinorVersionUpgrade?: pulumi.Input<boolean>;
    /**
     * The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) about the details.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The identifier of the CA certificate for the DB instance.
     */
    caCertIdentifier?: pulumi.Input<string>;
    /**
     * The identifier of the `aws.rds.Cluster` in which to launch this instance.
     */
    clusterIdentifier: pulumi.Input<string>;
    /**
     * Indicates whether to copy all of the user-defined tags from the DB instance to snapshots of the DB instance. Default `false`.
     */
    copyTagsToSnapshot?: pulumi.Input<boolean>;
    /**
     * The name of the DB parameter group to associate with this instance.
     */
    dbParameterGroupName?: pulumi.Input<string>;
    /**
     * A DB subnet group to associate with this DB instance. **NOTE:** This must match the `dbSubnetGroupName` of the attached `aws.rds.Cluster`.
     */
    dbSubnetGroupName?: pulumi.Input<string>;
    /**
     * The name of the database engine to be used for the RDS instance. Defaults to `aurora`. Valid Values: `aurora`, `aurora-mysql`, `aurora-postgresql`.
     * For information on the difference between the available Aurora MySQL engines
     * see [Comparison between Aurora MySQL 1 and Aurora MySQL 2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraMySQL.Updates.20180206.html)
     * in the Amazon RDS User Guide.
     */
    engine?: pulumi.Input<EngineType>;
    /**
     * The database engine version.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * The indentifier for the RDS instance, if omitted, this provider will assign a random, unique identifier.
     */
    identifier?: pulumi.Input<string>;
    /**
     * Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
     */
    identifierPrefix?: pulumi.Input<string>;
    /**
     * The instance class to use. For details on CPU
     * and memory, see [Scaling Aurora DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Aurora.Managing.html). Aurora uses `db.*` instance classes/types. Please see [AWS Documentation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) for currently available instance classes and complete details.
     */
    instanceClass: pulumi.Input<string | enums.rds.InstanceType>;
    /**
     * The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collecting Enhanced Monitoring metrics, specify 0. The default is 0. Valid Values: 0, 1, 5, 10, 15, 30, 60.
     */
    monitoringInterval?: pulumi.Input<number>;
    /**
     * The ARN for the IAM role that permits RDS to send
     * enhanced monitoring metrics to CloudWatch Logs. You can find more information on the [AWS Documentation](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.html)
     * what IAM permissions are needed to allow Enhanced Monitoring for RDS Instances.
     */
    monitoringRoleArn?: pulumi.Input<string>;
    /**
     * Specifies whether Performance Insights is enabled or not.
     */
    performanceInsightsEnabled?: pulumi.Input<boolean>;
    /**
     * ARN for the KMS key to encrypt Performance Insights data. When specifying `performanceInsightsKmsKeyId`, `performanceInsightsEnabled` needs to be set to true.
     */
    performanceInsightsKmsKeyId?: pulumi.Input<string>;
    /**
     * Amount of time in days to retain Performance Insights data. Valid values are `7`, `731` (2 years) or a multiple of `31`. When specifying `performanceInsightsRetentionPeriod`, `performanceInsightsEnabled` needs to be set to true. Defaults to '7'.
     */
    performanceInsightsRetentionPeriod?: pulumi.Input<number>;
    /**
     * The daily time range during which automated backups are created if automated backups are enabled. Eg: "04:00-09:00". **NOTE:** If `preferredBackupWindow` is set at the cluster level, this argument **must** be omitted.
     */
    preferredBackupWindow?: pulumi.Input<string>;
    /**
     * The window to perform maintenance in.
     * Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
     */
    preferredMaintenanceWindow?: pulumi.Input<string>;
    /**
     * Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoted to writer.
     */
    promotionTier?: pulumi.Input<number>;
    /**
     * Bool to control if instance is publicly accessible.
     * Default `false`. See the documentation on [Creating DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) for more
     * details on controlling this property.
     */
    publiclyAccessible?: pulumi.Input<boolean>;
    /**
     * A map of tags to assign to the instance. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
