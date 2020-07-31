# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GetExportConfigurationResult:
    """
    Properties that define a Continuous Export configuration.
    """
    def __init__(__self__, application_name=None, container_name=None, destination_account_id=None, destination_storage_location_id=None, destination_storage_subscription_id=None, destination_type=None, export_id=None, export_status=None, instrumentation_key=None, is_user_enabled=None, last_gap_time=None, last_success_time=None, last_user_update=None, notification_queue_enabled=None, permanent_error_reason=None, record_types=None, resource_group=None, storage_name=None, subscription_id=None):
        if application_name and not isinstance(application_name, str):
            raise TypeError("Expected argument 'application_name' to be a str")
        __self__.application_name = application_name
        """
        The name of the Application Insights component.
        """
        if container_name and not isinstance(container_name, str):
            raise TypeError("Expected argument 'container_name' to be a str")
        __self__.container_name = container_name
        """
        The name of the destination storage container.
        """
        if destination_account_id and not isinstance(destination_account_id, str):
            raise TypeError("Expected argument 'destination_account_id' to be a str")
        __self__.destination_account_id = destination_account_id
        """
        The name of destination account.
        """
        if destination_storage_location_id and not isinstance(destination_storage_location_id, str):
            raise TypeError("Expected argument 'destination_storage_location_id' to be a str")
        __self__.destination_storage_location_id = destination_storage_location_id
        """
        The destination account location ID.
        """
        if destination_storage_subscription_id and not isinstance(destination_storage_subscription_id, str):
            raise TypeError("Expected argument 'destination_storage_subscription_id' to be a str")
        __self__.destination_storage_subscription_id = destination_storage_subscription_id
        """
        The destination storage account subscription ID.
        """
        if destination_type and not isinstance(destination_type, str):
            raise TypeError("Expected argument 'destination_type' to be a str")
        __self__.destination_type = destination_type
        """
        The destination type.
        """
        if export_id and not isinstance(export_id, str):
            raise TypeError("Expected argument 'export_id' to be a str")
        __self__.export_id = export_id
        """
        The unique ID of the export configuration inside an Application Insights component. It is auto generated when the Continuous Export configuration is created.
        """
        if export_status and not isinstance(export_status, str):
            raise TypeError("Expected argument 'export_status' to be a str")
        __self__.export_status = export_status
        """
        This indicates current Continuous Export configuration status. The possible values are 'Preparing', 'Success', 'Failure'.
        """
        if instrumentation_key and not isinstance(instrumentation_key, str):
            raise TypeError("Expected argument 'instrumentation_key' to be a str")
        __self__.instrumentation_key = instrumentation_key
        """
        The instrumentation key of the Application Insights component.
        """
        if is_user_enabled and not isinstance(is_user_enabled, str):
            raise TypeError("Expected argument 'is_user_enabled' to be a str")
        __self__.is_user_enabled = is_user_enabled
        """
        This will be 'true' if the Continuous Export configuration is enabled, otherwise it will be 'false'.
        """
        if last_gap_time and not isinstance(last_gap_time, str):
            raise TypeError("Expected argument 'last_gap_time' to be a str")
        __self__.last_gap_time = last_gap_time
        """
        The last time the Continuous Export configuration started failing.
        """
        if last_success_time and not isinstance(last_success_time, str):
            raise TypeError("Expected argument 'last_success_time' to be a str")
        __self__.last_success_time = last_success_time
        """
        The last time data was successfully delivered to the destination storage container for this Continuous Export configuration.
        """
        if last_user_update and not isinstance(last_user_update, str):
            raise TypeError("Expected argument 'last_user_update' to be a str")
        __self__.last_user_update = last_user_update
        """
        Last time the Continuous Export configuration was updated.
        """
        if notification_queue_enabled and not isinstance(notification_queue_enabled, str):
            raise TypeError("Expected argument 'notification_queue_enabled' to be a str")
        __self__.notification_queue_enabled = notification_queue_enabled
        """
        Deprecated
        """
        if permanent_error_reason and not isinstance(permanent_error_reason, str):
            raise TypeError("Expected argument 'permanent_error_reason' to be a str")
        __self__.permanent_error_reason = permanent_error_reason
        """
        This is the reason the Continuous Export configuration started failing. It can be 'AzureStorageNotFound' or 'AzureStorageAccessDenied'.
        """
        if record_types and not isinstance(record_types, str):
            raise TypeError("Expected argument 'record_types' to be a str")
        __self__.record_types = record_types
        """
        This comma separated list of document types that will be exported. The possible values include 'Requests', 'Event', 'Exceptions', 'Metrics', 'PageViews', 'PageViewPerformance', 'Rdd', 'PerformanceCounters', 'Availability', 'Messages'.
        """
        if resource_group and not isinstance(resource_group, str):
            raise TypeError("Expected argument 'resource_group' to be a str")
        __self__.resource_group = resource_group
        """
        The resource group of the Application Insights component.
        """
        if storage_name and not isinstance(storage_name, str):
            raise TypeError("Expected argument 'storage_name' to be a str")
        __self__.storage_name = storage_name
        """
        The name of the destination storage account.
        """
        if subscription_id and not isinstance(subscription_id, str):
            raise TypeError("Expected argument 'subscription_id' to be a str")
        __self__.subscription_id = subscription_id
        """
        The subscription of the Application Insights component.
        """


