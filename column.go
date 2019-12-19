package model

import "fmt"

type Column struct {
	ColumnName    string
	Type          string
	Nullable      string
	TableName     string
	ColumnComment string
	Tag           string
}

func (c Column) GetTag(tagKey string) string {
	return fmt.Sprintf("`%s:\"%s\" json:\"%s\"`", tagKey, c.Tag, c.Tag)
}

func (c Column) GetGoType() string {
	v, ok := TypeMappingMysqlToGo[c.Type]
	if ok {
		return v
	}
	return ""
}

func (c Column) GetMysqlType() string {
	return c.Type
}

func (c Column) GetGoColumn(prefix string, ucFirst bool) string {
	return CamelCase(c.ColumnName, prefix, ucFirst)
}