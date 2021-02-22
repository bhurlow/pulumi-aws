// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an ElastiCache Global Replication Group resource, which manage a replication between 2 or more redis replication group in different regions.
 *
 * ## Example Usage
 * ### Global replication group with a single instance redis replication group
 *
 * To create a single shard primary with single read replica:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const primary = new aws.elasticache.ReplicationGroup("primary", {
 *     replicationGroupDescription: "test example",
 *     engine: "redis",
 *     engineVersion: "5.0.6",
 *     nodeType: "cache.m5.large",
 *     numberCacheClusters: 1,
 * });
 * const replicationGroup = new aws.elasticache.GlobalReplicationGroup("replicationGroup", {
 *     globalReplicationGroupIdSuffix: "example",
 *     primaryReplicationGroupId: primary.id,
 * });
 * ```
 *
 * ## Import
 *
 * ElastiCache Global Replication Groups can be imported using the `global_replication_group_id`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:elasticache/globalReplicationGroup:GlobalReplicationGroup my_global_replication_group okuqm-global-replication-group-1
 * ```
 */
export class GlobalReplicationGroup extends pulumi.CustomResource {
    /**
     * Get an existing GlobalReplicationGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GlobalReplicationGroupState, opts?: pulumi.CustomResourceOptions): GlobalReplicationGroup {
        return new GlobalReplicationGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:elasticache/globalReplicationGroup:GlobalReplicationGroup';

    /**
     * Returns true if the given object is an instance of GlobalReplicationGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GlobalReplicationGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GlobalReplicationGroup.__pulumiType;
    }

    /**
     * The full version number of the cache engine running on the members of this global replication group.
     */
    public /*out*/ readonly actualEngineVersion!: pulumi.Output<string>;
    /**
     * The ARN of the ElastiCache Global Replication Group.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A flag that indicate whether the encryption at rest is enabled.
     */
    public /*out*/ readonly atRestEncryptionEnabled!: pulumi.Output<boolean>;
    /**
     * A flag that indicate whether AuthToken (password) is enabled.
     */
    public /*out*/ readonly authTokenEnabled!: pulumi.Output<boolean>;
    /**
     * The instance class used. See AWS documentation for information on [supported node types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html) and [guidance on selecting node types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/nodes-select-size.html).
     */
    public /*out*/ readonly cacheNodeType!: pulumi.Output<string>;
    /**
     * Indicates whether the Global Datastore is cluster enabled.
     */
    public /*out*/ readonly clusterEnabled!: pulumi.Output<boolean>;
    /**
     * The name of the cache engine to be used for the clusters in this global replication group.
     */
    public /*out*/ readonly engine!: pulumi.Output<string>;
    /**
     * A user-created description for the global replication group.
     */
    public readonly globalReplicationGroupDescription!: pulumi.Output<string | undefined>;
    /**
     * The full ID of the global replication group.
     */
    public /*out*/ readonly globalReplicationGroupId!: pulumi.Output<string>;
    /**
     * The suffix name of a Global Datastore. If `globalReplicationGroupIdSuffix` is changed, creates a new resource.
     */
    public readonly globalReplicationGroupIdSuffix!: pulumi.Output<string>;
    /**
     * The ID of the primary cluster that accepts writes and will replicate updates to the secondary cluster. If `primaryReplicationGroupId` is changed, creates a new resource.
     */
    public readonly primaryReplicationGroupId!: pulumi.Output<string>;
    /**
     * A flag that indicates whether the encryption in transit is enabled.
     */
    public /*out*/ readonly transitEncryptionEnabled!: pulumi.Output<boolean>;

