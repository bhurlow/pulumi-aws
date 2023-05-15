// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an IAM OpenID Connect provider.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const _default = new aws.iam.OpenIdConnectProvider("default", {
 *     clientIdLists: ["266362248691-342342xasdasdasda-apps.googleusercontent.com"],
 *     thumbprintLists: ["cf23df2207d99a74fbe169e3eba035e633b65d94"],
 *     url: "https://accounts.google.com",
 * });
 * ```
 *
 * ## Import
 *
 * IAM OpenID Connect Providers can be imported using the `arn`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:iam/openIdConnectProvider:OpenIdConnectProvider default arn:aws:iam::123456789012:oidc-provider/accounts.google.com
 * ```
 */
export class OpenIdConnectProvider extends pulumi.CustomResource {
    /**
     * Get an existing OpenIdConnectProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OpenIdConnectProviderState, opts?: pulumi.CustomResourceOptions): OpenIdConnectProvider {
        return new OpenIdConnectProvider(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:iam/openIdConnectProvider:OpenIdConnectProvider';

    /**
     * Returns true if the given object is an instance of OpenIdConnectProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OpenIdConnectProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OpenIdConnectProvider.__pulumiType;
    }

    /**
     * The ARN assigned by AWS for this provider.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the clientId parameter on OAuth requests.)
     */
    public readonly clientIdLists!: pulumi.Output<string[]>;
    /**
     * Map of resource tags for the IAM OIDC provider. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
     */
    public readonly thumbprintLists!: pulumi.Output<string[]>;
    /**
     * The URL of the identity provider. Corresponds to the _iss_ claim.
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a OpenIdConnectProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OpenIdConnectProviderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OpenIdConnectProviderArgs | OpenIdConnectProviderState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OpenIdConnectProviderState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["clientIdLists"] = state ? state.clientIdLists : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["thumbprintLists"] = state ? state.thumbprintLists : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as OpenIdConnectProviderArgs | undefined;
            if ((!args || args.clientIdLists === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientIdLists'");
            }
            if ((!args || args.thumbprintLists === undefined) && !opts.urn) {
                throw new Error("Missing required property 'thumbprintLists'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["clientIdLists"] = args ? args.clientIdLists : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["thumbprintLists"] = args ? args.thumbprintLists : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OpenIdConnectProvider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OpenIdConnectProvider resources.
 */
export interface OpenIdConnectProviderState {
    /**
     * The ARN assigned by AWS for this provider.
     */
    arn?: pulumi.Input<string>;
    /**
     * A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the clientId parameter on OAuth requests.)
     */
    clientIdLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Map of resource tags for the IAM OIDC provider. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
     */
    thumbprintLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The URL of the identity provider. Corresponds to the _iss_ claim.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OpenIdConnectProvider resource.
 */
export interface OpenIdConnectProviderArgs {
    /**
     * A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the clientId parameter on OAuth requests.)
     */
    clientIdLists: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Map of resource tags for the IAM OIDC provider. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
     */
    thumbprintLists: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The URL of the identity provider. Corresponds to the _iss_ claim.
     */
    url: pulumi.Input<string>;
}
