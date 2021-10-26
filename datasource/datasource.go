package datasource

type DataSource interface {
	Close() error
}
