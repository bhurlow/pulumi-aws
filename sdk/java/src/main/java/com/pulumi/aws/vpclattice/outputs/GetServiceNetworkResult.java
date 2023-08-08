// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.vpclattice.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetServiceNetworkResult {
    /**
     * @return ARN of the Service Network.
     * 
     */
    private String arn;
    /**
     * @return Authentication type for the service network. Either `NONE` or `AWS_IAM`.
     * 
     */
    private String authType;
    /**
     * @return Date and time the service network was created.
     * 
     */
    private String createdAt;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Date and time the service network was last updated.
     * 
     */
    private String lastUpdatedAt;
    /**
     * @return Name of the service network.
     * 
     */
    private String name;
    /**
     * @return Number of services associated with this service network.
     * 
     */
    private Integer numberOfAssociatedServices;
    /**
     * @return Number of VPCs associated with this service network.
     * 
     */
    private Integer numberOfAssociatedVpcs;
    private String serviceNetworkIdentifier;
    private Map<String,String> tags;

    private GetServiceNetworkResult() {}
    /**
     * @return ARN of the Service Network.
     * 
     */
    public String arn() {
        return this.arn;
    }
    /**
     * @return Authentication type for the service network. Either `NONE` or `AWS_IAM`.
     * 
     */
    public String authType() {
        return this.authType;
    }
    /**
     * @return Date and time the service network was created.
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Date and time the service network was last updated.
     * 
     */
    public String lastUpdatedAt() {
        return this.lastUpdatedAt;
    }
    /**
     * @return Name of the service network.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Number of services associated with this service network.
     * 
     */
    public Integer numberOfAssociatedServices() {
        return this.numberOfAssociatedServices;
    }
    /**
     * @return Number of VPCs associated with this service network.
     * 
     */
    public Integer numberOfAssociatedVpcs() {
        return this.numberOfAssociatedVpcs;
    }
    public String serviceNetworkIdentifier() {
        return this.serviceNetworkIdentifier;
    }
    public Map<String,String> tags() {
        return this.tags;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetServiceNetworkResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String arn;
        private String authType;
        private String createdAt;
        private String id;
        private String lastUpdatedAt;
        private String name;
        private Integer numberOfAssociatedServices;
        private Integer numberOfAssociatedVpcs;
        private String serviceNetworkIdentifier;
        private Map<String,String> tags;
        public Builder() {}
        public Builder(GetServiceNetworkResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.arn = defaults.arn;
    	      this.authType = defaults.authType;
    	      this.createdAt = defaults.createdAt;
    	      this.id = defaults.id;
    	      this.lastUpdatedAt = defaults.lastUpdatedAt;
    	      this.name = defaults.name;
    	      this.numberOfAssociatedServices = defaults.numberOfAssociatedServices;
    	      this.numberOfAssociatedVpcs = defaults.numberOfAssociatedVpcs;
    	      this.serviceNetworkIdentifier = defaults.serviceNetworkIdentifier;
    	      this.tags = defaults.tags;
        }

        @CustomType.Setter
        public Builder arn(String arn) {
            this.arn = Objects.requireNonNull(arn);
            return this;
        }
        @CustomType.Setter
        public Builder authType(String authType) {
            this.authType = Objects.requireNonNull(authType);
            return this;
        }
        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            this.createdAt = Objects.requireNonNull(createdAt);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder lastUpdatedAt(String lastUpdatedAt) {
            this.lastUpdatedAt = Objects.requireNonNull(lastUpdatedAt);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder numberOfAssociatedServices(Integer numberOfAssociatedServices) {
            this.numberOfAssociatedServices = Objects.requireNonNull(numberOfAssociatedServices);
            return this;
        }
        @CustomType.Setter
        public Builder numberOfAssociatedVpcs(Integer numberOfAssociatedVpcs) {
            this.numberOfAssociatedVpcs = Objects.requireNonNull(numberOfAssociatedVpcs);
            return this;
        }
        @CustomType.Setter
        public Builder serviceNetworkIdentifier(String serviceNetworkIdentifier) {
            this.serviceNetworkIdentifier = Objects.requireNonNull(serviceNetworkIdentifier);
            return this;
        }
        @CustomType.Setter
        public Builder tags(Map<String,String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        public GetServiceNetworkResult build() {
            final var o = new GetServiceNetworkResult();
            o.arn = arn;
            o.authType = authType;
            o.createdAt = createdAt;
            o.id = id;
            o.lastUpdatedAt = lastUpdatedAt;
            o.name = name;
            o.numberOfAssociatedServices = numberOfAssociatedServices;
            o.numberOfAssociatedVpcs = numberOfAssociatedVpcs;
            o.serviceNetworkIdentifier = serviceNetworkIdentifier;
            o.tags = tags;
            return o;
        }
    }
}
