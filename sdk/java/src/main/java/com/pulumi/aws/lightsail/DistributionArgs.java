// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.lightsail;

import com.pulumi.aws.lightsail.inputs.DistributionCacheBehaviorArgs;
import com.pulumi.aws.lightsail.inputs.DistributionCacheBehaviorSettingsArgs;
import com.pulumi.aws.lightsail.inputs.DistributionDefaultCacheBehaviorArgs;
import com.pulumi.aws.lightsail.inputs.DistributionOriginArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DistributionArgs extends com.pulumi.resources.ResourceArgs {

    public static final DistributionArgs Empty = new DistributionArgs();

    /**
     * Bundle ID to use for the distribution.
     * 
     */
    @Import(name="bundleId", required=true)
    private Output<String> bundleId;

    /**
     * @return Bundle ID to use for the distribution.
     * 
     */
    public Output<String> bundleId() {
        return this.bundleId;
    }

    /**
     * An object that describes the cache behavior settings of the distribution. Detailed below
     * 
     */
    @Import(name="cacheBehaviorSettings")
    private @Nullable Output<DistributionCacheBehaviorSettingsArgs> cacheBehaviorSettings;

    /**
     * @return An object that describes the cache behavior settings of the distribution. Detailed below
     * 
     */
    public Optional<Output<DistributionCacheBehaviorSettingsArgs>> cacheBehaviorSettings() {
        return Optional.ofNullable(this.cacheBehaviorSettings);
    }

    /**
     * A set of configuration blocks that describe the per-path cache behavior of the distribution. Detailed below
     * 
     */
    @Import(name="cacheBehaviors")
    private @Nullable Output<List<DistributionCacheBehaviorArgs>> cacheBehaviors;

    /**
     * @return A set of configuration blocks that describe the per-path cache behavior of the distribution. Detailed below
     * 
     */
    public Optional<Output<List<DistributionCacheBehaviorArgs>>> cacheBehaviors() {
        return Optional.ofNullable(this.cacheBehaviors);
    }

    /**
     * The name of the SSL/TLS certificate attached to the distribution, if any.
     * 
     */
    @Import(name="certificateName")
    private @Nullable Output<String> certificateName;

    /**
     * @return The name of the SSL/TLS certificate attached to the distribution, if any.
     * 
     */
    public Optional<Output<String>> certificateName() {
        return Optional.ofNullable(this.certificateName);
    }

    /**
     * Object that describes the default cache behavior of the distribution. Detailed below
     * 
     */
    @Import(name="defaultCacheBehavior", required=true)
    private Output<DistributionDefaultCacheBehaviorArgs> defaultCacheBehavior;

    /**
     * @return Object that describes the default cache behavior of the distribution. Detailed below
     * 
     */
    public Output<DistributionDefaultCacheBehaviorArgs> defaultCacheBehavior() {
        return this.defaultCacheBehavior;
    }

    /**
     * The IP address type of the distribution. Default: `dualstack`.
     * 
     */
    @Import(name="ipAddressType")
    private @Nullable Output<String> ipAddressType;

    /**
     * @return The IP address type of the distribution. Default: `dualstack`.
     * 
     */
    public Optional<Output<String>> ipAddressType() {
        return Optional.ofNullable(this.ipAddressType);
    }

    /**
     * Indicates whether the distribution is enabled. Default: `true`.
     * 
     */
    @Import(name="isEnabled")
    private @Nullable Output<Boolean> isEnabled;

    /**
     * @return Indicates whether the distribution is enabled. Default: `true`.
     * 
     */
    public Optional<Output<Boolean>> isEnabled() {
        return Optional.ofNullable(this.isEnabled);
    }

    /**
     * Name of the distribution.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the distribution.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Object that describes the origin resource of the distribution, such as a Lightsail instance, bucket, or load balancer. Detailed below
     * 
     */
    @Import(name="origin", required=true)
    private Output<DistributionOriginArgs> origin;

    /**
     * @return Object that describes the origin resource of the distribution, such as a Lightsail instance, bucket, or load balancer. Detailed below
     * 
     */
    public Output<DistributionOriginArgs> origin() {
        return this.origin;
    }

    /**
     * Map of tags for the Lightsail Distribution. If
     * configured with a provider
     * `default_tags` configuration block
     * present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Map of tags for the Lightsail Distribution. If
     * configured with a provider
     * `default_tags` configuration block
     * present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private DistributionArgs() {}

    private DistributionArgs(DistributionArgs $) {
        this.bundleId = $.bundleId;
        this.cacheBehaviorSettings = $.cacheBehaviorSettings;
        this.cacheBehaviors = $.cacheBehaviors;
        this.certificateName = $.certificateName;
        this.defaultCacheBehavior = $.defaultCacheBehavior;
        this.ipAddressType = $.ipAddressType;
        this.isEnabled = $.isEnabled;
        this.name = $.name;
        this.origin = $.origin;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DistributionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DistributionArgs $;

        public Builder() {
            $ = new DistributionArgs();
        }

        public Builder(DistributionArgs defaults) {
            $ = new DistributionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param bundleId Bundle ID to use for the distribution.
         * 
         * @return builder
         * 
         */
        public Builder bundleId(Output<String> bundleId) {
            $.bundleId = bundleId;
            return this;
        }

        /**
         * @param bundleId Bundle ID to use for the distribution.
         * 
         * @return builder
         * 
         */
        public Builder bundleId(String bundleId) {
            return bundleId(Output.of(bundleId));
        }

        /**
         * @param cacheBehaviorSettings An object that describes the cache behavior settings of the distribution. Detailed below
         * 
         * @return builder
         * 
         */
        public Builder cacheBehaviorSettings(@Nullable Output<DistributionCacheBehaviorSettingsArgs> cacheBehaviorSettings) {
            $.cacheBehaviorSettings = cacheBehaviorSettings;
            return this;
        }

        /**
         * @param cacheBehaviorSettings An object that describes the cache behavior settings of the distribution. Detailed below
         * 
         * @return builder
         * 
         */
        public Builder cacheBehaviorSettings(DistributionCacheBehaviorSettingsArgs cacheBehaviorSettings) {
            return cacheBehaviorSettings(Output.of(cacheBehaviorSettings));
        }

        /**
         * @param cacheBehaviors A set of configuration blocks that describe the per-path cache behavior of the distribution. Detailed below
         * 
         * @return builder
         * 
         */
        public Builder cacheBehaviors(@Nullable Output<List<DistributionCacheBehaviorArgs>> cacheBehaviors) {
            $.cacheBehaviors = cacheBehaviors;
            return this;
        }

        /**
         * @param cacheBehaviors A set of configuration blocks that describe the per-path cache behavior of the distribution. Detailed below
         * 
         * @return builder
         * 
         */
        public Builder cacheBehaviors(List<DistributionCacheBehaviorArgs> cacheBehaviors) {
            return cacheBehaviors(Output.of(cacheBehaviors));
        }

        /**
         * @param cacheBehaviors A set of configuration blocks that describe the per-path cache behavior of the distribution. Detailed below
         * 
         * @return builder
         * 
         */
        public Builder cacheBehaviors(DistributionCacheBehaviorArgs... cacheBehaviors) {
            return cacheBehaviors(List.of(cacheBehaviors));
        }

        /**
         * @param certificateName The name of the SSL/TLS certificate attached to the distribution, if any.
         * 
         * @return builder
         * 
         */
        public Builder certificateName(@Nullable Output<String> certificateName) {
            $.certificateName = certificateName;
            return this;
        }

        /**
         * @param certificateName The name of the SSL/TLS certificate attached to the distribution, if any.
         * 
         * @return builder
         * 
         */
        public Builder certificateName(String certificateName) {
            return certificateName(Output.of(certificateName));
        }

        /**
         * @param defaultCacheBehavior Object that describes the default cache behavior of the distribution. Detailed below
         * 
         * @return builder
         * 
         */
        public Builder defaultCacheBehavior(Output<DistributionDefaultCacheBehaviorArgs> defaultCacheBehavior) {
            $.defaultCacheBehavior = defaultCacheBehavior;
            return this;
        }

        /**
         * @param defaultCacheBehavior Object that describes the default cache behavior of the distribution. Detailed below
         * 
         * @return builder
         * 
         */
        public Builder defaultCacheBehavior(DistributionDefaultCacheBehaviorArgs defaultCacheBehavior) {
            return defaultCacheBehavior(Output.of(defaultCacheBehavior));
        }

        /**
         * @param ipAddressType The IP address type of the distribution. Default: `dualstack`.
         * 
         * @return builder
         * 
         */
        public Builder ipAddressType(@Nullable Output<String> ipAddressType) {
            $.ipAddressType = ipAddressType;
            return this;
        }

        /**
         * @param ipAddressType The IP address type of the distribution. Default: `dualstack`.
         * 
         * @return builder
         * 
         */
        public Builder ipAddressType(String ipAddressType) {
            return ipAddressType(Output.of(ipAddressType));
        }

        /**
         * @param isEnabled Indicates whether the distribution is enabled. Default: `true`.
         * 
         * @return builder
         * 
         */
        public Builder isEnabled(@Nullable Output<Boolean> isEnabled) {
            $.isEnabled = isEnabled;
            return this;
        }

        /**
         * @param isEnabled Indicates whether the distribution is enabled. Default: `true`.
         * 
         * @return builder
         * 
         */
        public Builder isEnabled(Boolean isEnabled) {
            return isEnabled(Output.of(isEnabled));
        }

        /**
         * @param name Name of the distribution.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the distribution.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param origin Object that describes the origin resource of the distribution, such as a Lightsail instance, bucket, or load balancer. Detailed below
         * 
         * @return builder
         * 
         */
        public Builder origin(Output<DistributionOriginArgs> origin) {
            $.origin = origin;
            return this;
        }

        /**
         * @param origin Object that describes the origin resource of the distribution, such as a Lightsail instance, bucket, or load balancer. Detailed below
         * 
         * @return builder
         * 
         */
        public Builder origin(DistributionOriginArgs origin) {
            return origin(Output.of(origin));
        }

        /**
         * @param tags Map of tags for the Lightsail Distribution. If
         * configured with a provider
         * `default_tags` configuration block
         * present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Map of tags for the Lightsail Distribution. If
         * configured with a provider
         * `default_tags` configuration block
         * present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public DistributionArgs build() {
            $.bundleId = Objects.requireNonNull($.bundleId, "expected parameter 'bundleId' to be non-null");
            $.defaultCacheBehavior = Objects.requireNonNull($.defaultCacheBehavior, "expected parameter 'defaultCacheBehavior' to be non-null");
            $.origin = Objects.requireNonNull($.origin, "expected parameter 'origin' to be non-null");
            return $;
        }
    }

}
