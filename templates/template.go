package templates

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
	"github.com/stoewer/go-strcase"
	"github.com/tiagoposse/connect/filter"
	exclusion "github.com/tiagoposse/entoas-fields-exclusion"
)

var (
	// funcMap contains extra template functions used by ogent.
	funcMap = template.FuncMap{
		"convertTo":       convertTo,
		"eagerLoad":       eagerLoad,
		"edgeOperations":  entoas.EdgeOperations,
		"edgeViewName":    entoas.EdgeViewName,
		"fieldAnnotation": fieldAnnotation,
		"isComplexType": isComplexType,
		"hasParams":       hasParams,
		"hasRequestBody":  hasRequestBody,
		"httpRoute":       httpRoute,
		"httpVerb":        httpVerb,
		"isCreate":        isCreate,
		"returnTotal":     returnTotal,
		"isDelete":        isDelete,
		"isList":          isList,
		"isRead":          isRead,
		"isUpdate":        isUpdate,
		"kebab":           strcase.KebabCase,
		"nodeOperations":  entoas.NodeOperations,
		"ogenToEnt":       ogenToEnt,
		"replaceAll":      strings.ReplaceAll,
		"setFieldExpr":    setFieldExpr,
		"viewName":        entoas.ViewName,
		"viewNameEdge":    entoas.ViewNameEdge,
		"pagination":      pagination,
	}
)

func NewTemplates(tplDir string) *gen.Template {
	return gen.MustParse(gen.NewTemplate("wg").Funcs(funcMap).ParseFS(os.DirFS(tplDir), "*.tmpl"))
}

type (
	Edges []*Edge
	Edge  entoas.Edge
)

// eagerLoad returns the Go expression to eager load the required edges on the node operation.
func eagerLoad(n *gen.Type, op entoas.Operation) (string, error) {
	gs, err := entoas.GroupsForOperation(n.Annotations, op)
	if err != nil {
		return "", err
	}
	t, err := entoas.EdgeTree(n, gs)
	if err != nil {
		return "", err
	}
	if len(t) > 0 {
		es := make(Edges, len(t))
		for i, e := range t {
			es[i] = (*Edge)(e)
		}
		return es.entQuery(), nil
	}
	return "", nil
}

// entQuery runs entQuery on every Edge and appends them.
func (es Edges) entQuery() string {
	b := new(strings.Builder)
	for _, e := range es {
		b.WriteString(e.entQuery())
	}
	return b.String()
}

// EntQuery constructs the Go code to eager load all requested edges for the given one.
func (e Edge) entQuery() string {
	b := new(strings.Builder)
	fmt.Fprintf(b, ".%s(", strings.Title(e.EagerLoadField()))
	if len(e.Edges) > 0 {
		es := make(Edges, len(e.Edges))
		for i, e := range e.Edges {
			es[i] = (*Edge)(e)
		}
		fmt.Fprintf(
			b,
			"func (q *%s.%s) {\nq%s\n}",
			filepath.Base(e.Type.Config.Package),
			e.Type.QueryName(),
			es.entQuery(),
		)
	}
	b.WriteString(")")
	return b.String()
}

// hasParams returns if the given entoas.Operation has parameters.
func hasParams(op entoas.Operation) bool {
	return op != entoas.OpCreate
}

// hasRequestBody returns if the given entoas.Operation has a request body.
func hasRequestBody(op entoas.Operation) bool {
	return op == entoas.OpCreate || op == entoas.OpUpdate
}

// httpVerb returns the HTTP httpVerb for the given entoas.Operation.
func httpVerb(op entoas.Operation) (string, error) {
	switch op {
	case entoas.OpCreate:
		return http.MethodPost, nil
	case entoas.OpRead, entoas.OpList:
		return http.MethodGet, nil
	case entoas.OpUpdate:
		return http.MethodPatch, nil
	case entoas.OpDelete:
		return http.MethodDelete, nil
	}
	return "", fmt.Errorf("unknown operation: %q", op)
}

// httpRoute returns the HTTP endpoint for the given entoas.Operation.
func httpRoute(root string, op entoas.Operation) (string, error) {
	switch op {
	case entoas.OpCreate, entoas.OpList:
		return root, nil
	case entoas.OpRead, entoas.OpUpdate, entoas.OpDelete:
		return path.Join(root, "{id}"), nil
	}
	return "", fmt.Errorf("unknown operation: %q", op)
}

// isCreate returns if the given entoas.Operation is entoas.OpCreate.
func isCreate(op entoas.Operation) bool { return op == entoas.OpCreate }

// isRead returns if the given entoas.Operation is entoas.OpRead.
func isRead(op entoas.Operation) bool { return op == entoas.OpRead }

// isUpdate returns if the given entoas.Operation is entoas.OpUpdate.
func isUpdate(op entoas.Operation) bool { return op == entoas.OpUpdate }

// isDelete returns if the given entoas.Operation is entoas.OpDelete.
func isDelete(op entoas.Operation) bool { return op == entoas.OpDelete }

