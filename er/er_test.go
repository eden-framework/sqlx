package er_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/eden-framework/sqlx/er"
	"github.com/eden-framework/sqlx/generator/__examples__/database"
	"github.com/eden-framework/sqlx/postgresqlconnector"
)

func TestDatabaseERFromDB(t *testing.T) {
	er := er.DatabaseERFromDB(database.DBTest, &postgresqlconnector.PostgreSQLConnector{})
	data, _ := json.MarshalIndent(er, "", "  ")

	fmt.Println(string(data))
}
