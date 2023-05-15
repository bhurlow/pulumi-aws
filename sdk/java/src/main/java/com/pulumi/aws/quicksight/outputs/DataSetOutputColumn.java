// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.quicksight.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DataSetOutputColumn {
    /**
     * @return Field folder description.
     * 
     */
    private @Nullable String description;
    /**
     * @return Display name for the dataset.
     * 
     */
    private @Nullable String name;
    /**
     * @return Data type of the column.
     * 
     */
    private @Nullable String type;

    private DataSetOutputColumn() {}
    /**
     * @return Field folder description.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return Display name for the dataset.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return Data type of the column.
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DataSetOutputColumn defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String description;
        private @Nullable String name;
        private @Nullable String type;
        public Builder() {}
        public Builder(DataSetOutputColumn defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.name = defaults.name;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder type(@Nullable String type) {
            this.type = type;
            return this;
        }
        public DataSetOutputColumn build() {
            final var o = new DataSetOutputColumn();
            o.description = description;
            o.name = name;
            o.type = type;
            return o;
        }
    }
}
