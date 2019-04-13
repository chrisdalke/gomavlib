package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

var reMsgName = regexp.MustCompile("^[A-Z0-9_]+$")
var reTypeIsArray = regexp.MustCompile("^(.+?)\\[([0-9]+)\\]$")

var dialectTypeToGo = map[string]string{
	"double":   "float64",
	"uint64_t": "uint64",
	"int64_t":  "int64",
	"float":    "float32",
	"uint32_t": "uint32",
	"int32_t":  "int32",
	"uint16_t": "uint16",
	"int16_t":  "int16",
	"uint8_t":  "uint8",
	"int8_t":   "int8",
	"char":     "string",
}

func dialectFieldGoToDef(in string) string {
	re := regexp.MustCompile("([A-Z])")
	in = re.ReplaceAllString(in, "_${1}")
	return strings.ToLower(in[1:])
}

func dialectFieldDefToGo(in string) string {
	return dialectMsgDefToGo(in)
}

func dialectMsgDefToGo(in string) string {
	re := regexp.MustCompile("_[a-z]")
	in = strings.ToLower(in)
	in = re.ReplaceAllStringFunc(in, func(match string) string {
		return strings.ToUpper(match[1:2])
	})
	return strings.ToUpper(in[:1]) + in[1:]
}

func filterDesc(in string) string {
	return strings.Replace(in, "\n", "", -1)
}

type outEnumValue struct {
	Value       string
	Name        string
	Description string
}

type outEnum struct {
	Name        string
	Description string
	Values      []*outEnumValue
}

type outField struct {
	Description string
	Line        string
}

type outMessage struct {
	Name        string
	Description string
	Id          int
	Fields      []*outField
}

type outDefinition struct {
	Name     string
	Enums    []*outEnum
	Messages []*outMessage
}

var tpl = template.Must(template.New("").Parse(
	`// Autogenerated with dialgen, do not edit.
{{- if .Preamble }}
//
// {{ .Preamble }}
//
{{- end }}
package {{ .Name }}

import (
	"github.com/gswly/gomavlib"
)

// Dialect contains the dialect object that can be passed to the library.
var Dialect = dialect

var dialect = gomavlib.MustDialect({{.Version}}, []gomavlib.Message{
{{- range .Defs }}
    // {{ .Name }}
{{- range .Messages }}
    &Message{{ .Name }}{},
{{- end }}
{{- end }}
})

{{ range .Enums }}
// {{ .Description }}
type {{ .Name }} int

const (
{{- $pn := .Name }}
{{- range .Values }}
	// {{ .Description }}
	{{ .Name }} {{ $pn }} = {{ .Value }}
{{- end }}
)
{{ end }}

{{ range .Defs }}
// {{ .Name }}

{{ range .Messages }}
// {{ .Description }}
type Message{{ .Name }} struct {
{{- range .Fields }}
	// {{ .Description }}
    {{ .Line }}
{{- end }}
}

func (*Message{{ .Name }}) GetId() uint32 {
    return {{ .Id }}
}
{{ end }}
{{ end }}
`))

func main() {
	kingpin.CommandLine.Help = "Generate a Mavlink dialect library from a definition file.\n" +
		"Example: dialgen \\\n--output=dialect.go \\\nhttps://raw.githubusercontent.com/mavlink/mavlink/master/message_definitions/v1.0/common.xml"

	quiet := kingpin.Flag("quiet", "suppress info messages during execution.").Short('q').Bool()
	outfile := kingpin.Flag("output", "output file").Required().String()
	preamble := kingpin.Flag("preamble", "preamble comment").String()
	mainDefAddr := kingpin.Arg("xml", "a path or url pointing to a Mavlink dialect definition in XML format").Required().String()

	kingpin.Parse()

	err := do(*quiet, *outfile, *preamble, *mainDefAddr)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}

func do(quiet bool, outfile string, preamble string, mainDefAddr string) error {
	if strings.HasSuffix(outfile, ".go") == false {
		return fmt.Errorf("output file must end with .go")
	}

	version := ""
	defsProcessed := make(map[string]struct{})
	isRemote := func() bool {
		_, err := url.ParseRequestURI(mainDefAddr)
		return err == nil
	}()

	// parse all definitions recursively
	outDefs, err := definitionProcess(quiet, &version, defsProcessed, isRemote, mainDefAddr)
	if err != nil {
		return err
	}

	// merge enums together
	enums := make(map[string]*outEnum)
	for _, def := range outDefs {
		for _, defEnum := range def.Enums {
			if _, ok := enums[defEnum.Name]; !ok {
				enums[defEnum.Name] = &outEnum{
					Name:        defEnum.Name,
					Description: defEnum.Description,
				}
			}
			enum := enums[defEnum.Name]

			for _, v := range defEnum.Values {
				enum.Values = append(enum.Values, v)
			}
		}
	}

	// fill enum missing values
	for _, enum := range enums {
		nextVal := 0
		for _, v := range enum.Values {
			if v.Value != "" {
				nextVal, _ = strconv.Atoi(v.Value)
				nextVal++
			} else {
				v.Value = strconv.Itoa(nextVal)
				nextVal++
			}
		}
	}

	// create output folder
	dir, _ := filepath.Split(outfile)
	os.Mkdir(dir, 0755)

	// open file
	f, err := os.Create(outfile)
	if err != nil {
		return err
	}
	defer f.Close()

	// dump
	return tpl.Execute(f, map[string]interface{}{
		"Name": func() string {
			_, name := filepath.Split(mainDefAddr)
			return strings.TrimSuffix(name, ".xml")
		}(),
		"Preamble": preamble,
		"Version": func() int {
			ret, _ := strconv.Atoi(version)
			return ret
		}(),
		"Defs":  outDefs,
		"Enums": enums,
	})
}

