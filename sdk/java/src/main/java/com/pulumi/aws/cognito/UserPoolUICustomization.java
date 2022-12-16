// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.cognito;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.cognito.UserPoolUICustomizationArgs;
import com.pulumi.aws.cognito.inputs.UserPoolUICustomizationState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Cognito User Pool UI Customization resource.
 * 
 * &gt; **Note:** To use this resource, the user pool must have a domain associated with it. For more information, see the Amazon Cognito Developer Guide on [Customizing the Built-in Sign-In and Sign-up Webpages](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-app-ui-customization.html).
 * 
 * ## Example Usage
 * ### UI customization settings for a single client
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.cognito.UserPool;
 * import com.pulumi.aws.cognito.UserPoolDomain;
 * import com.pulumi.aws.cognito.UserPoolDomainArgs;
 * import com.pulumi.aws.cognito.UserPoolClient;
 * import com.pulumi.aws.cognito.UserPoolClientArgs;
 * import com.pulumi.aws.cognito.UserPoolUICustomization;
 * import com.pulumi.aws.cognito.UserPoolUICustomizationArgs;
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
 *         var exampleUserPool = new UserPool(&#34;exampleUserPool&#34;);
 * 
 *         var exampleUserPoolDomain = new UserPoolDomain(&#34;exampleUserPoolDomain&#34;, UserPoolDomainArgs.builder()        
 *             .domain(&#34;example&#34;)
 *             .userPoolId(exampleUserPool.id())
 *             .build());
 * 
 *         var exampleUserPoolClient = new UserPoolClient(&#34;exampleUserPoolClient&#34;, UserPoolClientArgs.builder()        
 *             .userPoolId(exampleUserPool.id())
 *             .build());
 * 
 *         var exampleUserPoolUICustomization = new UserPoolUICustomization(&#34;exampleUserPoolUICustomization&#34;, UserPoolUICustomizationArgs.builder()        
 *             .clientId(exampleUserPoolClient.id())
 *             .css(&#34;.label-customizable {font-weight: 400;}&#34;)
 *             .imageFile(Base64.getEncoder().encodeToString(Files.readAllBytes(Paths.get(&#34;logo.png&#34;))))
 *             .userPoolId(exampleUserPoolDomain.userPoolId())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### UI customization settings for all clients
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.cognito.UserPool;
 * import com.pulumi.aws.cognito.UserPoolDomain;
 * import com.pulumi.aws.cognito.UserPoolDomainArgs;
 * import com.pulumi.aws.cognito.UserPoolUICustomization;
 * import com.pulumi.aws.cognito.UserPoolUICustomizationArgs;
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
 *         var exampleUserPool = new UserPool(&#34;exampleUserPool&#34;);
 * 
 *         var exampleUserPoolDomain = new UserPoolDomain(&#34;exampleUserPoolDomain&#34;, UserPoolDomainArgs.builder()        
 *             .domain(&#34;example&#34;)
 *             .userPoolId(exampleUserPool.id())
 *             .build());
 * 
 *         var exampleUserPoolUICustomization = new UserPoolUICustomization(&#34;exampleUserPoolUICustomization&#34;, UserPoolUICustomizationArgs.builder()        
 *             .css(&#34;.label-customizable {font-weight: 400;}&#34;)
 *             .imageFile(Base64.getEncoder().encodeToString(Files.readAllBytes(Paths.get(&#34;logo.png&#34;))))
 *             .userPoolId(exampleUserPoolDomain.userPoolId())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Cognito User Pool UI Customizations can be imported using the `user_pool_id` and `client_id` separated by `,`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:cognito/userPoolUICustomization:UserPoolUICustomization example us-west-2_ZCTarbt5C,12bu4fuk3mlgqa2rtrujgp6egq
 * ```
 * 
 */
@ResourceType(type="aws:cognito/userPoolUICustomization:UserPoolUICustomization")
public class UserPoolUICustomization extends com.pulumi.resources.CustomResource {
    /**
     * The client ID for the client app. Defaults to `ALL`. If `ALL` is specified, the `css` and/or `image_file` settings will be used for every client that has no UI customization set previously.
     * 
     */
    @Export(name="clientId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientId;

    /**
     * @return The client ID for the client app. Defaults to `ALL`. If `ALL` is specified, the `css` and/or `image_file` settings will be used for every client that has no UI customization set previously.
     * 
     */
    public Output<Optional<String>> clientId() {
        return Codegen.optional(this.clientId);
    }
    /**
     * The creation date in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) for the UI customization.
     * 
     */
    @Export(name="creationDate", refs={String.class}, tree="[0]")
    private Output<String> creationDate;

    /**
     * @return The creation date in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) for the UI customization.
     * 
     */
    public Output<String> creationDate() {
        return this.creationDate;
    }
    /**
     * The CSS values in the UI customization, provided as a String. At least one of `css` or `image_file` is required.
     * 
     */
    @Export(name="css", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> css;

    /**
     * @return The CSS values in the UI customization, provided as a String. At least one of `css` or `image_file` is required.
     * 
     */
    public Output<Optional<String>> css() {
        return Codegen.optional(this.css);
    }
    /**
     * The CSS version number.
     * 
     */
    @Export(name="cssVersion", refs={String.class}, tree="[0]")
    private Output<String> cssVersion;

    /**
     * @return The CSS version number.
     * 
     */
    public Output<String> cssVersion() {
        return this.cssVersion;
    }
    /**
     * The uploaded logo image for the UI customization, provided as a base64-encoded String. Drift detection is not possible for this argument. At least one of `css` or `image_file` is required.
     * 
     */
    @Export(name="imageFile", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> imageFile;

    /**
     * @return The uploaded logo image for the UI customization, provided as a base64-encoded String. Drift detection is not possible for this argument. At least one of `css` or `image_file` is required.
     * 
     */
    public Output<Optional<String>> imageFile() {
        return Codegen.optional(this.imageFile);
    }
    /**
     * The logo image URL for the UI customization.
     * 
     */
    @Export(name="imageUrl", refs={String.class}, tree="[0]")
    private Output<String> imageUrl;

    /**
     * @return The logo image URL for the UI customization.
     * 
     */
    public Output<String> imageUrl() {
        return this.imageUrl;
    }
    /**
     * The last-modified date in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) for the UI customization.
     * 
     */
    @Export(name="lastModifiedDate", refs={String.class}, tree="[0]")
    private Output<String> lastModifiedDate;

    /**
     * @return The last-modified date in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) for the UI customization.
     * 
     */
    public Output<String> lastModifiedDate() {
        return this.lastModifiedDate;
    }
    /**
     * The user pool ID for the user pool.
     * 
     */
    @Export(name="userPoolId", refs={String.class}, tree="[0]")
    private Output<String> userPoolId;

    /**
     * @return The user pool ID for the user pool.
     * 
     */
    public Output<String> userPoolId() {
        return this.userPoolId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public UserPoolUICustomization(String name) {
        this(name, UserPoolUICustomizationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public UserPoolUICustomization(String name, UserPoolUICustomizationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public UserPoolUICustomization(String name, UserPoolUICustomizationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cognito/userPoolUICustomization:UserPoolUICustomization", name, args == null ? UserPoolUICustomizationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private UserPoolUICustomization(String name, Output<String> id, @Nullable UserPoolUICustomizationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:cognito/userPoolUICustomization:UserPoolUICustomization", name, state, makeResourceOptions(options, id));
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
    public static UserPoolUICustomization get(String name, Output<String> id, @Nullable UserPoolUICustomizationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new UserPoolUICustomization(name, id, state, options);
    }
}
