package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"text/template"

	// Embed is used to import the template files
	_ "embed"

	"github.com/Masterminds/sprig/v3"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"
	_ "github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/postgres/walker"
	"github.com/stackrox/rox/pkg/stringutils"
	"github.com/stackrox/rox/pkg/utils"
	"github.com/stackrox/rox/tools/generate-helpers/common/packagenames"
	"golang.org/x/tools/imports"
)

//go:embed store.go.tpl
var storeFile string

//go:embed store_test.go.tpl
var storeTestFile string

//go:embed index.go.tpl
var indexFile string

var (
	storeTemplate     = template.Must(template.New("gen").Funcs(funcMap).Funcs(sprig.TxtFuncMap()).Parse(autogenerated + storeFile))
	storeTestTemplate = template.Must(template.New("gen").Funcs(funcMap).Funcs(sprig.TxtFuncMap()).Parse(autogenerated + storeTestFile))
	indexTemplate     = template.Must(template.New("gen").Funcs(funcMap).Funcs(sprig.TxtFuncMap()).Parse(autogenerated + indexFile))
)

type properties struct {
	Type           string
	TrimmedType    string
	Table          string
	RegisteredType string

	SearchCategory string
	ObjectPathName string
	Singular       string
	WriteOptions   bool
	OptionsPath    string

	// Refs indicate the additional referentiol relationships. Each string is <table_name>:<proto_type>.
	// These are non-embedding relations, that is, this table is not embedded into referenced table to
	// construct the proto message.
	Refs []string

	// When set to true, it means that the schema represents a join table. The generation of mutating functions
	// such as inserts, updates, deletes, is skipped. This is because join tables should be filled from parents.
	JoinTable bool
}

func renderFile(templateMap map[string]interface{}, temp *template.Template, templateFileName string) error {
	buf := bytes.NewBuffer(nil)
	if err := temp.Execute(buf, templateMap); err != nil {
		return err
	}
	file := buf.Bytes()

	formatted, err := imports.Process(templateFileName, file, nil)
	if err != nil {
		return err
	}
	if err := os.WriteFile(templateFileName, formatted, 0644); err != nil {
		return err
	}
	return nil
}

func main() {
	c := &cobra.Command{
		Use: "generate store implementations",
	}

	var props properties
	c.Flags().StringVar(&props.Type, "type", "", "the (Go) name of the object")
	utils.Must(c.MarkFlagRequired("type"))

	c.Flags().StringVar(&props.RegisteredType, "registered-type", "", "the type this is registered in proto as storage.X")

	c.Flags().StringVar(&props.Table, "table", "", "the logical table of the objects")
	utils.Must(c.MarkFlagRequired("table"))

	c.Flags().StringVar(&props.Singular, "singular", "", "the singular name of the object")
	c.Flags().StringVar(&props.OptionsPath, "options-path", "/index/mappings", "path to write out the options to")
	c.Flags().StringVar(&props.SearchCategory, "search-category", "", "the search category to index under")
	c.Flags().StringSliceVar(&props.Refs, "references", []string{}, "additional foreign key references as <table_name:type>")
	c.Flags().BoolVar(&props.JoinTable, "join-table", false, "indicates the schema represents a join table. The generation of mutating functions is skipped")

	c.RunE = func(*cobra.Command, []string) error {
		typ := stringutils.OrDefault(props.RegisteredType, props.Type)
		fmt.Println("Generating for", typ)
		mt := proto.MessageType(typ)
		if mt == nil {
			log.Fatalf("could not find message for type: %s", typ)
		}

		schema := walker.Walk(mt, props.Table)
		if schema.NoPrimaryKey() {
			log.Fatal("No primary key defined, please check relevant proto file and ensure a primary key is specified using the \"sql:\"pk\"\" tag")
		}

		compileFKArgAndAttachToSchema(schema, props.Refs)

		templateMap := map[string]interface{}{
			"Type":           props.Type,
			"TrimmedType":    stringutils.GetAfter(props.Type, "."),
			"Table":          props.Table,
			"Schema":         schema,
			"SearchCategory": fmt.Sprintf("SearchCategory_%s", props.SearchCategory),
			"OptionsPath":    path.Join(packagenames.Rox, props.OptionsPath),
			"JoinTable":      props.JoinTable,
		}

		if err := renderFile(templateMap, storeTemplate, "store.go"); err != nil {
			return err
		}
		if err := renderFile(templateMap, storeTestTemplate, "store_test.go"); err != nil {
			return err
		}
		if props.SearchCategory != "" {
			if err := renderFile(templateMap, indexTemplate, "index.go"); err != nil {
				return err
			}
		}

		return nil
	}
	if err := c.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
