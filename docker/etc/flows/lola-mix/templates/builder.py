#!/usr/bin/env python3

from random import uniform
import jinja2
import argparse
import os

def write_config(tmpl, conf_dir, label, port, at_points):
    env = jinja2.Environment(loader=jinja2.FileSystemLoader(searchpath="./"))
    output = env.get_template(tmpl).render(label=label, port=port, at_points=at_points, indent=4*" ")
    with open(os.path.join(conf_dir, label + ".yaml"), "w") as fh:
        fh.write(output)

def make_at_points(num_at_points, distance):
    # Pick a random point in [0, distance-1] (with decisecond precision)
    t0 = float("{:.1f}".format(uniform(0, distance - 1)))

    # Place as many points as requested at distance "distance" starting at t0
    return [t0 + offset for offset in range(0, num_at_points * distance, distance)]

def main():
    parser = argparse.ArgumentParser(description="create trafic configuration files from a template")
    parser.add_argument("--base_port", required=True, type=int, help="base port")
    parser.add_argument("--dest_dir", required=True, type=str, help="folder where to store the produced configuration files")
    parser.add_argument("--num_files", required=True, type=int, help="number of files to create")
    parser.add_argument("--label", required=True, type=str, help="a label for the traffic flow (also the basename for the produced file)")
    parser.add_argument("--template", required=True, type=str, help="the template file name")
    parser.add_argument("--num_at_points", required=True, type=int, help="number of 'at' points to produce")
    parser.add_argument("--at_points_distance", required=True, type=int, help="distance (in seconds) between 'at' points")

    args = parser.parse_args()

    for i in range(0, args.num_files * 2, 2):
        port = args.base_port + i
        label = "%s-%d" % (args.label, port)
        at_points = make_at_points(args.num_at_points, args.at_points_distance)

        write_config(args.template, args.dest_dir, label, port, at_points)

if __name__ == "__main__":
    main()
