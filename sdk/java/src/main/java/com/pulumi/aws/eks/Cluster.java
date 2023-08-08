// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.eks;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.eks.ClusterArgs;
import com.pulumi.aws.eks.inputs.ClusterState;
import com.pulumi.aws.eks.outputs.ClusterCertificateAuthority;
import com.pulumi.aws.eks.outputs.ClusterEncryptionConfig;
import com.pulumi.aws.eks.outputs.ClusterIdentity;
import com.pulumi.aws.eks.outputs.ClusterKubernetesNetworkConfig;
import com.pulumi.aws.eks.outputs.ClusterOutpostConfig;
import com.pulumi.aws.eks.outputs.ClusterVpcConfig;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an EKS Cluster.
 * 
 * ## Example Usage
 * ### Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.eks.Cluster;
 * import com.pulumi.aws.eks.ClusterArgs;
 * import com.pulumi.aws.eks.inputs.ClusterVpcConfigArgs;
 * import com.pulumi.resources.CustomResourceOptions;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new Cluster(&#34;example&#34;, ClusterArgs.builder()        
 *             .roleArn(aws_iam_role.example().arn())
 *             .vpcConfig(ClusterVpcConfigArgs.builder()
 *                 .subnetIds(                
 *                     aws_subnet.example1().id(),
 *                     aws_subnet.example2().id())
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(                
 *                     aws_iam_role_policy_attachment.example-AmazonEKSClusterPolicy(),
 *                     aws_iam_role_policy_attachment.example-AmazonEKSVPCResourceController())
 *                 .build());
 * 
 *         ctx.export(&#34;endpoint&#34;, example.endpoint());
 *         ctx.export(&#34;kubeconfig-certificate-authority-data&#34;, example.certificateAuthority().applyValue(certificateAuthority -&gt; certificateAuthority.data()));
 *     }
 * }
 * ```
 * ### Example IAM Role for EKS Cluster
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.iam.IamFunctions;
 * import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
 * import com.pulumi.aws.iam.Role;
 * import com.pulumi.aws.iam.RoleArgs;
 * import com.pulumi.aws.iam.RolePolicyAttachment;
 * import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
 *             .statements(GetPolicyDocumentStatementArgs.builder()
 *                 .effect(&#34;Allow&#34;)
 *                 .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
 *                     .type(&#34;Service&#34;)
 *                     .identifiers(&#34;eks.amazonaws.com&#34;)
 *                     .build())
 *                 .actions(&#34;sts:AssumeRole&#34;)
 *                 .build())
 *             .build());
 * 
 *         var example = new Role(&#34;example&#34;, RoleArgs.builder()        
 *             .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -&gt; getPolicyDocumentResult.json()))
 *             .build());
 * 
 *         var example_AmazonEKSClusterPolicy = new RolePolicyAttachment(&#34;example-AmazonEKSClusterPolicy&#34;, RolePolicyAttachmentArgs.builder()        
 *             .policyArn(&#34;arn:aws:iam::aws:policy/AmazonEKSClusterPolicy&#34;)
 *             .role(example.name())
 *             .build());
 * 
 *         var example_AmazonEKSVPCResourceController = new RolePolicyAttachment(&#34;example-AmazonEKSVPCResourceController&#34;, RolePolicyAttachmentArgs.builder()        
 *             .policyArn(&#34;arn:aws:iam::aws:policy/AmazonEKSVPCResourceController&#34;)
 *             .role(example.name())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Enabling Control Plane Logging
 * 
 * [EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html) can be enabled via the `enabled_cluster_log_types` argument. To manage the CloudWatch Log Group retention period, the `aws.cloudwatch.LogGroup` resource can be used.
 * 
 * &gt; The below configuration uses [`dependsOn`](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson) to prevent ordering issues with EKS automatically creating the log group first and a variable for naming consistency. Other ordering and naming methodologies may be more appropriate for your environment.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.cloudwatch.LogGroup;
 * import com.pulumi.aws.cloudwatch.LogGroupArgs;
 * import com.pulumi.aws.eks.Cluster;
 * import com.pulumi.aws.eks.ClusterArgs;
 * import com.pulumi.resources.CustomResourceOptions;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var config = ctx.config();
 *         final var clusterName = config.get(&#34;clusterName&#34;).orElse(&#34;example&#34;);
 *         var exampleLogGroup = new LogGroup(&#34;exampleLogGroup&#34;, LogGroupArgs.builder()        
 *             .retentionInDays(7)
 *             .build());
 * 
 *         var exampleCluster = new Cluster(&#34;exampleCluster&#34;, ClusterArgs.builder()        
 *             .enabledClusterLogTypes(            
 *                 &#34;api&#34;,
 *                 &#34;audit&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(exampleLogGroup)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ### EKS Cluster on AWS Outpost
 * 
 * [Creating a local Amazon EKS cluster on an AWS Outpost](https://docs.aws.amazon.com/eks/latest/userguide/create-cluster-outpost.html)
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.iam.Role;
 * import com.pulumi.aws.iam.RoleArgs;
 * import com.pulumi.aws.eks.Cluster;
 * import com.pulumi.aws.eks.ClusterArgs;
 * import com.pulumi.aws.eks.inputs.ClusterVpcConfigArgs;
 * import com.pulumi.aws.eks.inputs.ClusterOutpostConfigArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var exampleRole = new Role(&#34;exampleRole&#34;, RoleArgs.builder()        
 *             .assumeRolePolicy(data.aws_iam_policy_document().example_assume_role_policy().json())
 *             .build());
 * 
 *         var exampleCluster = new Cluster(&#34;exampleCluster&#34;, ClusterArgs.builder()        
 *             .roleArn(exampleRole.arn())
 *             .vpcConfig(ClusterVpcConfigArgs.builder()
 *                 .endpointPrivateAccess(true)
 *                 .endpointPublicAccess(false)
 *                 .build())
 *             .outpostConfig(ClusterOutpostConfigArgs.builder()
 *                 .controlPlaneInstanceType(&#34;m5d.large&#34;)
 *                 .outpostArns(data.aws_outposts_outpost().example().arn())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * After adding inline IAM Policies (e.g., `aws.iam.RolePolicy` resource) or attaching IAM Policies (e.g., `aws.iam.Policy` resource and `aws.iam.RolePolicyAttachment` resource) with the desired permissions to the IAM Role, annotate the Kubernetes service account (e.g., `kubernetes_service_account` resource) and recreate any pods.
 * 
 * ## Import
 * 
 * terraform import {
 * 
 *  to = aws_eks_cluster.my_cluster
 * 
 *  id = &#34;my_cluster&#34; } Using `pulumi import`, import EKS Clusters using the `name`. For exampleconsole % pulumi import aws_eks_cluster.my_cluster my_cluster
 * 
 */
@ResourceType(type="aws:eks/cluster:Cluster")
public class Cluster extends com.pulumi.resources.CustomResource {
    /**
     * ARN of the cluster.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return ARN of the cluster.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    @Export(name="certificateAuthorities", refs={List.class,ClusterCertificateAuthority.class}, tree="[0,1]")
    private Output<List<ClusterCertificateAuthority>> certificateAuthorities;

    public Output<List<ClusterCertificateAuthority>> certificateAuthorities() {
        return this.certificateAuthorities;
    }
    /**
     * Attribute block containing `certificate-authority-data` for your cluster. Detailed below.
     * 
     */
    @Export(name="certificateAuthority", refs={ClusterCertificateAuthority.class}, tree="[0]")
    private Output<ClusterCertificateAuthority> certificateAuthority;

    /**
     * @return Attribute block containing `certificate-authority-data` for your cluster. Detailed below.
     * 
     */
    public Output<ClusterCertificateAuthority> certificateAuthority() {
        return this.certificateAuthority;
    }
    /**
     * The ID of your local Amazon EKS cluster on the AWS Outpost. This attribute isn&#39;t available for an AWS EKS cluster on AWS cloud.
     * 
     */
    @Export(name="clusterId", refs={String.class}, tree="[0]")
    private Output<String> clusterId;

    /**
     * @return The ID of your local Amazon EKS cluster on the AWS Outpost. This attribute isn&#39;t available for an AWS EKS cluster on AWS cloud.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }
    /**
     * Unix epoch timestamp in seconds for when the cluster was created.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return Unix epoch timestamp in seconds for when the cluster was created.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    @Export(name="defaultAddonsToRemoves", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> defaultAddonsToRemoves;

    public Output<Optional<List<String>>> defaultAddonsToRemoves() {
        return Codegen.optional(this.defaultAddonsToRemoves);
    }
    /**
     * List of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html).
     * 
     */
    @Export(name="enabledClusterLogTypes", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> enabledClusterLogTypes;

    /**
     * @return List of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html).
     * 
     */
    public Output<Optional<List<String>>> enabledClusterLogTypes() {
        return Codegen.optional(this.enabledClusterLogTypes);
    }
    /**
     * Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
     * 
     */
    @Export(name="encryptionConfig", refs={ClusterEncryptionConfig.class}, tree="[0]")
    private Output</* @Nullable */ ClusterEncryptionConfig> encryptionConfig;

    /**
     * @return Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
     * 
     */
    public Output<Optional<ClusterEncryptionConfig>> encryptionConfig() {
        return Codegen.optional(this.encryptionConfig);
    }
    /**
     * Endpoint for your Kubernetes API server.
     * 
     */
    @Export(name="endpoint", refs={String.class}, tree="[0]")
    private Output<String> endpoint;

    /**
     * @return Endpoint for your Kubernetes API server.
     * 
     */
    public Output<String> endpoint() {
        return this.endpoint;
    }
    /**
     * Attribute block containing identity provider information for your cluster. Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019. Detailed below.
     * * `kubernetes_network_config.service_ipv6_cidr` - The CIDR block that Kubernetes pod and service IP addresses are assigned from if you specified `ipv6` for ipFamily when you created the cluster. Kubernetes assigns service addresses from the unique local address range (fc00::/7) because you can&#39;t specify a custom IPv6 CIDR block when you create the cluster.
     * 
     */
    @Export(name="identities", refs={List.class,ClusterIdentity.class}, tree="[0,1]")
    private Output<List<ClusterIdentity>> identities;

    /**
     * @return Attribute block containing identity provider information for your cluster. Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019. Detailed below.
     * * `kubernetes_network_config.service_ipv6_cidr` - The CIDR block that Kubernetes pod and service IP addresses are assigned from if you specified `ipv6` for ipFamily when you created the cluster. Kubernetes assigns service addresses from the unique local address range (fc00::/7) because you can&#39;t specify a custom IPv6 CIDR block when you create the cluster.
     * 
     */
    public Output<List<ClusterIdentity>> identities() {
        return this.identities;
    }
    /**
     * Configuration block with kubernetes network configuration for the cluster. Detailed below. If removed, this provider will only perform drift detection if a configuration value is provided.
     * 
     */
    @Export(name="kubernetesNetworkConfig", refs={ClusterKubernetesNetworkConfig.class}, tree="[0]")
    private Output<ClusterKubernetesNetworkConfig> kubernetesNetworkConfig;

    /**
     * @return Configuration block with kubernetes network configuration for the cluster. Detailed below. If removed, this provider will only perform drift detection if a configuration value is provided.
     * 
     */
    public Output<ClusterKubernetesNetworkConfig> kubernetesNetworkConfig() {
        return this.kubernetesNetworkConfig;
    }
    /**
     * Name of the cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Configuration block representing the configuration of your local Amazon EKS cluster on an AWS Outpost. This block isn&#39;t available for creating Amazon EKS clusters on the AWS cloud.
     * 
     */
    @Export(name="outpostConfig", refs={ClusterOutpostConfig.class}, tree="[0]")
    private Output</* @Nullable */ ClusterOutpostConfig> outpostConfig;

    /**
     * @return Configuration block representing the configuration of your local Amazon EKS cluster on an AWS Outpost. This block isn&#39;t available for creating Amazon EKS clusters on the AWS cloud.
     * 
     */
    public Output<Optional<ClusterOutpostConfig>> outpostConfig() {
        return Codegen.optional(this.outpostConfig);
    }
    /**
     * Platform version for the cluster.
     * 
     */
    @Export(name="platformVersion", refs={String.class}, tree="[0]")
    private Output<String> platformVersion;

    /**
     * @return Platform version for the cluster.
     * 
     */
    public Output<String> platformVersion() {
        return this.platformVersion;
    }
    /**
     * ARN of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf. Ensure the resource configuration includes explicit dependencies on the IAM Role permissions by adding `depends_on` if using the `aws.iam.RolePolicy` resource or `aws.iam.RolePolicyAttachment` resource, otherwise EKS cannot delete EKS managed EC2 infrastructure such as Security Groups on EKS Cluster deletion.
     * 
     */
    @Export(name="roleArn", refs={String.class}, tree="[0]")
    private Output<String> roleArn;

    /**
     * @return ARN of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf. Ensure the resource configuration includes explicit dependencies on the IAM Role permissions by adding `depends_on` if using the `aws.iam.RolePolicy` resource or `aws.iam.RolePolicyAttachment` resource, otherwise EKS cannot delete EKS managed EC2 infrastructure such as Security Groups on EKS Cluster deletion.
     * 
     */
    public Output<String> roleArn() {
        return this.roleArn;
    }
    /**
     * Status of the EKS cluster. One of `CREATING`, `ACTIVE`, `DELETING`, `FAILED`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Status of the EKS cluster. One of `CREATING`, `ACTIVE`, `DELETING`, `FAILED`.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    @Export(name="tagsAll", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> tagsAll;

    /**
     * @return Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }
    /**
     * Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
     * 
     */
    @Export(name="version", refs={String.class}, tree="[0]")
    private Output<String> version;

    /**
     * @return Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
     * 
     */
    public Output<String> version() {
        return this.version;
    }
    /**
     * Configuration block for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Detailed below. Also contains attributes detailed in the Attributes section.
     * 
     * The following arguments are optional:
     * 
     */
    @Export(name="vpcConfig", refs={ClusterVpcConfig.class}, tree="[0]")
    private Output<ClusterVpcConfig> vpcConfig;

    /**
     * @return Configuration block for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Detailed below. Also contains attributes detailed in the Attributes section.
     * 
     * The following arguments are optional:
     * 
     */
    public Output<ClusterVpcConfig> vpcConfig() {
        return this.vpcConfig;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Cluster(String name) {
        this(name, ClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Cluster(String name, ClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Cluster(String name, ClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:eks/cluster:Cluster", name, args == null ? ClusterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Cluster(String name, Output<String> id, @Nullable ClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:eks/cluster:Cluster", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Cluster get(String name, Output<String> id, @Nullable ClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Cluster(name, id, state, options);
    }
}
