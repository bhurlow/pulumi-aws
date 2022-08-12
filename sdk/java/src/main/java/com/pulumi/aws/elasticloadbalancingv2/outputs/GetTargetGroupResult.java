// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.elasticloadbalancingv2.outputs;

import com.pulumi.aws.elasticloadbalancingv2.outputs.GetTargetGroupHealthCheck;
import com.pulumi.aws.elasticloadbalancingv2.outputs.GetTargetGroupStickiness;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetTargetGroupResult {
    private final String arn;
    private final String arnSuffix;
    private final Boolean connectionTermination;
    private final Integer deregistrationDelay;
    private final GetTargetGroupHealthCheck healthCheck;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private final String id;
    private final Boolean lambdaMultiValueHeadersEnabled;
    private final String loadBalancingAlgorithmType;
    private final String name;
    private final Integer port;
    private final String preserveClientIp;
    private final String protocol;
    private final String protocolVersion;
    private final Boolean proxyProtocolV2;
    private final Integer slowStart;
    private final GetTargetGroupStickiness stickiness;
    private final Map<String,String> tags;
    private final String targetType;
    private final String vpcId;

    @CustomType.Constructor
    private GetTargetGroupResult(
        @CustomType.Parameter("arn") String arn,
        @CustomType.Parameter("arnSuffix") String arnSuffix,
        @CustomType.Parameter("connectionTermination") Boolean connectionTermination,
        @CustomType.Parameter("deregistrationDelay") Integer deregistrationDelay,
        @CustomType.Parameter("healthCheck") GetTargetGroupHealthCheck healthCheck,
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("lambdaMultiValueHeadersEnabled") Boolean lambdaMultiValueHeadersEnabled,
        @CustomType.Parameter("loadBalancingAlgorithmType") String loadBalancingAlgorithmType,
        @CustomType.Parameter("name") String name,
        @CustomType.Parameter("port") Integer port,
        @CustomType.Parameter("preserveClientIp") String preserveClientIp,
        @CustomType.Parameter("protocol") String protocol,
        @CustomType.Parameter("protocolVersion") String protocolVersion,
        @CustomType.Parameter("proxyProtocolV2") Boolean proxyProtocolV2,
        @CustomType.Parameter("slowStart") Integer slowStart,
        @CustomType.Parameter("stickiness") GetTargetGroupStickiness stickiness,
        @CustomType.Parameter("tags") Map<String,String> tags,
        @CustomType.Parameter("targetType") String targetType,
        @CustomType.Parameter("vpcId") String vpcId) {
        this.arn = arn;
        this.arnSuffix = arnSuffix;
        this.connectionTermination = connectionTermination;
        this.deregistrationDelay = deregistrationDelay;
        this.healthCheck = healthCheck;
        this.id = id;
        this.lambdaMultiValueHeadersEnabled = lambdaMultiValueHeadersEnabled;
        this.loadBalancingAlgorithmType = loadBalancingAlgorithmType;
        this.name = name;
        this.port = port;
        this.preserveClientIp = preserveClientIp;
        this.protocol = protocol;
        this.protocolVersion = protocolVersion;
        this.proxyProtocolV2 = proxyProtocolV2;
        this.slowStart = slowStart;
        this.stickiness = stickiness;
        this.tags = tags;
        this.targetType = targetType;
        this.vpcId = vpcId;
    }

    public String arn() {
        return this.arn;
    }
    public String arnSuffix() {
        return this.arnSuffix;
    }
    public Boolean connectionTermination() {
        return this.connectionTermination;
    }
    public Integer deregistrationDelay() {
        return this.deregistrationDelay;
    }
    public GetTargetGroupHealthCheck healthCheck() {
        return this.healthCheck;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Boolean lambdaMultiValueHeadersEnabled() {
        return this.lambdaMultiValueHeadersEnabled;
    }
    public String loadBalancingAlgorithmType() {
        return this.loadBalancingAlgorithmType;
    }
    public String name() {
        return this.name;
    }
    public Integer port() {
        return this.port;
    }
    public String preserveClientIp() {
        return this.preserveClientIp;
    }
    public String protocol() {
        return this.protocol;
    }
    public String protocolVersion() {
        return this.protocolVersion;
    }
    public Boolean proxyProtocolV2() {
        return this.proxyProtocolV2;
    }
    public Integer slowStart() {
        return this.slowStart;
    }
    public GetTargetGroupStickiness stickiness() {
        return this.stickiness;
    }
    public Map<String,String> tags() {
        return this.tags;
    }
    public String targetType() {
        return this.targetType;
    }
    public String vpcId() {
        return this.vpcId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTargetGroupResult defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String arn;
        private String arnSuffix;
        private Boolean connectionTermination;
        private Integer deregistrationDelay;
        private GetTargetGroupHealthCheck healthCheck;
        private String id;
        private Boolean lambdaMultiValueHeadersEnabled;
        private String loadBalancingAlgorithmType;
        private String name;
        private Integer port;
        private String preserveClientIp;
        private String protocol;
        private String protocolVersion;
        private Boolean proxyProtocolV2;
        private Integer slowStart;
        private GetTargetGroupStickiness stickiness;
        private Map<String,String> tags;
        private String targetType;
        private String vpcId;

        public Builder() {
    	      // Empty
        }

        public Builder(GetTargetGroupResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.arn = defaults.arn;
    	      this.arnSuffix = defaults.arnSuffix;
    	      this.connectionTermination = defaults.connectionTermination;
    	      this.deregistrationDelay = defaults.deregistrationDelay;
    	      this.healthCheck = defaults.healthCheck;
    	      this.id = defaults.id;
    	      this.lambdaMultiValueHeadersEnabled = defaults.lambdaMultiValueHeadersEnabled;
    	      this.loadBalancingAlgorithmType = defaults.loadBalancingAlgorithmType;
    	      this.name = defaults.name;
    	      this.port = defaults.port;
    	      this.preserveClientIp = defaults.preserveClientIp;
    	      this.protocol = defaults.protocol;
    	      this.protocolVersion = defaults.protocolVersion;
    	      this.proxyProtocolV2 = defaults.proxyProtocolV2;
    	      this.slowStart = defaults.slowStart;
    	      this.stickiness = defaults.stickiness;
    	      this.tags = defaults.tags;
    	      this.targetType = defaults.targetType;
    	      this.vpcId = defaults.vpcId;
        }

        public Builder arn(String arn) {
            this.arn = Objects.requireNonNull(arn);
            return this;
        }
        public Builder arnSuffix(String arnSuffix) {
            this.arnSuffix = Objects.requireNonNull(arnSuffix);
            return this;
        }
        public Builder connectionTermination(Boolean connectionTermination) {
            this.connectionTermination = Objects.requireNonNull(connectionTermination);
            return this;
        }
        public Builder deregistrationDelay(Integer deregistrationDelay) {
            this.deregistrationDelay = Objects.requireNonNull(deregistrationDelay);
            return this;
        }
        public Builder healthCheck(GetTargetGroupHealthCheck healthCheck) {
            this.healthCheck = Objects.requireNonNull(healthCheck);
            return this;
        }
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder lambdaMultiValueHeadersEnabled(Boolean lambdaMultiValueHeadersEnabled) {
            this.lambdaMultiValueHeadersEnabled = Objects.requireNonNull(lambdaMultiValueHeadersEnabled);
            return this;
        }
        public Builder loadBalancingAlgorithmType(String loadBalancingAlgorithmType) {
            this.loadBalancingAlgorithmType = Objects.requireNonNull(loadBalancingAlgorithmType);
            return this;
        }
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public Builder port(Integer port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }
        public Builder preserveClientIp(String preserveClientIp) {
            this.preserveClientIp = Objects.requireNonNull(preserveClientIp);
            return this;
        }
        public Builder protocol(String protocol) {
            this.protocol = Objects.requireNonNull(protocol);
            return this;
        }
        public Builder protocolVersion(String protocolVersion) {
            this.protocolVersion = Objects.requireNonNull(protocolVersion);
            return this;
        }
        public Builder proxyProtocolV2(Boolean proxyProtocolV2) {
            this.proxyProtocolV2 = Objects.requireNonNull(proxyProtocolV2);
            return this;
        }
        public Builder slowStart(Integer slowStart) {
            this.slowStart = Objects.requireNonNull(slowStart);
            return this;
        }
        public Builder stickiness(GetTargetGroupStickiness stickiness) {
            this.stickiness = Objects.requireNonNull(stickiness);
            return this;
        }
        public Builder tags(Map<String,String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        public Builder targetType(String targetType) {
            this.targetType = Objects.requireNonNull(targetType);
            return this;
        }
        public Builder vpcId(String vpcId) {
            this.vpcId = Objects.requireNonNull(vpcId);
            return this;
        }        public GetTargetGroupResult build() {
            return new GetTargetGroupResult(arn, arnSuffix, connectionTermination, deregistrationDelay, healthCheck, id, lambdaMultiValueHeadersEnabled, loadBalancingAlgorithmType, name, port, preserveClientIp, protocol, protocolVersion, proxyProtocolV2, slowStart, stickiness, tags, targetType, vpcId);
        }
    }
}