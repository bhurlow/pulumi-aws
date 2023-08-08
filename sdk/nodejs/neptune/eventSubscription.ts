// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const defaultCluster = new aws.neptune.Cluster("defaultCluster", {
 *     clusterIdentifier: "neptune-cluster-demo",
 *     engine: "neptune",
 *     backupRetentionPeriod: 5,
 *     preferredBackupWindow: "07:00-09:00",
 *     skipFinalSnapshot: true,
 *     iamDatabaseAuthenticationEnabled: true,
 *     applyImmediately: true,
 * });
 * const example = new aws.neptune.ClusterInstance("example", {
 *     clusterIdentifier: defaultCluster.id,
 *     engine: "neptune",
 *     instanceClass: "db.r4.large",
 *     applyImmediately: true,
 * });
 * const defaultTopic = new aws.sns.Topic("defaultTopic", {});
 * const defaultEventSubscription = new aws.neptune.EventSubscription("defaultEventSubscription", {
 *     snsTopicArn: defaultTopic.arn,
 *     sourceType: "db-instance",
 *     sourceIds: [example.id],
 *     eventCategories: [
 *         "maintenance",
 *         "availability",
 *         "creation",
 *         "backup",
 *         "restoration",
 *         "recovery",
 *         "deletion",
 *         "failover",
 *         "failure",
 *         "notification",
 *         "configuration change",
 *         "read replica",
 *     ],
 *     tags: {
 *         env: "test",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_neptune_event_subscription.example
 *
 *  id = "my-event-subscription" } Using `pulumi import`, import `aws_neptune_event_subscription` using the event subscription name. For exampleconsole % pulumi import aws_neptune_event_subscription.example my-event-subscription
 */
export class EventSubscription extends pulumi.CustomResource {
    /**
     * Get an existing EventSubscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventSubscriptionState, opts?: pulumi.CustomResourceOptions): EventSubscription {
        return new EventSubscription(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:neptune/eventSubscription:EventSubscription';

    /**
     * Returns true if the given object is an instance of EventSubscription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventSubscription {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventSubscription.__pulumiType;
    }

    /**
     * The Amazon Resource Name of the Neptune event notification subscription.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The AWS customer account associated with the Neptune event notification subscription.
     */
    public /*out*/ readonly customerAwsId!: pulumi.Output<string>;
    /**
     * A boolean flag to enable/disable the subscription. Defaults to true.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * A list of event categories for a `sourceType` that you want to subscribe to. Run `aws neptune describe-event-categories` to find all the event categories.
     */
    public readonly eventCategories!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the Neptune event subscription. By default generated by this provider.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Neptune event subscription. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string>;
    /**
     * The ARN of the SNS topic to send events to.
     */
    public readonly snsTopicArn!: pulumi.Output<string>;
    /**
     * A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `sourceType` must also be specified.
     */
    public readonly sourceIds!: pulumi.Output<string[] | undefined>;
    /**
     * The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster` or `db-cluster-snapshot`. If not set, all sources will be subscribed to.
     */
    public readonly sourceType!: pulumi.Output<string | undefined>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a EventSubscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventSubscriptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventSubscriptionArgs | EventSubscriptionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EventSubscriptionState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["customerAwsId"] = state ? state.customerAwsId : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["eventCategories"] = state ? state.eventCategories : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namePrefix"] = state ? state.namePrefix : undefined;
            resourceInputs["snsTopicArn"] = state ? state.snsTopicArn : undefined;
            resourceInputs["sourceIds"] = state ? state.sourceIds : undefined;
            resourceInputs["sourceType"] = state ? state.sourceType : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as EventSubscriptionArgs | undefined;
            if ((!args || args.snsTopicArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'snsTopicArn'");
            }
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["eventCategories"] = args ? args.eventCategories : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namePrefix"] = args ? args.namePrefix : undefined;
            resourceInputs["snsTopicArn"] = args ? args.snsTopicArn : undefined;
            resourceInputs["sourceIds"] = args ? args.sourceIds : undefined;
            resourceInputs["sourceType"] = args ? args.sourceType : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["customerAwsId"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EventSubscription.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventSubscription resources.
 */
export interface EventSubscriptionState {
    /**
     * The Amazon Resource Name of the Neptune event notification subscription.
     */
    arn?: pulumi.Input<string>;
    /**
     * The AWS customer account associated with the Neptune event notification subscription.
     */
    customerAwsId?: pulumi.Input<string>;
    /**
     * A boolean flag to enable/disable the subscription. Defaults to true.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * A list of event categories for a `sourceType` that you want to subscribe to. Run `aws neptune describe-event-categories` to find all the event categories.
     */
    eventCategories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the Neptune event subscription. By default generated by this provider.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the Neptune event subscription. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The ARN of the SNS topic to send events to.
     */
    snsTopicArn?: pulumi.Input<string>;
    /**
     * A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `sourceType` must also be specified.
     */
    sourceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster` or `db-cluster-snapshot`. If not set, all sources will be subscribed to.
     */
    sourceType?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a EventSubscription resource.
 */
export interface EventSubscriptionArgs {
    /**
     * A boolean flag to enable/disable the subscription. Defaults to true.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * A list of event categories for a `sourceType` that you want to subscribe to. Run `aws neptune describe-event-categories` to find all the event categories.
     */
    eventCategories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the Neptune event subscription. By default generated by this provider.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the Neptune event subscription. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The ARN of the SNS topic to send events to.
     */
    snsTopicArn: pulumi.Input<string>;
    /**
     * A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `sourceType` must also be specified.
     */
    sourceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster` or `db-cluster-snapshot`. If not set, all sources will be subscribed to.
     */
    sourceType?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
