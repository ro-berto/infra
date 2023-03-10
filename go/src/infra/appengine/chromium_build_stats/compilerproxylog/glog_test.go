// Copyright 2014 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package compilerproxylog

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func timeAt(ts string) time.Time {
	t, err := time.Parse("2006/01/02 15:04:05.000000", ts)
	if err != nil {
		panic(err)
	}
	return t
}

func BenchmarkParseLogLine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ParseLogline([]byte("I0911 03:41:26.740641 27529 compiler_proxy.cc:1498] goma built revision bb0fd40be997d444c8e82b54ddd0361148ba1379@1408698434"))
		if err != nil {
			b.Fatalf("got error from ParseLogLine(...): %v", err)
		}
	}
}

func BenchmarkLogTimestamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, err := logTimestamp([]byte("0911 03:41:26.740641"))
		if err != nil {
			b.Fatalf("got error from logTimestamp(...): %v", err)
		}
	}
}

func TestGlogParser(t *testing.T) {
	logcontent := `Log file created at: 2014/09/11 03:41:26
Running on machine: usho.tok.corp.google.com
Log line format: [IWEF]mmdd hh:mm:ss.uuuuuu threadid file:line] msg
I0911 03:41:26.740263 27529 breakpad_linux.cc:51] initialized breakpad.
I0911 03:41:26.740641 27529 compiler_proxy.cc:1498] goma built revision bb0fd40be997d444c8e82b54ddd0361148ba1379@1408698434
I0911 03:41:26.740849 27529 compiler_proxy.cc:1508] goma flags:GOMA_API_KEY_FILE=/usr/local/google/home/ukai/.goma_api_key
GOMA_COMPILER_PROXY_DAEMON_MODE=true
GOMA_COMPILER_PROXY_HTTP_THREADS=3 (auto configured)
GOMA_COMPILER_PROXY_THREADS=12 (auto configured)
GOMA_MAX_INCLUDE_CACHE_SIZE=0 (auto configured)
GOMA_MAX_POOLED_INCLUDE_DIR_CACHE=128 (auto configured)
GOMA_USE_SSL=true
I0911 03:41:26.902179 27539 openssl_engine.cc:892] Certificate loaded from system: Certificate:
  Data:
  Version: 1 (0x0)
  Serial Number: 421 (0x1a5)
W0911 03:41:26.921205 27539 http_rpc.cc:454] compress failed: err=-5
I0911 03:41:26.997308 27539 openssl_engine.cc:769] We may reject if the domain is not listed in loaded CRLs.
E0911 03:42:26.123456 27539 compile_task.cc:3974] Task:1194 Try:0: Missing 5 input files.
F0911 03:42:27.789012 27539 compiler_info.cc:2128] Unknown compiler type
`
	rd := strings.NewReader(logcontent)
	gp, err := NewGlogParser(rd)
	if err != nil {
		t.Fatalf("NewGlogParser()=%v, %v; want=_, <nil>", gp, err)
	}
	if got, want := gp.Created.Format("2006-01-02T15:04:05"), "2014-09-11T03:41:26"; got != want {
		t.Errorf("gp.Created=%s (%v); want=%s", got, gp.Created, want)
	}
	if got, want := gp.Machine, "usho.tok.corp.google.com"; got != want {
		t.Errorf("gp.Machine=%q; want=%q", got, want)
	}
	var got []Logline
	for gp.Next() {
		got = append(got, gp.Logline())
	}
	if err := gp.Err(); err != nil {
		t.Fatalf("gp.Err()=%v; want=<nil>", err)
	}

	want := []Logline{
		{
			Level:     Info,
			Timestamp: timeAt("2014/09/11 03:41:26.740263"),
			ThreadID:  "27529",
			Lines:     []string{"breakpad_linux.cc:51] initialized breakpad."},
		},
		{
			Level:     Info,
			Timestamp: timeAt("2014/09/11 03:41:26.740641"),
			ThreadID:  "27529",
			Lines:     []string{"compiler_proxy.cc:1498] goma built revision bb0fd40be997d444c8e82b54ddd0361148ba1379@1408698434"},
		},
		{
			Level:     Info,
			Timestamp: timeAt("2014/09/11 03:41:26.740849"),
			ThreadID:  "27529",
			Lines: []string{
				"compiler_proxy.cc:1508] goma flags:GOMA_API_KEY_FILE=/usr/local/google/home/ukai/.goma_api_key",
				"GOMA_COMPILER_PROXY_DAEMON_MODE=true",
				"GOMA_COMPILER_PROXY_HTTP_THREADS=3 (auto configured)",
				"GOMA_COMPILER_PROXY_THREADS=12 (auto configured)",
				"GOMA_MAX_INCLUDE_CACHE_SIZE=0 (auto configured)",
				"GOMA_MAX_POOLED_INCLUDE_DIR_CACHE=128 (auto configured)",
				"GOMA_USE_SSL=true",
			},
		},
		{
			Level:     Info,
			Timestamp: timeAt("2014/09/11 03:41:26.902179"),
			ThreadID:  "27539",
			Lines: []string{
				"openssl_engine.cc:892] Certificate loaded from system: Certificate:",
				"  Data:",
				"  Version: 1 (0x0)",
				"  Serial Number: 421 (0x1a5)",
			},
		},
		{
			Level:     Warning,
			Timestamp: timeAt("2014/09/11 03:41:26.921205"),
			ThreadID:  "27539",
			Lines:     []string{"http_rpc.cc:454] compress failed: err=-5"},
		},
		{
			Level:     Info,
			Timestamp: timeAt("2014/09/11 03:41:26.997308"),
			ThreadID:  "27539",
			Lines:     []string{"openssl_engine.cc:769] We may reject if the domain is not listed in loaded CRLs."},
		},
		{
			Level:     Error,
			Timestamp: timeAt("2014/09/11 03:42:26.123456"),
			ThreadID:  "27539",
			Lines:     []string{"compile_task.cc:3974] Task:1194 Try:0: Missing 5 input files."},
		},
		{
			Level:     Fatal,
			Timestamp: timeAt("2014/09/11 03:42:27.789012"),
			ThreadID:  "27539",
			Lines:     []string{"compiler_info.cc:2128] Unknown compiler type"},
		},
	}

	if len(got) != len(want) {
		t.Errorf("GlogParser=%v; want=%v", got, want)
		t.Fatalf("GlogParser=%d; want=%d", len(got), len(want))
	}
	for i, gi := range got {
		wi := want[i]
		if !reflect.DeepEqual(gi, wi) {
			t.Errorf("%d: got=%v; want=%v", i, gi, wi)
		}
	}
}

