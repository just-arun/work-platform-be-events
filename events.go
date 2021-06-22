package platformevents

// for user events
type UserEvent string

const (
	UserEventCreate         UserEvent = "create-user"
	UserEventUpdate         UserEvent = "update-user"
	UserEventDisableUser    UserEvent = "disable-user"
	UserEventDeleteUser     UserEvent = "delete-user"
	UserEventUpdateUserType UserEvent = "update-user-type"
)


