// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Creates a new Amazon Redshift Serverless Workgroup.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.redshiftserverless.Workgroup("example", {
 *     namespaceName: "concurrency-scaling",
 *     workgroupName: "concurrency-scaling",
 * });
 * ```
 *
 * ## Import
 *
 * Using `pulumi import`, import Redshift Serverless Workgroups using the `workgroup_name`. For example:
 *
 * ```sh
 *  $ pulumi import aws:redshiftserverless/workgroup:Workgroup example example
 * ```
 */
export class Workgroup extends pulumi.CustomResource {
    /**
     * Get an existing Workgroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WorkgroupState, opts?: pulumi.CustomResourceOptions): Workgroup {
        return new Workgroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:redshiftserverless/workgroup:Workgroup';

    /**
     * Returns true if the given object is an instance of Workgroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Workgroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Workgroup.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of the Redshift Serverless Workgroup.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The base data warehouse capacity of the workgroup in Redshift Processing Units (RPUs).
     */
    public readonly baseCapacity!: pulumi.Output<number>;
    /**
     * An array of parameters to set for more control over a serverless database. See `Config Parameter` below.
     */
    public readonly configParameters!: pulumi.Output<outputs.redshiftserverless.WorkgroupConfigParameter[]>;
    /**
     * The endpoint that is created from the workgroup. See `Endpoint` below.
     */
    public /*out*/ readonly endpoints!: pulumi.Output<outputs.redshiftserverless.WorkgroupEndpoint[]>;
    /**
     * The value that specifies whether to turn on enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC instead of over the internet.
     */
    public readonly enhancedVpcRouting!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the namespace.
     */
    public readonly namespaceName!: pulumi.Output<string>;
    /**
     * A value that specifies whether the workgroup can be accessed from a public network.
     */
    public readonly publiclyAccessible!: pulumi.Output<boolean | undefined>;
    /**
     * An array of security group IDs to associate with the workgroup.
     */
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    /**
     * An array of VPC subnet IDs to associate with the workgroup. When set, must contain at least three subnets spanning three Availability Zones. A minimum number of IP addresses is required and scales with the Base Capacity. For more information, see the following [AWS document](https://docs.aws.amazon.com/redshift/latest/mgmt/serverless-known-issues.html).
     */
    public readonly subnetIds!: pulumi.Output<string[]>;
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
     * The Redshift Workgroup ID.
     */
    public /*out*/ readonly workgroupId!: pulumi.Output<string>;
    /**
     * The name of the workgroup.
     *
     * The following arguments are optional:
     */
    public readonly workgroupName!: pulumi.Output<string>;

    /**
     * Create a Workgroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkgroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkgroupArgs | WorkgroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WorkgroupState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["baseCapacity"] = state ? state.baseCapacity : undefined;
            resourceInputs["configParameters"] = state ? state.configParameters : undefined;
            resourceInputs["endpoints"] = state ? state.endpoints : undefined;
            resourceInputs["enhancedVpcRouting"] = state ? state.enhancedVpcRouting : undefined;
            resourceInputs["namespaceName"] = state ? state.namespaceName : undefined;
            resourceInputs["publiclyAccessible"] = state ? state.publiclyAccessible : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["subnetIds"] = state ? state.subnetIds : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["workgroupId"] = state ? state.workgroupId : undefined;
            resourceInputs["workgroupName"] = state ? state.workgroupName : undefined;
        } else {
            const args = argsOrState as WorkgroupArgs | undefined;
            if ((!args || args.namespaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if ((!args || args.workgroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workgroupName'");
            }
            resourceInputs["baseCapacity"] = args ? args.baseCapacity : undefined;
            resourceInputs["configParameters"] = args ? args.configParameters : undefined;
            resourceInputs["enhancedVpcRouting"] = args ? args.enhancedVpcRouting : undefined;
            resourceInputs["namespaceName"] = args ? args.namespaceName : undefined;
            resourceInputs["publiclyAccessible"] = args ? args.publiclyAccessible : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["subnetIds"] = args ? args.subnetIds : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["workgroupName"] = args ? args.workgroupName : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["endpoints"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["workgroupId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["tagsAll"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Workgroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Workgroup resources.
 */
export interface WorkgroupState {
    /**
     * Amazon Resource Name (ARN) of the Redshift Serverless Workgroup.
     */
    arn?: pulumi.Input<string>;
    /**
     * The base data warehouse capacity of the workgroup in Redshift Processing Units (RPUs).
     */
    baseCapacity?: pulumi.Input<number>;
    /**
     * An array of parameters to set for more control over a serverless database. See `Config Parameter` below.
     */
    configParameters?: pulumi.Input<pulumi.Input<inputs.redshiftserverless.WorkgroupConfigParameter>[]>;
    /**
     * The endpoint that is created from the workgroup. See `Endpoint` below.
     */
    endpoints?: pulumi.Input<pulumi.Input<inputs.redshiftserverless.WorkgroupEndpoint>[]>;
    /**
     * The value that specifies whether to turn on enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC instead of over the internet.
     */
    enhancedVpcRouting?: pulumi.Input<boolean>;
    /**
     * The name of the namespace.
     */
    namespaceName?: pulumi.Input<string>;
    /**
     * A value that specifies whether the workgroup can be accessed from a public network.
     */
    publiclyAccessible?: pulumi.Input<boolean>;
    /**
     * An array of security group IDs to associate with the workgroup.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An array of VPC subnet IDs to associate with the workgroup. When set, must contain at least three subnets spanning three Availability Zones. A minimum number of IP addresses is required and scales with the Base Capacity. For more information, see the following [AWS document](https://docs.aws.amazon.com/redshift/latest/mgmt/serverless-known-issues.html).
     */
    subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
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
     * The Redshift Workgroup ID.
     */
    workgroupId?: pulumi.Input<string>;
    /**
     * The name of the workgroup.
     *
     * The following arguments are optional:
     */
    workgroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Workgroup resource.
 */
export interface WorkgroupArgs {
    /**
     * The base data warehouse capacity of the workgroup in Redshift Processing Units (RPUs).
     */
    baseCapacity?: pulumi.Input<number>;
    /**
     * An array of parameters to set for more control over a serverless database. See `Config Parameter` below.
     */
    configParameters?: pulumi.Input<pulumi.Input<inputs.redshiftserverless.WorkgroupConfigParameter>[]>;
    /**
     * The value that specifies whether to turn on enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC instead of over the internet.
     */
    enhancedVpcRouting?: pulumi.Input<boolean>;
    /**
     * The name of the namespace.
     */
    namespaceName: pulumi.Input<string>;
    /**
     * A value that specifies whether the workgroup can be accessed from a public network.
     */
    publiclyAccessible?: pulumi.Input<boolean>;
    /**
     * An array of security group IDs to associate with the workgroup.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An array of VPC subnet IDs to associate with the workgroup. When set, must contain at least three subnets spanning three Availability Zones. A minimum number of IP addresses is required and scales with the Base Capacity. For more information, see the following [AWS document](https://docs.aws.amazon.com/redshift/latest/mgmt/serverless-known-issues.html).
     */
    subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the workgroup.
     *
     * The following arguments are optional:
     */
    workgroupName: pulumi.Input<string>;
}
