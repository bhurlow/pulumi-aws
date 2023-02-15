// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.dynamodb.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TableReplicaArgs extends com.pulumi.resources.ResourceArgs {

    public static final TableReplicaArgs Empty = new TableReplicaArgs();

    /**
     * ARN of the table
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return ARN of the table
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, `alias/aws/dynamodb`. **Note:** This attribute will _not_ be populated with the ARN of _default_ keys.
     * 
     */
    @Import(name="kmsKeyArn")
    private @Nullable Output<String> kmsKeyArn;

    /**
     * @return ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, `alias/aws/dynamodb`. **Note:** This attribute will _not_ be populated with the ARN of _default_ keys.
     * 
     */
    public Optional<Output<String>> kmsKeyArn() {
        return Optional.ofNullable(this.kmsKeyArn);
    }

    /**
     * Whether to enable Point In Time Recovery for the replica. Default is `false`.
     * 
     */
    @Import(name="pointInTimeRecovery")
    private @Nullable Output<Boolean> pointInTimeRecovery;

    /**
     * @return Whether to enable Point In Time Recovery for the replica. Default is `false`.
     * 
     */
    public Optional<Output<Boolean>> pointInTimeRecovery() {
        return Optional.ofNullable(this.pointInTimeRecovery);
    }

    /**
     * Whether to propagate the global table&#39;s tags to a replica. Default is `false`. Changes to tags only move in one direction: from global (source) to replica. In other words, tag drift on a replica will not trigger an update. Tag or replica changes on the global table, whether from drift or configuration changes, are propagated to replicas. Changing from `true` to `false` on a subsequent `apply` means replica tags are left as they were, unmanaged, not deleted.
     * 
     */
    @Import(name="propagateTags")
    private @Nullable Output<Boolean> propagateTags;

    /**
     * @return Whether to propagate the global table&#39;s tags to a replica. Default is `false`. Changes to tags only move in one direction: from global (source) to replica. In other words, tag drift on a replica will not trigger an update. Tag or replica changes on the global table, whether from drift or configuration changes, are propagated to replicas. Changing from `true` to `false` on a subsequent `apply` means replica tags are left as they were, unmanaged, not deleted.
     * 
     */
    public Optional<Output<Boolean>> propagateTags() {
        return Optional.ofNullable(this.propagateTags);
    }

    /**
     * Region name of the replica.
     * 
     */
    @Import(name="regionName", required=true)
    private Output<String> regionName;

    /**
     * @return Region name of the replica.
     * 
     */
    public Output<String> regionName() {
        return this.regionName;
    }

    /**
     * ARN of the Table Stream. Only available when `stream_enabled = true`
     * 
     */
    @Import(name="streamArn")
    private @Nullable Output<String> streamArn;

    /**
     * @return ARN of the Table Stream. Only available when `stream_enabled = true`
     * 
     */
    public Optional<Output<String>> streamArn() {
        return Optional.ofNullable(this.streamArn);
    }

    /**
     * Timestamp, in ISO 8601 format, for this stream. Note that this timestamp is not a unique identifier for the stream on its own. However, the combination of AWS customer ID, table name and this field is guaranteed to be unique. It can be used for creating CloudWatch Alarms. Only available when `stream_enabled = true`.
     * 
     */
    @Import(name="streamLabel")
    private @Nullable Output<String> streamLabel;

    /**
     * @return Timestamp, in ISO 8601 format, for this stream. Note that this timestamp is not a unique identifier for the stream on its own. However, the combination of AWS customer ID, table name and this field is guaranteed to be unique. It can be used for creating CloudWatch Alarms. Only available when `stream_enabled = true`.
     * 
     */
    public Optional<Output<String>> streamLabel() {
        return Optional.ofNullable(this.streamLabel);
    }

    private TableReplicaArgs() {}

    private TableReplicaArgs(TableReplicaArgs $) {
        this.arn = $.arn;
        this.kmsKeyArn = $.kmsKeyArn;
        this.pointInTimeRecovery = $.pointInTimeRecovery;
        this.propagateTags = $.propagateTags;
        this.regionName = $.regionName;
        this.streamArn = $.streamArn;
        this.streamLabel = $.streamLabel;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TableReplicaArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TableReplicaArgs $;

        public Builder() {
            $ = new TableReplicaArgs();
        }

        public Builder(TableReplicaArgs defaults) {
            $ = new TableReplicaArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn ARN of the table
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn ARN of the table
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param kmsKeyArn ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, `alias/aws/dynamodb`. **Note:** This attribute will _not_ be populated with the ARN of _default_ keys.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyArn(@Nullable Output<String> kmsKeyArn) {
            $.kmsKeyArn = kmsKeyArn;
            return this;
        }

        /**
         * @param kmsKeyArn ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, `alias/aws/dynamodb`. **Note:** This attribute will _not_ be populated with the ARN of _default_ keys.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyArn(String kmsKeyArn) {
            return kmsKeyArn(Output.of(kmsKeyArn));
        }

        /**
         * @param pointInTimeRecovery Whether to enable Point In Time Recovery for the replica. Default is `false`.
         * 
         * @return builder
         * 
         */
        public Builder pointInTimeRecovery(@Nullable Output<Boolean> pointInTimeRecovery) {
            $.pointInTimeRecovery = pointInTimeRecovery;
            return this;
        }

        /**
         * @param pointInTimeRecovery Whether to enable Point In Time Recovery for the replica. Default is `false`.
         * 
         * @return builder
         * 
         */
        public Builder pointInTimeRecovery(Boolean pointInTimeRecovery) {
            return pointInTimeRecovery(Output.of(pointInTimeRecovery));
        }

        /**
         * @param propagateTags Whether to propagate the global table&#39;s tags to a replica. Default is `false`. Changes to tags only move in one direction: from global (source) to replica. In other words, tag drift on a replica will not trigger an update. Tag or replica changes on the global table, whether from drift or configuration changes, are propagated to replicas. Changing from `true` to `false` on a subsequent `apply` means replica tags are left as they were, unmanaged, not deleted.
         * 
         * @return builder
         * 
         */
        public Builder propagateTags(@Nullable Output<Boolean> propagateTags) {
            $.propagateTags = propagateTags;
            return this;
        }

        /**
         * @param propagateTags Whether to propagate the global table&#39;s tags to a replica. Default is `false`. Changes to tags only move in one direction: from global (source) to replica. In other words, tag drift on a replica will not trigger an update. Tag or replica changes on the global table, whether from drift or configuration changes, are propagated to replicas. Changing from `true` to `false` on a subsequent `apply` means replica tags are left as they were, unmanaged, not deleted.
         * 
         * @return builder
         * 
         */
        public Builder propagateTags(Boolean propagateTags) {
            return propagateTags(Output.of(propagateTags));
        }

        /**
         * @param regionName Region name of the replica.
         * 
         * @return builder
         * 
         */
        public Builder regionName(Output<String> regionName) {
            $.regionName = regionName;
            return this;
        }

        /**
         * @param regionName Region name of the replica.
         * 
         * @return builder
         * 
         */
        public Builder regionName(String regionName) {
            return regionName(Output.of(regionName));
        }

        /**
         * @param streamArn ARN of the Table Stream. Only available when `stream_enabled = true`
         * 
         * @return builder
         * 
         */
        public Builder streamArn(@Nullable Output<String> streamArn) {
            $.streamArn = streamArn;
            return this;
        }

        /**
         * @param streamArn ARN of the Table Stream. Only available when `stream_enabled = true`
         * 
         * @return builder
         * 
         */
        public Builder streamArn(String streamArn) {
            return streamArn(Output.of(streamArn));
        }

        /**
         * @param streamLabel Timestamp, in ISO 8601 format, for this stream. Note that this timestamp is not a unique identifier for the stream on its own. However, the combination of AWS customer ID, table name and this field is guaranteed to be unique. It can be used for creating CloudWatch Alarms. Only available when `stream_enabled = true`.
         * 
         * @return builder
         * 
         */
        public Builder streamLabel(@Nullable Output<String> streamLabel) {
            $.streamLabel = streamLabel;
            return this;
        }

        /**
         * @param streamLabel Timestamp, in ISO 8601 format, for this stream. Note that this timestamp is not a unique identifier for the stream on its own. However, the combination of AWS customer ID, table name and this field is guaranteed to be unique. It can be used for creating CloudWatch Alarms. Only available when `stream_enabled = true`.
         * 
         * @return builder
         * 
         */
        public Builder streamLabel(String streamLabel) {
            return streamLabel(Output.of(streamLabel));
        }

        public TableReplicaArgs build() {
            $.regionName = Objects.requireNonNull($.regionName, "expected parameter 'regionName' to be non-null");
            return $;
        }
    }

}
