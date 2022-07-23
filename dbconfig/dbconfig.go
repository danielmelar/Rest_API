package dbconfig

const (
	PostgresDriver = "postgres"
	Host           = "localhost"
	Port           = "5432"
	User           = "postgres"
	Password       = "1234"
	DbName         = "test"
	TbName         = "livros"
)

//var ConnectionString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s"+
//	"tbname=%s, sslmode=disable", Host, User, Password, DbName, Port, TbName)
