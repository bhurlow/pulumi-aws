// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an SSM Maintenance Window resource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const production = new aws.ssm.MaintenanceWindow("production", {
 *     cutoff: 1,
 *     duration: 3,
 *     schedule: "cron(0 16 ? * TUE *)",
 * });
 * ```
 *
 * ## Import
 *
 * SSM
 *
 * Maintenance Windows can be imported using the `maintenance window id`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:ssm/maintenanceWindow:MaintenanceWindow imported-window mw-0123456789
 * ```
 */
export class MaintenanceWindow extends pulumi.CustomResource {
    /**
     * Get an existing MaintenanceWindow resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MaintenanceWindowState, opts?: pulumi.CustomResourceOptions): MaintenanceWindow {
        return new MaintenanceWindow(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ssm/maintenanceWindow:MaintenanceWindow';

    /**
     * Returns true if the given object is an instance of MaintenanceWindow.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MaintenanceWindow {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MaintenanceWindow.__pulumiType;
    }

    /**
     * Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
     */
    public readonly allowUnassociatedTargets!: pulumi.Output<boolean | undefined>;
    /**
     * The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
     */
    public readonly cutoff!: pulumi.Output<number>;
    /**
     * A description for the maintenance window.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The duration of the Maintenance Window in hours.
     */
    public readonly duration!: pulumi.Output<number>;
    /**
     * Whether the maintenance window is enabled. Default: `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to no longer run the maintenance window.
     */
    public readonly endDate!: pulumi.Output<string | undefined>;
    /**
     * The name of the maintenance window.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The schedule of the Maintenance Window in the form of a [cron or rate expression](https://docs.aws.amazon.com/systems-manager/latest/userguide/reference-cron-and-rate-expressions.html).
     */
    public readonly schedule!: pulumi.Output<string>;
    /**
     * The number of days to wait after the date and time specified by a CRON expression before running the maintenance window.
     */
    public readonly scheduleOffset!: pulumi.Output<number | undefined>;
    /**
     * Timezone for schedule in [Internet Assigned Numbers Authority (IANA) Time Zone Database format](https://www.iana.org/time-zones). For example: `America/Los_Angeles`, `etc/UTC`, or `Asia/Seoul`.
     */
    public readonly scheduleTimezone!: pulumi.Output<string | undefined>;
    /**
     * Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to begin the maintenance window.
     */
    public readonly startDate!: pulumi.Output<string | undefined>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a MaintenanceWindow resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MaintenanceWindowArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MaintenanceWindowArgs | MaintenanceWindowState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MaintenanceWindowState | undefined;
            resourceInputs["allowUnassociatedTargets"] = state ? state.allowUnassociatedTargets : undefined;
            resourceInputs["cutoff"] = state ? state.cutoff : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["duration"] = state ? state.duration : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["endDate"] = state ? state.endDate : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["schedule"] = state ? state.schedule : undefined;
            resourceInputs["scheduleOffset"] = state ? state.scheduleOffset : undefined;
            resourceInputs["scheduleTimezone"] = state ? state.scheduleTimezone : undefined;
            resourceInputs["startDate"] = state ? state.startDate : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as MaintenanceWindowArgs | undefined;
            if ((!args || args.cutoff === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cutoff'");
            }
            if ((!args || args.duration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'duration'");
            }
            if ((!args || args.schedule === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schedule'");
            }
            resourceInputs["allowUnassociatedTargets"] = args ? args.allowUnassociatedTargets : undefined;
            resourceInputs["cutoff"] = args ? args.cutoff : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["duration"] = args ? args.duration : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["endDate"] = args ? args.endDate : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["schedule"] = args ? args.schedule : undefined;
            resourceInputs["scheduleOffset"] = args ? args.scheduleOffset : undefined;
            resourceInputs["scheduleTimezone"] = args ? args.scheduleTimezone : undefined;
            resourceInputs["startDate"] = args ? args.startDate : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MaintenanceWindow.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MaintenanceWindow resources.
 */
export interface MaintenanceWindowState {
    /**
     * Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
     */
    allowUnassociatedTargets?: pulumi.Input<boolean>;
    /**
     * The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
     */
    cutoff?: pulumi.Input<number>;
    /**
     * A description for the maintenance window.
     */
    description?: pulumi.Input<string>;
    /**
     * The duration of the Maintenance Window in hours.
     */
    duration?: pulumi.Input<number>;
    /**
     * Whether the maintenance window is enabled. Default: `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to no longer run the maintenance window.
     */
    endDate?: pulumi.Input<string>;
    /**
     * The name of the maintenance window.
     */
    name?: pulumi.Input<string>;
    /**
     * The schedule of the Maintenance Window in the form of a [cron or rate expression](https://docs.aws.amazon.com/systems-manager/latest/userguide/reference-cron-and-rate-expressions.html).
     */
    schedule?: pulumi.Input<string>;
    /**
     * The number of days to wait after the date and time specified by a CRON expression before running the maintenance window.
     */
    scheduleOffset?: pulumi.Input<number>;
    /**
     * Timezone for schedule in [Internet Assigned Numbers Authority (IANA) Time Zone Database format](https://www.iana.org/time-zones). For example: `America/Los_Angeles`, `etc/UTC`, or `Asia/Seoul`.
     */
    scheduleTimezone?: pulumi.Input<string>;
    /**
     * Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to begin the maintenance window.
     */
    startDate?: pulumi.Input<string>;
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
 * The set of arguments for constructing a MaintenanceWindow resource.
 */
export interface MaintenanceWindowArgs {
    /**
     * Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
     */
    allowUnassociatedTargets?: pulumi.Input<boolean>;
    /**
     * The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
     */
    cutoff: pulumi.Input<number>;
    /**
     * A description for the maintenance window.
     */
    description?: pulumi.Input<string>;
    /**
     * The duration of the Maintenance Window in hours.
     */
    duration: pulumi.Input<number>;
    /**
     * Whether the maintenance window is enabled. Default: `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to no longer run the maintenance window.
     */
    endDate?: pulumi.Input<string>;
    /**
     * The name of the maintenance window.
     */
    name?: pulumi.Input<string>;
    /**
     * The schedule of the Maintenance Window in the form of a [cron or rate expression](https://docs.aws.amazon.com/systems-manager/latest/userguide/reference-cron-and-rate-expressions.html).
     */
    schedule: pulumi.Input<string>;
    /**
     * The number of days to wait after the date and time specified by a CRON expression before running the maintenance window.
     */
    scheduleOffset?: pulumi.Input<number>;
    /**
     * Timezone for schedule in [Internet Assigned Numbers Authority (IANA) Time Zone Database format](https://www.iana.org/time-zones). For example: `America/Los_Angeles`, `etc/UTC`, or `Asia/Seoul`.
     */
    scheduleTimezone?: pulumi.Input<string>;
    /**
     * Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to begin the maintenance window.
     */
    startDate?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
