// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides information about a Lambda function URL.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const config = new pulumi.Config();
 * const functionName = config.require("functionName");
 * const existing = aws.lambda.getFunctionUrl({
 *     functionName: functionName,
 * });
 * ```
 */
export function getFunctionUrl(args: GetFunctionUrlArgs, opts?: pulumi.InvokeOptions): Promise<GetFunctionUrlResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws:lambda/getFunctionUrl:getFunctionUrl", {
        "functionName": args.functionName,
        "qualifier": args.qualifier,
    }, opts);
}

/**
 * A collection of arguments for invoking getFunctionUrl.
 */
export interface GetFunctionUrlArgs {
    /**
     * he name (or ARN) of the Lambda function.
     */
    functionName: string;
    /**
     * Alias name or `"$LATEST"`.
     */
    qualifier?: string;
}

/**
 * A collection of values returned by getFunctionUrl.
 */
export interface GetFunctionUrlResult {
    /**
     * Type of authentication that the function URL uses.
     */
    readonly authorizationType: string;
    /**
     * The [cross-origin resource sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) settings for the function URL. See the `aws.lambda.FunctionUrl` resource documentation for more details.
     */
    readonly cors: outputs.lambda.GetFunctionUrlCor[];
    /**
     * When the function URL was created, in [ISO-8601 format](https://www.w3.org/TR/NOTE-datetime).
     */
    readonly creationTime: string;
    /**
     * ARN of the function.
     */
    readonly functionArn: string;
    readonly functionName: string;
    /**
     * HTTP URL endpoint for the function in the format `https://<url_id>.lambda-url.<region>.on.aws`.
     */
    readonly functionUrl: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * When the function URL configuration was last updated, in [ISO-8601 format](https://www.w3.org/TR/NOTE-datetime).
     */
    readonly lastModifiedTime: string;
    readonly qualifier?: string;
    /**
     * Generated ID for the endpoint.
     */
    readonly urlId: string;
}

export function getFunctionUrlOutput(args: GetFunctionUrlOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFunctionUrlResult> {
    return pulumi.output(args).apply(a => getFunctionUrl(a, opts))
}

/**
 * A collection of arguments for invoking getFunctionUrl.
 */
export interface GetFunctionUrlOutputArgs {
    /**
     * he name (or ARN) of the Lambda function.
     */
    functionName: pulumi.Input<string>;
    /**
     * Alias name or `"$LATEST"`.
     */
    qualifier?: pulumi.Input<string>;
}
