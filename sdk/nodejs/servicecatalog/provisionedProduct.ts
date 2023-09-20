// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * This resource provisions and manages a Service Catalog provisioned product.
 *
 * A provisioned product is a resourced instance of a product. For example, provisioning a product based on a CloudFormation template launches a CloudFormation stack and its underlying resources.
 *
 * Like this resource, the `awsServicecatalogRecord` data source also provides information about a provisioned product. Although a Service Catalog record provides some overlapping information with this resource, a record is tied to a provisioned product event, such as provisioning, termination, and updating.
 *
 * > **Tip:** If you include conflicted keys as tags, AWS will report an error, "Parameter validation failed: Missing required parameter in Tags[N]:Value".
 *
 * > **Tip:** A "provisioning artifact" is also referred to as a "version." A "distributor" is also referred to as a "vendor."
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.servicecatalog.ProvisionedProduct("example", {
 *     productName: "Example product",
 *     provisioningArtifactName: "Example version",
 *     provisioningParameters: [{
 *         key: "foo",
 *         value: "bar",
 *     }],
 *     tags: {
 *         foo: "bar",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Using `pulumi import`, import `aws_servicecatalog_provisioned_product` using the provisioned product ID. For example:
 *
 * ```sh
 *  $ pulumi import aws:servicecatalog/provisionedProduct:ProvisionedProduct example pp-dnigbtea24ste
 * ```
 */
export class ProvisionedProduct extends pulumi.CustomResource {
    /**
     * Get an existing ProvisionedProduct resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProvisionedProductState, opts?: pulumi.CustomResourceOptions): ProvisionedProduct {
        return new ProvisionedProduct(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:servicecatalog/provisionedProduct:ProvisionedProduct';

    /**
     * Returns true if the given object is an instance of ProvisionedProduct.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProvisionedProduct {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProvisionedProduct.__pulumiType;
    }

    /**
     * Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
     */
    public readonly acceptLanguage!: pulumi.Output<string | undefined>;
    /**
     * ARN of the provisioned product.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Set of CloudWatch dashboards that were created when provisioning the product.
     */
    public /*out*/ readonly cloudwatchDashboardNames!: pulumi.Output<string[]>;
    /**
     * Time when the provisioned product was created.
     */
    public /*out*/ readonly createdTime!: pulumi.Output<string>;
    /**
     * _Only applies to deleting._ If set to `true`, AWS Service Catalog stops managing the specified provisioned product even if it cannot delete the underlying resources. The default value is `false`.
     */
    public readonly ignoreErrors!: pulumi.Output<boolean | undefined>;
    /**
     * Record identifier of the last request performed on this provisioned product of the following types: `ProvisionedProduct`, `UpdateProvisionedProduct`, `ExecuteProvisionedProductPlan`, `TerminateProvisionedProduct`.
     */
    public /*out*/ readonly lastProvisioningRecordId!: pulumi.Output<string>;
    /**
     * Record identifier of the last request performed on this provisioned product.
     */
    public /*out*/ readonly lastRecordId!: pulumi.Output<string>;
    /**
     * Record identifier of the last successful request performed on this provisioned product of the following types: `ProvisionedProduct`, `UpdateProvisionedProduct`, `ExecuteProvisionedProductPlan`, `TerminateProvisionedProduct`.
     */
    public /*out*/ readonly lastSuccessfulProvisioningRecordId!: pulumi.Output<string>;
    /**
     * ARN of the launch role associated with the provisioned product.
     */
    public /*out*/ readonly launchRoleArn!: pulumi.Output<string>;
    /**
     * User-friendly name of the provisioned product.
     *
     * The following arguments are optional:
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Passed to CloudFormation. The SNS topic ARNs to which to publish stack-related events.
     */
    public readonly notificationArns!: pulumi.Output<string[] | undefined>;
    /**
     * The set of outputs for the product created.
     */
    public /*out*/ readonly outputs!: pulumi.Output<outputs.servicecatalog.ProvisionedProductOutput[]>;
    /**
     * Path identifier of the product. This value is optional if the product has a default path, and required if the product has more than one path. To list the paths for a product, use `aws.servicecatalog.getLaunchPaths`. When required, you must provide `pathId` or `pathName`, but not both.
     */
    public readonly pathId!: pulumi.Output<string>;
    /**
     * Name of the path. You must provide `pathId` or `pathName`, but not both.
     */
    public readonly pathName!: pulumi.Output<string | undefined>;
    /**
     * Product identifier. For example, `prod-abcdzk7xy33qa`. You must provide `productId` or `productName`, but not both.
     */
    public readonly productId!: pulumi.Output<string>;
    /**
     * Name of the product. You must provide `productId` or `productName`, but not both.
     */
    public readonly productName!: pulumi.Output<string | undefined>;
    /**
     * Identifier of the provisioning artifact. For example, `pa-4abcdjnxjj6ne`. You must provide the `provisioningArtifactId` or `provisioningArtifactName`, but not both.
     */
    public readonly provisioningArtifactId!: pulumi.Output<string>;
    /**
     * Name of the provisioning artifact. You must provide the `provisioningArtifactId` or `provisioningArtifactName`, but not both.
     */
    public readonly provisioningArtifactName!: pulumi.Output<string | undefined>;
    /**
     * Configuration block with parameters specified by the administrator that are required for provisioning the product. See details below.
     */
    public readonly provisioningParameters!: pulumi.Output<outputs.servicecatalog.ProvisionedProductProvisioningParameter[] | undefined>;
    /**
     * _Only applies to deleting._ Whether to delete the Service Catalog provisioned product but leave the CloudFormation stack, stack set, or the underlying resources of the deleted provisioned product. The default value is `false`.
     */
    public readonly retainPhysicalResources!: pulumi.Output<boolean | undefined>;
    /**
     * Configuration block with information about the provisioning preferences for a stack set. See details below.
     */
    public readonly stackSetProvisioningPreferences!: pulumi.Output<outputs.servicecatalog.ProvisionedProductStackSetProvisioningPreferences | undefined>;
    /**
     * Current status of the provisioned product. See meanings below.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Current status message of the provisioned product.
     */
    public /*out*/ readonly statusMessage!: pulumi.Output<string>;
    /**
     * Tags to apply to the provisioned product. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Type of provisioned product. Valid values are `CFN_STACK` and `CFN_STACKSET`.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ProvisionedProduct resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProvisionedProductArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProvisionedProductArgs | ProvisionedProductState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProvisionedProductState | undefined;
            resourceInputs["acceptLanguage"] = state ? state.acceptLanguage : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["cloudwatchDashboardNames"] = state ? state.cloudwatchDashboardNames : undefined;
            resourceInputs["createdTime"] = state ? state.createdTime : undefined;
            resourceInputs["ignoreErrors"] = state ? state.ignoreErrors : undefined;
            resourceInputs["lastProvisioningRecordId"] = state ? state.lastProvisioningRecordId : undefined;
            resourceInputs["lastRecordId"] = state ? state.lastRecordId : undefined;
            resourceInputs["lastSuccessfulProvisioningRecordId"] = state ? state.lastSuccessfulProvisioningRecordId : undefined;
            resourceInputs["launchRoleArn"] = state ? state.launchRoleArn : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notificationArns"] = state ? state.notificationArns : undefined;
            resourceInputs["outputs"] = state ? state.outputs : undefined;
            resourceInputs["pathId"] = state ? state.pathId : undefined;
            resourceInputs["pathName"] = state ? state.pathName : undefined;
            resourceInputs["productId"] = state ? state.productId : undefined;
            resourceInputs["productName"] = state ? state.productName : undefined;
            resourceInputs["provisioningArtifactId"] = state ? state.provisioningArtifactId : undefined;
            resourceInputs["provisioningArtifactName"] = state ? state.provisioningArtifactName : undefined;
            resourceInputs["provisioningParameters"] = state ? state.provisioningParameters : undefined;
            resourceInputs["retainPhysicalResources"] = state ? state.retainPhysicalResources : undefined;
            resourceInputs["stackSetProvisioningPreferences"] = state ? state.stackSetProvisioningPreferences : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["statusMessage"] = state ? state.statusMessage : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ProvisionedProductArgs | undefined;
            resourceInputs["acceptLanguage"] = args ? args.acceptLanguage : undefined;
            resourceInputs["ignoreErrors"] = args ? args.ignoreErrors : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notificationArns"] = args ? args.notificationArns : undefined;
            resourceInputs["pathId"] = args ? args.pathId : undefined;
            resourceInputs["pathName"] = args ? args.pathName : undefined;
            resourceInputs["productId"] = args ? args.productId : undefined;
            resourceInputs["productName"] = args ? args.productName : undefined;
            resourceInputs["provisioningArtifactId"] = args ? args.provisioningArtifactId : undefined;
            resourceInputs["provisioningArtifactName"] = args ? args.provisioningArtifactName : undefined;
            resourceInputs["provisioningParameters"] = args ? args.provisioningParameters : undefined;
            resourceInputs["retainPhysicalResources"] = args ? args.retainPhysicalResources : undefined;
            resourceInputs["stackSetProvisioningPreferences"] = args ? args.stackSetProvisioningPreferences : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["cloudwatchDashboardNames"] = undefined /*out*/;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["lastProvisioningRecordId"] = undefined /*out*/;
            resourceInputs["lastRecordId"] = undefined /*out*/;
            resourceInputs["lastSuccessfulProvisioningRecordId"] = undefined /*out*/;
            resourceInputs["launchRoleArn"] = undefined /*out*/;
            resourceInputs["outputs"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["statusMessage"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ProvisionedProduct.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProvisionedProduct resources.
 */
export interface ProvisionedProductState {
    /**
     * Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
     */
    acceptLanguage?: pulumi.Input<string>;
    /**
     * ARN of the provisioned product.
     */
    arn?: pulumi.Input<string>;
    /**
     * Set of CloudWatch dashboards that were created when provisioning the product.
     */
    cloudwatchDashboardNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Time when the provisioned product was created.
     */
    createdTime?: pulumi.Input<string>;
    /**
     * _Only applies to deleting._ If set to `true`, AWS Service Catalog stops managing the specified provisioned product even if it cannot delete the underlying resources. The default value is `false`.
     */
    ignoreErrors?: pulumi.Input<boolean>;
    /**
     * Record identifier of the last request performed on this provisioned product of the following types: `ProvisionedProduct`, `UpdateProvisionedProduct`, `ExecuteProvisionedProductPlan`, `TerminateProvisionedProduct`.
     */
    lastProvisioningRecordId?: pulumi.Input<string>;
    /**
     * Record identifier of the last request performed on this provisioned product.
     */
    lastRecordId?: pulumi.Input<string>;
    /**
     * Record identifier of the last successful request performed on this provisioned product of the following types: `ProvisionedProduct`, `UpdateProvisionedProduct`, `ExecuteProvisionedProductPlan`, `TerminateProvisionedProduct`.
     */
    lastSuccessfulProvisioningRecordId?: pulumi.Input<string>;
    /**
     * ARN of the launch role associated with the provisioned product.
     */
    launchRoleArn?: pulumi.Input<string>;
    /**
     * User-friendly name of the provisioned product.
     *
     * The following arguments are optional:
     */
    name?: pulumi.Input<string>;
    /**
     * Passed to CloudFormation. The SNS topic ARNs to which to publish stack-related events.
     */
    notificationArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The set of outputs for the product created.
     */
    outputs?: pulumi.Input<pulumi.Input<inputs.servicecatalog.ProvisionedProductOutput>[]>;
    /**
     * Path identifier of the product. This value is optional if the product has a default path, and required if the product has more than one path. To list the paths for a product, use `aws.servicecatalog.getLaunchPaths`. When required, you must provide `pathId` or `pathName`, but not both.
     */
    pathId?: pulumi.Input<string>;
    /**
     * Name of the path. You must provide `pathId` or `pathName`, but not both.
     */
    pathName?: pulumi.Input<string>;
    /**
     * Product identifier. For example, `prod-abcdzk7xy33qa`. You must provide `productId` or `productName`, but not both.
     */
    productId?: pulumi.Input<string>;
    /**
     * Name of the product. You must provide `productId` or `productName`, but not both.
     */
    productName?: pulumi.Input<string>;
    /**
     * Identifier of the provisioning artifact. For example, `pa-4abcdjnxjj6ne`. You must provide the `provisioningArtifactId` or `provisioningArtifactName`, but not both.
     */
    provisioningArtifactId?: pulumi.Input<string>;
    /**
     * Name of the provisioning artifact. You must provide the `provisioningArtifactId` or `provisioningArtifactName`, but not both.
     */
    provisioningArtifactName?: pulumi.Input<string>;
    /**
     * Configuration block with parameters specified by the administrator that are required for provisioning the product. See details below.
     */
    provisioningParameters?: pulumi.Input<pulumi.Input<inputs.servicecatalog.ProvisionedProductProvisioningParameter>[]>;
    /**
     * _Only applies to deleting._ Whether to delete the Service Catalog provisioned product but leave the CloudFormation stack, stack set, or the underlying resources of the deleted provisioned product. The default value is `false`.
     */
    retainPhysicalResources?: pulumi.Input<boolean>;
    /**
     * Configuration block with information about the provisioning preferences for a stack set. See details below.
     */
    stackSetProvisioningPreferences?: pulumi.Input<inputs.servicecatalog.ProvisionedProductStackSetProvisioningPreferences>;
    /**
     * Current status of the provisioned product. See meanings below.
     */
    status?: pulumi.Input<string>;
    /**
     * Current status message of the provisioned product.
     */
    statusMessage?: pulumi.Input<string>;
    /**
     * Tags to apply to the provisioned product. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Type of provisioned product. Valid values are `CFN_STACK` and `CFN_STACKSET`.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProvisionedProduct resource.
 */
export interface ProvisionedProductArgs {
    /**
     * Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
     */
    acceptLanguage?: pulumi.Input<string>;
    /**
     * _Only applies to deleting._ If set to `true`, AWS Service Catalog stops managing the specified provisioned product even if it cannot delete the underlying resources. The default value is `false`.
     */
    ignoreErrors?: pulumi.Input<boolean>;
    /**
     * User-friendly name of the provisioned product.
     *
     * The following arguments are optional:
     */
    name?: pulumi.Input<string>;
    /**
     * Passed to CloudFormation. The SNS topic ARNs to which to publish stack-related events.
     */
    notificationArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Path identifier of the product. This value is optional if the product has a default path, and required if the product has more than one path. To list the paths for a product, use `aws.servicecatalog.getLaunchPaths`. When required, you must provide `pathId` or `pathName`, but not both.
     */
    pathId?: pulumi.Input<string>;
    /**
     * Name of the path. You must provide `pathId` or `pathName`, but not both.
     */
    pathName?: pulumi.Input<string>;
    /**
     * Product identifier. For example, `prod-abcdzk7xy33qa`. You must provide `productId` or `productName`, but not both.
     */
    productId?: pulumi.Input<string>;
    /**
     * Name of the product. You must provide `productId` or `productName`, but not both.
     */
    productName?: pulumi.Input<string>;
    /**
     * Identifier of the provisioning artifact. For example, `pa-4abcdjnxjj6ne`. You must provide the `provisioningArtifactId` or `provisioningArtifactName`, but not both.
     */
    provisioningArtifactId?: pulumi.Input<string>;
    /**
     * Name of the provisioning artifact. You must provide the `provisioningArtifactId` or `provisioningArtifactName`, but not both.
     */
    provisioningArtifactName?: pulumi.Input<string>;
    /**
     * Configuration block with parameters specified by the administrator that are required for provisioning the product. See details below.
     */
    provisioningParameters?: pulumi.Input<pulumi.Input<inputs.servicecatalog.ProvisionedProductProvisioningParameter>[]>;
    /**
     * _Only applies to deleting._ Whether to delete the Service Catalog provisioned product but leave the CloudFormation stack, stack set, or the underlying resources of the deleted provisioned product. The default value is `false`.
     */
    retainPhysicalResources?: pulumi.Input<boolean>;
    /**
     * Configuration block with information about the provisioning preferences for a stack set. See details below.
     */
    stackSetProvisioningPreferences?: pulumi.Input<inputs.servicecatalog.ProvisionedProductStackSetProvisioningPreferences>;
    /**
     * Tags to apply to the provisioned product. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
