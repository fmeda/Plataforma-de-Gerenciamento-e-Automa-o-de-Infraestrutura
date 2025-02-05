import subprocess

def run_terraform(command):
    process = subprocess.Popen(command, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    stdout, stderr = process.communicate()
    if process.returncode != 0:
        print(f"Erro: {stderr.decode('utf-8')}")
    else:
        print(stdout.decode('utf-8'))

def create_aws_resources():
    terraform_init_command = ['terraform', 'init']
    terraform_apply_command = ['terraform', 'apply', '-auto-approve']
    run_terraform(terraform_init_command)
    run_terraform(terraform_apply_command)

if __name__ == "__main__":
    create_aws_resources()
