// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appsync.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GraphQLApiAdditionalAuthenticationProviderUserPoolConfig {
    /**
     * @return A regular expression for validating the incoming Amazon Cognito User Pool app client ID.
     * 
     */
    private final @Nullable String appIdClientRegex;
    /**
     * @return The AWS region in which the user pool was created.
     * 
     */
    private final @Nullable String awsRegion;
    /**
     * @return The user pool ID.
     * 
     */
    private final String userPoolId;

    @CustomType.Constructor
    private GraphQLApiAdditionalAuthenticationProviderUserPoolConfig(
        @CustomType.Parameter("appIdClientRegex") @Nullable String appIdClientRegex,
        @CustomType.Parameter("awsRegion") @Nullable String awsRegion,
        @CustomType.Parameter("userPoolId") String userPoolId) {
        this.appIdClientRegex = appIdClientRegex;
        this.awsRegion = awsRegion;
        this.userPoolId = userPoolId;
    }

    /**
     * @return A regular expression for validating the incoming Amazon Cognito User Pool app client ID.
     * 
     */
    public Optional<String> appIdClientRegex() {
        return Optional.ofNullable(this.appIdClientRegex);
    }
    /**
     * @return The AWS region in which the user pool was created.
     * 
     */
    public Optional<String> awsRegion() {
        return Optional.ofNullable(this.awsRegion);
    }
    /**
     * @return The user pool ID.
     * 
     */
    public String userPoolId() {
        return this.userPoolId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GraphQLApiAdditionalAuthenticationProviderUserPoolConfig defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String appIdClientRegex;
        private @Nullable String awsRegion;
        private String userPoolId;

        public Builder() {
    	      // Empty
        }

        public Builder(GraphQLApiAdditionalAuthenticationProviderUserPoolConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.appIdClientRegex = defaults.appIdClientRegex;
    	      this.awsRegion = defaults.awsRegion;
    	      this.userPoolId = defaults.userPoolId;
        }

        public Builder appIdClientRegex(@Nullable String appIdClientRegex) {
            this.appIdClientRegex = appIdClientRegex;
            return this;
        }
        public Builder awsRegion(@Nullable String awsRegion) {
            this.awsRegion = awsRegion;
            return this;
        }
        public Builder userPoolId(String userPoolId) {
            this.userPoolId = Objects.requireNonNull(userPoolId);
            return this;
        }        public GraphQLApiAdditionalAuthenticationProviderUserPoolConfig build() {
            return new GraphQLApiAdditionalAuthenticationProviderUserPoolConfig(appIdClientRegex, awsRegion, userPoolId);
        }
    }
}