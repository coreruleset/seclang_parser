"""Command-line interface for the CRS parser."""

import argparse
from .parser import parse_file


def parse_args():
    """Parse our command line arguments"""
    cmdline = argparse.ArgumentParser(description="CRS parser")
    cmdline.add_argument(
        "-f",
        "--files",
        metavar="FILE",
        required=True,
        nargs="*",
        help="files to read, if empty, stdin is used",
    )
    return cmdline.parse_args()

def cli(args: list[str] = None):
    """ Receives a list of files to parse """
    for file in args.files:
        with open(file, "r", encoding="utf-8") as f:
            parse_file(f.read())

    return True


def run():
    """Runs the example parser"""
    args = parse_args()
    return cli(args)


if __name__ == "__main__":
    run()

