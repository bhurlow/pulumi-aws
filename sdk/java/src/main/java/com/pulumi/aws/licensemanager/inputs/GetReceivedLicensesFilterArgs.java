// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.licensemanager.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class GetReceivedLicensesFilterArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetReceivedLicensesFilterArgs Empty = new GetReceivedLicensesFilterArgs();

    /**
     * Name of the field to filter by, as defined by
     * [the underlying AWS API](https://docs.aws.amazon.com/license-manager/latest/APIReference/API_ListReceivedLicenses.html#API_ListReceivedLicenses_RequestSyntax).
     * For example, if filtering using `ProductSKU`, use:
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the field to filter by, as defined by
     * [the underlying AWS API](https://docs.aws.amazon.com/license-manager/latest/APIReference/API_ListReceivedLicenses.html#API_ListReceivedLicenses_RequestSyntax).
     * For example, if filtering using `ProductSKU`, use:
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * Set of values that are accepted for the given field.
     * 
     */
    @Import(name="values", required=true)
    private Output<List<String>> values;

    /**
     * @return Set of values that are accepted for the given field.
     * 
     */
    public Output<List<String>> values() {
        return this.values;
    }

    private GetReceivedLicensesFilterArgs() {}

    private GetReceivedLicensesFilterArgs(GetReceivedLicensesFilterArgs $) {
        this.name = $.name;
        this.values = $.values;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetReceivedLicensesFilterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetReceivedLicensesFilterArgs $;

        public Builder() {
            $ = new GetReceivedLicensesFilterArgs();
        }

        public Builder(GetReceivedLicensesFilterArgs defaults) {
            $ = new GetReceivedLicensesFilterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Name of the field to filter by, as defined by
         * [the underlying AWS API](https://docs.aws.amazon.com/license-manager/latest/APIReference/API_ListReceivedLicenses.html#API_ListReceivedLicenses_RequestSyntax).
         * For example, if filtering using `ProductSKU`, use:
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the field to filter by, as defined by
         * [the underlying AWS API](https://docs.aws.amazon.com/license-manager/latest/APIReference/API_ListReceivedLicenses.html#API_ListReceivedLicenses_RequestSyntax).
         * For example, if filtering using `ProductSKU`, use:
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param values Set of values that are accepted for the given field.
         * 
         * @return builder
         * 
         */
        public Builder values(Output<List<String>> values) {
            $.values = values;
            return this;
        }

        /**
         * @param values Set of values that are accepted for the given field.
         * 
         * @return builder
         * 
         */
        public Builder values(List<String> values) {
            return values(Output.of(values));
        }

        /**
         * @param values Set of values that are accepted for the given field.
         * 
         * @return builder
         * 
         */
        public Builder values(String... values) {
            return values(List.of(values));
        }

        public GetReceivedLicensesFilterArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.values = Objects.requireNonNull($.values, "expected parameter 'values' to be non-null");
            return $;
        }
    }

}
