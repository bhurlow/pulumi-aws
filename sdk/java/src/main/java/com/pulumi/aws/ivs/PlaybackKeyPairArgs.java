// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ivs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PlaybackKeyPairArgs extends com.pulumi.resources.ResourceArgs {

    public static final PlaybackKeyPairArgs Empty = new PlaybackKeyPairArgs();

    /**
     * Playback Key Pair name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Playback Key Pair name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Public portion of a customer-generated key pair. Must be an ECDSA public key in PEM format.
     * 
     */
    @Import(name="publicKey", required=true)
    private Output<String> publicKey;

    /**
     * @return Public portion of a customer-generated key pair. Must be an ECDSA public key in PEM format.
     * 
     */
    public Output<String> publicKey() {
        return this.publicKey;
    }

    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private PlaybackKeyPairArgs() {}

    private PlaybackKeyPairArgs(PlaybackKeyPairArgs $) {
        this.name = $.name;
        this.publicKey = $.publicKey;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PlaybackKeyPairArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PlaybackKeyPairArgs $;

        public Builder() {
            $ = new PlaybackKeyPairArgs();
        }

        public Builder(PlaybackKeyPairArgs defaults) {
            $ = new PlaybackKeyPairArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Playback Key Pair name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Playback Key Pair name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param publicKey Public portion of a customer-generated key pair. Must be an ECDSA public key in PEM format.
         * 
         * @return builder
         * 
         */
        public Builder publicKey(Output<String> publicKey) {
            $.publicKey = publicKey;
            return this;
        }

        /**
         * @param publicKey Public portion of a customer-generated key pair. Must be an ECDSA public key in PEM format.
         * 
         * @return builder
         * 
         */
        public Builder publicKey(String publicKey) {
            return publicKey(Output.of(publicKey));
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public PlaybackKeyPairArgs build() {
            $.publicKey = Objects.requireNonNull($.publicKey, "expected parameter 'publicKey' to be non-null");
            return $;
        }
    }

}