func definitionProcess(quiet bool, version *string, defsProcessed map[string]struct{}, isRemote bool, defAddr string) ([]*outDefinition, error) {
	// skip already processed
	if _, ok := defsProcessed[defAddr]; ok {
		return nil, nil
	}
	defsProcessed[defAddr] = struct{}{}

	if quiet == false {
		fmt.Printf("definition %s\n", defAddr)
	}

	content, err := definitionGet(isRemote, defAddr)
	if err != nil {
		return nil, err
	}

	def, err := definitionDecode(content)
	if err != nil {
		return nil, fmt.Errorf("unable to decode: %s", err)
	}

	addrPath, addrName := filepath.Split(defAddr)

	var outDefs []*outDefinition

	// version
	if def.Version != "" {
		if *version != "" && *version != def.Version {
			return nil, fmt.Errorf("version defined twice (%s and %s)", def.Version, *version)
		}
		*version = def.Version
	}

	// includes
	for _, inc := range def.Includes {
		// prepend url to remote address
		if isRemote == true {
			inc = addrPath + inc
		}
		subDefs, err := definitionProcess(quiet, version, defsProcessed, isRemote, inc)
		if err != nil {
			return nil, err
		}
		outDefs = append(outDefs, subDefs...)
	}

	outDef := &outDefinition{
		Name: addrName,
	}

	// enums
	for _, enum := range def.Enums {
		oute := &outEnum{
			Name:        enum.Name,
			Description: filterDesc(enum.Description),
		}
		for _, val := range enum.Values {
			oute.Values = append(oute.Values, &outEnumValue{
				Value:       val.Value,
				Name:        val.Name,
				Description: filterDesc(val.Description),
			})
		}
		outDef.Enums = append(outDef.Enums, oute)
	}

	// messages
	for _, msg := range def.Messages {
		outMsg, err := messageProcess(msg)
		if err != nil {
			return nil, err
		}
		outDef.Messages = append(outDef.Messages, outMsg)
	}

	outDefs = append(outDefs, outDef)
	return outDefs, nil
}

func definitionGet(isRemote bool, defAddr string) ([]byte, error) {
	if isRemote == true {
		byt, err := urlDownload(defAddr)
		if err != nil {
			return nil, fmt.Errorf("unable to download: %s", err)
		}
		return byt, nil
	}

	byt, err := ioutil.ReadFile(defAddr)
	if err != nil {
		return nil, fmt.Errorf("unable to open: %s", err)
	}
	return byt, nil
}

func urlDownload(desturl string) ([]byte, error) {
	res, err := http.Get(desturl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad return code: %v", res.StatusCode)
	}

	byt, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return byt, nil
}

func messageProcess(msg *definitionMessage) (*outMessage, error) {
	if m := reMsgName.FindStringSubmatch(msg.Name); m == nil {
		return nil, fmt.Errorf("unsupported message name: %s", msg.Name)
	}

	outMsg := &outMessage{
		Name:        dialectMsgDefToGo(msg.Name),
		Description: filterDesc(msg.Description),
		Id:          msg.Id,
	}

	for _, f := range msg.Fields {
		outField, err := fieldProcess(f)
		if err != nil {
			return nil, err
		}
		outMsg.Fields = append(outMsg.Fields, outField)
	}

	return outMsg, nil
}

func fieldProcess(field *dialectField) (*outField, error) {
	outF := &outField{
		Description: filterDesc(field.Description),
	}
	tags := make(map[string]string)

	newname := dialectFieldDefToGo(field.Name)

	// name conversion is not univoque: add tag
	if dialectFieldGoToDef(newname) != field.Name {
		tags["mavname"] = field.Name
	}

	outF.Line += newname

	typ := field.Type
	arrayLen := ""

	if typ == "uint8_t_mavlink_version" {
		typ = "uint8_t"
	}

	// string or array
	if matches := reTypeIsArray.FindStringSubmatch(typ); matches != nil {
		// string
		if matches[1] == "char" {
			tags["mavlen"] = matches[2]
			typ = "char"
			// array
		} else {
			arrayLen = matches[2]
			typ = matches[1]
		}
	}

	// extension
	if field.Extension == true {
		tags["mavext"] = "true"
	}

	typ = dialectTypeToGo[typ]
	if typ == "" {
		return nil, fmt.Errorf("unknown type: %s", typ)
	}

	outF.Line += " "
	if arrayLen != "" {
		outF.Line += "[" + arrayLen + "]"
	}
	if field.Enum != "" {
		outF.Line += field.Enum
		tags["mavenum"] = typ
	} else {
		outF.Line += typ
	}

	if len(tags) > 0 {
		var tmp []string
		for k, v := range tags {
			tmp = append(tmp, fmt.Sprintf("%s:\"%s\"", k, v))
		}
		sort.Strings(tmp)
		outF.Line += " `" + strings.Join(tmp, " ") + "`"
	}
	return outF, nil
}
