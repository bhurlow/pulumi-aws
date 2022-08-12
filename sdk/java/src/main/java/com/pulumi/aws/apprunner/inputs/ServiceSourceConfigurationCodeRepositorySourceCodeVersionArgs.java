// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.apprunner.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class ServiceSourceConfigurationCodeRepositorySourceCodeVersionArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceSourceConfigurationCodeRepositorySourceCodeVersionArgs Empty = new ServiceSourceConfigurationCodeRepositorySourceCodeVersionArgs();

    /**
     * The type of version identifier. For a git-based repository, branches represent versions. Valid values: `BRANCH`.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return The type of version identifier. For a git-based repository, branches represent versions. Valid values: `BRANCH`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     * A source code version. For a git-based repository, a branch name maps to a specific version. App Runner uses the most recent commit to the branch.
     * 
     */
    @Import(name="value", required=true)
    private Output<String> value;

    /**
     * @return A source code version. For a git-based repository, a branch name maps to a specific version. App Runner uses the most recent commit to the branch.
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    private ServiceSourceConfigurationCodeRepositorySourceCodeVersionArgs() {}

    private ServiceSourceConfigurationCodeRepositorySourceCodeVersionArgs(ServiceSourceConfigurationCodeRepositorySourceCodeVersionArgs $) {
        this.type = $.type;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceSourceConfigurationCodeRepositorySourceCodeVersionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceSourceConfigurationCodeRepositorySourceCodeVersionArgs $;

        public Builder() {
            $ = new ServiceSourceConfigurationCodeRepositorySourceCodeVersionArgs();
        }

        public Builder(ServiceSourceConfigurationCodeRepositorySourceCodeVersionArgs defaults) {
            $ = new ServiceSourceConfigurationCodeRepositorySourceCodeVersionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param type The type of version identifier. For a git-based repository, branches represent versions. Valid values: `BRANCH`.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The type of version identifier. For a git-based repository, branches represent versions. Valid values: `BRANCH`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param value A source code version. For a git-based repository, a branch name maps to a specific version. App Runner uses the most recent commit to the branch.
         * 
         * @return builder
         * 
         */
        public Builder value(Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value A source code version. For a git-based repository, a branch name maps to a specific version. App Runner uses the most recent commit to the branch.
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public ServiceSourceConfigurationCodeRepositorySourceCodeVersionArgs build() {
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            $.value = Objects.requireNonNull($.value, "expected parameter 'value' to be non-null");
            return $;
        }
    }

}