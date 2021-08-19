// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a CloudFormation Stack resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const network = new aws.cloudformation.Stack("network", {
 *     parameters: {
 *         VPCCidr: "10.0.0.0/16",
 *     },
 *     templateBody: `{
 *   "Parameters" : {
 *     "VPCCidr" : {
 *       "Type" : "String",
 *       "Default" : "10.0.0.0/16",
 *       "Description" : "Enter the CIDR block for the VPC. Default is 10.0.0.0/16."
 *     }
 *   },
 *   "Resources" : {
 *     "myVpc": {
 *       "Type" : "AWS::EC2::VPC",
 *       "Properties" : {
 *         "CidrBlock" : { "Ref" : "VPCCidr" },
 *         "Tags" : [
 *           {"Key": "Name", "Value": "Primary_CF_VPC"}
 *         ]
 *       }
 *     }
 *   }
 * }
 * `,
 * });
 * ```
 *
 * ## Import
 *
 * Cloudformation Stacks can be imported using the `name`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:cloudformation/stack:Stack stack networking-stack
 * ```
 */
export class Stack extends pulumi.CustomResource {
    /**
     * Get an existing Stack resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StackState, opts?: pulumi.CustomResourceOptions): Stack {
        return new Stack(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloudformation/stack:Stack';

    /**
     * Returns true if the given object is an instance of Stack.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Stack {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Stack.__pulumiType;
    }

    /**
     * A list of capabilities.
     * Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, or `CAPABILITY_AUTO_EXPAND`
     */
    public readonly capabilities!: pulumi.Output<string[] | undefined>;
    /**
     * Set to true to disable rollback of the stack if stack creation failed.
     * Conflicts with `onFailure`.
     */
    public readonly disableRollback!: pulumi.Output<boolean | undefined>;
    /**
     * The ARN of an IAM role that AWS CloudFormation assumes to create the stack. If you don't specify a value, AWS CloudFormation uses the role that was previously associated with the stack. If no role is available, AWS CloudFormation uses a temporary session that is generated from your user credentials.
     */
    public readonly iamRoleArn!: pulumi.Output<string | undefined>;
    /**
     * Stack name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of SNS topic ARNs to publish stack related events.
     */
    public readonly notificationArns!: pulumi.Output<string[] | undefined>;
    /**
     * Action to be taken if stack creation fails. This must be
     * one of: `DO_NOTHING`, `ROLLBACK`, or `DELETE`. Conflicts with `disableRollback`.
     */
    public readonly onFailure!: pulumi.Output<string | undefined>;
    /**
     * A map of outputs from the stack.
     */
    public /*out*/ readonly outputs!: pulumi.Output<{[key: string]: string}>;
    /**
     * A map of Parameter structures that specify input parameters for the stack.
     */
    public readonly parameters!: pulumi.Output<{[key: string]: string}>;
    /**
     * Structure containing the stack policy body.
     * Conflicts w/ `policyUrl`.
     */
    public readonly policyBody!: pulumi.Output<string>;
    /**
     * Location of a file containing the stack policy.
     * Conflicts w/ `policyBody`.
     */
    public readonly policyUrl!: pulumi.Output<string | undefined>;
    /**
     * Map of resource tags to associate with this stack. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Structure containing the template body (max size: 51,200 bytes).
     */
    public readonly templateBody!: pulumi.Output<string>;
    /**
     * Location of a file containing the template body (max size: 460,800 bytes).
     */
    public readonly templateUrl!: pulumi.Output<string | undefined>;
    /**
     * The amount of time that can pass before the stack status becomes `CREATE_FAILED`.
     */
    public readonly timeoutInMinutes!: pulumi.Output<number | undefined>;

