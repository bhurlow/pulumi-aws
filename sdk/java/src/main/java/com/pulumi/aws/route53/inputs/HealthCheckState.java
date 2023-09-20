// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.route53.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class HealthCheckState extends com.pulumi.resources.ResourceArgs {

    public static final HealthCheckState Empty = new HealthCheckState();

    /**
     * The Amazon Resource Name (ARN) of the Health Check.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) of the Health Check.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The minimum number of child health checks that must be healthy for Route 53 to consider the parent health check to be healthy. Valid values are integers between 0 and 256, inclusive
     * 
     */
    @Import(name="childHealthThreshold")
    private @Nullable Output<Integer> childHealthThreshold;

    /**
     * @return The minimum number of child health checks that must be healthy for Route 53 to consider the parent health check to be healthy. Valid values are integers between 0 and 256, inclusive
     * 
     */
    public Optional<Output<Integer>> childHealthThreshold() {
        return Optional.ofNullable(this.childHealthThreshold);
    }

    /**
     * For a specified parent health check, a list of HealthCheckId values for the associated child health checks.
     * 
     */
    @Import(name="childHealthchecks")
    private @Nullable Output<List<String>> childHealthchecks;

    /**
     * @return For a specified parent health check, a list of HealthCheckId values for the associated child health checks.
     * 
     */
    public Optional<Output<List<String>>> childHealthchecks() {
        return Optional.ofNullable(this.childHealthchecks);
    }

    /**
     * The name of the CloudWatch alarm.
     * 
     */
    @Import(name="cloudwatchAlarmName")
    private @Nullable Output<String> cloudwatchAlarmName;

    /**
     * @return The name of the CloudWatch alarm.
     * 
     */
    public Optional<Output<String>> cloudwatchAlarmName() {
        return Optional.ofNullable(this.cloudwatchAlarmName);
    }

    /**
     * The CloudWatchRegion that the CloudWatch alarm was created in.
     * 
     */
    @Import(name="cloudwatchAlarmRegion")
    private @Nullable Output<String> cloudwatchAlarmRegion;

    /**
     * @return The CloudWatchRegion that the CloudWatch alarm was created in.
     * 
     */
    public Optional<Output<String>> cloudwatchAlarmRegion() {
        return Optional.ofNullable(this.cloudwatchAlarmRegion);
    }

    /**
     * A boolean value that stops Route 53 from performing health checks. When set to true, Route 53 will do the following depending on the type of health check:
     * * For health checks that check the health of endpoints, Route5 53 stops submitting requests to your application, server, or other resource.
     * * For calculated health checks, Route 53 stops aggregating the status of the referenced health checks.
     * * For health checks that monitor CloudWatch alarms, Route 53 stops monitoring the corresponding CloudWatch metrics.
     * 
     * &gt; **Note:** After you disable a health check, Route 53 considers the status of the health check to always be healthy. If you configured DNS failover, Route 53 continues to route traffic to the corresponding resources. If you want to stop routing traffic to a resource, change the value of `invert_healthcheck`.
     * 
     */
    @Import(name="disabled")
    private @Nullable Output<Boolean> disabled;

    /**
     * @return A boolean value that stops Route 53 from performing health checks. When set to true, Route 53 will do the following depending on the type of health check:
     * * For health checks that check the health of endpoints, Route5 53 stops submitting requests to your application, server, or other resource.
     * * For calculated health checks, Route 53 stops aggregating the status of the referenced health checks.
     * * For health checks that monitor CloudWatch alarms, Route 53 stops monitoring the corresponding CloudWatch metrics.
     * 
     * &gt; **Note:** After you disable a health check, Route 53 considers the status of the health check to always be healthy. If you configured DNS failover, Route 53 continues to route traffic to the corresponding resources. If you want to stop routing traffic to a resource, change the value of `invert_healthcheck`.
     * 
     */
    public Optional<Output<Boolean>> disabled() {
        return Optional.ofNullable(this.disabled);
    }

    /**
     * A boolean value that indicates whether Route53 should send the `fqdn` to the endpoint when performing the health check. This defaults to AWS&#39; defaults: when the `type` is &#34;HTTPS&#34; `enable_sni` defaults to `true`, when `type` is anything else `enable_sni` defaults to `false`.
     * 
     */
    @Import(name="enableSni")
    private @Nullable Output<Boolean> enableSni;

    /**
     * @return A boolean value that indicates whether Route53 should send the `fqdn` to the endpoint when performing the health check. This defaults to AWS&#39; defaults: when the `type` is &#34;HTTPS&#34; `enable_sni` defaults to `true`, when `type` is anything else `enable_sni` defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> enableSni() {
        return Optional.ofNullable(this.enableSni);
    }

    /**
     * The number of consecutive health checks that an endpoint must pass or fail.
     * 
     */
    @Import(name="failureThreshold")
    private @Nullable Output<Integer> failureThreshold;

    /**
     * @return The number of consecutive health checks that an endpoint must pass or fail.
     * 
     */
    public Optional<Output<Integer>> failureThreshold() {
        return Optional.ofNullable(this.failureThreshold);
    }

    /**
     * The fully qualified domain name of the endpoint to be checked. If a value is set for `ip_address`, the value set for `fqdn` will be passed in the `Host` header.
     * 
     */
    @Import(name="fqdn")
    private @Nullable Output<String> fqdn;

    /**
     * @return The fully qualified domain name of the endpoint to be checked. If a value is set for `ip_address`, the value set for `fqdn` will be passed in the `Host` header.
     * 
     */
    public Optional<Output<String>> fqdn() {
        return Optional.ofNullable(this.fqdn);
    }

    /**
     * The status of the health check when CloudWatch has insufficient data about the state of associated alarm. Valid values are `Healthy` , `Unhealthy` and `LastKnownStatus`.
     * 
     */
    @Import(name="insufficientDataHealthStatus")
    private @Nullable Output<String> insufficientDataHealthStatus;

    /**
     * @return The status of the health check when CloudWatch has insufficient data about the state of associated alarm. Valid values are `Healthy` , `Unhealthy` and `LastKnownStatus`.
     * 
     */
    public Optional<Output<String>> insufficientDataHealthStatus() {
        return Optional.ofNullable(this.insufficientDataHealthStatus);
    }

    /**
     * A boolean value that indicates whether the status of health check should be inverted. For example, if a health check is healthy but Inverted is True , then Route 53 considers the health check to be unhealthy.
     * 
     */
    @Import(name="invertHealthcheck")
    private @Nullable Output<Boolean> invertHealthcheck;

    /**
     * @return A boolean value that indicates whether the status of health check should be inverted. For example, if a health check is healthy but Inverted is True , then Route 53 considers the health check to be unhealthy.
     * 
     */
    public Optional<Output<Boolean>> invertHealthcheck() {
        return Optional.ofNullable(this.invertHealthcheck);
    }

    /**
     * The IP address of the endpoint to be checked.
     * 
     */
    @Import(name="ipAddress")
    private @Nullable Output<String> ipAddress;

    /**
     * @return The IP address of the endpoint to be checked.
     * 
     */
    public Optional<Output<String>> ipAddress() {
        return Optional.ofNullable(this.ipAddress);
    }

    /**
     * A Boolean value that indicates whether you want Route 53 to measure the latency between health checkers in multiple AWS regions and your endpoint and to display CloudWatch latency graphs in the Route 53 console.
     * 
     */
    @Import(name="measureLatency")
    private @Nullable Output<Boolean> measureLatency;

    /**
     * @return A Boolean value that indicates whether you want Route 53 to measure the latency between health checkers in multiple AWS regions and your endpoint and to display CloudWatch latency graphs in the Route 53 console.
     * 
     */
    public Optional<Output<Boolean>> measureLatency() {
        return Optional.ofNullable(this.measureLatency);
    }

    /**
     * The port of the endpoint to be checked.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return The port of the endpoint to be checked.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * This is a reference name used in Caller Reference
     * (helpful for identifying single health_check set amongst others)
     * 
     */
    @Import(name="referenceName")
    private @Nullable Output<String> referenceName;

    /**
     * @return This is a reference name used in Caller Reference
     * (helpful for identifying single health_check set amongst others)
     * 
     */
    public Optional<Output<String>> referenceName() {
        return Optional.ofNullable(this.referenceName);
    }

    /**
     * A list of AWS regions that you want Amazon Route 53 health checkers to check the specified endpoint from.
     * 
     */
    @Import(name="regions")
    private @Nullable Output<List<String>> regions;

    /**
     * @return A list of AWS regions that you want Amazon Route 53 health checkers to check the specified endpoint from.
     * 
     */
    public Optional<Output<List<String>>> regions() {
        return Optional.ofNullable(this.regions);
    }

    /**
     * The number of seconds between the time that Amazon Route 53 gets a response from your endpoint and the time that it sends the next health-check request.
     * 
     */
    @Import(name="requestInterval")
    private @Nullable Output<Integer> requestInterval;

    /**
     * @return The number of seconds between the time that Amazon Route 53 gets a response from your endpoint and the time that it sends the next health-check request.
     * 
     */
    public Optional<Output<Integer>> requestInterval() {
        return Optional.ofNullable(this.requestInterval);
    }

    /**
     * The path that you want Amazon Route 53 to request when performing health checks.
     * 
     */
    @Import(name="resourcePath")
    private @Nullable Output<String> resourcePath;

    /**
     * @return The path that you want Amazon Route 53 to request when performing health checks.
     * 
     */
    public Optional<Output<String>> resourcePath() {
        return Optional.ofNullable(this.resourcePath);
    }

    /**
     * The Amazon Resource Name (ARN) for the Route 53 Application Recovery Controller routing control. This is used when health check type is `RECOVERY_CONTROL`
     * 
     */
    @Import(name="routingControlArn")
    private @Nullable Output<String> routingControlArn;

    /**
     * @return The Amazon Resource Name (ARN) for the Route 53 Application Recovery Controller routing control. This is used when health check type is `RECOVERY_CONTROL`
     * 
     */
    public Optional<Output<String>> routingControlArn() {
        return Optional.ofNullable(this.routingControlArn);
    }

    /**
     * String searched in the first 5120 bytes of the response body for check to be considered healthy. Only valid with `HTTP_STR_MATCH` and `HTTPS_STR_MATCH`.
     * 
     */
    @Import(name="searchString")
    private @Nullable Output<String> searchString;

    /**
     * @return String searched in the first 5120 bytes of the response body for check to be considered healthy. Only valid with `HTTP_STR_MATCH` and `HTTPS_STR_MATCH`.
     * 
     */
    public Optional<Output<String>> searchString() {
        return Optional.ofNullable(this.searchString);
    }

    /**
     * A map of tags to assign to the health check. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the health check. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    /**
     * The protocol to use when performing health checks. Valid values are `HTTP`, `HTTPS`, `HTTP_STR_MATCH`, `HTTPS_STR_MATCH`, `TCP`, `CALCULATED`, `CLOUDWATCH_METRIC` and `RECOVERY_CONTROL`.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The protocol to use when performing health checks. Valid values are `HTTP`, `HTTPS`, `HTTP_STR_MATCH`, `HTTPS_STR_MATCH`, `TCP`, `CALCULATED`, `CLOUDWATCH_METRIC` and `RECOVERY_CONTROL`.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private HealthCheckState() {}

    private HealthCheckState(HealthCheckState $) {
        this.arn = $.arn;
        this.childHealthThreshold = $.childHealthThreshold;
        this.childHealthchecks = $.childHealthchecks;
        this.cloudwatchAlarmName = $.cloudwatchAlarmName;
        this.cloudwatchAlarmRegion = $.cloudwatchAlarmRegion;
        this.disabled = $.disabled;
        this.enableSni = $.enableSni;
        this.failureThreshold = $.failureThreshold;
        this.fqdn = $.fqdn;
        this.insufficientDataHealthStatus = $.insufficientDataHealthStatus;
        this.invertHealthcheck = $.invertHealthcheck;
        this.ipAddress = $.ipAddress;
        this.measureLatency = $.measureLatency;
        this.port = $.port;
        this.referenceName = $.referenceName;
        this.regions = $.regions;
        this.requestInterval = $.requestInterval;
        this.resourcePath = $.resourcePath;
        this.routingControlArn = $.routingControlArn;
        this.searchString = $.searchString;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HealthCheckState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HealthCheckState $;

        public Builder() {
            $ = new HealthCheckState();
        }

        public Builder(HealthCheckState defaults) {
            $ = new HealthCheckState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The Amazon Resource Name (ARN) of the Health Check.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The Amazon Resource Name (ARN) of the Health Check.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param childHealthThreshold The minimum number of child health checks that must be healthy for Route 53 to consider the parent health check to be healthy. Valid values are integers between 0 and 256, inclusive
         * 
         * @return builder
         * 
         */
        public Builder childHealthThreshold(@Nullable Output<Integer> childHealthThreshold) {
            $.childHealthThreshold = childHealthThreshold;
            return this;
        }

        /**
         * @param childHealthThreshold The minimum number of child health checks that must be healthy for Route 53 to consider the parent health check to be healthy. Valid values are integers between 0 and 256, inclusive
         * 
         * @return builder
         * 
         */
        public Builder childHealthThreshold(Integer childHealthThreshold) {
            return childHealthThreshold(Output.of(childHealthThreshold));
        }

        /**
         * @param childHealthchecks For a specified parent health check, a list of HealthCheckId values for the associated child health checks.
         * 
         * @return builder
         * 
         */
        public Builder childHealthchecks(@Nullable Output<List<String>> childHealthchecks) {
            $.childHealthchecks = childHealthchecks;
            return this;
        }

        /**
         * @param childHealthchecks For a specified parent health check, a list of HealthCheckId values for the associated child health checks.
         * 
         * @return builder
         * 
         */
        public Builder childHealthchecks(List<String> childHealthchecks) {
            return childHealthchecks(Output.of(childHealthchecks));
        }

        /**
         * @param childHealthchecks For a specified parent health check, a list of HealthCheckId values for the associated child health checks.
         * 
         * @return builder
         * 
         */
        public Builder childHealthchecks(String... childHealthchecks) {
            return childHealthchecks(List.of(childHealthchecks));
        }

        /**
         * @param cloudwatchAlarmName The name of the CloudWatch alarm.
         * 
         * @return builder
         * 
         */
        public Builder cloudwatchAlarmName(@Nullable Output<String> cloudwatchAlarmName) {
            $.cloudwatchAlarmName = cloudwatchAlarmName;
            return this;
        }

        /**
         * @param cloudwatchAlarmName The name of the CloudWatch alarm.
         * 
         * @return builder
         * 
         */
        public Builder cloudwatchAlarmName(String cloudwatchAlarmName) {
            return cloudwatchAlarmName(Output.of(cloudwatchAlarmName));
        }

        /**
         * @param cloudwatchAlarmRegion The CloudWatchRegion that the CloudWatch alarm was created in.
         * 
         * @return builder
         * 
         */
        public Builder cloudwatchAlarmRegion(@Nullable Output<String> cloudwatchAlarmRegion) {
            $.cloudwatchAlarmRegion = cloudwatchAlarmRegion;
            return this;
        }

        /**
         * @param cloudwatchAlarmRegion The CloudWatchRegion that the CloudWatch alarm was created in.
         * 
         * @return builder
         * 
         */
        public Builder cloudwatchAlarmRegion(String cloudwatchAlarmRegion) {
            return cloudwatchAlarmRegion(Output.of(cloudwatchAlarmRegion));
        }

        /**
         * @param disabled A boolean value that stops Route 53 from performing health checks. When set to true, Route 53 will do the following depending on the type of health check:
         * * For health checks that check the health of endpoints, Route5 53 stops submitting requests to your application, server, or other resource.
         * * For calculated health checks, Route 53 stops aggregating the status of the referenced health checks.
         * * For health checks that monitor CloudWatch alarms, Route 53 stops monitoring the corresponding CloudWatch metrics.
         * 
         * &gt; **Note:** After you disable a health check, Route 53 considers the status of the health check to always be healthy. If you configured DNS failover, Route 53 continues to route traffic to the corresponding resources. If you want to stop routing traffic to a resource, change the value of `invert_healthcheck`.
         * 
         * @return builder
         * 
         */
        public Builder disabled(@Nullable Output<Boolean> disabled) {
            $.disabled = disabled;
            return this;
        }

        /**
         * @param disabled A boolean value that stops Route 53 from performing health checks. When set to true, Route 53 will do the following depending on the type of health check:
         * * For health checks that check the health of endpoints, Route5 53 stops submitting requests to your application, server, or other resource.
         * * For calculated health checks, Route 53 stops aggregating the status of the referenced health checks.
         * * For health checks that monitor CloudWatch alarms, Route 53 stops monitoring the corresponding CloudWatch metrics.
         * 
         * &gt; **Note:** After you disable a health check, Route 53 considers the status of the health check to always be healthy. If you configured DNS failover, Route 53 continues to route traffic to the corresponding resources. If you want to stop routing traffic to a resource, change the value of `invert_healthcheck`.
         * 
         * @return builder
         * 
         */
        public Builder disabled(Boolean disabled) {
            return disabled(Output.of(disabled));
        }

        /**
         * @param enableSni A boolean value that indicates whether Route53 should send the `fqdn` to the endpoint when performing the health check. This defaults to AWS&#39; defaults: when the `type` is &#34;HTTPS&#34; `enable_sni` defaults to `true`, when `type` is anything else `enable_sni` defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableSni(@Nullable Output<Boolean> enableSni) {
            $.enableSni = enableSni;
            return this;
        }

        /**
         * @param enableSni A boolean value that indicates whether Route53 should send the `fqdn` to the endpoint when performing the health check. This defaults to AWS&#39; defaults: when the `type` is &#34;HTTPS&#34; `enable_sni` defaults to `true`, when `type` is anything else `enable_sni` defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder enableSni(Boolean enableSni) {
            return enableSni(Output.of(enableSni));
        }

        /**
         * @param failureThreshold The number of consecutive health checks that an endpoint must pass or fail.
         * 
         * @return builder
         * 
         */
        public Builder failureThreshold(@Nullable Output<Integer> failureThreshold) {
            $.failureThreshold = failureThreshold;
            return this;
        }

        /**
         * @param failureThreshold The number of consecutive health checks that an endpoint must pass or fail.
         * 
         * @return builder
         * 
         */
        public Builder failureThreshold(Integer failureThreshold) {
            return failureThreshold(Output.of(failureThreshold));
        }

        /**
         * @param fqdn The fully qualified domain name of the endpoint to be checked. If a value is set for `ip_address`, the value set for `fqdn` will be passed in the `Host` header.
         * 
         * @return builder
         * 
         */
        public Builder fqdn(@Nullable Output<String> fqdn) {
            $.fqdn = fqdn;
            return this;
        }

        /**
         * @param fqdn The fully qualified domain name of the endpoint to be checked. If a value is set for `ip_address`, the value set for `fqdn` will be passed in the `Host` header.
         * 
         * @return builder
         * 
         */
        public Builder fqdn(String fqdn) {
            return fqdn(Output.of(fqdn));
        }

        /**
         * @param insufficientDataHealthStatus The status of the health check when CloudWatch has insufficient data about the state of associated alarm. Valid values are `Healthy` , `Unhealthy` and `LastKnownStatus`.
         * 
         * @return builder
         * 
         */
        public Builder insufficientDataHealthStatus(@Nullable Output<String> insufficientDataHealthStatus) {
            $.insufficientDataHealthStatus = insufficientDataHealthStatus;
            return this;
        }

        /**
         * @param insufficientDataHealthStatus The status of the health check when CloudWatch has insufficient data about the state of associated alarm. Valid values are `Healthy` , `Unhealthy` and `LastKnownStatus`.
         * 
         * @return builder
         * 
         */
        public Builder insufficientDataHealthStatus(String insufficientDataHealthStatus) {
            return insufficientDataHealthStatus(Output.of(insufficientDataHealthStatus));
        }

        /**
         * @param invertHealthcheck A boolean value that indicates whether the status of health check should be inverted. For example, if a health check is healthy but Inverted is True , then Route 53 considers the health check to be unhealthy.
         * 
         * @return builder
         * 
         */
        public Builder invertHealthcheck(@Nullable Output<Boolean> invertHealthcheck) {
            $.invertHealthcheck = invertHealthcheck;
            return this;
        }

        /**
         * @param invertHealthcheck A boolean value that indicates whether the status of health check should be inverted. For example, if a health check is healthy but Inverted is True , then Route 53 considers the health check to be unhealthy.
         * 
         * @return builder
         * 
         */
        public Builder invertHealthcheck(Boolean invertHealthcheck) {
            return invertHealthcheck(Output.of(invertHealthcheck));
        }

        /**
         * @param ipAddress The IP address of the endpoint to be checked.
         * 
         * @return builder
         * 
         */
        public Builder ipAddress(@Nullable Output<String> ipAddress) {
            $.ipAddress = ipAddress;
            return this;
        }

        /**
         * @param ipAddress The IP address of the endpoint to be checked.
         * 
         * @return builder
         * 
         */
        public Builder ipAddress(String ipAddress) {
            return ipAddress(Output.of(ipAddress));
        }

        /**
         * @param measureLatency A Boolean value that indicates whether you want Route 53 to measure the latency between health checkers in multiple AWS regions and your endpoint and to display CloudWatch latency graphs in the Route 53 console.
         * 
         * @return builder
         * 
         */
        public Builder measureLatency(@Nullable Output<Boolean> measureLatency) {
            $.measureLatency = measureLatency;
            return this;
        }

        /**
         * @param measureLatency A Boolean value that indicates whether you want Route 53 to measure the latency between health checkers in multiple AWS regions and your endpoint and to display CloudWatch latency graphs in the Route 53 console.
         * 
         * @return builder
         * 
         */
        public Builder measureLatency(Boolean measureLatency) {
            return measureLatency(Output.of(measureLatency));
        }

        /**
         * @param port The port of the endpoint to be checked.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The port of the endpoint to be checked.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param referenceName This is a reference name used in Caller Reference
         * (helpful for identifying single health_check set amongst others)
         * 
         * @return builder
         * 
         */
        public Builder referenceName(@Nullable Output<String> referenceName) {
            $.referenceName = referenceName;
            return this;
        }

        /**
         * @param referenceName This is a reference name used in Caller Reference
         * (helpful for identifying single health_check set amongst others)
         * 
         * @return builder
         * 
         */
        public Builder referenceName(String referenceName) {
            return referenceName(Output.of(referenceName));
        }

        /**
         * @param regions A list of AWS regions that you want Amazon Route 53 health checkers to check the specified endpoint from.
         * 
         * @return builder
         * 
         */
        public Builder regions(@Nullable Output<List<String>> regions) {
            $.regions = regions;
            return this;
        }

        /**
         * @param regions A list of AWS regions that you want Amazon Route 53 health checkers to check the specified endpoint from.
         * 
         * @return builder
         * 
         */
        public Builder regions(List<String> regions) {
            return regions(Output.of(regions));
        }

        /**
         * @param regions A list of AWS regions that you want Amazon Route 53 health checkers to check the specified endpoint from.
         * 
         * @return builder
         * 
         */
        public Builder regions(String... regions) {
            return regions(List.of(regions));
        }

        /**
         * @param requestInterval The number of seconds between the time that Amazon Route 53 gets a response from your endpoint and the time that it sends the next health-check request.
         * 
         * @return builder
         * 
         */
        public Builder requestInterval(@Nullable Output<Integer> requestInterval) {
            $.requestInterval = requestInterval;
            return this;
        }

        /**
         * @param requestInterval The number of seconds between the time that Amazon Route 53 gets a response from your endpoint and the time that it sends the next health-check request.
         * 
         * @return builder
         * 
         */
        public Builder requestInterval(Integer requestInterval) {
            return requestInterval(Output.of(requestInterval));
        }

        /**
         * @param resourcePath The path that you want Amazon Route 53 to request when performing health checks.
         * 
         * @return builder
         * 
         */
        public Builder resourcePath(@Nullable Output<String> resourcePath) {
            $.resourcePath = resourcePath;
            return this;
        }

        /**
         * @param resourcePath The path that you want Amazon Route 53 to request when performing health checks.
         * 
         * @return builder
         * 
         */
        public Builder resourcePath(String resourcePath) {
            return resourcePath(Output.of(resourcePath));
        }

        /**
         * @param routingControlArn The Amazon Resource Name (ARN) for the Route 53 Application Recovery Controller routing control. This is used when health check type is `RECOVERY_CONTROL`
         * 
         * @return builder
         * 
         */
        public Builder routingControlArn(@Nullable Output<String> routingControlArn) {
            $.routingControlArn = routingControlArn;
            return this;
        }

        /**
         * @param routingControlArn The Amazon Resource Name (ARN) for the Route 53 Application Recovery Controller routing control. This is used when health check type is `RECOVERY_CONTROL`
         * 
         * @return builder
         * 
         */
        public Builder routingControlArn(String routingControlArn) {
            return routingControlArn(Output.of(routingControlArn));
        }

        /**
         * @param searchString String searched in the first 5120 bytes of the response body for check to be considered healthy. Only valid with `HTTP_STR_MATCH` and `HTTPS_STR_MATCH`.
         * 
         * @return builder
         * 
         */
        public Builder searchString(@Nullable Output<String> searchString) {
            $.searchString = searchString;
            return this;
        }

        /**
         * @param searchString String searched in the first 5120 bytes of the response body for check to be considered healthy. Only valid with `HTTP_STR_MATCH` and `HTTPS_STR_MATCH`.
         * 
         * @return builder
         * 
         */
        public Builder searchString(String searchString) {
            return searchString(Output.of(searchString));
        }

        /**
         * @param tags A map of tags to assign to the health check. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the health check. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        /**
         * @param type The protocol to use when performing health checks. Valid values are `HTTP`, `HTTPS`, `HTTP_STR_MATCH`, `HTTPS_STR_MATCH`, `TCP`, `CALCULATED`, `CLOUDWATCH_METRIC` and `RECOVERY_CONTROL`.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The protocol to use when performing health checks. Valid values are `HTTP`, `HTTPS`, `HTTP_STR_MATCH`, `HTTPS_STR_MATCH`, `TCP`, `CALCULATED`, `CLOUDWATCH_METRIC` and `RECOVERY_CONTROL`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public HealthCheckState build() {
            return $;
        }
    }

}
