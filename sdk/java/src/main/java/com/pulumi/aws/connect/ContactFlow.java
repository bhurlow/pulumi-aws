// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.connect;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.connect.ContactFlowArgs;
import com.pulumi.aws.connect.inputs.ContactFlowState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an Amazon Connect Contact Flow resource. For more information see
 * [Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html)
 * 
 * This resource embeds or references Contact Flows specified in Amazon Connect Contact Flow Language. For more information see
 * [Amazon Connect Flow language](https://docs.aws.amazon.com/connect/latest/adminguide/flow-language.html)
 * 
 * !&gt; **WARN:** Contact Flows exported from the Console [Contact Flow import/export](https://docs.aws.amazon.com/connect/latest/adminguide/contact-flow-import-export.html) are not in the Amazon Connect Contact Flow Language and can not be used with this resource. Instead, the recommendation is to use the AWS CLI [`describe-contact-flow`](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/connect/describe-contact-flow.html).
 * See example below which uses `jq` to extract the `Content` attribute and saves it to a local file.
 * 
 * ## Example Usage
 * ### Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.connect.ContactFlow;
 * import com.pulumi.aws.connect.ContactFlowArgs;
 * import static com.pulumi.codegen.internal.Serialization.*;
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
 *         var test = new ContactFlow(&#34;test&#34;, ContactFlowArgs.builder()        
 *             .instanceId(&#34;aaaaaaaa-bbbb-cccc-dddd-111111111111&#34;)
 *             .description(&#34;Test Contact Flow Description&#34;)
 *             .type(&#34;CONTACT_FLOW&#34;)
 *             .content(serializeJson(
 *                 jsonObject(
 *                     jsonProperty(&#34;Version&#34;, &#34;2019-10-30&#34;),
 *                     jsonProperty(&#34;StartAction&#34;, &#34;12345678-1234-1234-1234-123456789012&#34;),
 *                     jsonProperty(&#34;Actions&#34;, jsonArray(
 *                         jsonObject(
 *                             jsonProperty(&#34;Identifier&#34;, &#34;12345678-1234-1234-1234-123456789012&#34;),
 *                             jsonProperty(&#34;Type&#34;, &#34;MessageParticipant&#34;),
 *                             jsonProperty(&#34;Transitions&#34;, jsonObject(
 *                                 jsonProperty(&#34;NextAction&#34;, &#34;abcdef-abcd-abcd-abcd-abcdefghijkl&#34;),
 *                                 jsonProperty(&#34;Errors&#34;, jsonArray(
 *                                 )),
 *                                 jsonProperty(&#34;Conditions&#34;, jsonArray(
 *                                 ))
 *                             )),
 *                             jsonProperty(&#34;Parameters&#34;, jsonObject(
 *                                 jsonProperty(&#34;Text&#34;, &#34;Thanks for calling the sample flow!&#34;)
 *                             ))
 *                         ), 
 *                         jsonObject(
 *                             jsonProperty(&#34;Identifier&#34;, &#34;abcdef-abcd-abcd-abcd-abcdefghijkl&#34;),
 *                             jsonProperty(&#34;Type&#34;, &#34;DisconnectParticipant&#34;),
 *                             jsonProperty(&#34;Transitions&#34;, jsonObject(
 * 
 *                             )),
 *                             jsonProperty(&#34;Parameters&#34;, jsonObject(
 * 
 *                             ))
 *                         )
 *                     ))
 *                 )))
 *             .tags(Map.ofEntries(
 *                 Map.entry(&#34;Name&#34;, &#34;Test Contact Flow&#34;),
 *                 Map.entry(&#34;Application&#34;, &#34;Example&#34;),
 *                 Map.entry(&#34;Method&#34;, &#34;Create&#34;)
 *             ))
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
 *  to = aws_connect_contact_flow.example
 * 
 *  id = &#34;f1288a1f-6193-445a-b47e-af739b2:c1d4e5f6-1b3c-1b3c-1b3c-c1d4e5f6c1d4e5&#34; } Using `pulumi import`, import Amazon Connect Contact Flows using the `instance_id` and `contact_flow_id` separated by a colon (`:`). For exampleconsole % pulumi import aws_connect_contact_flow.example f1288a1f-6193-445a-b47e-af739b2:c1d4e5f6-1b3c-1b3c-1b3c-c1d4e5f6c1d4e5
 * 
 */
@ResourceType(type="aws:connect/contactFlow:ContactFlow")
public class ContactFlow extends com.pulumi.resources.CustomResource {
    /**
     * The Amazon Resource Name (ARN) of the Contact Flow.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The Amazon Resource Name (ARN) of the Contact Flow.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The identifier of the Contact Flow.
     * 
     */
    @Export(name="contactFlowId", refs={String.class}, tree="[0]")
    private Output<String> contactFlowId;

    /**
     * @return The identifier of the Contact Flow.
     * 
     */
    public Output<String> contactFlowId() {
        return this.contactFlowId;
    }
    /**
     * Specifies the content of the Contact Flow, provided as a JSON string, written in Amazon Connect Contact Flow Language. If defined, the `filename` argument cannot be used.
     * 
     */
    @Export(name="content", refs={String.class}, tree="[0]")
    private Output<String> content;

    /**
     * @return Specifies the content of the Contact Flow, provided as a JSON string, written in Amazon Connect Contact Flow Language. If defined, the `filename` argument cannot be used.
     * 
     */
    public Output<String> content() {
        return this.content;
    }
    /**
     * Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the Contact Flow source specified with `filename`.
     * 
     */
    @Export(name="contentHash", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> contentHash;

    /**
     * @return Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the Contact Flow source specified with `filename`.
     * 
     */
    public Output<Optional<String>> contentHash() {
        return Codegen.optional(this.contentHash);
    }
    /**
     * Specifies the description of the Contact Flow.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Specifies the description of the Contact Flow.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The path to the Contact Flow source within the local filesystem. Conflicts with `content`.
     * 
     */
    @Export(name="filename", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> filename;

    /**
     * @return The path to the Contact Flow source within the local filesystem. Conflicts with `content`.
     * 
     */
    public Output<Optional<String>> filename() {
        return Codegen.optional(this.filename);
    }
    /**
     * Specifies the identifier of the hosting Amazon Connect Instance.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return Specifies the identifier of the hosting Amazon Connect Instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * Specifies the name of the Contact Flow.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Specifies the name of the Contact Flow.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Tags to apply to the Contact Flow. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return Tags to apply to the Contact Flow. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
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
     * Specifies the type of the Contact Flow. Defaults to `CONTACT_FLOW`. Allowed Values are: `CONTACT_FLOW`, `CUSTOMER_QUEUE`, `CUSTOMER_HOLD`, `CUSTOMER_WHISPER`, `AGENT_HOLD`, `AGENT_WHISPER`, `OUTBOUND_WHISPER`, `AGENT_TRANSFER`, `QUEUE_TRANSFER`.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    /**
     * @return Specifies the type of the Contact Flow. Defaults to `CONTACT_FLOW`. Allowed Values are: `CONTACT_FLOW`, `CUSTOMER_QUEUE`, `CUSTOMER_HOLD`, `CUSTOMER_WHISPER`, `AGENT_HOLD`, `AGENT_WHISPER`, `OUTBOUND_WHISPER`, `AGENT_TRANSFER`, `QUEUE_TRANSFER`.
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ContactFlow(String name) {
        this(name, ContactFlowArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ContactFlow(String name, ContactFlowArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ContactFlow(String name, ContactFlowArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:connect/contactFlow:ContactFlow", name, args == null ? ContactFlowArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ContactFlow(String name, Output<String> id, @Nullable ContactFlowState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:connect/contactFlow:ContactFlow", name, state, makeResourceOptions(options, id));
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
    public static ContactFlow get(String name, Output<String> id, @Nullable ContactFlowState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ContactFlow(name, id, state, options);
    }
}
