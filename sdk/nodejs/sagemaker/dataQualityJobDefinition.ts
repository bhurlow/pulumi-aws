// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a SageMaker data quality job definition resource.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.sagemaker.DataQualityJobDefinition("test", {
 *     dataQualityAppSpecification: {
 *         imageUri: data.aws_sagemaker_prebuilt_ecr_image.monitor.registry_path,
 *     },
 *     dataQualityJobInput: {
 *         endpointInput: {
 *             endpointName: aws_sagemaker_endpoint.my_endpoint.name,
 *         },
 *     },
 *     dataQualityJobOutputConfig: {
 *         monitoringOutputs: {
 *             s3Output: {
 *                 s3Uri: `https://${aws_s3_bucket.my_bucket.bucket_regional_domain_name}/output`,
 *             },
 *         },
 *     },
 *     jobResources: {
 *         clusterConfig: {
 *             instanceCount: 1,
 *             instanceType: "ml.t3.medium",
 *             volumeSizeInGb: 20,
 *         },
 *     },
 *     roleArn: aws_iam_role.my_role.arn,
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_sagemaker_data_quality_job_definition.test_data_quality_job_definition
 *
 *  id = "data-quality-job-definition-foo" } Using `pulumi import`, import data quality job definitions using the `name`. For exampleconsole % pulumi import aws_sagemaker_data_quality_job_definition.test_data_quality_job_definition data-quality-job-definition-foo
 */
