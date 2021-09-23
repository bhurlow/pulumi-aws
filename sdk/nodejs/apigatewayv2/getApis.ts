// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides details about multiple Amazon API Gateway Version 2 APIs.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.apigatewayv2.getApis({
 *     protocolType: "HTTP",
 * }));
 * ```
 */
export function getApis(args?: GetApisArgs, opts?: pulumi.InvokeOptions): Promise<GetApisResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:apigatewayv2/getApis:getApis", {
        "name": args.name,
        "protocolType": args.protocolType,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getApis.
 */
export interface GetApisArgs {
    /**
     * The API name.
     */
    name?: string;
    /**
     * The API protocol.
     */
    protocolType?: string;
    /**
     * A map of tags, each pair of which must exactly match
     * a pair on the desired APIs.
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getApis.
 */
export interface GetApisResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Set of API identifiers.
     */
    readonly ids: string[];
    readonly name?: string;
    readonly protocolType?: string;
    readonly tags?: {[key: string]: string};
}
