// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.sesv2;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.sesv2.EmailIdentityFeedbackAttributesArgs;
import com.pulumi.aws.sesv2.inputs.EmailIdentityFeedbackAttributesState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Resource for managing an AWS SESv2 (Simple Email V2) Email Identity Feedback Attributes.
 * 
 * ## Example Usage
 * ### Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.sesv2.EmailIdentity;
 * import com.pulumi.aws.sesv2.EmailIdentityArgs;
 * import com.pulumi.aws.sesv2.EmailIdentityFeedbackAttributes;
 * import com.pulumi.aws.sesv2.EmailIdentityFeedbackAttributesArgs;
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
 *         var exampleEmailIdentity = new EmailIdentity(&#34;exampleEmailIdentity&#34;, EmailIdentityArgs.builder()        
 *             .emailIdentity(&#34;example.com&#34;)
 *             .build());
 * 
 *         var exampleEmailIdentityFeedbackAttributes = new EmailIdentityFeedbackAttributes(&#34;exampleEmailIdentityFeedbackAttributes&#34;, EmailIdentityFeedbackAttributesArgs.builder()        
 *             .emailIdentity(exampleEmailIdentity.emailIdentity())
 *             .emailForwardingEnabled(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * SESv2 (Simple Email V2) Email Identity Feedback Attributes can be imported using the `email_identity`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:sesv2/emailIdentityFeedbackAttributes:EmailIdentityFeedbackAttributes example example.com
 * ```
 * 
 */
@ResourceType(type="aws:sesv2/emailIdentityFeedbackAttributes:EmailIdentityFeedbackAttributes")
public class EmailIdentityFeedbackAttributes extends com.pulumi.resources.CustomResource {
    /**
     * Sets the feedback forwarding configuration for the identity.
     * 
     */
    @Export(name="emailForwardingEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> emailForwardingEnabled;

    /**
     * @return Sets the feedback forwarding configuration for the identity.
     * 
     */
    public Output<Optional<Boolean>> emailForwardingEnabled() {
        return Codegen.optional(this.emailForwardingEnabled);
    }
    /**
     * The email identity.
     * 
     */
    @Export(name="emailIdentity", refs={String.class}, tree="[0]")
    private Output<String> emailIdentity;

    /**
     * @return The email identity.
     * 
     */
    public Output<String> emailIdentity() {
        return this.emailIdentity;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EmailIdentityFeedbackAttributes(String name) {
        this(name, EmailIdentityFeedbackAttributesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EmailIdentityFeedbackAttributes(String name, EmailIdentityFeedbackAttributesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EmailIdentityFeedbackAttributes(String name, EmailIdentityFeedbackAttributesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:sesv2/emailIdentityFeedbackAttributes:EmailIdentityFeedbackAttributes", name, args == null ? EmailIdentityFeedbackAttributesArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EmailIdentityFeedbackAttributes(String name, Output<String> id, @Nullable EmailIdentityFeedbackAttributesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:sesv2/emailIdentityFeedbackAttributes:EmailIdentityFeedbackAttributes", name, state, makeResourceOptions(options, id));
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
    public static EmailIdentityFeedbackAttributes get(String name, Output<String> id, @Nullable EmailIdentityFeedbackAttributesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EmailIdentityFeedbackAttributes(name, id, state, options);
    }
}
