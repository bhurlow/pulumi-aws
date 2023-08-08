// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ARN of a State Machine in AWS Step
 * Function (SFN). By using this data source, you can reference a
 * state machine without having to hard code the ARNs as input.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.sfn.getStateMachine({
 *     name: "an_example_sfn_name",
 * });
 * ```
 */
export function getStateMachine(args: GetStateMachineArgs, opts?: pulumi.InvokeOptions): Promise<GetStateMachineResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws:sfn/getStateMachine:getStateMachine", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getStateMachine.
 */
export interface GetStateMachineArgs {
    /**
     * Friendly name of the state machine to match.
     */
    name: string;
}

/**
 * A collection of values returned by getStateMachine.
 */
export interface GetStateMachineResult {
    /**
     * Set to the arn of the state function.
     */
    readonly arn: string;
    /**
     * Date the state machine was created.
     */
    readonly creationDate: string;
    /**
     * Set to the state machine definition.
     */
    readonly definition: string;
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * The revision identifier for the state machine.
     */
    readonly revisionId: string;
    /**
     * Set to the roleArn used by the state function.
     */
    readonly roleArn: string;
    /**
     * Set to the current status of the state machine.
     */
    readonly status: string;
}
/**
 * Use this data source to get the ARN of a State Machine in AWS Step
 * Function (SFN). By using this data source, you can reference a
 * state machine without having to hard code the ARNs as input.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.sfn.getStateMachine({
 *     name: "an_example_sfn_name",
 * });
 * ```
 */
export function getStateMachineOutput(args: GetStateMachineOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStateMachineResult> {
    return pulumi.output(args).apply((a: any) => getStateMachine(a, opts))
}

/**
 * A collection of arguments for invoking getStateMachine.
 */
export interface GetStateMachineOutputArgs {
    /**
     * Friendly name of the state machine to match.
     */
    name: pulumi.Input<string>;
}
