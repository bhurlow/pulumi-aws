// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource for managing an AWS Managed Streaming for Kafka VPC Connection.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.msk.VpcConnection("test", {
 *     authentication: "SASL_IAM",
 *     targetClusterArn: "aws_msk_cluster.arn",
 *     vpcId: aws_vpc.test.id,
 *     clientSubnets: aws_subnet.test.map(__item => __item.id),
 *     securityGroups: [aws_security_group.test.id],
 * });
 * ```
 *
 * ## Import
 *
 * In TODO v1.5.0 and later, use an `import` block to import MSK configurations using the configuration ARN. For exampleterraform import {
 *
 *  to = aws_msk_vpc_connection.example
 *
 *  id = "arn:aws:kafka:eu-west-2:123456789012:vpc-connection/123456789012/example/38173259-79cd-4ee8-87f3-682ea6023f48-2" } Using `TODO import`, import MSK configurations using the configuration ARN. For exampleconsole % TODO import aws_msk_vpc_connection.example arn:aws:kafka:eu-west-2:123456789012:vpc-connection/123456789012/example/38173259-79cd-4ee8-87f3-682ea6023f48-2
 */
export class VpcConnection extends pulumi.CustomResource {
    /**
     * Get an existing VpcConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcConnectionState, opts?: pulumi.CustomResourceOptions): VpcConnection {
        return new VpcConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:msk/vpcConnection:VpcConnection';

    /**
     * Returns true if the given object is an instance of VpcConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcConnection.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of the VPC connection.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The authentication type for the client VPC connection. Specify one of these auth type strings: SASL_IAM, SASL_SCRAM, or TLS.
     */
    public readonly authentication!: pulumi.Output<string>;
    /**
     * The list of subnets in the client VPC to connect to.
     */
    public readonly clientSubnets!: pulumi.Output<string[]>;
    /**
     * The security groups to attach to the ENIs for the broker nodes.
     */
    public readonly securityGroups!: pulumi.Output<string[]>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The Amazon Resource Name (ARN) of the cluster.
     */
    public readonly targetClusterArn!: pulumi.Output<string>;
    /**
     * The VPC ID of the remote client.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a VpcConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcConnectionArgs | VpcConnectionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcConnectionState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["authentication"] = state ? state.authentication : undefined;
            resourceInputs["clientSubnets"] = state ? state.clientSubnets : undefined;
            resourceInputs["securityGroups"] = state ? state.securityGroups : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["targetClusterArn"] = state ? state.targetClusterArn : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as VpcConnectionArgs | undefined;
            if ((!args || args.authentication === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authentication'");
            }
            if ((!args || args.clientSubnets === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientSubnets'");
            }
            if ((!args || args.securityGroups === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroups'");
            }
            if ((!args || args.targetClusterArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetClusterArn'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["authentication"] = args ? args.authentication : undefined;
            resourceInputs["clientSubnets"] = args ? args.clientSubnets : undefined;
            resourceInputs["securityGroups"] = args ? args.securityGroups : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["targetClusterArn"] = args ? args.targetClusterArn : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(VpcConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcConnection resources.
 */
export interface VpcConnectionState {
    /**
     * Amazon Resource Name (ARN) of the VPC connection.
     */
    arn?: pulumi.Input<string>;
    /**
     * The authentication type for the client VPC connection. Specify one of these auth type strings: SASL_IAM, SASL_SCRAM, or TLS.
     */
    authentication?: pulumi.Input<string>;
    /**
     * The list of subnets in the client VPC to connect to.
     */
    clientSubnets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The security groups to attach to the ENIs for the broker nodes.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     *
     * @deprecated Please use `tags` instead.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Amazon Resource Name (ARN) of the cluster.
     */
    targetClusterArn?: pulumi.Input<string>;
    /**
     * The VPC ID of the remote client.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcConnection resource.
 */
export interface VpcConnectionArgs {
    /**
     * The authentication type for the client VPC connection. Specify one of these auth type strings: SASL_IAM, SASL_SCRAM, or TLS.
     */
    authentication: pulumi.Input<string>;
    /**
     * The list of subnets in the client VPC to connect to.
     */
    clientSubnets: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The security groups to attach to the ENIs for the broker nodes.
     */
    securityGroups: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Amazon Resource Name (ARN) of the cluster.
     */
    targetClusterArn: pulumi.Input<string>;
    /**
     * The VPC ID of the remote client.
     */
    vpcId: pulumi.Input<string>;
}
