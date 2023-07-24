// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vercel
{
    public static class GetFile
    {
        /// <summary>
        /// Provides information about a file on disk.
        /// 
        /// This will read a single file, providing metadata for use with a `vercel.Deployment`.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vercel = Pulumi.Vercel;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var exampleFile = Vercel.GetFile.Invoke(new()
        ///     {
        ///         Path = "index.html",
        ///     });
        /// 
        ///     var exampleProject = Vercel.GetProject.Invoke(new()
        ///     {
        ///         Name = "my-project",
        ///     });
        /// 
        ///     var exampleDeployment = new Vercel.Deployment("exampleDeployment", new()
        ///     {
        ///         ProjectId = exampleProject.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         Files = exampleFile.Apply(getFileResult =&gt; getFileResult.File),
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFileResult> InvokeAsync(GetFileArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFileResult>("vercel:index/getFile:getFile", args ?? new GetFileArgs(), options.WithDefaults());

        /// <summary>
        /// Provides information about a file on disk.
        /// 
        /// This will read a single file, providing metadata for use with a `vercel.Deployment`.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vercel = Pulumi.Vercel;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var exampleFile = Vercel.GetFile.Invoke(new()
        ///     {
        ///         Path = "index.html",
        ///     });
        /// 
        ///     var exampleProject = Vercel.GetProject.Invoke(new()
        ///     {
        ///         Name = "my-project",
        ///     });
        /// 
        ///     var exampleDeployment = new Vercel.Deployment("exampleDeployment", new()
        ///     {
        ///         ProjectId = exampleProject.Apply(getProjectResult =&gt; getProjectResult.Id),
        ///         Files = exampleFile.Apply(getFileResult =&gt; getFileResult.File),
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetFileResult> Invoke(GetFileInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFileResult>("vercel:index/getFile:getFile", args ?? new GetFileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFileArgs : global::Pulumi.InvokeArgs
    {
        [Input("path", required: true)]
        public string Path { get; set; } = null!;

        public GetFileArgs()
        {
        }
        public static new GetFileArgs Empty => new GetFileArgs();
    }

    public sealed class GetFileInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        public GetFileInvokeArgs()
        {
        }
        public static new GetFileInvokeArgs Empty => new GetFileInvokeArgs();
    }


    [OutputType]
    public sealed class GetFileResult
    {
        /// <summary>
        /// A map of filename to metadata about the file. The metadata contains the file size and hash, and allows a deployment to be created if the file changes.
        /// </summary>
        public readonly ImmutableDictionary<string, string> File;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string Id;
        public readonly string Path;

        [OutputConstructor]
        private GetFileResult(
            ImmutableDictionary<string, string> file,

            string id,

            string path)
        {
            File = file;
            Id = id;
            Path = path;
        }
    }
}
