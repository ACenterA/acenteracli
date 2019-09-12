package outputs

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/wallix/awless/cloud/properties"
	"github.com/wallix/awless/cloud/rdf"
	"github.com/wallix/awless/graph"
)

type rBuilder struct {
	id, typ string
	props   map[string]interface{}
}

func new(typ, id string) *rBuilder {
	r := &rBuilder{id: id, typ: typ, props: make(map[string]interface{})}
	return r.Prop(properties.Id, id)
}

func New(typ, id string) *rBuilder {
	r := &rBuilder{id: id, typ: typ, props: make(map[string]interface{})}
	return r.Prop(properties.Id, id)
}

func (b *rBuilder) Object(obj interface{}) *rBuilder {
	val := reflect.ValueOf(obj)
	for i := 0; i < val.Type().NumField(); i++ {
		// _ := val.Type().Field(i).Tag.Get("json") // was key
		fieldName := val.Type().Field(i).Name
		// fieldNameStr := fieldName[0:]
		// fmt.Println("f fieldname: ", fieldNameStr)
		if !(fieldName == "" || (fieldName != "" && fieldName != strings.Title(fieldName))) {
			// fmt.Println("TEST 1 : ", val.Type().Field(i).Name)
			// if fmt.Sprintf("%s", val.Type().Field(i).Name[0]) == strings.ToUpper(fmt.Sprintf("%v", val.Type().Field(i).Name[0])) {
			// fmt.Println("TEST 2a : ", val.Type().Field(i).Name)
			// v1 := val.FieldByName(val.Type().Field(i).Name).Interface()
			// fmt.Println("TEST 2c : ", fieldName)
			if !(fieldName == "id" || fieldName == "Id") {
				// if key != "" {
				// b.props[fieldName] = v1
				// fmt.Println("Added of ", fieldName)
				if _, ok := rdf.Properties[fieldName]; !ok {
					rdf.Labels[fieldName] = fieldName
					f := rdf.RdfProp{ID: fieldName, RdfType: "rdf:Property", RdfsLabel: fieldName, RdfsDefinedBy: "rdfs:Literal", RdfsDataType: "xsd:string"}
					rdf.Properties[fieldName] = f
				} else {

				}
				// }
			} else {
				// fmt.Println("TEST 2bad : ", fieldName)
			}
			// }
		} else {

		}
	}
	return b
}

func (b *rBuilder) Object1(obj interface{}) *rBuilder {
	val := reflect.ValueOf(obj)
	for i := 0; i < val.Type().NumField(); i++ {
		// _ := val.Type().Field(i).Tag.Get("json") // was key
		fieldName := val.Type().Field(i).Name
		// fieldNameStr := fieldName[0:]
		// fmt.Println("f fieldname: ", fieldNameStr)
		if !(fieldName == "" || (fieldName != "" && fieldName != strings.Title(fieldName))) {
			//fmt.Println("TEST 1 : ", val.Type().Field(i).Name)
			// if fmt.Sprintf("%s", val.Type().Field(i).Name[0]) == strings.ToUpper(fmt.Sprintf("%v", val.Type().Field(i).Name[0])) {
			// fmt.Println("TEST 2a : ", val.Type().Field(i).Name)
			// v1 := val.FieldByName(val.Type().Field(i).Name).Interface()
			// fmt.Println("TEST 2c : ", fieldName)
			if !(fieldName == "id" || fieldName == "Id") {
				// if key != "" {
				if _, ok := rdf.Properties[fieldName]; !ok {
					rdf.Labels[fieldName] = fieldName
					f := rdf.RdfProp{ID: fieldName, RdfType: "rdf:Property", RdfsLabel: fieldName, RdfsDefinedBy: "rdfs:Literal", RdfsDataType: "xsd:string"}
					rdf.Properties[fieldName] = f
				} else {

				}
				// }
			} else {
				//fmt.Println("TEST 2bad : ", fieldName)
			}
			// }
		} else {

		}
	}
	return b
}

func Project(id string) *rBuilder {
	return new("projectId", id)
}

func Name(id string) *rBuilder {
	return new("name", id)
}

func (b *rBuilder) Prop(key string, value interface{}) *rBuilder {
	b.props[key] = value
	return b
}

func (b *rBuilder) Build() *graph.Resource {
	res := graph.InitResource(b.typ, b.id)
	for k, v := range b.props {
		// fmt.Println("SET PROPERTIES OF :", k, " and ", v)
		res.Properties()[k] = v
	}

	return res
}

func AddParents(g *graph.Graph, relations ...string) {
	for _, rel := range relations {
		splits := strings.Split(rel, "->")
		if len(splits) != 2 {
			panic(fmt.Sprintf("invalid relation '%s'", rel))
		}
		r1 := graph.InitResource("", strings.TrimSpace(splits[0]))
		r2 := graph.InitResource("", strings.TrimSpace(splits[1]))
		err := g.AddParentRelation(r1, r2)
		if err != nil {
			panic(err)
		}
	}
}
