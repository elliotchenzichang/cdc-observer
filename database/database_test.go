package database

import (
	"cdc-observer/constant"
	dockerapi "cdc-observer/docker_api"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDatabaseAndAddNewTable(t *testing.T) {
	// Create a new Docker client
	dockerClient, err := dockerapi.NewDockerClient()
	assert.NoError(t, err, "Failed to create Docker client")

	// Start the MySQL container and get the assigned port
	ctx := context.Background()
	err = dockerClient.StartMySQLContainer(ctx)
	assert.NoError(t, err, "Failed to start MySQL container")
	defer func() {
		dockerClient.StopAllContainers(ctx)
		dockerClient.RemoveAllContainers(ctx)
	}()

	containerName := dockerClient.ContainerName(constant.MysqlImageName)
	port, err := dockerClient.ContainerPort(ctx, containerName)
	assert.NoError(t, err, "Failed to get MySQL container port")

	// Initialize the database
	db, err := NewDatabase(port)
	assert.NoError(t, err, "Failed to create new database")

	// Create a new table
	table, err := NewTableBuilder("test_table", db.dbClient).
		AddFieldInt("test_field_int").
		AddFieldVarchar("test_field_string").
		Submit()
	assert.NoError(t, err, "Failed to create new table")

	// Add the table to the database
	err = db.AddTable(table)
	assert.NoError(t, err, "Failed to add table to database")

	// Apply changes
	err = db.Apply()
	assert.NoError(t, err, "Failed to apply changes")

	// Add a row to the table
	r := NewRowBuilder().
		AddField("test_field_int", int64(1)).
		AddField("test_field_string", "test string").
		Submit()
	table.AddRow(r)
	assert.NoError(t, err, "Failed to add row to table")
}
