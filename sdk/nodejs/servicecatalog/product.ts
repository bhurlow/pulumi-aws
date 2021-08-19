// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Service Catalog Product.
 *
 * > **NOTE:** The user or role that uses this resources must have the `cloudformation:GetTemplate` IAM policy permission. This policy permission is required when using the `templatePhysicalId` argument.
 *
 * > A "provisioning artifact" is also referred to as a "version." A "distributor" is also referred to as a "vendor."
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.servicecatalog.Product("example", {
 *     owner: [aws_security_group.example.id],
 *     type: aws_subnet.main.id,
 *     provisioningArtifactParameters: {
 *         templateUrl: "https://s3.amazonaws.com/cf-templates-ozkq9d3hgiq2-us-east-1/temp1.json",
 *     },
 *     tags: {
 *         foo: "bar",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * `aws_servicecatalog_product` can be imported using the product ID, e.g.
 *
 * ```sh
 *  $ pulumi import aws:servicecatalog/product:Product example prod-dnigbtea24ste
 * ```
 */
export class Product extends pulumi.CustomResource {
    /**
     * Get an existing Product resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProductState, opts?: pulumi.CustomResourceOptions): Product {
        return new Product(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:servicecatalog/product:Product';

    /**
     * Returns true if the given object is an instance of Product.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Product {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Product.__pulumiType;
    }

    /**
     * Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
     */
    public readonly acceptLanguage!: pulumi.Output<string | undefined>;
    /**
     * ARN of the product.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Time when the product was created.
     */
    public /*out*/ readonly createdTime!: pulumi.Output<string>;
    /**
     * Description of the provisioning artifact (i.e., version), including how it differs from the previous provisioning artifact.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Distributor (i.e., vendor) of the product.
     */
    public readonly distributor!: pulumi.Output<string>;
    /**
     * Whether the product has a default path. If the product does not have a default path, call `ListLaunchPaths` to disambiguate between paths.  Otherwise, `ListLaunchPaths` is not required, and the output of ProductViewSummary can be used directly with `DescribeProvisioningParameters`.
     */
    public /*out*/ readonly hasDefaultPath!: pulumi.Output<boolean>;
    /**
     * Name of the provisioning artifact (for example, `v1`, `v2beta`). No spaces are allowed.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Owner of the product.
     */
    public readonly owner!: pulumi.Output<string>;
    /**
     * Configuration block for provisioning artifact (i.e., version) parameters. Detailed below.
     */
    public readonly provisioningArtifactParameters!: pulumi.Output<outputs.servicecatalog.ProductProvisioningArtifactParameters>;
    /**
     * Status of the product.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Support information about the product.
     */
    public readonly supportDescription!: pulumi.Output<string>;
    /**
     * Contact email for product support.
     */
    public readonly supportEmail!: pulumi.Output<string>;
    /**
     * Contact URL for product support.
     */
    public readonly supportUrl!: pulumi.Output<string>;
    /**
     * Tags to apply to the product. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Type of provisioning artifact. Valid values: `CLOUD_FORMATION_TEMPLATE`, `MARKETPLACE_AMI`, `MARKETPLACE_CAR` (Marketplace Clusters and AWS Resources).
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Product resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProductArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProductArgs | ProductState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProductState | undefined;
            inputs["acceptLanguage"] = state ? state.acceptLanguage : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["createdTime"] = state ? state.createdTime : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["distributor"] = state ? state.distributor : undefined;
            inputs["hasDefaultPath"] = state ? state.hasDefaultPath : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["owner"] = state ? state.owner : undefined;
            inputs["provisioningArtifactParameters"] = state ? state.provisioningArtifactParameters : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["supportDescription"] = state ? state.supportDescription : undefined;
            inputs["supportEmail"] = state ? state.supportEmail : undefined;
            inputs["supportUrl"] = state ? state.supportUrl : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ProductArgs | undefined;
            if ((!args || args.owner === undefined) && !opts.urn) {
                throw new Error("Missing required property 'owner'");
            }
            if ((!args || args.provisioningArtifactParameters === undefined) && !opts.urn) {
                throw new Error("Missing required property 'provisioningArtifactParameters'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            inputs["acceptLanguage"] = args ? args.acceptLanguage : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["distributor"] = args ? args.distributor : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["owner"] = args ? args.owner : undefined;
            inputs["provisioningArtifactParameters"] = args ? args.provisioningArtifactParameters : undefined;
            inputs["supportDescription"] = args ? args.supportDescription : undefined;
            inputs["supportEmail"] = args ? args.supportEmail : undefined;
            inputs["supportUrl"] = args ? args.supportUrl : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["createdTime"] = undefined /*out*/;
            inputs["hasDefaultPath"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["tagsAll"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Product.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Product resources.
 */
export interface ProductState {
    /**
     * Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
     */
    acceptLanguage?: pulumi.Input<string>;
    /**
     * ARN of the product.
     */
    arn?: pulumi.Input<string>;
    /**
     * Time when the product was created.
     */
    createdTime?: pulumi.Input<string>;
    /**
     * Description of the provisioning artifact (i.e., version), including how it differs from the previous provisioning artifact.
     */
    description?: pulumi.Input<string>;
    /**
     * Distributor (i.e., vendor) of the product.
     */
    distributor?: pulumi.Input<string>;
    /**
     * Whether the product has a default path. If the product does not have a default path, call `ListLaunchPaths` to disambiguate between paths.  Otherwise, `ListLaunchPaths` is not required, and the output of ProductViewSummary can be used directly with `DescribeProvisioningParameters`.
     */
    hasDefaultPath?: pulumi.Input<boolean>;
    /**
     * Name of the provisioning artifact (for example, `v1`, `v2beta`). No spaces are allowed.
     */
    name?: pulumi.Input<string>;
    /**
     * Owner of the product.
     */
    owner?: pulumi.Input<string>;
    /**
     * Configuration block for provisioning artifact (i.e., version) parameters. Detailed below.
     */
    provisioningArtifactParameters?: pulumi.Input<inputs.servicecatalog.ProductProvisioningArtifactParameters>;
    /**
     * Status of the product.
     */
    status?: pulumi.Input<string>;
    /**
     * Support information about the product.
     */
    supportDescription?: pulumi.Input<string>;
    /**
     * Contact email for product support.
     */
    supportEmail?: pulumi.Input<string>;
    /**
     * Contact URL for product support.
     */
    supportUrl?: pulumi.Input<string>;
    /**
     * Tags to apply to the product. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Type of provisioning artifact. Valid values: `CLOUD_FORMATION_TEMPLATE`, `MARKETPLACE_AMI`, `MARKETPLACE_CAR` (Marketplace Clusters and AWS Resources).
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Product resource.
 */
export interface ProductArgs {
    /**
     * Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
     */
    acceptLanguage?: pulumi.Input<string>;
    /**
     * Description of the provisioning artifact (i.e., version), including how it differs from the previous provisioning artifact.
     */
    description?: pulumi.Input<string>;
    /**
     * Distributor (i.e., vendor) of the product.
     */
    distributor?: pulumi.Input<string>;
    /**
     * Name of the provisioning artifact (for example, `v1`, `v2beta`). No spaces are allowed.
     */
    name?: pulumi.Input<string>;
    /**
     * Owner of the product.
     */
    owner: pulumi.Input<string>;
    /**
     * Configuration block for provisioning artifact (i.e., version) parameters. Detailed below.
     */
    provisioningArtifactParameters: pulumi.Input<inputs.servicecatalog.ProductProvisioningArtifactParameters>;
    /**
     * Support information about the product.
     */
    supportDescription?: pulumi.Input<string>;
    /**
     * Contact email for product support.
     */
    supportEmail?: pulumi.Input<string>;
    /**
     * Contact URL for product support.
     */
    supportUrl?: pulumi.Input<string>;
    /**
     * Tags to apply to the product. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Type of provisioning artifact. Valid values: `CLOUD_FORMATION_TEMPLATE`, `MARKETPLACE_AMI`, `MARKETPLACE_CAR` (Marketplace Clusters and AWS Resources).
     */
    type: pulumi.Input<string>;
}
