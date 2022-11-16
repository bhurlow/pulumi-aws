// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create an association between a VPC endpoint and a security group.
 *
 * > **NOTE on VPC Endpoints and VPC Endpoint Security Group Associations:** The provider provides
 * both a standalone VPC Endpoint Security Group Association (an association between a VPC endpoint
 * and a single `securityGroupId`) and a VPC Endpoint resource with a `securityGroupIds`
 * attribute. Do not use the same security group ID in both a VPC Endpoint resource and a VPC Endpoint Security
 * Group Association resource. Doing so will cause a conflict of associations and will overwrite the association.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const sgEc2 = new aws.ec2.SecurityGroupAssociation("sgEc2", {
 *     vpcEndpointId: aws_vpc_endpoint.ec2.id,
 *     securityGroupId: aws_security_group.sg.id,
 * });
 * ```
 */
export class SecurityGroupAssociation extends pulumi.CustomResource {
    /**
     * Get an existing SecurityGroupAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityGroupAssociationState, opts?: pulumi.CustomResourceOptions): SecurityGroupAssociation {
        return new SecurityGroupAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/securityGroupAssociation:SecurityGroupAssociation';

    /**
     * Returns true if the given object is an instance of SecurityGroupAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityGroupAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityGroupAssociation.__pulumiType;
    }

    /**
     * Whether this association should replace the association with the VPC's default security group that is created when no security groups are specified during VPC endpoint creation. At most 1 association per-VPC endpoint should be configured with `replaceDefaultAssociation = true`.
     */
    public readonly replaceDefaultAssociation!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the security group to be associated with the VPC endpoint.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * The ID of the VPC endpoint with which the security group will be associated.
     */
    public readonly vpcEndpointId!: pulumi.Output<string>;

    /**
     * Create a SecurityGroupAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityGroupAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityGroupAssociationArgs | SecurityGroupAssociationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityGroupAssociationState | undefined;
            resourceInputs["replaceDefaultAssociation"] = state ? state.replaceDefaultAssociation : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            resourceInputs["vpcEndpointId"] = state ? state.vpcEndpointId : undefined;
        } else {
            const args = argsOrState as SecurityGroupAssociationArgs | undefined;
            if ((!args || args.securityGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupId'");
            }
            if ((!args || args.vpcEndpointId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcEndpointId'");
            }
            resourceInputs["replaceDefaultAssociation"] = args ? args.replaceDefaultAssociation : undefined;
            resourceInputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            resourceInputs["vpcEndpointId"] = args ? args.vpcEndpointId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecurityGroupAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityGroupAssociation resources.
 */
export interface SecurityGroupAssociationState {
    /**
     * Whether this association should replace the association with the VPC's default security group that is created when no security groups are specified during VPC endpoint creation. At most 1 association per-VPC endpoint should be configured with `replaceDefaultAssociation = true`.
     */
    replaceDefaultAssociation?: pulumi.Input<boolean>;
    /**
     * The ID of the security group to be associated with the VPC endpoint.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * The ID of the VPC endpoint with which the security group will be associated.
     */
    vpcEndpointId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityGroupAssociation resource.
 */
export interface SecurityGroupAssociationArgs {
    /**
     * Whether this association should replace the association with the VPC's default security group that is created when no security groups are specified during VPC endpoint creation. At most 1 association per-VPC endpoint should be configured with `replaceDefaultAssociation = true`.
     */
    replaceDefaultAssociation?: pulumi.Input<boolean>;
    /**
     * The ID of the security group to be associated with the VPC endpoint.
     */
    securityGroupId: pulumi.Input<string>;
    /**
     * The ID of the VPC endpoint with which the security group will be associated.
     */
    vpcEndpointId: pulumi.Input<string>;
}
