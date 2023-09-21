data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./src/consumer/repositories",
    "--dialect", "postgres",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/15"
  migration {
    dir = "file://./src/consumer/migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}