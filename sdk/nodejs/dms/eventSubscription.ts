// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a DMS (Data Migration Service) event subscription resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.dms.EventSubscription("example", {
 *     enabled: true,
 *     eventCategories: [
 *         "creation",
 *         "failure",
 *     ],
 *     snsTopicArn: aws_sns_topic.example.arn,
 *     sourceIds: [aws_dms_replication_task.example.replication_task_id],
 *     sourceType: "replication-task",
 *     tags: {
 *         Name: "example",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Event subscriptions can be imported using the `name`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:dms/eventSubscription:EventSubscription test my-awesome-event-subscription
 * ```
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
    public static readonly __pulumiType = 'aws:dms/eventSubscription:EventSubscription';

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
     * Amazon Resource Name (ARN) of the DMS Event Subscription.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Whether the event subscription should be enabled.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * List of event categories to listen for, see `DescribeEventCategories` for a canonical list.
     */
    public readonly eventCategories!: pulumi.Output<string[]>;
    /**
     * Name of event subscription.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * SNS topic arn to send events on.
     */
    public readonly snsTopicArn!: pulumi.Output<string>;
    /**
     * Ids of sources to listen to.
     */
    public readonly sourceIds!: pulumi.Output<string[] | undefined>;
    /**
     * Type of source for events. Valid values: `replication-instance` or `replication-task`
     */
    public readonly sourceType!: pulumi.Output<string | undefined>;
    /**
     * Map of resource tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
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
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EventSubscriptionState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["eventCategories"] = state ? state.eventCategories : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["snsTopicArn"] = state ? state.snsTopicArn : undefined;
            inputs["sourceIds"] = state ? state.sourceIds : undefined;
            inputs["sourceType"] = state ? state.sourceType : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as EventSubscriptionArgs | undefined;
            if ((!args || args.eventCategories === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eventCategories'");
            }
            if ((!args || args.snsTopicArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'snsTopicArn'");
            }
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["eventCategories"] = args ? args.eventCategories : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["snsTopicArn"] = args ? args.snsTopicArn : undefined;
            inputs["sourceIds"] = args ? args.sourceIds : undefined;
            inputs["sourceType"] = args ? args.sourceType : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["tagsAll"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(EventSubscription.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventSubscription resources.
 */
export interface EventSubscriptionState {
    /**
     * Amazon Resource Name (ARN) of the DMS Event Subscription.
     */
    arn?: pulumi.Input<string>;
    /**
     * Whether the event subscription should be enabled.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * List of event categories to listen for, see `DescribeEventCategories` for a canonical list.
     */
    eventCategories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of event subscription.
     */
    name?: pulumi.Input<string>;
    /**
     * SNS topic arn to send events on.
     */
    snsTopicArn?: pulumi.Input<string>;
    /**
     * Ids of sources to listen to.
     */
    sourceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Type of source for events. Valid values: `replication-instance` or `replication-task`
     */
    sourceType?: pulumi.Input<string>;
    /**
     * Map of resource tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a EventSubscription resource.
 */
export interface EventSubscriptionArgs {
    /**
     * Whether the event subscription should be enabled.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * List of event categories to listen for, see `DescribeEventCategories` for a canonical list.
     */
    eventCategories: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of event subscription.
     */
    name?: pulumi.Input<string>;
    /**
     * SNS topic arn to send events on.
     */
    snsTopicArn: pulumi.Input<string>;
    /**
     * Ids of sources to listen to.
     */
    sourceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Type of source for events. Valid values: `replication-instance` or `replication-task`
     */
    sourceType?: pulumi.Input<string>;
    /**
     * Map of resource tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
