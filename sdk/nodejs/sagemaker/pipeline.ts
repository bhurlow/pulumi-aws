// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a SageMaker Pipeline resource.
 *
 * ## Example Usage
 * ### Basic usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.sagemaker.Pipeline("example", {
 *     pipelineName: "example",
 *     pipelineDisplayName: "example",
 *     roleArn: aws_iam_role.example.arn,
 *     pipelineDefinition: JSON.stringify({
 *         Version: "2020-12-01",
 *         Steps: [{
 *             Name: "Test",
 *             Type: "Fail",
 *             Arguments: {
 *                 ErrorMessage: "test",
 *             },
 *         }],
 *     }),
 * });
 * ```
 *
 * ## Import
 *
 * You can use `pulumi import` to import pipelines using `pipeline_name`. For exampleterraform import {
 *
 *  to = aws_sagemaker_pipeline.test_pipeline
 *
 *  id = "pipeline" } Using `pulumi import`, import pipelines using the `pipeline_name`. For exampleconsole % pulumi import aws_sagemaker_pipeline.test_pipeline pipeline
 */
export class Pipeline extends pulumi.CustomResource {
    /**
     * Get an existing Pipeline resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PipelineState, opts?: pulumi.CustomResourceOptions): Pipeline {
        return new Pipeline(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:sagemaker/pipeline:Pipeline';

    /**
     * Returns true if the given object is an instance of Pipeline.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Pipeline {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Pipeline.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this Pipeline.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * This is the configuration that controls the parallelism of the pipeline. If specified, it applies to all runs of this pipeline by default. see Parallelism Configuration details below.
     */
    public readonly parallelismConfiguration!: pulumi.Output<outputs.sagemaker.PipelineParallelismConfiguration | undefined>;
    /**
     * The [JSON pipeline definition](https://aws-sagemaker-mlops.github.io/sagemaker-model-building-pipeline-definition-JSON-schema/) of the pipeline.
     */
    public readonly pipelineDefinition!: pulumi.Output<string | undefined>;
    /**
     * The location of the pipeline definition stored in Amazon S3. If specified, SageMaker will retrieve the pipeline definition from this location. see Pipeline Definition S3 Location details below.
     */
    public readonly pipelineDefinitionS3Location!: pulumi.Output<outputs.sagemaker.PipelinePipelineDefinitionS3Location | undefined>;
    /**
     * A description of the pipeline.
     */
    public readonly pipelineDescription!: pulumi.Output<string | undefined>;
    /**
     * The display name of the pipeline.
     */
    public readonly pipelineDisplayName!: pulumi.Output<string>;
    /**
     * The name of the pipeline.
     */
    public readonly pipelineName!: pulumi.Output<string>;
    /**
     * The name of the Pipeline (must be unique).
     */
    public readonly roleArn!: pulumi.Output<string | undefined>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a Pipeline resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PipelineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PipelineArgs | PipelineState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PipelineState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["parallelismConfiguration"] = state ? state.parallelismConfiguration : undefined;
            resourceInputs["pipelineDefinition"] = state ? state.pipelineDefinition : undefined;
            resourceInputs["pipelineDefinitionS3Location"] = state ? state.pipelineDefinitionS3Location : undefined;
            resourceInputs["pipelineDescription"] = state ? state.pipelineDescription : undefined;
            resourceInputs["pipelineDisplayName"] = state ? state.pipelineDisplayName : undefined;
            resourceInputs["pipelineName"] = state ? state.pipelineName : undefined;
            resourceInputs["roleArn"] = state ? state.roleArn : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as PipelineArgs | undefined;
            if ((!args || args.pipelineDisplayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pipelineDisplayName'");
            }
            if ((!args || args.pipelineName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pipelineName'");
            }
            resourceInputs["parallelismConfiguration"] = args ? args.parallelismConfiguration : undefined;
            resourceInputs["pipelineDefinition"] = args ? args.pipelineDefinition : undefined;
            resourceInputs["pipelineDefinitionS3Location"] = args ? args.pipelineDefinitionS3Location : undefined;
            resourceInputs["pipelineDescription"] = args ? args.pipelineDescription : undefined;
            resourceInputs["pipelineDisplayName"] = args ? args.pipelineDisplayName : undefined;
            resourceInputs["pipelineName"] = args ? args.pipelineName : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Pipeline.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Pipeline resources.
 */
export interface PipelineState {
    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this Pipeline.
     */
    arn?: pulumi.Input<string>;
    /**
     * This is the configuration that controls the parallelism of the pipeline. If specified, it applies to all runs of this pipeline by default. see Parallelism Configuration details below.
     */
    parallelismConfiguration?: pulumi.Input<inputs.sagemaker.PipelineParallelismConfiguration>;
    /**
     * The [JSON pipeline definition](https://aws-sagemaker-mlops.github.io/sagemaker-model-building-pipeline-definition-JSON-schema/) of the pipeline.
     */
    pipelineDefinition?: pulumi.Input<string>;
    /**
     * The location of the pipeline definition stored in Amazon S3. If specified, SageMaker will retrieve the pipeline definition from this location. see Pipeline Definition S3 Location details below.
     */
    pipelineDefinitionS3Location?: pulumi.Input<inputs.sagemaker.PipelinePipelineDefinitionS3Location>;
    /**
     * A description of the pipeline.
     */
    pipelineDescription?: pulumi.Input<string>;
    /**
     * The display name of the pipeline.
     */
    pipelineDisplayName?: pulumi.Input<string>;
    /**
     * The name of the pipeline.
     */
    pipelineName?: pulumi.Input<string>;
    /**
     * The name of the Pipeline (must be unique).
     */
    roleArn?: pulumi.Input<string>;
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
 * The set of arguments for constructing a Pipeline resource.
 */
export interface PipelineArgs {
    /**
     * This is the configuration that controls the parallelism of the pipeline. If specified, it applies to all runs of this pipeline by default. see Parallelism Configuration details below.
     */
    parallelismConfiguration?: pulumi.Input<inputs.sagemaker.PipelineParallelismConfiguration>;
    /**
     * The [JSON pipeline definition](https://aws-sagemaker-mlops.github.io/sagemaker-model-building-pipeline-definition-JSON-schema/) of the pipeline.
     */
    pipelineDefinition?: pulumi.Input<string>;
    /**
     * The location of the pipeline definition stored in Amazon S3. If specified, SageMaker will retrieve the pipeline definition from this location. see Pipeline Definition S3 Location details below.
     */
    pipelineDefinitionS3Location?: pulumi.Input<inputs.sagemaker.PipelinePipelineDefinitionS3Location>;
    /**
     * A description of the pipeline.
     */
    pipelineDescription?: pulumi.Input<string>;
    /**
     * The display name of the pipeline.
     */
    pipelineDisplayName: pulumi.Input<string>;
    /**
     * The name of the pipeline.
     */
    pipelineName: pulumi.Input<string>;
    /**
     * The name of the Pipeline (must be unique).
     */
    roleArn?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
