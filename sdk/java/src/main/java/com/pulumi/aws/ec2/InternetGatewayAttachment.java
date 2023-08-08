// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ec2.InternetGatewayAttachmentArgs;
import com.pulumi.aws.ec2.inputs.InternetGatewayAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a resource to create a VPC Internet Gateway Attachment.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ec2.Vpc;
 * import com.pulumi.aws.ec2.VpcArgs;
 * import com.pulumi.aws.ec2.InternetGateway;
 * import com.pulumi.aws.ec2.InternetGatewayAttachment;
 * import com.pulumi.aws.ec2.InternetGatewayAttachmentArgs;
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
 *         var exampleVpc = new Vpc(&#34;exampleVpc&#34;, VpcArgs.builder()        
 *             .cidrBlock(&#34;10.1.0.0/16&#34;)
 *             .build());
 * 
 *         var exampleInternetGateway = new InternetGateway(&#34;exampleInternetGateway&#34;);
 * 
 *         var exampleInternetGatewayAttachment = new InternetGatewayAttachment(&#34;exampleInternetGatewayAttachment&#34;, InternetGatewayAttachmentArgs.builder()        
 *             .internetGatewayId(exampleInternetGateway.id())
 *             .vpcId(exampleVpc.id())
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
 *  to = aws_internet_gateway_attachment.example
 * 
 *  id = &#34;igw-c0a643a9:vpc-123456&#34; } Using `pulumi import`, import Internet Gateway Attachments using the `id`. For exampleconsole % pulumi import aws_internet_gateway_attachment.example igw-c0a643a9:vpc-123456
 * 
 */
@ResourceType(type="aws:ec2/internetGatewayAttachment:InternetGatewayAttachment")
public class InternetGatewayAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the internet gateway.
     * 
     */
    @Export(name="internetGatewayId", refs={String.class}, tree="[0]")
    private Output<String> internetGatewayId;

    /**
     * @return The ID of the internet gateway.
     * 
     */
    public Output<String> internetGatewayId() {
        return this.internetGatewayId;
    }
    /**
     * The ID of the VPC.
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output<String> vpcId;

    /**
     * @return The ID of the VPC.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InternetGatewayAttachment(String name) {
        this(name, InternetGatewayAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InternetGatewayAttachment(String name, InternetGatewayAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InternetGatewayAttachment(String name, InternetGatewayAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/internetGatewayAttachment:InternetGatewayAttachment", name, args == null ? InternetGatewayAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InternetGatewayAttachment(String name, Output<String> id, @Nullable InternetGatewayAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/internetGatewayAttachment:InternetGatewayAttachment", name, state, makeResourceOptions(options, id));
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
    public static InternetGatewayAttachment get(String name, Output<String> id, @Nullable InternetGatewayAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InternetGatewayAttachment(name, id, state, options);
    }
}
