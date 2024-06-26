// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumiverse.Vercel
{
    public static class Config
    {
        [global::System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("vercel");

        private static readonly __Value<string?> _apiToken = new __Value<string?>(() => __config.Get("apiToken"));
        /// <summary>
        /// The Vercel API Token to use. This can also be specified with the `VERCEL_API_TOKEN` shell environment variable. Tokens
        /// can be created from your [Vercel settings](https://vercel.com/account/tokens).
        /// </summary>
        public static string? ApiToken
        {
            get => _apiToken.Get();
            set => _apiToken.Set(value);
        }

        private static readonly __Value<string?> _team = new __Value<string?>(() => __config.Get("team"));
        /// <summary>
        /// The default Vercel Team to use when creating resources or reading data sources. This can be provided as either a team
        /// slug, or team ID. The slug and ID are both available from the Team Settings page in the Vercel dashboard.
        /// </summary>
        public static string? Team
        {
            get => _team.Get();
            set => _team.Set(value);
        }

    }
}
