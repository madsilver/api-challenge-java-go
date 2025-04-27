package entity

type User struct {
	Id      string `json:"id"`
	Name    string `json:"nome"`
	Age     int    `json:"idade"`
	Score   int    `json:"score"`
	Active  bool   `json:"ativo"`
	Country string `json:"pais"`
	Team    Team   `json:"equipe"`
	Logs    []Log  `json:"logs"`
}

type Project struct {
	Name      string `json:"nome"`
	Completed bool   `json:"concluido"`
}

type Log struct {
	Date   string `json:"data"`
	Action string `json:"acao"`
}

type Team struct {
	Name     string    `json:"nome"`
	Leader   bool      `json:"lider"`
	Projects []Project `json:"projetos"`
}

type Country struct {
	Country string `json:"country"`
	Total   int    `json:"total"`
}

type TeamInfo struct {
	Team              string  `json:"team"`
	TotalMembers      int     `json:"total_members"`
	Leaders           int     `json:"leaders"`
	CompletedProjects int     `json:"completed_projects"`
	ActivePercentage  float64 `json:"active_percentage"`
}

type Login struct {
	Date  string `json:"date"`
	Total int    `json:"total"`
}

type Evaluation struct {
	Status        int   `json:"status"`
	TimeMs        int64 `json:"time_ms"`
	ValidResponse bool  `json:"valid_response"`
}
