// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.medialive.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ChannelInputSpecification {
    private String codec;
    private String inputResolution;
    private String maximumBitrate;

    private ChannelInputSpecification() {}
    public String codec() {
        return this.codec;
    }
    public String inputResolution() {
        return this.inputResolution;
    }
    public String maximumBitrate() {
        return this.maximumBitrate;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ChannelInputSpecification defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String codec;
        private String inputResolution;
        private String maximumBitrate;
        public Builder() {}
        public Builder(ChannelInputSpecification defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.codec = defaults.codec;
    	      this.inputResolution = defaults.inputResolution;
    	      this.maximumBitrate = defaults.maximumBitrate;
        }

        @CustomType.Setter
        public Builder codec(String codec) {
            this.codec = Objects.requireNonNull(codec);
            return this;
        }
        @CustomType.Setter
        public Builder inputResolution(String inputResolution) {
            this.inputResolution = Objects.requireNonNull(inputResolution);
            return this;
        }
        @CustomType.Setter
        public Builder maximumBitrate(String maximumBitrate) {
            this.maximumBitrate = Objects.requireNonNull(maximumBitrate);
            return this;
        }
        public ChannelInputSpecification build() {
            final var o = new ChannelInputSpecification();
            o.codec = codec;
            o.inputResolution = inputResolution;
            o.maximumBitrate = maximumBitrate;
            return o;
        }
    }
}
