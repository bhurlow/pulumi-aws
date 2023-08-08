// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.quicksight;

import com.pulumi.aws.quicksight.inputs.NamespaceTimeoutsArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NamespaceArgs extends com.pulumi.resources.ResourceArgs {

    public static final NamespaceArgs Empty = new NamespaceArgs();

    /**
     * AWS account ID.
     * 
     */
    @Import(name="awsAccountId")
    private @Nullable Output<String> awsAccountId;

    /**
     * @return AWS account ID.
     * 
     */
    public Optional<Output<String>> awsAccountId() {
        return Optional.ofNullable(this.awsAccountId);
    }

    /**
     * User identity directory type. Defaults to `QUICKSIGHT`, the only current valid value.
     * 
     */
    @Import(name="identityStore")
    private @Nullable Output<String> identityStore;

    /**
     * @return User identity directory type. Defaults to `QUICKSIGHT`, the only current valid value.
     * 
     */
    public Optional<Output<String>> identityStore() {
        return Optional.ofNullable(this.identityStore);
    }

    /**
     * Name of the namespace.
     * 
     * The following arguments are optional:
     * 
     */
    @Import(name="namespace", required=true)
    private Output<String> namespace;

    /**
     * @return Name of the namespace.
     * 
     * The following arguments are optional:
     * 
     */
    public Output<String> namespace() {
        return this.namespace;
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

    @Import(name="timeouts")
    private @Nullable Output<NamespaceTimeoutsArgs> timeouts;

    public Optional<Output<NamespaceTimeoutsArgs>> timeouts() {
        return Optional.ofNullable(this.timeouts);
    }

    private NamespaceArgs() {}

    private NamespaceArgs(NamespaceArgs $) {
        this.awsAccountId = $.awsAccountId;
        this.identityStore = $.identityStore;
        this.namespace = $.namespace;
        this.tags = $.tags;
        this.timeouts = $.timeouts;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NamespaceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NamespaceArgs $;

        public Builder() {
            $ = new NamespaceArgs();
        }

        public Builder(NamespaceArgs defaults) {
            $ = new NamespaceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param awsAccountId AWS account ID.
         * 
         * @return builder
         * 
         */
        public Builder awsAccountId(@Nullable Output<String> awsAccountId) {
            $.awsAccountId = awsAccountId;
            return this;
        }

        /**
         * @param awsAccountId AWS account ID.
         * 
         * @return builder
         * 
         */
        public Builder awsAccountId(String awsAccountId) {
            return awsAccountId(Output.of(awsAccountId));
        }

        /**
         * @param identityStore User identity directory type. Defaults to `QUICKSIGHT`, the only current valid value.
         * 
         * @return builder
         * 
         */
        public Builder identityStore(@Nullable Output<String> identityStore) {
            $.identityStore = identityStore;
            return this;
        }

        /**
         * @param identityStore User identity directory type. Defaults to `QUICKSIGHT`, the only current valid value.
         * 
         * @return builder
         * 
         */
        public Builder identityStore(String identityStore) {
            return identityStore(Output.of(identityStore));
        }

        /**
         * @param namespace Name of the namespace.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder namespace(Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace Name of the namespace.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
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

        public Builder timeouts(@Nullable Output<NamespaceTimeoutsArgs> timeouts) {
            $.timeouts = timeouts;
            return this;
        }

        public Builder timeouts(NamespaceTimeoutsArgs timeouts) {
            return timeouts(Output.of(timeouts));
        }

        public NamespaceArgs build() {
            $.namespace = Objects.requireNonNull($.namespace, "expected parameter 'namespace' to be non-null");
            return $;
        }
    }

}
