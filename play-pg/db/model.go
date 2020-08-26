package db

// policy and condition share same version, a copy rather foreign key
type PolicyBasic struct {
	tableName struct{} `pg:"policy_basic"`

	// policy
	Id              int      `json:"id" pg:",pk"`
	Pid             int      `json:"pid" pg:"pid"`
	UUID            string   `json:"uuid" pg:"uuid"`
	ParentId        int      `json:"parent_id" pg:"parent_id"`
	Version         int      `json:"version" pg:"version"` // version starts from 1, and can't be zero for we use zero for NULL in DB
	Status          string   `json:"status" pg:"status"`
	Level           int      `json:"level" pg:"level"`
	PCode           string   `json:"p_code" pg:"p_code" pg:",unique"`
	PName           string   `json:"p_name" pg:"p_name"`
	PType           *int     `json:"p_type" pg:"p_type"` // enum
	PTypeS          string   `json:"p_type_s" pg:"p_type_s"`
	PDepartment     []int    `json:"p_department" pg:",array" pg:"p_department"`
	PDepartmentS    string   `json:"p_department_s" pg:"p_department_s"`
	PClass          *int     `json:"p_class" pg:"p_class"` // enum
	PClassS         string   `json:"p_class_s" pg:"p_class_s"`
	PFunctionType   []int    `json:"p_function_type" pg:",array" pg:"p_function_type"`
	PFunctionTypeS  string   `json:"p_function_type_s" pg:"p_function_type_s"`
	PSource         string   `json:"p_source" pg:"p_source"`
	PPublishDate    string   `json:"p_publish_date" pg:"p_publish_date"`
	PPublishRegion  []int    `json:"p_publish_region" pg:",array" pg:"p_publish_region"` // enum
	PPublishRegionS string   `json:"p_publish_region_s" pg:"p_publish_region_s"`
	PKeywords       []string `json:"p_keywords" pg:",array" pg:"p_keywords"`

	// subject
	SName     string `json:"s_name" pg:"s_name"`
	SAnalysis string `json:"s_analysis" pg:"s_analysis"`

	// subject declaration
	DName        string         `json:"d_name" pg:"d_name"`
	DAcceptDep   string         `json:"accept_dep" pg:"d_accept_dep"`
	DExecCircle  string         `json:"d_exec_circle" pg:"d_exec_circle"`
	DRequires    string         `json:"d_required" pg:"d_requires"`
	DAnalysis    string         `json:"d_analysis" pg:"d_analysis"`
	DDeclPlat    string         `json:"d_decl_plat" pg:"d_decl_plat"`
	DeclDuration []DeclDuration `json:"d_decl_duration" pg:"d_decl_duration"`
	DeclAid      []DeclAid      `json:"d_decl_aid" pg:"d_decl_aid" pg:",array"`

	// condition summary
	CSummary string `json:"c_summary" pg:"c_summary"`
	OpTime   string `json:"op_time" pg:"op_time"`

	UpgradeReason string `json:"upgrade_reason" pg:"upgrade_reason"`
}

type DeclDuration struct {
	Region   []int    `json:"region"  pg:",array"` // enums
	Title    string   `json:"title"`
	Duration []string `json:"duration" pg:",array"`
	Key      string   `json:"key"`
	Seq      int      `json:"seq"`
}

type DeclAid struct {
	Title  string `json:"title"`
	Key    string `json:"key"`
	Seq    int    `json:"seq"`
	Region []int  `json:"region" pg:",array"` // enums
	Policy string `json:"policy"`
}
