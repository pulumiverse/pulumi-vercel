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
    public static class GetProject
    {
        public static Task<GetProjectResult> InvokeAsync(GetProjectArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectResult>("vercel:index/getProject:getProject", args ?? new GetProjectArgs(), options.WithDefaults());

        public static Output<GetProjectResult> Invoke(GetProjectInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectResult>("vercel:index/getProject:getProject", args ?? new GetProjectInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("teamId")]
        public string? TeamId { get; set; }

        public GetProjectArgs()
        {
        }
        public static new GetProjectArgs Empty => new GetProjectArgs();
    }

    public sealed class GetProjectInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        public GetProjectInvokeArgs()
        {
        }
        public static new GetProjectInvokeArgs Empty => new GetProjectInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectResult
    {
        public readonly bool AutoAssignCustomDomains;
        public readonly bool AutomaticallyExposeSystemEnvironmentVariables;
        public readonly string BuildCommand;
        public readonly bool CustomerSuccessCodeVisibility;
        public readonly string DevCommand;
        public readonly bool DirectoryListing;
        public readonly ImmutableArray<Outputs.GetProjectEnvironmentResult> Environments;
        public readonly string Framework;
        public readonly bool FunctionFailover;
        public readonly Outputs.GetProjectGitCommentsResult GitComments;
        public readonly bool GitForkProtection;
        public readonly bool GitLfs;
        public readonly Outputs.GetProjectGitRepositoryResult GitRepository;
        public readonly string Id;
        public readonly string IgnoreCommand;
        public readonly string InstallCommand;
        public readonly string Name;
        public readonly Outputs.GetProjectOidcTokenConfigResult OidcTokenConfig;
        public readonly Outputs.GetProjectOptionsAllowlistResult OptionsAllowlist;
        public readonly string OutputDirectory;
        public readonly Outputs.GetProjectPasswordProtectionResult PasswordProtection;
        public readonly bool PreviewComments;
        public readonly bool PrioritiseProductionBuilds;
        public readonly bool ProtectionBypassForAutomation;
        public readonly string ProtectionBypassForAutomationSecret;
        public readonly bool PublicSource;
        public readonly Outputs.GetProjectResourceConfigResult ResourceConfig;
        public readonly string RootDirectory;
        public readonly string ServerlessFunctionRegion;
        public readonly string SkewProtection;
        public readonly string TeamId;
        public readonly Outputs.GetProjectTrustedIpsResult TrustedIps;
        public readonly Outputs.GetProjectVercelAuthenticationResult VercelAuthentication;

        [OutputConstructor]
        private GetProjectResult(
            bool autoAssignCustomDomains,

            bool automaticallyExposeSystemEnvironmentVariables,

            string buildCommand,

            bool customerSuccessCodeVisibility,

            string devCommand,

            bool directoryListing,

            ImmutableArray<Outputs.GetProjectEnvironmentResult> environments,

            string framework,

            bool functionFailover,

            Outputs.GetProjectGitCommentsResult gitComments,

            bool gitForkProtection,

            bool gitLfs,

            Outputs.GetProjectGitRepositoryResult gitRepository,

            string id,

            string ignoreCommand,

            string installCommand,

            string name,

            Outputs.GetProjectOidcTokenConfigResult oidcTokenConfig,

            Outputs.GetProjectOptionsAllowlistResult optionsAllowlist,

            string outputDirectory,

            Outputs.GetProjectPasswordProtectionResult passwordProtection,

            bool previewComments,

            bool prioritiseProductionBuilds,

            bool protectionBypassForAutomation,

            string protectionBypassForAutomationSecret,

            bool publicSource,

            Outputs.GetProjectResourceConfigResult resourceConfig,

            string rootDirectory,

            string serverlessFunctionRegion,

            string skewProtection,

            string teamId,

            Outputs.GetProjectTrustedIpsResult trustedIps,

            Outputs.GetProjectVercelAuthenticationResult vercelAuthentication)
        {
            AutoAssignCustomDomains = autoAssignCustomDomains;
            AutomaticallyExposeSystemEnvironmentVariables = automaticallyExposeSystemEnvironmentVariables;
            BuildCommand = buildCommand;
            CustomerSuccessCodeVisibility = customerSuccessCodeVisibility;
            DevCommand = devCommand;
            DirectoryListing = directoryListing;
            Environments = environments;
            Framework = framework;
            FunctionFailover = functionFailover;
            GitComments = gitComments;
            GitForkProtection = gitForkProtection;
            GitLfs = gitLfs;
            GitRepository = gitRepository;
            Id = id;
            IgnoreCommand = ignoreCommand;
            InstallCommand = installCommand;
            Name = name;
            OidcTokenConfig = oidcTokenConfig;
            OptionsAllowlist = optionsAllowlist;
            OutputDirectory = outputDirectory;
            PasswordProtection = passwordProtection;
            PreviewComments = previewComments;
            PrioritiseProductionBuilds = prioritiseProductionBuilds;
            ProtectionBypassForAutomation = protectionBypassForAutomation;
            ProtectionBypassForAutomationSecret = protectionBypassForAutomationSecret;
            PublicSource = publicSource;
            ResourceConfig = resourceConfig;
            RootDirectory = rootDirectory;
            ServerlessFunctionRegion = serverlessFunctionRegion;
            SkewProtection = skewProtection;
            TeamId = teamId;
            TrustedIps = trustedIps;
            VercelAuthentication = vercelAuthentication;
        }
    }
}
