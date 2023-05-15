// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a CloudTrail Event Data Store.
 *
 * More information about event data stores can be found in the [Event Data Store User Guide](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-event-data-store.html).
 *
 * > **Tip:** For an organization event data store you must create this resource in the management account.
 *
 * ## Example Usage
 * ### Basic
 *
 * The most simple event data store configuration requires us to only set the `name` attribute. The event data store will automatically capture all management events. To capture management events from all the regions, `multiRegionEnabled` must be `true`.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.cloudtrail.EventDataStore("example", {});
 * ```
 * ### Data Event Logging
 *
 * CloudTrail can log [Data Events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html) for certain services such as S3 bucket objects and Lambda function invocations. Additional information about data event configuration can be found in the following links:
 *
 * - [CloudTrail API AdvancedFieldSelector documentation](https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_AdvancedFieldSelector.html)
 * ### Log all DynamoDB PutEvent actions for a specific DynamoDB table
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const table = aws.dynamodb.getTable({
 *     name: "not-important-dynamodb-table",
 * });
 * // ... other configuration ...
 * const example = new aws.cloudtrail.EventDataStore("example", {advancedEventSelectors: [{
 *     name: "Log all DynamoDB PutEvent actions for a specific DynamoDB table",
 *     fieldSelectors: [
 *         {
 *             field: "eventCategory",
 *             equals: ["Data"],
 *         },
 *         {
 *             field: "resources.type",
 *             equals: ["AWS::DynamoDB::Table"],
 *         },
 *         {
 *             field: "eventName",
 *             equals: ["PutItem"],
 *         },
 *         {
 *             field: "resources.ARN",
 *             equals: [table.then(table => table.arn)],
 *         },
 *     ],
 * }]});
 * ```
 *
 * ## Import
 *
 * Event data stores can be imported using their `arn`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:cloudtrail/eventDataStore:EventDataStore example arn:aws:cloudtrail:us-east-1:123456789123:eventdatastore/22333815-4414-412c-b155-dd254033gfhf
 * ```
 */
export class EventDataStore extends pulumi.CustomResource {
    /**
     * Get an existing EventDataStore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventDataStoreState, opts?: pulumi.CustomResourceOptions): EventDataStore {
        return new EventDataStore(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloudtrail/eventDataStore:EventDataStore';

    /**
     * Returns true if the given object is an instance of EventDataStore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventDataStore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventDataStore.__pulumiType;
    }

    /**
     * The advanced event selectors to use to select the events for the data store. For more information about how to use advanced event selectors, see [Log events by using advanced event selectors](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced) in the CloudTrail User Guide.
     */
    public readonly advancedEventSelectors!: pulumi.Output<outputs.cloudtrail.EventDataStoreAdvancedEventSelector[]>;
    /**
     * ARN of the event data store.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Specifies the AWS KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by alias/, a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether the event data store includes events from all regions, or only from the region in which the event data store is created. Default: `true`.
     */
    public readonly multiRegionEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the event data store.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies whether an event data store collects events logged for an organization in AWS Organizations. Default: `false`.
     */
    public readonly organizationEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The retention period of the event data store, in days. You can set a retention period of up to 2555 days, the equivalent of seven years. Default: `2555`.
     */
    public readonly retentionPeriod!: pulumi.Output<number | undefined>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Specifies whether termination protection is enabled for the event data store. If termination protection is enabled, you cannot delete the event data store until termination protection is disabled. Default: `true`.
     */
    public readonly terminationProtectionEnabled!: pulumi.Output<boolean | undefined>;

    /**
     * Create a EventDataStore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EventDataStoreArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventDataStoreArgs | EventDataStoreState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EventDataStoreState | undefined;
            resourceInputs["advancedEventSelectors"] = state ? state.advancedEventSelectors : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            resourceInputs["multiRegionEnabled"] = state ? state.multiRegionEnabled : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationEnabled"] = state ? state.organizationEnabled : undefined;
            resourceInputs["retentionPeriod"] = state ? state.retentionPeriod : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["terminationProtectionEnabled"] = state ? state.terminationProtectionEnabled : undefined;
        } else {
            const args = argsOrState as EventDataStoreArgs | undefined;
            resourceInputs["advancedEventSelectors"] = args ? args.advancedEventSelectors : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["multiRegionEnabled"] = args ? args.multiRegionEnabled : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organizationEnabled"] = args ? args.organizationEnabled : undefined;
            resourceInputs["retentionPeriod"] = args ? args.retentionPeriod : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["terminationProtectionEnabled"] = args ? args.terminationProtectionEnabled : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EventDataStore.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventDataStore resources.
 */
export interface EventDataStoreState {
    /**
     * The advanced event selectors to use to select the events for the data store. For more information about how to use advanced event selectors, see [Log events by using advanced event selectors](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced) in the CloudTrail User Guide.
     */
    advancedEventSelectors?: pulumi.Input<pulumi.Input<inputs.cloudtrail.EventDataStoreAdvancedEventSelector>[]>;
    /**
     * ARN of the event data store.
     */
    arn?: pulumi.Input<string>;
    /**
     * Specifies the AWS KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by alias/, a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * Specifies whether the event data store includes events from all regions, or only from the region in which the event data store is created. Default: `true`.
     */
    multiRegionEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the event data store.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies whether an event data store collects events logged for an organization in AWS Organizations. Default: `false`.
     */
    organizationEnabled?: pulumi.Input<boolean>;
    /**
     * The retention period of the event data store, in days. You can set a retention period of up to 2555 days, the equivalent of seven years. Default: `2555`.
     */
    retentionPeriod?: pulumi.Input<number>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies whether termination protection is enabled for the event data store. If termination protection is enabled, you cannot delete the event data store until termination protection is disabled. Default: `true`.
     */
    terminationProtectionEnabled?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a EventDataStore resource.
 */
export interface EventDataStoreArgs {
    /**
     * The advanced event selectors to use to select the events for the data store. For more information about how to use advanced event selectors, see [Log events by using advanced event selectors](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced) in the CloudTrail User Guide.
     */
    advancedEventSelectors?: pulumi.Input<pulumi.Input<inputs.cloudtrail.EventDataStoreAdvancedEventSelector>[]>;
    /**
     * Specifies the AWS KMS key ID to use to encrypt the events delivered by CloudTrail. The value can be an alias name prefixed by alias/, a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * Specifies whether the event data store includes events from all regions, or only from the region in which the event data store is created. Default: `true`.
     */
    multiRegionEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the event data store.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies whether an event data store collects events logged for an organization in AWS Organizations. Default: `false`.
     */
    organizationEnabled?: pulumi.Input<boolean>;
    /**
     * The retention period of the event data store, in days. You can set a retention period of up to 2555 days, the equivalent of seven years. Default: `2555`.
     */
    retentionPeriod?: pulumi.Input<number>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies whether termination protection is enabled for the event data store. If termination protection is enabled, you cannot delete the event data store until termination protection is disabled. Default: `true`.
     */
    terminationProtectionEnabled?: pulumi.Input<boolean>;
}
