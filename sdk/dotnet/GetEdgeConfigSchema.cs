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
    public static class GetEdgeConfigSchema
    {
        /// <summary>
        /// An Edge Config Schema provides an existing Edge Config with a JSON schema. Use schema protection to prevent unexpected updates that may cause bugs or downtime.
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
        ///     var test = Vercel.GetEdgeConfigSchema.Invoke(new()
        ///     {
        ///         Id = "ecfg_xxxxxxxxxxxxxxxxxxxxxxxxxxxx",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetEdgeConfigSchemaResult> InvokeAsync(GetEdgeConfigSchemaArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEdgeConfigSchemaResult>("vercel:index/getEdgeConfigSchema:getEdgeConfigSchema", args ?? new GetEdgeConfigSchemaArgs(), options.WithDefaults());

        /// <summary>
        /// An Edge Config Schema provides an existing Edge Config with a JSON schema. Use schema protection to prevent unexpected updates that may cause bugs or downtime.
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
        ///     var test = Vercel.GetEdgeConfigSchema.Invoke(new()
        ///     {
        ///         Id = "ecfg_xxxxxxxxxxxxxxxxxxxxxxxxxxxx",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetEdgeConfigSchemaResult> Invoke(GetEdgeConfigSchemaInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEdgeConfigSchemaResult>("vercel:index/getEdgeConfigSchema:getEdgeConfigSchema", args ?? new GetEdgeConfigSchemaInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEdgeConfigSchemaArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Edge Config that the schema should be for.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// The ID of the team the Edge Config should exist under. Required when configuring a team resource if a default team has not been set in the provider.
        /// </summary>
        [Input("teamId")]
        public string? TeamId { get; set; }

        public GetEdgeConfigSchemaArgs()
        {
        }
        public static new GetEdgeConfigSchemaArgs Empty => new GetEdgeConfigSchemaArgs();
    }

    public sealed class GetEdgeConfigSchemaInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Edge Config that the schema should be for.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The ID of the team the Edge Config should exist under. Required when configuring a team resource if a default team has not been set in the provider.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        public GetEdgeConfigSchemaInvokeArgs()
        {
        }
        public static new GetEdgeConfigSchemaInvokeArgs Empty => new GetEdgeConfigSchemaInvokeArgs();
    }


    [OutputType]
    public sealed class GetEdgeConfigSchemaResult
    {
        /// <summary>
        /// A JSON schema that will be used to validate data in the Edge Config.
        /// </summary>
        public readonly string Definition;
        /// <summary>
        /// The ID of the Edge Config that the schema should be for.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of the team the Edge Config should exist under. Required when configuring a team resource if a default team has not been set in the provider.
        /// </summary>
        public readonly string TeamId;

        [OutputConstructor]
        private GetEdgeConfigSchemaResult(
            string definition,

            string id,

            string teamId)
        {
            Definition = definition;
            Id = id;
            TeamId = teamId;
        }
    }
}