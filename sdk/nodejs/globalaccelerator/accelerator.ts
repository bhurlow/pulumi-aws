// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Creates a Global Accelerator accelerator.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.globalaccelerator.Accelerator("example", {
 *     attributes: {
 *         flowLogsEnabled: true,
 *         flowLogsS3Bucket: "example-bucket",
 *         flowLogsS3Prefix: "flow-logs/",
 *     },
 *     enabled: true,
 *     ipAddressType: "IPV4",
 *     ipAddresses: ["1.2.3.4"],
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_globalaccelerator_accelerator.example
 *
 *  id = "arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx" } Using `pulumi import`, import Global Accelerator accelerators using the `arn`. For exampleconsole % pulumi import aws_globalaccelerator_accelerator.example arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
 */
export class Accelerator extends pulumi.CustomResource {
    /**
     * Get an existing Accelerator resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AcceleratorState, opts?: pulumi.CustomResourceOptions): Accelerator {
        return new Accelerator(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:globalaccelerator/accelerator:Accelerator';

    /**
     * Returns true if the given object is an instance of Accelerator.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Accelerator {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Accelerator.__pulumiType;
    }

    /**
     * The attributes of the accelerator. Fields documented below.
     */
    public readonly attributes!: pulumi.Output<outputs.globalaccelerator.AcceleratorAttributes | undefined>;
    /**
     * The DNS name of the accelerator. For example, `a5d53ff5ee6bca4ce.awsglobalaccelerator.com`.
     */
    public /*out*/ readonly dnsName!: pulumi.Output<string>;
    /**
     * The Domain Name System (DNS) name that Global Accelerator creates that points to a dual-stack accelerator's four static IP addresses: two IPv4 addresses and two IPv6 addresses. For example, `a1234567890abcdef.dualstack.awsglobalaccelerator.com`.
     */
    public /*out*/ readonly dualStackDnsName!: pulumi.Output<string>;
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
     * The value for the address type. Defaults to `IPV4`. Valid values: `IPV4`, `DUAL_STACK`.
     */
    public readonly ipAddressType!: pulumi.Output<string | undefined>;
    /**
     * The IP addresses to use for BYOIP accelerators. If not specified, the service assigns IP addresses. Valid values: 1 or 2 IPv4 addresses.
     */
    public readonly ipAddresses!: pulumi.Output<string[] | undefined>;
    /**
     * IP address set associated with the accelerator.
     */
    public /*out*/ readonly ipSets!: pulumi.Output<outputs.globalaccelerator.AcceleratorIpSet[]>;
    /**
     * The name of the accelerator.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a Accelerator resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AcceleratorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AcceleratorArgs | AcceleratorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AcceleratorState | undefined;
            resourceInputs["attributes"] = state ? state.attributes : undefined;
            resourceInputs["dnsName"] = state ? state.dnsName : undefined;
            resourceInputs["dualStackDnsName"] = state ? state.dualStackDnsName : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["hostedZoneId"] = state ? state.hostedZoneId : undefined;
            resourceInputs["ipAddressType"] = state ? state.ipAddressType : undefined;
            resourceInputs["ipAddresses"] = state ? state.ipAddresses : undefined;
            resourceInputs["ipSets"] = state ? state.ipSets : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as AcceleratorArgs | undefined;
            resourceInputs["attributes"] = args ? args.attributes : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["ipAddressType"] = args ? args.ipAddressType : undefined;
            resourceInputs["ipAddresses"] = args ? args.ipAddresses : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["dnsName"] = undefined /*out*/;
            resourceInputs["dualStackDnsName"] = undefined /*out*/;
            resourceInputs["hostedZoneId"] = undefined /*out*/;
            resourceInputs["ipSets"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Accelerator.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Accelerator resources.
 */
export interface AcceleratorState {
    /**
     * The attributes of the accelerator. Fields documented below.
     */
    attributes?: pulumi.Input<inputs.globalaccelerator.AcceleratorAttributes>;
    /**
     * The DNS name of the accelerator. For example, `a5d53ff5ee6bca4ce.awsglobalaccelerator.com`.
     */
    dnsName?: pulumi.Input<string>;
    /**
     * The Domain Name System (DNS) name that Global Accelerator creates that points to a dual-stack accelerator's four static IP addresses: two IPv4 addresses and two IPv6 addresses. For example, `a1234567890abcdef.dualstack.awsglobalaccelerator.com`.
     */
    dualStackDnsName?: pulumi.Input<string>;
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
     * The value for the address type. Defaults to `IPV4`. Valid values: `IPV4`, `DUAL_STACK`.
     */
    ipAddressType?: pulumi.Input<string>;
    /**
     * The IP addresses to use for BYOIP accelerators. If not specified, the service assigns IP addresses. Valid values: 1 or 2 IPv4 addresses.
     */
    ipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * IP address set associated with the accelerator.
     */
    ipSets?: pulumi.Input<pulumi.Input<inputs.globalaccelerator.AcceleratorIpSet>[]>;
    /**
     * The name of the accelerator.
     */
    name?: pulumi.Input<string>;
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
 * The set of arguments for constructing a Accelerator resource.
 */
export interface AcceleratorArgs {
    /**
     * The attributes of the accelerator. Fields documented below.
     */
    attributes?: pulumi.Input<inputs.globalaccelerator.AcceleratorAttributes>;
    /**
     * Indicates whether the accelerator is enabled. Defaults to `true`. Valid values: `true`, `false`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The value for the address type. Defaults to `IPV4`. Valid values: `IPV4`, `DUAL_STACK`.
     */
    ipAddressType?: pulumi.Input<string>;
    /**
     * The IP addresses to use for BYOIP accelerators. If not specified, the service assigns IP addresses. Valid values: 1 or 2 IPv4 addresses.
     */
    ipAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the accelerator.
     */
    name?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
