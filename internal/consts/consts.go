package consts

const (
	VERSION                = "0.0.1"
	EnvironmentDevelopment = "development"
	EnvironmentStaging     = "staging"
	EnvironmentProduction  = "production"

	DriverName = "oci8"

	// DefaultPidFilename is default filename of pid file
	DefaultPidFilename = "al_hilal_core.pid"

	// DefaultLockFilename is default filename of lock file
	DefaultLockFilename = "al_hilal_core.lock"

	DefaultTempDirName = "al_hilal_core_temp"

	// DefaultWorkdirName name of working directory
	DefaultWorkdirName = "al_hilal_core_data"

	// byte bool
	False byte = 0
	True  byte = 1
)
