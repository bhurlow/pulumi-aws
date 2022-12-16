// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.aws.controltower;

import com.pulumi.aws.Utilities;
import com.pulumi.aws.controltower.ControlTowerControlArgs;
import com.pulumi.aws.controltower.inputs.ControlTowerControlState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Allows the application of pre-defined controls to organizational units. For more information on usage, please see the
 * [AWS Control Tower User Guide](https://docs.aws.amazon.com/controltower/latest/userguide/enable-guardrails.html).
 * 
 * ## Import
 * 
 * Control Tower Controls can be imported using their `organizational_unit_arn/control_identifier`, e.g.,
 * 
 * ```sh
 *  $ pulumi import aws:controltower/controlTowerControl:ControlTowerControl example arn:aws:organizations::123456789101:ou/o-qqaejywet/ou-qg5o-ufbhdtv3,arn:aws:controltower:us-east-1::control/WTDSMKDKDNLE
 * ```
 * 
 */
@ResourceType(type="aws:controltower/controlTowerControl:ControlTowerControl")
public class ControlTowerControl extends com.pulumi.resources.CustomResource {
    /**
     * The ARN of the control. Only Strongly recommended and Elective controls are permitted, with the exception of the Region deny guardrail.
     * 
     */
    @Export(name="controlIdentifier", refs={String.class}, tree="[0]")
    private Output<String> controlIdentifier;

    /**
     * @return The ARN of the control. Only Strongly recommended and Elective controls are permitted, with the exception of the Region deny guardrail.
     * 
     */
    public Output<String> controlIdentifier() {
        return this.controlIdentifier;
    }
    /**
     * The ARN of the organizational unit.
     * 
     */
    @Export(name="targetIdentifier", refs={String.class}, tree="[0]")
    private Output<String> targetIdentifier;

    /**
     * @return The ARN of the organizational unit.
     * 
     */
    public Output<String> targetIdentifier() {
        return this.targetIdentifier;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ControlTowerControl(String name) {
        this(name, ControlTowerControlArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ControlTowerControl(String name, ControlTowerControlArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ControlTowerControl(String name, ControlTowerControlArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:controltower/controlTowerControl:ControlTowerControl", name, args == null ? ControlTowerControlArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ControlTowerControl(String name, Output<String> id, @Nullable ControlTowerControlState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("aws:controltower/controlTowerControl:ControlTowerControl", name, state, makeResourceOptions(options, id));
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
    public static ControlTowerControl get(String name, Output<String> id, @Nullable ControlTowerControlState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ControlTowerControl(name, id, state, options);
    }
}
