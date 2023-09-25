package cockroachdb

import (
	"context"
	"database/sql"
	"path/filepath"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)


func TestWithInitScript(t *testing.T) {
	ctx := context.Background()

	container, err := RunContainer(ctx,
		// testcontainers.WithImage("docker.io/postgres:15.2-alpine"),
		WithInitScripts(filepath.Join("testdata", "init.sql")),
		testcontainers.WithWaitStrategy(
			wait.ForHTTP("/health").WithPort("8080")),
			// wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	// explicitly set sslmode=disable because the container is not configured to use TLS
	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	assert.NoError(t, err)

	db, err := sql.Open("postgres", connStr)
	assert.NoError(t, err)
	assert.NotNil(t, db)
	defer db.Close()

	// database created in init script. See testdata/init-user-db.sh
	result, err := db.Exec("SELECT * FROM testdb;")
	assert.NoError(t, err)
	assert.NotNil(t, result)
}