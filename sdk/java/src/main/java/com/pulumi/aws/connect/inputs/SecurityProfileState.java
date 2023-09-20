// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.connect.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecurityProfileState extends com.pulumi.resources.ResourceArgs {

    public static final SecurityProfileState Empty = new SecurityProfileState();

    /**
     * The Amazon Resource Name (ARN) of the Security Profile.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) of the Security Profile.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Specifies the description of the Security Profile.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Specifies the description of the Security Profile.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Specifies the identifier of the hosting Amazon Connect Instance.
     * 
     */
    @Import(name="instanceId")
    private @Nullable Output<String> instanceId;

    /**
     * @return Specifies the identifier of the hosting Amazon Connect Instance.
     * 
     */
    public Optional<Output<String>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * Specifies the name of the Security Profile.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Specifies the name of the Security Profile.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The organization resource identifier for the security profile.
     * 
     */
    @Import(name="organizationResourceId")
    private @Nullable Output<String> organizationResourceId;

    /**
     * @return The organization resource identifier for the security profile.
     * 
     */
    public Optional<Output<String>> organizationResourceId() {
        return Optional.ofNullable(this.organizationResourceId);
    }

    /**
     * Specifies a list of permissions assigned to the security profile.
     * 
     */
    @Import(name="permissions")
    private @Nullable Output<List<String>> permissions;

    /**
     * @return Specifies a list of permissions assigned to the security profile.
     * 
     */
    public Optional<Output<List<String>>> permissions() {
        return Optional.ofNullable(this.permissions);
    }

    /**
     * The identifier for the Security Profile.
     * 
     */
    @Import(name="securityProfileId")
    private @Nullable Output<String> securityProfileId;

    /**
     * @return The identifier for the Security Profile.
     * 
     */
    public Optional<Output<String>> securityProfileId() {
        return Optional.ofNullable(this.securityProfileId);
    }

    /**
     * Tags to apply to the Security Profile. If configured with a provider
     * `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Tags to apply to the Security Profile. If configured with a provider
     * `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

    private SecurityProfileState() {}

    private SecurityProfileState(SecurityProfileState $) {
        this.arn = $.arn;
        this.description = $.description;
        this.instanceId = $.instanceId;
        this.name = $.name;
        this.organizationResourceId = $.organizationResourceId;
        this.permissions = $.permissions;
        this.securityProfileId = $.securityProfileId;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecurityProfileState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecurityProfileState $;

        public Builder() {
            $ = new SecurityProfileState();
        }

        public Builder(SecurityProfileState defaults) {
            $ = new SecurityProfileState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The Amazon Resource Name (ARN) of the Security Profile.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The Amazon Resource Name (ARN) of the Security Profile.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param description Specifies the description of the Security Profile.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Specifies the description of the Security Profile.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param instanceId Specifies the identifier of the hosting Amazon Connect Instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId Specifies the identifier of the hosting Amazon Connect Instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param name Specifies the name of the Security Profile.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Specifies the name of the Security Profile.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param organizationResourceId The organization resource identifier for the security profile.
         * 
         * @return builder
         * 
         */
        public Builder organizationResourceId(@Nullable Output<String> organizationResourceId) {
            $.organizationResourceId = organizationResourceId;
            return this;
        }

        /**
         * @param organizationResourceId The organization resource identifier for the security profile.
         * 
         * @return builder
         * 
         */
        public Builder organizationResourceId(String organizationResourceId) {
            return organizationResourceId(Output.of(organizationResourceId));
        }

        /**
         * @param permissions Specifies a list of permissions assigned to the security profile.
         * 
         * @return builder
         * 
         */
        public Builder permissions(@Nullable Output<List<String>> permissions) {
            $.permissions = permissions;
            return this;
        }

        /**
         * @param permissions Specifies a list of permissions assigned to the security profile.
         * 
         * @return builder
         * 
         */
        public Builder permissions(List<String> permissions) {
            return permissions(Output.of(permissions));
        }

        /**
         * @param permissions Specifies a list of permissions assigned to the security profile.
         * 
         * @return builder
         * 
         */
        public Builder permissions(String... permissions) {
            return permissions(List.of(permissions));
        }

        /**
         * @param securityProfileId The identifier for the Security Profile.
         * 
         * @return builder
         * 
         */
        public Builder securityProfileId(@Nullable Output<String> securityProfileId) {
            $.securityProfileId = securityProfileId;
            return this;
        }

        /**
         * @param securityProfileId The identifier for the Security Profile.
         * 
         * @return builder
         * 
         */
        public Builder securityProfileId(String securityProfileId) {
            return securityProfileId(Output.of(securityProfileId));
        }

        /**
         * @param tags Tags to apply to the Security Profile. If configured with a provider
         * `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Tags to apply to the Security Profile. If configured with a provider
         * `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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

        public SecurityProfileState build() {
            return $;
        }
    }

}
