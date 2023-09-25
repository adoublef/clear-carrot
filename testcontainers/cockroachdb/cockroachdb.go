package cockroachdb

import (
	"context"
	"fmt"
	"net"
	"path/filepath"
	"strings"

	"github.com/testcontainers/testcontainers-go"
)

const (
	defaultUser             = "root"
	defaultPassword         = ""
	defaultDatabase         = "defaultdb"
	defaultCockroachDBImage = "docker.io/cockroachdb/cockroach:latest-v22.2"
)

// CockroachDBContainer represents the postgres container type used in the module
type CockroachDBContainer struct {
	testcontainers.Container
	dbName   string
	user     string
	password string
}

// ConnectionString returns the connection string for the postgres container, using the default 5432 port, and
// obtaining the host and exposed port from the container. It also accepts a variadic list of extra arguments
// which will be appended to the connection string. The format of the extra arguments is the same as the
// connection string format, e.g. "connect_timeout=10" or "application_name=myapp"
func (c *CockroachDBContainer) ConnectionString(ctx context.Context, args ...string) (string, error) {
	containerPort, err := c.MappedPort(ctx, "26257/tcp")
	if err != nil {
		return "", err
	}

	host, err := c.Host(ctx)
	if err != nil {
		return "", err
	}

	extraArgs := strings.Join(args, "&")
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?%s", c.user, c.password, net.JoinHostPort(host, containerPort.Port()), c.dbName, extraArgs)
	return connStr, nil
}

// WithInitScripts sets the init scripts to be run when the container starts
func WithInitScripts(scripts ...string) testcontainers.CustomizeRequestOption {
	return func(req *testcontainers.GenericContainerRequest) {
		initScripts := []testcontainers.ContainerFile{}
		for _, script := range scripts {
			cf := testcontainers.ContainerFile{
				HostFilePath:      script,
				ContainerFilePath: "/docker-entrypoint-initdb.d/" + filepath.Base(script),
				FileMode:          0o755,
			}
			initScripts = append(initScripts, cf)
		}
		req.Files = append(req.Files, initScripts...)
	}
}

// RunContainer creates an instance of the postgres container type
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*CockroachDBContainer, error) {
	req := testcontainers.ContainerRequest{
		Image: defaultCockroachDBImage,
		Env: map[string]string{
			"COCKROACH_USER":     defaultUser,
			"COCKROACH_PASSWORD": defaultPassword,
			"COCKROACH_DATABASE": defaultDatabase, // defaults to the user name
		},
		ExposedPorts: []string{"26257/tcp", "8080/tcp"},
		Cmd:          []string{"start-single-node", "--insecure"},
		// NOTE -- for demo purposes only
		Mounts: testcontainers.ContainerMounts{
			{
				Source: testcontainers.GenericVolumeMountSource{
					Name: "cockroach-data",
				},
				Target: "/cockroach/cockroach-data",
			},
		},
	}

	genericContainerReq := testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	}

	for _, opt := range opts {
		opt.Customize(&genericContainerReq)
	}

	container, err := testcontainers.GenericContainer(ctx, genericContainerReq)
	if err != nil {
		return nil, err
	}

	user := req.Env["COCKROACH_USER"]
	password := req.Env["COCKROACH_PASSWORD"]
	dbName := req.Env["COCKROACH_DATABASE"]

	return &CockroachDBContainer{Container: container, dbName: dbName, password: password, user: user}, nil
}
