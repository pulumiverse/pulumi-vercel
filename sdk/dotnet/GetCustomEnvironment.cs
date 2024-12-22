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
    public static class GetCustomEnvironment
    {
        public static Task<GetCustomEnvironmentResult> InvokeAsync(GetCustomEnvironmentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCustomEnvironmentResult>("vercel:index/getCustomEnvironment:getCustomEnvironment", args ?? new GetCustomEnvironmentArgs(), options.WithDefaults());

        public static Output<GetCustomEnvironmentResult> Invoke(GetCustomEnvironmentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCustomEnvironmentResult>("vercel:index/getCustomEnvironment:getCustomEnvironment", args ?? new GetCustomEnvironmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCustomEnvironmentArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("projectId", required: true)]
        public string ProjectId { get; set; } = null!;

        [Input("teamId")]
        public string? TeamId { get; set; }

        public GetCustomEnvironmentArgs()
        {
        }
        public static new GetCustomEnvironmentArgs Empty => new GetCustomEnvironmentArgs();
    }

    public sealed class GetCustomEnvironmentInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        public GetCustomEnvironmentInvokeArgs()
        {
        }
        public static new GetCustomEnvironmentInvokeArgs Empty => new GetCustomEnvironmentInvokeArgs();
    }


    [OutputType]
    public sealed class GetCustomEnvironmentResult
    {
        public readonly Outputs.GetCustomEnvironmentBranchTrackingResult BranchTracking;
        public readonly string Description;
        public readonly string Id;
        public readonly string Name;
        public readonly string ProjectId;
        public readonly string TeamId;

        [OutputConstructor]
        private GetCustomEnvironmentResult(
            Outputs.GetCustomEnvironmentBranchTrackingResult branchTracking,

            string description,

            string id,

            string name,

            string projectId,

            string teamId)
        {
            BranchTracking = branchTracking;
            Description = description;
            Id = id;
            Name = name;
            ProjectId = projectId;
            TeamId = teamId;
        }
    }
}
