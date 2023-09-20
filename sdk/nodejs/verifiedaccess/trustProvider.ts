// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource for managing a Verified Access Trust Provider.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.verifiedaccess.TrustProvider("example", {
 *     policyReferenceName: "example",
 *     trustProviderType: "user",
 *     userTrustProviderType: "iam-identity-center",
 * });
 * ```
 *
 * ## Import
 *
 * In TODO v1.5.0 and later, use an `import` block to import Transfer Workflows using the `id`. For exampleterraform import {
 *
 *  to = aws_verifiedaccess_trust_provider.example
 *
 *  id = "vatp-8012925589" } Using `TODO import`, import Transfer Workflows using the
 *
 * `id`. For exampleconsole % TODO import aws_verifiedaccess_trust_provider.example vatp-8012925589
 */
export class TrustProvider extends pulumi.CustomResource {
    /**
     * Get an existing TrustProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TrustProviderState, opts?: pulumi.CustomResourceOptions): TrustProvider {
        return new TrustProvider(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:verifiedaccess/trustProvider:TrustProvider';

    /**
     * Returns true if the given object is an instance of TrustProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TrustProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TrustProvider.__pulumiType;
    }

    /**
     * A description for the AWS Verified Access trust provider.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A block of options for device identity based trust providers.
     */
    public readonly deviceOptions!: pulumi.Output<outputs.verifiedaccess.TrustProviderDeviceOptions | undefined>;
    /**
     * The type of device-based trust provider.
     */
    public readonly deviceTrustProviderType!: pulumi.Output<string | undefined>;
    /**
     * The OpenID Connect details for an oidc-type, user-identity based trust provider.
     */
    public readonly oidcOptions!: pulumi.Output<outputs.verifiedaccess.TrustProviderOidcOptions | undefined>;
    /**
     * The identifier to be used when working with policy rules.
     */
    public readonly policyReferenceName!: pulumi.Output<string>;
    /**
     * Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * @deprecated Please use `tags` instead.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The type of trust provider can be either user or device-based.
     *
     * The following arguments are optional:
     */
    public readonly trustProviderType!: pulumi.Output<string>;
    /**
     * The type of user-based trust provider.
     */
    public readonly userTrustProviderType!: pulumi.Output<string | undefined>;

    /**
     * Create a TrustProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TrustProviderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TrustProviderArgs | TrustProviderState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TrustProviderState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["deviceOptions"] = state ? state.deviceOptions : undefined;
            resourceInputs["deviceTrustProviderType"] = state ? state.deviceTrustProviderType : undefined;
            resourceInputs["oidcOptions"] = state ? state.oidcOptions : undefined;
            resourceInputs["policyReferenceName"] = state ? state.policyReferenceName : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["trustProviderType"] = state ? state.trustProviderType : undefined;
            resourceInputs["userTrustProviderType"] = state ? state.userTrustProviderType : undefined;
        } else {
            const args = argsOrState as TrustProviderArgs | undefined;
            if ((!args || args.policyReferenceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyReferenceName'");
            }
            if ((!args || args.trustProviderType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'trustProviderType'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["deviceOptions"] = args ? args.deviceOptions : undefined;
            resourceInputs["deviceTrustProviderType"] = args ? args.deviceTrustProviderType : undefined;
            resourceInputs["oidcOptions"] = args ? args.oidcOptions : undefined;
            resourceInputs["policyReferenceName"] = args ? args.policyReferenceName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["trustProviderType"] = args ? args.trustProviderType : undefined;
            resourceInputs["userTrustProviderType"] = args ? args.userTrustProviderType : undefined;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(TrustProvider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TrustProvider resources.
 */
export interface TrustProviderState {
    /**
     * A description for the AWS Verified Access trust provider.
     */
    description?: pulumi.Input<string>;
    /**
     * A block of options for device identity based trust providers.
     */
    deviceOptions?: pulumi.Input<inputs.verifiedaccess.TrustProviderDeviceOptions>;
    /**
     * The type of device-based trust provider.
     */
    deviceTrustProviderType?: pulumi.Input<string>;
    /**
     * The OpenID Connect details for an oidc-type, user-identity based trust provider.
     */
    oidcOptions?: pulumi.Input<inputs.verifiedaccess.TrustProviderOidcOptions>;
    /**
     * The identifier to be used when working with policy rules.
     */
    policyReferenceName?: pulumi.Input<string>;
    /**
     * Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * @deprecated Please use `tags` instead.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of trust provider can be either user or device-based.
     *
     * The following arguments are optional:
     */
    trustProviderType?: pulumi.Input<string>;
    /**
     * The type of user-based trust provider.
     */
    userTrustProviderType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TrustProvider resource.
 */
export interface TrustProviderArgs {
    /**
     * A description for the AWS Verified Access trust provider.
     */
    description?: pulumi.Input<string>;
    /**
     * A block of options for device identity based trust providers.
     */
    deviceOptions?: pulumi.Input<inputs.verifiedaccess.TrustProviderDeviceOptions>;
    /**
     * The type of device-based trust provider.
     */
    deviceTrustProviderType?: pulumi.Input<string>;
    /**
     * The OpenID Connect details for an oidc-type, user-identity based trust provider.
     */
    oidcOptions?: pulumi.Input<inputs.verifiedaccess.TrustProviderOidcOptions>;
    /**
     * The identifier to be used when working with policy rules.
     */
    policyReferenceName: pulumi.Input<string>;
    /**
     * Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of trust provider can be either user or device-based.
     *
     * The following arguments are optional:
     */
    trustProviderType: pulumi.Input<string>;
    /**
     * The type of user-based trust provider.
     */
    userTrustProviderType?: pulumi.Input<string>;
}
