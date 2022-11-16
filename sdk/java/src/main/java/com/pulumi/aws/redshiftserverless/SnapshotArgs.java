// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.redshiftserverless;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SnapshotArgs extends com.pulumi.resources.ResourceArgs {

    public static final SnapshotArgs Empty = new SnapshotArgs();

    /**
     * The namespace to create a snapshot for.
     * 
     */
    @Import(name="namespaceName", required=true)
    private Output<String> namespaceName;

    /**
     * @return The namespace to create a snapshot for.
     * 
     */
    public Output<String> namespaceName() {
        return this.namespaceName;
    }

    /**
     * How long to retain the created snapshot. Default value is `-1`.
     * 
     */
    @Import(name="retentionPeriod")
    private @Nullable Output<Integer> retentionPeriod;

    /**
     * @return How long to retain the created snapshot. Default value is `-1`.
     * 
     */
    public Optional<Output<Integer>> retentionPeriod() {
        return Optional.ofNullable(this.retentionPeriod);
    }

    /**
     * The name of the snapshot.
     * 
     */
    @Import(name="snapshotName", required=true)
    private Output<String> snapshotName;

    /**
     * @return The name of the snapshot.
     * 
     */
    public Output<String> snapshotName() {
        return this.snapshotName;
    }

    private SnapshotArgs() {}

    private SnapshotArgs(SnapshotArgs $) {
        this.namespaceName = $.namespaceName;
        this.retentionPeriod = $.retentionPeriod;
        this.snapshotName = $.snapshotName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SnapshotArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SnapshotArgs $;

        public Builder() {
            $ = new SnapshotArgs();
        }

        public Builder(SnapshotArgs defaults) {
            $ = new SnapshotArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param namespaceName The namespace to create a snapshot for.
         * 
         * @return builder
         * 
         */
        public Builder namespaceName(Output<String> namespaceName) {
            $.namespaceName = namespaceName;
            return this;
        }

        /**
         * @param namespaceName The namespace to create a snapshot for.
         * 
         * @return builder
         * 
         */
        public Builder namespaceName(String namespaceName) {
            return namespaceName(Output.of(namespaceName));
        }

        /**
         * @param retentionPeriod How long to retain the created snapshot. Default value is `-1`.
         * 
         * @return builder
         * 
         */
        public Builder retentionPeriod(@Nullable Output<Integer> retentionPeriod) {
            $.retentionPeriod = retentionPeriod;
            return this;
        }

        /**
         * @param retentionPeriod How long to retain the created snapshot. Default value is `-1`.
         * 
         * @return builder
         * 
         */
        public Builder retentionPeriod(Integer retentionPeriod) {
            return retentionPeriod(Output.of(retentionPeriod));
        }

        /**
         * @param snapshotName The name of the snapshot.
         * 
         * @return builder
         * 
         */
        public Builder snapshotName(Output<String> snapshotName) {
            $.snapshotName = snapshotName;
            return this;
        }

        /**
         * @param snapshotName The name of the snapshot.
         * 
         * @return builder
         * 
         */
        public Builder snapshotName(String snapshotName) {
            return snapshotName(Output.of(snapshotName));
        }

        public SnapshotArgs build() {
            $.namespaceName = Objects.requireNonNull($.namespaceName, "expected parameter 'namespaceName' to be non-null");
            $.snapshotName = Objects.requireNonNull($.snapshotName, "expected parameter 'snapshotName' to be non-null");
            return $;
        }
    }

}
