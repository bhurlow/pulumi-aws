// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides details about an Image Builder Image Pipeline.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.imagebuilder.getImagePipeline({
 *     arn: "arn:aws:imagebuilder:us-west-2:aws:image-pipeline/example",
 * }));
 * ```
 */
export function getImagePipeline(args: GetImagePipelineArgs, opts?: pulumi.InvokeOptions): Promise<GetImagePipelineResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:imagebuilder/getImagePipeline:getImagePipeline", {
        "arn": args.arn,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getImagePipeline.
 */
export interface GetImagePipelineArgs {
    /**
     * Amazon Resource Name (ARN) of the image pipeline.
     */
    arn: string;
    /**
     * Key-value map of resource tags for the image pipeline.
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getImagePipeline.
 */
export interface GetImagePipelineResult {
    readonly arn: string;
    /**
     * Date the image pipeline was created.
     */
    readonly dateCreated: string;
    /**
     * Date the image pipeline was last run.
     */
    readonly dateLastRun: string;
    /**
     * Date the image pipeline will run next.
     */
    readonly dateNextRun: string;
    /**
     * Date the image pipeline was updated.
     */
    readonly dateUpdated: string;
    /**
     * Description of the image pipeline.
     */
    readonly description: string;
    /**
     * Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
     */
    readonly distributionConfigurationArn: string;
    /**
     * Whether additional information about the image being created is collected.
     */
    readonly enhancedImageMetadataEnabled: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Amazon Resource Name (ARN) of the Image Builder Infrastructure Recipe.
     */
    readonly imageRecipeArn: string;
    /**
     * List of an object with image tests configuration.
     */
    readonly imageTestsConfigurations: outputs.imagebuilder.GetImagePipelineImageTestsConfiguration[];
    /**
     * Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
     */
    readonly infrastructureConfigurationArn: string;
    /**
     * Name of the image pipeline.
     */
    readonly name: string;
    /**
     * Platform of the image pipeline.
     */
    readonly platform: string;
    /**
     * List of an object with schedule settings.
     */
    readonly schedules: outputs.imagebuilder.GetImagePipelineSchedule[];
    /**
     * Status of the image pipeline.
     */
    readonly status: string;
    /**
     * Key-value map of resource tags for the image pipeline.
     */
    readonly tags: {[key: string]: string};
}
