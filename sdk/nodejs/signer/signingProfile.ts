// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Creates a Signer Signing Profile. A signing profile contains information about the code signing configuration parameters that can be used by a given code signing user.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testSp = new aws.signer.SigningProfile("test_sp", {
 *     platformId: "AWSLambda-SHA384-ECDSA",
 * });
 * const prodSp = new aws.signer.SigningProfile("prod_sp", {
 *     namePrefix: "prod_sp_",
 *     platformId: "AWSLambda-SHA384-ECDSA",
 *     signatureValidityPeriod: {
 *         type: "YEARS",
 *         value: 5,
 *     },
 *     tags: {
 *         tag1: "value1",
 *         tag2: "value2",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Signer signing profiles can be imported using the `name`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:signer/signingProfile:SigningProfile test_signer_signing_profile test_sp_DdW3Mk1foYL88fajut4mTVFGpuwfd4ACO6ANL0D1uIj7lrn8adK
 * ```
 */
export class SigningProfile extends pulumi.CustomResource {
    /**
     * Get an existing SigningProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SigningProfileState, opts?: pulumi.CustomResourceOptions): SigningProfile {
        return new SigningProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:signer/signingProfile:SigningProfile';

    /**
     * Returns true if the given object is an instance of SigningProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SigningProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SigningProfile.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) for the signing profile.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A unique signing profile name. By default generated by the provider. Signing profile names are immutable and cannot be reused after canceled.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A signing profile name prefix. The provider will generate a unique suffix. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    /**
     * A human-readable name for the signing platform associated with the signing profile.
     */
    public /*out*/ readonly platformDisplayName!: pulumi.Output<string>;
    /**
     * The ID of the platform that is used by the target signing profile.
     */
    public readonly platformId!: pulumi.Output<string>;
    /**
     * Revocation information for a signing profile.
     */
    public /*out*/ readonly revocationRecords!: pulumi.Output<outputs.signer.SigningProfileRevocationRecord[]>;
    /**
     * The validity period for a signing job.
     */
    public readonly signatureValidityPeriod!: pulumi.Output<outputs.signer.SigningProfileSignatureValidityPeriod>;
    /**
     * The status of the target signing profile.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A list of tags associated with the signing profile.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The current version of the signing profile.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;
    /**
     * The signing profile ARN, including the profile version.
     */
    public /*out*/ readonly versionArn!: pulumi.Output<string>;

    /**
     * Create a SigningProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SigningProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SigningProfileArgs | SigningProfileState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SigningProfileState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["platformDisplayName"] = state ? state.platformDisplayName : undefined;
            inputs["platformId"] = state ? state.platformId : undefined;
            inputs["revocationRecords"] = state ? state.revocationRecords : undefined;
            inputs["signatureValidityPeriod"] = state ? state.signatureValidityPeriod : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["version"] = state ? state.version : undefined;
            inputs["versionArn"] = state ? state.versionArn : undefined;
        } else {
            const args = argsOrState as SigningProfileArgs | undefined;
            if ((!args || args.platformId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'platformId'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["platformId"] = args ? args.platformId : undefined;
            inputs["signatureValidityPeriod"] = args ? args.signatureValidityPeriod : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["platformDisplayName"] = undefined /*out*/;
            inputs["revocationRecords"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
            inputs["versionArn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SigningProfile.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SigningProfile resources.
 */
export interface SigningProfileState {
    /**
     * The Amazon Resource Name (ARN) for the signing profile.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * A unique signing profile name. By default generated by the provider. Signing profile names are immutable and cannot be reused after canceled.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A signing profile name prefix. The provider will generate a unique suffix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * A human-readable name for the signing platform associated with the signing profile.
     */
    readonly platformDisplayName?: pulumi.Input<string>;
    /**
     * The ID of the platform that is used by the target signing profile.
     */
    readonly platformId?: pulumi.Input<string>;
    /**
     * Revocation information for a signing profile.
     */
    readonly revocationRecords?: pulumi.Input<pulumi.Input<inputs.signer.SigningProfileRevocationRecord>[]>;
    /**
     * The validity period for a signing job.
     */
    readonly signatureValidityPeriod?: pulumi.Input<inputs.signer.SigningProfileSignatureValidityPeriod>;
    /**
     * The status of the target signing profile.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * A list of tags associated with the signing profile.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The current version of the signing profile.
     */
    readonly version?: pulumi.Input<string>;
    /**
     * The signing profile ARN, including the profile version.
     */
    readonly versionArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SigningProfile resource.
 */
export interface SigningProfileArgs {
    /**
     * A unique signing profile name. By default generated by the provider. Signing profile names are immutable and cannot be reused after canceled.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A signing profile name prefix. The provider will generate a unique suffix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * The ID of the platform that is used by the target signing profile.
     */
    readonly platformId: pulumi.Input<string>;
    /**
     * The validity period for a signing job.
     */
    readonly signatureValidityPeriod?: pulumi.Input<inputs.signer.SigningProfileSignatureValidityPeriod>;
    /**
     * A list of tags associated with the signing profile.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
