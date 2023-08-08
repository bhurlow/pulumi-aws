// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage AWS Device Farm Network Profiles.
 * ∂
 * > **NOTE:** AWS currently has limited regional support for Device Farm (e.g., `us-west-2`). See [AWS Device Farm endpoints and quotas](https://docs.aws.amazon.com/general/latest/gr/devicefarm.html) for information on supported regions.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleProject = new aws.devicefarm.Project("exampleProject", {});
 * const exampleNetworkProfile = new aws.devicefarm.NetworkProfile("exampleNetworkProfile", {projectArn: exampleProject.arn});
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_devicefarm_network_profile.example
 *
 *  id = "arn:aws:devicefarm:us-west-2:123456789012:networkprofile:4fa784c7-ccb4-4dbf-ba4f-02198320daa1" } Using `pulumi import`, import DeviceFarm Network Profiles using their ARN. For exampleconsole % pulumi import aws_devicefarm_network_profile.example arn:aws:devicefarm:us-west-2:123456789012:networkprofile:4fa784c7-ccb4-4dbf-ba4f-02198320daa1
 */
export class NetworkProfile extends pulumi.CustomResource {
    /**
     * Get an existing NetworkProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkProfileState, opts?: pulumi.CustomResourceOptions): NetworkProfile {
        return new NetworkProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:devicefarm/networkProfile:NetworkProfile';

    /**
     * Returns true if the given object is an instance of NetworkProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkProfile.__pulumiType;
    }

    /**
     * The Amazon Resource Name of this network profile.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The description of the network profile.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The data throughput rate in bits per second, as an integer from `0` to `104857600`. Default value is `104857600`.
     */
    public readonly downlinkBandwidthBits!: pulumi.Output<number | undefined>;
    /**
     * Delay time for all packets to destination in milliseconds as an integer from `0` to `2000`.
     */
    public readonly downlinkDelayMs!: pulumi.Output<number | undefined>;
    /**
     * Time variation in the delay of received packets in milliseconds as an integer from `0` to `2000`.
     */
    public readonly downlinkJitterMs!: pulumi.Output<number | undefined>;
    /**
     * Proportion of received packets that fail to arrive from `0` to `100` percent.
     */
    public readonly downlinkLossPercent!: pulumi.Output<number | undefined>;
    /**
     * The name for the network profile.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ARN of the project for the network profile.
     */
    public readonly projectArn!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The type of network profile to create. Valid values are listed are `PRIVATE` and `CURATED`.
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * The data throughput rate in bits per second, as an integer from `0` to `104857600`. Default value is `104857600`.
     */
    public readonly uplinkBandwidthBits!: pulumi.Output<number | undefined>;
    /**
     * Delay time for all packets to destination in milliseconds as an integer from `0` to `2000`.
     */
    public readonly uplinkDelayMs!: pulumi.Output<number | undefined>;
    /**
     * Time variation in the delay of received packets in milliseconds as an integer from `0` to `2000`.
     */
    public readonly uplinkJitterMs!: pulumi.Output<number | undefined>;
    /**
     * Proportion of received packets that fail to arrive from `0` to `100` percent.
     */
    public readonly uplinkLossPercent!: pulumi.Output<number | undefined>;

