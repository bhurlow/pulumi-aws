// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ecs.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class TaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfig {
    /**
     * @return The authorization credential option to use. The authorization credential options can be provided using either the Amazon Resource Name (ARN) of an AWS Secrets Manager secret or AWS Systems Manager Parameter Store parameter. The ARNs refer to the stored credentials.
     * 
     */
    private final String credentialsParameter;
    /**
     * @return A fully qualified domain name hosted by an AWS Directory Service Managed Microsoft AD (Active Directory) or self-hosted AD on Amazon EC2.
     * 
     */
    private final String domain;

    @CustomType.Constructor
    private TaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfig(
        @CustomType.Parameter("credentialsParameter") String credentialsParameter,
        @CustomType.Parameter("domain") String domain) {
        this.credentialsParameter = credentialsParameter;
        this.domain = domain;
    }

    /**
     * @return The authorization credential option to use. The authorization credential options can be provided using either the Amazon Resource Name (ARN) of an AWS Secrets Manager secret or AWS Systems Manager Parameter Store parameter. The ARNs refer to the stored credentials.
     * 
     */
    public String credentialsParameter() {
        return this.credentialsParameter;
    }
    /**
     * @return A fully qualified domain name hosted by an AWS Directory Service Managed Microsoft AD (Active Directory) or self-hosted AD on Amazon EC2.
     * 
     */
    public String domain() {
        return this.domain;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(TaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfig defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String credentialsParameter;
        private String domain;

        public Builder() {
    	      // Empty
        }

        public Builder(TaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.credentialsParameter = defaults.credentialsParameter;
    	      this.domain = defaults.domain;
        }

        public Builder credentialsParameter(String credentialsParameter) {
            this.credentialsParameter = Objects.requireNonNull(credentialsParameter);
            return this;
        }
        public Builder domain(String domain) {
            this.domain = Objects.requireNonNull(domain);
            return this;
        }        public TaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfig build() {
            return new TaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfig(credentialsParameter, domain);
        }
    }
}