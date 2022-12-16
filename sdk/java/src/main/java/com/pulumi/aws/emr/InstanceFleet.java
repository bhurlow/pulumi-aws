// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.emr;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.emr.InstanceFleetArgs;
import com.pulumi.aws.emr.inputs.InstanceFleetState;
import com.pulumi.aws.emr.outputs.InstanceFleetInstanceTypeConfig;
import com.pulumi.aws.emr.outputs.InstanceFleetLaunchSpecifications;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an Elastic MapReduce Cluster Instance Fleet configuration.
 * See [Amazon Elastic MapReduce Documentation](https://aws.amazon.com/documentation/emr/) for more information.
 * 
 * &gt; **NOTE:** At this time, Instance Fleets cannot be destroyed through the API nor
 * web interface. Instance Fleets are destroyed when the EMR Cluster is destroyed.
 * the provider will resize any Instance Fleet to zero when destroying the resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.emr.InstanceFleet;
 * import com.pulumi.aws.emr.InstanceFleetArgs;
 * import com.pulumi.aws.emr.inputs.InstanceFleetInstanceTypeConfigArgs;
 * import com.pulumi.aws.emr.inputs.InstanceFleetLaunchSpecificationsArgs;
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
 *         var task = new InstanceFleet(&#34;task&#34;, InstanceFleetArgs.builder()        
 *             .clusterId(aws_emr_cluster.cluster().id())
 *             .instanceTypeConfigs(            
 *                 InstanceFleetInstanceTypeConfigArgs.builder()
 *                     .bidPriceAsPercentageOfOnDemandPrice(100)
 *                     .ebsConfigs(InstanceFleetInstanceTypeConfigEbsConfigArgs.builder()
 *                         .size(100)
 *                         .type(&#34;gp2&#34;)
 *                         .volumesPerInstance(1)
 *                         .build())
 *                     .instanceType(&#34;m4.xlarge&#34;)
 *                     .weightedCapacity(1)
 *                     .build(),
 *                 InstanceFleetInstanceTypeConfigArgs.builder()
 *                     .bidPriceAsPercentageOfOnDemandPrice(100)
 *                     .ebsConfigs(InstanceFleetInstanceTypeConfigEbsConfigArgs.builder()
 *                         .size(100)
 *                         .type(&#34;gp2&#34;)
 *                         .volumesPerInstance(1)
 *                         .build())
 *                     .instanceType(&#34;m4.2xlarge&#34;)
 *                     .weightedCapacity(2)
 *                     .build())
 *             .launchSpecifications(InstanceFleetLaunchSpecificationsArgs.builder()
 *                 .spotSpecifications(InstanceFleetLaunchSpecificationsSpotSpecificationArgs.builder()
 *                     .allocationStrategy(&#34;capacity-optimized&#34;)
 *                     .blockDurationMinutes(0)
 *                     .timeoutAction(&#34;TERMINATE_CLUSTER&#34;)
 *                     .timeoutDurationMinutes(10)
 *                     .build())
 *                 .build())
 *             .targetOnDemandCapacity(1)
 *             .targetSpotCapacity(1)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * EMR Instance Fleet can be imported with the EMR Cluster identifier and Instance Fleet identifier separated by a forward slash (`/`), e.g., console
 * 
 * ```sh
 *  $ pulumi import aws:emr/instanceFleet:InstanceFleet example j-123456ABCDEF/if-15EK4O09RZLNR
 * ```
 * 
 */
@ResourceType(type="aws:emr/instanceFleet:InstanceFleet")
public class InstanceFleet extends com.pulumi.resources.CustomResource {
    /**
     * ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="clusterId", refs={String.class}, tree="[0]")
    private Output<String> clusterId;

    /**
     * @return ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }
    /**
     * Configuration block for instance fleet
     * 
     */
    @Export(name="instanceTypeConfigs", refs={List.class,InstanceFleetInstanceTypeConfig.class}, tree="[0,1]")
    private Output</* @Nullable */ List<InstanceFleetInstanceTypeConfig>> instanceTypeConfigs;

    /**
     * @return Configuration block for instance fleet
     * 
     */
    public Output<Optional<List<InstanceFleetInstanceTypeConfig>>> instanceTypeConfigs() {
        return Codegen.optional(this.instanceTypeConfigs);
    }
    /**
     * Configuration block for launch specification
     * 
     */
    @Export(name="launchSpecifications", refs={InstanceFleetLaunchSpecifications.class}, tree="[0]")
    private Output</* @Nullable */ InstanceFleetLaunchSpecifications> launchSpecifications;

    /**
     * @return Configuration block for launch specification
     * 
     */
    public Output<Optional<InstanceFleetLaunchSpecifications>> launchSpecifications() {
        return Codegen.optional(this.launchSpecifications);
    }
    /**
     * Friendly name given to the instance fleet.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Friendly name given to the instance fleet.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    @Export(name="provisionedOnDemandCapacity", refs={Integer.class}, tree="[0]")
    private Output<Integer> provisionedOnDemandCapacity;

    public Output<Integer> provisionedOnDemandCapacity() {
        return this.provisionedOnDemandCapacity;
    }
    @Export(name="provisionedSpotCapacity", refs={Integer.class}, tree="[0]")
    private Output<Integer> provisionedSpotCapacity;

    public Output<Integer> provisionedSpotCapacity() {
        return this.provisionedSpotCapacity;
    }
    /**
     * The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand instances to provision.
     * 
     */
    @Export(name="targetOnDemandCapacity", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> targetOnDemandCapacity;

    /**
     * @return The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand instances to provision.
     * 
     */
    public Output<Optional<Integer>> targetOnDemandCapacity() {
        return Codegen.optional(this.targetOnDemandCapacity);
    }
    /**
     * The target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision.
     * 
     */
    @Export(name="targetSpotCapacity", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> targetSpotCapacity;

    /**
     * @return The target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision.
     * 
     */
    public Output<Optional<Integer>> targetSpotCapacity() {
        return Codegen.optional(this.targetSpotCapacity);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InstanceFleet(String name) {
        this(name, InstanceFleetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InstanceFleet(String name, InstanceFleetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InstanceFleet(String name, InstanceFleetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:emr/instanceFleet:InstanceFleet", name, args == null ? InstanceFleetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InstanceFleet(String name, Output<String> id, @Nullable InstanceFleetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:emr/instanceFleet:InstanceFleet", name, state, makeResourceOptions(options, id));
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
    public static InstanceFleet get(String name, Output<String> id, @Nullable InstanceFleetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InstanceFleet(name, id, state, options);
    }
}
