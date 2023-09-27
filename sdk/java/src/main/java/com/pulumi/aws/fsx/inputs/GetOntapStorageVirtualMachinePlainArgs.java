// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.fsx.inputs;

import com.pulumi.aws.fsx.inputs.GetOntapStorageVirtualMachineFilter;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetOntapStorageVirtualMachinePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetOntapStorageVirtualMachinePlainArgs Empty = new GetOntapStorageVirtualMachinePlainArgs();

    /**
     * Configuration block. Detailed below.
     * 
     */
    @Import(name="filters")
    private @Nullable List<GetOntapStorageVirtualMachineFilter> filters;

    /**
     * @return Configuration block. Detailed below.
     * 
     */
    public Optional<List<GetOntapStorageVirtualMachineFilter>> filters() {
        return Optional.ofNullable(this.filters);
    }

    /**
     * Identifier of the storage virtual machine (e.g. `svm-12345678`).
     * 
     */
    @Import(name="id")
    private @Nullable String id;

    /**
     * @return Identifier of the storage virtual machine (e.g. `svm-12345678`).
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }

    @Import(name="tags")
    private @Nullable Map<String,String> tags;

    public Optional<Map<String,String>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private GetOntapStorageVirtualMachinePlainArgs() {}

    private GetOntapStorageVirtualMachinePlainArgs(GetOntapStorageVirtualMachinePlainArgs $) {
        this.filters = $.filters;
        this.id = $.id;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetOntapStorageVirtualMachinePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetOntapStorageVirtualMachinePlainArgs $;

        public Builder() {
            $ = new GetOntapStorageVirtualMachinePlainArgs();
        }

        public Builder(GetOntapStorageVirtualMachinePlainArgs defaults) {
            $ = new GetOntapStorageVirtualMachinePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filters Configuration block. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder filters(@Nullable List<GetOntapStorageVirtualMachineFilter> filters) {
            $.filters = filters;
            return this;
        }

        /**
         * @param filters Configuration block. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder filters(GetOntapStorageVirtualMachineFilter... filters) {
            return filters(List.of(filters));
        }

        /**
         * @param id Identifier of the storage virtual machine (e.g. `svm-12345678`).
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable String id) {
            $.id = id;
            return this;
        }

        public Builder tags(@Nullable Map<String,String> tags) {
            $.tags = tags;
            return this;
        }

        public GetOntapStorageVirtualMachinePlainArgs build() {
            return $;
        }
    }

}
