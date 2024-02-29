// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.codedeploy;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.codedeploy.DeploymentConfigArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentConfigState;
import com.pulumi.aws.codedeploy.outputs.DeploymentConfigMinimumHealthyHosts;
import com.pulumi.aws.codedeploy.outputs.DeploymentConfigTrafficRoutingConfig;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a CodeDeploy deployment config for an application
 * 
 * ## Example Usage
 * ### Server Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.codedeploy.DeploymentConfig;
 * import com.pulumi.aws.codedeploy.DeploymentConfigArgs;
 * import com.pulumi.aws.codedeploy.inputs.DeploymentConfigMinimumHealthyHostsArgs;
 * import com.pulumi.aws.codedeploy.DeploymentGroup;
 * import com.pulumi.aws.codedeploy.DeploymentGroupArgs;
 * import com.pulumi.aws.codedeploy.inputs.DeploymentGroupEc2TagFilterArgs;
 * import com.pulumi.aws.codedeploy.inputs.DeploymentGroupTriggerConfigurationArgs;
 * import com.pulumi.aws.codedeploy.inputs.DeploymentGroupAutoRollbackConfigurationArgs;
 * import com.pulumi.aws.codedeploy.inputs.DeploymentGroupAlarmConfigurationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var fooDeploymentConfig = new DeploymentConfig(&#34;fooDeploymentConfig&#34;, DeploymentConfigArgs.builder()        
 *             .deploymentConfigName(&#34;test-deployment-config&#34;)
 *             .minimumHealthyHosts(DeploymentConfigMinimumHealthyHostsArgs.builder()
 *                 .type(&#34;HOST_COUNT&#34;)
 *                 .value(2)
 *                 .build())
 *             .build());
 * 
 *         var fooDeploymentGroup = new DeploymentGroup(&#34;fooDeploymentGroup&#34;, DeploymentGroupArgs.builder()        
 *             .appName(aws_codedeploy_app.foo_app().name())
 *             .deploymentGroupName(&#34;bar&#34;)
 *             .serviceRoleArn(aws_iam_role.foo_role().arn())
 *             .deploymentConfigName(fooDeploymentConfig.id())
 *             .ec2TagFilters(DeploymentGroupEc2TagFilterArgs.builder()
 *                 .key(&#34;filterkey&#34;)
 *                 .type(&#34;KEY_AND_VALUE&#34;)
 *                 .value(&#34;filtervalue&#34;)
 *                 .build())
 *             .triggerConfigurations(DeploymentGroupTriggerConfigurationArgs.builder()
 *                 .triggerEvents(&#34;DeploymentFailure&#34;)
 *                 .triggerName(&#34;foo-trigger&#34;)
 *                 .triggerTargetArn(&#34;foo-topic-arn&#34;)
 *                 .build())
 *             .autoRollbackConfiguration(DeploymentGroupAutoRollbackConfigurationArgs.builder()
 *                 .enabled(true)
 *                 .events(&#34;DEPLOYMENT_FAILURE&#34;)
 *                 .build())
 *             .alarmConfiguration(DeploymentGroupAlarmConfigurationArgs.builder()
 *                 .alarms(&#34;my-alarm-name&#34;)
 *                 .enabled(true)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Lambda Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.codedeploy.DeploymentConfig;
 * import com.pulumi.aws.codedeploy.DeploymentConfigArgs;
 * import com.pulumi.aws.codedeploy.inputs.DeploymentConfigTrafficRoutingConfigArgs;
 * import com.pulumi.aws.codedeploy.inputs.DeploymentConfigTrafficRoutingConfigTimeBasedLinearArgs;
 * import com.pulumi.aws.codedeploy.DeploymentGroup;
 * import com.pulumi.aws.codedeploy.DeploymentGroupArgs;
 * import com.pulumi.aws.codedeploy.inputs.DeploymentGroupAutoRollbackConfigurationArgs;
 * import com.pulumi.aws.codedeploy.inputs.DeploymentGroupAlarmConfigurationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var fooDeploymentConfig = new DeploymentConfig(&#34;fooDeploymentConfig&#34;, DeploymentConfigArgs.builder()        
 *             .deploymentConfigName(&#34;test-deployment-config&#34;)
 *             .computePlatform(&#34;Lambda&#34;)
 *             .trafficRoutingConfig(DeploymentConfigTrafficRoutingConfigArgs.builder()
 *                 .type(&#34;TimeBasedLinear&#34;)
 *                 .timeBasedLinear(DeploymentConfigTrafficRoutingConfigTimeBasedLinearArgs.builder()
 *                     .interval(10)
 *                     .percentage(10)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var fooDeploymentGroup = new DeploymentGroup(&#34;fooDeploymentGroup&#34;, DeploymentGroupArgs.builder()        
 *             .appName(aws_codedeploy_app.foo_app().name())
 *             .deploymentGroupName(&#34;bar&#34;)
 *             .serviceRoleArn(aws_iam_role.foo_role().arn())
 *             .deploymentConfigName(fooDeploymentConfig.id())
 *             .autoRollbackConfiguration(DeploymentGroupAutoRollbackConfigurationArgs.builder()
 *                 .enabled(true)
 *                 .events(&#34;DEPLOYMENT_STOP_ON_ALARM&#34;)
 *                 .build())
 *             .alarmConfiguration(DeploymentGroupAlarmConfigurationArgs.builder()
 *                 .alarms(&#34;my-alarm-name&#34;)
 *                 .enabled(true)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import CodeDeploy Deployment Configurations using the `deployment_config_name`. For example:
 * 
 * ```sh
 *  $ pulumi import aws:codedeploy/deploymentConfig:DeploymentConfig example my-deployment-config
 * ```
 * 
 */
@ResourceType(type="aws:codedeploy/deploymentConfig:DeploymentConfig")
public class DeploymentConfig extends com.pulumi.resources.CustomResource {
    /**
     * The ARN of the deployment config.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the deployment config.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The compute platform can be `Server`, `Lambda`, or `ECS`. Default is `Server`.
     * 
     */
    @Export(name="computePlatform", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> computePlatform;

    /**
     * @return The compute platform can be `Server`, `Lambda`, or `ECS`. Default is `Server`.
     * 
     */
    public Output<Optional<String>> computePlatform() {
        return Codegen.optional(this.computePlatform);
    }
    /**
     * The AWS Assigned deployment config id
     * 
     */
    @Export(name="deploymentConfigId", refs={String.class}, tree="[0]")
    private Output<String> deploymentConfigId;

    /**
     * @return The AWS Assigned deployment config id
     * 
     */
    public Output<String> deploymentConfigId() {
        return this.deploymentConfigId;
    }
    /**
     * The name of the deployment config.
     * 
     */
    @Export(name="deploymentConfigName", refs={String.class}, tree="[0]")
    private Output<String> deploymentConfigName;

    /**
     * @return The name of the deployment config.
     * 
     */
    public Output<String> deploymentConfigName() {
        return this.deploymentConfigName;
    }
    /**
     * A minimum_healthy_hosts block. Required for `Server` compute platform. Minimum Healthy Hosts are documented below.
     * 
     */
    @Export(name="minimumHealthyHosts", refs={DeploymentConfigMinimumHealthyHosts.class}, tree="[0]")
    private Output</* @Nullable */ DeploymentConfigMinimumHealthyHosts> minimumHealthyHosts;

    /**
     * @return A minimum_healthy_hosts block. Required for `Server` compute platform. Minimum Healthy Hosts are documented below.
     * 
     */
    public Output<Optional<DeploymentConfigMinimumHealthyHosts>> minimumHealthyHosts() {
        return Codegen.optional(this.minimumHealthyHosts);
    }
    /**
     * A traffic_routing_config block. Traffic Routing Config is documented below.
     * 
     */
    @Export(name="trafficRoutingConfig", refs={DeploymentConfigTrafficRoutingConfig.class}, tree="[0]")
    private Output</* @Nullable */ DeploymentConfigTrafficRoutingConfig> trafficRoutingConfig;

    /**
     * @return A traffic_routing_config block. Traffic Routing Config is documented below.
     * 
     */
    public Output<Optional<DeploymentConfigTrafficRoutingConfig>> trafficRoutingConfig() {
        return Codegen.optional(this.trafficRoutingConfig);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DeploymentConfig(String name) {
        this(name, DeploymentConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DeploymentConfig(String name, @Nullable DeploymentConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DeploymentConfig(String name, @Nullable DeploymentConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:codedeploy/deploymentConfig:DeploymentConfig", name, args == null ? DeploymentConfigArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DeploymentConfig(String name, Output<String> id, @Nullable DeploymentConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:codedeploy/deploymentConfig:DeploymentConfig", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static DeploymentConfig get(String name, Output<String> id, @Nullable DeploymentConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DeploymentConfig(name, id, state, options);
    }
}
