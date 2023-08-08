// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.pipes.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PipeTargetParametersEcsTaskParametersNetworkConfigurationAwsVpcConfigurationArgs extends com.pulumi.resources.ResourceArgs {

    public static final PipeTargetParametersEcsTaskParametersNetworkConfigurationAwsVpcConfigurationArgs Empty = new PipeTargetParametersEcsTaskParametersNetworkConfigurationAwsVpcConfigurationArgs();

    /**
     * Specifies whether the task&#39;s elastic network interface receives a public IP address. You can specify ENABLED only when LaunchType in EcsParameters is set to FARGATE. Valid Values: ENABLED, DISABLED.
     * 
     */
    @Import(name="assignPublicIp")
    private @Nullable Output<String> assignPublicIp;

    /**
     * @return Specifies whether the task&#39;s elastic network interface receives a public IP address. You can specify ENABLED only when LaunchType in EcsParameters is set to FARGATE. Valid Values: ENABLED, DISABLED.
     * 
     */
    public Optional<Output<String>> assignPublicIp() {
        return Optional.ofNullable(this.assignPublicIp);
    }

    /**
     * List of security groups associated with the stream. These security groups must all be in the same VPC. You can specify as many as five security groups. If you do not specify a security group, the default security group for the VPC is used.
     * 
     */
    @Import(name="securityGroups")
    private @Nullable Output<List<String>> securityGroups;

    /**
     * @return List of security groups associated with the stream. These security groups must all be in the same VPC. You can specify as many as five security groups. If you do not specify a security group, the default security group for the VPC is used.
     * 
     */
    public Optional<Output<List<String>>> securityGroups() {
        return Optional.ofNullable(this.securityGroups);
    }

    /**
     * List of the subnets associated with the stream. These subnets must all be in the same VPC. You can specify as many as 16 subnets.
     * 
     */
    @Import(name="subnets")
    private @Nullable Output<List<String>> subnets;

    /**
     * @return List of the subnets associated with the stream. These subnets must all be in the same VPC. You can specify as many as 16 subnets.
     * 
     */
    public Optional<Output<List<String>>> subnets() {
        return Optional.ofNullable(this.subnets);
    }

    private PipeTargetParametersEcsTaskParametersNetworkConfigurationAwsVpcConfigurationArgs() {}

    private PipeTargetParametersEcsTaskParametersNetworkConfigurationAwsVpcConfigurationArgs(PipeTargetParametersEcsTaskParametersNetworkConfigurationAwsVpcConfigurationArgs $) {
        this.assignPublicIp = $.assignPublicIp;
        this.securityGroups = $.securityGroups;
        this.subnets = $.subnets;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PipeTargetParametersEcsTaskParametersNetworkConfigurationAwsVpcConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PipeTargetParametersEcsTaskParametersNetworkConfigurationAwsVpcConfigurationArgs $;

        public Builder() {
            $ = new PipeTargetParametersEcsTaskParametersNetworkConfigurationAwsVpcConfigurationArgs();
        }

        public Builder(PipeTargetParametersEcsTaskParametersNetworkConfigurationAwsVpcConfigurationArgs defaults) {
            $ = new PipeTargetParametersEcsTaskParametersNetworkConfigurationAwsVpcConfigurationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param assignPublicIp Specifies whether the task&#39;s elastic network interface receives a public IP address. You can specify ENABLED only when LaunchType in EcsParameters is set to FARGATE. Valid Values: ENABLED, DISABLED.
         * 
         * @return builder
         * 
         */
        public Builder assignPublicIp(@Nullable Output<String> assignPublicIp) {
            $.assignPublicIp = assignPublicIp;
            return this;
        }

        /**
         * @param assignPublicIp Specifies whether the task&#39;s elastic network interface receives a public IP address. You can specify ENABLED only when LaunchType in EcsParameters is set to FARGATE. Valid Values: ENABLED, DISABLED.
         * 
         * @return builder
         * 
         */
        public Builder assignPublicIp(String assignPublicIp) {
            return assignPublicIp(Output.of(assignPublicIp));
        }

        /**
         * @param securityGroups List of security groups associated with the stream. These security groups must all be in the same VPC. You can specify as many as five security groups. If you do not specify a security group, the default security group for the VPC is used.
         * 
         * @return builder
         * 
         */
        public Builder securityGroups(@Nullable Output<List<String>> securityGroups) {
            $.securityGroups = securityGroups;
            return this;
        }

        /**
         * @param securityGroups List of security groups associated with the stream. These security groups must all be in the same VPC. You can specify as many as five security groups. If you do not specify a security group, the default security group for the VPC is used.
         * 
         * @return builder
         * 
         */
        public Builder securityGroups(List<String> securityGroups) {
            return securityGroups(Output.of(securityGroups));
        }

        /**
         * @param securityGroups List of security groups associated with the stream. These security groups must all be in the same VPC. You can specify as many as five security groups. If you do not specify a security group, the default security group for the VPC is used.
         * 
         * @return builder
         * 
         */
        public Builder securityGroups(String... securityGroups) {
            return securityGroups(List.of(securityGroups));
        }

        /**
         * @param subnets List of the subnets associated with the stream. These subnets must all be in the same VPC. You can specify as many as 16 subnets.
         * 
         * @return builder
         * 
         */
        public Builder subnets(@Nullable Output<List<String>> subnets) {
            $.subnets = subnets;
            return this;
        }

        /**
         * @param subnets List of the subnets associated with the stream. These subnets must all be in the same VPC. You can specify as many as 16 subnets.
         * 
         * @return builder
         * 
         */
        public Builder subnets(List<String> subnets) {
            return subnets(Output.of(subnets));
        }

        /**
         * @param subnets List of the subnets associated with the stream. These subnets must all be in the same VPC. You can specify as many as 16 subnets.
         * 
         * @return builder
         * 
         */
        public Builder subnets(String... subnets) {
            return subnets(List.of(subnets));
        }

        public PipeTargetParametersEcsTaskParametersNetworkConfigurationAwsVpcConfigurationArgs build() {
            return $;
        }
    }

}
