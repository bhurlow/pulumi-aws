// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.eks.inputs;

import com.pulumi.aws.eks.inputs.ClusterCertificateAuthorityArgs;
import com.pulumi.aws.eks.inputs.ClusterEncryptionConfigArgs;
import com.pulumi.aws.eks.inputs.ClusterIdentityArgs;
import com.pulumi.aws.eks.inputs.ClusterKubernetesNetworkConfigArgs;
import com.pulumi.aws.eks.inputs.ClusterOutpostConfigArgs;
import com.pulumi.aws.eks.inputs.ClusterVpcConfigArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClusterState extends com.pulumi.resources.ResourceArgs {

    public static final ClusterState Empty = new ClusterState();

    /**
     * ARN of the cluster.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return ARN of the cluster.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    @Import(name="certificateAuthorities")
    private @Nullable Output<List<ClusterCertificateAuthorityArgs>> certificateAuthorities;

    public Optional<Output<List<ClusterCertificateAuthorityArgs>>> certificateAuthorities() {
        return Optional.ofNullable(this.certificateAuthorities);
    }

    /**
     * Attribute block containing `certificate-authority-data` for your cluster. Detailed below.
     * 
     */
    @Import(name="certificateAuthority")
    private @Nullable Output<ClusterCertificateAuthorityArgs> certificateAuthority;

    /**
     * @return Attribute block containing `certificate-authority-data` for your cluster. Detailed below.
     * 
     */
    public Optional<Output<ClusterCertificateAuthorityArgs>> certificateAuthority() {
        return Optional.ofNullable(this.certificateAuthority);
    }

    /**
     * Unix epoch timestamp in seconds for when the cluster was created.
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return Unix epoch timestamp in seconds for when the cluster was created.
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    @Import(name="defaultAddonsToRemoves")
    private @Nullable Output<List<String>> defaultAddonsToRemoves;

    public Optional<Output<List<String>>> defaultAddonsToRemoves() {
        return Optional.ofNullable(this.defaultAddonsToRemoves);
    }

    /**
     * List of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html).
     * 
     */
    @Import(name="enabledClusterLogTypes")
    private @Nullable Output<List<String>> enabledClusterLogTypes;

    /**
     * @return List of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html).
     * 
     */
    public Optional<Output<List<String>>> enabledClusterLogTypes() {
        return Optional.ofNullable(this.enabledClusterLogTypes);
    }

    /**
     * Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
     * 
     */
    @Import(name="encryptionConfig")
    private @Nullable Output<ClusterEncryptionConfigArgs> encryptionConfig;

    /**
     * @return Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
     * 
     */
    public Optional<Output<ClusterEncryptionConfigArgs>> encryptionConfig() {
        return Optional.ofNullable(this.encryptionConfig);
    }

    /**
     * Endpoint for your Kubernetes API server.
     * 
     */
    @Import(name="endpoint")
    private @Nullable Output<String> endpoint;

    /**
     * @return Endpoint for your Kubernetes API server.
     * 
     */
    public Optional<Output<String>> endpoint() {
        return Optional.ofNullable(this.endpoint);
    }

    /**
     * Attribute block containing identity provider information for your cluster. Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019. Detailed below.
     * * `kubernetes_network_config.service_ipv6_cidr` - The CIDR block that Kubernetes pod and service IP addresses are assigned from if you specified `ipv6` for ipFamily when you created the cluster. Kubernetes assigns service addresses from the unique local address range (fc00::/7) because you can&#39;t specify a custom IPv6 CIDR block when you create the cluster.
     * 
     */
    @Import(name="identities")
    private @Nullable Output<List<ClusterIdentityArgs>> identities;

    /**
     * @return Attribute block containing identity provider information for your cluster. Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019. Detailed below.
     * * `kubernetes_network_config.service_ipv6_cidr` - The CIDR block that Kubernetes pod and service IP addresses are assigned from if you specified `ipv6` for ipFamily when you created the cluster. Kubernetes assigns service addresses from the unique local address range (fc00::/7) because you can&#39;t specify a custom IPv6 CIDR block when you create the cluster.
     * 
     */
    public Optional<Output<List<ClusterIdentityArgs>>> identities() {
        return Optional.ofNullable(this.identities);
    }

    /**
     * Configuration block with kubernetes network configuration for the cluster. Detailed below. If removed, this provider will only perform drift detection if a configuration value is provided.
     * 
     */
    @Import(name="kubernetesNetworkConfig")
    private @Nullable Output<ClusterKubernetesNetworkConfigArgs> kubernetesNetworkConfig;

    /**
     * @return Configuration block with kubernetes network configuration for the cluster. Detailed below. If removed, this provider will only perform drift detection if a configuration value is provided.
     * 
     */
    public Optional<Output<ClusterKubernetesNetworkConfigArgs>> kubernetesNetworkConfig() {
        return Optional.ofNullable(this.kubernetesNetworkConfig);
    }

    /**
     * Name of the cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Configuration block representing the configuration of your local Amazon EKS cluster on an AWS Outpost. This block isn&#39;t available for creating Amazon EKS clusters on the AWS cloud.
     * 
     */
    @Import(name="outpostConfig")
    private @Nullable Output<ClusterOutpostConfigArgs> outpostConfig;

    /**
     * @return Configuration block representing the configuration of your local Amazon EKS cluster on an AWS Outpost. This block isn&#39;t available for creating Amazon EKS clusters on the AWS cloud.
     * 
     */
    public Optional<Output<ClusterOutpostConfigArgs>> outpostConfig() {
        return Optional.ofNullable(this.outpostConfig);
    }

    /**
     * Platform version for the cluster.
     * 
     */
    @Import(name="platformVersion")
    private @Nullable Output<String> platformVersion;

    /**
     * @return Platform version for the cluster.
     * 
     */
    public Optional<Output<String>> platformVersion() {
        return Optional.ofNullable(this.platformVersion);
    }

    /**
     * ARN of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf. Ensure the resource configuration includes explicit dependencies on the IAM Role permissions by adding `depends_on` if using the `aws.iam.RolePolicy` resource or `aws.iam.RolePolicyAttachment` resource, otherwise EKS cannot delete EKS managed EC2 infrastructure such as Security Groups on EKS Cluster deletion.
     * 
     */
    @Import(name="roleArn")
    private @Nullable Output<String> roleArn;

    /**
     * @return ARN of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf. Ensure the resource configuration includes explicit dependencies on the IAM Role permissions by adding `depends_on` if using the `aws.iam.RolePolicy` resource or `aws.iam.RolePolicyAttachment` resource, otherwise EKS cannot delete EKS managed EC2 infrastructure such as Security Groups on EKS Cluster deletion.
     * 
     */
    public Optional<Output<String>> roleArn() {
        return Optional.ofNullable(this.roleArn);
    }

    /**
     * Status of the EKS cluster. One of `CREATING`, `ACTIVE`, `DELETING`, `FAILED`.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return Status of the EKS cluster. One of `CREATING`, `ACTIVE`, `DELETING`, `FAILED`.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    /**
     * Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    /**
     * Configuration block for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Detailed below. Also contains attributes detailed in the Attributes section.
     * 
     */
    @Import(name="vpcConfig")
    private @Nullable Output<ClusterVpcConfigArgs> vpcConfig;

    /**
     * @return Configuration block for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Detailed below. Also contains attributes detailed in the Attributes section.
     * 
     */
    public Optional<Output<ClusterVpcConfigArgs>> vpcConfig() {
        return Optional.ofNullable(this.vpcConfig);
    }

    private ClusterState() {}

    private ClusterState(ClusterState $) {
        this.arn = $.arn;
        this.certificateAuthorities = $.certificateAuthorities;
        this.certificateAuthority = $.certificateAuthority;
        this.createdAt = $.createdAt;
        this.defaultAddonsToRemoves = $.defaultAddonsToRemoves;
        this.enabledClusterLogTypes = $.enabledClusterLogTypes;
        this.encryptionConfig = $.encryptionConfig;
        this.endpoint = $.endpoint;
        this.identities = $.identities;
        this.kubernetesNetworkConfig = $.kubernetesNetworkConfig;
        this.name = $.name;
        this.outpostConfig = $.outpostConfig;
        this.platformVersion = $.platformVersion;
        this.roleArn = $.roleArn;
        this.status = $.status;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.version = $.version;
        this.vpcConfig = $.vpcConfig;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClusterState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClusterState $;

        public Builder() {
            $ = new ClusterState();
        }

        public Builder(ClusterState defaults) {
            $ = new ClusterState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn ARN of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn ARN of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        public Builder certificateAuthorities(@Nullable Output<List<ClusterCertificateAuthorityArgs>> certificateAuthorities) {
            $.certificateAuthorities = certificateAuthorities;
            return this;
        }

        public Builder certificateAuthorities(List<ClusterCertificateAuthorityArgs> certificateAuthorities) {
            return certificateAuthorities(Output.of(certificateAuthorities));
        }

        public Builder certificateAuthorities(ClusterCertificateAuthorityArgs... certificateAuthorities) {
            return certificateAuthorities(List.of(certificateAuthorities));
        }

        /**
         * @param certificateAuthority Attribute block containing `certificate-authority-data` for your cluster. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder certificateAuthority(@Nullable Output<ClusterCertificateAuthorityArgs> certificateAuthority) {
            $.certificateAuthority = certificateAuthority;
            return this;
        }

        /**
         * @param certificateAuthority Attribute block containing `certificate-authority-data` for your cluster. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder certificateAuthority(ClusterCertificateAuthorityArgs certificateAuthority) {
            return certificateAuthority(Output.of(certificateAuthority));
        }

        /**
         * @param createdAt Unix epoch timestamp in seconds for when the cluster was created.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt Unix epoch timestamp in seconds for when the cluster was created.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        public Builder defaultAddonsToRemoves(@Nullable Output<List<String>> defaultAddonsToRemoves) {
            $.defaultAddonsToRemoves = defaultAddonsToRemoves;
            return this;
        }

        public Builder defaultAddonsToRemoves(List<String> defaultAddonsToRemoves) {
            return defaultAddonsToRemoves(Output.of(defaultAddonsToRemoves));
        }

        public Builder defaultAddonsToRemoves(String... defaultAddonsToRemoves) {
            return defaultAddonsToRemoves(List.of(defaultAddonsToRemoves));
        }

        /**
         * @param enabledClusterLogTypes List of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html).
         * 
         * @return builder
         * 
         */
        public Builder enabledClusterLogTypes(@Nullable Output<List<String>> enabledClusterLogTypes) {
            $.enabledClusterLogTypes = enabledClusterLogTypes;
            return this;
        }

        /**
         * @param enabledClusterLogTypes List of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html).
         * 
         * @return builder
         * 
         */
        public Builder enabledClusterLogTypes(List<String> enabledClusterLogTypes) {
            return enabledClusterLogTypes(Output.of(enabledClusterLogTypes));
        }

        /**
         * @param enabledClusterLogTypes List of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html).
         * 
         * @return builder
         * 
         */
        public Builder enabledClusterLogTypes(String... enabledClusterLogTypes) {
            return enabledClusterLogTypes(List.of(enabledClusterLogTypes));
        }

        /**
         * @param encryptionConfig Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder encryptionConfig(@Nullable Output<ClusterEncryptionConfigArgs> encryptionConfig) {
            $.encryptionConfig = encryptionConfig;
            return this;
        }

        /**
         * @param encryptionConfig Configuration block with encryption configuration for the cluster. Only available on Kubernetes 1.13 and above clusters created after March 6, 2020. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder encryptionConfig(ClusterEncryptionConfigArgs encryptionConfig) {
            return encryptionConfig(Output.of(encryptionConfig));
        }

        /**
         * @param endpoint Endpoint for your Kubernetes API server.
         * 
         * @return builder
         * 
         */
        public Builder endpoint(@Nullable Output<String> endpoint) {
            $.endpoint = endpoint;
            return this;
        }

        /**
         * @param endpoint Endpoint for your Kubernetes API server.
         * 
         * @return builder
         * 
         */
        public Builder endpoint(String endpoint) {
            return endpoint(Output.of(endpoint));
        }

        /**
         * @param identities Attribute block containing identity provider information for your cluster. Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019. Detailed below.
         * * `kubernetes_network_config.service_ipv6_cidr` - The CIDR block that Kubernetes pod and service IP addresses are assigned from if you specified `ipv6` for ipFamily when you created the cluster. Kubernetes assigns service addresses from the unique local address range (fc00::/7) because you can&#39;t specify a custom IPv6 CIDR block when you create the cluster.
         * 
         * @return builder
         * 
         */
        public Builder identities(@Nullable Output<List<ClusterIdentityArgs>> identities) {
            $.identities = identities;
            return this;
        }

        /**
         * @param identities Attribute block containing identity provider information for your cluster. Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019. Detailed below.
         * * `kubernetes_network_config.service_ipv6_cidr` - The CIDR block that Kubernetes pod and service IP addresses are assigned from if you specified `ipv6` for ipFamily when you created the cluster. Kubernetes assigns service addresses from the unique local address range (fc00::/7) because you can&#39;t specify a custom IPv6 CIDR block when you create the cluster.
         * 
         * @return builder
         * 
         */
        public Builder identities(List<ClusterIdentityArgs> identities) {
            return identities(Output.of(identities));
        }

        /**
         * @param identities Attribute block containing identity provider information for your cluster. Only available on Kubernetes version 1.13 and 1.14 clusters created or upgraded on or after September 3, 2019. Detailed below.
         * * `kubernetes_network_config.service_ipv6_cidr` - The CIDR block that Kubernetes pod and service IP addresses are assigned from if you specified `ipv6` for ipFamily when you created the cluster. Kubernetes assigns service addresses from the unique local address range (fc00::/7) because you can&#39;t specify a custom IPv6 CIDR block when you create the cluster.
         * 
         * @return builder
         * 
         */
        public Builder identities(ClusterIdentityArgs... identities) {
            return identities(List.of(identities));
        }

        /**
         * @param kubernetesNetworkConfig Configuration block with kubernetes network configuration for the cluster. Detailed below. If removed, this provider will only perform drift detection if a configuration value is provided.
         * 
         * @return builder
         * 
         */
        public Builder kubernetesNetworkConfig(@Nullable Output<ClusterKubernetesNetworkConfigArgs> kubernetesNetworkConfig) {
            $.kubernetesNetworkConfig = kubernetesNetworkConfig;
            return this;
        }

        /**
         * @param kubernetesNetworkConfig Configuration block with kubernetes network configuration for the cluster. Detailed below. If removed, this provider will only perform drift detection if a configuration value is provided.
         * 
         * @return builder
         * 
         */
        public Builder kubernetesNetworkConfig(ClusterKubernetesNetworkConfigArgs kubernetesNetworkConfig) {
            return kubernetesNetworkConfig(Output.of(kubernetesNetworkConfig));
        }

        /**
         * @param name Name of the cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the cluster. Must be between 1-100 characters in length. Must begin with an alphanumeric character, and must only contain alphanumeric characters, dashes and underscores (`^[0-9A-Za-z][A-Za-z0-9\-_]+$`).
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param outpostConfig Configuration block representing the configuration of your local Amazon EKS cluster on an AWS Outpost. This block isn&#39;t available for creating Amazon EKS clusters on the AWS cloud.
         * 
         * @return builder
         * 
         */
        public Builder outpostConfig(@Nullable Output<ClusterOutpostConfigArgs> outpostConfig) {
            $.outpostConfig = outpostConfig;
            return this;
        }

        /**
         * @param outpostConfig Configuration block representing the configuration of your local Amazon EKS cluster on an AWS Outpost. This block isn&#39;t available for creating Amazon EKS clusters on the AWS cloud.
         * 
         * @return builder
         * 
         */
        public Builder outpostConfig(ClusterOutpostConfigArgs outpostConfig) {
            return outpostConfig(Output.of(outpostConfig));
        }

        /**
         * @param platformVersion Platform version for the cluster.
         * 
         * @return builder
         * 
         */
        public Builder platformVersion(@Nullable Output<String> platformVersion) {
            $.platformVersion = platformVersion;
            return this;
        }

        /**
         * @param platformVersion Platform version for the cluster.
         * 
         * @return builder
         * 
         */
        public Builder platformVersion(String platformVersion) {
            return platformVersion(Output.of(platformVersion));
        }

        /**
         * @param roleArn ARN of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf. Ensure the resource configuration includes explicit dependencies on the IAM Role permissions by adding `depends_on` if using the `aws.iam.RolePolicy` resource or `aws.iam.RolePolicyAttachment` resource, otherwise EKS cannot delete EKS managed EC2 infrastructure such as Security Groups on EKS Cluster deletion.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(@Nullable Output<String> roleArn) {
            $.roleArn = roleArn;
            return this;
        }

        /**
         * @param roleArn ARN of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf. Ensure the resource configuration includes explicit dependencies on the IAM Role permissions by adding `depends_on` if using the `aws.iam.RolePolicy` resource or `aws.iam.RolePolicyAttachment` resource, otherwise EKS cannot delete EKS managed EC2 infrastructure such as Security Groups on EKS Cluster deletion.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(String roleArn) {
            return roleArn(Output.of(roleArn));
        }

        /**
         * @param status Status of the EKS cluster. One of `CREATING`, `ACTIVE`, `DELETING`, `FAILED`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status Status of the EKS cluster. One of `CREATING`, `ACTIVE`, `DELETING`, `FAILED`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tagsAll Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        /**
         * @param version Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        /**
         * @param vpcConfig Configuration block for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Detailed below. Also contains attributes detailed in the Attributes section.
         * 
         * @return builder
         * 
         */
        public Builder vpcConfig(@Nullable Output<ClusterVpcConfigArgs> vpcConfig) {
            $.vpcConfig = vpcConfig;
            return this;
        }

        /**
         * @param vpcConfig Configuration block for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Detailed below. Also contains attributes detailed in the Attributes section.
         * 
         * @return builder
         * 
         */
        public Builder vpcConfig(ClusterVpcConfigArgs vpcConfig) {
            return vpcConfig(Output.of(vpcConfig));
        }

        public ClusterState build() {
            return $;
        }
    }

}
