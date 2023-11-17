package model

var MODEL_TYPE = map[string]interface{}{
	"profile":             Profile{},
	"credential":          Credential{},
	"role":                Role{},
	"temporaryCredential": TemporaryCredential{},
	"api":                 Api{},
	"apiRoleTeam":         ApiRoleTeam{},
	"project":             Project{},
	"projectProfile":      ProjectProfile{},
	"roleTeam":            RoleTeam{},
	"router":              Router{},
	"routerRoleTeam":      RouterRoleTeam{},
	"tableAction":         TableAction{},
	"tableActionRoleTeam": TableActionRoleTeam{},
	"team":                Team{},
	"teamProfile":         TeamProfile{},
}

var LIST_MODEL_TYPE = map[string]interface{}{
	"profile":             []Profile{},
	"credential":          []Credential{},
	"role":                []Role{},
	"temporaryCredential": []TemporaryCredential{},
	"api":                 []Api{},
	"apiRoleTeam":         []ApiRoleTeam{},
	"project":             []Project{},
	"projectProfile":      []ProjectProfile{},
	"roleTeam":            []RoleTeam{},
	"router":              []Router{},
	"routerRoleTeam":      []RouterRoleTeam{},
	"tableAction":         []TableAction{},
	"tableActionRoleTeam": []TableActionRoleTeam{},
	"team":                []Team{},
	"teamProfile":         []TeamProfile{},
}

type ROLE string

const (
	ADMIN ROLE = "admin"
	USER  ROLE = "user"
)

type OPTION_QUERY string

const (
	INSERT OPTION_QUERY = "insert"
	UPDATE OPTION_QUERY = "update"
	DELETE OPTION_QUERY = "delete"
)

type OPTION_CONVERT_FIELD string

const (
	TABLE_TO_MODEL OPTION_CONVERT_FIELD = "table_to_model"
	MODEL_TO_TABLE OPTION_CONVERT_FIELD = "model_to_table"
)
