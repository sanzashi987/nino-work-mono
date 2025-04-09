package model

const (
	STATIC     = "Static"
	API        = "API"
	FILE       = "File"
	PASSIVE    = "Passive"
	MYSQL      = "MySQL"
	ORACLE     = "Oracle"
	SQLSERVER  = "SQLServer"
	POSTGRESQL = "PostgreSQL"
)

var SourceTypeIntToString = [8]string{
	STATIC,
	API,
	FILE,
	PASSIVE,
	MYSQL,
	ORACLE,
	SQLSERVER,
	POSTGRESQL,
}

var SourceTypeStringToEnum = map[string]uint8{}

func init() {
	for index, str := range SourceTypeIntToString {
		SourceTypeStringToEnum[str] = uint8(index)
	}
}

type DataSourceModel struct {
	BaseModel
	Version    string
	SourceType uint8  `gorm:"index"`
	SourceInfo string `gorm:"type:blob"`
}

func (m DataSourceModel) TableName() string {
	return "data_sources"
}
