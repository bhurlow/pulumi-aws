// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cognito.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class GetUserGroupsGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetUserGroupsGroupArgs Empty = new GetUserGroupsGroupArgs();

    /**
     * Description of the user group.
     * 
     */
    @Import(name="description", required=true)
    private Output<String> description;

    /**
     * @return Description of the user group.
     * 
     */
    public Output<String> description() {
        return this.description;
    }

    /**
     * Name of the user group.
     * 
     */
    @Import(name="groupName", required=true)
    private Output<String> groupName;

    /**
     * @return Name of the user group.
     * 
     */
    public Output<String> groupName() {
        return this.groupName;
    }

    /**
     * Precedence of the user group.
     * 
     */
    @Import(name="precedence", required=true)
    private Output<Integer> precedence;

    /**
     * @return Precedence of the user group.
     * 
     */
    public Output<Integer> precedence() {
        return this.precedence;
    }

    /**
     * ARN of the IAM role to be associated with the user group.
     * 
     */
    @Import(name="roleArn", required=true)
    private Output<String> roleArn;

    /**
     * @return ARN of the IAM role to be associated with the user group.
     * 
     */
    public Output<String> roleArn() {
        return this.roleArn;
    }

    private GetUserGroupsGroupArgs() {}

    private GetUserGroupsGroupArgs(GetUserGroupsGroupArgs $) {
        this.description = $.description;
        this.groupName = $.groupName;
        this.precedence = $.precedence;
        this.roleArn = $.roleArn;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetUserGroupsGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetUserGroupsGroupArgs $;

        public Builder() {
            $ = new GetUserGroupsGroupArgs();
        }

        public Builder(GetUserGroupsGroupArgs defaults) {
            $ = new GetUserGroupsGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Description of the user group.
         * 
         * @return builder
         * 
         */
        public Builder description(Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the user group.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param groupName Name of the user group.
         * 
         * @return builder
         * 
         */
        public Builder groupName(Output<String> groupName) {
            $.groupName = groupName;
            return this;
        }

        /**
         * @param groupName Name of the user group.
         * 
         * @return builder
         * 
         */
        public Builder groupName(String groupName) {
            return groupName(Output.of(groupName));
        }

        /**
         * @param precedence Precedence of the user group.
         * 
         * @return builder
         * 
         */
        public Builder precedence(Output<Integer> precedence) {
            $.precedence = precedence;
            return this;
        }

        /**
         * @param precedence Precedence of the user group.
         * 
         * @return builder
         * 
         */
        public Builder precedence(Integer precedence) {
            return precedence(Output.of(precedence));
        }

        /**
         * @param roleArn ARN of the IAM role to be associated with the user group.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(Output<String> roleArn) {
            $.roleArn = roleArn;
            return this;
        }

        /**
         * @param roleArn ARN of the IAM role to be associated with the user group.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(String roleArn) {
            return roleArn(Output.of(roleArn));
        }

        public GetUserGroupsGroupArgs build() {
            if ($.description == null) {
                throw new MissingRequiredPropertyException("GetUserGroupsGroupArgs", "description");
            }
            if ($.groupName == null) {
                throw new MissingRequiredPropertyException("GetUserGroupsGroupArgs", "groupName");
            }
            if ($.precedence == null) {
                throw new MissingRequiredPropertyException("GetUserGroupsGroupArgs", "precedence");
            }
            if ($.roleArn == null) {
                throw new MissingRequiredPropertyException("GetUserGroupsGroupArgs", "roleArn");
            }
            return $;
        }
    }

}
