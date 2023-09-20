// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.gamelift.inputs;

import com.pulumi.aws.gamelift.inputs.GameServerGroupAutoScalingPolicyArgs;
import com.pulumi.aws.gamelift.inputs.GameServerGroupInstanceDefinitionArgs;
import com.pulumi.aws.gamelift.inputs.GameServerGroupLaunchTemplateArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GameServerGroupState extends com.pulumi.resources.ResourceArgs {

    public static final GameServerGroupState Empty = new GameServerGroupState();

    /**
     * The ARN of the GameLift Game Server Group.
     * 
     */
    @Import(name="arn")
    private @Nullable Output<String> arn;

    /**
     * @return The ARN of the GameLift Game Server Group.
     * 
     */
    public Optional<Output<String>> arn() {
        return Optional.ofNullable(this.arn);
    }

    /**
     * The ARN of the created EC2 Auto Scaling group.
     * 
     */
    @Import(name="autoScalingGroupArn")
    private @Nullable Output<String> autoScalingGroupArn;

    /**
     * @return The ARN of the created EC2 Auto Scaling group.
     * 
     */
    public Optional<Output<String>> autoScalingGroupArn() {
        return Optional.ofNullable(this.autoScalingGroupArn);
    }

    @Import(name="autoScalingPolicy")
    private @Nullable Output<GameServerGroupAutoScalingPolicyArgs> autoScalingPolicy;

    public Optional<Output<GameServerGroupAutoScalingPolicyArgs>> autoScalingPolicy() {
        return Optional.ofNullable(this.autoScalingPolicy);
    }

    /**
     * Indicates how GameLift FleetIQ balances the use of Spot Instances and On-Demand Instances.
     * Valid values: `SPOT_ONLY`, `SPOT_PREFERRED`, `ON_DEMAND_ONLY`. Defaults to `SPOT_PREFERRED`.
     * 
     */
    @Import(name="balancingStrategy")
    private @Nullable Output<String> balancingStrategy;

    /**
     * @return Indicates how GameLift FleetIQ balances the use of Spot Instances and On-Demand Instances.
     * Valid values: `SPOT_ONLY`, `SPOT_PREFERRED`, `ON_DEMAND_ONLY`. Defaults to `SPOT_PREFERRED`.
     * 
     */
    public Optional<Output<String>> balancingStrategy() {
        return Optional.ofNullable(this.balancingStrategy);
    }

    /**
     * Name of the game server group.
     * This value is used to generate unique ARN identifiers for the EC2 Auto Scaling group and the GameLift FleetIQ game server group.
     * 
     */
    @Import(name="gameServerGroupName")
    private @Nullable Output<String> gameServerGroupName;

    /**
     * @return Name of the game server group.
     * This value is used to generate unique ARN identifiers for the EC2 Auto Scaling group and the GameLift FleetIQ game server group.
     * 
     */
    public Optional<Output<String>> gameServerGroupName() {
        return Optional.ofNullable(this.gameServerGroupName);
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
    @Import(name="gameServerProtectionPolicy")
    private @Nullable Output<String> gameServerProtectionPolicy;

    /**
     * @return Indicates whether instances in the game server group are protected from early termination.
     * Unprotected instances that have active game servers running might be terminated during a scale-down event,
     * causing players to be dropped from the game.
     * Protected instances cannot be terminated while there are active game servers running except in the event
     * of a forced game server group deletion.
     * Valid values: `NO_PROTECTION`, `FULL_PROTECTION`. Defaults to `NO_PROTECTION`.
     * 
     */
    public Optional<Output<String>> gameServerProtectionPolicy() {
        return Optional.ofNullable(this.gameServerProtectionPolicy);
    }

    @Import(name="instanceDefinitions")
    private @Nullable Output<List<GameServerGroupInstanceDefinitionArgs>> instanceDefinitions;

    public Optional<Output<List<GameServerGroupInstanceDefinitionArgs>>> instanceDefinitions() {
        return Optional.ofNullable(this.instanceDefinitions);
    }

    @Import(name="launchTemplate")
    private @Nullable Output<GameServerGroupLaunchTemplateArgs> launchTemplate;

    public Optional<Output<GameServerGroupLaunchTemplateArgs>> launchTemplate() {
        return Optional.ofNullable(this.launchTemplate);
    }

    /**
     * The maximum number of instances allowed in the EC2 Auto Scaling group.
     * During automatic scaling events, GameLift FleetIQ and EC2 do not scale up the group above this maximum.
     * 
     */
    @Import(name="maxSize")
    private @Nullable Output<Integer> maxSize;

    /**
     * @return The maximum number of instances allowed in the EC2 Auto Scaling group.
     * During automatic scaling events, GameLift FleetIQ and EC2 do not scale up the group above this maximum.
     * 
     */
    public Optional<Output<Integer>> maxSize() {
        return Optional.ofNullable(this.maxSize);
    }

    /**
     * The minimum number of instances allowed in the EC2 Auto Scaling group.
     * During automatic scaling events, GameLift FleetIQ and EC2 do not scale down the group below this minimum.
     * 
     */
    @Import(name="minSize")
    private @Nullable Output<Integer> minSize;

    /**
     * @return The minimum number of instances allowed in the EC2 Auto Scaling group.
     * During automatic scaling events, GameLift FleetIQ and EC2 do not scale down the group below this minimum.
     * 
     */
    public Optional<Output<Integer>> minSize() {
        return Optional.ofNullable(this.minSize);
    }

    /**
     * ARN for an IAM role that allows Amazon GameLift to access your EC2 Auto Scaling groups.
     * 
     */
    @Import(name="roleArn")
    private @Nullable Output<String> roleArn;

    /**
     * @return ARN for an IAM role that allows Amazon GameLift to access your EC2 Auto Scaling groups.
     * 
     */
    public Optional<Output<String>> roleArn() {
        return Optional.ofNullable(this.roleArn);
    }

    /**
     * Key-value map of resource tags
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Import(name="tagsAll")
    private @Nullable Output<Map<String,String>> tagsAll;

    /**
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    public Optional<Output<Map<String,String>>> tagsAll() {
        return Optional.ofNullable(this.tagsAll);
    }

    /**
     * A list of VPC subnets to use with instances in the game server group.
     * By default, all GameLift FleetIQ-supported Availability Zones are used.
     * 
     */
    @Import(name="vpcSubnets")
    private @Nullable Output<List<String>> vpcSubnets;

    /**
     * @return A list of VPC subnets to use with instances in the game server group.
     * By default, all GameLift FleetIQ-supported Availability Zones are used.
     * 
     */
    public Optional<Output<List<String>>> vpcSubnets() {
        return Optional.ofNullable(this.vpcSubnets);
    }

    private GameServerGroupState() {}

    private GameServerGroupState(GameServerGroupState $) {
        this.arn = $.arn;
        this.autoScalingGroupArn = $.autoScalingGroupArn;
        this.autoScalingPolicy = $.autoScalingPolicy;
        this.balancingStrategy = $.balancingStrategy;
        this.gameServerGroupName = $.gameServerGroupName;
        this.gameServerProtectionPolicy = $.gameServerProtectionPolicy;
        this.instanceDefinitions = $.instanceDefinitions;
        this.launchTemplate = $.launchTemplate;
        this.maxSize = $.maxSize;
        this.minSize = $.minSize;
        this.roleArn = $.roleArn;
        this.tags = $.tags;
        this.tagsAll = $.tagsAll;
        this.vpcSubnets = $.vpcSubnets;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GameServerGroupState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GameServerGroupState $;

        public Builder() {
            $ = new GameServerGroupState();
        }

        public Builder(GameServerGroupState defaults) {
            $ = new GameServerGroupState(Objects.requireNonNull(defaults));
        }

        /**
         * @param arn The ARN of the GameLift Game Server Group.
         * 
         * @return builder
         * 
         */
        public Builder arn(@Nullable Output<String> arn) {
            $.arn = arn;
            return this;
        }

        /**
         * @param arn The ARN of the GameLift Game Server Group.
         * 
         * @return builder
         * 
         */
        public Builder arn(String arn) {
            return arn(Output.of(arn));
        }

        /**
         * @param autoScalingGroupArn The ARN of the created EC2 Auto Scaling group.
         * 
         * @return builder
         * 
         */
        public Builder autoScalingGroupArn(@Nullable Output<String> autoScalingGroupArn) {
            $.autoScalingGroupArn = autoScalingGroupArn;
            return this;
        }

        /**
         * @param autoScalingGroupArn The ARN of the created EC2 Auto Scaling group.
         * 
         * @return builder
         * 
         */
        public Builder autoScalingGroupArn(String autoScalingGroupArn) {
            return autoScalingGroupArn(Output.of(autoScalingGroupArn));
        }

        public Builder autoScalingPolicy(@Nullable Output<GameServerGroupAutoScalingPolicyArgs> autoScalingPolicy) {
            $.autoScalingPolicy = autoScalingPolicy;
            return this;
        }

        public Builder autoScalingPolicy(GameServerGroupAutoScalingPolicyArgs autoScalingPolicy) {
            return autoScalingPolicy(Output.of(autoScalingPolicy));
        }

        /**
         * @param balancingStrategy Indicates how GameLift FleetIQ balances the use of Spot Instances and On-Demand Instances.
         * Valid values: `SPOT_ONLY`, `SPOT_PREFERRED`, `ON_DEMAND_ONLY`. Defaults to `SPOT_PREFERRED`.
         * 
         * @return builder
         * 
         */
        public Builder balancingStrategy(@Nullable Output<String> balancingStrategy) {
            $.balancingStrategy = balancingStrategy;
            return this;
        }

        /**
         * @param balancingStrategy Indicates how GameLift FleetIQ balances the use of Spot Instances and On-Demand Instances.
         * Valid values: `SPOT_ONLY`, `SPOT_PREFERRED`, `ON_DEMAND_ONLY`. Defaults to `SPOT_PREFERRED`.
         * 
         * @return builder
         * 
         */
        public Builder balancingStrategy(String balancingStrategy) {
            return balancingStrategy(Output.of(balancingStrategy));
        }

        /**
         * @param gameServerGroupName Name of the game server group.
         * This value is used to generate unique ARN identifiers for the EC2 Auto Scaling group and the GameLift FleetIQ game server group.
         * 
         * @return builder
         * 
         */
        public Builder gameServerGroupName(@Nullable Output<String> gameServerGroupName) {
            $.gameServerGroupName = gameServerGroupName;
            return this;
        }

        /**
         * @param gameServerGroupName Name of the game server group.
         * This value is used to generate unique ARN identifiers for the EC2 Auto Scaling group and the GameLift FleetIQ game server group.
         * 
         * @return builder
         * 
         */
        public Builder gameServerGroupName(String gameServerGroupName) {
            return gameServerGroupName(Output.of(gameServerGroupName));
        }

        /**
         * @param gameServerProtectionPolicy Indicates whether instances in the game server group are protected from early termination.
         * Unprotected instances that have active game servers running might be terminated during a scale-down event,
         * causing players to be dropped from the game.
         * Protected instances cannot be terminated while there are active game servers running except in the event
         * of a forced game server group deletion.
         * Valid values: `NO_PROTECTION`, `FULL_PROTECTION`. Defaults to `NO_PROTECTION`.
         * 
         * @return builder
         * 
         */
        public Builder gameServerProtectionPolicy(@Nullable Output<String> gameServerProtectionPolicy) {
            $.gameServerProtectionPolicy = gameServerProtectionPolicy;
            return this;
        }

        /**
         * @param gameServerProtectionPolicy Indicates whether instances in the game server group are protected from early termination.
         * Unprotected instances that have active game servers running might be terminated during a scale-down event,
         * causing players to be dropped from the game.
         * Protected instances cannot be terminated while there are active game servers running except in the event
         * of a forced game server group deletion.
         * Valid values: `NO_PROTECTION`, `FULL_PROTECTION`. Defaults to `NO_PROTECTION`.
         * 
         * @return builder
         * 
         */
        public Builder gameServerProtectionPolicy(String gameServerProtectionPolicy) {
            return gameServerProtectionPolicy(Output.of(gameServerProtectionPolicy));
        }

        public Builder instanceDefinitions(@Nullable Output<List<GameServerGroupInstanceDefinitionArgs>> instanceDefinitions) {
            $.instanceDefinitions = instanceDefinitions;
            return this;
        }

        public Builder instanceDefinitions(List<GameServerGroupInstanceDefinitionArgs> instanceDefinitions) {
            return instanceDefinitions(Output.of(instanceDefinitions));
        }

        public Builder instanceDefinitions(GameServerGroupInstanceDefinitionArgs... instanceDefinitions) {
            return instanceDefinitions(List.of(instanceDefinitions));
        }

        public Builder launchTemplate(@Nullable Output<GameServerGroupLaunchTemplateArgs> launchTemplate) {
            $.launchTemplate = launchTemplate;
            return this;
        }

        public Builder launchTemplate(GameServerGroupLaunchTemplateArgs launchTemplate) {
            return launchTemplate(Output.of(launchTemplate));
        }

        /**
         * @param maxSize The maximum number of instances allowed in the EC2 Auto Scaling group.
         * During automatic scaling events, GameLift FleetIQ and EC2 do not scale up the group above this maximum.
         * 
         * @return builder
         * 
         */
        public Builder maxSize(@Nullable Output<Integer> maxSize) {
            $.maxSize = maxSize;
            return this;
        }

        /**
         * @param maxSize The maximum number of instances allowed in the EC2 Auto Scaling group.
         * During automatic scaling events, GameLift FleetIQ and EC2 do not scale up the group above this maximum.
         * 
         * @return builder
         * 
         */
        public Builder maxSize(Integer maxSize) {
            return maxSize(Output.of(maxSize));
        }

        /**
         * @param minSize The minimum number of instances allowed in the EC2 Auto Scaling group.
         * During automatic scaling events, GameLift FleetIQ and EC2 do not scale down the group below this minimum.
         * 
         * @return builder
         * 
         */
        public Builder minSize(@Nullable Output<Integer> minSize) {
            $.minSize = minSize;
            return this;
        }

        /**
         * @param minSize The minimum number of instances allowed in the EC2 Auto Scaling group.
         * During automatic scaling events, GameLift FleetIQ and EC2 do not scale down the group below this minimum.
         * 
         * @return builder
         * 
         */
        public Builder minSize(Integer minSize) {
            return minSize(Output.of(minSize));
        }

        /**
         * @param roleArn ARN for an IAM role that allows Amazon GameLift to access your EC2 Auto Scaling groups.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(@Nullable Output<String> roleArn) {
            $.roleArn = roleArn;
            return this;
        }

        /**
         * @param roleArn ARN for an IAM role that allows Amazon GameLift to access your EC2 Auto Scaling groups.
         * 
         * @return builder
         * 
         */
        public Builder roleArn(String roleArn) {
            return roleArn(Output.of(roleArn));
        }

        /**
         * @param tags Key-value map of resource tags
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Key-value map of resource tags
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(@Nullable Output<Map<String,String>> tagsAll) {
            $.tagsAll = tagsAll;
            return this;
        }

        /**
         * @return builder
         * 
         * @deprecated
         * Please use `tags` instead.
         * 
         */
        @Deprecated /* Please use `tags` instead. */
        public Builder tagsAll(Map<String,String> tagsAll) {
            return tagsAll(Output.of(tagsAll));
        }

        /**
         * @param vpcSubnets A list of VPC subnets to use with instances in the game server group.
         * By default, all GameLift FleetIQ-supported Availability Zones are used.
         * 
         * @return builder
         * 
         */
        public Builder vpcSubnets(@Nullable Output<List<String>> vpcSubnets) {
            $.vpcSubnets = vpcSubnets;
            return this;
        }

        /**
         * @param vpcSubnets A list of VPC subnets to use with instances in the game server group.
         * By default, all GameLift FleetIQ-supported Availability Zones are used.
         * 
         * @return builder
         * 
         */
        public Builder vpcSubnets(List<String> vpcSubnets) {
            return vpcSubnets(Output.of(vpcSubnets));
        }

        /**
         * @param vpcSubnets A list of VPC subnets to use with instances in the game server group.
         * By default, all GameLift FleetIQ-supported Availability Zones are used.
         * 
         * @return builder
         * 
         */
        public Builder vpcSubnets(String... vpcSubnets) {
            return vpcSubnets(List.of(vpcSubnets));
        }

        public GameServerGroupState build() {
            return $;
        }
    }

}
