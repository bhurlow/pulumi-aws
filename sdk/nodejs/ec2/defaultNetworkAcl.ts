// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage a VPC's default network ACL. This resource can manage the default network ACL of the default or a non-default VPC.
 *
 * > **NOTE:** This is an advanced resource with special caveats. Please read this document in its entirety before using this resource. The `aws.ec2.DefaultNetworkAcl` behaves differently from normal resources. This provider does not _create_ this resource but instead attempts to "adopt" it into management.
 *
 * Every VPC has a default network ACL that can be managed but not destroyed. When the provider first adopts the Default Network ACL, it **immediately removes all rules in the ACL**. It then proceeds to create any rules specified in the configuration. This step is required so that only the rules specified in the configuration are created.
 *
 * This resource treats its inline rules as absolute; only the rules defined inline are created, and any additions/removals external to this resource will result in diffs being shown. For these reasons, this resource is incompatible with the `aws.ec2.NetworkAclRule` resource.
 *
 * For more information about Network ACLs, see the AWS Documentation on [Network ACLs][aws-network-acls].
 *
 * ## Example Usage
 * ### Basic Example
 *
 * The following config gives the Default Network ACL the same rules that AWS includes but pulls the resource under management by this provider. This means that any ACL rules added or changed will be detected as drift.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const mainvpc = new aws.ec2.Vpc("mainvpc", {cidrBlock: "10.1.0.0/16"});
 * const _default = new aws.ec2.DefaultNetworkAcl("default", {
 *     defaultNetworkAclId: mainvpc.defaultNetworkAclId,
 *     ingress: [{
 *         protocol: "-1",
 *         ruleNo: 100,
 *         action: "allow",
 *         cidrBlock: "0.0.0.0/0",
 *         fromPort: 0,
 *         toPort: 0,
 *     }],
 *     egress: [{
 *         protocol: "-1",
 *         ruleNo: 100,
 *         action: "allow",
 *         cidrBlock: "0.0.0.0/0",
 *         fromPort: 0,
 *         toPort: 0,
 *     }],
 * });
 * ```
 * ### Example: Deny All Egress Traffic, Allow Ingress
 *
 * The following denies all Egress traffic by omitting any `egress` rules, while including the default `ingress` rule to allow all traffic.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const mainvpc = new aws.ec2.Vpc("mainvpc", {cidrBlock: "10.1.0.0/16"});
 * const _default = new aws.ec2.DefaultNetworkAcl("default", {
 *     defaultNetworkAclId: mainvpc.defaultNetworkAclId,
 *     ingress: [{
 *         protocol: "-1",
 *         ruleNo: 100,
 *         action: "allow",
 *         cidrBlock: aws_default_vpc.mainvpc.cidr_block,
 *         fromPort: 0,
 *         toPort: 0,
 *     }],
 * });
 * ```
 * ### Example: Deny All Traffic To Any Subnet In The Default Network ACL
 *
 * This config denies all traffic in the Default ACL. This can be useful if you want to lock down the VPC to force all resources to assign a non-default ACL.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const mainvpc = new aws.ec2.Vpc("mainvpc", {cidrBlock: "10.1.0.0/16"});
 * const _default = new aws.ec2.DefaultNetworkAcl("default", {defaultNetworkAclId: mainvpc.defaultNetworkAclId});
 * // no rules defined, deny all traffic in this ACL
 * ```
 * ### Managing Subnets In A Default Network ACL
 *
 * Within a VPC, all Subnets must be associated with a Network ACL. In order to "delete" the association between a Subnet and a non-default Network ACL, the association is destroyed by replacing it with an association between the Subnet and the Default ACL instead.
 *
 * When managing the Default Network ACL, you cannot "remove" Subnets. Instead, they must be reassigned to another Network ACL, or the Subnet itself must be destroyed. Because of these requirements, removing the `subnetIds` attribute from the configuration of a `aws.ec2.DefaultNetworkAcl` resource may result in a reoccurring plan, until the Subnets are reassigned to another Network ACL or are destroyed.
 *
 * Because Subnets are by default associated with the Default Network ACL, any non-explicit association will show up as a plan to remove the Subnet. For example: if you have a custom `aws.ec2.NetworkAcl` with two subnets attached, and you remove the `aws.ec2.NetworkAcl` resource, after successfully destroying this resource future plans will show a diff on the managed `aws.ec2.DefaultNetworkAcl`, as those two Subnets have been orphaned by the now destroyed network acl and thus adopted by the Default Network ACL. In order to avoid a reoccurring plan, they will need to be reassigned, destroyed, or added to the `subnetIds` attribute of the `aws.ec2.DefaultNetworkAcl` entry.
 *
 * As an alternative to the above, you can also specify the following lifecycle configuration in your `aws.ec2.DefaultNetworkAcl` resource:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * // ... other configuration ...
 * const _default = new aws.ec2.DefaultNetworkAcl("default", {});
 * ```
 * ### Removing `aws.ec2.DefaultNetworkAcl` From Your Configuration
 *
 * Each AWS VPC comes with a Default Network ACL that cannot be deleted. The `aws.ec2.DefaultNetworkAcl` allows you to manage this Network ACL, but the provider cannot destroy it. Removing this resource from your configuration will remove it from your statefile and management, **but will not destroy the Network ACL.** All Subnets associations and ingress or egress rules will be left as they are at the time of removal. You can resume managing them via the AWS Console.
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_default_network_acl.sample
 *
 *  id = "acl-7aaabd18" } Using `pulumi import`, import Default Network ACLs using the `id`. For exampleconsole % pulumi import aws_default_network_acl.sample acl-7aaabd18
 */
