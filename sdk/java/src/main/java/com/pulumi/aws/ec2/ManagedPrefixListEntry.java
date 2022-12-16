// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.ec2;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.ec2.ManagedPrefixListEntryArgs;
import com.pulumi.aws.ec2.inputs.ManagedPrefixListEntryState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a managed prefix list entry resource.
 * 
 * &gt; **NOTE on Managed Prefix Lists and Managed Prefix List Entries:** The provider
 * currently provides both a standalone Managed Prefix List Entry resource (a single entry),
 * and a Managed Prefix List resource with entries defined
 * in-line. At this time you cannot use a Managed Prefix List with in-line rules in
 * conjunction with any Managed Prefix List Entry resources. Doing so will cause a conflict
 * of entries and will overwrite entries.
 * 
 * &gt; **NOTE on Managed Prefix Lists with many entries:**  To improved execution times on larger
 * updates, if you plan to create a prefix list with more than 100 entries, it is **recommended**
 * that you use the inline `entry` block as part of the Managed Prefix List resource
 * resource instead.
 * 
 * ## Example Usage
 * 
 * Basic usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.ec2.ManagedPrefixList;
 * import com.pulumi.aws.ec2.ManagedPrefixListArgs;
 * import com.pulumi.aws.ec2.ManagedPrefixListEntry;
 * import com.pulumi.aws.ec2.ManagedPrefixListEntryArgs;
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
 *         var example = new ManagedPrefixList(&#34;example&#34;, ManagedPrefixListArgs.builder()        
 *             .addressFamily(&#34;IPv4&#34;)
 *             .maxEntries(5)
 *             .tags(Map.of(&#34;Env&#34;, &#34;live&#34;))
 *             .build());
 * 
 *         var entry1 = new ManagedPrefixListEntry(&#34;entry1&#34;, ManagedPrefixListEntryArgs.builder()        
 *             .cidr(aws_vpc.example().cidr_block())
 *             .description(&#34;Primary&#34;)
 *             .prefixListId(example.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Prefix List Entries can be imported using the `prefix_list_id` and `cidr` separated by a `,`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:ec2/managedPrefixListEntry:ManagedPrefixListEntry default pl-0570a1d2d725c16be,10.0.3.0/24
 * ```
 * 
 */
@ResourceType(type="aws:ec2/managedPrefixListEntry:ManagedPrefixListEntry")
public class ManagedPrefixListEntry extends com.pulumi.resources.CustomResource {
    /**
     * CIDR block of this entry.
     * 
     */
    @Export(name="cidr", refs={String.class}, tree="[0]")
    private Output<String> cidr;

    /**
     * @return CIDR block of this entry.
     * 
     */
    public Output<String> cidr() {
        return this.cidr;
    }
    /**
     * Description of this entry. Due to API limitations, updating only the description of an entry requires recreating the entry.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of this entry. Due to API limitations, updating only the description of an entry requires recreating the entry.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * CIDR block of this entry.
     * 
     */
    @Export(name="prefixListId", refs={String.class}, tree="[0]")
    private Output<String> prefixListId;

    /**
     * @return CIDR block of this entry.
     * 
     */
    public Output<String> prefixListId() {
        return this.prefixListId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ManagedPrefixListEntry(String name) {
        this(name, ManagedPrefixListEntryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ManagedPrefixListEntry(String name, ManagedPrefixListEntryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ManagedPrefixListEntry(String name, ManagedPrefixListEntryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/managedPrefixListEntry:ManagedPrefixListEntry", name, args == null ? ManagedPrefixListEntryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ManagedPrefixListEntry(String name, Output<String> id, @Nullable ManagedPrefixListEntryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:ec2/managedPrefixListEntry:ManagedPrefixListEntry", name, state, makeResourceOptions(options, id));
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
    public static ManagedPrefixListEntry get(String name, Output<String> id, @Nullable ManagedPrefixListEntryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ManagedPrefixListEntry(name, id, state, options);
    }
}
