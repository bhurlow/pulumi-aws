// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.account.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RegionState extends com.pulumi.resources.ResourceArgs {

    public static final RegionState Empty = new RegionState();

    /**
     * The ID of the target account when managing member accounts. Will manage current user&#39;s account by default if omitted. To use this parameter, the caller must be an identity in the organization&#39;s management account or a delegated administrator account. The specified account ID must also be a member account in the same organization. The organization must have all features enabled, and the organization must have trusted access enabled for the Account Management service, and optionally a delegated admin account assigned.
     * 
     */
    @Import(name="accountId")
    private @Nullable Output<String> accountId;

    /**
     * @return The ID of the target account when managing member accounts. Will manage current user&#39;s account by default if omitted. To use this parameter, the caller must be an identity in the organization&#39;s management account or a delegated administrator account. The specified account ID must also be a member account in the same organization. The organization must have all features enabled, and the organization must have trusted access enabled for the Account Management service, and optionally a delegated admin account assigned.
     * 
     */
    public Optional<Output<String>> accountId() {
        return Optional.ofNullable(this.accountId);
    }

    /**
     * Whether the region is enabled.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Whether the region is enabled.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * The region opt status.
     * 
     */
    @Import(name="optStatus")
    private @Nullable Output<String> optStatus;

    /**
     * @return The region opt status.
     * 
     */
    public Optional<Output<String>> optStatus() {
        return Optional.ofNullable(this.optStatus);
    }

    /**
     * The region name to manage.
     * 
     */
    @Import(name="regionName")
    private @Nullable Output<String> regionName;

    /**
     * @return The region name to manage.
     * 
     */
    public Optional<Output<String>> regionName() {
        return Optional.ofNullable(this.regionName);
    }

    private RegionState() {}

    private RegionState(RegionState $) {
        this.accountId = $.accountId;
        this.enabled = $.enabled;
        this.optStatus = $.optStatus;
        this.regionName = $.regionName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RegionState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RegionState $;

        public Builder() {
            $ = new RegionState();
        }

        public Builder(RegionState defaults) {
            $ = new RegionState(Objects.requireNonNull(defaults));
        }

        /**
         * @param accountId The ID of the target account when managing member accounts. Will manage current user&#39;s account by default if omitted. To use this parameter, the caller must be an identity in the organization&#39;s management account or a delegated administrator account. The specified account ID must also be a member account in the same organization. The organization must have all features enabled, and the organization must have trusted access enabled for the Account Management service, and optionally a delegated admin account assigned.
         * 
         * @return builder
         * 
         */
        public Builder accountId(@Nullable Output<String> accountId) {
            $.accountId = accountId;
            return this;
        }

        /**
         * @param accountId The ID of the target account when managing member accounts. Will manage current user&#39;s account by default if omitted. To use this parameter, the caller must be an identity in the organization&#39;s management account or a delegated administrator account. The specified account ID must also be a member account in the same organization. The organization must have all features enabled, and the organization must have trusted access enabled for the Account Management service, and optionally a delegated admin account assigned.
         * 
         * @return builder
         * 
         */
        public Builder accountId(String accountId) {
            return accountId(Output.of(accountId));
        }

        /**
         * @param enabled Whether the region is enabled.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Whether the region is enabled.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param optStatus The region opt status.
         * 
         * @return builder
         * 
         */
        public Builder optStatus(@Nullable Output<String> optStatus) {
            $.optStatus = optStatus;
            return this;
        }

        /**
         * @param optStatus The region opt status.
         * 
         * @return builder
         * 
         */
        public Builder optStatus(String optStatus) {
            return optStatus(Output.of(optStatus));
        }

        /**
         * @param regionName The region name to manage.
         * 
         * @return builder
         * 
         */
        public Builder regionName(@Nullable Output<String> regionName) {
            $.regionName = regionName;
            return this;
        }

        /**
         * @param regionName The region name to manage.
         * 
         * @return builder
         * 
         */
        public Builder regionName(String regionName) {
            return regionName(Output.of(regionName));
        }

        public RegionState build() {
            return $;
        }
    }

}