export class DefaultNetworkAcl extends pulumi.CustomResource {
    /**
     * Get an existing DefaultNetworkAcl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DefaultNetworkAclState, opts?: pulumi.CustomResourceOptions): DefaultNetworkAcl {
        return new DefaultNetworkAcl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/defaultNetworkAcl:DefaultNetworkAcl';

    /**
     * Returns true if the given object is an instance of DefaultNetworkAcl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DefaultNetworkAcl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DefaultNetworkAcl.__pulumiType;
    }

    /**
     * ARN of the Default Network ACL
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Network ACL ID to manage. This attribute is exported from `aws.ec2.Vpc`, or manually found via the AWS Console.
     *
     * The following arguments are optional:
     */
    public readonly defaultNetworkAclId!: pulumi.Output<string>;
    /**
     * Configuration block for an egress rule. Detailed below.
     */
    public readonly egress!: pulumi.Output<outputs.ec2.DefaultNetworkAclEgress[] | undefined>;
    /**
     * Configuration block for an ingress rule. Detailed below.
     */
    public readonly ingress!: pulumi.Output<outputs.ec2.DefaultNetworkAclIngress[] | undefined>;
    /**
     * ID of the AWS account that owns the Default Network ACL
     */
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    /**
     * List of Subnet IDs to apply the ACL to. See the notes above on Managing Subnets in the Default Network ACL
     */
    public readonly subnetIds!: pulumi.Output<string[] | undefined>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * ID of the associated VPC
     */
    public /*out*/ readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a DefaultNetworkAcl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DefaultNetworkAclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DefaultNetworkAclArgs | DefaultNetworkAclState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DefaultNetworkAclState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["defaultNetworkAclId"] = state ? state.defaultNetworkAclId : undefined;
            resourceInputs["egress"] = state ? state.egress : undefined;
            resourceInputs["ingress"] = state ? state.ingress : undefined;
            resourceInputs["ownerId"] = state ? state.ownerId : undefined;
            resourceInputs["subnetIds"] = state ? state.subnetIds : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as DefaultNetworkAclArgs | undefined;
            if ((!args || args.defaultNetworkAclId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultNetworkAclId'");
            }
            resourceInputs["defaultNetworkAclId"] = args ? args.defaultNetworkAclId : undefined;
            resourceInputs["egress"] = args ? args.egress : undefined;
            resourceInputs["ingress"] = args ? args.ingress : undefined;
            resourceInputs["subnetIds"] = args ? args.subnetIds : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["ownerId"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DefaultNetworkAcl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DefaultNetworkAcl resources.
 */
export interface DefaultNetworkAclState {
    /**
     * ARN of the Default Network ACL
     */
    arn?: pulumi.Input<string>;
    /**
     * Network ACL ID to manage. This attribute is exported from `aws.ec2.Vpc`, or manually found via the AWS Console.
     *
     * The following arguments are optional:
     */
    defaultNetworkAclId?: pulumi.Input<string>;
    /**
     * Configuration block for an egress rule. Detailed below.
     */
    egress?: pulumi.Input<pulumi.Input<inputs.ec2.DefaultNetworkAclEgress>[]>;
    /**
     * Configuration block for an ingress rule. Detailed below.
     */
    ingress?: pulumi.Input<pulumi.Input<inputs.ec2.DefaultNetworkAclIngress>[]>;
    /**
     * ID of the AWS account that owns the Default Network ACL
     */
    ownerId?: pulumi.Input<string>;
    /**
     * List of Subnet IDs to apply the ACL to. See the notes above on Managing Subnets in the Default Network ACL
     */
    subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * ID of the associated VPC
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DefaultNetworkAcl resource.
 */
export interface DefaultNetworkAclArgs {
    /**
     * Network ACL ID to manage. This attribute is exported from `aws.ec2.Vpc`, or manually found via the AWS Console.
     *
     * The following arguments are optional:
     */
    defaultNetworkAclId: pulumi.Input<string>;
    /**
     * Configuration block for an egress rule. Detailed below.
     */
    egress?: pulumi.Input<pulumi.Input<inputs.ec2.DefaultNetworkAclEgress>[]>;
    /**
     * Configuration block for an ingress rule. Detailed below.
     */
    ingress?: pulumi.Input<pulumi.Input<inputs.ec2.DefaultNetworkAclIngress>[]>;
    /**
     * List of Subnet IDs to apply the ACL to. See the notes above on Managing Subnets in the Default Network ACL
     */
    subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
