// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ssoadmin.inputs;

import com.pulumi.aws.ssoadmin.inputs.GetApplicationAssignmentsApplicationAssignment;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetApplicationAssignmentsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetApplicationAssignmentsPlainArgs Empty = new GetApplicationAssignmentsPlainArgs();

    /**
     * ARN of the application.
     * 
     */
    @Import(name="applicationArn", required=true)
    private String applicationArn;

    /**
     * @return ARN of the application.
     * 
     */
    public String applicationArn() {
        return this.applicationArn;
    }

    /**
     * List of principals assigned to the application. See the `application_assignments` attribute reference below.
     * 
     */
    @Import(name="applicationAssignments")
    private @Nullable List<GetApplicationAssignmentsApplicationAssignment> applicationAssignments;

    /**
     * @return List of principals assigned to the application. See the `application_assignments` attribute reference below.
     * 
     */
    public Optional<List<GetApplicationAssignmentsApplicationAssignment>> applicationAssignments() {
        return Optional.ofNullable(this.applicationAssignments);
    }

    private GetApplicationAssignmentsPlainArgs() {}

    private GetApplicationAssignmentsPlainArgs(GetApplicationAssignmentsPlainArgs $) {
        this.applicationArn = $.applicationArn;
        this.applicationAssignments = $.applicationAssignments;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetApplicationAssignmentsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetApplicationAssignmentsPlainArgs $;

        public Builder() {
            $ = new GetApplicationAssignmentsPlainArgs();
        }

        public Builder(GetApplicationAssignmentsPlainArgs defaults) {
            $ = new GetApplicationAssignmentsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param applicationArn ARN of the application.
         * 
         * @return builder
         * 
         */
        public Builder applicationArn(String applicationArn) {
            $.applicationArn = applicationArn;
            return this;
        }

        /**
         * @param applicationAssignments List of principals assigned to the application. See the `application_assignments` attribute reference below.
         * 
         * @return builder
         * 
         */
        public Builder applicationAssignments(@Nullable List<GetApplicationAssignmentsApplicationAssignment> applicationAssignments) {
            $.applicationAssignments = applicationAssignments;
            return this;
        }

        /**
         * @param applicationAssignments List of principals assigned to the application. See the `application_assignments` attribute reference below.
         * 
         * @return builder
         * 
         */
        public Builder applicationAssignments(GetApplicationAssignmentsApplicationAssignment... applicationAssignments) {
            return applicationAssignments(List.of(applicationAssignments));
        }

        public GetApplicationAssignmentsPlainArgs build() {
            $.applicationArn = Objects.requireNonNull($.applicationArn, "expected parameter 'applicationArn' to be non-null");
            return $;
        }
    }

}
