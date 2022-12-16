// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.dms;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.dms.ReplicationSubnetGroupArgs;
import com.pulumi.aws.dms.inputs.ReplicationSubnetGroupState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a DMS (Data Migration Service) replication subnet group resource. DMS replication subnet groups can be created, updated, deleted, and imported.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.dms.ReplicationSubnetGroup;
 * import com.pulumi.aws.dms.ReplicationSubnetGroupArgs;
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
 *         var test = new ReplicationSubnetGroup(&#34;test&#34;, ReplicationSubnetGroupArgs.builder()        
 *             .replicationSubnetGroupDescription(&#34;Test replication subnet group&#34;)
 *             .replicationSubnetGroupId(&#34;test-dms-replication-subnet-group-tf&#34;)
 *             .subnetIds(&#34;subnet-12345678&#34;)
 *             .tags(Map.of(&#34;Name&#34;, &#34;test&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Replication subnet groups can be imported using the `replication_subnet_group_id`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:dms/replicationSubnetGroup:ReplicationSubnetGroup test test-dms-replication-subnet-group-tf
 * ```
 * 
 */
@ResourceType(type="aws:dms/replicationSubnetGroup:ReplicationSubnetGroup")
public class ReplicationSubnetGroup extends com.pulumi.resources.CustomResource {
    @Export(name="replicationSubnetGroupArn", refs={String.class}, tree="[0]")
    private Output<String> replicationSubnetGroupArn;

    public Output<String> replicationSubnetGroupArn() {
        return this.replicationSubnetGroupArn;
    }
    /**
     * The description for the subnet group.
     * 
     */
    @Export(name="replicationSubnetGroupDescription", refs={String.class}, tree="[0]")
    private Output<String> replicationSubnetGroupDescription;

    /**
     * @return The description for the subnet group.
     * 
     */
    public Output<String> replicationSubnetGroupDescription() {
        return this.replicationSubnetGroupDescription;
    }
    /**
     * The name for the replication subnet group. This value is stored as a lowercase string.
     * 
     */
    @Export(name="replicationSubnetGroupId", refs={String.class}, tree="[0]")
    private Output<String> replicationSubnetGroupId;

    /**
     * @return The name for the replication subnet group. This value is stored as a lowercase string.
     * 
     */
    public Output<String> replicationSubnetGroupId() {
        return this.replicationSubnetGroupId;
    }
    /**
     * A list of the EC2 subnet IDs for the subnet group.
     * 
     */
    @Export(name="subnetIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> subnetIds;

    /**
     * @return A list of the EC2 subnet IDs for the subnet group.
     * 
     */
    public Output<List<String>> subnetIds() {
        return this.subnetIds;
    }
    /**
     * A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    @Export(name="tagsAll", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }
    /**
     * The ID of the VPC the subnet group is in.
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output<String> vpcId;

    /**
     * @return The ID of the VPC the subnet group is in.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ReplicationSubnetGroup(String name) {
        this(name, ReplicationSubnetGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ReplicationSubnetGroup(String name, ReplicationSubnetGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ReplicationSubnetGroup(String name, ReplicationSubnetGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:dms/replicationSubnetGroup:ReplicationSubnetGroup", name, args == null ? ReplicationSubnetGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ReplicationSubnetGroup(String name, Output<String> id, @Nullable ReplicationSubnetGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:dms/replicationSubnetGroup:ReplicationSubnetGroup", name, state, makeResourceOptions(options, id));
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
    public static ReplicationSubnetGroup get(String name, Output<String> id, @Nullable ReplicationSubnetGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ReplicationSubnetGroup(name, id, state, options);
    }
}
