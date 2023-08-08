// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Attaches a load balancer to an Auto Scaling group.
 *
 * > **NOTE on Auto Scaling Groups, Attachments and Traffic Source Attachments:** Pulumi provides standalone Attachment (for attaching Classic Load Balancers and Application Load Balancer, Gateway Load Balancer, or Network Load Balancer target groups) and Traffic Source Attachment (for attaching Load Balancers and VPC Lattice target groups) resources and an Auto Scaling Group resource with `loadBalancers`, `targetGroupArns` and `trafficSource` attributes. Do not use the same traffic source in more than one of these resources. Doing so will cause a conflict of attachments. A `lifecycle` configuration block can be used to suppress differences if necessary.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * // Create a new load balancer attachment
 * const example = new aws.autoscaling.Attachment("example", {
 *     autoscalingGroupName: aws_autoscaling_group.example.id,
 *     elb: aws_elb.example.id,
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * // Create a new ALB Target Group attachment
 * const example = new aws.autoscaling.Attachment("example", {
 *     autoscalingGroupName: aws_autoscaling_group.example.id,
 *     lbTargetGroupArn: aws_lb_target_group.example.arn,
 * });
 * ```
 */
export class Attachment extends pulumi.CustomResource {
    /**
     * Get an existing Attachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AttachmentState, opts?: pulumi.CustomResourceOptions): Attachment {
        return new Attachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:autoscaling/attachment:Attachment';

    /**
     * Returns true if the given object is an instance of Attachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Attachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Attachment.__pulumiType;
    }

    /**
     * Name of ASG to associate with the ELB.
     */
    public readonly autoscalingGroupName!: pulumi.Output<string>;
    /**
     * Name of the ELB.
     */
    public readonly elb!: pulumi.Output<string | undefined>;
    /**
     * ARN of a load balancer target group.
     */
    public readonly lbTargetGroupArn!: pulumi.Output<string | undefined>;

    /**
     * Create a Attachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AttachmentArgs | AttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AttachmentState | undefined;
            resourceInputs["autoscalingGroupName"] = state ? state.autoscalingGroupName : undefined;
            resourceInputs["elb"] = state ? state.elb : undefined;
            resourceInputs["lbTargetGroupArn"] = state ? state.lbTargetGroupArn : undefined;
        } else {
            const args = argsOrState as AttachmentArgs | undefined;
            if ((!args || args.autoscalingGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autoscalingGroupName'");
            }
            resourceInputs["autoscalingGroupName"] = args ? args.autoscalingGroupName : undefined;
            resourceInputs["elb"] = args ? args.elb : undefined;
            resourceInputs["lbTargetGroupArn"] = args ? args.lbTargetGroupArn : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Attachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Attachment resources.
 */
export interface AttachmentState {
    /**
     * Name of ASG to associate with the ELB.
     */
    autoscalingGroupName?: pulumi.Input<string>;
    /**
     * Name of the ELB.
     */
    elb?: pulumi.Input<string>;
    /**
     * ARN of a load balancer target group.
     */
    lbTargetGroupArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Attachment resource.
 */
export interface AttachmentArgs {
    /**
     * Name of ASG to associate with the ELB.
     */
    autoscalingGroupName: pulumi.Input<string>;
    /**
     * Name of the ELB.
     */
    elb?: pulumi.Input<string>;
    /**
     * ARN of a load balancer target group.
     */
    lbTargetGroupArn?: pulumi.Input<string>;
}
