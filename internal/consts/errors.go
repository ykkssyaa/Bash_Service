package consts

const (
	ErrorNotJSON              = "Content Type is not application/json"
	ErrorBadRequestWrongField = "Bad Request: wrong type provided for field "
	ErrorBadRequest           = "Bad Request "
	ErrorWrongId              = "Bad Request: wrong command id"
	ErrorWrongLimit           = "Bad Request: wrong limit value"
	ErrorWrongOffset          = "Bad Request: wrong offset value"

	ErrorEmptyScript = "Bad Request: empty script"

	ErrorCreateCommand = "Error with creating command"
	ErrorUpdateCommand = "Error with updating command"
	ErrorGetCommand    = "Error with getting command"
	ErrorExecCommand   = "Error with execution command"

	ErrorUpdateCacheCommand = "Error with updating command in cache"
	ErrorRemoveCacheCommand = "Error with removing command from cache"
	ErrorGetCacheCommand    = "Error with getting command from cache"
)
