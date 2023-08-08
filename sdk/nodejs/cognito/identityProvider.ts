// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Cognito User Identity Provider resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.cognito.UserPool("example", {autoVerifiedAttributes: ["email"]});
 * const exampleProvider = new aws.cognito.IdentityProvider("exampleProvider", {
 *     userPoolId: example.id,
 *     providerName: "Google",
 *     providerType: "Google",
 *     providerDetails: {
 *         authorize_scopes: "email",
 *         client_id: "your client_id",
 *         client_secret: "your client_secret",
 *     },
 *     attributeMapping: {
 *         email: "email",
 *         username: "sub",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_cognito_identity_provider.example
 *
 *  id = "us-west-2_abc123:CorpAD" } Using `pulumi import`, import `aws_cognito_identity_provider` resources using their User Pool ID and Provider Name. For exampleconsole % pulumi import aws_cognito_identity_provider.example us-west-2_abc123:CorpAD
 */
export class IdentityProvider extends pulumi.CustomResource {
    /**
     * Get an existing IdentityProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IdentityProviderState, opts?: pulumi.CustomResourceOptions): IdentityProvider {
        return new IdentityProvider(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cognito/identityProvider:IdentityProvider';

    /**
     * Returns true if the given object is an instance of IdentityProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IdentityProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IdentityProvider.__pulumiType;
    }

    /**
     * The map of attribute mapping of user pool attributes. [AttributeMapping in AWS API documentation](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateIdentityProvider.html#CognitoUserPools-CreateIdentityProvider-request-AttributeMapping)
     */
    public readonly attributeMapping!: pulumi.Output<{[key: string]: string}>;
    /**
     * The list of identity providers.
     */
    public readonly idpIdentifiers!: pulumi.Output<string[] | undefined>;
    /**
     * The map of identity details, such as access token
     */
    public readonly providerDetails!: pulumi.Output<{[key: string]: string}>;
    /**
     * The provider name
     */
    public readonly providerName!: pulumi.Output<string>;
    /**
     * The provider type.  [See AWS API for valid values](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateIdentityProvider.html#CognitoUserPools-CreateIdentityProvider-request-ProviderType)
     */
    public readonly providerType!: pulumi.Output<string>;
    /**
     * The user pool id
     */
    public readonly userPoolId!: pulumi.Output<string>;

    /**
     * Create a IdentityProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IdentityProviderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IdentityProviderArgs | IdentityProviderState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IdentityProviderState | undefined;
            resourceInputs["attributeMapping"] = state ? state.attributeMapping : undefined;
            resourceInputs["idpIdentifiers"] = state ? state.idpIdentifiers : undefined;
            resourceInputs["providerDetails"] = state ? state.providerDetails : undefined;
            resourceInputs["providerName"] = state ? state.providerName : undefined;
            resourceInputs["providerType"] = state ? state.providerType : undefined;
            resourceInputs["userPoolId"] = state ? state.userPoolId : undefined;
        } else {
            const args = argsOrState as IdentityProviderArgs | undefined;
            if ((!args || args.providerDetails === undefined) && !opts.urn) {
                throw new Error("Missing required property 'providerDetails'");
            }
            if ((!args || args.providerName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'providerName'");
            }
            if ((!args || args.providerType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'providerType'");
            }
            if ((!args || args.userPoolId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userPoolId'");
            }
            resourceInputs["attributeMapping"] = args ? args.attributeMapping : undefined;
            resourceInputs["idpIdentifiers"] = args ? args.idpIdentifiers : undefined;
            resourceInputs["providerDetails"] = args ? args.providerDetails : undefined;
            resourceInputs["providerName"] = args ? args.providerName : undefined;
            resourceInputs["providerType"] = args ? args.providerType : undefined;
            resourceInputs["userPoolId"] = args ? args.userPoolId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IdentityProvider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IdentityProvider resources.
 */
export interface IdentityProviderState {
    /**
     * The map of attribute mapping of user pool attributes. [AttributeMapping in AWS API documentation](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateIdentityProvider.html#CognitoUserPools-CreateIdentityProvider-request-AttributeMapping)
     */
    attributeMapping?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The list of identity providers.
     */
    idpIdentifiers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The map of identity details, such as access token
     */
    providerDetails?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The provider name
     */
    providerName?: pulumi.Input<string>;
    /**
     * The provider type.  [See AWS API for valid values](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateIdentityProvider.html#CognitoUserPools-CreateIdentityProvider-request-ProviderType)
     */
    providerType?: pulumi.Input<string>;
    /**
     * The user pool id
     */
    userPoolId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IdentityProvider resource.
 */
export interface IdentityProviderArgs {
    /**
     * The map of attribute mapping of user pool attributes. [AttributeMapping in AWS API documentation](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateIdentityProvider.html#CognitoUserPools-CreateIdentityProvider-request-AttributeMapping)
     */
    attributeMapping?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The list of identity providers.
     */
    idpIdentifiers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The map of identity details, such as access token
     */
    providerDetails: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The provider name
     */
    providerName: pulumi.Input<string>;
    /**
     * The provider type.  [See AWS API for valid values](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateIdentityProvider.html#CognitoUserPools-CreateIdentityProvider-request-ProviderType)
     */
    providerType: pulumi.Input<string>;
    /**
     * The user pool id
     */
    userPoolId: pulumi.Input<string>;
}
