# schema

This directory contains configuration to run [SchemaSpy](https://schemaspy.org/) to generate database documentation.

> [!IMPORTANT]
> Please update [schema.properties](./config/schema.properties) to match your current database setup.
> You can always override the connection parameters, also the schema, using the [SchemaSpy](https://schemaspy.org/) CLI flags.


As an example, when the MySQL server is running inside a Docker container using `--net-host`, you can do the following:

```console
docker run --rm \
  -it \
  --net=host \
  -v `pwd`/output:/output \
  -v `pwd`:/work schemaspy/schemaspy:6.2.4 \
  -configFile /work/config/schema.properties -imageformat png
```