class AwaitableGetExportConfigurationResult(GetExportConfigurationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetExportConfigurationResult(
            application_name=self.application_name,
            container_name=self.container_name,
            destination_account_id=self.destination_account_id,
            destination_storage_location_id=self.destination_storage_location_id,
            destination_storage_subscription_id=self.destination_storage_subscription_id,
            destination_type=self.destination_type,
            export_id=self.export_id,
            export_status=self.export_status,
            instrumentation_key=self.instrumentation_key,
            is_user_enabled=self.is_user_enabled,
            last_gap_time=self.last_gap_time,
            last_success_time=self.last_success_time,
            last_user_update=self.last_user_update,
            notification_queue_enabled=self.notification_queue_enabled,
            permanent_error_reason=self.permanent_error_reason,
            record_types=self.record_types,
            resource_group=self.resource_group,
            storage_name=self.storage_name,
            subscription_id=self.subscription_id)


def get_export_configuration(export_id=None, resource_group_name=None, resource_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str export_id: The Continuous Export configuration ID. This is unique within a Application Insights component.
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    :param str resource_name: The name of the Application Insights component resource.
    """
    __args__ = dict()
    __args__['exportId'] = export_id
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceName'] = resource_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:insights/v20150501:getExportConfiguration', __args__, opts=opts).value

    return AwaitableGetExportConfigurationResult(
        application_name=__ret__.get('ApplicationName'),
        container_name=__ret__.get('ContainerName'),
        destination_account_id=__ret__.get('DestinationAccountId'),
        destination_storage_location_id=__ret__.get('DestinationStorageLocationId'),
        destination_storage_subscription_id=__ret__.get('DestinationStorageSubscriptionId'),
        destination_type=__ret__.get('DestinationType'),
        export_id=__ret__.get('ExportId'),
        export_status=__ret__.get('ExportStatus'),
        instrumentation_key=__ret__.get('InstrumentationKey'),
        is_user_enabled=__ret__.get('IsUserEnabled'),
        last_gap_time=__ret__.get('LastGapTime'),
        last_success_time=__ret__.get('LastSuccessTime'),
        last_user_update=__ret__.get('LastUserUpdate'),
        notification_queue_enabled=__ret__.get('NotificationQueueEnabled'),
        permanent_error_reason=__ret__.get('PermanentErrorReason'),
        record_types=__ret__.get('RecordTypes'),
        resource_group=__ret__.get('ResourceGroup'),
        storage_name=__ret__.get('StorageName'),
        subscription_id=__ret__.get('SubscriptionId'))