// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.gamelift;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.gamelift.GameServerGroupArgs;
import com.pulumi.aws.gamelift.inputs.GameServerGroupState;
import com.pulumi.aws.gamelift.outputs.GameServerGroupAutoScalingPolicy;
import com.pulumi.aws.gamelift.outputs.GameServerGroupInstanceDefinition;
import com.pulumi.aws.gamelift.outputs.GameServerGroupLaunchTemplate;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an GameLift Game Server Group resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.gamelift.GameServerGroup;
 * import com.pulumi.aws.gamelift.GameServerGroupArgs;
 * import com.pulumi.aws.gamelift.inputs.GameServerGroupInstanceDefinitionArgs;
 * import com.pulumi.aws.gamelift.inputs.GameServerGroupLaunchTemplateArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var example = new GameServerGroup(&#34;example&#34;, GameServerGroupArgs.builder()        
 *             .gameServerGroupName(&#34;example&#34;)
 *             .instanceDefinitions(            
 *                 GameServerGroupInstanceDefinitionArgs.builder()
 *                     .instanceType(&#34;c5.large&#34;)
 *                     .build(),
 *                 GameServerGroupInstanceDefinitionArgs.builder()
 *                     .instanceType(&#34;c5a.large&#34;)
 *                     .build())
 *             .launchTemplate(GameServerGroupLaunchTemplateArgs.builder()
 *                 .id(aws_launch_template.example().id())
 *                 .build())
 *             .maxSize(1)
 *             .minSize(1)
 *             .roleArn(aws_iam_role.example().arn())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(aws_iam_role_policy_attachment.example())
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * Full usage:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.gamelift.GameServerGroup;
 * import com.pulumi.aws.gamelift.GameServerGroupArgs;
 * import com.pulumi.aws.gamelift.inputs.GameServerGroupAutoScalingPolicyArgs;
 * import com.pulumi.aws.gamelift.inputs.GameServerGroupAutoScalingPolicyTargetTrackingConfigurationArgs;
 * import com.pulumi.aws.gamelift.inputs.GameServerGroupInstanceDefinitionArgs;
 * import com.pulumi.aws.gamelift.inputs.GameServerGroupLaunchTemplateArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var example = new GameServerGroup(&#34;example&#34;, GameServerGroupArgs.builder()        
 *             .autoScalingPolicy(GameServerGroupAutoScalingPolicyArgs.builder()
 *                 .estimatedInstanceWarmup(60)
 *                 .targetTrackingConfiguration(GameServerGroupAutoScalingPolicyTargetTrackingConfigurationArgs.builder()
 *                     .targetValue(75)
 *                     .build())
 *                 .build())
 *             .balancingStrategy(&#34;SPOT_ONLY&#34;)
 *             .gameServerGroupName(&#34;example&#34;)
 *             .gameServerProtectionPolicy(&#34;FULL_PROTECTION&#34;)
 *             .instanceDefinitions(            
 *                 GameServerGroupInstanceDefinitionArgs.builder()
 *                     .instanceType(&#34;c5.large&#34;)
 *                     .weightedCapacity(&#34;1&#34;)
 *                     .build(),
 *                 GameServerGroupInstanceDefinitionArgs.builder()
 *                     .instanceType(&#34;c5.2xlarge&#34;)
 *                     .weightedCapacity(&#34;2&#34;)
 *                     .build())
 *             .launchTemplate(GameServerGroupLaunchTemplateArgs.builder()
 *                 .id(aws_launch_template.example().id())
 *                 .version(&#34;1&#34;)
 *                 .build())
 *             .maxSize(1)
 *             .minSize(1)
 *             .roleArn(aws_iam_role.example().arn())
 *             .tags(Map.of(&#34;Name&#34;, &#34;example&#34;))
 *             .vpcSubnets(            
 *                 &#34;subnet-12345678&#34;,
 *                 &#34;subnet-23456789&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(aws_iam_role_policy_attachment.example())
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ### Example IAM Role for GameLift Game Server Group
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.AwsFunctions;
 * import com.pulumi.aws.inputs.GetPartitionArgs;
 * import com.pulumi.aws.iam.IamFunctions;
 * import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
 * import com.pulumi.aws.iam.Role;
 * import com.pulumi.aws.iam.RoleArgs;
 * import com.pulumi.aws.iam.RolePolicyAttachment;
 * import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
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
 *         final var current = AwsFunctions.getPartition();
 * 
 *         final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
 *             .statements(GetPolicyDocumentStatementArgs.builder()
 *                 .effect(&#34;Allow&#34;)
 *                 .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
 *                     .type(&#34;Service&#34;)
 *                     .identifiers(                    
 *                         &#34;autoscaling.amazonaws.com&#34;,
 *                         &#34;gamelift.amazonaws.com&#34;)
 *                     .build())
 *                 .actions(&#34;sts:AssumeRole&#34;)
 *                 .build())
 *             .build());
 * 
 *         var exampleRole = new Role(&#34;exampleRole&#34;, RoleArgs.builder()        
 *             .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -&gt; getPolicyDocumentResult.json()))
 *             .build());
 * 
 *         var exampleRolePolicyAttachment = new RolePolicyAttachment(&#34;exampleRolePolicyAttachment&#34;, RolePolicyAttachmentArgs.builder()        
 *             .policyArn(String.format(&#34;arn:%s:iam::aws:policy/GameLiftGameServerGroupPolicy&#34;, current.applyValue(getPartitionResult -&gt; getPartitionResult.partition())))
 *             .role(exampleRole.name())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * In TODO v1.5.0 and later, use an `import` block to import GameLift Game Server Group using the `name`. For exampleterraform import {
 * 
 *  to = aws_gamelift_game_server_group.example
 * 
 *  id = &#34;example&#34; } Using `TODO import`, import GameLift Game Server Group using the `name`. For exampleconsole % TODO import aws_gamelift_game_server_group.example example
 * 
 */
@ResourceType(type="aws:gamelift/gameServerGroup:GameServerGroup")
public class GameServerGroup extends com.pulumi.resources.CustomResource {
    /**
     * The ARN of the GameLift Game Server Group.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the GameLift Game Server Group.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The ARN of the created EC2 Auto Scaling group.
     * 
     */
    @Export(name="autoScalingGroupArn", refs={String.class}, tree="[0]")
    private Output<String> autoScalingGroupArn;

    /**
     * @return The ARN of the created EC2 Auto Scaling group.
     * 
     */
    public Output<String> autoScalingGroupArn() {
        return this.autoScalingGroupArn;
    }
    @Export(name="autoScalingPolicy", refs={GameServerGroupAutoScalingPolicy.class}, tree="[0]")
    private Output</* @Nullable */ GameServerGroupAutoScalingPolicy> autoScalingPolicy;

    public Output<Optional<GameServerGroupAutoScalingPolicy>> autoScalingPolicy() {
        return Codegen.optional(this.autoScalingPolicy);
    }
    /**
     * Indicates how GameLift FleetIQ balances the use of Spot Instances and On-Demand Instances.
     * Valid values: `SPOT_ONLY`, `SPOT_PREFERRED`, `ON_DEMAND_ONLY`. Defaults to `SPOT_PREFERRED`.
     * 
     */
    @Export(name="balancingStrategy", refs={String.class}, tree="[0]")
    private Output<String> balancingStrategy;

    /**
     * @return Indicates how GameLift FleetIQ balances the use of Spot Instances and On-Demand Instances.
     * Valid values: `SPOT_ONLY`, `SPOT_PREFERRED`, `ON_DEMAND_ONLY`. Defaults to `SPOT_PREFERRED`.
     * 
     */
    public Output<String> balancingStrategy() {
        return this.balancingStrategy;
    }
    /**
     * Name of the game server group.
     * This value is used to generate unique ARN identifiers for the EC2 Auto Scaling group and the GameLift FleetIQ game server group.
     * 
     */
    @Export(name="gameServerGroupName", refs={String.class}, tree="[0]")
    private Output<String> gameServerGroupName;

    /**
     * @return Name of the game server group.
     * This value is used to generate unique ARN identifiers for the EC2 Auto Scaling group and the GameLift FleetIQ game server group.
     * 
     */
    public Output<String> gameServerGroupName() {
        return this.gameServerGroupName;
    }
    /**
     * Indicates whether instances in the game server group are protected from early termination.
     * Unprotected instances that have active game servers running might be terminated during a scale-down event,
     * causing players to be dropped from the game.
     * Protected instances cannot be terminated while there are active game servers running except in the event
     * of a forced game server group deletion.
     * Valid values: `NO_PROTECTION`, `FULL_PROTECTION`. Defaults to `NO_PROTECTION`.
     * 
     */
    @Export(name="gameServerProtectionPolicy", refs={String.class}, tree="[0]")
    private Output<String> gameServerProtectionPolicy;

    /**
     * @return Indicates whether instances in the game server group are protected from early termination.
     * Unprotected instances that have active game servers running might be terminated during a scale-down event,
     * causing players to be dropped from the game.
     * Protected instances cannot be terminated while there are active game servers running except in the event
     * of a forced game server group deletion.
     * Valid values: `NO_PROTECTION`, `FULL_PROTECTION`. Defaults to `NO_PROTECTION`.
     * 
     */
    public Output<String> gameServerProtectionPolicy() {
        return this.gameServerProtectionPolicy;
    }
    @Export(name="instanceDefinitions", refs={List.class,GameServerGroupInstanceDefinition.class}, tree="[0,1]")
    private Output<List<GameServerGroupInstanceDefinition>> instanceDefinitions;

    public Output<List<GameServerGroupInstanceDefinition>> instanceDefinitions() {
        return this.instanceDefinitions;
    }
    @Export(name="launchTemplate", refs={GameServerGroupLaunchTemplate.class}, tree="[0]")
    private Output<GameServerGroupLaunchTemplate> launchTemplate;

    public Output<GameServerGroupLaunchTemplate> launchTemplate() {
        return this.launchTemplate;
    }
    /**
     * The maximum number of instances allowed in the EC2 Auto Scaling group.
     * During automatic scaling events, GameLift FleetIQ and EC2 do not scale up the group above this maximum.
     * 
     */
    @Export(name="maxSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxSize;

    /**
     * @return The maximum number of instances allowed in the EC2 Auto Scaling group.
     * During automatic scaling events, GameLift FleetIQ and EC2 do not scale up the group above this maximum.
     * 
     */
    public Output<Integer> maxSize() {
        return this.maxSize;
    }
    /**
     * The minimum number of instances allowed in the EC2 Auto Scaling group.
     * During automatic scaling events, GameLift FleetIQ and EC2 do not scale down the group below this minimum.
     * 
     */
    @Export(name="minSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> minSize;

    /**
     * @return The minimum number of instances allowed in the EC2 Auto Scaling group.
     * During automatic scaling events, GameLift FleetIQ and EC2 do not scale down the group below this minimum.
     * 
     */
    public Output<Integer> minSize() {
        return this.minSize;
    }
    /**
     * ARN for an IAM role that allows Amazon GameLift to access your EC2 Auto Scaling groups.
     * 
     */
    @Export(name="roleArn", refs={String.class}, tree="[0]")
    private Output<String> roleArn;

    /**
     * @return ARN for an IAM role that allows Amazon GameLift to access your EC2 Auto Scaling groups.
     * 
     */
    public Output<String> roleArn() {
        return this.roleArn;
    }
    /**
     * Key-value map of resource tags
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Export(name="tagsAll", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> tagsAll;

    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }
    /**
     * A list of VPC subnets to use with instances in the game server group.
     * By default, all GameLift FleetIQ-supported Availability Zones are used.
     * 
     */
    @Export(name="vpcSubnets", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> vpcSubnets;

    /**
     * @return A list of VPC subnets to use with instances in the game server group.
     * By default, all GameLift FleetIQ-supported Availability Zones are used.
     * 
     */
    public Output<Optional<List<String>>> vpcSubnets() {
        return Codegen.optional(this.vpcSubnets);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GameServerGroup(String name) {
        this(name, GameServerGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GameServerGroup(String name, GameServerGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GameServerGroup(String name, GameServerGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:gamelift/gameServerGroup:GameServerGroup", name, args == null ? GameServerGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GameServerGroup(String name, Output<String> id, @Nullable GameServerGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:gamelift/gameServerGroup:GameServerGroup", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "tagsAll"
            ))
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
    public static GameServerGroup get(String name, Output<String> id, @Nullable GameServerGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GameServerGroup(name, id, state, options);
    }
}
