// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an EMR Containers (EMR on EKS) Virtual Cluster.
 *
 * ## Example Usage
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.emrcontainers.VirtualCluster("example", {containerProvider: {
 *     id: aws_eks_cluster.example.name,
 *     type: "EKS",
 *     info: {
 *         eksInfo: {
 *             namespace: "default",
 *         },
 *     },
 * }});
 * ```
 *
 * ## Import
 *
 * EKS Clusters can be imported using the `name`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:emrcontainers/virtualCluster:VirtualCluster example
 * ```
 */
export class VirtualCluster extends pulumi.CustomResource {
    /**
     * Get an existing VirtualCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualClusterState, opts?: pulumi.CustomResourceOptions): VirtualCluster {
        return new VirtualCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:emrcontainers/virtualCluster:VirtualCluster';

    /**
     * Returns true if the given object is an instance of VirtualCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualCluster.__pulumiType;
    }

    /**
     * ARN of the cluster.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Configuration block for the container provider associated with your cluster.
     */
    public readonly containerProvider!: pulumi.Output<outputs.emrcontainers.VirtualClusterContainerProvider>;
    /**
     * Name of the virtual cluster.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Key-value mapping of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a VirtualCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualClusterArgs | VirtualClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VirtualClusterState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["containerProvider"] = state ? state.containerProvider : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as VirtualClusterArgs | undefined;
            if ((!args || args.containerProvider === undefined) && !opts.urn) {
                throw new Error("Missing required property 'containerProvider'");
            }
            resourceInputs["containerProvider"] = args ? args.containerProvider : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["tagsAll"] = args ? args.tagsAll : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VirtualCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualCluster resources.
 */
export interface VirtualClusterState {
    /**
     * ARN of the cluster.
     */
    arn?: pulumi.Input<string>;
    /**
     * Configuration block for the container provider associated with your cluster.
     */
    containerProvider?: pulumi.Input<inputs.emrcontainers.VirtualClusterContainerProvider>;
    /**
     * Name of the virtual cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * Key-value mapping of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a VirtualCluster resource.
 */
export interface VirtualClusterArgs {
    /**
     * Configuration block for the container provider associated with your cluster.
     */
    containerProvider: pulumi.Input<inputs.emrcontainers.VirtualClusterContainerProvider>;
    /**
     * Name of the virtual cluster.
     */
    name?: pulumi.Input<string>;
    /**
     * Key-value mapping of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
