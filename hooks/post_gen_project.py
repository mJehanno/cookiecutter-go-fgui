import os
from git import Repo

repo = Repo.init(None)
repo.index.add(repo.untracked_files)
repo.index.commit("Init")

os.system("go mod tidy")
