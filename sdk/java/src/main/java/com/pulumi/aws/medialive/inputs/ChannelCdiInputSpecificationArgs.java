// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.medialive.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class ChannelCdiInputSpecificationArgs extends com.pulumi.resources.ResourceArgs {

    public static final ChannelCdiInputSpecificationArgs Empty = new ChannelCdiInputSpecificationArgs();

    /**
     * - Maximum CDI input resolution.
     * 
     */
    @Import(name="resolution", required=true)
    private Output<String> resolution;

    /**
     * @return - Maximum CDI input resolution.
     * 
     */
    public Output<String> resolution() {
        return this.resolution;
    }

    private ChannelCdiInputSpecificationArgs() {}

    private ChannelCdiInputSpecificationArgs(ChannelCdiInputSpecificationArgs $) {
        this.resolution = $.resolution;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ChannelCdiInputSpecificationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ChannelCdiInputSpecificationArgs $;

        public Builder() {
            $ = new ChannelCdiInputSpecificationArgs();
        }

        public Builder(ChannelCdiInputSpecificationArgs defaults) {
            $ = new ChannelCdiInputSpecificationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param resolution - Maximum CDI input resolution.
         * 
         * @return builder
         * 
         */
        public Builder resolution(Output<String> resolution) {
            $.resolution = resolution;
            return this;
        }

        /**
         * @param resolution - Maximum CDI input resolution.
         * 
         * @return builder
         * 
         */
        public Builder resolution(String resolution) {
            return resolution(Output.of(resolution));
        }

        public ChannelCdiInputSpecificationArgs build() {
            $.resolution = Objects.requireNonNull($.resolution, "expected parameter 'resolution' to be non-null");
            return $;
        }
    }

}
