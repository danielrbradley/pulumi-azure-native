// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.Authorization.V20170401
{
    /// <summary>
    /// The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
    /// </summary>
    [EnumType]
    public readonly struct LockLevel : IEquatable<LockLevel>
    {
        private readonly string _value;

        private LockLevel(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static LockLevel NotSpecified { get; } = new LockLevel("NotSpecified");
        public static LockLevel CanNotDelete { get; } = new LockLevel("CanNotDelete");
        public static LockLevel ReadOnly { get; } = new LockLevel("ReadOnly");

        public static bool operator ==(LockLevel left, LockLevel right) => left.Equals(right);
        public static bool operator !=(LockLevel left, LockLevel right) => !left.Equals(right);

        public static explicit operator string(LockLevel value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is LockLevel other && Equals(other);
        public bool Equals(LockLevel other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
