// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.glue.inputs;

import com.pulumi.aws.glue.inputs.CatalogDatabaseCreateTableDefaultPermissionArgs;
import com.pulumi.aws.glue.inputs.CatalogDatabaseTargetDatabaseArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CatalogDatabaseState extends com.pulumi.resources.ResourceArgs {

    public static final CatalogDatabaseState Empty = new CatalogDatabaseState();

    /**
     * ARN of the Glue Catalog Database.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return ARN of the Glue Catalog Database.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * ID of the Glue Catalog to create the database in. If omitted, this defaults to the AWS Account ID.
     * 
     */
    @Import(name="catalogId")
    private @Nullable Output<String> catalogId;

    /**
     * @return ID of the Glue Catalog to create the database in. If omitted, this defaults to the AWS Account ID.
     * 
     */
    public Optional<Output<String>> catalogId() {
        return Optional.ofNullable(this.catalogId);
    }

    /**
     * Creates a set of default permissions on the table for principals. See `create_table_default_permission` below.
     * 
     */
    @Import(name="createTableDefaultPermissions")
    private @Nullable Output<List<CatalogDatabaseCreateTableDefaultPermissionArgs>> createTableDefaultPermissions;

    /**
     * @return Creates a set of default permissions on the table for principals. See `create_table_default_permission` below.
     * 
     */
    public Optional<Output<List<CatalogDatabaseCreateTableDefaultPermissionArgs>>> createTableDefaultPermissions() {
        return Optional.ofNullable(this.createTableDefaultPermissions);
    }

    /**
     * Description of the database.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the database.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Location of the database (for example, an HDFS path).
     * 
     */
    @Import(name="locationUri")
    private @Nullable Output<String> locationUri;

    /**
     * @return Location of the database (for example, an HDFS path).
     * 
     */
    public Optional<Output<String>> locationUri() {
        return Optional.ofNullable(this.locationUri);
    }

    /**
     * Name of the database. The acceptable characters are lowercase letters, numbers, and the underscore character.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the database. The acceptable characters are lowercase letters, numbers, and the underscore character.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * List of key-value pairs that define parameters and properties of the database.
     * 
     */
    @Import(name="parameters")
    private @Nullable Output<Map<String,String>> parameters;

    /**
     * @return List of key-value pairs that define parameters and properties of the database.
     * 
     */
    public Optional<Output<Map<String,String>>> parameters() {
        return Optional.ofNullable(this.parameters);
    }

    /**
     * Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    /**
     * Configuration block for a target database for resource linking. See `target_database` below.
     * 
     */
    @Import(name="targetDatabase")
    private @Nullable Output<CatalogDatabaseTargetDatabaseArgs> targetDatabase;

    /**
     * @return Configuration block for a target database for resource linking. See `target_database` below.
     * 
     */
    public Optional<Output<CatalogDatabaseTargetDatabaseArgs>> targetDatabase() {
        return Optional.ofNullable(this.targetDatabase);
    }

    private CatalogDatabaseState() {}

    private CatalogDatabaseState(CatalogDatabaseState $) {
        this.arn = $.arn;
        this.catalogId = $.catalogId;
        this.createTableDefaultPermissions = $.createTableDefaultPermissions;
        this.description = $.description;
        this.locationUri = $.locationUri;
        this.name = $.name;
        this.parameters = $.parameters;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.targetDatabase = $.targetDatabase;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CatalogDatabaseState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CatalogDatabaseState $;

        public Builder() {
            $ = new CatalogDatabaseState();
        }

        public Builder(CatalogDatabaseState defaults) {
            $ = new CatalogDatabaseState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn ARN of the Glue Catalog Database.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn ARN of the Glue Catalog Database.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param catalogId ID of the Glue Catalog to create the database in. If omitted, this defaults to the AWS Account ID.
         * 
         * @return builder
         * 
         */
        public Builder catalogId(@Nullable Output<String> catalogId) {
            $.catalogId = catalogId;
            return this;
        }

        /**
         * @param catalogId ID of the Glue Catalog to create the database in. If omitted, this defaults to the AWS Account ID.
         * 
         * @return builder
         * 
         */
        public Builder catalogId(String catalogId) {
            return catalogId(Output.of(catalogId));
        }

        /**
         * @param createTableDefaultPermissions Creates a set of default permissions on the table for principals. See `create_table_default_permission` below.
         * 
         * @return builder
         * 
         */
        public Builder createTableDefaultPermissions(@Nullable Output<List<CatalogDatabaseCreateTableDefaultPermissionArgs>> createTableDefaultPermissions) {
            $.createTableDefaultPermissions = createTableDefaultPermissions;
            return this;
        }

        /**
         * @param createTableDefaultPermissions Creates a set of default permissions on the table for principals. See `create_table_default_permission` below.
         * 
         * @return builder
         * 
         */
        public Builder createTableDefaultPermissions(List<CatalogDatabaseCreateTableDefaultPermissionArgs> createTableDefaultPermissions) {
            return createTableDefaultPermissions(Output.of(createTableDefaultPermissions));
        }

        /**
         * @param createTableDefaultPermissions Creates a set of default permissions on the table for principals. See `create_table_default_permission` below.
         * 
         * @return builder
         * 
         */
        public Builder createTableDefaultPermissions(CatalogDatabaseCreateTableDefaultPermissionArgs... createTableDefaultPermissions) {
            return createTableDefaultPermissions(List.of(createTableDefaultPermissions));
        }

        /**
         * @param description Description of the database.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the database.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param locationUri Location of the database (for example, an HDFS path).
         * 
         * @return builder
         * 
         */
        public Builder locationUri(@Nullable Output<String> locationUri) {
            $.locationUri = locationUri;
            return this;
        }

        /**
         * @param locationUri Location of the database (for example, an HDFS path).
         * 
         * @return builder
         * 
         */
        public Builder locationUri(String locationUri) {
            return locationUri(Output.of(locationUri));
        }

        /**
         * @param name Name of the database. The acceptable characters are lowercase letters, numbers, and the underscore character.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the database. The acceptable characters are lowercase letters, numbers, and the underscore character.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param parameters List of key-value pairs that define parameters and properties of the database.
         * 
         * @return builder
         * 
         */
        public Builder parameters(@Nullable Output<Map<String,String>> parameters) {
            $.parameters = parameters;
            return this;
        }

        /**
         * @param parameters List of key-value pairs that define parameters and properties of the database.
         * 
         * @return builder
         * 
         */
        public Builder parameters(Map<String,String> parameters) {
            return parameters(Output.of(parameters));
        }

        /**
         * @param tags Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        /**
         * @param targetDatabase Configuration block for a target database for resource linking. See `target_database` below.
         * 
         * @return builder
         * 
         */
        public Builder targetDatabase(@Nullable Output<CatalogDatabaseTargetDatabaseArgs> targetDatabase) {
            $.targetDatabase = targetDatabase;
            return this;
        }

        /**
         * @param targetDatabase Configuration block for a target database for resource linking. See `target_database` below.
         * 
         * @return builder
         * 
         */
        public Builder targetDatabase(CatalogDatabaseTargetDatabaseArgs targetDatabase) {
            return targetDatabase(Output.of(targetDatabase));
        }

        public CatalogDatabaseState build() {
            return $;
        }
    }

}
