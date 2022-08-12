// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class GetLocalGatewayVirtualInterfaceGroupsFilterArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetLocalGatewayVirtualInterfaceGroupsFilterArgs Empty = new GetLocalGatewayVirtualInterfaceGroupsFilterArgs();

    /**
     * Name of the filter.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the filter.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * List of one or more values for the filter.
     * 
     */
    @Import(name="values", required=true)
    private Output<List<String>> values;

    /**
     * @return List of one or more values for the filter.
     * 
     */
    public Output<List<String>> values() {
        return this.values;
    }

    private GetLocalGatewayVirtualInterfaceGroupsFilterArgs() {}

    private GetLocalGatewayVirtualInterfaceGroupsFilterArgs(GetLocalGatewayVirtualInterfaceGroupsFilterArgs $) {
        this.name = $.name;
        this.values = $.values;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetLocalGatewayVirtualInterfaceGroupsFilterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLocalGatewayVirtualInterfaceGroupsFilterArgs $;

        public Builder() {
            $ = new GetLocalGatewayVirtualInterfaceGroupsFilterArgs();
        }

        public Builder(GetLocalGatewayVirtualInterfaceGroupsFilterArgs defaults) {
            $ = new GetLocalGatewayVirtualInterfaceGroupsFilterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Name of the filter.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the filter.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param values List of one or more values for the filter.
         * 
         * @return builder
         * 
         */
        public Builder values(Output<List<String>> values) {
            $.values = values;
            return this;
        }

        /**
         * @param values List of one or more values for the filter.
         * 
         * @return builder
         * 
         */
        public Builder values(List<String> values) {
            return values(Output.of(values));
        }

        /**
         * @param values List of one or more values for the filter.
         * 
         * @return builder
         * 
         */
        public Builder values(String... values) {
            return values(List.of(values));
        }

        public GetLocalGatewayVirtualInterfaceGroupsFilterArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.values = Objects.requireNonNull($.values, "expected parameter 'values' to be non-null");
            return $;
        }
    }

}