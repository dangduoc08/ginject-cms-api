package postgres

import (
	"fmt"
	"strings"

	"github.com/dangduoc08/ginject-cms-api/internal/common/slice"
	"github.com/dangduoc08/ginject/core"
	"gorm.io/gorm"
)

type PostgresProvider struct {
	DB *gorm.DB
}

func (instance PostgresProvider) NewProvider() core.Provider {

	return instance
}

func (instance PostgresProvider) CreateEnum(typeName string, values []string) {
	formatValues := slice.Map(values, func(el string, i int) string {
		return fmt.Sprintf("'%s'", el)
	})
	sql := fmt.Sprintf(`
		DO $$ BEGIN
			IF
				NOT EXISTS (SELECT oid FROM pg_type WHERE typname = '%v')
			THEN
				CREATE TYPE %v AS ENUM (%v);
			END IF;
		END $$;`, typeName, typeName, strings.Join(formatValues, ", "))

	resp := instance.DB.Exec(sql)
	if resp.Error != nil {
		panic(resp.Error)
	}
}

func (instance PostgresProvider) Count(tableName string) int {
	var count int
	if err := instance.DB.Raw(fmt.Sprintf("SELECT count(*) FROM %v", tableName)).Scan(&count).Error; err != nil {
		panic(err)
	}

	return count
}
