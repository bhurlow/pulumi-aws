// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an Amazon Connect Security Profile resource. For more information see
 * [Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.connect.SecurityProfile("example", {
 *     description: "example description",
 *     instanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
 *     permissions: [
 *         "BasicAgentAccess",
 *         "OutboundCallAccess",
 *     ],
 *     tags: {
 *         Name: "Example Security Profile",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Amazon Connect Security Profiles can be imported using the `instance_id` and `security_profile_id` separated by a colon (`:`), e.g.,
 *
 * ```sh
 *  $ pulumi import aws:connect/securityProfile:SecurityProfile example f1288a1f-6193-445a-b47e-af739b2:c1d4e5f6-1b3c-1b3c-1b3c-c1d4e5f6c1d4e5
 * ```
 */
export class SecurityProfile extends pulumi.CustomResource {
    /**
     * Get an existing SecurityProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityProfileState, opts?: pulumi.CustomResourceOptions): SecurityProfile {
        return new SecurityProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:connect/securityProfile:SecurityProfile';

    /**
     * Returns true if the given object is an instance of SecurityProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityProfile.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the Security Profile.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Specifies the description of the Security Profile.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the identifier of the hosting Amazon Connect Instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Specifies the name of the Security Profile.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The organization resource identifier for the security profile.
     */
    public /*out*/ readonly organizationResourceId!: pulumi.Output<string>;
    /**
     * Specifies a list of permissions assigned to the security profile.
     */
    public readonly permissions!: pulumi.Output<string[] | undefined>;
    /**
     * The identifier for the Security Profile.
     */
    public /*out*/ readonly securityProfileId!: pulumi.Output<string>;
    /**
     * Tags to apply to the Security Profile. If configured with a provider
     * [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a SecurityProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityProfileArgs | SecurityProfileState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityProfileState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["organizationResourceId"] = state ? state.organizationResourceId : undefined;
            inputs["permissions"] = state ? state.permissions : undefined;
            inputs["securityProfileId"] = state ? state.securityProfileId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as SecurityProfileArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["permissions"] = args ? args.permissions : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["organizationResourceId"] = undefined /*out*/;
            inputs["securityProfileId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SecurityProfile.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityProfile resources.
 */
export interface SecurityProfileState {
    /**
     * The Amazon Resource Name (ARN) of the Security Profile.
     */
    arn?: pulumi.Input<string>;
    /**
     * Specifies the description of the Security Profile.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the identifier of the hosting Amazon Connect Instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Specifies the name of the Security Profile.
     */
    name?: pulumi.Input<string>;
    /**
     * The organization resource identifier for the security profile.
     */
    organizationResourceId?: pulumi.Input<string>;
    /**
     * Specifies a list of permissions assigned to the security profile.
     */
    permissions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The identifier for the Security Profile.
     */
    securityProfileId?: pulumi.Input<string>;
    /**
     * Tags to apply to the Security Profile. If configured with a provider
     * [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a SecurityProfile resource.
 */
export interface SecurityProfileArgs {
    /**
     * Specifies the description of the Security Profile.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the identifier of the hosting Amazon Connect Instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Specifies the name of the Security Profile.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies a list of permissions assigned to the security profile.
     */
    permissions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Tags to apply to the Security Profile. If configured with a provider
     * [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
