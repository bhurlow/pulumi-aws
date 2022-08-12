// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.codedeploy.inputs;

import com.pulumi.aws.codedeploy.inputs.DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DeploymentGroupBlueGreenDeploymentConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final DeploymentGroupBlueGreenDeploymentConfigArgs Empty = new DeploymentGroupBlueGreenDeploymentConfigArgs();

    /**
     * Information about the action to take when newly provisioned instances are ready to receive traffic in a blue/green deployment (documented below).
     * 
     */
    @Import(name="deploymentReadyOption")
    private @Nullable Output<DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionArgs> deploymentReadyOption;

    /**
     * @return Information about the action to take when newly provisioned instances are ready to receive traffic in a blue/green deployment (documented below).
     * 
     */
    public Optional<Output<DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionArgs>> deploymentReadyOption() {
        return Optional.ofNullable(this.deploymentReadyOption);
    }

    /**
     * Information about how instances are provisioned for a replacement environment in a blue/green deployment (documented below).
     * 
     */
    @Import(name="greenFleetProvisioningOption")
    private @Nullable Output<DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionArgs> greenFleetProvisioningOption;

    /**
     * @return Information about how instances are provisioned for a replacement environment in a blue/green deployment (documented below).
     * 
     */
    public Optional<Output<DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionArgs>> greenFleetProvisioningOption() {
        return Optional.ofNullable(this.greenFleetProvisioningOption);
    }

    /**
     * Information about whether to terminate instances in the original fleet during a blue/green deployment (documented below).
     * 
     */
    @Import(name="terminateBlueInstancesOnDeploymentSuccess")
    private @Nullable Output<DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessArgs> terminateBlueInstancesOnDeploymentSuccess;

    /**
     * @return Information about whether to terminate instances in the original fleet during a blue/green deployment (documented below).
     * 
     */
    public Optional<Output<DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessArgs>> terminateBlueInstancesOnDeploymentSuccess() {
        return Optional.ofNullable(this.terminateBlueInstancesOnDeploymentSuccess);
    }

    private DeploymentGroupBlueGreenDeploymentConfigArgs() {}

    private DeploymentGroupBlueGreenDeploymentConfigArgs(DeploymentGroupBlueGreenDeploymentConfigArgs $) {
        this.deploymentReadyOption = $.deploymentReadyOption;
        this.greenFleetProvisioningOption = $.greenFleetProvisioningOption;
        this.terminateBlueInstancesOnDeploymentSuccess = $.terminateBlueInstancesOnDeploymentSuccess;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DeploymentGroupBlueGreenDeploymentConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DeploymentGroupBlueGreenDeploymentConfigArgs $;

        public Builder() {
            $ = new DeploymentGroupBlueGreenDeploymentConfigArgs();
        }

        public Builder(DeploymentGroupBlueGreenDeploymentConfigArgs defaults) {
            $ = new DeploymentGroupBlueGreenDeploymentConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param deploymentReadyOption Information about the action to take when newly provisioned instances are ready to receive traffic in a blue/green deployment (documented below).
         * 
         * @return builder
         * 
         */
        public Builder deploymentReadyOption(@Nullable Output<DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionArgs> deploymentReadyOption) {
            $.deploymentReadyOption = deploymentReadyOption;
            return this;
        }

        /**
         * @param deploymentReadyOption Information about the action to take when newly provisioned instances are ready to receive traffic in a blue/green deployment (documented below).
         * 
         * @return builder
         * 
         */
        public Builder deploymentReadyOption(DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionArgs deploymentReadyOption) {
            return deploymentReadyOption(Output.of(deploymentReadyOption));
        }

        /**
         * @param greenFleetProvisioningOption Information about how instances are provisioned for a replacement environment in a blue/green deployment (documented below).
         * 
         * @return builder
         * 
         */
        public Builder greenFleetProvisioningOption(@Nullable Output<DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionArgs> greenFleetProvisioningOption) {
            $.greenFleetProvisioningOption = greenFleetProvisioningOption;
            return this;
        }

        /**
         * @param greenFleetProvisioningOption Information about how instances are provisioned for a replacement environment in a blue/green deployment (documented below).
         * 
         * @return builder
         * 
         */
        public Builder greenFleetProvisioningOption(DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionArgs greenFleetProvisioningOption) {
            return greenFleetProvisioningOption(Output.of(greenFleetProvisioningOption));
        }

        /**
         * @param terminateBlueInstancesOnDeploymentSuccess Information about whether to terminate instances in the original fleet during a blue/green deployment (documented below).
         * 
         * @return builder
         * 
         */
        public Builder terminateBlueInstancesOnDeploymentSuccess(@Nullable Output<DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessArgs> terminateBlueInstancesOnDeploymentSuccess) {
            $.terminateBlueInstancesOnDeploymentSuccess = terminateBlueInstancesOnDeploymentSuccess;
            return this;
        }

        /**
         * @param terminateBlueInstancesOnDeploymentSuccess Information about whether to terminate instances in the original fleet during a blue/green deployment (documented below).
         * 
         * @return builder
         * 
         */
        public Builder terminateBlueInstancesOnDeploymentSuccess(DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessArgs terminateBlueInstancesOnDeploymentSuccess) {
            return terminateBlueInstancesOnDeploymentSuccess(Output.of(terminateBlueInstancesOnDeploymentSuccess));
        }

        public DeploymentGroupBlueGreenDeploymentConfigArgs build() {
            return $;
        }
    }

}