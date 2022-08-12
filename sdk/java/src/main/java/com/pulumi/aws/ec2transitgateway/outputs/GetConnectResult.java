// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2transitgateway.outputs;

import com.pulumi.aws.ec2transitgateway.outputs.GetConnectFilter;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetConnectResult {
    private final @Nullable List<GetConnectFilter> filters;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private final String id;
    /**
     * @return The tunnel protocol
     * 
     */
    private final String protocol;
    /**
     * @return Key-value tags for the EC2 Transit Gateway Connect
     * 
     */
    private final Map<String,String> tags;
    private final String transitGatewayConnectId;
    /**
     * @return EC2 Transit Gateway identifier
     * 
     */
    private final String transitGatewayId;
    /**
     * @return The underlaying VPC attachment
     * 
     */
    private final String transportAttachmentId;

    @CustomType.Constructor
    private GetConnectResult(
        @CustomType.Parameter("filters") @Nullable List<GetConnectFilter> filters,
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("protocol") String protocol,
        @CustomType.Parameter("tags") Map<String,String> tags,
        @CustomType.Parameter("transitGatewayConnectId") String transitGatewayConnectId,
        @CustomType.Parameter("transitGatewayId") String transitGatewayId,
        @CustomType.Parameter("transportAttachmentId") String transportAttachmentId) {
        this.filters = filters;
        this.id = id;
        this.protocol = protocol;
        this.tags = tags;
        this.transitGatewayConnectId = transitGatewayConnectId;
        this.transitGatewayId = transitGatewayId;
        this.transportAttachmentId = transportAttachmentId;
    }

    public List<GetConnectFilter> filters() {
        return this.filters == null ? List.of() : this.filters;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The tunnel protocol
     * 
     */
    public String protocol() {
        return this.protocol;
    }
    /**
     * @return Key-value tags for the EC2 Transit Gateway Connect
     * 
     */
    public Map<String,String> tags() {
        return this.tags;
    }
    public String transitGatewayConnectId() {
        return this.transitGatewayConnectId;
    }
    /**
     * @return EC2 Transit Gateway identifier
     * 
     */
    public String transitGatewayId() {
        return this.transitGatewayId;
    }
    /**
     * @return The underlaying VPC attachment
     * 
     */
    public String transportAttachmentId() {
        return this.transportAttachmentId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetConnectResult defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable List<GetConnectFilter> filters;
        private String id;
        private String protocol;
        private Map<String,String> tags;
        private String transitGatewayConnectId;
        private String transitGatewayId;
        private String transportAttachmentId;

        public Builder() {
    	      // Empty
        }

        public Builder(GetConnectResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.filters = defaults.filters;
    	      this.id = defaults.id;
    	      this.protocol = defaults.protocol;
    	      this.tags = defaults.tags;
    	      this.transitGatewayConnectId = defaults.transitGatewayConnectId;
    	      this.transitGatewayId = defaults.transitGatewayId;
    	      this.transportAttachmentId = defaults.transportAttachmentId;
        }

        public Builder filters(@Nullable List<GetConnectFilter> filters) {
            this.filters = filters;
            return this;
        }
        public Builder filters(GetConnectFilter... filters) {
            return filters(List.of(filters));
        }
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder protocol(String protocol) {
            this.protocol = Objects.requireNonNull(protocol);
            return this;
        }
        public Builder tags(Map<String,String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        public Builder transitGatewayConnectId(String transitGatewayConnectId) {
            this.transitGatewayConnectId = Objects.requireNonNull(transitGatewayConnectId);
            return this;
        }
        public Builder transitGatewayId(String transitGatewayId) {
            this.transitGatewayId = Objects.requireNonNull(transitGatewayId);
            return this;
        }
        public Builder transportAttachmentId(String transportAttachmentId) {
            this.transportAttachmentId = Objects.requireNonNull(transportAttachmentId);
            return this;
        }        public GetConnectResult build() {
            return new GetConnectResult(filters, id, protocol, tags, transitGatewayConnectId, transitGatewayId, transportAttachmentId);
        }
    }
}