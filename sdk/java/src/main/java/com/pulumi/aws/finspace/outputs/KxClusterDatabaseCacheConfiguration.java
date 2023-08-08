// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.finspace.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class KxClusterDatabaseCacheConfiguration {
    /**
     * @return Type of disk cache.
     * 
     */
    private String cacheType;
    /**
     * @return Paths within the database to cache.
     * 
     */
    private @Nullable List<String> dbPaths;

    private KxClusterDatabaseCacheConfiguration() {}
    /**
     * @return Type of disk cache.
     * 
     */
    public String cacheType() {
        return this.cacheType;
    }
    /**
     * @return Paths within the database to cache.
     * 
     */
    public List<String> dbPaths() {
        return this.dbPaths == null ? List.of() : this.dbPaths;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(KxClusterDatabaseCacheConfiguration defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String cacheType;
        private @Nullable List<String> dbPaths;
        public Builder() {}
        public Builder(KxClusterDatabaseCacheConfiguration defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cacheType = defaults.cacheType;
    	      this.dbPaths = defaults.dbPaths;
        }

        @CustomType.Setter
        public Builder cacheType(String cacheType) {
            this.cacheType = Objects.requireNonNull(cacheType);
            return this;
        }
        @CustomType.Setter
        public Builder dbPaths(@Nullable List<String> dbPaths) {
            this.dbPaths = dbPaths;
            return this;
        }
        public Builder dbPaths(String... dbPaths) {
            return dbPaths(List.of(dbPaths));
        }
        public KxClusterDatabaseCacheConfiguration build() {
            final var o = new KxClusterDatabaseCacheConfiguration();
            o.cacheType = cacheType;
            o.dbPaths = dbPaths;
            return o;
        }
    }
}
