// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Manages an Image Builder Infrastructure Configuration.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.imagebuilder.InfrastructureConfiguration("example", {
 *     description: "example description",
 *     instanceProfileName: aws_iam_instance_profile.example.name,
 *     instanceTypes: [
 *         "t2.nano",
 *         "t3.micro",
 *     ],
 *     keyPair: aws_key_pair.example.key_name,
 *     securityGroupIds: [aws_security_group.example.id],
 *     snsTopicArn: aws_sns_topic.example.arn,
 *     subnetId: aws_subnet.main.id,
 *     terminateInstanceOnFailure: true,
 *     logging: {
 *         s3Logs: {
 *             s3BucketName: aws_s3_bucket.example.bucket,
 *             s3KeyPrefix: "logs",
 *         },
 *     },
 *     tags: {
 *         foo: "bar",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Using `pulumi import`, import `aws_imagebuilder_infrastructure_configuration` using the Amazon Resource Name (ARN). For example:
 *
 * ```sh
 *  $ pulumi import aws:imagebuilder/infrastructureConfiguration:InfrastructureConfiguration example arn:aws:imagebuilder:us-east-1:123456789012:infrastructure-configuration/example
 * ```
 */
export class InfrastructureConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing InfrastructureConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InfrastructureConfigurationState, opts?: pulumi.CustomResourceOptions): InfrastructureConfiguration {
        return new InfrastructureConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:imagebuilder/infrastructureConfiguration:InfrastructureConfiguration';

    /**
     * Returns true if the given object is an instance of InfrastructureConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InfrastructureConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InfrastructureConfiguration.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of the configuration.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Date when the configuration was created.
     */
    public /*out*/ readonly dateCreated!: pulumi.Output<string>;
    /**
     * Date when the configuration was updated.
     */
    public /*out*/ readonly dateUpdated!: pulumi.Output<string>;
    /**
     * Description for the configuration.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Configuration block with instance metadata options for the HTTP requests that pipeline builds use to launch EC2 build and test instances. Detailed below.
     */
    public readonly instanceMetadataOptions!: pulumi.Output<outputs.imagebuilder.InfrastructureConfigurationInstanceMetadataOptions | undefined>;
    /**
     * Name of IAM Instance Profile.
     */
    public readonly instanceProfileName!: pulumi.Output<string>;
    /**
     * Set of EC2 Instance Types.
     */
    public readonly instanceTypes!: pulumi.Output<string[] | undefined>;
    /**
     * Name of EC2 Key Pair.
     */
    public readonly keyPair!: pulumi.Output<string | undefined>;
    /**
     * Configuration block with logging settings. Detailed below.
     */
    public readonly logging!: pulumi.Output<outputs.imagebuilder.InfrastructureConfigurationLogging | undefined>;
    /**
     * Name for the configuration.
     *
     * The following arguments are optional:
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Key-value map of resource tags to assign to infrastructure created by the configuration.
     */
    public readonly resourceTags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Set of EC2 Security Group identifiers.
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * Amazon Resource Name (ARN) of SNS Topic.
     */
    public readonly snsTopicArn!: pulumi.Output<string | undefined>;
    /**
     * EC2 Subnet identifier. Also requires `securityGroupIds` argument.
     */
    public readonly subnetId!: pulumi.Output<string | undefined>;
    /**
     * Key-value map of resource tags to assign to the configuration. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Enable if the instance should be terminated when the pipeline fails. Defaults to `false`.
     */
    public readonly terminateInstanceOnFailure!: pulumi.Output<boolean | undefined>;

