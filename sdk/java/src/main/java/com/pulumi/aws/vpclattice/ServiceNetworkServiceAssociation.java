// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.vpclattice;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.vpclattice.ServiceNetworkServiceAssociationArgs;
import com.pulumi.aws.vpclattice.inputs.ServiceNetworkServiceAssociationState;
import com.pulumi.aws.vpclattice.outputs.ServiceNetworkServiceAssociationDnsEntry;
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
 * Resource for managing an AWS VPC Lattice Service Network Service Association.
 * 
 * ## Example Usage
 * ### Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.vpclattice.ServiceNetworkServiceAssociation;
 * import com.pulumi.aws.vpclattice.ServiceNetworkServiceAssociationArgs;
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
 *         var example = new ServiceNetworkServiceAssociation(&#34;example&#34;, ServiceNetworkServiceAssociationArgs.builder()        
 *             .serviceIdentifier(aws_vpclattice_service.example().id())
 *             .serviceNetworkIdentifier(aws_vpclattice_service_network.example().id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import VPC Lattice Service Network Service Association using the `id`. For example:
 * 
 * ```sh
 *  $ pulumi import aws:vpclattice/serviceNetworkServiceAssociation:ServiceNetworkServiceAssociation example snsa-05e2474658a88f6ba
 * ```
 * 
 */
@ResourceType(type="aws:vpclattice/serviceNetworkServiceAssociation:ServiceNetworkServiceAssociation")
public class ServiceNetworkServiceAssociation extends com.pulumi.resources.CustomResource {
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
     * The custom domain name of the service.
     * 
     */
    @Export(name="customDomainName", refs={String.class}, tree="[0]")
    private Output<String> customDomainName;

    /**
     * @return The custom domain name of the service.
     * 
     */
    public Output<String> customDomainName() {
        return this.customDomainName;
    }
    /**
     * The DNS name of the service.
     * 
     */
    @Export(name="dnsEntries", refs={List.class,ServiceNetworkServiceAssociationDnsEntry.class}, tree="[0,1]")
    private Output<List<ServiceNetworkServiceAssociationDnsEntry>> dnsEntries;

    /**
     * @return The DNS name of the service.
     * 
     */
    public Output<List<ServiceNetworkServiceAssociationDnsEntry>> dnsEntries() {
        return this.dnsEntries;
    }
    /**
     * The ID or Amazon Resource Identifier (ARN) of the service.
     * 
     */
    @Export(name="serviceIdentifier", refs={String.class}, tree="[0]")
    private Output<String> serviceIdentifier;

    /**
     * @return The ID or Amazon Resource Identifier (ARN) of the service.
     * 
     */
    public Output<String> serviceIdentifier() {
        return this.serviceIdentifier;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceNetworkServiceAssociation(String name) {
        this(name, ServiceNetworkServiceAssociationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceNetworkServiceAssociation(String name, ServiceNetworkServiceAssociationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceNetworkServiceAssociation(String name, ServiceNetworkServiceAssociationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:vpclattice/serviceNetworkServiceAssociation:ServiceNetworkServiceAssociation", name, args == null ? ServiceNetworkServiceAssociationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ServiceNetworkServiceAssociation(String name, Output<String> id, @Nullable ServiceNetworkServiceAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:vpclattice/serviceNetworkServiceAssociation:ServiceNetworkServiceAssociation", name, state, makeResourceOptions(options, id));
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
    public static ServiceNetworkServiceAssociation get(String name, Output<String> id, @Nullable ServiceNetworkServiceAssociationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceNetworkServiceAssociation(name, id, state, options);
    }
}
