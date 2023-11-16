// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Data source for managing an AWS Bedrock Foundation Model.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testModels = aws.bedrockfoundation.getModels({});
 * const testModel = testModels.then(testModels => aws.bedrockfoundation.getModel({
 *     modelId: testModels.modelSummaries?.[0]?.modelId,
 * }));
 * ```
 */
export function getModel(args: GetModelArgs, opts?: pulumi.InvokeOptions): Promise<GetModelResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws:bedrockfoundation/getModel:getModel", {
        "modelId": args.modelId,
    }, opts);
}

/**
 * A collection of arguments for invoking getModel.
 */
export interface GetModelArgs {
    /**
     * Model identifier.
     */
    modelId: string;
}

/**
 * A collection of values returned by getModel.
 */
export interface GetModelResult {
    /**
     * Customizations that the model supports.
     */
    readonly customizationsSupporteds: string[];
    readonly id: string;
    /**
     * Inference types that the model supports.
     */
    readonly inferenceTypesSupporteds: string[];
    /**
     * Input modalities that the model supports.
     */
    readonly inputModalities: string[];
    /**
     * Model ARN.
     */
    readonly modelArn: string;
    readonly modelId: string;
    /**
     * Model name.
     */
    readonly modelName: string;
    /**
     * Output modalities that the model supports.
     */
    readonly outputModalities: string[];
    /**
     * Model provider name.
     */
    readonly providerName: string;
    /**
     * Indicates whether the model supports streaming.
     */
    readonly responseStreamingSupported: boolean;
}
/**
 * Data source for managing an AWS Bedrock Foundation Model.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testModels = aws.bedrockfoundation.getModels({});
 * const testModel = testModels.then(testModels => aws.bedrockfoundation.getModel({
 *     modelId: testModels.modelSummaries?.[0]?.modelId,
 * }));
 * ```
 */
export function getModelOutput(args: GetModelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetModelResult> {
    return pulumi.output(args).apply((a: any) => getModel(a, opts))
}

/**
 * A collection of arguments for invoking getModel.
 */
export interface GetModelOutputArgs {
    /**
     * Model identifier.
     */
    modelId: pulumi.Input<string>;
}
