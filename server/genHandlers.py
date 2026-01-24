import os
import sys
import subprocess
import tempfile

CFG_YAML_BUILDER = '''
package: {}
generate:
  chi-server: true
  models: true

output-options:
  skip-prune: true
'''

def generate(api_spec_path: str, output_dir: str, client_output_dir: str, client_driver_dir: str):
    domains = [i.replace(".yaml", "") for i in os.listdir(api_spec_path)]
    if len(domains) == 0:
        print("No input files found.")
        sys.exit(1)
    
    for domain in domains:
        print(f"Generating {domain.upper()}", end=" ")

        # Create temp config
        cfg = CFG_YAML_BUILDER.format(f"gen{domain}")
        cfg_filepath = ""
        with tempfile.NamedTemporaryFile(delete=False, suffix=".yaml") as cfg_temp_file:
            cfg_temp_file.write(cfg.encode("utf-8"))
            cfg_filepath = cfg_temp_file.name

        try:
            # Generate server

            output_api_path: str = output_dir.format(domain)    
            output_dirname: str = os.path.dirname(output_api_path)
            os.makedirs(output_dirname, exist_ok=True)

            subprocess.run([
                "oapi-codegen", 
                "--config", cfg_filepath, 
                "-o", f"{output_dirname}/api.go", f"{api_spec_path}/{domain}.yaml"
            ], check=True) 

        except subprocess.CalledProcessError as e:
            print("❌")
            sys.exit(1)
        print("✅")

if __name__ == '__main__':
    API_SPEC_DIR = '.apispec'

    print("\033[93m", "\033[1m", "\033[4m", "Generating API", "\033[0m", sep='')
    generate(API_SPEC_DIR, 'handlers/generated/gen{}/api.go', "../frontend/src/lib/genapi/{}.ts", "../frontend/src/lib/apis/{}.ts")
    print()