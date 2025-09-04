package zoxide

var ErrZoxideNotFound = &zoxideNotFoundError{}

type zoxideNotFoundError struct{}

func (e *zoxideNotFoundError) Error() string {
	return "zoxide command not found in PATH"
}
