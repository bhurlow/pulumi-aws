// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides authorization rules for AWS Client VPN endpoints. For more information on usage, please see the
 * [AWS Client VPN Administrator's Guide](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/what-is.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.ec2clientvpn.AuthorizationRule("example", {
 *     clientVpnEndpointId: aws_ec2_client_vpn_endpoint.example.id,
 *     targetNetworkCidr: aws_subnet.example.cidr_block,
 *     authorizeAllGroups: true,
 * });
 * ```
 *
 * ## Import
 *
 * Using the endpoint ID and target network CIDRterraform import {
 *
 *  to = aws_ec2_client_vpn_authorization_rule.example
 *
 *  id = "cvpn-endpoint-0ac3a1abbccddd666,10.1.0.0/24" } Using the endpoint ID, target network CIDR, and group nameterraform import {
 *
 *  to = aws_ec2_client_vpn_authorization_rule.example
 *
 *  id = "cvpn-endpoint-0ac3a1abbccddd666,10.1.0.0/24,team-a" } **Using `pulumi import` to import** AWS Client VPN authorization rules using the endpoint ID and target network CIDR. If there is a specific group name, include that also. All values are separated by a `,`. For exampleUsing the endpoint ID and target network CIDRconsole % pulumi import aws_ec2_client_vpn_authorization_rule.example cvpn-endpoint-0ac3a1abbccddd666,10.1.0.0/24 Using the endpoint ID, target network CIDR, and group nameconsole % pulumi import aws_ec2_client_vpn_authorization_rule.example cvpn-endpoint-0ac3a1abbccddd666,10.1.0.0/24,team-a
 */
export class AuthorizationRule extends pulumi.CustomResource {
    /**
     * Get an existing AuthorizationRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthorizationRuleState, opts?: pulumi.CustomResourceOptions): AuthorizationRule {
        return new AuthorizationRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2clientvpn/authorizationRule:AuthorizationRule';

    /**
     * Returns true if the given object is an instance of AuthorizationRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthorizationRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthorizationRule.__pulumiType;
    }

    /**
     * The ID of the group to which the authorization rule grants access. One of `accessGroupId` or `authorizeAllGroups` must be set.
     */
    public readonly accessGroupId!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether the authorization rule grants access to all clients. One of `accessGroupId` or `authorizeAllGroups` must be set.
     */
    public readonly authorizeAllGroups!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the Client VPN endpoint.
     */
    public readonly clientVpnEndpointId!: pulumi.Output<string>;
    /**
     * A brief description of the authorization rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The IPv4 address range, in CIDR notation, of the network to which the authorization rule applies.
     */
    public readonly targetNetworkCidr!: pulumi.Output<string>;

    /**
     * Create a AuthorizationRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthorizationRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthorizationRuleArgs | AuthorizationRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthorizationRuleState | undefined;
            resourceInputs["accessGroupId"] = state ? state.accessGroupId : undefined;
            resourceInputs["authorizeAllGroups"] = state ? state.authorizeAllGroups : undefined;
            resourceInputs["clientVpnEndpointId"] = state ? state.clientVpnEndpointId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["targetNetworkCidr"] = state ? state.targetNetworkCidr : undefined;
        } else {
            const args = argsOrState as AuthorizationRuleArgs | undefined;
            if ((!args || args.clientVpnEndpointId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientVpnEndpointId'");
            }
            if ((!args || args.targetNetworkCidr === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetNetworkCidr'");
            }
            resourceInputs["accessGroupId"] = args ? args.accessGroupId : undefined;
            resourceInputs["authorizeAllGroups"] = args ? args.authorizeAllGroups : undefined;
            resourceInputs["clientVpnEndpointId"] = args ? args.clientVpnEndpointId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["targetNetworkCidr"] = args ? args.targetNetworkCidr : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AuthorizationRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthorizationRule resources.
 */
export interface AuthorizationRuleState {
    /**
     * The ID of the group to which the authorization rule grants access. One of `accessGroupId` or `authorizeAllGroups` must be set.
     */
    accessGroupId?: pulumi.Input<string>;
    /**
     * Indicates whether the authorization rule grants access to all clients. One of `accessGroupId` or `authorizeAllGroups` must be set.
     */
    authorizeAllGroups?: pulumi.Input<boolean>;
    /**
     * The ID of the Client VPN endpoint.
     */
    clientVpnEndpointId?: pulumi.Input<string>;
    /**
     * A brief description of the authorization rule.
     */
    description?: pulumi.Input<string>;
    /**
     * The IPv4 address range, in CIDR notation, of the network to which the authorization rule applies.
     */
    targetNetworkCidr?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthorizationRule resource.
 */
export interface AuthorizationRuleArgs {
    /**
     * The ID of the group to which the authorization rule grants access. One of `accessGroupId` or `authorizeAllGroups` must be set.
     */
    accessGroupId?: pulumi.Input<string>;
    /**
     * Indicates whether the authorization rule grants access to all clients. One of `accessGroupId` or `authorizeAllGroups` must be set.
     */
    authorizeAllGroups?: pulumi.Input<boolean>;
    /**
     * The ID of the Client VPN endpoint.
     */
    clientVpnEndpointId: pulumi.Input<string>;
    /**
     * A brief description of the authorization rule.
     */
    description?: pulumi.Input<string>;
    /**
     * The IPv4 address range, in CIDR notation, of the network to which the authorization rule applies.
     */
    targetNetworkCidr: pulumi.Input<string>;
}
