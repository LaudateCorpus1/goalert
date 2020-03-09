// Code generated by "stringer -type TargetType"; DO NOT EDIT.

package assignment

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TargetTypeUnspecified-0]
	_ = x[TargetTypeEscalationPolicy-1]
	_ = x[TargetTypeNotificationPolicy-2]
	_ = x[TargetTypeRotation-3]
	_ = x[TargetTypeService-4]
	_ = x[TargetTypeSchedule-5]
	_ = x[TargetTypeCalendarSubscription-6]
	_ = x[TargetTypeUser-7]
	_ = x[TargetTypeNotificationChannel-8]
	_ = x[TargetTypeSlackChannel-9]
	_ = x[TargetTypeIntegrationKey-10]
	_ = x[TargetTypeUserOverride-11]
	_ = x[TargetTypeNotificationRule-12]
	_ = x[TargetTypeContactMethod-13]
	_ = x[TargetTypeHeartbeatMonitor-14]
}

const _TargetType_name = "TargetTypeUnspecifiedTargetTypeEscalationPolicyTargetTypeNotificationPolicyTargetTypeRotationTargetTypeServiceTargetTypeScheduleTargetTypeCalendarSubscriptionTargetTypeUserTargetTypeNotificationChannelTargetTypeSlackChannelTargetTypeIntegrationKeyTargetTypeUserOverrideTargetTypeNotificationRuleTargetTypeContactMethodTargetTypeHeartbeatMonitor"

var _TargetType_index = [...]uint16{0, 21, 47, 75, 93, 110, 128, 158, 172, 201, 223, 247, 269, 295, 318, 344}

func (i TargetType) String() string {
	if i < 0 || i >= TargetType(len(_TargetType_index)-1) {
		return "TargetType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TargetType_name[_TargetType_index[i]:_TargetType_index[i+1]]
}
