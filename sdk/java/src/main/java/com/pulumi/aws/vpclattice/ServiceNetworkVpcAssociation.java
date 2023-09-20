// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.vpclattice;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.vpclattice.ServiceNetworkVpcAssociationArgs;
import com.pulumi.aws.vpclattice.inputs.ServiceNetworkVpcAssociationState;
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
 * Resource for managing an AWS VPC Lattice Service Network VPC Association.
 * 
 * ## Example Usage
 * ### Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.vpclattice.ServiceNetworkVpcAssociation;
 * import com.pulumi.aws.vpclattice.ServiceNetworkVpcAssociationArgs;
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
 *         var example = new ServiceNetworkVpcAssociation(&#34;example&#34;, ServiceNetworkVpcAssociationArgs.builder()        
 *             .vpcIdentifier(aws_vpc.example().id())
 *             .serviceNetworkIdentifier(aws_vpclattice_service_network.example().id())
 *             .securityGroupIds(aws_security_group.example().id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import VPC Lattice Service Network VPC Association using the `id`. For example:
 * 
 * ```sh
 *  $ pulumi import aws:vpclattice/serviceNetworkVpcAssociation:ServiceNetworkVpcAssociation example snsa-05e2474658a88f6ba
 * ```
 * 
 */
@ResourceType(type="aws:vpclattice/serviceNetworkVpcAssociation:ServiceNetworkVpcAssociation")
public class ServiceNetworkVpcAssociation extends com.pulumi.resources.CustomResource {
    /**
     * The ARN of the Association.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the Association.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The account that created the association.
     * 
     */
    @Export(name="createdBy", refs={String.class}, tree="[0]")
    private Output<String> createdBy;

    /**
     * @return The account that created the association.
     * 
     */
    public Output<String> createdBy() {
        return this.createdBy;
    }
    /**
     * The IDs of the security groups.
     * 
     */
    @Export(name="securityGroupIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> securityGroupIds;

    /**
     * @return The IDs of the security groups.
     * 
     */
    public Output<Optional<List<String>>> securityGroupIds() {
        return Codegen.optional(this.securityGroupIds);
    }
    /**
     * The ID or Amazon Resource Identifier (ARN) of the service network. You must use the ARN if the resources specified in the operation are in different accounts.
     * The following arguments are optional:
     * 
     */
    @Export(name="serviceNetworkIdentifier", refs={String.class}, tree="[0]")
    private Output<String> serviceNetworkIdentifier;

    /**
     * @return The ID or Amazon Resource Identifier (ARN) of the service network. You must use the ARN if the resources specified in the operation are in different accounts.
     * The following arguments are optional:
     * 
     */
    public Output<String> serviceNetworkIdentifier() {
        return this.serviceNetworkIdentifier;
    }
    /**
     * The operations status. Valid Values are CREATE_IN_PROGRESS | ACTIVE | DELETE_IN_PROGRESS | CREATE_FAILED | DELETE_FAILED
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The operations status. Valid Values are CREATE_IN_PROGRESS | ACTIVE | DELETE_IN_PROGRESS | CREATE_FAILED | DELETE_FAILED
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     * @deprecated
     * Please use `tags` instead.
     * 
     */
    @Deprecated /* Please use `tags` instead. */
    @Export(name="tagsAll", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> tagsAll;

    /**
     * @return Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }
    /**
     * The ID of the VPC.
     * 
     */
    @Export(name="vpcIdentifier", refs={String.class}, tree="[0]")
    private Output<String> vpcIdentifier;

    /**
     * @return The ID of the VPC.
     * 
     */
    public Output<String> vpcIdentifier() {
        return this.vpcIdentifier;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceNetworkVpcAssociation(String name) {
        this(name, ServiceNetworkVpcAssociationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceNetworkVpcAssociation(String name, ServiceNetworkVpcAssociationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceNetworkVpcAssociation(String name, ServiceNetworkVpcAssociationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:vpclattice/serviceNetworkVpcAssociation:ServiceNetworkVpcAssociation", name, args == null ? ServiceNetworkVpcAssociationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServiceNetworkVpcAssociation(String name, Output<String> id, @Nullable ServiceNetworkVpcAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:vpclattice/serviceNetworkVpcAssociation:ServiceNetworkVpcAssociation", name, state, makeResourceOptions(options, id));
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
    public static ServiceNetworkVpcAssociation get(String name, Output<String> id, @Nullable ServiceNetworkVpcAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceNetworkVpcAssociation(name, id, state, options);
    }
}
