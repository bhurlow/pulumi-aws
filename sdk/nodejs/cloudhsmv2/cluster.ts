// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Creates an Amazon CloudHSM v2 cluster.
 *
 * For information about CloudHSM v2, see the
 * [AWS CloudHSM User Guide](https://docs.aws.amazon.com/cloudhsm/latest/userguide/introduction.html) and the [Amazon
 * CloudHSM API Reference][2].
 *
 * > **NOTE:** A CloudHSM Cluster can take several minutes to set up.
 * Practically no single attribute can be updated, except for `tags`.
 * If you need to delete a cluster, you have to remove its HSM modules first.
 * To initialize cluster, you have to add an HSM instance to the cluster, then sign CSR and upload it.
 *
 * ## Import
 *
 * terraform import {
 *
 *  to = aws_cloudhsm_v2_cluster.test_cluster
 *
 *  id = "cluster-aeb282a201" } Using `pulumi import`, import CloudHSM v2 Clusters using the cluster `id`. For exampleconsole % pulumi import aws_cloudhsm_v2_cluster.test_cluster cluster-aeb282a201
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloudhsmv2/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * The list of cluster certificates.
     * * `cluster_certificates.0.cluster_certificate` - The cluster certificate issued (signed) by the issuing certificate authority (CA) of the cluster's owner.
     * * `cluster_certificates.0.cluster_csr` - The certificate signing request (CSR). Available only in `UNINITIALIZED` state after an HSM instance is added to the cluster.
     * * `cluster_certificates.0.aws_hardware_certificate` - The HSM hardware certificate issued (signed) by AWS CloudHSM.
     * * `cluster_certificates.0.hsm_certificate` - The HSM certificate issued (signed) by the HSM hardware.
     * * `cluster_certificates.0.manufacturer_hardware_certificate` - The HSM hardware certificate issued (signed) by the hardware manufacturer.
     */
    public /*out*/ readonly clusterCertificates!: pulumi.Output<outputs.cloudhsmv2.ClusterClusterCertificate[]>;
    /**
     * The id of the CloudHSM cluster.
     */
    public /*out*/ readonly clusterId!: pulumi.Output<string>;
    /**
     * The state of the CloudHSM cluster.
     */
    public /*out*/ readonly clusterState!: pulumi.Output<string>;
    /**
     * The type of HSM module in the cluster. Currently, only `hsm1.medium` is supported.
     */
    public readonly hsmType!: pulumi.Output<string>;
    /**
     * The ID of the security group associated with the CloudHSM cluster.
     */
    public /*out*/ readonly securityGroupId!: pulumi.Output<string>;
    /**
     * ID of Cloud HSM v2 cluster backup to be restored.
     */
    public readonly sourceBackupIdentifier!: pulumi.Output<string | undefined>;
    /**
     * The IDs of subnets in which cluster will operate.
     */
    public readonly subnetIds!: pulumi.Output<string[]>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The id of the VPC that the CloudHSM cluster resides in.
     */
    public /*out*/ readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterState | undefined;
            resourceInputs["clusterCertificates"] = state ? state.clusterCertificates : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["clusterState"] = state ? state.clusterState : undefined;
            resourceInputs["hsmType"] = state ? state.hsmType : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            resourceInputs["sourceBackupIdentifier"] = state ? state.sourceBackupIdentifier : undefined;
            resourceInputs["subnetIds"] = state ? state.subnetIds : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if ((!args || args.hsmType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hsmType'");
            }
            if ((!args || args.subnetIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetIds'");
            }
            resourceInputs["hsmType"] = args ? args.hsmType : undefined;
            resourceInputs["sourceBackupIdentifier"] = args ? args.sourceBackupIdentifier : undefined;
            resourceInputs["subnetIds"] = args ? args.subnetIds : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["clusterCertificates"] = undefined /*out*/;
            resourceInputs["clusterId"] = undefined /*out*/;
            resourceInputs["clusterState"] = undefined /*out*/;
            resourceInputs["securityGroupId"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * The list of cluster certificates.
     * * `cluster_certificates.0.cluster_certificate` - The cluster certificate issued (signed) by the issuing certificate authority (CA) of the cluster's owner.
     * * `cluster_certificates.0.cluster_csr` - The certificate signing request (CSR). Available only in `UNINITIALIZED` state after an HSM instance is added to the cluster.
     * * `cluster_certificates.0.aws_hardware_certificate` - The HSM hardware certificate issued (signed) by AWS CloudHSM.
     * * `cluster_certificates.0.hsm_certificate` - The HSM certificate issued (signed) by the HSM hardware.
     * * `cluster_certificates.0.manufacturer_hardware_certificate` - The HSM hardware certificate issued (signed) by the hardware manufacturer.
     */
    clusterCertificates?: pulumi.Input<pulumi.Input<inputs.cloudhsmv2.ClusterClusterCertificate>[]>;
    /**
     * The id of the CloudHSM cluster.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The state of the CloudHSM cluster.
     */
    clusterState?: pulumi.Input<string>;
    /**
     * The type of HSM module in the cluster. Currently, only `hsm1.medium` is supported.
     */
    hsmType?: pulumi.Input<string>;
    /**
     * The ID of the security group associated with the CloudHSM cluster.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * ID of Cloud HSM v2 cluster backup to be restored.
     */
    sourceBackupIdentifier?: pulumi.Input<string>;
    /**
     * The IDs of subnets in which cluster will operate.
     */
    subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `defaultTags` configuration block.
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The id of the VPC that the CloudHSM cluster resides in.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * The type of HSM module in the cluster. Currently, only `hsm1.medium` is supported.
     */
    hsmType: pulumi.Input<string>;
    /**
     * ID of Cloud HSM v2 cluster backup to be restored.
     */
    sourceBackupIdentifier?: pulumi.Input<string>;
    /**
     * The IDs of subnets in which cluster will operate.
     */
    subnetIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
