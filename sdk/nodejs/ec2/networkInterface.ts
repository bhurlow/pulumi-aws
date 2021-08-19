// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an Elastic network interface (ENI) resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.ec2.NetworkInterface("test", {
 *     subnetId: aws_subnet.public_a.id,
 *     privateIps: ["10.0.0.50"],
 *     securityGroups: [aws_security_group.web.id],
 *     attachments: [{
 *         instance: aws_instance.test.id,
 *         deviceIndex: 1,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Network Interfaces can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:ec2/networkInterface:NetworkInterface test eni-e5aa89a3
 * ```
 */
export class NetworkInterface extends pulumi.CustomResource {
    /**
     * Get an existing NetworkInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkInterfaceState, opts?: pulumi.CustomResourceOptions): NetworkInterface {
        return new NetworkInterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/networkInterface:NetworkInterface';

    /**
     * Returns true if the given object is an instance of NetworkInterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkInterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkInterface.__pulumiType;
    }

    /**
     * Block to define the attachment of the ENI. Documented below.
     */
    public readonly attachments!: pulumi.Output<outputs.ec2.NetworkInterfaceAttachment[]>;
    /**
     * A description for the network interface.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Type of network interface to create. Set to `efa` for Elastic Fabric Adapter.
     */
    public readonly interfaceType!: pulumi.Output<string>;
    /**
     * The number of IPv6 addresses to assign to a network interface. You can't use this option if specifying specific `ipv6Addresses`. If your subnet has the AssignIpv6AddressOnCreation attribute set to `true`, you can specify `0` to override this setting.
     */
    public readonly ipv6AddressCount!: pulumi.Output<number>;
    /**
     * One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. You can't use this option if you're specifying `ipv6AddressCount`.
     */
    public readonly ipv6Addresses!: pulumi.Output<string[]>;
    /**
     * The MAC address of the network interface.
     */
    public /*out*/ readonly macAddress!: pulumi.Output<string>;
    public /*out*/ readonly outpostArn!: pulumi.Output<string>;
    /**
     * The private DNS name of the network interface (IPv4).
     */
    public /*out*/ readonly privateDnsName!: pulumi.Output<string>;
    public readonly privateIp!: pulumi.Output<string>;
    /**
     * List of private IPs to assign to the ENI.
     */
    public readonly privateIps!: pulumi.Output<string[]>;
    /**
     * Number of secondary private IPs to assign to the ENI. The total number of private IPs will be 1 + private_ips_count, as a primary private IP will be assiged to an ENI by default.
     */
    public readonly privateIpsCount!: pulumi.Output<number>;
    /**
     * List of security group IDs to assign to the ENI.
     */
    public readonly securityGroups!: pulumi.Output<string[]>;
    /**
     * Whether to enable source destination checking for the ENI. Default true.
     */
    public readonly sourceDestCheck!: pulumi.Output<boolean | undefined>;
    /**
     * Subnet ID to create the ENI in.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a NetworkInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkInterfaceArgs | NetworkInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkInterfaceState | undefined;
            inputs["attachments"] = state ? state.attachments : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["interfaceType"] = state ? state.interfaceType : undefined;
            inputs["ipv6AddressCount"] = state ? state.ipv6AddressCount : undefined;
            inputs["ipv6Addresses"] = state ? state.ipv6Addresses : undefined;
            inputs["macAddress"] = state ? state.macAddress : undefined;
            inputs["outpostArn"] = state ? state.outpostArn : undefined;
            inputs["privateDnsName"] = state ? state.privateDnsName : undefined;
            inputs["privateIp"] = state ? state.privateIp : undefined;
            inputs["privateIps"] = state ? state.privateIps : undefined;
            inputs["privateIpsCount"] = state ? state.privateIpsCount : undefined;
            inputs["securityGroups"] = state ? state.securityGroups : undefined;
            inputs["sourceDestCheck"] = state ? state.sourceDestCheck : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as NetworkInterfaceArgs | undefined;
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            inputs["attachments"] = args ? args.attachments : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["interfaceType"] = args ? args.interfaceType : undefined;
            inputs["ipv6AddressCount"] = args ? args.ipv6AddressCount : undefined;
            inputs["ipv6Addresses"] = args ? args.ipv6Addresses : undefined;
            inputs["privateIp"] = args ? args.privateIp : undefined;
            inputs["privateIps"] = args ? args.privateIps : undefined;
            inputs["privateIpsCount"] = args ? args.privateIpsCount : undefined;
            inputs["securityGroups"] = args ? args.securityGroups : undefined;
            inputs["sourceDestCheck"] = args ? args.sourceDestCheck : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["macAddress"] = undefined /*out*/;
            inputs["outpostArn"] = undefined /*out*/;
            inputs["privateDnsName"] = undefined /*out*/;
            inputs["tagsAll"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NetworkInterface.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkInterface resources.
 */
export interface NetworkInterfaceState {
    /**
     * Block to define the attachment of the ENI. Documented below.
     */
    attachments?: pulumi.Input<pulumi.Input<inputs.ec2.NetworkInterfaceAttachment>[]>;
    /**
     * A description for the network interface.
     */
    description?: pulumi.Input<string>;
    /**
     * Type of network interface to create. Set to `efa` for Elastic Fabric Adapter.
     */
    interfaceType?: pulumi.Input<string>;
    /**
     * The number of IPv6 addresses to assign to a network interface. You can't use this option if specifying specific `ipv6Addresses`. If your subnet has the AssignIpv6AddressOnCreation attribute set to `true`, you can specify `0` to override this setting.
     */
    ipv6AddressCount?: pulumi.Input<number>;
    /**
     * One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. You can't use this option if you're specifying `ipv6AddressCount`.
     */
    ipv6Addresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The MAC address of the network interface.
     */
    macAddress?: pulumi.Input<string>;
    outpostArn?: pulumi.Input<string>;
    /**
     * The private DNS name of the network interface (IPv4).
     */
    privateDnsName?: pulumi.Input<string>;
    privateIp?: pulumi.Input<string>;
    /**
     * List of private IPs to assign to the ENI.
     */
    privateIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Number of secondary private IPs to assign to the ENI. The total number of private IPs will be 1 + private_ips_count, as a primary private IP will be assiged to an ENI by default.
     */
    privateIpsCount?: pulumi.Input<number>;
    /**
     * List of security group IDs to assign to the ENI.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to enable source destination checking for the ENI. Default true.
     */
    sourceDestCheck?: pulumi.Input<boolean>;
    /**
     * Subnet ID to create the ENI in.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a NetworkInterface resource.
 */
export interface NetworkInterfaceArgs {
    /**
     * Block to define the attachment of the ENI. Documented below.
     */
    attachments?: pulumi.Input<pulumi.Input<inputs.ec2.NetworkInterfaceAttachment>[]>;
    /**
     * A description for the network interface.
     */
    description?: pulumi.Input<string>;
    /**
     * Type of network interface to create. Set to `efa` for Elastic Fabric Adapter.
     */
    interfaceType?: pulumi.Input<string>;
    /**
     * The number of IPv6 addresses to assign to a network interface. You can't use this option if specifying specific `ipv6Addresses`. If your subnet has the AssignIpv6AddressOnCreation attribute set to `true`, you can specify `0` to override this setting.
     */
    ipv6AddressCount?: pulumi.Input<number>;
    /**
     * One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. You can't use this option if you're specifying `ipv6AddressCount`.
     */
    ipv6Addresses?: pulumi.Input<pulumi.Input<string>[]>;
    privateIp?: pulumi.Input<string>;
    /**
     * List of private IPs to assign to the ENI.
     */
    privateIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Number of secondary private IPs to assign to the ENI. The total number of private IPs will be 1 + private_ips_count, as a primary private IP will be assiged to an ENI by default.
     */
    privateIpsCount?: pulumi.Input<number>;
    /**
     * List of security group IDs to assign to the ENI.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to enable source destination checking for the ENI. Default true.
     */
    sourceDestCheck?: pulumi.Input<boolean>;
    /**
     * Subnet ID to create the ENI in.
     */
    subnetId: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