export class DataQualityJobDefinition extends pulumi.CustomResource {
    /**
     * Get an existing DataQualityJobDefinition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataQualityJobDefinitionState, opts?: pulumi.CustomResourceOptions): DataQualityJobDefinition {
        return new DataQualityJobDefinition(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:sagemaker/dataQualityJobDefinition:DataQualityJobDefinition';

    /**
     * Returns true if the given object is an instance of DataQualityJobDefinition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataQualityJobDefinition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataQualityJobDefinition.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this data quality job definition.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Specifies the container that runs the monitoring job. Fields are documented below.
     */
    public readonly dataQualityAppSpecification!: pulumi.Output<outputs.sagemaker.DataQualityJobDefinitionDataQualityAppSpecification>;
    /**
     * Configures the constraints and baselines for the monitoring job. Fields are documented below.
     */
    public readonly dataQualityBaselineConfig!: pulumi.Output<outputs.sagemaker.DataQualityJobDefinitionDataQualityBaselineConfig | undefined>;
    /**
     * A list of inputs for the monitoring job. Fields are documented below.
     */
    public readonly dataQualityJobInput!: pulumi.Output<outputs.sagemaker.DataQualityJobDefinitionDataQualityJobInput>;
    /**
     * The output configuration for monitoring jobs. Fields are documented below.
     */
    public readonly dataQualityJobOutputConfig!: pulumi.Output<outputs.sagemaker.DataQualityJobDefinitionDataQualityJobOutputConfig>;
    /**
     * Identifies the resources to deploy for a monitoring job. Fields are documented below.
     */
    public readonly jobResources!: pulumi.Output<outputs.sagemaker.DataQualityJobDefinitionJobResources>;
    /**
     * The name of the data quality job definition. If omitted, the provider will assign a random, unique name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies networking configuration for the monitoring job. Fields are documented below.
     */
    public readonly networkConfig!: pulumi.Output<outputs.sagemaker.DataQualityJobDefinitionNetworkConfig | undefined>;
    /**
     * The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * A time limit for how long the monitoring job is allowed to run before stopping. Fields are documented below.
     */
    public readonly stoppingCondition!: pulumi.Output<outputs.sagemaker.DataQualityJobDefinitionStoppingCondition>;
    /**
     * A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a DataQualityJobDefinition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataQualityJobDefinitionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataQualityJobDefinitionArgs | DataQualityJobDefinitionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DataQualityJobDefinitionState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["dataQualityAppSpecification"] = state ? state.dataQualityAppSpecification : undefined;
            resourceInputs["dataQualityBaselineConfig"] = state ? state.dataQualityBaselineConfig : undefined;
            resourceInputs["dataQualityJobInput"] = state ? state.dataQualityJobInput : undefined;
            resourceInputs["dataQualityJobOutputConfig"] = state ? state.dataQualityJobOutputConfig : undefined;
            resourceInputs["jobResources"] = state ? state.jobResources : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkConfig"] = state ? state.networkConfig : undefined;
            resourceInputs["roleArn"] = state ? state.roleArn : undefined;
            resourceInputs["stoppingCondition"] = state ? state.stoppingCondition : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as DataQualityJobDefinitionArgs | undefined;
            if ((!args || args.dataQualityAppSpecification === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataQualityAppSpecification'");
            }
            if ((!args || args.dataQualityJobInput === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataQualityJobInput'");
            }
            if ((!args || args.dataQualityJobOutputConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataQualityJobOutputConfig'");
            }
            if ((!args || args.jobResources === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jobResources'");
            }
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            resourceInputs["dataQualityAppSpecification"] = args ? args.dataQualityAppSpecification : undefined;
            resourceInputs["dataQualityBaselineConfig"] = args ? args.dataQualityBaselineConfig : undefined;
            resourceInputs["dataQualityJobInput"] = args ? args.dataQualityJobInput : undefined;
            resourceInputs["dataQualityJobOutputConfig"] = args ? args.dataQualityJobOutputConfig : undefined;
            resourceInputs["jobResources"] = args ? args.jobResources : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkConfig"] = args ? args.networkConfig : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["stoppingCondition"] = args ? args.stoppingCondition : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DataQualityJobDefinition.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataQualityJobDefinition resources.
 */
export interface DataQualityJobDefinitionState {
    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this data quality job definition.
     */
    arn?: pulumi.Input<string>;
    /**
     * Specifies the container that runs the monitoring job. Fields are documented below.
     */
    dataQualityAppSpecification?: pulumi.Input<inputs.sagemaker.DataQualityJobDefinitionDataQualityAppSpecification>;
    /**
     * Configures the constraints and baselines for the monitoring job. Fields are documented below.
     */
    dataQualityBaselineConfig?: pulumi.Input<inputs.sagemaker.DataQualityJobDefinitionDataQualityBaselineConfig>;
    /**
     * A list of inputs for the monitoring job. Fields are documented below.
     */
    dataQualityJobInput?: pulumi.Input<inputs.sagemaker.DataQualityJobDefinitionDataQualityJobInput>;
    /**
     * The output configuration for monitoring jobs. Fields are documented below.
     */
    dataQualityJobOutputConfig?: pulumi.Input<inputs.sagemaker.DataQualityJobDefinitionDataQualityJobOutputConfig>;
    /**
     * Identifies the resources to deploy for a monitoring job. Fields are documented below.
     */
    jobResources?: pulumi.Input<inputs.sagemaker.DataQualityJobDefinitionJobResources>;
    /**
     * The name of the data quality job definition. If omitted, the provider will assign a random, unique name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies networking configuration for the monitoring job. Fields are documented below.
     */
    networkConfig?: pulumi.Input<inputs.sagemaker.DataQualityJobDefinitionNetworkConfig>;
    /**
     * The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
     */
    roleArn?: pulumi.Input<string>;
    /**
     * A time limit for how long the monitoring job is allowed to run before stopping. Fields are documented below.
     */
    stoppingCondition?: pulumi.Input<inputs.sagemaker.DataQualityJobDefinitionStoppingCondition>;
    /**
     * A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a DataQualityJobDefinition resource.
 */
export interface DataQualityJobDefinitionArgs {
    /**
     * Specifies the container that runs the monitoring job. Fields are documented below.
     */
    dataQualityAppSpecification: pulumi.Input<inputs.sagemaker.DataQualityJobDefinitionDataQualityAppSpecification>;
    /**
     * Configures the constraints and baselines for the monitoring job. Fields are documented below.
     */
    dataQualityBaselineConfig?: pulumi.Input<inputs.sagemaker.DataQualityJobDefinitionDataQualityBaselineConfig>;
    /**
     * A list of inputs for the monitoring job. Fields are documented below.
     */
    dataQualityJobInput: pulumi.Input<inputs.sagemaker.DataQualityJobDefinitionDataQualityJobInput>;
    /**
     * The output configuration for monitoring jobs. Fields are documented below.
     */
    dataQualityJobOutputConfig: pulumi.Input<inputs.sagemaker.DataQualityJobDefinitionDataQualityJobOutputConfig>;
    /**
     * Identifies the resources to deploy for a monitoring job. Fields are documented below.
     */
    jobResources: pulumi.Input<inputs.sagemaker.DataQualityJobDefinitionJobResources>;
    /**
     * The name of the data quality job definition. If omitted, the provider will assign a random, unique name.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies networking configuration for the monitoring job. Fields are documented below.
     */
    networkConfig?: pulumi.Input<inputs.sagemaker.DataQualityJobDefinitionNetworkConfig>;
    /**
     * The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
     */
    roleArn: pulumi.Input<string>;
    /**
     * A time limit for how long the monitoring job is allowed to run before stopping. Fields are documented below.
     */
    stoppingCondition?: pulumi.Input<inputs.sagemaker.DataQualityJobDefinitionStoppingCondition>;
    /**
     * A mapping of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
