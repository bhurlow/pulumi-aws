// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sagemaker.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class DomainDefaultUserSettingsJupyterServerAppSettingsCodeRepository {
    /**
     * @return The URL of the Git repository.
     * 
     */
    private String repositoryUrl;

    private DomainDefaultUserSettingsJupyterServerAppSettingsCodeRepository() {}
    /**
     * @return The URL of the Git repository.
     * 
     */
    public String repositoryUrl() {
        return this.repositoryUrl;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DomainDefaultUserSettingsJupyterServerAppSettingsCodeRepository defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String repositoryUrl;
        public Builder() {}
        public Builder(DomainDefaultUserSettingsJupyterServerAppSettingsCodeRepository defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.repositoryUrl = defaults.repositoryUrl;
        }

        @CustomType.Setter
        public Builder repositoryUrl(String repositoryUrl) {
            this.repositoryUrl = Objects.requireNonNull(repositoryUrl);
            return this;
        }
        public DomainDefaultUserSettingsJupyterServerAppSettingsCodeRepository build() {
            final var o = new DomainDefaultUserSettingsJupyterServerAppSettingsCodeRepository();
            o.repositoryUrl = repositoryUrl;
            return o;
        }
    }
}
