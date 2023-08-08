// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppConfig
{
    /// <summary>
    /// Provides an AppConfig Hosted Configuration Version resource.
    /// 
    /// ## Example Usage
    /// ### Freeform
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.AppConfig.HostedConfigurationVersion("example", new()
    ///     {
    ///         ApplicationId = aws_appconfig_application.Example.Id,
    ///         ConfigurationProfileId = aws_appconfig_configuration_profile.Example.Configuration_profile_id,
    ///         Description = "Example Freeform Hosted Configuration Version",
    ///         ContentType = "application/json",
    ///         Content = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["foo"] = "bar",
    ///             ["fruit"] = new[]
    ///             {
    ///                 "apple",
    ///                 "pear",
    ///                 "orange",
    ///             },
    ///             ["isThingEnabled"] = true,
    ///         }),
    ///     });
    /// 
    /// });
    /// ```
    /// ### Feature Flags
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.AppConfig.HostedConfigurationVersion("example", new()
    ///     {
    ///         ApplicationId = aws_appconfig_application.Example.Id,
    ///         ConfigurationProfileId = aws_appconfig_configuration_profile.Example.Configuration_profile_id,
    ///         Description = "Example Feature Flag Configuration Version",
    ///         ContentType = "application/json",
    ///         Content = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["flags"] = new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 ["foo"] = new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["name"] = "foo",
    ///                     ["_deprecation"] = new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["status"] = "planned",
    ///                     },
    ///                 },
    ///                 ["bar"] = new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["name"] = "bar",
    ///                     ["attributes"] = new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["someAttribute"] = new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             ["constraints"] = new Dictionary&lt;string, object?&gt;
    ///                             {
    ///                                 ["type"] = "string",
    ///                                 ["required"] = true,
    ///                             },
    ///                         },
    ///                         ["someOtherAttribute"] = new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             ["constraints"] = new Dictionary&lt;string, object?&gt;
    ///                             {
    ///                                 ["type"] = "number",
    ///                                 ["required"] = true,
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///             ["values"] = new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 ["foo"] = new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["enabled"] = "true",
    ///                 },
    ///                 ["bar"] = new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["enabled"] = "true",
    ///                     ["someAttribute"] = "Hello World",
    ///                     ["someOtherAttribute"] = 123,
    ///                 },
    ///             },
    ///             ["version"] = "1",
    ///         }),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_appconfig_hosted_configuration_version.example
    /// 
    ///  id = "71abcde/11xxxxx/2" } Using `pulumi import`, import AppConfig Hosted Configuration Versions using the application ID, configuration profile ID, and version number separated by a slash (`/`). For exampleconsole % pulumi import aws_appconfig_hosted_configuration_version.example 71abcde/11xxxxx/2
    /// </summary>
    [AwsResourceType("aws:appconfig/hostedConfigurationVersion:HostedConfigurationVersion")]
    public partial class HostedConfigurationVersion : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Application ID.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// ARN of the AppConfig  hosted configuration version.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Configuration profile ID.
        /// </summary>
        [Output("configurationProfileId")]
        public Output<string> ConfigurationProfileId { get; private set; } = null!;

        /// <summary>
        /// Content of the configuration or the configuration data.
        /// </summary>
        [Output("content")]
        public Output<string> Content { get; private set; } = null!;

        /// <summary>
        /// Standard MIME type describing the format of the configuration content. For more information, see [Content-Type](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17).
        /// </summary>
        [Output("contentType")]
        public Output<string> ContentType { get; private set; } = null!;

        /// <summary>
        /// Description of the configuration.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Version number of the hosted configuration.
        /// </summary>
        [Output("versionNumber")]
        public Output<int> VersionNumber { get; private set; } = null!;


        /// <summary>
        /// Create a HostedConfigurationVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HostedConfigurationVersion(string name, HostedConfigurationVersionArgs args, CustomResourceOptions? options = null)
            : base("aws:appconfig/hostedConfigurationVersion:HostedConfigurationVersion", name, args ?? new HostedConfigurationVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HostedConfigurationVersion(string name, Input<string> id, HostedConfigurationVersionState? state = null, CustomResourceOptions? options = null)
            : base("aws:appconfig/hostedConfigurationVersion:HostedConfigurationVersion", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "content",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing HostedConfigurationVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HostedConfigurationVersion Get(string name, Input<string> id, HostedConfigurationVersionState? state = null, CustomResourceOptions? options = null)
        {
            return new HostedConfigurationVersion(name, id, state, options);
        }
    }

    public sealed class HostedConfigurationVersionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application ID.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        /// <summary>
        /// Configuration profile ID.
        /// </summary>
        [Input("configurationProfileId", required: true)]
        public Input<string> ConfigurationProfileId { get; set; } = null!;

        [Input("content", required: true)]
        private Input<string>? _content;

        /// <summary>
        /// Content of the configuration or the configuration data.
        /// </summary>
        public Input<string>? Content
        {
            get => _content;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _content = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Standard MIME type describing the format of the configuration content. For more information, see [Content-Type](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17).
        /// </summary>
        [Input("contentType", required: true)]
        public Input<string> ContentType { get; set; } = null!;

        /// <summary>
        /// Description of the configuration.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        public HostedConfigurationVersionArgs()
        {
        }
        public static new HostedConfigurationVersionArgs Empty => new HostedConfigurationVersionArgs();
    }

    public sealed class HostedConfigurationVersionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application ID.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// ARN of the AppConfig  hosted configuration version.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Configuration profile ID.
        /// </summary>
        [Input("configurationProfileId")]
        public Input<string>? ConfigurationProfileId { get; set; }

        [Input("content")]
        private Input<string>? _content;

        /// <summary>
        /// Content of the configuration or the configuration data.
        /// </summary>
        public Input<string>? Content
        {
            get => _content;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _content = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Standard MIME type describing the format of the configuration content. For more information, see [Content-Type](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17).
        /// </summary>
        [Input("contentType")]
        public Input<string>? ContentType { get; set; }

        /// <summary>
        /// Description of the configuration.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Version number of the hosted configuration.
        /// </summary>
        [Input("versionNumber")]
        public Input<int>? VersionNumber { get; set; }

        public HostedConfigurationVersionState()
        {
        }
        public static new HostedConfigurationVersionState Empty => new HostedConfigurationVersionState();
    }
}
