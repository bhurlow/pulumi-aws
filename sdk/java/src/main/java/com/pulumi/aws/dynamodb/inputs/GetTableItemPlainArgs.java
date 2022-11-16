// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.dynamodb.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetTableItemPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetTableItemPlainArgs Empty = new GetTableItemPlainArgs();

    @Import(name="expressionAttributeNames")
    private @Nullable Map<String,String> expressionAttributeNames;

    public Optional<Map<String,String>> expressionAttributeNames() {
        return Optional.ofNullable(this.expressionAttributeNames);
    }

    /**
     * A map of attribute names to AttributeValue objects, representing the primary key of the item to retrieve.
     * For the primary key, you must provide all of the attributes. For example, with a simple primary key, you only need to provide a value for the partition key. For a composite primary key, you must provide values for both the partition key and the sort key.
     * 
     */
    @Import(name="key", required=true)
    private String key;

    /**
     * @return A map of attribute names to AttributeValue objects, representing the primary key of the item to retrieve.
     * For the primary key, you must provide all of the attributes. For example, with a simple primary key, you only need to provide a value for the partition key. For a composite primary key, you must provide values for both the partition key and the sort key.
     * 
     */
    public String key() {
        return this.key;
    }

    /**
     * A string that identifies one or more attributes to retrieve from the table. These attributes can include scalars, sets, or elements of a JSON document. The attributes in the expression must be separated by commas.
     * If no attribute names are specified, then all attributes are returned. If any of the requested attributes are not found, they do not appear in the result.
     * 
     */
    @Import(name="projectionExpression")
    private @Nullable String projectionExpression;

    /**
     * @return A string that identifies one or more attributes to retrieve from the table. These attributes can include scalars, sets, or elements of a JSON document. The attributes in the expression must be separated by commas.
     * If no attribute names are specified, then all attributes are returned. If any of the requested attributes are not found, they do not appear in the result.
     * 
     */
    public Optional<String> projectionExpression() {
        return Optional.ofNullable(this.projectionExpression);
    }

    /**
     * The name of the table containing the requested item.
     * 
     */
    @Import(name="tableName", required=true)
    private String tableName;

    /**
     * @return The name of the table containing the requested item.
     * 
     */
    public String tableName() {
        return this.tableName;
    }

    private GetTableItemPlainArgs() {}

    private GetTableItemPlainArgs(GetTableItemPlainArgs $) {
        this.expressionAttributeNames = $.expressionAttributeNames;
        this.key = $.key;
        this.projectionExpression = $.projectionExpression;
        this.tableName = $.tableName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetTableItemPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetTableItemPlainArgs $;

        public Builder() {
            $ = new GetTableItemPlainArgs();
        }

        public Builder(GetTableItemPlainArgs defaults) {
            $ = new GetTableItemPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder expressionAttributeNames(@Nullable Map<String,String> expressionAttributeNames) {
            $.expressionAttributeNames = expressionAttributeNames;
            return this;
        }

        /**
         * @param key A map of attribute names to AttributeValue objects, representing the primary key of the item to retrieve.
         * For the primary key, you must provide all of the attributes. For example, with a simple primary key, you only need to provide a value for the partition key. For a composite primary key, you must provide values for both the partition key and the sort key.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            $.key = key;
            return this;
        }

        /**
         * @param projectionExpression A string that identifies one or more attributes to retrieve from the table. These attributes can include scalars, sets, or elements of a JSON document. The attributes in the expression must be separated by commas.
         * If no attribute names are specified, then all attributes are returned. If any of the requested attributes are not found, they do not appear in the result.
         * 
         * @return builder
         * 
         */
        public Builder projectionExpression(@Nullable String projectionExpression) {
            $.projectionExpression = projectionExpression;
            return this;
        }

        /**
         * @param tableName The name of the table containing the requested item.
         * 
         * @return builder
         * 
         */
        public Builder tableName(String tableName) {
            $.tableName = tableName;
            return this;
        }

        public GetTableItemPlainArgs build() {
            $.key = Objects.requireNonNull($.key, "expected parameter 'key' to be non-null");
            $.tableName = Objects.requireNonNull($.tableName, "expected parameter 'tableName' to be non-null");
            return $;
        }
    }

}
