package mixer

import (
	"fmt"
	"os"
)

// TODO put basedir+filename together into a outfile
func doWriteConf(outFile string, defaultTmpl string, g GlobalDesc, c FlowDesc) error {
	tmpl, err := setupTemplate(c, defaultTmpl)
	if err != nil {
		return err
	}

	props, err := evalProps(g, c)
	if err != nil {
		return err
	}

	out, err := os.Create(outFile)
	if err != nil {
		return fmt.Errorf("cannot create %s: %v", outFile, err)
	}
	defer out.Close()

	err = tmpl.Execute(out, props)
	if err != nil {
		return fmt.Errorf("cannot execute template: %v", err)
	}

	return nil
}

func evalInstances(total Bytes, percent Ratio) uint {
	return uint((float64(total.Val) * 8 * percent.Val) / 64000)
}

func evalProps(g GlobalDesc, c FlowDesc) (map[string]string, error) {
	p := make(map[string]string)

	for k, v := range c.Props {
		p[k] = v
	}

	p["port"] = fmt.Sprintf("%d", c.PortsRange.First)
	p["instances"] = fmt.Sprintf("%d", evalInstances(g.TotalBandwidth, c.PercentBandwidth))
	p["time"] = fmt.Sprintf("%fs", g.TotalTime.Seconds())
	p["report_interval"] = fmt.Sprintf("%fs", g.ReportInterval.Seconds())

	return p, nil
}
