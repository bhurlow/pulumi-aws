// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.lakeformation.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PermissionsTableArgs extends com.pulumi.resources.ResourceArgs {

    public static final PermissionsTableArgs Empty = new PermissionsTableArgs();

    /**
     * Identifier for the Data Catalog. By default, it is the account ID of the caller.
     * 
     */
    @Import(name="catalogId")
    private @Nullable Output<String> catalogId;

    /**
     * @return Identifier for the Data Catalog. By default, it is the account ID of the caller.
     * 
     */
    public Optional<Output<String>> catalogId() {
        return Optional.ofNullable(this.catalogId);
    }

    /**
     * Name of the database for the table with columns resource. Unique to the Data Catalog.
     * 
     */
    @Import(name="databaseName", required=true)
    private Output<String> databaseName;

    /**
     * @return Name of the database for the table with columns resource. Unique to the Data Catalog.
     * 
     */
    public Output<String> databaseName() {
        return this.databaseName;
    }

    /**
     * Name of the table resource.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the table resource.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Whether to use a column wildcard. If `excluded_column_names` is included, `wildcard` must be set to `true` to avoid the provider reporting a difference.
     * 
     */
    @Import(name="wildcard")
    private @Nullable Output<Boolean> wildcard;

    /**
     * @return Whether to use a column wildcard. If `excluded_column_names` is included, `wildcard` must be set to `true` to avoid the provider reporting a difference.
     * 
     */
    public Optional<Output<Boolean>> wildcard() {
        return Optional.ofNullable(this.wildcard);
    }

    private PermissionsTableArgs() {}

    private PermissionsTableArgs(PermissionsTableArgs $) {
        this.catalogId = $.catalogId;
        this.databaseName = $.databaseName;
        this.name = $.name;
        this.wildcard = $.wildcard;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PermissionsTableArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PermissionsTableArgs $;

        public Builder() {
            $ = new PermissionsTableArgs();
        }

        public Builder(PermissionsTableArgs defaults) {
            $ = new PermissionsTableArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param catalogId Identifier for the Data Catalog. By default, it is the account ID of the caller.
         * 
         * @return builder
         * 
         */
        public Builder catalogId(@Nullable Output<String> catalogId) {
            $.catalogId = catalogId;
            return this;
        }

        /**
         * @param catalogId Identifier for the Data Catalog. By default, it is the account ID of the caller.
         * 
         * @return builder
         * 
         */
        public Builder catalogId(String catalogId) {
            return catalogId(Output.of(catalogId));
        }

        /**
         * @param databaseName Name of the database for the table with columns resource. Unique to the Data Catalog.
         * 
         * @return builder
         * 
         */
        public Builder databaseName(Output<String> databaseName) {
            $.databaseName = databaseName;
            return this;
        }

        /**
         * @param databaseName Name of the database for the table with columns resource. Unique to the Data Catalog.
         * 
         * @return builder
         * 
         */
        public Builder databaseName(String databaseName) {
            return databaseName(Output.of(databaseName));
        }

        /**
         * @param name Name of the table resource.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the table resource.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param wildcard Whether to use a column wildcard. If `excluded_column_names` is included, `wildcard` must be set to `true` to avoid the provider reporting a difference.
         * 
         * @return builder
         * 
         */
        public Builder wildcard(@Nullable Output<Boolean> wildcard) {
            $.wildcard = wildcard;
            return this;
        }

        /**
         * @param wildcard Whether to use a column wildcard. If `excluded_column_names` is included, `wildcard` must be set to `true` to avoid the provider reporting a difference.
         * 
         * @return builder
         * 
         */
        public Builder wildcard(Boolean wildcard) {
            return wildcard(Output.of(wildcard));
        }

        public PermissionsTableArgs build() {
            $.databaseName = Objects.requireNonNull($.databaseName, "expected parameter 'databaseName' to be non-null");
            return $;
        }
    }

}
