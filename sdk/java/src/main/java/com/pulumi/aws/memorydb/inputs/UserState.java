// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.memorydb.inputs;

import com.pulumi.aws.memorydb.inputs.UserAuthenticationModeArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UserState extends com.pulumi.resources.ResourceArgs {

    public static final UserState Empty = new UserState();

    /**
     * The access permissions string used for this user.
     * 
     */
    @Import(name="accessString")
    private @Nullable Output<String> accessString;

    /**
     * @return The access permissions string used for this user.
     * 
     */
    public Optional<Output<String>> accessString() {
        return Optional.ofNullable(this.accessString);
    }

    /**
     * The ARN of the user.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The ARN of the user.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * Denotes the user&#39;s authentication properties. Detailed below.
     * 
     */
    @Import(name="authenticationMode")
    private @Nullable Output<UserAuthenticationModeArgs> authenticationMode;

    /**
     * @return Denotes the user&#39;s authentication properties. Detailed below.
     * 
     */
    public Optional<Output<UserAuthenticationModeArgs>> authenticationMode() {
        return Optional.ofNullable(this.authenticationMode);
    }

    /**
     * The minimum engine version supported for the user.
     * 
     */
    @Import(name="minimumEngineVersion")
    private @Nullable Output<String> minimumEngineVersion;

    /**
     * @return The minimum engine version supported for the user.
     * 
     */
    public Optional<Output<String>> minimumEngineVersion() {
        return Optional.ofNullable(this.minimumEngineVersion);
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
     * Name of the MemoryDB user. Up to 40 characters.
     * 
     * The following arguments are optional:
     * 
     */
    @Import(name="userName")
    private @Nullable Output<String> userName;

    /**
     * @return Name of the MemoryDB user. Up to 40 characters.
     * 
     * The following arguments are optional:
     * 
     */
    public Optional<Output<String>> userName() {
        return Optional.ofNullable(this.userName);
    }

    private UserState() {}

    private UserState(UserState $) {
        this.accessString = $.accessString;
        this.arn = $.arn;
        this.authenticationMode = $.authenticationMode;
        this.minimumEngineVersion = $.minimumEngineVersion;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.userName = $.userName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserState $;

        public Builder() {
            $ = new UserState();
        }

        public Builder(UserState defaults) {
            $ = new UserState(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessString The access permissions string used for this user.
         * 
         * @return builder
         * 
         */
        public Builder accessString(@Nullable Output<String> accessString) {
            $.accessString = accessString;
            return this;
        }

        /**
         * @param accessString The access permissions string used for this user.
         * 
         * @return builder
         * 
         */
        public Builder accessString(String accessString) {
            return accessString(Output.of(accessString));
        }

        /**
         * @param arn The ARN of the user.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The ARN of the user.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param authenticationMode Denotes the user&#39;s authentication properties. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder authenticationMode(@Nullable Output<UserAuthenticationModeArgs> authenticationMode) {
            $.authenticationMode = authenticationMode;
            return this;
        }

        /**
         * @param authenticationMode Denotes the user&#39;s authentication properties. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder authenticationMode(UserAuthenticationModeArgs authenticationMode) {
            return authenticationMode(Output.of(authenticationMode));
        }

        /**
         * @param minimumEngineVersion The minimum engine version supported for the user.
         * 
         * @return builder
         * 
         */
        public Builder minimumEngineVersion(@Nullable Output<String> minimumEngineVersion) {
            $.minimumEngineVersion = minimumEngineVersion;
            return this;
        }

        /**
         * @param minimumEngineVersion The minimum engine version supported for the user.
         * 
         * @return builder
         * 
         */
        public Builder minimumEngineVersion(String minimumEngineVersion) {
            return minimumEngineVersion(Output.of(minimumEngineVersion));
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
         * @param userName Name of the MemoryDB user. Up to 40 characters.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder userName(@Nullable Output<String> userName) {
            $.userName = userName;
            return this;
        }

        /**
         * @param userName Name of the MemoryDB user. Up to 40 characters.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder userName(String userName) {
            return userName(Output.of(userName));
        }

        public UserState build() {
            return $;
        }
    }

}
