// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2.outputs;

import com.pulumi.aws.ec2.outputs.GetNatGatewayFilter;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetNatGatewayResult {
    /**
     * @return ID of the EIP allocated to the selected NAT Gateway.
     * 
     */
    private String allocationId;
    /**
     * @return The association ID of the Elastic IP address that&#39;s associated with the NAT Gateway. Only available when `connectivity_type` is `public`.
     * 
     */
    private String associationId;
    /**
     * @return Connectivity type of the NAT Gateway.
     * 
     */
    private String connectivityType;
    private @Nullable List<GetNatGatewayFilter> filters;
    private String id;
    /**
     * @return The ID of the ENI allocated to the selected NAT Gateway.
     * 
     */
    private String networkInterfaceId;
    /**
     * @return Private IP address of the selected NAT Gateway.
     * 
     */
    private String privateIp;
    /**
     * @return Public IP (EIP) address of the selected NAT Gateway.
     * 
     */
    private String publicIp;
    /**
     * @return Secondary allocation EIP IDs for the selected NAT Gateway.
     * 
     */
    private List<String> secondaryAllocationIds;
    /**
     * @return The number of secondary private IPv4 addresses assigned to the selected NAT Gateway.
     * 
     */
    private Integer secondaryPrivateIpAddressCount;
    /**
     * @return Secondary private IPv4 addresses assigned to the selected NAT Gateway.
     * 
     */
    private List<String> secondaryPrivateIpAddresses;
    private String state;
    private String subnetId;
    private Map<String,String> tags;
    private String vpcId;

    private GetNatGatewayResult() {}
    /**
     * @return ID of the EIP allocated to the selected NAT Gateway.
     * 
     */
    public String allocationId() {
        return this.allocationId;
    }
    /**
     * @return The association ID of the Elastic IP address that&#39;s associated with the NAT Gateway. Only available when `connectivity_type` is `public`.
     * 
     */
    public String associationId() {
        return this.associationId;
    }
    /**
     * @return Connectivity type of the NAT Gateway.
     * 
     */
    public String connectivityType() {
        return this.connectivityType;
    }
    public List<GetNatGatewayFilter> filters() {
        return this.filters == null ? List.of() : this.filters;
    }
    public String id() {
        return this.id;
    }
    /**
     * @return The ID of the ENI allocated to the selected NAT Gateway.
     * 
     */
    public String networkInterfaceId() {
        return this.networkInterfaceId;
    }
    /**
     * @return Private IP address of the selected NAT Gateway.
     * 
     */
    public String privateIp() {
        return this.privateIp;
    }
    /**
     * @return Public IP (EIP) address of the selected NAT Gateway.
     * 
     */
    public String publicIp() {
        return this.publicIp;
    }
    /**
     * @return Secondary allocation EIP IDs for the selected NAT Gateway.
     * 
     */
    public List<String> secondaryAllocationIds() {
        return this.secondaryAllocationIds;
    }
    /**
     * @return The number of secondary private IPv4 addresses assigned to the selected NAT Gateway.
     * 
     */
    public Integer secondaryPrivateIpAddressCount() {
        return this.secondaryPrivateIpAddressCount;
    }
    /**
     * @return Secondary private IPv4 addresses assigned to the selected NAT Gateway.
     * 
     */
    public List<String> secondaryPrivateIpAddresses() {
        return this.secondaryPrivateIpAddresses;
    }
    public String state() {
        return this.state;
    }
    public String subnetId() {
        return this.subnetId;
    }
    public Map<String,String> tags() {
        return this.tags;
    }
    public String vpcId() {
        return this.vpcId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetNatGatewayResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String allocationId;
        private String associationId;
        private String connectivityType;
        private @Nullable List<GetNatGatewayFilter> filters;
        private String id;
        private String networkInterfaceId;
        private String privateIp;
        private String publicIp;
        private List<String> secondaryAllocationIds;
        private Integer secondaryPrivateIpAddressCount;
        private List<String> secondaryPrivateIpAddresses;
        private String state;
        private String subnetId;
        private Map<String,String> tags;
        private String vpcId;
        public Builder() {}
        public Builder(GetNatGatewayResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allocationId = defaults.allocationId;
    	      this.associationId = defaults.associationId;
    	      this.connectivityType = defaults.connectivityType;
    	      this.filters = defaults.filters;
    	      this.id = defaults.id;
    	      this.networkInterfaceId = defaults.networkInterfaceId;
    	      this.privateIp = defaults.privateIp;
    	      this.publicIp = defaults.publicIp;
    	      this.secondaryAllocationIds = defaults.secondaryAllocationIds;
    	      this.secondaryPrivateIpAddressCount = defaults.secondaryPrivateIpAddressCount;
    	      this.secondaryPrivateIpAddresses = defaults.secondaryPrivateIpAddresses;
    	      this.state = defaults.state;
    	      this.subnetId = defaults.subnetId;
    	      this.tags = defaults.tags;
    	      this.vpcId = defaults.vpcId;
        }

        @CustomType.Setter
        public Builder allocationId(String allocationId) {
            this.allocationId = Objects.requireNonNull(allocationId);
            return this;
        }
        @CustomType.Setter
        public Builder associationId(String associationId) {
            this.associationId = Objects.requireNonNull(associationId);
            return this;
        }
        @CustomType.Setter
        public Builder connectivityType(String connectivityType) {
            this.connectivityType = Objects.requireNonNull(connectivityType);
            return this;
        }
        @CustomType.Setter
        public Builder filters(@Nullable List<GetNatGatewayFilter> filters) {
            this.filters = filters;
            return this;
        }
        public Builder filters(GetNatGatewayFilter... filters) {
            return filters(List.of(filters));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder networkInterfaceId(String networkInterfaceId) {
            this.networkInterfaceId = Objects.requireNonNull(networkInterfaceId);
            return this;
        }
        @CustomType.Setter
        public Builder privateIp(String privateIp) {
            this.privateIp = Objects.requireNonNull(privateIp);
            return this;
        }
        @CustomType.Setter
        public Builder publicIp(String publicIp) {
            this.publicIp = Objects.requireNonNull(publicIp);
            return this;
        }
        @CustomType.Setter
        public Builder secondaryAllocationIds(List<String> secondaryAllocationIds) {
            this.secondaryAllocationIds = Objects.requireNonNull(secondaryAllocationIds);
            return this;
        }
        public Builder secondaryAllocationIds(String... secondaryAllocationIds) {
            return secondaryAllocationIds(List.of(secondaryAllocationIds));
        }
        @CustomType.Setter
        public Builder secondaryPrivateIpAddressCount(Integer secondaryPrivateIpAddressCount) {
            this.secondaryPrivateIpAddressCount = Objects.requireNonNull(secondaryPrivateIpAddressCount);
            return this;
        }
        @CustomType.Setter
        public Builder secondaryPrivateIpAddresses(List<String> secondaryPrivateIpAddresses) {
            this.secondaryPrivateIpAddresses = Objects.requireNonNull(secondaryPrivateIpAddresses);
            return this;
        }
        public Builder secondaryPrivateIpAddresses(String... secondaryPrivateIpAddresses) {
            return secondaryPrivateIpAddresses(List.of(secondaryPrivateIpAddresses));
        }
        @CustomType.Setter
        public Builder state(String state) {
            this.state = Objects.requireNonNull(state);
            return this;
        }
        @CustomType.Setter
        public Builder subnetId(String subnetId) {
            this.subnetId = Objects.requireNonNull(subnetId);
            return this;
        }
        @CustomType.Setter
        public Builder tags(Map<String,String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        @CustomType.Setter
        public Builder vpcId(String vpcId) {
            this.vpcId = Objects.requireNonNull(vpcId);
            return this;
        }
        public GetNatGatewayResult build() {
            final var o = new GetNatGatewayResult();
            o.allocationId = allocationId;
            o.associationId = associationId;
            o.connectivityType = connectivityType;
            o.filters = filters;
            o.id = id;
            o.networkInterfaceId = networkInterfaceId;
            o.privateIp = privateIp;
            o.publicIp = publicIp;
            o.secondaryAllocationIds = secondaryAllocationIds;
            o.secondaryPrivateIpAddressCount = secondaryPrivateIpAddressCount;
            o.secondaryPrivateIpAddresses = secondaryPrivateIpAddresses;
            o.state = state;
            o.subnetId = subnetId;
            o.tags = tags;
            o.vpcId = vpcId;
            return o;
        }
    }
}
