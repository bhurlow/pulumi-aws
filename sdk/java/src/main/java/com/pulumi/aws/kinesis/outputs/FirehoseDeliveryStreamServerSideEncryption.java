// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.kinesis.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class FirehoseDeliveryStreamServerSideEncryption {
    /**
     * @return Whether to enable encryption at rest. Default is `false`.
     * 
     */
    private @Nullable Boolean enabled;
    /**
     * @return Amazon Resource Name (ARN) of the encryption key. Required when `key_type` is `CUSTOMER_MANAGED_CMK`.
     * 
     * The `extended_s3_configuration` object supports the same fields from s3_configuration as well as the following:
     * 
     */
    private @Nullable String keyArn;
    /**
     * @return Type of encryption key. Default is `AWS_OWNED_CMK`. Valid values are `AWS_OWNED_CMK` and `CUSTOMER_MANAGED_CMK`
     * 
     */
    private @Nullable String keyType;

    private FirehoseDeliveryStreamServerSideEncryption() {}
    /**
     * @return Whether to enable encryption at rest. Default is `false`.
     * 
     */
    public Optional<Boolean> enabled() {
        return Optional.ofNullable(this.enabled);
    }
    /**
     * @return Amazon Resource Name (ARN) of the encryption key. Required when `key_type` is `CUSTOMER_MANAGED_CMK`.
     * 
     * The `extended_s3_configuration` object supports the same fields from s3_configuration as well as the following:
     * 
     */
    public Optional<String> keyArn() {
        return Optional.ofNullable(this.keyArn);
    }
    /**
     * @return Type of encryption key. Default is `AWS_OWNED_CMK`. Valid values are `AWS_OWNED_CMK` and `CUSTOMER_MANAGED_CMK`
     * 
     */
    public Optional<String> keyType() {
        return Optional.ofNullable(this.keyType);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FirehoseDeliveryStreamServerSideEncryption defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean enabled;
        private @Nullable String keyArn;
        private @Nullable String keyType;
        public Builder() {}
        public Builder(FirehoseDeliveryStreamServerSideEncryption defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enabled = defaults.enabled;
    	      this.keyArn = defaults.keyArn;
    	      this.keyType = defaults.keyType;
        }

        @CustomType.Setter
        public Builder enabled(@Nullable Boolean enabled) {
            this.enabled = enabled;
            return this;
        }
        @CustomType.Setter
        public Builder keyArn(@Nullable String keyArn) {
            this.keyArn = keyArn;
            return this;
        }
        @CustomType.Setter
        public Builder keyType(@Nullable String keyType) {
            this.keyType = keyType;
            return this;
        }
        public FirehoseDeliveryStreamServerSideEncryption build() {
            final var o = new FirehoseDeliveryStreamServerSideEncryption();
            o.enabled = enabled;
            o.keyArn = keyArn;
            o.keyType = keyType;
            return o;
        }
    }
}
