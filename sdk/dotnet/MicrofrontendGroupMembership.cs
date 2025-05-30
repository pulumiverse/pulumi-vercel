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
    /// <summary>
    /// Provides a Microfrontend Group Membership resource.
    /// 
    /// A Microfrontend Group Membership is a definition of a Vercel Project being a part of a Microfrontend Group.
    /// 
    /// ## Import
    /// 
    /// If importing into a personal account, or with a team configured on the provider, simply use the record id.
    /// 
    /// - the microfrontend ID can be taken from the microfrontend settings page
    /// 
    /// - the project ID can be taken from the project settings page
    /// 
    /// ```sh
    /// $ pulumi import vercel:index/microfrontendGroupMembership:MicrofrontendGroupMembership example mfe_xxxxxxxxxxxxxxxxxxxxxxxxxxxx/pid_xxxxxxxxxxxxxxxxxxxxxxxxxxxx
    /// ```
    /// 
    /// Alternatively, you can import via the team_id and microfrontend_id.
    /// 
    /// - team_id can be found in the team `settings` tab in the Vercel UI.
    /// 
    /// - the microfrontend ID can be taken from the microfrontend settings page
    /// 
    /// - the project ID can be taken from the project settings page
    /// 
    /// ```sh
    /// $ pulumi import vercel:index/microfrontendGroupMembership:MicrofrontendGroupMembership example team_xxxxxxxxxxxxxxxxxxxxxxxx/mfe_xxxxxxxxxxxxxxxxxxxxxxxxxxxx/pid_xxxxxxxxxxxxxxxxxxxxxxxxxxxx
    /// ```
    /// </summary>
    [VercelResourceType("vercel:index/microfrontendGroupMembership:MicrofrontendGroupMembership")]
    public partial class MicrofrontendGroupMembership : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The default route for the project. Used for the screenshot of deployments.
        /// </summary>
        [Output("defaultRoute")]
        public Output<string> DefaultRoute { get; private set; } = null!;

        /// <summary>
        /// The ID of the microfrontend group.
        /// </summary>
        [Output("microfrontendGroupId")]
        public Output<string> MicrofrontendGroupId { get; private set; } = null!;

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Whether the project is route observability for this project. If dalse, the project will be route observability for all projects to the default project.
        /// </summary>
        [Output("routeObservabilityToThisProject")]
        public Output<bool> RouteObservabilityToThisProject { get; private set; } = null!;

        /// <summary>
        /// The team ID to add the microfrontend group to. Required when configuring a team resource if a default team has not been set in the provider.
        /// </summary>
        [Output("teamId")]
        public Output<string> TeamId { get; private set; } = null!;


        /// <summary>
        /// Create a MicrofrontendGroupMembership resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MicrofrontendGroupMembership(string name, MicrofrontendGroupMembershipArgs args, CustomResourceOptions? options = null)
            : base("vercel:index/microfrontendGroupMembership:MicrofrontendGroupMembership", name, args ?? new MicrofrontendGroupMembershipArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MicrofrontendGroupMembership(string name, Input<string> id, MicrofrontendGroupMembershipState? state = null, CustomResourceOptions? options = null)
            : base("vercel:index/microfrontendGroupMembership:MicrofrontendGroupMembership", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MicrofrontendGroupMembership resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MicrofrontendGroupMembership Get(string name, Input<string> id, MicrofrontendGroupMembershipState? state = null, CustomResourceOptions? options = null)
        {
            return new MicrofrontendGroupMembership(name, id, state, options);
        }
    }

    public sealed class MicrofrontendGroupMembershipArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default route for the project. Used for the screenshot of deployments.
        /// </summary>
        [Input("defaultRoute")]
        public Input<string>? DefaultRoute { get; set; }

        /// <summary>
        /// The ID of the microfrontend group.
        /// </summary>
        [Input("microfrontendGroupId", required: true)]
        public Input<string> MicrofrontendGroupId { get; set; } = null!;

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Whether the project is route observability for this project. If dalse, the project will be route observability for all projects to the default project.
        /// </summary>
        [Input("routeObservabilityToThisProject")]
        public Input<bool>? RouteObservabilityToThisProject { get; set; }

        /// <summary>
        /// The team ID to add the microfrontend group to. Required when configuring a team resource if a default team has not been set in the provider.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        public MicrofrontendGroupMembershipArgs()
        {
        }
        public static new MicrofrontendGroupMembershipArgs Empty => new MicrofrontendGroupMembershipArgs();
    }

    public sealed class MicrofrontendGroupMembershipState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default route for the project. Used for the screenshot of deployments.
        /// </summary>
        [Input("defaultRoute")]
        public Input<string>? DefaultRoute { get; set; }

        /// <summary>
        /// The ID of the microfrontend group.
        /// </summary>
        [Input("microfrontendGroupId")]
        public Input<string>? MicrofrontendGroupId { get; set; }

        /// <summary>
        /// The ID of the project.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Whether the project is route observability for this project. If dalse, the project will be route observability for all projects to the default project.
        /// </summary>
        [Input("routeObservabilityToThisProject")]
        public Input<bool>? RouteObservabilityToThisProject { get; set; }

        /// <summary>
        /// The team ID to add the microfrontend group to. Required when configuring a team resource if a default team has not been set in the provider.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        public MicrofrontendGroupMembershipState()
        {
        }
        public static new MicrofrontendGroupMembershipState Empty => new MicrofrontendGroupMembershipState();
    }
}
