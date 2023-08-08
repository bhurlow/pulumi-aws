// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.auditmanager.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AssessmentReportState extends com.pulumi.resources.ResourceArgs {

    public static final AssessmentReportState Empty = new AssessmentReportState();

    /**
     * Unique identifier of the assessment to create the report from.
     * 
     * The following arguments are optional:
     * 
     */
    @Import(name="assessmentId")
    private @Nullable Output<String> assessmentId;

    /**
     * @return Unique identifier of the assessment to create the report from.
     * 
     * The following arguments are optional:
     * 
     */
    public Optional<Output<String>> assessmentId() {
        return Optional.ofNullable(this.assessmentId);
    }

    /**
     * Name of the user who created the assessment report.
     * 
     */
    @Import(name="author")
    private @Nullable Output<String> author;

    /**
     * @return Name of the user who created the assessment report.
     * 
     */
    public Optional<Output<String>> author() {
        return Optional.ofNullable(this.author);
    }

    /**
     * Description of the assessment report.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the assessment report.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Name of the assessment report.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the assessment report.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Current status of the specified assessment report. Valid values are `COMPLETE`, `IN_PROGRESS`, and `FAILED`.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return Current status of the specified assessment report. Valid values are `COMPLETE`, `IN_PROGRESS`, and `FAILED`.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private AssessmentReportState() {}

    private AssessmentReportState(AssessmentReportState $) {
        this.assessmentId = $.assessmentId;
        this.author = $.author;
        this.description = $.description;
        this.name = $.name;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AssessmentReportState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AssessmentReportState $;

        public Builder() {
            $ = new AssessmentReportState();
        }

        public Builder(AssessmentReportState defaults) {
            $ = new AssessmentReportState(Objects.requireNonNull(defaults));
        }

        /**
         * @param assessmentId Unique identifier of the assessment to create the report from.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder assessmentId(@Nullable Output<String> assessmentId) {
            $.assessmentId = assessmentId;
            return this;
        }

        /**
         * @param assessmentId Unique identifier of the assessment to create the report from.
         * 
         * The following arguments are optional:
         * 
         * @return builder
         * 
         */
        public Builder assessmentId(String assessmentId) {
            return assessmentId(Output.of(assessmentId));
        }

        /**
         * @param author Name of the user who created the assessment report.
         * 
         * @return builder
         * 
         */
        public Builder author(@Nullable Output<String> author) {
            $.author = author;
            return this;
        }

        /**
         * @param author Name of the user who created the assessment report.
         * 
         * @return builder
         * 
         */
        public Builder author(String author) {
            return author(Output.of(author));
        }

        /**
         * @param description Description of the assessment report.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the assessment report.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name Name of the assessment report.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the assessment report.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param status Current status of the specified assessment report. Valid values are `COMPLETE`, `IN_PROGRESS`, and `FAILED`.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status Current status of the specified assessment report. Valid values are `COMPLETE`, `IN_PROGRESS`, and `FAILED`.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public AssessmentReportState build() {
            return $;
        }
    }

}
