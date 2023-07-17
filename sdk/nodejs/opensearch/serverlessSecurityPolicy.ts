// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource for managing an AWS OpenSearch Serverless Security Policy. See AWS documentation for [encryption policies](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-encryption.html#serverless-encryption-policies) and [network policies](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-network.html#serverless-network-policies).
 *
 * ## Example Usage
 *
 * ### Encryption Security Policy
 * ### Applies to a single collection
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.opensearch.ServerlessSecurityPolicy("example", {
 *     name: "example",
 *     type: "encryption",
 *     description: "encryption security policy for example-collection",
 *     policy: JSON.stringify({
 *         Rules: [{
 *             Resource: ["collection/example-collection"],
 *             ResourceType: "collection",
 *         }],
 *         AWSOwnedKey: true,
 *     }),
 * });
 * ```
 * ### Applies to multiple collections
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.opensearch.ServerlessSecurityPolicy("example", {
 *     name: "example",
 *     type: "encryption",
 *     description: "encryption security policy for collections that begin with \"example\"",
 *     policy: JSON.stringify({
 *         Rules: [{
 *             Resource: ["collection/example*"],
 *             ResourceType: "collection",
 *         }],
 *         AWSOwnedKey: true,
 *     }),
 * });
 * ```
 * ### Using a customer managed key
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.opensearch.ServerlessSecurityPolicy("example", {
 *     name: "example",
 *     type: "encryption",
 *     description: "encryption security policy using customer KMS key",
 *     policy: JSON.stringify({
 *         Rules: [{
 *             Resource: ["collection/customer-managed-key-collection"],
 *             ResourceType: "collection",
 *         }],
 *         AWSOwnedKey: false,
 *         KmsARN: "arn:aws:kms:us-east-1:123456789012:key/93fd6da4-a317-4c17-bfe9-382b5d988b36",
 *     }),
 * });
 * ```
 * ### Network Security Policy
 * ### Allow public access to the collection endpoint and the Dashboards endpoint
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.opensearch.ServerlessSecurityPolicy("example", {
 *     name: "example",
 *     type: "network",
 *     description: "Public access",
 *     policy: JSON.stringify([{
 *         Description: "Public access to collection and Dashboards endpoint for example collection",
 *         Rules: [
 *             {
 *                 ResourceType: "collection",
 *                 Resource: ["collection/example-collection"],
 *             },
 *             {
 *                 ResourceType: "dashboard",
 *                 Resource: ["collection/example-collection"],
 *             },
 *         ],
 *         AllowFromPublic: true,
 *     }]),
 * });
 * ```
 * ### Allow VPC access to the collection endpoint and the Dashboards endpoint
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.opensearch.ServerlessSecurityPolicy("example", {
 *     name: "example",
 *     type: "network",
 *     description: "VPC access",
 *     policy: JSON.stringify([{
 *         Description: "VPC access to collection and Dashboards endpoint for example collection",
 *         Rules: [
 *             {
 *                 ResourceType: "collection",
 *                 Resource: ["collection/example-collection"],
 *             },
 *             {
 *                 ResourceType: "dashboard",
 *                 Resource: ["collection/example-collection"],
 *             },
 *         ],
 *         AllowFromPublic: false,
 *         SourceVPCEs: ["vpce-050f79086ee71ac05"],
 *     }]),
 * });
 * ```
 * ### Mixed access for different collections
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.opensearch.ServerlessSecurityPolicy("example", {
 *     name: "example",
 *     type: "network",
 *     description: "Mixed access for marketing and sales",
 *     policy: JSON.stringify([
 *         {
 *             Description: "Marketing access",
 *             Rules: [
 *                 {
 *                     ResourceType: "collection",
 *                     Resource: ["collection/marketing*"],
 *                 },
 *                 {
 *                     ResourceType: "dashboard",
 *                     Resource: ["collection/marketing*"],
 *                 },
 *             ],
 *             AllowFromPublic: false,
 *             SourceVPCEs: ["vpce-050f79086ee71ac05"],
 *         },
 *         {
 *             Description: "Sales access",
 *             Rules: [{
 *                 ResourceType: "collection",
 *                 Resource: ["collection/finance"],
 *             }],
 *             AllowFromPublic: true,
 *         },
 *     ]),
 * });
 * ```
 *
 * ## Import
 *
 * OpenSearchServerless Security Policy can be imported using the `name` and `type` arguments separated by a slash (`/`), e.g.,
 *
 * ```sh
 *  $ pulumi import aws:opensearch/serverlessSecurityPolicy:ServerlessSecurityPolicy example example/encryption
 * ```
 */
export class ServerlessSecurityPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ServerlessSecurityPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerlessSecurityPolicyState, opts?: pulumi.CustomResourceOptions): ServerlessSecurityPolicy {
        return new ServerlessSecurityPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:opensearch/serverlessSecurityPolicy:ServerlessSecurityPolicy';

    /**
     * Returns true if the given object is an instance of ServerlessSecurityPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerlessSecurityPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerlessSecurityPolicy.__pulumiType;
    }

    /**
     * Description of the policy. Typically used to store information about the permissions defined in the policy.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the policy.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * JSON policy document to use as the content for the new policy
     */
    public readonly policy!: pulumi.Output<string>;
    /**
     * Version of the policy.
     */
    public /*out*/ readonly policyVersion!: pulumi.Output<string>;
    /**
     * Type of security policy. One of `encryption` or `network`.
     *
     * The following arguments are optional:
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a ServerlessSecurityPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerlessSecurityPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerlessSecurityPolicyArgs | ServerlessSecurityPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServerlessSecurityPolicyState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["policy"] = state ? state.policy : undefined;
            resourceInputs["policyVersion"] = state ? state.policyVersion : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ServerlessSecurityPolicyArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.policy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policy'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["policy"] = args ? args.policy : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["policyVersion"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServerlessSecurityPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServerlessSecurityPolicy resources.
 */
export interface ServerlessSecurityPolicyState {
    /**
     * Description of the policy. Typically used to store information about the permissions defined in the policy.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the policy.
     */
    name?: pulumi.Input<string>;
    /**
     * JSON policy document to use as the content for the new policy
     */
    policy?: pulumi.Input<string>;
    /**
     * Version of the policy.
     */
    policyVersion?: pulumi.Input<string>;
    /**
     * Type of security policy. One of `encryption` or `network`.
     *
     * The following arguments are optional:
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServerlessSecurityPolicy resource.
 */
export interface ServerlessSecurityPolicyArgs {
    /**
     * Description of the policy. Typically used to store information about the permissions defined in the policy.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the policy.
     */
    name: pulumi.Input<string>;
    /**
     * JSON policy document to use as the content for the new policy
     */
    policy: pulumi.Input<string>;
    /**
     * Type of security policy. One of `encryption` or `network`.
     *
     * The following arguments are optional:
     */
    type: pulumi.Input<string>;
}
