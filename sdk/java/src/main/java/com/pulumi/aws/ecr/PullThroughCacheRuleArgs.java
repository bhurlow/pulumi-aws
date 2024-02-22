// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ecr;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PullThroughCacheRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final PullThroughCacheRuleArgs Empty = new PullThroughCacheRuleArgs();

    /**
     * ARN of the Secret which will be used to authenticate against the registry.
     * 
     */
    @Import(name="credentialArn")
    private @Nullable Output<String> credentialArn;

    /**
     * @return ARN of the Secret which will be used to authenticate against the registry.
     * 
     */
    public Optional<Output<String>> credentialArn() {
        return Optional.ofNullable(this.credentialArn);
    }

    /**
     * The repository name prefix to use when caching images from the source registry.
     * 
     */
    @Import(name="ecrRepositoryPrefix", required=true)
    private Output<String> ecrRepositoryPrefix;

    /**
     * @return The repository name prefix to use when caching images from the source registry.
     * 
     */
    public Output<String> ecrRepositoryPrefix() {
        return this.ecrRepositoryPrefix;
    }

    /**
     * The registry URL of the upstream public registry to use as the source.
     * 
     */
    @Import(name="upstreamRegistryUrl", required=true)
    private Output<String> upstreamRegistryUrl;

    /**
     * @return The registry URL of the upstream public registry to use as the source.
     * 
     */
    public Output<String> upstreamRegistryUrl() {
        return this.upstreamRegistryUrl;
    }

    private PullThroughCacheRuleArgs() {}

    private PullThroughCacheRuleArgs(PullThroughCacheRuleArgs $) {
        this.credentialArn = $.credentialArn;
        this.ecrRepositoryPrefix = $.ecrRepositoryPrefix;
        this.upstreamRegistryUrl = $.upstreamRegistryUrl;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PullThroughCacheRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PullThroughCacheRuleArgs $;

        public Builder() {
            $ = new PullThroughCacheRuleArgs();
        }

        public Builder(PullThroughCacheRuleArgs defaults) {
            $ = new PullThroughCacheRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param credentialArn ARN of the Secret which will be used to authenticate against the registry.
         * 
         * @return builder
         * 
         */
        public Builder credentialArn(@Nullable Output<String> credentialArn) {
            $.credentialArn = credentialArn;
            return this;
        }

        /**
         * @param credentialArn ARN of the Secret which will be used to authenticate against the registry.
         * 
         * @return builder
         * 
         */
        public Builder credentialArn(String credentialArn) {
            return credentialArn(Output.of(credentialArn));
        }

        /**
         * @param ecrRepositoryPrefix The repository name prefix to use when caching images from the source registry.
         * 
         * @return builder
         * 
         */
        public Builder ecrRepositoryPrefix(Output<String> ecrRepositoryPrefix) {
            $.ecrRepositoryPrefix = ecrRepositoryPrefix;
            return this;
        }

        /**
         * @param ecrRepositoryPrefix The repository name prefix to use when caching images from the source registry.
         * 
         * @return builder
         * 
         */
        public Builder ecrRepositoryPrefix(String ecrRepositoryPrefix) {
            return ecrRepositoryPrefix(Output.of(ecrRepositoryPrefix));
        }

        /**
         * @param upstreamRegistryUrl The registry URL of the upstream public registry to use as the source.
         * 
         * @return builder
         * 
         */
        public Builder upstreamRegistryUrl(Output<String> upstreamRegistryUrl) {
            $.upstreamRegistryUrl = upstreamRegistryUrl;
            return this;
        }

        /**
         * @param upstreamRegistryUrl The registry URL of the upstream public registry to use as the source.
         * 
         * @return builder
         * 
         */
        public Builder upstreamRegistryUrl(String upstreamRegistryUrl) {
            return upstreamRegistryUrl(Output.of(upstreamRegistryUrl));
        }

        public PullThroughCacheRuleArgs build() {
            if ($.ecrRepositoryPrefix == null) {
                throw new MissingRequiredPropertyException("PullThroughCacheRuleArgs", "ecrRepositoryPrefix");
            }
            if ($.upstreamRegistryUrl == null) {
                throw new MissingRequiredPropertyException("PullThroughCacheRuleArgs", "upstreamRegistryUrl");
            }
            return $;
        }
    }

}
