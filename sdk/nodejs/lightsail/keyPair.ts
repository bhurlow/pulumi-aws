// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Lightsail Key Pair, for use with Lightsail Instances. These key pairs
 * are separate from EC2 Key Pairs, and must be created or imported for use with
 * Lightsail.
 *
 * > **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details
 *
 * ## Example Usage
 * ### Create New Key Pair
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * // Create a new Lightsail Key Pair
 * const lgKeyPair = new aws.lightsail.KeyPair("lgKeyPair", {});
 * ```
 * ### Create New Key Pair with PGP Encrypted Private Key
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const lgKeyPair = new aws.lightsail.KeyPair("lgKeyPair", {pgpKey: "keybase:keybaseusername"});
 * ```
 * ### Existing Public Key Import
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as fs from "fs";
 *
 * const lgKeyPair = new aws.lightsail.KeyPair("lgKeyPair", {publicKey: fs.readFileSync("~/.ssh/id_rsa.pub", "utf8")});
 * ```
 *
 * ## Import
 *
 * You cannot import Lightsail Key Pairs because the private and public key are only available on initial creation.
 */
export class KeyPair extends pulumi.CustomResource {
    /**
     * Get an existing KeyPair resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KeyPairState, opts?: pulumi.CustomResourceOptions): KeyPair {
        return new KeyPair(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:lightsail/keyPair:KeyPair';

    /**
     * Returns true if the given object is an instance of KeyPair.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KeyPair {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KeyPair.__pulumiType;
    }

    /**
     * The ARN of the Lightsail key pair.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The MD5 public key fingerprint for the encrypted private key.
     */
    public /*out*/ readonly encryptedFingerprint!: pulumi.Output<string>;
    /**
     * the private key material, base 64 encoded and encrypted with the given `pgpKey`. This is only populated when creating a new key and `pgpKey` is supplied.
     */
    public /*out*/ readonly encryptedPrivateKey!: pulumi.Output<string>;
    /**
     * The MD5 public key fingerprint as specified in section 4 of RFC 4716.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * The name of the Lightsail Key Pair. If omitted, a unique name will be generated by this provider
     */
    public readonly name!: pulumi.Output<string>;
    public readonly namePrefix!: pulumi.Output<string>;
    /**
     * An optional PGP key to encrypt the resulting private key material. Only used when creating a new key pair
     */
    public readonly pgpKey!: pulumi.Output<string | undefined>;
    /**
     * the private key, base64 encoded. This is only populated when creating a new key, and when no `pgpKey` is provided.
     */
    public /*out*/ readonly privateKey!: pulumi.Output<string>;
    /**
     * The public key material. This public key will be imported into Lightsail
     */
    public readonly publicKey!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the collection. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     *
     * > **NOTE:** a PGP key is not required, however it is strongly encouraged. Without a PGP key, the private key material will be stored in state unencrypted.`pgpKey` is ignored if `publicKey` is supplied.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * @deprecated Please use `tags` instead.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a KeyPair resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: KeyPairArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeyPairArgs | KeyPairState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KeyPairState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["encryptedFingerprint"] = state ? state.encryptedFingerprint : undefined;
            resourceInputs["encryptedPrivateKey"] = state ? state.encryptedPrivateKey : undefined;
            resourceInputs["fingerprint"] = state ? state.fingerprint : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namePrefix"] = state ? state.namePrefix : undefined;
            resourceInputs["pgpKey"] = state ? state.pgpKey : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
            resourceInputs["publicKey"] = state ? state.publicKey : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as KeyPairArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namePrefix"] = args ? args.namePrefix : undefined;
            resourceInputs["pgpKey"] = args ? args.pgpKey : undefined;
            resourceInputs["publicKey"] = args ? args.publicKey : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["encryptedFingerprint"] = undefined /*out*/;
            resourceInputs["encryptedPrivateKey"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["privateKey"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(KeyPair.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KeyPair resources.
 */
export interface KeyPairState {
    /**
     * The ARN of the Lightsail key pair.
     */
    arn?: pulumi.Input<string>;
    /**
     * The MD5 public key fingerprint for the encrypted private key.
     */
    encryptedFingerprint?: pulumi.Input<string>;
    /**
     * the private key material, base 64 encoded and encrypted with the given `pgpKey`. This is only populated when creating a new key and `pgpKey` is supplied.
     */
    encryptedPrivateKey?: pulumi.Input<string>;
    /**
     * The MD5 public key fingerprint as specified in section 4 of RFC 4716.
     */
    fingerprint?: pulumi.Input<string>;
    /**
     * The name of the Lightsail Key Pair. If omitted, a unique name will be generated by this provider
     */
    name?: pulumi.Input<string>;
    namePrefix?: pulumi.Input<string>;
    /**
     * An optional PGP key to encrypt the resulting private key material. Only used when creating a new key pair
     */
    pgpKey?: pulumi.Input<string>;
    /**
     * the private key, base64 encoded. This is only populated when creating a new key, and when no `pgpKey` is provided.
     */
    privateKey?: pulumi.Input<string>;
    /**
     * The public key material. This public key will be imported into Lightsail
     */
    publicKey?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the collection. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     *
     * > **NOTE:** a PGP key is not required, however it is strongly encouraged. Without a PGP key, the private key material will be stored in state unencrypted.`pgpKey` is ignored if `publicKey` is supplied.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * @deprecated Please use `tags` instead.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a KeyPair resource.
 */
export interface KeyPairArgs {
    /**
     * The name of the Lightsail Key Pair. If omitted, a unique name will be generated by this provider
     */
    name?: pulumi.Input<string>;
    namePrefix?: pulumi.Input<string>;
    /**
     * An optional PGP key to encrypt the resulting private key material. Only used when creating a new key pair
     */
    pgpKey?: pulumi.Input<string>;
    /**
     * The public key material. This public key will be imported into Lightsail
     */
    publicKey?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the collection. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     *
     * > **NOTE:** a PGP key is not required, however it is strongly encouraged. Without a PGP key, the private key material will be stored in state unencrypted.`pgpKey` is ignored if `publicKey` is supplied.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
