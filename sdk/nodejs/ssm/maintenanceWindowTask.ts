// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides an SSM Maintenance Window Task resource
 *
 * ## Example Usage
 * ### Automation Tasks
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.ssm.MaintenanceWindowTask("example", {
 *     maxConcurrency: "2",
 *     maxErrors: "1",
 *     priority: 1,
 *     taskArn: "AWS-RestartEC2Instance",
 *     taskType: "AUTOMATION",
 *     windowId: aws_ssm_maintenance_window.example.id,
 *     targets: [{
 *         key: "InstanceIds",
 *         values: [aws_instance.example.id],
 *     }],
 *     taskInvocationParameters: {
 *         automationParameters: {
 *             documentVersion: "$LATEST",
 *             parameters: [{
 *                 name: "InstanceId",
 *                 values: [aws_instance.example.id],
 *             }],
 *         },
 *     },
 * });
 * ```
 * ### Run Command Tasks
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.ssm.MaintenanceWindowTask("example", {
 *     maxConcurrency: "2",
 *     maxErrors: "1",
 *     priority: 1,
 *     taskArn: "AWS-RunShellScript",
 *     taskType: "RUN_COMMAND",
 *     windowId: aws_ssm_maintenance_window.example.id,
 *     targets: [{
 *         key: "InstanceIds",
 *         values: [aws_instance.example.id],
 *     }],
 *     taskInvocationParameters: {
 *         runCommandParameters: {
 *             outputS3Bucket: aws_s3_bucket.example.id,
 *             outputS3KeyPrefix: "output",
 *             serviceRoleArn: aws_iam_role.example.arn,
 *             timeoutSeconds: 600,
 *             notificationConfig: {
 *                 notificationArn: aws_sns_topic.example.arn,
 *                 notificationEvents: ["All"],
 *                 notificationType: "Command",
 *             },
 *             parameters: [{
 *                 name: "commands",
 *                 values: ["date"],
 *             }],
 *         },
 *     },
 * });
 * ```
 * ### Step Function Tasks
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.ssm.MaintenanceWindowTask("example", {
 *     maxConcurrency: "2",
 *     maxErrors: "1",
 *     priority: 1,
 *     taskArn: aws_sfn_activity.example.id,
 *     taskType: "STEP_FUNCTIONS",
 *     windowId: aws_ssm_maintenance_window.example.id,
 *     targets: [{
 *         key: "InstanceIds",
 *         values: [aws_instance.example.id],
 *     }],
 *     taskInvocationParameters: {
 *         stepFunctionsParameters: {
 *             input: "{\"key1\":\"value1\"}",
 *             name: "example",
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_ssm_maintenance_window_task.task
 *
 *  id = "<window_id>/<window_task_id>" } Using `pulumi import`, import AWS Maintenance Window Task using the `window_id` and `window_task_id` separated by `/`. For exampleconsole % pulumi import aws_ssm_maintenance_window_task.task <window_id>/<window_task_id>
 */
export class MaintenanceWindowTask extends pulumi.CustomResource {
    /**
     * Get an existing MaintenanceWindowTask resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MaintenanceWindowTaskState, opts?: pulumi.CustomResourceOptions): MaintenanceWindowTask {
        return new MaintenanceWindowTask(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ssm/maintenanceWindowTask:MaintenanceWindowTask';

    /**
     * Returns true if the given object is an instance of MaintenanceWindowTask.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MaintenanceWindowTask {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MaintenanceWindowTask.__pulumiType;
    }

    /**
     * The ARN of the maintenance window task.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Indicates whether tasks should continue to run after the cutoff time specified in the maintenance windows is reached. Valid values are `CONTINUE_TASK` and `CANCEL_TASK`.
     */
    public readonly cutoffBehavior!: pulumi.Output<string | undefined>;
    /**
     * The description of the maintenance window task.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The maximum number of targets this task can be run for in parallel.
     */
    public readonly maxConcurrency!: pulumi.Output<string>;
    /**
     * The maximum number of errors allowed before this task stops being scheduled.
     */
    public readonly maxErrors!: pulumi.Output<string>;
    /**
     * The name of the maintenance window task.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * The role that should be assumed when executing the task. If a role is not provided, Systems Manager uses your account's service-linked role. If no service-linked role for Systems Manager exists in your account, it is created for you.
     */
    public readonly serviceRoleArn!: pulumi.Output<string>;
    /**
     * The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
     */
    public readonly targets!: pulumi.Output<outputs.ssm.MaintenanceWindowTaskTarget[] | undefined>;
    /**
     * The ARN of the task to execute.
     */
    public readonly taskArn!: pulumi.Output<string>;
    /**
     * Configuration block with parameters for task execution.
     */
    public readonly taskInvocationParameters!: pulumi.Output<outputs.ssm.MaintenanceWindowTaskTaskInvocationParameters | undefined>;
    /**
     * The type of task being registered. Valid values: `AUTOMATION`, `LAMBDA`, `RUN_COMMAND` or `STEP_FUNCTIONS`.
     */
    public readonly taskType!: pulumi.Output<string>;
    /**
     * The Id of the maintenance window to register the task with.
     */
    public readonly windowId!: pulumi.Output<string>;
    /**
     * The ID of the maintenance window task.
     */
    public /*out*/ readonly windowTaskId!: pulumi.Output<string>;

    /**
     * Create a MaintenanceWindowTask resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MaintenanceWindowTaskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MaintenanceWindowTaskArgs | MaintenanceWindowTaskState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MaintenanceWindowTaskState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["cutoffBehavior"] = state ? state.cutoffBehavior : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["maxConcurrency"] = state ? state.maxConcurrency : undefined;
            resourceInputs["maxErrors"] = state ? state.maxErrors : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["serviceRoleArn"] = state ? state.serviceRoleArn : undefined;
            resourceInputs["targets"] = state ? state.targets : undefined;
            resourceInputs["taskArn"] = state ? state.taskArn : undefined;
            resourceInputs["taskInvocationParameters"] = state ? state.taskInvocationParameters : undefined;
            resourceInputs["taskType"] = state ? state.taskType : undefined;
            resourceInputs["windowId"] = state ? state.windowId : undefined;
            resourceInputs["windowTaskId"] = state ? state.windowTaskId : undefined;
        } else {
            const args = argsOrState as MaintenanceWindowTaskArgs | undefined;
            if ((!args || args.taskArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'taskArn'");
            }
            if ((!args || args.taskType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'taskType'");
            }
            if ((!args || args.windowId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'windowId'");
            }
            resourceInputs["cutoffBehavior"] = args ? args.cutoffBehavior : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["maxConcurrency"] = args ? args.maxConcurrency : undefined;
            resourceInputs["maxErrors"] = args ? args.maxErrors : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["serviceRoleArn"] = args ? args.serviceRoleArn : undefined;
            resourceInputs["targets"] = args ? args.targets : undefined;
            resourceInputs["taskArn"] = args ? args.taskArn : undefined;
            resourceInputs["taskInvocationParameters"] = args ? args.taskInvocationParameters : undefined;
            resourceInputs["taskType"] = args ? args.taskType : undefined;
            resourceInputs["windowId"] = args ? args.windowId : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["windowTaskId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MaintenanceWindowTask.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MaintenanceWindowTask resources.
 */
export interface MaintenanceWindowTaskState {
    /**
     * The ARN of the maintenance window task.
     */
    arn?: pulumi.Input<string>;
    /**
     * Indicates whether tasks should continue to run after the cutoff time specified in the maintenance windows is reached. Valid values are `CONTINUE_TASK` and `CANCEL_TASK`.
     */
    cutoffBehavior?: pulumi.Input<string>;
    /**
     * The description of the maintenance window task.
     */
    description?: pulumi.Input<string>;
    /**
     * The maximum number of targets this task can be run for in parallel.
     */
    maxConcurrency?: pulumi.Input<string>;
    /**
     * The maximum number of errors allowed before this task stops being scheduled.
     */
    maxErrors?: pulumi.Input<string>;
    /**
     * The name of the maintenance window task.
     */
    name?: pulumi.Input<string>;
    /**
     * The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
     */
    priority?: pulumi.Input<number>;
    /**
     * The role that should be assumed when executing the task. If a role is not provided, Systems Manager uses your account's service-linked role. If no service-linked role for Systems Manager exists in your account, it is created for you.
     */
    serviceRoleArn?: pulumi.Input<string>;
    /**
     * The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
     */
    targets?: pulumi.Input<pulumi.Input<inputs.ssm.MaintenanceWindowTaskTarget>[]>;
    /**
     * The ARN of the task to execute.
     */
    taskArn?: pulumi.Input<string>;
    /**
     * Configuration block with parameters for task execution.
     */
    taskInvocationParameters?: pulumi.Input<inputs.ssm.MaintenanceWindowTaskTaskInvocationParameters>;
    /**
     * The type of task being registered. Valid values: `AUTOMATION`, `LAMBDA`, `RUN_COMMAND` or `STEP_FUNCTIONS`.
     */
    taskType?: pulumi.Input<string>;
    /**
     * The Id of the maintenance window to register the task with.
     */
    windowId?: pulumi.Input<string>;
    /**
     * The ID of the maintenance window task.
     */
    windowTaskId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MaintenanceWindowTask resource.
 */
export interface MaintenanceWindowTaskArgs {
    /**
     * Indicates whether tasks should continue to run after the cutoff time specified in the maintenance windows is reached. Valid values are `CONTINUE_TASK` and `CANCEL_TASK`.
     */
    cutoffBehavior?: pulumi.Input<string>;
    /**
     * The description of the maintenance window task.
     */
    description?: pulumi.Input<string>;
    /**
     * The maximum number of targets this task can be run for in parallel.
     */
    maxConcurrency?: pulumi.Input<string>;
    /**
     * The maximum number of errors allowed before this task stops being scheduled.
     */
    maxErrors?: pulumi.Input<string>;
    /**
     * The name of the maintenance window task.
     */
    name?: pulumi.Input<string>;
    /**
     * The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
     */
    priority?: pulumi.Input<number>;
    /**
     * The role that should be assumed when executing the task. If a role is not provided, Systems Manager uses your account's service-linked role. If no service-linked role for Systems Manager exists in your account, it is created for you.
     */
    serviceRoleArn?: pulumi.Input<string>;
    /**
     * The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
     */
    targets?: pulumi.Input<pulumi.Input<inputs.ssm.MaintenanceWindowTaskTarget>[]>;
    /**
     * The ARN of the task to execute.
     */
    taskArn: pulumi.Input<string>;
    /**
     * Configuration block with parameters for task execution.
     */
    taskInvocationParameters?: pulumi.Input<inputs.ssm.MaintenanceWindowTaskTaskInvocationParameters>;
    /**
     * The type of task being registered. Valid values: `AUTOMATION`, `LAMBDA`, `RUN_COMMAND` or `STEP_FUNCTIONS`.
     */
    taskType: pulumi.Input<string>;
    /**
     * The Id of the maintenance window to register the task with.
     */
    windowId: pulumi.Input<string>;
}
