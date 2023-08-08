// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.finspace.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class KxClusterSavedownStorageConfigurationArgs extends com.pulumi.resources.ResourceArgs {

    public static final KxClusterSavedownStorageConfigurationArgs Empty = new KxClusterSavedownStorageConfigurationArgs();

    /**
     * Size of temporary storage in gigabytes. Must be between 10 and 16000.
     * 
     */
    @Import(name="size", required=true)
    private Output<Integer> size;

    /**
     * @return Size of temporary storage in gigabytes. Must be between 10 and 16000.
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }

    /**
     * Type of writeable storage space for temporarily storing your savedown data. The valid values are:
     * * SDS01 - This type represents 3000 IOPS and io2 ebs volume type.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Type of writeable storage space for temporarily storing your savedown data. The valid values are:
     * * SDS01 - This type represents 3000 IOPS and io2 ebs volume type.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private KxClusterSavedownStorageConfigurationArgs() {}

    private KxClusterSavedownStorageConfigurationArgs(KxClusterSavedownStorageConfigurationArgs $) {
        this.size = $.size;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KxClusterSavedownStorageConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KxClusterSavedownStorageConfigurationArgs $;

        public Builder() {
            $ = new KxClusterSavedownStorageConfigurationArgs();
        }

        public Builder(KxClusterSavedownStorageConfigurationArgs defaults) {
            $ = new KxClusterSavedownStorageConfigurationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param size Size of temporary storage in gigabytes. Must be between 10 and 16000.
         * 
         * @return builder
         * 
         */
        public Builder size(Output<Integer> size) {
            $.size = size;
            return this;
        }

        /**
         * @param size Size of temporary storage in gigabytes. Must be between 10 and 16000.
         * 
         * @return builder
         * 
         */
        public Builder size(Integer size) {
            return size(Output.of(size));
        }

        /**
         * @param type Type of writeable storage space for temporarily storing your savedown data. The valid values are:
         * * SDS01 - This type represents 3000 IOPS and io2 ebs volume type.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Type of writeable storage space for temporarily storing your savedown data. The valid values are:
         * * SDS01 - This type represents 3000 IOPS and io2 ebs volume type.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public KxClusterSavedownStorageConfigurationArgs build() {
            $.size = Objects.requireNonNull($.size, "expected parameter 'size' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
