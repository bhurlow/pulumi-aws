// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a CloudFront Field-level Encryption Config resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.cloudfront.FieldLevelEncryptionConfig("test", {
 *     comment: "test comment",
 *     contentTypeProfileConfig: {
 *         forwardWhenContentTypeIsUnknown: true,
 *         contentTypeProfiles: {
 *             items: [{
 *                 contentType: "application/x-www-form-urlencoded",
 *                 format: "URLEncoded",
 *             }],
 *         },
 *     },
 *     queryArgProfileConfig: {
 *         forwardWhenQueryArgProfileIsUnknown: true,
 *         queryArgProfiles: {
 *             items: [{
 *                 profileId: aws_cloudfront_field_level_encryption_profile.test.id,
 *                 queryArg: "Arg1",
 *             }],
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_cloudfront_field_level_encryption_config.config
 *
 *  id = "E74FTE3AEXAMPLE" } Using `pulumi import`, import Cloudfront Field Level Encryption Config using the `id`. For exampleconsole % pulumi import aws_cloudfront_field_level_encryption_config.config E74FTE3AEXAMPLE
 */
export class FieldLevelEncryptionConfig extends pulumi.CustomResource {
    /**
     * Get an existing FieldLevelEncryptionConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FieldLevelEncryptionConfigState, opts?: pulumi.CustomResourceOptions): FieldLevelEncryptionConfig {
        return new FieldLevelEncryptionConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloudfront/fieldLevelEncryptionConfig:FieldLevelEncryptionConfig';

    /**
     * Returns true if the given object is an instance of FieldLevelEncryptionConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FieldLevelEncryptionConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FieldLevelEncryptionConfig.__pulumiType;
    }

    /**
     * Internal value used by CloudFront to allow future updates to the Field Level Encryption Config.
     */
    public /*out*/ readonly callerReference!: pulumi.Output<string>;
    /**
     * An optional comment about the Field Level Encryption Config.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Content Type Profile Config specifies when to forward content if a content type isn't recognized and profiles to use as by default in a request if a query argument doesn't specify a profile to use.
     */
    public readonly contentTypeProfileConfig!: pulumi.Output<outputs.cloudfront.FieldLevelEncryptionConfigContentTypeProfileConfig>;
    /**
     * The current version of the Field Level Encryption Config. For example: `E2QWRUHAPOMQZL`.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Query Arg Profile Config that specifies when to forward content if a profile isn't found and the profile that can be provided as a query argument in a request.
     */
    public readonly queryArgProfileConfig!: pulumi.Output<outputs.cloudfront.FieldLevelEncryptionConfigQueryArgProfileConfig>;

    /**
     * Create a FieldLevelEncryptionConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FieldLevelEncryptionConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FieldLevelEncryptionConfigArgs | FieldLevelEncryptionConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FieldLevelEncryptionConfigState | undefined;
            resourceInputs["callerReference"] = state ? state.callerReference : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["contentTypeProfileConfig"] = state ? state.contentTypeProfileConfig : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["queryArgProfileConfig"] = state ? state.queryArgProfileConfig : undefined;
        } else {
            const args = argsOrState as FieldLevelEncryptionConfigArgs | undefined;
            if ((!args || args.contentTypeProfileConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'contentTypeProfileConfig'");
            }
            if ((!args || args.queryArgProfileConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'queryArgProfileConfig'");
            }
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["contentTypeProfileConfig"] = args ? args.contentTypeProfileConfig : undefined;
            resourceInputs["queryArgProfileConfig"] = args ? args.queryArgProfileConfig : undefined;
            resourceInputs["callerReference"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FieldLevelEncryptionConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FieldLevelEncryptionConfig resources.
 */
export interface FieldLevelEncryptionConfigState {
    /**
     * Internal value used by CloudFront to allow future updates to the Field Level Encryption Config.
     */
    callerReference?: pulumi.Input<string>;
    /**
     * An optional comment about the Field Level Encryption Config.
     */
    comment?: pulumi.Input<string>;
    /**
     * Content Type Profile Config specifies when to forward content if a content type isn't recognized and profiles to use as by default in a request if a query argument doesn't specify a profile to use.
     */
    contentTypeProfileConfig?: pulumi.Input<inputs.cloudfront.FieldLevelEncryptionConfigContentTypeProfileConfig>;
    /**
     * The current version of the Field Level Encryption Config. For example: `E2QWRUHAPOMQZL`.
     */
    etag?: pulumi.Input<string>;
    /**
     * Query Arg Profile Config that specifies when to forward content if a profile isn't found and the profile that can be provided as a query argument in a request.
     */
    queryArgProfileConfig?: pulumi.Input<inputs.cloudfront.FieldLevelEncryptionConfigQueryArgProfileConfig>;
}

/**
 * The set of arguments for constructing a FieldLevelEncryptionConfig resource.
 */
export interface FieldLevelEncryptionConfigArgs {
    /**
     * An optional comment about the Field Level Encryption Config.
     */
    comment?: pulumi.Input<string>;
    /**
     * Content Type Profile Config specifies when to forward content if a content type isn't recognized and profiles to use as by default in a request if a query argument doesn't specify a profile to use.
     */
    contentTypeProfileConfig: pulumi.Input<inputs.cloudfront.FieldLevelEncryptionConfigContentTypeProfileConfig>;
    /**
     * Query Arg Profile Config that specifies when to forward content if a profile isn't found and the profile that can be provided as a query argument in a request.
     */
    queryArgProfileConfig: pulumi.Input<inputs.cloudfront.FieldLevelEncryptionConfigQueryArgProfileConfig>;
}
