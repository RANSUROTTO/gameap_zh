package sqlite_test

import (
	"testing"

	"github.com/gameap/gameap/internal/repositories"
	"github.com/gameap/gameap/internal/repositories/sqlite"
	repotesting "github.com/gameap/gameap/internal/repositories/testing"
	"github.com/stretchr/testify/suite"
)

func TestDLQRepository(t *testing.T) {
	suite.Run(t, repotesting.NewDLQRepositorySuite(
		func(t *testing.T) repositories.DLQRepository {
			t.Helper()

			return sqlite.NewDLQRepository(SetupTestDB(t))
		},
	))
}
