// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a new Amazon Redshift Serverless Namespace.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.redshiftserverless.Namespace("example", {namespaceName: "concurrency-scaling"});
 * ```
 *
 * ## Import
 *
 * Using `pulumi import`, import Redshift Serverless Namespaces using the `namespace_name`. For example:
 *
 * ```sh
 *  $ pulumi import aws:redshiftserverless/namespace:Namespace example example
 * ```
 */
export class Namespace extends pulumi.CustomResource {
    /**
     * Get an existing Namespace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NamespaceState, opts?: pulumi.CustomResourceOptions): Namespace {
        return new Namespace(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:redshiftserverless/namespace:Namespace';

    /**
     * Returns true if the given object is an instance of Namespace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Namespace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Namespace.__pulumiType;
    }

    /**
     * The password of the administrator for the first database created in the namespace.
     */
    public readonly adminUserPassword!: pulumi.Output<string | undefined>;
    /**
     * The username of the administrator for the first database created in the namespace.
     */
    public readonly adminUsername!: pulumi.Output<string>;
    /**
     * Amazon Resource Name (ARN) of the Redshift Serverless Namespace.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The name of the first database created in the namespace.
     */
    public readonly dbName!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role to set as a default in the namespace. When specifying `defaultIamRoleArn`, it also must be part of `iamRoles`.
     */
    public readonly defaultIamRoleArn!: pulumi.Output<string | undefined>;
    /**
     * A list of IAM roles to associate with the namespace.
     */
    public readonly iamRoles!: pulumi.Output<string[]>;
    /**
     * The ARN of the Amazon Web Services Key Management Service key used to encrypt your data.
     */
    public readonly kmsKeyId!: pulumi.Output<string>;
    /**
     * The types of logs the namespace can export. Available export types are `userlog`, `connectionlog`, and `useractivitylog`.
     */
    public readonly logExports!: pulumi.Output<string[] | undefined>;
    /**
     * The Redshift Namespace ID.
     */
    public /*out*/ readonly namespaceId!: pulumi.Output<string>;
    /**
     * The name of the namespace.
     */
    public readonly namespaceName!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a Namespace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NamespaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NamespaceArgs | NamespaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NamespaceState | undefined;
            resourceInputs["adminUserPassword"] = state ? state.adminUserPassword : undefined;
            resourceInputs["adminUsername"] = state ? state.adminUsername : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["dbName"] = state ? state.dbName : undefined;
            resourceInputs["defaultIamRoleArn"] = state ? state.defaultIamRoleArn : undefined;
            resourceInputs["iamRoles"] = state ? state.iamRoles : undefined;
            resourceInputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            resourceInputs["logExports"] = state ? state.logExports : undefined;
            resourceInputs["namespaceId"] = state ? state.namespaceId : undefined;
            resourceInputs["namespaceName"] = state ? state.namespaceName : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as NamespaceArgs | undefined;
            if ((!args || args.namespaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespaceName'");
            }
            resourceInputs["adminUserPassword"] = args?.adminUserPassword ? pulumi.secret(args.adminUserPassword) : undefined;
            resourceInputs["adminUsername"] = args?.adminUsername ? pulumi.secret(args.adminUsername) : undefined;
            resourceInputs["dbName"] = args ? args.dbName : undefined;
            resourceInputs["defaultIamRoleArn"] = args ? args.defaultIamRoleArn : undefined;
            resourceInputs["iamRoles"] = args ? args.iamRoles : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["logExports"] = args ? args.logExports : undefined;
            resourceInputs["namespaceName"] = args ? args.namespaceName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["namespaceId"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["adminUserPassword", "adminUsername", "tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Namespace.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Namespace resources.
 */
export interface NamespaceState {
    /**
     * The password of the administrator for the first database created in the namespace.
     */
    adminUserPassword?: pulumi.Input<string>;
    /**
     * The username of the administrator for the first database created in the namespace.
     */
    adminUsername?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the Redshift Serverless Namespace.
     */
    arn?: pulumi.Input<string>;
    /**
     * The name of the first database created in the namespace.
     */
    dbName?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role to set as a default in the namespace. When specifying `defaultIamRoleArn`, it also must be part of `iamRoles`.
     */
    defaultIamRoleArn?: pulumi.Input<string>;
    /**
     * A list of IAM roles to associate with the namespace.
     */
    iamRoles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ARN of the Amazon Web Services Key Management Service key used to encrypt your data.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * The types of logs the namespace can export. Available export types are `userlog`, `connectionlog`, and `useractivitylog`.
     */
    logExports?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Redshift Namespace ID.
     */
    namespaceId?: pulumi.Input<string>;
    /**
     * The name of the namespace.
     */
    namespaceName?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Namespace resource.
 */
export interface NamespaceArgs {
    /**
     * The password of the administrator for the first database created in the namespace.
     */
    adminUserPassword?: pulumi.Input<string>;
    /**
     * The username of the administrator for the first database created in the namespace.
     */
    adminUsername?: pulumi.Input<string>;
    /**
     * The name of the first database created in the namespace.
     */
    dbName?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role to set as a default in the namespace. When specifying `defaultIamRoleArn`, it also must be part of `iamRoles`.
     */
    defaultIamRoleArn?: pulumi.Input<string>;
    /**
     * A list of IAM roles to associate with the namespace.
     */
    iamRoles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ARN of the Amazon Web Services Key Management Service key used to encrypt your data.
     */
    kmsKeyId?: pulumi.Input<string>;
    /**
     * The types of logs the namespace can export. Available export types are `userlog`, `connectionlog`, and `useractivitylog`.
     */
    logExports?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the namespace.
     */
    namespaceName: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
