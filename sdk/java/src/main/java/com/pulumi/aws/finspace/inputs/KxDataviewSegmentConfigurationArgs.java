// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.finspace.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class KxDataviewSegmentConfigurationArgs extends com.pulumi.resources.ResourceArgs {

    public static final KxDataviewSegmentConfigurationArgs Empty = new KxDataviewSegmentConfigurationArgs();

    /**
     * The database path of the data that you want to place on each selected volume. Each segment must have a unique database path for each volume.
     * 
     */
    @Import(name="dbPaths", required=true)
    private Output<List<String>> dbPaths;

    /**
     * @return The database path of the data that you want to place on each selected volume. Each segment must have a unique database path for each volume.
     * 
     */
    public Output<List<String>> dbPaths() {
        return this.dbPaths;
    }

    /**
     * The name of the volume that you want to attach to a dataview. This volume must be in the same availability zone as the dataview that you are attaching to.
     * 
     */
    @Import(name="volumeName", required=true)
    private Output<String> volumeName;

    /**
     * @return The name of the volume that you want to attach to a dataview. This volume must be in the same availability zone as the dataview that you are attaching to.
     * 
     */
    public Output<String> volumeName() {
        return this.volumeName;
    }

    private KxDataviewSegmentConfigurationArgs() {}

    private KxDataviewSegmentConfigurationArgs(KxDataviewSegmentConfigurationArgs $) {
        this.dbPaths = $.dbPaths;
        this.volumeName = $.volumeName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KxDataviewSegmentConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KxDataviewSegmentConfigurationArgs $;

        public Builder() {
            $ = new KxDataviewSegmentConfigurationArgs();
        }

        public Builder(KxDataviewSegmentConfigurationArgs defaults) {
            $ = new KxDataviewSegmentConfigurationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param dbPaths The database path of the data that you want to place on each selected volume. Each segment must have a unique database path for each volume.
         * 
         * @return builder
         * 
         */
        public Builder dbPaths(Output<List<String>> dbPaths) {
            $.dbPaths = dbPaths;
            return this;
        }

        /**
         * @param dbPaths The database path of the data that you want to place on each selected volume. Each segment must have a unique database path for each volume.
         * 
         * @return builder
         * 
         */
        public Builder dbPaths(List<String> dbPaths) {
            return dbPaths(Output.of(dbPaths));
        }

        /**
         * @param dbPaths The database path of the data that you want to place on each selected volume. Each segment must have a unique database path for each volume.
         * 
         * @return builder
         * 
         */
        public Builder dbPaths(String... dbPaths) {
            return dbPaths(List.of(dbPaths));
        }

        /**
         * @param volumeName The name of the volume that you want to attach to a dataview. This volume must be in the same availability zone as the dataview that you are attaching to.
         * 
         * @return builder
         * 
         */
        public Builder volumeName(Output<String> volumeName) {
            $.volumeName = volumeName;
            return this;
        }

        /**
         * @param volumeName The name of the volume that you want to attach to a dataview. This volume must be in the same availability zone as the dataview that you are attaching to.
         * 
         * @return builder
         * 
         */
        public Builder volumeName(String volumeName) {
            return volumeName(Output.of(volumeName));
        }

        public KxDataviewSegmentConfigurationArgs build() {
            $.dbPaths = Objects.requireNonNull($.dbPaths, "expected parameter 'dbPaths' to be non-null");
            $.volumeName = Objects.requireNonNull($.volumeName, "expected parameter 'volumeName' to be non-null");
            return $;
        }
    }

}
