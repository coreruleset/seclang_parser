# Copyright 2025 OWASP CRS Project
# SPDX-License-Identifier: Apache-2.0

"""Comprehensive parser tests matching the Golang test suite."""

from pathlib import Path
from antlr4 import FileStream, CommonTokenStream, ParseTreeWalker
import pytest
import yaml

from seclang_parser.SecLangLexer import SecLangLexer
from seclang_parser.SecLangParser import SecLangParser
from seclang_parser.errors import CustomErrorListener
from seclang_parser.listener import TreeShapeListener, ParserResult


# Load test data from shared YAML file
TEST_DATA_PATH = Path(__file__).parent.parent / "testdata" / "tests.yaml"
with open(TEST_DATA_PATH, "r", encoding="utf-8") as f:
    _test_data = yaml.safe_load(f)

CRS_TEST_FILES = _test_data["crs_test_files"]
PLUGINS_TEST_FILES = _test_data["plugins_test_files"]
GENERIC_TESTS = _test_data["generic_tests"]

# Convert expected_result dictionaries to ParserResult objects for check_output_tests
CHECK_OUTPUT_TESTS = {}
for file_path, test_data in _test_data["check_output_tests"].items():
    result_data = test_data["expected_result"]
    CHECK_OUTPUT_TESTS[file_path] = {
        "error_count": test_data["error_count"],
        "comment": test_data["comment"],
        "expected_result": ParserResult(**result_data),
    }


def parse_file_for_test(file_path: str) -> tuple[TreeShapeListener, int]:
    """Parse a file and return the listener and error count."""
    input_stream = FileStream(file_path, encoding="utf-8")
    lexer = SecLangLexer(input_stream)

    lexer_errors = CustomErrorListener()
    lexer.removeErrorListeners()
    lexer.addErrorListener(lexer_errors)

    parser_errors = CustomErrorListener()
    stream = CommonTokenStream(lexer)
    parser = SecLangParser(stream)
    parser.removeErrorListeners()
    parser.addErrorListener(parser_errors)

    parser.buildParseTrees = True
    tree = parser.configuration()

    listener = TreeShapeListener()
    walker = ParseTreeWalker()
    walker.walk(listener, tree)

    total_errors = len(lexer_errors.errors) + len(parser_errors.errors)

    if lexer_errors.errors:
        print(f"Lexer {len(lexer_errors.errors)} errors found")
        print(f"First error: {lexer_errors.errors[0]}")
    if parser_errors.errors:
        print(f"Parser {len(parser_errors.errors)} errors found")
        print(f"First error: {parser_errors.errors[0]}")

    return listener, total_errors


@pytest.mark.parametrize("file_path,test_data", GENERIC_TESTS.items())
def test_seclang(file_path: str, test_data: dict):
    """Test generic SecLang configuration files."""
    listener, total_errors = parse_file_for_test(file_path)
    assert total_errors == test_data["error_count"], (
        f"Error count mismatch for file {file_path} -> we want: {test_data['comment']}"
    )


@pytest.mark.parametrize("file_path,test_data", CHECK_OUTPUT_TESTS.items())
def test_seclang_output(file_path: str, test_data: dict):
    """Test SecLang files with expected output validation."""
    listener, total_errors = parse_file_for_test(file_path)
    assert total_errors == test_data["error_count"], (
        f"Error count mismatch for file {file_path} -> we want: {test_data['comment']}"
    )
    assert listener.results == test_data["expected_result"], (
        f"Expected result mismatch for file {file_path} -> we want: {test_data['expected_result']}"
    )


@pytest.mark.parametrize("file_path", CRS_TEST_FILES)
def test_crs_lang(file_path: str):
    """Test CRS (Core Rule Set) configuration files."""
    listener, total_errors = parse_file_for_test(file_path)
    assert total_errors == 0, f"Error count mismatch for file {file_path} -> we want no errors"


@pytest.mark.parametrize("file_path", PLUGINS_TEST_FILES)
def test_plugins(file_path: str):
    """Test plugin configuration files."""
    listener, total_errors = parse_file_for_test(file_path)
    assert total_errors == 0, f"Error count mismatch for file {file_path} -> we want no errors"
