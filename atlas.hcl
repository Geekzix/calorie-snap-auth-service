data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./internal/domain",
    "--dialect", "postgres",
  ]
}
env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/16/dev"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

env "neon" {
  url = getenv("CALOIRE_SNAP_NEON_DB_URL")
  migration {
    dir              = "file://migrations"
    revisions_schema = "public"
  }
}
