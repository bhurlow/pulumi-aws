// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.lb.outputs;

import com.pulumi.aws.lb.outputs.GetLoadBalancerAccessLogs;
import com.pulumi.aws.lb.outputs.GetLoadBalancerConnectionLog;
import com.pulumi.aws.lb.outputs.GetLoadBalancerSubnetMapping;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetLoadBalancerResult {
    private GetLoadBalancerAccessLogs accessLogs;
    private String arn;
    private String arnSuffix;
    private List<GetLoadBalancerConnectionLog> connectionLogs;
    private String customerOwnedIpv4Pool;
    private String desyncMitigationMode;
    private String dnsName;
    private String dnsRecordClientRoutingPolicy;
    private Boolean dropInvalidHeaderFields;
    private Boolean enableCrossZoneLoadBalancing;
    private Boolean enableDeletionProtection;
    private Boolean enableHttp2;
    private Boolean enableTlsVersionAndCipherSuiteHeaders;
    private Boolean enableWafFailOpen;
    private Boolean enableXffClientPort;
    private String enforceSecurityGroupInboundRulesOnPrivateLinkTraffic;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private Integer idleTimeout;
    private Boolean internal;
    private String ipAddressType;
    private String loadBalancerType;
    private String name;
    private Boolean preserveHostHeader;
    private List<String> securityGroups;
    private List<GetLoadBalancerSubnetMapping> subnetMappings;
    private List<String> subnets;
    private Map<String,String> tags;
    private String vpcId;
    private String xffHeaderProcessingMode;
    private String zoneId;

    private GetLoadBalancerResult() {}
    public GetLoadBalancerAccessLogs accessLogs() {
        return this.accessLogs;
    }
    public String arn() {
        return this.arn;
    }
    public String arnSuffix() {
        return this.arnSuffix;
    }
    public List<GetLoadBalancerConnectionLog> connectionLogs() {
        return this.connectionLogs;
    }
    public String customerOwnedIpv4Pool() {
        return this.customerOwnedIpv4Pool;
    }
    public String desyncMitigationMode() {
        return this.desyncMitigationMode;
    }
    public String dnsName() {
        return this.dnsName;
    }
    public String dnsRecordClientRoutingPolicy() {
        return this.dnsRecordClientRoutingPolicy;
    }
    public Boolean dropInvalidHeaderFields() {
        return this.dropInvalidHeaderFields;
    }
    public Boolean enableCrossZoneLoadBalancing() {
        return this.enableCrossZoneLoadBalancing;
    }
    public Boolean enableDeletionProtection() {
        return this.enableDeletionProtection;
    }
    public Boolean enableHttp2() {
        return this.enableHttp2;
    }
    public Boolean enableTlsVersionAndCipherSuiteHeaders() {
        return this.enableTlsVersionAndCipherSuiteHeaders;
    }
    public Boolean enableWafFailOpen() {
        return this.enableWafFailOpen;
    }
    public Boolean enableXffClientPort() {
        return this.enableXffClientPort;
    }
    public String enforceSecurityGroupInboundRulesOnPrivateLinkTraffic() {
        return this.enforceSecurityGroupInboundRulesOnPrivateLinkTraffic;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Integer idleTimeout() {
        return this.idleTimeout;
    }
    public Boolean internal() {
        return this.internal;
    }
    public String ipAddressType() {
        return this.ipAddressType;
    }
    public String loadBalancerType() {
        return this.loadBalancerType;
    }
    public String name() {
        return this.name;
    }
    public Boolean preserveHostHeader() {
        return this.preserveHostHeader;
    }
    public List<String> securityGroups() {
        return this.securityGroups;
    }
    public List<GetLoadBalancerSubnetMapping> subnetMappings() {
        return this.subnetMappings;
    }
    public List<String> subnets() {
        return this.subnets;
    }
    public Map<String,String> tags() {
        return this.tags;
    }
    public String vpcId() {
        return this.vpcId;
    }
    public String xffHeaderProcessingMode() {
        return this.xffHeaderProcessingMode;
    }
    public String zoneId() {
        return this.zoneId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLoadBalancerResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private GetLoadBalancerAccessLogs accessLogs;
        private String arn;
        private String arnSuffix;
        private List<GetLoadBalancerConnectionLog> connectionLogs;
        private String customerOwnedIpv4Pool;
        private String desyncMitigationMode;
        private String dnsName;
        private String dnsRecordClientRoutingPolicy;
        private Boolean dropInvalidHeaderFields;
        private Boolean enableCrossZoneLoadBalancing;
        private Boolean enableDeletionProtection;
        private Boolean enableHttp2;
        private Boolean enableTlsVersionAndCipherSuiteHeaders;
        private Boolean enableWafFailOpen;
        private Boolean enableXffClientPort;
        private String enforceSecurityGroupInboundRulesOnPrivateLinkTraffic;
        private String id;
        private Integer idleTimeout;
        private Boolean internal;
        private String ipAddressType;
        private String loadBalancerType;
        private String name;
        private Boolean preserveHostHeader;
        private List<String> securityGroups;
        private List<GetLoadBalancerSubnetMapping> subnetMappings;
        private List<String> subnets;
        private Map<String,String> tags;
        private String vpcId;
        private String xffHeaderProcessingMode;
        private String zoneId;
        public Builder() {}
        public Builder(GetLoadBalancerResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessLogs = defaults.accessLogs;
    	      this.arn = defaults.arn;
    	      this.arnSuffix = defaults.arnSuffix;
    	      this.connectionLogs = defaults.connectionLogs;
    	      this.customerOwnedIpv4Pool = defaults.customerOwnedIpv4Pool;
    	      this.desyncMitigationMode = defaults.desyncMitigationMode;
    	      this.dnsName = defaults.dnsName;
    	      this.dnsRecordClientRoutingPolicy = defaults.dnsRecordClientRoutingPolicy;
    	      this.dropInvalidHeaderFields = defaults.dropInvalidHeaderFields;
    	      this.enableCrossZoneLoadBalancing = defaults.enableCrossZoneLoadBalancing;
    	      this.enableDeletionProtection = defaults.enableDeletionProtection;
    	      this.enableHttp2 = defaults.enableHttp2;
    	      this.enableTlsVersionAndCipherSuiteHeaders = defaults.enableTlsVersionAndCipherSuiteHeaders;
    	      this.enableWafFailOpen = defaults.enableWafFailOpen;
    	      this.enableXffClientPort = defaults.enableXffClientPort;
    	      this.enforceSecurityGroupInboundRulesOnPrivateLinkTraffic = defaults.enforceSecurityGroupInboundRulesOnPrivateLinkTraffic;
    	      this.id = defaults.id;
    	      this.idleTimeout = defaults.idleTimeout;
    	      this.internal = defaults.internal;
    	      this.ipAddressType = defaults.ipAddressType;
    	      this.loadBalancerType = defaults.loadBalancerType;
    	      this.name = defaults.name;
    	      this.preserveHostHeader = defaults.preserveHostHeader;
    	      this.securityGroups = defaults.securityGroups;
    	      this.subnetMappings = defaults.subnetMappings;
    	      this.subnets = defaults.subnets;
    	      this.tags = defaults.tags;
    	      this.vpcId = defaults.vpcId;
    	      this.xffHeaderProcessingMode = defaults.xffHeaderProcessingMode;
    	      this.zoneId = defaults.zoneId;
        }

        @CustomType.Setter
        public Builder accessLogs(GetLoadBalancerAccessLogs accessLogs) {
            this.accessLogs = Objects.requireNonNull(accessLogs);
            return this;
        }
        @CustomType.Setter
        public Builder arn(String arn) {
            this.arn = Objects.requireNonNull(arn);
            return this;
        }
        @CustomType.Setter
        public Builder arnSuffix(String arnSuffix) {
            this.arnSuffix = Objects.requireNonNull(arnSuffix);
            return this;
        }
        @CustomType.Setter
        public Builder connectionLogs(List<GetLoadBalancerConnectionLog> connectionLogs) {
            this.connectionLogs = Objects.requireNonNull(connectionLogs);
            return this;
        }
        public Builder connectionLogs(GetLoadBalancerConnectionLog... connectionLogs) {
            return connectionLogs(List.of(connectionLogs));
        }
        @CustomType.Setter
        public Builder customerOwnedIpv4Pool(String customerOwnedIpv4Pool) {
            this.customerOwnedIpv4Pool = Objects.requireNonNull(customerOwnedIpv4Pool);
            return this;
        }
        @CustomType.Setter
        public Builder desyncMitigationMode(String desyncMitigationMode) {
            this.desyncMitigationMode = Objects.requireNonNull(desyncMitigationMode);
            return this;
        }
        @CustomType.Setter
        public Builder dnsName(String dnsName) {
            this.dnsName = Objects.requireNonNull(dnsName);
            return this;
        }
        @CustomType.Setter
        public Builder dnsRecordClientRoutingPolicy(String dnsRecordClientRoutingPolicy) {
            this.dnsRecordClientRoutingPolicy = Objects.requireNonNull(dnsRecordClientRoutingPolicy);
            return this;
        }
        @CustomType.Setter
        public Builder dropInvalidHeaderFields(Boolean dropInvalidHeaderFields) {
            this.dropInvalidHeaderFields = Objects.requireNonNull(dropInvalidHeaderFields);
            return this;
        }
        @CustomType.Setter
        public Builder enableCrossZoneLoadBalancing(Boolean enableCrossZoneLoadBalancing) {
            this.enableCrossZoneLoadBalancing = Objects.requireNonNull(enableCrossZoneLoadBalancing);
            return this;
        }
        @CustomType.Setter
        public Builder enableDeletionProtection(Boolean enableDeletionProtection) {
            this.enableDeletionProtection = Objects.requireNonNull(enableDeletionProtection);
            return this;
        }
        @CustomType.Setter
        public Builder enableHttp2(Boolean enableHttp2) {
            this.enableHttp2 = Objects.requireNonNull(enableHttp2);
            return this;
        }
        @CustomType.Setter
        public Builder enableTlsVersionAndCipherSuiteHeaders(Boolean enableTlsVersionAndCipherSuiteHeaders) {
            this.enableTlsVersionAndCipherSuiteHeaders = Objects.requireNonNull(enableTlsVersionAndCipherSuiteHeaders);
            return this;
        }
        @CustomType.Setter
        public Builder enableWafFailOpen(Boolean enableWafFailOpen) {
            this.enableWafFailOpen = Objects.requireNonNull(enableWafFailOpen);
            return this;
        }
        @CustomType.Setter
        public Builder enableXffClientPort(Boolean enableXffClientPort) {
            this.enableXffClientPort = Objects.requireNonNull(enableXffClientPort);
            return this;
        }
        @CustomType.Setter
        public Builder enforceSecurityGroupInboundRulesOnPrivateLinkTraffic(String enforceSecurityGroupInboundRulesOnPrivateLinkTraffic) {
            this.enforceSecurityGroupInboundRulesOnPrivateLinkTraffic = Objects.requireNonNull(enforceSecurityGroupInboundRulesOnPrivateLinkTraffic);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder idleTimeout(Integer idleTimeout) {
            this.idleTimeout = Objects.requireNonNull(idleTimeout);
            return this;
        }
        @CustomType.Setter
        public Builder internal(Boolean internal) {
            this.internal = Objects.requireNonNull(internal);
            return this;
        }
        @CustomType.Setter
        public Builder ipAddressType(String ipAddressType) {
            this.ipAddressType = Objects.requireNonNull(ipAddressType);
            return this;
        }
        @CustomType.Setter
        public Builder loadBalancerType(String loadBalancerType) {
            this.loadBalancerType = Objects.requireNonNull(loadBalancerType);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder preserveHostHeader(Boolean preserveHostHeader) {
            this.preserveHostHeader = Objects.requireNonNull(preserveHostHeader);
            return this;
        }
        @CustomType.Setter
        public Builder securityGroups(List<String> securityGroups) {
            this.securityGroups = Objects.requireNonNull(securityGroups);
            return this;
        }
        public Builder securityGroups(String... securityGroups) {
            return securityGroups(List.of(securityGroups));
        }
        @CustomType.Setter
        public Builder subnetMappings(List<GetLoadBalancerSubnetMapping> subnetMappings) {
            this.subnetMappings = Objects.requireNonNull(subnetMappings);
            return this;
        }
        public Builder subnetMappings(GetLoadBalancerSubnetMapping... subnetMappings) {
            return subnetMappings(List.of(subnetMappings));
        }
        @CustomType.Setter
        public Builder subnets(List<String> subnets) {
            this.subnets = Objects.requireNonNull(subnets);
            return this;
        }
        public Builder subnets(String... subnets) {
            return subnets(List.of(subnets));
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
        @CustomType.Setter
        public Builder xffHeaderProcessingMode(String xffHeaderProcessingMode) {
            this.xffHeaderProcessingMode = Objects.requireNonNull(xffHeaderProcessingMode);
            return this;
        }
        @CustomType.Setter
        public Builder zoneId(String zoneId) {
            this.zoneId = Objects.requireNonNull(zoneId);
            return this;
        }
        public GetLoadBalancerResult build() {
            final var _resultValue = new GetLoadBalancerResult();
            _resultValue.accessLogs = accessLogs;
            _resultValue.arn = arn;
            _resultValue.arnSuffix = arnSuffix;
            _resultValue.connectionLogs = connectionLogs;
            _resultValue.customerOwnedIpv4Pool = customerOwnedIpv4Pool;
            _resultValue.desyncMitigationMode = desyncMitigationMode;
            _resultValue.dnsName = dnsName;
            _resultValue.dnsRecordClientRoutingPolicy = dnsRecordClientRoutingPolicy;
            _resultValue.dropInvalidHeaderFields = dropInvalidHeaderFields;
            _resultValue.enableCrossZoneLoadBalancing = enableCrossZoneLoadBalancing;
            _resultValue.enableDeletionProtection = enableDeletionProtection;
            _resultValue.enableHttp2 = enableHttp2;
            _resultValue.enableTlsVersionAndCipherSuiteHeaders = enableTlsVersionAndCipherSuiteHeaders;
            _resultValue.enableWafFailOpen = enableWafFailOpen;
            _resultValue.enableXffClientPort = enableXffClientPort;
            _resultValue.enforceSecurityGroupInboundRulesOnPrivateLinkTraffic = enforceSecurityGroupInboundRulesOnPrivateLinkTraffic;
            _resultValue.id = id;
            _resultValue.idleTimeout = idleTimeout;
            _resultValue.internal = internal;
            _resultValue.ipAddressType = ipAddressType;
            _resultValue.loadBalancerType = loadBalancerType;
            _resultValue.name = name;
            _resultValue.preserveHostHeader = preserveHostHeader;
            _resultValue.securityGroups = securityGroups;
            _resultValue.subnetMappings = subnetMappings;
            _resultValue.subnets = subnets;
            _resultValue.tags = tags;
            _resultValue.vpcId = vpcId;
            _resultValue.xffHeaderProcessingMode = xffHeaderProcessingMode;
            _resultValue.zoneId = zoneId;
            return _resultValue;
        }
    }
}