func TestGlogParser_FullYear(t *testing.T) {
	logcontent := `Log file created at: 2021/10/24 19:13:18
Running on machine: build309-m9.golo.chromium.org
Running duration (h:mm:ss): 0:00:00
Log line format: [IWEF]yyyymmdd hh:mm:ss.uuuuuu threadid file:line] msg
I20211024 19:13:18.766117 317894144 goma_init.cc:89] google-internal goma client
I20211024 19:13:18.767921 317894144 goma_init.cc:112] goma flags:GOMA_BURST_MAX_SUBPROCS=8 (autoconfigured)
GOMA_BURST_MAX_SUBPROCS_HEAVY=2 (auto configured)
GOMA_BURST_MAX_SUBPROCS_LOW=8 (auto configured)
GOMA_CACHE_DIR=/opt/s/w/ir/cache/goma/data/Chromium_iOS_Goma_RBE_ToT
GOMA_COMPILER_INFO_POOL=2 (auto configured)
GOMA_COMPILER_PROXY_DAEMON_MODE=true
GOMA_COMPILER_PROXY_HTTP_THREADS=1 (auto configured)
GOMA_COMPILER_PROXY_LOCK_FILENAME=goma_compiler_proxy.lock
GOMA_COMPILER_PROXY_PORT=8088
GOMA_COMPILER_PROXY_SOCKET_NAME=goma.ipc
GOMA_COMPILER_PROXY_THREADS=4 (auto configured)
GOMA_DEPS_CACHE_FILE=goma_deps_cache
GOMA_DUMP_COUNTERZ_FILE=/opt/s/w/ir/x/t/goma_counterz
GOMA_DUMP_STATS_FILE=/opt/s/w/ir/x/t/goma_stats
GOMA_ENABLE_COUNTERZ=true
GOMA_FAIL_FAST=true
GOMA_HERMETIC=error
GOMA_INCLUDE_PROCESSOR_THREADS=4 (auto configured)
GOMA_LOG_CLEAN_INTERVAL=86400
GOMA_MAX_SUBPROCS=2 (auto configured)
GOMA_MAX_SUBPROCS_LOW=1 (auto configured)
GOMA_PING_TIMEOUT_SEC=60
GOMA_RPC_EXTRA_PARAMS=?tot
GOMA_SERVER_HOST=staging-goma.chromium.org
GOMA_TMP_DIR=/var/folders/px/zrfs83fx2db3xsp30wm25fz00000gm/T/goma_chrome-bot
GOMA_USE_SSL=true
I20211024 19:13:18.768012 317894144 goma_init.cc:0] don't panic on following line
Item updated count = 0
`
	rd := strings.NewReader(logcontent)
	gp, err := NewGlogParser(rd)
	if err != nil {
		t.Fatalf("NewGlogParser()=%v, %v; want=_, <nil>", gp, err)
	}
	if got, want := gp.Created.Format("2006-01-02T15:04:05"), "2021-10-24T19:13:18"; got != want {
		t.Errorf("gp.Created=%s (%v); want=%s", got, gp.Created, want)
	}
	if got, want := gp.Machine, "build309-m9.golo.chromium.org"; got != want {
		t.Errorf("gp.Machine=%q; want=%q", got, want)
	}
	var got []Logline
	for gp.Next() {
		got = append(got, gp.Logline())
	}
	if err := gp.Err(); err != nil {
		t.Fatalf("gp.Err()=%v; want=<nil>", err)
	}

	want := []Logline{
		{
			Level:     Info,
			Timestamp: timeAt("2021/10/24 19:13:18.766117"),
			ThreadID:  "317894144",
			Lines:     []string{"goma_init.cc:89] google-internal goma client"},
		},
		{
			Level:     Info,
			Timestamp: timeAt("2021/10/24 19:13:18.767921"),
			ThreadID:  "317894144",
			Lines: []string{
				"goma_init.cc:112] goma flags:GOMA_BURST_MAX_SUBPROCS=8 (autoconfigured)",
				"GOMA_BURST_MAX_SUBPROCS_HEAVY=2 (auto configured)",
				"GOMA_BURST_MAX_SUBPROCS_LOW=8 (auto configured)",
				"GOMA_CACHE_DIR=/opt/s/w/ir/cache/goma/data/Chromium_iOS_Goma_RBE_ToT",
				"GOMA_COMPILER_INFO_POOL=2 (auto configured)",
				"GOMA_COMPILER_PROXY_DAEMON_MODE=true",
				"GOMA_COMPILER_PROXY_HTTP_THREADS=1 (auto configured)",
				"GOMA_COMPILER_PROXY_LOCK_FILENAME=goma_compiler_proxy.lock",
				"GOMA_COMPILER_PROXY_PORT=8088",
				"GOMA_COMPILER_PROXY_SOCKET_NAME=goma.ipc",
				"GOMA_COMPILER_PROXY_THREADS=4 (auto configured)",
				"GOMA_DEPS_CACHE_FILE=goma_deps_cache",
				"GOMA_DUMP_COUNTERZ_FILE=/opt/s/w/ir/x/t/goma_counterz",
				"GOMA_DUMP_STATS_FILE=/opt/s/w/ir/x/t/goma_stats",
				"GOMA_ENABLE_COUNTERZ=true",
				"GOMA_FAIL_FAST=true",
				"GOMA_HERMETIC=error",
				"GOMA_INCLUDE_PROCESSOR_THREADS=4 (auto configured)",
				"GOMA_LOG_CLEAN_INTERVAL=86400",
				"GOMA_MAX_SUBPROCS=2 (auto configured)",
				"GOMA_MAX_SUBPROCS_LOW=1 (auto configured)",
				"GOMA_PING_TIMEOUT_SEC=60",
				"GOMA_RPC_EXTRA_PARAMS=?tot",
				"GOMA_SERVER_HOST=staging-goma.chromium.org",
				"GOMA_TMP_DIR=/var/folders/px/zrfs83fx2db3xsp30wm25fz00000gm/T/goma_chrome-bot",
				"GOMA_USE_SSL=true",
			},
		},
		{
			Level:     Info,
			Timestamp: timeAt("2021/10/24 19:13:18.768012"),
			ThreadID:  "317894144",
			Lines: []string{
				"goma_init.cc:0] don't panic on following line",
				"Item updated count = 0",
			},
		},
	}
	if len(got) != len(want) {
		t.Errorf("GlogParser=%v; want=%v", got, want)
		t.Fatalf("GlogParser=%d; want=%d", len(got), len(want))
	}
	for i, gi := range got {
		wi := want[i]
		if !reflect.DeepEqual(gi, wi) {
			t.Errorf("%d: got=%v; want=%v", i, gi, wi)
		}
	}
}
