// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ec2.TagArgs;
import com.pulumi.aws.ec2.inputs.TagState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Manages an individual EC2 resource tag. This resource should only be used in cases where EC2 resources are created outside the provider (e.g. AMIs), being shared via Resource Access Manager (RAM), or implicitly created by other means (e.g. Transit Gateway VPN Attachments).
 * 
 * &gt; **NOTE:** This tagging resource should not be combined with the providers resource for managing the parent resource. For example, using `aws.ec2.Vpc` and `aws.ec2.Tag` to manage tags of the same VPC will cause a perpetual difference where the `aws.ec2.Vpc` resource will try to remove the tag being added by the `aws.ec2.Tag` resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ec2transitgateway.TransitGateway;
 * import com.pulumi.aws.ec2.CustomerGateway;
 * import com.pulumi.aws.ec2.CustomerGatewayArgs;
 * import com.pulumi.aws.ec2.VpnConnection;
 * import com.pulumi.aws.ec2.VpnConnectionArgs;
 * import com.pulumi.aws.ec2.Tag;
 * import com.pulumi.aws.ec2.TagArgs;
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
 *         var exampleTransitGateway = new TransitGateway(&#34;exampleTransitGateway&#34;);
 * 
 *         var exampleCustomerGateway = new CustomerGateway(&#34;exampleCustomerGateway&#34;, CustomerGatewayArgs.builder()        
 *             .bgpAsn(65000)
 *             .ipAddress(&#34;172.0.0.1&#34;)
 *             .type(&#34;ipsec.1&#34;)
 *             .build());
 * 
 *         var exampleVpnConnection = new VpnConnection(&#34;exampleVpnConnection&#34;, VpnConnectionArgs.builder()        
 *             .customerGatewayId(exampleCustomerGateway.id())
 *             .transitGatewayId(exampleTransitGateway.id())
 *             .type(exampleCustomerGateway.type())
 *             .build());
 * 
 *         var exampleTag = new Tag(&#34;exampleTag&#34;, TagArgs.builder()        
 *             .resourceId(exampleVpnConnection.transitGatewayAttachmentId())
 *             .key(&#34;Name&#34;)
 *             .value(&#34;Hello World&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * `aws_ec2_tag` can be imported by using the EC2 resource identifier and key, separated by a comma (`,`), e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:ec2/tag:Tag example tgw-attach-1234567890abcdef,Name
 * ```
 * 
 */
@ResourceType(type="aws:ec2/tag:Tag")
public class Tag extends com.pulumi.resources.CustomResource {
    /**
     * The tag name.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    /**
     * @return The tag name.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * The ID of the EC2 resource to manage the tag for.
     * 
     */
    @Export(name="resourceId", refs={String.class}, tree="[0]")
    private Output<String> resourceId;

    /**
     * @return The ID of the EC2 resource to manage the tag for.
     * 
     */
    public Output<String> resourceId() {
        return this.resourceId;
    }
    /**
     * The value of the tag.
     * 
     */
    @Export(name="value", refs={String.class}, tree="[0]")
    private Output<String> value;

    /**
     * @return The value of the tag.
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Tag(String name) {
        this(name, TagArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Tag(String name, TagArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Tag(String name, TagArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/tag:Tag", name, args == null ? TagArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Tag(String name, Output<String> id, @Nullable TagState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/tag:Tag", name, state, makeResourceOptions(options, id));
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
    public static Tag get(String name, Output<String> id, @Nullable TagState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Tag(name, id, state, options);
    }
}
