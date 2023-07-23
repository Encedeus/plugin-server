server {
  host = "localhost"
  port = 3001
}

database {
  host = "localhost"
  port = 54322
  user = "postgres"
  name = "PluginServerDB"
  password = "root"
}

auth {
  jwt_secret_access = "i95FOB61kCoJjSt2SBSifhtwMHQ7Nasi"
  jwt_secret_refresh = "SjSt2fhtwi7BiFOS95MHQiasB61kCoJN"
}

smtp {
  host = "smtp.gmail.com"
  port = 587
}

cdn {
  dir = "./pfp"
}