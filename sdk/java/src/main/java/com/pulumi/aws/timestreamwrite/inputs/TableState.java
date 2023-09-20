// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.timestreamwrite.inputs;

import com.pulumi.aws.timestreamwrite.inputs.TableMagneticStoreWritePropertiesArgs;
import com.pulumi.aws.timestreamwrite.inputs.TableRetentionPropertiesArgs;
import com.pulumi.aws.timestreamwrite.inputs.TableSchemaArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TableState extends com.pulumi.resources.ResourceArgs {

    public static final TableState Empty = new TableState();

    /**
     * The ARN that uniquely identifies this table.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The ARN that uniquely identifies this table.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The name of the Timestream database.
     * 
     */
    @Import(name="databaseName")
    private @Nullable Output<String> databaseName;

    /**
     * @return The name of the Timestream database.
     * 
     */
    public Optional<Output<String>> databaseName() {
        return Optional.ofNullable(this.databaseName);
    }

    /**
     * Contains properties to set on the table when enabling magnetic store writes. See Magnetic Store Write Properties below for more details.
     * 
     */
    @Import(name="magneticStoreWriteProperties")
    private @Nullable Output<TableMagneticStoreWritePropertiesArgs> magneticStoreWriteProperties;

    /**
     * @return Contains properties to set on the table when enabling magnetic store writes. See Magnetic Store Write Properties below for more details.
     * 
     */
    public Optional<Output<TableMagneticStoreWritePropertiesArgs>> magneticStoreWriteProperties() {
        return Optional.ofNullable(this.magneticStoreWriteProperties);
    }

    /**
     * The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, `magnetic_store_retention_period_in_days` default to 73000 and `memory_store_retention_period_in_hours` defaults to 6.
     * 
     */
    @Import(name="retentionProperties")
    private @Nullable Output<TableRetentionPropertiesArgs> retentionProperties;

    /**
     * @return The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, `magnetic_store_retention_period_in_days` default to 73000 and `memory_store_retention_period_in_hours` defaults to 6.
     * 
     */
    public Optional<Output<TableRetentionPropertiesArgs>> retentionProperties() {
        return Optional.ofNullable(this.retentionProperties);
    }

    /**
     * The schema of the table. See Schema below for more details.
     * 
     */
    @Import(name="schema")
    private @Nullable Output<TableSchemaArgs> schema;

    /**
     * @return The schema of the table. See Schema below for more details.
     * 
     */
    public Optional<Output<TableSchemaArgs>> schema() {
        return Optional.ofNullable(this.schema);
    }

    /**
     * The name of the Timestream table.
     * 
     */
    @Import(name="tableName")
    private @Nullable Output<String> tableName;

    /**
     * @return The name of the Timestream table.
     * 
     */
    public Optional<Output<String>> tableName() {
        return Optional.ofNullable(this.tableName);
    }

    /**
     * Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

    private TableState() {}

    private TableState(TableState $) {
        this.arn = $.arn;
        this.databaseName = $.databaseName;
        this.magneticStoreWriteProperties = $.magneticStoreWriteProperties;
        this.retentionProperties = $.retentionProperties;
        this.schema = $.schema;
        this.tableName = $.tableName;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TableState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TableState $;

        public Builder() {
            $ = new TableState();
        }

        public Builder(TableState defaults) {
            $ = new TableState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The ARN that uniquely identifies this table.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The ARN that uniquely identifies this table.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param databaseName The name of the Timestream database.
         * 
         * @return builder
         * 
         */
        public Builder databaseName(@Nullable Output<String> databaseName) {
            $.databaseName = databaseName;
            return this;
        }

        /**
         * @param databaseName The name of the Timestream database.
         * 
         * @return builder
         * 
         */
        public Builder databaseName(String databaseName) {
            return databaseName(Output.of(databaseName));
        }

        /**
         * @param magneticStoreWriteProperties Contains properties to set on the table when enabling magnetic store writes. See Magnetic Store Write Properties below for more details.
         * 
         * @return builder
         * 
         */
        public Builder magneticStoreWriteProperties(@Nullable Output<TableMagneticStoreWritePropertiesArgs> magneticStoreWriteProperties) {
            $.magneticStoreWriteProperties = magneticStoreWriteProperties;
            return this;
        }

        /**
         * @param magneticStoreWriteProperties Contains properties to set on the table when enabling magnetic store writes. See Magnetic Store Write Properties below for more details.
         * 
         * @return builder
         * 
         */
        public Builder magneticStoreWriteProperties(TableMagneticStoreWritePropertiesArgs magneticStoreWriteProperties) {
            return magneticStoreWriteProperties(Output.of(magneticStoreWriteProperties));
        }

        /**
         * @param retentionProperties The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, `magnetic_store_retention_period_in_days` default to 73000 and `memory_store_retention_period_in_hours` defaults to 6.
         * 
         * @return builder
         * 
         */
        public Builder retentionProperties(@Nullable Output<TableRetentionPropertiesArgs> retentionProperties) {
            $.retentionProperties = retentionProperties;
            return this;
        }

        /**
         * @param retentionProperties The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, `magnetic_store_retention_period_in_days` default to 73000 and `memory_store_retention_period_in_hours` defaults to 6.
         * 
         * @return builder
         * 
         */
        public Builder retentionProperties(TableRetentionPropertiesArgs retentionProperties) {
            return retentionProperties(Output.of(retentionProperties));
        }

        /**
         * @param schema The schema of the table. See Schema below for more details.
         * 
         * @return builder
         * 
         */
        public Builder schema(@Nullable Output<TableSchemaArgs> schema) {
            $.schema = schema;
            return this;
        }

        /**
         * @param schema The schema of the table. See Schema below for more details.
         * 
         * @return builder
         * 
         */
        public Builder schema(TableSchemaArgs schema) {
            return schema(Output.of(schema));
        }

        /**
         * @param tableName The name of the Timestream table.
         * 
         * @return builder
         * 
         */
        public Builder tableName(@Nullable Output<String> tableName) {
            $.tableName = tableName;
            return this;
        }

        /**
         * @param tableName The name of the Timestream table.
         * 
         * @return builder
         * 
         */
        public Builder tableName(String tableName) {
            return tableName(Output.of(tableName));
        }

        /**
         * @param tags Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Map of tags to assign to this resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public TableState build() {
            return $;
        }
    }

}
