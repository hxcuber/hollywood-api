package migrations

//go:generate sqlboiler --wipe psql

import (
	_ "github.com/lib/pq"
)
