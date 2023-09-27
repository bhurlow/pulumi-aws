// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.fsx.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetOntapStorageVirtualMachineEndpointNf {
    private String dnsName;
    private List<String> ipAddresses;

    private GetOntapStorageVirtualMachineEndpointNf() {}
    public String dnsName() {
        return this.dnsName;
    }
    public List<String> ipAddresses() {
        return this.ipAddresses;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOntapStorageVirtualMachineEndpointNf defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String dnsName;
        private List<String> ipAddresses;
        public Builder() {}
        public Builder(GetOntapStorageVirtualMachineEndpointNf defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dnsName = defaults.dnsName;
    	      this.ipAddresses = defaults.ipAddresses;
        }

        @CustomType.Setter
        public Builder dnsName(String dnsName) {
            this.dnsName = Objects.requireNonNull(dnsName);
            return this;
        }
        @CustomType.Setter
        public Builder ipAddresses(List<String> ipAddresses) {
            this.ipAddresses = Objects.requireNonNull(ipAddresses);
            return this;
        }
        public Builder ipAddresses(String... ipAddresses) {
            return ipAddresses(List.of(ipAddresses));
        }
        public GetOntapStorageVirtualMachineEndpointNf build() {
            final var o = new GetOntapStorageVirtualMachineEndpointNf();
            o.dnsName = dnsName;
            o.ipAddresses = ipAddresses;
            return o;
        }
    }
}
