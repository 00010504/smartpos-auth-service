package config

const (
	DateTimeFormat = "2006-01-02 15:04:05"
	DateFormat     = "2006-01-02"

	ConsumerGroupID = "auth_service"

	SystemUserTypeID = "1fe92aa8-2a61-4bf1-b907-182b497584ad"
	AdminUserTypeID  = "9fb3ada6-a73b-4b81-9295-5c1605e54552"

	SystemTypeID = "1fe92aa8-2a61-4bf1-b907-182b497584ad"
	AdminTypeID  = "9fb3ada6-a73b-4b81-9295-5c1605e54552"

	UserPendingStatus  = "1fe92aa8-2a61-4bf1-b907-182b497584ad" // pending
	UserActivetStatus  = "9fb3ada6-a73b-4b81-9295-5c1605e54552" // active
	UserRejectedStatus = "0adc982c-749b-4446-9d36-d136a76b99ea" // rejected
	UserBlockesStatus  = "3e6eff54-dd23-4603-99f6-8f5fc24d19ff" // blocked
)
