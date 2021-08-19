// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a VPC resource.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const main = new aws.ec2.Vpc("main", {
 *     cidrBlock: "10.0.0.0/16",
 * });
 * ```
 *
 * Basic usage with tags:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const main = new aws.ec2.Vpc("main", {
 *     cidrBlock: "10.0.0.0/16",
 *     instanceTenancy: "default",
 *     tags: {
 *         Name: "main",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * VPCs can be imported using the `vpc id`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:ec2/vpc:Vpc test_vpc vpc-a01106c2
 * ```
 */
export class Vpc extends pulumi.CustomResource {
    /**
     * Get an existing Vpc resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcState, opts?: pulumi.CustomResourceOptions): Vpc {
        return new Vpc(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/vpc:Vpc';

    /**
     * Returns true if the given object is an instance of Vpc.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Vpc {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Vpc.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of VPC
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Requests an Amazon-provided IPv6 CIDR
     * block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or
     * the size of the CIDR block. Default is `false`.
     */
    public readonly assignGeneratedIpv6CidrBlock!: pulumi.Output<boolean | undefined>;
    /**
     * The CIDR block for the VPC.
     */
    public readonly cidrBlock!: pulumi.Output<string>;
    /**
     * The ID of the network ACL created by default on VPC creation
     */
    public /*out*/ readonly defaultNetworkAclId!: pulumi.Output<string>;
    /**
     * The ID of the route table created by default on VPC creation
     */
    public /*out*/ readonly defaultRouteTableId!: pulumi.Output<string>;
    /**
     * The ID of the security group created by default on VPC creation
     */
    public /*out*/ readonly defaultSecurityGroupId!: pulumi.Output<string>;
    public /*out*/ readonly dhcpOptionsId!: pulumi.Output<string>;
    /**
     * A boolean flag to enable/disable ClassicLink
     * for the VPC. Only valid in regions and accounts that support EC2 Classic.
     * See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
     */
    public readonly enableClassiclink!: pulumi.Output<boolean>;
    /**
     * A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
     * Only valid in regions and accounts that support EC2 Classic.
     */
    public readonly enableClassiclinkDnsSupport!: pulumi.Output<boolean>;
    /**
     * A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
     */
    public readonly enableDnsHostnames!: pulumi.Output<boolean>;
    /**
     * A boolean flag to enable/disable DNS support in the VPC. Defaults true.
     */
    public readonly enableDnsSupport!: pulumi.Output<boolean | undefined>;
    /**
     * A tenancy option for instances launched into the VPC. Default is `default`, which
     * makes your instances shared on the host. Using either of the other options (`dedicated` or `host`) costs at least $2/hr.
     */
    public readonly instanceTenancy!: pulumi.Output<string | undefined>;
    /**
     * The association ID for the IPv6 CIDR block.
     */
    public /*out*/ readonly ipv6AssociationId!: pulumi.Output<string>;
    /**
     * The IPv6 CIDR block.
     */
    public /*out*/ readonly ipv6CidrBlock!: pulumi.Output<string>;
    /**
     * The ID of the main route table associated with
     * this VPC. Note that you can change a VPC's main route table by using an
     * `aws.ec2.MainRouteTableAssociation`.
     */
    public /*out*/ readonly mainRouteTableId!: pulumi.Output<string>;
    /**
     * The ID of the AWS account that owns the VPC.
     */
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a Vpc resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcArgs | VpcState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["assignGeneratedIpv6CidrBlock"] = state ? state.assignGeneratedIpv6CidrBlock : undefined;
            inputs["cidrBlock"] = state ? state.cidrBlock : undefined;
            inputs["defaultNetworkAclId"] = state ? state.defaultNetworkAclId : undefined;
            inputs["defaultRouteTableId"] = state ? state.defaultRouteTableId : undefined;
            inputs["defaultSecurityGroupId"] = state ? state.defaultSecurityGroupId : undefined;
            inputs["dhcpOptionsId"] = state ? state.dhcpOptionsId : undefined;
            inputs["enableClassiclink"] = state ? state.enableClassiclink : undefined;
            inputs["enableClassiclinkDnsSupport"] = state ? state.enableClassiclinkDnsSupport : undefined;
            inputs["enableDnsHostnames"] = state ? state.enableDnsHostnames : undefined;
            inputs["enableDnsSupport"] = state ? state.enableDnsSupport : undefined;
            inputs["instanceTenancy"] = state ? state.instanceTenancy : undefined;
            inputs["ipv6AssociationId"] = state ? state.ipv6AssociationId : undefined;
            inputs["ipv6CidrBlock"] = state ? state.ipv6CidrBlock : undefined;
            inputs["mainRouteTableId"] = state ? state.mainRouteTableId : undefined;
            inputs["ownerId"] = state ? state.ownerId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as VpcArgs | undefined;
            if ((!args || args.cidrBlock === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cidrBlock'");
            }
            inputs["assignGeneratedIpv6CidrBlock"] = args ? args.assignGeneratedIpv6CidrBlock : undefined;
            inputs["cidrBlock"] = args ? args.cidrBlock : undefined;
            inputs["enableClassiclink"] = args ? args.enableClassiclink : undefined;
            inputs["enableClassiclinkDnsSupport"] = args ? args.enableClassiclinkDnsSupport : undefined;
            inputs["enableDnsHostnames"] = args ? args.enableDnsHostnames : undefined;
            inputs["enableDnsSupport"] = args ? args.enableDnsSupport : undefined;
            inputs["instanceTenancy"] = args ? args.instanceTenancy : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["defaultNetworkAclId"] = undefined /*out*/;
            inputs["defaultRouteTableId"] = undefined /*out*/;
            inputs["defaultSecurityGroupId"] = undefined /*out*/;
            inputs["dhcpOptionsId"] = undefined /*out*/;
            inputs["ipv6AssociationId"] = undefined /*out*/;
            inputs["ipv6CidrBlock"] = undefined /*out*/;
            inputs["mainRouteTableId"] = undefined /*out*/;
            inputs["ownerId"] = undefined /*out*/;
            inputs["tagsAll"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Vpc.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Vpc resources.
 */
export interface VpcState {
    /**
     * Amazon Resource Name (ARN) of VPC
     */
    arn?: pulumi.Input<string>;
    /**
     * Requests an Amazon-provided IPv6 CIDR
     * block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or
     * the size of the CIDR block. Default is `false`.
     */
    assignGeneratedIpv6CidrBlock?: pulumi.Input<boolean>;
    /**
     * The CIDR block for the VPC.
     */
    cidrBlock?: pulumi.Input<string>;
    /**
     * The ID of the network ACL created by default on VPC creation
     */
    defaultNetworkAclId?: pulumi.Input<string>;
    /**
     * The ID of the route table created by default on VPC creation
     */
    defaultRouteTableId?: pulumi.Input<string>;
    /**
     * The ID of the security group created by default on VPC creation
     */
    defaultSecurityGroupId?: pulumi.Input<string>;
    dhcpOptionsId?: pulumi.Input<string>;
    /**
     * A boolean flag to enable/disable ClassicLink
     * for the VPC. Only valid in regions and accounts that support EC2 Classic.
     * See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
     */
    enableClassiclink?: pulumi.Input<boolean>;
    /**
     * A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
     * Only valid in regions and accounts that support EC2 Classic.
     */
    enableClassiclinkDnsSupport?: pulumi.Input<boolean>;
    /**
     * A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
     */
    enableDnsHostnames?: pulumi.Input<boolean>;
    /**
     * A boolean flag to enable/disable DNS support in the VPC. Defaults true.
     */
    enableDnsSupport?: pulumi.Input<boolean>;
    /**
     * A tenancy option for instances launched into the VPC. Default is `default`, which
     * makes your instances shared on the host. Using either of the other options (`dedicated` or `host`) costs at least $2/hr.
     */
    instanceTenancy?: pulumi.Input<string>;
    /**
     * The association ID for the IPv6 CIDR block.
     */
    ipv6AssociationId?: pulumi.Input<string>;
    /**
     * The IPv6 CIDR block.
     */
    ipv6CidrBlock?: pulumi.Input<string>;
    /**
     * The ID of the main route table associated with
     * this VPC. Note that you can change a VPC's main route table by using an
     * `aws.ec2.MainRouteTableAssociation`.
     */
    mainRouteTableId?: pulumi.Input<string>;
    /**
     * The ID of the AWS account that owns the VPC.
     */
    ownerId?: pulumi.Input<string>;
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
 * The set of arguments for constructing a Vpc resource.
 */
export interface VpcArgs {
    /**
     * Requests an Amazon-provided IPv6 CIDR
     * block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or
     * the size of the CIDR block. Default is `false`.
     */
    assignGeneratedIpv6CidrBlock?: pulumi.Input<boolean>;
    /**
     * The CIDR block for the VPC.
     */
    cidrBlock: pulumi.Input<string>;
    /**
     * A boolean flag to enable/disable ClassicLink
     * for the VPC. Only valid in regions and accounts that support EC2 Classic.
     * See the [ClassicLink documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) for more information. Defaults false.
     */
    enableClassiclink?: pulumi.Input<boolean>;
    /**
     * A boolean flag to enable/disable ClassicLink DNS Support for the VPC.
     * Only valid in regions and accounts that support EC2 Classic.
     */
    enableClassiclinkDnsSupport?: pulumi.Input<boolean>;
    /**
     * A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
     */
    enableDnsHostnames?: pulumi.Input<boolean>;
    /**
     * A boolean flag to enable/disable DNS support in the VPC. Defaults true.
     */
    enableDnsSupport?: pulumi.Input<boolean>;
    /**
     * A tenancy option for instances launched into the VPC. Default is `default`, which
     * makes your instances shared on the host. Using either of the other options (`dedicated` or `host`) costs at least $2/hr.
     */
    instanceTenancy?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
