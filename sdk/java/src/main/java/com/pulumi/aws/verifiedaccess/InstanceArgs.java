// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.verifiedaccess;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceArgs Empty = new InstanceArgs();

    /**
     * A description for the AWS Verified Access Instance.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description for the AWS Verified Access Instance.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Enable or disable support for Federal Information Processing Standards (FIPS) on the AWS Verified Access Instance.
     * 
     */
    @Import(name="fipsEnabled")
    private @Nullable Output<Boolean> fipsEnabled;

    /**
     * @return Enable or disable support for Federal Information Processing Standards (FIPS) on the AWS Verified Access Instance.
     * 
     */
    public Optional<Output<Boolean>> fipsEnabled() {
        return Optional.ofNullable(this.fipsEnabled);
    }

    /**
     * Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private InstanceArgs() {}

    private InstanceArgs(InstanceArgs $) {
        this.description = $.description;
        this.fipsEnabled = $.fipsEnabled;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceArgs $;

        public Builder() {
            $ = new InstanceArgs();
        }

        public Builder(InstanceArgs defaults) {
            $ = new InstanceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description A description for the AWS Verified Access Instance.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description for the AWS Verified Access Instance.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param fipsEnabled Enable or disable support for Federal Information Processing Standards (FIPS) on the AWS Verified Access Instance.
         * 
         * @return builder
         * 
         */
        public Builder fipsEnabled(@Nullable Output<Boolean> fipsEnabled) {
            $.fipsEnabled = fipsEnabled;
            return this;
        }

        /**
         * @param fipsEnabled Enable or disable support for Federal Information Processing Standards (FIPS) on the AWS Verified Access Instance.
         * 
         * @return builder
         * 
         */
        public Builder fipsEnabled(Boolean fipsEnabled) {
            return fipsEnabled(Output.of(fipsEnabled));
        }

        /**
         * @param tags Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public InstanceArgs build() {
            return $;
        }
    }

}
