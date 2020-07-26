# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class ListAutomationAccountKeysResult:
    def __init__(__self__, keys=None):
        if keys and not isinstance(keys, list):
            raise TypeError("Expected argument 'keys' to be a list")
        __self__.keys = keys
        """
        Lists the automation keys.
        """


class AwaitableListAutomationAccountKeysResult(ListAutomationAccountKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListAutomationAccountKeysResult(
            keys=self.keys)


def list_automation_account_keys(automation_account_name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str automation_account_name: The name of the automation account.
    :param str resource_group_name: Name of an Azure Resource group.
    """
    __args__ = dict()
    __args__['automationAccountName'] = automation_account_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:automation:listAutomationAccountKeys', __args__, opts=opts).value

    return AwaitableListAutomationAccountKeysResult(
        keys=__ret__.get('keys'))