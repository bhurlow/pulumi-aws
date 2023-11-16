// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sagemaker.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class UserProfileUserSettingsCanvasAppSettingsKendraSettings {
    /**
     * @return Describes whether the document querying feature is enabled or disabled in the Canvas application. Valid values are `ENABLED` and `DISABLED`.
     * 
     */
    private @Nullable String status;

    private UserProfileUserSettingsCanvasAppSettingsKendraSettings() {}
    /**
     * @return Describes whether the document querying feature is enabled or disabled in the Canvas application. Valid values are `ENABLED` and `DISABLED`.
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(UserProfileUserSettingsCanvasAppSettingsKendraSettings defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String status;
        public Builder() {}
        public Builder(UserProfileUserSettingsCanvasAppSettingsKendraSettings defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder status(@Nullable String status) {
            this.status = status;
            return this;
        }
        public UserProfileUserSettingsCanvasAppSettingsKendraSettings build() {
            final var o = new UserProfileUserSettingsCanvasAppSettingsKendraSettings();
            o.status = status;
            return o;
        }
    }
}
