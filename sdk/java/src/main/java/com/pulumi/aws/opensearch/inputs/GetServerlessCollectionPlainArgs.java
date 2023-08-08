// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.opensearch.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetServerlessCollectionPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetServerlessCollectionPlainArgs Empty = new GetServerlessCollectionPlainArgs();

    /**
     * ID of the collection. Either `id` or `name` must be provided.
     * 
     */
    @Import(name="id")
    private @Nullable String id;

    /**
     * @return ID of the collection. Either `id` or `name` must be provided.
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }

    /**
     * Name of the collection. Either `name` or `id` must be provided.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return Name of the collection. Either `name` or `id` must be provided.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    private GetServerlessCollectionPlainArgs() {}

    private GetServerlessCollectionPlainArgs(GetServerlessCollectionPlainArgs $) {
        this.id = $.id;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetServerlessCollectionPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetServerlessCollectionPlainArgs $;

        public Builder() {
            $ = new GetServerlessCollectionPlainArgs();
        }

        public Builder(GetServerlessCollectionPlainArgs defaults) {
            $ = new GetServerlessCollectionPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param id ID of the collection. Either `id` or `name` must be provided.
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable String id) {
            $.id = id;
            return this;
        }

        /**
         * @param name Name of the collection. Either `name` or `id` must be provided.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        public GetServerlessCollectionPlainArgs build() {
            return $;
        }
    }

}
