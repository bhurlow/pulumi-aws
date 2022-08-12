// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class InstancePrivateDnsNameOptions {
    /**
     * @return Indicates whether to respond to DNS queries for instance hostnames with DNS A records.
     * 
     */
    private final @Nullable Boolean enableResourceNameDnsARecord;
    /**
     * @return Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records.
     * 
     */
    private final @Nullable Boolean enableResourceNameDnsAaaaRecord;
    /**
     * @return The type of hostname for Amazon EC2 instances. For IPv4 only subnets, an instance DNS name must be based on the instance IPv4 address. For IPv6 native subnets, an instance DNS name must be based on the instance ID. For dual-stack subnets, you can specify whether DNS names use the instance IPv4 address or the instance ID. Valid values: `ip-name` and `resource-name`.
     * 
     */
    private final @Nullable String hostnameType;

    @CustomType.Constructor
    private InstancePrivateDnsNameOptions(
        @CustomType.Parameter("enableResourceNameDnsARecord") @Nullable Boolean enableResourceNameDnsARecord,
        @CustomType.Parameter("enableResourceNameDnsAaaaRecord") @Nullable Boolean enableResourceNameDnsAaaaRecord,
        @CustomType.Parameter("hostnameType") @Nullable String hostnameType) {
        this.enableResourceNameDnsARecord = enableResourceNameDnsARecord;
        this.enableResourceNameDnsAaaaRecord = enableResourceNameDnsAaaaRecord;
        this.hostnameType = hostnameType;
    }

    /**
     * @return Indicates whether to respond to DNS queries for instance hostnames with DNS A records.
     * 
     */
    public Optional<Boolean> enableResourceNameDnsARecord() {
        return Optional.ofNullable(this.enableResourceNameDnsARecord);
    }
    /**
     * @return Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records.
     * 
     */
    public Optional<Boolean> enableResourceNameDnsAaaaRecord() {
        return Optional.ofNullable(this.enableResourceNameDnsAaaaRecord);
    }
    /**
     * @return The type of hostname for Amazon EC2 instances. For IPv4 only subnets, an instance DNS name must be based on the instance IPv4 address. For IPv6 native subnets, an instance DNS name must be based on the instance ID. For dual-stack subnets, you can specify whether DNS names use the instance IPv4 address or the instance ID. Valid values: `ip-name` and `resource-name`.
     * 
     */
    public Optional<String> hostnameType() {
        return Optional.ofNullable(this.hostnameType);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InstancePrivateDnsNameOptions defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable Boolean enableResourceNameDnsARecord;
        private @Nullable Boolean enableResourceNameDnsAaaaRecord;
        private @Nullable String hostnameType;

        public Builder() {
    	      // Empty
        }

        public Builder(InstancePrivateDnsNameOptions defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enableResourceNameDnsARecord = defaults.enableResourceNameDnsARecord;
    	      this.enableResourceNameDnsAaaaRecord = defaults.enableResourceNameDnsAaaaRecord;
    	      this.hostnameType = defaults.hostnameType;
        }

        public Builder enableResourceNameDnsARecord(@Nullable Boolean enableResourceNameDnsARecord) {
            this.enableResourceNameDnsARecord = enableResourceNameDnsARecord;
            return this;
        }
        public Builder enableResourceNameDnsAaaaRecord(@Nullable Boolean enableResourceNameDnsAaaaRecord) {
            this.enableResourceNameDnsAaaaRecord = enableResourceNameDnsAaaaRecord;
            return this;
        }
        public Builder hostnameType(@Nullable String hostnameType) {
            this.hostnameType = hostnameType;
            return this;
        }        public InstancePrivateDnsNameOptions build() {
            return new InstancePrivateDnsNameOptions(enableResourceNameDnsARecord, enableResourceNameDnsAaaaRecord, hostnameType);
        }
    }
}