package errors

type FpErrorCode string

const (
	RecordNotFound FpErrorCode = "RecordNotFound"
)

func (e FpErrorCode) String() string {
	return string(e)
}
