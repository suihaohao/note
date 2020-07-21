package use_neo4j

type Neo4jPath struct {
	Nodes        []Node         `json:"nodes"`
	Relationship []Relationship `json:"relationship"`
}

type Relationship struct {
	Id      int64                  `json:"id"`
	StartId int64                  `json:"start_id"`
	EndId   int64                  `json:"end_id"`
	Type    string                 `json:"type"`
	Props   map[string]interface{} `json:"props"`
}

type Node struct {
	Id     int64                  `json:"id"`
	Labels []string               `json:"labels"`
	Props  map[string]interface{} `json:"props"`
}
