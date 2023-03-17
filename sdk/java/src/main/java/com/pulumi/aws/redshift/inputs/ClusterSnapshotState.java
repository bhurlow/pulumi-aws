// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.redshift.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClusterSnapshotState extends com.pulumi.resources.ResourceArgs {

    public static final ClusterSnapshotState Empty = new ClusterSnapshotState();

    /**
     * Amazon Resource Name (ARN) of the snapshot.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) of the snapshot.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The cluster identifier for which you want a snapshot.
     * 
     */
    @Import(name="clusterIdentifier")
    private @Nullable Output<String> clusterIdentifier;

    /**
     * @return The cluster identifier for which you want a snapshot.
     * 
     */
    public Optional<Output<String>> clusterIdentifier() {
        return Optional.ofNullable(this.clusterIdentifier);
    }

    /**
     * The Key Management Service (KMS) key ID of the encryption key that was used to encrypt data in the cluster from which the snapshot was taken.
     * 
     */
    @Import(name="kmsKeyId")
    private @Nullable Output<String> kmsKeyId;

    /**
     * @return The Key Management Service (KMS) key ID of the encryption key that was used to encrypt data in the cluster from which the snapshot was taken.
     * 
     */
    public Optional<Output<String>> kmsKeyId() {
        return Optional.ofNullable(this.kmsKeyId);
    }

    /**
     * The number of days that a manual snapshot is retained. If the value is `-1`, the manual snapshot is retained indefinitely. Valid values are -1 and between `1` and `3653`.
     * 
     */
    @Import(name="manualSnapshotRetentionPeriod")
    private @Nullable Output<Integer> manualSnapshotRetentionPeriod;

    /**
     * @return The number of days that a manual snapshot is retained. If the value is `-1`, the manual snapshot is retained indefinitely. Valid values are -1 and between `1` and `3653`.
     * 
     */
    public Optional<Output<Integer>> manualSnapshotRetentionPeriod() {
        return Optional.ofNullable(this.manualSnapshotRetentionPeriod);
    }

    /**
     * For manual snapshots, the Amazon Web Services account used to create or copy the snapshot. For automatic snapshots, the owner of the cluster. The owner can perform all snapshot actions, such as sharing a manual snapshot.
     * 
     */
    @Import(name="ownerAccount")
    private @Nullable Output<String> ownerAccount;

    /**
     * @return For manual snapshots, the Amazon Web Services account used to create or copy the snapshot. For automatic snapshots, the owner of the cluster. The owner can perform all snapshot actions, such as sharing a manual snapshot.
     * 
     */
    public Optional<Output<String>> ownerAccount() {
        return Optional.ofNullable(this.ownerAccount);
    }

    /**
     * A unique identifier for the snapshot that you are requesting. This identifier must be unique for all snapshots within the Amazon Web Services account.
     * 
     */
    @Import(name="snapshotIdentifier")
    private @Nullable Output<String> snapshotIdentifier;

    /**
     * @return A unique identifier for the snapshot that you are requesting. This identifier must be unique for all snapshots within the Amazon Web Services account.
     * 
     */
    public Optional<Output<String>> snapshotIdentifier() {
        return Optional.ofNullable(this.snapshotIdentifier);
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

    private ClusterSnapshotState() {}

    private ClusterSnapshotState(ClusterSnapshotState $) {
        this.arn = $.arn;
        this.clusterIdentifier = $.clusterIdentifier;
        this.kmsKeyId = $.kmsKeyId;
        this.manualSnapshotRetentionPeriod = $.manualSnapshotRetentionPeriod;
        this.ownerAccount = $.ownerAccount;
        this.snapshotIdentifier = $.snapshotIdentifier;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClusterSnapshotState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClusterSnapshotState $;

        public Builder() {
            $ = new ClusterSnapshotState();
        }

        public Builder(ClusterSnapshotState defaults) {
            $ = new ClusterSnapshotState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn Amazon Resource Name (ARN) of the snapshot.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn Amazon Resource Name (ARN) of the snapshot.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param clusterIdentifier The cluster identifier for which you want a snapshot.
         * 
         * @return builder
         * 
         */
        public Builder clusterIdentifier(@Nullable Output<String> clusterIdentifier) {
            $.clusterIdentifier = clusterIdentifier;
            return this;
        }

        /**
         * @param clusterIdentifier The cluster identifier for which you want a snapshot.
         * 
         * @return builder
         * 
         */
        public Builder clusterIdentifier(String clusterIdentifier) {
            return clusterIdentifier(Output.of(clusterIdentifier));
        }

        /**
         * @param kmsKeyId The Key Management Service (KMS) key ID of the encryption key that was used to encrypt data in the cluster from which the snapshot was taken.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyId(@Nullable Output<String> kmsKeyId) {
            $.kmsKeyId = kmsKeyId;
            return this;
        }

        /**
         * @param kmsKeyId The Key Management Service (KMS) key ID of the encryption key that was used to encrypt data in the cluster from which the snapshot was taken.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyId(String kmsKeyId) {
            return kmsKeyId(Output.of(kmsKeyId));
        }

        /**
         * @param manualSnapshotRetentionPeriod The number of days that a manual snapshot is retained. If the value is `-1`, the manual snapshot is retained indefinitely. Valid values are -1 and between `1` and `3653`.
         * 
         * @return builder
         * 
         */
        public Builder manualSnapshotRetentionPeriod(@Nullable Output<Integer> manualSnapshotRetentionPeriod) {
            $.manualSnapshotRetentionPeriod = manualSnapshotRetentionPeriod;
            return this;
        }

        /**
         * @param manualSnapshotRetentionPeriod The number of days that a manual snapshot is retained. If the value is `-1`, the manual snapshot is retained indefinitely. Valid values are -1 and between `1` and `3653`.
         * 
         * @return builder
         * 
         */
        public Builder manualSnapshotRetentionPeriod(Integer manualSnapshotRetentionPeriod) {
            return manualSnapshotRetentionPeriod(Output.of(manualSnapshotRetentionPeriod));
        }

        /**
         * @param ownerAccount For manual snapshots, the Amazon Web Services account used to create or copy the snapshot. For automatic snapshots, the owner of the cluster. The owner can perform all snapshot actions, such as sharing a manual snapshot.
         * 
         * @return builder
         * 
         */
        public Builder ownerAccount(@Nullable Output<String> ownerAccount) {
            $.ownerAccount = ownerAccount;
            return this;
        }

        /**
         * @param ownerAccount For manual snapshots, the Amazon Web Services account used to create or copy the snapshot. For automatic snapshots, the owner of the cluster. The owner can perform all snapshot actions, such as sharing a manual snapshot.
         * 
         * @return builder
         * 
         */
        public Builder ownerAccount(String ownerAccount) {
            return ownerAccount(Output.of(ownerAccount));
        }

        /**
         * @param snapshotIdentifier A unique identifier for the snapshot that you are requesting. This identifier must be unique for all snapshots within the Amazon Web Services account.
         * 
         * @return builder
         * 
         */
        public Builder snapshotIdentifier(@Nullable Output<String> snapshotIdentifier) {
            $.snapshotIdentifier = snapshotIdentifier;
            return this;
        }

        /**
         * @param snapshotIdentifier A unique identifier for the snapshot that you are requesting. This identifier must be unique for all snapshots within the Amazon Web Services account.
         * 
         * @return builder
         * 
         */
        public Builder snapshotIdentifier(String snapshotIdentifier) {
            return snapshotIdentifier(Output.of(snapshotIdentifier));
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

        public ClusterSnapshotState build() {
            return $;
        }
    }

}
