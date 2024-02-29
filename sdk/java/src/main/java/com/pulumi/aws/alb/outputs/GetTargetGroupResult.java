// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.alb.outputs;

import com.pulumi.aws.alb.outputs.GetTargetGroupHealthCheck;
import com.pulumi.aws.alb.outputs.GetTargetGroupStickiness;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetTargetGroupResult {
    private String arn;
    private String arnSuffix;
    private Boolean connectionTermination;
    private String deregistrationDelay;
    private GetTargetGroupHealthCheck healthCheck;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private Boolean lambdaMultiValueHeadersEnabled;
    private List<String> loadBalancerArns;
    private String loadBalancingAlgorithmType;
    private String loadBalancingAnomalyMitigation;
    private String loadBalancingCrossZoneEnabled;
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

    private GetTargetGroupResult() {}
    public String arn() {
        return this.arn;
    }
    public String arnSuffix() {
        return this.arnSuffix;
    }
    public Boolean connectionTermination() {
        return this.connectionTermination;
    }
    public String deregistrationDelay() {
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
    public List<String> loadBalancerArns() {
        return this.loadBalancerArns;
    }
    public String loadBalancingAlgorithmType() {
        return this.loadBalancingAlgorithmType;
    }
    public String loadBalancingAnomalyMitigation() {
        return this.loadBalancingAnomalyMitigation;
    }
    public String loadBalancingCrossZoneEnabled() {
        return this.loadBalancingCrossZoneEnabled;
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
    @CustomType.Builder
    public static final class Builder {
        private String arn;
        private String arnSuffix;
        private Boolean connectionTermination;
        private String deregistrationDelay;
        private GetTargetGroupHealthCheck healthCheck;
        private String id;
        private Boolean lambdaMultiValueHeadersEnabled;
        private List<String> loadBalancerArns;
        private String loadBalancingAlgorithmType;
        private String loadBalancingAnomalyMitigation;
        private String loadBalancingCrossZoneEnabled;
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
        public Builder() {}
        public Builder(GetTargetGroupResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.arn = defaults.arn;
    	      this.arnSuffix = defaults.arnSuffix;
    	      this.connectionTermination = defaults.connectionTermination;
    	      this.deregistrationDelay = defaults.deregistrationDelay;
    	      this.healthCheck = defaults.healthCheck;
    	      this.id = defaults.id;
    	      this.lambdaMultiValueHeadersEnabled = defaults.lambdaMultiValueHeadersEnabled;
    	      this.loadBalancerArns = defaults.loadBalancerArns;
    	      this.loadBalancingAlgorithmType = defaults.loadBalancingAlgorithmType;
    	      this.loadBalancingAnomalyMitigation = defaults.loadBalancingAnomalyMitigation;
    	      this.loadBalancingCrossZoneEnabled = defaults.loadBalancingCrossZoneEnabled;
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

        @CustomType.Setter
        public Builder arn(String arn) {
            if (arn == null) {
              throw new MissingRequiredPropertyException("GetTargetGroupResult", "arn");
            }
            this.arn = arn;
            return this;
        }
        @CustomType.Setter
        public Builder arnSuffix(String arnSuffix) {
            if (arnSuffix == null) {
              throw new MissingRequiredPropertyException("GetTargetGroupResult", "arnSuffix");
            }
            this.arnSuffix = arnSuffix;
            return this;
        }
        @CustomType.Setter
        public Builder connectionTermination(Boolean connectionTermination) {
            if (connectionTermination == null) {
              throw new MissingRequiredPropertyException("GetTargetGroupResult", "connectionTermination");
            }
            this.connectionTermination = connectionTermination;
            return this;
        }
        @CustomType.Setter
        public Builder deregistrationDelay(String deregistrationDelay) {
            if (deregistrationDelay == null) {
              throw new MissingRequiredPropertyException("GetTargetGroupResult", "deregistrationDelay");
            }
            this.deregistrationDelay = deregistrationDelay;
            return this;
        }
        @CustomType.Setter
        public Builder healthCheck(GetTargetGroupHealthCheck healthCheck) {
            if (healthCheck == null) {
              throw new MissingRequiredPropertyException("GetTargetGroupResult", "healthCheck");
            }
            this.healthCheck = healthCheck;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetTargetGroupResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder lambdaMultiValueHeadersEnabled(Boolean lambdaMultiValueHeadersEnabled) {
            if (lambdaMultiValueHeadersEnabled == null) {
              throw new MissingRequiredPropertyException("GetTargetGroupResult", "lambdaMultiValueHeadersEnabled");
            }
            this.lambdaMultiValueHeadersEnabled = lambdaMultiValueHeadersEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder loadBalancerArns(List<String> loadBalancerArns) {
            if (loadBalancerArns == null) {
              throw new MissingRequiredPropertyException("GetTargetGroupResult", "loadBalancerArns");
            }
            this.loadBalancerArns = loadBalancerArns;
            return this;
        }
        public Builder loadBalancerArns(String... loadBalancerArns) {
            return loadBalancerArns(List.of(loadBalancerArns));
        }
        @CustomType.Setter
        public Builder loadBalancingAlgorithmType(String loadBalancingAlgorithmType) {
            if (loadBalancingAlgorithmType == null) {
              throw new MissingRequiredPropertyException("GetTargetGroupResult", "loadBalancingAlgorithmType");
            }
            this.loadBalancingAlgorithmType = loadBalancingAlgorithmType;
            return this;
        }
        @CustomType.Setter
        public Builder loadBalancingAnomalyMitigation(String loadBalancingAnomalyMitigation) {
            if (loadBalancingAnomalyMitigation == null) {
              throw new MissingRequiredPropertyException("GetTargetGroupResult", "loadBalancingAnomalyMitigation");
            }
            this.loadBalancingAnomalyMitigation = loadBalancingAnomalyMitigation;
            return this;
        }
        @CustomType.Setter
        public Builder loadBalancingCrossZoneEnabled(String loadBalancingCrossZoneEnabled) {
            if (loadBalancingCrossZoneEnabled == null) {
              throw new MissingRequiredPropertyException("GetTargetGroupResult", "loadBalancingCrossZoneEnabled");
            }
            this.loadBalancingCrossZoneEnabled = loadBalancingCrossZoneEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetTargetGroupResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder port(Integer port) {
            if (port == null) {
              throw new MissingRequiredPropertyException("GetTargetGroupResult", "port");
            }
            this.port = port;
            return this;
        }
        @CustomType.Setter
        public Builder preserveClientIp(String preserveClientIp) {
            if (preserveClientIp == null) {
              throw new MissingRequiredPropertyException("GetTargetGroupResult", "preserveClientIp");
            }
            this.preserveClientIp = preserveClientIp;
            return this;
        }
        @CustomType.Setter
        public Builder protocol(String protocol) {
            if (protocol == null) {
              throw new MissingRequiredPropertyException("GetTargetGroupResult", "protocol");
            }
            this.protocol = protocol;
            return this;
        }
        @CustomType.Setter
        public Builder protocolVersion(String protocolVersion) {
            if (protocolVersion == null) {
              throw new MissingRequiredPropertyException("GetTargetGroupResult", "protocolVersion");
            }
            this.protocolVersion = protocolVersion;
            return this;
        }
        @CustomType.Setter
        public Builder proxyProtocolV2(Boolean proxyProtocolV2) {
            if (proxyProtocolV2 == null) {
              throw new MissingRequiredPropertyException("GetTargetGroupResult", "proxyProtocolV2");
            }
            this.proxyProtocolV2 = proxyProtocolV2;
            return this;
        }
        @CustomType.Setter
        public Builder slowStart(Integer slowStart) {
            if (slowStart == null) {
              throw new MissingRequiredPropertyException("GetTargetGroupResult", "slowStart");
            }
            this.slowStart = slowStart;
            return this;
        }
        @CustomType.Setter
        public Builder stickiness(GetTargetGroupStickiness stickiness) {
            if (stickiness == null) {
              throw new MissingRequiredPropertyException("GetTargetGroupResult", "stickiness");
            }
            this.stickiness = stickiness;
            return this;
        }
        @CustomType.Setter
        public Builder tags(Map<String,String> tags) {
            if (tags == null) {
              throw new MissingRequiredPropertyException("GetTargetGroupResult", "tags");
            }
            this.tags = tags;
            return this;
        }
        @CustomType.Setter
        public Builder targetType(String targetType) {
            if (targetType == null) {
              throw new MissingRequiredPropertyException("GetTargetGroupResult", "targetType");
            }
            this.targetType = targetType;
            return this;
        }
        @CustomType.Setter
        public Builder vpcId(String vpcId) {
            if (vpcId == null) {
              throw new MissingRequiredPropertyException("GetTargetGroupResult", "vpcId");
            }
            this.vpcId = vpcId;
            return this;
        }
        public GetTargetGroupResult build() {
            final var _resultValue = new GetTargetGroupResult();
            _resultValue.arn = arn;
            _resultValue.arnSuffix = arnSuffix;
            _resultValue.connectionTermination = connectionTermination;
            _resultValue.deregistrationDelay = deregistrationDelay;
            _resultValue.healthCheck = healthCheck;
            _resultValue.id = id;
            _resultValue.lambdaMultiValueHeadersEnabled = lambdaMultiValueHeadersEnabled;
            _resultValue.loadBalancerArns = loadBalancerArns;
            _resultValue.loadBalancingAlgorithmType = loadBalancingAlgorithmType;
            _resultValue.loadBalancingAnomalyMitigation = loadBalancingAnomalyMitigation;
            _resultValue.loadBalancingCrossZoneEnabled = loadBalancingCrossZoneEnabled;
            _resultValue.name = name;
            _resultValue.port = port;
            _resultValue.preserveClientIp = preserveClientIp;
            _resultValue.protocol = protocol;
            _resultValue.protocolVersion = protocolVersion;
            _resultValue.proxyProtocolV2 = proxyProtocolV2;
            _resultValue.slowStart = slowStart;
            _resultValue.stickiness = stickiness;
            _resultValue.tags = tags;
            _resultValue.targetType = targetType;
            _resultValue.vpcId = vpcId;
            return _resultValue;
        }
    }
}
