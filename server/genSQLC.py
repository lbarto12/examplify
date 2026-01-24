import os
import sys
import subprocess

SQLC_YAML_ENTRY = '''
  - engine: "postgresql"
    queries: "{queries}/{queryset}"
    schema: "{migrations}"
    gen:
      go:
        out: "{output}"
'''

SQLC_DEFAULT_OUTPUT = 'sqlc/sqlgen'

def generate(queries_path: str, schema_path: str):

    # generate sqlc.yaml
    query_dirs = os.listdir(queries_path)
    
    print("Found the following Query Domains:", end='')
    print("", *query_dirs, sep="\n- ")

    print("Generating sqlc.yaml", end='')

    sqlcyaml = '''
version: "2"
sql:'''
    for queryset in query_dirs:
        sqlcyaml += SQLC_YAML_ENTRY.format(
            queries=queries_path,
            queryset=queryset,
            migrations=schema_path,
            output=SQLC_DEFAULT_OUTPUT,
        )

    with open('sqlc.yaml', 'w') as cfg_file:
        cfg_file.write(sqlcyaml)

    print("✅")

    try:
        print("Generating SQLC queries", end=' ')
        subprocess.run([
            "sqlc", "generate"
        ])
    except subprocess.CalledProcessError as e:
        print("❌")
        sys.exit(1)
    print("✅")


if __name__ == '__main__':
    print("\033[93m", "\033[1m", "\033[4m", "Generating SQLC Queries", "\033[0m", sep='')
    generate(
        queries_path='sqlc/queries',
        schema_path='sqlc/migrations'
    )