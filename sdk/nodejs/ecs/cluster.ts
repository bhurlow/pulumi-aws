// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides an ECS cluster.
 *
 * > **NOTE on Clusters and Cluster Capacity Providers:** this provider provides both a standalone `aws.ecs.ClusterCapacityProviders` resource, as well as allowing the capacity providers and default strategies to be managed in-line by the `aws.ecs.Cluster` resource. You cannot use a Cluster with in-line capacity providers in conjunction with the Capacity Providers resource, nor use more than one Capacity Providers resource with a single Cluster, as doing so will cause a conflict and will lead to mutual overwrites.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const foo = new aws.ecs.Cluster("foo", {settings: [{
 *     name: "containerInsights",
 *     value: "enabled",
 * }]});
 * ```
 * ### Example with Log Configuration
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleKey = new aws.kms.Key("exampleKey", {
 *     description: "example",
 *     deletionWindowInDays: 7,
 * });
 * const exampleLogGroup = new aws.cloudwatch.LogGroup("exampleLogGroup", {});
 * const test = new aws.ecs.Cluster("test", {configuration: {
 *     executeCommandConfiguration: {
 *         kmsKeyId: exampleKey.arn,
 *         logging: "OVERRIDE",
 *         logConfiguration: {
 *             cloudWatchEncryptionEnabled: true,
 *             cloudWatchLogGroupName: exampleLogGroup.name,
 *         },
 *     },
 * }});
 * ```
 * ### Example with Capacity Providers
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleCluster = new aws.ecs.Cluster("exampleCluster", {});
 * const exampleCapacityProvider = new aws.ecs.CapacityProvider("exampleCapacityProvider", {autoScalingGroupProvider: {
 *     autoScalingGroupArn: aws_autoscaling_group.example.arn,
 * }});
 * const exampleClusterCapacityProviders = new aws.ecs.ClusterCapacityProviders("exampleClusterCapacityProviders", {
 *     clusterName: exampleCluster.name,
 *     capacityProviders: [exampleCapacityProvider.name],
 *     defaultCapacityProviderStrategies: [{
 *         base: 1,
 *         weight: 100,
 *         capacityProvider: exampleCapacityProvider.name,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * ECS clusters can be imported using the `name`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:ecs/cluster:Cluster stateless stateless-app
 * ```
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ecs/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * ARN that identifies the cluster.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * List of short names of one or more capacity providers to associate with the cluster. Valid values also include `FARGATE` and `FARGATE_SPOT`.
     *
     * @deprecated Use the aws_ecs_cluster_capacity_providers resource instead
     */
    public readonly capacityProviders!: pulumi.Output<string[]>;
    /**
     * The execute command configuration for the cluster. Detailed below.
     */
    public readonly configuration!: pulumi.Output<outputs.ecs.ClusterConfiguration | undefined>;
    /**
     * Configuration block for capacity provider strategy to use by default for the cluster. Can be one or more. Detailed below.
     *
     * @deprecated Use the aws_ecs_cluster_capacity_providers resource instead
     */
    public readonly defaultCapacityProviderStrategies!: pulumi.Output<outputs.ecs.ClusterDefaultCapacityProviderStrategy[]>;
    /**
     * Name of the setting to manage. Valid values: `containerInsights`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Configures a default Service Connect namespace. Detailed below.
     */
    public readonly serviceConnectDefaults!: pulumi.Output<outputs.ecs.ClusterServiceConnectDefaults | undefined>;
    /**
     * Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Detailed below.
     */
    public readonly settings!: pulumi.Output<outputs.ecs.ClusterSetting[]>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["capacityProviders"] = state ? state.capacityProviders : undefined;
            resourceInputs["configuration"] = state ? state.configuration : undefined;
            resourceInputs["defaultCapacityProviderStrategies"] = state ? state.defaultCapacityProviderStrategies : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["serviceConnectDefaults"] = state ? state.serviceConnectDefaults : undefined;
            resourceInputs["settings"] = state ? state.settings : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            resourceInputs["capacityProviders"] = args ? args.capacityProviders : undefined;
            resourceInputs["configuration"] = args ? args.configuration : undefined;
            resourceInputs["defaultCapacityProviderStrategies"] = args ? args.defaultCapacityProviderStrategies : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["serviceConnectDefaults"] = args ? args.serviceConnectDefaults : undefined;
            resourceInputs["settings"] = args ? args.settings : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * ARN that identifies the cluster.
     */
    arn?: pulumi.Input<string>;
    /**
     * List of short names of one or more capacity providers to associate with the cluster. Valid values also include `FARGATE` and `FARGATE_SPOT`.
     *
     * @deprecated Use the aws_ecs_cluster_capacity_providers resource instead
     */
    capacityProviders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The execute command configuration for the cluster. Detailed below.
     */
    configuration?: pulumi.Input<inputs.ecs.ClusterConfiguration>;
    /**
     * Configuration block for capacity provider strategy to use by default for the cluster. Can be one or more. Detailed below.
     *
     * @deprecated Use the aws_ecs_cluster_capacity_providers resource instead
     */
    defaultCapacityProviderStrategies?: pulumi.Input<pulumi.Input<inputs.ecs.ClusterDefaultCapacityProviderStrategy>[]>;
    /**
     * Name of the setting to manage. Valid values: `containerInsights`.
     */
    name?: pulumi.Input<string>;
    /**
     * Configures a default Service Connect namespace. Detailed below.
     */
    serviceConnectDefaults?: pulumi.Input<inputs.ecs.ClusterServiceConnectDefaults>;
    /**
     * Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Detailed below.
     */
    settings?: pulumi.Input<pulumi.Input<inputs.ecs.ClusterSetting>[]>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * List of short names of one or more capacity providers to associate with the cluster. Valid values also include `FARGATE` and `FARGATE_SPOT`.
     *
     * @deprecated Use the aws_ecs_cluster_capacity_providers resource instead
     */
    capacityProviders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The execute command configuration for the cluster. Detailed below.
     */
    configuration?: pulumi.Input<inputs.ecs.ClusterConfiguration>;
    /**
     * Configuration block for capacity provider strategy to use by default for the cluster. Can be one or more. Detailed below.
     *
     * @deprecated Use the aws_ecs_cluster_capacity_providers resource instead
     */
    defaultCapacityProviderStrategies?: pulumi.Input<pulumi.Input<inputs.ecs.ClusterDefaultCapacityProviderStrategy>[]>;
    /**
     * Name of the setting to manage. Valid values: `containerInsights`.
     */
    name?: pulumi.Input<string>;
    /**
     * Configures a default Service Connect namespace. Detailed below.
     */
    serviceConnectDefaults?: pulumi.Input<inputs.ecs.ClusterServiceConnectDefaults>;
    /**
     * Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Detailed below.
     */
    settings?: pulumi.Input<pulumi.Input<inputs.ecs.ClusterSetting>[]>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
