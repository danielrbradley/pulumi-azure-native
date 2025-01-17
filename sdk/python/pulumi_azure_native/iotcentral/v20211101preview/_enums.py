# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AppSku',
    'NetworkAction',
    'PrivateEndpointServiceConnectionStatus',
    'PublicNetworkAccess',
    'SystemAssignedServiceIdentityType',
]


class AppSku(str, Enum):
    """
    The name of the SKU.
    """
    ST0 = "ST0"
    ST1 = "ST1"
    ST2 = "ST2"


class NetworkAction(str, Enum):
    """
    The default network action to apply.
    """
    ALLOW = "Allow"
    DENY = "Deny"


class PrivateEndpointServiceConnectionStatus(str, Enum):
    """
    Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
    """
    PENDING = "Pending"
    APPROVED = "Approved"
    REJECTED = "Rejected"


class PublicNetworkAccess(str, Enum):
    """
    Whether requests from the public network are allowed.
    """
    ENABLED = "Enabled"
    DISABLED = "Disabled"


class SystemAssignedServiceIdentityType(str, Enum):
    """
    Type of managed service identity (either system assigned, or none).
    """
    NONE = "None"
    SYSTEM_ASSIGNED = "SystemAssigned"
