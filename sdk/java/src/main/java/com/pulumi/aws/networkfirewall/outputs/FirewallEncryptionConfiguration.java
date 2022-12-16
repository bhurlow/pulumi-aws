// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.networkfirewall.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class FirewallEncryptionConfiguration {
    /**
     * @return The ID of the customer managed key. You can use any of the [key identifiers](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id) that KMS supports, unless you&#39;re using a key that&#39;s managed by another account. If you&#39;re using a key managed by another account, then specify the key ARN.
     * 
     */
    private @Nullable String keyId;
    /**
     * @return The type of AWS KMS key to use for encryption of your Network Firewall resources. Valid values are `CUSTOMER_KMS` and `AWS_OWNED_KMS_KEY`.
     * 
     */
    private String type;

    private FirewallEncryptionConfiguration() {}
    /**
     * @return The ID of the customer managed key. You can use any of the [key identifiers](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id) that KMS supports, unless you&#39;re using a key that&#39;s managed by another account. If you&#39;re using a key managed by another account, then specify the key ARN.
     * 
     */
    public Optional<String> keyId() {
        return Optional.ofNullable(this.keyId);
    }
    /**
     * @return The type of AWS KMS key to use for encryption of your Network Firewall resources. Valid values are `CUSTOMER_KMS` and `AWS_OWNED_KMS_KEY`.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FirewallEncryptionConfiguration defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String keyId;
        private String type;
        public Builder() {}
        public Builder(FirewallEncryptionConfiguration defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.keyId = defaults.keyId;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder keyId(@Nullable String keyId) {
            this.keyId = keyId;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public FirewallEncryptionConfiguration build() {
            final var o = new FirewallEncryptionConfiguration();
            o.keyId = keyId;
            o.type = type;
            return o;
        }
    }
}
