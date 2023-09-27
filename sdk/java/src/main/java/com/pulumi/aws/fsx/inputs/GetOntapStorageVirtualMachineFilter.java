// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.fsx.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class GetOntapStorageVirtualMachineFilter extends com.pulumi.resources.InvokeArgs {

    public static final GetOntapStorageVirtualMachineFilter Empty = new GetOntapStorageVirtualMachineFilter();

    /**
     * Name of the field to filter by, as defined by [the underlying AWS API](https://docs.aws.amazon.com/fsx/latest/APIReference/API_StorageVirtualMachineFilter.html).
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return Name of the field to filter by, as defined by [the underlying AWS API](https://docs.aws.amazon.com/fsx/latest/APIReference/API_StorageVirtualMachineFilter.html).
     * 
     */
    public String name() {
        return this.name;
    }

    /**
     * Set of values that are accepted for the given field. An SVM will be selected if any one of the given values matches.
     * 
     */
    @Import(name="values", required=true)
    private List<String> values;

    /**
     * @return Set of values that are accepted for the given field. An SVM will be selected if any one of the given values matches.
     * 
     */
    public List<String> values() {
        return this.values;
    }

    private GetOntapStorageVirtualMachineFilter() {}

    private GetOntapStorageVirtualMachineFilter(GetOntapStorageVirtualMachineFilter $) {
        this.name = $.name;
        this.values = $.values;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetOntapStorageVirtualMachineFilter defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetOntapStorageVirtualMachineFilter $;

        public Builder() {
            $ = new GetOntapStorageVirtualMachineFilter();
        }

        public Builder(GetOntapStorageVirtualMachineFilter defaults) {
            $ = new GetOntapStorageVirtualMachineFilter(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Name of the field to filter by, as defined by [the underlying AWS API](https://docs.aws.amazon.com/fsx/latest/APIReference/API_StorageVirtualMachineFilter.html).
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        /**
         * @param values Set of values that are accepted for the given field. An SVM will be selected if any one of the given values matches.
         * 
         * @return builder
         * 
         */
        public Builder values(List<String> values) {
            $.values = values;
            return this;
        }

        /**
         * @param values Set of values that are accepted for the given field. An SVM will be selected if any one of the given values matches.
         * 
         * @return builder
         * 
         */
        public Builder values(String... values) {
            return values(List.of(values));
        }

        public GetOntapStorageVirtualMachineFilter build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.values = Objects.requireNonNull($.values, "expected parameter 'values' to be non-null");
            return $;
        }
    }

}
