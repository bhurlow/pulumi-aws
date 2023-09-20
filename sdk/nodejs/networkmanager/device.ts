// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Creates a device in a global network. If you specify both a site ID and a location,
 * the location of the site is used for visualization in the Network Manager console.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.networkmanager.Device("example", {
 *     globalNetworkId: aws_networkmanager_global_network.example.id,
 *     siteId: aws_networkmanager_site.example.id,
 * });
 * ```
 *
 * ## Import
 *
 * Using `pulumi import`, import `aws_networkmanager_device` using the device ARN. For example:
 *
 * ```sh
 *  $ pulumi import aws:networkmanager/device:Device example arn:aws:networkmanager::123456789012:device/global-network-0d47f6t230mz46dy4/device-07f6fd08867abc123
 * ```
 */
export class Device extends pulumi.CustomResource {
    /**
     * Get an existing Device resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeviceState, opts?: pulumi.CustomResourceOptions): Device {
        return new Device(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:networkmanager/device:Device';

    /**
     * Returns true if the given object is an instance of Device.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Device {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Device.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the device.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The AWS location of the device. Documented below.
     */
    public readonly awsLocation!: pulumi.Output<outputs.networkmanager.DeviceAwsLocation | undefined>;
    /**
     * A description of the device.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the global network.
     */
    public readonly globalNetworkId!: pulumi.Output<string>;
    /**
     * The location of the device. Documented below.
     */
    public readonly location!: pulumi.Output<outputs.networkmanager.DeviceLocation | undefined>;
    /**
     * The model of device.
     */
    public readonly model!: pulumi.Output<string | undefined>;
    /**
     * The serial number of the device.
     */
    public readonly serialNumber!: pulumi.Output<string | undefined>;
    /**
     * The ID of the site.
     */
    public readonly siteId!: pulumi.Output<string | undefined>;
    /**
     * Key-value tags for the device. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The type of device.
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * The vendor of the device.
     */
    public readonly vendor!: pulumi.Output<string | undefined>;

    /**
     * Create a Device resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeviceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeviceArgs | DeviceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeviceState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["awsLocation"] = state ? state.awsLocation : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["globalNetworkId"] = state ? state.globalNetworkId : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["model"] = state ? state.model : undefined;
            resourceInputs["serialNumber"] = state ? state.serialNumber : undefined;
            resourceInputs["siteId"] = state ? state.siteId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vendor"] = state ? state.vendor : undefined;
        } else {
            const args = argsOrState as DeviceArgs | undefined;
            if ((!args || args.globalNetworkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'globalNetworkId'");
            }
            resourceInputs["awsLocation"] = args ? args.awsLocation : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["globalNetworkId"] = args ? args.globalNetworkId : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["model"] = args ? args.model : undefined;
            resourceInputs["serialNumber"] = args ? args.serialNumber : undefined;
            resourceInputs["siteId"] = args ? args.siteId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vendor"] = args ? args.vendor : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Device.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Device resources.
 */
export interface DeviceState {
    /**
     * The Amazon Resource Name (ARN) of the device.
     */
    arn?: pulumi.Input<string>;
    /**
     * The AWS location of the device. Documented below.
     */
    awsLocation?: pulumi.Input<inputs.networkmanager.DeviceAwsLocation>;
    /**
     * A description of the device.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the global network.
     */
    globalNetworkId?: pulumi.Input<string>;
    /**
     * The location of the device. Documented below.
     */
    location?: pulumi.Input<inputs.networkmanager.DeviceLocation>;
    /**
     * The model of device.
     */
    model?: pulumi.Input<string>;
    /**
     * The serial number of the device.
     */
    serialNumber?: pulumi.Input<string>;
    /**
     * The ID of the site.
     */
    siteId?: pulumi.Input<string>;
    /**
     * Key-value tags for the device. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of device.
     */
    type?: pulumi.Input<string>;
    /**
     * The vendor of the device.
     */
    vendor?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Device resource.
 */
export interface DeviceArgs {
    /**
     * The AWS location of the device. Documented below.
     */
    awsLocation?: pulumi.Input<inputs.networkmanager.DeviceAwsLocation>;
    /**
     * A description of the device.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the global network.
     */
    globalNetworkId: pulumi.Input<string>;
    /**
     * The location of the device. Documented below.
     */
    location?: pulumi.Input<inputs.networkmanager.DeviceLocation>;
    /**
     * The model of device.
     */
    model?: pulumi.Input<string>;
    /**
     * The serial number of the device.
     */
    serialNumber?: pulumi.Input<string>;
    /**
     * The ID of the site.
     */
    siteId?: pulumi.Input<string>;
    /**
     * Key-value tags for the device. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of device.
     */
    type?: pulumi.Input<string>;
    /**
     * The vendor of the device.
     */
    vendor?: pulumi.Input<string>;
}
