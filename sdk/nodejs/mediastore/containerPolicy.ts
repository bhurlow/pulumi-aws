// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a MediaStore Container Policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const currentRegion = aws.getRegion({});
 * const currentCallerIdentity = aws.getCallerIdentity({});
 * const exampleContainer = new aws.mediastore.Container("exampleContainer", {});
 * const examplePolicyDocument = aws.iam.getPolicyDocumentOutput({
 *     statements: [{
 *         sid: "MediaStoreFullAccess",
 *         effect: "Allow",
 *         principals: [{
 *             type: "AWS",
 *             identifiers: [currentCallerIdentity.then(currentCallerIdentity => `arn:aws:iam::${currentCallerIdentity.accountId}:root`)],
 *         }],
 *         actions: ["mediastore:*"],
 *         resources: [pulumi.all([currentRegion, currentCallerIdentity, exampleContainer.name]).apply(([currentRegion, currentCallerIdentity, name]) => `arn:aws:mediastore:${currentRegion.name}:${currentCallerIdentity.accountId}:container/${name}/*`)],
 *         conditions: [{
 *             test: "Bool",
 *             variable: "aws:SecureTransport",
 *             values: ["true"],
 *         }],
 *     }],
 * });
 * const exampleContainerPolicy = new aws.mediastore.ContainerPolicy("exampleContainerPolicy", {
 *     containerName: exampleContainer.name,
 *     policy: examplePolicyDocument.apply(examplePolicyDocument => examplePolicyDocument.json),
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_media_store_container_policy.example
 *
 *  id = "example" } Using `pulumi import`, import MediaStore Container Policy using the MediaStore Container Name. For exampleconsole % pulumi import aws_media_store_container_policy.example example
 */
export class ContainerPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ContainerPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContainerPolicyState, opts?: pulumi.CustomResourceOptions): ContainerPolicy {
        return new ContainerPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:mediastore/containerPolicy:ContainerPolicy';

    /**
     * Returns true if the given object is an instance of ContainerPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContainerPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContainerPolicy.__pulumiType;
    }

    /**
     * The name of the container.
     */
    public readonly containerName!: pulumi.Output<string>;
    /**
     * The contents of the policy.
     */
    public readonly policy!: pulumi.Output<string>;

    /**
     * Create a ContainerPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContainerPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContainerPolicyArgs | ContainerPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContainerPolicyState | undefined;
            resourceInputs["containerName"] = state ? state.containerName : undefined;
            resourceInputs["policy"] = state ? state.policy : undefined;
        } else {
            const args = argsOrState as ContainerPolicyArgs | undefined;
            if ((!args || args.containerName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'containerName'");
            }
            if ((!args || args.policy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policy'");
            }
            resourceInputs["containerName"] = args ? args.containerName : undefined;
            resourceInputs["policy"] = args ? args.policy : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ContainerPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ContainerPolicy resources.
 */
export interface ContainerPolicyState {
    /**
     * The name of the container.
     */
    containerName?: pulumi.Input<string>;
    /**
     * The contents of the policy.
     */
    policy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ContainerPolicy resource.
 */
export interface ContainerPolicyArgs {
    /**
     * The name of the container.
     */
    containerName: pulumi.Input<string>;
    /**
     * The contents of the policy.
     */
    policy: pulumi.Input<string>;
}
