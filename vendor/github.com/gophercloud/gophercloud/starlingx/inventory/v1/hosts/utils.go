/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2019 Wind River Systems, Inc. */

package hosts

import (
	"fmt"
	"strings"
)

// Idle returns whether the host is currently running a maintenance task that
// may prevent other actions from running.
func (in *Host) Idle() bool {
	return in.Task == nil || *in.Task == ""
}

// Stable returns whether the operational status and administrative states
// agree.  If they do not then the host is still transitioning to a different
// state.  The host must also be "Idle" to be considered stable.
func (in *Host) Stable() bool {
	return in.Idle() &&
		((in.AdministrativeState == AdminUnlocked && in.OperationalStatus == OperEnabled) ||
			(in.AdministrativeState == AdminLocked && in.OperationalStatus == OperDisabled))
}

// IsUnlockedEnabled is a convenience utility to determine whether a host is
// unlocked and enabled.
func (in *Host) IsUnlockedEnabled() bool {
	return in.AdministrativeState == AdminUnlocked &&
		in.OperationalStatus == OperEnabled &&
		in.Idle()
}

// IsLockedDisabled is a convenience utility to determine whether a host is
// locked and disabled.
func (in *Host) IsLockedDisabled() bool {
	return in.AdministrativeState == AdminLocked &&
		in.OperationalStatus == OperDisabled &&
		in.Idle()
}

// IsUnlockedAvailable is a convenience utility to determine whether a host is
// unlocked, enabled, and available.
func (in *Host) IsUnlockedAvailable() bool {
	return in.IsUnlockedEnabled() && in.AvailabilityStatus == AvailAvailable
}

// IsInventoryCollected is a convenience utility to determine whether initial
// inventory collection has completed on a host.  It is not safe to configure
// or unlock a host until inventory collection has completed.
func (in *Host) IsInventoryCollected() bool {
	return in.InventoryState != nil && *in.InventoryState == InventoryCollected
}

// State returns a string representation of the host's administrative,
// operational and availability state/status.
func (in *Host) State() string {
	task := "idle"
	if !in.Idle() {
		task = strings.ToLower(*in.Task)
	}
	return fmt.Sprintf("%s/%s/%s/%s",
		in.AdministrativeState, in.OperationalStatus, in.AvailabilityStatus, task)
}
