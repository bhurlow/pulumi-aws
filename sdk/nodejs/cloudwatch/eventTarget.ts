// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides an EventBridge Target resource.
 *
 * > **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.
 *
 * ## Example Usage
 * ### Kinesis Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const console = new aws.cloudwatch.EventRule("console", {
 *     description: "Capture all EC2 scaling events",
 *     eventPattern: JSON.stringify({
 *         source: ["aws.autoscaling"],
 *         "detail-type": [
 *             "EC2 Instance Launch Successful",
 *             "EC2 Instance Terminate Successful",
 *             "EC2 Instance Launch Unsuccessful",
 *             "EC2 Instance Terminate Unsuccessful",
 *         ],
 *     }),
 * });
 * const testStream = new aws.kinesis.Stream("testStream", {shardCount: 1});
 * const yada = new aws.cloudwatch.EventTarget("yada", {
 *     rule: console.name,
 *     arn: testStream.arn,
 *     runCommandTargets: [
 *         {
 *             key: "tag:Name",
 *             values: ["FooBar"],
 *         },
 *         {
 *             key: "InstanceIds",
 *             values: ["i-162058cd308bffec2"],
 *         },
 *     ],
 * });
 * ```
 * ### SSM Document Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const ssmLifecycleTrust = aws.iam.getPolicyDocument({
 *     statements: [{
 *         actions: ["sts:AssumeRole"],
 *         principals: [{
 *             type: "Service",
 *             identifiers: ["events.amazonaws.com"],
 *         }],
 *     }],
 * });
 * const stopInstance = new aws.ssm.Document("stopInstance", {
 *     documentType: "Command",
 *     content: JSON.stringify({
 *         schemaVersion: "1.2",
 *         description: "Stop an instance",
 *         parameters: {},
 *         runtimeConfig: {
 *             "aws:runShellScript": {
 *                 properties: [{
 *                     id: "0.aws:runShellScript",
 *                     runCommand: ["halt"],
 *                 }],
 *             },
 *         },
 *     }),
 * });
 * const ssmLifecyclePolicyDocument = aws.iam.getPolicyDocumentOutput({
 *     statements: [
 *         {
 *             effect: "Allow",
 *             actions: ["ssm:SendCommand"],
 *             resources: ["arn:aws:ec2:eu-west-1:1234567890:instance/*"],
 *             conditions: [{
 *                 test: "StringEquals",
 *                 variable: "ec2:ResourceTag/Terminate",
 *                 values: ["*"],
 *             }],
 *         },
 *         {
 *             effect: "Allow",
 *             actions: ["ssm:SendCommand"],
 *             resources: [stopInstance.arn],
 *         },
 *     ],
 * });
 * const ssmLifecycleRole = new aws.iam.Role("ssmLifecycleRole", {assumeRolePolicy: ssmLifecycleTrust.then(ssmLifecycleTrust => ssmLifecycleTrust.json)});
 * const ssmLifecyclePolicy = new aws.iam.Policy("ssmLifecyclePolicy", {policy: ssmLifecyclePolicyDocument.apply(ssmLifecyclePolicyDocument => ssmLifecyclePolicyDocument.json)});
 * const ssmLifecycleRolePolicyAttachment = new aws.iam.RolePolicyAttachment("ssmLifecycleRolePolicyAttachment", {
 *     policyArn: ssmLifecyclePolicy.arn,
 *     role: ssmLifecycleRole.name,
 * });
 * const stopInstancesEventRule = new aws.cloudwatch.EventRule("stopInstancesEventRule", {
 *     description: "Stop instances nightly",
 *     scheduleExpression: "cron(0 0 * * ? *)",
 * });
 * const stopInstancesEventTarget = new aws.cloudwatch.EventTarget("stopInstancesEventTarget", {
 *     arn: stopInstance.arn,
 *     rule: stopInstancesEventRule.name,
 *     roleArn: ssmLifecycleRole.arn,
 *     runCommandTargets: [{
 *         key: "tag:Terminate",
 *         values: ["midnight"],
 *     }],
 * });
 * ```
 * ### RunCommand Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const stopInstancesEventRule = new aws.cloudwatch.EventRule("stopInstancesEventRule", {
 *     description: "Stop instances nightly",
 *     scheduleExpression: "cron(0 0 * * ? *)",
 * });
 * const stopInstancesEventTarget = new aws.cloudwatch.EventTarget("stopInstancesEventTarget", {
 *     arn: `arn:aws:ssm:${_var.aws_region}::document/AWS-RunShellScript`,
 *     input: "{\"commands\":[\"halt\"]}",
 *     rule: stopInstancesEventRule.name,
 *     roleArn: aws_iam_role.ssm_lifecycle.arn,
 *     runCommandTargets: [{
 *         key: "tag:Terminate",
 *         values: ["midnight"],
 *     }],
 * });
 * ```
 * ### API Gateway target
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleEventRule = new aws.cloudwatch.EventRule("exampleEventRule", {});
 * // ...
 * const exampleDeployment = new aws.apigateway.Deployment("exampleDeployment", {restApi: aws_api_gateway_rest_api.example.id});
 * // ...
 * const exampleStage = new aws.apigateway.Stage("exampleStage", {
 *     restApi: aws_api_gateway_rest_api.example.id,
 *     deployment: exampleDeployment.id,
 * });
 * // ...
 * const exampleEventTarget = new aws.cloudwatch.EventTarget("exampleEventTarget", {
 *     arn: pulumi.interpolate`${exampleStage.executionArn}/GET`,
 *     rule: exampleEventRule.id,
 *     httpTarget: {
 *         queryStringParameters: {
 *             Body: "$.detail.body",
 *         },
 *         headerParameters: {
 *             Env: "Test",
 *         },
 *     },
 * });
 * ```
 * ### Cross-Account Event Bus target
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const assumeRole = aws.iam.getPolicyDocument({
 *     statements: [{
 *         effect: "Allow",
 *         principals: [{
 *             type: "Service",
 *             identifiers: ["events.amazonaws.com"],
 *         }],
 *         actions: ["sts:AssumeRole"],
 *     }],
 * });
 * const eventBusInvokeRemoteEventBusRole = new aws.iam.Role("eventBusInvokeRemoteEventBusRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
 * const eventBusInvokeRemoteEventBusPolicyDocument = aws.iam.getPolicyDocument({
 *     statements: [{
 *         effect: "Allow",
 *         actions: ["events:PutEvents"],
 *         resources: ["arn:aws:events:eu-west-1:1234567890:event-bus/My-Event-Bus"],
 *     }],
 * });
 * const eventBusInvokeRemoteEventBusPolicy = new aws.iam.Policy("eventBusInvokeRemoteEventBusPolicy", {policy: eventBusInvokeRemoteEventBusPolicyDocument.then(eventBusInvokeRemoteEventBusPolicyDocument => eventBusInvokeRemoteEventBusPolicyDocument.json)});
 * const eventBusInvokeRemoteEventBusRolePolicyAttachment = new aws.iam.RolePolicyAttachment("eventBusInvokeRemoteEventBusRolePolicyAttachment", {
 *     role: eventBusInvokeRemoteEventBusRole.name,
 *     policyArn: eventBusInvokeRemoteEventBusPolicy.arn,
 * });
 * const stopInstancesEventRule = new aws.cloudwatch.EventRule("stopInstancesEventRule", {
 *     description: "Stop instances nightly",
 *     scheduleExpression: "cron(0 0 * * ? *)",
 * });
 * const stopInstancesEventTarget = new aws.cloudwatch.EventTarget("stopInstancesEventTarget", {
 *     arn: "arn:aws:events:eu-west-1:1234567890:event-bus/My-Event-Bus",
 *     rule: stopInstancesEventRule.name,
 *     roleArn: eventBusInvokeRemoteEventBusRole.arn,
 * });
 * ```
 * ### Input Transformer Usage - JSON Object
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleEventRule = new aws.cloudwatch.EventRule("exampleEventRule", {});
 * // ...
 * const exampleEventTarget = new aws.cloudwatch.EventTarget("exampleEventTarget", {
 *     arn: aws_lambda_function.example.arn,
 *     rule: exampleEventRule.id,
 *     inputTransformer: {
 *         inputPaths: {
 *             instance: "$.detail.instance",
 *             status: "$.detail.status",
 *         },
 *         inputTemplate: `{
 *   "instance_id": <instance>,
 *   "instance_status": <status>
 * }
 * `,
 *     },
 * });
 * ```
 * ### Input Transformer Usage - Simple String
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleEventRule = new aws.cloudwatch.EventRule("exampleEventRule", {});
 * // ...
 * const exampleEventTarget = new aws.cloudwatch.EventTarget("exampleEventTarget", {
 *     arn: aws_lambda_function.example.arn,
 *     rule: exampleEventRule.id,
 *     inputTransformer: {
 *         inputPaths: {
 *             instance: "$.detail.instance",
 *             status: "$.detail.status",
 *         },
 *         inputTemplate: "\"<instance> is in state <status>\"",
 *     },
 * });
 * ```
 * ### Cloudwatch Log Group Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleLogGroup = new aws.cloudwatch.LogGroup("exampleLogGroup", {retentionInDays: 1});
 * const exampleEventRule = new aws.cloudwatch.EventRule("exampleEventRule", {
 *     description: "GuardDuty Findings",
 *     eventPattern: JSON.stringify({
 *         source: ["aws.guardduty"],
 *     }),
 *     tags: {
 *         Environment: "example",
 *     },
 * });
 * const exampleLogPolicy = aws.iam.getPolicyDocumentOutput({
 *     statements: [
 *         {
 *             effect: "Allow",
 *             actions: ["logs:CreateLogStream"],
 *             resources: [pulumi.interpolate`${exampleLogGroup.arn}:*`],
 *             principals: [{
 *                 type: "Service",
 *                 identifiers: [
 *                     "events.amazonaws.com",
 *                     "delivery.logs.amazonaws.com",
 *                 ],
 *             }],
 *         },
 *         {
 *             effect: "Allow",
 *             actions: ["logs:PutLogEvents"],
 *             resources: [pulumi.interpolate`${exampleLogGroup.arn}:*:*`],
 *             principals: [{
 *                 type: "Service",
 *                 identifiers: [
 *                     "events.amazonaws.com",
 *                     "delivery.logs.amazonaws.com",
 *                 ],
 *             }],
 *             conditions: [{
 *                 test: "ArnEquals",
 *                 values: [exampleEventRule.arn],
 *                 variable: "aws:SourceArn",
 *             }],
 *         },
 *     ],
 * });
 * const exampleLogResourcePolicy = new aws.cloudwatch.LogResourcePolicy("exampleLogResourcePolicy", {
 *     policyDocument: exampleLogPolicy.apply(exampleLogPolicy => exampleLogPolicy.json),
 *     policyName: "guardduty-log-publishing-policy",
 * });
 * const exampleEventTarget = new aws.cloudwatch.EventTarget("exampleEventTarget", {
 *     rule: exampleEventRule.name,
 *     arn: exampleLogGroup.arn,
 * });
 * ```
 *
 * ## Import
 *
 *  terraform import {
 *
 *  to = aws_cloudwatch_event_target.test-event-target
 *
 *  id = "rule-name/target-id" } Using `pulumi import`, import EventBridge Targets using `event_bus_name/rule-name/target-id` (if you omit `event_bus_name`, the `default` event bus will be used). For example:
 *
 * console % pulumi import aws_cloudwatch_event_target.test-event-target rule-name/target-id
 */
export class EventTarget extends pulumi.CustomResource {
    /**
     * Get an existing EventTarget resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventTargetState, opts?: pulumi.CustomResourceOptions): EventTarget {
        return new EventTarget(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloudwatch/eventTarget:EventTarget';

    /**
     * Returns true if the given object is an instance of EventTarget.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventTarget {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventTarget.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the target.
     */
    public readonly arn!: pulumi.Output<string>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
     */
    public readonly batchTarget!: pulumi.Output<outputs.cloudwatch.EventTargetBatchTarget | undefined>;
    /**
     * Parameters used when you are providing a dead letter config. Documented below. A maximum of 1 are allowed.
     */
    public readonly deadLetterConfig!: pulumi.Output<outputs.cloudwatch.EventTargetDeadLetterConfig | undefined>;
    /**
     * Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
     */
    public readonly ecsTarget!: pulumi.Output<outputs.cloudwatch.EventTargetEcsTarget | undefined>;
    /**
     * The name or ARN of the event bus to associate with the rule.
     * If you omit this, the `default` event bus is used.
     */
    public readonly eventBusName!: pulumi.Output<string | undefined>;
    /**
     * Parameters used when you are using the rule to invoke an API Gateway REST endpoint. Documented below. A maximum of 1 is allowed.
     */
    public readonly httpTarget!: pulumi.Output<outputs.cloudwatch.EventTargetHttpTarget | undefined>;
    /**
     * Valid JSON text passed to the target. Conflicts with `inputPath` and `inputTransformer`.
     */
    public readonly input!: pulumi.Output<string | undefined>;
    /**
     * The value of the [JSONPath](http://goessner.net/articles/JsonPath/) that is used for extracting part of the matched event when passing it to the target. Conflicts with `input` and `inputTransformer`.
     */
    public readonly inputPath!: pulumi.Output<string | undefined>;
    /**
     * Parameters used when you are providing a custom input to a target based on certain event data. Conflicts with `input` and `inputPath`.
     */
    public readonly inputTransformer!: pulumi.Output<outputs.cloudwatch.EventTargetInputTransformer | undefined>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
     */
    public readonly kinesisTarget!: pulumi.Output<outputs.cloudwatch.EventTargetKinesisTarget | undefined>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon Redshift Statement. Documented below. A maximum of 1 are allowed.
     */
    public readonly redshiftTarget!: pulumi.Output<outputs.cloudwatch.EventTargetRedshiftTarget | undefined>;
    /**
     * Parameters used when you are providing retry policies. Documented below. A maximum of 1 are allowed.
     */
    public readonly retryPolicy!: pulumi.Output<outputs.cloudwatch.EventTargetRetryPolicy | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecsTarget` is used or target in `arn` is EC2 instance, Kinesis data stream, Step Functions state machine, or Event Bus in different account or region.
     */
    public readonly roleArn!: pulumi.Output<string | undefined>;
    /**
     * The name of the rule you want to add targets to.
     *
     * The following arguments are optional:
     */
    public readonly rule!: pulumi.Output<string>;
    /**
     * Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
     */
    public readonly runCommandTargets!: pulumi.Output<outputs.cloudwatch.EventTargetRunCommandTarget[] | undefined>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
     */
    public readonly sqsTarget!: pulumi.Output<outputs.cloudwatch.EventTargetSqsTarget | undefined>;
    /**
     * The unique target assignment ID. If missing, will generate a random, unique id.
     */
    public readonly targetId!: pulumi.Output<string>;

    /**
     * Create a EventTarget resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventTargetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventTargetArgs | EventTargetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EventTargetState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["batchTarget"] = state ? state.batchTarget : undefined;
            resourceInputs["deadLetterConfig"] = state ? state.deadLetterConfig : undefined;
            resourceInputs["ecsTarget"] = state ? state.ecsTarget : undefined;
            resourceInputs["eventBusName"] = state ? state.eventBusName : undefined;
            resourceInputs["httpTarget"] = state ? state.httpTarget : undefined;
            resourceInputs["input"] = state ? state.input : undefined;
            resourceInputs["inputPath"] = state ? state.inputPath : undefined;
            resourceInputs["inputTransformer"] = state ? state.inputTransformer : undefined;
            resourceInputs["kinesisTarget"] = state ? state.kinesisTarget : undefined;
            resourceInputs["redshiftTarget"] = state ? state.redshiftTarget : undefined;
            resourceInputs["retryPolicy"] = state ? state.retryPolicy : undefined;
            resourceInputs["roleArn"] = state ? state.roleArn : undefined;
            resourceInputs["rule"] = state ? state.rule : undefined;
            resourceInputs["runCommandTargets"] = state ? state.runCommandTargets : undefined;
            resourceInputs["sqsTarget"] = state ? state.sqsTarget : undefined;
            resourceInputs["targetId"] = state ? state.targetId : undefined;
        } else {
            const args = argsOrState as EventTargetArgs | undefined;
            if ((!args || args.arn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'arn'");
            }
            if ((!args || args.rule === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rule'");
            }
            resourceInputs["arn"] = args ? args.arn : undefined;
            resourceInputs["batchTarget"] = args ? args.batchTarget : undefined;
            resourceInputs["deadLetterConfig"] = args ? args.deadLetterConfig : undefined;
            resourceInputs["ecsTarget"] = args ? args.ecsTarget : undefined;
            resourceInputs["eventBusName"] = args ? args.eventBusName : undefined;
            resourceInputs["httpTarget"] = args ? args.httpTarget : undefined;
            resourceInputs["input"] = args ? args.input : undefined;
            resourceInputs["inputPath"] = args ? args.inputPath : undefined;
            resourceInputs["inputTransformer"] = args ? args.inputTransformer : undefined;
            resourceInputs["kinesisTarget"] = args ? args.kinesisTarget : undefined;
            resourceInputs["redshiftTarget"] = args ? args.redshiftTarget : undefined;
            resourceInputs["retryPolicy"] = args ? args.retryPolicy : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["rule"] = args ? args.rule : undefined;
            resourceInputs["runCommandTargets"] = args ? args.runCommandTargets : undefined;
            resourceInputs["sqsTarget"] = args ? args.sqsTarget : undefined;
            resourceInputs["targetId"] = args ? args.targetId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EventTarget.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventTarget resources.
 */
export interface EventTargetState {
    /**
     * The Amazon Resource Name (ARN) of the target.
     */
    arn?: pulumi.Input<string>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
     */
    batchTarget?: pulumi.Input<inputs.cloudwatch.EventTargetBatchTarget>;
    /**
     * Parameters used when you are providing a dead letter config. Documented below. A maximum of 1 are allowed.
     */
    deadLetterConfig?: pulumi.Input<inputs.cloudwatch.EventTargetDeadLetterConfig>;
    /**
     * Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
     */
    ecsTarget?: pulumi.Input<inputs.cloudwatch.EventTargetEcsTarget>;
    /**
     * The name or ARN of the event bus to associate with the rule.
     * If you omit this, the `default` event bus is used.
     */
    eventBusName?: pulumi.Input<string>;
    /**
     * Parameters used when you are using the rule to invoke an API Gateway REST endpoint. Documented below. A maximum of 1 is allowed.
     */
    httpTarget?: pulumi.Input<inputs.cloudwatch.EventTargetHttpTarget>;
    /**
     * Valid JSON text passed to the target. Conflicts with `inputPath` and `inputTransformer`.
     */
    input?: pulumi.Input<string>;
    /**
     * The value of the [JSONPath](http://goessner.net/articles/JsonPath/) that is used for extracting part of the matched event when passing it to the target. Conflicts with `input` and `inputTransformer`.
     */
    inputPath?: pulumi.Input<string>;
    /**
     * Parameters used when you are providing a custom input to a target based on certain event data. Conflicts with `input` and `inputPath`.
     */
    inputTransformer?: pulumi.Input<inputs.cloudwatch.EventTargetInputTransformer>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
     */
    kinesisTarget?: pulumi.Input<inputs.cloudwatch.EventTargetKinesisTarget>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon Redshift Statement. Documented below. A maximum of 1 are allowed.
     */
    redshiftTarget?: pulumi.Input<inputs.cloudwatch.EventTargetRedshiftTarget>;
    /**
     * Parameters used when you are providing retry policies. Documented below. A maximum of 1 are allowed.
     */
    retryPolicy?: pulumi.Input<inputs.cloudwatch.EventTargetRetryPolicy>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecsTarget` is used or target in `arn` is EC2 instance, Kinesis data stream, Step Functions state machine, or Event Bus in different account or region.
     */
    roleArn?: pulumi.Input<string>;
    /**
     * The name of the rule you want to add targets to.
     *
     * The following arguments are optional:
     */
    rule?: pulumi.Input<string>;
    /**
     * Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
     */
    runCommandTargets?: pulumi.Input<pulumi.Input<inputs.cloudwatch.EventTargetRunCommandTarget>[]>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
     */
    sqsTarget?: pulumi.Input<inputs.cloudwatch.EventTargetSqsTarget>;
    /**
     * The unique target assignment ID. If missing, will generate a random, unique id.
     */
    targetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EventTarget resource.
 */
export interface EventTargetArgs {
    /**
     * The Amazon Resource Name (ARN) of the target.
     */
    arn: pulumi.Input<string>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
     */
    batchTarget?: pulumi.Input<inputs.cloudwatch.EventTargetBatchTarget>;
    /**
     * Parameters used when you are providing a dead letter config. Documented below. A maximum of 1 are allowed.
     */
    deadLetterConfig?: pulumi.Input<inputs.cloudwatch.EventTargetDeadLetterConfig>;
    /**
     * Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
     */
    ecsTarget?: pulumi.Input<inputs.cloudwatch.EventTargetEcsTarget>;
    /**
     * The name or ARN of the event bus to associate with the rule.
     * If you omit this, the `default` event bus is used.
     */
    eventBusName?: pulumi.Input<string>;
    /**
     * Parameters used when you are using the rule to invoke an API Gateway REST endpoint. Documented below. A maximum of 1 is allowed.
     */
    httpTarget?: pulumi.Input<inputs.cloudwatch.EventTargetHttpTarget>;
    /**
     * Valid JSON text passed to the target. Conflicts with `inputPath` and `inputTransformer`.
     */
    input?: pulumi.Input<string>;
    /**
     * The value of the [JSONPath](http://goessner.net/articles/JsonPath/) that is used for extracting part of the matched event when passing it to the target. Conflicts with `input` and `inputTransformer`.
     */
    inputPath?: pulumi.Input<string>;
    /**
     * Parameters used when you are providing a custom input to a target based on certain event data. Conflicts with `input` and `inputPath`.
     */
    inputTransformer?: pulumi.Input<inputs.cloudwatch.EventTargetInputTransformer>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
     */
    kinesisTarget?: pulumi.Input<inputs.cloudwatch.EventTargetKinesisTarget>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon Redshift Statement. Documented below. A maximum of 1 are allowed.
     */
    redshiftTarget?: pulumi.Input<inputs.cloudwatch.EventTargetRedshiftTarget>;
    /**
     * Parameters used when you are providing retry policies. Documented below. A maximum of 1 are allowed.
     */
    retryPolicy?: pulumi.Input<inputs.cloudwatch.EventTargetRetryPolicy>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecsTarget` is used or target in `arn` is EC2 instance, Kinesis data stream, Step Functions state machine, or Event Bus in different account or region.
     */
    roleArn?: pulumi.Input<string>;
    /**
     * The name of the rule you want to add targets to.
     *
     * The following arguments are optional:
     */
    rule: pulumi.Input<string>;
    /**
     * Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
     */
    runCommandTargets?: pulumi.Input<pulumi.Input<inputs.cloudwatch.EventTargetRunCommandTarget>[]>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
     */
    sqsTarget?: pulumi.Input<inputs.cloudwatch.EventTargetSqsTarget>;
    /**
     * The unique target assignment ID. If missing, will generate a random, unique id.
     */
    targetId?: pulumi.Input<string>;
}
