// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ebs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VolumeArgs extends com.pulumi.resources.ResourceArgs {

    public static final VolumeArgs Empty = new VolumeArgs();

    /**
     * The AZ where the EBS volume will exist.
     * 
     */
    @Import(name="availabilityZone", required=true)
    private Output<String> availabilityZone;

    /**
     * @return The AZ where the EBS volume will exist.
     * 
     */
    public Output<String> availabilityZone() {
        return this.availabilityZone;
    }

    /**
     * If true, the disk will be encrypted.
     * 
     */
    @Import(name="encrypted")
    private @Nullable Output<Boolean> encrypted;

    /**
     * @return If true, the disk will be encrypted.
     * 
     */
    public Optional<Output<Boolean>> encrypted() {
        return Optional.ofNullable(this.encrypted);
    }

    /**
     * If true, snapshot will be created before volume deletion. Any tags on the volume will be migrated to the snapshot. By default set to false
     * 
     */
    @Import(name="finalSnapshot")
    private @Nullable Output<Boolean> finalSnapshot;

    /**
     * @return If true, snapshot will be created before volume deletion. Any tags on the volume will be migrated to the snapshot. By default set to false
     * 
     */
    public Optional<Output<Boolean>> finalSnapshot() {
        return Optional.ofNullable(this.finalSnapshot);
    }

    /**
     * The amount of IOPS to provision for the disk. Only valid for `type` of `io1`, `io2` or `gp3`.
     * 
     */
    @Import(name="iops")
    private @Nullable Output<Integer> iops;

    /**
     * @return The amount of IOPS to provision for the disk. Only valid for `type` of `io1`, `io2` or `gp3`.
     * 
     */
    public Optional<Output<Integer>> iops() {
        return Optional.ofNullable(this.iops);
    }

    /**
     * The ARN for the KMS encryption key. When specifying `kms_key_id`, `encrypted` needs to be set to true. Note: The provider must be running with credentials which have the `GenerateDataKeyWithoutPlaintext` permission on the specified KMS key as required by the [EBS KMS CMK volume provisioning process](https://docs.aws.amazon.com/kms/latest/developerguide/services-ebs.html#ebs-cmk) to prevent a volume from being created and almost immediately deleted.
     * 
     */
    @Import(name="kmsKeyId")
    private @Nullable Output<String> kmsKeyId;

    /**
     * @return The ARN for the KMS encryption key. When specifying `kms_key_id`, `encrypted` needs to be set to true. Note: The provider must be running with credentials which have the `GenerateDataKeyWithoutPlaintext` permission on the specified KMS key as required by the [EBS KMS CMK volume provisioning process](https://docs.aws.amazon.com/kms/latest/developerguide/services-ebs.html#ebs-cmk) to prevent a volume from being created and almost immediately deleted.
     * 
     */
    public Optional<Output<String>> kmsKeyId() {
        return Optional.ofNullable(this.kmsKeyId);
    }

    /**
     * Specifies whether to enable Amazon EBS Multi-Attach. Multi-Attach is supported on `io1` and `io2` volumes.
     * 
     */
    @Import(name="multiAttachEnabled")
    private @Nullable Output<Boolean> multiAttachEnabled;

    /**
     * @return Specifies whether to enable Amazon EBS Multi-Attach. Multi-Attach is supported on `io1` and `io2` volumes.
     * 
     */
    public Optional<Output<Boolean>> multiAttachEnabled() {
        return Optional.ofNullable(this.multiAttachEnabled);
    }

    /**
     * The Amazon Resource Name (ARN) of the Outpost.
     * 
     */
    @Import(name="outpostArn")
    private @Nullable Output<String> outpostArn;

    /**
     * @return The Amazon Resource Name (ARN) of the Outpost.
     * 
     */
    public Optional<Output<String>> outpostArn() {
        return Optional.ofNullable(this.outpostArn);
    }

    /**
     * The size of the drive in GiBs.
     * 
     */
    @Import(name="size")
    private @Nullable Output<Integer> size;

    /**
     * @return The size of the drive in GiBs.
     * 
     */
    public Optional<Output<Integer>> size() {
        return Optional.ofNullable(this.size);
    }

    /**
     * A snapshot to base the EBS volume off of.
     * 
     */
    @Import(name="snapshotId")
    private @Nullable Output<String> snapshotId;

    /**
     * @return A snapshot to base the EBS volume off of.
     * 
     */
    public Optional<Output<String>> snapshotId() {
        return Optional.ofNullable(this.snapshotId);
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
     * The throughput that the volume supports, in MiB/s. Only valid for `type` of `gp3`.
     * 
     */
    @Import(name="throughput")
    private @Nullable Output<Integer> throughput;

    /**
     * @return The throughput that the volume supports, in MiB/s. Only valid for `type` of `gp3`.
     * 
     */
    public Optional<Output<Integer>> throughput() {
        return Optional.ofNullable(this.throughput);
    }

    /**
     * The type of EBS volume. Can be `standard`, `gp2`, `gp3`, `io1`, `io2`, `sc1` or `st1` (Default: `gp2`).
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The type of EBS volume. Can be `standard`, `gp2`, `gp3`, `io1`, `io2`, `sc1` or `st1` (Default: `gp2`).
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private VolumeArgs() {}

    private VolumeArgs(VolumeArgs $) {
        this.availabilityZone = $.availabilityZone;
        this.encrypted = $.encrypted;
        this.finalSnapshot = $.finalSnapshot;
        this.iops = $.iops;
        this.kmsKeyId = $.kmsKeyId;
        this.multiAttachEnabled = $.multiAttachEnabled;
        this.outpostArn = $.outpostArn;
        this.size = $.size;
        this.snapshotId = $.snapshotId;
        this.tags = $.tags;
        this.throughput = $.throughput;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VolumeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VolumeArgs $;

        public Builder() {
            $ = new VolumeArgs();
        }

        public Builder(VolumeArgs defaults) {
            $ = new VolumeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param availabilityZone The AZ where the EBS volume will exist.
         * 
         * @return builder
         * 
         */
        public Builder availabilityZone(Output<String> availabilityZone) {
            $.availabilityZone = availabilityZone;
            return this;
        }

        /**
         * @param availabilityZone The AZ where the EBS volume will exist.
         * 
         * @return builder
         * 
         */
        public Builder availabilityZone(String availabilityZone) {
            return availabilityZone(Output.of(availabilityZone));
        }

        /**
         * @param encrypted If true, the disk will be encrypted.
         * 
         * @return builder
         * 
         */
        public Builder encrypted(@Nullable Output<Boolean> encrypted) {
            $.encrypted = encrypted;
            return this;
        }

        /**
         * @param encrypted If true, the disk will be encrypted.
         * 
         * @return builder
         * 
         */
        public Builder encrypted(Boolean encrypted) {
            return encrypted(Output.of(encrypted));
        }

        /**
         * @param finalSnapshot If true, snapshot will be created before volume deletion. Any tags on the volume will be migrated to the snapshot. By default set to false
         * 
         * @return builder
         * 
         */
        public Builder finalSnapshot(@Nullable Output<Boolean> finalSnapshot) {
            $.finalSnapshot = finalSnapshot;
            return this;
        }

        /**
         * @param finalSnapshot If true, snapshot will be created before volume deletion. Any tags on the volume will be migrated to the snapshot. By default set to false
         * 
         * @return builder
         * 
         */
        public Builder finalSnapshot(Boolean finalSnapshot) {
            return finalSnapshot(Output.of(finalSnapshot));
        }

        /**
         * @param iops The amount of IOPS to provision for the disk. Only valid for `type` of `io1`, `io2` or `gp3`.
         * 
         * @return builder
         * 
         */
        public Builder iops(@Nullable Output<Integer> iops) {
            $.iops = iops;
            return this;
        }

        /**
         * @param iops The amount of IOPS to provision for the disk. Only valid for `type` of `io1`, `io2` or `gp3`.
         * 
         * @return builder
         * 
         */
        public Builder iops(Integer iops) {
            return iops(Output.of(iops));
        }

        /**
         * @param kmsKeyId The ARN for the KMS encryption key. When specifying `kms_key_id`, `encrypted` needs to be set to true. Note: The provider must be running with credentials which have the `GenerateDataKeyWithoutPlaintext` permission on the specified KMS key as required by the [EBS KMS CMK volume provisioning process](https://docs.aws.amazon.com/kms/latest/developerguide/services-ebs.html#ebs-cmk) to prevent a volume from being created and almost immediately deleted.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyId(@Nullable Output<String> kmsKeyId) {
            $.kmsKeyId = kmsKeyId;
            return this;
        }

        /**
         * @param kmsKeyId The ARN for the KMS encryption key. When specifying `kms_key_id`, `encrypted` needs to be set to true. Note: The provider must be running with credentials which have the `GenerateDataKeyWithoutPlaintext` permission on the specified KMS key as required by the [EBS KMS CMK volume provisioning process](https://docs.aws.amazon.com/kms/latest/developerguide/services-ebs.html#ebs-cmk) to prevent a volume from being created and almost immediately deleted.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyId(String kmsKeyId) {
            return kmsKeyId(Output.of(kmsKeyId));
        }

        /**
         * @param multiAttachEnabled Specifies whether to enable Amazon EBS Multi-Attach. Multi-Attach is supported on `io1` and `io2` volumes.
         * 
         * @return builder
         * 
         */
        public Builder multiAttachEnabled(@Nullable Output<Boolean> multiAttachEnabled) {
            $.multiAttachEnabled = multiAttachEnabled;
            return this;
        }

        /**
         * @param multiAttachEnabled Specifies whether to enable Amazon EBS Multi-Attach. Multi-Attach is supported on `io1` and `io2` volumes.
         * 
         * @return builder
         * 
         */
        public Builder multiAttachEnabled(Boolean multiAttachEnabled) {
            return multiAttachEnabled(Output.of(multiAttachEnabled));
        }

        /**
         * @param outpostArn The Amazon Resource Name (ARN) of the Outpost.
         * 
         * @return builder
         * 
         */
        public Builder outpostArn(@Nullable Output<String> outpostArn) {
            $.outpostArn = outpostArn;
            return this;
        }

        /**
         * @param outpostArn The Amazon Resource Name (ARN) of the Outpost.
         * 
         * @return builder
         * 
         */
        public Builder outpostArn(String outpostArn) {
            return outpostArn(Output.of(outpostArn));
        }

        /**
         * @param size The size of the drive in GiBs.
         * 
         * @return builder
         * 
         */
        public Builder size(@Nullable Output<Integer> size) {
            $.size = size;
            return this;
        }

        /**
         * @param size The size of the drive in GiBs.
         * 
         * @return builder
         * 
         */
        public Builder size(Integer size) {
            return size(Output.of(size));
        }

        /**
         * @param snapshotId A snapshot to base the EBS volume off of.
         * 
         * @return builder
         * 
         */
        public Builder snapshotId(@Nullable Output<String> snapshotId) {
            $.snapshotId = snapshotId;
            return this;
        }

        /**
         * @param snapshotId A snapshot to base the EBS volume off of.
         * 
         * @return builder
         * 
         */
        public Builder snapshotId(String snapshotId) {
            return snapshotId(Output.of(snapshotId));
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
         * @param throughput The throughput that the volume supports, in MiB/s. Only valid for `type` of `gp3`.
         * 
         * @return builder
         * 
         */
        public Builder throughput(@Nullable Output<Integer> throughput) {
            $.throughput = throughput;
            return this;
        }

        /**
         * @param throughput The throughput that the volume supports, in MiB/s. Only valid for `type` of `gp3`.
         * 
         * @return builder
         * 
         */
        public Builder throughput(Integer throughput) {
            return throughput(Output.of(throughput));
        }

        /**
         * @param type The type of EBS volume. Can be `standard`, `gp2`, `gp3`, `io1`, `io2`, `sc1` or `st1` (Default: `gp2`).
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of EBS volume. Can be `standard`, `gp2`, `gp3`, `io1`, `io2`, `sc1` or `st1` (Default: `gp2`).
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public VolumeArgs build() {
            $.availabilityZone = Objects.requireNonNull($.availabilityZone, "expected parameter 'availabilityZone' to be non-null");
            return $;
        }
    }

}
