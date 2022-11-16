// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an ElastiCache Security Group to control access to one or more cache
 * clusters.
 *
 * > **NOTE:** ElastiCache Security Groups are for use only when working with an
 * ElastiCache cluster **outside** of a VPC. If you are using a VPC, see the
 * ElastiCache Subnet Group resource.
 *
 * !> **WARNING:** With the retirement of EC2-Classic the `aws.elasticache.SecurityGroup` resource has been deprecated and will be removed in a future version.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const barSecurityGroup = new aws.ec2.SecurityGroup("barSecurityGroup", {});
 * const barElasticache_securityGroupSecurityGroup = new aws.elasticache.SecurityGroup("barElasticache/securityGroupSecurityGroup", {securityGroupNames: [barSecurityGroup.name]});
 * ```
 *
 * ## Import
 *
 * ElastiCache Security Groups can be imported by name, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:elasticache/securityGroup:SecurityGroup my_ec_security_group ec-security-group-1
 * ```
 */
export class SecurityGroup extends pulumi.CustomResource {
    /**
     * Get an existing SecurityGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityGroupState, opts?: pulumi.CustomResourceOptions): SecurityGroup {
        return new SecurityGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:elasticache/securityGroup:SecurityGroup';

    /**
     * Returns true if the given object is an instance of SecurityGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityGroup.__pulumiType;
    }

    /**
     * description for the cache security group. Defaults to "Managed by Pulumi".
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Name for the cache security group. This value is stored as a lowercase string.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * List of EC2 security group names to be
     * authorized for ingress to the cache security group
     */
    public readonly securityGroupNames!: pulumi.Output<string[]>;

    /**
     * Create a SecurityGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityGroupArgs | SecurityGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityGroupState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["securityGroupNames"] = state ? state.securityGroupNames : undefined;
        } else {
            const args = argsOrState as SecurityGroupArgs | undefined;
            if ((!args || args.securityGroupNames === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupNames'");
            }
            resourceInputs["description"] = (args ? args.description : undefined) ?? "Managed by Pulumi";
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["securityGroupNames"] = args ? args.securityGroupNames : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecurityGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityGroup resources.
 */
export interface SecurityGroupState {
    /**
     * description for the cache security group. Defaults to "Managed by Pulumi".
     */
    description?: pulumi.Input<string>;
    /**
     * Name for the cache security group. This value is stored as a lowercase string.
     */
    name?: pulumi.Input<string>;
    /**
     * List of EC2 security group names to be
     * authorized for ingress to the cache security group
     */
    securityGroupNames?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a SecurityGroup resource.
 */
export interface SecurityGroupArgs {
    /**
     * description for the cache security group. Defaults to "Managed by Pulumi".
     */
    description?: pulumi.Input<string>;
    /**
     * Name for the cache security group. This value is stored as a lowercase string.
     */
    name?: pulumi.Input<string>;
    /**
     * List of EC2 security group names to be
     * authorized for ingress to the cache security group
     */
    securityGroupNames: pulumi.Input<pulumi.Input<string>[]>;
}
