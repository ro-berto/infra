#!/usr/bin/env python3
# Copyright 2022 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

import argparse
import json
import os
import urllib.request

def do_latest():
  print(
      json.load(
          urllib.request.urlopen(
              'https://api.github.com/repos/xpack-dev-tools/openocd-xpack/'
              'releases/latest'
          ))['tag_name'].lstrip('v'))


_PLATFORMS = {
    'linux-amd64': 'linux-x64',
    'linux-arm64': 'linux-arm64',
    'mac-amd64': 'darwin-x64',
    'mac-arm64': 'darwin-arm64',
    'windows-amd64': 'win32-x64',
}


def get_download_url(version, platform):
  if platform not in _PLATFORMS:
    raise ValueError('unsupported platform {}'.format(platform))

  extension = '.zip' if 'windows' in platform else '.tar.gz'

  url = (
      'https://github.com/xpack-dev-tools/openocd-xpack/releases/download/'
      'v{version}/xpack-openocd-{version}-{platform}{extension}').format(
          version=version,
          platform=_PLATFORMS[platform],
          extension=extension,
      )

  manifest = {
      'url': [url],
      'ext': extension,
  }

  print(json.dumps(manifest))


def main():
  ap = argparse.ArgumentParser()
  sub = ap.add_subparsers()

  latest = sub.add_parser("latest")
  latest.set_defaults(func=lambda _opts: do_latest())

  download = sub.add_parser("get_url")
  download.set_defaults(
    func=lambda opts: get_download_url(
      os.environ['_3PP_VERSION'], os.environ['_3PP_PLATFORM']))

  opts = ap.parse_args()
  opts.func(opts)


if __name__ == '__main__':
  main()
