// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.memorydb;

import com.pulumi.aws.memorydb.inputs.UserAuthenticationModeArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UserArgs extends com.pulumi.resources.ResourceArgs {

    public static final UserArgs Empty = new UserArgs();

    /**
     * The access permissions string used for this user.
     * 
     */
    @Import(name="accessString", required=true)
    private Output<String> accessString;

    /**
     * @return The access permissions string used for this user.
     * 
     */
    public Output<String> accessString() {
        return this.accessString;
    }

    /**
     * Denotes the user&#39;s authentication properties. Detailed below.
     * 
     */
    @Import(name="authenticationMode", required=true)
    private Output<UserAuthenticationModeArgs> authenticationMode;

    /**
     * @return Denotes the user&#39;s authentication properties. Detailed below.
     * 
     */
    public Output<UserAuthenticationModeArgs> authenticationMode() {
        return this.authenticationMode;
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
     * Name of the MemoryDB user. Up to 40 characters.
     * 
     */
    @Import(name="userName", required=true)
    private Output<String> userName;

    /**
     * @return Name of the MemoryDB user. Up to 40 characters.
     * 
     */
    public Output<String> userName() {
        return this.userName;
    }

    private UserArgs() {}

    private UserArgs(UserArgs $) {
        this.accessString = $.accessString;
        this.authenticationMode = $.authenticationMode;
        this.tags = $.tags;
        this.userName = $.userName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserArgs $;

        public Builder() {
            $ = new UserArgs();
        }

        public Builder(UserArgs defaults) {
            $ = new UserArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessString The access permissions string used for this user.
         * 
         * @return builder
         * 
         */
        public Builder accessString(Output<String> accessString) {
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
         * @param authenticationMode Denotes the user&#39;s authentication properties. Detailed below.
         * 
         * @return builder
         * 
         */
        public Builder authenticationMode(Output<UserAuthenticationModeArgs> authenticationMode) {
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
         * @param userName Name of the MemoryDB user. Up to 40 characters.
         * 
         * @return builder
         * 
         */
        public Builder userName(Output<String> userName) {
            $.userName = userName;
            return this;
        }

        /**
         * @param userName Name of the MemoryDB user. Up to 40 characters.
         * 
         * @return builder
         * 
         */
        public Builder userName(String userName) {
            return userName(Output.of(userName));
        }

        public UserArgs build() {
            $.accessString = Objects.requireNonNull($.accessString, "expected parameter 'accessString' to be non-null");
            $.authenticationMode = Objects.requireNonNull($.authenticationMode, "expected parameter 'authenticationMode' to be non-null");
            $.userName = Objects.requireNonNull($.userName, "expected parameter 'userName' to be non-null");
            return $;
        }
    }

}
