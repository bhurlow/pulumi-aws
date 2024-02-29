// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cognito.outputs;

import com.pulumi.aws.cognito.outputs.GetUserGroupsGroup;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetUserGroupsResult {
    /**
     * @return List of groups. See `groups` below.
     * 
     */
    private @Nullable List<GetUserGroupsGroup> groups;
    /**
     * @return User pool identifier.
     * 
     */
    private String id;
    private String userPoolId;

    private GetUserGroupsResult() {}
    /**
     * @return List of groups. See `groups` below.
     * 
     */
    public List<GetUserGroupsGroup> groups() {
        return this.groups == null ? List.of() : this.groups;
    }
    /**
     * @return User pool identifier.
     * 
     */
    public String id() {
        return this.id;
    }
    public String userPoolId() {
        return this.userPoolId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetUserGroupsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<GetUserGroupsGroup> groups;
        private String id;
        private String userPoolId;
        public Builder() {}
        public Builder(GetUserGroupsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.groups = defaults.groups;
    	      this.id = defaults.id;
    	      this.userPoolId = defaults.userPoolId;
        }

        @CustomType.Setter
        public Builder groups(@Nullable List<GetUserGroupsGroup> groups) {

            this.groups = groups;
            return this;
        }
        public Builder groups(GetUserGroupsGroup... groups) {
            return groups(List.of(groups));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetUserGroupsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder userPoolId(String userPoolId) {
            if (userPoolId == null) {
              throw new MissingRequiredPropertyException("GetUserGroupsResult", "userPoolId");
            }
            this.userPoolId = userPoolId;
            return this;
        }
        public GetUserGroupsResult build() {
            final var _resultValue = new GetUserGroupsResult();
            _resultValue.groups = groups;
            _resultValue.id = id;
            _resultValue.userPoolId = userPoolId;
            return _resultValue;
        }
    }
}
