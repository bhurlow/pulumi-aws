// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an AWS Quantum Ledger Database (QLDB) resource
 *
 * > **NOTE:** Deletion protection is enabled by default. To successfully delete this resource via this provider, `deletionProtection = false` must be applied before attempting deletion.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const sample_ledger = new aws.qldb.Ledger("sample-ledger", {permissionsMode: "STANDARD"});
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_qldb_ledger.sample-ledger
 *
 *  id = "sample-ledger" } Using `pulumi import`, import QLDB Ledgers using the `name`. For exampleconsole % pulumi import aws_qldb_ledger.sample-ledger sample-ledger
 */
export class Ledger extends pulumi.CustomResource {
    /**
     * Get an existing Ledger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LedgerState, opts?: pulumi.CustomResourceOptions): Ledger {
        return new Ledger(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:qldb/ledger:Ledger';

    /**
     * Returns true if the given object is an instance of Ledger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ledger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ledger.__pulumiType;
    }

    /**
     * The ARN of the QLDB Ledger
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The deletion protection for the QLDB Ledger instance. By default it is `true`. To delete this resource via the provider, this value must be configured to `false` and applied first before attempting deletion.
     */
    public readonly deletionProtection!: pulumi.Output<boolean | undefined>;
    /**
     * The key in AWS Key Management Service (AWS KMS) to use for encryption of data at rest in the ledger. For more information, see the [AWS documentation](https://docs.aws.amazon.com/qldb/latest/developerguide/encryption-at-rest.html). Valid values are `"AWS_OWNED_KMS_KEY"` to use an AWS KMS key that is owned and managed by AWS on your behalf, or the ARN of a valid symmetric customer managed KMS key.
     */
    public readonly kmsKey!: pulumi.Output<string>;
    /**
     * The friendly name for the QLDB Ledger instance. By default generated by the provider.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The permissions mode for the QLDB ledger instance. Specify either `ALLOW_ALL` or `STANDARD`.
     */
    public readonly permissionsMode!: pulumi.Output<string>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a Ledger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LedgerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LedgerArgs | LedgerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LedgerState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            resourceInputs["kmsKey"] = state ? state.kmsKey : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["permissionsMode"] = state ? state.permissionsMode : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as LedgerArgs | undefined;
            if ((!args || args.permissionsMode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permissionsMode'");
            }
            resourceInputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            resourceInputs["kmsKey"] = args ? args.kmsKey : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["permissionsMode"] = args ? args.permissionsMode : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Ledger.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ledger resources.
 */
export interface LedgerState {
    /**
     * The ARN of the QLDB Ledger
     */
    arn?: pulumi.Input<string>;
    /**
     * The deletion protection for the QLDB Ledger instance. By default it is `true`. To delete this resource via the provider, this value must be configured to `false` and applied first before attempting deletion.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * The key in AWS Key Management Service (AWS KMS) to use for encryption of data at rest in the ledger. For more information, see the [AWS documentation](https://docs.aws.amazon.com/qldb/latest/developerguide/encryption-at-rest.html). Valid values are `"AWS_OWNED_KMS_KEY"` to use an AWS KMS key that is owned and managed by AWS on your behalf, or the ARN of a valid symmetric customer managed KMS key.
     */
    kmsKey?: pulumi.Input<string>;
    /**
     * The friendly name for the QLDB Ledger instance. By default generated by the provider.
     */
    name?: pulumi.Input<string>;
    /**
     * The permissions mode for the QLDB ledger instance. Specify either `ALLOW_ALL` or `STANDARD`.
     */
    permissionsMode?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Ledger resource.
 */
export interface LedgerArgs {
    /**
     * The deletion protection for the QLDB Ledger instance. By default it is `true`. To delete this resource via the provider, this value must be configured to `false` and applied first before attempting deletion.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * The key in AWS Key Management Service (AWS KMS) to use for encryption of data at rest in the ledger. For more information, see the [AWS documentation](https://docs.aws.amazon.com/qldb/latest/developerguide/encryption-at-rest.html). Valid values are `"AWS_OWNED_KMS_KEY"` to use an AWS KMS key that is owned and managed by AWS on your behalf, or the ARN of a valid symmetric customer managed KMS key.
     */
    kmsKey?: pulumi.Input<string>;
    /**
     * The friendly name for the QLDB Ledger instance. By default generated by the provider.
     */
    name?: pulumi.Input<string>;
    /**
     * The permissions mode for the QLDB ledger instance. Specify either `ALLOW_ALL` or `STANDARD`.
     */
    permissionsMode: pulumi.Input<string>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
