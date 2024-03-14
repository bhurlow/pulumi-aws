// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.securitylake.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class SubscriberSubscriberIdentity {
    /**
     * @return The AWS Regions where Security Lake is automatically enabled.
     * 
     */
    private String externalId;
    /**
     * @return Provides encryption details of Amazon Security Lake object.
     * 
     */
    private String principal;

    private SubscriberSubscriberIdentity() {}
    /**
     * @return The AWS Regions where Security Lake is automatically enabled.
     * 
     */
    public String externalId() {
        return this.externalId;
    }
    /**
     * @return Provides encryption details of Amazon Security Lake object.
     * 
     */
    public String principal() {
        return this.principal;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SubscriberSubscriberIdentity defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String externalId;
        private String principal;
        public Builder() {}
        public Builder(SubscriberSubscriberIdentity defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.externalId = defaults.externalId;
    	      this.principal = defaults.principal;
        }

        @CustomType.Setter
        public Builder externalId(String externalId) {
            if (externalId == null) {
              throw new MissingRequiredPropertyException("SubscriberSubscriberIdentity", "externalId");
            }
            this.externalId = externalId;
            return this;
        }
        @CustomType.Setter
        public Builder principal(String principal) {
            if (principal == null) {
              throw new MissingRequiredPropertyException("SubscriberSubscriberIdentity", "principal");
            }
            this.principal = principal;
            return this;
        }
        public SubscriberSubscriberIdentity build() {
            final var _resultValue = new SubscriberSubscriberIdentity();
            _resultValue.externalId = externalId;
            _resultValue.principal = principal;
            return _resultValue;
        }
    }
}
