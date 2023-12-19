// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.alb;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.alb.LoadBalancerArgs;
import com.pulumi.aws.alb.inputs.LoadBalancerState;
import com.pulumi.aws.alb.outputs.LoadBalancerAccessLogs;
import com.pulumi.aws.alb.outputs.LoadBalancerConnectionLogs;
import com.pulumi.aws.alb.outputs.LoadBalancerSubnetMapping;
import com.pulumi.core.Alias;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Load Balancer resource.
 * 
 * &gt; **Note:** `aws.alb.LoadBalancer` is known as `aws.lb.LoadBalancer`. The functionality is identical.
 * 
 * ## Example Usage
 * ### Specifying Elastic IPs
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.lb.LoadBalancer;
 * import com.pulumi.aws.lb.LoadBalancerArgs;
 * import com.pulumi.aws.lb.inputs.LoadBalancerSubnetMappingArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new LoadBalancer(&#34;example&#34;, LoadBalancerArgs.builder()        
 *             .loadBalancerType(&#34;network&#34;)
 *             .subnetMappings(            
 *                 LoadBalancerSubnetMappingArgs.builder()
 *                     .subnetId(aws_subnet.example1().id())
 *                     .allocationId(aws_eip.example1().id())
 *                     .build(),
 *                 LoadBalancerSubnetMappingArgs.builder()
 *                     .subnetId(aws_subnet.example2().id())
 *                     .allocationId(aws_eip.example2().id())
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Specifying private IP addresses for an internal-facing load balancer
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.lb.LoadBalancer;
 * import com.pulumi.aws.lb.LoadBalancerArgs;
 * import com.pulumi.aws.lb.inputs.LoadBalancerSubnetMappingArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new LoadBalancer(&#34;example&#34;, LoadBalancerArgs.builder()        
 *             .loadBalancerType(&#34;network&#34;)
 *             .subnetMappings(            
 *                 LoadBalancerSubnetMappingArgs.builder()
 *                     .subnetId(aws_subnet.example1().id())
 *                     .privateIpv4Address(&#34;10.0.1.15&#34;)
 *                     .build(),
 *                 LoadBalancerSubnetMappingArgs.builder()
 *                     .subnetId(aws_subnet.example2().id())
 *                     .privateIpv4Address(&#34;10.0.2.15&#34;)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import LBs using their ARN. For example:
 * 
 * ```sh
 *  $ pulumi import aws:alb/loadBalancer:LoadBalancer bar arn:aws:elasticloadbalancing:us-west-2:123456789012:loadbalancer/app/my-load-balancer/50dc6c495c0c9188
 * ```
 * 
 */
@ResourceType(type="aws:alb/loadBalancer:LoadBalancer")
public class LoadBalancer extends com.pulumi.resources.CustomResource {
    /**
     * An Access Logs block. Access Logs documented below.
     * 
     */
    @Export(name="accessLogs", refs={LoadBalancerAccessLogs.class}, tree="[0]")
    private Output</* @Nullable */ LoadBalancerAccessLogs> accessLogs;

    /**
     * @return An Access Logs block. Access Logs documented below.
     * 
     */
    public Output<Optional<LoadBalancerAccessLogs>> accessLogs() {
        return Codegen.optional(this.accessLogs);
    }
    /**
     * The ARN of the load balancer (matches `id`).
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the load balancer (matches `id`).
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The ARN suffix for use with CloudWatch Metrics.
     * 
     */
    @Export(name="arnSuffix", refs={String.class}, tree="[0]")
    private Output<String> arnSuffix;

    /**
     * @return The ARN suffix for use with CloudWatch Metrics.
     * 
     */
    public Output<String> arnSuffix() {
        return this.arnSuffix;
    }
    /**
     * A Connection Logs block. Connection Logs documented below. Only valid for Load Balancers of type `application`.
     * 
     */
    @Export(name="connectionLogs", refs={LoadBalancerConnectionLogs.class}, tree="[0]")
    private Output</* @Nullable */ LoadBalancerConnectionLogs> connectionLogs;

    /**
     * @return A Connection Logs block. Connection Logs documented below. Only valid for Load Balancers of type `application`.
     * 
     */
    public Output<Optional<LoadBalancerConnectionLogs>> connectionLogs() {
        return Codegen.optional(this.connectionLogs);
    }
    /**
     * The ID of the customer owned ipv4 pool to use for this load balancer.
     * 
     */
    @Export(name="customerOwnedIpv4Pool", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> customerOwnedIpv4Pool;

    /**
     * @return The ID of the customer owned ipv4 pool to use for this load balancer.
     * 
     */
    public Output<Optional<String>> customerOwnedIpv4Pool() {
        return Codegen.optional(this.customerOwnedIpv4Pool);
    }
    /**
     * Determines how the load balancer handles requests that might pose a security risk to an application due to HTTP desync. Valid values are `monitor`, `defensive` (default), `strictest`.
     * 
     */
    @Export(name="desyncMitigationMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> desyncMitigationMode;

    /**
     * @return Determines how the load balancer handles requests that might pose a security risk to an application due to HTTP desync. Valid values are `monitor`, `defensive` (default), `strictest`.
     * 
     */
    public Output<Optional<String>> desyncMitigationMode() {
        return Codegen.optional(this.desyncMitigationMode);
    }
    /**
     * The DNS name of the load balancer.
     * 
     */
    @Export(name="dnsName", refs={String.class}, tree="[0]")
    private Output<String> dnsName;

    /**
     * @return The DNS name of the load balancer.
     * 
     */
    public Output<String> dnsName() {
        return this.dnsName;
    }
    /**
     * Indicates how traffic is distributed among the load balancer Availability Zones. Possible values are `any_availability_zone` (default), `availability_zone_affinity`, or `partial_availability_zone_affinity`. See   [Availability Zone DNS affinity](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/network-load-balancers.html#zonal-dns-affinity) for additional details. Only valid for `network` type load balancers.
     * 
     */
    @Export(name="dnsRecordClientRoutingPolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dnsRecordClientRoutingPolicy;

    /**
     * @return Indicates how traffic is distributed among the load balancer Availability Zones. Possible values are `any_availability_zone` (default), `availability_zone_affinity`, or `partial_availability_zone_affinity`. See   [Availability Zone DNS affinity](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/network-load-balancers.html#zonal-dns-affinity) for additional details. Only valid for `network` type load balancers.
     * 
     */
    public Output<Optional<String>> dnsRecordClientRoutingPolicy() {
        return Codegen.optional(this.dnsRecordClientRoutingPolicy);
    }
    /**
     * Indicates whether HTTP headers with header fields that are not valid are removed by the load balancer (true) or routed to targets (false). The default is false. Elastic Load Balancing requires that message header names contain only alphanumeric characters and hyphens. Only valid for Load Balancers of type `application`.
     * 
     */
    @Export(name="dropInvalidHeaderFields", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dropInvalidHeaderFields;

    /**
     * @return Indicates whether HTTP headers with header fields that are not valid are removed by the load balancer (true) or routed to targets (false). The default is false. Elastic Load Balancing requires that message header names contain only alphanumeric characters and hyphens. Only valid for Load Balancers of type `application`.
     * 
     */
    public Output<Optional<Boolean>> dropInvalidHeaderFields() {
        return Codegen.optional(this.dropInvalidHeaderFields);
    }
    /**
     * If true, cross-zone load balancing of the load balancer will be enabled. For `network` and `gateway` type load balancers, this feature is disabled by default (`false`). For `application` load balancer this feature is always enabled (`true`) and cannot be disabled. Defaults to `false`.
     * 
     */
    @Export(name="enableCrossZoneLoadBalancing", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableCrossZoneLoadBalancing;

    /**
     * @return If true, cross-zone load balancing of the load balancer will be enabled. For `network` and `gateway` type load balancers, this feature is disabled by default (`false`). For `application` load balancer this feature is always enabled (`true`) and cannot be disabled. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> enableCrossZoneLoadBalancing() {
        return Codegen.optional(this.enableCrossZoneLoadBalancing);
    }
    /**
     * If true, deletion of the load balancer will be disabled via the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
     * 
     */
    @Export(name="enableDeletionProtection", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableDeletionProtection;

    /**
     * @return If true, deletion of the load balancer will be disabled via the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> enableDeletionProtection() {
        return Codegen.optional(this.enableDeletionProtection);
    }
    /**
     * Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
     * 
     */
    @Export(name="enableHttp2", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableHttp2;

    /**
     * @return Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> enableHttp2() {
        return Codegen.optional(this.enableHttp2);
    }
    /**
     * Indicates whether the two headers (`x-amzn-tls-version` and `x-amzn-tls-cipher-suite`), which contain information about the negotiated TLS version and cipher suite, are added to the client request before sending it to the target. Only valid for Load Balancers of type `application`. Defaults to `false`
     * 
     */
    @Export(name="enableTlsVersionAndCipherSuiteHeaders", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableTlsVersionAndCipherSuiteHeaders;

    /**
     * @return Indicates whether the two headers (`x-amzn-tls-version` and `x-amzn-tls-cipher-suite`), which contain information about the negotiated TLS version and cipher suite, are added to the client request before sending it to the target. Only valid for Load Balancers of type `application`. Defaults to `false`
     * 
     */
    public Output<Optional<Boolean>> enableTlsVersionAndCipherSuiteHeaders() {
        return Codegen.optional(this.enableTlsVersionAndCipherSuiteHeaders);
    }
    /**
     * Indicates whether to allow a WAF-enabled load balancer to route requests to targets if it is unable to forward the request to AWS WAF. Defaults to `false`.
     * 
     */
    @Export(name="enableWafFailOpen", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableWafFailOpen;

    /**
     * @return Indicates whether to allow a WAF-enabled load balancer to route requests to targets if it is unable to forward the request to AWS WAF. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> enableWafFailOpen() {
        return Codegen.optional(this.enableWafFailOpen);
    }
    /**
     * Indicates whether the X-Forwarded-For header should preserve the source port that the client used to connect to the load balancer in `application` load balancers. Defaults to `false`.
     * 
     */
    @Export(name="enableXffClientPort", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableXffClientPort;

    /**
     * @return Indicates whether the X-Forwarded-For header should preserve the source port that the client used to connect to the load balancer in `application` load balancers. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> enableXffClientPort() {
        return Codegen.optional(this.enableXffClientPort);
    }
    /**
     * Indicates whether inbound security group rules are enforced for traffic originating from a PrivateLink. Only valid for Load Balancers of type `network`. The possible values are `on` and `off`.
     * 
     */
    @Export(name="enforceSecurityGroupInboundRulesOnPrivateLinkTraffic", refs={String.class}, tree="[0]")
    private Output<String> enforceSecurityGroupInboundRulesOnPrivateLinkTraffic;

    /**
     * @return Indicates whether inbound security group rules are enforced for traffic originating from a PrivateLink. Only valid for Load Balancers of type `network`. The possible values are `on` and `off`.
     * 
     */
    public Output<String> enforceSecurityGroupInboundRulesOnPrivateLinkTraffic() {
        return this.enforceSecurityGroupInboundRulesOnPrivateLinkTraffic;
    }
    /**
     * The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
     * 
     */
    @Export(name="idleTimeout", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> idleTimeout;

    /**
     * @return The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
     * 
     */
    public Output<Optional<Integer>> idleTimeout() {
        return Codegen.optional(this.idleTimeout);
    }
    /**
     * If true, the LB will be internal. Defaults to `false`.
     * 
     */
    @Export(name="internal", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> internal;

    /**
     * @return If true, the LB will be internal. Defaults to `false`.
     * 
     */
    public Output<Boolean> internal() {
        return this.internal;
    }
    /**
     * The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`.
     * 
     */
    @Export(name="ipAddressType", refs={String.class}, tree="[0]")
    private Output<String> ipAddressType;

    /**
     * @return The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`.
     * 
     */
    public Output<String> ipAddressType() {
        return this.ipAddressType;
    }
    /**
     * The type of load balancer to create. Possible values are `application`, `gateway`, or `network`. The default value is `application`.
     * 
     */
    @Export(name="loadBalancerType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> loadBalancerType;

    /**
     * @return The type of load balancer to create. Possible values are `application`, `gateway`, or `network`. The default value is `application`.
     * 
     */
    public Output<Optional<String>> loadBalancerType() {
        return Codegen.optional(this.loadBalancerType);
    }
    /**
     * The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
     * must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
     * this provider will autogenerate a name beginning with `tf-lb`.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
     * must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
     * this provider will autogenerate a name beginning with `tf-lb`.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     * 
     */
    @Export(name="namePrefix", refs={String.class}, tree="[0]")
    private Output<String> namePrefix;

    /**
     * @return Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     * 
     */
    public Output<String> namePrefix() {
        return this.namePrefix;
    }
    /**
     * Indicates whether the Application Load Balancer should preserve the Host header in the HTTP request and send it to the target without any change. Defaults to `false`.
     * 
     */
    @Export(name="preserveHostHeader", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> preserveHostHeader;

    /**
     * @return Indicates whether the Application Load Balancer should preserve the Host header in the HTTP request and send it to the target without any change. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> preserveHostHeader() {
        return Codegen.optional(this.preserveHostHeader);
    }
    /**
     * A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application` or `network`. For load balancers of type `network` security groups cannot be added if none are currently present, and cannot all be removed once added. If either of these conditions are met, this will force a recreation of the resource.
     * 
     */
    @Export(name="securityGroups", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> securityGroups;

    /**
     * @return A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application` or `network`. For load balancers of type `network` security groups cannot be added if none are currently present, and cannot all be removed once added. If either of these conditions are met, this will force a recreation of the resource.
     * 
     */
    public Output<List<String>> securityGroups() {
        return this.securityGroups;
    }
    /**
     * A subnet mapping block as documented below. For Load Balancers of type `network` subnet mappings can only be added.
     * 
     */
    @Export(name="subnetMappings", refs={List.class,LoadBalancerSubnetMapping.class}, tree="[0,1]")
    private Output<List<LoadBalancerSubnetMapping>> subnetMappings;

    /**
     * @return A subnet mapping block as documented below. For Load Balancers of type `network` subnet mappings can only be added.
     * 
     */
    public Output<List<LoadBalancerSubnetMapping>> subnetMappings() {
        return this.subnetMappings;
    }
    /**
     * A list of subnet IDs to attach to the LB. For Load Balancers of type `network` subnets can only be added (see [Availability Zones](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/network-load-balancers.html#availability-zones)), deleting a subnet for load balancers of type `network` will force a recreation of the resource.
     * 
     */
    @Export(name="subnets", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> subnets;

    /**
     * @return A list of subnet IDs to attach to the LB. For Load Balancers of type `network` subnets can only be added (see [Availability Zones](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/network-load-balancers.html#availability-zones)), deleting a subnet for load balancers of type `network` will force a recreation of the resource.
     * 
     */
    public Output<List<String>> subnets() {
        return this.subnets;
    }
    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Export(name="tagsAll", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output<String> vpcId;

    public Output<String> vpcId() {
        return this.vpcId;
    }
    /**
     * Determines how the load balancer modifies the `X-Forwarded-For` header in the HTTP request before sending the request to the target. The possible values are `append`, `preserve`, and `remove`. Only valid for Load Balancers of type `application`. The default is `append`.
     * 
     */
    @Export(name="xffHeaderProcessingMode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> xffHeaderProcessingMode;

    /**
     * @return Determines how the load balancer modifies the `X-Forwarded-For` header in the HTTP request before sending the request to the target. The possible values are `append`, `preserve`, and `remove`. Only valid for Load Balancers of type `application`. The default is `append`.
     * 
     */
    public Output<Optional<String>> xffHeaderProcessingMode() {
        return Codegen.optional(this.xffHeaderProcessingMode);
    }
    /**
     * The canonical hosted zone ID of the load balancer (to be used in a Route 53 Alias record).
     * 
     */
    @Export(name="zoneId", refs={String.class}, tree="[0]")
    private Output<String> zoneId;

    /**
     * @return The canonical hosted zone ID of the load balancer (to be used in a Route 53 Alias record).
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LoadBalancer(String name) {
        this(name, LoadBalancerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LoadBalancer(String name, @Nullable LoadBalancerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LoadBalancer(String name, @Nullable LoadBalancerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:alb/loadBalancer:LoadBalancer", name, args == null ? LoadBalancerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LoadBalancer(String name, Output<String> id, @Nullable LoadBalancerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:alb/loadBalancer:LoadBalancer", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .aliases(List.of(
                Output.of(Alias.builder().type("aws:applicationloadbalancing/loadBalancer:LoadBalancer").build())
            ))
            .additionalSecretOutputs(List.of(
                "tagsAll"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static LoadBalancer get(String name, Output<String> id, @Nullable LoadBalancerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LoadBalancer(name, id, state, options);
    }
}
