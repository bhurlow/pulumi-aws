// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage AWS Device Farm Instance Profiles.
 * ∂
 * > **NOTE:** AWS currently has limited regional support for Device Farm (e.g., `us-west-2`). See [AWS Device Farm endpoints and quotas](https://docs.aws.amazon.com/general/latest/gr/devicefarm.html) for information on supported regions.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.devicefarm.InstanceProfile("example", {});
 * ```
 *
 * ## Import
 *
 * DeviceFarm Instance Profiles can be imported by their arn
 *
 * ```sh
 *  $ pulumi import aws:devicefarm/instanceProfile:InstanceProfile example arn:aws:devicefarm:us-west-2:123456789012:instanceprofile:4fa784c7-ccb4-4dbf-ba4f-02198320daa1
 * ```
 */
export class InstanceProfile extends pulumi.CustomResource {
    /**
     * Get an existing InstanceProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceProfileState, opts?: pulumi.CustomResourceOptions): InstanceProfile {
        return new InstanceProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:devicefarm/instanceProfile:InstanceProfile';

    /**
     * Returns true if the given object is an instance of InstanceProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceProfile.__pulumiType;
    }

    /**
     * The Amazon Resource Name of this instance profile.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The description of the instance profile.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * An array of strings that specifies the list of app packages that should not be cleaned up from the device after a test run.
     */
    public readonly excludeAppPackagesFromCleanups!: pulumi.Output<string[] | undefined>;
    /**
     * The name for the instance profile.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * When set to `true`, Device Farm removes app packages after a test run. The default value is `false` for private devices.
     */
    public readonly packageCleanup!: pulumi.Output<boolean | undefined>;
    /**
     * When set to `true`, Device Farm reboots the instance after a test run. The default value is `true`.
     */
    public readonly rebootAfterUse!: pulumi.Output<boolean | undefined>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a InstanceProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InstanceProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceProfileArgs | InstanceProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceProfileState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["excludeAppPackagesFromCleanups"] = state ? state.excludeAppPackagesFromCleanups : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["packageCleanup"] = state ? state.packageCleanup : undefined;
            resourceInputs["rebootAfterUse"] = state ? state.rebootAfterUse : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as InstanceProfileArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["excludeAppPackagesFromCleanups"] = args ? args.excludeAppPackagesFromCleanups : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["packageCleanup"] = args ? args.packageCleanup : undefined;
            resourceInputs["rebootAfterUse"] = args ? args.rebootAfterUse : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstanceProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceProfile resources.
 */
export interface InstanceProfileState {
    /**
     * The Amazon Resource Name of this instance profile.
     */
    arn?: pulumi.Input<string>;
    /**
     * The description of the instance profile.
     */
    description?: pulumi.Input<string>;
    /**
     * An array of strings that specifies the list of app packages that should not be cleaned up from the device after a test run.
     */
    excludeAppPackagesFromCleanups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name for the instance profile.
     */
    name?: pulumi.Input<string>;
    /**
     * When set to `true`, Device Farm removes app packages after a test run. The default value is `false` for private devices.
     */
    packageCleanup?: pulumi.Input<boolean>;
    /**
     * When set to `true`, Device Farm reboots the instance after a test run. The default value is `true`.
     */
    rebootAfterUse?: pulumi.Input<boolean>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a InstanceProfile resource.
 */
export interface InstanceProfileArgs {
    /**
     * The description of the instance profile.
     */
    description?: pulumi.Input<string>;
    /**
     * An array of strings that specifies the list of app packages that should not be cleaned up from the device after a test run.
     */
    excludeAppPackagesFromCleanups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name for the instance profile.
     */
    name?: pulumi.Input<string>;
    /**
     * When set to `true`, Device Farm removes app packages after a test run. The default value is `false` for private devices.
     */
    packageCleanup?: pulumi.Input<boolean>;
    /**
     * When set to `true`, Device Farm reboots the instance after a test run. The default value is `true`.
     */
    rebootAfterUse?: pulumi.Input<boolean>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
