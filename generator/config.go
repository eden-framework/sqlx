package generator

type Config struct {
	StructName string
	TableName  string
	Database   string

	WithComments        bool
	WithTableName       bool
	WithTableInterfaces bool
	WithMethods         bool

	FieldPrimaryKey   string
	FieldKeyDeletedAt string
	FieldKeyCreatedAt string
	FieldKeyUpdatedAt string
}

func (g *Config) SetDefaults() {
	if g.FieldKeyDeletedAt == "" {
		g.FieldKeyDeletedAt = "DeletedAt"
	}

	if g.FieldKeyCreatedAt == "" {
		g.FieldKeyCreatedAt = "CreatedAt"
	}

	if g.FieldKeyUpdatedAt == "" {
		g.FieldKeyUpdatedAt = "UpdatedAt"
	}

	if g.TableName == "" {
		g.TableName = toDefaultTableName(g.StructName)
	}
}