    /**
     * Create a GlobalReplicationGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GlobalReplicationGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GlobalReplicationGroupArgs | GlobalReplicationGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GlobalReplicationGroupState | undefined;
            inputs["actualEngineVersion"] = state ? state.actualEngineVersion : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["atRestEncryptionEnabled"] = state ? state.atRestEncryptionEnabled : undefined;
            inputs["authTokenEnabled"] = state ? state.authTokenEnabled : undefined;
            inputs["cacheNodeType"] = state ? state.cacheNodeType : undefined;
            inputs["clusterEnabled"] = state ? state.clusterEnabled : undefined;
            inputs["engine"] = state ? state.engine : undefined;
            inputs["globalReplicationGroupDescription"] = state ? state.globalReplicationGroupDescription : undefined;
            inputs["globalReplicationGroupId"] = state ? state.globalReplicationGroupId : undefined;
            inputs["globalReplicationGroupIdSuffix"] = state ? state.globalReplicationGroupIdSuffix : undefined;
            inputs["primaryReplicationGroupId"] = state ? state.primaryReplicationGroupId : undefined;
            inputs["transitEncryptionEnabled"] = state ? state.transitEncryptionEnabled : undefined;
        } else {
            const args = argsOrState as GlobalReplicationGroupArgs | undefined;
            if ((!args || args.globalReplicationGroupIdSuffix === undefined) && !opts.urn) {
                throw new Error("Missing required property 'globalReplicationGroupIdSuffix'");
            }
            if ((!args || args.primaryReplicationGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'primaryReplicationGroupId'");
            }
            inputs["globalReplicationGroupDescription"] = args ? args.globalReplicationGroupDescription : undefined;
            inputs["globalReplicationGroupIdSuffix"] = args ? args.globalReplicationGroupIdSuffix : undefined;
            inputs["primaryReplicationGroupId"] = args ? args.primaryReplicationGroupId : undefined;
            inputs["actualEngineVersion"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
            inputs["atRestEncryptionEnabled"] = undefined /*out*/;
            inputs["authTokenEnabled"] = undefined /*out*/;
            inputs["cacheNodeType"] = undefined /*out*/;
            inputs["clusterEnabled"] = undefined /*out*/;
            inputs["engine"] = undefined /*out*/;
            inputs["globalReplicationGroupId"] = undefined /*out*/;
            inputs["transitEncryptionEnabled"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(GlobalReplicationGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GlobalReplicationGroup resources.
 */
export interface GlobalReplicationGroupState {
    /**
     * The full version number of the cache engine running on the members of this global replication group.
     */
    readonly actualEngineVersion?: pulumi.Input<string>;
    /**
     * The ARN of the ElastiCache Global Replication Group.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * A flag that indicate whether the encryption at rest is enabled.
     */
    readonly atRestEncryptionEnabled?: pulumi.Input<boolean>;
    /**
     * A flag that indicate whether AuthToken (password) is enabled.
     */
    readonly authTokenEnabled?: pulumi.Input<boolean>;
    /**
     * The instance class used. See AWS documentation for information on [supported node types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html) and [guidance on selecting node types](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/nodes-select-size.html).
     */
    readonly cacheNodeType?: pulumi.Input<string>;
    /**
     * Indicates whether the Global Datastore is cluster enabled.
     */
    readonly clusterEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the cache engine to be used for the clusters in this global replication group.
     */
    readonly engine?: pulumi.Input<string>;
    /**
     * A user-created description for the global replication group.
     */
    readonly globalReplicationGroupDescription?: pulumi.Input<string>;
    /**
     * The full ID of the global replication group.
     */
    readonly globalReplicationGroupId?: pulumi.Input<string>;
    /**
     * The suffix name of a Global Datastore. If `globalReplicationGroupIdSuffix` is changed, creates a new resource.
     */
    readonly globalReplicationGroupIdSuffix?: pulumi.Input<string>;
    /**
     * The ID of the primary cluster that accepts writes and will replicate updates to the secondary cluster. If `primaryReplicationGroupId` is changed, creates a new resource.
     */
    readonly primaryReplicationGroupId?: pulumi.Input<string>;
    /**
     * A flag that indicates whether the encryption in transit is enabled.
     */
    readonly transitEncryptionEnabled?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a GlobalReplicationGroup resource.
 */
export interface GlobalReplicationGroupArgs {
    /**
     * A user-created description for the global replication group.
     */
    readonly globalReplicationGroupDescription?: pulumi.Input<string>;
    /**
     * The suffix name of a Global Datastore. If `globalReplicationGroupIdSuffix` is changed, creates a new resource.
     */
    readonly globalReplicationGroupIdSuffix: pulumi.Input<string>;
    /**
     * The ID of the primary cluster that accepts writes and will replicate updates to the secondary cluster. If `primaryReplicationGroupId` is changed, creates a new resource.
     */
    readonly primaryReplicationGroupId: pulumi.Input<string>;
}
