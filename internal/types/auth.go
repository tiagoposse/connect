package types

import (
	authz "github.com/tiagoposse/go-auth/controller"
)

type DisabledReason int64

const (
	ReasonProviderDisabled DisabledReason = iota
	ReasonProviderDeleted
	ReasonAdminDeleted
	ReasonAdminDisabled
)

func (s DisabledReason) String() string {
	switch s {
	case ReasonProviderDisabled:
		return "ProviderDisabled"
	case ReasonProviderDeleted:
		return "ProviderDeleted"
	case ReasonAdminDeleted:
		return "AdminDeleted"
	case ReasonAdminDisabled:
		return "AdminDisabled"
	}
	return "unknown"
}

const (
	AdminUsersWrite       authz.Scope = "admin.users.write"
	AdminUsersReadOnly    authz.Scope = "admin.users.readonly"
	AdminDevicesWrite     authz.Scope = "admin.groups.write"
	AdminDevicesReadOnly  authz.Scope = "admin.groups.readonly"
	AdminGroupsWrite      authz.Scope = "admin.devices.write"
	AdminGroupsReadOnly   authz.Scope = "admin.devices.readonly"
	AdminSettingsWrite    authz.Scope = "admin.settings.write"
	AdminSettingsReadOnly authz.Scope = "admin.settings.readonly"
	AdminAll              authz.Scope = "admin.*"
	UserDevicesWrite      authz.Scope = "user.devices.write"
	UserApiKeyWrite       authz.Scope = "user.apikey.write"
	UserDevicesReadOnly   authz.Scope = "user.devices.readonly"
	UserAll               authz.Scope = "user.*"
)

var AllScopes = authz.Scopes{
	AdminUsersWrite,
	AdminUsersReadOnly,
	AdminDevicesWrite,
	AdminDevicesReadOnly,
	AdminGroupsWrite,
	AdminGroupsReadOnly,
	AdminSettingsWrite,
	AdminSettingsReadOnly,
	AdminAll,
	UserDevicesWrite,
	UserApiKeyWrite,
	UserDevicesReadOnly,
	UserAll,
}
