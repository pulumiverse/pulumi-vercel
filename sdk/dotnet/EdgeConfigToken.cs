// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Vercel
{
    [VercelResourceType("vercel:index/edgeConfigToken:EdgeConfigToken")]
    public partial class EdgeConfigToken : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A connection string is a URL that connects a project to an Edge Config. The variable can be called anything, but our
        /// Edge Config client SDK will search for process.env.EDGE_CONFIG by default.
        /// </summary>
        [Output("connectionString")]
        public Output<string> ConnectionString { get; private set; } = null!;

        /// <summary>
        /// The ID of the Edge Config store.
        /// </summary>
        [Output("edgeConfigId")]
        public Output<string> EdgeConfigId { get; private set; } = null!;

        /// <summary>
        /// The label of the Edge Config Token.
        /// </summary>
        [Output("label")]
        public Output<string> Label { get; private set; } = null!;

        /// <summary>
        /// The ID of the team the Edge Config should exist under. Required when configuring a team resource if a default team has
        /// not been set in the provider.
        /// </summary>
        [Output("teamId")]
        public Output<string> TeamId { get; private set; } = null!;

        /// <summary>
        /// A read access token used for authenticating against the Edge Config's endpoint for high volume, low-latency requests.
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;


        /// <summary>
        /// Create a EdgeConfigToken resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EdgeConfigToken(string name, EdgeConfigTokenArgs args, CustomResourceOptions? options = null)
            : base("vercel:index/edgeConfigToken:EdgeConfigToken", name, args ?? new EdgeConfigTokenArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EdgeConfigToken(string name, Input<string> id, EdgeConfigTokenState? state = null, CustomResourceOptions? options = null)
            : base("vercel:index/edgeConfigToken:EdgeConfigToken", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EdgeConfigToken resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EdgeConfigToken Get(string name, Input<string> id, EdgeConfigTokenState? state = null, CustomResourceOptions? options = null)
        {
            return new EdgeConfigToken(name, id, state, options);
        }
    }

    public sealed class EdgeConfigTokenArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Edge Config store.
        /// </summary>
        [Input("edgeConfigId", required: true)]
        public Input<string> EdgeConfigId { get; set; } = null!;

        /// <summary>
        /// The label of the Edge Config Token.
        /// </summary>
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        /// <summary>
        /// The ID of the team the Edge Config should exist under. Required when configuring a team resource if a default team has
        /// not been set in the provider.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        public EdgeConfigTokenArgs()
        {
        }
        public static new EdgeConfigTokenArgs Empty => new EdgeConfigTokenArgs();
    }

    public sealed class EdgeConfigTokenState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A connection string is a URL that connects a project to an Edge Config. The variable can be called anything, but our
        /// Edge Config client SDK will search for process.env.EDGE_CONFIG by default.
        /// </summary>
        [Input("connectionString")]
        public Input<string>? ConnectionString { get; set; }

        /// <summary>
        /// The ID of the Edge Config store.
        /// </summary>
        [Input("edgeConfigId")]
        public Input<string>? EdgeConfigId { get; set; }

        /// <summary>
        /// The label of the Edge Config Token.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// The ID of the team the Edge Config should exist under. Required when configuring a team resource if a default team has
        /// not been set in the provider.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        /// <summary>
        /// A read access token used for authenticating against the Edge Config's endpoint for high volume, low-latency requests.
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        public EdgeConfigTokenState()
        {
        }
        public static new EdgeConfigTokenState Empty => new EdgeConfigTokenState();
    }
}
