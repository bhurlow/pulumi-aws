// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an IAM Virtual MFA Device.
 *
 * > **Note:** All attributes will be stored in the raw state as plain-text.
 * **Note:** A virtual MFA device cannot be directly associated with an IAM User from the provider.
 *   To associate the virtual MFA device with a user and enable it, use the code returned in either `base32StringSeed` or `qrCodePng` to generate TOTP authentication codes.
 *   The authentication codes can then be used with the AWS CLI command [`aws iam enable-mfa-device`](https://docs.aws.amazon.com/cli/latest/reference/iam/enable-mfa-device.html) or the AWS API call [`EnableMFADevice`](https://docs.aws.amazon.com/IAM/latest/APIReference/API_EnableMFADevice.html).
 *
 * ## Example Usage
 *
 * **Using certs on file:**
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.iam.VirtualMfaDevice("example", {virtualMfaDeviceName: "example"});
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_iam_virtual_mfa_device.example
 *
 *  id = "arn:aws:iam::123456789012:mfa/example" } Using `pulumi import`, import IAM Virtual MFA Devices using the `arn`. For exampleconsole % pulumi import aws_iam_virtual_mfa_device.example arn:aws:iam::123456789012:mfa/example
 */
export class VirtualMfaDevice extends pulumi.CustomResource {
    /**
     * Get an existing VirtualMfaDevice resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualMfaDeviceState, opts?: pulumi.CustomResourceOptions): VirtualMfaDevice {
        return new VirtualMfaDevice(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:iam/virtualMfaDevice:VirtualMfaDevice';

    /**
     * Returns true if the given object is an instance of VirtualMfaDevice.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualMfaDevice {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualMfaDevice.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) specifying the virtual mfa device.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The base32 seed defined as specified in [RFC3548](https://tools.ietf.org/html/rfc3548.txt). The `base32StringSeed` is base64-encoded.
     */
    public /*out*/ readonly base32StringSeed!: pulumi.Output<string>;
    /**
     * The date and time when the virtual MFA device was enabled.
     */
    public /*out*/ readonly enableDate!: pulumi.Output<string>;
    /**
     * The path for the virtual MFA device.
     */
    public readonly path!: pulumi.Output<string | undefined>;
    /**
     * A QR code PNG image that encodes `otpauth://totp/$virtualMFADeviceName@$AccountName?secret=$Base32String` where `$virtualMFADeviceName` is one of the create call arguments. AccountName is the user name if set (otherwise, the account ID), and Base32String is the seed in base32 format.
     */
    public /*out*/ readonly qrCodePng!: pulumi.Output<string>;
    /**
     * Map of resource tags for the virtual mfa device. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The associated IAM User name if the virtual MFA device is enabled.
     */
    public /*out*/ readonly userName!: pulumi.Output<string>;
    /**
     * The name of the virtual MFA device. Use with path to uniquely identify a virtual MFA device.
     */
    public readonly virtualMfaDeviceName!: pulumi.Output<string>;

    /**
     * Create a VirtualMfaDevice resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualMfaDeviceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualMfaDeviceArgs | VirtualMfaDeviceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VirtualMfaDeviceState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["base32StringSeed"] = state ? state.base32StringSeed : undefined;
            resourceInputs["enableDate"] = state ? state.enableDate : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["qrCodePng"] = state ? state.qrCodePng : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
            resourceInputs["virtualMfaDeviceName"] = state ? state.virtualMfaDeviceName : undefined;
        } else {
            const args = argsOrState as VirtualMfaDeviceArgs | undefined;
            if ((!args || args.virtualMfaDeviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'virtualMfaDeviceName'");
            }
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["virtualMfaDeviceName"] = args ? args.virtualMfaDeviceName : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["base32StringSeed"] = undefined /*out*/;
            resourceInputs["enableDate"] = undefined /*out*/;
            resourceInputs["qrCodePng"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["userName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VirtualMfaDevice.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualMfaDevice resources.
 */
export interface VirtualMfaDeviceState {
    /**
     * The Amazon Resource Name (ARN) specifying the virtual mfa device.
     */
    arn?: pulumi.Input<string>;
    /**
     * The base32 seed defined as specified in [RFC3548](https://tools.ietf.org/html/rfc3548.txt). The `base32StringSeed` is base64-encoded.
     */
    base32StringSeed?: pulumi.Input<string>;
    /**
     * The date and time when the virtual MFA device was enabled.
     */
    enableDate?: pulumi.Input<string>;
    /**
     * The path for the virtual MFA device.
     */
    path?: pulumi.Input<string>;
    /**
     * A QR code PNG image that encodes `otpauth://totp/$virtualMFADeviceName@$AccountName?secret=$Base32String` where `$virtualMFADeviceName` is one of the create call arguments. AccountName is the user name if set (otherwise, the account ID), and Base32String is the seed in base32 format.
     */
    qrCodePng?: pulumi.Input<string>;
    /**
     * Map of resource tags for the virtual mfa device. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The associated IAM User name if the virtual MFA device is enabled.
     */
    userName?: pulumi.Input<string>;
    /**
     * The name of the virtual MFA device. Use with path to uniquely identify a virtual MFA device.
     */
    virtualMfaDeviceName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VirtualMfaDevice resource.
 */
export interface VirtualMfaDeviceArgs {
    /**
     * The path for the virtual MFA device.
     */
    path?: pulumi.Input<string>;
    /**
     * Map of resource tags for the virtual mfa device. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the virtual MFA device. Use with path to uniquely identify a virtual MFA device.
     */
    virtualMfaDeviceName: pulumi.Input<string>;
}
