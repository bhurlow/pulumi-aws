// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudWatch
{
    /// <summary>
    /// Provides an EventBridge connection resource.
    /// 
    /// &gt; **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.
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
    ///     var test = new Aws.CloudWatch.EventConnection("test", new()
    ///     {
    ///         AuthParameters = new Aws.CloudWatch.Inputs.EventConnectionAuthParametersArgs
    ///         {
    ///             ApiKey = new Aws.CloudWatch.Inputs.EventConnectionAuthParametersApiKeyArgs
    ///             {
    ///                 Key = "x-signature",
    ///                 Value = "1234",
    ///             },
    ///         },
    ///         AuthorizationType = "API_KEY",
    ///         Description = "A connection description",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Basic Authorization
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test = new Aws.CloudWatch.EventConnection("test", new()
    ///     {
    ///         AuthParameters = new Aws.CloudWatch.Inputs.EventConnectionAuthParametersArgs
    ///         {
    ///             Basic = new Aws.CloudWatch.Inputs.EventConnectionAuthParametersBasicArgs
    ///             {
    ///                 Password = "Pass1234!",
    ///                 Username = "user",
    ///             },
    ///         },
    ///         AuthorizationType = "BASIC",
    ///         Description = "A connection description",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_cloudwatch_event_connection.test
    /// 
    ///  id = "ngrok-connection" } Using `pulumi import`, import EventBridge EventBridge connection using the `name`. For exampleconsole % pulumi import aws_cloudwatch_event_connection.test ngrok-connection
    /// </summary>
    [AwsResourceType("aws:cloudwatch/eventConnection:EventConnection")]
    public partial class EventConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the connection.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Parameters used for authorization. A maximum of 1 are allowed. Documented below.
        /// </summary>
        [Output("authParameters")]
        public Output<Outputs.EventConnectionAuthParameters> AuthParameters { get; private set; } = null!;

        /// <summary>
        /// Choose the type of authorization to use for the connection. One of `API_KEY`,`BASIC`,`OAUTH_CLIENT_CREDENTIALS`.
        /// </summary>
        [Output("authorizationType")]
        public Output<string> AuthorizationType { get; private set; } = null!;

        /// <summary>
        /// Enter a description for the connection. Maximum of 512 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the new connection. Maximum of 64 characters consisting of numbers, lower/upper case letters, .,-,_.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the secret created from the authorization parameters specified for the connection.
        /// </summary>
        [Output("secretArn")]
        public Output<string> SecretArn { get; private set; } = null!;


        /// <summary>
        /// Create a EventConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventConnection(string name, EventConnectionArgs args, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/eventConnection:EventConnection", name, args ?? new EventConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventConnection(string name, Input<string> id, EventConnectionState? state = null, CustomResourceOptions? options = null)
            : base("aws:cloudwatch/eventConnection:EventConnection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EventConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventConnection Get(string name, Input<string> id, EventConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new EventConnection(name, id, state, options);
        }
    }

    public sealed class EventConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Parameters used for authorization. A maximum of 1 are allowed. Documented below.
        /// </summary>
        [Input("authParameters", required: true)]
        public Input<Inputs.EventConnectionAuthParametersArgs> AuthParameters { get; set; } = null!;

        /// <summary>
        /// Choose the type of authorization to use for the connection. One of `API_KEY`,`BASIC`,`OAUTH_CLIENT_CREDENTIALS`.
        /// </summary>
        [Input("authorizationType", required: true)]
        public Input<string> AuthorizationType { get; set; } = null!;

        /// <summary>
        /// Enter a description for the connection. Maximum of 512 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the new connection. Maximum of 64 characters consisting of numbers, lower/upper case letters, .,-,_.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public EventConnectionArgs()
        {
        }
        public static new EventConnectionArgs Empty => new EventConnectionArgs();
    }

    public sealed class EventConnectionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the connection.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Parameters used for authorization. A maximum of 1 are allowed. Documented below.
        /// </summary>
        [Input("authParameters")]
        public Input<Inputs.EventConnectionAuthParametersGetArgs>? AuthParameters { get; set; }

        /// <summary>
        /// Choose the type of authorization to use for the connection. One of `API_KEY`,`BASIC`,`OAUTH_CLIENT_CREDENTIALS`.
        /// </summary>
        [Input("authorizationType")]
        public Input<string>? AuthorizationType { get; set; }

        /// <summary>
        /// Enter a description for the connection. Maximum of 512 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the new connection. Maximum of 64 characters consisting of numbers, lower/upper case letters, .,-,_.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the secret created from the authorization parameters specified for the connection.
        /// </summary>
        [Input("secretArn")]
        public Input<string>? SecretArn { get; set; }

        public EventConnectionState()
        {
        }
        public static new EventConnectionState Empty => new EventConnectionState();
    }
}
