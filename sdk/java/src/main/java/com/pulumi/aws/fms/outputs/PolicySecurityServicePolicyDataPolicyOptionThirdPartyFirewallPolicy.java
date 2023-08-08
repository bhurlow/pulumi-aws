// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.fms.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicy {
    /**
     * @return Defines the deployment model to use for the third-party firewall policy. Valid values are `CENTRALIZED` and `DISTRIBUTED`.
     * 
     */
    private @Nullable String firewallDeploymentModel;

    private PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicy() {}
    /**
     * @return Defines the deployment model to use for the third-party firewall policy. Valid values are `CENTRALIZED` and `DISTRIBUTED`.
     * 
     */
    public Optional<String> firewallDeploymentModel() {
        return Optional.ofNullable(this.firewallDeploymentModel);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String firewallDeploymentModel;
        public Builder() {}
        public Builder(PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.firewallDeploymentModel = defaults.firewallDeploymentModel;
        }

        @CustomType.Setter
        public Builder firewallDeploymentModel(@Nullable String firewallDeploymentModel) {
            this.firewallDeploymentModel = firewallDeploymentModel;
            return this;
        }
        public PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicy build() {
            final var o = new PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicy();
            o.firewallDeploymentModel = firewallDeploymentModel;
            return o;
        }
    }
}
