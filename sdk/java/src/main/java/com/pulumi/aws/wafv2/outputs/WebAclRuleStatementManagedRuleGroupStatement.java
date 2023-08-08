// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.wafv2.outputs;

import com.pulumi.aws.wafv2.outputs.WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfig;
import com.pulumi.aws.wafv2.outputs.WebAclRuleStatementManagedRuleGroupStatementRuleActionOverride;
import com.pulumi.aws.wafv2.outputs.WebAclRuleStatementManagedRuleGroupStatementScopeDownStatement;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class WebAclRuleStatementManagedRuleGroupStatement {
    /**
     * @return Additional information that&#39;s used by a managed rule group. Only one rule attribute is allowed in each config. See `managed_rule_group_configs` for more details
     * 
     */
    private @Nullable List<WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfig> managedRuleGroupConfigs;
    /**
     * @return Name of the managed rule group.
     * 
     */
    private String name;
    /**
     * @return Action settings to use in the place of the rule actions that are configured inside the rule group. You specify one override for each rule whose action you want to change. See `rule_action_override` below for details.
     * 
     */
    private @Nullable List<WebAclRuleStatementManagedRuleGroupStatementRuleActionOverride> ruleActionOverrides;
    /**
     * @return Narrows the scope of the statement to matching web requests. This can be any nestable statement, and you can nest statements at any level below this scope-down statement. See `statement` above for details.
     * 
     */
    private @Nullable WebAclRuleStatementManagedRuleGroupStatementScopeDownStatement scopeDownStatement;
    /**
     * @return Name of the managed rule group vendor.
     * 
     */
    private String vendorName;
    /**
     * @return Version of the managed rule group. You can set `Version_1.0` or `Version_1.1` etc. If you want to use the default version, do not set anything.
     * 
     */
    private @Nullable String version;

    private WebAclRuleStatementManagedRuleGroupStatement() {}
    /**
     * @return Additional information that&#39;s used by a managed rule group. Only one rule attribute is allowed in each config. See `managed_rule_group_configs` for more details
     * 
     */
    public List<WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfig> managedRuleGroupConfigs() {
        return this.managedRuleGroupConfigs == null ? List.of() : this.managedRuleGroupConfigs;
    }
    /**
     * @return Name of the managed rule group.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Action settings to use in the place of the rule actions that are configured inside the rule group. You specify one override for each rule whose action you want to change. See `rule_action_override` below for details.
     * 
     */
    public List<WebAclRuleStatementManagedRuleGroupStatementRuleActionOverride> ruleActionOverrides() {
        return this.ruleActionOverrides == null ? List.of() : this.ruleActionOverrides;
    }
    /**
     * @return Narrows the scope of the statement to matching web requests. This can be any nestable statement, and you can nest statements at any level below this scope-down statement. See `statement` above for details.
     * 
     */
    public Optional<WebAclRuleStatementManagedRuleGroupStatementScopeDownStatement> scopeDownStatement() {
        return Optional.ofNullable(this.scopeDownStatement);
    }
    /**
     * @return Name of the managed rule group vendor.
     * 
     */
    public String vendorName() {
        return this.vendorName;
    }
    /**
     * @return Version of the managed rule group. You can set `Version_1.0` or `Version_1.1` etc. If you want to use the default version, do not set anything.
     * 
     */
    public Optional<String> version() {
        return Optional.ofNullable(this.version);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(WebAclRuleStatementManagedRuleGroupStatement defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfig> managedRuleGroupConfigs;
        private String name;
        private @Nullable List<WebAclRuleStatementManagedRuleGroupStatementRuleActionOverride> ruleActionOverrides;
        private @Nullable WebAclRuleStatementManagedRuleGroupStatementScopeDownStatement scopeDownStatement;
        private String vendorName;
        private @Nullable String version;
        public Builder() {}
        public Builder(WebAclRuleStatementManagedRuleGroupStatement defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.managedRuleGroupConfigs = defaults.managedRuleGroupConfigs;
    	      this.name = defaults.name;
    	      this.ruleActionOverrides = defaults.ruleActionOverrides;
    	      this.scopeDownStatement = defaults.scopeDownStatement;
    	      this.vendorName = defaults.vendorName;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder managedRuleGroupConfigs(@Nullable List<WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfig> managedRuleGroupConfigs) {
            this.managedRuleGroupConfigs = managedRuleGroupConfigs;
            return this;
        }
        public Builder managedRuleGroupConfigs(WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfig... managedRuleGroupConfigs) {
            return managedRuleGroupConfigs(List.of(managedRuleGroupConfigs));
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder ruleActionOverrides(@Nullable List<WebAclRuleStatementManagedRuleGroupStatementRuleActionOverride> ruleActionOverrides) {
            this.ruleActionOverrides = ruleActionOverrides;
            return this;
        }
        public Builder ruleActionOverrides(WebAclRuleStatementManagedRuleGroupStatementRuleActionOverride... ruleActionOverrides) {
            return ruleActionOverrides(List.of(ruleActionOverrides));
        }
        @CustomType.Setter
        public Builder scopeDownStatement(@Nullable WebAclRuleStatementManagedRuleGroupStatementScopeDownStatement scopeDownStatement) {
            this.scopeDownStatement = scopeDownStatement;
            return this;
        }
        @CustomType.Setter
        public Builder vendorName(String vendorName) {
            this.vendorName = Objects.requireNonNull(vendorName);
            return this;
        }
        @CustomType.Setter
        public Builder version(@Nullable String version) {
            this.version = version;
            return this;
        }
        public WebAclRuleStatementManagedRuleGroupStatement build() {
            final var o = new WebAclRuleStatementManagedRuleGroupStatement();
            o.managedRuleGroupConfigs = managedRuleGroupConfigs;
            o.name = name;
            o.ruleActionOverrides = ruleActionOverrides;
            o.scopeDownStatement = scopeDownStatement;
            o.vendorName = vendorName;
            o.version = version;
            return o;
        }
    }
}
