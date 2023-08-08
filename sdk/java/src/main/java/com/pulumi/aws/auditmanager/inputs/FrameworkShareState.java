// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.auditmanager.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FrameworkShareState extends com.pulumi.resources.ResourceArgs {

    public static final FrameworkShareState Empty = new FrameworkShareState();

    /**
     * Comment from the sender about the share request.
     * 
     */
    @Import(name="comment")
    private @Nullable Output<String> comment;

    /**
     * @return Comment from the sender about the share request.
     * 
     */
    public Optional<Output<String>> comment() {
        return Optional.ofNullable(this.comment);
    }

    /**
     * Amazon Web Services account of the recipient.
     * 
     */
    @Import(name="destinationAccount")
    private @Nullable Output<String> destinationAccount;

    /**
     * @return Amazon Web Services account of the recipient.
     * 
     */
    public Optional<Output<String>> destinationAccount() {
        return Optional.ofNullable(this.destinationAccount);
    }

    /**
     * Amazon Web Services region of the recipient.
     * 
     */
    @Import(name="destinationRegion")
    private @Nullable Output<String> destinationRegion;

    /**
     * @return Amazon Web Services region of the recipient.
     * 
     */
    public Optional<Output<String>> destinationRegion() {
        return Optional.ofNullable(this.destinationRegion);
    }

    /**
     * Unique identifier for the shared custom framework.
     * 
     * The following arguments are optional:
     * 
     */
    @Import(name="frameworkId")
    private @Nullable Output<String> frameworkId;

    /**
     * @return Unique identifier for the shared custom framework.
     * 
     * The following arguments are optional:
     * 
     */
    public Optional<Output<String>> frameworkId() {
        return Optional.ofNullable(this.frameworkId);
    }

    /**
     * Status of the share request.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return Status of the share request.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private FrameworkShareState() {}

    private FrameworkShareState(FrameworkShareState $) {
        this.comment = $.comment;
        this.destinationAccount = $.destinationAccount;
        this.destinationRegion = $.destinationRegion;
        this.frameworkId = $.frameworkId;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FrameworkShareState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FrameworkShareState $;

        public Builder() {
            $ = new FrameworkShareState();
        }

        public Builder(FrameworkShareState defaults) {
            $ = new FrameworkShareState(Objects.requireNonNull(defaults));
        }

        /**
         * @param comment Comment from the sender about the share request.
         * 
         * @return builder
         * 
         */
        public Builder comment(@Nullable Output<String> comment) {
            $.comment = comment;
            return this;
        }

        /**
         * @param comment Comment from the sender about the share request.
         * 
         * @return builder
         * 
         */
        public Builder comment(String comment) {
            return comment(Output.of(comment));
        }

        /**
         * @param destinationAccount Amazon Web Services account of the recipient.
         * 
         * @return builder
         * 
         */
        public Builder destinationAccount(@Nullable Output<String> destinationAccount) {
            $.destinationAccount = destinationAccount;
            return this;
        }

        /**
         * @param destinationAccount Amazon Web Services account of the recipient.
         * 
         * @return builder
         * 
         */
        public Builder destinationAccount(String destinationAccount) {
            return destinationAccount(Output.of(destinationAccount));
        }

        /**
         * @param destinationRegion Amazon Web Services region of the recipient.
         * 
         * @return builder
         * 
         */
        public Builder destinationRegion(@Nullable Output<String> destinationRegion) {
            $.destinationRegion = destinationRegion;
            return this;
        }

        /**
         * @param destinationRegion Amazon Web Services region of the recipient.
         * 
         * @return builder
         * 
         */
        public Builder destinationRegion(String destinationRegion) {
            return destinationRegion(Output.of(destinationRegion));
        }

        /**
         * @param frameworkId Unique identifier for the shared custom framework.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder frameworkId(@Nullable Output<String> frameworkId) {
            $.frameworkId = frameworkId;
            return this;
        }

        /**
         * @param frameworkId Unique identifier for the shared custom framework.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder frameworkId(String frameworkId) {
            return frameworkId(Output.of(frameworkId));
        }

        /**
         * @param status Status of the share request.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status Status of the share request.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public FrameworkShareState build() {
            return $;
        }
    }

}
