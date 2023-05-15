// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Cloud9 EC2 Development Environment.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.cloud9.EnvironmentEC2("example", {instanceType: "t2.micro"});
 * ```
 *
 * Get the URL of the Cloud9 environment after creation:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.cloud9.EnvironmentEC2("example", {instanceType: "t2.micro"});
 * const cloud9Instance = aws.ec2.getInstanceOutput({
 *     filters: [{
 *         name: "tag:aws:cloud9:environment",
 *         values: [example.id],
 *     }],
 * });
 * export const cloud9Url = pulumi.interpolate`https://${_var.region}.console.aws.amazon.com/cloud9/ide/${example.id}`;
 * ```
 *
 * Allocate a static IP to the Cloud9 environment:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.cloud9.EnvironmentEC2("example", {instanceType: "t2.micro"});
 * const cloud9Instance = aws.ec2.getInstanceOutput({
 *     filters: [{
 *         name: "tag:aws:cloud9:environment",
 *         values: [example.id],
 *     }],
 * });
 * const cloud9Eip = new aws.ec2.Eip("cloud9Eip", {
 *     instance: cloud9Instance.apply(cloud9Instance => cloud9Instance.id),
 *     vpc: true,
 * });
 * export const cloud9PublicIp = cloud9Eip.publicIp;
 * ```
 */
export class EnvironmentEC2 extends pulumi.CustomResource {
    /**
     * Get an existing EnvironmentEC2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EnvironmentEC2State, opts?: pulumi.CustomResourceOptions): EnvironmentEC2 {
        return new EnvironmentEC2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloud9/environmentEC2:EnvironmentEC2';

    /**
     * Returns true if the given object is an instance of EnvironmentEC2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EnvironmentEC2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EnvironmentEC2.__pulumiType;
    }

    /**
     * The ARN of the environment.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The number of minutes until the running instance is shut down after the environment has last been used.
     */
    public readonly automaticStopTimeMinutes!: pulumi.Output<number | undefined>;
    /**
     * The connection type used for connecting to an Amazon EC2 environment. Valid values are `CONNECT_SSH` and `CONNECT_SSM`. For more information please refer [AWS documentation for Cloud9](https://docs.aws.amazon.com/cloud9/latest/user-guide/ec2-ssm.html).
     */
    public readonly connectionType!: pulumi.Output<string | undefined>;
    /**
     * The description of the environment.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The identifier for the Amazon Machine Image (AMI) that's used to create the EC2 instance. Valid values are
     * * `amazonlinux-1-x86_64`
     * * `amazonlinux-2-x86_64`
     * * `ubuntu-18.04-x86_64`
     * * `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-1-x86_64`
     * * `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-2-x86_64`
     * * `resolve:ssm:/aws/service/cloud9/amis/ubuntu-18.04-x86_64`
     */
    public readonly imageId!: pulumi.Output<string | undefined>;
    /**
     * The type of instance to connect to the environment, e.g., `t2.micro`.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * The name of the environment.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ARN of the environment owner. This can be ARN of any AWS IAM principal. Defaults to the environment's creator.
     */
    public readonly ownerArn!: pulumi.Output<string>;
    /**
     * The ID of the subnet in Amazon VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
     */
    public readonly subnetId!: pulumi.Output<string | undefined>;
    /**
     * Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The type of the environment (e.g., `ssh` or `ec2`)
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a EnvironmentEC2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnvironmentEC2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EnvironmentEC2Args | EnvironmentEC2State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EnvironmentEC2State | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["automaticStopTimeMinutes"] = state ? state.automaticStopTimeMinutes : undefined;
            resourceInputs["connectionType"] = state ? state.connectionType : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["imageId"] = state ? state.imageId : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["ownerArn"] = state ? state.ownerArn : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as EnvironmentEC2Args | undefined;
            if ((!args || args.instanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceType'");
            }
            resourceInputs["automaticStopTimeMinutes"] = args ? args.automaticStopTimeMinutes : undefined;
            resourceInputs["connectionType"] = args ? args.connectionType : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["imageId"] = args ? args.imageId : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ownerArn"] = args ? args.ownerArn : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EnvironmentEC2.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EnvironmentEC2 resources.
 */
export interface EnvironmentEC2State {
    /**
     * The ARN of the environment.
     */
    arn?: pulumi.Input<string>;
    /**
     * The number of minutes until the running instance is shut down after the environment has last been used.
     */
    automaticStopTimeMinutes?: pulumi.Input<number>;
    /**
     * The connection type used for connecting to an Amazon EC2 environment. Valid values are `CONNECT_SSH` and `CONNECT_SSM`. For more information please refer [AWS documentation for Cloud9](https://docs.aws.amazon.com/cloud9/latest/user-guide/ec2-ssm.html).
     */
    connectionType?: pulumi.Input<string>;
    /**
     * The description of the environment.
     */
    description?: pulumi.Input<string>;
    /**
     * The identifier for the Amazon Machine Image (AMI) that's used to create the EC2 instance. Valid values are
     * * `amazonlinux-1-x86_64`
     * * `amazonlinux-2-x86_64`
     * * `ubuntu-18.04-x86_64`
     * * `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-1-x86_64`
     * * `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-2-x86_64`
     * * `resolve:ssm:/aws/service/cloud9/amis/ubuntu-18.04-x86_64`
     */
    imageId?: pulumi.Input<string>;
    /**
     * The type of instance to connect to the environment, e.g., `t2.micro`.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * The name of the environment.
     */
    name?: pulumi.Input<string>;
    /**
     * The ARN of the environment owner. This can be ARN of any AWS IAM principal. Defaults to the environment's creator.
     */
    ownerArn?: pulumi.Input<string>;
    /**
     * The ID of the subnet in Amazon VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of the environment (e.g., `ssh` or `ec2`)
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EnvironmentEC2 resource.
 */
export interface EnvironmentEC2Args {
    /**
     * The number of minutes until the running instance is shut down after the environment has last been used.
     */
    automaticStopTimeMinutes?: pulumi.Input<number>;
    /**
     * The connection type used for connecting to an Amazon EC2 environment. Valid values are `CONNECT_SSH` and `CONNECT_SSM`. For more information please refer [AWS documentation for Cloud9](https://docs.aws.amazon.com/cloud9/latest/user-guide/ec2-ssm.html).
     */
    connectionType?: pulumi.Input<string>;
    /**
     * The description of the environment.
     */
    description?: pulumi.Input<string>;
    /**
     * The identifier for the Amazon Machine Image (AMI) that's used to create the EC2 instance. Valid values are
     * * `amazonlinux-1-x86_64`
     * * `amazonlinux-2-x86_64`
     * * `ubuntu-18.04-x86_64`
     * * `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-1-x86_64`
     * * `resolve:ssm:/aws/service/cloud9/amis/amazonlinux-2-x86_64`
     * * `resolve:ssm:/aws/service/cloud9/amis/ubuntu-18.04-x86_64`
     */
    imageId?: pulumi.Input<string>;
    /**
     * The type of instance to connect to the environment, e.g., `t2.micro`.
     */
    instanceType: pulumi.Input<string>;
    /**
     * The name of the environment.
     */
    name?: pulumi.Input<string>;
    /**
     * The ARN of the environment owner. This can be ARN of any AWS IAM principal. Defaults to the environment's creator.
     */
    ownerArn?: pulumi.Input<string>;
    /**
     * The ID of the subnet in Amazon VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
