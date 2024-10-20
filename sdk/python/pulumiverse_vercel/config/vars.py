# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

import types

__config__ = pulumi.Config('vercel')


class _ExportableConfig(types.ModuleType):
    @property
    def api_token(self) -> Optional[str]:
        """
        The Vercel API Token to use. This can also be specified with the `VERCEL_API_TOKEN` shell environment variable. Tokens
        can be created from your [Vercel settings](https://vercel.com/account/tokens).
        """
        return __config__.get('apiToken')

    @property
    def team(self) -> Optional[str]:
        """
        The default Vercel Team to use when creating resources or reading data sources. This can be provided as either a team
        slug, or team ID. The slug and ID are both available from the Team Settings page in the Vercel dashboard.
        """
        return __config__.get('team')

