// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.keyspaces.inputs;

import com.pulumi.aws.keyspaces.inputs.TableCapacitySpecificationArgs;
import com.pulumi.aws.keyspaces.inputs.TableClientSideTimestampsArgs;
import com.pulumi.aws.keyspaces.inputs.TableCommentArgs;
import com.pulumi.aws.keyspaces.inputs.TableEncryptionSpecificationArgs;
import com.pulumi.aws.keyspaces.inputs.TablePointInTimeRecoveryArgs;
import com.pulumi.aws.keyspaces.inputs.TableSchemaDefinitionArgs;
import com.pulumi.aws.keyspaces.inputs.TableTtlArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TableState extends com.pulumi.resources.ResourceArgs {

    public static final TableState Empty = new TableState();

    /**
     * The ARN of the table.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The ARN of the table.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Specifies the read/write throughput capacity mode for the table.
     * 
     */
    @Import(name="capacitySpecification")
    private @Nullable Output<TableCapacitySpecificationArgs> capacitySpecification;

    /**
     * @return Specifies the read/write throughput capacity mode for the table.
     * 
     */
    public Optional<Output<TableCapacitySpecificationArgs>> capacitySpecification() {
        return Optional.ofNullable(this.capacitySpecification);
    }

    /**
     * Enables client-side timestamps for the table. By default, the setting is disabled.
     * 
     */
    @Import(name="clientSideTimestamps")
    private @Nullable Output<TableClientSideTimestampsArgs> clientSideTimestamps;

    /**
     * @return Enables client-side timestamps for the table. By default, the setting is disabled.
     * 
     */
    public Optional<Output<TableClientSideTimestampsArgs>> clientSideTimestamps() {
        return Optional.ofNullable(this.clientSideTimestamps);
    }

    /**
     * A description of the table.
     * 
     */
    @Import(name="comment")
    private @Nullable Output<TableCommentArgs> comment;

    /**
     * @return A description of the table.
     * 
     */
    public Optional<Output<TableCommentArgs>> comment() {
        return Optional.ofNullable(this.comment);
    }

    /**
     * The default Time to Live setting in seconds for the table. More information can be found in the [Developer Guide](https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL-how-it-works.html#ttl-howitworks_default_ttl).
     * 
     */
    @Import(name="defaultTimeToLive")
    private @Nullable Output<Integer> defaultTimeToLive;

    /**
     * @return The default Time to Live setting in seconds for the table. More information can be found in the [Developer Guide](https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL-how-it-works.html#ttl-howitworks_default_ttl).
     * 
     */
    public Optional<Output<Integer>> defaultTimeToLive() {
        return Optional.ofNullable(this.defaultTimeToLive);
    }

    /**
     * Specifies how the encryption key for encryption at rest is managed for the table. More information can be found in the [Developer Guide](https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html).
     * 
     */
    @Import(name="encryptionSpecification")
    private @Nullable Output<TableEncryptionSpecificationArgs> encryptionSpecification;

    /**
     * @return Specifies how the encryption key for encryption at rest is managed for the table. More information can be found in the [Developer Guide](https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html).
     * 
     */
    public Optional<Output<TableEncryptionSpecificationArgs>> encryptionSpecification() {
        return Optional.ofNullable(this.encryptionSpecification);
    }

    /**
     * The name of the keyspace that the table is going to be created in.
     * 
     */
    @Import(name="keyspaceName")
    private @Nullable Output<String> keyspaceName;

    /**
     * @return The name of the keyspace that the table is going to be created in.
     * 
     */
    public Optional<Output<String>> keyspaceName() {
        return Optional.ofNullable(this.keyspaceName);
    }

    /**
     * Specifies if point-in-time recovery is enabled or disabled for the table. More information can be found in the [Developer Guide](https://docs.aws.amazon.com/keyspaces/latest/devguide/PointInTimeRecovery.html).
     * 
     */
    @Import(name="pointInTimeRecovery")
    private @Nullable Output<TablePointInTimeRecoveryArgs> pointInTimeRecovery;

    /**
     * @return Specifies if point-in-time recovery is enabled or disabled for the table. More information can be found in the [Developer Guide](https://docs.aws.amazon.com/keyspaces/latest/devguide/PointInTimeRecovery.html).
     * 
     */
    public Optional<Output<TablePointInTimeRecoveryArgs>> pointInTimeRecovery() {
        return Optional.ofNullable(this.pointInTimeRecovery);
    }

    /**
     * Describes the schema of the table.
     * 
     */
    @Import(name="schemaDefinition")
    private @Nullable Output<TableSchemaDefinitionArgs> schemaDefinition;

    /**
     * @return Describes the schema of the table.
     * 
     */
    public Optional<Output<TableSchemaDefinitionArgs>> schemaDefinition() {
        return Optional.ofNullable(this.schemaDefinition);
    }

    /**
     * The name of the table.
     * 
     * The following arguments are optional:
     * 
     */
    @Import(name="tableName")
    private @Nullable Output<String> tableName;

    /**
     * @return The name of the table.
     * 
     * The following arguments are optional:
     * 
     */
    public Optional<Output<String>> tableName() {
        return Optional.ofNullable(this.tableName);
    }

    /**
     * A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    /**
     * Enables Time to Live custom settings for the table. More information can be found in the [Developer Guide](https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL.html).
     * 
     */
    @Import(name="ttl")
    private @Nullable Output<TableTtlArgs> ttl;

    /**
     * @return Enables Time to Live custom settings for the table. More information can be found in the [Developer Guide](https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL.html).
     * 
     */
    public Optional<Output<TableTtlArgs>> ttl() {
        return Optional.ofNullable(this.ttl);
    }

    private TableState() {}

    private TableState(TableState $) {
        this.arn = $.arn;
        this.capacitySpecification = $.capacitySpecification;
        this.clientSideTimestamps = $.clientSideTimestamps;
        this.comment = $.comment;
        this.defaultTimeToLive = $.defaultTimeToLive;
        this.encryptionSpecification = $.encryptionSpecification;
        this.keyspaceName = $.keyspaceName;
        this.pointInTimeRecovery = $.pointInTimeRecovery;
        this.schemaDefinition = $.schemaDefinition;
        this.tableName = $.tableName;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.ttl = $.ttl;
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
         * @param arn The ARN of the table.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The ARN of the table.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param capacitySpecification Specifies the read/write throughput capacity mode for the table.
         * 
         * @return builder
         * 
         */
        public Builder capacitySpecification(@Nullable Output<TableCapacitySpecificationArgs> capacitySpecification) {
            $.capacitySpecification = capacitySpecification;
            return this;
        }

        /**
         * @param capacitySpecification Specifies the read/write throughput capacity mode for the table.
         * 
         * @return builder
         * 
         */
        public Builder capacitySpecification(TableCapacitySpecificationArgs capacitySpecification) {
            return capacitySpecification(Output.of(capacitySpecification));
        }

        /**
         * @param clientSideTimestamps Enables client-side timestamps for the table. By default, the setting is disabled.
         * 
         * @return builder
         * 
         */
        public Builder clientSideTimestamps(@Nullable Output<TableClientSideTimestampsArgs> clientSideTimestamps) {
            $.clientSideTimestamps = clientSideTimestamps;
            return this;
        }

        /**
         * @param clientSideTimestamps Enables client-side timestamps for the table. By default, the setting is disabled.
         * 
         * @return builder
         * 
         */
        public Builder clientSideTimestamps(TableClientSideTimestampsArgs clientSideTimestamps) {
            return clientSideTimestamps(Output.of(clientSideTimestamps));
        }

        /**
         * @param comment A description of the table.
         * 
         * @return builder
         * 
         */
        public Builder comment(@Nullable Output<TableCommentArgs> comment) {
            $.comment = comment;
            return this;
        }

        /**
         * @param comment A description of the table.
         * 
         * @return builder
         * 
         */
        public Builder comment(TableCommentArgs comment) {
            return comment(Output.of(comment));
        }

        /**
         * @param defaultTimeToLive The default Time to Live setting in seconds for the table. More information can be found in the [Developer Guide](https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL-how-it-works.html#ttl-howitworks_default_ttl).
         * 
         * @return builder
         * 
         */
        public Builder defaultTimeToLive(@Nullable Output<Integer> defaultTimeToLive) {
            $.defaultTimeToLive = defaultTimeToLive;
            return this;
        }

        /**
         * @param defaultTimeToLive The default Time to Live setting in seconds for the table. More information can be found in the [Developer Guide](https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL-how-it-works.html#ttl-howitworks_default_ttl).
         * 
         * @return builder
         * 
         */
        public Builder defaultTimeToLive(Integer defaultTimeToLive) {
            return defaultTimeToLive(Output.of(defaultTimeToLive));
        }

        /**
         * @param encryptionSpecification Specifies how the encryption key for encryption at rest is managed for the table. More information can be found in the [Developer Guide](https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html).
         * 
         * @return builder
         * 
         */
        public Builder encryptionSpecification(@Nullable Output<TableEncryptionSpecificationArgs> encryptionSpecification) {
            $.encryptionSpecification = encryptionSpecification;
            return this;
        }

        /**
         * @param encryptionSpecification Specifies how the encryption key for encryption at rest is managed for the table. More information can be found in the [Developer Guide](https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html).
         * 
         * @return builder
         * 
         */
        public Builder encryptionSpecification(TableEncryptionSpecificationArgs encryptionSpecification) {
            return encryptionSpecification(Output.of(encryptionSpecification));
        }

        /**
         * @param keyspaceName The name of the keyspace that the table is going to be created in.
         * 
         * @return builder
         * 
         */
        public Builder keyspaceName(@Nullable Output<String> keyspaceName) {
            $.keyspaceName = keyspaceName;
            return this;
        }

        /**
         * @param keyspaceName The name of the keyspace that the table is going to be created in.
         * 
         * @return builder
         * 
         */
        public Builder keyspaceName(String keyspaceName) {
            return keyspaceName(Output.of(keyspaceName));
        }

        /**
         * @param pointInTimeRecovery Specifies if point-in-time recovery is enabled or disabled for the table. More information can be found in the [Developer Guide](https://docs.aws.amazon.com/keyspaces/latest/devguide/PointInTimeRecovery.html).
         * 
         * @return builder
         * 
         */
        public Builder pointInTimeRecovery(@Nullable Output<TablePointInTimeRecoveryArgs> pointInTimeRecovery) {
            $.pointInTimeRecovery = pointInTimeRecovery;
            return this;
        }

        /**
         * @param pointInTimeRecovery Specifies if point-in-time recovery is enabled or disabled for the table. More information can be found in the [Developer Guide](https://docs.aws.amazon.com/keyspaces/latest/devguide/PointInTimeRecovery.html).
         * 
         * @return builder
         * 
         */
        public Builder pointInTimeRecovery(TablePointInTimeRecoveryArgs pointInTimeRecovery) {
            return pointInTimeRecovery(Output.of(pointInTimeRecovery));
        }

        /**
         * @param schemaDefinition Describes the schema of the table.
         * 
         * @return builder
         * 
         */
        public Builder schemaDefinition(@Nullable Output<TableSchemaDefinitionArgs> schemaDefinition) {
            $.schemaDefinition = schemaDefinition;
            return this;
        }

        /**
         * @param schemaDefinition Describes the schema of the table.
         * 
         * @return builder
         * 
         */
        public Builder schemaDefinition(TableSchemaDefinitionArgs schemaDefinition) {
            return schemaDefinition(Output.of(schemaDefinition));
        }

        /**
         * @param tableName The name of the table.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder tableName(@Nullable Output<String> tableName) {
            $.tableName = tableName;
            return this;
        }

        /**
         * @param tableName The name of the table.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder tableName(String tableName) {
            return tableName(Output.of(tableName));
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
         */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        /**
         * @param ttl Enables Time to Live custom settings for the table. More information can be found in the [Developer Guide](https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL.html).
         * 
         * @return builder
         * 
         */
        public Builder ttl(@Nullable Output<TableTtlArgs> ttl) {
            $.ttl = ttl;
            return this;
        }

        /**
         * @param ttl Enables Time to Live custom settings for the table. More information can be found in the [Developer Guide](https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL.html).
         * 
         * @return builder
         * 
         */
        public Builder ttl(TableTtlArgs ttl) {
            return ttl(Output.of(ttl));
        }

        public TableState build() {
            return $;
        }
    }

}
