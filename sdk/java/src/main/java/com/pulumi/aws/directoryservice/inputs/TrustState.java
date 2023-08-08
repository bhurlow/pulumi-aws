// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.directoryservice.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TrustState extends com.pulumi.resources.ResourceArgs {

    public static final TrustState Empty = new TrustState();

    /**
     * Set of IPv4 addresses for the DNS server associated with the remote Directory.
     * Can contain between 1 and 4 values.
     * 
     */
    @Import(name="conditionalForwarderIpAddrs")
    private @Nullable Output<List<String>> conditionalForwarderIpAddrs;

    /**
     * @return Set of IPv4 addresses for the DNS server associated with the remote Directory.
     * Can contain between 1 and 4 values.
     * 
     */
    public Optional<Output<List<String>>> conditionalForwarderIpAddrs() {
        return Optional.ofNullable(this.conditionalForwarderIpAddrs);
    }

    /**
     * Date and time when the Trust was created.
     * 
     */
    @Import(name="createdDateTime")
    private @Nullable Output<String> createdDateTime;

    /**
     * @return Date and time when the Trust was created.
     * 
     */
    public Optional<Output<String>> createdDateTime() {
        return Optional.ofNullable(this.createdDateTime);
    }

    /**
     * Whether to delete the conditional forwarder when deleting the Trust relationship.
     * 
     */
    @Import(name="deleteAssociatedConditionalForwarder")
    private @Nullable Output<Boolean> deleteAssociatedConditionalForwarder;

    /**
     * @return Whether to delete the conditional forwarder when deleting the Trust relationship.
     * 
     */
    public Optional<Output<Boolean>> deleteAssociatedConditionalForwarder() {
        return Optional.ofNullable(this.deleteAssociatedConditionalForwarder);
    }

    /**
     * ID of the Directory.
     * 
     */
    @Import(name="directoryId")
    private @Nullable Output<String> directoryId;

    /**
     * @return ID of the Directory.
     * 
     */
    public Optional<Output<String>> directoryId() {
        return Optional.ofNullable(this.directoryId);
    }

    /**
     * Date and time when the Trust was last updated.
     * 
     */
    @Import(name="lastUpdatedDateTime")
    private @Nullable Output<String> lastUpdatedDateTime;

    /**
     * @return Date and time when the Trust was last updated.
     * 
     */
    public Optional<Output<String>> lastUpdatedDateTime() {
        return Optional.ofNullable(this.lastUpdatedDateTime);
    }

    /**
     * Fully qualified domain name of the remote Directory.
     * 
     */
    @Import(name="remoteDomainName")
    private @Nullable Output<String> remoteDomainName;

    /**
     * @return Fully qualified domain name of the remote Directory.
     * 
     */
    public Optional<Output<String>> remoteDomainName() {
        return Optional.ofNullable(this.remoteDomainName);
    }

    /**
     * Whether to enable selective authentication.
     * Valid values are `Enabled` and `Disabled`.
     * Default value is `Disabled`.
     * 
     */
    @Import(name="selectiveAuth")
    private @Nullable Output<String> selectiveAuth;

    /**
     * @return Whether to enable selective authentication.
     * Valid values are `Enabled` and `Disabled`.
     * Default value is `Disabled`.
     * 
     */
    public Optional<Output<String>> selectiveAuth() {
        return Optional.ofNullable(this.selectiveAuth);
    }

    /**
     * Date and time when the Trust state in `trust_state` was last updated.
     * 
     */
    @Import(name="stateLastUpdatedDateTime")
    private @Nullable Output<String> stateLastUpdatedDateTime;

    /**
     * @return Date and time when the Trust state in `trust_state` was last updated.
     * 
     */
    public Optional<Output<String>> stateLastUpdatedDateTime() {
        return Optional.ofNullable(this.stateLastUpdatedDateTime);
    }

    /**
     * The direction of the Trust relationship.
     * Valid values are `One-Way: Outgoing`, `One-Way: Incoming`, and `Two-Way`.
     * 
     */
    @Import(name="trustDirection")
    private @Nullable Output<String> trustDirection;

    /**
     * @return The direction of the Trust relationship.
     * Valid values are `One-Way: Outgoing`, `One-Way: Incoming`, and `Two-Way`.
     * 
     */
    public Optional<Output<String>> trustDirection() {
        return Optional.ofNullable(this.trustDirection);
    }

    /**
     * Password for the Trust.
     * Does not need to match the passwords for either Directory.
     * Can contain upper- and lower-case letters, numbers, and punctuation characters.
     * May be up to 128 characters long.
     * 
     */
    @Import(name="trustPassword")
    private @Nullable Output<String> trustPassword;

    /**
     * @return Password for the Trust.
     * Does not need to match the passwords for either Directory.
     * Can contain upper- and lower-case letters, numbers, and punctuation characters.
     * May be up to 128 characters long.
     * 
     */
    public Optional<Output<String>> trustPassword() {
        return Optional.ofNullable(this.trustPassword);
    }

    /**
     * State of the Trust relationship.
     * One of `Created`, `VerifyFailed`,`Verified`, `UpdateFailed`,`Updated`,`Deleted`, or `Failed`.
     * 
     */
    @Import(name="trustState")
    private @Nullable Output<String> trustState;

    /**
     * @return State of the Trust relationship.
     * One of `Created`, `VerifyFailed`,`Verified`, `UpdateFailed`,`Updated`,`Deleted`, or `Failed`.
     * 
     */
    public Optional<Output<String>> trustState() {
        return Optional.ofNullable(this.trustState);
    }

    /**
     * Reason for the Trust state set in `trust_state`.
     * 
     */
    @Import(name="trustStateReason")
    private @Nullable Output<String> trustStateReason;

    /**
     * @return Reason for the Trust state set in `trust_state`.
     * 
     */
    public Optional<Output<String>> trustStateReason() {
        return Optional.ofNullable(this.trustStateReason);
    }

    /**
     * Type of the Trust relationship.
     * Valid values are `Forest` and `External`.
     * Default value is `Forest`.
     * 
     */
    @Import(name="trustType")
    private @Nullable Output<String> trustType;

    /**
     * @return Type of the Trust relationship.
     * Valid values are `Forest` and `External`.
     * Default value is `Forest`.
     * 
     */
    public Optional<Output<String>> trustType() {
        return Optional.ofNullable(this.trustType);
    }

    private TrustState() {}

    private TrustState(TrustState $) {
        this.conditionalForwarderIpAddrs = $.conditionalForwarderIpAddrs;
        this.createdDateTime = $.createdDateTime;
        this.deleteAssociatedConditionalForwarder = $.deleteAssociatedConditionalForwarder;
        this.directoryId = $.directoryId;
        this.lastUpdatedDateTime = $.lastUpdatedDateTime;
        this.remoteDomainName = $.remoteDomainName;
        this.selectiveAuth = $.selectiveAuth;
        this.stateLastUpdatedDateTime = $.stateLastUpdatedDateTime;
        this.trustDirection = $.trustDirection;
        this.trustPassword = $.trustPassword;
        this.trustState = $.trustState;
        this.trustStateReason = $.trustStateReason;
        this.trustType = $.trustType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TrustState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TrustState $;

        public Builder() {
            $ = new TrustState();
        }

        public Builder(TrustState defaults) {
            $ = new TrustState(Objects.requireNonNull(defaults));
        }

        /**
         * @param conditionalForwarderIpAddrs Set of IPv4 addresses for the DNS server associated with the remote Directory.
         * Can contain between 1 and 4 values.
         * 
         * @return builder
         * 
         */
        public Builder conditionalForwarderIpAddrs(@Nullable Output<List<String>> conditionalForwarderIpAddrs) {
            $.conditionalForwarderIpAddrs = conditionalForwarderIpAddrs;
            return this;
        }

        /**
         * @param conditionalForwarderIpAddrs Set of IPv4 addresses for the DNS server associated with the remote Directory.
         * Can contain between 1 and 4 values.
         * 
         * @return builder
         * 
         */
        public Builder conditionalForwarderIpAddrs(List<String> conditionalForwarderIpAddrs) {
            return conditionalForwarderIpAddrs(Output.of(conditionalForwarderIpAddrs));
        }

        /**
         * @param conditionalForwarderIpAddrs Set of IPv4 addresses for the DNS server associated with the remote Directory.
         * Can contain between 1 and 4 values.
         * 
         * @return builder
         * 
         */
        public Builder conditionalForwarderIpAddrs(String... conditionalForwarderIpAddrs) {
            return conditionalForwarderIpAddrs(List.of(conditionalForwarderIpAddrs));
        }

        /**
         * @param createdDateTime Date and time when the Trust was created.
         * 
         * @return builder
         * 
         */
        public Builder createdDateTime(@Nullable Output<String> createdDateTime) {
            $.createdDateTime = createdDateTime;
            return this;
        }

        /**
         * @param createdDateTime Date and time when the Trust was created.
         * 
         * @return builder
         * 
         */
        public Builder createdDateTime(String createdDateTime) {
            return createdDateTime(Output.of(createdDateTime));
        }

        /**
         * @param deleteAssociatedConditionalForwarder Whether to delete the conditional forwarder when deleting the Trust relationship.
         * 
         * @return builder
         * 
         */
        public Builder deleteAssociatedConditionalForwarder(@Nullable Output<Boolean> deleteAssociatedConditionalForwarder) {
            $.deleteAssociatedConditionalForwarder = deleteAssociatedConditionalForwarder;
            return this;
        }

        /**
         * @param deleteAssociatedConditionalForwarder Whether to delete the conditional forwarder when deleting the Trust relationship.
         * 
         * @return builder
         * 
         */
        public Builder deleteAssociatedConditionalForwarder(Boolean deleteAssociatedConditionalForwarder) {
            return deleteAssociatedConditionalForwarder(Output.of(deleteAssociatedConditionalForwarder));
        }

        /**
         * @param directoryId ID of the Directory.
         * 
         * @return builder
         * 
         */
        public Builder directoryId(@Nullable Output<String> directoryId) {
            $.directoryId = directoryId;
            return this;
        }

        /**
         * @param directoryId ID of the Directory.
         * 
         * @return builder
         * 
         */
        public Builder directoryId(String directoryId) {
            return directoryId(Output.of(directoryId));
        }

        /**
         * @param lastUpdatedDateTime Date and time when the Trust was last updated.
         * 
         * @return builder
         * 
         */
        public Builder lastUpdatedDateTime(@Nullable Output<String> lastUpdatedDateTime) {
            $.lastUpdatedDateTime = lastUpdatedDateTime;
            return this;
        }

        /**
         * @param lastUpdatedDateTime Date and time when the Trust was last updated.
         * 
         * @return builder
         * 
         */
        public Builder lastUpdatedDateTime(String lastUpdatedDateTime) {
            return lastUpdatedDateTime(Output.of(lastUpdatedDateTime));
        }

        /**
         * @param remoteDomainName Fully qualified domain name of the remote Directory.
         * 
         * @return builder
         * 
         */
        public Builder remoteDomainName(@Nullable Output<String> remoteDomainName) {
            $.remoteDomainName = remoteDomainName;
            return this;
        }

        /**
         * @param remoteDomainName Fully qualified domain name of the remote Directory.
         * 
         * @return builder
         * 
         */
        public Builder remoteDomainName(String remoteDomainName) {
            return remoteDomainName(Output.of(remoteDomainName));
        }

        /**
         * @param selectiveAuth Whether to enable selective authentication.
         * Valid values are `Enabled` and `Disabled`.
         * Default value is `Disabled`.
         * 
         * @return builder
         * 
         */
        public Builder selectiveAuth(@Nullable Output<String> selectiveAuth) {
            $.selectiveAuth = selectiveAuth;
            return this;
        }

        /**
         * @param selectiveAuth Whether to enable selective authentication.
         * Valid values are `Enabled` and `Disabled`.
         * Default value is `Disabled`.
         * 
         * @return builder
         * 
         */
        public Builder selectiveAuth(String selectiveAuth) {
            return selectiveAuth(Output.of(selectiveAuth));
        }

        /**
         * @param stateLastUpdatedDateTime Date and time when the Trust state in `trust_state` was last updated.
         * 
         * @return builder
         * 
         */
        public Builder stateLastUpdatedDateTime(@Nullable Output<String> stateLastUpdatedDateTime) {
            $.stateLastUpdatedDateTime = stateLastUpdatedDateTime;
            return this;
        }

        /**
         * @param stateLastUpdatedDateTime Date and time when the Trust state in `trust_state` was last updated.
         * 
         * @return builder
         * 
         */
        public Builder stateLastUpdatedDateTime(String stateLastUpdatedDateTime) {
            return stateLastUpdatedDateTime(Output.of(stateLastUpdatedDateTime));
        }

        /**
         * @param trustDirection The direction of the Trust relationship.
         * Valid values are `One-Way: Outgoing`, `One-Way: Incoming`, and `Two-Way`.
         * 
         * @return builder
         * 
         */
        public Builder trustDirection(@Nullable Output<String> trustDirection) {
            $.trustDirection = trustDirection;
            return this;
        }

        /**
         * @param trustDirection The direction of the Trust relationship.
         * Valid values are `One-Way: Outgoing`, `One-Way: Incoming`, and `Two-Way`.
         * 
         * @return builder
         * 
         */
        public Builder trustDirection(String trustDirection) {
            return trustDirection(Output.of(trustDirection));
        }

        /**
         * @param trustPassword Password for the Trust.
         * Does not need to match the passwords for either Directory.
         * Can contain upper- and lower-case letters, numbers, and punctuation characters.
         * May be up to 128 characters long.
         * 
         * @return builder
         * 
         */
        public Builder trustPassword(@Nullable Output<String> trustPassword) {
            $.trustPassword = trustPassword;
            return this;
        }

        /**
         * @param trustPassword Password for the Trust.
         * Does not need to match the passwords for either Directory.
         * Can contain upper- and lower-case letters, numbers, and punctuation characters.
         * May be up to 128 characters long.
         * 
         * @return builder
         * 
         */
        public Builder trustPassword(String trustPassword) {
            return trustPassword(Output.of(trustPassword));
        }

        /**
         * @param trustState State of the Trust relationship.
         * One of `Created`, `VerifyFailed`,`Verified`, `UpdateFailed`,`Updated`,`Deleted`, or `Failed`.
         * 
         * @return builder
         * 
         */
        public Builder trustState(@Nullable Output<String> trustState) {
            $.trustState = trustState;
            return this;
        }

        /**
         * @param trustState State of the Trust relationship.
         * One of `Created`, `VerifyFailed`,`Verified`, `UpdateFailed`,`Updated`,`Deleted`, or `Failed`.
         * 
         * @return builder
         * 
         */
        public Builder trustState(String trustState) {
            return trustState(Output.of(trustState));
        }

        /**
         * @param trustStateReason Reason for the Trust state set in `trust_state`.
         * 
         * @return builder
         * 
         */
        public Builder trustStateReason(@Nullable Output<String> trustStateReason) {
            $.trustStateReason = trustStateReason;
            return this;
        }

        /**
         * @param trustStateReason Reason for the Trust state set in `trust_state`.
         * 
         * @return builder
         * 
         */
        public Builder trustStateReason(String trustStateReason) {
            return trustStateReason(Output.of(trustStateReason));
        }

        /**
         * @param trustType Type of the Trust relationship.
         * Valid values are `Forest` and `External`.
         * Default value is `Forest`.
         * 
         * @return builder
         * 
         */
        public Builder trustType(@Nullable Output<String> trustType) {
            $.trustType = trustType;
            return this;
        }

        /**
         * @param trustType Type of the Trust relationship.
         * Valid values are `Forest` and `External`.
         * Default value is `Forest`.
         * 
         * @return builder
         * 
         */
        public Builder trustType(String trustType) {
            return trustType(Output.of(trustType));
        }

        public TrustState build() {
            return $;
        }
    }

}
