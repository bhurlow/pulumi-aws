// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Route53RecoveryControl
{
    /// <summary>
    /// Provides an AWS Route 53 Recovery Control Config Routing Control.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.Route53RecoveryControl.RoutingControl("example", new()
    ///     {
    ///         ClusterArn = "arn:aws:route53-recovery-control::881188118811:cluster/8d47920e-d789-437d-803a-2dcc4b204393",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.Route53RecoveryControl.RoutingControl("example", new()
    ///     {
    ///         ClusterArn = "arn:aws:route53-recovery-control::881188118811:cluster/8d47920e-d789-437d-803a-2dcc4b204393",
    ///         ControlPanelArn = "arn:aws:route53-recovery-control::428113431245:controlpanel/abd5fbfc052d4844a082dbf400f61da8",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_route53recoverycontrolconfig_routing_control.mycontrol
    /// 
    ///  id = "arn:aws:route53-recovery-control::313517334327:controlpanel/abd5fbfc052d4844a082dbf400f61da8/routingcontrol/d5d90e587870494b" } Using `pulumi import`, import Route53 Recovery Control Config Routing Control using the routing control arn. For exampleconsole % pulumi import aws_route53recoverycontrolconfig_routing_control.mycontrol arn:aws:route53-recovery-control::313517334327:controlpanel/abd5fbfc052d4844a082dbf400f61da8/routingcontrol/d5d90e587870494b
    /// </summary>
    [AwsResourceType("aws:route53recoverycontrol/routingControl:RoutingControl")]
    public partial class RoutingControl : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ARN of the routing control.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// ARN of the cluster in which this routing control will reside.
        /// </summary>
        [Output("clusterArn")]
        public Output<string> ClusterArn { get; private set; } = null!;

        /// <summary>
        /// ARN of the control panel in which this routing control will reside.
        /// </summary>
        [Output("controlPanelArn")]
        public Output<string> ControlPanelArn { get; private set; } = null!;

        /// <summary>
        /// The name describing the routing control.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Status of routing control. `PENDING` when it is being created/updated, `PENDING_DELETION` when it is being deleted, and `DEPLOYED` otherwise.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a RoutingControl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RoutingControl(string name, RoutingControlArgs args, CustomResourceOptions? options = null)
            : base("aws:route53recoverycontrol/routingControl:RoutingControl", name, args ?? new RoutingControlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RoutingControl(string name, Input<string> id, RoutingControlState? state = null, CustomResourceOptions? options = null)
            : base("aws:route53recoverycontrol/routingControl:RoutingControl", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RoutingControl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RoutingControl Get(string name, Input<string> id, RoutingControlState? state = null, CustomResourceOptions? options = null)
        {
            return new RoutingControl(name, id, state, options);
        }
    }

    public sealed class RoutingControlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the cluster in which this routing control will reside.
        /// </summary>
        [Input("clusterArn", required: true)]
        public Input<string> ClusterArn { get; set; } = null!;

        /// <summary>
        /// ARN of the control panel in which this routing control will reside.
        /// </summary>
        [Input("controlPanelArn")]
        public Input<string>? ControlPanelArn { get; set; }

        /// <summary>
        /// The name describing the routing control.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public RoutingControlArgs()
        {
        }
        public static new RoutingControlArgs Empty => new RoutingControlArgs();
    }

    public sealed class RoutingControlState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the routing control.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// ARN of the cluster in which this routing control will reside.
        /// </summary>
        [Input("clusterArn")]
        public Input<string>? ClusterArn { get; set; }

        /// <summary>
        /// ARN of the control panel in which this routing control will reside.
        /// </summary>
        [Input("controlPanelArn")]
        public Input<string>? ControlPanelArn { get; set; }

        /// <summary>
        /// The name describing the routing control.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Status of routing control. `PENDING` when it is being created/updated, `PENDING_DELETION` when it is being deleted, and `DEPLOYED` otherwise.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public RoutingControlState()
        {
        }
        public static new RoutingControlState Empty => new RoutingControlState();
    }
}
