// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a CloudFront Field-level Encryption Profile resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as fs from "fs";
 *
 * const example = new aws.cloudfront.PublicKey("example", {
 *     comment: "test public key",
 *     encodedKey: fs.readFileSync("public_key.pem"),
 * });
 * const test = new aws.cloudfront.FieldLevelEncryptionProfile("test", {
 *     comment: "test comment",
 *     encryptionEntities: {
 *         items: [{
 *             publicKeyId: example.id,
 *             providerId: "test provider",
 *             fieldPatterns: {
 *                 items: ["DateOfBirth"],
 *             },
 *         }],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_cloudfront_field_level_encryption_profile.profile
 *
 *  id = "K3D5EWEUDCCXON" } Using `pulumi import`, import Cloudfront Field Level Encryption Profile using the `id`. For exampleconsole % pulumi import aws_cloudfront_field_level_encryption_profile.profile K3D5EWEUDCCXON
 */
export class FieldLevelEncryptionProfile extends pulumi.CustomResource {
    /**
     * Get an existing FieldLevelEncryptionProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FieldLevelEncryptionProfileState, opts?: pulumi.CustomResourceOptions): FieldLevelEncryptionProfile {
        return new FieldLevelEncryptionProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloudfront/fieldLevelEncryptionProfile:FieldLevelEncryptionProfile';

    /**
     * Returns true if the given object is an instance of FieldLevelEncryptionProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FieldLevelEncryptionProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FieldLevelEncryptionProfile.__pulumiType;
    }

    /**
     * Internal value used by CloudFront to allow future updates to the Field Level Encryption Profile.
     */
    public /*out*/ readonly callerReference!: pulumi.Output<string>;
    /**
     * An optional comment about the Field Level Encryption Profile.
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * The encryption entities config block for field-level encryption profiles that contains an attribute `items` which includes the encryption key and field pattern specifications.
     */
    public readonly encryptionEntities!: pulumi.Output<outputs.cloudfront.FieldLevelEncryptionProfileEncryptionEntities>;
    /**
     * The current version of the Field Level Encryption Profile. For example: `E2QWRUHAPOMQZL`.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name of the Field Level Encryption Profile.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a FieldLevelEncryptionProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FieldLevelEncryptionProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FieldLevelEncryptionProfileArgs | FieldLevelEncryptionProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FieldLevelEncryptionProfileState | undefined;
            resourceInputs["callerReference"] = state ? state.callerReference : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["encryptionEntities"] = state ? state.encryptionEntities : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as FieldLevelEncryptionProfileArgs | undefined;
            if ((!args || args.encryptionEntities === undefined) && !opts.urn) {
                throw new Error("Missing required property 'encryptionEntities'");
            }
            resourceInputs["comment"] = args ? args.comment : undefined;
            resourceInputs["encryptionEntities"] = args ? args.encryptionEntities : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["callerReference"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FieldLevelEncryptionProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FieldLevelEncryptionProfile resources.
 */
export interface FieldLevelEncryptionProfileState {
    /**
     * Internal value used by CloudFront to allow future updates to the Field Level Encryption Profile.
     */
    callerReference?: pulumi.Input<string>;
    /**
     * An optional comment about the Field Level Encryption Profile.
     */
    comment?: pulumi.Input<string>;
    /**
     * The encryption entities config block for field-level encryption profiles that contains an attribute `items` which includes the encryption key and field pattern specifications.
     */
    encryptionEntities?: pulumi.Input<inputs.cloudfront.FieldLevelEncryptionProfileEncryptionEntities>;
    /**
     * The current version of the Field Level Encryption Profile. For example: `E2QWRUHAPOMQZL`.
     */
    etag?: pulumi.Input<string>;
    /**
     * The name of the Field Level Encryption Profile.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FieldLevelEncryptionProfile resource.
 */
export interface FieldLevelEncryptionProfileArgs {
    /**
     * An optional comment about the Field Level Encryption Profile.
     */
    comment?: pulumi.Input<string>;
    /**
     * The encryption entities config block for field-level encryption profiles that contains an attribute `items` which includes the encryption key and field pattern specifications.
     */
    encryptionEntities: pulumi.Input<inputs.cloudfront.FieldLevelEncryptionProfileEncryptionEntities>;
    /**
     * The name of the Field Level Encryption Profile.
     */
    name?: pulumi.Input<string>;
}
