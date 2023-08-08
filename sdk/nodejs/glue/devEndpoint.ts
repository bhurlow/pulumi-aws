// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Glue Development Endpoint resource.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const examplePolicyDocument = aws.iam.getPolicyDocument({
 *     statements: [{
 *         actions: ["sts:AssumeRole"],
 *         principals: [{
 *             type: "Service",
 *             identifiers: ["glue.amazonaws.com"],
 *         }],
 *     }],
 * });
 * const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json)});
 * const exampleDevEndpoint = new aws.glue.DevEndpoint("exampleDevEndpoint", {roleArn: exampleRole.arn});
 * const example_AWSGlueServiceRole = new aws.iam.RolePolicyAttachment("example-AWSGlueServiceRole", {
 *     policyArn: "arn:aws:iam::aws:policy/service-role/AWSGlueServiceRole",
 *     role: exampleRole.name,
 * });
 * ```
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_glue_dev_endpoint.example
 *
 *  id = "foo" } Using `pulumi import`, import a Glue Development Endpoint using the `name`. For exampleconsole % pulumi import aws_glue_dev_endpoint.example foo
 */
export class DevEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing DevEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DevEndpointState, opts?: pulumi.CustomResourceOptions): DevEndpoint {
        return new DevEndpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:glue/devEndpoint:DevEndpoint';

    /**
     * Returns true if the given object is an instance of DevEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DevEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DevEndpoint.__pulumiType;
    }

    /**
     * A map of arguments used to configure the endpoint.
     */
    public readonly arguments!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ARN of the endpoint.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The AWS availability zone where this endpoint is located.
     */
    public /*out*/ readonly availabilityZone!: pulumi.Output<string>;
    /**
     * Path to one or more Java Jars in an S3 bucket that should be loaded in this endpoint.
     */
    public readonly extraJarsS3Path!: pulumi.Output<string | undefined>;
    /**
     * Path(s) to one or more Python libraries in an S3 bucket that should be loaded in this endpoint. Multiple values must be complete paths separated by a comma.
     */
    public readonly extraPythonLibsS3Path!: pulumi.Output<string | undefined>;
    /**
     * The reason for a current failure in this endpoint.
     */
    public /*out*/ readonly failureReason!: pulumi.Output<string>;
    /**
     * Specifies the versions of Python and Apache Spark to use. Defaults to AWS Glue version 0.9.
     */
    public readonly glueVersion!: pulumi.Output<string | undefined>;
    /**
     * The name of this endpoint. It must be unique in your account.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The number of AWS Glue Data Processing Units (DPUs) to allocate to this endpoint. Conflicts with `workerType`.
     */
    public readonly numberOfNodes!: pulumi.Output<number | undefined>;
    /**
     * The number of workers of a defined worker type that are allocated to this endpoint. This field is available only when you choose worker type G.1X or G.2X.
     */
    public readonly numberOfWorkers!: pulumi.Output<number | undefined>;
    /**
     * A private IP address to access the endpoint within a VPC, if this endpoint is created within one.
     */
    public /*out*/ readonly privateAddress!: pulumi.Output<string>;
    /**
     * The public IP address used by this endpoint. The PublicAddress field is present only when you create a non-VPC endpoint.
     */
    public /*out*/ readonly publicAddress!: pulumi.Output<string>;
    /**
     * The public key to be used by this endpoint for authentication.
     */
    public readonly publicKey!: pulumi.Output<string | undefined>;
    /**
     * A list of public keys to be used by this endpoint for authentication.
     */
    public readonly publicKeys!: pulumi.Output<string[] | undefined>;
    /**
     * The IAM role for this endpoint.
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * The name of the Security Configuration structure to be used with this endpoint.
     */
    public readonly securityConfiguration!: pulumi.Output<string | undefined>;
    /**
     * Security group IDs for the security groups to be used by this endpoint.
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * The current status of this endpoint.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The subnet ID for the new endpoint to use.
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
     * he ID of the VPC used by this endpoint.
     */
    public /*out*/ readonly vpcId!: pulumi.Output<string>;
    /**
     * The type of predefined worker that is allocated to this endpoint. Accepts a value of Standard, G.1X, or G.2X.
     */
    public readonly workerType!: pulumi.Output<string | undefined>;
    /**
     * The YARN endpoint address used by this endpoint.
     */
    public /*out*/ readonly yarnEndpointAddress!: pulumi.Output<string>;
    /**
     * The Apache Zeppelin port for the remote Apache Spark interpreter.
     */
    public /*out*/ readonly zeppelinRemoteSparkInterpreterPort!: pulumi.Output<number>;

    /**
     * Create a DevEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DevEndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DevEndpointArgs | DevEndpointState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DevEndpointState | undefined;
            resourceInputs["arguments"] = state ? state.arguments : undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            resourceInputs["extraJarsS3Path"] = state ? state.extraJarsS3Path : undefined;
            resourceInputs["extraPythonLibsS3Path"] = state ? state.extraPythonLibsS3Path : undefined;
            resourceInputs["failureReason"] = state ? state.failureReason : undefined;
            resourceInputs["glueVersion"] = state ? state.glueVersion : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["numberOfNodes"] = state ? state.numberOfNodes : undefined;
            resourceInputs["numberOfWorkers"] = state ? state.numberOfWorkers : undefined;
            resourceInputs["privateAddress"] = state ? state.privateAddress : undefined;
            resourceInputs["publicAddress"] = state ? state.publicAddress : undefined;
            resourceInputs["publicKey"] = state ? state.publicKey : undefined;
            resourceInputs["publicKeys"] = state ? state.publicKeys : undefined;
            resourceInputs["roleArn"] = state ? state.roleArn : undefined;
            resourceInputs["securityConfiguration"] = state ? state.securityConfiguration : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["workerType"] = state ? state.workerType : undefined;
            resourceInputs["yarnEndpointAddress"] = state ? state.yarnEndpointAddress : undefined;
            resourceInputs["zeppelinRemoteSparkInterpreterPort"] = state ? state.zeppelinRemoteSparkInterpreterPort : undefined;
        } else {
            const args = argsOrState as DevEndpointArgs | undefined;
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            resourceInputs["arguments"] = args ? args.arguments : undefined;
            resourceInputs["extraJarsS3Path"] = args ? args.extraJarsS3Path : undefined;
            resourceInputs["extraPythonLibsS3Path"] = args ? args.extraPythonLibsS3Path : undefined;
            resourceInputs["glueVersion"] = args ? args.glueVersion : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["numberOfNodes"] = args ? args.numberOfNodes : undefined;
            resourceInputs["numberOfWorkers"] = args ? args.numberOfWorkers : undefined;
            resourceInputs["publicKey"] = args ? args.publicKey : undefined;
            resourceInputs["publicKeys"] = args ? args.publicKeys : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["securityConfiguration"] = args ? args.securityConfiguration : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["workerType"] = args ? args.workerType : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["availabilityZone"] = undefined /*out*/;
            resourceInputs["failureReason"] = undefined /*out*/;
            resourceInputs["privateAddress"] = undefined /*out*/;
            resourceInputs["publicAddress"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
            resourceInputs["yarnEndpointAddress"] = undefined /*out*/;
            resourceInputs["zeppelinRemoteSparkInterpreterPort"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DevEndpoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DevEndpoint resources.
 */
export interface DevEndpointState {
    /**
     * A map of arguments used to configure the endpoint.
     */
    arguments?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ARN of the endpoint.
     */
    arn?: pulumi.Input<string>;
    /**
     * The AWS availability zone where this endpoint is located.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Path to one or more Java Jars in an S3 bucket that should be loaded in this endpoint.
     */
    extraJarsS3Path?: pulumi.Input<string>;
    /**
     * Path(s) to one or more Python libraries in an S3 bucket that should be loaded in this endpoint. Multiple values must be complete paths separated by a comma.
     */
    extraPythonLibsS3Path?: pulumi.Input<string>;
    /**
     * The reason for a current failure in this endpoint.
     */
    failureReason?: pulumi.Input<string>;
    /**
     * Specifies the versions of Python and Apache Spark to use. Defaults to AWS Glue version 0.9.
     */
    glueVersion?: pulumi.Input<string>;
    /**
     * The name of this endpoint. It must be unique in your account.
     */
    name?: pulumi.Input<string>;
    /**
     * The number of AWS Glue Data Processing Units (DPUs) to allocate to this endpoint. Conflicts with `workerType`.
     */
    numberOfNodes?: pulumi.Input<number>;
    /**
     * The number of workers of a defined worker type that are allocated to this endpoint. This field is available only when you choose worker type G.1X or G.2X.
     */
    numberOfWorkers?: pulumi.Input<number>;
    /**
     * A private IP address to access the endpoint within a VPC, if this endpoint is created within one.
     */
    privateAddress?: pulumi.Input<string>;
    /**
     * The public IP address used by this endpoint. The PublicAddress field is present only when you create a non-VPC endpoint.
     */
    publicAddress?: pulumi.Input<string>;
    /**
     * The public key to be used by this endpoint for authentication.
     */
    publicKey?: pulumi.Input<string>;
    /**
     * A list of public keys to be used by this endpoint for authentication.
     */
    publicKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The IAM role for this endpoint.
     */
    roleArn?: pulumi.Input<string>;
    /**
     * The name of the Security Configuration structure to be used with this endpoint.
     */
    securityConfiguration?: pulumi.Input<string>;
    /**
     * Security group IDs for the security groups to be used by this endpoint.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The current status of this endpoint.
     */
    status?: pulumi.Input<string>;
    /**
     * The subnet ID for the new endpoint to use.
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
     * he ID of the VPC used by this endpoint.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The type of predefined worker that is allocated to this endpoint. Accepts a value of Standard, G.1X, or G.2X.
     */
    workerType?: pulumi.Input<string>;
    /**
     * The YARN endpoint address used by this endpoint.
     */
    yarnEndpointAddress?: pulumi.Input<string>;
    /**
     * The Apache Zeppelin port for the remote Apache Spark interpreter.
     */
    zeppelinRemoteSparkInterpreterPort?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a DevEndpoint resource.
 */
export interface DevEndpointArgs {
    /**
     * A map of arguments used to configure the endpoint.
     */
    arguments?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Path to one or more Java Jars in an S3 bucket that should be loaded in this endpoint.
     */
    extraJarsS3Path?: pulumi.Input<string>;
    /**
     * Path(s) to one or more Python libraries in an S3 bucket that should be loaded in this endpoint. Multiple values must be complete paths separated by a comma.
     */
    extraPythonLibsS3Path?: pulumi.Input<string>;
    /**
     * Specifies the versions of Python and Apache Spark to use. Defaults to AWS Glue version 0.9.
     */
    glueVersion?: pulumi.Input<string>;
    /**
     * The name of this endpoint. It must be unique in your account.
     */
    name?: pulumi.Input<string>;
    /**
     * The number of AWS Glue Data Processing Units (DPUs) to allocate to this endpoint. Conflicts with `workerType`.
     */
    numberOfNodes?: pulumi.Input<number>;
    /**
     * The number of workers of a defined worker type that are allocated to this endpoint. This field is available only when you choose worker type G.1X or G.2X.
     */
    numberOfWorkers?: pulumi.Input<number>;
    /**
     * The public key to be used by this endpoint for authentication.
     */
    publicKey?: pulumi.Input<string>;
    /**
     * A list of public keys to be used by this endpoint for authentication.
     */
    publicKeys?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The IAM role for this endpoint.
     */
    roleArn: pulumi.Input<string>;
    /**
     * The name of the Security Configuration structure to be used with this endpoint.
     */
    securityConfiguration?: pulumi.Input<string>;
    /**
     * Security group IDs for the security groups to be used by this endpoint.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The subnet ID for the new endpoint to use.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of predefined worker that is allocated to this endpoint. Accepts a value of Standard, G.1X, or G.2X.
     */
    workerType?: pulumi.Input<string>;
}
