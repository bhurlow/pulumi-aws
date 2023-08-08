// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.medialive.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class MultiplexProgramMultiplexProgramSettingsServiceDescriptorArgs extends com.pulumi.resources.ResourceArgs {

    public static final MultiplexProgramMultiplexProgramSettingsServiceDescriptorArgs Empty = new MultiplexProgramMultiplexProgramSettingsServiceDescriptorArgs();

    /**
     * Unique provider name.
     * 
     */
    @Import(name="providerName", required=true)
    private Output<String> providerName;

    /**
     * @return Unique provider name.
     * 
     */
    public Output<String> providerName() {
        return this.providerName;
    }

    /**
     * Unique service name.
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return Unique service name.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    private MultiplexProgramMultiplexProgramSettingsServiceDescriptorArgs() {}

    private MultiplexProgramMultiplexProgramSettingsServiceDescriptorArgs(MultiplexProgramMultiplexProgramSettingsServiceDescriptorArgs $) {
        this.providerName = $.providerName;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MultiplexProgramMultiplexProgramSettingsServiceDescriptorArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MultiplexProgramMultiplexProgramSettingsServiceDescriptorArgs $;

        public Builder() {
            $ = new MultiplexProgramMultiplexProgramSettingsServiceDescriptorArgs();
        }

        public Builder(MultiplexProgramMultiplexProgramSettingsServiceDescriptorArgs defaults) {
            $ = new MultiplexProgramMultiplexProgramSettingsServiceDescriptorArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param providerName Unique provider name.
         * 
         * @return builder
         * 
         */
        public Builder providerName(Output<String> providerName) {
            $.providerName = providerName;
            return this;
        }

        /**
         * @param providerName Unique provider name.
         * 
         * @return builder
         * 
         */
        public Builder providerName(String providerName) {
            return providerName(Output.of(providerName));
        }

        /**
         * @param serviceName Unique service name.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName Unique service name.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public MultiplexProgramMultiplexProgramSettingsServiceDescriptorArgs build() {
            $.providerName = Objects.requireNonNull($.providerName, "expected parameter 'providerName' to be non-null");
            $.serviceName = Objects.requireNonNull($.serviceName, "expected parameter 'serviceName' to be non-null");
            return $;
        }
    }

}