// isList returns if the given entoas.Operation is entoas.OpList.
func isList(op entoas.Operation) bool { return op == entoas.OpList }

// OAS Schema types.
const (
	Array   = "array"
	Integer = "integer"
	Number  = "number"
	String  = "string"
	Boolean = "boolean"
)

// OAS Schema formats.
const (
	None     = ""
	UUID     = "uuid"
	Date     = "date"
	Time     = "time"
	DateTime = "date-time"
	Duration = "duration"
	URI      = "uri"
	IPv4     = "ipv4"
	IPv6     = "ipv6"
	Byte     = "byte"
	Password = "password"
	Int64    = "int64"
	Int32    = "int32"
	Float    = "float"
	Double   = "double"
)

// setFieldExpr returns a Go expression to set the field on a response.
func setFieldExpr(f *gen.Field, schema, rec, ident string) (string, error) {
	if !f.Optional && !f.Nillable {
		expr := fmt.Sprintf("%s.%s", ident, f.StructField())
		if f.IsEnum() {
			expr = convertTo(schema+f.StructField(), expr)
		}

		expr = entToOgen(f, expr, schema)
		return fmt.Sprintf("%s.%s = %s", rec, f.StructField(), expr), nil
	}
	t, err := entoas.OgenSchema(f)
	if err != nil {
		return "", err
	}
	buf := new(strings.Builder)
	// Enums need special handling.
	if f.IsEnum() {
		fmt.Fprintf(buf, "NewOpt%s%s(%s)",
			schema, f.StructField(),
			convertTo(schema+f.StructField(), fmt.Sprintf("%s.%s", ident, f.StructField())),
		)
		return buf.String(), nil
	}

	var opt string
	switch t.Type {
	case Integer:
		switch t.Format {
		case Int32:
			opt = "Int32"
		case Int64:
			opt = "Int64"
		case None:
			opt = "Int"
		default:
			return "", fmt.Errorf("unexpected type: %q", t.Format)
		}
	case Number:
		switch t.Format {
		case Float:
			opt = "Float32"
		case Double, None:
			opt = "Float64"
		case Int32:
			opt = "Int32"
		case Int64:
			opt = "Int64"
		default:
			return "", fmt.Errorf("unexpected type: %q", t.Format)
		}
	case String:
		switch t.Format {
		case Byte:
			return fmt.Sprintf("%s.%s = %s.%s", rec, f.StructField(), ident, f.StructField()), nil
		case DateTime:
			opt = "DateTime"
		case Date:
			opt = "Date"
		case Time:
			opt = "Time"
		case Duration:
			opt = "Duration"
		case UUID:
			opt = "UUID"
		case IPv4, IPv6:
			opt = "IP"
		case URI:
			opt = "URL"
		case Password, None:
			opt = "String"
		default:
			return "", fmt.Errorf("unexpected type: %q", t.Format)
		}
	case Boolean:
		switch t.Format {
		case None:
			opt = "Bool"
		default:
			return "", fmt.Errorf("unexpected type: %q", t.Format)
		}
	default:
		return "", fmt.Errorf("unexpected type: %q", t.Format)
	}

	if f.Nillable {
		fmt.Fprintf(buf, "%s.%s = Opt%s{}\n", rec, f.StructField(), opt)
		fmt.Fprintf(buf, "if %s.%s != nil { %s.%s.SetTo(*%s.%s) }",
			ident, f.StructField(),
			rec, f.StructField(),
			ident, f.StructField(),
		)
	} else {
		expr := entToOgen(f, fmt.Sprintf("%s.%s", ident, f.StructField()), schema)
		fmt.Fprintf(buf, "%s.%s = NewOpt%s(%s)", rec, f.StructField(), opt, expr)
	}
	return buf.String(), nil
}

func convertTo(typ, expr string) string {
	return fmt.Sprintf("%s(%s)", typ, expr)
}

func ogenToEnt(f *gen.Field, expr string, name string) string {
	if f.Type.RType != nil {
		switch f.Type.Type {
		case field.TypeString:
			return fmt.Sprintf("%s{}.ParseString(%s)", f.Type.RType, expr)
		}
	}

	switch f.Type.Type {
	case field.TypeUint8:
		return fmt.Sprintf("uint8(%s)", expr)
	case field.TypeInt8:
		return fmt.Sprintf("int8(%s)", expr)
	case field.TypeUint16:
		return fmt.Sprintf("uint16(%s)", expr)
	case field.TypeInt16:
		return fmt.Sprintf("int16(%s)", expr)
	case field.TypeUint32:
		return fmt.Sprintf("uint32(%s)", expr)
	case field.TypeUint:
		return fmt.Sprintf("uint(%s)", expr)
	case field.TypeUint64:
		return fmt.Sprintf("uint64(%s)", expr)
	case field.TypeJSON, field.TypeOther:
		// bs, _ := json.Marshal(f.Type)
		// fmt.Println(string(bs))
		// fmt.Printf("%#v\n", f.Type.RType)
		// if regexp.MustCompile(`\*.+`).MatchString(f.Type.RType.Ident) {
		// 	return fmt.Sprintf("JsonConvert(%s, &%s{}).(*%s)", expr, f.Type.RType.Ident, f.Type.RType.Ident)
		// }

		// switch f.Type.RType.Ident {
		// case "[]string":
		// 	return fmt.Sprintf("*JsonConvert(%s, &%s{}).(*%s)", expr, f.Type.RType.Ident, f.Type.RType.Ident)
		// default:
		return fmt.Sprintf("*JsonConvert(%s, &%s{}).(*%s)", expr, f.Type.RType.Ident, f.Type.RType.Ident)
		// }
	default:
		return expr
	}
}

