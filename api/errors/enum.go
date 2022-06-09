package errors

type FpErrorCode string

const (
	NoResults      FpErrorCode = "NoResults"
	RecordNotFound FpErrorCode = "RecordNotFound"
)

func (e FpErrorCode) String() string {
	return string(e)
}
