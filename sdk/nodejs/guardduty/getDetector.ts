// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Retrieve information about a GuardDuty detector.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.guardduty.getDetector({ async: true }));
 * ```
 */
export function getDetector(args?: GetDetectorArgs, opts?: pulumi.InvokeOptions): Promise<GetDetectorResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:guardduty/getDetector:getDetector", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getDetector.
 */
export interface GetDetectorArgs {
    /**
     * The ID of the detector.
     */
    readonly id?: string;
}

/**
 * A collection of values returned by getDetector.
 */
export interface GetDetectorResult {
    /**
     * The frequency of notifications sent about subsequent finding occurrences.
     */
    readonly findingPublishingFrequency: string;
    readonly id: string;
    /**
     * The service-linked role that grants GuardDuty access to the resources in the AWS account.
     */
    readonly serviceRoleArn: string;
    /**
     * The current status of the detector.
     */
    readonly status: string;
}
