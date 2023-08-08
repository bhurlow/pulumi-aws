// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.opensearch.inputs;

import com.pulumi.aws.opensearch.inputs.DomainOffPeakWindowOptionsOffPeakWindowArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DomainOffPeakWindowOptionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final DomainOffPeakWindowOptionsArgs Empty = new DomainOffPeakWindowOptionsArgs();

    /**
     * Enabled disabled toggle for off-peak update window.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Enabled disabled toggle for off-peak update window.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    @Import(name="offPeakWindow")
    private @Nullable Output<DomainOffPeakWindowOptionsOffPeakWindowArgs> offPeakWindow;

    public Optional<Output<DomainOffPeakWindowOptionsOffPeakWindowArgs>> offPeakWindow() {
        return Optional.ofNullable(this.offPeakWindow);
    }

    private DomainOffPeakWindowOptionsArgs() {}

    private DomainOffPeakWindowOptionsArgs(DomainOffPeakWindowOptionsArgs $) {
        this.enabled = $.enabled;
        this.offPeakWindow = $.offPeakWindow;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DomainOffPeakWindowOptionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DomainOffPeakWindowOptionsArgs $;

        public Builder() {
            $ = new DomainOffPeakWindowOptionsArgs();
        }

        public Builder(DomainOffPeakWindowOptionsArgs defaults) {
            $ = new DomainOffPeakWindowOptionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enabled Enabled disabled toggle for off-peak update window.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Enabled disabled toggle for off-peak update window.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        public Builder offPeakWindow(@Nullable Output<DomainOffPeakWindowOptionsOffPeakWindowArgs> offPeakWindow) {
            $.offPeakWindow = offPeakWindow;
            return this;
        }

        public Builder offPeakWindow(DomainOffPeakWindowOptionsOffPeakWindowArgs offPeakWindow) {
            return offPeakWindow(Output.of(offPeakWindow));
        }

        public DomainOffPeakWindowOptionsArgs build() {
            return $;
        }
    }

}
