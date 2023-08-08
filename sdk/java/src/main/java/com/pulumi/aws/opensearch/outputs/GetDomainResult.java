// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.opensearch.outputs;

import com.pulumi.aws.opensearch.outputs.GetDomainAdvancedSecurityOption;
import com.pulumi.aws.opensearch.outputs.GetDomainAutoTuneOption;
import com.pulumi.aws.opensearch.outputs.GetDomainClusterConfig;
import com.pulumi.aws.opensearch.outputs.GetDomainCognitoOption;
import com.pulumi.aws.opensearch.outputs.GetDomainEbsOption;
import com.pulumi.aws.opensearch.outputs.GetDomainEncryptionAtRest;
import com.pulumi.aws.opensearch.outputs.GetDomainLogPublishingOption;
import com.pulumi.aws.opensearch.outputs.GetDomainNodeToNodeEncryption;
import com.pulumi.aws.opensearch.outputs.GetDomainOffPeakWindowOptions;
import com.pulumi.aws.opensearch.outputs.GetDomainSnapshotOption;
import com.pulumi.aws.opensearch.outputs.GetDomainVpcOption;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetDomainResult {
    /**
     * @return Policy document attached to the domain.
     * 
     */
    private String accessPolicies;
    /**
     * @return Key-value string pairs to specify advanced configuration options.
     * 
     */
    private Map<String,String> advancedOptions;
    /**
     * @return Status of the OpenSearch domain&#39;s advanced security options. The block consists of the following attributes:
     * 
     */
    private List<GetDomainAdvancedSecurityOption> advancedSecurityOptions;
    /**
     * @return ARN of the domain.
     * 
     */
    private String arn;
    /**
     * @return Configuration of the Auto-Tune options of the domain.
     * 
     */
    private List<GetDomainAutoTuneOption> autoTuneOptions;
    /**
     * @return Cluster configuration of the domain.
     * 
     */
    private List<GetDomainClusterConfig> clusterConfigs;
    /**
     * @return Domain Amazon Cognito Authentication options for Dashboard.
     * 
     */
    private List<GetDomainCognitoOption> cognitoOptions;
    /**
     * @return Status of the creation of the domain.
     * 
     */
    private Boolean created;
    /**
     * @return Domain-specific endpoint used to access the [Dashboard application](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/dashboards.html).
     * 
     */
    private String dashboardEndpoint;
    /**
     * @return Status of the deletion of the domain.
     * 
     */
    private Boolean deleted;
    /**
     * @return Unique identifier for the domain.
     * 
     */
    private String domainId;
    private String domainName;
    /**
     * @return EBS Options for the instances in the domain.
     * 
     */
    private List<GetDomainEbsOption> ebsOptions;
    /**
     * @return Domain encryption at rest related options.
     * 
     */
    private List<GetDomainEncryptionAtRest> encryptionAtRests;
    /**
     * @return Domain-specific endpoint used to submit index, search, and data upload requests.
     * 
     */
    private String endpoint;
    /**
     * @return OpenSearch version for the domain.
     * 
     */
    private String engineVersion;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return (**Deprecated**) Domain-specific endpoint for kibana without https scheme. Use the `dashboard_endpoint` attribute instead.
     * 
     * @deprecated
     * use &#39;dashboard_endpoint&#39; attribute instead
     * 
     */
    @Deprecated /* use 'dashboard_endpoint' attribute instead */
    private String kibanaEndpoint;
    /**
     * @return Domain log publishing related options.
     * 
     */
    private List<GetDomainLogPublishingOption> logPublishingOptions;
    /**
     * @return Domain in transit encryption related options.
     * 
     */
    private List<GetDomainNodeToNodeEncryption> nodeToNodeEncryptions;
    /**
     * @return Off Peak update options
     * 
     */
    private @Nullable GetDomainOffPeakWindowOptions offPeakWindowOptions;
    /**
     * @return Status of a configuration change in the domain.
     * 
     */
    private Boolean processing;
    /**
     * @return Domain snapshot related options.
     * 
     */
    private List<GetDomainSnapshotOption> snapshotOptions;
    /**
     * @return Tags assigned to the domain.
     * 
     */
    private Map<String,String> tags;
    /**
     * @return VPC Options for private OpenSearch domains.
     * 
     */
    private List<GetDomainVpcOption> vpcOptions;

    private GetDomainResult() {}
    /**
     * @return Policy document attached to the domain.
     * 
     */
    public String accessPolicies() {
        return this.accessPolicies;
    }
    /**
     * @return Key-value string pairs to specify advanced configuration options.
     * 
     */
    public Map<String,String> advancedOptions() {
        return this.advancedOptions;
    }
    /**
     * @return Status of the OpenSearch domain&#39;s advanced security options. The block consists of the following attributes:
     * 
     */
    public List<GetDomainAdvancedSecurityOption> advancedSecurityOptions() {
        return this.advancedSecurityOptions;
    }
    /**
     * @return ARN of the domain.
     * 
     */
    public String arn() {
        return this.arn;
    }
    /**
     * @return Configuration of the Auto-Tune options of the domain.
     * 
     */
    public List<GetDomainAutoTuneOption> autoTuneOptions() {
        return this.autoTuneOptions;
    }
    /**
     * @return Cluster configuration of the domain.
     * 
     */
    public List<GetDomainClusterConfig> clusterConfigs() {
        return this.clusterConfigs;
    }
    /**
     * @return Domain Amazon Cognito Authentication options for Dashboard.
     * 
     */
    public List<GetDomainCognitoOption> cognitoOptions() {
        return this.cognitoOptions;
    }
    /**
     * @return Status of the creation of the domain.
     * 
     */
    public Boolean created() {
        return this.created;
    }
    /**
     * @return Domain-specific endpoint used to access the [Dashboard application](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/dashboards.html).
     * 
     */
    public String dashboardEndpoint() {
        return this.dashboardEndpoint;
    }
    /**
     * @return Status of the deletion of the domain.
     * 
     */
    public Boolean deleted() {
        return this.deleted;
    }
    /**
     * @return Unique identifier for the domain.
     * 
     */
    public String domainId() {
        return this.domainId;
    }
    public String domainName() {
        return this.domainName;
    }
    /**
     * @return EBS Options for the instances in the domain.
     * 
     */
    public List<GetDomainEbsOption> ebsOptions() {
        return this.ebsOptions;
    }
    /**
     * @return Domain encryption at rest related options.
     * 
     */
    public List<GetDomainEncryptionAtRest> encryptionAtRests() {
        return this.encryptionAtRests;
    }
    /**
     * @return Domain-specific endpoint used to submit index, search, and data upload requests.
     * 
     */
    public String endpoint() {
        return this.endpoint;
    }
    /**
     * @return OpenSearch version for the domain.
     * 
     */
    public String engineVersion() {
        return this.engineVersion;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return (**Deprecated**) Domain-specific endpoint for kibana without https scheme. Use the `dashboard_endpoint` attribute instead.
     * 
     * @deprecated
     * use &#39;dashboard_endpoint&#39; attribute instead
     * 
     */
    @Deprecated /* use 'dashboard_endpoint' attribute instead */
    public String kibanaEndpoint() {
        return this.kibanaEndpoint;
    }
    /**
     * @return Domain log publishing related options.
     * 
     */
    public List<GetDomainLogPublishingOption> logPublishingOptions() {
        return this.logPublishingOptions;
    }
    /**
     * @return Domain in transit encryption related options.
     * 
     */
    public List<GetDomainNodeToNodeEncryption> nodeToNodeEncryptions() {
        return this.nodeToNodeEncryptions;
    }
    /**
     * @return Off Peak update options
     * 
     */
    public Optional<GetDomainOffPeakWindowOptions> offPeakWindowOptions() {
        return Optional.ofNullable(this.offPeakWindowOptions);
    }
    /**
     * @return Status of a configuration change in the domain.
     * 
     */
    public Boolean processing() {
        return this.processing;
    }
    /**
     * @return Domain snapshot related options.
     * 
     */
    public List<GetDomainSnapshotOption> snapshotOptions() {
        return this.snapshotOptions;
    }
    /**
     * @return Tags assigned to the domain.
     * 
     */
    public Map<String,String> tags() {
        return this.tags;
    }
    /**
     * @return VPC Options for private OpenSearch domains.
     * 
     */
    public List<GetDomainVpcOption> vpcOptions() {
        return this.vpcOptions;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDomainResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accessPolicies;
        private Map<String,String> advancedOptions;
        private List<GetDomainAdvancedSecurityOption> advancedSecurityOptions;
        private String arn;
        private List<GetDomainAutoTuneOption> autoTuneOptions;
        private List<GetDomainClusterConfig> clusterConfigs;
        private List<GetDomainCognitoOption> cognitoOptions;
        private Boolean created;
        private String dashboardEndpoint;
        private Boolean deleted;
        private String domainId;
        private String domainName;
        private List<GetDomainEbsOption> ebsOptions;
        private List<GetDomainEncryptionAtRest> encryptionAtRests;
        private String endpoint;
        private String engineVersion;
        private String id;
        private String kibanaEndpoint;
        private List<GetDomainLogPublishingOption> logPublishingOptions;
        private List<GetDomainNodeToNodeEncryption> nodeToNodeEncryptions;
        private @Nullable GetDomainOffPeakWindowOptions offPeakWindowOptions;
        private Boolean processing;
        private List<GetDomainSnapshotOption> snapshotOptions;
        private Map<String,String> tags;
        private List<GetDomainVpcOption> vpcOptions;
        public Builder() {}
        public Builder(GetDomainResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessPolicies = defaults.accessPolicies;
    	      this.advancedOptions = defaults.advancedOptions;
    	      this.advancedSecurityOptions = defaults.advancedSecurityOptions;
    	      this.arn = defaults.arn;
    	      this.autoTuneOptions = defaults.autoTuneOptions;
    	      this.clusterConfigs = defaults.clusterConfigs;
    	      this.cognitoOptions = defaults.cognitoOptions;
    	      this.created = defaults.created;
    	      this.dashboardEndpoint = defaults.dashboardEndpoint;
    	      this.deleted = defaults.deleted;
    	      this.domainId = defaults.domainId;
    	      this.domainName = defaults.domainName;
    	      this.ebsOptions = defaults.ebsOptions;
    	      this.encryptionAtRests = defaults.encryptionAtRests;
    	      this.endpoint = defaults.endpoint;
    	      this.engineVersion = defaults.engineVersion;
    	      this.id = defaults.id;
    	      this.kibanaEndpoint = defaults.kibanaEndpoint;
    	      this.logPublishingOptions = defaults.logPublishingOptions;
    	      this.nodeToNodeEncryptions = defaults.nodeToNodeEncryptions;
    	      this.offPeakWindowOptions = defaults.offPeakWindowOptions;
    	      this.processing = defaults.processing;
    	      this.snapshotOptions = defaults.snapshotOptions;
    	      this.tags = defaults.tags;
    	      this.vpcOptions = defaults.vpcOptions;
        }

        @CustomType.Setter
        public Builder accessPolicies(String accessPolicies) {
            this.accessPolicies = Objects.requireNonNull(accessPolicies);
            return this;
        }
        @CustomType.Setter
        public Builder advancedOptions(Map<String,String> advancedOptions) {
            this.advancedOptions = Objects.requireNonNull(advancedOptions);
            return this;
        }
        @CustomType.Setter
        public Builder advancedSecurityOptions(List<GetDomainAdvancedSecurityOption> advancedSecurityOptions) {
            this.advancedSecurityOptions = Objects.requireNonNull(advancedSecurityOptions);
            return this;
        }
        public Builder advancedSecurityOptions(GetDomainAdvancedSecurityOption... advancedSecurityOptions) {
            return advancedSecurityOptions(List.of(advancedSecurityOptions));
        }
        @CustomType.Setter
        public Builder arn(String arn) {
            this.arn = Objects.requireNonNull(arn);
            return this;
        }
        @CustomType.Setter
        public Builder autoTuneOptions(List<GetDomainAutoTuneOption> autoTuneOptions) {
            this.autoTuneOptions = Objects.requireNonNull(autoTuneOptions);
            return this;
        }
        public Builder autoTuneOptions(GetDomainAutoTuneOption... autoTuneOptions) {
            return autoTuneOptions(List.of(autoTuneOptions));
        }
        @CustomType.Setter
        public Builder clusterConfigs(List<GetDomainClusterConfig> clusterConfigs) {
            this.clusterConfigs = Objects.requireNonNull(clusterConfigs);
            return this;
        }
        public Builder clusterConfigs(GetDomainClusterConfig... clusterConfigs) {
            return clusterConfigs(List.of(clusterConfigs));
        }
        @CustomType.Setter
        public Builder cognitoOptions(List<GetDomainCognitoOption> cognitoOptions) {
            this.cognitoOptions = Objects.requireNonNull(cognitoOptions);
            return this;
        }
        public Builder cognitoOptions(GetDomainCognitoOption... cognitoOptions) {
            return cognitoOptions(List.of(cognitoOptions));
        }
        @CustomType.Setter
        public Builder created(Boolean created) {
            this.created = Objects.requireNonNull(created);
            return this;
        }
        @CustomType.Setter
        public Builder dashboardEndpoint(String dashboardEndpoint) {
            this.dashboardEndpoint = Objects.requireNonNull(dashboardEndpoint);
            return this;
        }
        @CustomType.Setter
        public Builder deleted(Boolean deleted) {
            this.deleted = Objects.requireNonNull(deleted);
            return this;
        }
        @CustomType.Setter
        public Builder domainId(String domainId) {
            this.domainId = Objects.requireNonNull(domainId);
            return this;
        }
        @CustomType.Setter
        public Builder domainName(String domainName) {
            this.domainName = Objects.requireNonNull(domainName);
            return this;
        }
        @CustomType.Setter
        public Builder ebsOptions(List<GetDomainEbsOption> ebsOptions) {
            this.ebsOptions = Objects.requireNonNull(ebsOptions);
            return this;
        }
        public Builder ebsOptions(GetDomainEbsOption... ebsOptions) {
            return ebsOptions(List.of(ebsOptions));
        }
        @CustomType.Setter
        public Builder encryptionAtRests(List<GetDomainEncryptionAtRest> encryptionAtRests) {
            this.encryptionAtRests = Objects.requireNonNull(encryptionAtRests);
            return this;
        }
        public Builder encryptionAtRests(GetDomainEncryptionAtRest... encryptionAtRests) {
            return encryptionAtRests(List.of(encryptionAtRests));
        }
        @CustomType.Setter
        public Builder endpoint(String endpoint) {
            this.endpoint = Objects.requireNonNull(endpoint);
            return this;
        }
        @CustomType.Setter
        public Builder engineVersion(String engineVersion) {
            this.engineVersion = Objects.requireNonNull(engineVersion);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder kibanaEndpoint(String kibanaEndpoint) {
            this.kibanaEndpoint = Objects.requireNonNull(kibanaEndpoint);
            return this;
        }
        @CustomType.Setter
        public Builder logPublishingOptions(List<GetDomainLogPublishingOption> logPublishingOptions) {
            this.logPublishingOptions = Objects.requireNonNull(logPublishingOptions);
            return this;
        }
        public Builder logPublishingOptions(GetDomainLogPublishingOption... logPublishingOptions) {
            return logPublishingOptions(List.of(logPublishingOptions));
        }
        @CustomType.Setter
        public Builder nodeToNodeEncryptions(List<GetDomainNodeToNodeEncryption> nodeToNodeEncryptions) {
            this.nodeToNodeEncryptions = Objects.requireNonNull(nodeToNodeEncryptions);
            return this;
        }
        public Builder nodeToNodeEncryptions(GetDomainNodeToNodeEncryption... nodeToNodeEncryptions) {
            return nodeToNodeEncryptions(List.of(nodeToNodeEncryptions));
        }
        @CustomType.Setter
        public Builder offPeakWindowOptions(@Nullable GetDomainOffPeakWindowOptions offPeakWindowOptions) {
            this.offPeakWindowOptions = offPeakWindowOptions;
            return this;
        }
        @CustomType.Setter
        public Builder processing(Boolean processing) {
            this.processing = Objects.requireNonNull(processing);
            return this;
        }
        @CustomType.Setter
        public Builder snapshotOptions(List<GetDomainSnapshotOption> snapshotOptions) {
            this.snapshotOptions = Objects.requireNonNull(snapshotOptions);
            return this;
        }
        public Builder snapshotOptions(GetDomainSnapshotOption... snapshotOptions) {
            return snapshotOptions(List.of(snapshotOptions));
        }
        @CustomType.Setter
        public Builder tags(Map<String,String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        @CustomType.Setter
        public Builder vpcOptions(List<GetDomainVpcOption> vpcOptions) {
            this.vpcOptions = Objects.requireNonNull(vpcOptions);
            return this;
        }
        public Builder vpcOptions(GetDomainVpcOption... vpcOptions) {
            return vpcOptions(List.of(vpcOptions));
        }
        public GetDomainResult build() {
            final var o = new GetDomainResult();
            o.accessPolicies = accessPolicies;
            o.advancedOptions = advancedOptions;
            o.advancedSecurityOptions = advancedSecurityOptions;
            o.arn = arn;
            o.autoTuneOptions = autoTuneOptions;
            o.clusterConfigs = clusterConfigs;
            o.cognitoOptions = cognitoOptions;
            o.created = created;
            o.dashboardEndpoint = dashboardEndpoint;
            o.deleted = deleted;
            o.domainId = domainId;
            o.domainName = domainName;
            o.ebsOptions = ebsOptions;
            o.encryptionAtRests = encryptionAtRests;
            o.endpoint = endpoint;
            o.engineVersion = engineVersion;
            o.id = id;
            o.kibanaEndpoint = kibanaEndpoint;
            o.logPublishingOptions = logPublishingOptions;
            o.nodeToNodeEncryptions = nodeToNodeEncryptions;
            o.offPeakWindowOptions = offPeakWindowOptions;
            o.processing = processing;
            o.snapshotOptions = snapshotOptions;
            o.tags = tags;
            o.vpcOptions = vpcOptions;
            return o;
        }
    }
}
