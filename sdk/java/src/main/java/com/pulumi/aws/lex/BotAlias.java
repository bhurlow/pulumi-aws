// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.lex;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.lex.BotAliasArgs;
import com.pulumi.aws.lex.inputs.BotAliasState;
import com.pulumi.aws.lex.outputs.BotAliasConversationLogs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an Amazon Lex Bot Alias resource. For more information see
 * [Amazon Lex: How It Works](https://docs.aws.amazon.com/lex/latest/dg/how-it-works.html)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.lex.BotAlias;
 * import com.pulumi.aws.lex.BotAliasArgs;
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
 *         var orderFlowersProd = new BotAlias(&#34;orderFlowersProd&#34;, BotAliasArgs.builder()        
 *             .botName(&#34;OrderFlowers&#34;)
 *             .botVersion(&#34;1&#34;)
 *             .description(&#34;Production Version of the OrderFlowers Bot.&#34;)
 *             .name(&#34;OrderFlowersProd&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Bot aliases can be imported using an ID with the format `bot_name:bot_alias_name`.
 * 
 * ```sh
 *  $ pulumi import aws:lex/botAlias:BotAlias order_flowers_prod OrderFlowers:OrderFlowersProd
 * ```
 * 
 */
@ResourceType(type="aws:lex/botAlias:BotAlias")
public class BotAlias extends com.pulumi.resources.CustomResource {
    /**
     * The ARN of the bot alias.
     * 
     */
    @Export(name="arn", refs={String.class}, tree="[0]")
    private Output<String> arn;

    /**
     * @return The ARN of the bot alias.
     * 
     */
    public Output<String> arn() {
        return this.arn;
    }
    /**
     * The name of the bot.
     * 
     */
    @Export(name="botName", refs={String.class}, tree="[0]")
    private Output<String> botName;

    /**
     * @return The name of the bot.
     * 
     */
    public Output<String> botName() {
        return this.botName;
    }
    /**
     * The name of the bot.
     * 
     */
    @Export(name="botVersion", refs={String.class}, tree="[0]")
    private Output<String> botVersion;

    /**
     * @return The name of the bot.
     * 
     */
    public Output<String> botVersion() {
        return this.botVersion;
    }
    /**
     * Checksum of the bot alias.
     * 
     */
    @Export(name="checksum", refs={String.class}, tree="[0]")
    private Output<String> checksum;

    /**
     * @return Checksum of the bot alias.
     * 
     */
    public Output<String> checksum() {
        return this.checksum;
    }
    /**
     * The settings that determine how Amazon Lex uses conversation logs for the alias. Attributes are documented under conversation_logs.
     * 
     */
    @Export(name="conversationLogs", refs={BotAliasConversationLogs.class}, tree="[0]")
    private Output</* @Nullable */ BotAliasConversationLogs> conversationLogs;

    /**
     * @return The settings that determine how Amazon Lex uses conversation logs for the alias. Attributes are documented under conversation_logs.
     * 
     */
    public Output<Optional<BotAliasConversationLogs>> conversationLogs() {
        return Codegen.optional(this.conversationLogs);
    }
    /**
     * The date that the bot alias was created.
     * 
     */
    @Export(name="createdDate", refs={String.class}, tree="[0]")
    private Output<String> createdDate;

    /**
     * @return The date that the bot alias was created.
     * 
     */
    public Output<String> createdDate() {
        return this.createdDate;
    }
    /**
     * A description of the alias. Must be less than or equal to 200 characters in length.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the alias. Must be less than or equal to 200 characters in length.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The date that the bot alias was updated. When you create a resource, the creation date and the last updated date are the same.
     * 
     */
    @Export(name="lastUpdatedDate", refs={String.class}, tree="[0]")
    private Output<String> lastUpdatedDate;

    /**
     * @return The date that the bot alias was updated. When you create a resource, the creation date and the last updated date are the same.
     * 
     */
    public Output<String> lastUpdatedDate() {
        return this.lastUpdatedDate;
    }
    /**
     * The name of the alias. The name is not case sensitive. Must be less than or equal to 100 characters in length.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the alias. The name is not case sensitive. Must be less than or equal to 100 characters in length.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BotAlias(String name) {
        this(name, BotAliasArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BotAlias(String name, BotAliasArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BotAlias(String name, BotAliasArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:lex/botAlias:BotAlias", name, args == null ? BotAliasArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BotAlias(String name, Output<String> id, @Nullable BotAliasState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:lex/botAlias:BotAlias", name, state, makeResourceOptions(options, id));
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
    public static BotAlias get(String name, Output<String> id, @Nullable BotAliasState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BotAlias(name, id, state, options);
    }
}
