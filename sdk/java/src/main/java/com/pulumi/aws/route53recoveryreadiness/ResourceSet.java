// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.route53recoveryreadiness;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.route53recoveryreadiness.ResourceSetArgs;
import com.pulumi.aws.route53recoveryreadiness.inputs.ResourceSetState;
import com.pulumi.aws.route53recoveryreadiness.outputs.ResourceSetResource;
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
 * Provides an AWS Route 53 Recovery Readiness Resource Set.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.route53recoveryreadiness.ResourceSet;
 * import com.pulumi.aws.route53recoveryreadiness.ResourceSetArgs;
 * import com.pulumi.aws.route53recoveryreadiness.inputs.ResourceSetResourceArgs;
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
 *         var example = new ResourceSet(&#34;example&#34;, ResourceSetArgs.builder()        
 *             .resourceSetName(my_cw_alarm_set)
 *             .resourceSetType(&#34;AWS::CloudWatch::Alarm&#34;)
 *             .resources(ResourceSetResourceArgs.builder()
 *                 .resourceArn(aws_cloudwatch_metric_alarm.example().arn())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Route53 Recovery Readiness resource set name can be imported via the resource set name, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:route53recoveryreadiness/resourceSet:ResourceSet my-cw-alarm-set
 * ```
 * 
 */
@ResourceType(type="aws:route53recoveryreadiness/resourceSet:ResourceSet")
public class ResourceSet extends com.pulumi.resources.CustomResource {
    /**
     * NLB resource ARN.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return NLB resource ARN.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Unique name describing the resource set.
     * 
     */
    @Export(name="resourceSetName", refs={String.class}, tree="[0]")
    private Output<String> resourceSetName;

    /**
     * @return Unique name describing the resource set.
     * 
     */
    public Output<String> resourceSetName() {
        return this.resourceSetName;
    }
    /**
     * Type of the resources in the resource set.
     * 
     */
    @Export(name="resourceSetType", refs={String.class}, tree="[0]")
    private Output<String> resourceSetType;

    /**
     * @return Type of the resources in the resource set.
     * 
     */
    public Output<String> resourceSetType() {
        return this.resourceSetType;
    }
    /**
     * List of resources to add to this resource set. See below.
     * 
     */
    @Export(name="resources", refs={List.class,ResourceSetResource.class}, tree="[0,1]")
    private Output<List<ResourceSetResource>> resources;

    /**
     * @return List of resources to add to this resource set. See below.
     * 
     */
    public Output<List<ResourceSetResource>> resources() {
        return this.resources;
    }
    /**
     * Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
     * 
     */
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
    public ResourceSet(String name) {
        this(name, ResourceSetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ResourceSet(String name, ResourceSetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ResourceSet(String name, ResourceSetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:route53recoveryreadiness/resourceSet:ResourceSet", name, args == null ? ResourceSetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ResourceSet(String name, Output<String> id, @Nullable ResourceSetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:route53recoveryreadiness/resourceSet:ResourceSet", name, state, makeResourceOptions(options, id));
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
    public static ResourceSet get(String name, Output<String> id, @Nullable ResourceSetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ResourceSet(name, id, state, options);
    }
}
