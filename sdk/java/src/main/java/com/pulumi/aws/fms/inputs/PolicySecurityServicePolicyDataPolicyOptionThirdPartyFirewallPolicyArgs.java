// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.fms.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicyArgs Empty = new PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicyArgs();

    /**
     * Defines the deployment model to use for the third-party firewall policy. Valid values are `CENTRALIZED` and `DISTRIBUTED`.
     * 
     */
    @Import(name="firewallDeploymentModel")
    private @Nullable Output<String> firewallDeploymentModel;

    /**
     * @return Defines the deployment model to use for the third-party firewall policy. Valid values are `CENTRALIZED` and `DISTRIBUTED`.
     * 
     */
    public Optional<Output<String>> firewallDeploymentModel() {
        return Optional.ofNullable(this.firewallDeploymentModel);
    }

    private PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicyArgs() {}

    private PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicyArgs(PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicyArgs $) {
        this.firewallDeploymentModel = $.firewallDeploymentModel;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicyArgs $;

        public Builder() {
            $ = new PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicyArgs();
        }

        public Builder(PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicyArgs defaults) {
            $ = new PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param firewallDeploymentModel Defines the deployment model to use for the third-party firewall policy. Valid values are `CENTRALIZED` and `DISTRIBUTED`.
         * 
         * @return builder
         * 
         */
        public Builder firewallDeploymentModel(@Nullable Output<String> firewallDeploymentModel) {
            $.firewallDeploymentModel = firewallDeploymentModel;
            return this;
        }

        /**
         * @param firewallDeploymentModel Defines the deployment model to use for the third-party firewall policy. Valid values are `CENTRALIZED` and `DISTRIBUTED`.
         * 
         * @return builder
         * 
         */
        public Builder firewallDeploymentModel(String firewallDeploymentModel) {
            return firewallDeploymentModel(Output.of(firewallDeploymentModel));
        }

        public PolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicyArgs build() {
            return $;
        }
    }

}
