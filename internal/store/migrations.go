package store

// migrations contains the SQL statements to create the application schema.
var migrations = []string{
	`CREATE TABLE IF NOT EXISTS landmarks (
		id              INTEGER PRIMARY KEY AUTOINCREMENT,
		latitude        REAL NOT NULL,
		longitude       REAL NOT NULL,
		image_filename  TEXT NOT NULL,
		year_built      INTEGER NOT NULL,
		year_destroyed  INTEGER
	)`,
	`CREATE TABLE IF NOT EXISTS landmark_translations (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
		landmark_id INTEGER NOT NULL,
		locale      TEXT NOT NULL,
		name        TEXT NOT NULL,
		description TEXT NOT NULL,
		history     TEXT NOT NULL,
		FOREIGN KEY (landmark_id) REFERENCES landmarks(id),
		UNIQUE(landmark_id, locale)
	)`,
	`CREATE TABLE IF NOT EXISTS categories (
		id    INTEGER PRIMARY KEY AUTOINCREMENT,
		slug  TEXT NOT NULL UNIQUE,
		color TEXT NOT NULL
	)`,
	`CREATE TABLE IF NOT EXISTS category_translations (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
		category_id INTEGER NOT NULL,
		locale      TEXT NOT NULL,
		name        TEXT NOT NULL,
		FOREIGN KEY (category_id) REFERENCES categories(id),
		UNIQUE(category_id, locale)
	)`,
	`ALTER TABLE landmarks ADD COLUMN category_id INTEGER REFERENCES categories(id) DEFAULT 1`,
	`ALTER TABLE landmarks ADD COLUMN highlighted INTEGER NOT NULL DEFAULT 0`,
}
