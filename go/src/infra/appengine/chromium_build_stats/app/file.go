// Copyright 2014 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

// file.go provides /file endpoints.

import (
	"io"
	"net/http"
	"strings"

	"cloud.google.com/go/storage"
	"google.golang.org/appengine/v2/log"
	"google.golang.org/appengine/v2/user"

	"infra/appengine/chromium_build_stats/logstore"
)

func init() {
	http.Handle("/file/", http.StripPrefix("/file/", http.HandlerFunc(fileHandler)))
}

// fileHandler handles /<path> to access gs://chrome-goma-log/<path>.
func fileHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	user := user.Current(ctx)
	if user == nil {
		http.Error(w, "Login required", http.StatusUnauthorized)
		return
	}
	if !strings.HasSuffix(user.Email, "@google.com") {
		http.Error(w, "Unauthorized to access", http.StatusUnauthorized)
		return
	}

	path := req.URL.Path

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "failed to create storage client: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Close()

	reader, err := logstore.Fetch(ctx, client, path)
	if err != nil {
		log.Errorf(ctx, "failed to fetch %s: %v", path, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer reader.Close()

	w.Header().Set("Content-Type", reader.Attrs.ContentType)
	w.Header().Set("Content-Encoding", reader.Attrs.ContentEncoding)

	_, err = io.Copy(w, reader)
	if err != nil {
		log.Errorf(ctx, "failed to copy %s: %v", path, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