    /**
     * Create a InfrastructureConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InfrastructureConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InfrastructureConfigurationArgs | InfrastructureConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InfrastructureConfigurationState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["dateCreated"] = state ? state.dateCreated : undefined;
            resourceInputs["dateUpdated"] = state ? state.dateUpdated : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["instanceMetadataOptions"] = state ? state.instanceMetadataOptions : undefined;
            resourceInputs["instanceProfileName"] = state ? state.instanceProfileName : undefined;
            resourceInputs["instanceTypes"] = state ? state.instanceTypes : undefined;
            resourceInputs["keyPair"] = state ? state.keyPair : undefined;
            resourceInputs["logging"] = state ? state.logging : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["resourceTags"] = state ? state.resourceTags : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["snsTopicArn"] = state ? state.snsTopicArn : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["terminateInstanceOnFailure"] = state ? state.terminateInstanceOnFailure : undefined;
        } else {
            const args = argsOrState as InfrastructureConfigurationArgs | undefined;
            if ((!args || args.instanceProfileName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceProfileName'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["instanceMetadataOptions"] = args ? args.instanceMetadataOptions : undefined;
            resourceInputs["instanceProfileName"] = args ? args.instanceProfileName : undefined;
            resourceInputs["instanceTypes"] = args ? args.instanceTypes : undefined;
            resourceInputs["keyPair"] = args ? args.keyPair : undefined;
            resourceInputs["logging"] = args ? args.logging : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceTags"] = args ? args.resourceTags : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["snsTopicArn"] = args ? args.snsTopicArn : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["terminateInstanceOnFailure"] = args ? args.terminateInstanceOnFailure : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["dateCreated"] = undefined /*out*/;
            resourceInputs["dateUpdated"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(InfrastructureConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InfrastructureConfiguration resources.
 */
export interface InfrastructureConfigurationState {
    /**
     * Amazon Resource Name (ARN) of the configuration.
     */
    arn?: pulumi.Input<string>;
    /**
     * Date when the configuration was created.
     */
    dateCreated?: pulumi.Input<string>;
    /**
     * Date when the configuration was updated.
     */
    dateUpdated?: pulumi.Input<string>;
    /**
     * Description for the configuration.
     */
    description?: pulumi.Input<string>;
    /**
     * Configuration block with instance metadata options for the HTTP requests that pipeline builds use to launch EC2 build and test instances. Detailed below.
     */
    instanceMetadataOptions?: pulumi.Input<inputs.imagebuilder.InfrastructureConfigurationInstanceMetadataOptions>;
    /**
     * Name of IAM Instance Profile.
     */
    instanceProfileName?: pulumi.Input<string>;
    /**
     * Set of EC2 Instance Types.
     */
    instanceTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of EC2 Key Pair.
     */
    keyPair?: pulumi.Input<string>;
    /**
     * Configuration block with logging settings. Detailed below.
     */
    logging?: pulumi.Input<inputs.imagebuilder.InfrastructureConfigurationLogging>;
    /**
     * Name for the configuration.
     *
     * The following arguments are optional:
     */
    name?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags to assign to infrastructure created by the configuration.
     */
    resourceTags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Set of EC2 Security Group identifiers.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Amazon Resource Name (ARN) of SNS Topic.
     */
    snsTopicArn?: pulumi.Input<string>;
    /**
     * EC2 Subnet identifier. Also requires `securityGroupIds` argument.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags to assign to the configuration. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Enable if the instance should be terminated when the pipeline fails. Defaults to `false`.
     */
    terminateInstanceOnFailure?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a InfrastructureConfiguration resource.
 */
export interface InfrastructureConfigurationArgs {
    /**
     * Description for the configuration.
     */
    description?: pulumi.Input<string>;
    /**
     * Configuration block with instance metadata options for the HTTP requests that pipeline builds use to launch EC2 build and test instances. Detailed below.
     */
    instanceMetadataOptions?: pulumi.Input<inputs.imagebuilder.InfrastructureConfigurationInstanceMetadataOptions>;
    /**
     * Name of IAM Instance Profile.
     */
    instanceProfileName: pulumi.Input<string>;
    /**
     * Set of EC2 Instance Types.
     */
    instanceTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of EC2 Key Pair.
     */
    keyPair?: pulumi.Input<string>;
    /**
     * Configuration block with logging settings. Detailed below.
     */
    logging?: pulumi.Input<inputs.imagebuilder.InfrastructureConfigurationLogging>;
    /**
     * Name for the configuration.
     *
     * The following arguments are optional:
     */
    name?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags to assign to infrastructure created by the configuration.
     */
    resourceTags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Set of EC2 Security Group identifiers.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Amazon Resource Name (ARN) of SNS Topic.
     */
    snsTopicArn?: pulumi.Input<string>;
    /**
     * EC2 Subnet identifier. Also requires `securityGroupIds` argument.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags to assign to the configuration. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Enable if the instance should be terminated when the pipeline fails. Defaults to `false`.
     */
    terminateInstanceOnFailure?: pulumi.Input<boolean>;
}
