// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.wafregional;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.wafregional.XssMatchSetArgs;
import com.pulumi.aws.wafregional.inputs.XssMatchSetState;
import com.pulumi.aws.wafregional.outputs.XssMatchSetXssMatchTuple;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a WAF Regional XSS Match Set Resource for use with Application Load Balancer.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.wafregional.XssMatchSet;
 * import com.pulumi.aws.wafregional.XssMatchSetArgs;
 * import com.pulumi.aws.wafregional.inputs.XssMatchSetXssMatchTupleArgs;
 * import com.pulumi.aws.wafregional.inputs.XssMatchSetXssMatchTupleFieldToMatchArgs;
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
 *         var xssMatchSet = new XssMatchSet(&#34;xssMatchSet&#34;, XssMatchSetArgs.builder()        
 *             .xssMatchTuples(            
 *                 XssMatchSetXssMatchTupleArgs.builder()
 *                     .fieldToMatch(XssMatchSetXssMatchTupleFieldToMatchArgs.builder()
 *                         .type(&#34;URI&#34;)
 *                         .build())
 *                     .textTransformation(&#34;NONE&#34;)
 *                     .build(),
 *                 XssMatchSetXssMatchTupleArgs.builder()
 *                     .fieldToMatch(XssMatchSetXssMatchTupleFieldToMatchArgs.builder()
 *                         .type(&#34;QUERY_STRING&#34;)
 *                         .build())
 *                     .textTransformation(&#34;NONE&#34;)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * AWS WAF Regional XSS Match can be imported using the `id`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:wafregional/xssMatchSet:XssMatchSet example 12345abcde
 * ```
 * 
 */
@ResourceType(type="aws:wafregional/xssMatchSet:XssMatchSet")
public class XssMatchSet extends com.pulumi.resources.CustomResource {
    /**
     * The name of the set
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the set
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The parts of web requests that you want to inspect for cross-site scripting attacks.
     * 
     */
    @Export(name="xssMatchTuples", refs={List.class,XssMatchSetXssMatchTuple.class}, tree="[0,1]")
    private Output</* @Nullable */ List<XssMatchSetXssMatchTuple>> xssMatchTuples;

    /**
     * @return The parts of web requests that you want to inspect for cross-site scripting attacks.
     * 
     */
    public Output<Optional<List<XssMatchSetXssMatchTuple>>> xssMatchTuples() {
        return Codegen.optional(this.xssMatchTuples);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public XssMatchSet(String name) {
        this(name, XssMatchSetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public XssMatchSet(String name, @Nullable XssMatchSetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public XssMatchSet(String name, @Nullable XssMatchSetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:wafregional/xssMatchSet:XssMatchSet", name, args == null ? XssMatchSetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private XssMatchSet(String name, Output<String> id, @Nullable XssMatchSetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:wafregional/xssMatchSet:XssMatchSet", name, state, makeResourceOptions(options, id));
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
    public static XssMatchSet get(String name, Output<String> id, @Nullable XssMatchSetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new XssMatchSet(name, id, state, options);
    }
}
