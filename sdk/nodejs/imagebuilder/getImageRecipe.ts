// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides details about an Image Builder Image Recipe.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.imagebuilder.getImageRecipe({
 *     arn: "arn:aws:imagebuilder:us-east-1:aws:image-recipe/example/1.0.0",
 * }));
 * ```
 */
export function getImageRecipe(args: GetImageRecipeArgs, opts?: pulumi.InvokeOptions): Promise<GetImageRecipeResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:imagebuilder/getImageRecipe:getImageRecipe", {
        "arn": args.arn,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getImageRecipe.
 */
export interface GetImageRecipeArgs {
    /**
     * Amazon Resource Name (ARN) of the image recipe.
     */
    arn: string;
    /**
     * Key-value map of resource tags for the image recipe.
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getImageRecipe.
 */
export interface GetImageRecipeResult {
    readonly arn: string;
    /**
     * Set of objects with block device mappings for the the image recipe.
     */
    readonly blockDeviceMappings: outputs.imagebuilder.GetImageRecipeBlockDeviceMapping[];
    /**
     * List of objects with components for the image recipe.
     */
    readonly components: outputs.imagebuilder.GetImageRecipeComponent[];
    /**
     * Date the image recipe was created.
     */
    readonly dateCreated: string;
    /**
     * Description of the image recipe.
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Name of the image recipe.
     */
    readonly name: string;
    /**
     * Owner of the image recipe.
     */
    readonly owner: string;
    /**
     * Platform of the image recipe.
     */
    readonly parentImage: string;
    /**
     * Platform of the image recipe.
     */
    readonly platform: string;
    /**
     * Key-value map of resource tags for the image recipe.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Version of the image recipe.
     */
    readonly version: string;
    /**
     * The working directory used during build and test workflows.
     */
    readonly workingDirectory: string;
}
