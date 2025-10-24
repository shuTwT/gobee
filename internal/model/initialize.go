package model

type InitializeRequest struct {
	DBType string `json:"db_type" binding:"required"`

	// SQLite
	SqliteFile string `json:"sqlite_file"`

	// MySQL
	MysqlHost     string `json:"mysql_host"`
	MysqlPort     string `json:"mysql_port"`
	MysqlUser     string `json:"mysql_user"`
	MysqlPassword string `json:"mysql_password"`
	MysqlDatabase string `json:"mysql_database"`

	// PostgreSQL
	PgHost     string `json:"pg_host"`
	PgPort     string `json:"pg_port"`
	PgUser     string `json:"pg_user"`
	PgPassword string `json:"pg_password"`
	PgDatabase string `json:"pg_database"`

	// Admin Account
	AdminUsername       string `json:"admin_username" binding:"required"`
	AdminEmail          string `json:"admin_email" binding:"required,email"`
	AdminPassword       string `json:"admin_password" binding:"required,min=8"`
	AdminPasswordConfirm string `json:"admin_password_confirm" binding:"required,min=8"`
}