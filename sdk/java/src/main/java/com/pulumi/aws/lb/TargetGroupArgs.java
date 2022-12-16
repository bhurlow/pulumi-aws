// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.lb;

import com.pulumi.aws.lb.inputs.TargetGroupHealthCheckArgs;
import com.pulumi.aws.lb.inputs.TargetGroupStickinessArgs;
import com.pulumi.aws.lb.inputs.TargetGroupTargetFailoverArgs;
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


public final class TargetGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final TargetGroupArgs Empty = new TargetGroupArgs();

    /**
     * Whether to terminate connections at the end of the deregistration timeout on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#deregistration-delay) for more information. Default is `false`.
     * 
     */
    @Import(name="connectionTermination")
    private @Nullable Output<Boolean> connectionTermination;

    /**
     * @return Whether to terminate connections at the end of the deregistration timeout on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#deregistration-delay) for more information. Default is `false`.
     * 
     */
    public Optional<Output<Boolean>> connectionTermination() {
        return Optional.ofNullable(this.connectionTermination);
    }

    /**
     * Amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. The range is 0-3600 seconds. The default value is 300 seconds.
     * 
     */
    @Import(name="deregistrationDelay")
    private @Nullable Output<Integer> deregistrationDelay;

    /**
     * @return Amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. The range is 0-3600 seconds. The default value is 300 seconds.
     * 
     */
    public Optional<Output<Integer>> deregistrationDelay() {
        return Optional.ofNullable(this.deregistrationDelay);
    }

    /**
     * Health Check configuration block. Detailed below.
     * 
     */
    @Import(name="healthCheck")
    private @Nullable Output<TargetGroupHealthCheckArgs> healthCheck;

    /**
     * @return Health Check configuration block. Detailed below.
     * 
     */
    public Optional<Output<TargetGroupHealthCheckArgs>> healthCheck() {
        return Optional.ofNullable(this.healthCheck);
    }

    /**
     * The type of IP addresses used by the target group, only supported when target type is set to `ip`. Possible values are `ipv4` or `ipv6`.
     * 
     */
    @Import(name="ipAddressType")
    private @Nullable Output<String> ipAddressType;

    /**
     * @return The type of IP addresses used by the target group, only supported when target type is set to `ip`. Possible values are `ipv4` or `ipv6`.
     * 
     */
    public Optional<Output<String>> ipAddressType() {
        return Optional.ofNullable(this.ipAddressType);
    }

    /**
     * Whether the request and response headers exchanged between the load balancer and the Lambda function include arrays of values or strings. Only applies when `target_type` is `lambda`. Default is `false`.
     * 
     */
    @Import(name="lambdaMultiValueHeadersEnabled")
    private @Nullable Output<Boolean> lambdaMultiValueHeadersEnabled;

    /**
     * @return Whether the request and response headers exchanged between the load balancer and the Lambda function include arrays of values or strings. Only applies when `target_type` is `lambda`. Default is `false`.
     * 
     */
    public Optional<Output<Boolean>> lambdaMultiValueHeadersEnabled() {
        return Optional.ofNullable(this.lambdaMultiValueHeadersEnabled);
    }

    /**
     * Determines how the load balancer selects targets when routing requests. Only applicable for Application Load Balancer Target Groups. The value is `round_robin` or `least_outstanding_requests`. The default is `round_robin`.
     * 
     */
    @Import(name="loadBalancingAlgorithmType")
    private @Nullable Output<String> loadBalancingAlgorithmType;

    /**
     * @return Determines how the load balancer selects targets when routing requests. Only applicable for Application Load Balancer Target Groups. The value is `round_robin` or `least_outstanding_requests`. The default is `round_robin`.
     * 
     */
    public Optional<Output<String>> loadBalancingAlgorithmType() {
        return Optional.ofNullable(this.loadBalancingAlgorithmType);
    }

    /**
     * Name of the target group. If omitted, this provider will assign a random, unique name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the target group. If omitted, this provider will assign a random, unique name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`. Cannot be longer than 6 characters.
     * 
     */
    @Import(name="namePrefix")
    private @Nullable Output<String> namePrefix;

    /**
     * @return Creates a unique name beginning with the specified prefix. Conflicts with `name`. Cannot be longer than 6 characters.
     * 
     */
    public Optional<Output<String>> namePrefix() {
        return Optional.ofNullable(this.namePrefix);
    }

    /**
     * The port the load balancer uses when performing health checks on targets. Default is traffic-port.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return The port the load balancer uses when performing health checks on targets. Default is traffic-port.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * Whether client IP preservation is enabled. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#client-ip-preservation) for more information.
     * 
     */
    @Import(name="preserveClientIp")
    private @Nullable Output<String> preserveClientIp;

    /**
     * @return Whether client IP preservation is enabled. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#client-ip-preservation) for more information.
     * 
     */
    public Optional<Output<String>> preserveClientIp() {
        return Optional.ofNullable(this.preserveClientIp);
    }

    /**
     * Protocol the load balancer uses when performing health checks on targets. Must be either `TCP`, `HTTP`, or `HTTPS`. The TCP protocol is not supported for health checks if the protocol of the target group is HTTP or HTTPS. Defaults to HTTP.
     * 
     */
    @Import(name="protocol")
    private @Nullable Output<String> protocol;

    /**
     * @return Protocol the load balancer uses when performing health checks on targets. Must be either `TCP`, `HTTP`, or `HTTPS`. The TCP protocol is not supported for health checks if the protocol of the target group is HTTP or HTTPS. Defaults to HTTP.
     * 
     */
    public Optional<Output<String>> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    /**
     * Only applicable when `protocol` is `HTTP` or `HTTPS`. The protocol version. Specify `GRPC` to send requests to targets using gRPC. Specify `HTTP2` to send requests to targets using HTTP/2. The default is `HTTP1`, which sends requests to targets using HTTP/1.1
     * 
     */
    @Import(name="protocolVersion")
    private @Nullable Output<String> protocolVersion;

    /**
     * @return Only applicable when `protocol` is `HTTP` or `HTTPS`. The protocol version. Specify `GRPC` to send requests to targets using gRPC. Specify `HTTP2` to send requests to targets using HTTP/2. The default is `HTTP1`, which sends requests to targets using HTTP/1.1
     * 
     */
    public Optional<Output<String>> protocolVersion() {
        return Optional.ofNullable(this.protocolVersion);
    }

    /**
     * Whether to enable support for proxy protocol v2 on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#proxy-protocol) for more information. Default is `false`.
     * 
     */
    @Import(name="proxyProtocolV2")
    private @Nullable Output<Boolean> proxyProtocolV2;

    /**
     * @return Whether to enable support for proxy protocol v2 on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#proxy-protocol) for more information. Default is `false`.
     * 
     */
    public Optional<Output<Boolean>> proxyProtocolV2() {
        return Optional.ofNullable(this.proxyProtocolV2);
    }

    /**
     * Amount time for targets to warm up before the load balancer sends them a full share of requests. The range is 30-900 seconds or 0 to disable. The default value is 0 seconds.
     * 
     */
    @Import(name="slowStart")
    private @Nullable Output<Integer> slowStart;

    /**
     * @return Amount time for targets to warm up before the load balancer sends them a full share of requests. The range is 30-900 seconds or 0 to disable. The default value is 0 seconds.
     * 
     */
    public Optional<Output<Integer>> slowStart() {
        return Optional.ofNullable(this.slowStart);
    }

    /**
     * Stickiness configuration block. Detailed below.
     * 
     */
    @Import(name="stickiness")
    private @Nullable Output<TargetGroupStickinessArgs> stickiness;

    /**
     * @return Stickiness configuration block. Detailed below.
     * 
     */
    public Optional<Output<TargetGroupStickinessArgs>> stickiness() {
        return Optional.ofNullable(this.stickiness);
    }

    /**
     * Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Target failover block. Only applicable for Gateway Load Balancer target groups. See target_failover for more information.
     * 
     */
    @Import(name="targetFailovers")
    private @Nullable Output<List<TargetGroupTargetFailoverArgs>> targetFailovers;

    /**
     * @return Target failover block. Only applicable for Gateway Load Balancer target groups. See target_failover for more information.
     * 
     */
    public Optional<Output<List<TargetGroupTargetFailoverArgs>>> targetFailovers() {
        return Optional.ofNullable(this.targetFailovers);
    }

    /**
     * Type of target that you must specify when registering targets with this target group. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_CreateTargetGroup.html) for supported values. The default is `instance`.
     * 
     */
    @Import(name="targetType")
    private @Nullable Output<String> targetType;

    /**
     * @return Type of target that you must specify when registering targets with this target group. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_CreateTargetGroup.html) for supported values. The default is `instance`.
     * 
     */
    public Optional<Output<String>> targetType() {
        return Optional.ofNullable(this.targetType);
    }

    /**
     * Identifier of the VPC in which to create the target group. Required when `target_type` is `instance`, `ip` or `alb`. Does not apply when `target_type` is `lambda`.
     * 
     */
    @Import(name="vpcId")
    private @Nullable Output<String> vpcId;

    /**
     * @return Identifier of the VPC in which to create the target group. Required when `target_type` is `instance`, `ip` or `alb`. Does not apply when `target_type` is `lambda`.
     * 
     */
    public Optional<Output<String>> vpcId() {
        return Optional.ofNullable(this.vpcId);
    }

    private TargetGroupArgs() {}

    private TargetGroupArgs(TargetGroupArgs $) {
        this.connectionTermination = $.connectionTermination;
        this.deregistrationDelay = $.deregistrationDelay;
        this.healthCheck = $.healthCheck;
        this.ipAddressType = $.ipAddressType;
        this.lambdaMultiValueHeadersEnabled = $.lambdaMultiValueHeadersEnabled;
        this.loadBalancingAlgorithmType = $.loadBalancingAlgorithmType;
        this.name = $.name;
        this.namePrefix = $.namePrefix;
        this.port = $.port;
        this.preserveClientIp = $.preserveClientIp;
        this.protocol = $.protocol;
        this.protocolVersion = $.protocolVersion;
        this.proxyProtocolV2 = $.proxyProtocolV2;
        this.slowStart = $.slowStart;
        this.stickiness = $.stickiness;
        this.tags = $.tags;
        this.targetFailovers = $.targetFailovers;
        this.targetType = $.targetType;
        this.vpcId = $.vpcId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TargetGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TargetGroupArgs $;

        public Builder() {
            $ = new TargetGroupArgs();
        }

        public Builder(TargetGroupArgs defaults) {
            $ = new TargetGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param connectionTermination Whether to terminate connections at the end of the deregistration timeout on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#deregistration-delay) for more information. Default is `false`.
         * 
         * @return builder
         * 
         */
        public Builder connectionTermination(@Nullable Output<Boolean> connectionTermination) {
            $.connectionTermination = connectionTermination;
            return this;
        }

        /**
         * @param connectionTermination Whether to terminate connections at the end of the deregistration timeout on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#deregistration-delay) for more information. Default is `false`.
         * 
         * @return builder
         * 
         */
        public Builder connectionTermination(Boolean connectionTermination) {
            return connectionTermination(Output.of(connectionTermination));
        }

        /**
         * @param deregistrationDelay Amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. The range is 0-3600 seconds. The default value is 300 seconds.
         * 
         * @return builder
         * 
         */
        public Builder deregistrationDelay(@Nullable Output<Integer> deregistrationDelay) {
            $.deregistrationDelay = deregistrationDelay;
            return this;
        }

        /**
         * @param deregistrationDelay Amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. The range is 0-3600 seconds. The default value is 300 seconds.
         * 
         * @return builder
         * 
         */
        public Builder deregistrationDelay(Integer deregistrationDelay) {
            return deregistrationDelay(Output.of(deregistrationDelay));
        }

        /**
         * @param healthCheck Health Check configuration block. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder healthCheck(@Nullable Output<TargetGroupHealthCheckArgs> healthCheck) {
            $.healthCheck = healthCheck;
            return this;
        }

        /**
         * @param healthCheck Health Check configuration block. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder healthCheck(TargetGroupHealthCheckArgs healthCheck) {
            return healthCheck(Output.of(healthCheck));
        }

        /**
         * @param ipAddressType The type of IP addresses used by the target group, only supported when target type is set to `ip`. Possible values are `ipv4` or `ipv6`.
         * 
         * @return builder
         * 
         */
        public Builder ipAddressType(@Nullable Output<String> ipAddressType) {
            $.ipAddressType = ipAddressType;
            return this;
        }

        /**
         * @param ipAddressType The type of IP addresses used by the target group, only supported when target type is set to `ip`. Possible values are `ipv4` or `ipv6`.
         * 
         * @return builder
         * 
         */
        public Builder ipAddressType(String ipAddressType) {
            return ipAddressType(Output.of(ipAddressType));
        }

        /**
         * @param lambdaMultiValueHeadersEnabled Whether the request and response headers exchanged between the load balancer and the Lambda function include arrays of values or strings. Only applies when `target_type` is `lambda`. Default is `false`.
         * 
         * @return builder
         * 
         */
        public Builder lambdaMultiValueHeadersEnabled(@Nullable Output<Boolean> lambdaMultiValueHeadersEnabled) {
            $.lambdaMultiValueHeadersEnabled = lambdaMultiValueHeadersEnabled;
            return this;
        }

        /**
         * @param lambdaMultiValueHeadersEnabled Whether the request and response headers exchanged between the load balancer and the Lambda function include arrays of values or strings. Only applies when `target_type` is `lambda`. Default is `false`.
         * 
         * @return builder
         * 
         */
        public Builder lambdaMultiValueHeadersEnabled(Boolean lambdaMultiValueHeadersEnabled) {
            return lambdaMultiValueHeadersEnabled(Output.of(lambdaMultiValueHeadersEnabled));
        }

        /**
         * @param loadBalancingAlgorithmType Determines how the load balancer selects targets when routing requests. Only applicable for Application Load Balancer Target Groups. The value is `round_robin` or `least_outstanding_requests`. The default is `round_robin`.
         * 
         * @return builder
         * 
         */
        public Builder loadBalancingAlgorithmType(@Nullable Output<String> loadBalancingAlgorithmType) {
            $.loadBalancingAlgorithmType = loadBalancingAlgorithmType;
            return this;
        }

        /**
         * @param loadBalancingAlgorithmType Determines how the load balancer selects targets when routing requests. Only applicable for Application Load Balancer Target Groups. The value is `round_robin` or `least_outstanding_requests`. The default is `round_robin`.
         * 
         * @return builder
         * 
         */
        public Builder loadBalancingAlgorithmType(String loadBalancingAlgorithmType) {
            return loadBalancingAlgorithmType(Output.of(loadBalancingAlgorithmType));
        }

        /**
         * @param name Name of the target group. If omitted, this provider will assign a random, unique name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the target group. If omitted, this provider will assign a random, unique name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param namePrefix Creates a unique name beginning with the specified prefix. Conflicts with `name`. Cannot be longer than 6 characters.
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(@Nullable Output<String> namePrefix) {
            $.namePrefix = namePrefix;
            return this;
        }

        /**
         * @param namePrefix Creates a unique name beginning with the specified prefix. Conflicts with `name`. Cannot be longer than 6 characters.
         * 
         * @return builder
         * 
         */
        public Builder namePrefix(String namePrefix) {
            return namePrefix(Output.of(namePrefix));
        }

        /**
         * @param port The port the load balancer uses when performing health checks on targets. Default is traffic-port.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The port the load balancer uses when performing health checks on targets. Default is traffic-port.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param preserveClientIp Whether client IP preservation is enabled. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#client-ip-preservation) for more information.
         * 
         * @return builder
         * 
         */
        public Builder preserveClientIp(@Nullable Output<String> preserveClientIp) {
            $.preserveClientIp = preserveClientIp;
            return this;
        }

        /**
         * @param preserveClientIp Whether client IP preservation is enabled. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#client-ip-preservation) for more information.
         * 
         * @return builder
         * 
         */
        public Builder preserveClientIp(String preserveClientIp) {
            return preserveClientIp(Output.of(preserveClientIp));
        }

        /**
         * @param protocol Protocol the load balancer uses when performing health checks on targets. Must be either `TCP`, `HTTP`, or `HTTPS`. The TCP protocol is not supported for health checks if the protocol of the target group is HTTP or HTTPS. Defaults to HTTP.
         * 
         * @return builder
         * 
         */
        public Builder protocol(@Nullable Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        /**
         * @param protocol Protocol the load balancer uses when performing health checks on targets. Must be either `TCP`, `HTTP`, or `HTTPS`. The TCP protocol is not supported for health checks if the protocol of the target group is HTTP or HTTPS. Defaults to HTTP.
         * 
         * @return builder
         * 
         */
        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        /**
         * @param protocolVersion Only applicable when `protocol` is `HTTP` or `HTTPS`. The protocol version. Specify `GRPC` to send requests to targets using gRPC. Specify `HTTP2` to send requests to targets using HTTP/2. The default is `HTTP1`, which sends requests to targets using HTTP/1.1
         * 
         * @return builder
         * 
         */
        public Builder protocolVersion(@Nullable Output<String> protocolVersion) {
            $.protocolVersion = protocolVersion;
            return this;
        }

        /**
         * @param protocolVersion Only applicable when `protocol` is `HTTP` or `HTTPS`. The protocol version. Specify `GRPC` to send requests to targets using gRPC. Specify `HTTP2` to send requests to targets using HTTP/2. The default is `HTTP1`, which sends requests to targets using HTTP/1.1
         * 
         * @return builder
         * 
         */
        public Builder protocolVersion(String protocolVersion) {
            return protocolVersion(Output.of(protocolVersion));
        }

        /**
         * @param proxyProtocolV2 Whether to enable support for proxy protocol v2 on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#proxy-protocol) for more information. Default is `false`.
         * 
         * @return builder
         * 
         */
        public Builder proxyProtocolV2(@Nullable Output<Boolean> proxyProtocolV2) {
            $.proxyProtocolV2 = proxyProtocolV2;
            return this;
        }

        /**
         * @param proxyProtocolV2 Whether to enable support for proxy protocol v2 on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#proxy-protocol) for more information. Default is `false`.
         * 
         * @return builder
         * 
         */
        public Builder proxyProtocolV2(Boolean proxyProtocolV2) {
            return proxyProtocolV2(Output.of(proxyProtocolV2));
        }

        /**
         * @param slowStart Amount time for targets to warm up before the load balancer sends them a full share of requests. The range is 30-900 seconds or 0 to disable. The default value is 0 seconds.
         * 
         * @return builder
         * 
         */
        public Builder slowStart(@Nullable Output<Integer> slowStart) {
            $.slowStart = slowStart;
            return this;
        }

        /**
         * @param slowStart Amount time for targets to warm up before the load balancer sends them a full share of requests. The range is 30-900 seconds or 0 to disable. The default value is 0 seconds.
         * 
         * @return builder
         * 
         */
        public Builder slowStart(Integer slowStart) {
            return slowStart(Output.of(slowStart));
        }

        /**
         * @param stickiness Stickiness configuration block. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder stickiness(@Nullable Output<TargetGroupStickinessArgs> stickiness) {
            $.stickiness = stickiness;
            return this;
        }

        /**
         * @param stickiness Stickiness configuration block. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder stickiness(TargetGroupStickinessArgs stickiness) {
            return stickiness(Output.of(stickiness));
        }

        /**
         * @param tags Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param targetFailovers Target failover block. Only applicable for Gateway Load Balancer target groups. See target_failover for more information.
         * 
         * @return builder
         * 
         */
        public Builder targetFailovers(@Nullable Output<List<TargetGroupTargetFailoverArgs>> targetFailovers) {
            $.targetFailovers = targetFailovers;
            return this;
        }

        /**
         * @param targetFailovers Target failover block. Only applicable for Gateway Load Balancer target groups. See target_failover for more information.
         * 
         * @return builder
         * 
         */
        public Builder targetFailovers(List<TargetGroupTargetFailoverArgs> targetFailovers) {
            return targetFailovers(Output.of(targetFailovers));
        }

        /**
         * @param targetFailovers Target failover block. Only applicable for Gateway Load Balancer target groups. See target_failover for more information.
         * 
         * @return builder
         * 
         */
        public Builder targetFailovers(TargetGroupTargetFailoverArgs... targetFailovers) {
            return targetFailovers(List.of(targetFailovers));
        }

        /**
         * @param targetType Type of target that you must specify when registering targets with this target group. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_CreateTargetGroup.html) for supported values. The default is `instance`.
         * 
         * @return builder
         * 
         */
        public Builder targetType(@Nullable Output<String> targetType) {
            $.targetType = targetType;
            return this;
        }

        /**
         * @param targetType Type of target that you must specify when registering targets with this target group. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_CreateTargetGroup.html) for supported values. The default is `instance`.
         * 
         * @return builder
         * 
         */
        public Builder targetType(String targetType) {
            return targetType(Output.of(targetType));
        }

        /**
         * @param vpcId Identifier of the VPC in which to create the target group. Required when `target_type` is `instance`, `ip` or `alb`. Does not apply when `target_type` is `lambda`.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(@Nullable Output<String> vpcId) {
            $.vpcId = vpcId;
            return this;
        }

        /**
         * @param vpcId Identifier of the VPC in which to create the target group. Required when `target_type` is `instance`, `ip` or `alb`. Does not apply when `target_type` is `lambda`.
         * 
         * @return builder
         * 
         */
        public Builder vpcId(String vpcId) {
            return vpcId(Output.of(vpcId));
        }

        public TargetGroupArgs build() {
            return $;
        }
    }

}
