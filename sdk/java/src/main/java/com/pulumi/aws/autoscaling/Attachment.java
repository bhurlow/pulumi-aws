// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.autoscaling;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.autoscaling.AttachmentArgs;
import com.pulumi.aws.autoscaling.inputs.AttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an Auto Scaling Attachment resource.
 * 
 * &gt; **NOTE on Auto Scaling Groups and ASG Attachments:** This provider currently provides
 * both a standalone `aws.autoscaling.Attachment` resource
 * (describing an ASG attached to an ELB or ALB), and an `aws.autoscaling.Group`
 * with `load_balancers` and `target_group_arns` defined in-line. These two methods are not
 * mutually-exclusive. If `aws.autoscaling.Attachment` resources are used, either alone or with inline
 * `load_balancers` or `target_group_arns`, the `aws.autoscaling.Group` resource must be configured
 * to ignore changes to the `load_balancers` and `target_group_arns` arguments.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.autoscaling.Attachment;
 * import com.pulumi.aws.autoscaling.AttachmentArgs;
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
 *         var asgAttachmentBar = new Attachment(&#34;asgAttachmentBar&#34;, AttachmentArgs.builder()        
 *             .autoscalingGroupName(aws_autoscaling_group.asg().id())
 *             .elb(aws_elb.bar().id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.autoscaling.Attachment;
 * import com.pulumi.aws.autoscaling.AttachmentArgs;
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
 *         var asgAttachmentBar = new Attachment(&#34;asgAttachmentBar&#34;, AttachmentArgs.builder()        
 *             .autoscalingGroupName(aws_autoscaling_group.asg().id())
 *             .lbTargetGroupArn(aws_lb_target_group.test().arn())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## With An AutoScaling Group Resource
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.aws.autoscaling.Group;
 * import com.pulumi.aws.autoscaling.Attachment;
 * import com.pulumi.aws.autoscaling.AttachmentArgs;
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
 *         var asg = new Group(&#34;asg&#34;);
 * 
 *         var asgAttachmentBar = new Attachment(&#34;asgAttachmentBar&#34;, AttachmentArgs.builder()        
 *             .autoscalingGroupName(asg.id())
 *             .elb(aws_elb.test().id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="aws:autoscaling/attachment:Attachment")
public class Attachment extends com.pulumi.resources.CustomResource {
    /**
     * ARN of an ALB Target Group.
     * 
     * @deprecated
     * Use lb_target_group_arn instead
     * 
     */
    @Deprecated /* Use lb_target_group_arn instead */
    @Export(name="albTargetGroupArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> albTargetGroupArn;

    /**
     * @return ARN of an ALB Target Group.
     * 
     */
    public Output<Optional<String>> albTargetGroupArn() {
        return Codegen.optional(this.albTargetGroupArn);
    }
    /**
     * Name of ASG to associate with the ELB.
     * 
     */
    @Export(name="autoscalingGroupName", refs={String.class}, tree="[0]")
    private Output<String> autoscalingGroupName;

    /**
     * @return Name of ASG to associate with the ELB.
     * 
     */
    public Output<String> autoscalingGroupName() {
        return this.autoscalingGroupName;
    }
    /**
     * Name of the ELB.
     * 
     */
    @Export(name="elb", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> elb;

    /**
     * @return Name of the ELB.
     * 
     */
    public Output<Optional<String>> elb() {
        return Codegen.optional(this.elb);
    }
    /**
     * ARN of a load balancer target group.
     * 
     */
    @Export(name="lbTargetGroupArn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> lbTargetGroupArn;

    /**
     * @return ARN of a load balancer target group.
     * 
     */
    public Output<Optional<String>> lbTargetGroupArn() {
        return Codegen.optional(this.lbTargetGroupArn);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Attachment(String name) {
        this(name, AttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Attachment(String name, AttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Attachment(String name, AttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:autoscaling/attachment:Attachment", name, args == null ? AttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Attachment(String name, Output<String> id, @Nullable AttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:autoscaling/attachment:Attachment", name, state, makeResourceOptions(options, id));
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
    public static Attachment get(String name, Output<String> id, @Nullable AttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Attachment(name, id, state, options);
    }
}
