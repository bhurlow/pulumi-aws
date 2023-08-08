// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a SageMaker App resource.
 *
 * ## Example Usage
 * ### Basic usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.sagemaker.App("example", {
 *     domainId: aws_sagemaker_domain.example.id,
 *     userProfileName: aws_sagemaker_user_profile.example.user_profile_name,
 *     appName: "example",
 *     appType: "JupyterServer",
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_sagemaker_app.example
 *
 *  id = "arn:aws:sagemaker:us-west-2:012345678912:app/domain-id/user-profile-name/app-type/app-name" } Using `pulumi import`, import SageMaker Apps using the `id`. For exampleconsole % pulumi import aws_sagemaker_app.example arn:aws:sagemaker:us-west-2:012345678912:app/domain-id/user-profile-name/app-type/app-name
 */
export class App extends pulumi.CustomResource {
    /**
     * Get an existing App resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppState, opts?: pulumi.CustomResourceOptions): App {
        return new App(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:sagemaker/app:App';

    /**
     * Returns true if the given object is an instance of App.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is App {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === App.__pulumiType;
    }

    /**
     * The name of the app.
     */
    public readonly appName!: pulumi.Output<string>;
    /**
     * The type of app. Valid values are `JupyterServer`, `KernelGateway`, `RStudioServerPro`, `RSessionGateway` and `TensorBoard`.
     */
    public readonly appType!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the app.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The domain ID.
     */
    public readonly domainId!: pulumi.Output<string>;
    /**
     * The instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance.See Resource Spec below.
     */
    public readonly resourceSpec!: pulumi.Output<outputs.sagemaker.AppResourceSpec>;
    /**
     * The name of the space. At least one of `userProfileName` or `spaceName` required.
     */
    public readonly spaceName!: pulumi.Output<string | undefined>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The user profile name. At least one of `userProfileName` or `spaceName` required.
     */
    public readonly userProfileName!: pulumi.Output<string | undefined>;

    /**
     * Create a App resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppArgs | AppState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppState | undefined;
            resourceInputs["appName"] = state ? state.appName : undefined;
            resourceInputs["appType"] = state ? state.appType : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["domainId"] = state ? state.domainId : undefined;
            resourceInputs["resourceSpec"] = state ? state.resourceSpec : undefined;
            resourceInputs["spaceName"] = state ? state.spaceName : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["userProfileName"] = state ? state.userProfileName : undefined;
        } else {
            const args = argsOrState as AppArgs | undefined;
            if ((!args || args.appName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appName'");
            }
            if ((!args || args.appType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appType'");
            }
            if ((!args || args.domainId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainId'");
            }
            resourceInputs["appName"] = args ? args.appName : undefined;
            resourceInputs["appType"] = args ? args.appType : undefined;
            resourceInputs["domainId"] = args ? args.domainId : undefined;
            resourceInputs["resourceSpec"] = args ? args.resourceSpec : undefined;
            resourceInputs["spaceName"] = args ? args.spaceName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["userProfileName"] = args ? args.userProfileName : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(App.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering App resources.
 */
export interface AppState {
    /**
     * The name of the app.
     */
    appName?: pulumi.Input<string>;
    /**
     * The type of app. Valid values are `JupyterServer`, `KernelGateway`, `RStudioServerPro`, `RSessionGateway` and `TensorBoard`.
     */
    appType?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the app.
     */
    arn?: pulumi.Input<string>;
    /**
     * The domain ID.
     */
    domainId?: pulumi.Input<string>;
    /**
     * The instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance.See Resource Spec below.
     */
    resourceSpec?: pulumi.Input<inputs.sagemaker.AppResourceSpec>;
    /**
     * The name of the space. At least one of `userProfileName` or `spaceName` required.
     */
    spaceName?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The user profile name. At least one of `userProfileName` or `spaceName` required.
     */
    userProfileName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a App resource.
 */
export interface AppArgs {
    /**
     * The name of the app.
     */
    appName: pulumi.Input<string>;
    /**
     * The type of app. Valid values are `JupyterServer`, `KernelGateway`, `RStudioServerPro`, `RSessionGateway` and `TensorBoard`.
     */
    appType: pulumi.Input<string>;
    /**
     * The domain ID.
     */
    domainId: pulumi.Input<string>;
    /**
     * The instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance.See Resource Spec below.
     */
    resourceSpec?: pulumi.Input<inputs.sagemaker.AppResourceSpec>;
    /**
     * The name of the space. At least one of `userProfileName` or `spaceName` required.
     */
    spaceName?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The user profile name. At least one of `userProfileName` or `spaceName` required.
     */
    userProfileName?: pulumi.Input<string>;
}
