package fieldtype

type FieldType string

const (
	INT                = "number"
	VARCHAR            = "text"
	TEXT               = "text"
	DATE               = "date"
	TINYINT            = "number"
	SMALLINT           = "number"
	MEDIUMINT          = "number"
	BIGINT             = "number"
	DECIMAL            = "text"
	FLOAT              = "text"
	DOUBLE             = "text"
	REAL               = "text"
	BIT                = "checkbox"
	BOOLEAN            = "checkbox"
	SERIAL             = "number"
	DATETIME           = "datetime-local"
	TIMESTAMP          = "number"
	TIME               = "time"
	YEAR               = "number"
	CHAR               = "text"
	TINYTEXT           = "text"
	MEDIUMTEXT         = "text"
	LONGTEXT           = "text"
	BINARY             = "file"
	VARBINARY          = "file"
	TINYBLOB           = "file"
	BLOB               = "file"
	MEDIUMBLOB         = "file"
	LONGBLOB           = "file"
	ENUM               = "text"
	SET                = "text"
	GEOMETRY           = "text"
	POINT              = "text"
	LINESTRING         = "text"
	POLYGON            = "text"
	MULTIPOINT         = "text"
	MULTILINESTRING    = "text"
	MULTIPOLYGON       = "text"
	GEOMETRYCOLLECTION = "text"
	JSON               = "text"
	TEL                = "tel"
)
