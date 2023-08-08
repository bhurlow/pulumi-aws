// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.datasync.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TaskOptionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final TaskOptionsArgs Empty = new TaskOptionsArgs();

    /**
     * A file metadata that shows the last time a file was accessed (that is when the file was read or written to). If set to `BEST_EFFORT`, the DataSync Task attempts to preserve the original (that is, the version before sync `PREPARING` phase) `atime` attribute on all source files. Valid values: `BEST_EFFORT`, `NONE`. Default: `BEST_EFFORT`.
     * 
     */
    @Import(name="atime")
    private @Nullable Output<String> atime;

    /**
     * @return A file metadata that shows the last time a file was accessed (that is when the file was read or written to). If set to `BEST_EFFORT`, the DataSync Task attempts to preserve the original (that is, the version before sync `PREPARING` phase) `atime` attribute on all source files. Valid values: `BEST_EFFORT`, `NONE`. Default: `BEST_EFFORT`.
     * 
     */
    public Optional<Output<String>> atime() {
        return Optional.ofNullable(this.atime);
    }

    /**
     * Limits the bandwidth utilized. For example, to set a maximum of 1 MB, set this value to `1048576`. Value values: `-1` or greater. Default: `-1` (unlimited).
     * 
     */
    @Import(name="bytesPerSecond")
    private @Nullable Output<Integer> bytesPerSecond;

    /**
     * @return Limits the bandwidth utilized. For example, to set a maximum of 1 MB, set this value to `1048576`. Value values: `-1` or greater. Default: `-1` (unlimited).
     * 
     */
    public Optional<Output<Integer>> bytesPerSecond() {
        return Optional.ofNullable(this.bytesPerSecond);
    }

    /**
     * Group identifier of the file&#39;s owners. Valid values: `BOTH`, `INT_VALUE`, `NAME`, `NONE`. Default: `INT_VALUE` (preserve integer value of the ID).
     * 
     */
    @Import(name="gid")
    private @Nullable Output<String> gid;

    /**
     * @return Group identifier of the file&#39;s owners. Valid values: `BOTH`, `INT_VALUE`, `NAME`, `NONE`. Default: `INT_VALUE` (preserve integer value of the ID).
     * 
     */
    public Optional<Output<String>> gid() {
        return Optional.ofNullable(this.gid);
    }

    /**
     * Determines the type of logs that DataSync publishes to a log stream in the Amazon CloudWatch log group that you provide. Valid values: `OFF`, `BASIC`, `TRANSFER`. Default: `OFF`.
     * 
     */
    @Import(name="logLevel")
    private @Nullable Output<String> logLevel;

    /**
     * @return Determines the type of logs that DataSync publishes to a log stream in the Amazon CloudWatch log group that you provide. Valid values: `OFF`, `BASIC`, `TRANSFER`. Default: `OFF`.
     * 
     */
    public Optional<Output<String>> logLevel() {
        return Optional.ofNullable(this.logLevel);
    }

    /**
     * A file metadata that indicates the last time a file was modified (written to) before the sync `PREPARING` phase. Value values: `NONE`, `PRESERVE`. Default: `PRESERVE`.
     * 
     */
    @Import(name="mtime")
    private @Nullable Output<String> mtime;

    /**
     * @return A file metadata that indicates the last time a file was modified (written to) before the sync `PREPARING` phase. Value values: `NONE`, `PRESERVE`. Default: `PRESERVE`.
     * 
     */
    public Optional<Output<String>> mtime() {
        return Optional.ofNullable(this.mtime);
    }

    /**
     * Specifies whether object tags are maintained when transferring between object storage systems. If you want your DataSync task to ignore object tags, specify the NONE value. Valid values: `PRESERVE`, `NONE`. Default value: `PRESERVE`.
     * 
     */
    @Import(name="objectTags")
    private @Nullable Output<String> objectTags;

    /**
     * @return Specifies whether object tags are maintained when transferring between object storage systems. If you want your DataSync task to ignore object tags, specify the NONE value. Valid values: `PRESERVE`, `NONE`. Default value: `PRESERVE`.
     * 
     */
    public Optional<Output<String>> objectTags() {
        return Optional.ofNullable(this.objectTags);
    }

    /**
     * Determines whether files at the destination should be overwritten or preserved when copying files. Valid values: `ALWAYS`, `NEVER`. Default: `ALWAYS`.
     * 
     */
    @Import(name="overwriteMode")
    private @Nullable Output<String> overwriteMode;

    /**
     * @return Determines whether files at the destination should be overwritten or preserved when copying files. Valid values: `ALWAYS`, `NEVER`. Default: `ALWAYS`.
     * 
     */
    public Optional<Output<String>> overwriteMode() {
        return Optional.ofNullable(this.overwriteMode);
    }

    /**
     * Determines which users or groups can access a file for a specific purpose such as reading, writing, or execution of the file. Valid values: `NONE`, `PRESERVE`. Default: `PRESERVE`.
     * 
     */
    @Import(name="posixPermissions")
    private @Nullable Output<String> posixPermissions;

    /**
     * @return Determines which users or groups can access a file for a specific purpose such as reading, writing, or execution of the file. Valid values: `NONE`, `PRESERVE`. Default: `PRESERVE`.
     * 
     */
    public Optional<Output<String>> posixPermissions() {
        return Optional.ofNullable(this.posixPermissions);
    }

    /**
     * Whether files deleted in the source should be removed or preserved in the destination file system. Valid values: `PRESERVE`, `REMOVE`. Default: `PRESERVE`.
     * 
     */
    @Import(name="preserveDeletedFiles")
    private @Nullable Output<String> preserveDeletedFiles;

    /**
     * @return Whether files deleted in the source should be removed or preserved in the destination file system. Valid values: `PRESERVE`, `REMOVE`. Default: `PRESERVE`.
     * 
     */
    public Optional<Output<String>> preserveDeletedFiles() {
        return Optional.ofNullable(this.preserveDeletedFiles);
    }

    /**
     * Whether the DataSync Task should preserve the metadata of block and character devices in the source files system, and recreate the files with that device name and metadata on the destination. The DataSync Task can’t sync the actual contents of such devices, because many of the devices are non-terminal and don’t return an end of file (EOF) marker. Valid values: `NONE`, `PRESERVE`. Default: `NONE` (ignore special devices).
     * 
     */
    @Import(name="preserveDevices")
    private @Nullable Output<String> preserveDevices;

    /**
     * @return Whether the DataSync Task should preserve the metadata of block and character devices in the source files system, and recreate the files with that device name and metadata on the destination. The DataSync Task can’t sync the actual contents of such devices, because many of the devices are non-terminal and don’t return an end of file (EOF) marker. Valid values: `NONE`, `PRESERVE`. Default: `NONE` (ignore special devices).
     * 
     */
    public Optional<Output<String>> preserveDevices() {
        return Optional.ofNullable(this.preserveDevices);
    }

    /**
     * Determines which components of the SMB security descriptor are copied from source to destination objects. This value is only used for transfers between SMB and Amazon FSx for Windows File Server locations, or between two Amazon FSx for Windows File Server locations. Valid values: `NONE`, `OWNER_DACL`, `OWNER_DACL_SACL`. Default: `OWNER_DACL`.
     * 
     */
    @Import(name="securityDescriptorCopyFlags")
    private @Nullable Output<String> securityDescriptorCopyFlags;

    /**
     * @return Determines which components of the SMB security descriptor are copied from source to destination objects. This value is only used for transfers between SMB and Amazon FSx for Windows File Server locations, or between two Amazon FSx for Windows File Server locations. Valid values: `NONE`, `OWNER_DACL`, `OWNER_DACL_SACL`. Default: `OWNER_DACL`.
     * 
     */
    public Optional<Output<String>> securityDescriptorCopyFlags() {
        return Optional.ofNullable(this.securityDescriptorCopyFlags);
    }

    /**
     * Determines whether tasks should be queued before executing the tasks. Valid values: `ENABLED`, `DISABLED`. Default `ENABLED`.
     * 
     */
    @Import(name="taskQueueing")
    private @Nullable Output<String> taskQueueing;

    /**
     * @return Determines whether tasks should be queued before executing the tasks. Valid values: `ENABLED`, `DISABLED`. Default `ENABLED`.
     * 
     */
    public Optional<Output<String>> taskQueueing() {
        return Optional.ofNullable(this.taskQueueing);
    }

    /**
     * Determines whether DataSync transfers only the data and metadata that differ between the source and the destination location, or whether DataSync transfers all the content from the source, without comparing to the destination location. Valid values: `CHANGED`, `ALL`. Default: `CHANGED`
     * 
     */
    @Import(name="transferMode")
    private @Nullable Output<String> transferMode;

    /**
     * @return Determines whether DataSync transfers only the data and metadata that differ between the source and the destination location, or whether DataSync transfers all the content from the source, without comparing to the destination location. Valid values: `CHANGED`, `ALL`. Default: `CHANGED`
     * 
     */
    public Optional<Output<String>> transferMode() {
        return Optional.ofNullable(this.transferMode);
    }

    /**
     * User identifier of the file&#39;s owners. Valid values: `BOTH`, `INT_VALUE`, `NAME`, `NONE`. Default: `INT_VALUE` (preserve integer value of the ID).
     * 
     */
    @Import(name="uid")
    private @Nullable Output<String> uid;

    /**
     * @return User identifier of the file&#39;s owners. Valid values: `BOTH`, `INT_VALUE`, `NAME`, `NONE`. Default: `INT_VALUE` (preserve integer value of the ID).
     * 
     */
    public Optional<Output<String>> uid() {
        return Optional.ofNullable(this.uid);
    }

    /**
     * Whether a data integrity verification should be performed at the end of a task execution after all data and metadata have been transferred. Valid values: `NONE`, `POINT_IN_TIME_CONSISTENT`, `ONLY_FILES_TRANSFERRED`. Default: `POINT_IN_TIME_CONSISTENT`.
     * 
     */
    @Import(name="verifyMode")
    private @Nullable Output<String> verifyMode;

    /**
     * @return Whether a data integrity verification should be performed at the end of a task execution after all data and metadata have been transferred. Valid values: `NONE`, `POINT_IN_TIME_CONSISTENT`, `ONLY_FILES_TRANSFERRED`. Default: `POINT_IN_TIME_CONSISTENT`.
     * 
     */
    public Optional<Output<String>> verifyMode() {
        return Optional.ofNullable(this.verifyMode);
    }

    private TaskOptionsArgs() {}

    private TaskOptionsArgs(TaskOptionsArgs $) {
        this.atime = $.atime;
        this.bytesPerSecond = $.bytesPerSecond;
        this.gid = $.gid;
        this.logLevel = $.logLevel;
        this.mtime = $.mtime;
        this.objectTags = $.objectTags;
        this.overwriteMode = $.overwriteMode;
        this.posixPermissions = $.posixPermissions;
        this.preserveDeletedFiles = $.preserveDeletedFiles;
        this.preserveDevices = $.preserveDevices;
        this.securityDescriptorCopyFlags = $.securityDescriptorCopyFlags;
        this.taskQueueing = $.taskQueueing;
        this.transferMode = $.transferMode;
        this.uid = $.uid;
        this.verifyMode = $.verifyMode;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TaskOptionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TaskOptionsArgs $;

        public Builder() {
            $ = new TaskOptionsArgs();
        }

        public Builder(TaskOptionsArgs defaults) {
            $ = new TaskOptionsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param atime A file metadata that shows the last time a file was accessed (that is when the file was read or written to). If set to `BEST_EFFORT`, the DataSync Task attempts to preserve the original (that is, the version before sync `PREPARING` phase) `atime` attribute on all source files. Valid values: `BEST_EFFORT`, `NONE`. Default: `BEST_EFFORT`.
         * 
         * @return builder
         * 
         */
        public Builder atime(@Nullable Output<String> atime) {
            $.atime = atime;
            return this;
        }

        /**
         * @param atime A file metadata that shows the last time a file was accessed (that is when the file was read or written to). If set to `BEST_EFFORT`, the DataSync Task attempts to preserve the original (that is, the version before sync `PREPARING` phase) `atime` attribute on all source files. Valid values: `BEST_EFFORT`, `NONE`. Default: `BEST_EFFORT`.
         * 
         * @return builder
         * 
         */
        public Builder atime(String atime) {
            return atime(Output.of(atime));
        }

        /**
         * @param bytesPerSecond Limits the bandwidth utilized. For example, to set a maximum of 1 MB, set this value to `1048576`. Value values: `-1` or greater. Default: `-1` (unlimited).
         * 
         * @return builder
         * 
         */
        public Builder bytesPerSecond(@Nullable Output<Integer> bytesPerSecond) {
            $.bytesPerSecond = bytesPerSecond;
            return this;
        }

        /**
         * @param bytesPerSecond Limits the bandwidth utilized. For example, to set a maximum of 1 MB, set this value to `1048576`. Value values: `-1` or greater. Default: `-1` (unlimited).
         * 
         * @return builder
         * 
         */
        public Builder bytesPerSecond(Integer bytesPerSecond) {
            return bytesPerSecond(Output.of(bytesPerSecond));
        }

        /**
         * @param gid Group identifier of the file&#39;s owners. Valid values: `BOTH`, `INT_VALUE`, `NAME`, `NONE`. Default: `INT_VALUE` (preserve integer value of the ID).
         * 
         * @return builder
         * 
         */
        public Builder gid(@Nullable Output<String> gid) {
            $.gid = gid;
            return this;
        }

        /**
         * @param gid Group identifier of the file&#39;s owners. Valid values: `BOTH`, `INT_VALUE`, `NAME`, `NONE`. Default: `INT_VALUE` (preserve integer value of the ID).
         * 
         * @return builder
         * 
         */
        public Builder gid(String gid) {
            return gid(Output.of(gid));
        }

        /**
         * @param logLevel Determines the type of logs that DataSync publishes to a log stream in the Amazon CloudWatch log group that you provide. Valid values: `OFF`, `BASIC`, `TRANSFER`. Default: `OFF`.
         * 
         * @return builder
         * 
         */
        public Builder logLevel(@Nullable Output<String> logLevel) {
            $.logLevel = logLevel;
            return this;
        }

        /**
         * @param logLevel Determines the type of logs that DataSync publishes to a log stream in the Amazon CloudWatch log group that you provide. Valid values: `OFF`, `BASIC`, `TRANSFER`. Default: `OFF`.
         * 
         * @return builder
         * 
         */
        public Builder logLevel(String logLevel) {
            return logLevel(Output.of(logLevel));
        }

        /**
         * @param mtime A file metadata that indicates the last time a file was modified (written to) before the sync `PREPARING` phase. Value values: `NONE`, `PRESERVE`. Default: `PRESERVE`.
         * 
         * @return builder
         * 
         */
        public Builder mtime(@Nullable Output<String> mtime) {
            $.mtime = mtime;
            return this;
        }

        /**
         * @param mtime A file metadata that indicates the last time a file was modified (written to) before the sync `PREPARING` phase. Value values: `NONE`, `PRESERVE`. Default: `PRESERVE`.
         * 
         * @return builder
         * 
         */
        public Builder mtime(String mtime) {
            return mtime(Output.of(mtime));
        }

        /**
         * @param objectTags Specifies whether object tags are maintained when transferring between object storage systems. If you want your DataSync task to ignore object tags, specify the NONE value. Valid values: `PRESERVE`, `NONE`. Default value: `PRESERVE`.
         * 
         * @return builder
         * 
         */
        public Builder objectTags(@Nullable Output<String> objectTags) {
            $.objectTags = objectTags;
            return this;
        }

        /**
         * @param objectTags Specifies whether object tags are maintained when transferring between object storage systems. If you want your DataSync task to ignore object tags, specify the NONE value. Valid values: `PRESERVE`, `NONE`. Default value: `PRESERVE`.
         * 
         * @return builder
         * 
         */
        public Builder objectTags(String objectTags) {
            return objectTags(Output.of(objectTags));
        }

        /**
         * @param overwriteMode Determines whether files at the destination should be overwritten or preserved when copying files. Valid values: `ALWAYS`, `NEVER`. Default: `ALWAYS`.
         * 
         * @return builder
         * 
         */
        public Builder overwriteMode(@Nullable Output<String> overwriteMode) {
            $.overwriteMode = overwriteMode;
            return this;
        }

        /**
         * @param overwriteMode Determines whether files at the destination should be overwritten or preserved when copying files. Valid values: `ALWAYS`, `NEVER`. Default: `ALWAYS`.
         * 
         * @return builder
         * 
         */
        public Builder overwriteMode(String overwriteMode) {
            return overwriteMode(Output.of(overwriteMode));
        }

        /**
         * @param posixPermissions Determines which users or groups can access a file for a specific purpose such as reading, writing, or execution of the file. Valid values: `NONE`, `PRESERVE`. Default: `PRESERVE`.
         * 
         * @return builder
         * 
         */
        public Builder posixPermissions(@Nullable Output<String> posixPermissions) {
            $.posixPermissions = posixPermissions;
            return this;
        }

        /**
         * @param posixPermissions Determines which users or groups can access a file for a specific purpose such as reading, writing, or execution of the file. Valid values: `NONE`, `PRESERVE`. Default: `PRESERVE`.
         * 
         * @return builder
         * 
         */
        public Builder posixPermissions(String posixPermissions) {
            return posixPermissions(Output.of(posixPermissions));
        }

        /**
         * @param preserveDeletedFiles Whether files deleted in the source should be removed or preserved in the destination file system. Valid values: `PRESERVE`, `REMOVE`. Default: `PRESERVE`.
         * 
         * @return builder
         * 
         */
        public Builder preserveDeletedFiles(@Nullable Output<String> preserveDeletedFiles) {
            $.preserveDeletedFiles = preserveDeletedFiles;
            return this;
        }

        /**
         * @param preserveDeletedFiles Whether files deleted in the source should be removed or preserved in the destination file system. Valid values: `PRESERVE`, `REMOVE`. Default: `PRESERVE`.
         * 
         * @return builder
         * 
         */
        public Builder preserveDeletedFiles(String preserveDeletedFiles) {
            return preserveDeletedFiles(Output.of(preserveDeletedFiles));
        }

        /**
         * @param preserveDevices Whether the DataSync Task should preserve the metadata of block and character devices in the source files system, and recreate the files with that device name and metadata on the destination. The DataSync Task can’t sync the actual contents of such devices, because many of the devices are non-terminal and don’t return an end of file (EOF) marker. Valid values: `NONE`, `PRESERVE`. Default: `NONE` (ignore special devices).
         * 
         * @return builder
         * 
         */
        public Builder preserveDevices(@Nullable Output<String> preserveDevices) {
            $.preserveDevices = preserveDevices;
            return this;
        }

        /**
         * @param preserveDevices Whether the DataSync Task should preserve the metadata of block and character devices in the source files system, and recreate the files with that device name and metadata on the destination. The DataSync Task can’t sync the actual contents of such devices, because many of the devices are non-terminal and don’t return an end of file (EOF) marker. Valid values: `NONE`, `PRESERVE`. Default: `NONE` (ignore special devices).
         * 
         * @return builder
         * 
         */
        public Builder preserveDevices(String preserveDevices) {
            return preserveDevices(Output.of(preserveDevices));
        }

        /**
         * @param securityDescriptorCopyFlags Determines which components of the SMB security descriptor are copied from source to destination objects. This value is only used for transfers between SMB and Amazon FSx for Windows File Server locations, or between two Amazon FSx for Windows File Server locations. Valid values: `NONE`, `OWNER_DACL`, `OWNER_DACL_SACL`. Default: `OWNER_DACL`.
         * 
         * @return builder
         * 
         */
        public Builder securityDescriptorCopyFlags(@Nullable Output<String> securityDescriptorCopyFlags) {
            $.securityDescriptorCopyFlags = securityDescriptorCopyFlags;
            return this;
        }

        /**
         * @param securityDescriptorCopyFlags Determines which components of the SMB security descriptor are copied from source to destination objects. This value is only used for transfers between SMB and Amazon FSx for Windows File Server locations, or between two Amazon FSx for Windows File Server locations. Valid values: `NONE`, `OWNER_DACL`, `OWNER_DACL_SACL`. Default: `OWNER_DACL`.
         * 
         * @return builder
         * 
         */
        public Builder securityDescriptorCopyFlags(String securityDescriptorCopyFlags) {
            return securityDescriptorCopyFlags(Output.of(securityDescriptorCopyFlags));
        }

        /**
         * @param taskQueueing Determines whether tasks should be queued before executing the tasks. Valid values: `ENABLED`, `DISABLED`. Default `ENABLED`.
         * 
         * @return builder
         * 
         */
        public Builder taskQueueing(@Nullable Output<String> taskQueueing) {
            $.taskQueueing = taskQueueing;
            return this;
        }

        /**
         * @param taskQueueing Determines whether tasks should be queued before executing the tasks. Valid values: `ENABLED`, `DISABLED`. Default `ENABLED`.
         * 
         * @return builder
         * 
         */
        public Builder taskQueueing(String taskQueueing) {
            return taskQueueing(Output.of(taskQueueing));
        }

        /**
         * @param transferMode Determines whether DataSync transfers only the data and metadata that differ between the source and the destination location, or whether DataSync transfers all the content from the source, without comparing to the destination location. Valid values: `CHANGED`, `ALL`. Default: `CHANGED`
         * 
         * @return builder
         * 
         */
        public Builder transferMode(@Nullable Output<String> transferMode) {
            $.transferMode = transferMode;
            return this;
        }

        /**
         * @param transferMode Determines whether DataSync transfers only the data and metadata that differ between the source and the destination location, or whether DataSync transfers all the content from the source, without comparing to the destination location. Valid values: `CHANGED`, `ALL`. Default: `CHANGED`
         * 
         * @return builder
         * 
         */
        public Builder transferMode(String transferMode) {
            return transferMode(Output.of(transferMode));
        }

        /**
         * @param uid User identifier of the file&#39;s owners. Valid values: `BOTH`, `INT_VALUE`, `NAME`, `NONE`. Default: `INT_VALUE` (preserve integer value of the ID).
         * 
         * @return builder
         * 
         */
        public Builder uid(@Nullable Output<String> uid) {
            $.uid = uid;
            return this;
        }

        /**
         * @param uid User identifier of the file&#39;s owners. Valid values: `BOTH`, `INT_VALUE`, `NAME`, `NONE`. Default: `INT_VALUE` (preserve integer value of the ID).
         * 
         * @return builder
         * 
         */
        public Builder uid(String uid) {
            return uid(Output.of(uid));
        }

        /**
         * @param verifyMode Whether a data integrity verification should be performed at the end of a task execution after all data and metadata have been transferred. Valid values: `NONE`, `POINT_IN_TIME_CONSISTENT`, `ONLY_FILES_TRANSFERRED`. Default: `POINT_IN_TIME_CONSISTENT`.
         * 
         * @return builder
         * 
         */
        public Builder verifyMode(@Nullable Output<String> verifyMode) {
            $.verifyMode = verifyMode;
            return this;
        }

        /**
         * @param verifyMode Whether a data integrity verification should be performed at the end of a task execution after all data and metadata have been transferred. Valid values: `NONE`, `POINT_IN_TIME_CONSISTENT`, `ONLY_FILES_TRANSFERRED`. Default: `POINT_IN_TIME_CONSISTENT`.
         * 
         * @return builder
         * 
         */
        public Builder verifyMode(String verifyMode) {
            return verifyMode(Output.of(verifyMode));
        }

        public TaskOptionsArgs build() {
            return $;
        }
    }

}
