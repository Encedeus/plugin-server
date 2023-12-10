server {
  host = "localhost"
  port = 3001
}

database {
  host     = "localhost"
  port     = 5678
  user     = "postgres"
  name     = "PluginServerDB"
  password = "root"
}

auth {
  jwt_secret_access  = "i95FOB61kCoJjSt2SBSifhtwMHQ7Nasi"
  jwt_secret_refresh = "SjSt2fhtwi7BiFOS95MHQiasB61kCoJN"
}

smtp {
  host     = "smtp.gmail.com"
  port     = 587
  name     = "noreply.encedeus@gmail.com"
  password = "fxzvqjqtxgfjmceh"
}

storage {
  dir = "./pfp"
}

validation {
  max_email_len = 32

  max_name_len = 32
  min_name_len = 3

  max_pass_len = 64
  min_pass_len = 8

  max_plugin_name_len = 32
  min_plugin_name_len = 2

  max_release_name_len = 25
  min_release_name_len = 2
}

