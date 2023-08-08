// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource for managing an AWS FinSpace Kx User.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleKey = new aws.kms.Key("exampleKey", {
 *     description: "Example KMS Key",
 *     deletionWindowInDays: 7,
 * });
 * const exampleKxEnvironment = new aws.finspace.KxEnvironment("exampleKxEnvironment", {kmsKeyId: exampleKey.arn});
 * const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: JSON.stringify({
 *     Version: "2012-10-17",
 *     Statement: [{
 *         Action: "sts:AssumeRole",
 *         Effect: "Allow",
 *         Sid: "",
 *         Principal: {
 *             Service: "ec2.amazonaws.com",
 *         },
 *     }],
 * })});
 * const exampleKxUser = new aws.finspace.KxUser("exampleKxUser", {
 *     environmentId: exampleKxEnvironment.id,
 *     iamRole: exampleRole.arn,
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_finspace_kx_user.example
 *
 *  id = "n3ceo7wqxoxcti5tujqwzs,my-tf-kx-user" } Using `pulumi import`, import an AWS FinSpace Kx User using the `id` (environment ID and user name, comma-delimited). For exampleconsole % pulumi import aws_finspace_kx_user.example n3ceo7wqxoxcti5tujqwzs,my-tf-kx-user
 */
export class KxUser extends pulumi.CustomResource {
    /**
     * Get an existing KxUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KxUserState, opts?: pulumi.CustomResourceOptions): KxUser {
        return new KxUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:finspace/kxUser:KxUser';

    /**
     * Returns true if the given object is an instance of KxUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KxUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KxUser.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) identifier of the KX user.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Unique identifier for the KX environment.
     */
    public readonly environmentId!: pulumi.Output<string>;
    /**
     * IAM role ARN to be associated with the user.
     *
     * The following arguments are optional:
     */
    public readonly iamRole!: pulumi.Output<string>;
    /**
     * A unique identifier for the user.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a KxUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KxUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KxUserArgs | KxUserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KxUserState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["environmentId"] = state ? state.environmentId : undefined;
            resourceInputs["iamRole"] = state ? state.iamRole : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as KxUserArgs | undefined;
            if ((!args || args.environmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environmentId'");
            }
            if ((!args || args.iamRole === undefined) && !opts.urn) {
                throw new Error("Missing required property 'iamRole'");
            }
            resourceInputs["environmentId"] = args ? args.environmentId : undefined;
            resourceInputs["iamRole"] = args ? args.iamRole : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(KxUser.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KxUser resources.
 */
export interface KxUserState {
    /**
     * Amazon Resource Name (ARN) identifier of the KX user.
     */
    arn?: pulumi.Input<string>;
    /**
     * Unique identifier for the KX environment.
     */
    environmentId?: pulumi.Input<string>;
    /**
     * IAM role ARN to be associated with the user.
     *
     * The following arguments are optional:
     */
    iamRole?: pulumi.Input<string>;
    /**
     * A unique identifier for the user.
     */
    name?: pulumi.Input<string>;
    /**
     * Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a KxUser resource.
 */
export interface KxUserArgs {
    /**
     * Unique identifier for the KX environment.
     */
    environmentId: pulumi.Input<string>;
    /**
     * IAM role ARN to be associated with the user.
     *
     * The following arguments are optional:
     */
    iamRole: pulumi.Input<string>;
    /**
     * A unique identifier for the user.
     */
    name?: pulumi.Input<string>;
    /**
     * Key-value mapping of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
