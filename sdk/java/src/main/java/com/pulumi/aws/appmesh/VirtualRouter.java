// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.appmesh;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.appmesh.VirtualRouterArgs;
import com.pulumi.aws.appmesh.inputs.VirtualRouterState;
import com.pulumi.aws.appmesh.outputs.VirtualRouterSpec;
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
 * Provides an AWS App Mesh virtual router resource.
 * 
 * ## Breaking Changes
 * 
 * Because of backward incompatible API changes (read [here](https://github.com/awslabs/aws-app-mesh-examples/issues/92) and [here](https://github.com/awslabs/aws-app-mesh-examples/issues/94)), `aws.appmesh.VirtualRouter` resource definitions created with provider versions earlier than v2.3.0 will need to be modified:
 * 
 * * Remove service `service_names` from the `spec` argument. AWS has created a `aws.appmesh.VirtualService` resource for each service name. Import these resource using `pulumi import`.
 * 
 * * Add a `listener` configuration block to the `spec` argument.
 * 
 * The state associated with existing resources will automatically be migrated.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.appmesh.VirtualRouter;
 * import com.pulumi.aws.appmesh.VirtualRouterArgs;
 * import com.pulumi.aws.appmesh.inputs.VirtualRouterSpecArgs;
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
 *         var serviceb = new VirtualRouter(&#34;serviceb&#34;, VirtualRouterArgs.builder()        
 *             .meshName(aws_appmesh_mesh.simple().id())
 *             .spec(VirtualRouterSpecArgs.builder()
 *                 .listeners(VirtualRouterSpecListenerArgs.builder()
 *                     .portMapping(VirtualRouterSpecListenerPortMappingArgs.builder()
 *                         .port(8080)
 *                         .protocol(&#34;http&#34;)
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Using `pulumi import`, import App Mesh virtual routers using `mesh_name` together with the virtual router&#39;s `name`. For example:
 * 
 * ```sh
 *  $ pulumi import aws:appmesh/virtualRouter:VirtualRouter serviceb simpleapp/serviceB
 * ```
 * 
 */
@ResourceType(type="aws:appmesh/virtualRouter:VirtualRouter")
public class VirtualRouter extends com.pulumi.resources.CustomResource {
    /**
     * ARN of the virtual router.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return ARN of the virtual router.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Creation date of the virtual router.
     * 
     */
    @Export(name="createdDate", refs={String.class}, tree="[0]")
    private Output<String> createdDate;

    /**
     * @return Creation date of the virtual router.
     * 
     */
    public Output<String> createdDate() {
        return this.createdDate;
    }
    /**
     * Last update date of the virtual router.
     * 
     */
    @Export(name="lastUpdatedDate", refs={String.class}, tree="[0]")
    private Output<String> lastUpdatedDate;

    /**
     * @return Last update date of the virtual router.
     * 
     */
    public Output<String> lastUpdatedDate() {
        return this.lastUpdatedDate;
    }
    /**
     * Name of the service mesh in which to create the virtual router. Must be between 1 and 255 characters in length.
     * 
     */
    @Export(name="meshName", refs={String.class}, tree="[0]")
    private Output<String> meshName;

    /**
     * @return Name of the service mesh in which to create the virtual router. Must be between 1 and 255 characters in length.
     * 
     */
    public Output<String> meshName() {
        return this.meshName;
    }
    /**
     * AWS account ID of the service mesh&#39;s owner. Defaults to the account ID the AWS provider is currently connected to.
     * 
     */
    @Export(name="meshOwner", refs={String.class}, tree="[0]")
    private Output<String> meshOwner;

    /**
     * @return AWS account ID of the service mesh&#39;s owner. Defaults to the account ID the AWS provider is currently connected to.
     * 
     */
    public Output<String> meshOwner() {
        return this.meshOwner;
    }
    /**
     * Name to use for the virtual router. Must be between 1 and 255 characters in length.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name to use for the virtual router. Must be between 1 and 255 characters in length.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Resource owner&#39;s AWS account ID.
     * 
     */
    @Export(name="resourceOwner", refs={String.class}, tree="[0]")
    private Output<String> resourceOwner;

    /**
     * @return Resource owner&#39;s AWS account ID.
     * 
     */
    public Output<String> resourceOwner() {
        return this.resourceOwner;
    }
    /**
     * Virtual router specification to apply.
     * 
     */
    @Export(name="spec", refs={VirtualRouterSpec.class}, tree="[0]")
    private Output<VirtualRouterSpec> spec;

    /**
     * @return Virtual router specification to apply.
     * 
     */
    public Output<VirtualRouterSpec> spec() {
        return this.spec;
    }
    /**
     * Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
    public VirtualRouter(String name) {
        this(name, VirtualRouterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VirtualRouter(String name, VirtualRouterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VirtualRouter(String name, VirtualRouterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:appmesh/virtualRouter:VirtualRouter", name, args == null ? VirtualRouterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VirtualRouter(String name, Output<String> id, @Nullable VirtualRouterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:appmesh/virtualRouter:VirtualRouter", name, state, makeResourceOptions(options, id));
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
    public static VirtualRouter get(String name, Output<String> id, @Nullable VirtualRouterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VirtualRouter(name, id, state, options);
    }
}
