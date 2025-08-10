package repository

import (
	"github.com/huandu/go-sqlbuilder"
	"strings"
)

const Builder = sqlbuilder.PostgreSQL

func buildOnConflictFields(ib *sqlbuilder.InsertBuilder, onConflictFields []string, updateField ...string) {
	if len(updateField) == 0 || len(onConflictFields) == 0 {
		return
	}

	ib.SQL(` ON CONFLICT ("` + strings.Join(onConflictFields, "\", \"") + `") DO UPDATE SET`)

	for i, field := range updateField {
		if i != 0 {
			ib.SQL(",")
		}
		ib.SQL(`"` + field + `" = excluded."` + field + `"`)
	}
}
