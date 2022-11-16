// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an AutoScaling Lifecycle Hook resource.
 *
 * > **NOTE:** This provider has two types of ways you can add lifecycle hooks - via
 * the `initialLifecycleHook` attribute from the
 * `aws.autoscaling.Group`
 * resource, or via this one. Hooks added via this resource will not be added
 * until the autoscaling group has been created, and depending on your
 * capacity
 * settings, after the initial instances have been launched, creating unintended
 * behavior. If you need hooks to run on all instances, add them with
 * `initialLifecycleHook` in
 * `aws.autoscaling.Group`,
 * but take care to not duplicate those hooks with this resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const foobarGroup = new aws.autoscaling.Group("foobarGroup", {
 *     availabilityZones: ["us-west-2a"],
 *     healthCheckType: "EC2",
 *     terminationPolicies: ["OldestInstance"],
 *     tags: [{
 *         key: "Foo",
 *         value: "foo-bar",
 *         propagateAtLaunch: true,
 *     }],
 * });
 * const foobarLifecycleHook = new aws.autoscaling.LifecycleHook("foobarLifecycleHook", {
 *     autoscalingGroupName: foobarGroup.name,
 *     defaultResult: "CONTINUE",
 *     heartbeatTimeout: 2000,
 *     lifecycleTransition: "autoscaling:EC2_INSTANCE_LAUNCHING",
 *     notificationMetadata: `{
 *   "foo": "bar"
 * }
 * `,
 *     notificationTargetArn: "arn:aws:sqs:us-east-1:444455556666:queue1*",
 *     roleArn: "arn:aws:iam::123456789012:role/S3Access",
 * });
 * ```
 *
 * ## Import
 *
 * AutoScaling Lifecycle Hooks can be imported using the role autoscaling_group_name and name separated by `/`.
 *
 * ```sh
 *  $ pulumi import aws:autoscaling/lifecycleHook:LifecycleHook test-lifecycle-hook asg-name/lifecycle-hook-name
 * ```
 */
export class LifecycleHook extends pulumi.CustomResource {
    /**
     * Get an existing LifecycleHook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LifecycleHookState, opts?: pulumi.CustomResourceOptions): LifecycleHook {
        return new LifecycleHook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:autoscaling/lifecycleHook:LifecycleHook';

    /**
     * Returns true if the given object is an instance of LifecycleHook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LifecycleHook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LifecycleHook.__pulumiType;
    }

    /**
     * Name of the Auto Scaling group to which you want to assign the lifecycle hook
     */
    public readonly autoscalingGroupName!: pulumi.Output<string>;
    /**
     * Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. The value for this parameter can be either CONTINUE or ABANDON. The default value for this parameter is ABANDON.
     */
    public readonly defaultResult!: pulumi.Output<string>;
    /**
     * Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the DefaultResult parameter
     */
    public readonly heartbeatTimeout!: pulumi.Output<number | undefined>;
    /**
     * Instance state to which you want to attach the lifecycle hook. For a list of lifecycle hook types, see [describe-lifecycle-hook-types](https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-lifecycle-hook-types.html#examples)
     */
    public readonly lifecycleTransition!: pulumi.Output<string>;
    /**
     * Name of the lifecycle hook.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Contains additional information that you want to include any time Auto Scaling sends a message to the notification target.
     */
    public readonly notificationMetadata!: pulumi.Output<string | undefined>;
    /**
     * ARN of the notification target that Auto Scaling will use to notify you when an instance is in the transition state for the lifecycle hook. This ARN target can be either an SQS queue or an SNS topic.
     */
    public readonly notificationTargetArn!: pulumi.Output<string | undefined>;
    /**
     * ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
     */
    public readonly roleArn!: pulumi.Output<string | undefined>;

    /**
     * Create a LifecycleHook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LifecycleHookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LifecycleHookArgs | LifecycleHookState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LifecycleHookState | undefined;
            resourceInputs["autoscalingGroupName"] = state ? state.autoscalingGroupName : undefined;
            resourceInputs["defaultResult"] = state ? state.defaultResult : undefined;
            resourceInputs["heartbeatTimeout"] = state ? state.heartbeatTimeout : undefined;
            resourceInputs["lifecycleTransition"] = state ? state.lifecycleTransition : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notificationMetadata"] = state ? state.notificationMetadata : undefined;
            resourceInputs["notificationTargetArn"] = state ? state.notificationTargetArn : undefined;
            resourceInputs["roleArn"] = state ? state.roleArn : undefined;
        } else {
            const args = argsOrState as LifecycleHookArgs | undefined;
            if ((!args || args.autoscalingGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autoscalingGroupName'");
            }
            if ((!args || args.lifecycleTransition === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lifecycleTransition'");
            }
            resourceInputs["autoscalingGroupName"] = args ? args.autoscalingGroupName : undefined;
            resourceInputs["defaultResult"] = args ? args.defaultResult : undefined;
            resourceInputs["heartbeatTimeout"] = args ? args.heartbeatTimeout : undefined;
            resourceInputs["lifecycleTransition"] = args ? args.lifecycleTransition : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notificationMetadata"] = args ? args.notificationMetadata : undefined;
            resourceInputs["notificationTargetArn"] = args ? args.notificationTargetArn : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LifecycleHook.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LifecycleHook resources.
 */
export interface LifecycleHookState {
    /**
     * Name of the Auto Scaling group to which you want to assign the lifecycle hook
     */
    autoscalingGroupName?: pulumi.Input<string>;
    /**
     * Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. The value for this parameter can be either CONTINUE or ABANDON. The default value for this parameter is ABANDON.
     */
    defaultResult?: pulumi.Input<string>;
    /**
     * Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the DefaultResult parameter
     */
    heartbeatTimeout?: pulumi.Input<number>;
    /**
     * Instance state to which you want to attach the lifecycle hook. For a list of lifecycle hook types, see [describe-lifecycle-hook-types](https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-lifecycle-hook-types.html#examples)
     */
    lifecycleTransition?: pulumi.Input<string>;
    /**
     * Name of the lifecycle hook.
     */
    name?: pulumi.Input<string>;
    /**
     * Contains additional information that you want to include any time Auto Scaling sends a message to the notification target.
     */
    notificationMetadata?: pulumi.Input<string>;
    /**
     * ARN of the notification target that Auto Scaling will use to notify you when an instance is in the transition state for the lifecycle hook. This ARN target can be either an SQS queue or an SNS topic.
     */
    notificationTargetArn?: pulumi.Input<string>;
    /**
     * ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
     */
    roleArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LifecycleHook resource.
 */
export interface LifecycleHookArgs {
    /**
     * Name of the Auto Scaling group to which you want to assign the lifecycle hook
     */
    autoscalingGroupName: pulumi.Input<string>;
    /**
     * Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. The value for this parameter can be either CONTINUE or ABANDON. The default value for this parameter is ABANDON.
     */
    defaultResult?: pulumi.Input<string>;
    /**
     * Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the DefaultResult parameter
     */
    heartbeatTimeout?: pulumi.Input<number>;
    /**
     * Instance state to which you want to attach the lifecycle hook. For a list of lifecycle hook types, see [describe-lifecycle-hook-types](https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-lifecycle-hook-types.html#examples)
     */
    lifecycleTransition: pulumi.Input<string>;
    /**
     * Name of the lifecycle hook.
     */
    name?: pulumi.Input<string>;
    /**
     * Contains additional information that you want to include any time Auto Scaling sends a message to the notification target.
     */
    notificationMetadata?: pulumi.Input<string>;
    /**
     * ARN of the notification target that Auto Scaling will use to notify you when an instance is in the transition state for the lifecycle hook. This ARN target can be either an SQS queue or an SNS topic.
     */
    notificationTargetArn?: pulumi.Input<string>;
    /**
     * ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
     */
    roleArn?: pulumi.Input<string>;
}
