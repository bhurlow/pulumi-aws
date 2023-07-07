// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Creates a Global Accelerator custom routing accelerator.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.globalaccelerator.CustomRoutingAccelerator("example", {
 *     attributes: {
 *         flowLogsEnabled: true,
 *         flowLogsS3Bucket: "example-bucket",
 *         flowLogsS3Prefix: "flow-logs/",
 *     },
 *     enabled: true,
 *     ipAddressType: "IPV4",
 *     ipAddresses: ["1.2.3.4"],
 *     name: "Example",
 * });
 * ```
 *
 * ## Import
 *
 * Global Accelerator custom routing accelerators can be imported using the `arn`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:globalaccelerator/customRoutingAccelerator:CustomRoutingAccelerator example arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
 * ```
 */
export class CustomRoutingAccelerator extends pulumi.CustomResource {
    /**
     * Get an existing CustomRoutingAccelerator resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomRoutingAcceleratorState, opts?: pulumi.CustomResourceOptions): CustomRoutingAccelerator {
        return new CustomRoutingAccelerator(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:globalaccelerator/customRoutingAccelerator:CustomRoutingAccelerator';

    /**
     * Returns true if the given object is an instance of CustomRoutingAccelerator.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomRoutingAccelerator {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomRoutingAccelerator.__pulumiType;
    }

    /**
     * The attributes of the accelerator. Fields documented below.
     */
    public readonly attributes!: pulumi.Output<outputs.globalaccelerator.CustomRoutingAcceleratorAttributes | undefined>;
    /**
     * The DNS name of the accelerator. For example, `a5d53ff5ee6bca4ce.awsglobalaccelerator.com`.
     */
    public /*out*/ readonly dnsName!: pulumi.Output<string>;
    /**
     * Indicates whether the accelerator is enabled. Defaults to `true`. Valid values: `true`, `false`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * -  The Global Accelerator Route 53 zone ID that can be used to
     * route an [Alias Resource Record Set](https://docs.aws.amazon.com/Route53/latest/APIReference/API_AliasTarget.html) to the Global Accelerator. This attribute
     * is simply an alias for the zone ID `Z2BJ6XQ5FK7U4H`.
     */
    public /*out*/ readonly hostedZoneId!: pulumi.Output<string>;
    /**
     * The IP address type that an accelerator supports. For a custom routing accelerator, the value must be `"IPV4"`.
     */
    public readonly ipAddressType!: pulumi.Output<string | undefined>;
    /**
     * The IP addresses to use for BYOIP accelerators. If not specified, the service assigns IP addresses. Valid values: 1 or 2 IPv4 addresses.
     */
    public readonly ipAddresses!: pulumi.Output<string[] | undefined>;
    /**
     * IP address set associated with the accelerator.
     */
    public /*out*/ readonly ipSets!: pulumi.Output<outputs.globalaccelerator.CustomRoutingAcceleratorIpSet[]>;
    /**
     * The name of a custom routing accelerator.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     *
     * **attributes** supports the following attributes:
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a CustomRoutingAccelerator resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomRoutingAcceleratorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomRoutingAcceleratorArgs | CustomRoutingAcceleratorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomRoutingAcceleratorState | undefined;
            resourceInputs["attributes"] = state ? state.attributes : undefined;
            resourceInputs["dnsName"] = state ? state.dnsName : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["hostedZoneId"] = state ? state.hostedZoneId : undefined;
            resourceInputs["ipAddressType"] = state ? state.ipAddressType : undefined;
            resourceInputs["ipAddresses"] = state ? state.ipAddresses : undefined;
            resourceInputs["ipSets"] = state ? state.ipSets : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as CustomRoutingAcceleratorArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["attributes"] = args ? args.attributes : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["ipAddressType"] = args ? args.ipAddressType : undefined;
            resourceInputs["ipAddresses"] = args ? args.ipAddresses : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["dnsName"] = undefined /*out*/;
            resourceInputs["hostedZoneId"] = undefined /*out*/;
            resourceInputs["ipSets"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomRoutingAccelerator.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomRoutingAccelerator resources.
 */
export interface CustomRoutingAcceleratorState {
    /**
     * The attributes of the accelerator. Fields documented below.
     */
    attributes?: pulumi.Input<inputs.globalaccelerator.CustomRoutingAcceleratorAttributes>;
    /**
     * The DNS name of the accelerator. For example, `a5d53ff5ee6bca4ce.awsglobalaccelerator.com`.
     */
    dnsName?: pulumi.Input<string>;
    /**
     * Indicates whether the accelerator is enabled. Defaults to `true`. Valid values: `true`, `false`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * -  The Global Accelerator Route 53 zone ID that can be used to
     * route an [Alias Resource Record Set](https://docs.aws.amazon.com/Route53/latest/APIReference/API_AliasTarget.html) to the Global Accelerator. This attribute
     * is simply an alias for the zone ID `Z2BJ6XQ5FK7U4H`.
     */
    hostedZoneId?: pulumi.Input<string>;
    /**
     * The IP address type that an accelerator supports. For a custom routing accelerator, the value must be `"IPV4"`.
     */
    ipAddressType?: pulumi.Input<string>;
    /**
     * The IP addresses to use for BYOIP accelerators. If not specified, the service assigns IP addresses. Valid values: 1 or 2 IPv4 addresses.
     */
    ipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * IP address set associated with the accelerator.
     */
    ipSets?: pulumi.Input<pulumi.Input<inputs.globalaccelerator.CustomRoutingAcceleratorIpSet>[]>;
    /**
     * The name of a custom routing accelerator.
     */
    name?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     *
     * **attributes** supports the following attributes:
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a CustomRoutingAccelerator resource.
 */
export interface CustomRoutingAcceleratorArgs {
    /**
     * The attributes of the accelerator. Fields documented below.
     */
    attributes?: pulumi.Input<inputs.globalaccelerator.CustomRoutingAcceleratorAttributes>;
    /**
     * Indicates whether the accelerator is enabled. Defaults to `true`. Valid values: `true`, `false`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The IP address type that an accelerator supports. For a custom routing accelerator, the value must be `"IPV4"`.
     */
    ipAddressType?: pulumi.Input<string>;
    /**
     * The IP addresses to use for BYOIP accelerators. If not specified, the service assigns IP addresses. Valid values: 1 or 2 IPv4 addresses.
     */
    ipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of a custom routing accelerator.
     */
    name: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     *
     * **attributes** supports the following attributes:
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