    /**
     * Create a Stack resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: StackArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StackArgs | StackState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StackState | undefined;
            inputs["capabilities"] = state ? state.capabilities : undefined;
            inputs["disableRollback"] = state ? state.disableRollback : undefined;
            inputs["iamRoleArn"] = state ? state.iamRoleArn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notificationArns"] = state ? state.notificationArns : undefined;
            inputs["onFailure"] = state ? state.onFailure : undefined;
            inputs["outputs"] = state ? state.outputs : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["policyBody"] = state ? state.policyBody : undefined;
            inputs["policyUrl"] = state ? state.policyUrl : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["templateBody"] = state ? state.templateBody : undefined;
            inputs["templateUrl"] = state ? state.templateUrl : undefined;
            inputs["timeoutInMinutes"] = state ? state.timeoutInMinutes : undefined;
        } else {
            const args = argsOrState as StackArgs | undefined;
            inputs["capabilities"] = args ? args.capabilities : undefined;
            inputs["disableRollback"] = args ? args.disableRollback : undefined;
            inputs["iamRoleArn"] = args ? args.iamRoleArn : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notificationArns"] = args ? args.notificationArns : undefined;
            inputs["onFailure"] = args ? args.onFailure : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["policyBody"] = args ? args.policyBody : undefined;
            inputs["policyUrl"] = args ? args.policyUrl : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["templateBody"] = args ? args.templateBody : undefined;
            inputs["templateUrl"] = args ? args.templateUrl : undefined;
            inputs["timeoutInMinutes"] = args ? args.timeoutInMinutes : undefined;
            inputs["outputs"] = undefined /*out*/;
            inputs["tagsAll"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Stack.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Stack resources.
 */
export interface StackState {
    /**
     * A list of capabilities.
     * Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, or `CAPABILITY_AUTO_EXPAND`
     */
    capabilities?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set to true to disable rollback of the stack if stack creation failed.
     * Conflicts with `onFailure`.
     */
    disableRollback?: pulumi.Input<boolean>;
    /**
     * The ARN of an IAM role that AWS CloudFormation assumes to create the stack. If you don't specify a value, AWS CloudFormation uses the role that was previously associated with the stack. If no role is available, AWS CloudFormation uses a temporary session that is generated from your user credentials.
     */
    iamRoleArn?: pulumi.Input<string>;
    /**
     * Stack name.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of SNS topic ARNs to publish stack related events.
     */
    notificationArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Action to be taken if stack creation fails. This must be
     * one of: `DO_NOTHING`, `ROLLBACK`, or `DELETE`. Conflicts with `disableRollback`.
     */
    onFailure?: pulumi.Input<string>;
    /**
     * A map of outputs from the stack.
     */
    outputs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of Parameter structures that specify input parameters for the stack.
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Structure containing the stack policy body.
     * Conflicts w/ `policyUrl`.
     */
    policyBody?: pulumi.Input<string>;
    /**
     * Location of a file containing the stack policy.
     * Conflicts w/ `policyBody`.
     */
    policyUrl?: pulumi.Input<string>;
    /**
     * Map of resource tags to associate with this stack. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Structure containing the template body (max size: 51,200 bytes).
     */
    templateBody?: pulumi.Input<string>;
    /**
     * Location of a file containing the template body (max size: 460,800 bytes).
     */
    templateUrl?: pulumi.Input<string>;
    /**
     * The amount of time that can pass before the stack status becomes `CREATE_FAILED`.
     */
    timeoutInMinutes?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Stack resource.
 */
export interface StackArgs {
    /**
     * A list of capabilities.
     * Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, or `CAPABILITY_AUTO_EXPAND`
     */
    capabilities?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set to true to disable rollback of the stack if stack creation failed.
     * Conflicts with `onFailure`.
     */
    disableRollback?: pulumi.Input<boolean>;
    /**
     * The ARN of an IAM role that AWS CloudFormation assumes to create the stack. If you don't specify a value, AWS CloudFormation uses the role that was previously associated with the stack. If no role is available, AWS CloudFormation uses a temporary session that is generated from your user credentials.
     */
    iamRoleArn?: pulumi.Input<string>;
    /**
     * Stack name.
     */
    name?: pulumi.Input<string>;
    /**
     * A list of SNS topic ARNs to publish stack related events.
     */
    notificationArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Action to be taken if stack creation fails. This must be
     * one of: `DO_NOTHING`, `ROLLBACK`, or `DELETE`. Conflicts with `disableRollback`.
     */
    onFailure?: pulumi.Input<string>;
    /**
     * A map of Parameter structures that specify input parameters for the stack.
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Structure containing the stack policy body.
     * Conflicts w/ `policyUrl`.
     */
    policyBody?: pulumi.Input<string>;
    /**
     * Location of a file containing the stack policy.
     * Conflicts w/ `policyBody`.
     */
    policyUrl?: pulumi.Input<string>;
    /**
     * Map of resource tags to associate with this stack. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Structure containing the template body (max size: 51,200 bytes).
     */
    templateBody?: pulumi.Input<string>;
    /**
     * Location of a file containing the template body (max size: 460,800 bytes).
     */
    templateUrl?: pulumi.Input<string>;
    /**
     * The amount of time that can pass before the stack status becomes `CREATE_FAILED`.
     */
    timeoutInMinutes?: pulumi.Input<number>;
}
