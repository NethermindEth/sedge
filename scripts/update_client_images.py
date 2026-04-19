#!/usr/bin/env python3
import os
import re
import platform
import stat
import subprocess
import tempfile
import sys
from urllib.request import urlretrieve
from packaging.version import parse as parse_version, InvalidVersion

# Ensure packaging is available. If not, it won't work anyway, but python3 usually has it or it can be pip installed.
# We could check for it explicitly, but earlier tests showed it was available.

YAML_FILE = os.path.join(os.path.dirname(__file__), '..', 'configs', 'client_images.yaml')

def download_crane():
    """Downloads crane executable for the current platform."""
    sys_type = platform.system().lower()
    machine = platform.machine().lower()
    
    if sys_type == 'darwin':
        os_name = 'Darwin'
    elif sys_type == 'linux':
        os_name = 'Linux'
    else:
        print(f"Unsupported OS: {sys_type}")
        sys.exit(1)
        
    if machine in ['x86_64', 'amd64']:
        arch = 'x86_64'
    elif machine in ['arm64', 'aarch64']:
        arch = 'arm64'
    else:
        print(f"Unsupported architecture: {machine}")
        sys.exit(1)

    url = f"https://github.com/google/go-containerregistry/releases/latest/download/go-containerregistry_{os_name}_{arch}.tar.gz"
    
    temp_dir = tempfile.mkdtemp()
    tar_path = os.path.join(temp_dir, 'crane.tar.gz')
    crane_path = os.path.join(temp_dir, 'crane')
    
    print(f"Downloading crane from {url}...")
    urlretrieve(url, tar_path)
    
    subprocess.run(['tar', '-xzf', tar_path, '-C', temp_dir, 'crane'], check=True)
    os.chmod(crane_path, os.stat(crane_path).st_mode | stat.S_IEXEC)
    
    return crane_path

def get_latest_tag(crane_bin, image_name, current_version):
    try:
        result = subprocess.run([crane_bin, 'ls', image_name], capture_output=True, text=True, check=True)
        tags = result.stdout.strip().split('\n')
    except subprocess.CalledProcessError as e:
        print(f"Failed to fetch tags for {image_name}: {e}")
        return current_version
    
    valid_tags = []
    for tag in tags:
        if not tag or tag == 'latest':
            continue
            
        clean_tag = tag
        if clean_tag.startswith('v'): clean_tag = clean_tag[1:]
        elif clean_tag.startswith('multiarch-v'): clean_tag = clean_tag[11:]
        elif clean_tag.startswith('client-'): clean_tag = clean_tag[7:]
             
        # Target semver-like formats (e.g. 1.2.3 or 1.2.3-beta.1)
        if not re.match(r'^\d+\.\d+\.\d+(-[a-zA-Z0-9\.]+)?$', clean_tag):
            continue
            
        try:
            v = parse_version(clean_tag)
            valid_tags.append((v, tag))
        except InvalidVersion:
            pass

    if not valid_tags:
        return current_version
        
    valid_tags.sort(key=lambda x: x[0], reverse=True)
    
    # Try parsing current version to see if it's considered un-stable
    current_is_prerelease = False
    clean_curr = current_version
    for prefix in ['v', 'multiarch-v', 'client-']:
        if clean_curr.startswith(prefix):
            clean_curr = clean_curr[len(prefix):]
    
    try:
        current_is_prerelease = parse_version(clean_curr).is_prerelease
    except:
        pass
    
    # Prefer stable versions if the current version is also stable
    if not current_is_prerelease:
        stable_tags = [t for t in valid_tags if not t[0].is_prerelease]
        if stable_tags:
            return stable_tags[0][1]
            
    return valid_tags[0][1]

def main():
    if not os.path.exists(YAML_FILE):
        print(f"Cannot find {YAML_FILE}")
        sys.exit(1)

    print(f"Reading {YAML_FILE}...")
    with open(YAML_FILE, 'r') as f:
        content = f.read()

    # Find blocks of name + version
    # It assumes formatting:
    #    name: <image_name>
    #    version: <version_tag>
    
    pattern = re.compile(r'(\s+name:\s+([^\n]+)\n\s+version:\s+)([^\n]+)')
    
    # Collect unique images to check
    # Avoid duplicate calls to crane for the same image
    images_to_check = {}
    for match in pattern.finditer(content):
        img_name = match.group(2).strip()
        curr_ver = match.group(3).strip()
        if img_name not in images_to_check:
             images_to_check[img_name] = curr_ver

    crane_bin = download_crane()
    updates = {}
    print("Checking for latest versions...")

    try:
        for img, curr_ver in images_to_check.items():
            if "surge-testnet" in curr_ver or "master" in curr_ver:
                print(f"[{img}] Skipping custom tag '{curr_ver}'")
                continue
            
            latest = get_latest_tag(crane_bin, img, curr_ver)
            if latest != curr_ver:
                print(f"[{img}] {curr_ver} -> {latest}")
                updates[img] = latest
            else:
                print(f"[{img}] is up to date ({curr_ver})")
    finally:
        # Cleanup
        temp_dir = os.path.dirname(crane_bin)
        subprocess.run(['rm', '-rf', temp_dir])

    if not updates:
        print("Everything is already up to date.")
        sys.exit(0)

    # Replace file contents
    def replacer(match):
        prefix_and_name = match.group(1)
        img_name = match.group(2).strip()
        old_ver = match.group(3)
        if img_name in updates:
            new_ver = updates[img_name]
            return f"{prefix_and_name}{new_ver}"
        return match.group(0)

    new_content = pattern.sub(replacer, content)

    with open(YAML_FILE, 'w') as f:
        f.write(new_content)

    print(f"\nSuccessfully updated {YAML_FILE} with the latest stable versions.")

if __name__ == '__main__':
    main()
