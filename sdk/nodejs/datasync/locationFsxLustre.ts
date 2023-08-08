// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an AWS DataSync FSx Lustre Location.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.datasync.LocationFsxLustre("example", {
 *     fsxFilesystemArn: aws_fsx_lustre_file_system.example.arn,
 *     securityGroupArns: [aws_security_group.example.arn],
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_datasync_location_fsx_lustre_file_system.example
 *
 *  id = "arn:aws:datasync:us-west-2:123456789012:location/loc-12345678901234567#arn:aws:fsx:us-west-2:476956259333:file-system/fs-08e04cd442c1bb94a" } Using `pulumi import`, import `aws_datasync_location_fsx_lustre_file_system` using the `DataSync-ARN#FSx-Lustre-ARN`. For exampleconsole % pulumi import aws_datasync_location_fsx_lustre_file_system.example arn:aws:datasync:us-west-2:123456789012:location/loc-12345678901234567#arn:aws:fsx:us-west-2:476956259333:file-system/fs-08e04cd442c1bb94a
 */
export class LocationFsxLustre extends pulumi.CustomResource {
    /**
     * Get an existing LocationFsxLustre resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LocationFsxLustreState, opts?: pulumi.CustomResourceOptions): LocationFsxLustre {
        return new LocationFsxLustre(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:datasync/locationFsxLustre:LocationFsxLustre';

    /**
     * Returns true if the given object is an instance of LocationFsxLustre.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LocationFsxLustre {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LocationFsxLustre.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of the DataSync Location.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The time that the FSx for Lustre location was created.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) for the FSx for Lustre file system.
     */
    public readonly fsxFilesystemArn!: pulumi.Output<string>;
    /**
     * The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Lustre file system.
     */
    public readonly securityGroupArns!: pulumi.Output<string[]>;
    /**
     * Subdirectory to perform actions as source or destination.
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
    /**
     * The URL of the FSx for Lustre location that was described.
     */
    public /*out*/ readonly uri!: pulumi.Output<string>;

    /**
     * Create a LocationFsxLustre resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LocationFsxLustreArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LocationFsxLustreArgs | LocationFsxLustreState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LocationFsxLustreState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["creationTime"] = state ? state.creationTime : undefined;
            resourceInputs["fsxFilesystemArn"] = state ? state.fsxFilesystemArn : undefined;
            resourceInputs["securityGroupArns"] = state ? state.securityGroupArns : undefined;
            resourceInputs["subdirectory"] = state ? state.subdirectory : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["uri"] = state ? state.uri : undefined;
        } else {
            const args = argsOrState as LocationFsxLustreArgs | undefined;
            if ((!args || args.fsxFilesystemArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'fsxFilesystemArn'");
            }
            if ((!args || args.securityGroupArns === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupArns'");
            }
            resourceInputs["fsxFilesystemArn"] = args ? args.fsxFilesystemArn : undefined;
            resourceInputs["securityGroupArns"] = args ? args.securityGroupArns : undefined;
            resourceInputs["subdirectory"] = args ? args.subdirectory : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["uri"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LocationFsxLustre.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LocationFsxLustre resources.
 */
export interface LocationFsxLustreState {
    /**
     * Amazon Resource Name (ARN) of the DataSync Location.
     */
    arn?: pulumi.Input<string>;
    /**
     * The time that the FSx for Lustre location was created.
     */
    creationTime?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) for the FSx for Lustre file system.
     */
    fsxFilesystemArn?: pulumi.Input<string>;
    /**
     * The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Lustre file system.
     */
    securityGroupArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Subdirectory to perform actions as source or destination.
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
    /**
     * The URL of the FSx for Lustre location that was described.
     */
    uri?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LocationFsxLustre resource.
 */
export interface LocationFsxLustreArgs {
    /**
     * The Amazon Resource Name (ARN) for the FSx for Lustre file system.
     */
    fsxFilesystemArn: pulumi.Input<string>;
    /**
     * The Amazon Resource Names (ARNs) of the security groups that are to use to configure the FSx for Lustre file system.
     */
    securityGroupArns: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Subdirectory to perform actions as source or destination.
     */
    subdirectory?: pulumi.Input<string>;
    /**
     * Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
