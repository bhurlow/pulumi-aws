// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.lb.inputs;

import com.pulumi.aws.lb.inputs.ListenerDefaultActionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ListenerState extends com.pulumi.resources.ResourceArgs {

    public static final ListenerState Empty = new ListenerState();

    /**
     * Name of the Application-Layer Protocol Negotiation (ALPN) policy. Can be set if `protocol` is `TLS`. Valid values are `HTTP1Only`, `HTTP2Only`, `HTTP2Optional`, `HTTP2Preferred`, and `None`.
     * 
     */
    @Import(name="alpnPolicy")
    private @Nullable Output<String> alpnPolicy;

    /**
     * @return Name of the Application-Layer Protocol Negotiation (ALPN) policy. Can be set if `protocol` is `TLS`. Valid values are `HTTP1Only`, `HTTP2Only`, `HTTP2Optional`, `HTTP2Preferred`, and `None`.
     * 
     */
    public Optional<Output<String>> alpnPolicy() {
        return Optional.ofNullable(this.alpnPolicy);
    }

    /**
     * ARN of the target group.
     * 
     * The following arguments are optional:
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return ARN of the target group.
     * 
     * The following arguments are optional:
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `aws.lb.ListenerCertificate` resource.
     * 
     */
    @Import(name="certificateArn")
    private @Nullable Output<String> certificateArn;

    /**
     * @return ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `aws.lb.ListenerCertificate` resource.
     * 
     */
    public Optional<Output<String>> certificateArn() {
        return Optional.ofNullable(this.certificateArn);
    }

    /**
     * Configuration block for default actions. Detailed below.
     * 
     */
    @Import(name="defaultActions")
    private @Nullable Output<List<ListenerDefaultActionArgs>> defaultActions;

    /**
     * @return Configuration block for default actions. Detailed below.
     * 
     */
    public Optional<Output<List<ListenerDefaultActionArgs>>> defaultActions() {
        return Optional.ofNullable(this.defaultActions);
    }

    /**
     * ARN of the load balancer.
     * 
     * The following arguments are optional:
     * 
     */
    @Import(name="loadBalancerArn")
    private @Nullable Output<String> loadBalancerArn;

    /**
     * @return ARN of the load balancer.
     * 
     * The following arguments are optional:
     * 
     */
    public Optional<Output<String>> loadBalancerArn() {
        return Optional.ofNullable(this.loadBalancerArn);
    }

    /**
     * Port on which the load balancer is listening. Not valid for Gateway Load Balancers.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return Port on which the load balancer is listening. Not valid for Gateway Load Balancers.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * Protocol for connections from clients to the load balancer. For Application Load Balancers, valid values are `HTTP` and `HTTPS`, with a default of `HTTP`. For Network Load Balancers, valid values are `TCP`, `TLS`, `UDP`, and `TCP_UDP`. Not valid to use `UDP` or `TCP_UDP` if dual-stack mode is enabled. Not valid for Gateway Load Balancers.
     * 
     */
    @Import(name="protocol")
    private @Nullable Output<String> protocol;

    /**
     * @return Protocol for connections from clients to the load balancer. For Application Load Balancers, valid values are `HTTP` and `HTTPS`, with a default of `HTTP`. For Network Load Balancers, valid values are `TCP`, `TLS`, `UDP`, and `TCP_UDP`. Not valid to use `UDP` or `TCP_UDP` if dual-stack mode is enabled. Not valid for Gateway Load Balancers.
     * 
     */
    public Optional<Output<String>> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    /**
     * Name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
     * 
     */
    @Import(name="sslPolicy")
    private @Nullable Output<String> sslPolicy;

    /**
     * @return Name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
     * 
     */
    public Optional<Output<String>> sslPolicy() {
        return Optional.ofNullable(this.sslPolicy);
    }

    /**
     * A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     * &gt; **NOTE::** Please note that listeners that are attached to Application Load Balancers must use either `HTTP` or `HTTPS` protocols while listeners that are attached to Network Load Balancers must use the `TCP` protocol.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     * &gt; **NOTE::** Please note that listeners that are attached to Application Load Balancers must use either `HTTP` or `HTTPS` protocols while listeners that are attached to Network Load Balancers must use the `TCP` protocol.
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

    private ListenerState() {}

    private ListenerState(ListenerState $) {
        this.alpnPolicy = $.alpnPolicy;
        this.arn = $.arn;
        this.certificateArn = $.certificateArn;
        this.defaultActions = $.defaultActions;
        this.loadBalancerArn = $.loadBalancerArn;
        this.port = $.port;
        this.protocol = $.protocol;
        this.sslPolicy = $.sslPolicy;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ListenerState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ListenerState $;

        public Builder() {
            $ = new ListenerState();
        }

        public Builder(ListenerState defaults) {
            $ = new ListenerState(Objects.requireNonNull(defaults));
        }

        /**
         * @param alpnPolicy Name of the Application-Layer Protocol Negotiation (ALPN) policy. Can be set if `protocol` is `TLS`. Valid values are `HTTP1Only`, `HTTP2Only`, `HTTP2Optional`, `HTTP2Preferred`, and `None`.
         * 
         * @return builder
         * 
         */
        public Builder alpnPolicy(@Nullable Output<String> alpnPolicy) {
            $.alpnPolicy = alpnPolicy;
            return this;
        }

        /**
         * @param alpnPolicy Name of the Application-Layer Protocol Negotiation (ALPN) policy. Can be set if `protocol` is `TLS`. Valid values are `HTTP1Only`, `HTTP2Only`, `HTTP2Optional`, `HTTP2Preferred`, and `None`.
         * 
         * @return builder
         * 
         */
        public Builder alpnPolicy(String alpnPolicy) {
            return alpnPolicy(Output.of(alpnPolicy));
        }

        /**
         * @param arn ARN of the target group.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn ARN of the target group.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param certificateArn ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `aws.lb.ListenerCertificate` resource.
         * 
         * @return builder
         * 
         */
        public Builder certificateArn(@Nullable Output<String> certificateArn) {
            $.certificateArn = certificateArn;
            return this;
        }

        /**
         * @param certificateArn ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `aws.lb.ListenerCertificate` resource.
         * 
         * @return builder
         * 
         */
        public Builder certificateArn(String certificateArn) {
            return certificateArn(Output.of(certificateArn));
        }

        /**
         * @param defaultActions Configuration block for default actions. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder defaultActions(@Nullable Output<List<ListenerDefaultActionArgs>> defaultActions) {
            $.defaultActions = defaultActions;
            return this;
        }

        /**
         * @param defaultActions Configuration block for default actions. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder defaultActions(List<ListenerDefaultActionArgs> defaultActions) {
            return defaultActions(Output.of(defaultActions));
        }

        /**
         * @param defaultActions Configuration block for default actions. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder defaultActions(ListenerDefaultActionArgs... defaultActions) {
            return defaultActions(List.of(defaultActions));
        }

        /**
         * @param loadBalancerArn ARN of the load balancer.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerArn(@Nullable Output<String> loadBalancerArn) {
            $.loadBalancerArn = loadBalancerArn;
            return this;
        }

        /**
         * @param loadBalancerArn ARN of the load balancer.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder loadBalancerArn(String loadBalancerArn) {
            return loadBalancerArn(Output.of(loadBalancerArn));
        }

        /**
         * @param port Port on which the load balancer is listening. Not valid for Gateway Load Balancers.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port Port on which the load balancer is listening. Not valid for Gateway Load Balancers.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param protocol Protocol for connections from clients to the load balancer. For Application Load Balancers, valid values are `HTTP` and `HTTPS`, with a default of `HTTP`. For Network Load Balancers, valid values are `TCP`, `TLS`, `UDP`, and `TCP_UDP`. Not valid to use `UDP` or `TCP_UDP` if dual-stack mode is enabled. Not valid for Gateway Load Balancers.
         * 
         * @return builder
         * 
         */
        public Builder protocol(@Nullable Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        /**
         * @param protocol Protocol for connections from clients to the load balancer. For Application Load Balancers, valid values are `HTTP` and `HTTPS`, with a default of `HTTP`. For Network Load Balancers, valid values are `TCP`, `TLS`, `UDP`, and `TCP_UDP`. Not valid to use `UDP` or `TCP_UDP` if dual-stack mode is enabled. Not valid for Gateway Load Balancers.
         * 
         * @return builder
         * 
         */
        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        /**
         * @param sslPolicy Name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
         * 
         * @return builder
         * 
         */
        public Builder sslPolicy(@Nullable Output<String> sslPolicy) {
            $.sslPolicy = sslPolicy;
            return this;
        }

        /**
         * @param sslPolicy Name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
         * 
         * @return builder
         * 
         */
        public Builder sslPolicy(String sslPolicy) {
            return sslPolicy(Output.of(sslPolicy));
        }

        /**
         * @param tags A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * &gt; **NOTE::** Please note that listeners that are attached to Application Load Balancers must use either `HTTP` or `HTTPS` protocols while listeners that are attached to Network Load Balancers must use the `TCP` protocol.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * &gt; **NOTE::** Please note that listeners that are attached to Application Load Balancers must use either `HTTP` or `HTTPS` protocols while listeners that are attached to Network Load Balancers must use the `TCP` protocol.
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

        public ListenerState build() {
            return $;
        }
    }

}
