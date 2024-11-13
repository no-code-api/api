package migrations

type Migration interface {
	GetId() string
	GetDescription() string
	Operations() []string
}
