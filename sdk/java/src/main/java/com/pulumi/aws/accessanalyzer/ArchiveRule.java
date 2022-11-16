// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.accessanalyzer;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.accessanalyzer.ArchiveRuleArgs;
import com.pulumi.aws.accessanalyzer.inputs.ArchiveRuleState;
import com.pulumi.aws.accessanalyzer.outputs.ArchiveRuleFilter;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Resource for managing an AWS AccessAnalyzer Archive Rule.
 * 
 * ## Example Usage
 * ### Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.accessanalyzer.ArchiveRule;
 * import com.pulumi.aws.accessanalyzer.ArchiveRuleArgs;
 * import com.pulumi.aws.accessanalyzer.inputs.ArchiveRuleFilterArgs;
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
 *         var example = new ArchiveRule(&#34;example&#34;, ArchiveRuleArgs.builder()        
 *             .analyzerName(&#34;example-analyzer&#34;)
 *             .filters(            
 *                 ArchiveRuleFilterArgs.builder()
 *                     .criteria(&#34;condition.aws:UserId&#34;)
 *                     .eqs(&#34;userid&#34;)
 *                     .build(),
 *                 ArchiveRuleFilterArgs.builder()
 *                     .criteria(&#34;error&#34;)
 *                     .exists(true)
 *                     .build(),
 *                 ArchiveRuleFilterArgs.builder()
 *                     .criteria(&#34;isPublic&#34;)
 *                     .eqs(&#34;false&#34;)
 *                     .build())
 *             .ruleName(&#34;example-rule&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * AccessAnalyzer ArchiveRule can be imported using the `analyzer_name/rule_name`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:accessanalyzer/archiveRule:ArchiveRule example example-analyzer/example-rule
 * ```
 * 
 */
@ResourceType(type="aws:accessanalyzer/archiveRule:ArchiveRule")
public class ArchiveRule extends com.pulumi.resources.CustomResource {
    /**
     * Analyzer name.
     * 
     */
    @Export(name="analyzerName", type=String.class, parameters={})
    private Output<String> analyzerName;

    /**
     * @return Analyzer name.
     * 
     */
    public Output<String> analyzerName() {
        return this.analyzerName;
    }
    /**
     * Filter criteria for the archive rule. See Filter for more details.
     * 
     */
    @Export(name="filters", type=List.class, parameters={ArchiveRuleFilter.class})
    private Output<List<ArchiveRuleFilter>> filters;

    /**
     * @return Filter criteria for the archive rule. See Filter for more details.
     * 
     */
    public Output<List<ArchiveRuleFilter>> filters() {
        return this.filters;
    }
    /**
     * Rule name.
     * 
     */
    @Export(name="ruleName", type=String.class, parameters={})
    private Output<String> ruleName;

    /**
     * @return Rule name.
     * 
     */
    public Output<String> ruleName() {
        return this.ruleName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ArchiveRule(String name) {
        this(name, ArchiveRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ArchiveRule(String name, ArchiveRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ArchiveRule(String name, ArchiveRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:accessanalyzer/archiveRule:ArchiveRule", name, args == null ? ArchiveRuleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ArchiveRule(String name, Output<String> id, @Nullable ArchiveRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:accessanalyzer/archiveRule:ArchiveRule", name, state, makeResourceOptions(options, id));
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
    public static ArchiveRule get(String name, Output<String> id, @Nullable ArchiveRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ArchiveRule(name, id, state, options);
    }
}
