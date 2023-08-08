// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SsmContacts
{
    /// <summary>
    /// Resource for managing an AWS SSM Contact.
    /// 
    /// ## Example Usage
    /// ### Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.SsmContacts.Contact("example", new()
    ///     {
    ///         Alias = "alias",
    ///         Type = "PERSONAL",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             aws_ssmincidents_replication_set.Example,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Usage With All Fields
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Aws.SsmContacts.Contact("example", new()
    ///     {
    ///         Alias = "alias",
    ///         DisplayName = "displayName",
    ///         Type = "ESCALATION",
    ///         Tags = 
    ///         {
    ///             { "key", "value" },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             aws_ssmincidents_replication_set.Example,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform import {
    /// 
    ///  to = aws_ssmcontacts_contact.example
    /// 
    ///  id = "{ARNValue}" } Using `pulumi import`, import SSM Contact using the `ARN`. For exampleconsole % pulumi import aws_ssmcontacts_contact.example {ARNValue}
    /// </summary>
    [AwsResourceType("aws:ssmcontacts/contact:Contact")]
    public partial class Contact : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A unique and identifiable alias for the contact or escalation plan.
        /// </summary>
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the contact or escalation plan.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Full friendly name of the contact or escalation plan.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Map of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        /// <summary>
        /// The type of contact engaged. A single contact is type PERSONAL and an escalation
        /// plan is type ESCALATION.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Contact resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Contact(string name, ContactArgs args, CustomResourceOptions? options = null)
            : base("aws:ssmcontacts/contact:Contact", name, args ?? new ContactArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Contact(string name, Input<string> id, ContactState? state = null, CustomResourceOptions? options = null)
            : base("aws:ssmcontacts/contact:Contact", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Contact resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Contact Get(string name, Input<string> id, ContactState? state = null, CustomResourceOptions? options = null)
        {
            return new Contact(name, id, state, options);
        }
    }

    public sealed class ContactArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique and identifiable alias for the contact or escalation plan.
        /// </summary>
        [Input("alias", required: true)]
        public Input<string> Alias { get; set; } = null!;

        /// <summary>
        /// Full friendly name of the contact or escalation plan.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of contact engaged. A single contact is type PERSONAL and an escalation
        /// plan is type ESCALATION.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ContactArgs()
        {
        }
        public static new ContactArgs Empty => new ContactArgs();
    }

    public sealed class ContactState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique and identifiable alias for the contact or escalation plan.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the contact or escalation plan.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Full friendly name of the contact or escalation plan.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Map of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;

        /// <summary>
        /// Map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        /// </summary>
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        /// <summary>
        /// The type of contact engaged. A single contact is type PERSONAL and an escalation
        /// plan is type ESCALATION.
        /// 
        /// The following arguments are optional:
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ContactState()
        {
        }
        public static new ContactState Empty => new ContactState();
    }
}
