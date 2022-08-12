// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.organizations.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class OrganizationRootPolicyType {
    /**
     * @return The status of the policy type as it relates to the associated root
     * 
     */
    private final @Nullable String status;
    private final @Nullable String type;

    @CustomType.Constructor
    private OrganizationRootPolicyType(
        @CustomType.Parameter("status") @Nullable String status,
        @CustomType.Parameter("type") @Nullable String type) {
        this.status = status;
        this.type = type;
    }

    /**
     * @return The status of the policy type as it relates to the associated root
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(OrganizationRootPolicyType defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String status;
        private @Nullable String type;

        public Builder() {
    	      // Empty
        }

        public Builder(OrganizationRootPolicyType defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.status = defaults.status;
    	      this.type = defaults.type;
        }

        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        public Builder type(@Nullable String type) {
            this.type = type;
            return this;
        }        public OrganizationRootPolicyType build() {
            return new OrganizationRootPolicyType(status, type);
        }
    }
}