# YAML Schemas Definitions

## Manifest

- **version** (string, required): Version of the object.
- **name** (string, required): Name of the object.
- **profiles** (array of strings, required): List of profiles.
- **upgrade** (string): Upgrade status.
- **hardware_requirements** (object): Hardware requirements, including:
  - **min_cpu_cores** (integer, required, >=0): Minimum CPU cores.
  - **min_ram** (integer, required, >=0): Minimum RAM.
  - **min_free_space** (integer, required, >=0): Minimum free space.
  - **stop_if_requirements_are_not_met** (boolean, required): Flag to stop if requirements aren't met.
- **plugin** (object): Plugin details, including:
  - **image** (string): Plugin image.
- _No additional properties are allowed_

## Profile

- **name** (string): Profile name.
- **monitoring** (object, required): Monitoring details, including:
  - **targets** (array of objects, required): List of targets, each with:
    - **service** (string, required): Name of the docker-compose service
    - **port** (integer, required 1 <= port <= 65535): Port serving the metrics
    - **path** (string, required): Metrics path
- **hardware_requirements_overrides** (object): Overrides of the Manifest's hardware requirements, including:
  - **min_cpu_cores** (integer, required, >=0): Minimum CPU cores.
  - **min_ram** (integer, required, >=0): Minimum RAM.
  - **min_free_space** (integer, required, >=0): Minimum free space.
  - **stop_if_requirements_are_not_met** (boolean, required): Flag to stop if requirements aren't met.
- **plugin_overrides** (object): Overrides of the Manifest's plugin details, including:
  - **image** (string, required): Pre-built docker image name ready to be pulled.
- **options** (array of objects): List of options, each with:
  - **name** (string, required): Option name.
  - **target** (string, required): Option target.
  - **type** (string, required): Option type.
  - **default** (any): Default value.
  - **help** (string, required): Help text.
  - **validate** (object): Validation rules, including re2_regex, format, uri_scheme, min_value, max_value, and options.
- **api** (object): AVS Node API details, including:
  - **service** (string, required): Name of the docker-compose service exposing the API.
  - **port** (integer, required 1 <= port <= 65535): Port serving the API.
- _No additional properties are allowed._
