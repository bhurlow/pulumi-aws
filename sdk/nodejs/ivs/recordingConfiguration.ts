// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource for managing an AWS IVS (Interactive Video) Recording Configuration.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.ivs.RecordingConfiguration("example", {destinationConfiguration: {
 *     s3: {
 *         bucketName: "ivs-stream-archive",
 *     },
 * }});
 * ```
 *
 * ## Import
 *
 * IVS (Interactive Video) Recording Configuration can be imported using the ARN, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:ivs/recordingConfiguration:RecordingConfiguration example arn:aws:ivs:us-west-2:326937407773:recording-configuration/KAk1sHBl2L47
 * ```
 */
export class RecordingConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing RecordingConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RecordingConfigurationState, opts?: pulumi.CustomResourceOptions): RecordingConfiguration {
        return new RecordingConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ivs/recordingConfiguration:RecordingConfiguration';

    /**
     * Returns true if the given object is an instance of RecordingConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RecordingConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RecordingConfiguration.__pulumiType;
    }

    /**
     * ARN of the Recording Configuration.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Object containing destination configuration for where recorded video will be stored.
     */
    public readonly destinationConfiguration!: pulumi.Output<outputs.ivs.RecordingConfigurationDestinationConfiguration>;
    /**
     * Recording Configuration name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * If a broadcast disconnects and then reconnects within the specified interval, the multiple streams will be considered a single broadcast and merged together.
     */
    public readonly recordingReconnectWindowSeconds!: pulumi.Output<number>;
    /**
     * The current state of the Recording Configuration.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Object containing information to enable/disable the recording of thumbnails for a live session and modify the interval at which thumbnails are generated for the live session.
     */
    public readonly thumbnailConfiguration!: pulumi.Output<outputs.ivs.RecordingConfigurationThumbnailConfiguration>;

    /**
     * Create a RecordingConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RecordingConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RecordingConfigurationArgs | RecordingConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RecordingConfigurationState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["destinationConfiguration"] = state ? state.destinationConfiguration : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["recordingReconnectWindowSeconds"] = state ? state.recordingReconnectWindowSeconds : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["thumbnailConfiguration"] = state ? state.thumbnailConfiguration : undefined;
        } else {
            const args = argsOrState as RecordingConfigurationArgs | undefined;
            if ((!args || args.destinationConfiguration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationConfiguration'");
            }
            resourceInputs["destinationConfiguration"] = args ? args.destinationConfiguration : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["recordingReconnectWindowSeconds"] = args ? args.recordingReconnectWindowSeconds : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["thumbnailConfiguration"] = args ? args.thumbnailConfiguration : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RecordingConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RecordingConfiguration resources.
 */
export interface RecordingConfigurationState {
    /**
     * ARN of the Recording Configuration.
     */
    arn?: pulumi.Input<string>;
    /**
     * Object containing destination configuration for where recorded video will be stored.
     */
    destinationConfiguration?: pulumi.Input<inputs.ivs.RecordingConfigurationDestinationConfiguration>;
    /**
     * Recording Configuration name.
     */
    name?: pulumi.Input<string>;
    /**
     * If a broadcast disconnects and then reconnects within the specified interval, the multiple streams will be considered a single broadcast and merged together.
     */
    recordingReconnectWindowSeconds?: pulumi.Input<number>;
    /**
     * The current state of the Recording Configuration.
     */
    state?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Object containing information to enable/disable the recording of thumbnails for a live session and modify the interval at which thumbnails are generated for the live session.
     */
    thumbnailConfiguration?: pulumi.Input<inputs.ivs.RecordingConfigurationThumbnailConfiguration>;
}

/**
 * The set of arguments for constructing a RecordingConfiguration resource.
 */
export interface RecordingConfigurationArgs {
    /**
     * Object containing destination configuration for where recorded video will be stored.
     */
    destinationConfiguration: pulumi.Input<inputs.ivs.RecordingConfigurationDestinationConfiguration>;
    /**
     * Recording Configuration name.
     */
    name?: pulumi.Input<string>;
    /**
     * If a broadcast disconnects and then reconnects within the specified interval, the multiple streams will be considered a single broadcast and merged together.
     */
    recordingReconnectWindowSeconds?: pulumi.Input<number>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Object containing information to enable/disable the recording of thumbnails for a live session and modify the interval at which thumbnails are generated for the live session.
     */
    thumbnailConfiguration?: pulumi.Input<inputs.ivs.RecordingConfigurationThumbnailConfiguration>;
}
