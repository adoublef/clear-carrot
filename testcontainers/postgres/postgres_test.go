package postgres

import (
	"context"
	"database/sql"
	"path/filepath"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	dbname   = "test-db"
	user     = "postgres"
	password = "password"
)

func TestWithInitScript(t *testing.T) {
	ctx := context.Background()

	// container, err := postgres.RunContainer(ctx,
	// 	testcontainers.WithImage("docker.io/postgres:15.2-alpine"),
	// 	postgres.WithInitScripts(filepath.Join("testdata", "init.sql")),
	// 	postgres.WithDatabase(dbname),
	// 	postgres.WithUsername(user),
	// 	postgres.WithPassword(password),
	// 	testcontainers.WithWaitStrategy(
	// 		wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	// )
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// t.Cleanup(func() {
	// 	if err := container.Terminate(ctx); err != nil {
	// 		t.Fatalf("failed to terminate container: %s", err)
	// 	}
	// })

	// explicitly set sslmode=disable because the container is not configured to use TLS
	// connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	// assert.NoError(t, err)

	// db, err := sql.Open("postgres", connStr)
	// assert.NoError(t, err)
	// assert.NotNil(t, db)
	// defer db.Close()

	// database created in init script. See testdata/init-user-db.sh
	// result, err := db.Exec("SELECT * FROM testdb;")
	// assert.NoError(t, err)
	// assert.NotNil(t, result)

	t.Run("select * from testdb", WithClient(ctx, func(db *sql.DB) {
		// database created in init script. See testdata/init-user-db.sh
		result, err := db.Exec("SELECT * FROM testdb;")
		assert.NoError(t, err)
		assert.NotNil(t, result)
	}))
}

func WithClient(ctx context.Context, fn func(db *sql.DB)) func(t *testing.T) {
	return func(t *testing.T) {
		container, err := postgres.RunContainer(ctx,
			testcontainers.WithImage("docker.io/postgres:15.2-alpine"),
			postgres.WithInitScripts(filepath.Join("testdata", "init.sql")),
			postgres.WithDatabase(dbname),
			postgres.WithUsername(user),
			postgres.WithPassword(password),
			testcontainers.WithWaitStrategy(
				wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)),
		)
		if err != nil {
			t.Fatal(err)
		}

		t.Cleanup(func() {
			if err := container.Terminate(ctx); err != nil {
				t.Fatalf("failed to terminate container: %s", err)
			}
		})

		connStr, err := container.ConnectionString(ctx, "sslmode=disable")
		assert.NoError(t, err)

		db, err := sql.Open("postgres", connStr)
		assert.NoError(t, err)
		assert.NotNil(t, db)
		defer db.Close()

		fn(db)
	}
}
