// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

import {RoutingRule} from "@/s3";

export interface ClusterCertificateAuthority {
    /**
     * Base64 encoded certificate data required to communicate with your cluster. Add this to the `certificate-authority-data` section of the `kubeconfig` file for your cluster.
     */
    data?: pulumi.Input<string>;
}

export interface ClusterEncryptionConfig {
    /**
     * Configuration block with provider for encryption. Detailed below.
     */
    provider: pulumi.Input<inputs.eks.ClusterEncryptionConfigProvider>;
    /**
     * List of strings with resources to be encrypted. Valid values: `secrets`.
     */
    resources: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ClusterEncryptionConfigProvider {
    /**
     * ARN of the Key Management Service (KMS) customer master key (CMK). The CMK must be symmetric, created in the same region as the cluster, and if the CMK was created in a different account, the user must have access to the CMK. For more information, see [Allowing Users in Other Accounts to Use a CMK in the AWS Key Management Service Developer Guide](https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-modifying-external-accounts.html).
     */
    keyArn: pulumi.Input<string>;
}

export interface ClusterIdentity {
    /**
     * Nested block containing [OpenID Connect](https://openid.net/connect/) identity provider information for the cluster. Detailed below.
     */
    oidcs?: pulumi.Input<pulumi.Input<inputs.eks.ClusterIdentityOidc>[]>;
}

export interface ClusterIdentityOidc {
    /**
     * Issuer URL for the OpenID Connect identity provider.
     */
    issuer?: pulumi.Input<string>;
}

export interface ClusterKubernetesNetworkConfig {
    /**
     * The IP family used to assign Kubernetes pod and service addresses. Valid values are `ipv4` (default) and `ipv6`. You can only specify an IP family when you create a cluster, changing this value will force a new cluster to be created.
     */
    ipFamily?: pulumi.Input<string>;
    /**
     * The CIDR block to assign Kubernetes pod and service IP addresses from. If you don't specify a block, Kubernetes assigns addresses from either the 10.100.0.0/16 or 172.20.0.0/16 CIDR blocks. We recommend that you specify a block that does not overlap with resources in other networks that are peered or connected to your VPC. You can only specify a custom CIDR block when you create a cluster, changing this value will force a new cluster to be created. The block must meet the following requirements:
     */
    serviceIpv4Cidr?: pulumi.Input<string>;
    serviceIpv6Cidr?: pulumi.Input<string>;
}

export interface ClusterOutpostConfig {
    /**
     * The Amazon EC2 instance type that you want to use for your local Amazon EKS cluster on Outposts. The instance type that you specify is used for all Kubernetes control plane instances. The instance type can't be changed after cluster creation. Choose an instance type based on the number of nodes that your cluster will have. If your cluster will have:
     */
    controlPlaneInstanceType: pulumi.Input<string>;
    /**
     * An object representing the placement configuration for all the control plane instances of your local Amazon EKS cluster on AWS Outpost.
     * The following arguments are supported in the `controlPlanePlacement` configuration block:
     */
    controlPlanePlacement?: pulumi.Input<inputs.eks.ClusterOutpostConfigControlPlanePlacement>;
    /**
     * The ARN of the Outpost that you want to use for your local Amazon EKS cluster on Outposts. This argument is a list of arns, but only a single Outpost ARN is supported currently.
     */
    outpostArns: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ClusterOutpostConfigControlPlanePlacement {
    /**
     * The name of the placement group for the Kubernetes control plane instances. This setting can't be changed after cluster creation.
     */
    groupName: pulumi.Input<string>;
}

export interface ClusterVpcConfig {
    /**
     * Cluster security group that was created by Amazon EKS for the cluster. Managed node groups use this security group for control-plane-to-data-plane communication.
     */
    clusterSecurityGroupId?: pulumi.Input<string>;
    /**
     * Whether the Amazon EKS private API server endpoint is enabled. Default is `false`.
     */
    endpointPrivateAccess?: pulumi.Input<boolean>;
    /**
     * Whether the Amazon EKS public API server endpoint is enabled. Default is `true`.
     */
    endpointPublicAccess?: pulumi.Input<boolean>;
    /**
     * List of CIDR blocks. Indicates which CIDR blocks can access the Amazon EKS public API server endpoint when enabled. EKS defaults this to a list with `0.0.0.0/0`. The provider will only perform drift detection of its value when present in a configuration.
     */
    publicAccessCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of security group IDs for the cross-account elastic network interfaces that Amazon EKS creates to use to allow communication between your worker nodes and the Kubernetes control plane.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of subnet IDs. Must be in at least two different availability zones. Amazon EKS creates cross-account elastic network interfaces in these subnets to allow communication between your worker nodes and the Kubernetes control plane.
     */
    subnetIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of the VPC associated with your cluster.
     */
    vpcId?: pulumi.Input<string>;
}

export interface FargateProfileSelector {
    /**
     * Key-value map of Kubernetes labels for selection.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Kubernetes namespace for selection.
     */
    namespace: pulumi.Input<string>;
}

export interface IdentityProviderConfigOidc {
    /**
     * Client ID for the OpenID Connect identity provider.
     */
    clientId: pulumi.Input<string>;
    /**
     * The JWT claim that the provider will use to return groups.
     */
    groupsClaim?: pulumi.Input<string>;
    /**
     * A prefix that is prepended to group claims e.g., `oidc:`.
     */
    groupsPrefix?: pulumi.Input<string>;
    /**
     * The name of the identity provider config.
     */
    identityProviderConfigName: pulumi.Input<string>;
    /**
     * Issuer URL for the OpenID Connect identity provider.
     */
    issuerUrl: pulumi.Input<string>;
    /**
     * The key value pairs that describe required claims in the identity token.
     */
    requiredClaims?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The JWT claim that the provider will use as the username.
     */
    usernameClaim?: pulumi.Input<string>;
    /**
     * A prefix that is prepended to username claims.
     */
    usernamePrefix?: pulumi.Input<string>;
}

export interface NodeGroupLaunchTemplate {
    /**
     * Identifier of the EC2 Launch Template. Conflicts with `name`.
     */
    id?: pulumi.Input<string>;
    /**
     * Name of the EC2 Launch Template. Conflicts with `id`.
     */
    name?: pulumi.Input<string>;
    /**
     * EC2 Launch Template version number. While the API accepts values like `$Default` and `$Latest`, the API will convert the value to the associated version number (e.g., `1`) on read and the provider will show a difference on next plan. Using the `defaultVersion` or `latestVersion` attribute of the `aws.ec2.LaunchTemplate` resource or data source is recommended for this argument.
     */
    version: pulumi.Input<string>;
}

export interface NodeGroupRemoteAccess {
    /**
     * EC2 Key Pair name that provides access for SSH communication with the worker nodes in the EKS Node Group. If you specify this configuration, but do not specify `sourceSecurityGroupIds` when you create an EKS Node Group, port 22 on the worker nodes is opened to the Internet (0.0.0.0/0).
     */
    ec2SshKey?: pulumi.Input<string>;
    /**
     * Set of EC2 Security Group IDs to allow SSH access (port 22) from on the worker nodes. If you specify `ec2SshKey`, but do not specify this configuration when you create an EKS Node Group, port 22 on the worker nodes is opened to the Internet (0.0.0.0/0).
     */
    sourceSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface NodeGroupResource {
    /**
     * List of objects containing information about AutoScaling Groups.
     */
    autoscalingGroups?: pulumi.Input<pulumi.Input<inputs.eks.NodeGroupResourceAutoscalingGroup>[]>;
    /**
     * Identifier of the remote access EC2 Security Group.
     */
    remoteAccessSecurityGroupId?: pulumi.Input<string>;
}

export interface NodeGroupResourceAutoscalingGroup {
    /**
     * Name of the EC2 Launch Template. Conflicts with `id`.
     */
    name?: pulumi.Input<string>;
}

export interface NodeGroupScalingConfig {
    /**
     * Desired number of worker nodes.
     */
    desiredSize: pulumi.Input<number>;
    /**
     * Maximum number of worker nodes.
     */
    maxSize: pulumi.Input<number>;
    /**
     * Minimum number of worker nodes.
     */
    minSize: pulumi.Input<number>;
}

export interface NodeGroupTaint {
    /**
     * The effect of the taint. Valid values: `NO_SCHEDULE`, `NO_EXECUTE`, `PREFER_NO_SCHEDULE`.
     */
    effect: pulumi.Input<string>;
    /**
     * The key of the taint. Maximum length of 63.
     */
    key: pulumi.Input<string>;
    /**
     * The value of the taint. Maximum length of 63.
     */
    value?: pulumi.Input<string>;
}

export interface NodeGroupUpdateConfig {
    /**
     * Desired max number of unavailable worker nodes during node group update.
     */
    maxUnavailable?: pulumi.Input<number>;
    /**
     * Desired max percentage of unavailable worker nodes during node group update.
     */
    maxUnavailablePercentage?: pulumi.Input<number>;
}
