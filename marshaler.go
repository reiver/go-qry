package qry

type Marshaler interface {
	MarshalQRY() ([]byte, error)
}
