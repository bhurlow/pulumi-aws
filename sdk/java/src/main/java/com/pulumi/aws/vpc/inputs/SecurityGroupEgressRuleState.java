// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.vpc.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecurityGroupEgressRuleState extends com.pulumi.resources.ResourceArgs {

    public static final SecurityGroupEgressRuleState Empty = new SecurityGroupEgressRuleState();

    /**
     * The Amazon Resource Name (ARN) of the security group rule.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) of the security group rule.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The destination IPv4 CIDR range.
     * 
     */
    @Import(name="cidrIpv4")
    private @Nullable Output<String> cidrIpv4;

    /**
     * @return The destination IPv4 CIDR range.
     * 
     */
    public Optional<Output<String>> cidrIpv4() {
        return Optional.ofNullable(this.cidrIpv4);
    }

    /**
     * The destination IPv6 CIDR range.
     * 
     */
    @Import(name="cidrIpv6")
    private @Nullable Output<String> cidrIpv6;

    /**
     * @return The destination IPv6 CIDR range.
     * 
     */
    public Optional<Output<String>> cidrIpv6() {
        return Optional.ofNullable(this.cidrIpv6);
    }

    /**
     * The security group rule description.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The security group rule description.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type.
     * 
     */
    @Import(name="fromPort")
    private @Nullable Output<Integer> fromPort;

    /**
     * @return The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type.
     * 
     */
    public Optional<Output<Integer>> fromPort() {
        return Optional.ofNullable(this.fromPort);
    }

    /**
     * The IP protocol name or number. Use `-1` to specify all protocols. Note that if `ip_protocol` is set to `-1`, it translates to all protocols, all port ranges, and `from_port` and `to_port` values should not be defined.
     * 
     */
    @Import(name="ipProtocol")
    private @Nullable Output<String> ipProtocol;

    /**
     * @return The IP protocol name or number. Use `-1` to specify all protocols. Note that if `ip_protocol` is set to `-1`, it translates to all protocols, all port ranges, and `from_port` and `to_port` values should not be defined.
     * 
     */
    public Optional<Output<String>> ipProtocol() {
        return Optional.ofNullable(this.ipProtocol);
    }

    /**
     * The ID of the destination prefix list.
     * 
     */
    @Import(name="prefixListId")
    private @Nullable Output<String> prefixListId;

    /**
     * @return The ID of the destination prefix list.
     * 
     */
    public Optional<Output<String>> prefixListId() {
        return Optional.ofNullable(this.prefixListId);
    }

    /**
     * The destination security group that is referenced in the rule.
     * 
     */
    @Import(name="referencedSecurityGroupId")
    private @Nullable Output<String> referencedSecurityGroupId;

    /**
     * @return The destination security group that is referenced in the rule.
     * 
     */
    public Optional<Output<String>> referencedSecurityGroupId() {
        return Optional.ofNullable(this.referencedSecurityGroupId);
    }

    /**
     * The ID of the security group.
     * 
     */
    @Import(name="securityGroupId")
    private @Nullable Output<String> securityGroupId;

    /**
     * @return The ID of the security group.
     * 
     */
    public Optional<Output<String>> securityGroupId() {
        return Optional.ofNullable(this.securityGroupId);
    }

    /**
     * The ID of the security group rule.
     * 
     */
    @Import(name="securityGroupRuleId")
    private @Nullable Output<String> securityGroupRuleId;

    /**
     * @return The ID of the security group rule.
     * 
     */
    public Optional<Output<String>> securityGroupRuleId() {
        return Optional.ofNullable(this.securityGroupRuleId);
    }

    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.
     * 
     */
    @Import(name="toPort")
    private @Nullable Output<Integer> toPort;

    /**
     * @return The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.
     * 
     */
    public Optional<Output<Integer>> toPort() {
        return Optional.ofNullable(this.toPort);
    }

    private SecurityGroupEgressRuleState() {}

    private SecurityGroupEgressRuleState(SecurityGroupEgressRuleState $) {
        this.arn = $.arn;
        this.cidrIpv4 = $.cidrIpv4;
        this.cidrIpv6 = $.cidrIpv6;
        this.description = $.description;
        this.fromPort = $.fromPort;
        this.ipProtocol = $.ipProtocol;
        this.prefixListId = $.prefixListId;
        this.referencedSecurityGroupId = $.referencedSecurityGroupId;
        this.securityGroupId = $.securityGroupId;
        this.securityGroupRuleId = $.securityGroupRuleId;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.toPort = $.toPort;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecurityGroupEgressRuleState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecurityGroupEgressRuleState $;

        public Builder() {
            $ = new SecurityGroupEgressRuleState();
        }

        public Builder(SecurityGroupEgressRuleState defaults) {
            $ = new SecurityGroupEgressRuleState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The Amazon Resource Name (ARN) of the security group rule.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The Amazon Resource Name (ARN) of the security group rule.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param cidrIpv4 The destination IPv4 CIDR range.
         * 
         * @return builder
         * 
         */
        public Builder cidrIpv4(@Nullable Output<String> cidrIpv4) {
            $.cidrIpv4 = cidrIpv4;
            return this;
        }

        /**
         * @param cidrIpv4 The destination IPv4 CIDR range.
         * 
         * @return builder
         * 
         */
        public Builder cidrIpv4(String cidrIpv4) {
            return cidrIpv4(Output.of(cidrIpv4));
        }

        /**
         * @param cidrIpv6 The destination IPv6 CIDR range.
         * 
         * @return builder
         * 
         */
        public Builder cidrIpv6(@Nullable Output<String> cidrIpv6) {
            $.cidrIpv6 = cidrIpv6;
            return this;
        }

        /**
         * @param cidrIpv6 The destination IPv6 CIDR range.
         * 
         * @return builder
         * 
         */
        public Builder cidrIpv6(String cidrIpv6) {
            return cidrIpv6(Output.of(cidrIpv6));
        }

        /**
         * @param description The security group rule description.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The security group rule description.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param fromPort The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type.
         * 
         * @return builder
         * 
         */
        public Builder fromPort(@Nullable Output<Integer> fromPort) {
            $.fromPort = fromPort;
            return this;
        }

        /**
         * @param fromPort The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type.
         * 
         * @return builder
         * 
         */
        public Builder fromPort(Integer fromPort) {
            return fromPort(Output.of(fromPort));
        }

        /**
         * @param ipProtocol The IP protocol name or number. Use `-1` to specify all protocols. Note that if `ip_protocol` is set to `-1`, it translates to all protocols, all port ranges, and `from_port` and `to_port` values should not be defined.
         * 
         * @return builder
         * 
         */
        public Builder ipProtocol(@Nullable Output<String> ipProtocol) {
            $.ipProtocol = ipProtocol;
            return this;
        }

        /**
         * @param ipProtocol The IP protocol name or number. Use `-1` to specify all protocols. Note that if `ip_protocol` is set to `-1`, it translates to all protocols, all port ranges, and `from_port` and `to_port` values should not be defined.
         * 
         * @return builder
         * 
         */
        public Builder ipProtocol(String ipProtocol) {
            return ipProtocol(Output.of(ipProtocol));
        }

        /**
         * @param prefixListId The ID of the destination prefix list.
         * 
         * @return builder
         * 
         */
        public Builder prefixListId(@Nullable Output<String> prefixListId) {
            $.prefixListId = prefixListId;
            return this;
        }

        /**
         * @param prefixListId The ID of the destination prefix list.
         * 
         * @return builder
         * 
         */
        public Builder prefixListId(String prefixListId) {
            return prefixListId(Output.of(prefixListId));
        }

        /**
         * @param referencedSecurityGroupId The destination security group that is referenced in the rule.
         * 
         * @return builder
         * 
         */
        public Builder referencedSecurityGroupId(@Nullable Output<String> referencedSecurityGroupId) {
            $.referencedSecurityGroupId = referencedSecurityGroupId;
            return this;
        }

        /**
         * @param referencedSecurityGroupId The destination security group that is referenced in the rule.
         * 
         * @return builder
         * 
         */
        public Builder referencedSecurityGroupId(String referencedSecurityGroupId) {
            return referencedSecurityGroupId(Output.of(referencedSecurityGroupId));
        }

        /**
         * @param securityGroupId The ID of the security group.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupId(@Nullable Output<String> securityGroupId) {
            $.securityGroupId = securityGroupId;
            return this;
        }

        /**
         * @param securityGroupId The ID of the security group.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupId(String securityGroupId) {
            return securityGroupId(Output.of(securityGroupId));
        }

        /**
         * @param securityGroupRuleId The ID of the security group rule.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupRuleId(@Nullable Output<String> securityGroupRuleId) {
            $.securityGroupRuleId = securityGroupRuleId;
            return this;
        }

        /**
         * @param securityGroupRuleId The ID of the security group rule.
         * 
         * @return builder
         * 
         */
        public Builder securityGroupRuleId(String securityGroupRuleId) {
            return securityGroupRuleId(Output.of(securityGroupRuleId));
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
         * @param toPort The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.
         * 
         * @return builder
         * 
         */
        public Builder toPort(@Nullable Output<Integer> toPort) {
            $.toPort = toPort;
            return this;
        }

        /**
         * @param toPort The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.
         * 
         * @return builder
         * 
         */
        public Builder toPort(Integer toPort) {
            return toPort(Output.of(toPort));
        }

        public SecurityGroupEgressRuleState build() {
            return $;
        }
    }

}
