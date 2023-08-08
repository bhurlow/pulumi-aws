// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2.outputs;

import com.pulumi.aws.ec2.outputs.GetDedicatedHostFilter;
import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetDedicatedHostResult {
    /**
     * @return ARN of the Dedicated Host.
     * 
     */
    private String arn;
    /**
     * @return The ID of the Outpost hardware asset on which the Dedicated Host is allocated.
     * 
     */
    private String assetId;
    /**
     * @return Whether auto-placement is on or off.
     * 
     */
    private String autoPlacement;
    /**
     * @return Availability Zone of the Dedicated Host.
     * 
     */
    private String availabilityZone;
    /**
     * @return Number of cores on the Dedicated Host.
     * 
     */
    private Integer cores;
    private @Nullable List<GetDedicatedHostFilter> filters;
    private String hostId;
    /**
     * @return Whether host recovery is enabled or disabled for the Dedicated Host.
     * 
     */
    private String hostRecovery;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Instance family supported by the Dedicated Host. For example, &#34;m5&#34;.
     * 
     */
    private String instanceFamily;
    /**
     * @return Instance type supported by the Dedicated Host. For example, &#34;m5.large&#34;. If the host supports multiple instance types, no instanceType is returned.
     * 
     */
    private String instanceType;
    /**
     * @return ARN of the AWS Outpost on which the Dedicated Host is allocated.
     * 
     */
    private String outpostArn;
    /**
     * @return ID of the AWS account that owns the Dedicated Host.
     * 
     */
    private String ownerId;
    /**
     * @return Number of sockets on the Dedicated Host.
     * 
     */
    private Integer sockets;
    private Map<String,String> tags;
    /**
     * @return Total number of vCPUs on the Dedicated Host.
     * 
     */
    private Integer totalVcpus;

    private GetDedicatedHostResult() {}
    /**
     * @return ARN of the Dedicated Host.
     * 
     */
    public String arn() {
        return this.arn;
    }
    /**
     * @return The ID of the Outpost hardware asset on which the Dedicated Host is allocated.
     * 
     */
    public String assetId() {
        return this.assetId;
    }
    /**
     * @return Whether auto-placement is on or off.
     * 
     */
    public String autoPlacement() {
        return this.autoPlacement;
    }
    /**
     * @return Availability Zone of the Dedicated Host.
     * 
     */
    public String availabilityZone() {
        return this.availabilityZone;
    }
    /**
     * @return Number of cores on the Dedicated Host.
     * 
     */
    public Integer cores() {
        return this.cores;
    }
    public List<GetDedicatedHostFilter> filters() {
        return this.filters == null ? List.of() : this.filters;
    }
    public String hostId() {
        return this.hostId;
    }
    /**
     * @return Whether host recovery is enabled or disabled for the Dedicated Host.
     * 
     */
    public String hostRecovery() {
        return this.hostRecovery;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Instance family supported by the Dedicated Host. For example, &#34;m5&#34;.
     * 
     */
    public String instanceFamily() {
        return this.instanceFamily;
    }
    /**
     * @return Instance type supported by the Dedicated Host. For example, &#34;m5.large&#34;. If the host supports multiple instance types, no instanceType is returned.
     * 
     */
    public String instanceType() {
        return this.instanceType;
    }
    /**
     * @return ARN of the AWS Outpost on which the Dedicated Host is allocated.
     * 
     */
    public String outpostArn() {
        return this.outpostArn;
    }
    /**
     * @return ID of the AWS account that owns the Dedicated Host.
     * 
     */
    public String ownerId() {
        return this.ownerId;
    }
    /**
     * @return Number of sockets on the Dedicated Host.
     * 
     */
    public Integer sockets() {
        return this.sockets;
    }
    public Map<String,String> tags() {
        return this.tags;
    }
    /**
     * @return Total number of vCPUs on the Dedicated Host.
     * 
     */
    public Integer totalVcpus() {
        return this.totalVcpus;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDedicatedHostResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String arn;
        private String assetId;
        private String autoPlacement;
        private String availabilityZone;
        private Integer cores;
        private @Nullable List<GetDedicatedHostFilter> filters;
        private String hostId;
        private String hostRecovery;
        private String id;
        private String instanceFamily;
        private String instanceType;
        private String outpostArn;
        private String ownerId;
        private Integer sockets;
        private Map<String,String> tags;
        private Integer totalVcpus;
        public Builder() {}
        public Builder(GetDedicatedHostResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.arn = defaults.arn;
    	      this.assetId = defaults.assetId;
    	      this.autoPlacement = defaults.autoPlacement;
    	      this.availabilityZone = defaults.availabilityZone;
    	      this.cores = defaults.cores;
    	      this.filters = defaults.filters;
    	      this.hostId = defaults.hostId;
    	      this.hostRecovery = defaults.hostRecovery;
    	      this.id = defaults.id;
    	      this.instanceFamily = defaults.instanceFamily;
    	      this.instanceType = defaults.instanceType;
    	      this.outpostArn = defaults.outpostArn;
    	      this.ownerId = defaults.ownerId;
    	      this.sockets = defaults.sockets;
    	      this.tags = defaults.tags;
    	      this.totalVcpus = defaults.totalVcpus;
        }

        @CustomType.Setter
        public Builder arn(String arn) {
            this.arn = Objects.requireNonNull(arn);
            return this;
        }
        @CustomType.Setter
        public Builder assetId(String assetId) {
            this.assetId = Objects.requireNonNull(assetId);
            return this;
        }
        @CustomType.Setter
        public Builder autoPlacement(String autoPlacement) {
            this.autoPlacement = Objects.requireNonNull(autoPlacement);
            return this;
        }
        @CustomType.Setter
        public Builder availabilityZone(String availabilityZone) {
            this.availabilityZone = Objects.requireNonNull(availabilityZone);
            return this;
        }
        @CustomType.Setter
        public Builder cores(Integer cores) {
            this.cores = Objects.requireNonNull(cores);
            return this;
        }
        @CustomType.Setter
        public Builder filters(@Nullable List<GetDedicatedHostFilter> filters) {
            this.filters = filters;
            return this;
        }
        public Builder filters(GetDedicatedHostFilter... filters) {
            return filters(List.of(filters));
        }
        @CustomType.Setter
        public Builder hostId(String hostId) {
            this.hostId = Objects.requireNonNull(hostId);
            return this;
        }
        @CustomType.Setter
        public Builder hostRecovery(String hostRecovery) {
            this.hostRecovery = Objects.requireNonNull(hostRecovery);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder instanceFamily(String instanceFamily) {
            this.instanceFamily = Objects.requireNonNull(instanceFamily);
            return this;
        }
        @CustomType.Setter
        public Builder instanceType(String instanceType) {
            this.instanceType = Objects.requireNonNull(instanceType);
            return this;
        }
        @CustomType.Setter
        public Builder outpostArn(String outpostArn) {
            this.outpostArn = Objects.requireNonNull(outpostArn);
            return this;
        }
        @CustomType.Setter
        public Builder ownerId(String ownerId) {
            this.ownerId = Objects.requireNonNull(ownerId);
            return this;
        }
        @CustomType.Setter
        public Builder sockets(Integer sockets) {
            this.sockets = Objects.requireNonNull(sockets);
            return this;
        }
        @CustomType.Setter
        public Builder tags(Map<String,String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        @CustomType.Setter
        public Builder totalVcpus(Integer totalVcpus) {
            this.totalVcpus = Objects.requireNonNull(totalVcpus);
            return this;
        }
        public GetDedicatedHostResult build() {
            final var o = new GetDedicatedHostResult();
            o.arn = arn;
            o.assetId = assetId;
            o.autoPlacement = autoPlacement;
            o.availabilityZone = availabilityZone;
            o.cores = cores;
            o.filters = filters;
            o.hostId = hostId;
            o.hostRecovery = hostRecovery;
            o.id = id;
            o.instanceFamily = instanceFamily;
            o.instanceType = instanceType;
            o.outpostArn = outpostArn;
            o.ownerId = ownerId;
            o.sockets = sockets;
            o.tags = tags;
            o.totalVcpus = totalVcpus;
            return o;
        }
    }
}
