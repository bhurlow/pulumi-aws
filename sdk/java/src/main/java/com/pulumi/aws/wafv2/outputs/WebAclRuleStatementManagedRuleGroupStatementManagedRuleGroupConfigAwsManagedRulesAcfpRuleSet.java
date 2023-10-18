// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.wafv2.outputs;

import com.pulumi.aws.wafv2.outputs.WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetRequestInspection;
import com.pulumi.aws.wafv2.outputs.WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetResponseInspection;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSet {
    /**
     * @return The path of the account creation endpoint for your application. This is the page on your website that accepts the completed registration form for a new user. This page must accept POST requests.
     * 
     */
    private String creationPath;
    /**
     * @return Whether or not to allow the use of regular expressions in the login page path.
     * 
     */
    private @Nullable Boolean enableRegexInPath;
    /**
     * @return The path of the account registration endpoint for your application. This is the page on your website that presents the registration form to new users. This page must accept GET text/html requests.
     * 
     */
    private String registrationPagePath;
    /**
     * @return The criteria for inspecting login requests, used by the ATP rule group to validate credentials usage. See `request_inspection` for more details.
     * 
     */
    private WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetRequestInspection requestInspection;
    /**
     * @return The criteria for inspecting responses to login requests, used by the ATP rule group to track login failure rates. Note that Response Inspection is available only on web ACLs that protect CloudFront distributions. See `response_inspection` for more details.
     * 
     */
    private @Nullable WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetResponseInspection responseInspection;

    private WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSet() {}
    /**
     * @return The path of the account creation endpoint for your application. This is the page on your website that accepts the completed registration form for a new user. This page must accept POST requests.
     * 
     */
    public String creationPath() {
        return this.creationPath;
    }
    /**
     * @return Whether or not to allow the use of regular expressions in the login page path.
     * 
     */
    public Optional<Boolean> enableRegexInPath() {
        return Optional.ofNullable(this.enableRegexInPath);
    }
    /**
     * @return The path of the account registration endpoint for your application. This is the page on your website that presents the registration form to new users. This page must accept GET text/html requests.
     * 
     */
    public String registrationPagePath() {
        return this.registrationPagePath;
    }
    /**
     * @return The criteria for inspecting login requests, used by the ATP rule group to validate credentials usage. See `request_inspection` for more details.
     * 
     */
    public WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetRequestInspection requestInspection() {
        return this.requestInspection;
    }
    /**
     * @return The criteria for inspecting responses to login requests, used by the ATP rule group to track login failure rates. Note that Response Inspection is available only on web ACLs that protect CloudFront distributions. See `response_inspection` for more details.
     * 
     */
    public Optional<WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetResponseInspection> responseInspection() {
        return Optional.ofNullable(this.responseInspection);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSet defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String creationPath;
        private @Nullable Boolean enableRegexInPath;
        private String registrationPagePath;
        private WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetRequestInspection requestInspection;
        private @Nullable WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetResponseInspection responseInspection;
        public Builder() {}
        public Builder(WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSet defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.creationPath = defaults.creationPath;
    	      this.enableRegexInPath = defaults.enableRegexInPath;
    	      this.registrationPagePath = defaults.registrationPagePath;
    	      this.requestInspection = defaults.requestInspection;
    	      this.responseInspection = defaults.responseInspection;
        }

        @CustomType.Setter
        public Builder creationPath(String creationPath) {
            this.creationPath = Objects.requireNonNull(creationPath);
            return this;
        }
        @CustomType.Setter
        public Builder enableRegexInPath(@Nullable Boolean enableRegexInPath) {
            this.enableRegexInPath = enableRegexInPath;
            return this;
        }
        @CustomType.Setter
        public Builder registrationPagePath(String registrationPagePath) {
            this.registrationPagePath = Objects.requireNonNull(registrationPagePath);
            return this;
        }
        @CustomType.Setter
        public Builder requestInspection(WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetRequestInspection requestInspection) {
            this.requestInspection = Objects.requireNonNull(requestInspection);
            return this;
        }
        @CustomType.Setter
        public Builder responseInspection(@Nullable WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetResponseInspection responseInspection) {
            this.responseInspection = responseInspection;
            return this;
        }
        public WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSet build() {
            final var o = new WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSet();
            o.creationPath = creationPath;
            o.enableRegexInPath = enableRegexInPath;
            o.registrationPagePath = registrationPagePath;
            o.requestInspection = requestInspection;
            o.responseInspection = responseInspection;
            return o;
        }
    }
}
