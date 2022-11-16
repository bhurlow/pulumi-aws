// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.organizations.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AccountState extends com.pulumi.resources.ResourceArgs {

    public static final AccountState Empty = new AccountState();

    /**
     * The ARN for this account.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The ARN for this account.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * If true, a deletion event will close the account. Otherwise, it will only remove from the organization. This is not supported for GovCloud accounts.
     * 
     */
    @Import(name="closeOnDeletion")
    private @Nullable Output<Boolean> closeOnDeletion;

    /**
     * @return If true, a deletion event will close the account. Otherwise, it will only remove from the organization. This is not supported for GovCloud accounts.
     * 
     */
    public Optional<Output<Boolean>> closeOnDeletion() {
        return Optional.ofNullable(this.closeOnDeletion);
    }

    /**
     * Whether to also create a GovCloud account. The GovCloud account is tied to the main (commercial) account this resource creates. If `true`, the GovCloud account ID is available in the `govcloud_id` attribute. The only way to manage the GovCloud account with the provider is to subsequently import the account using this resource.
     * 
     */
    @Import(name="createGovcloud")
    private @Nullable Output<Boolean> createGovcloud;

    /**
     * @return Whether to also create a GovCloud account. The GovCloud account is tied to the main (commercial) account this resource creates. If `true`, the GovCloud account ID is available in the `govcloud_id` attribute. The only way to manage the GovCloud account with the provider is to subsequently import the account using this resource.
     * 
     */
    public Optional<Output<Boolean>> createGovcloud() {
        return Optional.ofNullable(this.createGovcloud);
    }

    /**
     * Email address of the owner to assign to the new member account. This email address must not already be associated with another AWS account.
     * 
     */
    @Import(name="email")
    private @Nullable Output<String> email;

    /**
     * @return Email address of the owner to assign to the new member account. This email address must not already be associated with another AWS account.
     * 
     */
    public Optional<Output<String>> email() {
        return Optional.ofNullable(this.email);
    }

    /**
     * ID for a GovCloud account created with the account.
     * 
     */
    @Import(name="govcloudId")
    private @Nullable Output<String> govcloudId;

    /**
     * @return ID for a GovCloud account created with the account.
     * 
     */
    public Optional<Output<String>> govcloudId() {
        return Optional.ofNullable(this.govcloudId);
    }

    /**
     * If set to `ALLOW`, the new account enables IAM users and roles to access account billing information if they have the required permissions. If set to `DENY`, then only the root user (and no roles) of the new account can access account billing information. If this is unset, the AWS API will default this to `ALLOW`. If the resource is created and this option is changed, it will try to recreate the account.
     * 
     */
    @Import(name="iamUserAccessToBilling")
    private @Nullable Output<String> iamUserAccessToBilling;

    /**
     * @return If set to `ALLOW`, the new account enables IAM users and roles to access account billing information if they have the required permissions. If set to `DENY`, then only the root user (and no roles) of the new account can access account billing information. If this is unset, the AWS API will default this to `ALLOW`. If the resource is created and this option is changed, it will try to recreate the account.
     * 
     */
    public Optional<Output<String>> iamUserAccessToBilling() {
        return Optional.ofNullable(this.iamUserAccessToBilling);
    }

    @Import(name="joinedMethod")
    private @Nullable Output<String> joinedMethod;

    public Optional<Output<String>> joinedMethod() {
        return Optional.ofNullable(this.joinedMethod);
    }

    @Import(name="joinedTimestamp")
    private @Nullable Output<String> joinedTimestamp;

    public Optional<Output<String>> joinedTimestamp() {
        return Optional.ofNullable(this.joinedTimestamp);
    }

    /**
     * Friendly name for the member account.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Friendly name for the member account.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Parent Organizational Unit ID or Root ID for the account. Defaults to the Organization default Root ID. A configuration must be present for this argument to perform drift detection.
     * 
     */
    @Import(name="parentId")
    private @Nullable Output<String> parentId;

    /**
     * @return Parent Organizational Unit ID or Root ID for the account. Defaults to the Organization default Root ID. A configuration must be present for this argument to perform drift detection.
     * 
     */
    public Optional<Output<String>> parentId() {
        return Optional.ofNullable(this.parentId);
    }

    /**
     * The name of an IAM role that Organizations automatically preconfigures in the new member account. This role trusts the root account, allowing users in the root account to assume the role, as permitted by the root account administrator. The role has administrator permissions in the new member account. The Organizations API provides no method for reading this information after account creation, so the provider cannot perform drift detection on its value and will always show a difference for a configured value after import unless `ignoreChanges` is used.
     * 
     */
    @Import(name="roleName")
    private @Nullable Output<String> roleName;

    /**
     * @return The name of an IAM role that Organizations automatically preconfigures in the new member account. This role trusts the root account, allowing users in the root account to assume the role, as permitted by the root account administrator. The role has administrator permissions in the new member account. The Organizations API provides no method for reading this information after account creation, so the provider cannot perform drift detection on its value and will always show a difference for a configured value after import unless `ignoreChanges` is used.
     * 
     */
    public Optional<Output<String>> roleName() {
        return Optional.ofNullable(this.roleName);
    }

    @Import(name="status")
    private @Nullable Output<String> status;

    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

    private AccountState() {}

    private AccountState(AccountState $) {
        this.arn = $.arn;
        this.closeOnDeletion = $.closeOnDeletion;
        this.createGovcloud = $.createGovcloud;
        this.email = $.email;
        this.govcloudId = $.govcloudId;
        this.iamUserAccessToBilling = $.iamUserAccessToBilling;
        this.joinedMethod = $.joinedMethod;
        this.joinedTimestamp = $.joinedTimestamp;
        this.name = $.name;
        this.parentId = $.parentId;
        this.roleName = $.roleName;
        this.status = $.status;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AccountState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AccountState $;

        public Builder() {
            $ = new AccountState();
        }

        public Builder(AccountState defaults) {
            $ = new AccountState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The ARN for this account.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The ARN for this account.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param closeOnDeletion If true, a deletion event will close the account. Otherwise, it will only remove from the organization. This is not supported for GovCloud accounts.
         * 
         * @return builder
         * 
         */
        public Builder closeOnDeletion(@Nullable Output<Boolean> closeOnDeletion) {
            $.closeOnDeletion = closeOnDeletion;
            return this;
        }

        /**
         * @param closeOnDeletion If true, a deletion event will close the account. Otherwise, it will only remove from the organization. This is not supported for GovCloud accounts.
         * 
         * @return builder
         * 
         */
        public Builder closeOnDeletion(Boolean closeOnDeletion) {
            return closeOnDeletion(Output.of(closeOnDeletion));
        }

        /**
         * @param createGovcloud Whether to also create a GovCloud account. The GovCloud account is tied to the main (commercial) account this resource creates. If `true`, the GovCloud account ID is available in the `govcloud_id` attribute. The only way to manage the GovCloud account with the provider is to subsequently import the account using this resource.
         * 
         * @return builder
         * 
         */
        public Builder createGovcloud(@Nullable Output<Boolean> createGovcloud) {
            $.createGovcloud = createGovcloud;
            return this;
        }

        /**
         * @param createGovcloud Whether to also create a GovCloud account. The GovCloud account is tied to the main (commercial) account this resource creates. If `true`, the GovCloud account ID is available in the `govcloud_id` attribute. The only way to manage the GovCloud account with the provider is to subsequently import the account using this resource.
         * 
         * @return builder
         * 
         */
        public Builder createGovcloud(Boolean createGovcloud) {
            return createGovcloud(Output.of(createGovcloud));
        }

        /**
         * @param email Email address of the owner to assign to the new member account. This email address must not already be associated with another AWS account.
         * 
         * @return builder
         * 
         */
        public Builder email(@Nullable Output<String> email) {
            $.email = email;
            return this;
        }

        /**
         * @param email Email address of the owner to assign to the new member account. This email address must not already be associated with another AWS account.
         * 
         * @return builder
         * 
         */
        public Builder email(String email) {
            return email(Output.of(email));
        }

        /**
         * @param govcloudId ID for a GovCloud account created with the account.
         * 
         * @return builder
         * 
         */
        public Builder govcloudId(@Nullable Output<String> govcloudId) {
            $.govcloudId = govcloudId;
            return this;
        }

        /**
         * @param govcloudId ID for a GovCloud account created with the account.
         * 
         * @return builder
         * 
         */
        public Builder govcloudId(String govcloudId) {
            return govcloudId(Output.of(govcloudId));
        }

        /**
         * @param iamUserAccessToBilling If set to `ALLOW`, the new account enables IAM users and roles to access account billing information if they have the required permissions. If set to `DENY`, then only the root user (and no roles) of the new account can access account billing information. If this is unset, the AWS API will default this to `ALLOW`. If the resource is created and this option is changed, it will try to recreate the account.
         * 
         * @return builder
         * 
         */
        public Builder iamUserAccessToBilling(@Nullable Output<String> iamUserAccessToBilling) {
            $.iamUserAccessToBilling = iamUserAccessToBilling;
            return this;
        }

        /**
         * @param iamUserAccessToBilling If set to `ALLOW`, the new account enables IAM users and roles to access account billing information if they have the required permissions. If set to `DENY`, then only the root user (and no roles) of the new account can access account billing information. If this is unset, the AWS API will default this to `ALLOW`. If the resource is created and this option is changed, it will try to recreate the account.
         * 
         * @return builder
         * 
         */
        public Builder iamUserAccessToBilling(String iamUserAccessToBilling) {
            return iamUserAccessToBilling(Output.of(iamUserAccessToBilling));
        }

        public Builder joinedMethod(@Nullable Output<String> joinedMethod) {
            $.joinedMethod = joinedMethod;
            return this;
        }

        public Builder joinedMethod(String joinedMethod) {
            return joinedMethod(Output.of(joinedMethod));
        }

        public Builder joinedTimestamp(@Nullable Output<String> joinedTimestamp) {
            $.joinedTimestamp = joinedTimestamp;
            return this;
        }

        public Builder joinedTimestamp(String joinedTimestamp) {
            return joinedTimestamp(Output.of(joinedTimestamp));
        }

        /**
         * @param name Friendly name for the member account.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Friendly name for the member account.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param parentId Parent Organizational Unit ID or Root ID for the account. Defaults to the Organization default Root ID. A configuration must be present for this argument to perform drift detection.
         * 
         * @return builder
         * 
         */
        public Builder parentId(@Nullable Output<String> parentId) {
            $.parentId = parentId;
            return this;
        }

        /**
         * @param parentId Parent Organizational Unit ID or Root ID for the account. Defaults to the Organization default Root ID. A configuration must be present for this argument to perform drift detection.
         * 
         * @return builder
         * 
         */
        public Builder parentId(String parentId) {
            return parentId(Output.of(parentId));
        }

        /**
         * @param roleName The name of an IAM role that Organizations automatically preconfigures in the new member account. This role trusts the root account, allowing users in the root account to assume the role, as permitted by the root account administrator. The role has administrator permissions in the new member account. The Organizations API provides no method for reading this information after account creation, so the provider cannot perform drift detection on its value and will always show a difference for a configured value after import unless `ignoreChanges` is used.
         * 
         * @return builder
         * 
         */
        public Builder roleName(@Nullable Output<String> roleName) {
            $.roleName = roleName;
            return this;
        }

        /**
         * @param roleName The name of an IAM role that Organizations automatically preconfigures in the new member account. This role trusts the root account, allowing users in the root account to assume the role, as permitted by the root account administrator. The role has administrator permissions in the new member account. The Organizations API provides no method for reading this information after account creation, so the provider cannot perform drift detection on its value and will always show a difference for a configured value after import unless `ignoreChanges` is used.
         * 
         * @return builder
         * 
         */
        public Builder roleName(String roleName) {
            return roleName(Output.of(roleName));
        }

        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public AccountState build() {
            return $;
        }
    }

}
