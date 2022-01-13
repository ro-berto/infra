// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

module.exports = (config) => {
    config.set({
        frameworks: ['mocha', 'chai'],
        files: [
            'src/**/*.test.ts',
        ],
        preprocessors: {
            '**/*.ts': ['esbuild'],
        },
        reporters: ['mocha'],
        browsers: ['ChromeHeadless'],
    });
};
