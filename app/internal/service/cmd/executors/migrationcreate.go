package executors

import (
	"bytes"
	"context"
	"fmt"
	mainarg "github.com/dzamyatin/atomWebsite/internal/service/arg"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"
	"text/template"
	"time"
)

type ArgMigrationCreate struct {
	mainarg.Arg
	Name string        `arg:"-n,required" help:"Name of migration"`
	Type typeMigration `arg:"-t,required" help:"Type of migration (sql, go)"`
}

type typeMigration string

const (
	typeMigrationGo  = typeMigration("go")
	typeMigrationSQL = typeMigration("sql")
)

type MigrationCreateCommand struct {
	logger *zap.Logger
}

func NewMigrationCreateCommand(logger *zap.Logger) *MigrationCreateCommand {
	return &MigrationCreateCommand{logger: logger}
}

const templateSql = `-- +goose Up
SELECT 1;
-- +goose Down
SELECT 1;
`

const templateGo = `package migration

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(Up, Down)
}

func Up(_ context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("SELECT 1;")
	if err != nil {
		return err
	}
	return nil
}

func Down(_ context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("SELECT 1;")
	if err != nil {
		return err
	}
	return nil
}
`

func (m *MigrationCreateCommand) Execute(ctx context.Context, u ArgMigrationCreate) error {
	fileName := fmt.Sprintf(
		"%v_%v",
		time.Now().Unix(),
		strings.Replace(strings.ToLower(u.Name),
			" ",
			"_",
			-1,
		),
	)

	var err error

	switch u.Type {
	case typeMigrationGo:
		err = m.generateGoTemplate(fileName)
	case typeMigrationSQL:
		fallthrough
	default:
		err = m.generateSQLTemplate(fileName)
	}

	return errors.Wrap(err, "generate template")
}

func (m *MigrationCreateCommand) generateGoTemplate(fileName string) error {

	f, err := os.Create("internal/migration/" + fileName + ".go")

	if err != nil {
		return errors.Wrap(err, "error creating migration file")
	}

	defer f.Close()

	tpl, err := template.New("").Parse(templateGo)

	if err != nil {
		return errors.Wrap(err, "error parsing template")
	}

	buf := bytes.NewBufferString("")

	err = tpl.Execute(buf, nil)

	if err != nil {
		return errors.Wrap(err, "error executing template")
	}

	fset := token.NewFileSet()
	a, err := parser.ParseFile(fset, "", buf.Bytes(), 0)

	if err != nil {
		return errors.Wrap(err, "error parsing file")
	}

	for _, d := range a.Decls {
		funcDec, ok := d.(*ast.FuncDecl)
		if !ok {
			continue
		}

		if funcDec.Name.Name == "Up" {
			funcDec.Name.Name = funcDec.Name.Name + fileName
			continue
		}

		if funcDec.Name.Name == "Down" {
			funcDec.Name.Name = funcDec.Name.Name + fileName
			continue
		}

		if funcDec.Name.Name == "init" {
			for _, initSmtp := range funcDec.Body.List {
				expr, ok := initSmtp.(*ast.ExprStmt)
				if !ok {
					continue
				}

				callExpr, ok := expr.X.(*ast.CallExpr)
				if !ok {
					continue
				}

				for _, arg := range callExpr.Args {
					a, ok := arg.(*ast.Ident)
					if !ok {
						continue
					}
					if a.Name == "Up" {
						a.Name = a.Name + fileName
					}
					if a.Name == "Down" {
						a.Name = a.Name + fileName
					}
				}
			}
		}
	}

	err = printer.Fprint(f, token.NewFileSet(), a)

	return nil
}

func (m *MigrationCreateCommand) generateSQLTemplate(fileName string) error {
	f, err := os.Create("internal/migration/sql/" + fileName + ".sql")

	if err != nil {
		return errors.Wrap(err, "error creating migration file")
	}

	defer f.Close()

	_, err = f.WriteString(templateSql)

	return errors.Wrap(err, "error executing template")
}
