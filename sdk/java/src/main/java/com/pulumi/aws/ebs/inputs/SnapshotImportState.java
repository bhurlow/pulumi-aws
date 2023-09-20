// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ebs.inputs;

import com.pulumi.aws.ebs.inputs.SnapshotImportClientDataArgs;
import com.pulumi.aws.ebs.inputs.SnapshotImportDiskContainerArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SnapshotImportState extends com.pulumi.resources.ResourceArgs {

    public static final SnapshotImportState Empty = new SnapshotImportState();

    /**
     * Amazon Resource Name (ARN) of the EBS Snapshot.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return Amazon Resource Name (ARN) of the EBS Snapshot.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The client-specific data. Detailed below.
     * 
     */
    @Import(name="clientData")
    private @Nullable Output<SnapshotImportClientDataArgs> clientData;

    /**
     * @return The client-specific data. Detailed below.
     * 
     */
    public Optional<Output<SnapshotImportClientDataArgs>> clientData() {
        return Optional.ofNullable(this.clientData);
    }

    /**
     * The data encryption key identifier for the snapshot.
     * 
     */
    @Import(name="dataEncryptionKeyId")
    private @Nullable Output<String> dataEncryptionKeyId;

    /**
     * @return The data encryption key identifier for the snapshot.
     * 
     */
    public Optional<Output<String>> dataEncryptionKeyId() {
        return Optional.ofNullable(this.dataEncryptionKeyId);
    }

    /**
     * The description string for the import snapshot task.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description string for the import snapshot task.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Information about the disk container. Detailed below.
     * 
     */
    @Import(name="diskContainer")
    private @Nullable Output<SnapshotImportDiskContainerArgs> diskContainer;

    /**
     * @return Information about the disk container. Detailed below.
     * 
     */
    public Optional<Output<SnapshotImportDiskContainerArgs>> diskContainer() {
        return Optional.ofNullable(this.diskContainer);
    }

    /**
     * Specifies whether the destination snapshot of the imported image should be encrypted. The default KMS key for EBS is used unless you specify a non-default KMS key using KmsKeyId.
     * 
     */
    @Import(name="encrypted")
    private @Nullable Output<Boolean> encrypted;

    /**
     * @return Specifies whether the destination snapshot of the imported image should be encrypted. The default KMS key for EBS is used unless you specify a non-default KMS key using KmsKeyId.
     * 
     */
    public Optional<Output<Boolean>> encrypted() {
        return Optional.ofNullable(this.encrypted);
    }

    /**
     * An identifier for the symmetric KMS key to use when creating the encrypted snapshot. This parameter is only required if you want to use a non-default KMS key; if this parameter is not specified, the default KMS key for EBS is used. If a KmsKeyId is specified, the Encrypted flag must also be set.
     * 
     */
    @Import(name="kmsKeyId")
    private @Nullable Output<String> kmsKeyId;

    /**
     * @return An identifier for the symmetric KMS key to use when creating the encrypted snapshot. This parameter is only required if you want to use a non-default KMS key; if this parameter is not specified, the default KMS key for EBS is used. If a KmsKeyId is specified, the Encrypted flag must also be set.
     * 
     */
    public Optional<Output<String>> kmsKeyId() {
        return Optional.ofNullable(this.kmsKeyId);
    }

    @Import(name="outpostArn")
    private @Nullable Output<String> outpostArn;

    public Optional<Output<String>> outpostArn() {
        return Optional.ofNullable(this.outpostArn);
    }

    /**
     * Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
     * 
     */
    @Import(name="ownerAlias")
    private @Nullable Output<String> ownerAlias;

    /**
     * @return Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
     * 
     */
    public Optional<Output<String>> ownerAlias() {
        return Optional.ofNullable(this.ownerAlias);
    }

    /**
     * The AWS account ID of the EBS snapshot owner.
     * 
     */
    @Import(name="ownerId")
    private @Nullable Output<String> ownerId;

    /**
     * @return The AWS account ID of the EBS snapshot owner.
     * 
     */
    public Optional<Output<String>> ownerId() {
        return Optional.ofNullable(this.ownerId);
    }

    /**
     * Indicates whether to permanently restore an archived snapshot.
     * 
     */
    @Import(name="permanentRestore")
    private @Nullable Output<Boolean> permanentRestore;

    /**
     * @return Indicates whether to permanently restore an archived snapshot.
     * 
     */
    public Optional<Output<Boolean>> permanentRestore() {
        return Optional.ofNullable(this.permanentRestore);
    }

    /**
     * The name of the IAM Role the VM Import/Export service will assume. This role needs certain permissions. See https://docs.aws.amazon.com/vm-import/latest/userguide/vmie_prereqs.html#vmimport-role. Default: `vmimport`
     * 
     */
    @Import(name="roleName")
    private @Nullable Output<String> roleName;

    /**
     * @return The name of the IAM Role the VM Import/Export service will assume. This role needs certain permissions. See https://docs.aws.amazon.com/vm-import/latest/userguide/vmie_prereqs.html#vmimport-role. Default: `vmimport`
     * 
     */
    public Optional<Output<String>> roleName() {
        return Optional.ofNullable(this.roleName);
    }

    /**
     * The name of the storage tier. Valid values are `archive` and `standard`. Default value is `standard`.
     * 
     */
    @Import(name="storageTier")
    private @Nullable Output<String> storageTier;

    /**
     * @return The name of the storage tier. Valid values are `archive` and `standard`. Default value is `standard`.
     * 
     */
    public Optional<Output<String>> storageTier() {
        return Optional.ofNullable(this.storageTier);
    }

    /**
     * A map of tags to assign to the snapshot.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the snapshot.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    /**
     * Specifies the number of days for which to temporarily restore an archived snapshot. Required for temporary restores only. The snapshot will be automatically re-archived after this period.
     * 
     */
    @Import(name="temporaryRestoreDays")
    private @Nullable Output<Integer> temporaryRestoreDays;

    /**
     * @return Specifies the number of days for which to temporarily restore an archived snapshot. Required for temporary restores only. The snapshot will be automatically re-archived after this period.
     * 
     */
    public Optional<Output<Integer>> temporaryRestoreDays() {
        return Optional.ofNullable(this.temporaryRestoreDays);
    }

    @Import(name="volumeId")
    private @Nullable Output<String> volumeId;

    public Optional<Output<String>> volumeId() {
        return Optional.ofNullable(this.volumeId);
    }

    /**
     * The size of the drive in GiBs.
     * 
     */
    @Import(name="volumeSize")
    private @Nullable Output<Integer> volumeSize;

    /**
     * @return The size of the drive in GiBs.
     * 
     */
    public Optional<Output<Integer>> volumeSize() {
        return Optional.ofNullable(this.volumeSize);
    }

    private SnapshotImportState() {}

    private SnapshotImportState(SnapshotImportState $) {
        this.arn = $.arn;
        this.clientData = $.clientData;
        this.dataEncryptionKeyId = $.dataEncryptionKeyId;
        this.description = $.description;
        this.diskContainer = $.diskContainer;
        this.encrypted = $.encrypted;
        this.kmsKeyId = $.kmsKeyId;
        this.outpostArn = $.outpostArn;
        this.ownerAlias = $.ownerAlias;
        this.ownerId = $.ownerId;
        this.permanentRestore = $.permanentRestore;
        this.roleName = $.roleName;
        this.storageTier = $.storageTier;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.temporaryRestoreDays = $.temporaryRestoreDays;
        this.volumeId = $.volumeId;
        this.volumeSize = $.volumeSize;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SnapshotImportState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SnapshotImportState $;

        public Builder() {
            $ = new SnapshotImportState();
        }

        public Builder(SnapshotImportState defaults) {
            $ = new SnapshotImportState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn Amazon Resource Name (ARN) of the EBS Snapshot.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn Amazon Resource Name (ARN) of the EBS Snapshot.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param clientData The client-specific data. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder clientData(@Nullable Output<SnapshotImportClientDataArgs> clientData) {
            $.clientData = clientData;
            return this;
        }

        /**
         * @param clientData The client-specific data. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder clientData(SnapshotImportClientDataArgs clientData) {
            return clientData(Output.of(clientData));
        }

        /**
         * @param dataEncryptionKeyId The data encryption key identifier for the snapshot.
         * 
         * @return builder
         * 
         */
        public Builder dataEncryptionKeyId(@Nullable Output<String> dataEncryptionKeyId) {
            $.dataEncryptionKeyId = dataEncryptionKeyId;
            return this;
        }

        /**
         * @param dataEncryptionKeyId The data encryption key identifier for the snapshot.
         * 
         * @return builder
         * 
         */
        public Builder dataEncryptionKeyId(String dataEncryptionKeyId) {
            return dataEncryptionKeyId(Output.of(dataEncryptionKeyId));
        }

        /**
         * @param description The description string for the import snapshot task.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description string for the import snapshot task.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param diskContainer Information about the disk container. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder diskContainer(@Nullable Output<SnapshotImportDiskContainerArgs> diskContainer) {
            $.diskContainer = diskContainer;
            return this;
        }

        /**
         * @param diskContainer Information about the disk container. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder diskContainer(SnapshotImportDiskContainerArgs diskContainer) {
            return diskContainer(Output.of(diskContainer));
        }

        /**
         * @param encrypted Specifies whether the destination snapshot of the imported image should be encrypted. The default KMS key for EBS is used unless you specify a non-default KMS key using KmsKeyId.
         * 
         * @return builder
         * 
         */
        public Builder encrypted(@Nullable Output<Boolean> encrypted) {
            $.encrypted = encrypted;
            return this;
        }

        /**
         * @param encrypted Specifies whether the destination snapshot of the imported image should be encrypted. The default KMS key for EBS is used unless you specify a non-default KMS key using KmsKeyId.
         * 
         * @return builder
         * 
         */
        public Builder encrypted(Boolean encrypted) {
            return encrypted(Output.of(encrypted));
        }

        /**
         * @param kmsKeyId An identifier for the symmetric KMS key to use when creating the encrypted snapshot. This parameter is only required if you want to use a non-default KMS key; if this parameter is not specified, the default KMS key for EBS is used. If a KmsKeyId is specified, the Encrypted flag must also be set.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyId(@Nullable Output<String> kmsKeyId) {
            $.kmsKeyId = kmsKeyId;
            return this;
        }

        /**
         * @param kmsKeyId An identifier for the symmetric KMS key to use when creating the encrypted snapshot. This parameter is only required if you want to use a non-default KMS key; if this parameter is not specified, the default KMS key for EBS is used. If a KmsKeyId is specified, the Encrypted flag must also be set.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyId(String kmsKeyId) {
            return kmsKeyId(Output.of(kmsKeyId));
        }

        public Builder outpostArn(@Nullable Output<String> outpostArn) {
            $.outpostArn = outpostArn;
            return this;
        }

        public Builder outpostArn(String outpostArn) {
            return outpostArn(Output.of(outpostArn));
        }

        /**
         * @param ownerAlias Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
         * 
         * @return builder
         * 
         */
        public Builder ownerAlias(@Nullable Output<String> ownerAlias) {
            $.ownerAlias = ownerAlias;
            return this;
        }

        /**
         * @param ownerAlias Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
         * 
         * @return builder
         * 
         */
        public Builder ownerAlias(String ownerAlias) {
            return ownerAlias(Output.of(ownerAlias));
        }

        /**
         * @param ownerId The AWS account ID of the EBS snapshot owner.
         * 
         * @return builder
         * 
         */
        public Builder ownerId(@Nullable Output<String> ownerId) {
            $.ownerId = ownerId;
            return this;
        }

        /**
         * @param ownerId The AWS account ID of the EBS snapshot owner.
         * 
         * @return builder
         * 
         */
        public Builder ownerId(String ownerId) {
            return ownerId(Output.of(ownerId));
        }

        /**
         * @param permanentRestore Indicates whether to permanently restore an archived snapshot.
         * 
         * @return builder
         * 
         */
        public Builder permanentRestore(@Nullable Output<Boolean> permanentRestore) {
            $.permanentRestore = permanentRestore;
            return this;
        }

        /**
         * @param permanentRestore Indicates whether to permanently restore an archived snapshot.
         * 
         * @return builder
         * 
         */
        public Builder permanentRestore(Boolean permanentRestore) {
            return permanentRestore(Output.of(permanentRestore));
        }

        /**
         * @param roleName The name of the IAM Role the VM Import/Export service will assume. This role needs certain permissions. See https://docs.aws.amazon.com/vm-import/latest/userguide/vmie_prereqs.html#vmimport-role. Default: `vmimport`
         * 
         * @return builder
         * 
         */
        public Builder roleName(@Nullable Output<String> roleName) {
            $.roleName = roleName;
            return this;
        }

        /**
         * @param roleName The name of the IAM Role the VM Import/Export service will assume. This role needs certain permissions. See https://docs.aws.amazon.com/vm-import/latest/userguide/vmie_prereqs.html#vmimport-role. Default: `vmimport`
         * 
         * @return builder
         * 
         */
        public Builder roleName(String roleName) {
            return roleName(Output.of(roleName));
        }

        /**
         * @param storageTier The name of the storage tier. Valid values are `archive` and `standard`. Default value is `standard`.
         * 
         * @return builder
         * 
         */
        public Builder storageTier(@Nullable Output<String> storageTier) {
            $.storageTier = storageTier;
            return this;
        }

        /**
         * @param storageTier The name of the storage tier. Valid values are `archive` and `standard`. Default value is `standard`.
         * 
         * @return builder
         * 
         */
        public Builder storageTier(String storageTier) {
            return storageTier(Output.of(storageTier));
        }

        /**
         * @param tags A map of tags to assign to the snapshot.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags A map of tags to assign to the snapshot.
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
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @param tagsAll A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
         * 
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        /**
         * @param temporaryRestoreDays Specifies the number of days for which to temporarily restore an archived snapshot. Required for temporary restores only. The snapshot will be automatically re-archived after this period.
         * 
         * @return builder
         * 
         */
        public Builder temporaryRestoreDays(@Nullable Output<Integer> temporaryRestoreDays) {
            $.temporaryRestoreDays = temporaryRestoreDays;
            return this;
        }

        /**
         * @param temporaryRestoreDays Specifies the number of days for which to temporarily restore an archived snapshot. Required for temporary restores only. The snapshot will be automatically re-archived after this period.
         * 
         * @return builder
         * 
         */
        public Builder temporaryRestoreDays(Integer temporaryRestoreDays) {
            return temporaryRestoreDays(Output.of(temporaryRestoreDays));
        }

        public Builder volumeId(@Nullable Output<String> volumeId) {
            $.volumeId = volumeId;
            return this;
        }

        public Builder volumeId(String volumeId) {
            return volumeId(Output.of(volumeId));
        }

        /**
         * @param volumeSize The size of the drive in GiBs.
         * 
         * @return builder
         * 
         */
        public Builder volumeSize(@Nullable Output<Integer> volumeSize) {
            $.volumeSize = volumeSize;
            return this;
        }

        /**
         * @param volumeSize The size of the drive in GiBs.
         * 
         * @return builder
         * 
         */
        public Builder volumeSize(Integer volumeSize) {
            return volumeSize(Output.of(volumeSize));
        }

        public SnapshotImportState build() {
            return $;
        }
    }

}