func entToOgen(f *gen.Field, expr, schemaName string) string {
	if f.Type.RType != nil {
		switch f.Type.Type {
		case field.TypeString:
			return fmt.Sprintf("%s.String()", expr)
		}
	}

	switch f.Type.Type {
	case field.TypeInt8, field.TypeUint8,
		field.TypeInt16, field.TypeUint16:
		return fmt.Sprintf("int32(%s)", expr)
	case field.TypeUint64, field.TypeUint32, field.TypeUint:
		return fmt.Sprintf("int64(%s)", expr) // TODO: possibly losing information here for uint64
	case field.TypeJSON, field.TypeOther:
		if oasAnnotations, ok := f.Annotations["EntOAS"]; ok {
			if schema, ok := oasAnnotations.(map[string]interface{})["Schema"]; ok && schema != nil {
				s := schema.(map[string]interface{})
				if s["type"] == "array" {
					items := s["items"].(map[string]interface{})
					if _, ok := items["enum"]; ok {
						return fmt.Sprintf("*JsonConvert(%s, &[]%s%sItem{}).(*[]%s%sItem)", expr, schemaName, f.StructField(), schemaName, f.StructField())
					} else if items["type"] == "object" {
						return fmt.Sprintf("*JsonConvert(%s, &[]%s%sItem{}).(*[]%s%sItem)", expr, schemaName, f.StructField(), schemaName, f.StructField())
					} else {
						return fmt.Sprintf("*JsonConvert(%s, &[]%s{}).(*[]%s)", expr, items["type"].(string), items["type"].(string))
					}
				}
			}
		}

		return expr
	default:
		return expr
	}
}

type ExclusionAnnotation = exclusion.Annotation

type Annotation struct {
	ExclusionAnnotation
	entoas.Annotation
}

// fieldAnnotation returns the Annotation on the given gen.Field.
func fieldAnnotation(f *gen.Field) (*Annotation, error) {
	return annotation(f.Annotations)
}

// annotation decodes the Annotation from the given gen.Annotations.
func annotation(as gen.Annotations) (*Annotation, error) {
	ant := &Annotation{}
	if as != nil {
		if as[ant.Annotation.Name()] != nil {
			if err := ant.Annotation.Decode(as[ant.Annotation.Name()]); err != nil {
				return nil, err
			}
		}

		if as[ant.ExclusionAnnotation.Name()] != nil {
			if err := ant.ExclusionAnnotation.Decode(as[ant.ExclusionAnnotation.Name()]); err != nil {
				return nil, err
			}
		}
	}

	return ant, nil
}

func pagination(as gen.Annotations, name string) string {
	ant := &filter.Annotation{}
	if as != nil {
		if as[ant.Name()] != nil {
			if err := ant.Decode(as[ant.Name()]); err != nil {
				panic(err)
			}

		}
	}

	switch name {
	case "page":
		if ant.Page == nil {
			return "Page"
		} else {
			return toCamelCase(ant.Page.Name)
		}
	case "itemsPerPage":
		if ant.ItemsPerPage == nil {
			return "ItemsPerPage"
		} else {
			return toCamelCase(ant.ItemsPerPage.Name)
		}
	case "total":
		return toCamelCase(ant.ReturnTotal.Name)
	}
	return ""
}

func toCamelCase(s string) string {
	capitalizeNext := true
	var result []rune

	for _, v := range s {
		if unicode.IsSpace(v) || v == '_' || v == '-' {
			capitalizeNext = true
		} else if capitalizeNext {
			result = append(result, unicode.ToUpper(v))
			capitalizeNext = false
		} else {
			result = append(result, unicode.ToLower(v))
		}
	}

	return string(result)
}

func returnTotal(as gen.Annotations, nodeName string) bool {
	ant := &filter.Annotation{}
	if as != nil {
		if as[ant.Name()] != nil {
			if err := ant.Decode(as[ant.Name()]); err != nil {
				panic(err)
			}
		}
	}

	return ant.ReturnTotal != nil
}

func isComplexType(f *gen.Field) bool {
	switch f.Type.Type {
	case field.TypeJSON, field.TypeOther:
		if oasAnnotations, ok := f.Annotations["EntOAS"]; ok {
			if schema, ok := oasAnnotations.(map[string]interface{})["Schema"]; ok && schema != nil {
				s := schema.(map[string]interface{})
				if s["type"] == "array" {
					return true
				}
			}
		}
	}

	return false
}
