// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.glue.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetCatalogTableTargetTable {
    /**
     * @return ID of the Glue Catalog and database where the table metadata resides. If omitted, this defaults to the current AWS Account ID.
     * 
     */
    private String catalogId;
    /**
     * @return Name of the metadata database where the table metadata resides.
     * 
     */
    private String databaseName;
    /**
     * @return Name of the table.
     * 
     */
    private String name;

    private GetCatalogTableTargetTable() {}
    /**
     * @return ID of the Glue Catalog and database where the table metadata resides. If omitted, this defaults to the current AWS Account ID.
     * 
     */
    public String catalogId() {
        return this.catalogId;
    }
    /**
     * @return Name of the metadata database where the table metadata resides.
     * 
     */
    public String databaseName() {
        return this.databaseName;
    }
    /**
     * @return Name of the table.
     * 
     */
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCatalogTableTargetTable defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String catalogId;
        private String databaseName;
        private String name;
        public Builder() {}
        public Builder(GetCatalogTableTargetTable defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.catalogId = defaults.catalogId;
    	      this.databaseName = defaults.databaseName;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder catalogId(String catalogId) {
            this.catalogId = Objects.requireNonNull(catalogId);
            return this;
        }
        @CustomType.Setter
        public Builder databaseName(String databaseName) {
            this.databaseName = Objects.requireNonNull(databaseName);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public GetCatalogTableTargetTable build() {
            final var o = new GetCatalogTableTargetTable();
            o.catalogId = catalogId;
            o.databaseName = databaseName;
            o.name = name;
            return o;
        }
    }
}
