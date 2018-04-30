package mixer

import (
	"fmt"
	"os"
)

// TODO scavenger too uses this.  we might want to re-evaluate factorisation
func doWriteConf(outFile string, defaultTmpl string, g GlobalDesc, c FlowDesc, flowBitrate float64) error {
	tmpl, err := setupTemplate(c, defaultTmpl)
	if err != nil {
		return err
	}

	props, err := evalProps(g, c, flowBitrate)
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

func evalInstances(total Bytes, percent Ratio, flowBitrate float64) uint {
	return uint((float64(total.Val) * 8 * percent.Val) / flowBitrate)
}

func evalProps(g GlobalDesc, c FlowDesc, flowBitrate float64) (map[string]string, error) {
	p := make(map[string]string)

	for k, v := range c.Props {
		p[k] = v
	}

	p["port"] = fmt.Sprintf("%d", c.PortsRange.First)
	p["instances"] = fmt.Sprintf("%d", evalInstances(g.TotalBandwidth, c.PercentBandwidth, flowBitrate))
	p["time"] = fmt.Sprintf("%fs", g.TotalTime.Seconds())
	p["report_interval"] = fmt.Sprintf("%fs", g.ReportInterval.Seconds())

	return p, nil
}
