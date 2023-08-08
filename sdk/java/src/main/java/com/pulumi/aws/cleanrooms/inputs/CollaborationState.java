// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cleanrooms.inputs;

import com.pulumi.aws.cleanrooms.inputs.CollaborationDataEncryptionMetadataArgs;
import com.pulumi.aws.cleanrooms.inputs.CollaborationMemberArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CollaborationState extends com.pulumi.resources.ResourceArgs {

    public static final CollaborationState Empty = new CollaborationState();

    /**
     * The arn of the collaboration
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The arn of the collaboration
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The date and time the collaboration was created
     * * `member status` - For each member included in the collaboration an additional computed attribute of status is added. These values [may be
     *   ound here](https://docs.aws.amazon.com/clean-rooms/latest/apireference/API_MemberSummary.html#API-Type-MemberSummary-status)
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return The date and time the collaboration was created
     * * `member status` - For each member included in the collaboration an additional computed attribute of status is added. These values [may be
     *   ound here](https://docs.aws.amazon.com/clean-rooms/latest/apireference/API_MemberSummary.html#API-Type-MemberSummary-status)
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * The name for the member record for the collaboration creator.
     * 
     */
    @Import(name="creatorDisplayName")
    private @Nullable Output<String> creatorDisplayName;

    /**
     * @return The name for the member record for the collaboration creator.
     * 
     */
    public Optional<Output<String>> creatorDisplayName() {
        return Optional.ofNullable(this.creatorDisplayName);
    }

    /**
     * The list of member abilities for the creator of the collaboration.  Valid v
     * lues [may be found here](https://docs.aws.amazon.com/clean-rooms/latest/apireference/API_CreateCollaboration.html#API-CreateCollaboration-re
     * uest-creatorMemberAbilities)
     * 
     */
    @Import(name="creatorMemberAbilities")
    private @Nullable Output<List<String>> creatorMemberAbilities;

    /**
     * @return The list of member abilities for the creator of the collaboration.  Valid v
     * lues [may be found here](https://docs.aws.amazon.com/clean-rooms/latest/apireference/API_CreateCollaboration.html#API-CreateCollaboration-re
     * uest-creatorMemberAbilities)
     * 
     */
    public Optional<Output<List<String>>> creatorMemberAbilities() {
        return Optional.ofNullable(this.creatorMemberAbilities);
    }

    /**
     * a collection of settings which determine how the [c3r client](https://docs
     * aws.amazon.com/clean-rooms/latest/userguide/crypto-computing.html) will encrypt data for use within this collaboration
     * * `data_encryption_metadata.allow_clear_text` - (Required - Forces new resource) - Indicates whether encrypted tables can contain cleartext data. This is a boolea
     *   field.
     * * `data_encryption_metadata.allow_duplicates` - (Required - Forces new resource ) - Indicates whether Fingerprint columns can contain duplicate entries. This is a
     *   boolean field.
     * * `data_encryption_metadata.allow_joins_on_columns_with_different_names` - (Required - Forces new resource) - Indicates whether Fingerprint columns can be joined
     *   n any other Fingerprint column with a different name. This is a boolean field.
     * * `data_encryption_metadata.preserve_nulls` - (Required - Forces new resource) - Indicates whether NULL values are to be copied as NULL to encrypted tables (true)
     *   or cryptographically processed (false).
     * 
     */
    @Import(name="dataEncryptionMetadata")
    private @Nullable Output<CollaborationDataEncryptionMetadataArgs> dataEncryptionMetadata;

    /**
     * @return a collection of settings which determine how the [c3r client](https://docs
     * aws.amazon.com/clean-rooms/latest/userguide/crypto-computing.html) will encrypt data for use within this collaboration
     * * `data_encryption_metadata.allow_clear_text` - (Required - Forces new resource) - Indicates whether encrypted tables can contain cleartext data. This is a boolea
     *   field.
     * * `data_encryption_metadata.allow_duplicates` - (Required - Forces new resource ) - Indicates whether Fingerprint columns can contain duplicate entries. This is a
     *   boolean field.
     * * `data_encryption_metadata.allow_joins_on_columns_with_different_names` - (Required - Forces new resource) - Indicates whether Fingerprint columns can be joined
     *   n any other Fingerprint column with a different name. This is a boolean field.
     * * `data_encryption_metadata.preserve_nulls` - (Required - Forces new resource) - Indicates whether NULL values are to be copied as NULL to encrypted tables (true)
     *   or cryptographically processed (false).
     * 
     */
    public Optional<Output<CollaborationDataEncryptionMetadataArgs>> dataEncryptionMetadata() {
        return Optional.ofNullable(this.dataEncryptionMetadata);
    }

    /**
     * A description for a collaboration.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description for a collaboration.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Additional members of the collaboration which will be invited to join the collaboration.
     * * `member.account_id` - (Required - Forces new resource) - The account id for the invited member
     * * `member.display_name` - (Required - Forces new resource) - The display name for the invited member
     * * `member.member_abilities` - (Required - Forces new resource) - The list of abilities for the invited member. Valid values [may be found here](https://docs.aws.amazon.com/clean-rooms/latest/apireference/API_CreateCollaboration.html#API-CreateCollaboration-request-creatorMemberAbiliti
     *   s
     * 
     */
    @Import(name="members")
    private @Nullable Output<List<CollaborationMemberArgs>> members;

    /**
     * @return Additional members of the collaboration which will be invited to join the collaboration.
     * * `member.account_id` - (Required - Forces new resource) - The account id for the invited member
     * * `member.display_name` - (Required - Forces new resource) - The display name for the invited member
     * * `member.member_abilities` - (Required - Forces new resource) - The list of abilities for the invited member. Valid values [may be found here](https://docs.aws.amazon.com/clean-rooms/latest/apireference/API_CreateCollaboration.html#API-CreateCollaboration-request-creatorMemberAbiliti
     *   s
     * 
     */
    public Optional<Output<List<CollaborationMemberArgs>>> members() {
        return Optional.ofNullable(this.members);
    }

    /**
     * The name of the collaboration.  Collaboration names do not need to be unique.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the collaboration.  Collaboration names do not need to be unique.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Determines if members of the collaboration can enable query logs within their own
     * emberships. Valid values [may be found here](https://docs.aws.amazon.com/clean-rooms/latest/apireference/API_CreateCollaboration.html#API-Cr
     * ateCollaboration-request-queryLogStatus).
     * 
     */
    @Import(name="queryLogStatus")
    private @Nullable Output<String> queryLogStatus;

    /**
     * @return Determines if members of the collaboration can enable query logs within their own
     * emberships. Valid values [may be found here](https://docs.aws.amazon.com/clean-rooms/latest/apireference/API_CreateCollaboration.html#API-Cr
     * ateCollaboration-request-queryLogStatus).
     * 
     */
    public Optional<Output<String>> queryLogStatus() {
        return Optional.ofNullable(this.queryLogStatus);
    }

    /**
     * Key value pairs which tag the collaboration.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key value pairs which tag the collaboration.
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    @Import(name="updateTime")
    private @Nullable Output<String> updateTime;

    public Optional<Output<String>> updateTime() {
        return Optional.ofNullable(this.updateTime);
    }

    private CollaborationState() {}

    private CollaborationState(CollaborationState $) {
        this.arn = $.arn;
        this.createTime = $.createTime;
        this.creatorDisplayName = $.creatorDisplayName;
        this.creatorMemberAbilities = $.creatorMemberAbilities;
        this.dataEncryptionMetadata = $.dataEncryptionMetadata;
        this.description = $.description;
        this.members = $.members;
        this.name = $.name;
        this.queryLogStatus = $.queryLogStatus;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.updateTime = $.updateTime;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CollaborationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CollaborationState $;

        public Builder() {
            $ = new CollaborationState();
        }

        public Builder(CollaborationState defaults) {
            $ = new CollaborationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The arn of the collaboration
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The arn of the collaboration
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param createTime The date and time the collaboration was created
         * * `member status` - For each member included in the collaboration an additional computed attribute of status is added. These values [may be
         *   ound here](https://docs.aws.amazon.com/clean-rooms/latest/apireference/API_MemberSummary.html#API-Type-MemberSummary-status)
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime The date and time the collaboration was created
         * * `member status` - For each member included in the collaboration an additional computed attribute of status is added. These values [may be
         *   ound here](https://docs.aws.amazon.com/clean-rooms/latest/apireference/API_MemberSummary.html#API-Type-MemberSummary-status)
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param creatorDisplayName The name for the member record for the collaboration creator.
         * 
         * @return builder
         * 
         */
        public Builder creatorDisplayName(@Nullable Output<String> creatorDisplayName) {
            $.creatorDisplayName = creatorDisplayName;
            return this;
        }

        /**
         * @param creatorDisplayName The name for the member record for the collaboration creator.
         * 
         * @return builder
         * 
         */
        public Builder creatorDisplayName(String creatorDisplayName) {
            return creatorDisplayName(Output.of(creatorDisplayName));
        }

        /**
         * @param creatorMemberAbilities The list of member abilities for the creator of the collaboration.  Valid v
         * lues [may be found here](https://docs.aws.amazon.com/clean-rooms/latest/apireference/API_CreateCollaboration.html#API-CreateCollaboration-re
         * uest-creatorMemberAbilities)
         * 
         * @return builder
         * 
         */
        public Builder creatorMemberAbilities(@Nullable Output<List<String>> creatorMemberAbilities) {
            $.creatorMemberAbilities = creatorMemberAbilities;
            return this;
        }

        /**
         * @param creatorMemberAbilities The list of member abilities for the creator of the collaboration.  Valid v
         * lues [may be found here](https://docs.aws.amazon.com/clean-rooms/latest/apireference/API_CreateCollaboration.html#API-CreateCollaboration-re
         * uest-creatorMemberAbilities)
         * 
         * @return builder
         * 
         */
        public Builder creatorMemberAbilities(List<String> creatorMemberAbilities) {
            return creatorMemberAbilities(Output.of(creatorMemberAbilities));
        }

        /**
         * @param creatorMemberAbilities The list of member abilities for the creator of the collaboration.  Valid v
         * lues [may be found here](https://docs.aws.amazon.com/clean-rooms/latest/apireference/API_CreateCollaboration.html#API-CreateCollaboration-re
         * uest-creatorMemberAbilities)
         * 
         * @return builder
         * 
         */
        public Builder creatorMemberAbilities(String... creatorMemberAbilities) {
            return creatorMemberAbilities(List.of(creatorMemberAbilities));
        }

        /**
         * @param dataEncryptionMetadata a collection of settings which determine how the [c3r client](https://docs
         * aws.amazon.com/clean-rooms/latest/userguide/crypto-computing.html) will encrypt data for use within this collaboration
         * * `data_encryption_metadata.allow_clear_text` - (Required - Forces new resource) - Indicates whether encrypted tables can contain cleartext data. This is a boolea
         *   field.
         * * `data_encryption_metadata.allow_duplicates` - (Required - Forces new resource ) - Indicates whether Fingerprint columns can contain duplicate entries. This is a
         *   boolean field.
         * * `data_encryption_metadata.allow_joins_on_columns_with_different_names` - (Required - Forces new resource) - Indicates whether Fingerprint columns can be joined
         *   n any other Fingerprint column with a different name. This is a boolean field.
         * * `data_encryption_metadata.preserve_nulls` - (Required - Forces new resource) - Indicates whether NULL values are to be copied as NULL to encrypted tables (true)
         *   or cryptographically processed (false).
         * 
         * @return builder
         * 
         */
        public Builder dataEncryptionMetadata(@Nullable Output<CollaborationDataEncryptionMetadataArgs> dataEncryptionMetadata) {
            $.dataEncryptionMetadata = dataEncryptionMetadata;
            return this;
        }

        /**
         * @param dataEncryptionMetadata a collection of settings which determine how the [c3r client](https://docs
         * aws.amazon.com/clean-rooms/latest/userguide/crypto-computing.html) will encrypt data for use within this collaboration
         * * `data_encryption_metadata.allow_clear_text` - (Required - Forces new resource) - Indicates whether encrypted tables can contain cleartext data. This is a boolea
         *   field.
         * * `data_encryption_metadata.allow_duplicates` - (Required - Forces new resource ) - Indicates whether Fingerprint columns can contain duplicate entries. This is a
         *   boolean field.
         * * `data_encryption_metadata.allow_joins_on_columns_with_different_names` - (Required - Forces new resource) - Indicates whether Fingerprint columns can be joined
         *   n any other Fingerprint column with a different name. This is a boolean field.
         * * `data_encryption_metadata.preserve_nulls` - (Required - Forces new resource) - Indicates whether NULL values are to be copied as NULL to encrypted tables (true)
         *   or cryptographically processed (false).
         * 
         * @return builder
         * 
         */
        public Builder dataEncryptionMetadata(CollaborationDataEncryptionMetadataArgs dataEncryptionMetadata) {
            return dataEncryptionMetadata(Output.of(dataEncryptionMetadata));
        }

        /**
         * @param description A description for a collaboration.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description for a collaboration.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param members Additional members of the collaboration which will be invited to join the collaboration.
         * * `member.account_id` - (Required - Forces new resource) - The account id for the invited member
         * * `member.display_name` - (Required - Forces new resource) - The display name for the invited member
         * * `member.member_abilities` - (Required - Forces new resource) - The list of abilities for the invited member. Valid values [may be found here](https://docs.aws.amazon.com/clean-rooms/latest/apireference/API_CreateCollaboration.html#API-CreateCollaboration-request-creatorMemberAbiliti
         *   s
         * 
         * @return builder
         * 
         */
        public Builder members(@Nullable Output<List<CollaborationMemberArgs>> members) {
            $.members = members;
            return this;
        }

        /**
         * @param members Additional members of the collaboration which will be invited to join the collaboration.
         * * `member.account_id` - (Required - Forces new resource) - The account id for the invited member
         * * `member.display_name` - (Required - Forces new resource) - The display name for the invited member
         * * `member.member_abilities` - (Required - Forces new resource) - The list of abilities for the invited member. Valid values [may be found here](https://docs.aws.amazon.com/clean-rooms/latest/apireference/API_CreateCollaboration.html#API-CreateCollaboration-request-creatorMemberAbiliti
         *   s
         * 
         * @return builder
         * 
         */
        public Builder members(List<CollaborationMemberArgs> members) {
            return members(Output.of(members));
        }

        /**
         * @param members Additional members of the collaboration which will be invited to join the collaboration.
         * * `member.account_id` - (Required - Forces new resource) - The account id for the invited member
         * * `member.display_name` - (Required - Forces new resource) - The display name for the invited member
         * * `member.member_abilities` - (Required - Forces new resource) - The list of abilities for the invited member. Valid values [may be found here](https://docs.aws.amazon.com/clean-rooms/latest/apireference/API_CreateCollaboration.html#API-CreateCollaboration-request-creatorMemberAbiliti
         *   s
         * 
         * @return builder
         * 
         */
        public Builder members(CollaborationMemberArgs... members) {
            return members(List.of(members));
        }

        /**
         * @param name The name of the collaboration.  Collaboration names do not need to be unique.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the collaboration.  Collaboration names do not need to be unique.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param queryLogStatus Determines if members of the collaboration can enable query logs within their own
         * emberships. Valid values [may be found here](https://docs.aws.amazon.com/clean-rooms/latest/apireference/API_CreateCollaboration.html#API-Cr
         * ateCollaboration-request-queryLogStatus).
         * 
         * @return builder
         * 
         */
        public Builder queryLogStatus(@Nullable Output<String> queryLogStatus) {
            $.queryLogStatus = queryLogStatus;
            return this;
        }

        /**
         * @param queryLogStatus Determines if members of the collaboration can enable query logs within their own
         * emberships. Valid values [may be found here](https://docs.aws.amazon.com/clean-rooms/latest/apireference/API_CreateCollaboration.html#API-Cr
         * ateCollaboration-request-queryLogStatus).
         * 
         * @return builder
         * 
         */
        public Builder queryLogStatus(String queryLogStatus) {
            return queryLogStatus(Output.of(queryLogStatus));
        }

        /**
         * @param tags Key value pairs which tag the collaboration.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key value pairs which tag the collaboration.
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        public Builder updateTime(@Nullable Output<String> updateTime) {
            $.updateTime = updateTime;
            return this;
        }

        public Builder updateTime(String updateTime) {
            return updateTime(Output.of(updateTime));
        }

        public CollaborationState build() {
            return $;
        }
    }

}
