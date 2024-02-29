// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides an AWS Config Configuration Recorder. Please note that this resource **does not start** the created recorder automatically.
 *
 * > **Note:** _Starting_ the Configuration Recorder requires a delivery channel (while delivery channel creation requires Configuration Recorder). This is why `aws.cfg.RecorderStatus` is a separate resource.
 *
 * ## Example Usage
 * ### Basic Usage
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
 *             identifiers: ["config.amazonaws.com"],
 *         }],
 *         actions: ["sts:AssumeRole"],
 *     }],
 * });
 * const role = new aws.iam.Role("role", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
 * const foo = new aws.cfg.Recorder("foo", {roleArn: role.arn});
 * ```
 * ### Exclude Resources Types Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const foo = new aws.cfg.Recorder("foo", {
 *     roleArn: aws_iam_role.r.arn,
 *     recordingGroup: {
 *         allSupported: false,
 *         exclusionByResourceTypes: [{
 *             resourceTypes: ["AWS::EC2::Instance"],
 *         }],
 *         recordingStrategies: [{
 *             useOnly: "EXCLUSION_BY_RESOURCE_TYPES",
 *         }],
 *     },
 * });
 * ```
 * ### Periodic Recording
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const foo = new aws.cfg.Recorder("foo", {
 *     roleArn: aws_iam_role.r.arn,
 *     recordingGroup: {
 *         allSupported: false,
 *         includeGlobalResourceTypes: false,
 *         resourceTypes: [
 *             "AWS::EC2::Instance",
 *             "AWS::EC2::NetworkInterface",
 *         ],
 *     },
 *     recordingMode: {
 *         recordingFrequency: "CONTINUOUS",
 *         recordingModeOverride: {
 *             description: "Only record EC2 network interfaces daily",
 *             resourceTypes: ["AWS::EC2::NetworkInterface"],
 *             recordingFrequency: "DAILY",
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Using `pulumi import`, import Configuration Recorder using the name. For example:
 *
 * ```sh
 *  $ pulumi import aws:cfg/recorder:Recorder foo example
 * ```
 */
export class Recorder extends pulumi.CustomResource {
    /**
     * Get an existing Recorder resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RecorderState, opts?: pulumi.CustomResourceOptions): Recorder {
        return new Recorder(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cfg/recorder:Recorder';

    /**
     * Returns true if the given object is an instance of Recorder.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Recorder {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Recorder.__pulumiType;
    }

    /**
     * The name of the recorder. Defaults to `default`. Changing it recreates the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Recording group - see below.
     */
    public readonly recordingGroup!: pulumi.Output<outputs.cfg.RecorderRecordingGroup>;
    /**
     * Recording mode - see below.
     */
    public readonly recordingMode!: pulumi.Output<outputs.cfg.RecorderRecordingMode>;
    /**
     * Amazon Resource Name (ARN) of the IAM role. Used to make read or write requests to the delivery channel and to describe the AWS resources associated with the account. See [AWS Docs](http://docs.aws.amazon.com/config/latest/developerguide/iamrole-permissions.html) for more details.
     */
    public readonly roleArn!: pulumi.Output<string>;

    /**
     * Create a Recorder resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RecorderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RecorderArgs | RecorderState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RecorderState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["recordingGroup"] = state ? state.recordingGroup : undefined;
            resourceInputs["recordingMode"] = state ? state.recordingMode : undefined;
            resourceInputs["roleArn"] = state ? state.roleArn : undefined;
        } else {
            const args = argsOrState as RecorderArgs | undefined;
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["recordingGroup"] = args ? args.recordingGroup : undefined;
            resourceInputs["recordingMode"] = args ? args.recordingMode : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Recorder.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Recorder resources.
 */
export interface RecorderState {
    /**
     * The name of the recorder. Defaults to `default`. Changing it recreates the resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Recording group - see below.
     */
    recordingGroup?: pulumi.Input<inputs.cfg.RecorderRecordingGroup>;
    /**
     * Recording mode - see below.
     */
    recordingMode?: pulumi.Input<inputs.cfg.RecorderRecordingMode>;
    /**
     * Amazon Resource Name (ARN) of the IAM role. Used to make read or write requests to the delivery channel and to describe the AWS resources associated with the account. See [AWS Docs](http://docs.aws.amazon.com/config/latest/developerguide/iamrole-permissions.html) for more details.
     */
    roleArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Recorder resource.
 */
export interface RecorderArgs {
    /**
     * The name of the recorder. Defaults to `default`. Changing it recreates the resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Recording group - see below.
     */
    recordingGroup?: pulumi.Input<inputs.cfg.RecorderRecordingGroup>;
    /**
     * Recording mode - see below.
     */
    recordingMode?: pulumi.Input<inputs.cfg.RecorderRecordingMode>;
    /**
     * Amazon Resource Name (ARN) of the IAM role. Used to make read or write requests to the delivery channel and to describe the AWS resources associated with the account. See [AWS Docs](http://docs.aws.amazon.com/config/latest/developerguide/iamrole-permissions.html) for more details.
     */
    roleArn: pulumi.Input<string>;
}
