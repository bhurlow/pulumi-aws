// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Route 53 Resolver query logging configuration association resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.route53.ResolverQueryLogConfigAssociation("example", {
 *     resolverQueryLogConfigId: aws_route53_resolver_query_log_config.example.id,
 *     resourceId: aws_vpc.example.id,
 * });
 * ```
 *
 * ## Import
 *
 *  Route 53 Resolver query logging configuration associations can be imported using the Route 53 Resolver query logging configuration association ID, e.g.
 *
 * ```sh
 *  $ pulumi import aws:route53/resolverQueryLogConfigAssociation:ResolverQueryLogConfigAssociation example rqlca-b320624fef3c4d70
 * ```
 */
export class ResolverQueryLogConfigAssociation extends pulumi.CustomResource {
    /**
     * Get an existing ResolverQueryLogConfigAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResolverQueryLogConfigAssociationState, opts?: pulumi.CustomResourceOptions): ResolverQueryLogConfigAssociation {
        return new ResolverQueryLogConfigAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:route53/resolverQueryLogConfigAssociation:ResolverQueryLogConfigAssociation';

    /**
     * Returns true if the given object is an instance of ResolverQueryLogConfigAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResolverQueryLogConfigAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResolverQueryLogConfigAssociation.__pulumiType;
    }

    /**
     * The ID of the Route 53 Resolver query logging configuration that you want to associate a VPC with.
     */
    public readonly resolverQueryLogConfigId!: pulumi.Output<string>;
    /**
     * The ID of a VPC that you want this query logging configuration to log queries for.
     */
    public readonly resourceId!: pulumi.Output<string>;

    /**
     * Create a ResolverQueryLogConfigAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResolverQueryLogConfigAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResolverQueryLogConfigAssociationArgs | ResolverQueryLogConfigAssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ResolverQueryLogConfigAssociationState | undefined;
            inputs["resolverQueryLogConfigId"] = state ? state.resolverQueryLogConfigId : undefined;
            inputs["resourceId"] = state ? state.resourceId : undefined;
        } else {
            const args = argsOrState as ResolverQueryLogConfigAssociationArgs | undefined;
            if ((!args || args.resolverQueryLogConfigId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resolverQueryLogConfigId'");
            }
            if ((!args || args.resourceId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceId'");
            }
            inputs["resolverQueryLogConfigId"] = args ? args.resolverQueryLogConfigId : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ResolverQueryLogConfigAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ResolverQueryLogConfigAssociation resources.
 */
export interface ResolverQueryLogConfigAssociationState {
    /**
     * The ID of the Route 53 Resolver query logging configuration that you want to associate a VPC with.
     */
    readonly resolverQueryLogConfigId?: pulumi.Input<string>;
    /**
     * The ID of a VPC that you want this query logging configuration to log queries for.
     */
    readonly resourceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ResolverQueryLogConfigAssociation resource.
 */
export interface ResolverQueryLogConfigAssociationArgs {
    /**
     * The ID of the Route 53 Resolver query logging configuration that you want to associate a VPC with.
     */
    readonly resolverQueryLogConfigId: pulumi.Input<string>;
    /**
     * The ID of a VPC that you want this query logging configuration to log queries for.
     */
    readonly resourceId: pulumi.Input<string>;
}
