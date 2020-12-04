// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource attaches a security group to an Elastic Network Interface (ENI).
 * It can be used to attach a security group to any existing ENI, be it a
 * secondary ENI or one attached as the primary interface on an instance.
 *
 * > **NOTE on instances, interfaces, and security groups:** This provider currently
 * provides the capability to assign security groups via the [`aws.ec2.Instance`][1]
 * and the [`aws.ec2.NetworkInterface`][2] resources. Using this resource in
 * conjunction with security groups provided in-line in those resources will cause
 * conflicts, and will lead to spurious diffs and undefined behavior - please use
 * one or the other.
 *
 * [1]: https://www.terraform.io/docs/providers/aws/d/instance.html
 * [2]: https://www.terraform.io/docs/providers/aws/r/network_interface.html
 *
 * ## Example Usage
 *
 * The following provides a very basic example of setting up an instance (provided
 * by `instance`) in the default security group, creating a security group
 * (provided by `sg`) and then attaching the security group to the instance's
 * primary network interface via the `aws.ec2.NetworkInterfaceSecurityGroupAttachment` resource,
 * named `sgAttachment`:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const ami = aws.getAmi({
 *     mostRecent: true,
 *     filters: [{
 *         name: "name",
 *         values: ["amzn-ami-hvm-*"],
 *     }],
 *     owners: ["amazon"],
 * });
 * const instance = new aws.ec2.Instance("instance", {
 *     instanceType: "t2.micro",
 *     ami: ami.then(ami => ami.id),
 *     tags: {
 *         type: "test-instance",
 *     },
 * });
 * const sg = new aws.ec2.SecurityGroup("sg", {tags: {
 *     type: "test-security-group",
 * }});
 * const sgAttachment = new aws.ec2.NetworkInterfaceSecurityGroupAttachment("sgAttachment", {
 *     securityGroupId: sg.id,
 *     networkInterfaceId: instance.primaryNetworkInterfaceId,
 * });
 * ```
 *
 * In this example, `instance` is provided by the `aws.ec2.Instance` data source,
 * fetching an external instance, possibly not managed by this provider.
 * `sgAttachment` then attaches to the output instance's `networkInterfaceId`:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const instance = aws.ec2.getInstance({
 *     instanceId: "i-1234567890abcdef0",
 * });
 * const sg = new aws.ec2.SecurityGroup("sg", {tags: {
 *     type: "test-security-group",
 * }});
 * const sgAttachment = new aws.ec2.NetworkInterfaceSecurityGroupAttachment("sgAttachment", {
 *     securityGroupId: sg.id,
 *     networkInterfaceId: instance.then(instance => instance.networkInterfaceId),
 * });
 * ```
 * ## Output Reference
 *
 * There are no outputs for this resource.
 */
export class NetworkInterfaceSecurityGroupAttachment extends pulumi.CustomResource {
    /**
     * Get an existing NetworkInterfaceSecurityGroupAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkInterfaceSecurityGroupAttachmentState, opts?: pulumi.CustomResourceOptions): NetworkInterfaceSecurityGroupAttachment {
        return new NetworkInterfaceSecurityGroupAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/networkInterfaceSecurityGroupAttachment:NetworkInterfaceSecurityGroupAttachment';

    /**
     * Returns true if the given object is an instance of NetworkInterfaceSecurityGroupAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkInterfaceSecurityGroupAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkInterfaceSecurityGroupAttachment.__pulumiType;
    }

    /**
     * The ID of the network interface to attach to.
     */
    public readonly networkInterfaceId!: pulumi.Output<string>;
    /**
     * The ID of the security group.
     */
    public readonly securityGroupId!: pulumi.Output<string>;

    /**
     * Create a NetworkInterfaceSecurityGroupAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkInterfaceSecurityGroupAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkInterfaceSecurityGroupAttachmentArgs | NetworkInterfaceSecurityGroupAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NetworkInterfaceSecurityGroupAttachmentState | undefined;
            inputs["networkInterfaceId"] = state ? state.networkInterfaceId : undefined;
            inputs["securityGroupId"] = state ? state.securityGroupId : undefined;
        } else {
            const args = argsOrState as NetworkInterfaceSecurityGroupAttachmentArgs | undefined;
            if ((!args || args.networkInterfaceId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'networkInterfaceId'");
            }
            if ((!args || args.securityGroupId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'securityGroupId'");
            }
            inputs["networkInterfaceId"] = args ? args.networkInterfaceId : undefined;
            inputs["securityGroupId"] = args ? args.securityGroupId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(NetworkInterfaceSecurityGroupAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkInterfaceSecurityGroupAttachment resources.
 */
export interface NetworkInterfaceSecurityGroupAttachmentState {
    /**
     * The ID of the network interface to attach to.
     */
    readonly networkInterfaceId?: pulumi.Input<string>;
    /**
     * The ID of the security group.
     */
    readonly securityGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkInterfaceSecurityGroupAttachment resource.
 */
export interface NetworkInterfaceSecurityGroupAttachmentArgs {
    /**
     * The ID of the network interface to attach to.
     */
    readonly networkInterfaceId: pulumi.Input<string>;
    /**
     * The ID of the security group.
     */
    readonly securityGroupId: pulumi.Input<string>;
}
