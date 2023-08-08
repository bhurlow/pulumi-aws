// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.networkmanager;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.networkmanager.SiteToSiteVpnAttachmentArgs;
import com.pulumi.aws.networkmanager.inputs.SiteToSiteVpnAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Resource for managing an AWS NetworkManager SiteToSiteAttachment.
 * 
 * ## Example Usage
 * ### Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.networkmanager.SiteToSiteVpnAttachment;
 * import com.pulumi.aws.networkmanager.SiteToSiteVpnAttachmentArgs;
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
 *         var example = new SiteToSiteVpnAttachment(&#34;example&#34;, SiteToSiteVpnAttachmentArgs.builder()        
 *             .coreNetworkId(awscc_networkmanager_core_network.example().id())
 *             .vpnConnectionArn(aws_vpn_connection.example().arn())
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
 *  to = aws_networkmanager_site_to_site_vpn_attachment.example
 * 
 *  id = &#34;attachment-0f8fa60d2238d1bd8&#34; } Using `pulumi import`, import `aws_networkmanager_site_to_site_vpn_attachment` using the attachment ID. For exampleconsole % pulumi import aws_networkmanager_site_to_site_vpn_attachment.example attachment-0f8fa60d2238d1bd8
 * 
 */
@ResourceType(type="aws:networkmanager/siteToSiteVpnAttachment:SiteToSiteVpnAttachment")
public class SiteToSiteVpnAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The ARN of the attachment.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the attachment.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The policy rule number associated with the attachment.
     * 
     */
    @Export(name="attachmentPolicyRuleNumber", refs={Integer.class}, tree="[0]")
    private Output<Integer> attachmentPolicyRuleNumber;

    /**
     * @return The policy rule number associated with the attachment.
     * 
     */
    public Output<Integer> attachmentPolicyRuleNumber() {
        return this.attachmentPolicyRuleNumber;
    }
    /**
     * The type of attachment.
     * 
     */
    @Export(name="attachmentType", refs={String.class}, tree="[0]")
    private Output<String> attachmentType;

    /**
     * @return The type of attachment.
     * 
     */
    public Output<String> attachmentType() {
        return this.attachmentType;
    }
    /**
     * The ARN of a core network.
     * 
     */
    @Export(name="coreNetworkArn", refs={String.class}, tree="[0]")
    private Output<String> coreNetworkArn;

    /**
     * @return The ARN of a core network.
     * 
     */
    public Output<String> coreNetworkArn() {
        return this.coreNetworkArn;
    }
    /**
     * The ID of a core network for the VPN attachment.
     * 
     */
    @Export(name="coreNetworkId", refs={String.class}, tree="[0]")
    private Output<String> coreNetworkId;

    /**
     * @return The ID of a core network for the VPN attachment.
     * 
     */
    public Output<String> coreNetworkId() {
        return this.coreNetworkId;
    }
    /**
     * The Region where the edge is located.
     * 
     */
    @Export(name="edgeLocation", refs={String.class}, tree="[0]")
    private Output<String> edgeLocation;

    /**
     * @return The Region where the edge is located.
     * 
     */
    public Output<String> edgeLocation() {
        return this.edgeLocation;
    }
    /**
     * The ID of the attachment account owner.
     * 
     */
    @Export(name="ownerAccountId", refs={String.class}, tree="[0]")
    private Output<String> ownerAccountId;

    /**
     * @return The ID of the attachment account owner.
     * 
     */
    public Output<String> ownerAccountId() {
        return this.ownerAccountId;
    }
    /**
     * The attachment resource ARN.
     * 
     */
    @Export(name="resourceArn", refs={String.class}, tree="[0]")
    private Output<String> resourceArn;

    /**
     * @return The attachment resource ARN.
     * 
     */
    public Output<String> resourceArn() {
        return this.resourceArn;
    }
    /**
     * The name of the segment attachment.
     * 
     */
    @Export(name="segmentName", refs={String.class}, tree="[0]")
    private Output<String> segmentName;

    /**
     * @return The name of the segment attachment.
     * 
     */
    public Output<String> segmentName() {
        return this.segmentName;
    }
    /**
     * The state of the attachment.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return The state of the attachment.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Key-value tags for the attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Key-value tags for the attachment. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * The ARN of the site-to-site VPN connection.
     * 
     * The following arguments are optional:
     * 
     */
    @Export(name="vpnConnectionArn", refs={String.class}, tree="[0]")
    private Output<String> vpnConnectionArn;

    /**
     * @return The ARN of the site-to-site VPN connection.
     * 
     * The following arguments are optional:
     * 
     */
    public Output<String> vpnConnectionArn() {
        return this.vpnConnectionArn;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SiteToSiteVpnAttachment(String name) {
        this(name, SiteToSiteVpnAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SiteToSiteVpnAttachment(String name, SiteToSiteVpnAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SiteToSiteVpnAttachment(String name, SiteToSiteVpnAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:networkmanager/siteToSiteVpnAttachment:SiteToSiteVpnAttachment", name, args == null ? SiteToSiteVpnAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SiteToSiteVpnAttachment(String name, Output<String> id, @Nullable SiteToSiteVpnAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:networkmanager/siteToSiteVpnAttachment:SiteToSiteVpnAttachment", name, state, makeResourceOptions(options, id));
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
    public static SiteToSiteVpnAttachment get(String name, Output<String> id, @Nullable SiteToSiteVpnAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SiteToSiteVpnAttachment(name, id, state, options);
    }
}
