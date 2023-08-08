// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Manages a SMB Location within AWS DataSync.
 *
 * > **NOTE:** The DataSync Agents must be available before creating this resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.datasync.LocationSmb("example", {
 *     serverHostname: "smb.example.com",
 *     subdirectory: "/exported/path",
 *     user: "Guest",
 *     password: "ANotGreatPassword",
 *     agentArns: [aws_datasync_agent.example.arn],
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_datasync_location_smb.example
 *
 *  id = "arn:aws:datasync:us-east-1:123456789012:location/loc-12345678901234567" } Using `pulumi import`, import `aws_datasync_location_smb` using the Amazon Resource Name (ARN). For exampleconsole % pulumi import aws_datasync_location_smb.example arn:aws:datasync:us-east-1:123456789012:location/loc-12345678901234567
 */
export class LocationSmb extends pulumi.CustomResource {
    /**
     * Get an existing LocationSmb resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LocationSmbState, opts?: pulumi.CustomResourceOptions): LocationSmb {
        return new LocationSmb(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:datasync/locationSmb:LocationSmb';

    /**
     * Returns true if the given object is an instance of LocationSmb.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LocationSmb {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LocationSmb.__pulumiType;
    }

    /**
     * A list of DataSync Agent ARNs with which this location will be associated.
     */
    public readonly agentArns!: pulumi.Output<string[]>;
    /**
     * Amazon Resource Name (ARN) of the DataSync Location.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The name of the Windows domain the SMB server belongs to.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
     */
    public readonly mountOptions!: pulumi.Output<outputs.datasync.LocationSmbMountOptions | undefined>;
    /**
     * The password of the user who can mount the share and has file permissions in the SMB.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
     */
    public readonly serverHostname!: pulumi.Output<string>;
    /**
     * Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
     */
    public readonly subdirectory!: pulumi.Output<string>;
    /**
     * Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    public /*out*/ readonly uri!: pulumi.Output<string>;
    /**
     * The user who can mount the share and has file and folder permissions in the SMB share.
     */
    public readonly user!: pulumi.Output<string>;

    /**
     * Create a LocationSmb resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LocationSmbArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LocationSmbArgs | LocationSmbState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LocationSmbState | undefined;
            resourceInputs["agentArns"] = state ? state.agentArns : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["mountOptions"] = state ? state.mountOptions : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["serverHostname"] = state ? state.serverHostname : undefined;
            resourceInputs["subdirectory"] = state ? state.subdirectory : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["uri"] = state ? state.uri : undefined;
            resourceInputs["user"] = state ? state.user : undefined;
        } else {
            const args = argsOrState as LocationSmbArgs | undefined;
            if ((!args || args.agentArns === undefined) && !opts.urn) {
                throw new Error("Missing required property 'agentArns'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.serverHostname === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverHostname'");
            }
            if ((!args || args.subdirectory === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subdirectory'");
            }
            if ((!args || args.user === undefined) && !opts.urn) {
                throw new Error("Missing required property 'user'");
            }
            resourceInputs["agentArns"] = args ? args.agentArns : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["mountOptions"] = args ? args.mountOptions : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["serverHostname"] = args ? args.serverHostname : undefined;
            resourceInputs["subdirectory"] = args ? args.subdirectory : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["user"] = args ? args.user : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["uri"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(LocationSmb.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LocationSmb resources.
 */
export interface LocationSmbState {
    /**
     * A list of DataSync Agent ARNs with which this location will be associated.
     */
    agentArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Amazon Resource Name (ARN) of the DataSync Location.
     */
    arn?: pulumi.Input<string>;
    /**
     * The name of the Windows domain the SMB server belongs to.
     */
    domain?: pulumi.Input<string>;
    /**
     * Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
     */
    mountOptions?: pulumi.Input<inputs.datasync.LocationSmbMountOptions>;
    /**
     * The password of the user who can mount the share and has file permissions in the SMB.
     */
    password?: pulumi.Input<string>;
    /**
     * Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
     */
    serverHostname?: pulumi.Input<string>;
    /**
     * Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
     */
    subdirectory?: pulumi.Input<string>;
    /**
     * Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    uri?: pulumi.Input<string>;
    /**
     * The user who can mount the share and has file and folder permissions in the SMB share.
     */
    user?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LocationSmb resource.
 */
export interface LocationSmbArgs {
    /**
     * A list of DataSync Agent ARNs with which this location will be associated.
     */
    agentArns: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the Windows domain the SMB server belongs to.
     */
    domain?: pulumi.Input<string>;
    /**
     * Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
     */
    mountOptions?: pulumi.Input<inputs.datasync.LocationSmbMountOptions>;
    /**
     * The password of the user who can mount the share and has file permissions in the SMB.
     */
    password: pulumi.Input<string>;
    /**
     * Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
     */
    serverHostname: pulumi.Input<string>;
    /**
     * Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
     */
    subdirectory: pulumi.Input<string>;
    /**
     * Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The user who can mount the share and has file and folder permissions in the SMB share.
     */
    user: pulumi.Input<string>;
}
