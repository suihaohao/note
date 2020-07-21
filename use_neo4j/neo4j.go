package use_neo4j

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"strings"
)

func GetLabelInfo(label string, props string) (interface{}, error) {
	driver := GetNeo4jConn()
	session, err := (*driver).Session(neo4j.AccessModeRead)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer session.Close()
	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []Node
		var result neo4j.Result
		if result, err = tx.Run(fmt.Sprintf("MATCH (n:%s%s) RETURN n", label, props), nil); err != nil {
			fmt.Println("tx.Run", err)
			return nil, err
		}
		for result.Next() {
			neo4jNode := result.Record().GetByIndex(0).(neo4j.Node)
			neoNode := Node{
				Id:     neo4jNode.Id(),
				Labels: neo4jNode.Labels(),
				Props:  neo4jNode.Props(),
			}
			list = append(list, neoNode)
		}
		if err = result.Err(); err != nil {
			fmt.Errorf("result.Err: %v", err)
			return nil, err
		}
		return list, nil
	})
	if err != nil {
		fmt.Errorf("result %v", err)
		return nil, err
	}
	return result, nil
}

func GetRelationship(relationship string) (interface{}, error) {
	driver := GetNeo4jConn()
	session, err := (*driver).Session(neo4j.AccessModeRead)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer session.Close()
	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []Neo4jPath
		var result neo4j.Result
		if result, err = tx.Run(fmt.Sprintf("MATCH n=()-[r:%s]->() RETURN n", relationship), nil); err != nil {
			fmt.Println("tx.Run", err)
			return nil, err
		}
		for result.Next() {
			neo4jPath := Neo4jPath{}
			neoPath := result.Record().GetByIndex(0).(neo4j.Path)
			for _, item := range neoPath.Relationships() {
				neoRelationship := Relationship{
					Id:      item.Id(),
					StartId: item.StartId(),
					EndId:   item.EndId(),
					Type:    item.Type(),
					Props:   item.Props(),
				}
				neo4jPath.Relationship = append(neo4jPath.Relationship, neoRelationship)
			}
			for _, item := range neoPath.Nodes() {
				neoNode := Node{
					Id:     item.Id(),
					Labels: item.Labels(),
					Props:  item.Props(),
				}
				neo4jPath.Nodes = append(neo4jPath.Nodes, neoNode)
			}
			list = append(list, neo4jPath)
		}
		if err = result.Err(); err != nil {
			fmt.Errorf("result.Err: %v", err)
			return nil, err
		}
		return list, nil
	})
	if err != nil {
		fmt.Errorf("result %v", err)
		return nil, err
	}
	return result, nil
}

func propsMapToNeo4jQueryString(props map[string]string) string {
	str := strings.Builder{}
	str.WriteString("{")
	for k, v := range props {
		str.WriteString(k + ":'" + v + "'")
	}
	str.WriteString("}")
	return str.String()
}

