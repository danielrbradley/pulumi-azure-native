# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult:
    """
    The list of invoice section properties with create subscription permission.
    """
    def __init__(__self__, next_link=None, value=None):
        if next_link and not isinstance(next_link, str):
            raise TypeError("Expected argument 'next_link' to be a str")
        __self__.next_link = next_link
        """
        The link (url) to the next page of results.
        """
        if value and not isinstance(value, list):
            raise TypeError("Expected argument 'value' to be a list")
        __self__.value = value
        """
        The list of invoice section properties with create subscription permission.
        """


class AwaitableListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult(ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult(
            next_link=self.next_link,
            value=self.value)


def list_billing_account_invoice_sections_by_create_subscription_permission(billing_account_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str billing_account_name: The ID that uniquely identifies a billing account.
    """
    __args__ = dict()
    __args__['billingAccountName'] = billing_account_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:billing/v20200501:listBillingAccountInvoiceSectionsByCreateSubscriptionPermission', __args__, opts=opts).value

    return AwaitableListBillingAccountInvoiceSectionsByCreateSubscriptionPermissionResult(
        next_link=__ret__.get('nextLink'),
        value=__ret__.get('value'))