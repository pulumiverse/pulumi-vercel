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
    public static class GetProjectDirectory
    {
        public static Task<GetProjectDirectoryResult> InvokeAsync(GetProjectDirectoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectDirectoryResult>("vercel:index/getProjectDirectory:getProjectDirectory", args ?? new GetProjectDirectoryArgs(), options.WithDefaults());

        public static Output<GetProjectDirectoryResult> Invoke(GetProjectDirectoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectDirectoryResult>("vercel:index/getProjectDirectory:getProjectDirectory", args ?? new GetProjectDirectoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectDirectoryArgs : global::Pulumi.InvokeArgs
    {
        [Input("path", required: true)]
        public string Path { get; set; } = null!;

        public GetProjectDirectoryArgs()
        {
        }
        public static new GetProjectDirectoryArgs Empty => new GetProjectDirectoryArgs();
    }

    public sealed class GetProjectDirectoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        public GetProjectDirectoryInvokeArgs()
        {
        }
        public static new GetProjectDirectoryInvokeArgs Empty => new GetProjectDirectoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectDirectoryResult
    {
        public readonly ImmutableDictionary<string, string> Files;
        public readonly string Id;
        public readonly string Path;

        [OutputConstructor]
        private GetProjectDirectoryResult(
            ImmutableDictionary<string, string> files,

            string id,

            string path)
        {
            Files = files;
            Id = id;
            Path = path;
        }
    }
}
