# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class ListSearchServiceAdminKeysResult:
    """
    Response containing the primary and secondary admin API keys for a given Azure Cognitive Search service.
    """
    def __init__(__self__, primary_key=None, secondary_key=None):
        if primary_key and not isinstance(primary_key, str):
            raise TypeError("Expected argument 'primary_key' to be a str")
        __self__.primary_key = primary_key
        """
        The primary admin API key of the Search service.
        """
        if secondary_key and not isinstance(secondary_key, str):
            raise TypeError("Expected argument 'secondary_key' to be a str")
        __self__.secondary_key = secondary_key
        """
        The secondary admin API key of the Search service.
        """


class AwaitableListSearchServiceAdminKeysResult(ListSearchServiceAdminKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListSearchServiceAdminKeysResult(
            primary_key=self.primary_key,
            secondary_key=self.secondary_key)


def list_search_service_admin_keys(resource_group_name=None, search_service_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str resource_group_name: The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
    :param str search_service_name: The name of the Azure Cognitive Search service associated with the specified resource group.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['searchServiceName'] = search_service_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:search:listSearchServiceAdminKeys', __args__, opts=opts).value

    return AwaitableListSearchServiceAdminKeysResult(
        primary_key=__ret__.get('primaryKey'),
        secondary_key=__ret__.get('secondaryKey'))