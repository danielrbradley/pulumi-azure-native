# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class ListBotConnectionServiceProvidersResult:
    """
    The list of bot service providers response.
    """
    def __init__(__self__, next_link=None, value=None):
        if next_link and not isinstance(next_link, str):
            raise TypeError("Expected argument 'next_link' to be a str")
        __self__.next_link = next_link
        """
        The link used to get the next page of bot service providers.
        """
        if value and not isinstance(value, list):
            raise TypeError("Expected argument 'value' to be a list")
        __self__.value = value
        """
        Gets the list of bot service providers and their properties.
        """


class AwaitableListBotConnectionServiceProvidersResult(ListBotConnectionServiceProvidersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListBotConnectionServiceProvidersResult(
            next_link=self.next_link,
            value=self.value)


def list_bot_connection_service_providers(opts=None):
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:botservice/v20200602:listBotConnectionServiceProviders', __args__, opts=opts).value

    return AwaitableListBotConnectionServiceProvidersResult(
        next_link=__ret__.get('nextLink'),
        value=__ret__.get('value'))