// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.glue.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SchemaState extends com.pulumi.resources.ResourceArgs {

    public static final SchemaState Empty = new SchemaState();

    /**
     * Amazon Resource Name (ARN) of the schema.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) of the schema.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The compatibility mode of the schema. Values values are: `NONE`, `DISABLED`, `BACKWARD`, `BACKWARD_ALL`, `FORWARD`, `FORWARD_ALL`, `FULL`, and `FULL_ALL`.
     * 
     */
    @Import(name="compatibility")
    private @Nullable Output<String> compatibility;

    /**
     * @return The compatibility mode of the schema. Values values are: `NONE`, `DISABLED`, `BACKWARD`, `BACKWARD_ALL`, `FORWARD`, `FORWARD_ALL`, `FULL`, and `FULL_ALL`.
     * 
     */
    public Optional<Output<String>> compatibility() {
        return Optional.ofNullable(this.compatibility);
    }

    /**
     * The data format of the schema definition. Valid values are `AVRO`, `JSON` and `PROTOBUF`.
     * 
     */
    @Import(name="dataFormat")
    private @Nullable Output<String> dataFormat;

    /**
     * @return The data format of the schema definition. Valid values are `AVRO`, `JSON` and `PROTOBUF`.
     * 
     */
    public Optional<Output<String>> dataFormat() {
        return Optional.ofNullable(this.dataFormat);
    }

    /**
     * A description of the schema.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description of the schema.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The latest version of the schema associated with the returned schema definition.
     * 
     */
    @Import(name="latestSchemaVersion")
    private @Nullable Output<Integer> latestSchemaVersion;

    /**
     * @return The latest version of the schema associated with the returned schema definition.
     * 
     */
    public Optional<Output<Integer>> latestSchemaVersion() {
        return Optional.ofNullable(this.latestSchemaVersion);
    }

    /**
     * The next version of the schema associated with the returned schema definition.
     * 
     */
    @Import(name="nextSchemaVersion")
    private @Nullable Output<Integer> nextSchemaVersion;

    /**
     * @return The next version of the schema associated with the returned schema definition.
     * 
     */
    public Optional<Output<Integer>> nextSchemaVersion() {
        return Optional.ofNullable(this.nextSchemaVersion);
    }

    /**
     * The ARN of the Glue Registry to create the schema in.
     * 
     */
    @Import(name="registryArn")
    private @Nullable Output<String> registryArn;

    /**
     * @return The ARN of the Glue Registry to create the schema in.
     * 
     */
    public Optional<Output<String>> registryArn() {
        return Optional.ofNullable(this.registryArn);
    }

    /**
     * The name of the Glue Registry.
     * 
     */
    @Import(name="registryName")
    private @Nullable Output<String> registryName;

    /**
     * @return The name of the Glue Registry.
     * 
     */
    public Optional<Output<String>> registryName() {
        return Optional.ofNullable(this.registryName);
    }

    /**
     * The version number of the checkpoint (the last time the compatibility mode was changed).
     * 
     */
    @Import(name="schemaCheckpoint")
    private @Nullable Output<Integer> schemaCheckpoint;

    /**
     * @return The version number of the checkpoint (the last time the compatibility mode was changed).
     * 
     */
    public Optional<Output<Integer>> schemaCheckpoint() {
        return Optional.ofNullable(this.schemaCheckpoint);
    }

    /**
     * The schema definition using the `data_format` setting for `schema_name`.
     * 
     */
    @Import(name="schemaDefinition")
    private @Nullable Output<String> schemaDefinition;

    /**
     * @return The schema definition using the `data_format` setting for `schema_name`.
     * 
     */
    public Optional<Output<String>> schemaDefinition() {
        return Optional.ofNullable(this.schemaDefinition);
    }

    /**
     * The Name of the schema.
     * 
     */
    @Import(name="schemaName")
    private @Nullable Output<String> schemaName;

    /**
     * @return The Name of the schema.
     * 
     */
    public Optional<Output<String>> schemaName() {
        return Optional.ofNullable(this.schemaName);
    }

    /**
     * Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

    private SchemaState() {}

    private SchemaState(SchemaState $) {
        this.arn = $.arn;
        this.compatibility = $.compatibility;
        this.dataFormat = $.dataFormat;
        this.description = $.description;
        this.latestSchemaVersion = $.latestSchemaVersion;
        this.nextSchemaVersion = $.nextSchemaVersion;
        this.registryArn = $.registryArn;
        this.registryName = $.registryName;
        this.schemaCheckpoint = $.schemaCheckpoint;
        this.schemaDefinition = $.schemaDefinition;
        this.schemaName = $.schemaName;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SchemaState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SchemaState $;

        public Builder() {
            $ = new SchemaState();
        }

        public Builder(SchemaState defaults) {
            $ = new SchemaState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn Amazon Resource Name (ARN) of the schema.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn Amazon Resource Name (ARN) of the schema.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param compatibility The compatibility mode of the schema. Values values are: `NONE`, `DISABLED`, `BACKWARD`, `BACKWARD_ALL`, `FORWARD`, `FORWARD_ALL`, `FULL`, and `FULL_ALL`.
         * 
         * @return builder
         * 
         */
        public Builder compatibility(@Nullable Output<String> compatibility) {
            $.compatibility = compatibility;
            return this;
        }

        /**
         * @param compatibility The compatibility mode of the schema. Values values are: `NONE`, `DISABLED`, `BACKWARD`, `BACKWARD_ALL`, `FORWARD`, `FORWARD_ALL`, `FULL`, and `FULL_ALL`.
         * 
         * @return builder
         * 
         */
        public Builder compatibility(String compatibility) {
            return compatibility(Output.of(compatibility));
        }

        /**
         * @param dataFormat The data format of the schema definition. Valid values are `AVRO`, `JSON` and `PROTOBUF`.
         * 
         * @return builder
         * 
         */
        public Builder dataFormat(@Nullable Output<String> dataFormat) {
            $.dataFormat = dataFormat;
            return this;
        }

        /**
         * @param dataFormat The data format of the schema definition. Valid values are `AVRO`, `JSON` and `PROTOBUF`.
         * 
         * @return builder
         * 
         */
        public Builder dataFormat(String dataFormat) {
            return dataFormat(Output.of(dataFormat));
        }

        /**
         * @param description A description of the schema.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description of the schema.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param latestSchemaVersion The latest version of the schema associated with the returned schema definition.
         * 
         * @return builder
         * 
         */
        public Builder latestSchemaVersion(@Nullable Output<Integer> latestSchemaVersion) {
            $.latestSchemaVersion = latestSchemaVersion;
            return this;
        }

        /**
         * @param latestSchemaVersion The latest version of the schema associated with the returned schema definition.
         * 
         * @return builder
         * 
         */
        public Builder latestSchemaVersion(Integer latestSchemaVersion) {
            return latestSchemaVersion(Output.of(latestSchemaVersion));
        }

        /**
         * @param nextSchemaVersion The next version of the schema associated with the returned schema definition.
         * 
         * @return builder
         * 
         */
        public Builder nextSchemaVersion(@Nullable Output<Integer> nextSchemaVersion) {
            $.nextSchemaVersion = nextSchemaVersion;
            return this;
        }

        /**
         * @param nextSchemaVersion The next version of the schema associated with the returned schema definition.
         * 
         * @return builder
         * 
         */
        public Builder nextSchemaVersion(Integer nextSchemaVersion) {
            return nextSchemaVersion(Output.of(nextSchemaVersion));
        }

        /**
         * @param registryArn The ARN of the Glue Registry to create the schema in.
         * 
         * @return builder
         * 
         */
        public Builder registryArn(@Nullable Output<String> registryArn) {
            $.registryArn = registryArn;
            return this;
        }

        /**
         * @param registryArn The ARN of the Glue Registry to create the schema in.
         * 
         * @return builder
         * 
         */
        public Builder registryArn(String registryArn) {
            return registryArn(Output.of(registryArn));
        }

        /**
         * @param registryName The name of the Glue Registry.
         * 
         * @return builder
         * 
         */
        public Builder registryName(@Nullable Output<String> registryName) {
            $.registryName = registryName;
            return this;
        }

        /**
         * @param registryName The name of the Glue Registry.
         * 
         * @return builder
         * 
         */
        public Builder registryName(String registryName) {
            return registryName(Output.of(registryName));
        }

        /**
         * @param schemaCheckpoint The version number of the checkpoint (the last time the compatibility mode was changed).
         * 
         * @return builder
         * 
         */
        public Builder schemaCheckpoint(@Nullable Output<Integer> schemaCheckpoint) {
            $.schemaCheckpoint = schemaCheckpoint;
            return this;
        }

        /**
         * @param schemaCheckpoint The version number of the checkpoint (the last time the compatibility mode was changed).
         * 
         * @return builder
         * 
         */
        public Builder schemaCheckpoint(Integer schemaCheckpoint) {
            return schemaCheckpoint(Output.of(schemaCheckpoint));
        }

        /**
         * @param schemaDefinition The schema definition using the `data_format` setting for `schema_name`.
         * 
         * @return builder
         * 
         */
        public Builder schemaDefinition(@Nullable Output<String> schemaDefinition) {
            $.schemaDefinition = schemaDefinition;
            return this;
        }

        /**
         * @param schemaDefinition The schema definition using the `data_format` setting for `schema_name`.
         * 
         * @return builder
         * 
         */
        public Builder schemaDefinition(String schemaDefinition) {
            return schemaDefinition(Output.of(schemaDefinition));
        }

        /**
         * @param schemaName The Name of the schema.
         * 
         * @return builder
         * 
         */
        public Builder schemaName(@Nullable Output<String> schemaName) {
            $.schemaName = schemaName;
            return this;
        }

        /**
         * @param schemaName The Name of the schema.
         * 
         * @return builder
         * 
         */
        public Builder schemaName(String schemaName) {
            return schemaName(Output.of(schemaName));
        }

        /**
         * @param tags Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public SchemaState build() {
            return $;
        }
    }

}
