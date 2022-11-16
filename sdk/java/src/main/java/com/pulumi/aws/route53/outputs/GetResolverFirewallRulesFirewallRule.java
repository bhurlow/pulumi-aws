// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.route53.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetResolverFirewallRulesFirewallRule {
    /**
     * @return The action that DNS Firewall should take on a DNS query when it matches one of the domains in the rule&#39;s domain list.
     * 
     */
    private String action;
    /**
     * @return The DNS record&#39;s type.
     * 
     */
    private String blockOverrideDnsType;
    /**
     * @return The custom DNS record to send back in response to the query.
     * 
     */
    private String blockOverrideDomain;
    /**
     * @return The recommended amount of time, in seconds, for the DNS resolver or web browser to cache the provided override record.
     * 
     */
    private Integer blockOverrideTtl;
    /**
     * @return The way that you want DNS Firewall to block the request.
     * 
     */
    private String blockResponse;
    /**
     * @return The date and time that the rule was created, in Unix time format and Coordinated Universal Time (UTC).
     * 
     */
    private String creationTime;
    /**
     * @return A unique string defined by you to identify the request.
     * 
     */
    private String creatorRequestId;
    /**
     * @return The ID of the domain list that&#39;s used in the rule.
     * 
     */
    private String firewallDomainListId;
    /**
     * @return The unique identifier of the firewall rule group that you want to retrieve the rules for.
     * 
     */
    private String firewallRuleGroupId;
    /**
     * @return The date and time that the rule was last modified, in Unix time format and Coordinated Universal Time (UTC).
     * 
     */
    private String modificationTime;
    /**
     * @return The name of the rule.
     * 
     */
    private String name;
    /**
     * @return The setting that determines the processing order of the rules in a rule group.
     * 
     */
    private Integer priority;

    private GetResolverFirewallRulesFirewallRule() {}
    /**
     * @return The action that DNS Firewall should take on a DNS query when it matches one of the domains in the rule&#39;s domain list.
     * 
     */
    public String action() {
        return this.action;
    }
    /**
     * @return The DNS record&#39;s type.
     * 
     */
    public String blockOverrideDnsType() {
        return this.blockOverrideDnsType;
    }
    /**
     * @return The custom DNS record to send back in response to the query.
     * 
     */
    public String blockOverrideDomain() {
        return this.blockOverrideDomain;
    }
    /**
     * @return The recommended amount of time, in seconds, for the DNS resolver or web browser to cache the provided override record.
     * 
     */
    public Integer blockOverrideTtl() {
        return this.blockOverrideTtl;
    }
    /**
     * @return The way that you want DNS Firewall to block the request.
     * 
     */
    public String blockResponse() {
        return this.blockResponse;
    }
    /**
     * @return The date and time that the rule was created, in Unix time format and Coordinated Universal Time (UTC).
     * 
     */
    public String creationTime() {
        return this.creationTime;
    }
    /**
     * @return A unique string defined by you to identify the request.
     * 
     */
    public String creatorRequestId() {
        return this.creatorRequestId;
    }
    /**
     * @return The ID of the domain list that&#39;s used in the rule.
     * 
     */
    public String firewallDomainListId() {
        return this.firewallDomainListId;
    }
    /**
     * @return The unique identifier of the firewall rule group that you want to retrieve the rules for.
     * 
     */
    public String firewallRuleGroupId() {
        return this.firewallRuleGroupId;
    }
    /**
     * @return The date and time that the rule was last modified, in Unix time format and Coordinated Universal Time (UTC).
     * 
     */
    public String modificationTime() {
        return this.modificationTime;
    }
    /**
     * @return The name of the rule.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The setting that determines the processing order of the rules in a rule group.
     * 
     */
    public Integer priority() {
        return this.priority;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetResolverFirewallRulesFirewallRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String action;
        private String blockOverrideDnsType;
        private String blockOverrideDomain;
        private Integer blockOverrideTtl;
        private String blockResponse;
        private String creationTime;
        private String creatorRequestId;
        private String firewallDomainListId;
        private String firewallRuleGroupId;
        private String modificationTime;
        private String name;
        private Integer priority;
        public Builder() {}
        public Builder(GetResolverFirewallRulesFirewallRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.action = defaults.action;
    	      this.blockOverrideDnsType = defaults.blockOverrideDnsType;
    	      this.blockOverrideDomain = defaults.blockOverrideDomain;
    	      this.blockOverrideTtl = defaults.blockOverrideTtl;
    	      this.blockResponse = defaults.blockResponse;
    	      this.creationTime = defaults.creationTime;
    	      this.creatorRequestId = defaults.creatorRequestId;
    	      this.firewallDomainListId = defaults.firewallDomainListId;
    	      this.firewallRuleGroupId = defaults.firewallRuleGroupId;
    	      this.modificationTime = defaults.modificationTime;
    	      this.name = defaults.name;
    	      this.priority = defaults.priority;
        }

        @CustomType.Setter
        public Builder action(String action) {
            this.action = Objects.requireNonNull(action);
            return this;
        }
        @CustomType.Setter
        public Builder blockOverrideDnsType(String blockOverrideDnsType) {
            this.blockOverrideDnsType = Objects.requireNonNull(blockOverrideDnsType);
            return this;
        }
        @CustomType.Setter
        public Builder blockOverrideDomain(String blockOverrideDomain) {
            this.blockOverrideDomain = Objects.requireNonNull(blockOverrideDomain);
            return this;
        }
        @CustomType.Setter
        public Builder blockOverrideTtl(Integer blockOverrideTtl) {
            this.blockOverrideTtl = Objects.requireNonNull(blockOverrideTtl);
            return this;
        }
        @CustomType.Setter
        public Builder blockResponse(String blockResponse) {
            this.blockResponse = Objects.requireNonNull(blockResponse);
            return this;
        }
        @CustomType.Setter
        public Builder creationTime(String creationTime) {
            this.creationTime = Objects.requireNonNull(creationTime);
            return this;
        }
        @CustomType.Setter
        public Builder creatorRequestId(String creatorRequestId) {
            this.creatorRequestId = Objects.requireNonNull(creatorRequestId);
            return this;
        }
        @CustomType.Setter
        public Builder firewallDomainListId(String firewallDomainListId) {
            this.firewallDomainListId = Objects.requireNonNull(firewallDomainListId);
            return this;
        }
        @CustomType.Setter
        public Builder firewallRuleGroupId(String firewallRuleGroupId) {
            this.firewallRuleGroupId = Objects.requireNonNull(firewallRuleGroupId);
            return this;
        }
        @CustomType.Setter
        public Builder modificationTime(String modificationTime) {
            this.modificationTime = Objects.requireNonNull(modificationTime);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder priority(Integer priority) {
            this.priority = Objects.requireNonNull(priority);
            return this;
        }
        public GetResolverFirewallRulesFirewallRule build() {
            final var o = new GetResolverFirewallRulesFirewallRule();
            o.action = action;
            o.blockOverrideDnsType = blockOverrideDnsType;
            o.blockOverrideDomain = blockOverrideDomain;
            o.blockOverrideTtl = blockOverrideTtl;
            o.blockResponse = blockResponse;
            o.creationTime = creationTime;
            o.creatorRequestId = creatorRequestId;
            o.firewallDomainListId = firewallDomainListId;
            o.firewallRuleGroupId = firewallRuleGroupId;
            o.modificationTime = modificationTime;
            o.name = name;
            o.priority = priority;
            return o;
        }
    }
}
