// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Data source for managing AWS Bedrock Foundation Models.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = aws.bedrockfoundation.getModels({});
 * ```
 * ### Filter by Inference Type
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = aws.bedrockfoundation.getModels({
 *     byInferenceType: "ON_DEMAND",
 * });
 * ```
 */
export function getModels(args?: GetModelsArgs, opts?: pulumi.InvokeOptions): Promise<GetModelsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws:bedrockfoundation/getModels:getModels", {
        "byCustomizationType": args.byCustomizationType,
        "byInferenceType": args.byInferenceType,
        "byOutputModality": args.byOutputModality,
        "byProvider": args.byProvider,
        "modelSummaries": args.modelSummaries,
    }, opts);
}

/**
 * A collection of arguments for invoking getModels.
 */
export interface GetModelsArgs {
    /**
     * Customization type to filter on. Valid values are `FINE_TUNING`.
     */
    byCustomizationType?: string;
    /**
     * Inference type to filter on. Valid values are `ON_DEMAND` and `PROVISIONED`.
     */
    byInferenceType?: string;
    /**
     * Output modality to filter on. Valid values are `TEXT`, `IMAGE`, and `EMBEDDING`.
     */
    byOutputModality?: string;
    /**
     * Model provider to filter on.
     */
    byProvider?: string;
    /**
     * List of model summary objects. See `modelSummaries`.
     */
    modelSummaries?: inputs.bedrockfoundation.GetModelsModelSummary[];
}

/**
 * A collection of values returned by getModels.
 */
export interface GetModelsResult {
    readonly byCustomizationType?: string;
    readonly byInferenceType?: string;
    readonly byOutputModality?: string;
    readonly byProvider?: string;
    /**
     * AWS region.
     */
    readonly id: string;
    /**
     * List of model summary objects. See `modelSummaries`.
     */
    readonly modelSummaries?: outputs.bedrockfoundation.GetModelsModelSummary[];
}
/**
 * Data source for managing AWS Bedrock Foundation Models.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = aws.bedrockfoundation.getModels({});
 * ```
 * ### Filter by Inference Type
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = aws.bedrockfoundation.getModels({
 *     byInferenceType: "ON_DEMAND",
 * });
 * ```
 */
export function getModelsOutput(args?: GetModelsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetModelsResult> {
    return pulumi.output(args).apply((a: any) => getModels(a, opts))
}

/**
 * A collection of arguments for invoking getModels.
 */
export interface GetModelsOutputArgs {
    /**
     * Customization type to filter on. Valid values are `FINE_TUNING`.
     */
    byCustomizationType?: pulumi.Input<string>;
    /**
     * Inference type to filter on. Valid values are `ON_DEMAND` and `PROVISIONED`.
     */
    byInferenceType?: pulumi.Input<string>;
    /**
     * Output modality to filter on. Valid values are `TEXT`, `IMAGE`, and `EMBEDDING`.
     */
    byOutputModality?: pulumi.Input<string>;
    /**
     * Model provider to filter on.
     */
    byProvider?: pulumi.Input<string>;
    /**
     * List of model summary objects. See `modelSummaries`.
     */
    modelSummaries?: pulumi.Input<pulumi.Input<inputs.bedrockfoundation.GetModelsModelSummaryArgs>[]>;
}
