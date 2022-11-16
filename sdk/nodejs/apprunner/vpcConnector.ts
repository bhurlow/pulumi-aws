// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an App Runner VPC Connector.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const connector = new aws.apprunner.VpcConnector("connector", {
 *     securityGroups: [
 *         "sg1",
 *         "sg2",
 *     ],
 *     subnets: [
 *         "subnet1",
 *         "subnet2",
 *     ],
 *     vpcConnectorName: "name",
 * });
 * ```
 *
 * ## Import
 *
 * App Runner vpc connector can be imported by using the `arn`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:apprunner/vpcConnector:VpcConnector example arn:aws:apprunner:us-east-1:1234567890:vpcconnector/example/1/0a03292a89764e5882c41d8f991c82fe
 * ```
 */
export class VpcConnector extends pulumi.CustomResource {
    /**
     * Get an existing VpcConnector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcConnectorState, opts?: pulumi.CustomResourceOptions): VpcConnector {
        return new VpcConnector(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apprunner/vpcConnector:VpcConnector';

    /**
     * Returns true if the given object is an instance of VpcConnector.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcConnector {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcConnector.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * List of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets. If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.
     */
    public readonly securityGroups!: pulumi.Output<string[]>;
    /**
     * Current state of the VPC connector. If the status of a connector revision is INACTIVE, it was deleted and can't be used. Inactive connector revisions are permanently removed some time after they are deleted.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * List of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC. Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.
     */
    public readonly subnets!: pulumi.Output<string[]>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string}>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Name for the VPC connector.
     */
    public readonly vpcConnectorName!: pulumi.Output<string>;
    /**
     * The revision of VPC connector. It's unique among all the active connectors ("Status": "ACTIVE") that share the same Name.
     */
    public /*out*/ readonly vpcConnectorRevision!: pulumi.Output<number>;

    /**
     * Create a VpcConnector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcConnectorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcConnectorArgs | VpcConnectorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcConnectorState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["securityGroups"] = state ? state.securityGroups : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subnets"] = state ? state.subnets : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["vpcConnectorName"] = state ? state.vpcConnectorName : undefined;
            resourceInputs["vpcConnectorRevision"] = state ? state.vpcConnectorRevision : undefined;
        } else {
            const args = argsOrState as VpcConnectorArgs | undefined;
            if ((!args || args.securityGroups === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroups'");
            }
            if ((!args || args.subnets === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnets'");
            }
            if ((!args || args.vpcConnectorName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcConnectorName'");
            }
            resourceInputs["securityGroups"] = args ? args.securityGroups : undefined;
            resourceInputs["subnets"] = args ? args.subnets : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcConnectorName"] = args ? args.vpcConnectorName : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["vpcConnectorRevision"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcConnector.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcConnector resources.
 */
export interface VpcConnectorState {
    arn?: pulumi.Input<string>;
    /**
     * List of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets. If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Current state of the VPC connector. If the status of a connector revision is INACTIVE, it was deleted and can't be used. Inactive connector revisions are permanently removed some time after they are deleted.
     */
    status?: pulumi.Input<string>;
    /**
     * List of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC. Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.
     */
    subnets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name for the VPC connector.
     */
    vpcConnectorName?: pulumi.Input<string>;
    /**
     * The revision of VPC connector. It's unique among all the active connectors ("Status": "ACTIVE") that share the same Name.
     */
    vpcConnectorRevision?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a VpcConnector resource.
 */
export interface VpcConnectorArgs {
    /**
     * List of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets. If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.
     */
    securityGroups: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC. Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.
     */
    subnets: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Key-value map of resource tags. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name for the VPC connector.
     */
    vpcConnectorName: pulumi.Input<string>;
}
