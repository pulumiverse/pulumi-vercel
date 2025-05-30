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
    /// Provides a Firewall Bypass Rule
    /// 
    /// Firewall Bypass Rules configure sets of domains and ip address to prevent bypass Vercel's system mitigations for.  The hosts used in a bypass rule must be a production domain assigned to the associated project.  Requests that bypass system mitigations will incur usage.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vercel = Pulumiverse.Vercel;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Vercel.Project("example", new()
    ///     {
    ///         Name = "firewall-bypass-example",
    ///     });
    /// 
    ///     var bypassTargeted = new Vercel.FirewallBypass("bypass_targeted", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         SourceIp = "5.6.7.8",
    ///         Domain = "my-production-domain.com",
    ///     });
    /// 
    ///     var bypassCidr = new Vercel.FirewallBypass("bypass_cidr", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         SourceIp = "52.33.44.0/24",
    ///         Domain = "my-production-domain.com",
    ///     });
    /// 
    ///     var bypassAll = new Vercel.FirewallBypass("bypass_all", new()
    ///     {
    ///         ProjectId = example.Id,
    ///         SourceIp = "52.33.44.0/24",
    ///         Domain = "*",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import vercel:index/firewallBypass:FirewallBypass example team_xxxxxxxxxxxxxxxxxxxxxxxx/prj_xxxxxxxxxxxxxxxxxxxxxxxxxxxx#mybypasshost.com#3.4.5.0/24
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import vercel:index/firewallBypass:FirewallBypass example team_xxxxxxxxxxxxxxxxxxxxxxxx/prj_xxxxxxxxxxxxxxxxxxxxxxxxxxxx#3.4.5.0/24
    /// ```
    /// </summary>
    [VercelResourceType("vercel:index/firewallBypass:FirewallBypass")]
    public partial class FirewallBypass : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The domain to configure the bypass rule for.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// The ID of the Project to assign the bypass rule to
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The source IP address to configure the bypass rule for.
        /// </summary>
        [Output("sourceIp")]
        public Output<string> SourceIp { get; private set; } = null!;

        /// <summary>
        /// The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
        /// </summary>
        [Output("teamId")]
        public Output<string> TeamId { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallBypass resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallBypass(string name, FirewallBypassArgs args, CustomResourceOptions? options = null)
            : base("vercel:index/firewallBypass:FirewallBypass", name, args ?? new FirewallBypassArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallBypass(string name, Input<string> id, FirewallBypassState? state = null, CustomResourceOptions? options = null)
            : base("vercel:index/firewallBypass:FirewallBypass", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FirewallBypass resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallBypass Get(string name, Input<string> id, FirewallBypassState? state = null, CustomResourceOptions? options = null)
        {
            return new FirewallBypass(name, id, state, options);
        }
    }

    public sealed class FirewallBypassArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain to configure the bypass rule for.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// The ID of the Project to assign the bypass rule to
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The source IP address to configure the bypass rule for.
        /// </summary>
        [Input("sourceIp", required: true)]
        public Input<string> SourceIp { get; set; } = null!;

        /// <summary>
        /// The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        public FirewallBypassArgs()
        {
        }
        public static new FirewallBypassArgs Empty => new FirewallBypassArgs();
    }

    public sealed class FirewallBypassState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain to configure the bypass rule for.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The ID of the Project to assign the bypass rule to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The source IP address to configure the bypass rule for.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        public FirewallBypassState()
        {
        }
        public static new FirewallBypassState Empty => new FirewallBypassState();
    }
}
