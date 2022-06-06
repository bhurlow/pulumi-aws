// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage a GuardDuty detector.
 *
 * > **NOTE:** Deleting this resource is equivalent to "disabling" GuardDuty for an AWS region, which removes all existing findings. You can set the `enable` attribute to `false` to instead "suspend" monitoring and feedback reporting while keeping existing data. See the [Suspending or Disabling Amazon GuardDuty documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_suspend-disable.html) for more information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const myDetector = new aws.guardduty.Detector("MyDetector", {
 *     datasources: {
 *         kubernetes: {
 *             auditLogs: {
 *                 enable: false,
 *             },
 *         },
 *         s3Logs: {
 *             enable: true,
 *         },
 *     },
 *     enable: true,
 * });
 * ```
 *
 * ## Import
 *
 * GuardDuty detectors can be imported using the detector ID, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:guardduty/detector:Detector MyDetector 00b00fd5aecc0ab60a708659477e9617
 * ```
 */
export class Detector extends pulumi.CustomResource {
    /**
     * Get an existing Detector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DetectorState, opts?: pulumi.CustomResourceOptions): Detector {
        return new Detector(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:guardduty/detector:Detector';

    /**
     * Returns true if the given object is an instance of Detector.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Detector {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Detector.__pulumiType;
    }

    /**
     * The AWS account ID of the GuardDuty detector
     */
    public /*out*/ readonly accountId!: pulumi.Output<string>;
    /**
     * Amazon Resource Name (ARN) of the GuardDuty detector
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Describes which data sources will be enabled for the detector. See Data Sources below for more details.
     */
    public readonly datasources!: pulumi.Output<outputs.guardduty.DetectorDatasources>;
    /**
     * If true, enables [S3 protection](https://docs.aws.amazon.com/guardduty/latest/ug/s3-protection.html).
     * Defaults to `true`.
     */
    public readonly enable!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty primary account and cannot be modified, otherwise defaults to `SIX_HOURS`. For standalone and GuardDuty primary accounts, it must be configured in this provider to enable drift detection. Valid values for standalone and primary accounts: `FIFTEEN_MINUTES`, `ONE_HOUR`, `SIX_HOURS`. See [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html#guardduty_findings_cloudwatch_notification_frequency) for more information.
     */
    public readonly findingPublishingFrequency!: pulumi.Output<string>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a Detector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DetectorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DetectorArgs | DetectorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DetectorState | undefined;
            resourceInputs["accountId"] = state ? state.accountId : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["datasources"] = state ? state.datasources : undefined;
            resourceInputs["enable"] = state ? state.enable : undefined;
            resourceInputs["findingPublishingFrequency"] = state ? state.findingPublishingFrequency : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as DetectorArgs | undefined;
            resourceInputs["datasources"] = args ? args.datasources : undefined;
            resourceInputs["enable"] = args ? args.enable : undefined;
            resourceInputs["findingPublishingFrequency"] = args ? args.findingPublishingFrequency : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["accountId"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Detector.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Detector resources.
 */
export interface DetectorState {
    /**
     * The AWS account ID of the GuardDuty detector
     */
    accountId?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the GuardDuty detector
     */
    arn?: pulumi.Input<string>;
    /**
     * Describes which data sources will be enabled for the detector. See Data Sources below for more details.
     */
    datasources?: pulumi.Input<inputs.guardduty.DetectorDatasources>;
    /**
     * If true, enables [S3 protection](https://docs.aws.amazon.com/guardduty/latest/ug/s3-protection.html).
     * Defaults to `true`.
     */
    enable?: pulumi.Input<boolean>;
    /**
     * Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty primary account and cannot be modified, otherwise defaults to `SIX_HOURS`. For standalone and GuardDuty primary accounts, it must be configured in this provider to enable drift detection. Valid values for standalone and primary accounts: `FIFTEEN_MINUTES`, `ONE_HOUR`, `SIX_HOURS`. See [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html#guardduty_findings_cloudwatch_notification_frequency) for more information.
     */
    findingPublishingFrequency?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Detector resource.
 */
export interface DetectorArgs {
    /**
     * Describes which data sources will be enabled for the detector. See Data Sources below for more details.
     */
    datasources?: pulumi.Input<inputs.guardduty.DetectorDatasources>;
    /**
     * If true, enables [S3 protection](https://docs.aws.amazon.com/guardduty/latest/ug/s3-protection.html).
     * Defaults to `true`.
     */
    enable?: pulumi.Input<boolean>;
    /**
     * Specifies the frequency of notifications sent for subsequent finding occurrences. If the detector is a GuardDuty member account, the value is determined by the GuardDuty primary account and cannot be modified, otherwise defaults to `SIX_HOURS`. For standalone and GuardDuty primary accounts, it must be configured in this provider to enable drift detection. Valid values for standalone and primary accounts: `FIFTEEN_MINUTES`, `ONE_HOUR`, `SIX_HOURS`. See [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings_cloudwatch.html#guardduty_findings_cloudwatch_notification_frequency) for more information.
     */
    findingPublishingFrequency?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
