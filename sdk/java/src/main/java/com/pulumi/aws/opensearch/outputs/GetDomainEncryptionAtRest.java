// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.opensearch.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetDomainEncryptionAtRest {
    /**
     * @return Enabled disabled toggle for off-peak update window
     * 
     */
    private Boolean enabled;
    /**
     * @return KMS key id used to encrypt data at rest.
     * 
     */
    private String kmsKeyId;

    private GetDomainEncryptionAtRest() {}
    /**
     * @return Enabled disabled toggle for off-peak update window
     * 
     */
    public Boolean enabled() {
        return this.enabled;
    }
    /**
     * @return KMS key id used to encrypt data at rest.
     * 
     */
    public String kmsKeyId() {
        return this.kmsKeyId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDomainEncryptionAtRest defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean enabled;
        private String kmsKeyId;
        public Builder() {}
        public Builder(GetDomainEncryptionAtRest defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enabled = defaults.enabled;
    	      this.kmsKeyId = defaults.kmsKeyId;
        }

        @CustomType.Setter
        public Builder enabled(Boolean enabled) {
            this.enabled = Objects.requireNonNull(enabled);
            return this;
        }
        @CustomType.Setter
        public Builder kmsKeyId(String kmsKeyId) {
            this.kmsKeyId = Objects.requireNonNull(kmsKeyId);
            return this;
        }
        public GetDomainEncryptionAtRest build() {
            final var o = new GetDomainEncryptionAtRest();
            o.enabled = enabled;
            o.kmsKeyId = kmsKeyId;
            return o;
        }
    }
}
