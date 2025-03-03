# The "local" environment represents our local testings.
env "local" {
	url = "sqlite://turtle_cms.db"
	migration {
		dir = "file://db/migrations"
	}
}
