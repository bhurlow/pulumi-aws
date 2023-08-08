// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a SageMaker model resource.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const assumeRole = aws.iam.getPolicyDocument({
 *     statements: [{
 *         actions: ["sts:AssumeRole"],
 *         principals: [{
 *             type: "Service",
 *             identifiers: ["sagemaker.amazonaws.com"],
 *         }],
 *     }],
 * });
 * const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
 * const test = aws.sagemaker.getPrebuiltEcrImage({
 *     repositoryName: "kmeans",
 * });
 * const exampleModel = new aws.sagemaker.Model("exampleModel", {
 *     executionRoleArn: exampleRole.arn,
 *     primaryContainer: {
 *         image: test.then(test => test.registryPath),
 *     },
 * });
 * ```
 * ## Inference Execution Config
 *
 * * `mode` - (Required) How containers in a multi-container are run. The following values are valid `Serial` and `Direct`.
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_sagemaker_model.test_model
 *
 *  id = "model-foo" } Using `pulumi import`, import models using the `name`. For exampleconsole % pulumi import aws_sagemaker_model.test_model model-foo
 */
export class Model extends pulumi.CustomResource {
    /**
     * Get an existing Model resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ModelState, opts?: pulumi.CustomResourceOptions): Model {
        return new Model(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:sagemaker/model:Model';

    /**
     * Returns true if the given object is an instance of Model.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Model {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Model.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this model.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Specifies containers in the inference pipeline. If not specified, the `primaryContainer` argument is required. Fields are documented below.
     */
    public readonly containers!: pulumi.Output<outputs.sagemaker.ModelContainer[] | undefined>;
    /**
     * Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
     */
    public readonly enableNetworkIsolation!: pulumi.Output<boolean | undefined>;
    /**
     * A role that SageMaker can assume to access model artifacts and docker images for deployment.
     */
    public readonly executionRoleArn!: pulumi.Output<string>;
    /**
     * Specifies details of how containers in a multi-container endpoint are called. see Inference Execution Config.
     */
    public readonly inferenceExecutionConfig!: pulumi.Output<outputs.sagemaker.ModelInferenceExecutionConfig>;
    /**
     * The name of the model (must be unique). If omitted, this provider will assign a random, unique name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the `container` argument is required. Fields are documented below.
     */
    public readonly primaryContainer!: pulumi.Output<outputs.sagemaker.ModelPrimaryContainer | undefined>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     *
     * The `primaryContainer` and `container` block both support:
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.
     */
    public readonly vpcConfig!: pulumi.Output<outputs.sagemaker.ModelVpcConfig | undefined>;

    /**
     * Create a Model resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ModelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ModelArgs | ModelState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ModelState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["containers"] = state ? state.containers : undefined;
            resourceInputs["enableNetworkIsolation"] = state ? state.enableNetworkIsolation : undefined;
            resourceInputs["executionRoleArn"] = state ? state.executionRoleArn : undefined;
            resourceInputs["inferenceExecutionConfig"] = state ? state.inferenceExecutionConfig : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["primaryContainer"] = state ? state.primaryContainer : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["vpcConfig"] = state ? state.vpcConfig : undefined;
        } else {
            const args = argsOrState as ModelArgs | undefined;
            if ((!args || args.executionRoleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'executionRoleArn'");
            }
            resourceInputs["containers"] = args ? args.containers : undefined;
            resourceInputs["enableNetworkIsolation"] = args ? args.enableNetworkIsolation : undefined;
            resourceInputs["executionRoleArn"] = args ? args.executionRoleArn : undefined;
            resourceInputs["inferenceExecutionConfig"] = args ? args.inferenceExecutionConfig : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["primaryContainer"] = args ? args.primaryContainer : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcConfig"] = args ? args.vpcConfig : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Model.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Model resources.
 */
export interface ModelState {
    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this model.
     */
    arn?: pulumi.Input<string>;
    /**
     * Specifies containers in the inference pipeline. If not specified, the `primaryContainer` argument is required. Fields are documented below.
     */
    containers?: pulumi.Input<pulumi.Input<inputs.sagemaker.ModelContainer>[]>;
    /**
     * Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
     */
    enableNetworkIsolation?: pulumi.Input<boolean>;
    /**
     * A role that SageMaker can assume to access model artifacts and docker images for deployment.
     */
    executionRoleArn?: pulumi.Input<string>;
    /**
     * Specifies details of how containers in a multi-container endpoint are called. see Inference Execution Config.
     */
    inferenceExecutionConfig?: pulumi.Input<inputs.sagemaker.ModelInferenceExecutionConfig>;
    /**
     * The name of the model (must be unique). If omitted, this provider will assign a random, unique name.
     */
    name?: pulumi.Input<string>;
    /**
     * The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the `container` argument is required. Fields are documented below.
     */
    primaryContainer?: pulumi.Input<inputs.sagemaker.ModelPrimaryContainer>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     *
     * The `primaryContainer` and `container` block both support:
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.
     */
    vpcConfig?: pulumi.Input<inputs.sagemaker.ModelVpcConfig>;
}

/**
 * The set of arguments for constructing a Model resource.
 */
export interface ModelArgs {
    /**
     * Specifies containers in the inference pipeline. If not specified, the `primaryContainer` argument is required. Fields are documented below.
     */
    containers?: pulumi.Input<pulumi.Input<inputs.sagemaker.ModelContainer>[]>;
    /**
     * Isolates the model container. No inbound or outbound network calls can be made to or from the model container.
     */
    enableNetworkIsolation?: pulumi.Input<boolean>;
    /**
     * A role that SageMaker can assume to access model artifacts and docker images for deployment.
     */
    executionRoleArn: pulumi.Input<string>;
    /**
     * Specifies details of how containers in a multi-container endpoint are called. see Inference Execution Config.
     */
    inferenceExecutionConfig?: pulumi.Input<inputs.sagemaker.ModelInferenceExecutionConfig>;
    /**
     * The name of the model (must be unique). If omitted, this provider will assign a random, unique name.
     */
    name?: pulumi.Input<string>;
    /**
     * The primary docker image containing inference code that is used when the model is deployed for predictions.  If not specified, the `container` argument is required. Fields are documented below.
     */
    primaryContainer?: pulumi.Input<inputs.sagemaker.ModelPrimaryContainer>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     *
     * The `primaryContainer` and `container` block both support:
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the VPC that you want your model to connect to. VpcConfig is used in hosting services and in batch transform.
     */
    vpcConfig?: pulumi.Input<inputs.sagemaker.ModelVpcConfig>;
}
