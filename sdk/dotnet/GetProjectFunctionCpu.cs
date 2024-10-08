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
    public static class GetProjectFunctionCpu
    {
        /// <summary>
        /// &gt; This data source has been deprecated and no longer works. Please use the `vercel.Project` data source and its `resource_config` attribute instead.
        /// 
        /// Provides information about a Project's Function CPU setting.
        /// 
        /// This controls the maximum amount of CPU utilization your Serverless Functions can use while executing. Standard is optimal for most frontend workloads. You can override this per function using the vercel.json file.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vercel = Pulumi.Vercel;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var exampleProject = Vercel.GetProject.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        ///     var exampleProjectFunctionCpu = Vercel.GetProjectFunctionCpu.Invoke(new()
        ///     {
        ///         ProjectId = exampleProject.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetProjectFunctionCpuResult> InvokeAsync(GetProjectFunctionCpuArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectFunctionCpuResult>("vercel:index/getProjectFunctionCpu:getProjectFunctionCpu", args ?? new GetProjectFunctionCpuArgs(), options.WithDefaults());

        /// <summary>
        /// &gt; This data source has been deprecated and no longer works. Please use the `vercel.Project` data source and its `resource_config` attribute instead.
        /// 
        /// Provides information about a Project's Function CPU setting.
        /// 
        /// This controls the maximum amount of CPU utilization your Serverless Functions can use while executing. Standard is optimal for most frontend workloads. You can override this per function using the vercel.json file.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vercel = Pulumi.Vercel;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var exampleProject = Vercel.GetProject.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        ///     var exampleProjectFunctionCpu = Vercel.GetProjectFunctionCpu.Invoke(new()
        ///     {
        ///         ProjectId = exampleProject.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetProjectFunctionCpuResult> Invoke(GetProjectFunctionCpuInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectFunctionCpuResult>("vercel:index/getProjectFunctionCpu:getProjectFunctionCpu", args ?? new GetProjectFunctionCpuInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectFunctionCpuArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Project to read the Function CPU setting for.
        /// </summary>
        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        /// <summary>
        /// The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
        /// </summary>
        [Input("teamId")]
        public string? TeamId { get; set; }

        public GetProjectFunctionCpuArgs()
        {
        }
        public static new GetProjectFunctionCpuArgs Empty => new GetProjectFunctionCpuArgs();
    }

    public sealed class GetProjectFunctionCpuInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Project to read the Function CPU setting for.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        public GetProjectFunctionCpuInvokeArgs()
        {
        }
        public static new GetProjectFunctionCpuInvokeArgs Empty => new GetProjectFunctionCpuInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectFunctionCpuResult
    {
        /// <summary>
        /// The amount of CPU available to your Serverless Functions. Should be one of 'basic' (0.6vCPU), 'standard' (1vCPU) or 'performance' (1.7vCPUs).
        /// </summary>
        public readonly string Cpu;
        /// <summary>
        /// The ID of the resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of the Project to read the Function CPU setting for.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// The ID of the team the Project exists under. Required when configuring a team resource if a default team has not been set in the provider.
        /// </summary>
        public readonly string TeamId;

        [OutputConstructor]
        private GetProjectFunctionCpuResult(
            string cpu,

            string id,

            string projectId,

            string teamId)
        {
            Cpu = cpu;
            Id = id;
            ProjectId = projectId;
            TeamId = teamId;
        }
    }
}
