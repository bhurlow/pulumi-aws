// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.globalaccelerator;

import com.pulumi.aws.globalaccelerator.inputs.CustomRoutingListenerPortRangeArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class CustomRoutingListenerArgs extends com.pulumi.resources.ResourceArgs {

    public static final CustomRoutingListenerArgs Empty = new CustomRoutingListenerArgs();

    /**
     * The Amazon Resource Name (ARN) of a custom routing accelerator.
     * 
     */
    @Import(name="acceleratorArn", required=true)
    private Output<String> acceleratorArn;

    /**
     * @return The Amazon Resource Name (ARN) of a custom routing accelerator.
     * 
     */
    public Output<String> acceleratorArn() {
        return this.acceleratorArn;
    }

    /**
     * The list of port ranges for the connections from clients to the accelerator. Fields documented below.
     * 
     */
    @Import(name="portRanges", required=true)
    private Output<List<CustomRoutingListenerPortRangeArgs>> portRanges;

    /**
     * @return The list of port ranges for the connections from clients to the accelerator. Fields documented below.
     * 
     */
    public Output<List<CustomRoutingListenerPortRangeArgs>> portRanges() {
        return this.portRanges;
    }

    private CustomRoutingListenerArgs() {}

    private CustomRoutingListenerArgs(CustomRoutingListenerArgs $) {
        this.acceleratorArn = $.acceleratorArn;
        this.portRanges = $.portRanges;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CustomRoutingListenerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CustomRoutingListenerArgs $;

        public Builder() {
            $ = new CustomRoutingListenerArgs();
        }

        public Builder(CustomRoutingListenerArgs defaults) {
            $ = new CustomRoutingListenerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param acceleratorArn The Amazon Resource Name (ARN) of a custom routing accelerator.
         * 
         * @return builder
         * 
         */
        public Builder acceleratorArn(Output<String> acceleratorArn) {
            $.acceleratorArn = acceleratorArn;
            return this;
        }

        /**
         * @param acceleratorArn The Amazon Resource Name (ARN) of a custom routing accelerator.
         * 
         * @return builder
         * 
         */
        public Builder acceleratorArn(String acceleratorArn) {
            return acceleratorArn(Output.of(acceleratorArn));
        }

        /**
         * @param portRanges The list of port ranges for the connections from clients to the accelerator. Fields documented below.
         * 
         * @return builder
         * 
         */
        public Builder portRanges(Output<List<CustomRoutingListenerPortRangeArgs>> portRanges) {
            $.portRanges = portRanges;
            return this;
        }

        /**
         * @param portRanges The list of port ranges for the connections from clients to the accelerator. Fields documented below.
         * 
         * @return builder
         * 
         */
        public Builder portRanges(List<CustomRoutingListenerPortRangeArgs> portRanges) {
            return portRanges(Output.of(portRanges));
        }

        /**
         * @param portRanges The list of port ranges for the connections from clients to the accelerator. Fields documented below.
         * 
         * @return builder
         * 
         */
        public Builder portRanges(CustomRoutingListenerPortRangeArgs... portRanges) {
            return portRanges(List.of(portRanges));
        }

        public CustomRoutingListenerArgs build() {
            $.acceleratorArn = Objects.requireNonNull($.acceleratorArn, "expected parameter 'acceleratorArn' to be non-null");
            $.portRanges = Objects.requireNonNull($.portRanges, "expected parameter 'portRanges' to be non-null");
            return $;
        }
    }

}
