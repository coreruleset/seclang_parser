# Copyright 2025 OWASP CRS Project
# SPDX-License-Identifier: Apache-2.0

class Configuration:
    def __init__(self, marker=None, directives=None):
        self.marker = marker
        self.directives = directives or []

class ConfigurationList:
    def __init__(self):
        self.configurations = []

    def add_configuration(self, configuration):
        self.configurations.append(configuration)
