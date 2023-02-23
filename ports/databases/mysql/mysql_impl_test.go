package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/RyaWcksn/chibiko/configs"
	"github.com/RyaWcksn/chibiko/entities"
	"github.com/RyaWcksn/chibiko/pkgs/containers"
	"github.com/RyaWcksn/chibiko/pkgs/databases"
)

func TestDBImpl_Save(t *testing.T) {

	ctx := context.Background()

	mysqlContainer, err := containers.SetupMysqlContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Clean up the container after test is complete
	t.Cleanup(func() {
		if err := mysqlContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	// Assertion
	host, _ := mysqlContainer.Host(ctx)
	p, _ := mysqlContainer.MappedPort(ctx, "3306/tcp")
	port := p.Int()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?tls=skip-verify",
		"root", "password", host, port, "database")

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		t.Errorf("error pinging db: %+v\n", err)
	}

	dbConf := configs.Database{
		Host:     host,
		Port:     port,
		Database: "database",
		Password: "password",
		Username: "root",
	}

	mysqlInstance := databases.NewMysqlConnection(dbConf)
	connection := mysqlInstance.DBConnect()

	type fields struct {
		sql *sql.DB
	}
	type args struct {
		ctx    context.Context
		entity *entities.SaveDatabase
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantId  int64
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				sql: connection,
			},
			args: args{
				ctx: ctx,
				entity: &entities.SaveDatabase{
					Url: "https://google.com",
				},
			},
			wantId:  1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql := NewSql(tt.fields.sql)
			gotId, err := sql.Save(tt.args.ctx, tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("DBImpl.Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotId != tt.wantId {
				t.Errorf("DBImpl.Save() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}
