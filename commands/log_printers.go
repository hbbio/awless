package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/fatih/color"
	"github.com/hbbio/awless/console"
	"github.com/hbbio/awless/logger"
	"github.com/hbbio/awless/template"
)

type logPrinter interface {
	print(*template.Execution) error
}

type fullLogPrinter struct {
	w io.Writer
}

func (p *fullLogPrinter) print(t *template.Execution) error {
	writeMultilineLogHeader(t, p.w)

	if t.Message != "" {
		fmt.Fprintf(p.w, "\t%s\n\n", t.Message)
	}

	for _, cmd := range t.CommandNodesIterator() {
		var status string
		if cmd.CmdErr != nil {
			status = renderRedFn("KO")
		} else {
			status = renderGreenFn("OK")
		}

		var line string
		if v, ok := cmd.CmdResult.(string); ok && v != "" {
			line = fmt.Sprintf("    %s\t%s\t[%s]", status, cmd.String(), v)
		} else {
			line = fmt.Sprintf("    %s\t%s", status, cmd.String())
		}

		fmt.Fprintln(p.w, line)
		logger.New("", 0, p.w).MultiLineError(cmd.Err())
	}
	return nil
}

type statLogPrinter struct {
	w io.Writer
}

func (p *statLogPrinter) print(t *template.Execution) error {
	writeLogHeader(t, p.w)

	if t.Message != "" {
		fmt.Fprintf(p.w, "\n\t%s\n", t.Message)
	}

	return nil
}

type shortLogPrinter struct {
	w io.Writer
}

func (p *shortLogPrinter) print(t *template.Execution) error {
	writeLogHeader(t, p.w)
	return nil
}

type rawJSONPrinter struct {
	w io.Writer
}

func (p *rawJSONPrinter) print(t *template.Execution) error {
	if err := json.NewEncoder(p.w).Encode(t); err != nil {
		return fmt.Errorf("json printer: %s", err)
	}
	return nil
}

type idOnlyPrinter struct {
	w io.Writer
}

func (p *idOnlyPrinter) print(t *template.Execution) error {
	fmt.Fprint(p.w, t.ID)
	return nil
}

func writeLogHeader(t *template.Execution, w io.Writer) {
	stats := t.Stats()

	fmt.Fprint(w, renderYellowFn(t.ID))
	if stats.KOCount == 0 {
		color.New(color.FgGreen).Fprint(w, " OK")
	} else {
		color.New(color.FgRed).Fprint(w, " KO")
	}

	fmt.Fprintf(w, " (%s ago)", console.HumanizeTime(t.Date()))

	if t.Author != "" {
		fmt.Fprintf(w, " by %s", renderBlueFn(t.Author))
	}
	if t.Profile != "" {
		fmt.Fprintf(w, " with profile %s", renderBlueFn(t.Profile))
	}
	if t.Locale != "" {
		fmt.Fprintf(w, " in %s", renderBlueFn(t.Locale))
	}
	if !template.IsRevertible(t.Template) {
		fmt.Fprintf(w, " (not revertible)")
	}
}

func writeMultilineLogHeader(t *template.Execution, w io.Writer) {
	color.New(color.FgYellow).Fprintf(w, "id %s", t.ID)
	if !template.IsRevertible(t.Template) {
		fmt.Fprintln(w, " (not revertible)")
	} else {
		fmt.Fprintln(w)
	}

	fmt.Fprintf(w, "Date: %s\n", t.Date().Format(time.RFC1123Z))
	if t.Author != "" {
		fmt.Fprintf(w, "Author: %s\n", t.Author)
	}
	if t.Profile != "" {
		fmt.Fprintf(w, "Profile: %s\n", t.Profile)
	}
	if t.Locale != "" {
		fmt.Fprintf(w, "Region: %s\n", t.Locale)
	}
	fmt.Fprintln(w)
}
