// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sagemaker;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.sagemaker.ImageVersionArgs;
import com.pulumi.aws.sagemaker.inputs.ImageVersionState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a SageMaker Image Version resource.
 * 
 * ## Example Usage
 * ### Basic usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.sagemaker.ImageVersion;
 * import com.pulumi.aws.sagemaker.ImageVersionArgs;
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
 *         var test = new ImageVersion(&#34;test&#34;, ImageVersionArgs.builder()        
 *             .imageName(aws_sagemaker_image.test().id())
 *             .baseImage(&#34;012345678912.dkr.ecr.us-west-2.amazonaws.com/image:latest&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * terraform import {
 * 
 *  to = aws_sagemaker_image_version.test_image
 * 
 *  id = &#34;my-code-repo&#34; } Using `pulumi import`, import SageMaker Image Versions using the `name`. For exampleconsole % pulumi import aws_sagemaker_image_version.test_image my-code-repo
 * 
 */
@ResourceType(type="aws:sagemaker/imageVersion:ImageVersion")
public class ImageVersion extends com.pulumi.resources.CustomResource {
    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this Image Version.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) assigned by AWS to this Image Version.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The registry path of the container image on which this image version is based.
     * 
     */
    @Export(name="baseImage", refs={String.class}, tree="[0]")
    private Output<String> baseImage;

    /**
     * @return The registry path of the container image on which this image version is based.
     * 
     */
    public Output<String> baseImage() {
        return this.baseImage;
    }
    /**
     * The registry path of the container image that contains this image version.
     * 
     */
    @Export(name="containerImage", refs={String.class}, tree="[0]")
    private Output<String> containerImage;

    /**
     * @return The registry path of the container image that contains this image version.
     * 
     */
    public Output<String> containerImage() {
        return this.containerImage;
    }
    /**
     * The Amazon Resource Name (ARN) of the image the version is based on.
     * 
     */
    @Export(name="imageArn", refs={String.class}, tree="[0]")
    private Output<String> imageArn;

    /**
     * @return The Amazon Resource Name (ARN) of the image the version is based on.
     * 
     */
    public Output<String> imageArn() {
        return this.imageArn;
    }
    /**
     * The name of the image. Must be unique to your account.
     * 
     */
    @Export(name="imageName", refs={String.class}, tree="[0]")
    private Output<String> imageName;

    /**
     * @return The name of the image. Must be unique to your account.
     * 
     */
    public Output<String> imageName() {
        return this.imageName;
    }
    @Export(name="version", refs={Integer.class}, tree="[0]")
    private Output<Integer> version;

    public Output<Integer> version() {
        return this.version;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ImageVersion(String name) {
        this(name, ImageVersionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ImageVersion(String name, ImageVersionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ImageVersion(String name, ImageVersionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:sagemaker/imageVersion:ImageVersion", name, args == null ? ImageVersionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ImageVersion(String name, Output<String> id, @Nullable ImageVersionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:sagemaker/imageVersion:ImageVersion", name, state, makeResourceOptions(options, id));
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
    public static ImageVersion get(String name, Output<String> id, @Nullable ImageVersionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ImageVersion(name, id, state, options);
    }
}
