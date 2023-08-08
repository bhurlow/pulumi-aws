// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.finspace.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class KxClusterCacheStorageConfigurationArgs extends com.pulumi.resources.ResourceArgs {

    public static final KxClusterCacheStorageConfigurationArgs Empty = new KxClusterCacheStorageConfigurationArgs();

    /**
     * Size of cache in Gigabytes.
     * 
     */
    @Import(name="size", required=true)
    private Output<Integer> size;

    /**
     * @return Size of cache in Gigabytes.
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }

    /**
     * Type of KDB database. The following types are available:
     * * HDB - Historical Database. The data is only accessible with read-only permissions from one of the FinSpace managed KX databases mounted to the cluster.
     * * RDB - Realtime Database. This type of database captures all the data from a ticker plant and stores it in memory until the end of day, after which it writes all of its data to a disk and reloads the HDB. This cluster type requires local storage for temporary storage of data during the savedown process. If you specify this field in your request, you must provide the `savedownStorageConfiguration` parameter.
     * * GATEWAY - A gateway cluster allows you to access data across processes in kdb systems. It allows you to create your own routing logic using the initialization scripts and custom code. This type of cluster does not require a  writable local storage.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Type of KDB database. The following types are available:
     * * HDB - Historical Database. The data is only accessible with read-only permissions from one of the FinSpace managed KX databases mounted to the cluster.
     * * RDB - Realtime Database. This type of database captures all the data from a ticker plant and stores it in memory until the end of day, after which it writes all of its data to a disk and reloads the HDB. This cluster type requires local storage for temporary storage of data during the savedown process. If you specify this field in your request, you must provide the `savedownStorageConfiguration` parameter.
     * * GATEWAY - A gateway cluster allows you to access data across processes in kdb systems. It allows you to create your own routing logic using the initialization scripts and custom code. This type of cluster does not require a  writable local storage.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private KxClusterCacheStorageConfigurationArgs() {}

    private KxClusterCacheStorageConfigurationArgs(KxClusterCacheStorageConfigurationArgs $) {
        this.size = $.size;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KxClusterCacheStorageConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KxClusterCacheStorageConfigurationArgs $;

        public Builder() {
            $ = new KxClusterCacheStorageConfigurationArgs();
        }

        public Builder(KxClusterCacheStorageConfigurationArgs defaults) {
            $ = new KxClusterCacheStorageConfigurationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param size Size of cache in Gigabytes.
         * 
         * @return builder
         * 
         */
        public Builder size(Output<Integer> size) {
            $.size = size;
            return this;
        }

        /**
         * @param size Size of cache in Gigabytes.
         * 
         * @return builder
         * 
         */
        public Builder size(Integer size) {
            return size(Output.of(size));
        }

        /**
         * @param type Type of KDB database. The following types are available:
         * * HDB - Historical Database. The data is only accessible with read-only permissions from one of the FinSpace managed KX databases mounted to the cluster.
         * * RDB - Realtime Database. This type of database captures all the data from a ticker plant and stores it in memory until the end of day, after which it writes all of its data to a disk and reloads the HDB. This cluster type requires local storage for temporary storage of data during the savedown process. If you specify this field in your request, you must provide the `savedownStorageConfiguration` parameter.
         * * GATEWAY - A gateway cluster allows you to access data across processes in kdb systems. It allows you to create your own routing logic using the initialization scripts and custom code. This type of cluster does not require a  writable local storage.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Type of KDB database. The following types are available:
         * * HDB - Historical Database. The data is only accessible with read-only permissions from one of the FinSpace managed KX databases mounted to the cluster.
         * * RDB - Realtime Database. This type of database captures all the data from a ticker plant and stores it in memory until the end of day, after which it writes all of its data to a disk and reloads the HDB. This cluster type requires local storage for temporary storage of data during the savedown process. If you specify this field in your request, you must provide the `savedownStorageConfiguration` parameter.
         * * GATEWAY - A gateway cluster allows you to access data across processes in kdb systems. It allows you to create your own routing logic using the initialization scripts and custom code. This type of cluster does not require a  writable local storage.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public KxClusterCacheStorageConfigurationArgs build() {
            $.size = Objects.requireNonNull($.size, "expected parameter 'size' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
