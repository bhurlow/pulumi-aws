// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an App Runner AutoScaling Configuration Version.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.apprunner.AutoScalingConfigurationVersion("example", {
 *     autoScalingConfigurationName: "example",
 *     maxConcurrency: 50,
 *     maxSize: 10,
 *     minSize: 2,
 *     tags: {
 *         Name: "example-apprunner-autoscaling",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Using `pulumi import`, import App Runner AutoScaling Configuration Versions using the `arn`. For example:
 *
 * ```sh
 *  $ pulumi import aws:apprunner/autoScalingConfigurationVersion:AutoScalingConfigurationVersion example "arn:aws:apprunner:us-east-1:1234567890:autoscalingconfiguration/example/1/69bdfe0115224b0db49398b7beb68e0f
 * ```
 */
export class AutoScalingConfigurationVersion extends pulumi.CustomResource {
    /**
     * Get an existing AutoScalingConfigurationVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AutoScalingConfigurationVersionState, opts?: pulumi.CustomResourceOptions): AutoScalingConfigurationVersion {
        return new AutoScalingConfigurationVersion(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apprunner/autoScalingConfigurationVersion:AutoScalingConfigurationVersion';

    /**
     * Returns true if the given object is an instance of AutoScalingConfigurationVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AutoScalingConfigurationVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AutoScalingConfigurationVersion.__pulumiType;
    }

    /**
     * ARN of this auto scaling configuration version.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Name of the auto scaling configuration.
     */
    public readonly autoScalingConfigurationName!: pulumi.Output<string>;
    /**
     * The revision of this auto scaling configuration.
     */
    public /*out*/ readonly autoScalingConfigurationRevision!: pulumi.Output<number>;
    /**
     * Whether the auto scaling configuration has the highest `autoScalingConfigurationRevision` among all configurations that share the same `autoScalingConfigurationName`.
     */
    public /*out*/ readonly latest!: pulumi.Output<boolean>;
    /**
     * Maximal number of concurrent requests that you want an instance to process. When the number of concurrent requests goes over this limit, App Runner scales up your service.
     */
    public readonly maxConcurrency!: pulumi.Output<number | undefined>;
    /**
     * Maximal number of instances that App Runner provisions for your service.
     */
    public readonly maxSize!: pulumi.Output<number | undefined>;
    /**
     * Minimal number of instances that App Runner provisions for your service.
     */
    public readonly minSize!: pulumi.Output<number | undefined>;
    /**
     * Current state of the auto scaling configuration. An INACTIVE configuration revision has been deleted and can't be used. It is permanently removed some time after deletion.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a AutoScalingConfigurationVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AutoScalingConfigurationVersionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AutoScalingConfigurationVersionArgs | AutoScalingConfigurationVersionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AutoScalingConfigurationVersionState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["autoScalingConfigurationName"] = state ? state.autoScalingConfigurationName : undefined;
            resourceInputs["autoScalingConfigurationRevision"] = state ? state.autoScalingConfigurationRevision : undefined;
            resourceInputs["latest"] = state ? state.latest : undefined;
            resourceInputs["maxConcurrency"] = state ? state.maxConcurrency : undefined;
            resourceInputs["maxSize"] = state ? state.maxSize : undefined;
            resourceInputs["minSize"] = state ? state.minSize : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as AutoScalingConfigurationVersionArgs | undefined;
            if ((!args || args.autoScalingConfigurationName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autoScalingConfigurationName'");
            }
            resourceInputs["autoScalingConfigurationName"] = args ? args.autoScalingConfigurationName : undefined;
            resourceInputs["maxConcurrency"] = args ? args.maxConcurrency : undefined;
            resourceInputs["maxSize"] = args ? args.maxSize : undefined;
            resourceInputs["minSize"] = args ? args.minSize : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["autoScalingConfigurationRevision"] = undefined /*out*/;
            resourceInputs["latest"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(AutoScalingConfigurationVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AutoScalingConfigurationVersion resources.
 */
export interface AutoScalingConfigurationVersionState {
    /**
     * ARN of this auto scaling configuration version.
     */
    arn?: pulumi.Input<string>;
    /**
     * Name of the auto scaling configuration.
     */
    autoScalingConfigurationName?: pulumi.Input<string>;
    /**
     * The revision of this auto scaling configuration.
     */
    autoScalingConfigurationRevision?: pulumi.Input<number>;
    /**
     * Whether the auto scaling configuration has the highest `autoScalingConfigurationRevision` among all configurations that share the same `autoScalingConfigurationName`.
     */
    latest?: pulumi.Input<boolean>;
    /**
     * Maximal number of concurrent requests that you want an instance to process. When the number of concurrent requests goes over this limit, App Runner scales up your service.
     */
    maxConcurrency?: pulumi.Input<number>;
    /**
     * Maximal number of instances that App Runner provisions for your service.
     */
    maxSize?: pulumi.Input<number>;
    /**
     * Minimal number of instances that App Runner provisions for your service.
     */
    minSize?: pulumi.Input<number>;
    /**
     * Current state of the auto scaling configuration. An INACTIVE configuration revision has been deleted and can't be used. It is permanently removed some time after deletion.
     */
    status?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a AutoScalingConfigurationVersion resource.
 */
export interface AutoScalingConfigurationVersionArgs {
    /**
     * Name of the auto scaling configuration.
     */
    autoScalingConfigurationName: pulumi.Input<string>;
    /**
     * Maximal number of concurrent requests that you want an instance to process. When the number of concurrent requests goes over this limit, App Runner scales up your service.
     */
    maxConcurrency?: pulumi.Input<number>;
    /**
     * Maximal number of instances that App Runner provisions for your service.
     */
    maxSize?: pulumi.Input<number>;
    /**
     * Minimal number of instances that App Runner provisions for your service.
     */
    minSize?: pulumi.Input<number>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
