// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.networkfirewall.inputs;

import com.pulumi.aws.networkfirewall.inputs.FirewallPolicyFirewallPolicyPolicyVariablesRuleVariableIpSetArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class FirewallPolicyFirewallPolicyPolicyVariablesRuleVariableArgs extends com.pulumi.resources.ResourceArgs {

    public static final FirewallPolicyFirewallPolicyPolicyVariablesRuleVariableArgs Empty = new FirewallPolicyFirewallPolicyPolicyVariablesRuleVariableArgs();

    /**
     * A configuration block that defines a set of IP addresses. See IP Set below for details.
     * 
     */
    @Import(name="ipSet", required=true)
    private Output<FirewallPolicyFirewallPolicyPolicyVariablesRuleVariableIpSetArgs> ipSet;

    /**
     * @return A configuration block that defines a set of IP addresses. See IP Set below for details.
     * 
     */
    public Output<FirewallPolicyFirewallPolicyPolicyVariablesRuleVariableIpSetArgs> ipSet() {
        return this.ipSet;
    }

    /**
     * An alphanumeric string to identify the `ip_set`. Valid values: `HOME_NET`
     * 
     */
    @Import(name="key", required=true)
    private Output<String> key;

    /**
     * @return An alphanumeric string to identify the `ip_set`. Valid values: `HOME_NET`
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    private FirewallPolicyFirewallPolicyPolicyVariablesRuleVariableArgs() {}

    private FirewallPolicyFirewallPolicyPolicyVariablesRuleVariableArgs(FirewallPolicyFirewallPolicyPolicyVariablesRuleVariableArgs $) {
        this.ipSet = $.ipSet;
        this.key = $.key;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FirewallPolicyFirewallPolicyPolicyVariablesRuleVariableArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FirewallPolicyFirewallPolicyPolicyVariablesRuleVariableArgs $;

        public Builder() {
            $ = new FirewallPolicyFirewallPolicyPolicyVariablesRuleVariableArgs();
        }

        public Builder(FirewallPolicyFirewallPolicyPolicyVariablesRuleVariableArgs defaults) {
            $ = new FirewallPolicyFirewallPolicyPolicyVariablesRuleVariableArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ipSet A configuration block that defines a set of IP addresses. See IP Set below for details.
         * 
         * @return builder
         * 
         */
        public Builder ipSet(Output<FirewallPolicyFirewallPolicyPolicyVariablesRuleVariableIpSetArgs> ipSet) {
            $.ipSet = ipSet;
            return this;
        }

        /**
         * @param ipSet A configuration block that defines a set of IP addresses. See IP Set below for details.
         * 
         * @return builder
         * 
         */
        public Builder ipSet(FirewallPolicyFirewallPolicyPolicyVariablesRuleVariableIpSetArgs ipSet) {
            return ipSet(Output.of(ipSet));
        }

        /**
         * @param key An alphanumeric string to identify the `ip_set`. Valid values: `HOME_NET`
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key An alphanumeric string to identify the `ip_set`. Valid values: `HOME_NET`
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        public FirewallPolicyFirewallPolicyPolicyVariablesRuleVariableArgs build() {
            $.ipSet = Objects.requireNonNull($.ipSet, "expected parameter 'ipSet' to be non-null");
            $.key = Objects.requireNonNull($.key, "expected parameter 'key' to be non-null");
            return $;
        }
    }

}
