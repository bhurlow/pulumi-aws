// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides information about a Lambda Code Signing Config. A code signing configuration defines a list of allowed signing profiles and defines the code-signing validation policy (action to be taken if deployment validation checks fail).
 *
 * For information about Lambda code signing configurations and how to use them, see [configuring code signing for Lambda functions](https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const existingCsc = pulumi.output(aws.lambda.getCodeSigningConfig({
 *     arn: `arn:aws:lambda:${var_aws_region}:${var_aws_account}:code-signing-config:csc-0f6c334abcdea4d8b`,
 * }));
 * ```
 */
export function getCodeSigningConfig(args: GetCodeSigningConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetCodeSigningConfigResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:lambda/getCodeSigningConfig:getCodeSigningConfig", {
        "arn": args.arn,
    }, opts);
}

/**
 * A collection of arguments for invoking getCodeSigningConfig.
 */
export interface GetCodeSigningConfigArgs {
    /**
     * The Amazon Resource Name (ARN) of the code signing configuration.
     */
    arn: string;
}

/**
 * A collection of values returned by getCodeSigningConfig.
 */
export interface GetCodeSigningConfigResult {
    /**
     * List of allowed publishers as signing profiles for this code signing configuration.
     */
    readonly allowedPublishers: outputs.lambda.GetCodeSigningConfigAllowedPublisher[];
    readonly arn: string;
    /**
     * Unique identifier for the code signing configuration.
     */
    readonly configId: string;
    /**
     * Code signing configuration description.
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The date and time that the code signing configuration was last modified.
     */
    readonly lastModified: string;
    /**
     * List of code signing policies that control the validation failure action for signature mismatch or expiry.
     */
    readonly policies: outputs.lambda.GetCodeSigningConfigPolicy[];
}