    /**
     * Create a NetworkProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkProfileArgs | NetworkProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkProfileState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["downlinkBandwidthBits"] = state ? state.downlinkBandwidthBits : undefined;
            resourceInputs["downlinkDelayMs"] = state ? state.downlinkDelayMs : undefined;
            resourceInputs["downlinkJitterMs"] = state ? state.downlinkJitterMs : undefined;
            resourceInputs["downlinkLossPercent"] = state ? state.downlinkLossPercent : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["projectArn"] = state ? state.projectArn : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["uplinkBandwidthBits"] = state ? state.uplinkBandwidthBits : undefined;
            resourceInputs["uplinkDelayMs"] = state ? state.uplinkDelayMs : undefined;
            resourceInputs["uplinkJitterMs"] = state ? state.uplinkJitterMs : undefined;
            resourceInputs["uplinkLossPercent"] = state ? state.uplinkLossPercent : undefined;
        } else {
            const args = argsOrState as NetworkProfileArgs | undefined;
            if ((!args || args.projectArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectArn'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["downlinkBandwidthBits"] = args ? args.downlinkBandwidthBits : undefined;
            resourceInputs["downlinkDelayMs"] = args ? args.downlinkDelayMs : undefined;
            resourceInputs["downlinkJitterMs"] = args ? args.downlinkJitterMs : undefined;
            resourceInputs["downlinkLossPercent"] = args ? args.downlinkLossPercent : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectArn"] = args ? args.projectArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["uplinkBandwidthBits"] = args ? args.uplinkBandwidthBits : undefined;
            resourceInputs["uplinkDelayMs"] = args ? args.uplinkDelayMs : undefined;
            resourceInputs["uplinkJitterMs"] = args ? args.uplinkJitterMs : undefined;
            resourceInputs["uplinkLossPercent"] = args ? args.uplinkLossPercent : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetworkProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkProfile resources.
 */
export interface NetworkProfileState {
    /**
     * The Amazon Resource Name of this network profile.
     */
    arn?: pulumi.Input<string>;
    /**
     * The description of the network profile.
     */
    description?: pulumi.Input<string>;
    /**
     * The data throughput rate in bits per second, as an integer from `0` to `104857600`. Default value is `104857600`.
     */
    downlinkBandwidthBits?: pulumi.Input<number>;
    /**
     * Delay time for all packets to destination in milliseconds as an integer from `0` to `2000`.
     */
    downlinkDelayMs?: pulumi.Input<number>;
    /**
     * Time variation in the delay of received packets in milliseconds as an integer from `0` to `2000`.
     */
    downlinkJitterMs?: pulumi.Input<number>;
    /**
     * Proportion of received packets that fail to arrive from `0` to `100` percent.
     */
    downlinkLossPercent?: pulumi.Input<number>;
    /**
     * The name for the network profile.
     */
    name?: pulumi.Input<string>;
    /**
     * The ARN of the project for the network profile.
     */
    projectArn?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of network profile to create. Valid values are listed are `PRIVATE` and `CURATED`.
     */
    type?: pulumi.Input<string>;
    /**
     * The data throughput rate in bits per second, as an integer from `0` to `104857600`. Default value is `104857600`.
     */
    uplinkBandwidthBits?: pulumi.Input<number>;
    /**
     * Delay time for all packets to destination in milliseconds as an integer from `0` to `2000`.
     */
    uplinkDelayMs?: pulumi.Input<number>;
    /**
     * Time variation in the delay of received packets in milliseconds as an integer from `0` to `2000`.
     */
    uplinkJitterMs?: pulumi.Input<number>;
    /**
     * Proportion of received packets that fail to arrive from `0` to `100` percent.
     */
    uplinkLossPercent?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a NetworkProfile resource.
 */
export interface NetworkProfileArgs {
    /**
     * The description of the network profile.
     */
    description?: pulumi.Input<string>;
    /**
     * The data throughput rate in bits per second, as an integer from `0` to `104857600`. Default value is `104857600`.
     */
    downlinkBandwidthBits?: pulumi.Input<number>;
    /**
     * Delay time for all packets to destination in milliseconds as an integer from `0` to `2000`.
     */
    downlinkDelayMs?: pulumi.Input<number>;
    /**
     * Time variation in the delay of received packets in milliseconds as an integer from `0` to `2000`.
     */
    downlinkJitterMs?: pulumi.Input<number>;
    /**
     * Proportion of received packets that fail to arrive from `0` to `100` percent.
     */
    downlinkLossPercent?: pulumi.Input<number>;
    /**
     * The name for the network profile.
     */
    name?: pulumi.Input<string>;
    /**
     * The ARN of the project for the network profile.
     */
    projectArn: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of network profile to create. Valid values are listed are `PRIVATE` and `CURATED`.
     */
    type?: pulumi.Input<string>;
    /**
     * The data throughput rate in bits per second, as an integer from `0` to `104857600`. Default value is `104857600`.
     */
    uplinkBandwidthBits?: pulumi.Input<number>;
    /**
     * Delay time for all packets to destination in milliseconds as an integer from `0` to `2000`.
     */
    uplinkDelayMs?: pulumi.Input<number>;
    /**
     * Time variation in the delay of received packets in milliseconds as an integer from `0` to `2000`.
     */
    uplinkJitterMs?: pulumi.Input<number>;
    /**
     * Proportion of received packets that fail to arrive from `0` to `100` percent.
     */
    uplinkLossPercent?: pulumi.Input<number>;
}
