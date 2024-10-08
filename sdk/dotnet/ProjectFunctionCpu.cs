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
    /// &gt; This resource has been deprecated and no longer works. Please use the `vercel.Project` resource and its `resource_config` attribute instead.
    /// 
    /// Provides a Function CPU resource for a Project.
    /// 
    /// This controls the maximum amount of CPU utilization your Serverless Functions can use while executing. Standard is optimal for most frontend workloads. You can override this per function using the vercel.json file.
    /// 
    /// A new Deployment is required for your changes to take effect.
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
    ///     var exampleProject = new Vercel.Project("exampleProject");
    /// 
    ///     var exampleProjectFunctionCpu = new Vercel.ProjectFunctionCpu("exampleProjectFunctionCpu", new()
    ///     {
    ///         ProjectId = exampleProject.Id,
    ///         Cpu = "performance",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// You can import via the team_id and project_id.
    /// 
    /// - team_id can be found in the team `settings` tab in the Vercel UI.
    /// 
    /// - project_id can be found in the project `settings` tab in the Vercel UI.
    /// 
    /// ```sh
    /// $ pulumi import vercel:index/projectFunctionCpu:ProjectFunctionCpu example team_xxxxxxxxxxxxxxxxxxxxxxxx/prj_xxxxxxxxxxxxxxxxxxxxxxxxxxxx
    /// ```
    /// </summary>
    [VercelResourceType("vercel:index/projectFunctionCpu:ProjectFunctionCpu")]
    public partial class ProjectFunctionCpu : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The amount of CPU available to your Serverless Functions. Should be one of 'basic' (0.6vCPU), 'standard' (1vCPU) or 'performance' (1.7vCPUs).
        /// </summary>
        [Output("cpu")]
        public Output<string> Cpu { get; private set; } = null!;

        /// <summary>
        /// The ID of the Project to adjust the CPU for.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
        /// </summary>
        [Output("teamId")]
        public Output<string> TeamId { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectFunctionCpu resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectFunctionCpu(string name, ProjectFunctionCpuArgs args, CustomResourceOptions? options = null)
            : base("vercel:index/projectFunctionCpu:ProjectFunctionCpu", name, args ?? new ProjectFunctionCpuArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectFunctionCpu(string name, Input<string> id, ProjectFunctionCpuState? state = null, CustomResourceOptions? options = null)
            : base("vercel:index/projectFunctionCpu:ProjectFunctionCpu", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectFunctionCpu resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectFunctionCpu Get(string name, Input<string> id, ProjectFunctionCpuState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectFunctionCpu(name, id, state, options);
        }
    }

    public sealed class ProjectFunctionCpuArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The amount of CPU available to your Serverless Functions. Should be one of 'basic' (0.6vCPU), 'standard' (1vCPU) or 'performance' (1.7vCPUs).
        /// </summary>
        [Input("cpu", required: true)]
        public Input<string> Cpu { get; set; } = null!;

        /// <summary>
        /// The ID of the Project to adjust the CPU for.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        public ProjectFunctionCpuArgs()
        {
        }
        public static new ProjectFunctionCpuArgs Empty => new ProjectFunctionCpuArgs();
    }

    public sealed class ProjectFunctionCpuState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The amount of CPU available to your Serverless Functions. Should be one of 'basic' (0.6vCPU), 'standard' (1vCPU) or 'performance' (1.7vCPUs).
        /// </summary>
        [Input("cpu")]
        public Input<string>? Cpu { get; set; }

        /// <summary>
        /// The ID of the Project to adjust the CPU for.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        public ProjectFunctionCpuState()
        {
        }
        public static new ProjectFunctionCpuState Empty => new ProjectFunctionCpuState();
    }
}
