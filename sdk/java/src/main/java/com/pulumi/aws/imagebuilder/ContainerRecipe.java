// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.imagebuilder;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.imagebuilder.ContainerRecipeArgs;
import com.pulumi.aws.imagebuilder.inputs.ContainerRecipeState;
import com.pulumi.aws.imagebuilder.outputs.ContainerRecipeComponent;
import com.pulumi.aws.imagebuilder.outputs.ContainerRecipeInstanceConfiguration;
import com.pulumi.aws.imagebuilder.outputs.ContainerRecipeTargetRepository;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an Image Builder Container Recipe.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.imagebuilder.ContainerRecipe;
 * import com.pulumi.aws.imagebuilder.ContainerRecipeArgs;
 * import com.pulumi.aws.imagebuilder.inputs.ContainerRecipeTargetRepositoryArgs;
 * import com.pulumi.aws.imagebuilder.inputs.ContainerRecipeComponentArgs;
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
 *         var example = new ContainerRecipe(&#34;example&#34;, ContainerRecipeArgs.builder()        
 *             .version(&#34;1.0.0&#34;)
 *             .containerType(&#34;DOCKER&#34;)
 *             .parentImage(&#34;arn:aws:imagebuilder:eu-central-1:aws:image/amazon-linux-x86-latest/x.x.x&#34;)
 *             .targetRepository(ContainerRecipeTargetRepositoryArgs.builder()
 *                 .repositoryName(aws_ecr_repository.example().name())
 *                 .service(&#34;ECR&#34;)
 *                 .build())
 *             .components(ContainerRecipeComponentArgs.builder()
 *                 .componentArn(aws_imagebuilder_component.example().arn())
 *                 .parameters(                
 *                     ContainerRecipeComponentParameterArgs.builder()
 *                         .name(&#34;Parameter1&#34;)
 *                         .value(&#34;Value1&#34;)
 *                         .build(),
 *                     ContainerRecipeComponentParameterArgs.builder()
 *                         .name(&#34;Parameter2&#34;)
 *                         .value(&#34;Value2&#34;)
 *                         .build())
 *                 .build())
 *             .dockerfileTemplateData(&#34;&#34;&#34;
 * FROM {{{ imagebuilder:parentImage }}}
 * {{{ imagebuilder:environments }}}
 * {{{ imagebuilder:components }}}
 *             &#34;&#34;&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * `aws_imagebuilder_container_recipe` resources can be imported by using the Amazon Resource Name (ARN), e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:imagebuilder/containerRecipe:ContainerRecipe example arn:aws:imagebuilder:us-east-1:123456789012:container-recipe/example/1.0.0
 * ```
 * 
 */
@ResourceType(type="aws:imagebuilder/containerRecipe:ContainerRecipe")
public class ContainerRecipe extends com.pulumi.resources.CustomResource {
    /**
     * (Required) Amazon Resource Name (ARN) of the container recipe.
     * 
     */
    @Export(name="arn", type=String.class, parameters={})
    private Output<String> arn;

    /**
     * @return (Required) Amazon Resource Name (ARN) of the container recipe.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * Ordered configuration block(s) with components for the container recipe. Detailed below.
     * 
     */
    @Export(name="components", type=List.class, parameters={ContainerRecipeComponent.class})
    private Output<List<ContainerRecipeComponent>> components;

    /**
     * @return Ordered configuration block(s) with components for the container recipe. Detailed below.
     * 
     */
    public Output<List<ContainerRecipeComponent>> components() {
        return this.components;
    }
    /**
     * The type of the container to create. Valid values: `DOCKER`.
     * 
     */
    @Export(name="containerType", type=String.class, parameters={})
    private Output<String> containerType;

    /**
     * @return The type of the container to create. Valid values: `DOCKER`.
     * 
     */
    public Output<String> containerType() {
        return this.containerType;
    }
    /**
     * Date the container recipe was created.
     * 
     */
    @Export(name="dateCreated", type=String.class, parameters={})
    private Output<String> dateCreated;

    /**
     * @return Date the container recipe was created.
     * 
     */
    public Output<String> dateCreated() {
        return this.dateCreated;
    }
    /**
     * The description of the container recipe.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the container recipe.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The Dockerfile template used to build the image as an inline data blob.
     * 
     */
    @Export(name="dockerfileTemplateData", type=String.class, parameters={})
    private Output<String> dockerfileTemplateData;

    /**
     * @return The Dockerfile template used to build the image as an inline data blob.
     * 
     */
    public Output<String> dockerfileTemplateData() {
        return this.dockerfileTemplateData;
    }
    /**
     * The Amazon S3 URI for the Dockerfile that will be used to build the container image.
     * 
     */
    @Export(name="dockerfileTemplateUri", type=String.class, parameters={})
    private Output</* @Nullable */ String> dockerfileTemplateUri;

    /**
     * @return The Amazon S3 URI for the Dockerfile that will be used to build the container image.
     * 
     */
    public Output<Optional<String>> dockerfileTemplateUri() {
        return Codegen.optional(this.dockerfileTemplateUri);
    }
    /**
     * Whether to encrypt the volume. Defaults to unset, which is the value inherited from the parent image.
     * 
     */
    @Export(name="encrypted", type=Boolean.class, parameters={})
    private Output<Boolean> encrypted;

    /**
     * @return Whether to encrypt the volume. Defaults to unset, which is the value inherited from the parent image.
     * 
     */
    public Output<Boolean> encrypted() {
        return this.encrypted;
    }
    /**
     * Configuration block used to configure an instance for building and testing container images. Detailed below.
     * 
     */
    @Export(name="instanceConfiguration", type=ContainerRecipeInstanceConfiguration.class, parameters={})
    private Output</* @Nullable */ ContainerRecipeInstanceConfiguration> instanceConfiguration;

    /**
     * @return Configuration block used to configure an instance for building and testing container images. Detailed below.
     * 
     */
    public Output<Optional<ContainerRecipeInstanceConfiguration>> instanceConfiguration() {
        return Codegen.optional(this.instanceConfiguration);
    }
    /**
     * Amazon Resource Name (ARN) of the Key Management Service (KMS) Key for encryption.
     * 
     */
    @Export(name="kmsKeyId", type=String.class, parameters={})
    private Output</* @Nullable */ String> kmsKeyId;

    /**
     * @return Amazon Resource Name (ARN) of the Key Management Service (KMS) Key for encryption.
     * 
     */
    public Output<Optional<String>> kmsKeyId() {
        return Codegen.optional(this.kmsKeyId);
    }
    /**
     * The name of the component parameter.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the component parameter.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Owner of the container recipe.
     * 
     */
    @Export(name="owner", type=String.class, parameters={})
    private Output<String> owner;

    /**
     * @return Owner of the container recipe.
     * 
     */
    public Output<String> owner() {
        return this.owner;
    }
    /**
     * The base image for the container recipe.
     * 
     */
    @Export(name="parentImage", type=String.class, parameters={})
    private Output<String> parentImage;

    /**
     * @return The base image for the container recipe.
     * 
     */
    public Output<String> parentImage() {
        return this.parentImage;
    }
    /**
     * Platform of the container recipe.
     * 
     */
    @Export(name="platform", type=String.class, parameters={})
    private Output<String> platform;

    /**
     * @return Platform of the container recipe.
     * 
     */
    public Output<String> platform() {
        return this.platform;
    }
    /**
     * Key-value map of resource tags for the container recipe. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value map of resource tags for the container recipe. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
     * 
     */
    @Export(name="tagsAll", type=Map.class, parameters={String.class, String.class})
    private Output<Map<String,String>> tagsAll;

    /**
     * @return A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
     * 
     */
    public Output<Map<String,String>> tagsAll() {
        return this.tagsAll;
    }
    /**
     * The destination repository for the container image. Detailed below.
     * 
     */
    @Export(name="targetRepository", type=ContainerRecipeTargetRepository.class, parameters={})
    private Output<ContainerRecipeTargetRepository> targetRepository;

    /**
     * @return The destination repository for the container image. Detailed below.
     * 
     */
    public Output<ContainerRecipeTargetRepository> targetRepository() {
        return this.targetRepository;
    }
    /**
     * Version of the container recipe.
     * 
     */
    @Export(name="version", type=String.class, parameters={})
    private Output<String> version;

    /**
     * @return Version of the container recipe.
     * 
     */
    public Output<String> version() {
        return this.version;
    }
    /**
     * The working directory to be used during build and test workflows.
     * 
     */
    @Export(name="workingDirectory", type=String.class, parameters={})
    private Output</* @Nullable */ String> workingDirectory;

    /**
     * @return The working directory to be used during build and test workflows.
     * 
     */
    public Output<Optional<String>> workingDirectory() {
        return Codegen.optional(this.workingDirectory);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ContainerRecipe(String name) {
        this(name, ContainerRecipeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ContainerRecipe(String name, ContainerRecipeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ContainerRecipe(String name, ContainerRecipeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:imagebuilder/containerRecipe:ContainerRecipe", name, args == null ? ContainerRecipeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ContainerRecipe(String name, Output<String> id, @Nullable ContainerRecipeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:imagebuilder/containerRecipe:ContainerRecipe", name, state, makeResourceOptions(options, id));
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
    public static ContainerRecipe get(String name, Output<String> id, @Nullable ContainerRecipeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ContainerRecipe(name, id, state, options);
    }
}