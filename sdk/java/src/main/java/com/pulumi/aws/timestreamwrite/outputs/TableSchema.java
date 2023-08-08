// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.timestreamwrite.outputs;

import com.pulumi.aws.timestreamwrite.outputs.TableSchemaCompositePartitionKey;
import com.pulumi.core.annotations.CustomType;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class TableSchema {
    /**
     * @return A non-empty list of partition keys defining the attributes used to partition the table data. The order of the list determines the partition hierarchy. The name and type of each partition key as well as the partition key order cannot be changed after the table is created. However, the enforcement level of each partition key can be changed. See Composite Partition Key below for more details.
     * 
     */
    private @Nullable TableSchemaCompositePartitionKey compositePartitionKey;

    private TableSchema() {}
    /**
     * @return A non-empty list of partition keys defining the attributes used to partition the table data. The order of the list determines the partition hierarchy. The name and type of each partition key as well as the partition key order cannot be changed after the table is created. However, the enforcement level of each partition key can be changed. See Composite Partition Key below for more details.
     * 
     */
    public Optional<TableSchemaCompositePartitionKey> compositePartitionKey() {
        return Optional.ofNullable(this.compositePartitionKey);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(TableSchema defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable TableSchemaCompositePartitionKey compositePartitionKey;
        public Builder() {}
        public Builder(TableSchema defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.compositePartitionKey = defaults.compositePartitionKey;
        }

        @CustomType.Setter
        public Builder compositePartitionKey(@Nullable TableSchemaCompositePartitionKey compositePartitionKey) {
            this.compositePartitionKey = compositePartitionKey;
            return this;
        }
        public TableSchema build() {
            final var o = new TableSchema();
            o.compositePartitionKey = compositePartitionKey;
            return o;
        }
    }
}
