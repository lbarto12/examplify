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

API_CLIENT_DRIVER_CODE = '''\
import {{ createApiClient }} from "../genapi/{}";
import {{ PUBLIC_API_BASE_PUBLIC }} from '$env/static/public';
import type {{ AnyZodiosRequestOptions, ZodiosEndpointDefinitions, ZodiosPlugin }} from "@zodios/core";
import type {{ DeepReadonlyObject }} from "@zodios/core/lib/utils.types";

let client = createApiClient(PUBLIC_API_BASE_PUBLIC, {{}});

// Auth plugin using cookieStore (browser only)
const authPlugin: ZodiosPlugin = {{
  request: async (
    api: ZodiosEndpointDefinitions,
    config: DeepReadonlyObject<AnyZodiosRequestOptions>
  ) => {{
    const auth = await cookieStore.get("auth");
    if (!auth) return config;

    const mutableConfig = {{ ...(config as AnyZodiosRequestOptions) }};
    mutableConfig.headers = {{
      ...mutableConfig.headers,
      Authorization: `Bearer ${{auth.value}}`,
    }};

    return mutableConfig;
  }}
}};

export function getClient() {{
  client.use(authPlugin);
  return client;
}}
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

            # Generate client 
            output_client_path: str = client_output_dir.format(domain)
            print("here:", output_client_path)
            output_dirname = os.path.dirname(output_client_path)
            os.makedirs(output_dirname, exist_ok=True)

            subprocess.run([
                "bunx",
                "openapi-zod-client",
                f"{api_spec_path}/{domain}.yaml",
                "-o", f"{output_dirname}/{domain}.ts", "--with-docs"
            ], check=True)

            # Generate driver
            client_driver_path = client_driver_dir.format(domain)
            output_dirname = os.path.dirname(client_driver_path)
            os.makedirs(output_dirname, exist_ok=True)
            with open(f"{output_dirname}/{domain}.svelte.ts", "w") as f:
                f.write(API_CLIENT_DRIVER_CODE.format(domain))

        except subprocess.CalledProcessError as e:
            print("❌")
            sys.exit(1)
        print("✅")

if __name__ == '__main__':
    API_SPEC_DIR = '.apispec'

    print("\033[93m", "\033[1m", "\033[4m", "Generating API", "\033[0m", sep='')
    generate(API_SPEC_DIR, 'handlers/generated/gen{}/api.go', "../frontend/src/lib/genapi/{}.ts", "../frontend/src/lib/apis/{}.ts")
    print()