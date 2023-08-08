// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Data source for managing an AWS SFN (Step Functions) State Machine Versions.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = aws.sfn.getStateMachineVersions({
 *     statemachineArn: aws_sfn_state_machine.test.arn,
 * });
 * ```
 */
export function getStateMachineVersions(args: GetStateMachineVersionsArgs, opts?: pulumi.InvokeOptions): Promise<GetStateMachineVersionsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws:sfn/getStateMachineVersions:getStateMachineVersions", {
        "statemachineArn": args.statemachineArn,
    }, opts);
}

/**
 * A collection of arguments for invoking getStateMachineVersions.
 */
export interface GetStateMachineVersionsArgs {
    /**
     * ARN of the State Machine.
     */
    statemachineArn: string;
}

/**
 * A collection of values returned by getStateMachineVersions.
 */
export interface GetStateMachineVersionsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly statemachineArn: string;
    /**
     * ARN List identifying the statemachine versions.
     */
    readonly statemachineVersions: string[];
}
/**
 * Data source for managing an AWS SFN (Step Functions) State Machine Versions.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = aws.sfn.getStateMachineVersions({
 *     statemachineArn: aws_sfn_state_machine.test.arn,
 * });
 * ```
 */
export function getStateMachineVersionsOutput(args: GetStateMachineVersionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStateMachineVersionsResult> {
    return pulumi.output(args).apply((a: any) => getStateMachineVersions(a, opts))
}

/**
 * A collection of arguments for invoking getStateMachineVersions.
 */
export interface GetStateMachineVersionsOutputArgs {
    /**
     * ARN of the State Machine.
     */
    statemachineArn: pulumi.Input<string>;
}
