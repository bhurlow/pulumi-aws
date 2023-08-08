// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
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
 * const testSp = new aws.signer.SigningProfile("testSp", {platformId: "AWSLambda-SHA384-ECDSA"});
 * const prodSp = new aws.signer.SigningProfile("prodSp", {
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
 * terraform import {
 *
 *  to = aws_signer_signing_profile.test_signer_signing_profile
 *
 *  id = "test_sp_DdW3Mk1foYL88fajut4mTVFGpuwfd4ACO6ANL0D1uIj7lrn8adK" } Using `pulumi import`, import Signer signing profiles using the `name`. For exampleconsole % pulumi import aws_signer_signing_profile.test_signer_signing_profile test_sp_DdW3Mk1foYL88fajut4mTVFGpuwfd4ACO6ANL0D1uIj7lrn8adK
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
    public readonly signingMaterial!: pulumi.Output<outputs.signer.SigningProfileSigningMaterial>;
    /**
     * The status of the target signing profile.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A list of tags associated with the signing profile. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SigningProfileState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namePrefix"] = state ? state.namePrefix : undefined;
            resourceInputs["platformDisplayName"] = state ? state.platformDisplayName : undefined;
            resourceInputs["platformId"] = state ? state.platformId : undefined;
            resourceInputs["revocationRecords"] = state ? state.revocationRecords : undefined;
            resourceInputs["signatureValidityPeriod"] = state ? state.signatureValidityPeriod : undefined;
            resourceInputs["signingMaterial"] = state ? state.signingMaterial : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
            resourceInputs["versionArn"] = state ? state.versionArn : undefined;
        } else {
            const args = argsOrState as SigningProfileArgs | undefined;
            if ((!args || args.platformId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'platformId'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namePrefix"] = args ? args.namePrefix : undefined;
            resourceInputs["platformId"] = args ? args.platformId : undefined;
            resourceInputs["signatureValidityPeriod"] = args ? args.signatureValidityPeriod : undefined;
            resourceInputs["signingMaterial"] = args ? args.signingMaterial : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["platformDisplayName"] = undefined /*out*/;
            resourceInputs["revocationRecords"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
            resourceInputs["versionArn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SigningProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SigningProfile resources.
 */
export interface SigningProfileState {
    /**
     * The Amazon Resource Name (ARN) for the signing profile.
     */
    arn?: pulumi.Input<string>;
    /**
     * A unique signing profile name. By default generated by the provider. Signing profile names are immutable and cannot be reused after canceled.
     */
    name?: pulumi.Input<string>;
    /**
     * A signing profile name prefix. The provider will generate a unique suffix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * A human-readable name for the signing platform associated with the signing profile.
     */
    platformDisplayName?: pulumi.Input<string>;
    /**
     * The ID of the platform that is used by the target signing profile.
     */
    platformId?: pulumi.Input<string>;
    /**
     * Revocation information for a signing profile.
     */
    revocationRecords?: pulumi.Input<pulumi.Input<inputs.signer.SigningProfileRevocationRecord>[]>;
    /**
     * The validity period for a signing job.
     */
    signatureValidityPeriod?: pulumi.Input<inputs.signer.SigningProfileSignatureValidityPeriod>;
    signingMaterial?: pulumi.Input<inputs.signer.SigningProfileSigningMaterial>;
    /**
     * The status of the target signing profile.
     */
    status?: pulumi.Input<string>;
    /**
     * A list of tags associated with the signing profile. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The current version of the signing profile.
     */
    version?: pulumi.Input<string>;
    /**
     * The signing profile ARN, including the profile version.
     */
    versionArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SigningProfile resource.
 */
export interface SigningProfileArgs {
    /**
     * A unique signing profile name. By default generated by the provider. Signing profile names are immutable and cannot be reused after canceled.
     */
    name?: pulumi.Input<string>;
    /**
     * A signing profile name prefix. The provider will generate a unique suffix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The ID of the platform that is used by the target signing profile.
     */
    platformId: pulumi.Input<string>;
    /**
     * The validity period for a signing job.
     */
    signatureValidityPeriod?: pulumi.Input<inputs.signer.SigningProfileSignatureValidityPeriod>;
    signingMaterial?: pulumi.Input<inputs.signer.SigningProfileSigningMaterial>;
    /**
     * A list of tags associated with the signing profile. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
