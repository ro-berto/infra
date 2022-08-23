create {
  platform_re: "linux-(amd64|arm64)"

  source {
    git {
      repo: "https://sourceware.org/git/valgrind.git"
      tag_pattern: "VALGRIND_%s"
      version_join: "_"
    }
  }
  build {}
}

create {
  platform_re: "linux-.*"
}

upload { pkg_prefix: "tools" }